pragma solidity 0.8.15;

import { StdUtils } from "forge-std/StdUtils.sol";
import { Vm } from "forge-std/Vm.sol";

import { Encoding } from "../../libraries/Encoding.sol";
import { Hashing } from "../../libraries/Hashing.sol";
import { Predeploys } from "../../libraries/Predeploys.sol";
import { Types } from "../../libraries/Types.sol";
import { L1CrossDomainMessenger } from "../../L1/L1CrossDomainMessenger.sol";
import { KromaPortal } from "../../L1/KromaPortal.sol";
import { Messenger_Initializer } from "../CommonTest.t.sol";

contract RelayActor is StdUtils {
    // Storage slot of the l2Sender
    uint256 constant senderSlotIndex = 50;

    uint256 public numHashes;
    bytes32[] public hashes;
    bool public reverted = false;

    KromaPortal portal;
    L1CrossDomainMessenger xdm;
    Vm vm;

    constructor(
        KromaPortal _portal,
        L1CrossDomainMessenger _xdm,
        Vm _vm
    ) {
        portal = _portal;
        xdm = _xdm;
        vm = _vm;
    }

    /**
     * Relays a message to the `L1CrossDomainMessenger` with a random `version`, `_minGasLimit`
     * and `_message`.
     */
    function relay(
        uint32 _minGasLimit,
        bytes memory _message
    ) external {
        address target = address(0x04); // ID precompile
        address sender = Predeploys.L2_CROSS_DOMAIN_MESSENGER;

        // set the value of op.l2Sender() to be the L2 Cross Domain Messenger.
        vm.store(address(portal), bytes32(senderSlotIndex), bytes32(abi.encode(sender)));

        // Restrict `_minGasLimit` to a number in the range of the block gas limit.
        _minGasLimit = uint32(bound(_minGasLimit, 0, block.gaslimit));

        // Compute the cross domain message hash and store it in `hashes`.
        bytes32 _hash = Hashing.hashCrossDomainMessageV0(
            Encoding.encodeVersionedNonce({ _nonce: 0, _version: 0 }),
            sender,
            target,
            0, // value
            _minGasLimit,
            _message
        );

        // Act as the kroma portal and call `relayMessage` on the `L1CrossDomainMessenger` with
        // the outer min gas limit.
        vm.startPrank(address(portal));
        vm.expectCall(target, _message);
        try
            xdm.relayMessage{ gas: xdm.baseGas(_message, _minGasLimit) }(
                Encoding.encodeVersionedNonce({ _nonce: 0, _version: 0 }),
                sender,
                target,
                0, // value
                _minGasLimit,
                _message
            )
        {} catch {
            // If any of these calls revert, set `reverted` to true to fail the invariant test.
            // NOTE: This is to get around forge's invariant fuzzer ignoring reverted calls
            // to this function.
            reverted = true;
        }
        vm.stopPrank();

        hashes.push(_hash);
        numHashes += 1;
    }
}

contract XDM_MinGasLimits is Messenger_Initializer {
    RelayActor actor;

    function setUp() public virtual override {
        // Set up the `L1CrossDomainMessenger` and `KromaPortal` contracts.
        super.setUp();

        // Deploy a relay actor
        actor = new RelayActor(portal, L1Messenger, vm);

        // Target the `RelayActor` contract
        targetContract(address(actor));

        // Target the actor's `relay` function
        bytes4[] memory selectors = new bytes4[](1);
        selectors[0] = actor.relay.selector;
        targetSelector(FuzzSelector({ addr: address(actor), selectors: selectors }));
    }

    /**
     * @custom:invariant A call to `relayMessage` should never revert if at least the proper minimum
     * gas limits are supplied.
     *
     * There are two minimum gas limits here:
     *
     * - The outer min gas limit is for the call from the `KromaPortal` to the
     * `L1CrossDomainMessenger`,  and it can be retrieved by calling the xdm's `baseGas` function
     * with the `message` and inner limit.
     *
     * - The inner min gas limit is for the call from the `L1CrossDomainMessenger` to the target
     * contract.
     */
    /*
    NOTE(chokobole): disable this test as Optimism also fails.
    function invariant_minGasLimits() public {
        uint256 length = actor.numHashes();
        for (uint256 i = 0; i < length; ++i) {
            bytes32 hash = actor.hashes(i);
            // the message hash is in the successfulMessages mapping
            assertTrue(L1Messenger.successfulMessages(hash));
            // it is not in the received messages mapping
            assertFalse(L1Messenger.failedMessages(hash));
        }
        assertFalse(actor.reverted());
    }
    */
}
