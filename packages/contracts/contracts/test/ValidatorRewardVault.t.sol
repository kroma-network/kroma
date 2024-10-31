// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { Predeploys } from "../libraries/Predeploys.sol";
import { Types } from "../libraries/Types.sol";
import { ValidatorRewardVault } from "../L2/ValidatorRewardVault.sol";
import { StandardBridge } from "../universal/StandardBridge.sol";
import { AddressAliasHelper } from "../vendor/AddressAliasHelper.sol";
import { Bridge_Initializer } from "./CommonTest.t.sol";

contract ValidatorRewardVault_Test is Bridge_Initializer {
    ValidatorRewardVault internal vault =
        ValidatorRewardVault(payable(Predeploys.VALIDATOR_REWARD_VAULT));
    address internal constant recipient = address(256);
    uint256 internal constant l2BlockNumber = 1800;

    uint256 internal vaultBalance;
    uint256 internal rewardDivider = 0;

    event Rewarded(address indexed validator, uint256 indexed l2BlockNumber, uint256 amount);

    event Withdrawal(uint256 value, address to, address from);

    function setUp() public override {
        super.setUp();

        rewardDivider = oracle.FINALIZATION_PERIOD_SECONDS() / (submissionInterval * l2BlockTime);

        vm.etch(
            Predeploys.VALIDATOR_REWARD_VAULT,
            address(new ValidatorRewardVault(address(pool), rewardDivider)).code
        );
        vm.label(Predeploys.VALIDATOR_REWARD_VAULT, "ValidatorRewardVault");

        vaultBalance = rewardDivider * (vault.MIN_WITHDRAWAL_AMOUNT() + 10 ether);
    }

    function test_minWithdrawalAmount_succeeds() external {
        assertEq(vault.MIN_WITHDRAWAL_AMOUNT(), 0);
    }

    function test_constructor_succeeds() external {
        assertEq(vault.RECIPIENT(), address(0));
        assertEq(vault.VALIDATOR_POOL(), address(pool));
        assertEq(vault.REWARD_DIVIDER(), rewardDivider);
    }

    function test_receive_succeeds() external {
        uint256 balance = address(vault).balance;

        vm.prank(alice);
        (bool success, ) = address(vault).call{ value: 100 }(hex"");

        assertEq(success, true);
        assertEq(address(vault).balance, balance + 100);
    }

    function test_reward_succeeds() external {
        vm.deal(address(vault), vaultBalance);

        uint256 reserved = vault.totalReserved();
        uint256 balance = vault.balanceOf(recipient);
        uint256 rewardAmount = (address(vault).balance - reserved) / rewardDivider;

        vm.expectEmit(true, true, false, false, address(Predeploys.VALIDATOR_REWARD_VAULT));
        emit Rewarded(recipient, l2BlockNumber, rewardAmount);
        vm.prank(AddressAliasHelper.applyL1ToL2Alias(address(pool)));
        vault.reward(recipient, l2BlockNumber);
        // Check the balance was increased.
        assertEq(vault.balanceOf(recipient), balance + rewardAmount);
        assertEq(vault.totalReserved(), reserved + rewardAmount);
    }

    function test_reward_senderNotValidatorPool_reverts() external {
        vm.expectRevert("ValidatorRewardVault: function can only be called from the ValidatorPool");
        vault.reward(address(0), 0);
    }

    function test_reward_zeroValidatorAddress_reverts() external {
        vm.expectRevert("ValidatorRewardVault: validator address cannot be 0");
        vm.prank(AddressAliasHelper.applyL1ToL2Alias(address(pool)));
        vault.reward(address(0), 0);
    }

    function test_reward_alreadyPaidBlockNumber_reverts() external {
        vm.prank(AddressAliasHelper.applyL1ToL2Alias(address(pool)));
        vault.reward(recipient, l2BlockNumber);

        vm.expectRevert(
            "ValidatorRewardVault: the reward has already been paid for the L2 block number"
        );
        vm.prank(AddressAliasHelper.applyL1ToL2Alias(address(pool)));
        vault.reward(recipient, l2BlockNumber);
    }

    function test_withdraw_succeeds() external {
        vm.deal(address(vault), vaultBalance);
        vm.prank(AddressAliasHelper.applyL1ToL2Alias(address(pool)));
        vault.reward(recipient, l2BlockNumber);

        uint256 amount = vault.balanceOf(recipient);
        uint256 reserved = vault.totalReserved();

        // No ether has been withdrawn yet
        assertEq(vault.totalProcessed(), 0);

        vm.expectEmit(true, true, true, true, address(Predeploys.VALIDATOR_REWARD_VAULT));
        emit Withdrawal(amount, recipient, recipient);

        // The entire vault's balance is withdrawn
        vm.expectCall(
            Predeploys.L2_STANDARD_BRIDGE,
            amount,
            abi.encodeWithSelector(
                StandardBridge.bridgeETHTo.selector,
                recipient,
                35_000,
                bytes("")
            )
        );

        vm.prank(recipient);
        vault.withdraw();
        // The withdrawal was successful
        assertEq(vault.totalProcessed(), amount);
        // Check the total determined reward amount was decreased.
        assertEq(vault.totalReserved(), reserved - amount);
    }

    function test_withdrawL2_succeeds() external {
        vm.deal(address(vault), vaultBalance);
        vm.prank(AddressAliasHelper.applyL1ToL2Alias(address(pool)));
        vault.reward(recipient, l2BlockNumber);

        uint256 amount = vault.balanceOf(recipient);
        uint256 reserved = vault.totalReserved();

        // No ether has been withdrawn yet
        assertEq(vault.totalProcessed(), 0);

        vm.expectEmit(true, true, true, true, address(Predeploys.VALIDATOR_REWARD_VAULT));
        emit Withdrawal(amount, recipient, recipient);

        uint256 prevBalance = recipient.balance;
        vm.prank(recipient);
        vault.withdrawToL2();
        // The withdrawal was successful
        assertEq(vault.totalProcessed(), amount);
        // Check the total determined reward amount was decreased.
        assertEq(vault.totalReserved(), reserved - amount);

        assertEq(recipient.balance, prevBalance + amount);
    }

    function test_withdrawToProtocolVault_succeeds() external {
        vm.deal(address(vault), vaultBalance);
        vm.prank(AddressAliasHelper.applyL1ToL2Alias(address(pool)));
        vault.reward(recipient, l2BlockNumber);

        uint256 reserved = vault.totalReserved();
        uint256 amount = vaultBalance - reserved;

        // No ether has been withdrawn yet
        assertEq(vault.totalProcessed(), 0);
        assertGt(reserved, 0);

        vm.expectEmit(true, true, true, true, address(Predeploys.VALIDATOR_REWARD_VAULT));
        emit Withdrawal(amount, Predeploys.PROTOCOL_VAULT, recipient);

        uint256 prevBalance = Predeploys.PROTOCOL_VAULT.balance;
        vm.prank(recipient);
        vault.withdrawToProtocolVault();

        // The withdrawal was successful
        assertEq(vault.totalProcessed(), amount);
        assertEq(vault.totalReserved(), reserved);
        assertEq(Predeploys.PROTOCOL_VAULT.balance, prevBalance + amount);
    }

    function test_balanceOf_succeeds() external {
        vm.deal(address(vault), vaultBalance);
        vm.prank(AddressAliasHelper.applyL1ToL2Alias(address(pool)));
        vault.reward(recipient, l2BlockNumber);

        uint256 expectedBalance = vaultBalance / vault.REWARD_DIVIDER();
        assertEq(vault.balanceOf(recipient), expectedBalance);
    }
}
