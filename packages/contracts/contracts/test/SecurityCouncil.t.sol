// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { Colosseum } from "../L1/Colosseum.sol";
import { Types } from "../libraries/Types.sol";
import { SecurityCouncil_Initializer } from "./CommonTest.t.sol";

contract SecurityCouncilTest is SecurityCouncil_Initializer {
    /**
     *  Events
     */
    event TransactionSubmitted(address indexed sender, uint256 indexed transactionId);
    event TransactionConfirmed(address indexed sender, uint256 indexed transactionId);
    event TransactionExecuted(address indexed sender, uint256 indexed transactionId);
    event ConfirmationRevoked(address indexed sender, uint256 indexed transactionId);
    event ValidationRequested(
        uint256 indexed transactionId,
        bytes32 outputRoot,
        uint256 l2BlockNumber
    );
    event DeletionRequested(uint256 indexed transactionId, uint256 indexed outputIndex);

    uint256 blockNumber = 0;
    address private txTarget;
    uint256 private txValue;
    bytes private txData;

    function setUp() public virtual override {
        super.setUp();
        // dummy transaction data
        txTarget = makeAddr("target");
        txValue = 0;
        txData = bytes("anydata");

        _roll(1);
        // minting to guardians
        vm.prank(owner);
        securityCouncilToken.safeMint(guardian1, baseUri);
        vm.prank(owner);
        securityCouncilToken.safeMint(guardian2, baseUri);
        vm.prank(owner);
        securityCouncilToken.safeMint(guardian3, baseUri);

        assertEq(securityCouncilToken.balanceOf(guardian1), 1);
        assertEq(securityCouncilToken.balanceOf(guardian2), 1);
        assertEq(securityCouncilToken.balanceOf(guardian3), 1);
        assertEq(securityCouncilToken.getVotes(guardian1), 1);
        assertEq(securityCouncilToken.getVotes(guardian2), 1);
        assertEq(securityCouncilToken.getVotes(guardian3), 1);
        assertEq(securityCouncilToken.owner(), owner);

        _roll(1);
    }

    function test_initialize_succeeds() external {
        assertEq(address(securityCouncil.COLOSSEUM()), colosseum);
        assertEq(address(securityCouncil.GOVERNOR()), address(upgradeGovernor));
    }

    function test_submitTransaction_onlyTokenOwner_reverts() external {
        vm.prank(makeAddr("not governance"));
        vm.expectRevert("TokenMultiSigWallet: only allowed to governance token owner");
        securityCouncil.submitTransaction(txTarget, txValue, txData);
    }

    function test_submitTransaction_targetInvalid_reverts() external {
        vm.startPrank(guardian1);
        vm.expectRevert("TokenMultiSigWallet: address is not valid");
        securityCouncil.submitTransaction(address(0), txValue, txData);
        vm.stopPrank();
    }

    function test_submitTransaction_transactionExists_reverts() external {
        vm.startPrank(guardian1);
        securityCouncil.submitTransaction(txTarget, txValue, txData);
        vm.expectRevert("TokenMultiSigWallet: transaction already exists");
        securityCouncil.submitTransaction(txTarget, txValue, txData);
        vm.stopPrank();
    }

    function test_submitTransaction_succeeds() external {
        vm.startPrank(guardian1);
        vm.expectEmit(true, true, false, true);
        uint256 transactionId = securityCouncil.generateTransactionId(txTarget, txValue, txData);
        emit TransactionSubmitted(guardian1, transactionId);
        uint256 resTransactionId = securityCouncil.submitTransaction(txTarget, txValue, txData);
        assertEq(transactionId, resTransactionId);
        (,bool executed,,) = securityCouncil.transactions(transactionId);
        assertEq(executed, false);
        vm.stopPrank();
    }

    function test_confirmTransaction_onlyTokenOwner_reverts() external {
        vm.prank(guardian1);
        uint256 transactionId = securityCouncil.submitTransaction(txTarget, txValue, txData);

        vm.startPrank(makeAddr("not governance"));
        vm.expectRevert("TokenMultiSigWallet: only allowed to governance token owner");
        securityCouncil.confirmTransaction(transactionId);
        vm.stopPrank();
    }

    function test_confirmTransaction_nonExistent_reverts() external {
        vm.startPrank(guardian1);
        vm.expectRevert("TokenMultiSigWallet: transaction does not exist");
        securityCouncil.confirmTransaction(0);
        vm.stopPrank();
    }

    function test_confirmTransaction_alreadyConfirmed_reverts() external {
        vm.startPrank(guardian1);
        uint256 transactionId = securityCouncil.submitTransaction(txTarget, txValue, txData);
        securityCouncil.confirmTransaction(transactionId);
        vm.expectRevert("TokenMultiSigWallet: already confirmed");
        securityCouncil.confirmTransaction(transactionId);
        vm.stopPrank();
    }

    function test_confirmTransaction_succeeds() external {
        vm.startPrank(guardian1);
        uint256 transactionId = securityCouncil.submitTransaction(txTarget, txValue, txData);
        vm.expectEmit(true, true, false, true);
        emit TransactionConfirmed(guardian1, transactionId);
        uint256 confirmCountOld = securityCouncil.getConfirmationCount(transactionId);
        securityCouncil.confirmTransaction(transactionId);
        uint256 confirmCountNew = securityCouncil.getConfirmationCount(transactionId);
        assertEq(confirmCountOld + securityCouncil.getVotes(guardian1), confirmCountNew);
        bool confirmed = securityCouncil.isConfirmedBy(transactionId, guardian1);
        assertTrue(confirmed);
        vm.stopPrank();
    }

    function test_executeTransaction_succeeds() external {
        vm.startPrank(guardian1);
        uint256 transactionId = securityCouncil.submitTransaction(txTarget, txValue, txData);
        securityCouncil.confirmTransaction(transactionId);
        vm.stopPrank();

        vm.startPrank(guardian2);
        vm.expectEmit(true, true, false, true);
        emit TransactionConfirmed(guardian2, transactionId);
        vm.expectEmit(true, true, false, true);
        emit TransactionExecuted(guardian2, transactionId);
        securityCouncil.confirmTransaction(transactionId);
        vm.stopPrank();
    }

    function test_requestValidation_reverts() external {
        vm.prank(makeAddr("not colosseum"));
        vm.expectRevert("SecurityCouncil: only the colosseum contract can be a sender");
        securityCouncil.requestValidation(bytes32("dummy output root"), txValue, txData);
    }

    function test_requestValidation_succeeds() external {
        vm.startPrank(colosseum);
        // request output validation
        uint128 l2BlockNumber = 3;
        vm.expectEmit(true, false, false, true);
        uint256 transactionId = securityCouncil.generateTransactionId(colosseum, txValue, txData);
        emit ValidationRequested(transactionId, bytes32("dummy output root"), l2BlockNumber);
        securityCouncil.requestValidation(bytes32("dummy output root"), l2BlockNumber, txData);
        vm.stopPrank();
    }

    function test_requestValidation_execute_succeeds() external {
        vm.startPrank(colosseum);
        // request output validation
        uint256 l2BlockNumber = 3;
        vm.expectEmit(true, false, false, true);
        uint256 transactionId = securityCouncil.generateTransactionId(colosseum, txValue, txData);
        emit ValidationRequested(transactionId, bytes32("dummy output root"), l2BlockNumber);
        securityCouncil.requestValidation(bytes32("dummy output root"), l2BlockNumber, txData);

        // check transaction not executed
        Types.MultiSigTransaction memory t;
        (t.target, t.executed, t.value, t.data) = securityCouncil.transactions(transactionId);
        assertEq(t.executed, false);
        vm.stopPrank();

        // confirm transaction to execute
        vm.prank(guardian1);
        securityCouncil.confirmTransaction(transactionId);
        vm.prank(guardian2);
        securityCouncil.confirmTransaction(transactionId);

        // check transaction confirmed
        assertTrue(securityCouncil.isConfirmedBy(transactionId, guardian1));
        assertTrue(securityCouncil.isConfirmedBy(transactionId, guardian2));

        // check transaction executed
        (t.target, t.executed, t.value, t.data) = securityCouncil.transactions(transactionId);
        assertEq(t.executed, true);
    }

    function test_requestDeletion_succeeds() external {
        vm.startPrank(guardian1);
        // request output deletion
        uint256 outputIndex = 1;
        vm.expectEmit(true, true, false, true);
        bytes memory message = abi.encodeWithSelector(
            Colosseum.forceDeleteOutput.selector,
            outputIndex
        );
        uint256 transactionId = securityCouncil.generateTransactionId(securityCouncil.COLOSSEUM(), 0, message);
        emit DeletionRequested(transactionId, outputIndex);
        securityCouncil.requestDeletion(outputIndex, false);

        // check transaction not executed
        Types.MultiSigTransaction memory t;
        (t.target, t.executed, t.value, t.data) = securityCouncil.transactions(transactionId);
        assertEq(t.executed, false);
        vm.stopPrank();

        // confirm transaction to execute
        vm.prank(guardian2);
        securityCouncil.confirmTransaction(transactionId);

        // check transaction confirmed
        assertTrue(securityCouncil.isConfirmedBy(transactionId, guardian1));
        assertTrue(securityCouncil.isConfirmedBy(transactionId, guardian2));

        // check transaction executed
        (t.target, t.executed, t.value, t.data) = securityCouncil.transactions(transactionId);
        assertEq(t.executed, true);
    }

    function test_requestDeletion_alreadyRequested_reverts() external {
        vm.startPrank(guardian1);
        // request output deletion
        uint256 outputIndex = 1;
        vm.expectEmit(true, true, false, true);
        bytes memory message = abi.encodeWithSelector(
            Colosseum.forceDeleteOutput.selector,
            outputIndex
        );
        uint256 transactionId = securityCouncil.generateTransactionId(securityCouncil.COLOSSEUM(), 0, message);
        emit DeletionRequested(transactionId, outputIndex);
        securityCouncil.requestDeletion(outputIndex, false);

        // try to request the same output index
        vm.expectRevert("SecurityCouncil: the output has already been requested to be deleted");
        securityCouncil.requestDeletion(outputIndex, false);
        vm.stopPrank();
    }

    function test_requestDeletion_force_succeeds() external {
        vm.startPrank(guardian1);
        // request output deletion
        uint256 outputIndex = 1;
        vm.expectEmit(true, true, false, true);
        bytes memory message = abi.encodeWithSelector(
            Colosseum.forceDeleteOutput.selector,
            outputIndex
        );
        uint256 transactionId = securityCouncil.generateTransactionId(securityCouncil.COLOSSEUM(), 0, message);
        emit DeletionRequested(transactionId, outputIndex);
        securityCouncil.requestDeletion(outputIndex, false);

        _roll(1);

        // try to request the same output index
        vm.expectEmit(true, true, false, true);
        transactionId = securityCouncil.generateTransactionId(securityCouncil.COLOSSEUM(), 0, message);
        emit DeletionRequested(transactionId, outputIndex);
        securityCouncil.requestDeletion(outputIndex, true);
        vm.stopPrank();
    }

    function _roll(uint256 addBlockNumber) private {
        blockNumber += addBlockNumber;
        vm.roll(blockNumber);
    }
}
