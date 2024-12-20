// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { Predeploys } from "../libraries/Predeploys.sol";
import { ProtocolVault } from "../L2/ProtocolVault.sol";
import { L1FeeVault } from "../L2/L1FeeVault.sol";
import { StandardBridge } from "../universal/StandardBridge.sol";
import { Bridge_Initializer } from "./CommonTest.t.sol";

// Test the implementations of the FeeVault
contract FeeVault_Test is Bridge_Initializer {
    ProtocolVault protocolVault = ProtocolVault(payable(Predeploys.PROTOCOL_VAULT));
    L1FeeVault l1FeeVault = L1FeeVault(payable(Predeploys.KROMA_L1_FEE_VAULT));

    address constant recipient = address(0x10000);
    address constant caller = address(0x10001);

    event Withdrawal(uint256 value, address to, address from);

    function setUp() public override {
        super.setUp();
        vm.etch(Predeploys.PROTOCOL_VAULT, address(new ProtocolVault(recipient)).code);
        vm.etch(Predeploys.KROMA_L1_FEE_VAULT, address(new L1FeeVault(recipient)).code);

        vm.label(Predeploys.PROTOCOL_VAULT, "ProtocolVault");
        vm.label(Predeploys.KROMA_L1_FEE_VAULT, "L1FeeVault");
    }

    function test_constructor_succeeds() external {
        assertEq(protocolVault.RECIPIENT(), recipient);
        assertEq(l1FeeVault.RECIPIENT(), recipient);
    }

    function test_minWithdrawalAmount_succeeds() external {
        assertEq(protocolVault.MIN_WITHDRAWAL_AMOUNT(), 0);
        assertEq(l1FeeVault.MIN_WITHDRAWAL_AMOUNT(), 0);
    }

    function test_withdrawToL2_succeeds() external {
        uint256 reward = 1 ether;
        vm.deal(address(l1FeeVault), reward);
        assertEq(payable(Predeploys.KROMA_L1_FEE_VAULT).balance, reward);

        uint256 prevBalance = payable(recipient).balance;

        // No ether has been withdrawn yet
        assertEq(l1FeeVault.totalProcessed(), 0);

        vm.expectEmit(true, true, true, true, address(Predeploys.KROMA_L1_FEE_VAULT));
        emit Withdrawal(reward, recipient, recipient);

        // Withdraw to L2
        vm.prank(recipient);
        l1FeeVault.withdrawToL2();

        assertEq(l1FeeVault.totalProcessed(), reward);
        assertEq(payable(recipient).balance, prevBalance + reward);
    }

    function test_withdraw_fromOtherEOA_reverts() external {
        vm.expectRevert("FeeVault: the only recipient can call");
        vm.prank(caller);
        l1FeeVault.withdraw();
    }

    function test_withdrawToL2_fromOtherEOA_reverts() external {
        vm.expectRevert("FeeVault: the only recipient can call");
        vm.prank(caller);
        l1FeeVault.withdrawToL2();
    }
}
