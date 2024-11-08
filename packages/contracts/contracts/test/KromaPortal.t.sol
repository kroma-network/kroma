// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { stdError } from "forge-std/Test.sol";

import { KromaPortal } from "../L1/KromaPortal.sol";
import { L2OutputOracle } from "../L1/L2OutputOracle.sol";
import { ResourceMetering } from "../L1/ResourceMetering.sol";
import { Hashing } from "../libraries/Hashing.sol";
import { Types } from "../libraries/Types.sol";
import { Proxy } from "../universal/Proxy.sol";
import { AddressAliasHelper } from "../vendor/AddressAliasHelper.sol";
import { Portal_Initializer, CommonTest, NextImpl } from "./CommonTest.t.sol";

contract KromaPortal_Test is Portal_Initializer {
    event Paused(address);
    event Unpaused(address);

    function test_constructor_succeeds() external {
        assertEq(address(portal.L2_ORACLE()), address(oracle));
        assertEq(portal.l2Sender(), 0x000000000000000000000000000000000000dEaD);
        assertEq(portal.paused(), false);
    }

    /**
     * @notice The KromaPortal can be paused by the GUARDIAN
     */
    function test_pause_succeeds() external {
        address guardian = portal.GUARDIAN();

        assertEq(portal.paused(), false);

        vm.expectEmit(true, true, true, true, address(portal));
        emit Paused(guardian);

        vm.prank(guardian);
        portal.pause();

        assertEq(portal.paused(), true);
    }

    /**
     * @notice The KromaPortal reverts when an account that is not the
     *         GUARDIAN calls `pause()`
     */
    function test_pause_onlyGuardian_reverts() external {
        assertEq(portal.paused(), false);

        assertTrue(portal.GUARDIAN() != alice);
        vm.expectRevert("KromaPortal: only guardian can pause");
        vm.prank(alice);
        portal.pause();

        assertEq(portal.paused(), false);
    }

    /**
     * @notice The KromaPortal can be unpaused by the GUARDIAN
     */
    function test_unpause_succeeds() external {
        address guardian = portal.GUARDIAN();

        vm.prank(guardian);
        portal.pause();
        assertEq(portal.paused(), true);

        vm.expectEmit(true, true, true, true, address(portal));
        emit Unpaused(guardian);
        vm.prank(guardian);
        portal.unpause();

        assertEq(portal.paused(), false);
    }

    /**
     * @notice The KromaPortal reverts when an account that is not
     *         the GUARDIAN calls `unpause()`
     */
    function test_unpause_onlyGuardian_reverts() external {
        address guardian = portal.GUARDIAN();

        vm.prank(guardian);
        portal.pause();
        assertEq(portal.paused(), true);

        assertTrue(portal.GUARDIAN() != alice);
        vm.expectRevert("KromaPortal: only guardian can unpause");
        vm.prank(alice);
        portal.unpause();

        assertEq(portal.paused(), true);
    }

    function test_receive_succeeds() external {
        vm.expectEmit(true, true, false, true);
        emitTransactionDeposited(alice, alice, 100, 100, 100_000, false, hex"");

        // give alice money and send as an eoa
        vm.deal(alice, 2 ** 64);
        vm.prank(alice, alice);
        (bool s, ) = address(portal).call{ value: 100 }(hex"");

        assert(s);
        assertEq(address(portal).balance, 100);
    }

    // Test: depositTransaction fails when contract creation has a non-zero destination address
    function test_depositTransaction_contractCreation_reverts() external {
        // contract creation must have a target of address(0)
        vm.expectRevert("KromaPortal: must send to address(0) when creating a contract");
        portal.depositTransaction(address(1), 1, 0, true, hex"");
    }

    /**
     * @notice Prevent gasless deposits from being force processed in L2 by
     *         ensuring that they have a large enough gas limit set.
     */
    function test_depositTransaction_smallGasLimit_reverts() external {
        vm.expectRevert("KromaPortal: gas limit must cover instrinsic gas cost");
        portal.depositTransaction({
            _to: address(1),
            _value: 0,
            _gasLimit: 0,
            _isCreation: false,
            _data: hex""
        });
    }

    // Test: depositTransaction should emit the correct log when an EOA deposits a tx with 0 value
    function test_depositTransaction_noValueEOA_succeeds() external {
        // EOA emulation
        vm.prank(address(this), address(this));
        vm.expectEmit(true, true, false, true);
        emitTransactionDeposited(
            address(this),
            NON_ZERO_ADDRESS,
            ZERO_VALUE,
            ZERO_VALUE,
            NON_ZERO_GASLIMIT,
            false,
            NON_ZERO_DATA
        );

        portal.depositTransaction(
            NON_ZERO_ADDRESS,
            ZERO_VALUE,
            NON_ZERO_GASLIMIT,
            false,
            NON_ZERO_DATA
        );
    }

    // Test: depositTransaction should emit the correct log when a contract deposits a tx with 0 value
    function test_depositTransaction_noValueContract_succeeds() external {
        vm.expectEmit(true, true, false, true);
        emitTransactionDeposited(
            AddressAliasHelper.applyL1ToL2Alias(address(this)),
            NON_ZERO_ADDRESS,
            ZERO_VALUE,
            ZERO_VALUE,
            NON_ZERO_GASLIMIT,
            false,
            NON_ZERO_DATA
        );

        portal.depositTransaction(
            NON_ZERO_ADDRESS,
            ZERO_VALUE,
            NON_ZERO_GASLIMIT,
            false,
            NON_ZERO_DATA
        );
    }

    // Test: depositTransaction should emit the correct log when an EOA deposits a contract creation with 0 value
    function test_depositTransaction_createWithZeroValueForEOA_succeeds() external {
        // EOA emulation
        vm.prank(address(this), address(this));

        vm.expectEmit(true, true, false, true);
        emitTransactionDeposited(
            address(this),
            ZERO_ADDRESS,
            ZERO_VALUE,
            ZERO_VALUE,
            NON_ZERO_GASLIMIT,
            true,
            NON_ZERO_DATA
        );

        portal.depositTransaction(ZERO_ADDRESS, ZERO_VALUE, NON_ZERO_GASLIMIT, true, NON_ZERO_DATA);
    }

    // Test: depositTransaction should emit the correct log when a contract deposits a contract creation with 0 value
    function test_depositTransaction_createWithZeroValueForContract_succeeds() external {
        vm.expectEmit(true, true, false, true);
        emitTransactionDeposited(
            AddressAliasHelper.applyL1ToL2Alias(address(this)),
            ZERO_ADDRESS,
            ZERO_VALUE,
            ZERO_VALUE,
            NON_ZERO_GASLIMIT,
            true,
            NON_ZERO_DATA
        );

        portal.depositTransaction(ZERO_ADDRESS, ZERO_VALUE, NON_ZERO_GASLIMIT, true, NON_ZERO_DATA);
    }

    // Test: depositTransaction should increase its eth balance when an EOA deposits a transaction with ETH
    function test_depositTransaction_withEthValueFromEOA_succeeds() external {
        // EOA emulation
        vm.prank(address(this), address(this));

        vm.expectEmit(true, true, false, true);
        emitTransactionDeposited(
            address(this),
            NON_ZERO_ADDRESS,
            NON_ZERO_VALUE,
            ZERO_VALUE,
            NON_ZERO_GASLIMIT,
            false,
            NON_ZERO_DATA
        );

        portal.depositTransaction{ value: NON_ZERO_VALUE }(
            NON_ZERO_ADDRESS,
            ZERO_VALUE,
            NON_ZERO_GASLIMIT,
            false,
            NON_ZERO_DATA
        );
        assertEq(address(portal).balance, NON_ZERO_VALUE);
    }

    // Test: depositTransaction should increase its eth balance when a contract deposits a transaction with ETH
    function test_depositTransaction_withEthValueFromContract_succeeds() external {
        vm.expectEmit(true, true, false, true);
        emitTransactionDeposited(
            AddressAliasHelper.applyL1ToL2Alias(address(this)),
            NON_ZERO_ADDRESS,
            NON_ZERO_VALUE,
            ZERO_VALUE,
            NON_ZERO_GASLIMIT,
            false,
            NON_ZERO_DATA
        );

        portal.depositTransaction{ value: NON_ZERO_VALUE }(
            NON_ZERO_ADDRESS,
            ZERO_VALUE,
            NON_ZERO_GASLIMIT,
            false,
            NON_ZERO_DATA
        );
    }

    // Test: depositTransaction should increase its eth balance when an EOA deposits a contract creation with ETH
    function test_depositTransaction_withEthValueAndEOAContractCreation_succeeds() external {
        // EOA emulation
        vm.prank(address(this), address(this));

        vm.expectEmit(true, true, false, true);
        emitTransactionDeposited(
            address(this),
            ZERO_ADDRESS,
            NON_ZERO_VALUE,
            ZERO_VALUE,
            NON_ZERO_GASLIMIT,
            true,
            hex""
        );

        portal.depositTransaction{ value: NON_ZERO_VALUE }(
            ZERO_ADDRESS,
            ZERO_VALUE,
            NON_ZERO_GASLIMIT,
            true,
            hex""
        );
        assertEq(address(portal).balance, NON_ZERO_VALUE);
    }

    // Test: depositTransaction should increase its eth balance when a contract deposits a contract creation with ETH
    function test_depositTransaction_withEthValueAndContractContractCreation_succeeds() external {
        vm.expectEmit(true, true, false, true);
        emitTransactionDeposited(
            AddressAliasHelper.applyL1ToL2Alias(address(this)),
            ZERO_ADDRESS,
            NON_ZERO_VALUE,
            ZERO_VALUE,
            NON_ZERO_GASLIMIT,
            true,
            NON_ZERO_DATA
        );

        portal.depositTransaction{ value: NON_ZERO_VALUE }(
            ZERO_ADDRESS,
            ZERO_VALUE,
            NON_ZERO_GASLIMIT,
            true,
            NON_ZERO_DATA
        );
        assertEq(address(portal).balance, NON_ZERO_VALUE);
    }

    function test_simple_isOutputFinalized_succeeds() external {
        uint256 ts = block.timestamp;
        vm.mockCall(
            address(portal.L2_ORACLE()),
            abi.encodeWithSelector(L2OutputOracle.getL2Output.selector),
            abi.encode(
                Types.CheckpointOutput(
                    trusted,
                    bytes32(uint256(1)),
                    uint128(ts),
                    uint128(startingBlockNumber)
                )
            )
        );

        // warp to the finalization period
        vm.warp(ts + oracle.FINALIZATION_PERIOD_SECONDS());
        assertEq(portal.isOutputFinalized(0), false);

        // warp past the finalization period
        vm.warp(ts + oracle.FINALIZATION_PERIOD_SECONDS() + 1);
        assertEq(portal.isOutputFinalized(0), true);
    }

    function test_isOutputFinalized_succeeds() external {
        uint256 checkpoint = oracle.nextBlockNumber();
        uint256 nextOutputIndex = oracle.nextOutputIndex();
        vm.roll(checkpoint);
        warpToSubmitTime();
        vm.prank(trusted);
        oracle.submitL2Output(keccak256(abi.encode(2)), checkpoint, 0, 0);

        // warp to the final second of the finalization period
        uint256 finalizationHorizon = block.timestamp + oracle.FINALIZATION_PERIOD_SECONDS();
        vm.warp(finalizationHorizon);
        // The checkpointed block should not be finalized until 1 second from now.
        assertEq(portal.isOutputFinalized(nextOutputIndex), false);
        // Nor should a block after it
        vm.expectRevert(stdError.indexOOBError);
        assertEq(portal.isOutputFinalized(nextOutputIndex + 1), false);

        // warp past the finalization period
        vm.warp(finalizationHorizon + 1);
        // It should now be finalized.
        assertEq(portal.isOutputFinalized(nextOutputIndex), true);
        // But not the block after it.
        vm.expectRevert(stdError.indexOOBError);
        assertEq(portal.isOutputFinalized(nextOutputIndex + 1), false);
    }
}

contract KromaPortal_FinalizeWithdrawal_Test is Portal_Initializer {
    // Reusable default values for a test withdrawal
    Types.WithdrawalTransaction _defaultTx;

    uint256 _submittedOutputIndex;
    uint256 _submittedBlockNumber;
    bytes32 _stateRoot;
    bytes32 _storageRoot;
    bytes32 _outputRoot;
    bytes32 _withdrawalHash;
    bytes[] _withdrawalProof;
    Types.OutputRootProof internal _outputRootProof;

    // Use a constructor to set the storage vars above, so as to minimize the number of ffi calls.
    constructor() {
        super.setUp();
        _defaultTx = Types.WithdrawalTransaction({
            nonce: 0,
            sender: alice,
            target: bob,
            value: 100,
            gasLimit: 100_000,
            data: hex""
        });
        // Get withdrawal proof data we can use for testing.
        (_stateRoot, _storageRoot, _outputRoot, _withdrawalHash, _withdrawalProof) = ffi
            .getProveWithdrawalTransactionInputs(_defaultTx, true);

        // Setup a dummy output root proof for reuse.
        _outputRootProof = Types.OutputRootProof({
            version: bytes32(uint256(0)),
            stateRoot: _stateRoot,
            messagePasserStorageRoot: _storageRoot,
            blockHash: bytes32(uint256(0)),
            nextBlockHash: bytes32(uint256(0))
        });
        _submittedBlockNumber = oracle.nextBlockNumber();
        _submittedOutputIndex = oracle.nextOutputIndex();
    }

    // Get the system into a nice ready-to-use state.
    function setUp() public override {
        // Configure the oracle to return the output root we've prepared.
        warpToSubmitTime();
        vm.prank(trusted);
        oracle.submitL2Output(_outputRoot, _submittedBlockNumber, 0, 0);

        // Warp beyond the finalization period for the block we've submitted.
        vm.warp(
            oracle.getL2Output(_submittedOutputIndex).timestamp +
                oracle.FINALIZATION_PERIOD_SECONDS() +
                1
        );
        // Fund the portal so that we can withdraw ETH.
        vm.deal(address(portal), 0xFFFFFFFF);
    }

    // Utility function used in the subsequent test. This is necessary to assert that the
    // reentrant call will revert.
    function callPortalAndExpectRevert() external payable {
        vm.expectRevert("KromaPortal: can only trigger one withdrawal per transaction");
        // Arguments here don't matter, as the require check is the first thing that happens.
        // We assume that this has already been proven.
        portal.finalizeWithdrawalTransaction(_defaultTx);
        // Assert that the withdrawal was not finalized.
        assertFalse(portal.finalizedWithdrawals(Hashing.hashWithdrawal(_defaultTx)));
    }

    /**
     * @notice Proving withdrawal transactions should revert when paused
     */
    function test_proveWithdrawalTransaction_paused_reverts() external {
        vm.prank(portal.GUARDIAN());
        portal.pause();

        vm.expectRevert("KromaPortal: paused");
        portal.proveWithdrawalTransaction({
            _tx: _defaultTx,
            _l2OutputIndex: _submittedOutputIndex,
            _outputRootProof: _outputRootProof,
            _withdrawalProof: _withdrawalProof
        });
    }

    // Test: proveWithdrawalTransaction cannot prove a withdrawal with itself (the KromaPortal) as the target.
    function test_proveWithdrawalTransaction_onSelfCall_reverts() external {
        _defaultTx.target = address(portal);
        vm.expectRevert("KromaPortal: you cannot send messages to the portal contract");
        portal.proveWithdrawalTransaction(
            _defaultTx,
            _submittedOutputIndex,
            _outputRootProof,
            _withdrawalProof
        );
    }

    // Test: proveWithdrawalTransaction reverts if the outputRootProof does not match the output root
    function test_proveWithdrawalTransaction_onInvalidOutputRootProof_reverts() external {
        // Modify the version to invalidate the withdrawal proof.
        _outputRootProof.version = bytes32(uint256(MAX_OUTPUT_ROOT_PROOF_VERSION + 1));
        vm.expectRevert("Hashing: unknown output root proof version");
        portal.proveWithdrawalTransaction(
            _defaultTx,
            _submittedOutputIndex,
            _outputRootProof,
            _withdrawalProof
        );
    }

    // Test: proveWithdrawalTransaction reverts if the passed transaction's withdrawalHash has
    // already been proven.
    function test_proveWithdrawalTransaction_replayProve_reverts() external {
        vm.expectEmit(true, true, true, true);
        emit WithdrawalProven(_withdrawalHash, alice, bob);
        portal.proveWithdrawalTransaction(
            _defaultTx,
            _submittedOutputIndex,
            _outputRootProof,
            _withdrawalProof
        );

        vm.expectRevert("KromaPortal: withdrawal hash has already been proven");
        portal.proveWithdrawalTransaction(
            _defaultTx,
            _submittedOutputIndex,
            _outputRootProof,
            _withdrawalProof
        );
    }

    // Test: proveWithdrawalTransaction succeeds if the passed transaction's withdrawalHash has
    // already been proven AND the output root has changed AND the l2BlockNumber stays the same.
    function test_proveWithdrawalTransaction_replayProveChangedOutputRoot_succeeds() external {
        vm.expectEmit(true, true, true, true);
        emit WithdrawalProven(_withdrawalHash, alice, bob);
        portal.proveWithdrawalTransaction(
            _defaultTx,
            _submittedOutputIndex,
            _outputRootProof,
            _withdrawalProof
        );

        // Compute the storage slot of the outputRoot corresponding to the `withdrawalHash`
        // inside of the `provenWithdrawal`s mapping.
        bytes32 slot;
        assembly {
            mstore(0x00, sload(_withdrawalHash.slot))
            mstore(0x20, 52) // 52 is the slot of the `provenWithdrawals` mapping in KromaPortal
            slot := keccak256(0x00, 0x40)
        }

        // Store a different output root within the `provenWithdrawals` mapping without
        // touching the l2BlockNumber or timestamp.
        vm.store(address(portal), slot, bytes32(0));

        // Warp ahead 1 second
        vm.warp(block.timestamp + 1);

        // Even though we have already proven this withdrawalHash, we should be allowed to re-submit
        // our proof with a changed outputRoot
        vm.expectEmit(true, true, true, true);
        emit WithdrawalProven(_withdrawalHash, alice, bob);
        portal.proveWithdrawalTransaction(
            _defaultTx,
            _submittedOutputIndex,
            _outputRootProof,
            _withdrawalProof
        );

        // Ensure that the withdrawal was updated within the mapping
        (, uint128 timestamp, ) = portal.provenWithdrawals(_withdrawalHash);
        assertEq(timestamp, block.timestamp);
    }

    // Test: proveWithdrawalTransaction succeeds if the passed transaction's withdrawalHash has
    // already been proven AND the output root + output index + l2BlockNumber changes.
    function test_proveWithdrawalTransaction_replayProveChangedOutputRootAndOutputIndex_succeeds()
        external
    {
        vm.expectEmit(true, true, true, true);
        emit WithdrawalProven(_withdrawalHash, alice, bob);
        portal.proveWithdrawalTransaction(
            _defaultTx,
            _submittedOutputIndex,
            _outputRootProof,
            _withdrawalProof
        );

        // Compute the storage slot of the outputRoot corresponding to the `withdrawalHash`
        // inside of the `provenWithdrawal`s mapping.
        bytes32 slot;
        assembly {
            mstore(0x00, sload(_withdrawalHash.slot))
            mstore(0x20, 52) // 52 is the slot of the `provenWithdrawals` mapping in KromaPortal
            slot := keccak256(0x00, 0x40)
        }

        // Store a dummy output root within the `provenWithdrawals` mapping without touching the
        // l2BlockNumber or timestamp.
        vm.store(address(portal), slot, bytes32(0));

        // Fetch the checkpoint output at `_submittedOutputIndex` from the L2OutputOracle
        Types.CheckpointOutput memory output = portal.L2_ORACLE().getL2Output(
            _submittedOutputIndex
        );

        // Propose the same output root again, creating the same output at a different index + l2BlockNumber.
        vm.startPrank(trusted);
        portal.L2_ORACLE().submitL2Output(
            output.outputRoot,
            portal.L2_ORACLE().nextBlockNumber(),
            blockhash(block.number),
            block.number
        );
        vm.stopPrank();

        // Warp ahead 1 second
        vm.warp(block.timestamp + 1);

        // Even though we have already proven this withdrawalHash, we should be allowed to re-submit
        // our proof with a changed outputRoot + a different output index
        vm.expectEmit(true, true, true, true);
        emit WithdrawalProven(_withdrawalHash, alice, bob);
        portal.proveWithdrawalTransaction(
            _defaultTx,
            _submittedOutputIndex + 1,
            _outputRootProof,
            _withdrawalProof
        );

        // Ensure that the withdrawal was updated within the mapping
        (, uint128 timestamp, ) = portal.provenWithdrawals(_withdrawalHash);
        assertEq(timestamp, block.timestamp);
    }

    // Test: proveWithdrawalTransaction succeeds and emits the WithdrawalProven event.
    function test_proveWithdrawalTransaction_validWithdrawalProof_succeeds() external {
        vm.expectEmit(true, true, true, true);
        emit WithdrawalProven(_withdrawalHash, alice, bob);
        portal.proveWithdrawalTransaction(
            _defaultTx,
            _submittedOutputIndex,
            _outputRootProof,
            _withdrawalProof
        );
    }

    // Test: proveWithdrawalTransaction succeeds when nextBlockHash is not zero.
    function test_proveWithdrawalTransaction_nextBlockHashNotZero_succeeds() external {
        // Get modified proof inputs when isKromaMPT is false.
        (_stateRoot, _storageRoot, _outputRoot, _withdrawalHash, _withdrawalProof) = ffi
            .getProveWithdrawalTransactionInputs(_defaultTx, false);

        // Create the output root proof with non zero nextBlockHash
        _outputRootProof = Types.OutputRootProof({
            version: bytes32(uint256(0)),
            stateRoot: _stateRoot,
            messagePasserStorageRoot: _storageRoot,
            blockHash: bytes32(uint256(0)),
            nextBlockHash: bytes32(uint256(1))
        });

        // Setup the Oracle to return the outputRoot
        vm.mockCall(
            address(oracle),
            abi.encodeWithSelector(oracle.getL2Output.selector),
            abi.encode(trusted, _outputRoot, block.timestamp, _submittedBlockNumber)
        );

        // Prove the withdrawal transaction
        portal.proveWithdrawalTransaction(
            _defaultTx,
            _submittedOutputIndex,
            _outputRootProof,
            _withdrawalProof
        );
    }

    // Test: finalizeWithdrawalTransaction succeeds and emits the WithdrawalFinalized event.
    function test_finalizeWithdrawalTransaction_provenWithdrawalHash_succeeds() external {
        uint256 bobBalanceBefore = address(bob).balance;

        vm.expectEmit(true, true, true, true);
        emit WithdrawalProven(_withdrawalHash, alice, bob);
        portal.proveWithdrawalTransaction(
            _defaultTx,
            _submittedOutputIndex,
            _outputRootProof,
            _withdrawalProof
        );

        vm.warp(block.timestamp + oracle.FINALIZATION_PERIOD_SECONDS() + 1);
        vm.expectEmit(true, true, false, true);
        emit WithdrawalFinalized(_withdrawalHash, true);
        portal.finalizeWithdrawalTransaction(_defaultTx);

        assert(address(bob).balance == bobBalanceBefore + 100);
    }

    /**
     * @notice Finalizing withdrawal transactions should revert when paused
     */
    function test_finalizeWithdrawalTransaction_paused_reverts() external {
        vm.prank(portal.GUARDIAN());
        portal.pause();

        vm.expectRevert("KromaPortal: paused");
        portal.finalizeWithdrawalTransaction(_defaultTx);
    }

    // Test: finalizeWithdrawalTransaction reverts if the withdrawal has not been proven.
    function test_finalizeWithdrawalTransaction_ifWithdrawalNotProven_reverts() external {
        uint256 bobBalanceBefore = address(bob).balance;

        vm.expectRevert("KromaPortal: withdrawal has not been proven yet");
        portal.finalizeWithdrawalTransaction(_defaultTx);

        assert(address(bob).balance == bobBalanceBefore);
    }

    // Test: finalizeWithdrawalTransaction reverts if withdrawal not proven long enough ago.
    function test_finalizeWithdrawalTransaction_ifWithdrawalProofNotOldEnough_reverts() external {
        uint256 bobBalanceBefore = address(bob).balance;

        vm.expectEmit(true, true, true, true);
        emit WithdrawalProven(_withdrawalHash, alice, bob);
        portal.proveWithdrawalTransaction(
            _defaultTx,
            _submittedOutputIndex,
            _outputRootProof,
            _withdrawalProof
        );

        // Mock a call where the resulting output root is anything but the original output root. In
        // this case we just use bytes32(uint256(1)).
        vm.mockCall(
            address(portal.L2_ORACLE()),
            abi.encodeWithSelector(L2OutputOracle.getL2Output.selector),
            abi.encode(bytes32(uint256(1)), _submittedBlockNumber)
        );

        vm.expectRevert("KromaPortal: proven withdrawal finalization period has not elapsed");
        portal.finalizeWithdrawalTransaction(_defaultTx);

        assert(address(bob).balance == bobBalanceBefore);
    }

    // Test: finalizeWithdrawalTransaction reverts if the provenWithdrawal's timestamp is less
    // than the L2 output oracle's starting timestamp
    function test_finalizeWithdrawalTransaction_timestampLessThanL2OracleStart_reverts() external {
        uint256 bobBalanceBefore = address(bob).balance;

        // Prove our withdrawal
        vm.expectEmit(true, true, true, true);
        emit WithdrawalProven(_withdrawalHash, alice, bob);
        portal.proveWithdrawalTransaction(
            _defaultTx,
            _submittedOutputIndex,
            _outputRootProof,
            _withdrawalProof
        );

        // Warp to after the finalization period
        vm.warp(block.timestamp + oracle.FINALIZATION_PERIOD_SECONDS() + 1);

        // Mock a startingTimestamp change on the L2 Oracle
        vm.mockCall(
            address(portal.L2_ORACLE()),
            abi.encodeWithSignature("startingTimestamp()"),
            abi.encode(block.timestamp + 1)
        );

        // Attempt to finalize the withdrawal
        vm.expectRevert("KromaPortal: withdrawal timestamp less than L2 Oracle starting timestamp");
        portal.finalizeWithdrawalTransaction(_defaultTx);

        // Ensure that bob's balance has remained the same
        assertEq(bobBalanceBefore, address(bob).balance);
    }

    // Test: finalizeWithdrawalTransaction reverts if the output root proven is not the same as the
    // output root at the time of finalization.
    function test_finalizeWithdrawalTransaction_ifOutputRootChanges_reverts() external {
        uint256 bobBalanceBefore = address(bob).balance;

        // Prove our withdrawal
        vm.expectEmit(true, true, true, true);
        emit WithdrawalProven(_withdrawalHash, alice, bob);
        portal.proveWithdrawalTransaction(
            _defaultTx,
            _submittedOutputIndex,
            _outputRootProof,
            _withdrawalProof
        );

        // Warp to after the finalization period
        vm.warp(block.timestamp + oracle.FINALIZATION_PERIOD_SECONDS() + 1);

        // Mock an outputRoot change on the checkpoint output before attempting
        // to finalize the withdrawal.
        vm.mockCall(
            address(portal.L2_ORACLE()),
            abi.encodeWithSelector(L2OutputOracle.getL2Output.selector),
            abi.encode(
                Types.CheckpointOutput(
                    trusted,
                    bytes32(uint256(0)),
                    uint128(block.timestamp),
                    uint128(_submittedBlockNumber)
                )
            )
        );

        // Attempt to finalize the withdrawal
        vm.expectRevert("KromaPortal: output root proven is not the same as current output root");
        portal.finalizeWithdrawalTransaction(_defaultTx);

        // Ensure that bob's balance has remained the same
        assertEq(bobBalanceBefore, address(bob).balance);
    }

    // Test: finalizeWithdrawalTransaction reverts if the checkpoint output's timestamp has
    // not passed the finalization period.
    function test_finalizeWithdrawalTransaction_ifOutputTimestampIsNotFinalized_reverts() external {
        uint256 bobBalanceBefore = address(bob).balance;

        // Prove our withdrawal
        vm.expectEmit(true, true, true, true);
        emit WithdrawalProven(_withdrawalHash, alice, bob);
        portal.proveWithdrawalTransaction(
            _defaultTx,
            _submittedOutputIndex,
            _outputRootProof,
            _withdrawalProof
        );

        // Warp to after the finalization period
        vm.warp(block.timestamp + oracle.FINALIZATION_PERIOD_SECONDS() + 1);

        // Mock a timestamp change on the checkpoint output that has not passed the
        // finalization period.
        vm.mockCall(
            address(portal.L2_ORACLE()),
            abi.encodeWithSelector(L2OutputOracle.getL2Output.selector),
            abi.encode(
                Types.CheckpointOutput(
                    trusted,
                    _outputRoot,
                    uint128(block.timestamp + 1),
                    uint128(_submittedBlockNumber)
                )
            )
        );

        // Attempt to finalize the withdrawal
        vm.expectRevert("KromaPortal: checkpoint output finalization period has not elapsed");
        portal.finalizeWithdrawalTransaction(_defaultTx);

        // Ensure that bob's balance has remained the same
        assertEq(bobBalanceBefore, address(bob).balance);
    }

    // Test: finalizeWithdrawalTransaction fails because the target reverts,
    // and emits the WithdrawalFinalized event with success=false.
    function test_finalizeWithdrawalTransaction_targetFails_fails() external {
        uint256 bobBalanceBefore = address(bob).balance;
        vm.etch(bob, hex"fe"); // Contract with just the invalid opcode.

        vm.expectEmit(true, true, true, true);
        emit WithdrawalProven(_withdrawalHash, alice, bob);
        portal.proveWithdrawalTransaction(
            _defaultTx,
            _submittedOutputIndex,
            _outputRootProof,
            _withdrawalProof
        );

        vm.warp(block.timestamp + oracle.FINALIZATION_PERIOD_SECONDS() + 1);
        vm.expectEmit(true, true, true, true);
        emit WithdrawalFinalized(_withdrawalHash, false);
        portal.finalizeWithdrawalTransaction(_defaultTx);

        assert(address(bob).balance == bobBalanceBefore);
    }

    // Test: finalizeWithdrawalTransaction reverts if the finalization period has not yet passed.
    function test_finalizeWithdrawalTransaction_onRecentWithdrawal_reverts() external {
        // Setup the Oracle to return an output with a recent timestamp
        uint256 recentTimestamp = block.timestamp - 1000;
        vm.mockCall(
            address(portal.L2_ORACLE()),
            abi.encodeWithSelector(L2OutputOracle.getL2Output.selector),
            abi.encode(
                Types.CheckpointOutput(
                    trusted,
                    _outputRoot,
                    uint128(recentTimestamp),
                    uint128(_submittedBlockNumber)
                )
            )
        );

        portal.proveWithdrawalTransaction(
            _defaultTx,
            _submittedOutputIndex,
            _outputRootProof,
            _withdrawalProof
        );

        vm.expectRevert("KromaPortal: proven withdrawal finalization period has not elapsed");
        portal.finalizeWithdrawalTransaction(_defaultTx);
    }

    // Test: finalizeWithdrawalTransaction reverts if the withdrawal has already been finalized.
    function test_finalizeWithdrawalTransaction_onReplay_reverts() external {
        vm.expectEmit(true, true, true, true);
        emit WithdrawalProven(_withdrawalHash, alice, bob);
        portal.proveWithdrawalTransaction(
            _defaultTx,
            _submittedOutputIndex,
            _outputRootProof,
            _withdrawalProof
        );

        vm.warp(block.timestamp + oracle.FINALIZATION_PERIOD_SECONDS() + 1);
        vm.expectEmit(true, true, true, true);
        emit WithdrawalFinalized(_withdrawalHash, true);
        portal.finalizeWithdrawalTransaction(_defaultTx);

        vm.expectRevert("KromaPortal: withdrawal has already been finalized");
        portal.finalizeWithdrawalTransaction(_defaultTx);
    }

    // Test: finalizeWithdrawalTransaction reverts if insufficient gas is supplied.
    function test_finalizeWithdrawalTransaction_onInsufficientGas_reverts() external {
        // This number was identified through trial and error.
        uint256 gasLimit = 150_000;
        Types.WithdrawalTransaction memory insufficientGasTx = Types.WithdrawalTransaction({
            nonce: 0,
            sender: alice,
            target: bob,
            value: 100,
            gasLimit: gasLimit,
            data: hex""
        });

        // Get updated proof inputs.
        (bytes32 stateRoot, bytes32 storageRoot, , , bytes[] memory withdrawalProof) = ffi
            .getProveWithdrawalTransactionInputs(insufficientGasTx, true);
        Types.OutputRootProof memory outputRootProof = Types.OutputRootProof({
            version: bytes32(uint256(0)),
            stateRoot: stateRoot,
            messagePasserStorageRoot: storageRoot,
            blockHash: bytes32(uint256(0)),
            nextBlockHash: bytes32(uint256(0))
        });

        vm.mockCall(
            address(portal.L2_ORACLE()),
            abi.encodeWithSelector(L2OutputOracle.getL2Output.selector),
            abi.encode(
                Types.CheckpointOutput(
                    trusted,
                    Hashing.hashOutputRootProof(outputRootProof),
                    uint128(block.timestamp),
                    uint128(_submittedBlockNumber)
                )
            )
        );

        portal.proveWithdrawalTransaction(
            insufficientGasTx,
            _submittedOutputIndex,
            outputRootProof,
            withdrawalProof
        );

        vm.warp(block.timestamp + oracle.FINALIZATION_PERIOD_SECONDS() + 1);
        vm.expectRevert("SafeCall: Not enough gas");
        portal.finalizeWithdrawalTransaction{ gas: gasLimit }(insufficientGasTx);
    }

    // Test: finalizeWithdrawalTransaction reverts if a sub-call attempts to finalize another
    // withdrawal.
    function test_finalizeWithdrawalTransaction_onReentrancy_reverts() external {
        uint256 bobBalanceBefore = address(bob).balance;

        // Copy and modify the default test values to attempt a reentrant call by first calling to
        // this contract's callPortalAndExpectRevert() function above.
        Types.WithdrawalTransaction memory _testTx = _defaultTx;
        _testTx.target = address(this);
        _testTx.data = abi.encodeWithSelector(this.callPortalAndExpectRevert.selector);

        // Get modified proof inputs.
        (
            bytes32 stateRoot,
            bytes32 storageRoot,
            bytes32 outputRoot,
            bytes32 withdrawalHash,
            bytes[] memory withdrawalProof
        ) = ffi.getProveWithdrawalTransactionInputs(_testTx, true);
        Types.OutputRootProof memory outputRootProof = Types.OutputRootProof({
            version: bytes32(uint256(0)),
            stateRoot: stateRoot,
            messagePasserStorageRoot: storageRoot,
            blockHash: bytes32(uint256(0)),
            nextBlockHash: bytes32(uint256(0))
        });

        // Setup the Oracle to return the outputRoot we want as well as a finalized timestamp.
        uint256 finalizedTimestamp = block.timestamp - oracle.FINALIZATION_PERIOD_SECONDS() - 1;
        vm.mockCall(
            address(portal.L2_ORACLE()),
            abi.encodeWithSelector(L2OutputOracle.getL2Output.selector),
            abi.encode(
                Types.CheckpointOutput(
                    trusted,
                    outputRoot,
                    uint128(finalizedTimestamp),
                    uint128(_submittedBlockNumber)
                )
            )
        );

        vm.expectEmit(true, true, true, true);
        emit WithdrawalProven(withdrawalHash, alice, address(this));
        portal.proveWithdrawalTransaction(
            _testTx,
            _submittedBlockNumber,
            outputRootProof,
            withdrawalProof
        );

        vm.warp(block.timestamp + oracle.FINALIZATION_PERIOD_SECONDS() + 1);
        vm.expectCall(address(this), _testTx.data);
        vm.expectEmit(true, true, true, true);
        emit WithdrawalFinalized(withdrawalHash, true);
        portal.finalizeWithdrawalTransaction(_testTx);

        // Ensure that bob's balance was not changed by the reentrant call.
        assert(address(bob).balance == bobBalanceBefore);
    }

    function testDiff_finalizeWithdrawalTransaction_succeeds(
        address _sender,
        address _target,
        uint256 _value,
        uint256 _gasLimit,
        bytes memory _data
    ) external {
        vm.assume(
            _target != address(portal) && // Cannot call the kroma portal or a contract
                _target.code.length == 0 && // No accounts with code
                _target != CONSOLE && // The console has no code but behaves like a contract
                uint160(_target) > 9 // No precompiles (or zero address)
        );

        // Total ETH supply is currently about 120M ETH.
        uint256 value = bound(_value, 0, 200_000_000 ether);
        vm.deal(address(portal), value);

        uint256 gasLimit = bound(_gasLimit, 0, 50_000_000);
        uint256 nonce = messagePasser.messageNonce();

        // Get a withdrawal transaction and mock proof from the differential testing script.
        Types.WithdrawalTransaction memory _tx = Types.WithdrawalTransaction({
            nonce: nonce,
            sender: _sender,
            target: _target,
            value: value,
            gasLimit: gasLimit,
            data: _data
        });
        (
            bytes32 stateRoot,
            bytes32 storageRoot,
            bytes32 outputRoot,
            bytes32 withdrawalHash,
            bytes[] memory withdrawalProof
        ) = ffi.getProveWithdrawalTransactionInputs(_tx, true);

        // Create the output root proof
        Types.OutputRootProof memory proof = Types.OutputRootProof({
            version: bytes32(uint256(0)),
            stateRoot: stateRoot,
            messagePasserStorageRoot: storageRoot,
            blockHash: bytes32(uint256(0)),
            nextBlockHash: bytes32(uint256(0))
        });

        // Ensure the values returned from ffi are correct
        assertEq(outputRoot, Hashing.hashOutputRootProof(proof));
        assertEq(withdrawalHash, Hashing.hashWithdrawal(_tx));

        // Setup the Oracle to return the outputRoot
        vm.mockCall(
            address(oracle),
            abi.encodeWithSelector(oracle.getL2Output.selector),
            abi.encode(address(0), outputRoot, block.timestamp, 100)
        );

        // Prove the withdrawal transaction
        portal.proveWithdrawalTransaction(
            _tx,
            1, // l2OutputIndex
            proof,
            withdrawalProof
        );
        (bytes32 _root, , ) = portal.provenWithdrawals(withdrawalHash);
        assertTrue(_root != bytes32(0));

        // Warp past the finalization period
        vm.warp(block.timestamp + oracle.FINALIZATION_PERIOD_SECONDS() + 1);

        // Finalize the withdrawal transaction
        vm.expectCallMinGas(_tx.target, _tx.value, uint64(_tx.gasLimit), _tx.data);
        portal.finalizeWithdrawalTransaction(_tx);
        assertTrue(portal.finalizedWithdrawals(withdrawalHash));
    }
}

contract KromaPortalUpgradeable_Test is Portal_Initializer {
    Proxy internal proxy;
    uint64 initialBlockNum;

    function setUp() public override {
        super.setUp();
        initialBlockNum = uint64(block.number);
        proxy = Proxy(payable(address(portal)));
    }

    function test_params_initValuesOnProxy_succeeds() external {
        KromaPortal p = KromaPortal(payable(address(proxy)));
        (uint128 prevBaseFee, uint64 prevBoughtGas, uint64 prevBlockNum) = p.params();
        ResourceMetering.ResourceConfig memory rcfg = systemConfig.resourceConfig();

        assertEq(prevBaseFee, rcfg.minimumBaseFee);
        assertEq(prevBoughtGas, 0);
        assertEq(prevBlockNum, initialBlockNum);
    }

    function test_initialize_cannotInitProxy_reverts() external {
        vm.expectRevert("Initializable: contract is already initialized");
        KromaPortal(payable(proxy)).initialize(false);
    }

    function test_initialize_cannotInitImpl_reverts() external {
        vm.expectRevert("Initializable: contract is already initialized");
        KromaPortal(portalImpl).initialize(false);
    }

    function test_upgradeToAndCall_upgrading_succeeds() external {
        // Check an unused slot before upgrading.
        bytes32 slot21Before = vm.load(address(portal), bytes32(uint256(21)));
        assertEq(bytes32(0), slot21Before);

        NextImpl nextImpl = new NextImpl();
        vm.startPrank(multisig);
        proxy.upgradeToAndCall(
            address(nextImpl),
            abi.encodeWithSelector(NextImpl.initialize.selector)
        );
        assertEq(proxy.implementation(), address(nextImpl));

        // Verify that the NextImpl contract initialized its values according as expected
        bytes32 slot21After = vm.load(address(portal), bytes32(uint256(21)));
        bytes32 slot21Expected = NextImpl(address(portal)).slot21Init();
        assertEq(slot21Expected, slot21After);
    }
}

/**
 * @title KromaPortalResourceFuzz_Test
 * @dev Test various values of the resource metering config to ensure that deposits cannot be
 *         broken by changing the config.
 */
contract KromaPortalResourceFuzz_Test is Portal_Initializer {
    /**
     * @dev The max gas limit observed throughout this test. Setting this too high can cause
     *      the test to take too long to run.
     */
    uint256 constant MAX_GAS_LIMIT = 30_000_000;

    /**
     * @dev Test that various values of the resource metering config will not break deposits.
     */
    function testFuzz_systemConfigDeposit_succeeds(
        uint32 _maxResourceLimit,
        uint8 _elasticityMultiplier,
        uint8 _baseFeeMaxChangeDenominator,
        uint32 _minimumBaseFee,
        uint32 _systemTxMaxGas,
        uint128 _maximumBaseFee,
        uint64 _gasLimit,
        uint64 _prevBoughtGas,
        uint128 _prevBaseFee,
        uint8 _blockDiff
    ) external {
        // Get the set system gas limit
        uint64 gasLimit = systemConfig.gasLimit();
        // Bound resource config
        _maxResourceLimit = uint32(bound(_maxResourceLimit, 21000, MAX_GAS_LIMIT / 8));
        _gasLimit = uint64(bound(_gasLimit, 21000, _maxResourceLimit));
        _prevBaseFee = uint128(bound(_prevBaseFee, 0, 5 gwei));
        // Prevent values that would cause reverts
        vm.assume(gasLimit >= _gasLimit);
        vm.assume(_minimumBaseFee < _maximumBaseFee);
        vm.assume(_baseFeeMaxChangeDenominator > 1);
        vm.assume(uint256(_maxResourceLimit) + uint256(_systemTxMaxGas) <= gasLimit);
        vm.assume(_elasticityMultiplier > 0);
        vm.assume(
            ((_maxResourceLimit / _elasticityMultiplier) * _elasticityMultiplier) ==
                _maxResourceLimit
        );
        _prevBoughtGas = uint64(bound(_prevBoughtGas, 0, _maxResourceLimit - _gasLimit));
        _blockDiff = uint8(bound(_blockDiff, 0, 3));

        // Create a resource config to mock the call to the system config with
        ResourceMetering.ResourceConfig memory rcfg = ResourceMetering.ResourceConfig({
            maxResourceLimit: _maxResourceLimit,
            elasticityMultiplier: _elasticityMultiplier,
            baseFeeMaxChangeDenominator: _baseFeeMaxChangeDenominator,
            minimumBaseFee: _minimumBaseFee,
            systemTxMaxGas: _systemTxMaxGas,
            maximumBaseFee: _maximumBaseFee
        });
        vm.mockCall(
            address(systemConfig),
            abi.encodeWithSelector(systemConfig.resourceConfig.selector),
            abi.encode(rcfg)
        );

        // Set the resource params
        uint256 _prevBlockNum = block.number - _blockDiff;
        vm.store(
            address(portal),
            bytes32(uint256(1)),
            bytes32((_prevBlockNum << 192) | (uint256(_prevBoughtGas) << 128) | _prevBaseFee)
        );
        // Ensure that the storage setting is correct
        (uint128 prevBaseFee, uint64 prevBoughtGas, uint64 prevBlockNum) = portal.params();
        assertEq(prevBaseFee, _prevBaseFee);
        assertEq(prevBoughtGas, _prevBoughtGas);
        assertEq(prevBlockNum, _prevBlockNum);

        // Do a deposit, should not revert
        portal.depositTransaction{ gas: MAX_GAS_LIMIT }({
            _to: address(0x20),
            _value: 0x40,
            _gasLimit: _gasLimit,
            _isCreation: false,
            _data: hex""
        });
    }
}
