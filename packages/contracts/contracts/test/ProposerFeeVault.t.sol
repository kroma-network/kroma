// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { Predeploys } from "../libraries/Predeploys.sol";
import { ProposerFeeVault } from "../L2/ProposerFeeVault.sol";
import { StandardBridge } from "../universal/StandardBridge.sol";
import { Bridge_Initializer } from "./CommonTest.t.sol";

contract ProposerFeeVault_Test is Bridge_Initializer {
    ProposerFeeVault vault = ProposerFeeVault(payable(Predeploys.PROPOSER_FEE_WALLET));
    address constant recipient = address(256);

    event Withdrawal(uint256 value, address to, address from);

    function setUp() public override {
        super.setUp();
        vm.etch(Predeploys.PROPOSER_FEE_WALLET, address(new ProposerFeeVault(recipient)).code);
        vm.label(Predeploys.PROPOSER_FEE_WALLET, "ProposerFeeVault");
    }

    function test_minWithdrawalAmount_succeeds() external {
        assertEq(vault.MIN_WITHDRAWAL_AMOUNT(), 10 ether);
    }

    function test_constructor_succeeds() external {
        assertEq(vault.RECIPIENT(), recipient);
    }

    function test_receive_succeeds() external {
        uint256 balance = address(vault).balance;

        vm.prank(alice);
        (bool success, ) = address(vault).call{ value: 100 }(hex"");

        assertEq(success, true);
        assertEq(address(vault).balance, balance + 100);
    }

    function test_withdraw_notEnough_reverts() external {
        assert(address(vault).balance < vault.MIN_WITHDRAWAL_AMOUNT());

        vm.expectRevert(
            "FeeVault: withdrawal amount must be greater than minimum withdrawal amount"
        );
        vault.withdraw();
    }

    function test_withdraw_succeeds() external {
        uint256 amount = vault.MIN_WITHDRAWAL_AMOUNT() + 1;
        vm.deal(address(vault), amount);

        // No ether has been withdrawn yet
        assertEq(vault.totalProcessed(), 0);

        vm.expectEmit(true, true, true, true, address(Predeploys.PROPOSER_FEE_WALLET));
        emit Withdrawal(address(vault).balance, vault.RECIPIENT(), address(this));

        // The entire vault's balance is withdrawn
        vm.expectCall(
            Predeploys.L2_STANDARD_BRIDGE,
            address(vault).balance,
            abi.encodeWithSelector(
                StandardBridge.bridgeETHTo.selector,
                vault.RECIPIENT(),
                35_000,
                bytes("")
            )
        );

        vault.withdraw();

        // The withdrawal was successful
        assertEq(vault.totalProcessed(), amount);
    }
}
