// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { AssetManager } from "../AssetManager.sol";
import { L2OutputOracle } from "../L2OutputOracle.sol";

/**
 * @title IValidatorManager
 * @notice Interface for ValidatorManager contract.
 */
interface IValidatorManager {
    /**
     * @notice Enum of the status of a validator.
     *
     * Below is the possible conditions of each status. "initiated" means the validator has been
     * initiated at least once, "activated" means the validator has been activated and added to the
     * validator tree. "MIN_REGISTER_AMOUNT" means the total assets of the validator exceeds
     * MIN_REGISTER_AMOUNT, "MIN_ACTIVATE_AMOUNT" means the same.
     *
     * +------------+-----------+-----------+---------------------+---------------------+
     * | Status     | initiated | activated | MIN_REGISTER_AMOUNT | MIN_ACTIVATE_AMOUNT |
     * +------------+-----------+-----------+---------------------+---------------------+
     * | NONE       | X         | X         | X                   | X                   |
     * | EXITED     | O         | O/X       | X                   | O/X                 |
     * | REGISTERED | O         | X         | O                   | X                   |
     * | READY      | O         | X         | O                   | O                   |
     * | INACTIVE   | O         | O         | O                   | X                   |
     * | ACTIVE     | O         | O         | O                   | O                   |
     * +------------+-----------+-----------+---------------------+---------------------+
     */
    enum ValidatorStatus {
        NONE,
        EXITED,
        REGISTERED,
        READY,
        INACTIVE,
        ACTIVE
    }

    /**
     * @notice Constructs the constructor parameters of ValidatorManager contract.
     *
     * @custom:field _l2Oracle                     Address of the L2OutputOracle contract.
     * @custom:field _assetManager                 Address of the AssetManager contract.
     * @custom:field _trustedValidator             Address of the trusted validator.
     * @custom:field _commissionChangeDelaySeconds The delay to finalize the commission rate change
     *                                             in seconds.
     * @custom:field _roundDurationSeconds         The duration of one submission round in seconds.
     * @custom:field _softJailPeriodSeconds        The minimum duration to get out of jail in
     *                                             seconds in output non-submissions penalty.
     * @custom:field _hardJailPeriodSeconds        The minimum duration to get out of jail in
     *                                             seconds in slashing penalty.
     * @custom:field _jailThreshold                The maximum allowed number of output
     *                                             non-submissions before jailed.
     * @custom:field _maxOutputFinalizations       Max number of finalized outputs.
     * @custom:field _baseReward                   Base reward for the validator.
     * @custom:field _minRegisterAmount            Minimum amount to register as a validator.
     * @custom:field _minActivateAmount            Minimum amount to activate a validator.
     * @custom:field _mptFirstOutputIndex          First output index after the MPT transition.
                                                   Only TrustedValidator is allowed to submit output.
                                                   Challenges for this outputIndex are also restricted.
     */
    struct ConstructorParams {
        L2OutputOracle _l2Oracle;
        AssetManager _assetManager;
        address _trustedValidator;
        uint128 _commissionChangeDelaySeconds;
        uint128 _roundDurationSeconds;
        uint128 _softJailPeriodSeconds;
        uint128 _hardJailPeriodSeconds;
        uint128 _jailThreshold;
        uint128 _maxOutputFinalizations;
        uint128 _baseReward;
        uint128 _minRegisterAmount;
        uint128 _minActivateAmount;
        uint256 _mptFirstOutputIndex;
    }

    /**
     * @notice Constructs the information of a validator.
     *
     * @custom:field isInitiated                 Whether the validator is initiated.
     * @custom:field noSubmissionCount           Number of counts that the validator did not submit
     *                                           the output in priority round.
     * @custom:field commissionRate              Commission rate of validator.
     * @custom:field pendingCommissionRate       Pending commission rate of validator.
     * @custom:field commissionChangeInitiatedAt Timestamp of commission change initialization.
     */
    struct Validator {
        bool isInitiated;
        uint8 noSubmissionCount;
        uint8 commissionRate;
        uint8 pendingCommissionRate;
        uint128 commissionChangeInitiatedAt;
    }

    /**
     * @notice Emitted when registers as a validator.
     *
     * @param validator      Address of the validator.
     * @param activated      If the validator is activated or not.
     * @param commissionRate The commission rate the validator sets.
     * @param assets         The number of assets the validator deposits.
     */
    event ValidatorRegistered(
        address indexed validator,
        bool activated,
        uint8 commissionRate,
        uint128 assets
    );

    /**
     * @notice Emitted when a validator activated, which means added to the validator tree.
     *
     * @param validator   Address of the validator.
     * @param activatedAt The timestamp when the validator activated.
     */
    event ValidatorActivated(address indexed validator, uint256 activatedAt);

    /**
     * @notice Emitted when a validator stops, which means removed from the validator tree.
     *
     * @param validator Address of the validator.
     * @param stopsAt   The timestamp when the validator stops.
     */
    event ValidatorStopped(address indexed validator, uint256 stopsAt);

    /**
     * @notice Emitted when a validator initiated commission rate change.
     *
     * @param validator         Address of the validator.
     * @param oldCommissionRate The old commission rate.
     * @param newCommissionRate The new commission rate.
     */
    event ValidatorCommissionChangeInitiated(
        address indexed validator,
        uint8 oldCommissionRate,
        uint8 newCommissionRate
    );

    /**
     * @notice Emitted when a validator finalized commission rate change.
     *
     * @param validator         Address of the validator.
     * @param oldCommissionRate The old commission rate.
     * @param newCommissionRate The new commission rate.
     */
    event ValidatorCommissionChangeFinalized(
        address indexed validator,
        uint8 oldCommissionRate,
        uint8 newCommissionRate
    );

    /**
     * @notice Emitted when a validator is jailed.
     *
     * @param validator Address of the validator.
     * @param expiresAt The expiration timestamp of the jail.
     */
    event ValidatorJailed(address indexed validator, uint128 expiresAt);

    /**
     * @notice Emitted when a validator is unjailed.
     *
     * @param validator Address of the validator.
     */
    event ValidatorUnjailed(address indexed validator);

    /**
     * @notice Emitted when the output reward is distributed.
     *
     * @param outputIndex     Index of the L2 checkpoint output.
     * @param validator       Address of the validator whose vault is rewarded.
     * @param validatorReward The amount of validator reward.
     * @param baseReward      The amount of base reward for KRO delegators.
     * @param boostedReward   The amount of boosted reward for KGH delegators.
     */
    event RewardDistributed(
        uint256 indexed outputIndex,
        address indexed validator,
        uint128 validatorReward,
        uint128 baseReward,
        uint128 boostedReward
    );

    /**
     * @notice Emitted when challenge reward for challenge winner is distributed.
     *
     * @param outputIndex Index of the L2 checkpoint output.
     * @param recipient   Address of the reward recipient.
     * @param amount      The amount of challenge reward.
     */
    event ChallengeRewardDistributed(
        uint256 indexed outputIndex,
        address indexed recipient,
        uint128 amount
    );

    /**
     * @notice Emitted when the validator is slashed.
     *
     * @param outputIndex Index of the L2 checkpoint output.
     * @param loser       Address of the challenge loser.
     * @param amount      The amount of KRO slashed.
     */
    event Slashed(uint256 indexed outputIndex, address indexed loser, uint128 amount);

    /**
     * @notice Emitted when the slash is reverted.
     *
     * @param outputIndex Index of the L2 checkpoint output.
     * @param loser       Address of the challenge original loser.
     * @param amount      The amount of KRO refunded to the loser.
     */
    event SlashReverted(uint256 indexed outputIndex, address indexed loser, uint128 amount);

    /**
     * @notice Reverts when caller is not allowed.
     */
    error NotAllowedCaller();

    /**
     * @notice Reverts when constructor parameters are invalid.
     */
    error InvalidConstructorParams();

    /**
     * @notice Reverts when the status of validator is improper.
     */
    error ImproperValidatorStatus();

    /**
     * @notice Reverts when the asset is insufficient.
     */
    error InsufficientAsset();

    /**
     * @notice Reverts when the commission rate exceeds the max value.
     */
    error MaxCommissionRateExceeded();

    /**
     * @notice Reverts when try to change commission rate with same value as previous.
     */
    error SameCommissionRate();

    /**
     * @notice Reverts when the commission rate change has not been initiated.
     */
    error NotInitiatedCommissionChange();

    /**
     * @notice Reverts when the delay of commission rate change finalization has not elapsed.
     */
    error NotElapsedCommissionChangeDelay();

    /**
     * @notice Reverts when try to unjail before jail period elapsed.
     */
    error NotElapsedJailPeriod();

    /**
     * @notice Reverts if the validator is not selected priority validator.
     */
    error NotSelectedPriorityValidator();

    /**
     * @notice Reverts if the output index is restricted since it's first output after MPT transition.
     */
    error MptFirstOutputRestricted();

    /**
     * @notice Registers as a validator with assets at least MIN_REGISTER_AMOUNT. The validator with
     *         assets more than MIN_ACTIVATE_AMOUNT can be activated at the same time.
     *
     * @param assets          The amount of assets to deposit.
     * @param commissionRate  The commission rate the validator sets.
     * @param withdrawAccount An account where assets can be withdrawn to. Only this account can
     *                        withdraw the assets.
     */
    function registerValidator(
        uint128 assets,
        uint8 commissionRate,
        address withdrawAccount
    ) external;

    /**
     * @notice Activates a validator and adds the validator to validator tree. To submit outputs,
     *         the validator should be activated.
     */
    function activateValidator() external;

    /**
     * @notice Tries to activate a validator and adds the validator to validator tree. To submit
     *         outputs, the validator should be activated. This function can only be called by
     *         AssetManager.
     *
     * @param validator Address of the validator.
     */
    function tryActivateValidator(address validator) external;

    /**
     * @notice Handles some essential actions such as reward distribution, jail handling, next
     *         priority validator selection after output submission. This function can only be
     *         called by L2OutputOracle.
     *
     * @param outputIndex Index of the L2 checkpoint output submitted.
     */
    function afterSubmitL2Output(uint256 outputIndex) external;

    /**
     * @notice Initiates the commission rate change of a validator. An exited or jailed validator
     *         cannot initiate it.
     *
     * @param newCommissionRate The new commission rate to apply.
     */
    function initCommissionChange(uint8 newCommissionRate) external;

    /**
     * @notice Finalizes the commission rate change of a validator. An exited or jailed validator
     *         cannot finalize it, and a validator can finalize it after
     *         COMMISION_CHANGE_DELAY_SECONDS elapsed since the initialization of commission change.
     */
    function finalizeCommissionChange() external;

    /**
     * @notice Attempts to unjail a validator. Only the validator who wants to unjail can call
     *         itself.
     */
    function tryUnjail() external;

    /**
     * @notice Call ASSET_MANAGER.bondValidatorKro(). This function is only called by the Colosseum
     *         contract.
     *
     * @param validator Address of the validator.
     */
    function bondValidatorKro(address validator) external;

    /**
     * @notice Call ASSET_MANAGER.unbondValidatorKro(). This function is only called by the
     *         Colosseum contract.
     *
     * @param validator Address of the validator.
     */
    function unbondValidatorKro(address validator) external;

    /**
     * @notice Slash KRO from the vault of the challenge loser and move the slashing asset to
     *         pending challenge reward before output rewarded, after directly to winner's asset.
     *         Since the behavior could threaten the security of the chain, the loser is sent to
     *         jail for HARD_JAIL_PERIOD_SECONDS. This function is only called by the Colosseum
     *         contract.
     *
     * @param outputIndex The index of output challenged.
     * @param winner      Address of the challenge winner.
     * @param loser       Address of the challenge loser.
     */
    function slash(uint256 outputIndex, address winner, address loser) external;

    /**
     * @notice Revert slash. This function is only called by the Colosseum contract.
     *
     * @param outputIndex The index of output challenged.
     * @param loser       Address of the challenge loser.
     */
    function revertSlash(uint256 outputIndex, address loser) external;

    /**
     * @notice Updates the validator tree.
     *
     * @param validator Address of the validator.
     * @param tryRemove Flag to try remove the validator from validator tree.
     */
    function updateValidatorTree(address validator, bool tryRemove) external;

    /**
     * @notice Returns the no submission count of given validator.
     *
     * @param validator Address of the validator.
     *
     * @return The no submission count of given validator.
     */
    function noSubmissionCount(address validator) external view returns (uint8);

    /**
     * @notice Returns the commission rate of given validator.
     *
     * @param validator Address of the validator.
     *
     * @return The commission rate of given validator.
     */
    function getCommissionRate(address validator) external view returns (uint8);

    /**
     * @notice Returns the pending commission rate of given validator.
     *
     * @param validator Address of the validator.
     *
     * @return The pending commission rate of given validator.
     */
    function getPendingCommissionRate(address validator) external view returns (uint8);

    /**
     * @notice Returns when commission change of given validator can be finalized.
     *
     * @param validator Address of the validator.
     *
     * @return When commission change of given validator can be finalized.
     */
    function canFinalizeCommissionChangeAt(address validator) external view returns (uint128);

    /**
     * @notice Checks the eligibility to submit L2 checkpoint output during output submission.
     *         Note that only the validator whose status is ACTIVE can submit output. This function
     *         can only be called by L2OutputOracle during output submission.
     *
     * @param validator Address of the output submitter.
     */
    function checkSubmissionEligibility(address validator) external view;

    /**
     * @notice Checks the eligibility to create challenge to this outputIndex.
     *
     * @param outputIndex Index of the L2 checkpoint output.
     */
    function checkChallengeEligibility(uint256 outputIndex) external view;

    /**
     * @notice Determines who can submit the L2 checkpoint output for the current round.
     *
     * @return Address of the validator who can submit the L2 checkpoint output for the current
     *         round.
     */
    function nextValidator() external view returns (address);

    /**
     * @notice Returns the status of the validator corresponding to the given address.
     *
     * @param validator Address of the validator.
     *
     * @return The status of the validator corresponding to the given address.
     */
    function getStatus(address validator) external view returns (ValidatorStatus);

    /**
     * @notice Returns if the given validator is in jail or not.
     *
     * @param validator Address of the validator.
     *
     * @return If the given validator is in jail or not.
     */
    function inJail(address validator) external view returns (bool);

    /**
     * @notice Returns the jail expiration timestamp of given validator.
     *
     * @param validator Address of the jailed validator.
     *
     * @return The jail expiration timestamp of given validator.
     */
    function jailExpiresAt(address validator) external view returns (uint128);

    /**
     * @notice Returns if the status of the given validator is active.
     *
     * @param validator Address of the validator.
     *
     * @return If the status of the given validator is active.
     */
    function isActive(address validator) external view returns (bool);

    /**
     * @notice Returns the weight of given validator. It not activated, returns 0.
     *         Note that `weight / activatedValidatorTotalWeight()` is the probability that the
     *         validator is selected as a priority validator.
     *
     * @param validator Address of the validator.
     *
     * @return The weight of given validator.
     */
    function getWeight(address validator) external view returns (uint120);

    /**
     * @notice Returns the number of activated validators.
     *
     * @return The number of activated validators.
     */
    function activatedValidatorCount() external view returns (uint32);

    /**
     * @notice Returns the total weight of activated validators.
     *
     * @return The total weight of activated validators.
     */
    function activatedValidatorTotalWeight() external view returns (uint120);
}
