pragma solidity 0.8.15;

import { Types } from "../../libraries/Types.sol";
import { Portal_Initializer } from "../CommonTest.t.sol";

contract KromaPortal_Invariant_Harness is Portal_Initializer {
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

    function setUp() public virtual override {
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
            .getProveWithdrawalTransactionInputs(_defaultTx);

        // Setup a dummy output root proof for reuse.
        _outputRootProof = Types.OutputRootProof({
            version: bytes32(uint256(1)),
            stateRoot: _stateRoot,
            messagePasserStorageRoot: _storageRoot,
            blockHash: bytes32(uint256(0)),
            nextBlockHash: bytes32(uint256(0))
        });
        _submittedBlockNumber = oracle.nextBlockNumber();
        _submittedOutputIndex = oracle.nextOutputIndex();

        // Configure the oracle to return the output root we've prepared.
        vm.warp(oracle.computeL2Timestamp(_submittedBlockNumber) + 1);
        vm.prank(trusted);
        oracle.submitL2Output(_outputRoot, _submittedBlockNumber, 0, 0, minBond);

        // Warp beyond the finalization period for the block we've submitted.
        vm.warp(
            oracle.getL2Output(_submittedOutputIndex).timestamp +
                oracle.FINALIZATION_PERIOD_SECONDS() +
                1
        );
        // Fund the portal so that we can withdraw ETH.
        vm.deal(address(portal), 0xFFFFFFFF);
    }
}

contract KromaPortal_CannotTimeTravel is KromaPortal_Invariant_Harness {
    function setUp() public override {
        super.setUp();

        // Prove the withdrawal transaction
        portal.proveWithdrawalTransaction(
            _defaultTx,
            _submittedOutputIndex,
            _outputRootProof,
            _withdrawalProof
        );

        // Set the target contract to the portal proxy
        targetContract(address(portal));
        // Exclude the proxy multisig from the senders so that the proxy cannot be upgraded
        excludeSender(address(multisig));
    }

    /**
     * @custom:invariant `finalizeWithdrawalTransaction` should revert if the finalization
     * period has not elapsed.
     *
     * A withdrawal that has been proven should not be able to be finalized until after
     * the finalization period has elapsed.
     */
    function invariant_cannotFinalizeBeforePeriodHasPassed() external {
        vm.expectRevert("KromaPortal: proven withdrawal finalization period has not elapsed");
        portal.finalizeWithdrawalTransaction(_defaultTx);
    }
}

contract KromaPortal_CannotFinalizeTwice is KromaPortal_Invariant_Harness {
    function setUp() public override {
        super.setUp();

        // Prove the withdrawal transaction
        portal.proveWithdrawalTransaction(
            _defaultTx,
            _submittedOutputIndex,
            _outputRootProof,
            _withdrawalProof
        );

        // Warp past the finalization period.
        vm.warp(block.timestamp + oracle.FINALIZATION_PERIOD_SECONDS() + 1);

        // Finalize the withdrawal transaction.
        portal.finalizeWithdrawalTransaction(_defaultTx);

        // Set the target contract to the portal proxy
        targetContract(address(portal));
        // Exclude the proxy multisig from the senders so that the proxy cannot be upgraded
        excludeSender(address(multisig));
    }

    /**
     * @custom:invariant `finalizeWithdrawalTransaction` should revert if the withdrawal
     * has already been finalized.
     *
     * Ensures that there is no chain of calls that can be made that allows a withdrawal
     * to be finalized twice.
     */
    function invariant_cannotFinalizeTwice() external {
        vm.expectRevert("KromaPortal: withdrawal has already been finalized");
        portal.finalizeWithdrawalTransaction(_defaultTx);
    }
}

contract KromaPortal_CanAlwaysFinalizeAfterWindow is KromaPortal_Invariant_Harness {
    function setUp() public override {
        super.setUp();

        // Prove the withdrawal transaction
        portal.proveWithdrawalTransaction(
            _defaultTx,
            _submittedOutputIndex,
            _outputRootProof,
            _withdrawalProof
        );

        // Warp past the finalization period.
        vm.warp(block.timestamp + oracle.FINALIZATION_PERIOD_SECONDS() + 1);

        // Set the target contract to the portal proxy
        targetContract(address(portal));
        // Exclude the proxy multisig from the senders so that the proxy cannot be upgraded
        excludeSender(address(multisig));
    }

    /**
     * @custom:invariant A withdrawal should **always** be able to be finalized
     * `FINALIZATION_PERIOD_SECONDS` after it was successfully proven.
     *
     * This invariant asserts that there is no chain of calls that can be made that
     * will prevent a withdrawal from being finalized exactly `FINALIZATION_PERIOD_SECONDS`
     * after it was successfully proven.
     */
    function invariant_canAlwaysFinalize() external {
        uint256 bobBalanceBefore = address(bob).balance;

        portal.finalizeWithdrawalTransaction(_defaultTx);

        assertEq(address(bob).balance, bobBalanceBefore + _defaultTx.value);
    }
}
