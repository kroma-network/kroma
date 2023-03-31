// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { ERC20 } from "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import { stdStorage, StdStorage } from "forge-std/Test.sol";

import { Hashing } from "../libraries/Hashing.sol";
import { Types } from "../libraries/Types.sol";
import { Predeploys } from "../libraries/Predeploys.sol";
import { L2ToL1MessagePasser } from "../L2/L2ToL1MessagePasser.sol";
import { CrossDomainMessenger } from "../universal/CrossDomainMessenger.sol";
import { KanvasMintableERC20 } from "../universal/KanvasMintableERC20.sol";
import { StandardBridge } from "../universal/StandardBridge.sol";
import { Bridge_Initializer } from "./CommonTest.t.sol";

contract L2StandardBridge_Test is Bridge_Initializer {
    using stdStorage for StdStorage;

    function test_initialize_succeeds() external {
        assertEq(address(L2Bridge.MESSENGER()), address(L2Messenger));
        assertEq(address(L1Bridge.OTHER_BRIDGE()), address(L2Bridge));
        assertEq(address(L2Bridge.OTHER_BRIDGE()), address(L1Bridge));
    }

    // receive
    // - can accept ETH
    function test_receive_succeeds() external {
        assertEq(address(messagePasser).balance, 0);
        vm.prank(alice, alice);
        (bool success, ) = address(L2Bridge).call{ value: 100 }(hex"");
        assertEq(success, true);
        assertEq(address(messagePasser).balance, 100);
    }
}

contract PreBridgeERC20 is Bridge_Initializer {
    function _preBridgeERC20() internal {
        // Alice has 100 L2Token
        deal(address(L2Token), alice, 100, true);
        assertEq(L2Token.balanceOf(alice), 100);
        uint256 nonce = L2Messenger.messageNonce();
        bytes memory message = abi.encodeWithSelector(
            StandardBridge.finalizeBridgeERC20.selector,
            address(L1Token),
            address(L2Token),
            alice,
            alice,
            100,
            hex""
        );
        uint64 baseGas = L2Messenger.baseGas(message, 1000);
        bytes memory withdrawalData = abi.encodeWithSelector(
            CrossDomainMessenger.relayMessage.selector,
            nonce,
            address(L2Bridge),
            address(L1Bridge),
            0,
            1000,
            message
        );
        bytes32 withdrawalHash = Hashing.hashWithdrawal(
            Types.WithdrawalTransaction({
                nonce: nonce,
                sender: address(L2Messenger),
                target: address(L1Messenger),
                value: 0,
                gasLimit: baseGas,
                data: withdrawalData
            })
        );

        vm.expectCall(
            address(L2Bridge),
            abi.encodeWithSelector(
                L2Bridge.bridgeERC20.selector,
                address(L2Token),
                address(L1Token),
                100,
                1000,
                hex""
            )
        );

        vm.expectCall(
            address(L2Messenger),
            abi.encodeWithSelector(
                CrossDomainMessenger.sendMessage.selector,
                address(L1Bridge),
                message,
                1000
            )
        );

        vm.expectCall(
            Predeploys.L2_TO_L1_MESSAGE_PASSER,
            abi.encodeWithSelector(
                L2ToL1MessagePasser.initiateWithdrawal.selector,
                address(L1Messenger),
                baseGas,
                withdrawalData
            )
        );

        // The L2Bridge should burn the tokens
        vm.expectCall(
            address(L2Token),
            abi.encodeWithSelector(KanvasMintableERC20.burn.selector, alice, 100)
        );

        vm.expectEmit(true, true, true, true);
        emit ERC20BridgeInitiated(address(L2Token), address(L1Token), alice, alice, 100, hex"");

        vm.expectEmit(true, true, true, true);
        emit MessagePassed(
            nonce,
            address(L2Messenger),
            address(L1Messenger),
            0,
            baseGas,
            withdrawalData,
            withdrawalHash
        );

        // SentMessage event emitted by the CrossDomainMessenger
        vm.expectEmit(true, true, true, true);
        emit SentMessage(address(L1Bridge), address(L2Bridge), 0, message, nonce, 1000);

        vm.prank(alice, alice);
    }
}

contract L2StandardBridge_BridgeERC20_Test is PreBridgeERC20 {
    // BridgeERC20
    // - token is burned
    // - emits ERC20BridgeInitiated
    // - calls Withdrawer.initiateWithdrawal
    function test_bridgeERC20_succeeds() external {
        _preBridgeERC20();
        L2Bridge.bridgeERC20(address(L2Token), address(L1Token), 100, 1000, hex"");

        assertEq(L2Token.balanceOf(alice), 0);
    }
}

contract PreBridgeERC20To is Bridge_Initializer {
    function _preBridgeERC20To() internal {
        deal(address(L2Token), alice, 100, true);
        assertEq(L2Token.balanceOf(alice), 100);
        uint256 nonce = L2Messenger.messageNonce();
        bytes memory message = abi.encodeWithSelector(
            StandardBridge.finalizeBridgeERC20.selector,
            address(L1Token),
            address(L2Token),
            alice,
            bob,
            100,
            hex""
        );
        uint64 baseGas = L2Messenger.baseGas(message, 1000);
        bytes memory withdrawalData = abi.encodeWithSelector(
            CrossDomainMessenger.relayMessage.selector,
            nonce,
            address(L2Bridge),
            address(L1Bridge),
            0,
            1000,
            message
        );
        bytes32 withdrawalHash = Hashing.hashWithdrawal(
            Types.WithdrawalTransaction({
                nonce: nonce,
                sender: address(L2Messenger),
                target: address(L1Messenger),
                value: 0,
                gasLimit: baseGas,
                data: withdrawalData
            })
        );

        vm.expectEmit(true, true, true, true, address(L2Bridge));
        emit ERC20BridgeInitiated(address(L2Token), address(L1Token), alice, bob, 100, hex"");

        vm.expectEmit(true, true, true, true, address(messagePasser));
        emit MessagePassed(
            nonce,
            address(L2Messenger),
            address(L1Messenger),
            0,
            baseGas,
            withdrawalData,
            withdrawalHash
        );

        // SentMessage event emitted by the CrossDomainMessenger
        vm.expectEmit(true, true, true, true, address(L2Messenger));
        emit SentMessage(address(L1Bridge), address(L2Bridge), 0, message, nonce, 1000);

        vm.expectCall(
            address(L2Bridge),
            abi.encodeWithSelector(
                L2Bridge.bridgeERC20To.selector,
                address(L2Token),
                address(L1Token),
                bob,
                100,
                1000,
                hex""
            )
        );

        vm.expectCall(
            address(L2Messenger),
            abi.encodeWithSelector(
                CrossDomainMessenger.sendMessage.selector,
                address(L1Bridge),
                message,
                1000
            )
        );

        vm.expectCall(
            Predeploys.L2_TO_L1_MESSAGE_PASSER,
            abi.encodeWithSelector(
                L2ToL1MessagePasser.initiateWithdrawal.selector,
                address(L1Messenger),
                baseGas,
                withdrawalData
            )
        );

        // The L2Bridge should burn the tokens
        vm.expectCall(
            address(L2Token),
            abi.encodeWithSelector(KanvasMintableERC20.burn.selector, alice, 100)
        );

        vm.prank(alice, alice);
    }
}

contract L2StandardBridge_BridgeERC20To_Test is PreBridgeERC20To {
    // bridgeERC20To
    // - token is burned
    // - emits ERC20BridgeInitiated w/ correct recipient
    // - calls Withdrawer.initiateWithdrawal
    function test_bridgeERC20To_succeeds() external {
        _preBridgeERC20To();
        L2Bridge.bridgeERC20To(address(L2Token), address(L1Token), bob, 100, 1000, hex"");
        assertEq(L2Token.balanceOf(alice), 0);
    }
}

contract L2StandardBridge_Bridge_Test is Bridge_Initializer {
    // finalizeDeposit
    // - only callable by l1Bridge
    // - supported token pair emits ERC20BridgeFinalized
    // - invalid deposit calls Withdrawer.initiateWithdrawal
    function test_finalizeBridgeERC20_succeeds() external {
        vm.mockCall(
            address(L2Bridge.MESSENGER()),
            abi.encodeWithSelector(CrossDomainMessenger.xDomainMessageSender.selector),
            abi.encode(address(L2Bridge.OTHER_BRIDGE()))
        );

        vm.expectCall(
            address(L2Token),
            abi.encodeWithSelector(KanvasMintableERC20.mint.selector, alice, 100)
        );

        vm.expectEmit(true, true, true, true, address(L2Bridge));
        emit ERC20BridgeFinalized(address(L2Token), address(L1Token), alice, alice, 100, hex"");

        vm.prank(address(L2Messenger));
        L2Bridge.finalizeBridgeERC20(address(L2Token), address(L1Token), alice, alice, 100, hex"");
    }

    function test_finalizeBridgeETH_incorrectValue_reverts() external {
        vm.mockCall(
            address(L2Bridge.MESSENGER()),
            abi.encodeWithSelector(CrossDomainMessenger.xDomainMessageSender.selector),
            abi.encode(address(L2Bridge.OTHER_BRIDGE()))
        );
        vm.deal(address(L2Messenger), 100);
        vm.prank(address(L2Messenger));
        vm.expectRevert("StandardBridge: amount sent does not match amount required");
        L2Bridge.finalizeBridgeETH{ value: 50 }(alice, alice, 100, hex"");
    }

    function test_finalizeBridgeETH_sendToSelf_reverts() external {
        vm.mockCall(
            address(L2Bridge.MESSENGER()),
            abi.encodeWithSelector(CrossDomainMessenger.xDomainMessageSender.selector),
            abi.encode(address(L2Bridge.OTHER_BRIDGE()))
        );
        vm.deal(address(L2Messenger), 100);
        vm.prank(address(L2Messenger));
        vm.expectRevert("StandardBridge: cannot send to self");
        L2Bridge.finalizeBridgeETH{ value: 100 }(alice, address(L2Bridge), 100, hex"");
    }

    function test_finalizeBridgeETH_sendToMessenger_reverts() external {
        vm.mockCall(
            address(L2Bridge.MESSENGER()),
            abi.encodeWithSelector(CrossDomainMessenger.xDomainMessageSender.selector),
            abi.encode(address(L2Bridge.OTHER_BRIDGE()))
        );
        vm.deal(address(L2Messenger), 100);
        vm.prank(address(L2Messenger));
        vm.expectRevert("StandardBridge: cannot send to messenger");
        L2Bridge.finalizeBridgeETH{ value: 100 }(alice, address(L2Messenger), 100, hex"");
    }
}

contract L2StandardBridge_FinalizeBridgeETH_Test is Bridge_Initializer {
    function test_finalizeBridgeETH_succeeds() external {
        address messenger = address(L2Bridge.MESSENGER());
        vm.mockCall(
            messenger,
            abi.encodeWithSelector(CrossDomainMessenger.xDomainMessageSender.selector),
            abi.encode(address(L2Bridge.OTHER_BRIDGE()))
        );
        vm.deal(messenger, 100);
        vm.prank(messenger);

        vm.expectEmit(true, true, true, true);
        emit ETHBridgeFinalized(alice, alice, 100, hex"");

        L2Bridge.finalizeBridgeETH{ value: 100 }(alice, alice, 100, hex"");
    }
}
