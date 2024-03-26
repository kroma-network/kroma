// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

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
     * @param commissionMaxChangeRate The max change rate of the commission the validator sets.
     * @param assets                  The number of assets the validator self-delegates.
     */
    event ValidatorRegistered(
        address indexed validator,
        bool indexed started,
        uint8 commissionRate,
        uint8 commissionMaxChangeRate,
        uint128 assets
    );

    /**
     * @notice Emitted when a validator starts.
     *
     * @param validator Address of the validator.
     * @param startsAt  The timestamp when the validator starts.
     */
    event ValidatorStarted(address indexed validator, uint256 startsAt);

    /**
     * @notice Emitted when a validator changed commission rate.
     *
     * @param validator         Address of the validator.
     * @param oldCommissionRate The old commission rate.
     * @param newCommissionRate The new commission rate.
     */
    event ValidatorCommissionRateChanged(
        address validator,
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
    event ValidatorUnjailed(address validator);

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
     * @notice Emitted when pending challenge reward for challenge winner is distributed.
     *
     * @param recipient Address of the reward recipient.
     * @param amount    The amount of challenge reward.
     */
    event ChallengeRewardDistributed(address indexed recipient, uint128 amount);

    /**
     * @notice Emitted when the validator is slashed.
     *
     * @param loser  Address of the loser at the challenge.
     * @param winner Address of the winner at the challenge.
     * @param amount The amount of KRO slashed.
     */
    event Slashed(address indexed loser, address indexed winner, uint128 amount);

    /**
     * @notice Registers as a validator with assets at least MIN_REGISTER_AMOUNT. The validator with
     *         assets more than MIN_START_AMOUNT can be started at the same time.
     *
     * @param assets                  The amount of assets to self-delegate.
     * @param commissionRate          The commission rate the validator sets.
     * @param commissionMaxChangeRate The max change rate of commission the validator sets.
     */
    function registerValidator(
        uint128 assets,
        uint8 commissionRate,
        uint8 commissionMaxChangeRate
    ) external;

    /**
     * @notice Starts a validator and adds the validator to weight tree. To submit outputs, the
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
     * @notice Attempts to unjail a validator. Only the validator who wants to unjail can call.
     */
    function tryUnjail() external;

    /**
     * @notice Updates the weight tree of the validator.
     *
     * @param validator Address of the validator.
     * @param tryRemove Flag to try remove the validator from weight tree.
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
}
