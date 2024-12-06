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
     * @notice Minimum amount to register as a validator. It should be equal or more than
     *         ASSET_MANAGER.BOND_AMOUNT.
     */
    uint128 public immutable MIN_REGISTER_AMOUNT;

    /**
     * @notice Minimum amount to activate a validator and add it to the validator tree.
     *         Note that only the active validators can submit outputs.
     */
    uint128 public immutable MIN_ACTIVATE_AMOUNT;

    /**
     * @notice The delay to finalize the commission rate change of the validator (in seconds).
     */
    uint128 public immutable COMMISSION_CHANGE_DELAY_SECONDS;

    /**
     * @notice The duration of a submission round for one output (in seconds).
     *         Note that there are two submission rounds for an output: PRIORITY ROUND and PUBLIC
     *         ROUND.
     */
    uint128 public immutable ROUND_DURATION_SECONDS;

    /**
     * @notice The minimum duration to get out of jail in output non-submissions penalty (in seconds).
     */
    uint128 public immutable SOFT_JAIL_PERIOD_SECONDS;

    /**
     * @notice The maximum duration to get out of jail in slashing penalty (in seconds).
     */
    uint128 public immutable HARD_JAIL_PERIOD_SECONDS;

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
     * @notice The first output index after MPT transition, only allowed to be submitted by TRUSTED_VALIDATOR.
     *         Challenging to this output is also restricted to prevent unintended challenges from
     *         nodes that haven't upgraded.
     */
    uint256 public immutable MPT_FIRST_OUTPUT_INDEX;

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
     * @notice A modifier that only allows AssetManager contract to call.
     */
    modifier onlyAssetManager() {
        if (msg.sender != address(ASSET_MANAGER)) revert NotAllowedCaller();
        _;
    }

    /**
     * @notice A modifier that only allows TrustedValidator to call.
     */
    modifier onlyTrustedValidator() {
        if (msg.sender != TRUSTED_VALIDATOR) revert NotAllowedCaller();
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
        if (_constructorParams._minRegisterAmount > _constructorParams._minActivateAmount)
            revert InvalidConstructorParams();

        L2_ORACLE = _constructorParams._l2Oracle;
        ASSET_MANAGER = _constructorParams._assetManager;
        TRUSTED_VALIDATOR = _constructorParams._trustedValidator;
        MIN_REGISTER_AMOUNT = _constructorParams._minRegisterAmount;
        MIN_ACTIVATE_AMOUNT = _constructorParams._minActivateAmount;
        COMMISSION_CHANGE_DELAY_SECONDS = _constructorParams._commissionChangeDelaySeconds;
        // Note that this value MUST be (SUBMISSION_INTERVAL * L2_BLOCK_TIME) / 2.
        ROUND_DURATION_SECONDS = _constructorParams._roundDurationSeconds;
        SOFT_JAIL_PERIOD_SECONDS = _constructorParams._softJailPeriodSeconds;
        HARD_JAIL_PERIOD_SECONDS = _constructorParams._hardJailPeriodSeconds;
        JAIL_THRESHOLD = _constructorParams._jailThreshold;
        MAX_OUTPUT_FINALIZATIONS = _constructorParams._maxOutputFinalizations;
        BASE_REWARD = _constructorParams._baseReward;
        MPT_FIRST_OUTPUT_INDEX = _constructorParams._mptFirstOutputIndex;
    }

    /**
     * @inheritdoc IValidatorManager
     */
    function registerValidator(
        uint128 assets,
        uint8 commissionRate,
        address withdrawAccount
    ) external {
        if (msg.sender.code.length > 0 || msg.sender != tx.origin) revert NotAllowedCaller();
        if (getStatus(msg.sender) != ValidatorStatus.NONE) revert ImproperValidatorStatus();
        if (assets < MIN_REGISTER_AMOUNT) revert InsufficientAsset();
        if (commissionRate > COMMISSION_RATE_DENOM) revert MaxCommissionRateExceeded();

        Validator storage validatorInfo = _validatorInfo[msg.sender];
        validatorInfo.isInitiated = true;
        validatorInfo.commissionRate = commissionRate;

        ASSET_MANAGER.depositToRegister(msg.sender, assets, withdrawAccount);

        bool ready = assets >= MIN_ACTIVATE_AMOUNT;
        if (ready) {
            _activateValidator(msg.sender);
        }

        emit ValidatorRegistered(msg.sender, ready, commissionRate, assets);
    }

    /**
     * @inheritdoc IValidatorManager
     */
    function activateValidator() external {
        if (getStatus(msg.sender) != ValidatorStatus.READY || inJail(msg.sender))
            revert ImproperValidatorStatus();

        _activateValidator(msg.sender);
    }

    /**
     * @inheritdoc IValidatorManager
     */
    function tryActivateValidator(address validator) external onlyAssetManager {
        if (getStatus(validator) == ValidatorStatus.READY && !inJail(validator))
            _activateValidator(validator);
    }

    /**
     * @inheritdoc IValidatorManager
     */
    function afterSubmitL2Output(uint256 outputIndex) external onlyL2OutputOracle {
        _distributeReward();

        // Bond validator KRO to reserve slashing amount.
        address submitter = L2_ORACLE.getSubmitter(outputIndex);
        ASSET_MANAGER.bondValidatorKro(submitter);

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
    function initCommissionChange(uint8 newCommissionRate) external {
        if (getStatus(msg.sender) < ValidatorStatus.REGISTERED || inJail(msg.sender))
            revert ImproperValidatorStatus();

        if (newCommissionRate > COMMISSION_RATE_DENOM) revert MaxCommissionRateExceeded();

        Validator storage validatorInfo = _validatorInfo[msg.sender];
        uint8 oldCommissionRate = validatorInfo.commissionRate;
        if (newCommissionRate == oldCommissionRate) revert SameCommissionRate();

        validatorInfo.pendingCommissionRate = newCommissionRate;
        validatorInfo.commissionChangeInitiatedAt = uint128(block.timestamp);

        emit ValidatorCommissionChangeInitiated(msg.sender, oldCommissionRate, newCommissionRate);
    }

    /**
     * @inheritdoc IValidatorManager
     */
    function finalizeCommissionChange() external {
        if (getStatus(msg.sender) < ValidatorStatus.REGISTERED || inJail(msg.sender))
            revert ImproperValidatorStatus();

        uint128 canFinalizeAt = canFinalizeCommissionChangeAt(msg.sender);
        if (canFinalizeAt == COMMISSION_CHANGE_DELAY_SECONDS) revert NotInitiatedCommissionChange();
        if (block.timestamp < canFinalizeAt) revert NotElapsedCommissionChangeDelay();

        Validator storage validatorInfo = _validatorInfo[msg.sender];
        uint8 oldCommissionRate = validatorInfo.commissionRate;
        uint8 newCommissionRate = validatorInfo.pendingCommissionRate;

        validatorInfo.commissionRate = newCommissionRate;
        validatorInfo.pendingCommissionRate = 0;
        validatorInfo.commissionChangeInitiatedAt = 0;

        emit ValidatorCommissionChangeFinalized(msg.sender, oldCommissionRate, newCommissionRate);
    }

    /**
     * @inheritdoc IValidatorManager
     */
    function tryUnjail() external {
        if (!inJail(msg.sender)) revert ImproperValidatorStatus();
        if (_jail[msg.sender] > block.timestamp) revert NotElapsedJailPeriod();

        _resetNoSubmissionCount(msg.sender);
        delete _jail[msg.sender];

        emit ValidatorUnjailed(msg.sender);

        if (getStatus(msg.sender) == ValidatorStatus.READY) {
            _activateValidator(msg.sender);
        }
    }

    /**
     * @inheritdoc IValidatorManager
     */
    function bondValidatorKro(address validator) external onlyColosseum {
        ASSET_MANAGER.bondValidatorKro(validator);
    }

    /**
     * @inheritdoc IValidatorManager
     */
    function unbondValidatorKro(address validator) external onlyColosseum {
        ASSET_MANAGER.unbondValidatorKro(validator);
    }

    /**
     * @inheritdoc IValidatorManager
     */
    function slash(uint256 outputIndex, address winner, address loser) external onlyColosseum {
        uint128 challengeReward = ASSET_MANAGER.decreaseBalanceWithChallenge(loser);

        emit Slashed(outputIndex, loser, challengeReward);

        _sendToJail(loser, false);

        if (L2_ORACLE.nextFinalizeOutputIndex() <= outputIndex) {
            // If output is not rewarded yet, add slashing asset to the pending challenge reward.
            unchecked {
                _pendingChallengeReward[outputIndex] += challengeReward;
            }
        } else {
            // If output is already rewarded, add slashing asset to the winner's asset directly.
            challengeReward = ASSET_MANAGER.increaseBalanceWithChallenge(winner, challengeReward);
            updateValidatorTree(winner, false);

            emit ChallengeRewardDistributed(outputIndex, winner, challengeReward);
        }
    }

    /**
     * @inheritdoc IValidatorManager
     */
    function revertSlash(uint256 outputIndex, address loser) external onlyColosseum {
        uint128 challengeReward = ASSET_MANAGER.revertDecreaseBalanceWithChallenge(loser);
        unchecked {
            _pendingChallengeReward[outputIndex] -= challengeReward;
        }

        emit SlashReverted(outputIndex, loser, challengeReward);

        if (inJail(loser)) {
            // Revert jail expiration timestamp of the original loser.
            uint128 expiresAt = _jail[loser] - HARD_JAIL_PERIOD_SECONDS;
            if (block.timestamp < expiresAt) {
                _jail[loser] = expiresAt;

                emit ValidatorJailed(loser, expiresAt);
            } else {
                delete _jail[loser];

                emit ValidatorUnjailed(loser);

                if (getStatus(loser) == ValidatorStatus.READY) {
                    _activateValidator(loser);
                }
            }
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

        if (!isActive(validator)) revert ImproperValidatorStatus();
    }

    /**
     * @inheritdoc IValidatorManager
     */
    function checkChallengeEligibility(uint256 outputIndex) external view onlyColosseum {
        if (MPT_FIRST_OUTPUT_INDEX == outputIndex) {
            revert MptFirstOutputRestricted();
        }
    }

    /**
     * @inheritdoc IValidatorManager
     */
    function getCommissionRate(address validator) external view returns (uint8) {
        return _validatorInfo[validator].commissionRate;
    }

    /**
     * @inheritdoc IValidatorManager
     */
    function getPendingCommissionRate(address validator) external view returns (uint8) {
        return _validatorInfo[validator].pendingCommissionRate;
    }

    /**
     * @inheritdoc IValidatorManager
     */
    function activatedValidatorCount() external view returns (uint32) {
        return _validatorTree.counter - _validatorTree.removed;
    }

    /**
     * @inheritdoc IValidatorManager
     */
    function getWeight(address validator) external view returns (uint120) {
        return _validatorTree.nodes[_validatorTree.nodeMap[validator]].weight;
    }

    /**
     * @inheritdoc IValidatorManager
     */
    function jailExpiresAt(address validator) external view returns (uint128) {
        return _jail[validator];
    }

    /**
     * @inheritdoc IValidatorManager
     */
    function updateValidatorTree(address validator, bool tryRemove) public {
        ValidatorStatus status = getStatus(validator);
        if (tryRemove && (status == ValidatorStatus.EXITED || status == ValidatorStatus.INACTIVE)) {
            if (_validatorTree.remove(validator)) emit ValidatorStopped(validator, block.timestamp);
        } else if (status >= ValidatorStatus.INACTIVE) {
            _validatorTree.update(validator, uint120(ASSET_MANAGER.reflectiveWeight(validator)));
        }
    }

    /**
     * @inheritdoc IValidatorManager
     */
    function nextValidator() public view returns (address) {
        if (MPT_FIRST_OUTPUT_INDEX == L2_ORACLE.nextOutputIndex()) {
            return TRUSTED_VALIDATOR;
        }

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
            return ValidatorStatus.EXITED;
        }

        bool activated = _validatorTree.nodeMap[validator] > 0;

        if (ASSET_MANAGER.reflectiveWeight(validator) < MIN_ACTIVATE_AMOUNT) {
            if (!activated) {
                return ValidatorStatus.REGISTERED;
            }
            return ValidatorStatus.INACTIVE;
        }

        if (!activated) {
            return ValidatorStatus.READY;
        }
        return ValidatorStatus.ACTIVE;
    }

    /**
     * @inheritdoc IValidatorManager
     */
    function inJail(address validator) public view returns (bool) {
        return _jail[validator] != 0;
    }

    /**
     * @inheritdoc IValidatorManager
     */
    function isActive(address validator) public view returns (bool) {
        if (getStatus(validator) == ValidatorStatus.ACTIVE) return true;
        return false;
    }

    /**
     * @inheritdoc IValidatorManager
     */
    function noSubmissionCount(address validator) public view returns (uint8) {
        return _validatorInfo[validator].noSubmissionCount;
    }

    /**
     * @inheritdoc IValidatorManager
     */
    function canFinalizeCommissionChangeAt(address validator) public view returns (uint128) {
        return
            _validatorInfo[validator].commissionChangeInitiatedAt + COMMISSION_CHANGE_DELAY_SECONDS;
    }

    /**
     * @inheritdoc IValidatorManager
     */
    function activatedValidatorTotalWeight() public view returns (uint120) {
        return _validatorTree.nodes[_validatorTree.root].weightSum;
    }

    /**
     * @notice Private function to activate a validator and adds the validator to validator tree.
     *
     * @param validator Address of the validator.
     */
    function _activateValidator(address validator) private {
        _validatorTree.insert(validator, uint120(ASSET_MANAGER.reflectiveWeight(validator)));

        emit ValidatorActivated(validator, block.timestamp);
    }

    /**
     * @notice Private function to add output submission rewards to the vaults of finalized output
     *         submitters.
     *
     * @return Whether the reward distribution is done at least once or not.
     */
    function _distributeReward() private returns (bool) {
        uint256 outputIndex = L2_ORACLE.nextFinalizeOutputIndex();
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

                emit RewardDistributed(
                    outputIndex,
                    submitter,
                    validatorReward,
                    baseReward,
                    boostedReward
                );

                uint128 challengeReward = _pendingChallengeReward[outputIndex];
                if (challengeReward > 0) {
                    challengeReward = ASSET_MANAGER.increaseBalanceWithChallenge(
                        submitter,
                        challengeReward
                    );
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
            L2_ORACLE.setNextFinalizeOutputIndex(outputIndex);

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
     * @return The amount of base reward, excluding base reward for the validator.
     * @return The amount of boosted reward.
     * @return The amount of reward from commission and base reward for the validator.
     */
    function _calculateReward(address validator) internal view returns (uint128, uint128, uint128) {
        if (validator == ASSET_MANAGER.SECURITY_COUNCIL()) {
            return (0, 0, BASE_REWARD);
        }

        uint128 commissionRate = _validatorInfo[validator].commissionRate;
        uint128 boostedReward = _getBoostedReward(validator);
        uint128 baseReward;
        uint128 validatorReward;

        unchecked {
            validatorReward = (BASE_REWARD + boostedReward).mulDiv(
                commissionRate,
                COMMISSION_RATE_DENOM
            );
            baseReward = BASE_REWARD.mulDiv(
                COMMISSION_RATE_DENOM - commissionRate,
                COMMISSION_RATE_DENOM
            );
            boostedReward = boostedReward.mulDiv(
                COMMISSION_RATE_DENOM - commissionRate,
                COMMISSION_RATE_DENOM
            );

            uint128 validatorKro = ASSET_MANAGER.totalValidatorKro(validator);
            uint128 totalKro = ASSET_MANAGER.totalKroAssets(validator);
            uint128 validatorBaseReward = baseReward.mulDiv(validatorKro, totalKro + validatorKro);
            // Exclude the base reward for the validator from total base reward given to KRO delegators.
            baseReward -= validatorBaseReward;
            validatorReward += validatorBaseReward;
        }

        return (baseReward, boostedReward, validatorReward);
    }

    /**
     * @notice Updates next priority validator address. Validators with more delegation tokens have
     *         a higher probability of being selected. The random weight selection is based on the
     *         last finalized output root.
     */
    function _updatePriorityValidator() private {
        uint120 weightSum = activatedValidatorTotalWeight();
        uint256 nextFinalizeOutputIndex = L2_ORACLE.nextFinalizeOutputIndex();

        if (weightSum > 0 && nextFinalizeOutputIndex > 0) {
            Types.CheckpointOutput memory output = L2_ORACLE.getL2Output(
                nextFinalizeOutputIndex - 1
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
     *         submission round but did not submit the output. The period to get out of jail is
     *         SOFT_JAIL_PERIOD_SECONDS.
     */
    function _tryJail() private {
        if (_nextPriorityValidator == address(0)) return;

        if (_validatorInfo[_nextPriorityValidator].noSubmissionCount >= JAIL_THRESHOLD) {
            _sendToJail(_nextPriorityValidator, true);
        } else {
            unchecked {
                _validatorInfo[_nextPriorityValidator].noSubmissionCount++;
            }
        }
    }

    /**
     * @notice Send the given validator to the jail and remove from the validator tree.
     *
     * @param validator Address of the validator.
     * @param isSoft    Whether the jail is soft or hard.
     */
    function _sendToJail(address validator, bool isSoft) private {
        uint128 jailSeconds = isSoft ? SOFT_JAIL_PERIOD_SECONDS : HARD_JAIL_PERIOD_SECONDS;
        uint128 expiresAt = _jail[validator].max(uint128(block.timestamp)) + jailSeconds;
        _jail[validator] = expiresAt;

        emit ValidatorJailed(validator, expiresAt);

        if (_validatorTree.remove(validator)) emit ValidatorStopped(validator, block.timestamp);
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
