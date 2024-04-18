// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { Constants } from "../libraries/Constants.sol";
import { Types } from "../libraries/Types.sol";
import { Proxy } from "../universal/Proxy.sol";
import { IValidatorManager } from "../L1/interfaces/IValidatorManager.sol";
import { L2OutputOracle } from "../L1/L2OutputOracle.sol";
import { ValidatorManager } from "../L1/ValidatorManager.sol";
import { ValidatorPool } from "../L1/ValidatorPool.sol";
import { L2OutputOracle_ValidatorSystemUpgrade_Initializer } from "./CommonTest.t.sol";

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

    function replaceOutput(uint256 outputIndex) external {
        l2Outputs[outputIndex].submitter = msg.sender;
        l2Outputs[outputIndex].outputRoot = bytes32(0);
    }

    function mockSetLatestFinalizedOutputIndex(uint256 l2OutputIndex) external {
        latestFinalizedOutputIndex = l2OutputIndex;
    }
}

contract MockValidatorManager is ValidatorManager {
    constructor(ConstructorParams memory _constructorParams) ValidatorManager(_constructorParams) {}

    function updatePriorityValidator(address validator) external {
        _nextPriorityValidator = validator;
    }

    function nextPriorityValidator() external view returns (address) {
        return _nextPriorityValidator;
    }

    function commissionRateChangedAt(address validator) external view returns (uint128) {
        return _validatorInfo[validator].commissionRateChangedAt;
    }
}

contract ValidatorManagerTest is L2OutputOracle_ValidatorSystemUpgrade_Initializer {
    MockL2OutputOracle mockOracle;
    MockValidatorManager mockValMan;
    uint128 public VKRO_PER_KGH;

    event ValidatorRegistered(
        address indexed validator,
        bool started,
        uint8 commissionRate,
        uint8 commissionMaxChangeRate,
        uint128 assets
    );

    event ValidatorStarted(address indexed validator, uint256 startsAt);

    event ValidatorCommissionRateChanged(
        address indexed validator,
        uint8 oldCommissionRate,
        uint8 newCommissionRate
    );

    event ValidatorJailed(address indexed validator, uint128 expiresAt);

    event ValidatorUnjailed(address indexed validator);

    event RewardDistributed(
        address indexed validator,
        uint128 validatorReward,
        uint128 baseReward,
        uint128 boostedReward
    );

    function _submitL2Output(uint256 l2BlockNumber, bool isPublicRound) private {
        uint256 outputIndex = oracle.nextOutputIndex();
        if (!isPublicRound) {
            vm.prank(valMan.nextValidator());
        }
        mockOracle.addOutput(l2BlockNumber);
        vm.prank(address(oracle));
        valMan.afterSubmitL2Output(outputIndex);
    }

    function _setUpHundredKghDelegation(address validator, uint256 startingTokenId) private {
        uint256[] memory tokenIds = new uint256[](100);
        for (uint256 i = startingTokenId; i < 100 + startingTokenId; i++) {
            kgh.mint(validator, i);
            vm.prank(validator);
            kgh.approve(address(assetMan), i);
            tokenIds[i - startingTokenId] = i;
        }
        vm.prank(validator);
        assetMan.delegateKghBatch(validator, tokenIds);
    }

    function setUp() public override {
        super.setUp();

        address oracleAddress = address(oracle);
        MockL2OutputOracle mockOracleImpl = new MockL2OutputOracle(
            pool,
            valMan,
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

        address valManAddress = address(valMan);
        MockValidatorManager mockValManImpl = new MockValidatorManager(constructorParams);
        vm.prank(multisig);
        Proxy(payable(valManAddress)).upgradeTo(address(mockValManImpl));
        mockValMan = MockValidatorManager(valManAddress);

        VKRO_PER_KGH = assetMan.VKRO_PER_KGH();
    }

    function test_constructor_succeeds() external {
        assertEq(address(valMan.L2_ORACLE()), address(oracle));
        assertEq(address(valMan.ASSET_MANAGER()), address(assetMan));
        assertEq(valMan.TRUSTED_VALIDATOR(), trusted);
        assertEq(valMan.MIN_REGISTER_AMOUNT(), minRegisterAmount);
        assertEq(valMan.MIN_START_AMOUNT(), minStartAmount);
        assertEq(valMan.COMMISSION_RATE_MIN_CHANGE_SECONDS(), commissionRateMinChangeSeconds);
        assertEq(valMan.ROUND_DURATION_SECONDS(), roundDuration);
        assertEq(valMan.JAIL_PERIOD_SECONDS(), jailPeriodSeconds);
        assertEq(valMan.JAIL_THRESHOLD(), jailThreshold);
        assertEq(valMan.MAX_OUTPUT_FINALIZATIONS(), maxOutputFinalizations);
        assertEq(valMan.BASE_REWARD(), baseReward);
    }

    function test_constructor_smallMinStartAmount_reverts() external {
        constructorParams._minRegisterAmount = minStartAmount + 1;
        vm.expectRevert(IValidatorManager.InvalidConstructorParams.selector);
        new MockValidatorManager(constructorParams);
    }

    function test_registerValidator_canSubmitOutput_succeeds() external {
        uint256 trustedBalance = assetToken.balanceOf(trusted);
        uint32 count = valMan.startedValidatorCount();

        uint128 assets = minStartAmount;
        uint8 commissionRate = 10;
        uint8 commissionMaxChangeRate = 5;

        vm.startPrank(trusted);
        assetToken.approve(address(assetMan), uint256(assets));
        vm.expectEmit(true, true, false, true, address(valMan));
        emit ValidatorRegistered(trusted, true, commissionRate, commissionMaxChangeRate, assets);
        valMan.registerValidator(assets, commissionRate, commissionMaxChangeRate);
        vm.stopPrank();

        assertEq(assetToken.balanceOf(trusted), trustedBalance - assets);
        assertEq(assetMan.totalKroAssets(trusted), assets);
        assertEq(valMan.getCommissionRate(trusted), commissionRate);
        assertEq(valMan.getCommissionMaxChangeRate(trusted), commissionMaxChangeRate);
        assertEq(mockValMan.commissionRateChangedAt(trusted), block.timestamp);

        assertTrue(
            valMan.getStatus(trusted) == IValidatorManager.ValidatorStatus.CAN_SUBMIT_OUTPUT
        );
        assertEq(valMan.startedValidatorCount(), count + 1);
        assertEq(valMan.getWeight(trusted), assets);
    }

    function test_registerValidator_active_succeeds() external {
        uint32 count = valMan.startedValidatorCount();

        uint128 assets = minStartAmount - 1;
        uint8 commissionRate = 10;
        uint8 commissionMaxChangeRate = 5;

        vm.startPrank(trusted);
        assetToken.approve(address(assetMan), uint256(assets));
        vm.expectEmit(true, true, false, true, address(valMan));
        emit ValidatorRegistered(trusted, false, commissionRate, commissionMaxChangeRate, assets);
        valMan.registerValidator(assets, commissionRate, commissionMaxChangeRate);
        vm.stopPrank();

        assertTrue(valMan.getStatus(trusted) == IValidatorManager.ValidatorStatus.ACTIVE);
        assertEq(valMan.startedValidatorCount(), count);
        assertEq(valMan.getWeight(trusted), 0);
    }

    function test_registerValidator_alreadyInitiated_reverts() external {
        uint128 assets = minStartAmount;

        _registerValidator(trusted, assets);

        vm.startPrank(trusted);
        assetToken.approve(address(assetMan), uint256(assets));
        vm.expectRevert(IValidatorManager.ImproperValidatorStatus.selector);
        valMan.registerValidator(assets, 10, 5);
    }

    function test_registerValidator_smallAsset_reverts() external {
        uint128 assets = minRegisterAmount - 1;

        vm.startPrank(trusted);
        assetToken.approve(address(assetMan), uint256(assets));
        vm.expectRevert(IValidatorManager.InsufficientAsset.selector);
        valMan.registerValidator(assets, 10, 5);
    }

    function test_registerValidator_largeCommissionRate_reverts() external {
        uint128 assets = minRegisterAmount;

        vm.startPrank(trusted);
        assetToken.approve(address(assetMan), uint256(assets));
        vm.expectRevert(IValidatorManager.MaxCommissionRateExceeded.selector);
        valMan.registerValidator(assets, 101, 5);
    }

    function test_registerValidator_largeCommissionMaxChangeRate_reverts() external {
        uint128 assets = minRegisterAmount;

        vm.startPrank(trusted);
        assetToken.approve(address(assetMan), uint256(assets));
        vm.expectRevert(IValidatorManager.MaxCommissionChangeRateExceeded.selector);
        valMan.registerValidator(assets, 10, 101);
    }

    function test_startValidator_succeeds() external {
        uint32 count = valMan.startedValidatorCount();

        _registerValidator(trusted, minStartAmount - 1);
        vm.startPrank(asserter);
        assetToken.approve(address(assetMan), 1);
        assetMan.delegate(trusted, 1);
        vm.stopPrank();
        assertTrue(valMan.getStatus(trusted) == IValidatorManager.ValidatorStatus.CAN_START);

        vm.prank(trusted);
        vm.expectEmit(true, false, false, true, address(valMan));
        emit ValidatorStarted(trusted, block.timestamp);
        valMan.startValidator();

        assertTrue(
            valMan.getStatus(trusted) == IValidatorManager.ValidatorStatus.CAN_SUBMIT_OUTPUT
        );
        assertEq(valMan.startedValidatorCount(), count + 1);
        assertEq(valMan.getWeight(trusted), minStartAmount);
    }

    function test_startValidator_notValidator_reverts() external {
        assertTrue(valMan.getStatus(trusted) == IValidatorManager.ValidatorStatus.NONE);

        vm.prank(trusted);
        vm.expectRevert(IValidatorManager.ImproperValidatorStatus.selector);
        valMan.startValidator();
    }

    function test_startValidator_active_reverts() external {
        _registerValidator(trusted, minStartAmount - 1);

        vm.prank(trusted);
        vm.expectRevert(IValidatorManager.ImproperValidatorStatus.selector);
        valMan.startValidator();
    }

    function test_startValidator_inactive_reverts() external {
        _registerValidator(trusted, minStartAmount);
        uint128 kroShares = assetMan.getKroTotalShareBalance(trusted, trusted);
        vm.prank(trusted);
        assetMan.initUndelegate(trusted, kroShares);

        assertTrue(valMan.getStatus(trusted) == IValidatorManager.ValidatorStatus.INACTIVE);

        vm.prank(trusted);
        vm.expectRevert(IValidatorManager.ImproperValidatorStatus.selector);
        valMan.startValidator();
    }

    function test_startValidator_alreadyStarted_reverts() external {
        _registerValidator(trusted, minStartAmount);

        vm.prank(trusted);
        vm.expectRevert(IValidatorManager.ImproperValidatorStatus.selector);
        valMan.startValidator();
    }

    function test_afterSubmitL2Output_distributeReward_succeeds() external {
        // register validator with commission rate 10%
        _registerValidator(trusted, minStartAmount);

        // submit all outputs which interact with ValidatorPool
        for (uint256 i; i <= terminateOutputIndex; i++) {
            vm.prank(trusted);
            mockOracle.addOutput(i * oracle.SUBMISSION_INTERVAL());
        }

        vm.warp(oracle.finalizedAt(terminateOutputIndex));
        mockOracle.mockSetLatestFinalizedOutputIndex(terminateOutputIndex);

        // delegate 100 KGHs
        _setUpHundredKghDelegation(trusted, 1);

        assertEq(assetMan.totalKghNum(trusted), 100);
        // 20e18 * 0.9 will be calculated as 18000000000000000000
        baseReward = 18000000000000000000;
        // 8 * arctan(0.01 * kghNum) * 1e18 * 0.9 will be calculated as 5654856240000663092
        uint128 boostedReward = 5654856240000663092;
        // 20e18 * 0.1 + 8 * arctan(0.01 * kghNum) * 1e18 * 0.1 will be calculated as 2565485624000066309
        uint128 validatorReward = 2565485624000066309;

        // submit the first output which interacts with ValidatorManager
        _submitL2Output(oracle.nextBlockNumber(), false);

        // jump to the finalization time of the first output of ValidatorManager
        vm.warp(oracle.finalizedAt(terminateOutputIndex + 1));

        // submit one more output and distribute reward
        uint256 outputIndex = oracle.nextOutputIndex();
        vm.prank(valMan.nextValidator());
        mockOracle.addOutput(oracle.nextBlockNumber());
        vm.prank(address(oracle));
        vm.expectEmit(true, false, false, true, address(valMan));
        emit RewardDistributed(trusted, validatorReward, baseReward, boostedReward);
        valMan.afterSubmitL2Output(outputIndex);

        uint128 kroReward = assetMan.totalKroAssets(trusted) - minStartAmount - 100 * VKRO_PER_KGH;
        vm.prank(trusted);
        uint128 oneKghReward = assetMan.previewKghUndelegate(trusted, 1) - VKRO_PER_KGH;

        assertEq(kroReward, baseReward);
        assertEq(oneKghReward, boostedReward / 100);
    }

    function test_afterSubmitL2Output_notUpdatePriorityValidator_succeeds() external {
        // submit all outputs which interact with ValidatorPool
        for (uint256 i; i <= terminateOutputIndex; i++) {
            vm.prank(trusted);
            mockOracle.addOutput(i * oracle.SUBMISSION_INTERVAL());
        }

        _registerValidator(trusted, minStartAmount - 1);

        assertEq(mockValMan.nextPriorityValidator(), address(0));

        // submit output
        warpToSubmitTime();
        _submitL2Output(oracle.nextBlockNumber(), false);

        assertEq(mockValMan.nextPriorityValidator(), address(0));
    }

    function test_afterSubmitL2Output_updatePriorityValidator_succeeds() external {
        // deposit funds
        _registerValidator(asserter, minStartAmount);
        _registerValidator(trusted, minStartAmount);

        // submit all outputs which interact with ValidatorPool
        for (uint256 i; i <= terminateOutputIndex; i++) {
            vm.prank(trusted);
            mockOracle.addOutput(i * oracle.SUBMISSION_INTERVAL());
        }

        vm.warp(oracle.finalizedAt(terminateOutputIndex));
        mockOracle.mockSetLatestFinalizedOutputIndex(terminateOutputIndex);

        // check if next priority validator is not set in ValidatorManager
        assertTrue(mockValMan.nextPriorityValidator() == address(0));
        assertEq(valMan.nextValidator(), trusted);

        // submit the first output which interacts with ValidatorManager
        _submitL2Output(oracle.nextBlockNumber(), false);

        // check if lastest finalized output is not updated
        assertEq(oracle.latestFinalizedOutputIndex(), terminateOutputIndex);
        // check if next priority validator is set in ValidatorManager
        address nextValidator = mockValMan.nextPriorityValidator();
        assertTrue(nextValidator != address(0));
        assertTrue(valMan.nextValidator() == nextValidator);

        // jump to the finalization time of the first output of ValidatorManager
        vm.warp(oracle.finalizedAt(terminateOutputIndex + 1));
        _submitL2Output(oracle.nextBlockNumber(), false);

        // check if lastest finalized output is updated
        assertEq(oracle.latestFinalizedOutputIndex(), terminateOutputIndex + 1);

        // submit 10 outputs
        uint256 tries = 10;
        bool changed = false;
        nextValidator = valMan.nextValidator();

        for (uint256 i; i < tries; i++) {
            // submit next output and finalize prev output
            warpToSubmitTime();
            _submitL2Output(oracle.nextBlockNumber(), false);

            // check the next validator has changed
            address newValidator = valMan.nextValidator();
            if (nextValidator != newValidator) {
                changed = true;
                break;
            }
        }

        assertTrue(changed);
    }

    function test_afterSubmitL2Output_tryJail_succeeds() public {
        // deposit funds
        _registerValidator(asserter, minStartAmount);
        _registerValidator(trusted, minStartAmount);

        for (uint256 i; i < jailThreshold; i++) {
            mockValMan.updatePriorityValidator(asserter);

            // warp to public round
            vm.warp(oracle.nextOutputMinL2Timestamp() + roundDuration + 1);
            vm.prank(trusted);
            _submitL2Output(oracle.nextBlockNumber(), true);

            assertEq(valMan.noSubmissionCount(asserter), i + 1);
            assertFalse(valMan.inJail(asserter));
        }

        mockValMan.updatePriorityValidator(asserter);

        // warp to public round
        vm.warp(oracle.nextOutputMinL2Timestamp() + roundDuration + 1);
        uint256 outputIndex = oracle.nextOutputIndex();

        vm.prank(trusted);
        mockOracle.addOutput(oracle.nextBlockNumber());
        vm.prank(address(oracle));
        vm.expectEmit(true, false, false, true, address(valMan));
        emit ValidatorJailed(asserter, uint128(block.timestamp) + jailPeriodSeconds);
        valMan.afterSubmitL2Output(outputIndex);

        assertEq(valMan.noSubmissionCount(asserter), jailThreshold);
        assertTrue(valMan.inJail(asserter));
    }

    function test_afterSubmitL2Output_resetNoSubmissionCount_succeeds() external {
        // deposit funds
        _registerValidator(asserter, minStartAmount);
        _registerValidator(trusted, minStartAmount);

        mockValMan.updatePriorityValidator(asserter);

        // warp to public round
        vm.warp(oracle.nextOutputMinL2Timestamp() + roundDuration + 1);
        assertEq(valMan.nextValidator(), Constants.VALIDATOR_PUBLIC_ROUND_ADDRESS);
        vm.prank(trusted);
        _submitL2Output(oracle.nextBlockNumber(), true);

        assertEq(valMan.noSubmissionCount(asserter), 1);
        assertFalse(valMan.inJail(asserter));

        mockValMan.updatePriorityValidator(asserter);

        // warp to priority round
        warpToSubmitTime();
        _submitL2Output(oracle.nextBlockNumber(), false);

        assertEq(valMan.noSubmissionCount(asserter), 0);
    }

    function test_afterSubmitL2Output_senderNotL2OO_reverts() external {
        vm.prank(trusted);
        vm.expectRevert(IValidatorManager.NotAllowedCaller.selector);
        valMan.afterSubmitL2Output(0);
    }

    function test_changeCommissionRate_succeeds() public {
        _registerValidator(asserter, minStartAmount);

        uint8 commissionRate = valMan.getCommissionRate(asserter);
        uint8 commissionMaxChangeRate = valMan.getCommissionMaxChangeRate(asserter);
        uint8 newCommissionRate = commissionRate + commissionMaxChangeRate;

        vm.warp(
            mockValMan.commissionRateChangedAt(asserter) +
                valMan.COMMISSION_RATE_MIN_CHANGE_SECONDS()
        );
        vm.prank(asserter);
        vm.expectEmit(false, false, false, true, address(valMan));
        emit ValidatorCommissionRateChanged(asserter, commissionRate, newCommissionRate);
        valMan.changeCommissionRate(newCommissionRate);

        assertEq(valMan.getCommissionRate(asserter), newCommissionRate);
    }

    function test_changeCommissionRate_twice_succeeds() external {
        test_changeCommissionRate_succeeds();

        uint8 commissionRate = valMan.getCommissionRate(asserter);
        uint8 commissionMaxChangeRate = valMan.getCommissionMaxChangeRate(asserter);
        uint8 newCommissionRate = commissionRate - commissionMaxChangeRate;

        vm.warp(
            mockValMan.commissionRateChangedAt(asserter) +
                valMan.COMMISSION_RATE_MIN_CHANGE_SECONDS()
        );
        vm.prank(asserter);
        valMan.changeCommissionRate(newCommissionRate);

        assertEq(valMan.getCommissionRate(asserter), newCommissionRate);
    }

    function test_changeCommissionRate_inactive_reverts() external {
        _registerValidator(trusted, minStartAmount);
        uint128 kroShares = assetMan.getKroTotalShareBalance(trusted, trusted);
        vm.prank(trusted);
        assetMan.initUndelegate(trusted, kroShares);

        assertTrue(valMan.getStatus(trusted) == IValidatorManager.ValidatorStatus.INACTIVE);

        vm.prank(asserter);
        vm.expectRevert(IValidatorManager.ImproperValidatorStatus.selector);
        valMan.changeCommissionRate(15);
    }

    function test_changeCommissionRate_minChangeSecNotElapsed_reverts() external {
        _registerValidator(asserter, minStartAmount);

        vm.prank(asserter);
        vm.expectRevert(IValidatorManager.NotElapsedCommissionChangePeriod.selector);
        valMan.changeCommissionRate(15);
    }

    function test_changeCommissionRate_largeCommissionRate_reverts() external {
        _registerValidator(asserter, minStartAmount);

        vm.warp(
            mockValMan.commissionRateChangedAt(asserter) +
                valMan.COMMISSION_RATE_MIN_CHANGE_SECONDS()
        );
        vm.prank(asserter);
        vm.expectRevert(IValidatorManager.MaxCommissionRateExceeded.selector);
        valMan.changeCommissionRate(101);
    }

    function test_changeCommissionRate_sameCommissionRate_reverts() external {
        _registerValidator(asserter, minStartAmount);

        uint8 commissionRate = valMan.getCommissionRate(asserter);

        vm.warp(
            mockValMan.commissionRateChangedAt(asserter) +
                valMan.COMMISSION_RATE_MIN_CHANGE_SECONDS()
        );
        vm.prank(asserter);
        vm.expectRevert(IValidatorManager.SameCommissionRate.selector);
        valMan.changeCommissionRate(commissionRate);
    }

    function test_changeCommissionRate_largeChangeRate_reverts() external {
        _registerValidator(asserter, minStartAmount);

        uint8 commissionRate = valMan.getCommissionRate(asserter);
        uint8 commissionMaxChangeRate = valMan.getCommissionMaxChangeRate(asserter);
        uint8 newCommissionRate = commissionRate + commissionMaxChangeRate + 1;

        vm.warp(
            mockValMan.commissionRateChangedAt(asserter) +
                valMan.COMMISSION_RATE_MIN_CHANGE_SECONDS()
        );
        vm.prank(asserter);
        vm.expectRevert(IValidatorManager.CommissionChangeRateExceeded.selector);
        valMan.changeCommissionRate(newCommissionRate);
    }

    function test_tryUnjail_succeeds() external {
        test_afterSubmitL2Output_tryJail_succeeds();

        vm.warp(valMan.jailExpiresAt(asserter));
        vm.prank(asserter);
        vm.expectEmit(false, false, false, true, address(valMan));
        emit ValidatorUnjailed(asserter);
        valMan.tryUnjail(asserter, false);

        assertEq(valMan.noSubmissionCount(asserter), 0);
        assertFalse(valMan.inJail(asserter));
    }

    function test_tryUnjail_notInJail_reverts() external {
        vm.prank(asserter);
        vm.expectRevert(IValidatorManager.ImproperValidatorStatus.selector);
        valMan.tryUnjail(asserter, false);
    }

    function test_tryUnjail_periodNotElapsed_reverts() external {
        test_afterSubmitL2Output_tryJail_succeeds();

        vm.prank(asserter);
        vm.expectRevert(IValidatorManager.NotElapsedJailPeriod.selector);
        valMan.tryUnjail(asserter, false);
    }

    function test_slash_succeeds() external {
        uint32 count = valMan.startedValidatorCount();
        // deposit funds
        _registerValidator(asserter, minStartAmount);
        _registerValidator(challenger, minStartAmount);
        assertEq(valMan.startedValidatorCount(), count + 2);

        // submit all outputs which interact with ValidatorPool
        for (uint256 i; i <= terminateOutputIndex; i++) {
            vm.prank(trusted);
            mockOracle.addOutput(i * oracle.SUBMISSION_INTERVAL());
        }

        vm.warp(oracle.finalizedAt(terminateOutputIndex));
        mockOracle.mockSetLatestFinalizedOutputIndex(terminateOutputIndex);

        // delegate KGHs
        _setUpHundredKghDelegation(asserter, 1);
        _setUpHundredKghDelegation(challenger, 101);
        assertEq(assetMan.totalKghNum(asserter), 100);
        assertEq(assetMan.totalKghNum(challenger), 100);

        // submit the first output which interacts with ValidatorManager
        mockValMan.updatePriorityValidator(asserter);
        _submitL2Output(oracle.nextBlockNumber(), false);

        vm.startPrank(address(colosseum));
        // suppose that the challenge is successful, so the winner is challenger
        valMan.slash(oracle.latestOutputIndex(), challenger, asserter);
        vm.stopPrank();
        // this will be done by the l2 output oracle contract in the real environment
        vm.startPrank(challenger);
        mockOracle.replaceOutput(oracle.latestOutputIndex());
        vm.stopPrank();

        // jump to the finalization time of the challenged output
        vm.warp(oracle.finalizedAt(oracle.latestOutputIndex()));
        // submit one more output to distribute reward
        _submitL2Output(oracle.nextBlockNumber(), false);

        // Total slashingAmount is 2e18.
        uint128 asserterSlashedAmount = minStartAmount +
            100 *
            VKRO_PER_KGH -
            assetMan.totalKroAssets(asserter);
        assertEq(asserterSlashedAmount, 2e18);

        // KRO reward by slashing for challenger is calculated as
        // 16e17 * ((totalKro - totalKroInKgh) / (totalKro - totalKroInKgh + boostedReward)),
        // which is 1495796931079677248, with tax taken by security council.
        // Adding this to the original reward 18e18 is 19495796931079677248.
        uint128 challengerKroAmount = assetMan.totalKroAssets(challenger) -
            minStartAmount -
            100 *
            VKRO_PER_KGH;
        assertEq(challengerKroAmount, 19495796931079677248);

        // Asserter KGH reward should be 0.
        vm.prank(asserter);
        uint128 oneKghRewardForAsserter = assetMan.previewKghUndelegate(asserter, 1) - VKRO_PER_KGH;
        assertEq(oneKghRewardForAsserter, 0);

        // Challenger KGH reward should be 16e17 * (boostedReward / (totalKro - totalKroInKgh + boostedReward)),
        // which is 716823441482183, with tax taken by security council and validator commission.
        // Adding this to the original reward 5654856240000663092 is 57265385841488813.
        vm.prank(challenger);
        uint128 oneKghRewardForChallenger = assetMan.previewKghUndelegate(challenger, 101) -
            VKRO_PER_KGH;
        assertEq(oneKghRewardForChallenger, 57265385841488813);

        assertEq(assetMan.ASSET_TOKEN().balanceOf(guardian), 4e17);
        assertEq(valMan.startedValidatorCount(), count + 1);
    }

    function test_slash_notColosseum_reverts() external {
        vm.prank(address(1));
        vm.expectRevert(IValidatorManager.NotAllowedCaller.selector);
        valMan.slash(1, challenger, asserter);
    }

    function test_checkSubmissionEligibility_priorityRound_succeeds() external {
        address nextValidator = valMan.nextValidator();
        _registerValidator(nextValidator, minStartAmount);

        vm.prank(address(oracle));
        valMan.checkSubmissionEligibility(nextValidator);
    }

    function test_checkSubmissionEligibility_publicRound_succeeds() external {
        mockValMan.updatePriorityValidator(asserter);

        _registerValidator(trusted, minStartAmount);

        // warp to public round
        vm.warp(oracle.nextOutputMinL2Timestamp() + roundDuration + 1);
        vm.prank(address(oracle));
        valMan.checkSubmissionEligibility(trusted);
    }

    function test_checkSubmissionEligibility_senderNotL2OO_reverts() external {
        vm.expectRevert(IValidatorManager.NotAllowedCaller.selector);
        valMan.checkSubmissionEligibility(trusted);
    }

    function test_checkSubmissionEligibility_notSelected_reverts() external {
        mockValMan.updatePriorityValidator(asserter);

        vm.prank(address(oracle));
        vm.expectRevert(IValidatorManager.NotSelectedPriorityValidator.selector);
        valMan.checkSubmissionEligibility(trusted);
    }

    function test_checkSubmissionEligibility_notSatisfyCondition_reverts() external {
        address nextValidator = valMan.nextValidator();

        vm.prank(address(oracle));
        vm.expectRevert(IValidatorManager.ImproperValidatorStatus.selector);
        valMan.checkSubmissionEligibility(nextValidator);
    }

    function test_checkSubmissionEligibility_publicRound_notSatisfyCondition_reverts() external {
        mockValMan.updatePriorityValidator(asserter);

        // warp to public round
        vm.warp(oracle.nextOutputMinL2Timestamp() + roundDuration + 1);
        vm.prank(address(oracle));
        vm.expectRevert(IValidatorManager.ImproperValidatorStatus.selector);
        valMan.checkSubmissionEligibility(trusted);
    }

    function test_getStatus_active_succeeds() external {
        _registerValidator(trusted, minStartAmount);
        assertEq(valMan.getWeight(trusted), minStartAmount);

        uint128 minUndelegateShares = assetMan.previewDelegate(trusted, 1);
        vm.prank(trusted);
        assetMan.initUndelegate(trusted, minUndelegateShares);
        assertTrue(valMan.getStatus(trusted) == IValidatorManager.ValidatorStatus.ACTIVE);
        assertEq(valMan.getWeight(trusted), 0);
    }

    function test_startedValidatorTotalWeight_succeeds() external {
        uint32 count = valMan.startedValidatorCount();
        uint120 totalWeight = valMan.startedValidatorTotalWeight();
        _registerValidator(trusted, minStartAmount);
        assertEq(valMan.startedValidatorCount(), count + 1);
        assertEq(valMan.startedValidatorTotalWeight(), totalWeight + minStartAmount);

        count = valMan.startedValidatorCount();
        totalWeight = valMan.startedValidatorTotalWeight();
        _registerValidator(asserter, minStartAmount);
        assertEq(valMan.startedValidatorCount(), count + 1);
        assertEq(valMan.startedValidatorTotalWeight(), totalWeight + minStartAmount);

        count = valMan.startedValidatorCount();
        totalWeight = valMan.startedValidatorTotalWeight();
        vm.startPrank(challenger);
        assetToken.approve(address(assetMan), 10);
        assetMan.delegate(asserter, 10);
        vm.stopPrank();
        assertEq(valMan.startedValidatorCount(), count);
        assertEq(valMan.startedValidatorTotalWeight(), totalWeight + 10);
    }
}
