// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { stdError } from "forge-std/Test.sol";

import { Types } from "../libraries/Types.sol";
import { KromaPortal } from "../L1/KromaPortal.sol";
import { L2OutputOracle } from "../L1/L2OutputOracle.sol";
import { ValidatorPool } from "../L1/ValidatorPool.sol";
import { ZKMerkleTrie } from "../L1/ZKMerkleTrie.sol";
import { Proxy } from "../universal/Proxy.sol";
import { L2OutputOracle_Initializer } from "./CommonTest.t.sol";

contract MockL2OutputOracle is L2OutputOracle {
    constructor(ValidatorPool _validatorPool, address _colosseum)
        L2OutputOracle(1800, 2, 0, 0, _validatorPool, _colosseum, 7 days)
    {}

    function addOutput(bytes32 _outputRoot, uint256 _l2BlockNumber) external payable {
        l2Outputs.push(
            Types.CheckpointOutput({
                submitter: msg.sender,
                outputRoot: _outputRoot,
                timestamp: uint128(block.timestamp),
                l2BlockNumber: uint128(_l2BlockNumber)
            })
        );
    }

    function replaceOutput(uint256 _outputIndex, bytes32 _outputRoot) external {
        l2Outputs[_outputIndex] = Types.CheckpointOutput({
            submitter: msg.sender,
            outputRoot: _outputRoot,
            timestamp: uint128(block.timestamp),
            l2BlockNumber: l2Outputs[_outputIndex].l2BlockNumber
        });
    }
}

// Test the implementations of the ValidatorPool
contract ValidatorPoolTest is L2OutputOracle_Initializer {
    MockL2OutputOracle mockOracle;
    KromaPortal portal;

    uint256 internal finalizationPeriodSeconds;

    event Bonded(
        address indexed submitter,
        uint256 indexed outputIndex,
        uint128 amount,
        uint128 expiresAt
    );

    event BondIncreased(address indexed challenger, uint256 indexed outputIndex, uint128 amount);

    event Unbonded(uint256 indexed outputIndex, address indexed recipient, uint128 amount);

    function setUp() public override {
        super.setUp();

        finalizationPeriodSeconds = oracle.FINALIZATION_PERIOD_SECONDS();

        address oracleAddress = address(oracle);
        MockL2OutputOracle mockOracleImpl = new MockL2OutputOracle(pool, address(challenger));
        vm.prank(multisig);
        Proxy(payable(oracleAddress)).upgradeTo(address(mockOracleImpl));
        mockOracle = MockL2OutputOracle(oracleAddress);

        portal = new KromaPortal({
            _l2Oracle: oracle,
            _guardian: guardian,
            _paused: false,
            _config: systemConfig,
            _zkMerkleTrie: ZKMerkleTrie(address(0))
        });
    }

    function test_constructor_succeeds() external {
        assertEq(address(pool.L2_ORACLE()), address(oracle));
        assertEq(pool.TRUSTED_VALIDATOR(), trusted);
        assertEq(pool.MIN_BOND_AMOUNT(), minBond);
    }

    function test_deposit_succeeds() public {
        uint256 trustedBalance = trusted.balance;

        vm.prank(trusted);
        pool.deposit{ value: minBond }();
        assertEq(pool.balanceOf(trusted), minBond);
        assertEq(trusted.balance, trustedBalance - minBond);
        assertTrue(pool.isValidator(trusted));
        assertEq(pool.validatorCount(), 1);

        vm.prank(asserter);
        pool.deposit{ value: minBond }();
        assertEq(pool.balanceOf(asserter), minBond);
        assertTrue(pool.isValidator(asserter));
        assertEq(pool.validatorCount(), 2);
    }

    function test_deposit_alreadyValidator_succeeds() external {
        test_deposit_succeeds();

        uint256 count = pool.validatorCount();
        address nextValidator = pool.nextValidator();
        uint256 deposits = pool.balanceOf(nextValidator);

        uint256 prevBalance = nextValidator.balance;
        uint256 depositAmount = 1;

        vm.prank(nextValidator);
        pool.deposit{ value: depositAmount }();
        assertEq(pool.balanceOf(trusted), deposits + depositAmount);
        assertEq(nextValidator.balance, prevBalance - depositAmount);
        assertTrue(pool.isValidator(trusted));
        assertEq(pool.validatorCount(), count);
    }

    function test_deposit_insufficientBalances_reverts() external {
        vm.deal(asserter, 0);
        vm.prank(asserter);
        vm.expectRevert();
        pool.deposit{ value: minBond }();
    }

    function test_withdraw_loseValidatorEligibility_succeeds() external {
        test_deposit_succeeds();

        uint256 count = pool.validatorCount();
        address nextValidator = pool.nextValidator();
        uint256 deposits = pool.balanceOf(nextValidator);

        uint256 prevBalance = nextValidator.balance;
        uint256 withdrawalAmount = 1;

        vm.prank(nextValidator);
        pool.withdraw(withdrawalAmount);
        assertEq(pool.balanceOf(nextValidator), deposits - withdrawalAmount);
        assertEq(nextValidator.balance, prevBalance + withdrawalAmount);
        assertFalse(pool.isValidator(nextValidator));
        assertEq(pool.validatorCount(), count - 1);
    }

    function test_withdraw_all_succeeds() external {
        test_deposit_succeeds();

        uint256 count = pool.validatorCount();
        address nextValidator = pool.nextValidator();
        uint256 deposits = pool.balanceOf(nextValidator);

        uint256 prevBalance = nextValidator.balance;

        vm.prank(nextValidator);
        pool.withdraw(deposits);
        assertEq(pool.balanceOf(nextValidator), 0);
        assertEq(nextValidator.balance, prevBalance + deposits);
        assertFalse(pool.isValidator(nextValidator));
        assertEq(pool.validatorCount(), count - 1);
    }

    function test_withdraw_maintainValidatorEligibility_succeeds() external {
        uint256 trustedBalance = trusted.balance;
        uint256 depositAmount = minBond * 2;

        vm.prank(trusted);
        pool.deposit{ value: depositAmount }();
        assertEq(pool.balanceOf(trusted), depositAmount);
        assertEq(trusted.balance, trustedBalance - depositAmount);
        assertTrue(pool.isValidator(trusted));
        assertEq(pool.validatorCount(), 1);

        trustedBalance = trusted.balance;
        uint256 withdrawalAmount = minBond;

        vm.prank(trusted);
        pool.withdraw(withdrawalAmount);
        assertEq(pool.balanceOf(trusted), withdrawalAmount);
        assertEq(trusted.balance, trustedBalance + withdrawalAmount);
        assertTrue(pool.isValidator(trusted));
        assertEq(pool.validatorCount(), 1);
    }

    function test_createBond_succeeds() public {
        test_deposit_succeeds();

        uint256 nextOutputIndex = oracle.nextOutputIndex();
        uint256 nextBlockNumber = (nextOutputIndex + 1) * submissionInterval;
        bytes32 outputRoot = keccak256(abi.encode(nextBlockNumber));
        address validator = pool.nextValidator();
        vm.prank(validator);
        mockOracle.addOutput(outputRoot, nextBlockNumber);

        uint128 expiresAt = uint128(block.timestamp + finalizationPeriodSeconds);
        vm.prank(address(oracle));
        vm.expectEmit(true, true, false, false);
        emit Bonded(validator, nextOutputIndex, uint128(minBond), expiresAt);
        pool.createBond(nextOutputIndex, uint128(minBond), expiresAt);
        assertEq(pool.balanceOf(validator), 0);
        assertFalse(pool.isValidator(validator));
        assertEq(pool.getBond(nextOutputIndex).expiresAt, expiresAt);
    }

    function test_createBond_unbondBefore_succeeds() external {
        test_createBond_succeeds();

        Types.CheckpointOutput memory firstOutput = oracle.getL2Output(0);
        Types.Bond memory firstBond = pool.getBond(0);
        // warp to the expiration time of the first bond.
        vm.warp(firstBond.expiresAt);

        uint256 nextOutputIndex = oracle.nextOutputIndex();
        uint256 nextBlockNumber = (nextOutputIndex + 1) * submissionInterval;
        bytes32 outputRoot = keccak256(abi.encode(nextBlockNumber));
        address validator = pool.nextValidator();

        // deposit again & append new output
        vm.startPrank(validator);
        mockOracle.addOutput(outputRoot, nextBlockNumber);
        pool.deposit{ value: minBond }();
        vm.stopPrank();

        uint128 expiresAt = uint128(block.timestamp + finalizationPeriodSeconds);
        vm.prank(address(oracle));
        vm.expectEmit(true, true, false, false);
        emit Unbonded(0, firstOutput.submitter, uint128(firstBond.amount));
        pool.createBond(nextOutputIndex, uint128(minBond), expiresAt);
        assertEq(pool.balanceOf(firstOutput.submitter), minBond);

        // check whether bond is deleted
        vm.expectRevert("ValidatorPool: the bond does not exist");
        pool.getBond(0);
    }

    function test_createBond_senderNotL2OO_reverts() external {
        test_deposit_succeeds();

        vm.prank(trusted);
        vm.expectRevert("ValidatorPool: sender is not L2OutputOracle");
        pool.createBond(0, uint128(minBond), uint128(block.timestamp + finalizationPeriodSeconds));
    }

    function test_createBond_zeroAmount_reverts() external {
        test_deposit_succeeds();

        vm.prank(address(oracle));
        vm.expectRevert("ValidatorPool: the bond amount is too small");
        pool.createBond(0, 0, uint128(block.timestamp + finalizationPeriodSeconds));
    }

    function test_createBond_existsBond_reverts() external {
        test_createBond_succeeds();

        uint256 latestOutputIndex = oracle.latestOutputIndex();
        Types.Bond memory bond = pool.getBond(latestOutputIndex);
        assertTrue(bond.expiresAt > 0);

        Types.CheckpointOutput memory output = oracle.getL2Output(latestOutputIndex);

        vm.prank(output.submitter);
        pool.deposit{ value: minBond }();

        vm.prank(address(oracle));
        vm.expectRevert("ValidatorPool: bond of the given output index already exists");
        pool.createBond(
            latestOutputIndex,
            uint128(minBond),
            uint128(block.timestamp + finalizationPeriodSeconds)
        );
    }

    function test_createBond_insufficientBalances_reverts() external {
        uint256 nextOutputIndex = oracle.nextOutputIndex();
        uint256 nextBlockNumber = (nextOutputIndex + 1) * submissionInterval;
        bytes32 outputRoot = keccak256(abi.encode(nextBlockNumber));
        address validator = pool.nextValidator();
        vm.prank(validator);
        mockOracle.addOutput(outputRoot, nextBlockNumber);

        uint128 expiresAt = uint128(block.timestamp + finalizationPeriodSeconds);
        vm.prank(address(oracle));
        vm.expectRevert("ValidatorPool: insufficient balances");
        pool.createBond(nextOutputIndex, uint128(minBond), expiresAt);
    }

    function test_unbond_succeeds() public {
        test_createBond_succeeds();

        uint256 latestOutputIndex = oracle.latestOutputIndex();
        Types.CheckpointOutput memory output = oracle.getL2Output(latestOutputIndex);
        Types.Bond memory bond = pool.getBond(latestOutputIndex);
        assertTrue(bond.expiresAt > 0);

        // warp to the expiration time
        vm.warp(bond.expiresAt);

        vm.prank(trusted);
        vm.expectEmit(true, true, false, false);
        emit Unbonded(latestOutputIndex, output.submitter, uint128(bond.amount));
        pool.unbond();
        assertEq(pool.balanceOf(output.submitter), minBond);
    }

    function test_unbond_notExpired_reverts() external {
        test_createBond_succeeds();

        vm.expectRevert("ValidatorPool: no bond that can be unbond");
        pool.unbond();
    }

    function test_unbond_noBond_reverts() external {
        vm.expectRevert(stdError.indexOOBError);
        pool.unbond();
    }

    function test_increaseBond_succeeds() public {
        test_createBond_succeeds();

        uint256 latestOutputIndex = oracle.latestOutputIndex();
        Types.Bond memory prevBond = pool.getBond(latestOutputIndex);

        vm.prank(challenger);
        pool.deposit{ value: prevBond.amount }();

        vm.prank(oracle.COLOSSEUM());
        vm.expectEmit(true, true, false, false);
        emit BondIncreased(challenger, latestOutputIndex, prevBond.amount);
        pool.increaseBond(challenger, latestOutputIndex);

        // check bond state
        assertEq(pool.getBond(latestOutputIndex).amount, prevBond.amount * 2);
        assertEq(pool.balanceOf(challenger), 0);
    }

    function test_increaseBond_noBond_reverts() external {
        vm.prank(oracle.COLOSSEUM());
        vm.expectRevert("ValidatorPool: the bond does not exist");
        pool.increaseBond(challenger, 0);
    }

    function test_increaseBond_insufficientBalances_reverts() external {
        test_createBond_succeeds();

        uint256 latestOutputIndex = oracle.latestOutputIndex();

        vm.prank(oracle.COLOSSEUM());
        vm.expectRevert("ValidatorPool: insufficient balances");
        pool.increaseBond(challenger, latestOutputIndex);
    }

    function test_getBond_succeeds() external {
        test_createBond_succeeds();

        uint256 latestOutputIndex = oracle.latestOutputIndex();
        Types.Bond memory bond = pool.getBond(latestOutputIndex);

        assertTrue(bond.amount > 0);
        assertTrue(bond.expiresAt > block.timestamp);
    }

    function test_getBond_noBond_reverts() external {
        vm.expectRevert("ValidatorPool: the bond does not exist");
        pool.getBond(0);
    }

    function test_balanceOf_succeeds() external {
        vm.prank(trusted);
        pool.deposit{ value: 1 }();

        assertEq(pool.balanceOf(trusted), 1);
        assertEq(pool.balanceOf(asserter), 0);
        assertEq(pool.balanceOf(challenger), 0);
    }

    function test_isValidator_succeeds() external {
        vm.prank(trusted);
        pool.deposit{ value: minBond }();
        vm.prank(asserter);
        pool.deposit{ value: minBond - 1 }();

        assertTrue(pool.isValidator(trusted));
        assertFalse(pool.isValidator(asserter));
        assertFalse(pool.isValidator(challenger));
    }

    function test_validatorCount_succeeds() external {
        vm.prank(trusted);
        pool.deposit{ value: minBond }();
        assertEq(pool.validatorCount(), 1);

        vm.prank(asserter);
        pool.deposit{ value: minBond }();
        assertEq(pool.validatorCount(), 2);

        vm.prank(challenger);
        pool.deposit{ value: minBond - 1 }();
        assertEq(pool.validatorCount(), 2);
    }

    function test_nextValidator_succeeds() external {
        address prev = pool.nextValidator();
        assertEq(prev, trusted);

        test_deposit_succeeds();

        mockOracle.addOutput(keccak256(abi.encode(1)), 1);

        uint256 diffCount = 0;
        for (uint256 i = 0; i < 10; i++) {
            if (pool.nextValidator() != prev) {
                diffCount++;
            }
            prev = pool.nextValidator();
        }
        assertGt(diffCount, 0);
    }
}
