// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { Types } from "../../libraries/Types.sol";

interface IColosseum {
    struct ZkVmProof {
        bytes32 zkVmProgramVKey;
        bytes publicValues;
        bytes proofBytes;
    }

    enum ChallengeStatus {
        NONE,
        CHALLENGER_TURN,
        ASSERTER_TURN,
        CHALLENGER_TIMEOUT,
        ASSERTER_TIMEOUT,
        READY_TO_PROVE
    }

    error AlreadyVerifiedPublicInput();
    error AssertionAlreadyCreated();
    error CannotCancelChallenge();
    error ImproperChallengeStatus();
    error ImproperChallengeStatusToCancel();
    error ImproperValidatorStatus();
    error InvalidAddressGiven();
    error InvalidOutputGiven();
    error InvalidPublicInputHash();
    error L1Reorged();
    error NotAllowedCaller();
    error NotAllowedGenesisOutput();
    error OnlyChallengerCanCancel();
    error OutputAlreadyDeleted();
    error OutputAlreadyFinalized();
    error OutputNotDeleted();

    event AssertionCreated(
        uint256 indexed outputIndex,
        address indexed asserter,
        uint256 timestamp
    );
    event Bisected(
        uint256 indexed outputIndex,
        address indexed challenger,
        ChallengeStatus status,
        uint256 timestamp
    );
    event ChallengeCanceled(
        uint256 indexed outputIndex,
        address indexed challenger,
        uint256 timestamp
    );
    event ChallengeCreated(
        uint256 indexed outputIndex,
        address indexed asserter,
        address indexed challenger,
        uint256 timestamp
    );
    event ChallengeDismissed(
        uint256 indexed outputIndex,
        address indexed challenger,
        uint256 timestamp
    );
    event ChallengerTimedOut(
        uint256 indexed outputIndex,
        address indexed challenger,
        uint256 timestamp
    );
    event Initialized(uint8 version);
    event OutputForceDeleted(
        uint256 indexed outputIndex,
        address indexed asseter,
        uint256 timestamp
    );
    event Proven(uint256 indexed outputIndex, address indexed challenger, uint256 timestamp);
    event ReadyToProve(uint256 indexed outputIndex, address indexed challenger);

    function L2_ORACLE() external view returns (address);

    function L2_ORACLE_SUBMISSION_INTERVAL() external view returns (uint256);

    function SECURITY_COUNCIL() external view returns (address);

    function ZK_PROOF_VERIFIER() external view returns (address);

    function GUARDIAN_PERIOD() external view returns (uint256);

    function MAX_CLOCK_DURATION_SECONDS() external view returns (uint256);

    function bisect(
        uint256 _outputIndex,
        address _challenger,
        uint256 _pos,
        bytes32 _output
    ) external;

    function cancelChallenge(uint256 _outputIndex) external;

    function createAssertion(uint256 _outputIndex, address asserter) external;

    function createChallenge(
        uint256 _outputIndex,
        bytes32 _l1BlockHash,
        uint256 _l1BlockNumber
    ) external;

    function deletedOutputs(
        uint256
    )
        external
        view
        returns (address submitter, bytes32 outputRoot, uint128 timestamp, uint128 l2BlockNumber);

    function dismissChallenge(
        uint256 _outputIndex,
        address _challenger,
        address _asserter,
        bytes32 _outputRoot,
        bytes32 _publicInputHash
    ) external;

    function forceDeleteOutput(uint256 _outputIndex) external;

    function getChallenge(
        uint256 _outputIndex,
        address _challenger
    ) external view returns (Types.Challenge memory);

    function getAssertion(uint256 _outputIndex) external view returns (Types.AssertionView memory);

    function isFinalized(uint256 _outputIndex) external view returns (bool);

    function getAssertionStatus(uint256 _outputIndex) external view returns (Types.AssertionStatus);

    function initialize() external;

    function proveFaultWithZkVm(
        uint256 _outputIndex,
        uint256 _pos,
        ZkVmProof memory _zkVmProof
    ) external;

    function verifiedPublicInputs(bytes32) external view returns (bool);

    function version() external view returns (string memory);
}
