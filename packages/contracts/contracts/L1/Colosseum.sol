// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { Initializable } from "@openzeppelin/contracts/proxy/utils/Initializable.sol";
import { Math } from "@openzeppelin/contracts/utils/math/Math.sol";

import { Hashing } from "../libraries/Hashing.sol";
import { Types } from "../libraries/Types.sol";
import { Semver } from "../universal/Semver.sol";
import { L2OutputOracle } from "./L2OutputOracle.sol";
import { ZKVerifier } from "./ZKVerifier.sol";

contract Colosseum is Initializable, Semver {
    /**
     * @notice Enum of challenge status.
     *
     * See the https://github.com/kroma-network/kroma/blob/dev/specs/challenge.md#state-diagram for more details.
     *
     * Belows are possible state transitions at current implementation.
     * TODO: add PROOF_VERIFIED state.
     *
     *  1) NONE               → createChallenge()                   → ASSERTER_TURN
     *  2) ASSERTER_TURN      → bisect()                            → ASSERTER_TURN
     *  3) ASSERTER_TURN      → on asserter timeout                 → ASSERTER_TIMEOUT
     *  4) CHALLENGER_TURN    → when isAbleToBisect() returns false → READY_TO_PROVE
     *  5) CHALLENGER_TURN    → on challenger timeout               → CHALLENGER_TIMEOUT
     *  6) ASSERTER_TIMEOUT   → asserterTimeout()                   → CLOSED
     *  7) ASSERTER_TIMEOUT   → on challenger timeout               → CHALLENGER_TIMEOUT
     *  8) READY_TO_PROVE     → when proveFault() succeeds          → CLOSED
     *  9) READY_TO_PROVE     → on challenger timeout               → CHALLENGER_TIMEOUT
     * 10) CHALLENGER_TIMEOUT → challengerTimeout()                 → CLOSED
     */
    enum ChallengeStatus {
        NONE,
        CHALLENGER_TURN,
        ASSERTER_TURN,
        CHALLENGER_TIMEOUT,
        ASSERTER_TIMEOUT,
        READY_TO_PROVE
    }

    /**
     * @notice Enum of turn validation modes.
     *
     * @custom:value NORMAL Represents a normal validation mode.
     * @custom:value PROOF  Represents a proof validation mode.
     */
    enum ValidateTurnMode {
        NORMAL,
        PROOF
    }

    /**
     * @notice Address of the L2OutputOracle.
     */
    L2OutputOracle public immutable L2_ORACLE;

    /**
     * @notice Address of the ZKVerifier.
     */
    ZKVerifier public immutable ZK_VERIFIER;

    /**
     * @notice Timeout seconds for the challenge.
     */
    uint256 public immutable CHALLENGE_TIMEOUT;

    /**
     * @notice The interval in L2 blocks at which checkpoints must be
     *         submitted on L2OutputOracle contract.
     */
    uint256 public immutable L2_ORACLE_SUBMISSION_INTERVAL;

    /**
     * @notice The chain id.
     */
    uint256 public immutable CHAIN_ID;

    /**
     * @notice The dummy transaction hash. This is used to pad if the
     *         number of transactions is less than MAX_TXS. This is same as:
     *         unsignedTx = {
     *           nonce: 0,
     *           gasLimit: 0,
     *           gasPrice: 0,
     *           value: 0,
     *           data: '0x',
     *           chainId: CHAIN_ID,
     *         }
     *         signature = sign(unsignedTx, 0x1)
     *         dummyHash = keccak256(rlp({
     *           ...unsignedTx,
     *           signature,
     *         }))
     */
    bytes32 public immutable DUMMY_HASH;

    /**
     * @notice The maximum number of transactions
     */
    uint256 public immutable MAX_TXS;

    /**
     * @notice Length of segment array for each turn.
     */
    mapping(uint256 => uint256) internal segmentsLengths;

    /**
     * @notice A mapping of challenge.
     */
    mapping(uint256 => Types.Challenge) public challenges;

    /**
     * @notice The number of the most recent challenge created by this contract.
     */
    uint256 public latestChallengeId;

    /**
     * @notice Emitted when the challenge is created.
     *
     * @param challengeId Identifier of the challenge.
     * @param challenger  Address of the creator.
     * @param outputIndex Index of invalid output.
     * @param timestamp   The timestamp when created.
     */
    event ChallengeCreated(
        uint256 indexed challengeId,
        address indexed challenger,
        uint256 indexed outputIndex,
        uint256 timestamp
    );

    event Bisected(uint256 indexed challengeId, uint256 turn, uint256 timestamp);
    event ProofCompleted(uint256 indexed challengeId, uint256 outputIndex);
    event Closed(uint256 indexed challengeId, uint256 turn, uint256 timestamp);

    /**
     * @custom:semver 0.1.0
     *
     * @param _l2Oracle                  Address of the L2OutputOracle contract.
     * @param _zkVerifier                Address of the ZKVerifier contract.
     * @param _submissionInterval        Interval in blocks at which checkpoints must be submitted.
     * @param _challengeTimeout          Timeout seconds for the challenge.
     * @param _chainId                   Chain ID.
     * @param _dummyHash                 Dummy hash.
     * @param _segmentsLengths           Lengths of segments.
     */
    constructor(
        L2OutputOracle _l2Oracle,
        ZKVerifier _zkVerifier,
        uint256 _submissionInterval,
        uint256 _challengeTimeout,
        uint256 _chainId,
        bytes32 _dummyHash,
        uint256 _maxTxs,
        uint256[] memory _segmentsLengths
    ) Semver(0, 1, 0) {
        L2_ORACLE = _l2Oracle;
        ZK_VERIFIER = _zkVerifier;
        CHALLENGE_TIMEOUT = _challengeTimeout;
        L2_ORACLE_SUBMISSION_INTERVAL = _submissionInterval;
        CHAIN_ID = _chainId;
        DUMMY_HASH = _dummyHash;
        MAX_TXS = _maxTxs;
        initialize(_segmentsLengths);
    }

    /**
     * @notice Initializer.
     */
    function initialize(uint256[] memory _segmentsLengths) public initializer {
        _setSegmentsLengths(_segmentsLengths);
    }

    /**
     * @notice Creates a challenge against an invalid output.
     *
     * @param _outputIndex Index of invalid output.
     * @param _segments    Array of the segment. A segment is the first output root of a specific range.
     */
    function createChallenge(uint256 _outputIndex, bytes32[] calldata _segments) external {
        // TODO(pangssu): Currently, only one task can be opened. It is necessary to
        //                consider holding multiple challenges at the same time.
        require(!isInProgress(), "Colosseum: previous challenge is in progress");

        Types.CheckpointOutput memory targetOutput = L2_ORACLE.getL2Output(_outputIndex);
        require(targetOutput.l2BlockNumber != 0, "Colosseum: output not found");
        Types.CheckpointOutput memory prevOutput;
        // TODO(chokobole): Enable dispute resolution including genesis output root.
        if (_outputIndex > 0) {
            prevOutput = L2_ORACLE.getL2Output(_outputIndex - 1);
        }

        _validateSegments(1, prevOutput.outputRoot, targetOutput.outputRoot, _segments);

        uint256 challengeId = latestChallengeId + 1;
        Types.Challenge storage challenge = challenges[challengeId];

        _updateSegments(
            challenge,
            _segments,
            prevOutput.l2BlockNumber,
            targetOutput.l2BlockNumber - prevOutput.l2BlockNumber
        );

        challenge.outputIndex = _outputIndex;
        challenge.turn = 1;
        challenge.current = msg.sender;
        challenge.next = L2_ORACLE.VALIDATOR();
        challenge.timeoutAt = block.timestamp + CHALLENGE_TIMEOUT;

        latestChallengeId = challengeId;

        emit ChallengeCreated(challengeId, msg.sender, _outputIndex, block.timestamp);
    }

    /**
     * @notice Selects an invalid section and submit segments of that section.
     *
     * @param _pos         Position of invalid section.
     * @param _segments    Array of the segment. A segment is the first output root of a specific range.
     */
    function bisect(uint256 _pos, bytes32[] calldata _segments) external {
        uint256 challengeId = latestChallengeId;
        Types.Challenge storage challenge = challenges[challengeId];

        _validateTurn(challenge, ValidateTurnMode.NORMAL);

        uint256 newTurn = challenge.turn + 1;

        _validateSegments(
            newTurn,
            challenge.segments[_pos],
            challenge.segments[_pos + 1],
            _segments
        );

        uint256 segSize = _calcSegSize(challenge.segSize, newTurn - 1);
        uint256 segStart = challenge.segStart + _pos * segSize;

        _updateSegments(challenge, _segments, segStart, segSize);

        challenge.turn = newTurn;
        challenge.next = challenge.current;
        challenge.current = msg.sender;
        challenge.timeoutAt = block.timestamp + CHALLENGE_TIMEOUT;

        emit Bisected(challengeId, newTurn, block.timestamp);
    }

    /**
     * @notice Proves that a specific output is invalid using ZKP.
     *         This function can only be called in the READY_TO_PROVE state.
     *
     * @param _pos                Position of invalid section.
     * @param _srcOutputRootProof TBD
     * @param _dstOutputRootProof TBD
     * @param _publicInput        TBD
     * @param _rlps               TBD
     * @param _proof              TBD
     * @param _pair               TBD
     */
    function proveFault(
        uint256 _pos,
        Types.OutputRootProof calldata _srcOutputRootProof,
        Types.OutputRootProof calldata _dstOutputRootProof,
        Types.PublicInput calldata _publicInput,
        Types.BlockHeaderRLP calldata _rlps,
        uint256[] calldata _proof,
        uint256[] calldata _pair
    ) external {
        uint256 challengeId = latestChallengeId;
        Types.Challenge storage challenge = challenges[challengeId];
        _validateTurn(challenge, ValidateTurnMode.PROOF);

        bytes32 srcOutputRoot = Hashing.hashOutputRootProof(_srcOutputRootProof);
        require(
            challenge.segments[_pos] == srcOutputRoot,
            "Colosseum: the source segment must be matched"
        );

        bytes32 dstOutputRoot = Hashing.hashOutputRootProof(_dstOutputRootProof);
        require(
            challenge.segments[_pos + 1] != dstOutputRoot,
            "Colosseum: the destination segment must not be matched"
        );

        require(
            _srcOutputRootProof.nextBlockHash == _dstOutputRootProof.blockHash,
            "Colosseum: the block hash must be matched"
        );

        bytes32 blockHash = Hashing.hashBlockHeader(_publicInput, _rlps);
        require(
            _srcOutputRootProof.nextBlockHash == blockHash,
            "Colosseum: the block hash must be matched"
        );

        bytes32[] memory dummyHashes;
        if (_publicInput.txHashes.length < MAX_TXS) {
            dummyHashes = Hashing.generateDummyHashes(
                DUMMY_HASH,
                MAX_TXS - _publicInput.txHashes.length
            );
        }

        // TODO(chokobole): check transaction root hash
        // TODO(chokobole): use this public input hash require(_proof[4] == _publicInputHash);
        bytes32 _publicInputHash = Hashing.hashPublicInput(
            _srcOutputRootProof.stateRoot,
            _publicInput,
            CHAIN_ID,
            dummyHashes
        );

        require(ZK_VERIFIER.verify(_proof, _pair), "Colosseum: invalid proof");

        L2_ORACLE.deleteL2Outputs(challenge.outputIndex);
        emit ProofCompleted(challengeId, challenge.outputIndex);
        _closeChallenge(challengeId, challenge);
    }

    /**
     * @notice Closes the challenge because the asserter timed out.
     *         The winner is the challenger.
     */
    function asserterTimeout() external {
        uint256 challengeId = latestChallengeId;
        Types.Challenge storage challenge = challenges[challengeId];
        _validateTurn(challenge, ValidateTurnMode.NORMAL);
        L2_ORACLE.deleteL2Outputs(challenge.outputIndex);
        _closeChallenge(challengeId, challenge);
    }

    /**
     * @notice Closes the challenge because the challenger timed out.
     *         The winner is the asserter.
     *
     * @param _challengeId Identifier of the challenge.
     */
    function challengerTimeout(uint256 _challengeId) external {
        Types.Challenge storage challenge = challenges[_challengeId];
        _validateTurn(challenge, ValidateTurnMode.NORMAL);
        _closeChallenge(_challengeId, challenge);
    }

    /**
     * @notice Reverts if it's not sender's turn.
     *
     * @param _challenge The challenge data.
     * @param _mode      Turn validation mode.
     */
    function _validateTurn(Types.Challenge storage _challenge, ValidateTurnMode _mode)
        private
        view
    {
        ChallengeStatus status = _challengeStatus(_challenge);
        require(status != ChallengeStatus.NONE, "Colosseum: challenge not found");

        address expectedSender = _isTimedOut(status) ? _challenge.current : _challenge.next;
        require(expectedSender == msg.sender, "Colosseum: not your turn");

        if (_mode == ValidateTurnMode.PROOF) {
            require(status == ChallengeStatus.READY_TO_PROVE, "Colosseum: not ready to prove");
        }
    }

    /**
     * @notice Reverts if the given segments are invalid.
     *
     * @param _turn      The current turn.
     * @param _prevFirst The first segment of previous turn.
     * @param _prevLast  The last segment of previous turn.
     * @param _segments  Array of the segment.
     */
    function _validateSegments(
        uint256 _turn,
        bytes32 _prevFirst,
        bytes32 _prevLast,
        bytes32[] memory _segments
    ) private view {
        uint256 segLen = _segments.length;

        require(getSegmentsLength(_turn) == segLen, "Colosseum: invalid segments length");
        require(_prevFirst == _segments[0], "Colosseum: the first segment must be matched");
        require(
            _prevLast != _segments[segLen - 1],
            "Colosseum: the last segment must not be matched"
        );
    }

    /**
     * @notice Updates the segment information for a given challenge.
     *
     * @param _challenge The challenge data.
     * @param _segments  Array of the segment.
     * @param _segStart  The L2 block number of the first segment.
     * @param _segSize   The number of L2 blocks.
     */
    function _updateSegments(
        Types.Challenge storage _challenge,
        bytes32[] memory _segments,
        uint256 _segStart,
        uint256 _segSize
    ) private {
        _challenge.segments = _segments;
        _challenge.segStart = _segStart;
        _challenge.segSize = _segSize;
    }

    /**
     * @notice Validates and updates the lengths of segments.
     *
     * @param _segmentsLengths Lengths of segments.
     */
    function _setSegmentsLengths(uint256[] memory _segmentsLengths) private {
        // _segmentsLengths length should be an even number in order to let challenger submit
        // invalidity proof at the last turn.
        require(
            _segmentsLengths.length % 2 == 0,
            "Colosseum: length of segments lengths cannot be odd number."
        );

        uint256 sum = 1;
        for (uint256 i = 0; i < _segmentsLengths.length; i++) {
            segmentsLengths[i] = _segmentsLengths[i];
            sum = sum * (_segmentsLengths[i] - 1);
        }

        require(sum == L2_ORACLE_SUBMISSION_INTERVAL, "Colosseum: invalid segments lengths");
    }

    /**
     * @notice Slashes loser's deposit and close the challenge.
     *
     * @param _challengeId Identifier of the challenge.
     * @param _challenge   The challenge data.
     */
    function _closeChallenge(uint256 _challengeId, Types.Challenge storage _challenge) private {
        delete challenges[_challengeId];

        emit Closed(_challengeId, _challenge.turn, block.timestamp);
    }

    /**
     * @notice Returns the number of L2 blocks of a given turn.
     *
     * @param _prevSegSize The number of L2 blocks submitted by the previous turn.
     * @param _turn        The current turn.
     *
     * @return The number of L2 blocks.
     */
    function _calcSegSize(uint256 _prevSegSize, uint256 _turn) private view returns (uint256) {
        return _prevSegSize / (getSegmentsLength(_turn) - 1);
    }

    /**
     * @notice Determines whether a given timestamp is past.
     *
     * @param _sec The timestamp to check.
     *
     * @return Whether it's in the past.
     */
    function _isPast(uint256 _sec) private view returns (bool) {
        return block.timestamp > _sec;
    }

    /**
     * @notice Determines whether a given challenge timed out.
     *
     * @param _status The challenge status.
     *
     * @return Whether the challenge timed out.
     */
    function _isTimedOut(ChallengeStatus _status) private pure returns (bool) {
        return
            _status == ChallengeStatus.CHALLENGER_TIMEOUT ||
            _status == ChallengeStatus.ASSERTER_TIMEOUT;
    }

    /**
     * @notice Determines if bisection is possible.
     *
     * @param _segSize The number of L2 blocks.
     * @param _turn    The current turn.
     *
     * @return Whether bisection is possible.
     */
    function _isAbleToBisect(uint256 _segSize, uint256 _turn) private view returns (bool) {
        return _calcSegSize(_segSize, _turn) > 1;
    }

    /**
     * @notice Returns status of a given challenge.
     *
     * @param _challenge The challenge data.
     *
     * @return The status of challenge.
     */
    function _challengeStatus(Types.Challenge storage _challenge)
        private
        view
        returns (ChallengeStatus)
    {
        // If challenge does not exist or is closed, the turn is 0.
        if (_challenge.turn == 0) {
            return ChallengeStatus.NONE;
        }

        // Challenger turns are odd numbers and asserter turns are even numbers.
        bool isChallengerTurn = (_challenge.turn + 1) % 2 == 1;

        // Check if it is a timed out challenge.
        if (_isPast(_challenge.timeoutAt)) {
            // timeout on challenger turn
            if (isChallengerTurn) {
                return ChallengeStatus.CHALLENGER_TIMEOUT;
            }

            // If the challenger doesn't close a challenge even though the asserter has timed out,
            // the challenger is considered to have timed out.
            if (_isPast(_challenge.timeoutAt + CHALLENGE_TIMEOUT)) {
                return ChallengeStatus.CHALLENGER_TIMEOUT;
            }

            // timeout on asserter turn
            return ChallengeStatus.ASSERTER_TIMEOUT;
        }

        // If bisection is not possible, the Challenger must execute the fault proof.
        if (!_isAbleToBisect(_challenge.segSize, _challenge.turn)) {
            return ChallengeStatus.READY_TO_PROVE;
        }

        return isChallengerTurn ? ChallengeStatus.CHALLENGER_TURN : ChallengeStatus.ASSERTER_TURN;
    }

    /**
     * @notice Returns the segment length required for that turn.
     *
     * @param _turn The challenge turn.
     *
     * @return The segments length.
     */
    function getSegmentsLength(uint256 _turn) public view returns (uint256) {
        require(_turn > 0, "Colosseum: invalid turn");
        return segmentsLengths[_turn - 1];
    }

    /**
     * @notice Returns status of the current challenge.
     *
     * @return The status of the current challenge.
     */
    function getStatusInProgress() public view returns (ChallengeStatus) {
        return _challengeStatus(challenges[latestChallengeId]);
    }

    /**
     * @notice Returns challenge data of the current challenge.
     *
     * @return The current challenge data
     */
    function getChallengeInProgress() public view returns (Types.Challenge memory) {
        Types.Challenge memory challenge = challenges[latestChallengeId];

        require(challenge.turn > 0, "Colosseum: challenge not found");

        return challenge;
    }

    /**
     * @notice Determines whether bisection is possible.
     *
     * @return Whether bisection is possible.
     */
    function isAbleToBisect() public view returns (bool) {
        Types.Challenge memory challenge = getChallengeInProgress();
        return _isAbleToBisect(challenge.segSize, challenge.turn);
    }

    /**
     * @notice Determines whether the current challenge is in progress.
     *
     * @return Whether the challenge is in progress.
     */
    function isInProgress() public view returns (bool) {
        ChallengeStatus status = getStatusInProgress();

        // If the challenger turn times out, there is no motivation for the asserter
        // to progress the challenge. The asserter only pay gas. So the challenger
        // timeout status is considered to be closed, too.
        return status != ChallengeStatus.NONE && status != ChallengeStatus.CHALLENGER_TIMEOUT;
    }

    /**
     * @notice Determines whether a given address is a participant in the current challenge.
     *
     * @param _address Address.
     *
     * @return Whether a given address is a participant.
     */
    function isChallengeRelated(address _address) public view returns (bool) {
        Types.Challenge memory challenge = challenges[latestChallengeId];
        return challenge.current == _address || challenge.next == _address;
    }
}
