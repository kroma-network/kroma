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
    L1FeeVault l1FeeVault = L1FeeVault(payable(Predeploys.L1_FEE_VAULT));

    address constant recipient = address(0x10000);

    function setUp() public override {
        super.setUp();
        vm.etch(Predeploys.PROTOCOL_VAULT, address(new ProtocolVault(recipient)).code);
        vm.etch(Predeploys.L1_FEE_VAULT, address(new L1FeeVault(recipient)).code);

        vm.label(Predeploys.PROTOCOL_VAULT, "ProtocolVault");
        vm.label(Predeploys.L1_FEE_VAULT, "L1FeeVault");
    }

    function test_constructor_succeeds() external {
        assertEq(protocolVault.RECIPIENT(), recipient);
        assertEq(l1FeeVault.RECIPIENT(), recipient);
    }

    function test_minWithdrawalAmount_succeeds() external {
        assertEq(protocolVault.MIN_WITHDRAWAL_AMOUNT(), 10 ether);
        assertEq(l1FeeVault.MIN_WITHDRAWAL_AMOUNT(), 10 ether);
    }
}
