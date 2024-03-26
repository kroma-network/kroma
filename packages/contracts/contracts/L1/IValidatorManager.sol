// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

/**
 * @title IValidatorManager
 * @notice Interface for ValidatorManager contract.
 */
interface IValidatorManager {
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
     * @notice Slash KRO at the vault of the challenge loser.
     *
     * @param loser       Address of the challenge loser.
     * @param outputIndex The index of output challenged.
     */
    function slash(address loser, uint256 outputIndex) external;
}
