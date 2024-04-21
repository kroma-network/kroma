// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { Constants } from "../libraries/Constants.sol";
import { Types } from "../libraries/Types.sol";
import { Proxy } from "../universal/Proxy.sol";
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

contract ValidatorManagerTest is ValidatorSystemUpgrade_Initializer {
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

    event ValidatorStopped(address indexed validator, uint256 stopsAt);

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

    event Slashed(uint256 indexed outputIndex, address indexed loser, uint128 amount);

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

        // Submit until terminateOutputIndex and set it latest finalized output
        vm.prank(trusted);
        pool.deposit{ value: trusted.balance }();
        for (uint256 i = oracle.nextOutputIndex(); i <= terminateOutputIndex; i++) {
            _submitL2OutputV1();
        }
        vm.warp(oracle.finalizedAt(terminateOutputIndex));
        mockOracle.mockSetLatestFinalizedOutputIndex(terminateOutputIndex);
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

    function test_startValidator_inJail_reverts() external {
        test_afterSubmitL2Output_tryJail_succeeds();

        // Undelegate all assets of jailed validator
        uint128 kroShares = assetMan.getKroTotalShareBalance(asserter, asserter);
        vm.prank(asserter);
        vm.expectEmit(true, false, false, true, address(valMan));
        emit ValidatorStopped(asserter, block.timestamp);
        assetMan.initUndelegate(asserter, kroShares);
        assertTrue(valMan.getStatus(asserter) == IValidatorManager.ValidatorStatus.INACTIVE);

        // Delegate to re-start validator
        vm.startPrank(asserter);
        assetToken.approve(address(assetMan), minStartAmount);
        assetMan.delegate(asserter, minStartAmount);
        vm.stopPrank();
        assertTrue(valMan.getStatus(asserter) == IValidatorManager.ValidatorStatus.CAN_START);

        vm.prank(asserter);
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
        // Register validator with commission rate 10%
        _registerValidator(trusted, minStartAmount);

        // Delegate 100 KGHs
        uint128 kghCounts = 100;
        _setUpHundredKghDelegation(trusted, 1);
        assertEq(assetMan.totalKghNum(trusted), kghCounts);

        // Submit the first output which interacts with ValidatorManager
        _submitL2OutputV2(false);

        // Jump to the finalization time of the first output of ValidatorManager
        vm.warp(oracle.finalizedAt(terminateOutputIndex + 1));

        // Boosted reward with 100 kgh delegation
        uint128 boostedReward = 6283173600000736769;
        uint128 validatorReward = (baseReward * 10) / 100 + (boostedReward * 10) / 100;
        baseReward = (baseReward * 90) / 100;
        boostedReward = (boostedReward * 90) / 100;

        // Submit one more output and distribute reward
        uint256 outputIndex = oracle.nextOutputIndex();
        vm.prank(valMan.nextValidator());
        mockOracle.addOutput(oracle.nextBlockNumber());
        vm.prank(address(oracle));
        vm.expectEmit(true, false, false, true, address(valMan));
        emit RewardDistributed(trusted, validatorReward, baseReward, boostedReward);
        valMan.afterSubmitL2Output(outputIndex);

        uint128 kroReward = assetMan.totalKroAssets(trusted) -
            minStartAmount -
            kghCounts *
            VKRO_PER_KGH;
        vm.prank(trusted);
        uint128 oneKghReward = assetMan.previewKghUndelegate(trusted, 1) - VKRO_PER_KGH;

        assertEq(kroReward, baseReward);
        assertEq(oneKghReward, boostedReward / kghCounts);

        // Check validator tree updated with rewards
        assertEq(
            valMan.getWeight(trusted),
            minStartAmount +
                kghManager.totalKroInKgh(1) *
                kghCounts +
                baseReward +
                boostedReward +
                validatorReward
        );

        assertEq(oracle.latestFinalizedOutputIndex(), terminateOutputIndex + 1);
    }

    function test_afterSubmitL2Output_updatePriorityValidator_succeeds() external {
        // Register as a validator
        _registerValidator(asserter, minStartAmount);
        _registerValidator(trusted, minStartAmount);

        // Check if next priority validator is not set in ValidatorManager
        assertTrue(mockValMan.nextPriorityValidator() == address(0));
        assertEq(valMan.nextValidator(), trusted);

        // Submit the first output which interacts with ValidatorManager
        _submitL2OutputV2(false);

        // Check if lastest finalized output is not updated
        assertEq(oracle.latestFinalizedOutputIndex(), terminateOutputIndex);
        // Check if next priority validator is set in ValidatorManager
        address nextValidator = mockValMan.nextPriorityValidator();
        assertTrue(nextValidator != address(0));
        assertTrue(valMan.nextValidator() == nextValidator);

        // Jump to the finalization time of the first output of ValidatorManager
        vm.warp(oracle.finalizedAt(terminateOutputIndex + 1));
        vm.startPrank(nextValidator);
        _submitL2OutputV2(true);
        vm.stopPrank();

        // Check if lastest finalized output is updated
        assertEq(oracle.latestFinalizedOutputIndex(), terminateOutputIndex + 1);

        // Submit 10 outputs
        uint256 tries = 10;
        bool changed = false;
        nextValidator = valMan.nextValidator();

        for (uint256 i; i < tries; i++) {
            // Submit next output and finalize prev output
            warpToSubmitTime();
            _submitL2OutputV2(false);

            // Check the next validator has changed
            address newValidator = valMan.nextValidator();
            if (nextValidator != newValidator) {
                changed = true;
                break;
            }
        }

        assertTrue(changed);
    }

    function test_afterSubmitL2Output_tryJail_succeeds() public {
        // Register as a validator
        _registerValidator(asserter, minStartAmount);
        _registerValidator(trusted, minStartAmount);

        vm.startPrank(trusted);
        for (uint256 i; i < jailThreshold; i++) {
            mockValMan.updatePriorityValidator(asserter);

            // Warp to public round
            vm.warp(oracle.nextOutputMinL2Timestamp() + roundDuration + 1);
            _submitL2OutputV2(true);

            assertEq(valMan.noSubmissionCount(asserter), i + 1);
            assertFalse(valMan.inJail(asserter));
        }
        vm.stopPrank();

        mockValMan.updatePriorityValidator(asserter);

        // Warp to public round
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
        // Register as a validator
        _registerValidator(asserter, minStartAmount);
        _registerValidator(trusted, minStartAmount);

        mockValMan.updatePriorityValidator(asserter);

        // Warp to public round
        vm.warp(oracle.nextOutputMinL2Timestamp() + roundDuration + 1);
        assertEq(valMan.nextValidator(), Constants.VALIDATOR_PUBLIC_ROUND_ADDRESS);
        vm.startPrank(trusted);
        _submitL2OutputV2(true);
        vm.stopPrank();

        assertEq(valMan.noSubmissionCount(asserter), 1);
        assertFalse(valMan.inJail(asserter));

        mockValMan.updatePriorityValidator(asserter);

        // Warp to priority round
        warpToSubmitTime();
        _submitL2OutputV2(false);

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

    function test_changeCommissionRate_inJail_reverts() external {
        test_afterSubmitL2Output_tryJail_succeeds();

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

    function test_tryUnjail_senderNotSelf_reverts() external {
        test_afterSubmitL2Output_tryJail_succeeds();

        vm.prank(trusted);
        vm.expectRevert(IValidatorManager.NotAllowedCaller.selector);
        valMan.tryUnjail(asserter, false);
    }

    function test_tryUnjail_force_senderNotColosseum_reverts() external {
        test_afterSubmitL2Output_tryJail_succeeds();

        vm.prank(asserter);
        vm.expectRevert(IValidatorManager.NotAllowedCaller.selector);
        valMan.tryUnjail(asserter, true);
    }

    function test_tryUnjail_periodNotElapsed_reverts() external {
        test_afterSubmitL2Output_tryJail_succeeds();

        vm.prank(asserter);
        vm.expectRevert(IValidatorManager.NotElapsedJailPeriod.selector);
        valMan.tryUnjail(asserter, false);
    }

    function test_slash_succeeds() external {
        uint32 count = valMan.startedValidatorCount();
        // Register as a validator
        _registerValidator(asserter, minStartAmount);
        _registerValidator(challenger, minStartAmount);
        assertEq(valMan.startedValidatorCount(), count + 2);

        // Delegate KGHs
        uint128 kghCounts = 100;
        _setUpHundredKghDelegation(asserter, 1);
        _setUpHundredKghDelegation(challenger, 1 + kghCounts);
        assertEq(assetMan.totalKghNum(asserter), kghCounts);
        assertEq(assetMan.totalKghNum(challenger), kghCounts);

        // Submit the first output which interacts with ValidatorManager
        mockValMan.updatePriorityValidator(asserter);
        warpToSubmitTime();
        _submitL2OutputV2(false);
        uint256 challengedOutputIndex = oracle.latestOutputIndex();

        // Suppose that the challenge is successful, so the winner is challenger
        uint128 slashingAmount = (minStartAmount * slashingRate) / assetMan.SLASHING_RATE_DENOM();
        vm.prank(address(colosseum));
        vm.expectEmit(true, true, false, true, address(valMan));
        emit Slashed(challengedOutputIndex, asserter, slashingAmount);
        valMan.slash(challengedOutputIndex, challenger, asserter);

        // Asserter in jail after slashed
        assertTrue(valMan.inJail(asserter));

        // This will be done by the l2 output oracle contract in the real environment
        vm.prank(challenger);
        mockOracle.replaceOutput(challengedOutputIndex);

        // Jump to the finalization time of the challenged output
        vm.warp(oracle.finalizedAt(challengedOutputIndex));
        vm.startPrank(challenger);
        // Submit one more output to distribute reward
        _submitL2OutputV2(true);
        vm.stopPrank();

        // Asserter asset decreased by slashingAmount
        uint128 asserterTotalKro = assetMan.totalKroAssets(asserter) -
            kghCounts *
            kghManager.totalKroInKgh(1);
        assertEq(asserterTotalKro, minStartAmount - slashingAmount);
        // Asserter has 0 rewards
        assertEq(assetMan.reflectiveWeight(asserter), assetMan.totalKroAssets(asserter));

        // Asserter removed from validator tree
        assertEq(valMan.startedValidatorCount(), count + 1);
        assertTrue(valMan.getStatus(asserter) == IValidatorManager.ValidatorStatus.ACTIVE);

        // Security council balance of asset token increased by tax
        uint128 taxAmount = (slashingAmount * assetMan.TAX_NUMERATOR()) /
            assetMan.TAX_DENOMINATOR();
        assertEq(assetToken.balanceOf(assetMan.SECURITY_COUNCIL()), taxAmount);

        // Challenger asset increased by output reward and challenge reward
        // Boosted reward with 100 kgh delegation
        uint128 boostedReward = 6283173600000736769;
        uint128 challengeReward = slashingAmount - taxAmount;
        uint128 challengerAsset = assetMan.reflectiveWeight(challenger);
        assertEq(
            challengerAsset,
            minStartAmount +
                kghCounts *
                kghManager.totalKroInKgh(1) +
                baseReward +
                boostedReward -
                1 + // Boosted reward is reduced by 1 when distributed to validator and delegators
                challengeReward -
                1 // Challenge reward is reduced by 1 when distributed to each assets in validator vault
        );
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

        // Warp to public round
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

        warpToSubmitTime();
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

        // Warp to public round
        vm.warp(oracle.nextOutputMinL2Timestamp() + roundDuration + 1);
        vm.prank(address(oracle));
        vm.expectRevert(IValidatorManager.ImproperValidatorStatus.selector);
        valMan.checkSubmissionEligibility(trusted);
    }

    function test_checkSubmissionEligibility_inJail_reverts() external {
        test_afterSubmitL2Output_tryJail_succeeds();

        mockValMan.updatePriorityValidator(asserter);

        vm.prank(address(oracle));
        vm.expectRevert(IValidatorManager.ImproperValidatorStatus.selector);
        valMan.checkSubmissionEligibility(asserter);
    }

    function test_checkSubmissionEligibility_publicRound_inJail_reverts() external {
        test_afterSubmitL2Output_tryJail_succeeds();

        mockValMan.updatePriorityValidator(trusted);

        // Warp to public round
        vm.warp(oracle.nextOutputMinL2Timestamp() + roundDuration + 1);
        vm.prank(address(oracle));
        vm.expectRevert(IValidatorManager.ImproperValidatorStatus.selector);
        valMan.checkSubmissionEligibility(asserter);
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
