// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { Initializable } from "@openzeppelin/contracts/proxy/utils/Initializable.sol";
import { Math } from "@openzeppelin/contracts/utils/math/Math.sol";

import { Types } from "../libraries/Types.sol";
import { Semver } from "../universal/Semver.sol";
import { L2OutputOracle } from "./L2OutputOracle.sol";
import { ZKVerifier } from "./ZKVerifier.sol";

contract Colosseum is Initializable, Semver {
    /**
     * @notice enum of challenge status.
     *
     * See the https://github.com/wemixkanvas/kanvas/blob/dev/specs/challenge.md#state-diagram for more details.
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
     * @param _challengeTimeout          Timeout seconds for the challenge.
     * @param _segmentsLengths           Lengths of segments.
     */
    constructor(
        L2OutputOracle _l2Oracle,
        ZKVerifier _zkVerifier,
        uint256 _submissionInterval,
        uint256 _challengeTimeout,
        uint256[] memory _segmentsLengths
    ) Semver(0, 1, 0) {
        L2_ORACLE = _l2Oracle;
        ZK_VERIFIER = _zkVerifier;
        CHALLENGE_TIMEOUT = _challengeTimeout;
        L2_ORACLE_SUBMISSION_INTERVAL = _submissionInterval;
        initialize(_segmentsLengths);
    }

    /**
     * @notice Initializer.
     */
    function initialize(uint256[] memory _segmentsLengths) public initializer {
        _setSegmentsLengths(_segmentsLengths);
    }

    function createChallenge(uint256 _outputIndex, bytes32[] calldata _segments) external payable {
        // TODO(pangssu): Currently, only one task can be opened. It is necessary to
        //                consider holding multiple challenges at the same time.
        require(!isInProgress(), "Colosseum: previous challenge is in progress");

        Types.CheckpointOutput memory targetOutput = L2_ORACLE.getL2Output(_outputIndex);
        require(targetOutput.l2BlockNumber != 0, "Colosseum: output not found");
        Types.CheckpointOutput memory prevOutput;
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

    function bisect(uint256 _pos, bytes32[] calldata _segments) external payable {
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

    function proveFault(
        uint256 _pos,
        bytes32 _outputRoot,
        uint256[] calldata _proof,
        uint256[] calldata _pair
    ) external payable {
        uint256 challengeId = latestChallengeId;
        Types.Challenge storage challenge = challenges[challengeId];
        _validateTurn(challenge, ValidateTurnMode.PROOF);

        require(
            challenge.segments[_pos + 1] != _outputRoot,
            "Colosseum: the last segment must not be matched"
        );
        require(ZK_VERIFIER.verify(_proof, _pair), "Colosseum: invalid proof");

        L2_ORACLE.deleteL2Outputs(challenge.outputIndex);
        emit ProofCompleted(challengeId, challenge.outputIndex);
        _closeChallenge(challengeId, challenge);
    }

    function asserterTimeout() external {
        uint256 challengeId = latestChallengeId;
        Types.Challenge storage challenge = challenges[challengeId];
        _validateTurn(challenge, ValidateTurnMode.NORMAL);
        L2_ORACLE.deleteL2Outputs(challenge.outputIndex);
        _closeChallenge(challengeId, challenge);
    }

    function challengerTimeout(uint256 _challengeId) external {
        Types.Challenge storage challenge = challenges[_challengeId];
        _validateTurn(challenge, ValidateTurnMode.NORMAL);
        _closeChallenge(_challengeId, challenge);
    }

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

    function _closeChallenge(uint256 _challengeId, Types.Challenge storage _challenge) private {
        delete challenges[_challengeId];

        emit Closed(_challengeId, _challenge.turn, block.timestamp);
    }

    function _calcSegSize(uint256 _prevSegSize, uint256 _turn) private view returns (uint256) {
        return _prevSegSize / (getSegmentsLength(_turn) - 1);
    }

    function _isPast(uint256 _sec) private view returns (bool) {
        return block.timestamp > _sec;
    }

    function _isTimedOut(ChallengeStatus status) private pure returns (bool) {
        return
            status == ChallengeStatus.CHALLENGER_TIMEOUT ||
            status == ChallengeStatus.ASSERTER_TIMEOUT;
    }

    function _isAbleToBisect(uint256 _segSize, uint256 _turn) private view returns (bool) {
        return _calcSegSize(_segSize, _turn) > 1;
    }

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

    function getSegmentsLength(uint256 _turn) public view returns (uint256) {
        require(_turn > 0, "Colosseum: invalid turn");
        return segmentsLengths[_turn - 1];
    }

    function getStatusInProgress() public view returns (ChallengeStatus) {
        return _challengeStatus(challenges[latestChallengeId]);
    }

    function getChallengeInProgress() public view returns (Types.Challenge memory) {
        Types.Challenge memory challenge = challenges[latestChallengeId];

        require(challenge.turn > 0, "Colosseum: challenge not found");

        return challenge;
    }

    function isAbleToBisect() public view returns (bool) {
        Types.Challenge memory challenge = getChallengeInProgress();
        return _isAbleToBisect(challenge.segSize, challenge.turn);
    }

    function isInProgress() public view returns (bool) {
        ChallengeStatus status = getStatusInProgress();

        // If the challenger turn times out, there is no motivation for the asserter
        // to progress the challenge. The asserter only pay gas. So the challenger
        // timeout status is considered to be closed, too.
        return status != ChallengeStatus.NONE && status != ChallengeStatus.CHALLENGER_TIMEOUT;
    }

    function isChallengeRelated(address _account) public view returns (bool) {
        Types.Challenge memory challenge = challenges[latestChallengeId];
        return challenge.current == _account || challenge.next == _account;
    }
}
