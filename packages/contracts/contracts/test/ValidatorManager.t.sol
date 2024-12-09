// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { Constants } from "../libraries/Constants.sol";
import { Types } from "../libraries/Types.sol";
import { Proxy } from "../universal/Proxy.sol";
import { IAssetManager } from "../L1/interfaces/IAssetManager.sol";
import { IValidatorManager } from "../L1/interfaces/IValidatorManager.sol";
import { L2OutputOracle } from "../L1/L2OutputOracle.sol";
import { ValidatorManager } from "../L1/ValidatorManager.sol";
import { ValidatorPool } from "../L1/ValidatorPool.sol";
import { ValidatorSystemUpgrade_Initializer } from "./CommonTest.t.sol";

contract MockL2OutputOracle is L2OutputOracle {
    constructor(
        ValidatorPool _validatorPool,
        ValidatorManager _validatorManager,
        address _colosseum,
        uint256 _submissionInterval,
        uint256 _l2BlockTime,
        uint256 _startingBlockNumber,
        uint256 _startingTimestamp,
        uint256 _finalizationPeriodSeconds
    )
        L2OutputOracle(
            _validatorPool,
            _validatorManager,
            _colosseum,
            _submissionInterval,
            _l2BlockTime,
            _startingBlockNumber,
            _startingTimestamp,
            _finalizationPeriodSeconds
        )
    {}

    function addOutput(uint256 l2BlockNumber) external {
        l2Outputs.push(
            Types.CheckpointOutput({
                submitter: msg.sender,
                outputRoot: keccak256(abi.encode(l2BlockNumber)),
                timestamp: uint128(block.timestamp),
                l2BlockNumber: uint128(l2BlockNumber)
            })
        );
    }

    function replaceOutput(address validator, uint256 outputIndex) external {
        l2Outputs[outputIndex].submitter = validator;
        l2Outputs[outputIndex].outputRoot = bytes32(0);
    }

    function mockSetNextFinalizeOutputIndex(uint256 l2OutputIndex) external {
        nextFinalizeOutputIndex = l2OutputIndex;
    }
}

contract MockValidatorManager is ValidatorManager {
    constructor(ConstructorParams memory _constructorParams) ValidatorManager(_constructorParams) {}

    function updatePriorityValidator(address validator) external {
        _nextPriorityValidator = validator;
    }

    function setJailExpiresAt(address validator, uint128 expiresAt) external {
        _jail[validator] = expiresAt;
    }

    function nextPriorityValidator() external view returns (address) {
        return _nextPriorityValidator;
    }

    function getBoostedReward(address validator) external view returns (uint128) {
        return _getBoostedReward(validator);
    }
}

contract ValidatorManagerTest is ValidatorSystemUpgrade_Initializer {
    MockL2OutputOracle mockOracle;
    MockValidatorManager mockValMgr;

    event ValidatorRegistered(
        address indexed validator,
        bool activated,
        uint8 commissionRate,
        uint128 assets
    );

    event ValidatorActivated(address indexed validator, uint256 activatedAt);

    event ValidatorStopped(address indexed validator, uint256 stopsAt);

    event ValidatorCommissionChangeInitiated(
        address indexed validator,
        uint8 oldCommissionRate,
        uint8 newCommissionRate
    );

    event ValidatorCommissionChangeFinalized(
        address indexed validator,
        uint8 oldCommissionRate,
        uint8 newCommissionRate
    );

    event ValidatorJailed(address indexed validator, uint128 expiresAt);

    event ValidatorUnjailed(address indexed validator);

    event RewardDistributed(
        uint256 indexed outputIndex,
        address indexed validator,
        uint128 validatorReward,
        uint128 baseReward,
        uint128 boostedReward
    );

    event Slashed(uint256 indexed outputIndex, address indexed loser, uint128 amount);

    event SlashReverted(uint256 indexed outputIndex, address indexed loser, uint128 amount);

    function _setUpKghDelegation(
        address validator,
        uint256 startingTokenId,
        uint128 kghCounts
    ) private {
        uint256[] memory tokenIds = new uint256[](kghCounts);
        for (uint256 i = startingTokenId; i < startingTokenId + kghCounts; i++) {
            kgh.mint(validator, i);
            vm.prank(validator);
            kgh.approve(address(assetMgr), i);
            tokenIds[i - startingTokenId] = i;
        }
        vm.prank(validator);
        assetMgr.delegateKghBatch(validator, tokenIds);
    }

    function _withdraw(address validator, uint128 amount) private {
        vm.warp(assetMgr.canWithdrawAt(validator) + 1);
        vm.prank(withdrawAcc);
        assetMgr.withdraw(validator, amount);
    }

    function setUp() public override {
        super.setUp();

        address oracleAddress = address(oracle);
        MockL2OutputOracle mockOracleImpl = new MockL2OutputOracle(
            pool,
            valMgr,
            address(colosseum),
            submissionInterval,
            l2BlockTime,
            startingBlockNumber,
            startingTimestamp,
            finalizationPeriodSeconds
        );
        vm.prank(multisig);
        Proxy(payable(oracleAddress)).upgradeTo(address(mockOracleImpl));
        mockOracle = MockL2OutputOracle(oracleAddress);

        address valMgrAddress = address(valMgr);
        MockValidatorManager mockValMgrImpl = new MockValidatorManager(constructorParams);
        vm.prank(multisig);
        Proxy(payable(valMgrAddress)).upgradeTo(address(mockValMgrImpl));
        mockValMgr = MockValidatorManager(valMgrAddress);

        // Submit until terminateOutputIndex and set next output index to be finalized after it
        vm.prank(trusted);
        pool.deposit{ value: trusted.balance }();
        for (uint256 i = oracle.nextOutputIndex(); i <= terminateOutputIndex; i++) {
            _submitL2OutputV1();
        }
        vm.warp(oracle.finalizedAt(terminateOutputIndex));
        mockOracle.mockSetNextFinalizeOutputIndex(terminateOutputIndex + 1);
    }

    function test_constructor_succeeds() external {
        assertEq(address(valMgr.L2_ORACLE()), address(oracle));
        assertEq(address(valMgr.ASSET_MANAGER()), address(assetMgr));
        assertEq(valMgr.TRUSTED_VALIDATOR(), trusted);
        assertEq(valMgr.MIN_REGISTER_AMOUNT(), minRegisterAmount);
        assertEq(valMgr.MIN_ACTIVATE_AMOUNT(), minActivateAmount);
        assertEq(valMgr.COMMISSION_CHANGE_DELAY_SECONDS(), commissionChangeDelaySeconds);
        assertEq(valMgr.ROUND_DURATION_SECONDS(), roundDuration);
        assertEq(valMgr.SOFT_JAIL_PERIOD_SECONDS(), softJailPeriodSeconds);
        assertEq(valMgr.HARD_JAIL_PERIOD_SECONDS(), hardJailPeriodSeconds);
        assertEq(valMgr.JAIL_THRESHOLD(), jailThreshold);
        assertEq(valMgr.MAX_OUTPUT_FINALIZATIONS(), maxOutputFinalizations);
        assertEq(valMgr.BASE_REWARD(), baseReward);
        assertEq(valMgr.MPT_FIRST_OUTPUT_INDEX(), mptFirstOutputIndex);
    }

    function test_constructor_smallMinActivateAmount_reverts() external {
        constructorParams._minRegisterAmount = minActivateAmount + 1;
        vm.expectRevert(IValidatorManager.InvalidConstructorParams.selector);
        new MockValidatorManager(constructorParams);
    }

    function test_registerValidator_active_succeeds() external {
        uint256 trustedBalance = assetToken.balanceOf(trusted);
        uint32 count = valMgr.activatedValidatorCount();

        uint128 assets = minActivateAmount;
        uint8 commissionRate = 10;

        vm.startPrank(trusted, trusted);
        assetToken.approve(address(assetMgr), uint256(assets));
        vm.expectEmit(true, false, false, true, address(valMgr));
        emit ValidatorActivated(trusted, block.timestamp);
        vm.expectEmit(true, true, false, true, address(valMgr));
        emit ValidatorRegistered(trusted, true, commissionRate, assets);
        valMgr.registerValidator(assets, commissionRate, withdrawAcc);
        vm.stopPrank();

        assertEq(assetToken.balanceOf(trusted), trustedBalance - assets);
        assertEq(assetMgr.totalValidatorKro(trusted), assets);
        assertEq(valMgr.getCommissionRate(trusted), commissionRate);
        assertEq(assetMgr.getWithdrawAccount(trusted), withdrawAcc);

        assertTrue(valMgr.getStatus(trusted) == IValidatorManager.ValidatorStatus.ACTIVE);
        assertEq(valMgr.activatedValidatorCount(), count + 1);
        assertEq(valMgr.getWeight(trusted), assets);
    }

    function test_registerValidator_registered_succeeds() external {
        uint32 count = valMgr.activatedValidatorCount();

        uint128 assets = minActivateAmount - 1;
        uint8 commissionRate = 10;

        vm.startPrank(trusted, trusted);
        assetToken.approve(address(assetMgr), uint256(assets));
        vm.expectEmit(true, true, false, true, address(valMgr));
        emit ValidatorRegistered(trusted, false, commissionRate, assets);
        valMgr.registerValidator(assets, commissionRate, withdrawAcc);
        vm.stopPrank();

        assertTrue(valMgr.getStatus(trusted) == IValidatorManager.ValidatorStatus.REGISTERED);
        assertEq(valMgr.activatedValidatorCount(), count);
        assertEq(valMgr.getWeight(trusted), 0);
    }

    function test_registerValidator_fromContract_reverts() external {
        vm.prank(address(oracle), address(oracle));
        vm.expectRevert(IValidatorManager.NotAllowedCaller.selector);
        valMgr.registerValidator(minActivateAmount, 10, withdrawAcc);
    }

    function test_registerValidator_differentOrigin_reverts() external {
        vm.prank(trusted, asserter);
        vm.expectRevert(IValidatorManager.NotAllowedCaller.selector);
        valMgr.registerValidator(minActivateAmount, 10, withdrawAcc);
    }

    function test_registerValidator_alreadyInitiated_reverts() external {
        uint128 assets = minActivateAmount;

        _registerValidator(trusted, assets);

        vm.startPrank(trusted, trusted);
        assetToken.approve(address(assetMgr), uint256(assets));
        vm.expectRevert(IValidatorManager.ImproperValidatorStatus.selector);
        valMgr.registerValidator(assets, 10, withdrawAcc);
    }

    function test_registerValidator_smallAsset_reverts() external {
        uint128 assets = minRegisterAmount - 1;

        vm.startPrank(trusted, trusted);
        assetToken.approve(address(assetMgr), uint256(assets));
        vm.expectRevert(IValidatorManager.InsufficientAsset.selector);
        valMgr.registerValidator(assets, 10, withdrawAcc);
    }

    function test_registerValidator_largeCommissionRate_reverts() external {
        uint128 assets = minRegisterAmount;

        vm.startPrank(trusted, trusted);
        assetToken.approve(address(assetMgr), uint256(assets));
        vm.expectRevert(IValidatorManager.MaxCommissionRateExceeded.selector);
        valMgr.registerValidator(assets, 101, withdrawAcc);
    }

    function test_registerValidator_withdrawZeroAddr_reverts() external {
        uint128 assets = minRegisterAmount;

        vm.startPrank(trusted, trusted);
        assetToken.approve(address(assetMgr), uint256(assets));
        vm.expectRevert(IAssetManager.ZeroAddress.selector);
        valMgr.registerValidator(assets, 10, address(0));
    }

    function test_activateValidator_notValidator_reverts() external {
        assertTrue(valMgr.getStatus(trusted) == IValidatorManager.ValidatorStatus.NONE);

        vm.prank(trusted);
        vm.expectRevert(IValidatorManager.ImproperValidatorStatus.selector);
        valMgr.activateValidator();
    }

    function test_activateValidator_registered_reverts() external {
        _registerValidator(trusted, minActivateAmount - 1);

        vm.prank(trusted);
        vm.expectRevert(IValidatorManager.ImproperValidatorStatus.selector);
        valMgr.activateValidator();
    }

    function test_activateValidator_exited_reverts() external {
        _registerValidator(trusted, minRegisterAmount);
        assertTrue(valMgr.getStatus(trusted) == IValidatorManager.ValidatorStatus.REGISTERED);

        _withdraw(trusted, 1);
        assertTrue(valMgr.getStatus(trusted) == IValidatorManager.ValidatorStatus.EXITED);

        vm.prank(trusted);
        vm.expectRevert(IValidatorManager.ImproperValidatorStatus.selector);
        valMgr.activateValidator();
    }

    function test_activateValidator_inJail_reverts() external {
        test_afterSubmitL2Output_tryJail_succeeds();

        // Withdraw all assets of jailed validator
        uint128 validatorKro = assetMgr.totalValidatorKro(asserter);
        _withdraw(asserter, validatorKro);
        assertEq(assetMgr.totalValidatorKro(asserter), 0);
        assertTrue(valMgr.getStatus(asserter) == IValidatorManager.ValidatorStatus.EXITED);

        // Delegate to re-activate validator
        vm.startPrank(asserter);
        assetToken.approve(address(assetMgr), minActivateAmount);
        assetMgr.deposit(minActivateAmount);
        vm.stopPrank();
        assertTrue(valMgr.getStatus(asserter) == IValidatorManager.ValidatorStatus.READY);

        vm.prank(asserter);
        vm.expectRevert(IValidatorManager.ImproperValidatorStatus.selector);
        valMgr.activateValidator();
    }

    function test_activateValidator_alreadyActivated_reverts() external {
        _registerValidator(trusted, minActivateAmount);

        vm.prank(trusted);
        vm.expectRevert(IValidatorManager.ImproperValidatorStatus.selector);
        valMgr.activateValidator();
    }

    function test_tryActivateValidator_succeeds() external {
        uint32 count = valMgr.activatedValidatorCount();

        _registerValidator(trusted, minActivateAmount - 1);
        assertTrue(valMgr.getStatus(trusted) == IValidatorManager.ValidatorStatus.REGISTERED);

        vm.startPrank(trusted);
        assetToken.approve(address(assetMgr), 1);
        vm.expectEmit(true, false, false, true, address(valMgr));
        emit ValidatorActivated(trusted, block.timestamp);
        assetMgr.deposit(1);
        vm.stopPrank();

        assertTrue(valMgr.getStatus(trusted) == IValidatorManager.ValidatorStatus.ACTIVE);
        assertEq(valMgr.activatedValidatorCount(), count + 1);
        assertEq(valMgr.getWeight(trusted), minActivateAmount);
    }

    function test_afterSubmitL2Output_distributeReward_succeeds() external {
        // Register validator with commission rate 10%
        _registerValidator(trusted, minActivateAmount);

        // Delegate 1 KGHs
        uint128 kghCounts = 1;
        _setUpKghDelegation(trusted, 1, 1);
        assertEq(assetMgr.totalKghNum(trusted), kghCounts);

        // Delegate KRO from 1 delegator
        uint128 delegateAsset = minActivateAmount;
        vm.startPrank(delegator);
        assetToken.approve(address(assetMgr), uint256(delegateAsset));
        assetMgr.delegate(trusted, delegateAsset);
        vm.stopPrank();

        assertEq(assetMgr.totalValidatorKro(trusted), minActivateAmount);
        assertEq(assetMgr.totalKroAssets(trusted), delegateAsset);

        // Submit the first output which interacts with ValidatorManager
        _submitL2OutputV2(false);

        // check KRO bonded
        assertEq(assetMgr.totalValidatorKroBonded(trusted), bondAmount);

        // Jump to the finalization time of the first output of ValidatorManager
        vm.warp(oracle.finalizedAt(terminateOutputIndex + 1));

        vm.startPrank(trusted);
        _submitL2OutputV2(true); // distribute reward 1 time

        uint128 expectedBaseReward = ((baseReward * 90) / 100) / 2; // delegator base reward
        uint128 boostedReward = mockValMgr.getBoostedReward(trusted);
        uint128 expectedBoostedReward = (boostedReward * 90) / 100;
        uint128 expectedValidatorReward = (((baseReward + boostedReward) * 10) / 100) + // commission
            ((baseReward * 90) / 100) /
            2;

        assertEq(assetMgr.totalKroAssets(trusted), delegateAsset + expectedBaseReward);
        assertEq(assetMgr.getKghReward(trusted, trusted), expectedBoostedReward);
        assertEq(assetMgr.totalValidatorKro(trusted), minActivateAmount + expectedValidatorReward);

        // check KRO bonded
        assertEq(assetMgr.totalValidatorKroBonded(trusted), bondAmount);

        // Check validator tree updated with rewards
        assertEq(
            valMgr.getWeight(trusted),
            minActivateAmount + delegateAsset + expectedBaseReward + expectedValidatorReward
        );
    }

    function test_afterSubmitL2Output_distributeRewardToSC_succeeds() external {
        // Register validator with commission rate 10%
        _registerValidator(trusted, minActivateAmount);

        // Submit the first output which interacts with ValidatorManager
        _submitL2OutputV2(false);

        // Change the output submitter to SC
        uint256 firstOutputIndex = terminateOutputIndex + 1;
        mockOracle.replaceOutput(assetMgr.SECURITY_COUNCIL(), firstOutputIndex);

        // Jump to the finalization time of the first output of ValidatorManager
        vm.warp(oracle.finalizedAt(firstOutputIndex));

        vm.startPrank(trusted);
        _submitL2OutputV2(true); // distribute reward 1 time

        // Check if the reward is transferred to SC directly
        assertEq(assetToken.balanceOf(assetMgr.SECURITY_COUNCIL()), baseReward);
    }

    function test_afterSubmitL2Output_updatePriorityValidator_succeeds() external {
        // Register as a validator
        _registerValidator(asserter, minActivateAmount);
        _registerValidator(trusted, minActivateAmount);

        // Check if next priority validator is not set in ValidatorManager
        assertTrue(mockValMgr.nextPriorityValidator() == address(0));
        assertEq(valMgr.nextValidator(), trusted);

        // Submit the first output which interacts with ValidatorManager
        _submitL2OutputV2(false);

        // Check if next finalize output is not updated
        assertEq(oracle.nextFinalizeOutputIndex(), terminateOutputIndex + 1);
        // Check if next priority validator is set in ValidatorManager
        address nextValidator = mockValMgr.nextPriorityValidator();
        assertTrue(nextValidator != address(0));
        assertTrue(valMgr.nextValidator() == nextValidator);

        // Jump to the finalization time of the first output of ValidatorManager
        vm.warp(oracle.finalizedAt(terminateOutputIndex + 1));
        vm.startPrank(nextValidator);
        _submitL2OutputV2(true);
        vm.stopPrank();

        // Check if next finalize output is updated
        assertEq(oracle.nextFinalizeOutputIndex(), terminateOutputIndex + 2);

        // Submit 10 outputs
        uint256 tries = 10;
        bool changed = false;
        nextValidator = valMgr.nextValidator();

        for (uint256 i; i < tries; i++) {
            // Submit next output and finalize prev output
            warpToSubmitTime();
            _submitL2OutputV2(false);

            // Check the next validator has changed
            address newValidator = valMgr.nextValidator();
            if (nextValidator != newValidator) {
                changed = true;
                break;
            }
        }

        assertTrue(changed);
    }

    function test_afterSubmitL2Output_tryJail_succeeds() public {
        // Register as a validator
        _registerValidator(asserter, minActivateAmount);
        _registerValidator(trusted, 10 * minActivateAmount);

        vm.startPrank(trusted);
        for (uint256 i; i < jailThreshold; i++) {
            mockValMgr.updatePriorityValidator(asserter);

            // Warp to public round
            vm.warp(oracle.nextOutputMinL2Timestamp() + roundDuration + 1);
            _submitL2OutputV2(true);

            assertEq(valMgr.noSubmissionCount(asserter), i + 1);
            assertFalse(valMgr.inJail(asserter));
        }
        vm.stopPrank();

        mockValMgr.updatePriorityValidator(asserter);

        // Warp to public round
        vm.warp(oracle.nextOutputMinL2Timestamp() + roundDuration + 1);
        vm.startPrank(trusted);
        vm.expectEmit(true, false, false, true, address(valMgr));
        emit ValidatorJailed(asserter, uint128(block.timestamp) + softJailPeriodSeconds);
        vm.expectEmit(true, false, false, true, address(valMgr));
        emit ValidatorStopped(asserter, block.timestamp);
        _submitL2OutputV2(true);
        vm.stopPrank();

        assertEq(valMgr.noSubmissionCount(asserter), jailThreshold);
        assertTrue(valMgr.inJail(asserter));
        assertEq(valMgr.getWeight(asserter), 0);
    }

    function test_afterSubmitL2Output_resetNoSubmissionCount_succeeds() external {
        // Register as a validator
        _registerValidator(asserter, minActivateAmount);
        _registerValidator(trusted, minActivateAmount);

        mockValMgr.updatePriorityValidator(asserter);

        // Warp to public round
        vm.warp(oracle.nextOutputMinL2Timestamp() + roundDuration + 1);
        assertEq(valMgr.nextValidator(), Constants.VALIDATOR_PUBLIC_ROUND_ADDRESS);
        vm.startPrank(trusted);
        _submitL2OutputV2(true);
        vm.stopPrank();

        assertEq(valMgr.noSubmissionCount(asserter), 1);
        assertFalse(valMgr.inJail(asserter));

        mockValMgr.updatePriorityValidator(asserter);

        // Warp to priority round
        warpToSubmitTime();
        _submitL2OutputV2(false);

        assertEq(valMgr.noSubmissionCount(asserter), 0);
    }

    function test_afterSubmitL2Output_senderNotL2OO_reverts() external {
        vm.prank(trusted);
        vm.expectRevert(IValidatorManager.NotAllowedCaller.selector);
        valMgr.afterSubmitL2Output(0);
    }

    function test_initCommissionChange_succeeds() public {
        _registerValidator(asserter, minActivateAmount);

        uint8 commissionRate = valMgr.getCommissionRate(asserter);
        uint8 newCommissionRate = commissionRate + 1;

        vm.prank(asserter);
        vm.expectEmit(true, false, false, true, address(valMgr));
        emit ValidatorCommissionChangeInitiated(asserter, commissionRate, newCommissionRate);
        valMgr.initCommissionChange(newCommissionRate);

        assertEq(valMgr.getPendingCommissionRate(asserter), newCommissionRate);
        assertEq(
            valMgr.canFinalizeCommissionChangeAt(asserter),
            block.timestamp + valMgr.COMMISSION_CHANGE_DELAY_SECONDS()
        );
    }

    function test_initCommissionChange_exited_reverts() external {
        _registerValidator(trusted, minActivateAmount);

        uint128 validatorKro = assetMgr.totalValidatorKro(trusted);
        _withdraw(trusted, validatorKro);
        assertTrue(valMgr.getStatus(trusted) == IValidatorManager.ValidatorStatus.EXITED);

        vm.prank(trusted);
        vm.expectRevert(IValidatorManager.ImproperValidatorStatus.selector);
        valMgr.initCommissionChange(15);
    }

    function test_initCommissionChange_inJail_reverts() external {
        test_afterSubmitL2Output_tryJail_succeeds();

        vm.prank(asserter);
        vm.expectRevert(IValidatorManager.ImproperValidatorStatus.selector);
        valMgr.initCommissionChange(15);
    }

    function test_initCommissionChange_largeCommissionRate_reverts() external {
        _registerValidator(asserter, minActivateAmount);

        vm.prank(asserter);
        vm.expectRevert(IValidatorManager.MaxCommissionRateExceeded.selector);
        valMgr.initCommissionChange(101);
    }

    function test_initCommissionChange_sameCommissionRate_reverts() external {
        _registerValidator(asserter, minActivateAmount);

        uint8 commissionRate = valMgr.getCommissionRate(asserter);
        vm.prank(asserter);
        vm.expectRevert(IValidatorManager.SameCommissionRate.selector);
        valMgr.initCommissionChange(commissionRate);
    }

    function test_finalizeCommissionChange_succeeds() public {
        test_initCommissionChange_succeeds();

        uint8 oldCommissionRate = valMgr.getCommissionRate(asserter);
        uint8 newCommissionRate = valMgr.getPendingCommissionRate(asserter);

        vm.warp(valMgr.canFinalizeCommissionChangeAt(asserter));
        vm.prank(asserter);
        vm.expectEmit(true, false, false, true, address(valMgr));
        emit ValidatorCommissionChangeFinalized(asserter, oldCommissionRate, newCommissionRate);
        valMgr.finalizeCommissionChange();

        assertEq(valMgr.getCommissionRate(asserter), newCommissionRate);
        assertEq(valMgr.getPendingCommissionRate(asserter), 0);
        assertEq(
            valMgr.canFinalizeCommissionChangeAt(asserter),
            valMgr.COMMISSION_CHANGE_DELAY_SECONDS()
        );
    }

    function test_finalizeCommissionChange_exited_reverts() external {
        _registerValidator(trusted, minActivateAmount);

        vm.prank(trusted);
        valMgr.initCommissionChange(15);

        uint128 validatorKro = assetMgr.totalValidatorKro(trusted);
        _withdraw(trusted, validatorKro);
        assertTrue(valMgr.getStatus(trusted) == IValidatorManager.ValidatorStatus.EXITED);

        vm.prank(trusted);
        vm.expectRevert(IValidatorManager.ImproperValidatorStatus.selector);
        valMgr.finalizeCommissionChange();
    }

    function test_finalizeCommissionChange_inJail_reverts() external {
        test_afterSubmitL2Output_tryJail_succeeds();

        vm.prank(asserter);
        vm.expectRevert(IValidatorManager.ImproperValidatorStatus.selector);
        valMgr.finalizeCommissionChange();
    }

    function test_finalizeCommissionChange_changeDelayNotElapsed_reverts() external {
        test_initCommissionChange_succeeds();

        vm.prank(asserter);
        vm.expectRevert(IValidatorManager.NotElapsedCommissionChangeDelay.selector);
        valMgr.finalizeCommissionChange();
    }

    function test_finalizeCommissionChange_notInitiated_reverts() external {
        _registerValidator(trusted, minActivateAmount);

        vm.prank(trusted);
        vm.expectRevert(IValidatorManager.NotInitiatedCommissionChange.selector);
        valMgr.finalizeCommissionChange();
    }

    function test_tryUnjail_succeeds() external {
        test_afterSubmitL2Output_tryJail_succeeds();
        assertTrue(valMgr.getStatus(asserter) == IValidatorManager.ValidatorStatus.READY);

        vm.warp(valMgr.jailExpiresAt(asserter));
        vm.prank(asserter);
        vm.expectEmit(true, false, false, false, address(valMgr));
        emit ValidatorUnjailed(asserter);
        vm.expectEmit(true, false, false, true, address(valMgr));
        emit ValidatorActivated(asserter, block.timestamp);
        valMgr.tryUnjail();

        assertEq(valMgr.noSubmissionCount(asserter), 0);
        assertFalse(valMgr.inJail(asserter));
        assertTrue(valMgr.getStatus(asserter) == IValidatorManager.ValidatorStatus.ACTIVE);
    }

    function test_tryUnjail_notInJail_reverts() external {
        vm.prank(asserter);
        vm.expectRevert(IValidatorManager.ImproperValidatorStatus.selector);
        valMgr.tryUnjail();
    }

    function test_tryUnjail_senderNotSelf_reverts() external {
        test_afterSubmitL2Output_tryJail_succeeds();

        vm.prank(trusted);
        vm.expectRevert(IValidatorManager.ImproperValidatorStatus.selector);
        valMgr.tryUnjail();
    }

    function test_tryUnjail_periodNotElapsed_reverts() external {
        test_afterSubmitL2Output_tryJail_succeeds();

        vm.prank(asserter);
        vm.expectRevert(IValidatorManager.NotElapsedJailPeriod.selector);
        valMgr.tryUnjail();
    }

    function test_bondValidatorKro_succeeds() external {
        _registerValidator(trusted, minActivateAmount);
        assertEq(assetMgr.totalValidatorKroBonded(trusted), 0);

        vm.prank(address(colosseum));
        valMgr.bondValidatorKro(trusted);

        assertEq(assetMgr.totalValidatorKroBonded(trusted), bondAmount);
    }

    function test_bondValidatorKro_notColosseum_reverts() external {
        _registerValidator(trusted, minActivateAmount);

        vm.prank(asserter);
        vm.expectRevert(IValidatorManager.NotAllowedCaller.selector);
        valMgr.bondValidatorKro(trusted);
    }

    function test_unbondValidatorKro_succeeds() external {
        _registerValidator(trusted, minActivateAmount);

        vm.prank(trusted);
        _submitL2OutputV2(false);

        assertEq(assetMgr.totalValidatorKroBonded(trusted), bondAmount);

        vm.prank(address(colosseum));
        valMgr.unbondValidatorKro(trusted);

        assertEq(assetMgr.totalValidatorKroBonded(trusted), 0);
    }

    function test_unbondValidatorKro_notColosseum_reverts() external {
        _registerValidator(trusted, minActivateAmount);

        vm.prank(trusted);
        _submitL2OutputV2(false);

        assertEq(assetMgr.totalValidatorKroBonded(trusted), bondAmount);

        vm.prank(asserter);
        vm.expectRevert(IValidatorManager.NotAllowedCaller.selector);
        valMgr.unbondValidatorKro(trusted);
    }

    function test_slash_succeeds() external {
        uint32 count = valMgr.activatedValidatorCount();
        // Register as a validator
        _registerValidator(asserter, minActivateAmount);
        _registerValidator(challenger, minActivateAmount);
        assertEq(valMgr.activatedValidatorCount(), count + 2);

        // Delegate KGHs
        uint128 kghCounts = 1;
        uint128 startingTokenId = 1;
        _setUpKghDelegation(asserter, startingTokenId, kghCounts);
        _setUpKghDelegation(challenger, startingTokenId + kghCounts, kghCounts);
        assertEq(assetMgr.totalKghNum(asserter), kghCounts);
        assertEq(assetMgr.totalKghNum(challenger), kghCounts);

        // Submit the first output which interacts with ValidatorManager
        mockValMgr.updatePriorityValidator(asserter);
        warpToSubmitTime();
        _submitL2OutputV2(false);
        uint256 challengedOutputIndex = oracle.latestOutputIndex();

        // Suppose that the challenge is successful, so the winner is challenger
        uint128 slashingAmount = bondAmount;
        vm.startPrank(address(colosseum));
        valMgr.bondValidatorKro(challenger); // bond for creating challenge
        vm.expectEmit(true, true, false, true, address(valMgr));
        emit Slashed(challengedOutputIndex, asserter, slashingAmount);
        vm.expectEmit(true, false, false, true, address(valMgr));
        emit ValidatorJailed(asserter, uint128(block.timestamp) + hardJailPeriodSeconds);
        vm.expectEmit(true, false, false, true, address(valMgr));
        emit ValidatorStopped(asserter, block.timestamp);
        valMgr.slash(challengedOutputIndex, challenger, asserter);
        vm.stopPrank();

        // This will be done by the l2 output oracle contract in the real environment
        vm.prank(challenger);
        mockOracle.replaceOutput(challenger, challengedOutputIndex);

        // Jump to the finalization time of the challenged output
        vm.warp(oracle.finalizedAt(challengedOutputIndex));
        vm.startPrank(challenger);
        // Submit one more output to distribute reward
        _submitL2OutputV2(true);
        vm.stopPrank();

        // Asserter in jail after slashed
        assertTrue(valMgr.inJail(asserter));
        // Asserter removed from validator tree
        assertEq(valMgr.activatedValidatorCount(), count + 1);
        assertTrue(valMgr.getStatus(asserter) == IValidatorManager.ValidatorStatus.REGISTERED);

        // Asserter asset decreased by slashingAmount
        uint128 asserterTotalValidatorKro = assetMgr.totalValidatorKro(asserter);
        assertEq(asserterTotalValidatorKro, minActivateAmount - slashingAmount);
        assertEq(assetMgr.totalValidatorKroBonded(asserter), 0);

        // Security council balance of asset token increased by tax
        uint128 taxAmount = (slashingAmount * assetMgr.TAX_NUMERATOR()) /
            assetMgr.TAX_DENOMINATOR();
        assertEq(assetToken.balanceOf(assetMgr.SECURITY_COUNCIL()), taxAmount);

        // Challenger asset increased by output reward and challenge reward
        // Boosted reward with 1 kgh delegation
        uint128 boostedReward = mockValMgr.getBoostedReward(challenger);
        uint128 challengeReward = slashingAmount - taxAmount;
        uint128 challengerKro = assetMgr.totalValidatorKro(challenger);
        assertEq(
            challengerKro,
            minActivateAmount + baseReward + (boostedReward / 10) + challengeReward
        );
    }

    function test_slash_alreadyInJail_succeeds() external {
        _registerValidator(asserter, minActivateAmount);

        // Submit the first output which interacts with ValidatorManager
        mockValMgr.updatePriorityValidator(asserter);
        warpToSubmitTime();
        _submitL2OutputV2(false);
        uint256 challengedOutputIndex = oracle.latestOutputIndex();

        // Before slashed, send to jail
        uint128 firstJailExpiresAt = uint128(block.timestamp) + softJailPeriodSeconds;
        mockValMgr.setJailExpiresAt(asserter, firstJailExpiresAt);

        // After jail expired, slash
        vm.warp(firstJailExpiresAt + 1);
        vm.prank(address(colosseum));
        valMgr.slash(challengedOutputIndex, challenger, asserter);

        assertEq(valMgr.jailExpiresAt(asserter), firstJailExpiresAt + 1 + hardJailPeriodSeconds);
    }

    function test_slash_notColosseum_reverts() external {
        vm.prank(address(1));
        vm.expectRevert(IValidatorManager.NotAllowedCaller.selector);
        valMgr.slash(1, challenger, asserter);
    }

    function test_revertSlash_succeeds() external {
        uint32 count = valMgr.activatedValidatorCount();
        // Register as a validator
        _registerValidator(asserter, minActivateAmount);
        _registerValidator(challenger, minActivateAmount);
        assertEq(valMgr.activatedValidatorCount(), count + 2);

        // Delegate KGHs
        uint128 kghCounts = 1;
        uint128 startingTokenId = 1;
        _setUpKghDelegation(asserter, startingTokenId, kghCounts);
        _setUpKghDelegation(challenger, startingTokenId + kghCounts, kghCounts);
        assertEq(assetMgr.totalKghNum(asserter), kghCounts);
        assertEq(assetMgr.totalKghNum(challenger), kghCounts);

        // Submit the first output which interacts with ValidatorManager
        mockValMgr.updatePriorityValidator(asserter);
        warpToSubmitTime();
        _submitL2OutputV2(false);
        uint256 challengedOutputIndex = oracle.latestOutputIndex();

        // Suppose that the challenge is successful, so the winner is challenger
        uint128 slashingAmount = bondAmount;
        vm.startPrank(address(colosseum));
        valMgr.bondValidatorKro(challenger); // bond for creating challenge
        vm.expectEmit(true, true, false, true, address(valMgr));
        emit Slashed(challengedOutputIndex, asserter, slashingAmount);
        vm.expectEmit(true, false, false, true, address(valMgr));
        emit ValidatorJailed(asserter, uint128(block.timestamp) + hardJailPeriodSeconds);
        vm.expectEmit(true, false, false, true, address(valMgr));
        emit ValidatorStopped(asserter, block.timestamp);
        valMgr.slash(challengedOutputIndex, challenger, asserter);
        vm.stopPrank();

        // Asserter in jail after slashed
        assertTrue(valMgr.inJail(asserter));

        // Revert slash
        vm.startPrank(address(colosseum));
        vm.expectEmit(true, true, true, true, address(valMgr));
        emit SlashReverted(challengedOutputIndex, asserter, slashingAmount);
        vm.expectEmit(true, true, false, false, address(valMgr));
        emit ValidatorUnjailed(asserter);
        valMgr.revertSlash(challengedOutputIndex, asserter);

        assertEq(assetMgr.totalValidatorKro(asserter), minActivateAmount);
        assertEq(assetMgr.totalValidatorKroBonded(asserter), bondAmount);
        assertFalse(valMgr.inJail(asserter));
        assertTrue(valMgr.getStatus(asserter) == IValidatorManager.ValidatorStatus.ACTIVE);
    }

    function test_revertSlash_stillInJail_succeeds() external {
        _registerValidator(asserter, minActivateAmount);

        // Submit the first output which interacts with ValidatorManager
        mockValMgr.updatePriorityValidator(asserter);
        warpToSubmitTime();
        _submitL2OutputV2(false);
        uint256 challengedOutputIndex = oracle.latestOutputIndex();

        // Before slashed, send to jail
        uint128 firstJailExpiresAt = uint128(block.timestamp) + softJailPeriodSeconds;
        mockValMgr.setJailExpiresAt(asserter, firstJailExpiresAt);

        // Before jail expired, slash
        vm.prank(address(colosseum));
        valMgr.slash(challengedOutputIndex, challenger, asserter);

        assertEq(valMgr.jailExpiresAt(asserter), firstJailExpiresAt + hardJailPeriodSeconds);

        // Revert slash
        vm.prank(address(colosseum));
        valMgr.revertSlash(challengedOutputIndex, asserter);

        assertEq(valMgr.jailExpiresAt(asserter), firstJailExpiresAt);
    }

    function test_revertSlash_notColosseum_reverts() external {
        vm.prank(trusted);
        vm.expectRevert(IValidatorManager.NotAllowedCaller.selector);
        valMgr.revertSlash(1, trusted);
    }

    function test_checkSubmissionEligibility_priorityRound_succeeds() external {
        address nextValidator = valMgr.nextValidator();
        _registerValidator(nextValidator, minActivateAmount);

        vm.prank(address(oracle));
        valMgr.checkSubmissionEligibility(nextValidator);
    }

    function test_checkSubmissionEligibility_publicRound_succeeds() external {
        mockValMgr.updatePriorityValidator(asserter);

        _registerValidator(trusted, minActivateAmount);

        // Warp to public round
        vm.warp(oracle.nextOutputMinL2Timestamp() + roundDuration + 1);
        vm.prank(address(oracle));
        valMgr.checkSubmissionEligibility(trusted);
    }

    function test_checkSubmissionEligibility_senderNotL2OO_reverts() external {
        vm.expectRevert(IValidatorManager.NotAllowedCaller.selector);
        valMgr.checkSubmissionEligibility(trusted);
    }

    function test_checkSubmissionEligibility_notSelected_reverts() external {
        mockValMgr.updatePriorityValidator(asserter);

        warpToSubmitTime();
        vm.prank(address(oracle));
        vm.expectRevert(IValidatorManager.NotSelectedPriorityValidator.selector);
        valMgr.checkSubmissionEligibility(trusted);
    }

    function test_checkSubmissionEligibility_notSatisfyCondition_reverts() external {
        address nextValidator = valMgr.nextValidator();

        vm.prank(address(oracle));
        vm.expectRevert(IValidatorManager.ImproperValidatorStatus.selector);
        valMgr.checkSubmissionEligibility(nextValidator);
    }

    function test_checkSubmissionEligibility_publicRound_notSatisfyCondition_reverts() external {
        mockValMgr.updatePriorityValidator(asserter);

        // Warp to public round
        vm.warp(oracle.nextOutputMinL2Timestamp() + roundDuration + 1);
        vm.prank(address(oracle));
        vm.expectRevert(IValidatorManager.ImproperValidatorStatus.selector);
        valMgr.checkSubmissionEligibility(trusted);
    }

    function test_checkSubmissionEligibility_inJail_reverts() external {
        test_afterSubmitL2Output_tryJail_succeeds();

        mockValMgr.updatePriorityValidator(asserter);

        vm.prank(address(oracle));
        vm.expectRevert(IValidatorManager.ImproperValidatorStatus.selector);
        valMgr.checkSubmissionEligibility(asserter);
    }

    function test_checkSubmissionEligibility_publicRound_inJail_reverts() external {
        test_afterSubmitL2Output_tryJail_succeeds();

        mockValMgr.updatePriorityValidator(trusted);

        // Warp to public round
        vm.warp(oracle.nextOutputMinL2Timestamp() + roundDuration + 1);
        vm.prank(address(oracle));
        vm.expectRevert(IValidatorManager.ImproperValidatorStatus.selector);
        valMgr.checkSubmissionEligibility(asserter);
    }

    function test_getStatus_registered_succeeds() external {
        _registerValidator(trusted, minActivateAmount);
        assertEq(valMgr.getWeight(trusted), minActivateAmount);

        _withdraw(trusted, 1);
        assertTrue(valMgr.getStatus(trusted) == IValidatorManager.ValidatorStatus.REGISTERED);
        assertEq(valMgr.getWeight(trusted), 0);
    }

    function test_activatedValidatorTotalWeight_succeeds() external {
        uint32 count = valMgr.activatedValidatorCount();
        uint120 totalWeight = valMgr.activatedValidatorTotalWeight();
        _registerValidator(trusted, minActivateAmount);
        assertEq(valMgr.activatedValidatorCount(), count + 1);
        assertEq(valMgr.activatedValidatorTotalWeight(), totalWeight + minActivateAmount);

        count = valMgr.activatedValidatorCount();
        totalWeight = valMgr.activatedValidatorTotalWeight();
        _registerValidator(asserter, minActivateAmount);
        assertEq(valMgr.activatedValidatorCount(), count + 1);
        assertEq(valMgr.activatedValidatorTotalWeight(), totalWeight + minActivateAmount);

        count = valMgr.activatedValidatorCount();
        totalWeight = valMgr.activatedValidatorTotalWeight();
        vm.startPrank(challenger);
        assetToken.approve(address(assetMgr), 10);
        assetMgr.delegate(asserter, 10);
        vm.stopPrank();
        assertEq(valMgr.activatedValidatorCount(), count);
        assertEq(valMgr.activatedValidatorTotalWeight(), totalWeight + 10);
    }
}

contract ValidatorManager_MptTransition_Test is ValidatorSystemUpgrade_Initializer {
    MockL2OutputOracle mockOracle;
    MockValidatorManager mockValMgr;
    function setUp() public override {
        super.setUp();

        address oracleAddress = address(oracle);
        MockL2OutputOracle mockOracleImpl = new MockL2OutputOracle(
            pool,
            valMgr,
            address(colosseum),
            submissionInterval,
            l2BlockTime,
            startingBlockNumber,
            startingTimestamp,
            finalizationPeriodSeconds
        );
        vm.prank(multisig);
        Proxy(payable(oracleAddress)).upgradeTo(address(mockOracleImpl));
        mockOracle = MockL2OutputOracle(oracleAddress);

        // upgrade validatorManager with new mptFirstOutputIndex param
        mptFirstOutputIndex = 10;
        constructorParams._mptFirstOutputIndex = mptFirstOutputIndex;
        address valMgrAddress = address(valMgr);
        MockValidatorManager mockValMgrImpl = new MockValidatorManager(constructorParams);
        vm.prank(multisig);
        Proxy(payable(valMgrAddress)).upgradeTo(address(mockValMgrImpl));
        mockValMgr = MockValidatorManager(valMgrAddress);

        // Submit until terminateOutputIndex and set next output index to be finalized after it
        vm.prank(trusted);
        pool.deposit{ value: trusted.balance }();
        for (uint256 i = oracle.nextOutputIndex(); i <= terminateOutputIndex; i++) {
            _submitL2OutputV1();
        }

        // Since the first output of V2 is always for trusted validator,
        // so for accurate test, we should pass the first V2 output
        uint128 assets = minActivateAmount;
        _registerValidator(asserter, assets);
        _registerValidator(trusted, assets);
        for (uint256 i = oracle.nextOutputIndex(); i < mptFirstOutputIndex; i++) {
            warpToSubmitTime();
            _submitL2OutputV2(false);
        }
        vm.warp(oracle.finalizedAt(mptFirstOutputIndex - 1));
        mockOracle.mockSetNextFinalizeOutputIndex(mptFirstOutputIndex);
    }

    function test_submitL2Output_mptFirstOutput_privateRound_trustedValidator_succeeds() public {
        warpToSubmitTime();
        assertEq(valMgr.nextValidator(), trusted);
        _submitL2OutputV2(false);
    }

    function test_submitL2Output_mptFirstOutput_publicRound_trustedValidator_succeeds() public {
        // Warp to public round
        vm.warp(oracle.nextOutputMinL2Timestamp() + roundDuration + 1);
        vm.startPrank(trusted, trusted);
        _submitL2OutputV2(true);
    }

    function test_submitL2Output_mptFirstOutput_publicRound_notTrustedValidator_reverts() public {
        // Warp to public round
        vm.warp(oracle.nextOutputMinL2Timestamp() + roundDuration + 1);

        uint256 nextBlockNumber = oracle.nextBlockNumber();
        bytes32 outputRoot = keccak256(abi.encode(nextBlockNumber));
        vm.startPrank(asserter, asserter);
        vm.expectRevert(IValidatorManager.NotSelectedPriorityValidator.selector);
        oracle.submitL2Output(outputRoot, nextBlockNumber, 0, 0);
    }

    function test_submitL2Output_mptFirstOutput_afterUpgrade_notTrustedValidator_succeeds() public {
        uint128 newMptFirstOutputIndex = 1000;
        // Warp to public round
        vm.warp(oracle.nextOutputMinL2Timestamp() + roundDuration + 1);

        // upgrade validatorManager with newMptFirstOutput
        address valMgrAddress = address(valMgr);
        constructorParams._mptFirstOutputIndex = newMptFirstOutputIndex;
        MockValidatorManager mockValMgrImpl = new MockValidatorManager(constructorParams);
        vm.prank(multisig);
        Proxy(payable(valMgrAddress)).upgradeTo(address(mockValMgrImpl));
        mockValMgr = MockValidatorManager(valMgrAddress);

        uint256 nextBlockNumber = oracle.nextBlockNumber();
        bytes32 outputRoot = keccak256(abi.encode(nextBlockNumber));
        vm.startPrank(asserter, asserter);
        oracle.submitL2Output(outputRoot, nextBlockNumber, 0, 0);
    }
}
