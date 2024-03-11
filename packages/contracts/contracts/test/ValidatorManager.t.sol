// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { Constants } from "../libraries/Constants.sol";
import { Types } from "../libraries/Types.sol";
import { Proxy } from "../universal/Proxy.sol";
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
    constructor(
        ConstructorParams memory _constructorParams,
        address _trustedValidator,
        uint128 _commissionRateMinChangeSeconds,
        uint128 _roundDurationSeconds,
        uint128 _jailPeriodSeconds,
        uint128 _jailThreshold
    )
        ValidatorManager(
            _constructorParams,
            _trustedValidator,
            _commissionRateMinChangeSeconds,
            _roundDurationSeconds,
            _jailPeriodSeconds,
            _jailThreshold
        )
    {}

    function updatePriorityValidator(address validator) external {
        _nextPriorityValidator = validator;
    }

    function nextPriorityValidator() external view returns (address) {
        return _nextPriorityValidator;
    }

    function commissionRateChangedAt(address validator) external view returns (uint128) {
        return _vaults[validator].reward.commissionRateChangedAt;
    }
}

contract ValidatorManagerTest is L2OutputOracle_ValidatorSystemUpgrade_Initializer {
    MockL2OutputOracle mockOracle;
    MockValidatorManager mockValMan;

    event ValidatorRegistered(
        address indexed validator,
        bool indexed started,
        uint8 commissionRate,
        uint8 commissionMaxChangeRate,
        uint128 assets
    );

    event ValidatorStarted(address indexed validator, uint256 startsAt);

    event ValidatorCommissionRateChanged(
        address validator,
        uint8 oldCommissionRate,
        uint8 newCommissionRate
    );

    event ValidatorJailed(address indexed validator, uint128 expiresAt);

    event ValidatorUnjailed(address validator);

    function _submitL2Output(uint256 l2BlockNumber, bool isPublicRound) private {
        uint256 outputIndex = oracle.nextOutputIndex();
        if (!isPublicRound) {
            vm.prank(valMan.nextValidator());
        }
        mockOracle.addOutput(l2BlockNumber);
        vm.prank(address(oracle));
        valMan.afterSubmitL2Output(outputIndex);
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
        MockValidatorManager mockValManImpl = new MockValidatorManager(
            constructorParams,
            trusted,
            commissionRateMinChangeSeconds,
            uint128(roundDuration),
            jailPeriodSeconds,
            jailThreshold
        );
        vm.prank(multisig);
        Proxy(payable(valManAddress)).upgradeTo(address(mockValManImpl));
        mockValMan = MockValidatorManager(valManAddress);
    }

    function test_constructor_succeeds() external {
        assertEq(valMan.TRUSTED_VALIDATOR(), trusted);
        assertEq(valMan.COMMISSION_RATE_MIN_CHANGE_SECONDS(), commissionRateMinChangeSeconds);
        assertEq(valMan.ROUND_DURATION_SECONDS(), roundDuration);
        assertEq(valMan.JAIL_PERIOD_SECONDS(), jailPeriodSeconds);
        assertEq(valMan.JAIL_THRESHOLD(), jailThreshold);
    }

    function test_registerValidator_canSubmitOutput_succeeds() external {
        uint256 trustedBalance = assetToken.balanceOf(trusted);
        uint32 count = valMan.startedValidatorCount();

        uint128 assets = minStartAmount;
        uint8 commissionRate = 10;
        uint8 commissionMaxChangeRate = 5;

        vm.startPrank(trusted);
        assetToken.approve(address(valMan), uint256(assets));
        vm.expectEmit(true, true, false, true, address(valMan));
        emit ValidatorRegistered(trusted, true, commissionRate, commissionMaxChangeRate, assets);
        valMan.registerValidator(assets, commissionRate, commissionMaxChangeRate);
        vm.stopPrank();

        assertEq(assetToken.balanceOf(trusted), trustedBalance - assets);
        assertEq(valMan.totalKroAssets(trusted), assets);
        assertEq(valMan.getCommissionRate(trusted), commissionRate);
        assertEq(valMan.getCommissionMaxChangeRate(trusted), commissionMaxChangeRate);
        assertEq(mockValMan.commissionRateChangedAt(trusted), block.timestamp);

        assertTrue(valMan.getStatus(trusted) == ValidatorManager.ValidatorStatus.CAN_SUBMIT_OUTPUT);
        assertEq(valMan.startedValidatorCount(), count + 1);
        assertEq(valMan.getWeight(trusted), assets);
    }

    function test_registerValidator_active_succeeds() external {
        uint32 count = valMan.startedValidatorCount();

        uint128 assets = minStartAmount - 1;
        uint8 commissionRate = 10;
        uint8 commissionMaxChangeRate = 5;

        vm.startPrank(trusted);
        assetToken.approve(address(valMan), uint256(assets));
        vm.expectEmit(true, true, false, true, address(valMan));
        emit ValidatorRegistered(trusted, false, commissionRate, commissionMaxChangeRate, assets);
        valMan.registerValidator(assets, commissionRate, commissionMaxChangeRate);
        vm.stopPrank();

        assertTrue(valMan.getStatus(trusted) == ValidatorManager.ValidatorStatus.ACTIVE);
        assertEq(valMan.startedValidatorCount(), count);
        assertEq(valMan.getWeight(trusted), 0);
    }

    function test_registerValidator_alreadyInitiated_reverts() external {
        uint128 assets = minStartAmount;

        _registerValidator(trusted, assets);

        vm.startPrank(trusted);
        assetToken.approve(address(valMan), uint256(assets));
        vm.expectRevert("ValidatorManager: already initiated validator");
        valMan.registerValidator(assets, 10, 5);
    }

    function test_registerValidator_smallAsset_reverts() external {
        uint128 assets = minRegisterAmount - 1;

        vm.startPrank(trusted);
        assetToken.approve(address(valMan), uint256(assets));
        vm.expectRevert("ValidatorManager: need to register with at least min register amount");
        valMan.registerValidator(assets, 10, 5);
    }

    function test_registerValidator_largeCommissionRate_reverts() external {
        uint128 assets = minRegisterAmount;

        vm.startPrank(trusted);
        assetToken.approve(address(valMan), uint256(assets));
        vm.expectRevert("ValidatorManager: the max value of commission rate has been exceeded");
        valMan.registerValidator(assets, 101, 5);
    }

    function test_registerValidator_largeCommissionMaxChangeRate_reverts() external {
        uint128 assets = minRegisterAmount;

        vm.startPrank(trusted);
        assetToken.approve(address(valMan), uint256(assets));
        vm.expectRevert(
            "ValidatorManager: the max value of commission rate max change rate has been exceeded"
        );
        valMan.registerValidator(assets, 10, 101);
    }

    function test_startValidator_succeeds() external {
        uint32 count = valMan.startedValidatorCount();

        _registerValidator(trusted, minStartAmount - 1);
        vm.startPrank(asserter);
        assetToken.approve(address(valMan), 1);
        valMan.delegate(trusted, 1);
        vm.stopPrank();
        assertTrue(valMan.getStatus(trusted) == ValidatorManager.ValidatorStatus.CAN_START);

        vm.prank(trusted);
        vm.expectEmit(true, false, false, true, address(valMan));
        emit ValidatorStarted(trusted, block.timestamp);
        valMan.startValidator();

        assertTrue(valMan.getStatus(trusted) == ValidatorManager.ValidatorStatus.CAN_SUBMIT_OUTPUT);
        assertEq(valMan.startedValidatorCount(), count + 1);
        assertEq(valMan.getWeight(trusted), minStartAmount);
    }

    function test_startValidator_notValidator_reverts() external {
        assertTrue(valMan.getStatus(trusted) == ValidatorManager.ValidatorStatus.NONE);

        vm.prank(trusted);
        vm.expectRevert("ValidatorManager: validator start condition is not met");
        valMan.startValidator();
    }

    function test_startValidator_active_reverts() external {
        _registerValidator(trusted, minStartAmount - 1);

        vm.prank(trusted);
        vm.expectRevert("ValidatorManager: validator start condition is not met");
        valMan.startValidator();
    }

    function test_startValidator_inactive_reverts() external {
        _registerValidator(trusted, minStartAmount);
        uint128 kroShares = valMan.getKroTotalShareBalance(trusted, trusted);
        vm.prank(trusted);
        valMan.initUndelegate(trusted, kroShares);

        assertTrue(valMan.getStatus(trusted) == ValidatorManager.ValidatorStatus.INACTIVE);

        vm.prank(trusted);
        vm.expectRevert("ValidatorManager: validator start condition is not met");
        valMan.startValidator();
    }

    function test_startValidator_alreadyStarted_reverts() external {
        _registerValidator(trusted, minStartAmount);

        vm.prank(trusted);
        vm.expectRevert("ValidatorManager: validator start condition is not met");
        valMan.startValidator();
    }

    function test_afterSubmitL2Output_notUpdatePriorityValidator_succeeds() external {
        _registerValidator(trusted, minStartAmount);

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
        for (uint256 i; i <= poolLastOutputIndex; i++) {
            vm.prank(trusted);
            mockOracle.addOutput(i * oracle.SUBMISSION_INTERVAL());
        }

        vm.warp(oracle.finalizedAt(poolLastOutputIndex));
        mockOracle.mockSetLatestFinalizedOutputIndex(poolLastOutputIndex);

        // check if next priority validator is not set in ValidatorManager
        assertTrue(mockValMan.nextPriorityValidator() == address(0));
        assertEq(valMan.nextValidator(), trusted);

        // submit the first output which interacts with ValidatorManager
        _submitL2Output(oracle.nextBlockNumber(), false);

        // check if lastest finalized output is not updated
        assertEq(oracle.latestFinalizedOutputIndex(), poolLastOutputIndex);
        // check if next priority validator is set in ValidatorManager
        address nextValidator = mockValMan.nextPriorityValidator();
        assertTrue(nextValidator != address(0));
        assertTrue(valMan.nextValidator() == nextValidator);

        // jump to the finalization time of the first output of ValidatorManager
        vm.warp(oracle.finalizedAt(poolLastOutputIndex + 1));
        _submitL2Output(oracle.nextBlockNumber(), false);

        // check if lastest finalized output is updated
        assertEq(oracle.latestFinalizedOutputIndex(), poolLastOutputIndex + 1);

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
            assertFalse(valMan.getStatus(asserter) == ValidatorManager.ValidatorStatus.IN_JAIL);
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
        assertTrue(valMan.getStatus(asserter) == ValidatorManager.ValidatorStatus.IN_JAIL);
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
        assertFalse(valMan.getStatus(asserter) == ValidatorManager.ValidatorStatus.IN_JAIL);

        mockValMan.updatePriorityValidator(asserter);

        // warp to priority round
        warpToSubmitTime();
        _submitL2Output(oracle.nextBlockNumber(), false);

        assertEq(valMan.noSubmissionCount(asserter), 0);
    }

    function test_afterSubmitL2Output_senderNotL2OO_reverts() external {
        vm.prank(trusted);
        vm.expectRevert("ValidatorManager: Only L2OutputOracle can call this function");
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
        uint128 kroShares = valMan.getKroTotalShareBalance(trusted, trusted);
        vm.prank(trusted);
        valMan.initUndelegate(trusted, kroShares);

        assertTrue(valMan.getStatus(trusted) == ValidatorManager.ValidatorStatus.INACTIVE);

        vm.prank(asserter);
        vm.expectRevert("ValidatorManager: cannot change commission rate of inactive validator");
        valMan.changeCommissionRate(15);
    }

    function test_changeCommissionRate_minChangeSecNotElapsed_reverts() external {
        _registerValidator(asserter, minStartAmount);

        vm.prank(asserter);
        vm.expectRevert("ValidatorManager: min change seconds of commission rate has not elapsed");
        valMan.changeCommissionRate(15);
    }

    function test_changeCommissionRate_largeCommissionRate_reverts() external {
        _registerValidator(asserter, minStartAmount);

        vm.warp(
            mockValMan.commissionRateChangedAt(asserter) +
                valMan.COMMISSION_RATE_MIN_CHANGE_SECONDS()
        );
        vm.prank(asserter);
        vm.expectRevert("ValidatorManager: the max value of commission rate has been exceeded");
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
        vm.expectRevert("ValidatorManager: cannot change to the same value");
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
        vm.expectRevert("ValidatorManager: max change rate of commission rate has been exceeded");
        valMan.changeCommissionRate(newCommissionRate);
    }

    function test_tryUnjail_succeeds() external {
        test_afterSubmitL2Output_tryJail_succeeds();

        vm.warp(valMan.jailExpiresAt(asserter));
        vm.prank(asserter);
        vm.expectEmit(false, false, false, true, address(valMan));
        emit ValidatorUnjailed(asserter);
        valMan.tryUnjail();

        assertEq(valMan.noSubmissionCount(asserter), 0);
        assertFalse(valMan.getStatus(asserter) == ValidatorManager.ValidatorStatus.IN_JAIL);
    }

    function test_tryUnjail_notInJail_reverts() external {
        vm.prank(asserter);
        vm.expectRevert("ValidatorManager: not in jail");
        valMan.tryUnjail();
    }

    function test_tryUnjail_periodNotElapsed_reverts() external {
        test_afterSubmitL2Output_tryJail_succeeds();

        vm.prank(asserter);
        vm.expectRevert("ValidatorManager: jail period has not elasped");
        valMan.tryUnjail();
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
        vm.expectRevert("ValidatorManager: Only L2OutputOracle can call this function");
        valMan.checkSubmissionEligibility(trusted);
    }

    function test_checkSubmissionEligibility_notSelected_reverts() external {
        mockValMan.updatePriorityValidator(asserter);

        vm.prank(address(oracle));
        vm.expectRevert("ValidatorManager: only the next selected validator can submit output");
        valMan.checkSubmissionEligibility(trusted);
    }

    function test_checkSubmissionEligibility_notSatisfyCondition_reverts() external {
        address nextValidator = valMan.nextValidator();

        vm.prank(address(oracle));
        vm.expectRevert(
            "ValidatorManager: validator should satisfy the condition to submit output"
        );
        valMan.checkSubmissionEligibility(nextValidator);
    }

    function test_checkSubmissionEligibility_publicRound_notSatisfyCondition_reverts() external {
        mockValMan.updatePriorityValidator(asserter);

        // warp to public round
        vm.warp(oracle.nextOutputMinL2Timestamp() + roundDuration + 1);
        vm.prank(address(oracle));
        vm.expectRevert(
            "ValidatorManager: validator should satisfy the condition to submit output"
        );
        valMan.checkSubmissionEligibility(trusted);
    }

    function test_getStatus_active_succeeds() external {
        _registerValidator(trusted, minStartAmount);
        uint128 minUndelegateShares = valMan.previewDelegate(trusted, 1);
        vm.prank(trusted);
        valMan.initUndelegate(trusted, minUndelegateShares);
        assertTrue(valMan.getStatus(trusted) == ValidatorManager.ValidatorStatus.ACTIVE);
    }

    function startedValidatorTotalWeight_succeeds() external {
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
        assetToken.approve(address(valMan), 10);
        valMan.delegate(asserter, 10);
        vm.stopPrank();
        assertEq(valMan.startedValidatorCount(), count);
        assertEq(valMan.startedValidatorTotalWeight(), totalWeight + 10);
    }
}
