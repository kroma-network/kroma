// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { Types } from "../../libraries/Types.sol";
import { IValidatorPool } from "./IValidatorPool.sol";
import { IValidatorManager } from "./IValidatorManager.sol";

interface IL2OutputOracle {
    event Initialized(uint8 version);
    event OutputReplaced(
        uint256 indexed outputIndex,
        address indexed newSubmitter,
        bytes32 newOutputRoot
    );
    event OutputSubmitted(
        bytes32 indexed outputRoot,
        uint256 indexed l2OutputIndex,
        uint256 indexed l2BlockNumber,
        uint256 l1Timestamp
    );

    function COLOSSEUM() external view returns (address);

    function FINALIZATION_PERIOD_SECONDS() external view returns (uint256);

    function L2_BLOCK_TIME() external view returns (uint256);

    function SUBMISSION_INTERVAL() external view returns (uint256);

    function VALIDATOR_MANAGER() external view returns (IValidatorManager);

    function VALIDATOR_POOL() external view returns (IValidatorPool);

    function computeL2Timestamp(uint256 _l2BlockNumber) external view returns (uint256);

    function finalizedAt(uint256 _outputIndex) external view returns (uint256);

    function getL2Output(
        uint256 _l2OutputIndex
    ) external view returns (Types.CheckpointOutput memory);

    function getL2OutputAfter(
        uint256 _l2BlockNumber
    ) external view returns (Types.CheckpointOutput memory);

    function getL2OutputIndexAfter(uint256 _l2BlockNumber) external view returns (uint256);

    function getSubmitter(uint256 _outputIndex) external view returns (address);

    function initialize(uint256 _startingBlockNumber, uint256 _startingTimestamp) external;

    function isFinalized(uint256 _outputIndex) external view returns (bool);

    function latestBlockNumber() external view returns (uint256);

    function latestFinalizedOutputIndex() external view returns (uint256);

    function latestOutputIndex() external view returns (uint256);

    function nextBlockNumber() external view returns (uint256);

    function nextFinalizeOutputIndex() external view returns (uint256);

    function nextOutputIndex() external view returns (uint256);

    function nextOutputMinL2Timestamp() external view returns (uint256);

    function replaceL2Output(
        uint256 _l2OutputIndex,
        bytes32 _newOutputRoot,
        address _submitter
    ) external;

    function setNextFinalizeOutputIndex(uint256 _outputIndex) external;

    function startingBlockNumber() external view returns (uint256);

    function startingTimestamp() external view returns (uint256);

    function submitL2Output(
        bytes32 _outputRoot,
        uint256 _l2BlockNumber,
        bytes32 _l1BlockHash,
        uint256 _l1BlockNumber
    ) external payable;

    function version() external view returns (string memory);
}
