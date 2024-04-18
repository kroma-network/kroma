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
     * initiated at least once, "started" means the validator has been started and added to the
     * validator tree. "MIN_REGISTER_AMOUNT" means the total assets of the validator exceeds
     * MIN_REGISTER_AMOUNT, "MIN_START_AMOUNT" means the same.
     *
     * +-------------------+-----------+---------+---------------------+------------------+
     * | Status            | initiated | started | MIN_REGISTER_AMOUNT | MIN_START_AMOUNT |
     * +-------------------+-----------+---------+---------------------+------------------+
     * | NONE              | X         | X       | X                   | X                |
     * | INACTIVE          | O         | X       | X                   | O/X              |
     * | ACTIVE            | O         | X       | O                   | X                |
     * | CAN_START         | O         | X       | O                   | O                |
     * | STARTED           | O         | O       | O                   | X                |
     * | CAN_SUBMIT_OUTPUT | O         | O       | O                   | O                |
     * +-------------------+-----------+---------+---------------------+------------------+
     */
    enum ValidatorStatus {
        NONE,
        INACTIVE,
        ACTIVE,
        CAN_START,
        STARTED,
        CAN_SUBMIT_OUTPUT
    }

    /**
     * @notice Constructs the constructor parameters of ValidatorManager contract.
     *
     * @custom:field _l2Oracle                       Address of the L2OutputOracle contract.
     * @custom:field _assetManager                   Address of the AssetManager contract.
     * @custom:field _trustedValidator               Address of the trusted validator.
     * @custom:field _commissionRateMinChangeSeconds The minimum duration to change the commission
     *                                               rate in seconds.
     * @custom:field _roundDurationSeconds           The duration of one submission round in
     *                                               seconds.
     * @custom:field _jailPeriodSeconds              The minimum duration to get out of jail in
     *                                               seconds.
     * @custom:field _jailThreshold                  The maximum allowed number of output
     *                                               non-submissions before jailed.
     * @custom:field _maxOutputFinalizations         Max number of finalized outputs.
     * @custom:field _baseReward                     Base reward for the validator.
     * @custom:field _minRegisterAmount              Minimum amount to register as a validator.
     * @custom:field _minStartAmount                 Minimum amount to start submitting outputs.
     */
    struct ConstructorParams {
        L2OutputOracle _l2Oracle;
        AssetManager _assetManager;
        address _trustedValidator;
        uint128 _commissionRateMinChangeSeconds;
        uint128 _roundDurationSeconds;
        uint128 _jailPeriodSeconds;
        uint128 _jailThreshold;
        uint128 _maxOutputFinalizations;
        uint128 _baseReward;
        uint128 _minRegisterAmount;
        uint128 _minStartAmount;
    }

    /**
     * @notice Constructs the information of a validator.
     *
     * @custom:field isInitiated             Whether the validator is initiated.
     * @custom:field noSubmissionCount       Number of counts that the validator did not submit the
     *                                       output in priority round.
     * @custom:field commissionRate          Commission rate of validator.
     * @custom:field commissionMaxChangeRate Maximum changeable commission rate at once.
     * @custom:field commissionRateChangedAt Last timestamp when the commission rate was changed.
     */
    struct Validator {
        bool isInitiated;
        uint8 noSubmissionCount;
        uint8 commissionRate;
        uint8 commissionMaxChangeRate;
        uint128 commissionRateChangedAt;
    }

    /**
     * @notice Emitted when registers as a validator.
     *
     * @param validator               Address of the validator.
     * @param started                 If the validator is started or not.
     * @param commissionRate          The commission rate the validator sets.
     * @param commissionMaxChangeRate Maximum changeable commission rate at once.
     * @param assets                  The number of assets the validator self-delegates.
     */
    event ValidatorRegistered(
        address indexed validator,
        bool started,
        uint8 commissionRate,
        uint8 commissionMaxChangeRate,
        uint128 assets
    );

    /**
     * @notice Emitted when a validator starts, which means added to the validator tree.
     *
     * @param validator Address of the validator.
     * @param startsAt  The timestamp when the validator starts.
     */
    event ValidatorStarted(address indexed validator, uint256 startsAt);

    /**
     * @notice Emitted when a validator stops, which means removed from the validator tree.
     *
     * @param validator Address of the validator.
     * @param stopsAt   The timestamp when the validator stops.
     */
    event ValidatorStopped(address indexed validator, uint256 stopsAt);

    /**
     * @notice Emitted when a validator changed commission rate.
     *
     * @param validator         Address of the validator.
     * @param oldCommissionRate The old commission rate.
     * @param newCommissionRate The new commission rate.
     */
    event ValidatorCommissionRateChanged(
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
     * @param validator       Address of the validator whose vault is rewarded.
     * @param validatorReward The amount of validator reward.
     * @param baseReward      The amount of base reward for KRO delegators.
     * @param boostedReward   The amount of boosted reward for KGH delegators.
     */
    event RewardDistributed(
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
     * @notice Reverts when the commission max change rate exceeds the max value.
     */
    error MaxCommissionChangeRateExceeded();

    /**
     * @notice Reverts when try to change commission rate with same value as previous.
     */
    error SameCommissionRate();

    /**
     * @notice Reverts when try to change commission rate beyond max change rate.
     */
    error CommissionChangeRateExceeded();

    /**
     * @notice Reverts when the min change seconds of commission has not elapsed.
     */
    error NotElapsedCommissionChangePeriod();

    /**
     * @notice Reverts when try to unjail before jail period elapsed.
     */
    error NotElapsedJailPeriod();

    /**
     * @notice Reverts if the validator is not selected priority validator.
     */
    error NotSelectedPriorityValidator();

    /**
     * @notice Registers as a validator with assets at least MIN_REGISTER_AMOUNT. The validator with
     *         assets more than MIN_START_AMOUNT can be started at the same time.
     *
     * @param assets                  The amount of assets to self-delegate.
     * @param commissionRate          The commission rate the validator sets.
     * @param commissionMaxChangeRate Maximum changeable commission rate at once.
     */
    function registerValidator(
        uint128 assets,
        uint8 commissionRate,
        uint8 commissionMaxChangeRate
    ) external;

    /**
     * @notice Starts a validator and adds the validator to validator tree. To submit outputs, the
     *         validator should start.
     */
    function startValidator() external;

    /**
     * @notice Handles some essential actions such as reward distribution, jail handling, next
     *         priority validator selection after output submission. This function can only be
     *         called by L2OutputOracle.
     *
     * @param outputIndex Index of the L2 checkpoint output submitted.
     */
    function afterSubmitL2Output(uint256 outputIndex) external;

    /**
     * @notice Changes the commission rate of a validator. An inactive validator cannot change it,
     *         and a validator can change it after COMMISION_RATE_MIN_CHANGE_SECONDS elapsed since
     *         the last changed time. Also, the validator can only make changes within the
     *         commissionMaxChangeRate that the validator set initially.
     *
     * @param newCommissionRate The new commission rate to apply.
     */
    function changeCommissionRate(uint8 newCommissionRate) external;

    /**
     * @notice Attempts to unjail a validator. Only Colosseum can set force to true, otherwise only
     *         the validator who wants to unjail can call it.
     *
     * @param validator Address of the validator.
     * @param force     To unjail forcefully or not.
     */
    function tryUnjail(address validator, bool force) external;

    /**
     * @notice Slash KRO from the vault of the challenge loser and move the slashing asset to
     *         pending challenge reward before output rewarded, after directly to winner's asset.
     *         Since the behavior could threaten the security of the chain, the loser is sent to
     *         jail.
     *
     * @param outputIndex The index of output challenged.
     * @param winner      Address of the challenge winner.
     * @param loser       Address of the challenge loser.
     */
    function slash(uint256 outputIndex, address winner, address loser) external;

    /**
     * @notice Updates the validator tree.
     *
     * @param validator Address of the validator.
     * @param tryRemove Flag to try remove the validator from validator tree.
     */
    function updateValidatorTree(address validator, bool tryRemove) external;

    /**
     * @notice Checks the eligibility to submit L2 checkpoint output during output submission.
     *         Note that only the validator whose status is CAN_SUBMIT_OUTPUT can submit output.
     *         This function can only be called by L2OutputOracle during output submission.
     *
     * @param validator Address of the output submitter.
     */
    function checkSubmissionEligibility(address validator) external view;

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
     * @notice Asserts that the given validator satisfies output submission condition.
     *
     * @param validator Address of the validator.
     */
    function assertCanSubmitOutput(address validator) external view;
}
