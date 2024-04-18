// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { Math } from "@openzeppelin/contracts/utils/math/Math.sol";

import { Atan2 } from "../libraries/Atan2.sol";
import { BalancedWeightTree } from "../libraries/BalancedWeightTree.sol";
import { Constants } from "../libraries/Constants.sol";
import { Types } from "../libraries/Types.sol";
import { Uint128Math } from "../libraries/Uint128Math.sol";
import { ISemver } from "../universal/ISemver.sol";
import { AssetManager } from "./AssetManager.sol";
import { IValidatorManager } from "./interfaces/IValidatorManager.sol";
import { L2OutputOracle } from "./L2OutputOracle.sol";

/**
 * @custom:proxied
 * @title ValidatorManager
 * @notice The ValidatorManager manages validator set and determines the next validator who can
 *         submit the checkpoint output to L2OutputOracle.
 */
contract ValidatorManager is ISemver, IValidatorManager {
    using BalancedWeightTree for BalancedWeightTree.Tree;
    using Uint128Math for uint128;
    using Math for uint256;

    /**
     * @notice The denominator for the commission rate.
     */
    uint128 public constant COMMISSION_RATE_DENOM = 100;

    /**
     * @notice The numerator for the boosted reward.
     */
    uint128 public constant BOOSTED_REWARD_NUMERATOR = 40;

    /**
     * @notice The denominator for the boosted reward.
     */
    uint128 public constant BOOSTED_REWARD_DENOM = 100;

    /**
     * @notice Address of the L2OutputOracle contract. Can be updated via upgrade.
     */
    L2OutputOracle public immutable L2_ORACLE;

    /**
     * @notice The address of AssetManager contract. Can be updated via upgrade.
     */
    AssetManager public immutable ASSET_MANAGER;

    /**
     * @notice The address of the trusted validator.
     */
    address public immutable TRUSTED_VALIDATOR;

    /**
     * @notice Minimum amount to register as a validator.
     */
    uint128 public immutable MIN_REGISTER_AMOUNT;

    /**
     * @notice Minimum amount to start a validator and add it to the validator tree.
     *         Note that only the started validators can submit outputs.
     */
    uint128 public immutable MIN_START_AMOUNT;

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
     * @notice The max number of outputs to be finalized at once when distributing rewards.
     */
    uint128 public immutable MAX_OUTPUT_FINALIZATIONS;

    /**
     * @notice Amount of base reward for the validator.
     */
    uint128 public immutable BASE_REWARD;

    /**
     * @notice Address of the next validator with priority for submitting output.
     */
    address internal _nextPriorityValidator;

    /**
     * @notice Weighted tree to store and calculate the probability to be selected as an output submitter.
     */
    BalancedWeightTree.Tree internal _validatorTree;

    /**
     * @notice A mapping of the validator to the validator information.
     */
    mapping(address => Validator) internal _validatorInfo;

    /**
     * @notice A mapping of the jailed validator to the jail expiration timestamp.
     */
    mapping(address => uint128) internal _jail;

    /**
     * @notice A mapping of output index challenged successfully to pending challenge rewards.
     */
    mapping(uint256 => uint128) internal _pendingChallengeReward;

    /**
     * @notice A modifier that only allows L2OutputOracle contract to call.
     */
    modifier onlyL2OutputOracle() {
        if (msg.sender != address(L2_ORACLE)) revert NotAllowedCaller();
        _;
    }

    /**
     * @notice A modifier that only allows Colosseum contract to call.
     */
    modifier onlyColosseum() {
        if (msg.sender != L2_ORACLE.COLOSSEUM()) revert NotAllowedCaller();
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
     * @param _constructorParams The constructor parameters.
     */
    constructor(ConstructorParams memory _constructorParams) {
        if (_constructorParams._minRegisterAmount > _constructorParams._minStartAmount)
            revert InvalidConstructorParams();

        L2_ORACLE = _constructorParams._l2Oracle;
        ASSET_MANAGER = _constructorParams._assetManager;
        TRUSTED_VALIDATOR = _constructorParams._trustedValidator;
        MIN_REGISTER_AMOUNT = _constructorParams._minRegisterAmount;
        MIN_START_AMOUNT = _constructorParams._minStartAmount;
        COMMISSION_RATE_MIN_CHANGE_SECONDS = _constructorParams._commissionRateMinChangeSeconds;
        // Note that this value MUST be (SUBMISSION_INTERVAL * L2_BLOCK_TIME) / 2.
        ROUND_DURATION_SECONDS = _constructorParams._roundDurationSeconds;
        JAIL_PERIOD_SECONDS = _constructorParams._jailPeriodSeconds;
        JAIL_THRESHOLD = _constructorParams._jailThreshold;
        MAX_OUTPUT_FINALIZATIONS = _constructorParams._maxOutputFinalizations;
        BASE_REWARD = _constructorParams._baseReward;
    }

    /**
     * @inheritdoc IValidatorManager
     */
    function registerValidator(
        uint128 assets,
        uint8 commissionRate,
        uint8 commissionMaxChangeRate
    ) external {
        if (getStatus(msg.sender) != ValidatorStatus.NONE) revert ImproperValidatorStatus();

        if (assets < MIN_REGISTER_AMOUNT) revert InsufficientAsset();

        if (commissionRate > COMMISSION_RATE_DENOM) revert MaxCommissionRateExceeded();

        if (commissionMaxChangeRate > COMMISSION_RATE_DENOM)
            revert MaxCommissionChangeRateExceeded();

        Validator storage validatorInfo = _validatorInfo[msg.sender];
        validatorInfo.isInitiated = true;
        validatorInfo.commissionRate = commissionRate;
        validatorInfo.commissionMaxChangeRate = commissionMaxChangeRate;
        validatorInfo.commissionRateChangedAt = uint128(block.timestamp);

        ASSET_MANAGER.delegateToRegister(msg.sender, assets);

        bool canStart = assets >= MIN_START_AMOUNT;
        if (canStart) {
            _validatorTree.insert(msg.sender, uint120(ASSET_MANAGER.reflectiveWeight(msg.sender)));
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
        if (getStatus(msg.sender) != ValidatorStatus.CAN_START || inJail(msg.sender))
            revert ImproperValidatorStatus();

        _validatorTree.insert(msg.sender, uint120(ASSET_MANAGER.reflectiveWeight(msg.sender)));

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
        if (getStatus(msg.sender) < ValidatorStatus.ACTIVE || inJail(msg.sender))
            revert ImproperValidatorStatus();

        Validator storage validatorInfo = _validatorInfo[msg.sender];

        if (
            validatorInfo.commissionRateChangedAt + COMMISSION_RATE_MIN_CHANGE_SECONDS >
            block.timestamp
        ) revert NotElapsedCommissionChangePeriod();

        if (newCommissionRate > COMMISSION_RATE_DENOM) revert MaxCommissionRateExceeded();

        uint8 oldCommissionRate = validatorInfo.commissionRate;
        if (newCommissionRate == oldCommissionRate) revert SameCommissionRate();

        uint8 changeRange = newCommissionRate > oldCommissionRate
            ? newCommissionRate - oldCommissionRate
            : oldCommissionRate - newCommissionRate;
        if (changeRange > validatorInfo.commissionMaxChangeRate)
            revert CommissionChangeRateExceeded();

        validatorInfo.commissionRate = newCommissionRate;
        validatorInfo.commissionRateChangedAt = uint128(block.timestamp);

        emit ValidatorCommissionRateChanged(msg.sender, oldCommissionRate, newCommissionRate);
    }

    /**
     * @inheritdoc IValidatorManager
     */
    function tryUnjail(address validator, bool force) external {
        if (!inJail(validator)) revert ImproperValidatorStatus();

        if (force) {
            if (msg.sender != L2_ORACLE.COLOSSEUM()) revert NotAllowedCaller();
        } else {
            if (msg.sender != validator) revert NotAllowedCaller();
            if (_jail[validator] > block.timestamp) revert NotElapsedJailPeriod();

            _resetNoSubmissionCount(validator);
        }

        delete _jail[validator];

        emit ValidatorUnjailed(validator);
    }

    /**
     * @inheritdoc IValidatorManager
     */
    function slash(uint256 outputIndex, address winner, address loser) external onlyColosseum {
        uint128 amountToSlash = ASSET_MANAGER.modifyBalanceWithSlashing(loser, 0, true);
        updateValidatorTree(loser, true);

        emit Slashed(outputIndex, loser, amountToSlash);

        _sendToJail(loser);

        if (L2_ORACLE.latestFinalizedOutputIndex() < outputIndex) {
            // If output is not rewarded yet, add slashing asset to the pending challenge reward.
            unchecked {
                _pendingChallengeReward[outputIndex] += amountToSlash;
            }
        } else {
            // If output is already rewarded, add slashing asset to the winner's asset directly.
            ASSET_MANAGER.modifyBalanceWithSlashing(winner, amountToSlash, false);
            updateValidatorTree(winner, false);

            emit ChallengeRewardDistributed(outputIndex, winner, amountToSlash);
        }
    }

    /**
     * @inheritdoc IValidatorManager
     */
    function checkSubmissionEligibility(address validator) external view onlyL2OutputOracle {
        address _nextValidator = nextValidator();
        if (
            _nextValidator != Constants.VALIDATOR_PUBLIC_ROUND_ADDRESS &&
            validator != _nextValidator
        ) revert NotSelectedPriorityValidator();

        _assertCanSubmitOutput(validator);
    }

    /**
     * @inheritdoc IValidatorManager
     */
    function assertCanSubmitOutput(address validator) external view {
        _assertCanSubmitOutput(validator);
    }

    /**
     * @notice Returns the commission rate of given validator.
     *
     * @param validator Address of the validator.
     *
     * @return The commission rate of given validator.
     */
    function getCommissionRate(address validator) external view returns (uint8) {
        return _validatorInfo[validator].commissionRate;
    }

    /**
     * @notice Returns the commission max change rate of given validator.
     *
     * @param validator Address of the validator.
     *
     * @return The commission max change rate of given validator.
     */
    function getCommissionMaxChangeRate(address validator) external view returns (uint8) {
        return _validatorInfo[validator].commissionMaxChangeRate;
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
     * @notice Returns the weight of given validator. It not started, returns 0.
     *         Note that `weight / startedValidatorTotalWeight()` is the probability that the
     *         validator is selected as a priority validator.
     *
     * @param validator Address of the validator.
     *
     * @return The weight of given validator.
     */
    function getWeight(address validator) external view returns (uint120) {
        return _validatorTree.nodes[_validatorTree.nodeMap[validator]].weight;
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
    function updateValidatorTree(address validator, bool tryRemove) public {
        ValidatorStatus status = getStatus(validator);
        if (tryRemove && status == ValidatorStatus.STARTED) {
            _validatorTree.remove(validator);
            emit ValidatorStopped(validator, block.timestamp);
        } else if (status >= ValidatorStatus.STARTED) {
            _validatorTree.update(validator, uint120(ASSET_MANAGER.reflectiveWeight(validator)));
        }
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
            return TRUSTED_VALIDATOR;
        }
    }

    /**
     * @inheritdoc IValidatorManager
     */
    function getStatus(address validator) public view returns (ValidatorStatus) {
        if (!_validatorInfo[validator].isInitiated) {
            return ValidatorStatus.NONE;
        }

        if (ASSET_MANAGER.totalValidatorKro(validator) < MIN_REGISTER_AMOUNT) {
            return ValidatorStatus.INACTIVE;
        }

        bool started = _validatorTree.nodeMap[validator] > 0;

        // To prevent all MIN_START_AMOUNT is fulfilled with KRO in KGH which is not subject to slash,
        // enable to start the validator when real asset satisfies the threshold.
        if (
            ASSET_MANAGER.reflectiveWeight(validator) - ASSET_MANAGER.totalKroInKgh(validator) <
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
     * @inheritdoc IValidatorManager
     */
    function inJail(address validator) public view returns (bool) {
        return _jail[validator] != 0;
    }

    /**
     * @notice Returns the no submission count of given validator.
     *
     * @param validator Address of the validator.
     *
     * @return The no submission count of given validator.
     */
    function noSubmissionCount(address validator) public view returns (uint8) {
        return _validatorInfo[validator].noSubmissionCount;
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
     * @notice Private function to add output submission rewards to the vaults of finalized output
     *         submitters.
     *
     * @return Whether the reward distribution is done at least once or not.
     */
    function _distributeReward() private returns (bool) {
        uint256 outputIndex = L2_ORACLE.latestFinalizedOutputIndex() + 1;
        uint256 latestOutputIndex = L2_ORACLE.latestOutputIndex();

        if (!L2_ORACLE.VALIDATOR_POOL().isTerminated(outputIndex)) {
            return false;
        }

        uint128 finalizedOutputNum = 0;
        address submitter;

        while (finalizedOutputNum < MAX_OUTPUT_FINALIZATIONS && outputIndex <= latestOutputIndex) {
            if (L2_ORACLE.isFinalized(outputIndex)) {
                submitter = L2_ORACLE.getSubmitter(outputIndex);

                (
                    uint128 baseReward,
                    uint128 boostedReward,
                    uint128 validatorReward
                ) = _calculateReward(submitter);

                ASSET_MANAGER.increaseBalanceWithReward(
                    submitter,
                    baseReward,
                    boostedReward,
                    validatorReward
                );

                emit RewardDistributed(submitter, validatorReward, baseReward, boostedReward);

                uint128 challengeReward = _pendingChallengeReward[outputIndex];
                if (challengeReward > 0) {
                    ASSET_MANAGER.modifyBalanceWithSlashing(submitter, challengeReward, false);
                    delete _pendingChallengeReward[outputIndex];

                    emit ChallengeRewardDistributed(outputIndex, submitter, challengeReward);
                }

                updateValidatorTree(submitter, false);

                unchecked {
                    ++outputIndex;
                    ++finalizedOutputNum;
                }
            } else {
                break;
            }
        }

        if (finalizedOutputNum > 0) {
            L2_ORACLE.setLatestFinalizedOutputIndex(outputIndex - 1);

            return true;
        }

        return false;
    }

    /**
     * @notice Internal function to get the boosted reward with the number of KGH.
     *
     * @param validator Address of the validator.
     *
     * @return The boosted reward with the number of KGH.
     */
    function _getBoostedReward(address validator) internal view returns (uint128) {
        uint128 numKgh = ASSET_MANAGER.totalKghNum(validator);
        uint128 coefficient = BASE_REWARD.mulDiv(BOOSTED_REWARD_NUMERATOR, BOOSTED_REWARD_DENOM);
        return uint128(Atan2.atan2(numKgh, 100).mulDiv(coefficient, 1 << 40));
    }

    /**
     * @notice Internal function to calculate the reward of the validator when distributing reward.
     *
     * @param validator Address of the validator.
     *
     * @return The amount of base reward.
     * @return The amount of boosted reward.
     * @return The amount of validator reward.
     */
    function _calculateReward(address validator) internal view returns (uint128, uint128, uint128) {
        uint128 commissionRate = _validatorInfo[validator].commissionRate;
        uint128 boostedReward = _getBoostedReward(validator);
        uint128 baseReward;
        uint128 validatorReward;

        unchecked {
            baseReward = BASE_REWARD.mulDiv(
                COMMISSION_RATE_DENOM - commissionRate,
                COMMISSION_RATE_DENOM
            );
            boostedReward = boostedReward.mulDiv(
                COMMISSION_RATE_DENOM - commissionRate,
                COMMISSION_RATE_DENOM
            );
            validatorReward = (BASE_REWARD + boostedReward).mulDiv(
                commissionRate,
                COMMISSION_RATE_DENOM
            );
        }

        return (baseReward, boostedReward, validatorReward);
    }

    /**
     * @notice Internal function to assert that the given validator satisfies output submission
     *         condition.
     *
     * @param validator Address of the validator.
     */
    function _assertCanSubmitOutput(address validator) internal view {
        if (getStatus(validator) != ValidatorStatus.CAN_SUBMIT_OUTPUT || inJail(validator))
            revert ImproperValidatorStatus();
    }

    /**
     * @notice Updates next priority validator address. Validators with more delegation tokens have
     *         a higher probability of being selected. The random weight selection is based on the
     *         last finalized output root.
     */
    function _updatePriorityValidator() private {
        uint120 weightSum = startedValidatorTotalWeight();

        if (weightSum > 0) {
            uint256 latestFinalizedOutputIndex = L2_ORACLE.latestFinalizedOutputIndex();
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
            if (_validatorInfo[_nextPriorityValidator].noSubmissionCount >= JAIL_THRESHOLD) {
                _sendToJail(_nextPriorityValidator);
            } else {
                unchecked {
                    _validatorInfo[_nextPriorityValidator].noSubmissionCount++;
                }
            }
        }
    }

    /**
     * @notice Send the given validator to the jail. If the validator is already in jail, the
     *         expiration timestamp is updated.
     *
     * @param validator Address of the validator.
     */
    function _sendToJail(address validator) private {
        uint128 expiresAt = uint128(block.timestamp + JAIL_PERIOD_SECONDS);
        _jail[validator] = expiresAt;

        emit ValidatorJailed(validator, expiresAt);
    }

    /**
     * @notice Attempts to reset non-submission count of a validator.
     *
     * @param validator Address of the validator.
     */
    function _resetNoSubmissionCount(address validator) private {
        if (noSubmissionCount(validator) > 0) {
            _validatorInfo[validator].noSubmissionCount = 0;
        }
    }
}
