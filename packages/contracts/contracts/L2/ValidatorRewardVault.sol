// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { L2StandardBridge } from "../L2/L2StandardBridge.sol";
import { Predeploys } from "../libraries/Predeploys.sol";
import { SafeCall } from "../libraries/SafeCall.sol";
import { FeeVault } from "../universal/FeeVault.sol";
import { ISemver } from "../universal/ISemver.sol";
import { AddressAliasHelper } from "../vendor/AddressAliasHelper.sol";

/**
 * @custom:proxied
 * @custom:predeploy 0x4200000000000000000000000000000000000008
 * @title ValidatorRewardVault
 * @notice The ValidatorRewardVault accumulates transaction fees and pays rewards to validators.
 */
contract ValidatorRewardVault is FeeVault, ISemver {
    /**
     * @notice Address of the ValidatorPool contract on L1.
     */
    address public immutable VALIDATOR_POOL;

    /**
     * @notice A value to divide the vault balance by when determining the reward amount.
     */
    uint256 public immutable REWARD_DIVIDER;

    /**
     * @notice The reward balance that the validator is eligible to receive.
     */
    mapping(address => uint256) internal rewards;

    /**
     * @notice A mapping of whether the reward corresponding to the L2 block number has been paid.
     */
    mapping(uint256 => bool) internal isPaid;

    /**
     * @notice The amount of determined as rewards.
     */
    uint256 public totalReserved;

    /**
     * @notice Emitted when the balance of a validator has increased.
     *
     * @param validator     Address of the validator.
     * @param l2BlockNumber The L2 block number of the output root.
     * @param amount        Amount of the reward.
     */
    event Rewarded(address indexed validator, uint256 indexed l2BlockNumber, uint256 amount);

    /**
     * @notice Semantic version.
     * @custom:semver 1.0.1
     */
    string public constant version = "1.0.1";

    /**
     * @notice Constructs the ValidatorRewardVault contract.
     *
     * @param _validatorPool Address of the ValidatorPool contract on L1.
     * @param _rewardDivider A value to divide the vault balance by when determining the reward amount.
     */
    constructor(address _validatorPool, uint256 _rewardDivider) FeeVault(address(0), 0) {
        VALIDATOR_POOL = _validatorPool;
        REWARD_DIVIDER = _rewardDivider;
    }

    /**
     * @notice Rewards the validator for submitting the output.
     *         ValidatorPool contract on L1 calls this function over the portal when output is finalized.
     *
     * @param _validator     Address of the validator.
     * @param _l2BlockNumber The L2 block number of the output root.
     */
    function reward(address _validator, uint256 _l2BlockNumber) external {
        require(
            AddressAliasHelper.undoL1ToL2Alias(msg.sender) == VALIDATOR_POOL,
            "ValidatorRewardVault: function can only be called from the ValidatorPool"
        );

        require(_validator != address(0), "ValidatorRewardVault: validator address cannot be 0");

        require(
            !isPaid[_l2BlockNumber],
            "ValidatorRewardVault: the reward has already been paid for the L2 block number"
        );

        uint256 amount = _determineRewardAmount();

        unchecked {
            totalReserved += amount;
            rewards[_validator] += amount;
        }

        isPaid[_l2BlockNumber] = true;

        emit Rewarded(_validator, _l2BlockNumber, amount);
    }

    /**
     * @notice Checks if the withdrawal is possible, and returns the withdrawal amount.
     *         When a withdrawal is available, it resets the balance and updates the total processed amount.
     */
    function _processWithdrawal() internal override returns (uint256) {
        uint256 amount = rewards[msg.sender];
        require(
            amount >= MIN_WITHDRAWAL_AMOUNT,
            "ValidatorRewardVault: withdrawal amount must be greater than minimum withdrawal amount"
        );

        rewards[msg.sender] = 0;
        unchecked {
            totalReserved -= amount;
            totalProcessed += amount;
        }

        emit Withdrawal(amount, msg.sender, msg.sender);
        return amount;
    }

    /**
     * @notice Withdraws all of the sender's balance.
     *         Reverts if the balance is less than the minimum withdrawal amount.
     */
    function withdraw() external override {
        uint256 amount = _processWithdrawal();

        L2StandardBridge(payable(Predeploys.L2_STANDARD_BRIDGE)).bridgeETHTo{ value: amount }(
            msg.sender,
            WITHDRAWAL_MIN_GAS,
            bytes("")
        );
    }

    /**
     * @notice Withdraws all of the sender's balance to L2.
     *         Reverts if the balance is less than the minimum withdrawal amount.
     */
    function withdrawToL2() external override {
        uint256 amount = _processWithdrawal();

        bool success = SafeCall.call(msg.sender, gasleft(), amount, bytes(""));
        require(success, "ValidatorRewardVault: ETH transfer failed");
    }

    /**
     * @notice Withdraws all remaining contract balance excluding reserved amount to ProtocolVault.
     *         Reverts if the available withdraw amount is 0.
     *         Note that this function is added not to use this contract anymore.
     */
    function withdrawToProtocolVault() external {
        uint256 amount = address(this).balance - totalReserved;
        require(amount > 0, "ValidatorRewardVault: withdrawal amount must be greater than zero");

        unchecked {
            totalProcessed += amount;
        }

        emit Withdrawal(amount, Predeploys.PROTOCOL_VAULT, msg.sender);

        bool success = SafeCall.call(Predeploys.PROTOCOL_VAULT, gasleft(), amount, bytes(""));
        require(success, "ValidatorRewardVault: ETH transfer failed");
    }

    /**
     * @notice Determines the reward amount.
     *
     * @return Amount of the reward.
     */
    function _determineRewardAmount() internal view returns (uint256) {
        return (address(this).balance - totalReserved) / REWARD_DIVIDER;
    }

    /**
     * @notice Returns the reward balance of the given address.
     *
     * @param _addr Address to lookup.
     *
     * @return The reward balance of the given address.
     */
    function balanceOf(address _addr) external view returns (uint256) {
        return rewards[_addr];
    }
}
