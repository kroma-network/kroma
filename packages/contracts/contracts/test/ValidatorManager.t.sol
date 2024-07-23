// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { IERC20 } from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import { IERC721 } from "@openzeppelin/contracts/token/ERC721/IERC721.sol";

import { Constants } from "../libraries/Constants.sol";
import { Types } from "../libraries/Types.sol";
import { Proxy } from "../universal/Proxy.sol";
import { IValidatorManager } from "../L1/interfaces/IValidatorManager.sol";
import { L2OutputOracle } from "../L1/L2OutputOracle.sol";
import { ValidatorManager } from "../L1/ValidatorManager.sol";
import { ValidatorPool } from "../L1/ValidatorPool.sol";
import { MockAssetManager } from "./AssetManager.t.sol";
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

    function mockSetNextFinalizeOutputIndex(uint256 l2OutputIndex) external {
        nextFinalizeOutputIndex = l2OutputIndex;
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
}

contract ValidatorManagerTest is ValidatorSystemUpgrade_Initializer {
    MockL2OutputOracle mockOracle;
    MockValidatorManager mockValMgr;
    MockAssetManager mockAssetMgr;
    uint128 public VKRO_PER_KGH;

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

    function _setUpHundredKghDelegation(address validator, uint256 startingTokenId) private {
        uint256[] memory tokenIds = new uint256[](100);
        for (uint256 i = startingTokenId; i < 100 + startingTokenId; i++) {
            kgh.mint(validator, i);
            vm.prank(validator);
            kgh.approve(address(assetMgr), i);
            tokenIds[i - startingTokenId] = i;
        }
        vm.prank(validator);
        assetMgr.delegateKghBatch(validator, tokenIds);
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

        address assetMgrAddress = address(assetMgr);
        MockAssetManager mockAssetMgrImpl = new MockAssetManager(
            IERC20(assetToken),
            IERC721(kgh),
            guardian,
            validatorRewardVault,
            valMgr,
            minDelegationPeriod,
            bondAmount
        );
        vm.prank(multisig);
        Proxy(payable(assetMgrAddress)).upgradeTo(address(mockAssetMgrImpl));
        mockAssetMgr = MockAssetManager(assetMgrAddress);

        VKRO_PER_KGH = assetMgr.VKRO_PER_KGH();

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
        assertEq(valMgr.JAIL_PERIOD_SECONDS(), jailPeriodSeconds);
        assertEq(valMgr.JAIL_THRESHOLD(), jailThreshold);
        assertEq(valMgr.MAX_OUTPUT_FINALIZATIONS(), maxOutputFinalizations);
        assertEq(valMgr.BASE_REWARD(), baseReward);
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

        vm.startPrank(trusted);
        assetToken.approve(address(assetMgr), uint256(assets));
        vm.expectEmit(true, false, false, true, address(valMgr));
        emit ValidatorActivated(trusted, block.timestamp);
        vm.expectEmit(true, true, false, true, address(valMgr));
        emit ValidatorRegistered(trusted, true, commissionRate, assets);
        valMgr.registerValidator(assets, commissionRate, withdrawAcc);
        vm.stopPrank();

        assertEq(assetToken.balanceOf(trusted), trustedBalance - assets);
        assertEq(assetMgr.totalKroAssets(trusted), assets);
        assertEq(valMgr.getCommissionRate(trusted), commissionRate);
        assertEq(valMgr.getWithdrawAccount(trusted), withdrawAcc);

        assertTrue(valMgr.getStatus(trusted) == IValidatorManager.ValidatorStatus.ACTIVE);
        assertEq(valMgr.activatedValidatorCount(), count + 1);
        assertEq(valMgr.getWeight(trusted), assets);
    }

    function test_registerValidator_registered_succeeds() external {
        uint32 count = valMgr.activatedValidatorCount();

        uint128 assets = minActivateAmount - 1;
        uint8 commissionRate = 10;

        vm.startPrank(trusted);
        assetToken.approve(address(assetMgr), uint256(assets));
        vm.expectEmit(true, true, false, true, address(valMgr));
        emit ValidatorRegistered(trusted, false, commissionRate, assets);
        valMgr.registerValidator(assets, commissionRate, withdrawAcc);
        vm.stopPrank();

        assertTrue(valMgr.getStatus(trusted) == IValidatorManager.ValidatorStatus.REGISTERED);
        assertEq(valMgr.activatedValidatorCount(), count);
        assertEq(valMgr.getWeight(trusted), 0);
    }

    function test_registerValidator_alreadyInitiated_reverts() external {
        uint128 assets = minActivateAmount;

        _registerValidator(trusted, assets);

        vm.startPrank(trusted);
        assetToken.approve(address(assetMgr), uint256(assets));
        vm.expectRevert(IValidatorManager.ImproperValidatorStatus.selector);
        valMgr.registerValidator(assets, 10, withdrawAcc);
    }

    function test_registerValidator_smallAsset_reverts() external {
        uint128 assets = minRegisterAmount - 1;

        vm.startPrank(trusted);
        assetToken.approve(address(assetMgr), uint256(assets));
        vm.expectRevert(IValidatorManager.InsufficientAsset.selector);
        valMgr.registerValidator(assets, 10, withdrawAcc);
    }

    function test_registerValidator_largeCommissionRate_reverts() external {
        uint128 assets = minRegisterAmount;

        vm.startPrank(trusted);
        assetToken.approve(address(assetMgr), uint256(assets));
        vm.expectRevert(IValidatorManager.MaxCommissionRateExceeded.selector);
        valMgr.registerValidator(assets, 101, withdrawAcc);
    }

    function test_registerValidator_withdrawZeroAddr_reverts() external {
        uint128 assets = minRegisterAmount;

        vm.startPrank(trusted);
        assetToken.approve(address(assetMgr), uint256(assets));
        vm.expectRevert(IValidatorManager.ZeroAddress.selector);
        valMgr.registerValidator(assets, 10, address(0));
    }

    function test_activateValidator_succeeds() external {
        uint32 count = valMgr.activatedValidatorCount();

        _registerValidator(trusted, minActivateAmount - 1);
        vm.startPrank(asserter);
        assetToken.approve(address(assetMgr), 1);
        assetMgr.delegate(trusted, 1);
        vm.stopPrank();
        assertTrue(valMgr.getStatus(trusted) == IValidatorManager.ValidatorStatus.READY);

        vm.prank(trusted);
        vm.expectEmit(true, false, false, true, address(valMgr));
        emit ValidatorActivated(trusted, block.timestamp);
        valMgr.activateValidator();

        assertTrue(valMgr.getStatus(trusted) == IValidatorManager.ValidatorStatus.ACTIVE);
        assertEq(valMgr.activatedValidatorCount(), count + 1);
        assertEq(valMgr.getWeight(trusted), minActivateAmount);
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
        _registerValidator(trusted, minActivateAmount);
        uint128 kroShares = assetMgr.getKroTotalShareBalance(trusted, trusted);
        vm.prank(trusted);
        assetMgr.initUndelegate(trusted, kroShares);

        assertTrue(valMgr.getStatus(trusted) == IValidatorManager.ValidatorStatus.EXITED);

        vm.prank(trusted);
        vm.expectRevert(IValidatorManager.ImproperValidatorStatus.selector);
        valMgr.activateValidator();
    }

    function test_activateValidator_inJail_reverts() external {
        test_afterSubmitL2Output_tryJail_succeeds();

        // Undelegate all assets of jailed validator
        uint128 kroShares = assetMgr.getKroTotalShareBalance(asserter, asserter);
        vm.prank(asserter);
        assetMgr.initUndelegate(asserter, kroShares);
        assertTrue(valMgr.getStatus(asserter) == IValidatorManager.ValidatorStatus.EXITED);

        // Delegate to re-activate validator
        vm.startPrank(asserter);
        assetToken.approve(address(assetMgr), minActivateAmount);
        assetMgr.delegate(asserter, minActivateAmount);
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

    function test_afterSubmitL2Output_distributeReward_succeeds() external {
        // Register validator with commission rate 10%
        _registerValidator(trusted, minActivateAmount);

        // Delegate 100 KGHs
        uint128 kghCounts = 100;
        _setUpHundredKghDelegation(trusted, 1);
        assertEq(assetMgr.totalKghNum(trusted), kghCounts);

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
        vm.prank(valMgr.nextValidator());
        mockOracle.addOutput(oracle.nextBlockNumber());
        vm.prank(address(oracle));
        vm.expectEmit(true, false, false, true, address(valMgr));
        emit RewardDistributed(
            terminateOutputIndex + 1,
            trusted,
            validatorReward,
            baseReward,
            boostedReward
        );
        valMgr.afterSubmitL2Output(outputIndex);

        uint128 kroReward = assetMgr.totalKroAssets(trusted) -
            minActivateAmount -
            kghCounts *
            VKRO_PER_KGH;
        vm.prank(trusted);
        uint128 oneKghReward = mockAssetMgr.convertToKghAssets(trusted, trusted, 1) - VKRO_PER_KGH;

        assertEq(kroReward, baseReward);
        assertEq(oneKghReward, boostedReward / kghCounts);

        // Check validator tree updated with rewards
        assertEq(
            valMgr.getWeight(trusted),
            minActivateAmount +
                kghManager.totalKroInKgh(1) *
                kghCounts +
                baseReward +
                boostedReward +
                validatorReward
        );

        assertEq(oracle.nextFinalizeOutputIndex(), terminateOutputIndex + 2);
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
        _registerValidator(trusted, minActivateAmount);

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
        uint256 outputIndex = oracle.nextOutputIndex();

        vm.prank(trusted);
        mockOracle.addOutput(oracle.nextBlockNumber());
        vm.prank(address(oracle));
        vm.expectEmit(true, false, false, true, address(valMgr));
        emit ValidatorJailed(asserter, uint128(block.timestamp) + jailPeriodSeconds);
        vm.expectEmit(true, false, false, true, address(valMgr));
        emit ValidatorStopped(asserter, block.timestamp);
        valMgr.afterSubmitL2Output(outputIndex);

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

        uint128 kroShares = assetMgr.getKroTotalShareBalance(trusted, trusted);
        vm.prank(trusted);
        assetMgr.initUndelegate(trusted, kroShares);
        assertTrue(valMgr.getStatus(trusted) == IValidatorManager.ValidatorStatus.EXITED);

        vm.prank(asserter);
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

        uint128 kroShares = assetMgr.getKroTotalShareBalance(trusted, trusted);
        vm.prank(trusted);
        assetMgr.initUndelegate(trusted, kroShares);
        assertTrue(valMgr.getStatus(trusted) == IValidatorManager.ValidatorStatus.EXITED);

        vm.prank(asserter);
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

    function test_tryUnjail_succeeds() external {
        test_afterSubmitL2Output_tryJail_succeeds();
        assertTrue(valMgr.getStatus(asserter) == IValidatorManager.ValidatorStatus.READY);

        vm.warp(valMgr.jailExpiresAt(asserter));
        vm.prank(asserter);
        vm.expectEmit(true, false, false, false, address(valMgr));
        emit ValidatorUnjailed(asserter);
        vm.expectEmit(true, false, false, true, address(valMgr));
        emit ValidatorActivated(asserter, block.timestamp);
        valMgr.tryUnjail(asserter, false);

        assertEq(valMgr.noSubmissionCount(asserter), 0);
        assertFalse(valMgr.inJail(asserter));
        assertTrue(valMgr.getStatus(asserter) == IValidatorManager.ValidatorStatus.ACTIVE);
    }

    function test_tryUnjail_notInJail_reverts() external {
        vm.prank(asserter);
        vm.expectRevert(IValidatorManager.ImproperValidatorStatus.selector);
        valMgr.tryUnjail(asserter, false);
    }

    function test_tryUnjail_senderNotSelf_reverts() external {
        test_afterSubmitL2Output_tryJail_succeeds();

        vm.prank(trusted);
        vm.expectRevert(IValidatorManager.NotAllowedCaller.selector);
        valMgr.tryUnjail(asserter, false);
    }

    function test_tryUnjail_force_senderNotColosseum_reverts() external {
        test_afterSubmitL2Output_tryJail_succeeds();

        vm.prank(asserter);
        vm.expectRevert(IValidatorManager.NotAllowedCaller.selector);
        valMgr.tryUnjail(asserter, true);
    }

    function test_tryUnjail_periodNotElapsed_reverts() external {
        test_afterSubmitL2Output_tryJail_succeeds();

        vm.prank(asserter);
        vm.expectRevert(IValidatorManager.NotElapsedJailPeriod.selector);
        valMgr.tryUnjail(asserter, false);
    }

    function test_slash_succeeds() external {
        uint32 count = valMgr.activatedValidatorCount();
        // Register as a validator
        _registerValidator(asserter, minActivateAmount);
        _registerValidator(challenger, minActivateAmount);
        assertEq(valMgr.activatedValidatorCount(), count + 2);

        // Delegate KGHs
        uint128 kghCounts = 100;
        _setUpHundredKghDelegation(asserter, 1);
        _setUpHundredKghDelegation(challenger, 1 + kghCounts);
        assertEq(assetMgr.totalKghNum(asserter), kghCounts);
        assertEq(assetMgr.totalKghNum(challenger), kghCounts);

        // Submit the first output which interacts with ValidatorManager
        mockValMgr.updatePriorityValidator(asserter);
        warpToSubmitTime();
        _submitL2OutputV2(false);
        uint256 challengedOutputIndex = oracle.latestOutputIndex();

        // Suppose that the challenge is successful, so the winner is challenger
        uint128 slashingAmount = (minActivateAmount * slashingRate) /
            assetMgr.SLASHING_RATE_DENOM();
        vm.prank(address(colosseum));
        vm.expectEmit(true, true, false, true, address(valMgr));
        emit Slashed(challengedOutputIndex, asserter, slashingAmount);
        vm.expectEmit(true, false, false, true, address(valMgr));
        emit ValidatorJailed(asserter, uint128(block.timestamp) + jailPeriodSeconds);
        vm.expectEmit(true, false, false, true, address(valMgr));
        emit ValidatorStopped(asserter, block.timestamp);
        valMgr.slash(challengedOutputIndex, challenger, asserter);

        // This will be done by the l2 output oracle contract in the real environment
        vm.prank(challenger);
        mockOracle.replaceOutput(challengedOutputIndex);

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
        uint128 asserterTotalKro = assetMgr.totalKroAssets(asserter) -
            kghCounts *
            kghManager.totalKroInKgh(1);
        assertEq(asserterTotalKro, minActivateAmount - slashingAmount);
        assertEq(assetMgr.totalValidatorKro(asserter), asserterTotalKro);
        // Asserter has 0 rewards
        assertEq(assetMgr.reflectiveWeight(asserter), assetMgr.totalKroAssets(asserter));

        // Security council balance of asset token increased by tax
        uint128 taxAmount = (slashingAmount * assetMgr.TAX_NUMERATOR()) /
            assetMgr.TAX_DENOMINATOR();
        assertEq(assetToken.balanceOf(assetMgr.SECURITY_COUNCIL()), taxAmount);

        // Challenger asset increased by output reward and challenge reward
        // Boosted reward with 100 kgh delegation
        uint128 boostedReward = 6283173600000736769;
        uint128 challengeReward = slashingAmount - taxAmount;
        uint128 challengerAsset = assetMgr.reflectiveWeight(challenger);
        assertEq(
            challengerAsset,
            minActivateAmount +
                kghCounts *
                kghManager.totalKroInKgh(1) +
                baseReward +
                boostedReward -
                1 + // Boosted reward is reduced by 1 when distributed to validator and delegators
                challengeReward -
                1 // Challenge reward is reduced by 1 when distributed to each assets in validator vault
        );
        assertGt(assetMgr.totalValidatorKro(challenger), minActivateAmount);
    }

    function test_slash_notColosseum_reverts() external {
        vm.prank(address(1));
        vm.expectRevert(IValidatorManager.NotAllowedCaller.selector);
        valMgr.slash(1, challenger, asserter);
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

        uint128 minUndelegateShares = assetMgr.previewDelegate(trusted, 1);
        vm.prank(trusted);
        assetMgr.initUndelegate(trusted, minUndelegateShares);
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
