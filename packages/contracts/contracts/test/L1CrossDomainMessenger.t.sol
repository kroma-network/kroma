// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { Encoding } from "../libraries/Encoding.sol";
import { Hashing } from "../libraries/Hashing.sol";
import { Predeploys } from "../libraries/Predeploys.sol";
import { KromaPortal } from "../L1/KromaPortal.sol";
import { L1CrossDomainMessenger } from "../L1/L1CrossDomainMessenger.sol";
import { L2OutputOracle } from "../L1/L2OutputOracle.sol";
import { AddressAliasHelper } from "../vendor/AddressAliasHelper.sol";
import { Messenger_Initializer, Reverter, ConfigurableCaller } from "./CommonTest.t.sol";
import { L2OutputOracle_Initializer } from "./L2OutputOracle.t.sol";

contract L1CrossDomainMessenger_Test is Messenger_Initializer {
    // Receiver address for testing
    address recipient = address(0xabbaacdc);

    // Storage slot of the l2Sender
    uint256 constant senderSlotIndex = 50;

    // the version is encoded in the nonce
    function test_messageVersion_succeeds() external {
        (, uint16 version) = Encoding.decodeVersionedNonce(L1Messenger.messageNonce());
        assertEq(version, L1Messenger.MESSAGE_VERSION());
    }

    // sendMessage: should be able to send a single message
    function test_sendMessage_succeeds() external {
        // deposit transaction on the kroma portal should be called
        vm.expectCall(
            address(portal),
            abi.encodeWithSelector(
                KromaPortal.depositTransaction.selector,
                Predeploys.L2_CROSS_DOMAIN_MESSENGER,
                0,
                L1Messenger.baseGas(hex"ff", 100),
                false,
                Encoding.encodeCrossDomainMessage(
                    L1Messenger.messageNonce(),
                    alice,
                    recipient,
                    0,
                    100,
                    hex"ff"
                )
            )
        );

        // TransactionDeposited event
        vm.expectEmit(true, true, true, true);
        emitTransactionDeposited(
            AddressAliasHelper.applyL1ToL2Alias(address(L1Messenger)),
            Predeploys.L2_CROSS_DOMAIN_MESSENGER,
            0,
            0,
            L1Messenger.baseGas(hex"ff", 100),
            false,
            Encoding.encodeCrossDomainMessage(
                L1Messenger.messageNonce(),
                alice,
                recipient,
                0,
                100,
                hex"ff"
            )
        );

        // SentMessage event
        vm.expectEmit(true, true, true, true);
        emit SentMessage(recipient, alice, 0, hex"ff", L1Messenger.messageNonce(), 100);

        vm.prank(alice);
        L1Messenger.sendMessage(recipient, hex"ff", uint32(100));
    }

    // sendMessage: should be able to send the same message twice
    function test_sendMessage_twice_succeeds() external {
        uint256 nonce = L1Messenger.messageNonce();
        L1Messenger.sendMessage(recipient, hex"aa", uint32(500_000));
        L1Messenger.sendMessage(recipient, hex"aa", uint32(500_000));
        // the nonce increments for each message sent
        assertEq(nonce + 2, L1Messenger.messageNonce());
    }

    function test_xDomainSender_notSet_reverts() external {
        vm.expectRevert("CrossDomainMessenger: xDomainMessageSender is not set");
        L1Messenger.xDomainMessageSender();
    }

    function test_relayMessage_v2_reverts() external {
        address target = address(0xabcd);
        address sender = Predeploys.L2_CROSS_DOMAIN_MESSENGER;

        // Set the value of portal.l2Sender() to be the L2 Cross Domain Messenger.
        vm.store(address(portal), bytes32(senderSlotIndex), bytes32(abi.encode(sender)));

        // Expect a revert.
        vm.expectRevert("CrossDomainMessenger: only version 0 messages is supported at this time");

        // Try to relay a v2 message.
        vm.prank(address(portal));
        L2Messenger.relayMessage(
            Encoding.encodeVersionedNonce({ _nonce: 0, _version: 2 }), // nonce
            sender,
            target,
            0, // value
            0,
            hex"1111"
        );
    }

    // relayMessage: should send a successful call to the target contract
    function test_relayMessage_succeeds() external {
        address target = address(0xabcd);
        address sender = Predeploys.L2_CROSS_DOMAIN_MESSENGER;

        vm.expectCall(target, hex"1111");

        // set the value of portal.l2Sender() to be the L2 Cross Domain Messenger.
        vm.store(address(portal), bytes32(senderSlotIndex), bytes32(abi.encode(sender)));
        vm.prank(address(portal));

        vm.expectEmit(true, true, true, true);

        bytes32 hash = Hashing.hashCrossDomainMessage(
            Encoding.encodeVersionedNonce({ _nonce: 0, _version: 0 }),
            sender,
            target,
            0,
            0,
            hex"1111"
        );

        emit RelayedMessage(hash);

        L1Messenger.relayMessage(
            Encoding.encodeVersionedNonce({ _nonce: 0, _version: 0 }), // nonce
            sender,
            target,
            0, // value
            0,
            hex"1111"
        );

        // the message hash is in the successfulMessages mapping
        assert(L1Messenger.successfulMessages(hash));
        // it is not in the received messages mapping
        assertEq(L1Messenger.failedMessages(hash), false);
    }

    // relayMessage: should revert if attempting to relay a message sent to an L1 system contract
    function test_relayMessage_toSystemContract_reverts() external {
        // set the target to be the KromaPortal
        address target = address(portal);
        address sender = Predeploys.L2_CROSS_DOMAIN_MESSENGER;
        bytes memory message = hex"1111";

        vm.prank(address(portal));
        vm.expectRevert("CrossDomainMessenger: message cannot be replayed");
        L1Messenger.relayMessage(
            Encoding.encodeVersionedNonce({ _nonce: 0, _version: 0 }),
            sender,
            target,
            0,
            0,
            message
        );

        vm.store(address(portal), 0, bytes32(abi.encode(sender)));
        vm.expectRevert("CrossDomainMessenger: message cannot be replayed");
        L1Messenger.relayMessage(
            Encoding.encodeVersionedNonce({ _nonce: 0, _version: 0 }),
            sender,
            target,
            0,
            0,
            message
        );
    }

    // relayMessage: should revert if eth is sent from a contract other than the standard bridge
    function test_replayMessage_withValue_reverts() external {
        address target = address(0xabcd);
        address sender = Predeploys.L2_CROSS_DOMAIN_MESSENGER;
        bytes memory message = hex"1111";

        vm.expectRevert(
            "CrossDomainMessenger: value must be zero unless message is from a system address"
        );
        L1Messenger.relayMessage{ value: 100 }(
            Encoding.encodeVersionedNonce({ _nonce: 0, _version: 0 }),
            sender,
            target,
            0,
            0,
            message
        );
    }

    // relayMessage: the xDomainMessageSender is reset to the original value
    function test_xDomainMessageSender_reset_succeeds() external {
        vm.expectRevert("CrossDomainMessenger: xDomainMessageSender is not set");
        L1Messenger.xDomainMessageSender();

        address sender = Predeploys.L2_CROSS_DOMAIN_MESSENGER;

        vm.store(address(portal), bytes32(senderSlotIndex), bytes32(abi.encode(sender)));
        vm.prank(address(portal));
        L1Messenger.relayMessage(
            Encoding.encodeVersionedNonce({ _nonce: 0, _version: 0 }),
            address(0),
            address(0),
            0,
            0,
            hex""
        );

        vm.expectRevert("CrossDomainMessenger: xDomainMessageSender is not set");
        L1Messenger.xDomainMessageSender();
    }

    // relayMessage: should send a successful call to the target contract after the first message
    // fails and ETH gets stuck, but the second message succeeds
    function test_relayMessage_retryAfterFailure_succeeds() external {
        address target = address(0xabcd);
        address sender = Predeploys.L2_CROSS_DOMAIN_MESSENGER;
        uint256 value = 100;

        vm.expectCall(target, hex"1111");

        bytes32 hash = Hashing.hashCrossDomainMessage(
            Encoding.encodeVersionedNonce({ _nonce: 0, _version: 0 }),
            sender,
            target,
            value,
            0,
            hex"1111"
        );

        vm.store(address(portal), bytes32(senderSlotIndex), bytes32(abi.encode(sender)));
        vm.etch(target, address(new Reverter()).code);
        vm.deal(address(portal), value);
        vm.prank(address(portal));
        L1Messenger.relayMessage{ value: value }(
            Encoding.encodeVersionedNonce({ _nonce: 0, _version: 0 }), // nonce
            sender,
            target,
            value,
            0,
            hex"1111"
        );

        assertEq(address(L1Messenger).balance, value);
        assertEq(address(target).balance, 0);
        assertEq(L1Messenger.successfulMessages(hash), false);
        assertEq(L1Messenger.failedMessages(hash), true);

        vm.expectEmit(true, true, true, true);

        emit RelayedMessage(hash);

        vm.etch(target, address(0).code);
        vm.prank(address(sender));
        L1Messenger.relayMessage(
            Encoding.encodeVersionedNonce({ _nonce: 0, _version: 0 }), // nonce
            sender,
            target,
            value,
            0,
            hex"1111"
        );

        assertEq(address(L1Messenger).balance, 0);
        assertEq(address(target).balance, value);
        assertEq(L1Messenger.successfulMessages(hash), true);
        assertEq(L1Messenger.failedMessages(hash), true);
    }
}
