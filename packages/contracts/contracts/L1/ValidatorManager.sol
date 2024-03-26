// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { BalancedWeightTree } from "../libraries/BalancedWeightTree.sol";
import { Constants } from "../libraries/Constants.sol";
import { Types } from "../libraries/Types.sol";
import { ISemver } from "../universal/ISemver.sol";
import { AssetManager } from "./AssetManager.sol";
import { IValidatorManager } from "./IValidatorManager.sol";
import { L2OutputOracle } from "./L2OutputOracle.sol";

/**
 * @custom:proxied
 * @title ValidatorManager
 * @notice The ValidatorManager manages validator set and determines the next validator who can
 *         submit the checkpoint output to L2OutputOracle.
 */
contract ValidatorManager is ISemver, IValidatorManager, AssetManager {
    using BalancedWeightTree for BalancedWeightTree.Tree;

    /**
     * @notice Enum of the status of a validator.
     *
     * Below is the possible conditions of each status. "initiated" means the validator has been
     * initiated at least once, "started" means the validator has been started and added to the
     * weight tree. "MIN_REGISTER_AMOUNT" means the total assets of the validator exceeds
     * MIN_REGISTER_AMOUNT, "MIN_START_AMOUNT" means the same.
     *
     * +-------------------+-----------+---------+---------------------+------------------+
     * | Status            | initiated | started | MIN_REGISTER_AMOUNT | MIN_START_AMOUNT |
     * +-------------------+-----------+---------+---------------------+------------------+
     * | NONE              | X         | X       | X                   | X                |
     * | INACTIVE          | O         | X       | X                   | O/X              |
     * | IN_JAIL           | O         | O/X     | O/X                 | O/X              |
     * | ACTIVE            | O         | X       | O                   | X                |
     * | CAN_START         | O         | X       | O                   | O                |
     * | STARTED           | O         | O       | O                   | X                |
     * | CAN_SUBMIT_OUTPUT | O         | O       | O                   | O                |
     * +-------------------+-----------+---------+---------------------+------------------+
     */
    enum ValidatorStatus {
        NONE,
        INACTIVE,
        IN_JAIL,
        ACTIVE,
        CAN_START,
        STARTED,
        CAN_SUBMIT_OUTPUT
    }

    /**
     * @notice The address of the trusted validator.
     */
    address public immutable TRUSTED_VALIDATOR;

    /**
     * @notice The minimum duration to change the commission rate of the validator (in seconds).
     */
    uint128 public immutable COMMISSION_RATE_MIN_CHANGE_SECONDS;

    /**
     * @notice The duration of a submission round for one output (in seconds).
     *         Note that there are two submission rounds for an output: PRIORITY ROUND and PUBLIC
     *         ROUND.
     */
    uint128 public immutable ROUND_DURATION_SECONDS;

    /**
     * @notice The minimum duration to get out of jail (in seconds).
     */
    uint128 public immutable JAIL_PERIOD_SECONDS;

    /**
     * @notice Maximum allowed number of output non-submissions in priority round before the
     *         validator goes to jail.
     */
    uint128 public immutable JAIL_THRESHOLD;

    /**
     * @notice Address of the next validator with priority for submitting output.
     */
    address internal _nextPriorityValidator;

    /**
     * @notice A mapping of the jailed validator to the jail expiration timestamp.
     */
    mapping(address => uint128) internal _jail;

    /**
     * @notice A modifier that only allows L2OutputOracle contract to call.
     */
    modifier onlyL2OutputOracle() {
        require(
            msg.sender == address(L2_ORACLE),
            "ValidatorManager: only L2OutputOracle can call this function"
        );
        _;
    }

    /**
     * @notice A modifier that only allows Colosseum contract to call.
     */
    modifier onlyColosseum() {
        require(
            msg.sender == L2_ORACLE.COLOSSEUM(),
            "ValidatorManager: only Colosseum can call this function"
        );
        _;
    }

    /**
     * @notice Semantic version.
     * @custom:semver 1.0.0
     */
    string public constant version = "1.0.0";

    /**
     * @notice Constructs the ValidatorManager contract.
     *
     * @param _constructorParams              Constructor parameters for AssetManager.
     * @param _trustedValidator               Address of the trusted validator.
     * @param _commissionRateMinChangeSeconds The minimum duration to change the commission rate in
     *                                        seconds.
     * @param _roundDurationSeconds           The duration of one submission round in seconds.
     * @param _jailPeriodSeconds              The minimum duration to get out of jail in seconds.
     * @param _jailThreshold                  The maximum allowed number of output non-submissions
     *                                        before jailed.
     */
    constructor(
        ConstructorParams memory _constructorParams,
        address _trustedValidator,
        uint128 _commissionRateMinChangeSeconds,
        uint128 _roundDurationSeconds,
        uint128 _jailPeriodSeconds,
        uint128 _jailThreshold
    ) AssetManager(_constructorParams) {
        TRUSTED_VALIDATOR = _trustedValidator;
        COMMISSION_RATE_MIN_CHANGE_SECONDS = _commissionRateMinChangeSeconds;

        // Note that this value MUST be (SUBMISSION_INTERVAL * L2_BLOCK_TIME) / 2.
        ROUND_DURATION_SECONDS = _roundDurationSeconds;

        JAIL_PERIOD_SECONDS = _jailPeriodSeconds;
        JAIL_THRESHOLD = _jailThreshold;
    }

    /**
     * @inheritdoc IValidatorManager
     */
    function registerValidator(
        uint128 assets,
        uint8 commissionRate,
        uint8 commissionMaxChangeRate
    ) external {
        require(
            getStatus(msg.sender) == ValidatorStatus.NONE,
            "ValidatorManager: already initiated validator"
        );

        require(
            assets >= MIN_REGISTER_AMOUNT,
            "ValidatorManager: need to register with at least min register amount"
        );

        require(
            commissionRate <= COMMISSION_RATE_DENOM,
            "ValidatorManager: the max value of commission rate has been exceeded"
        );

        require(
            commissionMaxChangeRate <= COMMISSION_RATE_DENOM,
            "ValidatorManager: the max value of commission rate max change rate has been exceeded"
        );

        Vault storage vault = _vaults[msg.sender];
        vault.isInitiated = true;
        vault.reward.commissionRate = commissionRate;
        vault.reward.commissionMaxChangeRate = commissionMaxChangeRate;
        vault.reward.commissionRateChangedAt = uint128(block.timestamp);

        _delegate(msg.sender, msg.sender, assets, false);

        bool canStart = assets >= MIN_START_AMOUNT;
        if (canStart) {
            _validatorTree.insert(msg.sender, uint120(_reflectiveWeight(vault)));
        }

        emit ValidatorRegistered(
            msg.sender,
            canStart,
            commissionRate,
            commissionMaxChangeRate,
            assets
        );
    }

    /**
     * @inheritdoc IValidatorManager
     */
    function startValidator() external {
        require(
            getStatus(msg.sender) == ValidatorStatus.CAN_START,
            "ValidatorManager: validator start condition is not met"
        );

        _validatorTree.insert(msg.sender, uint120(_reflectiveWeight(_vaults[msg.sender])));

        emit ValidatorStarted(msg.sender, block.timestamp);
    }

    /**
     * @inheritdoc IValidatorManager
     */
    function afterSubmitL2Output(uint256 outputIndex) external onlyL2OutputOracle {
        _distributeReward();

        address submitter = L2_ORACLE.getSubmitter(outputIndex);
        if (submitter == _nextPriorityValidator) {
            _resetNoSubmissionCount(submitter);
        } else {
            _tryJail();
        }

        // Select the next priority validator.
        _updatePriorityValidator();
    }

    /**
     * @inheritdoc IValidatorManager
     */
    function changeCommissionRate(uint8 newCommissionRate) external {
        require(
            getStatus(msg.sender) > ValidatorStatus.INACTIVE,
            "ValidatorManager: cannot change commission rate of inactive validator"
        );

        Reward storage reward = _vaults[msg.sender].reward;

        require(
            reward.commissionRateChangedAt + COMMISSION_RATE_MIN_CHANGE_SECONDS <= block.timestamp,
            "ValidatorManager: min change seconds of commission rate has not elapsed"
        );

        require(
            newCommissionRate <= COMMISSION_RATE_DENOM,
            "ValidatorManager: the max value of commission rate has been exceeded"
        );

        uint8 oldCommissionRate = reward.commissionRate;
        require(
            newCommissionRate != oldCommissionRate,
            "ValidatorManager: cannot change to the same value"
        );

        uint8 changeRange;
        if (newCommissionRate > oldCommissionRate) {
            changeRange = newCommissionRate - oldCommissionRate;
        } else {
            changeRange = oldCommissionRate - newCommissionRate;
        }
        require(
            changeRange <= reward.commissionMaxChangeRate,
            "ValidatorManager: max change rate of commission rate has been exceeded"
        );

        reward.commissionRate = newCommissionRate;
        reward.commissionRateChangedAt = uint128(block.timestamp);

        emit ValidatorCommissionRateChanged(msg.sender, oldCommissionRate, newCommissionRate);
    }

    /**
     * @inheritdoc IValidatorManager
     */
    function tryUnjail() external {
        require(getStatus(msg.sender) == ValidatorStatus.IN_JAIL, "ValidatorManager: not in jail");

        require(
            _jail[msg.sender] <= block.timestamp,
            "ValidatorManager: jail period has not elasped"
        );

        delete _jail[msg.sender];

        _resetNoSubmissionCount(msg.sender);

        emit ValidatorUnjailed(msg.sender);
    }

    /**
     * @inheritdoc IValidatorManager
     */
    function slash(address loser, uint256 outputIndex) external onlyColosseum {
        _slash(loser, outputIndex);
    }

    /**
     * @inheritdoc IValidatorManager
     */
    function checkSubmissionEligibility(address validator) external view onlyL2OutputOracle {
        address _nextValidator = nextValidator();

        if (_nextValidator != Constants.VALIDATOR_PUBLIC_ROUND_ADDRESS) {
            require(
                validator == _nextValidator,
                "ValidatorManager: only the next selected validator can submit output"
            );
        }

        require(
            getStatus(validator) == ValidatorStatus.CAN_SUBMIT_OUTPUT,
            "ValidatorManager: validator should satisfy the condition to submit output"
        );
    }

    /**
     * @notice Returns the commission rate of given validator.
     *
     * @param validator Address of the validator.
     *
     * @return The commission rate of given validator.
     */
    function getCommissionRate(address validator) external view returns (uint8) {
        return _vaults[validator].reward.commissionRate;
    }

    /**
     * @notice Returns the commission max change rate of given validator.
     *
     * @param validator Address of the validator.
     *
     * @return The commission max change rate of given validator.
     */
    function getCommissionMaxChangeRate(address validator) external view returns (uint8) {
        return _vaults[validator].reward.commissionMaxChangeRate;
    }

    /**
     * @notice Returns the number of started validators.
     *
     * @return The number of started validators.
     */
    function startedValidatorCount() external view returns (uint32) {
        return _validatorTree.counter - _validatorTree.removed;
    }

    /**
     * @notice Returns the jail expiration timestamp of given validator.
     *
     * @param validator Address of the jailed validator.
     *
     * @return The jail expiration timestamp of given validator.
     */
    function jailExpiresAt(address validator) external view returns (uint128) {
        return _jail[validator];
    }

    /**
     * @inheritdoc IValidatorManager
     */
    function nextValidator() public view returns (address) {
        if (_nextPriorityValidator != address(0)) {
            uint256 l2Timestamp = L2_ORACLE.nextOutputMinL2Timestamp();
            if (block.timestamp >= l2Timestamp) {
                uint256 elapsed = block.timestamp - l2Timestamp;
                // If the current time exceeds one round time, it is a public round.
                if (elapsed > ROUND_DURATION_SECONDS) {
                    return Constants.VALIDATOR_PUBLIC_ROUND_ADDRESS;
                }
            }

            return _nextPriorityValidator;
        } else {
            // Only trusted validator can submit output right after the switch of validator system.
            return TRUSTED_VALIDATOR;
        }
    }

    /**
     * @notice Returns the status of the validator corresponding to the given address.
     *
     * @param validator Address of the validator.
     *
     * @return The status of the validator corresponding to the given address.
     */
    function getStatus(address validator) public view returns (ValidatorStatus) {
        if (!_vaults[validator].isInitiated) {
            return ValidatorStatus.NONE;
        }

        if (_jail[validator] != 0) {
            return ValidatorStatus.IN_JAIL;
        }

        if (_vaults[validator].asset.validatorKro < MIN_REGISTER_AMOUNT) {
            return ValidatorStatus.INACTIVE;
        }

        bool started = _validatorTree.nodeMap[validator] > 0;

        // To prevent all MIN_START_AMOUNT is fulfilled with KRO in KGH which is not subject to slash,
        // enable to start the validator when real asset satisfies the threshold.
        if (
            _reflectiveWeight(_vaults[validator]) - _vaults[validator].asset.totalKroInKgh <
            MIN_START_AMOUNT
        ) {
            if (!started) {
                return ValidatorStatus.ACTIVE;
            }
            return ValidatorStatus.STARTED;
        }

        if (!started) {
            return ValidatorStatus.CAN_START;
        }
        return ValidatorStatus.CAN_SUBMIT_OUTPUT;
    }

    /**
     * @notice Returns the no submission count of given validator.
     *
     * @param validator Address of the validator.
     *
     * @return The no submission count of given validator.
     */
    function noSubmissionCount(address validator) public view returns (uint8) {
        return _vaults[validator].noSubmissionCount;
    }

    /**
     * @notice Returns the total weight of started validators.
     *
     * @return The total weight of started validators.
     */
    function startedValidatorTotalWeight() public view returns (uint120) {
        return _validatorTree.nodes[_validatorTree.root].weightSum;
    }

    /**
     * @notice Returns the weight of given validator. It not started, returns 0.
     *         Note that `weight / startedValidatorTotalWeight()` is the probability that the
     *         validator is selected as a priority validator.
     *
     * @param validator Address of the validator.
     *
     * @return The weight of given validator.
     */
    function getWeight(address validator) public view returns (uint120) {
        return _validatorTree.nodes[_validatorTree.nodeMap[validator]].weight;
    }

    /**
     * @notice Updates next priority validator address. Validators with more delegation tokens have
     *         a higher probability of being selected. The random weight selection is based on the
     *         last finalized output root.
     */
    function _updatePriorityValidator() private {
        uint120 weightSum = startedValidatorTotalWeight();
        uint256 latestFinalizedOutputIndex = L2_ORACLE.latestFinalizedOutputIndex();

        if (weightSum > 0 && latestFinalizedOutputIndex > 0) {
            Types.CheckpointOutput memory output = L2_ORACLE.getL2Output(
                latestFinalizedOutputIndex
            );

            uint120 weight = uint120(
                uint256(
                    keccak256(
                        abi.encodePacked(
                            output.outputRoot,
                            block.number,
                            block.coinbase,
                            block.difficulty,
                            blockhash(block.number - 1)
                        )
                    )
                )
            ) % weightSum;

            _nextPriorityValidator = _validatorTree.select(weight);
        } else {
            _nextPriorityValidator = address(0);
        }
    }

    /**
     * @notice Attempts to jail a validator who was selected as a priority validator for this
     *         submission round but did not submit the output.
     */
    function _tryJail() private {
        if (_nextPriorityValidator != address(0)) {
            if (_vaults[_nextPriorityValidator].noSubmissionCount >= JAIL_THRESHOLD) {
                uint128 expiresAt = uint128(block.timestamp + JAIL_PERIOD_SECONDS);
                _jail[_nextPriorityValidator] = expiresAt;

                emit ValidatorJailed(_nextPriorityValidator, expiresAt);
            } else {
                unchecked {
                    _vaults[_nextPriorityValidator].noSubmissionCount++;
                }
            }
        }
    }

    /**
     * @notice Attempts to reset non-submission count of a validator.
     *
     * @param validator Address of the validator.
     */
    function _resetNoSubmissionCount(address validator) private {
        if (noSubmissionCount(validator) > 0) {
            _vaults[validator].noSubmissionCount = 0;
        }
    }
}
