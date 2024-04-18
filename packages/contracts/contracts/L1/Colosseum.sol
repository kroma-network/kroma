// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { Initializable } from "@openzeppelin/contracts/proxy/utils/Initializable.sol";

import { Hashing } from "../libraries/Hashing.sol";
import { Predeploys } from "../libraries/Predeploys.sol";
import { Types } from "../libraries/Types.sol";
import { ISemver } from "../universal/ISemver.sol";
import { IZKMerkleTrie } from "./IZKMerkleTrie.sol";
import { L2OutputOracle } from "./L2OutputOracle.sol";
import { SecurityCouncil } from "./SecurityCouncil.sol";
import { ZKVerifier } from "./ZKVerifier.sol";

contract Colosseum is Initializable, ISemver {
    /**
     * @notice The constant value for the first turn.
     */
    uint8 internal constant TURN_INIT = 1;

    /**
     * @notice The constant value for the delete output root.
     */
    bytes32 internal constant DELETED_OUTPUT_ROOT = bytes32(0);

    /**
     * @notice Enum of the challenge status.
     *
     * See the https://github.com/kroma-network/kroma/blob/dev/specs/challenge.md#state-diagram
     * for more details.
     *
     * Belows are possible state transitions at current implementation.
     *
     *  1) NONE               → createChallenge()                   → ASSERTER_TURN
     *  2) ASSERTER_TURN      → bisect()                            → CHALLENGER_TURN
     *  3) ASSERTER_TURN      → on bisection timeout                → ASSERTER_TIMEOUT
     *  4) CHALLENGER_TURN    → bisect()                            → ASSERTER_TURN
     *  5) CHALLENGER_TURN    → when isAbleToBisect() returns false → READY_TO_PROVE
     *  6) CHALLENGER_TURN    → on bisection timeout                → CHALLENGER_TIMEOUT
     *  7) ASSERTER_TIMEOUT   → when proveFault() succeeds          → NONE
     *  8) ASSERTER_TIMEOUT   → on proving timeout                  → CHALLENGER_TIMEOUT
     *  9) READY_TO_PROVE     → when proveFault() succeeds          → NONE
     * 10) READY_TO_PROVE     → on proving timeout                  → CHALLENGER_TIMEOUT
     * 11) CHALLENGER_TIMEOUT → challengerTimeout()                 → NONE
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
     * @notice Address of the L2OutputOracle.
     */
    L2OutputOracle public immutable L2_ORACLE;

    /**
     * @notice Address of the ZKVerifier.
     */
    ZKVerifier public immutable ZK_VERIFIER;

    /**
     * @notice The period seconds for which challenges can be created per each output.
     */
    uint256 public immutable CREATION_PERIOD_SECONDS;

    /**
     * @notice Timeout seconds for the bisection.
     */
    uint256 public immutable BISECTION_TIMEOUT;

    /**
     * @notice Timeout seconds for the proving.
     */
    uint256 public immutable PROVING_TIMEOUT;

    /**
     * @notice The interval in L2 blocks at which checkpoints must be
     *         submitted on L2OutputOracle contract.
     */
    uint256 public immutable L2_ORACLE_SUBMISSION_INTERVAL;

    /**
     * @notice The dummy transaction hash. This is used to pad if the
     *         number of transactions is less than MAX_TXS. This is same as:
     *         unsignedTx = {
     *           nonce: 0,
     *           gasLimit: 0,
     *           gasPrice: 0,
     *           to: address(0),
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
     * @notice Address that has the ability to approve the challenge.
     */
    address public immutable SECURITY_COUNCIL;

    /**
     * @notice Address that has the ability to verify the merkle proof.
     */
    address public immutable ZK_MERKLE_TRIE;

    /**
     * @notice Length of segment array for each turn.
     */
    mapping(uint256 => uint256) internal segmentsLengths;

    /**
     * @notice A mapping of the challenge.
     */
    mapping(uint256 => mapping(address => Types.Challenge)) public challenges;

    /**
     * @notice A mapping indicating whether a public input is verified or not.
     */
    mapping(bytes32 => bool) public verifiedPublicInputs;

    /**
     * @notice Emitted when the challenge is created.
     *
     * @param outputIndex Index of the L2 checkpoint output.
     * @param asserter    Address of the asserter.
     * @param challenger  Address of the challenger.
     * @param timestamp   The timestamp when created.
     */
    event ChallengeCreated(
        uint256 indexed outputIndex,
        address indexed asserter,
        address indexed challenger,
        uint256 timestamp
    );

    /**
     * @notice Emitted when segments are bisected.
     *
     * @param outputIndex Index of the L2 checkpoint output.
     * @param challenger  Address of the challenger.
     * @param turn        The current turn.
     * @param timestamp   The timestamp when bisected.
     */
    event Bisected(
        uint256 indexed outputIndex,
        address indexed challenger,
        uint8 turn,
        uint256 timestamp
    );

    /**
     * @notice Emitted when it is ready to be proved.
     *
     * @param outputIndex Index of the L2 checkpoint output.
     * @param challenger  Address of the challenger.
     */
    event ReadyToProve(uint256 indexed outputIndex, address indexed challenger);

    /**
     * @notice Emitted when proven fault.
     *
     * @param outputIndex Index of the L2 checkpoint output.
     * @param challenger  Address of the challenger.
     * @param timestamp   The timestamp when proven.
     */
    event Proven(uint256 indexed outputIndex, address indexed challenger, uint256 timestamp);

    /**
     * @notice Emitted when challenge is dismissed.
     *
     * @param outputIndex Index of the L2 checkpoint output.
     * @param challenger  Address of the challenger.
     * @param timestamp   The timestamp when dismissed.
     */
    event ChallengeDismissed(
        uint256 indexed outputIndex,
        address indexed challenger,
        uint256 timestamp
    );

    /**
     * @notice Emitted when challenge is canceled.
     *
     * @param outputIndex Index of the L2 checkpoint output.
     * @param challenger  Address of the challenger.
     * @param timestamp   The timestamp when canceled.
     */
    event ChallengeCanceled(
        uint256 indexed outputIndex,
        address indexed challenger,
        uint256 timestamp
    );

    /**
     * @notice Emitted when challenger timed out.
     *
     * @param outputIndex Index of the L2 checkpoint output.
     * @param challenger  Address of the challenger.
     * @param timestamp   The timestamp when deleted.
     */
    event ChallengerTimedOut(
        uint256 indexed outputIndex,
        address indexed challenger,
        uint256 timestamp
    );

    /**
     * @notice Reverts when caller is not allowed.
     */
    error NotAllowedCaller();

    /**
     * @notice Reverts when output is already finalized.
     */
    error OutputAlreadyFinalized();

    /**
     * @notice Reverts when output is already deleted.
     */
    error OutputAlreadyDeleted();

    /**
     * @notice Reverts when output is not deleted.
     */
    error OutputNotDeleted();

    /**
     * @notice Reverts when output is genesis output.
     */
    error NotAllowedGenesisOutput();

    /**
     * @notice Reverts when the status of challenge is improper.
     */
    error ImproperChallengeStatus();

    /**
     * @notice Reverts when the creation period is already passed.
     */
    error CreationPeriodPassed();

    /**
     * @notice Reverts when L1 is reorged.
     */
    error L1Reorged();

    /**
     * @notice Reverts when the public input has already been used to prove fault.
     */
    error AlreadyUsedPublicInput();

    /**
     * @notice Reverts when the ZK proof is invalid.
     */
    error InvalidZKProof();

    /**
     * @notice Reverts when the inclusion proof is invalid.
     */
    error InvalidInclusionProof();

    /**
     * @notice Reverts when segments length is invalid.
     */
    error InvalidSegmentsLength();

    /**
     * @notice Reverts when the first segment is mismatched.
     */
    error FirstSegmentMismatched();

    /**
     * @notice Reverts when the last segment is matched.
     */
    error LastSegmentMatched();

    /**
     * @notice Reverts when the block hash is mismatched.
     */
    error BlockHashMismatched();

    /**
     * @notice Reverts when the state root is mismatched.
     */
    error StateRootMismatched();

    /**
     * @notice Reverts when turn is invalid.
     */
    error InvalidTurn();

    /**
     * @notice Reverts when challenge cannot be cancelled.
     */
    error CannotCancelChallenge();

    /**
     * @notice Reverts when try to rollback output to zero hash.
     */
    error CannotRollbackOutputToZero();

    /**
     * @notice A modifier that only allows the security council to call
     */
    modifier onlySecurityCouncil() {
        if (msg.sender != SECURITY_COUNCIL) revert NotAllowedCaller();
        _;
    }

    /**
     * @notice Reverts if the output of given index is already finalized.
     *
     * @param _outputIndex Index of the L2 checkpoint output.
     */
    modifier outputNotFinalized(uint256 _outputIndex) {
        if (L2_ORACLE.isFinalized(_outputIndex)) revert OutputAlreadyFinalized();
        _;
    }

    /**
     * @notice Semantic version.
     * @custom:semver 1.1.0
     */
    string public constant version = "1.1.0";

    /**
     * @notice Constructs the Colosseum contract.
     *
     * @param _l2Oracle              Address of the L2OutputOracle contract.
     * @param _zkVerifier            Address of the ZKVerifier contract.
     * @param _submissionInterval    Interval in blocks at which checkpoints must be submitted.
     * @param _creationPeriodSeconds Seconds The period seconds for which challenges can be created per each output.
     * @param _bisectionTimeout      Timeout seconds for the bisection.
     * @param _provingTimeout        Timeout seconds for the proving.
     * @param _dummyHash             Dummy hash.
     * @param _maxTxs                Number of max transactions per block.
     * @param _segmentsLengths       Lengths of segments.
     * @param _securityCouncil       Address of security council.
     * @param _zkMerkleTrie          Address of zk merkle trie.
     */
    constructor(
        L2OutputOracle _l2Oracle,
        ZKVerifier _zkVerifier,
        uint256 _submissionInterval,
        uint256 _creationPeriodSeconds,
        uint256 _bisectionTimeout,
        uint256 _provingTimeout,
        bytes32 _dummyHash,
        uint256 _maxTxs,
        uint256[] memory _segmentsLengths,
        address _securityCouncil,
        address _zkMerkleTrie
    ) {
        L2_ORACLE = _l2Oracle;
        ZK_VERIFIER = _zkVerifier;
        CREATION_PERIOD_SECONDS = _creationPeriodSeconds;
        BISECTION_TIMEOUT = _bisectionTimeout;
        PROVING_TIMEOUT = _provingTimeout;
        L2_ORACLE_SUBMISSION_INTERVAL = _submissionInterval;
        DUMMY_HASH = _dummyHash;
        MAX_TXS = _maxTxs;
        SECURITY_COUNCIL = _securityCouncil;
        ZK_MERKLE_TRIE = _zkMerkleTrie;
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
     * @param _outputIndex   Index of the invalid L2 checkpoint output.
     * @param _l1BlockHash   The block hash of L1 at the time the output L2 block was created.
     * @param _l1BlockNumber The block number of L1 with the specified L1 block hash.
     * @param _segments      Array of the segment. A segment is the first output root of a specific range.
     */
    function createChallenge(
        uint256 _outputIndex,
        bytes32 _l1BlockHash,
        uint256 _l1BlockNumber,
        bytes32[] calldata _segments
    ) external {
        if (_outputIndex == 0) revert NotAllowedGenesisOutput();

        Types.Challenge storage challenge = challenges[_outputIndex][msg.sender];

        if (challenge.turn >= TURN_INIT) {
            ChallengeStatus status = _challengeStatus(challenge);
            if (status != ChallengeStatus.CHALLENGER_TIMEOUT) revert ImproperChallengeStatus();

            _challengerTimeout(_outputIndex, msg.sender);
        }

        Types.CheckpointOutput memory targetOutput = L2_ORACLE.getL2Output(_outputIndex);

        if (targetOutput.timestamp + CREATION_PERIOD_SECONDS < block.timestamp)
            revert CreationPeriodPassed();

        if (targetOutput.outputRoot == DELETED_OUTPUT_ROOT) revert OutputAlreadyDeleted();

        if (msg.sender == targetOutput.submitter) revert NotAllowedCaller();

        if (_l1BlockHash != bytes32(0) && blockhash(_l1BlockNumber) != bytes32(0)) {
            // Like L2OutputOracle, it reverts transactions when L1 reorged.
            if (blockhash(_l1BlockNumber) != _l1BlockHash) revert L1Reorged();
        }

        Types.CheckpointOutput memory prevOutput = L2_ORACLE.getL2Output(_outputIndex - 1);

        // If the previous output has been deleted, the first segment will not be compared with the previous output.
        if (prevOutput.outputRoot == DELETED_OUTPUT_ROOT) {
            _validateSegments(TURN_INIT, _segments[0], targetOutput.outputRoot, _segments);
        } else {
            _validateSegments(TURN_INIT, prevOutput.outputRoot, targetOutput.outputRoot, _segments);
        }

        // Switch validator system after validator pool contract terminated.
        if (L2_ORACLE.VALIDATOR_POOL().isTerminated(_outputIndex)) {
            // Only the validators who satisfy output submission condition can create challenge.
            L2_ORACLE.VALIDATOR_MANAGER().assertCanSubmitOutput(msg.sender);
        } else {
            L2_ORACLE.VALIDATOR_POOL().addPendingBond(_outputIndex, msg.sender);
        }

        _updateSegments(
            challenge,
            _segments,
            targetOutput.l2BlockNumber - L2_ORACLE_SUBMISSION_INTERVAL,
            L2_ORACLE_SUBMISSION_INTERVAL
        );
        challenge.turn = TURN_INIT;
        challenge.asserter = targetOutput.submitter;
        challenge.challenger = msg.sender;
        _updateTimeout(challenge);

        emit ChallengeCreated(_outputIndex, targetOutput.submitter, msg.sender, block.timestamp);
    }

    /**
     * @notice Selects an invalid section and submit segments of that section.
     *
     * @param _outputIndex Index of the L2 checkpoint output.
     * @param _challenger  Address of the challenger.
     * @param _pos         Position of the last valid segment.
     * @param _segments    Array of the segment. A segment is the first output root of a specific range.
     */
    function bisect(
        uint256 _outputIndex,
        address _challenger,
        uint256 _pos,
        bytes32[] calldata _segments
    ) external outputNotFinalized(_outputIndex) {
        Types.Challenge storage challenge = challenges[_outputIndex][_challenger];
        ChallengeStatus status = _challengeStatus(challenge);

        if (_cancelIfOutputDeleted(_outputIndex, challenge.challenger, status)) {
            return;
        }

        address expectedSender;
        if (status == ChallengeStatus.CHALLENGER_TURN) {
            expectedSender = challenge.challenger;
        } else if (status == ChallengeStatus.ASSERTER_TURN) {
            expectedSender = challenge.asserter;
        }
        if (msg.sender != expectedSender) revert NotAllowedCaller();

        uint8 newTurn = challenge.turn + 1;

        _validateSegments(
            newTurn,
            challenge.segments[_pos],
            challenge.segments[_pos + 1],
            _segments
        );

        uint256 segSize = _nextSegSize(challenge);
        uint256 segStart = challenge.segStart + _pos * segSize;

        _updateSegments(challenge, _segments, segStart, segSize);

        challenge.turn = newTurn;
        _updateTimeout(challenge);

        emit Bisected(_outputIndex, _challenger, newTurn, block.timestamp);

        if (!_isAbleToBisect(challenge)) {
            emit ReadyToProve(_outputIndex, _challenger);
        }
    }

    /**
     * @notice Proves that a specific output is invalid using ZKP.
     *         This function can only be called in the READY_TO_PROVE and ASSERTER_TIMEOUT states.
     *
     * @param _outputIndex Index of the L2 checkpoint output.
     * @param _pos         Position of the last valid segment.
     * @param _proof       Proof for public input validation.
     * @param _zkproof     Halo2 proofs composed of points and scalars.
     *                     See https://zcash.github.io/halo2/design/implementation/proofs.html.
     * @param _pair        Aggregated multi-opening proofs and public inputs. (Currently only 2 public inputs)
     */
    function proveFault(
        uint256 _outputIndex,
        uint256 _pos,
        Types.PublicInputProof calldata _proof,
        uint256[] calldata _zkproof,
        uint256[] calldata _pair
    ) external outputNotFinalized(_outputIndex) {
        Types.Challenge storage challenge = challenges[_outputIndex][msg.sender];
        ChallengeStatus status = _challengeStatus(challenge);

        if (_cancelIfOutputDeleted(_outputIndex, challenge.challenger, status)) {
            return;
        }

        if (status != ChallengeStatus.READY_TO_PROVE && status != ChallengeStatus.ASSERTER_TIMEOUT)
            revert ImproperChallengeStatus();

        bytes32 srcOutputRoot = Hashing.hashOutputRootProof(_proof.srcOutputRootProof);
        bytes32 dstOutputRoot = Hashing.hashOutputRootProof(_proof.dstOutputRootProof);

        _validateOutputRootProof(
            _pos,
            challenge,
            srcOutputRoot,
            dstOutputRoot,
            _proof.srcOutputRootProof,
            _proof.dstOutputRootProof
        );
        _validatePublicInput(
            _proof.srcOutputRootProof,
            _proof.dstOutputRootProof,
            _proof.publicInput,
            _proof.rlps
        );
        _validateWithdrawalStorageRoot(
            _proof.merkleProof,
            _proof.l2ToL1MessagePasserBalance,
            _proof.l2ToL1MessagePasserCodeHash,
            _proof.dstOutputRootProof.messagePasserStorageRoot,
            _proof.dstOutputRootProof.stateRoot
        );

        bytes32 publicInputHash = _hashPublicInput(
            _proof.srcOutputRootProof.stateRoot,
            _proof.publicInput
        );

        if (verifiedPublicInputs[publicInputHash]) revert AlreadyUsedPublicInput();

        if (!ZK_VERIFIER.verify(_zkproof, _pair, publicInputHash)) revert InvalidZKProof();
        emit Proven(_outputIndex, msg.sender, block.timestamp);

        // Scope to call the security council, to avoid stack too deep.
        {
            Types.CheckpointOutput memory output = L2_ORACLE.getL2Output(_outputIndex);

            bytes memory callbackData = abi.encodeWithSelector(
                this.dismissChallenge.selector,
                _outputIndex,
                msg.sender,
                challenge.asserter,
                output.outputRoot,
                publicInputHash
            );

            // Request outputRoot validation to security council
            SecurityCouncil(SECURITY_COUNCIL).requestValidation(
                output.outputRoot,
                output.l2BlockNumber,
                callbackData
            );
        }

        // Switch validator system after validator pool contract terminated.
        if (L2_ORACLE.VALIDATOR_POOL().isTerminated(_outputIndex)) {
            // Slash the asseter's asset and move it to pending challenge reward for the output.
            L2_ORACLE.VALIDATOR_MANAGER().slash(_outputIndex, msg.sender, challenge.asserter);
        } else {
            // The challenger's bond is also included in the bond for that output.
            L2_ORACLE.VALIDATOR_POOL().increaseBond(_outputIndex, msg.sender);
        }

        verifiedPublicInputs[publicInputHash] = true;
        delete challenges[_outputIndex][msg.sender];

        // Delete output root.
        L2_ORACLE.replaceL2Output(_outputIndex, DELETED_OUTPUT_ROOT, msg.sender);
    }

    /**
     * @notice Calls a private function that deletes the challenge because the challenger has timed out.
     *         Reverts if the challenger hasn't timed out.
     *
     * @param _outputIndex Index of the L2 checkpoint output.
     * @param _challenger  Address of the challenger.
     */
    function challengerTimeout(uint256 _outputIndex, address _challenger) external {
        Types.Challenge storage challenge = challenges[_outputIndex][_challenger];
        ChallengeStatus status = _challengeStatus(challenge);

        if (status != ChallengeStatus.CHALLENGER_TIMEOUT) revert ImproperChallengeStatus();

        _challengerTimeout(_outputIndex, _challenger);
    }

    /**
     * @notice Cancels the challenge.
     *         Reverts if is not possible to cancel the sender's challenge for the given output index.
     *
     * @param _outputIndex Index of the L2 checkpoint output.
     */
    function cancelChallenge(uint256 _outputIndex) external {
        Types.Challenge storage challenge = challenges[_outputIndex][msg.sender];
        ChallengeStatus status = _challengeStatus(challenge);

        if (status == ChallengeStatus.NONE) revert ImproperChallengeStatus();

        if (!_cancelIfOutputDeleted(_outputIndex, challenge.challenger, status))
            revert CannotCancelChallenge();
    }

    /**
     * @notice Dismisses the challenge and rollback l2 output.
     *         This function can only be called by Security Council contract.
     *
     * @param _outputIndex      Index of the L2 checkpoint output.
     * @param _challenger       Address of the challenger.
     * @param _asserter         Address of the asserter.
     * @param _outputRoot       The L2 output root to rollback.
     * @param _publicInputHash  Hash of public input.
     */
    function dismissChallenge(
        uint256 _outputIndex,
        address _challenger,
        address _asserter,
        bytes32 _outputRoot,
        bytes32 _publicInputHash
    ) external onlySecurityCouncil outputNotFinalized(_outputIndex) {
        if (_outputRoot == DELETED_OUTPUT_ROOT) revert CannotRollbackOutputToZero();
        if (L2_ORACLE.getL2Output(_outputIndex).outputRoot != DELETED_OUTPUT_ROOT)
            revert OutputNotDeleted();
        verifiedPublicInputs[_publicInputHash] = false;

        // Rollback output root.
        L2_ORACLE.replaceL2Output(_outputIndex, _outputRoot, _asserter);

        // Switch validator system after validator pool contract terminated.
        if (L2_ORACLE.VALIDATOR_POOL().isTerminated(_outputIndex)) {
            // Unjail asserter.
            L2_ORACLE.VALIDATOR_MANAGER().tryUnjail(_asserter, true);
        }

        emit ChallengeDismissed(_outputIndex, _challenger, block.timestamp);
    }

    /**
     * @notice Deletes the L2 output root forcefully by the Security Council
     *         when zk-proving is not possible due to an undeniable bug.
     *
     * @param _outputIndex Index of the L2 checkpoint output.
     */
    function forceDeleteOutput(
        uint256 _outputIndex
    ) external onlySecurityCouncil outputNotFinalized(_outputIndex) {
        // Check if the output is deleted.
        Types.CheckpointOutput memory output = L2_ORACLE.getL2Output(_outputIndex);
        if (output.outputRoot == DELETED_OUTPUT_ROOT) revert OutputAlreadyDeleted();

        // Delete output root.
        L2_ORACLE.replaceL2Output(_outputIndex, DELETED_OUTPUT_ROOT, SECURITY_COUNCIL);

        // Switch validator system after validator pool contract terminated.
        if (L2_ORACLE.VALIDATOR_POOL().isTerminated(_outputIndex)) {
            // Slash the asseter's asset and move it to pending challenge reward for the output.
            L2_ORACLE.VALIDATOR_MANAGER().slash(_outputIndex, SECURITY_COUNCIL, output.submitter);
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
        uint8 _turn,
        bytes32 _prevFirst,
        bytes32 _prevLast,
        bytes32[] memory _segments
    ) private view {
        uint256 segLen = _segments.length;

        if (getSegmentsLength(_turn) != segLen) revert InvalidSegmentsLength();
        if (_prevFirst != _segments[0]) revert FirstSegmentMismatched();
        if (_prevLast == _segments[segLen - 1]) revert LastSegmentMatched();
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
     * @notice Updates timestamp of the challenge timeout.
     *
     * @param _challenge The challenge data to update.
     */
    function _updateTimeout(Types.Challenge storage _challenge) private {
        if (!_isAbleToBisect(_challenge)) {
            _challenge.timeoutAt = uint64(block.timestamp + PROVING_TIMEOUT);
        } else {
            _challenge.timeoutAt = uint64(block.timestamp + BISECTION_TIMEOUT);
        }
    }

    /**
     * @notice Validates and updates the lengths of segments.
     *
     * @param _segmentsLengths Lengths of segments.
     */
    function _setSegmentsLengths(uint256[] memory _segmentsLengths) private {
        // _segmentsLengths length should be an even number in order to let challenger submit
        // invalidity proof at the last turn.
        if (_segmentsLengths.length % 2 != 0) revert InvalidSegmentsLength();

        uint256 sum = 1;
        for (uint256 i = 0; i < _segmentsLengths.length; ) {
            segmentsLengths[i] = _segmentsLengths[i];
            sum = sum * (_segmentsLengths[i] - 1);

            unchecked {
                ++i;
            }
        }

        if (sum != L2_ORACLE_SUBMISSION_INTERVAL) revert InvalidSegmentsLength();
    }

    /**
     * @notice Checks if the L2ToL1MesagePasser account is included in the given state root.
     *
     * @param _merkleProof                 Merkle proof of L2ToL1MessagePasser account against the state root.
     * @param _l2ToL1MessagePasserBalance  Balance of the L2ToL1MessagePasser account.
     * @param _l2ToL1MessagePasserCodeHash Codehash of the L2ToL1MessagePasser account.
     * @param _messagePasserStorageRoot    Storage root of the L2ToL1MessagePasser account.
     * @param _stateRoot                   State root.
     */
    function _validateWithdrawalStorageRoot(
        bytes[] calldata _merkleProof,
        bytes32 _l2ToL1MessagePasserBalance,
        bytes32 _l2ToL1MessagePasserCodeHash,
        bytes32 _messagePasserStorageRoot,
        bytes32 _stateRoot
    ) private view {
        // TODO(chokobole): Can we fix the codeHash?
        bytes memory l2ToL1MessagePasserAccount = abi.encodePacked(
            uint256(0), // nonce
            _l2ToL1MessagePasserBalance, // balance,
            _l2ToL1MessagePasserCodeHash, // codeHash,
            _messagePasserStorageRoot // storage root
        );

        if (
            !IZKMerkleTrie(ZK_MERKLE_TRIE).verifyInclusionProof(
                bytes32(bytes20(Predeploys.L2_TO_L1_MESSAGE_PASSER)),
                l2ToL1MessagePasserAccount,
                _merkleProof,
                _stateRoot
            )
        ) revert InvalidInclusionProof();
    }

    /**
     * @notice Validates the output root proofs.
     *
     * @param _pos                Position of the last valid segment.
     * @param _challenge          The challenge data.
     * @param _srcOutputRoot      The source output root.
     * @param _dstOutputRoot      The destination output root.
     * @param _srcOutputRootProof Proof of the source output root.
     * @param _dstOutputRootProof Proof of the destination output root.
     */
    function _validateOutputRootProof(
        uint256 _pos,
        Types.Challenge storage _challenge,
        bytes32 _srcOutputRoot,
        bytes32 _dstOutputRoot,
        Types.OutputRootProof calldata _srcOutputRootProof,
        Types.OutputRootProof calldata _dstOutputRootProof
    ) private view {
        if (_challenge.segments[_pos] != _srcOutputRoot) revert FirstSegmentMismatched();

        // If asserter timeout, the bisection of segments may not have ended.
        // Therefore, segment validation only proceeds when bisection is not possible.
        if (!_isAbleToBisect(_challenge)) {
            if (_challenge.segments[_pos + 1] == _dstOutputRoot) revert LastSegmentMatched();
        }

        if (_srcOutputRootProof.nextBlockHash != _dstOutputRootProof.blockHash)
            revert BlockHashMismatched();
    }

    /**
     * @notice Checks if the public input is valid.
     *         Reverts if public input is invalid.
     *
     * @param _srcOutputRootProof Proof of the source output root.
     * @param _dstOutputRootProof Proof of the destination output root.
     * @param _publicInput        Ingredients to compute the public input used by ZK proof verification.
     * @param _rlps               Pre-encoded RLPs to compute the next block hash of the source output root proof.
     */
    function _validatePublicInput(
        Types.OutputRootProof calldata _srcOutputRootProof,
        Types.OutputRootProof calldata _dstOutputRootProof,
        Types.PublicInput calldata _publicInput,
        Types.BlockHeaderRLP calldata _rlps
    ) private pure {
        // TODO(chokobole): check withdrawal storage root of _dstOutputRootProof against state root of _dstOutputRootProof.
        if (_publicInput.stateRoot != _dstOutputRootProof.stateRoot) revert StateRootMismatched();

        // parentBeaconRoot is non-zero for Cancun block
        bytes32 blockHash = _publicInput.parentBeaconRoot != bytes32(0)
            ? Hashing.hashBlockHeaderCancun(_publicInput, _rlps)
            : Hashing.hashBlockHeaderShanghai(_publicInput, _rlps);

        if (_srcOutputRootProof.nextBlockHash != blockHash) revert BlockHashMismatched();
    }

    /**
     * @notice Cancels the challenge if the output root to be challenged has already been deleted.
     *         If the output root has been deleted, delete the challenge. Note that before validator
     *         system upgrade, also refund the challenger's pending bond in validator pool.
     *         Reverts when challenger is timed out or called by non-challenger.
     *
     * @param _outputIndex Index of the L2 checkpoint output.
     * @param _challenger  Address of the challenger.
     * @param _status      Current status of the challenge.
     *
     * @return Whether the challenge was canceled.
     */
    function _cancelIfOutputDeleted(
        uint256 _outputIndex,
        address _challenger,
        ChallengeStatus _status
    ) private returns (bool) {
        bytes32 outputRoot = L2_ORACLE.getL2Output(_outputIndex).outputRoot;
        if (outputRoot != DELETED_OUTPUT_ROOT) {
            return false;
        }

        // If the output is deleted, the asserter does not need to do anything further.
        if (msg.sender != _challenger) revert NotAllowedCaller();

        if (_status == ChallengeStatus.CHALLENGER_TIMEOUT) revert ImproperChallengeStatus();

        delete challenges[_outputIndex][msg.sender];
        emit ChallengeCanceled(_outputIndex, msg.sender, block.timestamp);

        // Switch validator system after validator pool contract terminated.
        if (!L2_ORACLE.VALIDATOR_POOL().isTerminated(_outputIndex)) {
            L2_ORACLE.VALIDATOR_POOL().releasePendingBond(_outputIndex, msg.sender, msg.sender);
        }

        return true;
    }

    /**
     * @notice Deletes the challenge because the challenger timed out.
     *         The winner is the asserter, and challenger loses his asset.
     *
     * @param _outputIndex Index of the L2 checkpoint output.
     * @param _challenger  Address of the challenger.
     */
    function _challengerTimeout(uint256 _outputIndex, address _challenger) private {
        delete challenges[_outputIndex][_challenger];
        emit ChallengerTimedOut(_outputIndex, _challenger, block.timestamp);

        // Switch validator system after validator pool contract terminated.
        if (L2_ORACLE.VALIDATOR_POOL().isTerminated(_outputIndex)) {
            address submitter = L2_ORACLE.getSubmitter(_outputIndex);
            L2_ORACLE.VALIDATOR_MANAGER().slash(_outputIndex, submitter, _challenger);
            return;
        }

        // After output is finalized, the challenger's bond is included in the balance of output submitter.
        if (L2_ORACLE.isFinalized(_outputIndex)) {
            address submitter = L2_ORACLE.getSubmitter(_outputIndex);
            L2_ORACLE.VALIDATOR_POOL().releasePendingBond(_outputIndex, _challenger, submitter);
        } else {
            // Because the challenger lost, the challenger's bond is included in the bond for that output.
            L2_ORACLE.VALIDATOR_POOL().increaseBond(_outputIndex, _challenger);
        }
    }

    /**
     * @notice Hashes the public input with padding dummy transactions.
     *
     * @param _prevStateRoot Previous state root.
     * @param _publicInput   Ingredients to compute the public input used by ZK proof verification.
     *
     * @return Hash of public input.
     */
    function _hashPublicInput(
        bytes32 _prevStateRoot,
        Types.PublicInput calldata _publicInput
    ) private view returns (bytes32) {
        bytes32[] memory dummyHashes;
        if (_publicInput.txHashes.length < MAX_TXS) {
            dummyHashes = Hashing.generateDummyHashes(
                DUMMY_HASH,
                MAX_TXS - _publicInput.txHashes.length
            );
        }

        // NOTE(chokobole): We cannot calculate the Ethereum transaction root solely
        // based on transaction hashes. It is necessary to have access to the original
        // transactions. Considering the imposed constraints and the difficulty
        // of providing a preimage that would generate the desired public input hash
        // from an attacker's perspective, we have decided to omit the verification
        // using the transaction root.
        return Hashing.hashPublicInput(_prevStateRoot, _publicInput, dummyHashes);
    }

    /**
     * @notice Returns the number of L2 blocks for the next turn.
     *
     * @param _challenge The current challenge data.
     *
     * @return The number of L2 blocks for the next turn.
     */
    function _nextSegSize(Types.Challenge storage _challenge) private view returns (uint256) {
        uint8 turn = _challenge.turn;
        return _challenge.segSize / (getSegmentsLength(turn) - 1);
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
     * @notice Determines if bisection is possible.
     *
     * @param _challenge The current challenge data.
     *
     * @return Whether bisection is possible.
     */
    function _isAbleToBisect(Types.Challenge storage _challenge) private view returns (bool) {
        return _nextSegSize(_challenge) > 1;
    }

    /**
     * @notice Determines if the next turn is the challenger's turn.
     *         Note that challenger turns are odd numbers and asserter turns are even numbers.
     *
     * @param _turn The current turn.
     *
     * @return Whether the next turn is the challenger's turn.
     */
    function _isNextForChallenger(uint8 _turn) private pure returns (bool) {
        // If the _turn value is even, it means that the asserter has completed its turn,
        // so the next turn will be the challenger's turn.
        return _turn % 2 == 0;
    }

    /**
     * @notice Returns status of a given challenge.
     *
     * @param _challenge The challenge data.
     *
     * @return The status of the challenge.
     */
    function _challengeStatus(
        Types.Challenge storage _challenge
    ) private view returns (ChallengeStatus) {
        if (_challenge.turn < TURN_INIT) {
            return ChallengeStatus.NONE;
        }

        bool isChallengerTurn = _isNextForChallenger(_challenge.turn);

        // Check if it's a timed out challenge.
        if (_isPast(_challenge.timeoutAt)) {
            // timeout on challenger turn
            if (isChallengerTurn) {
                return ChallengeStatus.CHALLENGER_TIMEOUT;
            }

            // If the asserter times out and the challenger does not prove fault,
            // the challenger is assumed to have timed out.
            if (_isPast(_challenge.timeoutAt + PROVING_TIMEOUT)) {
                return ChallengeStatus.CHALLENGER_TIMEOUT;
            }

            // timeout on asserter turn
            return ChallengeStatus.ASSERTER_TIMEOUT;
        }

        // If bisection is not possible, the Challenger must execute the fault proof.
        if (!_isAbleToBisect(_challenge)) {
            return ChallengeStatus.READY_TO_PROVE;
        }

        return isChallengerTurn ? ChallengeStatus.CHALLENGER_TURN : ChallengeStatus.ASSERTER_TURN;
    }

    /**
     * @notice Returns the challenge corresponding to the given L2 output index.
     *
     * @param _outputIndex Index of the L2 checkpoint output.
     * @param _challenger  Address of the challenger.
     *
     * @return The challenge data.
     */
    function getChallenge(
        uint256 _outputIndex,
        address _challenger
    ) external view returns (Types.Challenge memory) {
        return challenges[_outputIndex][_challenger];
    }

    /**
     * @notice Returns the challenge status corresponding to the given L2 output index.
     *
     * @param _outputIndex Index of the L2 checkpoint output.
     * @param _challenger  Address of the challenger.
     *
     * @return The status of the challenge.
     */
    function getStatus(
        uint256 _outputIndex,
        address _challenger
    ) external view returns (ChallengeStatus) {
        Types.Challenge storage challenge = challenges[_outputIndex][_challenger];
        return _challengeStatus(challenge);
    }

    /**
     * @notice Returns the segment length required for that turn.
     *
     * @param _turn The challenge turn.
     *
     * @return The segments length.
     */
    function getSegmentsLength(uint8 _turn) public view returns (uint256) {
        if (_turn < TURN_INIT) revert InvalidTurn();
        return segmentsLengths[_turn - 1];
    }

    /**
     * @notice Determines whether bisection is possible in the challenge corresponding to the given
     *         L2 output index.
     *
     * @param _outputIndex Index of the L2 checkpoint output.
     * @param _challenger  Address of the challenger.
     *
     * @return Whether bisection is possible.
     */
    function isAbleToBisect(uint256 _outputIndex, address _challenger) public view returns (bool) {
        Types.Challenge storage challenge = challenges[_outputIndex][_challenger];
        return _isAbleToBisect(challenge);
    }

    /**
     * @notice Determines whether current timestamp is in challenge creation period corresponding to the given L2 output index.
     *
     * @param _outputIndex Index of the L2 checkpoint output.
     *
     * @return Whether current timestamp is in challenge creation period.
     */
    function isInCreationPeriod(uint256 _outputIndex) external view returns (bool) {
        Types.CheckpointOutput memory targetOutput = L2_ORACLE.getL2Output(_outputIndex);
        return targetOutput.timestamp + CREATION_PERIOD_SECONDS >= block.timestamp;
    }
}
