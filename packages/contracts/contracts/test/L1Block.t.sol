// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

// Testing utilities
import { CommonTest } from "./CommonTest.t.sol";

// Libraries
import { Encoding } from "../libraries/Encoding.sol";
import { Predeploys } from "../libraries/Predeploys.sol";

// Target contract
import { GovernanceToken } from "../governance/GovernanceToken.sol";
import { Predeploys } from "../libraries/Predeploys.sol";
import { KromaL1Block } from "../L2/KromaL1Block.sol";
import { L1Block } from "../L2/L1Block.sol";

contract L1BlockTest is CommonTest {
    KromaL1Block kromaL1Block;
    address depositor;
    bytes32 immutable NON_ZERO_HASH = keccak256(abi.encode(1));

    /// @dev Sets up the test suite.
    function setUp() public virtual override {
        super.setUp();
        kromaL1Block = new KromaL1Block();
        depositor = kromaL1Block.DEPOSITOR_ACCOUNT();
        vm.prank(depositor);
        kromaL1Block.setL1BlockValues({
            _number: uint64(1),
            _timestamp: uint64(2),
            _basefee: 3,
            _hash: NON_ZERO_HASH,
            _sequenceNumber: uint64(4),
            _batcherHash: bytes32(0),
            _l1FeeOverhead: 2,
            _l1FeeScalar: 3,
            _validatorRewardScalar: 1
        });
    }
}

contract L1BlockBedrock_Test is L1BlockTest {
    // @dev Tests that `setL1BlockValues` updates the values correctly.
    function testFuzz_updatesValues_succeeds(
        uint64 n,
        uint64 t,
        uint256 b,
        bytes32 h,
        uint64 s,
        bytes32 bt,
        uint256 fo,
        uint256 fs,
        uint256 vrr
    ) external {
        vrr = bound(vrr, 0, 10000);
        vm.prank(depositor);
        kromaL1Block.setL1BlockValues(n, t, b, h, s, bt, fo, fs, vrr);
        assertEq(kromaL1Block.number(), n);
        assertEq(kromaL1Block.timestamp(), t);
        assertEq(kromaL1Block.basefee(), b);
        assertEq(kromaL1Block.hash(), h);
        assertEq(kromaL1Block.sequenceNumber(), s);
        assertEq(kromaL1Block.batcherHash(), bt);
        assertEq(kromaL1Block.l1FeeOverhead(), fo);
        assertEq(kromaL1Block.l1FeeScalar(), fs);
        assertEq(kromaL1Block.validatorRewardScalar(), vrr);
    }

    /// @dev Tests that `setL1BlockValues` can set max values.
    function test_updateValues_succeeds() external {
        vm.prank(depositor);
        kromaL1Block.setL1BlockValues({
            _number: type(uint64).max,
            _timestamp: type(uint64).max,
            _basefee: type(uint256).max,
            _hash: keccak256(abi.encode(1)),
            _sequenceNumber: type(uint64).max,
            _batcherHash: bytes32(type(uint256).max),
            _l1FeeOverhead: type(uint256).max,
            _l1FeeScalar: type(uint256).max,
            _validatorRewardScalar: 10000
        });
    }
}

contract L1BlockEcotone_Test is L1BlockTest {
    /// @dev Tests that setL1BlockValuesEcotone updates the values appropriately.
    function testFuzz_setL1BlockValuesEcotone_succeeds(
        uint32 baseFeeScalar,
        uint32 blobBaseFeeScalar,
        uint64 sequenceNumber,
        uint64 timestamp,
        uint64 number,
        uint256 baseFee,
        uint256 blobBaseFee,
        bytes32 hash,
        bytes32 batcherHash,
        uint256 validatorRewardScalar
    ) external {
        bytes memory functionCallDataPacked = Encoding.encodeSetL1BlockValuesEcotone(
            baseFeeScalar,
            blobBaseFeeScalar,
            sequenceNumber,
            timestamp,
            number,
            baseFee,
            blobBaseFee,
            hash,
            batcherHash,
            validatorRewardScalar
        );

        vm.prank(depositor);
        (bool success, ) = address(kromaL1Block).call(functionCallDataPacked);
        assertTrue(success, "Function call failed");

        assertEq(kromaL1Block.baseFeeScalar(), baseFeeScalar);
        assertEq(kromaL1Block.blobBaseFeeScalar(), blobBaseFeeScalar);
        assertEq(kromaL1Block.sequenceNumber(), sequenceNumber);
        assertEq(kromaL1Block.timestamp(), timestamp);
        assertEq(kromaL1Block.number(), number);
        assertEq(kromaL1Block.basefee(), baseFee);
        assertEq(kromaL1Block.blobBaseFee(), blobBaseFee);
        assertEq(kromaL1Block.hash(), hash);
        assertEq(kromaL1Block.batcherHash(), batcherHash);
        assertEq(kromaL1Block.validatorRewardScalar(), validatorRewardScalar);

        // ensure we didn't accidentally pollute the 128 bits of the sequencenum+scalars slot that
        // should be empty
        bytes32 scalarsSlot = vm.load(address(kromaL1Block), bytes32(uint256(3)));
        bytes32 mask128 = hex"FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF00000000000000000000000000000000";

        assertEq(0, scalarsSlot & mask128);

        // ensure we didn't accidentally pollute the 128 bits of the number & timestamp slot that
        // should be empty
        bytes32 numberTimestampSlot = vm.load(address(kromaL1Block), bytes32(uint256(0)));
        assertEq(0, numberTimestampSlot & mask128);
    }

    /// @dev Tests that `setL1BlockValuesEcotone` succeeds if sender address is the depositor
    function test_setL1BlockValuesEcotone_isDepositor_succeeds() external {
        bytes memory functionCallDataPacked = Encoding.encodeSetL1BlockValuesEcotone(
            type(uint32).max,
            type(uint32).max,
            type(uint64).max,
            type(uint64).max,
            type(uint64).max,
            type(uint256).max,
            type(uint256).max,
            bytes32(type(uint256).max),
            bytes32(type(uint256).max),
            type(uint256).max
        );

        vm.prank(depositor);
        (bool success, ) = address(kromaL1Block).call(functionCallDataPacked);
        assertTrue(success, "function call failed");
    }

    /// @dev Tests that `setL1BlockValuesEcotone` fails if sender address is not the depositor
    function test_setL1BlockValuesEcotone_notDepositor_fails() external {
        bytes memory functionCallDataPacked = Encoding.encodeSetL1BlockValuesEcotone(
            type(uint32).max,
            type(uint32).max,
            type(uint64).max,
            type(uint64).max,
            type(uint64).max,
            type(uint256).max,
            type(uint256).max,
            bytes32(type(uint256).max),
            bytes32(type(uint256).max),
            type(uint256).max
        );

        (bool success, bytes memory data) = address(kromaL1Block).call(functionCallDataPacked);
        assertTrue(!success, "function call should have failed");
        // make sure return value is the expected function selector for "NotDepositor()"
        bytes memory expReturn = hex"3cc50b45";
        assertEq(data, expReturn);
    }
}

contract L1BlockKromaMPT_Test is CommonTest {
    L1Block l1Block;
    address depositor;
    bytes32 immutable NON_ZERO_HASH = keccak256(abi.encode(1));

    /// @dev Sets up the test suite.
    function setUp() public virtual override {
        super.setUp();
        l1Block = new L1Block();
        depositor = l1Block.DEPOSITOR_ACCOUNT();
    }

    /// @dev Tests that setL1BlockValuesEcotone updates the values appropriately without validatorRewardScalar.
    function testFuzz_setL1BlockValuesEcotone_withoutValidatorRewardScalar_succeeds(
        uint32 baseFeeScalar,
        uint32 blobBaseFeeScalar,
        uint64 sequenceNumber,
        uint64 timestamp,
        uint64 number,
        uint256 baseFee,
        uint256 blobBaseFee,
        bytes32 hash,
        bytes32 batcherHash
    ) external {
        bytes memory functionCallDataPacked = Encoding.encodeSetL1BlockValuesKromaMPT(
            baseFeeScalar,
            blobBaseFeeScalar,
            sequenceNumber,
            timestamp,
            number,
            baseFee,
            blobBaseFee,
            hash,
            batcherHash
        );

        vm.prank(depositor);
        (bool success, ) = address(l1Block).call(functionCallDataPacked);
        assertTrue(success, "Function call failed");

        assertEq(l1Block.baseFeeScalar(), baseFeeScalar);
        assertEq(l1Block.blobBaseFeeScalar(), blobBaseFeeScalar);
        assertEq(l1Block.sequenceNumber(), sequenceNumber);
        assertEq(l1Block.timestamp(), timestamp);
        assertEq(l1Block.number(), number);
        assertEq(l1Block.basefee(), baseFee);
        assertEq(l1Block.blobBaseFee(), blobBaseFee);
        assertEq(l1Block.hash(), hash);
        assertEq(l1Block.batcherHash(), batcherHash);

        // ensure we didn't accidentally pollute the 128 bits of the sequencenum+scalars slot that
        // should be empty
        bytes32 scalarsSlot = vm.load(address(l1Block), bytes32(uint256(3)));
        bytes32 mask128 = hex"FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF00000000000000000000000000000000";

        assertEq(0, scalarsSlot & mask128);

        // ensure we didn't accidentally pollute the 128 bits of the number & timestamp slot that
        // should be empty
        bytes32 numberTimestampSlot = vm.load(address(l1Block), bytes32(uint256(0)));
        assertEq(0, numberTimestampSlot & mask128);
    }

    /// @dev Tests that `setL1BlockValuesEcotone` succeeds if sender address is the depositor
    function test_setL1BlockValuesEcotone_isDepositor_succeeds() external {
        bytes memory functionCallDataPacked = Encoding.encodeSetL1BlockValuesKromaMPT(
            type(uint32).max,
            type(uint32).max,
            type(uint64).max,
            type(uint64).max,
            type(uint64).max,
            type(uint256).max,
            type(uint256).max,
            bytes32(type(uint256).max),
            bytes32(type(uint256).max)
        );

        vm.prank(depositor);
        (bool success, ) = address(l1Block).call(functionCallDataPacked);
        assertTrue(success, "function call failed");
    }

    /// @dev Tests that `setL1BlockValuesEcotone` fails if sender address is not the depositor
    function test_setL1BlockValuesEcotone_notDepositor_fails() external {
        bytes memory functionCallDataPacked = Encoding.encodeSetL1BlockValuesKromaMPT(
            type(uint32).max,
            type(uint32).max,
            type(uint64).max,
            type(uint64).max,
            type(uint64).max,
            type(uint256).max,
            type(uint256).max,
            bytes32(type(uint256).max),
            bytes32(type(uint256).max)
        );

        (bool success, bytes memory data) = address(l1Block).call(functionCallDataPacked);
        assertTrue(!success, "function call should have failed");
        // make sure return value is the expected function selector for "NotDepositor()"
        bytes memory expReturn = hex"3cc50b45";
        assertEq(data, expReturn);
    }
}
