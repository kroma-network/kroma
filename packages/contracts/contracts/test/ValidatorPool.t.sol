// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { Constants } from "../libraries/Constants.sol";
import { Types } from "../libraries/Types.sol";
import { KromaPortal } from "../L1/KromaPortal.sol";
import { L2OutputOracle } from "../L1/L2OutputOracle.sol";
import { ValidatorPool } from "../L1/ValidatorPool.sol";
import { ZKMerkleTrie } from "../L1/ZKMerkleTrie.sol";
import { ValidatorRewardVault } from "../L2/ValidatorRewardVault.sol";
import { Predeploys } from "../libraries/Predeploys.sol";
import { Proxy } from "../universal/Proxy.sol";
import { L2OutputOracle_Initializer } from "./CommonTest.t.sol";

contract MockL2OutputOracle is L2OutputOracle {
    constructor(ValidatorPool _validatorPool, address _colosseum)
        L2OutputOracle(_validatorPool, _colosseum, 1800, 2, 0, 0, 7 days)
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

    event BondIncreased(uint256 indexed outputIndex, address indexed challenger, uint128 amount);
    event PendingBondAdded(uint256 indexed outputIndex, address indexed challenger, uint128 amount);
    event PendingBondReleased(
        uint256 indexed outputIndex,
        address indexed challenger,
        address indexed recipient,
        uint128 amount
    );
    event Unbonded(uint256 indexed outputIndex, address indexed recipient, uint128 amount);

    function setUp() public override {
        super.setUp();

        finalizationPeriodSeconds = oracle.FINALIZATION_PERIOD_SECONDS();

        address oracleAddress = address(oracle);
        MockL2OutputOracle mockOracleImpl = new MockL2OutputOracle(pool, address(challenger));
        vm.prank(multisig);
        Proxy(payable(oracleAddress)).upgradeTo(address(mockOracleImpl));
        mockOracle = MockL2OutputOracle(oracleAddress);

        portal = pool.PORTAL();
    }

    function test_constructor_succeeds() external {
        assertEq(address(pool.L2_ORACLE()), address(oracle));
        assertEq(pool.TRUSTED_VALIDATOR(), trusted);
        assertEq(pool.REQUIRED_BOND_AMOUNT(), requiredBondAmount);
        assertEq(pool.MAX_UNBOND(), maxUnbond);
        assertEq(pool.ROUND_DURATION(), roundDuration);
    }

    function test_deposit_succeeds() public {
        uint256 trustedBalance = trusted.balance;

        vm.prank(trusted);
        pool.deposit{ value: requiredBondAmount }();
        assertEq(pool.balanceOf(trusted), requiredBondAmount);
        assertEq(trusted.balance, trustedBalance - requiredBondAmount);
        assertTrue(pool.isValidator(trusted));
        assertEq(pool.validatorCount(), 1);

        vm.prank(asserter);
        pool.deposit{ value: requiredBondAmount }();
        assertEq(pool.balanceOf(asserter), requiredBondAmount);
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
        pool.deposit{ value: requiredBondAmount }();
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
        uint256 depositAmount = requiredBondAmount * 2;

        vm.prank(trusted);
        pool.deposit{ value: depositAmount }();
        assertEq(pool.balanceOf(trusted), depositAmount);
        assertEq(trusted.balance, trustedBalance - depositAmount);
        assertTrue(pool.isValidator(trusted));
        assertEq(pool.validatorCount(), 1);

        trustedBalance = trusted.balance;
        uint256 withdrawalAmount = requiredBondAmount;

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
        uint256 nextBlockNumber = oracle.nextBlockNumber();
        bytes32 outputRoot = keccak256(abi.encode(nextBlockNumber));
        address validator = pool.nextValidator();

        warpToSubmitTime(nextBlockNumber);

        vm.prank(validator);
        mockOracle.addOutput(outputRoot, nextBlockNumber);

        uint128 expiresAt = uint128(block.timestamp + finalizationPeriodSeconds);
        vm.prank(address(oracle));
        vm.expectEmit(true, true, false, true, address(pool));
        emit Bonded(validator, nextOutputIndex, uint128(requiredBondAmount), expiresAt);
        pool.createBond(nextOutputIndex, expiresAt);
        assertEq(pool.balanceOf(validator), 0);
        assertFalse(pool.isValidator(validator));
        assertEq(pool.getBond(nextOutputIndex).amount, uint128(requiredBondAmount));
        assertEq(pool.getBond(nextOutputIndex).expiresAt, expiresAt);
    }

    function test_createBond_unbondBefore_succeeds() external {
        test_createBond_succeeds();

        Types.CheckpointOutput memory firstOutput = oracle.getL2Output(0);
        Types.Bond memory firstBond = pool.getBond(0);
        // warp to the expiration time of the first bond.
        vm.warp(firstBond.expiresAt);

        uint256 nextOutputIndex = oracle.nextOutputIndex();
        uint256 nextBlockNumber = oracle.nextBlockNumber();
        bytes32 outputRoot = keccak256(abi.encode(nextBlockNumber));
        address validator = pool.nextValidator();
        if (validator == Constants.VALIDATOR_PUBLIC_ROUND_ADDRESS) {
            validator = asserter;
        }

        // deposit again & append new output
        vm.startPrank(validator);
        mockOracle.addOutput(outputRoot, nextBlockNumber);
        pool.deposit{ value: requiredBondAmount }();
        vm.stopPrank();

        uint128 expiresAt = uint128(block.timestamp + finalizationPeriodSeconds);
        vm.prank(address(oracle));
        vm.expectEmit(true, true, false, true, address(pool));
        emit Unbonded(0, firstOutput.submitter, uint128(firstBond.amount));
        pool.createBond(nextOutputIndex, expiresAt);
        assertEq(pool.balanceOf(firstOutput.submitter), requiredBondAmount);

        // check whether bond is deleted
        vm.expectRevert("ValidatorPool: the bond does not exist");
        pool.getBond(0);
    }

    function test_createBond_senderNotL2OO_reverts() external {
        test_deposit_succeeds();

        vm.prank(trusted);
        vm.expectRevert("ValidatorPool: sender is not L2OutputOracle");
        pool.createBond(0, uint128(block.timestamp + finalizationPeriodSeconds));
    }

    function test_createBond_existsBond_reverts() external {
        test_createBond_succeeds();

        uint256 outputIndex = oracle.latestOutputIndex();
        Types.Bond memory bond = pool.getBond(outputIndex);
        assertTrue(bond.expiresAt > 0);

        Types.CheckpointOutput memory output = oracle.getL2Output(outputIndex);

        vm.prank(output.submitter);
        pool.deposit{ value: requiredBondAmount }();

        vm.prank(address(oracle));
        vm.expectRevert("ValidatorPool: bond of the given output index already exists");
        pool.createBond(outputIndex, uint128(block.timestamp + finalizationPeriodSeconds));
    }

    function test_createBond_insufficientBalances_reverts() external {
        uint256 nextOutputIndex = oracle.nextOutputIndex();
        uint256 nextBlockNumber = oracle.nextBlockNumber();
        bytes32 outputRoot = keccak256(abi.encode(nextBlockNumber));
        address validator = pool.nextValidator();

        warpToSubmitTime(nextBlockNumber);

        vm.prank(validator);
        mockOracle.addOutput(outputRoot, nextBlockNumber);

        uint128 expiresAt = uint128(block.timestamp + finalizationPeriodSeconds);
        vm.prank(address(oracle));
        vm.expectRevert("ValidatorPool: insufficient balances");
        pool.createBond(nextOutputIndex, expiresAt);
    }

    function test_unbond_succeeds() public {
        test_createBond_succeeds();

        uint256 outputIndex = oracle.latestOutputIndex();
        Types.CheckpointOutput memory output = oracle.getL2Output(outputIndex);
        Types.Bond memory bond = pool.getBond(outputIndex);

        // warp to the time the output is finalized and the bond is expires.
        vm.warp(bond.expiresAt);

        vm.expectEmit(true, true, false, true, address(pool));
        emit Unbonded(outputIndex, output.submitter, bond.amount);
        vm.expectCall(
            address(pool.PORTAL()),
            abi.encodeWithSelector(
                KromaPortal.depositTransactionByValidatorPool.selector,
                Predeploys.VALIDATOR_REWARD_VAULT,
                pool.VAULT_REWARD_GAS_LIMIT(),
                abi.encodeWithSelector(
                    ValidatorRewardVault.reward.selector,
                    output.submitter,
                    output.l2BlockNumber
                )
            )
        );
        vm.prank(trusted);
        pool.unbond();
        assertEq(pool.balanceOf(output.submitter), uint256(bond.amount));
    }

    function test_unbond_multipleBonds_succeeds() public {
        uint256 tries = 2;
        uint256 deposit = requiredBondAmount * tries;
        vm.prank(trusted);
        pool.deposit{ value: deposit }();

        // submit 2 outputs, only trusted can submit outputs before at least one unbond.
        uint256 blockNumber = 0;
        uint128 expiresAt = 0;
        for (uint256 i = 0; i < tries; i++) {
            blockNumber = oracle.nextBlockNumber();
            warpToSubmitTime(blockNumber);
            expiresAt = uint128(block.timestamp + finalizationPeriodSeconds);
            assertEq(pool.nextValidator(), trusted);
            vm.prank(trusted);
            mockOracle.addOutput(keccak256(abi.encode(blockNumber)), blockNumber);
            vm.prank(address(oracle));
            pool.createBond(i, expiresAt);
            assertEq(pool.balanceOf(trusted), deposit - requiredBondAmount * (i + 1));
        }

        uint256 firstOutputIndex = 0;
        Types.CheckpointOutput memory firstOutput = oracle.getL2Output(firstOutputIndex);
        Types.Bond memory firstBond = pool.getBond(firstOutputIndex);

        uint256 secondOutputIndex = 1;
        Types.CheckpointOutput memory secondOutput = oracle.getL2Output(secondOutputIndex);
        Types.Bond memory secondBond = pool.getBond(secondOutputIndex);

        // warp to the time the second output is finalized and the two bonds are expired.
        vm.warp(secondBond.expiresAt);

        vm.expectEmit(true, true, false, true, address(pool));
        emit Unbonded(firstOutputIndex, firstOutput.submitter, firstBond.amount);
        vm.expectCall(
            address(pool.PORTAL()),
            abi.encodeWithSelector(
                KromaPortal.depositTransactionByValidatorPool.selector,
                Predeploys.VALIDATOR_REWARD_VAULT,
                pool.VAULT_REWARD_GAS_LIMIT(),
                abi.encodeWithSelector(
                    ValidatorRewardVault.reward.selector,
                    firstOutput.submitter,
                    firstOutput.l2BlockNumber
                )
            )
        );
        vm.expectEmit(true, true, false, true, address(pool));
        emit Unbonded(secondOutputIndex, secondOutput.submitter, secondBond.amount);
        vm.expectCall(
            address(pool.PORTAL()),
            abi.encodeWithSelector(
                KromaPortal.depositTransactionByValidatorPool.selector,
                Predeploys.VALIDATOR_REWARD_VAULT,
                pool.VAULT_REWARD_GAS_LIMIT(),
                abi.encodeWithSelector(
                    ValidatorRewardVault.reward.selector,
                    secondOutput.submitter,
                    secondOutput.l2BlockNumber
                )
            )
        );
        vm.prank(trusted);
        pool.unbond();

        // check whether bonds are deleted and trusted balance has increased.
        for (uint256 i = 0; i < tries; i++) {
            vm.expectRevert("ValidatorPool: the bond does not exist");
            pool.getBond(i);
        }
        assertEq(pool.balanceOf(trusted), deposit);
    }

    function test_unbond_maxUnbond_succeeds() public {
        uint256 tries = maxUnbond + 1;
        uint256 deposit = requiredBondAmount * tries;
        vm.prank(trusted);
        pool.deposit{ value: deposit }();

        // submit (maxUnbond + 1) outputs, only trusted can submit outputs before at least one unbond.
        uint256 blockNumber = 0;
        uint128 expiresAt = 0;
        for (uint256 i = 0; i < tries; i++) {
            blockNumber = oracle.nextBlockNumber();
            warpToSubmitTime(blockNumber);
            expiresAt = uint128(block.timestamp + finalizationPeriodSeconds);
            assertEq(pool.nextValidator(), trusted);
            vm.prank(trusted);
            mockOracle.addOutput(keccak256(abi.encode(blockNumber)), blockNumber);
            vm.prank(address(oracle));
            pool.createBond(i, expiresAt);
            assertEq(pool.balanceOf(trusted), deposit - requiredBondAmount * (i + 1));
        }

        uint256 outputIndex = oracle.latestOutputIndex();
        Types.Bond memory bond = pool.getBond(outputIndex);

        // warp to the time the latest output is finalized and all bonds are expired.
        vm.warp(bond.expiresAt);

        vm.prank(trusted);
        pool.unbond();

        // check whether maxUnbond number of bonds are deleted and the last one is not.
        for (uint256 i = 0; i < tries - 1; i++) {
            vm.expectRevert("ValidatorPool: the bond does not exist");
            pool.getBond(i);
        }
        bond = pool.getBond(tries - 1);
        assertEq(bond.amount, requiredBondAmount);
    }

    function test_unbond_notExpired_reverts() external {
        test_createBond_succeeds();

        vm.expectRevert("ValidatorPool: no bond that can be unbond");
        pool.unbond();
    }

    function test_unbond_noBond_reverts() external {
        vm.expectRevert("ValidatorPool: no bond that can be unbond");
        pool.unbond();
    }

    function test_addPendingBond_succeeds() public {
        test_createBond_succeeds();

        uint256 outputIndex = oracle.latestOutputIndex();
        Types.Bond memory bond = pool.getBond(outputIndex);

        vm.prank(challenger);
        pool.deposit{ value: bond.amount }();

        vm.prank(oracle.COLOSSEUM());
        vm.expectEmit(true, true, false, true, address(pool));
        emit PendingBondAdded(outputIndex, challenger, bond.amount);
        pool.addPendingBond(outputIndex, challenger);

        // check bond state
        assertEq(pool.getPendingBond(outputIndex, challenger), bond.amount);
        assertEq(pool.balanceOf(challenger), 0);
    }

    function test_addPendingBond_noBond_reverts() external {
        vm.prank(oracle.COLOSSEUM());
        vm.expectRevert("ValidatorPool: the output is already finalized");
        pool.addPendingBond(0, challenger);
    }

    function test_addPendingBond_insufficientBalances_reverts() external {
        test_createBond_succeeds();

        uint256 outputIndex = oracle.latestOutputIndex();

        vm.prank(oracle.COLOSSEUM());
        vm.expectRevert("ValidatorPool: insufficient balances");
        pool.addPendingBond(outputIndex, challenger);
    }

    function test_increaseBond_succeeds() public {
        test_addPendingBond_succeeds();

        uint256 prevScBalance = pool.balanceOf(pool.SECURITY_COUNCIL());

        uint256 outputIndex = oracle.latestOutputIndex();
        Types.Bond memory prevBond = pool.getBond(outputIndex);
        uint128 pendingBond = pool.getPendingBond(outputIndex, challenger);
        uint128 tax = pendingBond * 20 / 100; // 20% tax
        uint128 increased = pendingBond - tax;

        vm.prank(oracle.COLOSSEUM());
        vm.expectEmit(true, true, false, false);
        emit BondIncreased(outputIndex, challenger, increased);
        pool.increaseBond(outputIndex, challenger);

        // check bond state
        assertEq(pool.getBond(outputIndex).amount, prevBond.amount + increased);
        assertEq(pool.balanceOf(challenger), 0);

        // check security council balance
        assertEq(pool.balanceOf(pool.SECURITY_COUNCIL()), prevScBalance + tax);
    }

    function test_increaseBond_noBond_reverts() external {
        vm.prank(oracle.COLOSSEUM());
        vm.expectRevert("ValidatorPool: the output is already finalized");
        pool.increaseBond(0, challenger);
    }

    function test_increaseBond_noPendingBond_reverts() external {
        test_createBond_succeeds();

        vm.prank(oracle.COLOSSEUM());
        vm.expectRevert("ValidatorPool: the pending bond does not exist");
        pool.increaseBond(0, challenger);
    }

    function test_releasePendingBond_succeeds() external {
        test_addPendingBond_succeeds();

        uint256 outputIndex = oracle.latestOutputIndex();
        uint128 pendingBond = pool.getPendingBond(outputIndex, challenger);

        vm.prank(oracle.COLOSSEUM());
        vm.expectEmit(true, true, false, true, address(pool));
        emit PendingBondReleased(outputIndex, challenger, challenger, pendingBond);
        pool.releasePendingBond(outputIndex, challenger, challenger);

        assertEq(pool.balanceOf(challenger), pendingBond);

        vm.expectRevert("ValidatorPool: the pending bond does not exist");
        pool.getPendingBond(outputIndex, challenger);
    }

    function test_releasePendingBond_noPendingBond_succeeds() external {
        vm.prank(oracle.COLOSSEUM());
        vm.expectRevert("ValidatorPool: the pending bond does not exist");
        pool.releasePendingBond(0, challenger, challenger);
    }

    function test_getBond_succeeds() external {
        test_createBond_succeeds();

        uint256 outputIndex = oracle.latestOutputIndex();
        Types.Bond memory bond = pool.getBond(outputIndex);

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
        pool.deposit{ value: requiredBondAmount }();
        vm.prank(asserter);
        pool.deposit{ value: requiredBondAmount - 1 }();

        assertTrue(pool.isValidator(trusted));
        assertFalse(pool.isValidator(asserter));
        assertFalse(pool.isValidator(challenger));
    }

    function test_validatorCount_succeeds() external {
        vm.prank(trusted);
        pool.deposit{ value: requiredBondAmount }();
        assertEq(pool.validatorCount(), 1);

        vm.prank(asserter);
        pool.deposit{ value: requiredBondAmount }();
        assertEq(pool.validatorCount(), 2);

        vm.prank(challenger);
        pool.deposit{ value: requiredBondAmount - 1 }();
        assertEq(pool.validatorCount(), 2);
    }

    function test_nextValidator_succeeds() external {
        // deposit funds
        vm.prank(trusted);
        pool.deposit{ value: requiredBondAmount * 10 }();
        vm.prank(asserter);
        pool.deposit{ value: requiredBondAmount * 10 }();

        address prev = pool.nextValidator();
        assertEq(prev, trusted);

        uint256 tries = 10;
        uint256 outputIndex = 0;
        uint256 blockNumber = 0;
        uint128 expiresAt = 0;

        // submit 10 outputs
        for (uint256 i = 0; i < tries; i++) {
            outputIndex = oracle.nextOutputIndex();
            blockNumber = oracle.nextBlockNumber();
            warpToSubmitTime(blockNumber);
            expiresAt = uint128(block.timestamp + finalizationPeriodSeconds);
            vm.prank(pool.nextValidator());
            mockOracle.addOutput(keccak256(abi.encode(blockNumber)), blockNumber);
            vm.prank(address(oracle));
            pool.createBond(outputIndex, expiresAt);
        }

        // warp to first finalization time and submit new output
        warpToSubmitTime(oracle.nextBlockNumber());
        outputIndex = oracle.nextOutputIndex();
        blockNumber = (expiresAt / oracle.L2_BLOCK_TIME()) - 1;
        vm.warp(expiresAt);
        expiresAt = uint128(block.timestamp + finalizationPeriodSeconds);
        vm.prank(pool.nextValidator());
        mockOracle.addOutput(keccak256(abi.encode(blockNumber)), blockNumber);
        vm.prank(address(oracle));
        pool.createBond(outputIndex, expiresAt);

        bool changed = false;
        for (uint256 i = 0; i < tries - 1; i++) {
            // check the next validator has changed
            if (pool.nextValidator() != prev) {
                changed = true;
                break;
            }

            prev = pool.nextValidator();
            // submit next output and finalize prev output
            outputIndex = oracle.nextOutputIndex();
            blockNumber = oracle.nextBlockNumber();
            warpToSubmitTime(blockNumber);
            expiresAt = uint128(block.timestamp + finalizationPeriodSeconds);
            vm.prank(pool.nextValidator());
            mockOracle.addOutput(keccak256(abi.encode(blockNumber)), blockNumber);
            vm.prank(address(oracle));
            pool.createBond(outputIndex, expiresAt);
        }

        assertTrue(changed, "the next validator has not changed");

        // warp to public round
        uint256 l2Timestamp = oracle.computeL2Timestamp(oracle.nextBlockNumber() + 1);
        vm.warp(l2Timestamp + roundDuration + 1);
        assertEq(pool.nextValidator(), Constants.VALIDATOR_PUBLIC_ROUND_ADDRESS);
    }

    function test_securityCouncilCannotBeValidator_succeeds() external {
        address sc = pool.SECURITY_COUNCIL();
        uint256 depositAmount = pool.REQUIRED_BOND_AMOUNT() * 100;
        vm.deal(sc, depositAmount + 1 ether);

        vm.prank(sc);
        pool.deposit{value: depositAmount}();
        assertEq(pool.balanceOf(sc), depositAmount);
        assertFalse(pool.isValidator(sc));
    }
}
