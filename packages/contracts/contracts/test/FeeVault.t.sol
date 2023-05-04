// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { Predeploys } from "../libraries/Predeploys.sol";
import { ProtocolVault } from "../L2/ProtocolVault.sol";
import { ProposerRewardVault } from "../L2/ProposerRewardVault.sol";
import { StandardBridge } from "../universal/StandardBridge.sol";
import { Bridge_Initializer } from "./CommonTest.t.sol";

// Test the implementations of the FeeVault
contract FeeVault_Test is Bridge_Initializer {
    ProtocolVault protocolVault = ProtocolVault(payable(Predeploys.PROTOCOL_VAULT));
    ProposerRewardVault proposerRewardVault = ProposerRewardVault(payable(Predeploys.PROPOSER_REWARD_VAULT));

    address constant recipient = address(0x10000);

    function setUp() public override {
        super.setUp();
        vm.etch(Predeploys.PROTOCOL_VAULT, address(new ProtocolVault(recipient)).code);
        vm.etch(Predeploys.PROPOSER_REWARD_VAULT, address(new ProposerRewardVault(recipient)).code);

        vm.label(Predeploys.PROTOCOL_VAULT, "ProtocolVault");
        vm.label(Predeploys.PROPOSER_REWARD_VAULT, "ProposerRewardVault");
    }

    function test_constructor_succeeds() external {
        assertEq(protocolVault.RECIPIENT(), recipient);
        assertEq(proposerRewardVault.RECIPIENT(), recipient);
    }

    function test_minWithdrawalAmount_succeeds() external {
        assertEq(protocolVault.MIN_WITHDRAWAL_AMOUNT(), 10 ether);
        assertEq(proposerRewardVault.MIN_WITHDRAWAL_AMOUNT(), 10 ether);
    }
}
