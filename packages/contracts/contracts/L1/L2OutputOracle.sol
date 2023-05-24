// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { Initializable } from "@openzeppelin/contracts/proxy/utils/Initializable.sol";

import { Types } from "../libraries/Types.sol";
import { Semver } from "../universal/Semver.sol";
import { ValidatorPool } from "./ValidatorPool.sol";

/**
 * @custom:proxied
 * @title L2OutputOracle
 * @notice The L2OutputOracle contains an array of L2 state outputs, where each output is a
 *         commitment to the state of the L2 chain. Other contracts like the KromaPortal use
 *         these outputs to verify information about the state of L2.
 */
contract L2OutputOracle is Initializable, Semver {
    /**
     * @notice The interval in L2 blocks at which checkpoints must be submitted. Although this is
     *         immutable, it can safely be modified by upgrading the implementation contract.
     */
    uint256 public immutable SUBMISSION_INTERVAL;

    /**
     * @notice The time between L2 blocks in seconds. Once set, this value MUST NOT be modified.
     */
    uint256 public immutable L2_BLOCK_TIME;

    /**
     * @notice The address of the colosseum contract. Can be updated via upgrade.
     */
    address public immutable COLOSSEUM;

    /**
     * @notice The address of the validator pool contract. Can be updated via upgrade.
     */
    ValidatorPool public immutable VALIDATOR_POOL;

    /**
     * @notice Minimum time (in seconds) that must elapse before a withdrawal can be finalized.
     */
    uint256 public immutable FINALIZATION_PERIOD_SECONDS;

    /**
     * @notice The number of the first L2 block recorded in this contract.
     */
    uint256 public startingBlockNumber;

    /**
     * @notice The timestamp of the first L2 block recorded in this contract.
     */
    uint256 public startingTimestamp;

    /**
     * @notice Array of L2 checkpoint outputs.
     */
    Types.CheckpointOutput[] internal l2Outputs;

    /**
     * @notice Emitted when an output is submitted.
     *
     * @param outputRoot    The output root.
     * @param l2OutputIndex The index of the output in the l2Outputs array.
     * @param l2BlockNumber The L2 block number of the output root.
     * @param l1Timestamp   The L1 timestamp when submitted.
     */
    event OutputSubmitted(
        bytes32 indexed outputRoot,
        uint256 indexed l2OutputIndex,
        uint256 indexed l2BlockNumber,
        uint256 l1Timestamp
    );

    /**
     * @notice Emitted when an output is replaced.
     *
     * @param outputIndex   Replaced L2 output index.
     * @param newOutputRoot L2 output root after replacement.
     */
    event OutputReplaced(uint256 indexed outputIndex, bytes32 newOutputRoot);

    /**
     * @custom:semver 0.1.0
     *
     * @param _submissionInterval        Interval in blocks at which checkpoints must be submitted.
     * @param _l2BlockTime               The time per L2 block, in seconds.
     * @param _startingBlockNumber       The number of the first L2 block.
     * @param _startingTimestamp         The timestamp of the first L2 block.
     * @param _validatorPool             The address of the ValidatorPool contract.
     * @param _colosseum                 The address of the Colosseum contract.
     * @param _finalizationPeriodSeconds Output finalization time in seconds.
     */
    constructor(
        uint256 _submissionInterval,
        uint256 _l2BlockTime,
        uint256 _startingBlockNumber,
        uint256 _startingTimestamp,
        ValidatorPool _validatorPool,
        address _colosseum,
        uint256 _finalizationPeriodSeconds
    ) Semver(0, 1, 0) {
        require(_l2BlockTime > 0, "L2OutputOracle: L2 block time must be greater than 0");
        require(
            _submissionInterval > _l2BlockTime,
            "L2OutputOracle: submission interval must be greater than L2 block time"
        );

        SUBMISSION_INTERVAL = _submissionInterval;
        L2_BLOCK_TIME = _l2BlockTime;
        VALIDATOR_POOL = _validatorPool;
        COLOSSEUM = _colosseum;
        FINALIZATION_PERIOD_SECONDS = _finalizationPeriodSeconds;

        initialize(_startingBlockNumber, _startingTimestamp);
    }

    /**
     * @notice Initializer.
     *
     * @param _startingBlockNumber Block number for the first recorded L2 block.
     * @param _startingTimestamp   Timestamp for the first recorded L2 block.
     */
    function initialize(uint256 _startingBlockNumber, uint256 _startingTimestamp)
        public
        initializer
    {
        require(
            _startingTimestamp <= block.timestamp,
            "L2OutputOracle: starting L2 timestamp must be less than current time"
        );

        startingTimestamp = _startingTimestamp;
        startingBlockNumber = _startingBlockNumber;
    }

    /**
     * @notice Replaces the output that corresponds to the given output index.
     *         Only the Colosseum contract can replace an output.
     *
     * @param _l2OutputIndex Index of the L2 output to be replaced.
     * @param _newOutputRoot The L2 output root to replace the existing one.
     * @param _submitter     Address of the L2 output submitter.
     */
    function replaceL2Output(
        uint256 _l2OutputIndex,
        bytes32 _newOutputRoot,
        address _submitter
    ) external {
        require(
            msg.sender == COLOSSEUM,
            "L2OutputOracle: only the colosseum contract can replace an output"
        );

        require(_submitter != address(0), "L2OutputOracle: submitter address cannot be zero");

        // Make sure we're not *increasing* the length of the array.
        require(
            _l2OutputIndex < l2Outputs.length,
            "L2OutputOracle: cannot replace an output after the latest output index"
        );

        // Do not allow replacing any outputs that have already been finalized.
        require(
            block.timestamp - l2Outputs[_l2OutputIndex].timestamp < FINALIZATION_PERIOD_SECONDS,
            "L2OutputOracle: cannot replace an output that has already been finalized"
        );

        l2Outputs[_l2OutputIndex].outputRoot = _newOutputRoot;
        l2Outputs[_l2OutputIndex].submitter = _submitter;

        emit OutputReplaced(_l2OutputIndex, _newOutputRoot);
    }

    /**
     * @notice Accepts an outputRoot and the timestamp of the corresponding L2 block. The timestamp
     *         must be equal to the current value returned by `nextTimestamp()` in order to be
     *         accepted. This function may only be called by the validator.
     *
     * @param _outputRoot    The L2 output of the checkpoint block.
     * @param _l2BlockNumber The L2 block number that resulted in _outputRoot.
     * @param _l1BlockHash   A block hash which must be included in the current chain.
     * @param _l1BlockNumber The block number with the specified block hash.
     * @param _bondAmount    Amount of bond.
     */
    function submitL2Output(
        bytes32 _outputRoot,
        uint256 _l2BlockNumber,
        bytes32 _l1BlockHash,
        uint256 _l1BlockNumber,
        uint256 _bondAmount
    ) external payable {
        require(msg.sender == VALIDATOR_POOL.nextValidator(), "L2OutputOracle: not your turn");

        require(
            _l2BlockNumber == nextBlockNumber(),
            "L2OutputOracle: block number must be equal to next expected block number"
        );

        require(
            computeL2Timestamp(_l2BlockNumber) < block.timestamp,
            "L2OutputOracle: cannot submit L2 output in the future"
        );

        require(
            _outputRoot != bytes32(0),
            "L2OutputOracle: L2 checkpoint output cannot be the zero hash"
        );

        if (_l1BlockHash != bytes32(0)) {
            // This check allows the validator to submit an output based on a given L1 block,
            // without fear that it will be reorged out.
            // It will also revert if the blockheight provided is more than 256 blocks behind the
            // chain tip (as the hash will return as zero). This does open the door to a griefing
            // attack in which the validator's submission is censored until the block is no longer
            // retrievable, if the validator is experiencing this attack it can simply leave out the
            // blockhash value, and delay submission until it is confident that the L1 block is
            // finalized.
            require(
                blockhash(_l1BlockNumber) == _l1BlockHash,
                "L2OutputOracle: block hash does not match the hash at the expected height"
            );
        }

        uint256 outputIndex = nextOutputIndex();

        l2Outputs.push(
            Types.CheckpointOutput({
                submitter: msg.sender,
                outputRoot: _outputRoot,
                timestamp: uint128(block.timestamp),
                l2BlockNumber: uint128(_l2BlockNumber)
            })
        );

        emit OutputSubmitted(_outputRoot, outputIndex, _l2BlockNumber, block.timestamp);

        VALIDATOR_POOL.createBond(
            outputIndex,
            uint128(_bondAmount),
            uint128(block.timestamp + FINALIZATION_PERIOD_SECONDS)
        );
    }

    /**
     * @notice Returns an output by index. Exists because Solidity's array access will return a
     *         tuple instead of a struct.
     *
     * @param _l2OutputIndex Index of the output to return.
     *
     * @return The output at the given index.
     */
    function getL2Output(uint256 _l2OutputIndex)
        external
        view
        returns (Types.CheckpointOutput memory)
    {
        return l2Outputs[_l2OutputIndex];
    }

    /**
     * @notice Returns the index of the L2 output that checkpoints a given L2 block number. Uses a
     *         binary search to find the first output greater than or equal to the given block.
     *
     * @param _l2BlockNumber L2 block number to find a checkpoint for.
     *
     * @return Index of the first checkpoint that commits to the given L2 block number.
     */
    function getL2OutputIndexAfter(uint256 _l2BlockNumber) public view returns (uint256) {
        // Make sure an output for this block number has actually been submitted.
        require(
            _l2BlockNumber <= latestBlockNumber(),
            "L2OutputOracle: cannot get output for a block that has not been submitted"
        );

        // Make sure there's at least one output submitted.
        require(
            l2Outputs.length > 0,
            "L2OutputOracle: cannot get output as no outputs have been submitted yet"
        );

        // Find the output via binary search, guaranteed to exist.
        uint256 lo = 0;
        uint256 hi = l2Outputs.length;
        while (lo < hi) {
            uint256 mid = (lo + hi) / 2;
            if (l2Outputs[mid].l2BlockNumber < _l2BlockNumber) {
                lo = mid + 1;
            } else {
                hi = mid;
            }
        }

        return lo;
    }

    /**
     * @notice Returns the L2 checkpoint output that checkpoints a given L2 block number. Uses a
     *         binary search to find the first output greater than or equal to the given block.
     *
     * @param _l2BlockNumber L2 block number to find a checkpoint for.
     *
     * @return First checkpoint that commits to the given L2 block number.
     */
    function getL2OutputAfter(uint256 _l2BlockNumber)
        external
        view
        returns (Types.CheckpointOutput memory)
    {
        return l2Outputs[getL2OutputIndexAfter(_l2BlockNumber)];
    }

    /**
     * @notice Returns the number of outputs that have been submitted. Will revert if no outputs
     *         have been submitted yet.
     *
     * @return The number of outputs that have been submitted.
     */
    function latestOutputIndex() external view returns (uint256) {
        return l2Outputs.length - 1;
    }

    /**
     * @notice Returns the index of the next output to be submitted.
     *
     * @return The index of the next output to be submitted.
     */
    function nextOutputIndex() public view returns (uint256) {
        return l2Outputs.length;
    }

    /**
     * @notice Returns the block number of the latest submitted L2 checkpoint output. If no outputs
     *         have been submitted yet then this function will return the starting block number.
     *
     * @return Latest submitted L2 block number.
     */
    function latestBlockNumber() public view returns (uint256) {
        return
            l2Outputs.length == 0
                ? startingBlockNumber
                : l2Outputs[l2Outputs.length - 1].l2BlockNumber;
    }

    /**
     * @notice Computes the block number of the next L2 block that needs to be checkpointed.
     *
     * @return Next L2 block number.
     */
    function nextBlockNumber() public view returns (uint256) {
        return latestBlockNumber() + SUBMISSION_INTERVAL;
    }

    /**
     * @notice Returns the L2 timestamp corresponding to a given L2 block number.
     *
     * @param _l2BlockNumber The L2 block number of the target block.
     *
     * @return L2 timestamp of the given block.
     */
    function computeL2Timestamp(uint256 _l2BlockNumber) public view returns (uint256) {
        return startingTimestamp + ((_l2BlockNumber - startingBlockNumber) * L2_BLOCK_TIME);
    }

    /**
     * @notice Returns the address of the L2 output submitter.
     *
     * @param _outputIndex Index of the output.
     *
     * @return Address of the submitter.
     */
    function getSubmitter(uint256 _outputIndex) external view returns (address) {
        return l2Outputs[_outputIndex].submitter;
    }
}
