// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { IERC20 } from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import { IERC721 } from "@openzeppelin/contracts/token/ERC721/IERC721.sol";

import { AssetManager } from "../L1/AssetManager.sol";
import { ValidatorManager } from "../L1/ValidatorManager.sol";
import { IAssetManager } from "../L1/interfaces/IAssetManager.sol";
import { IValidatorManager } from "../L1/interfaces/IValidatorManager.sol";
import { Proxy } from "../universal/Proxy.sol";
import { ValidatorSystemUpgrade_Initializer } from "./CommonTest.t.sol";
import { MockL2OutputOracle } from "./ValidatorManager.t.sol";

contract MockAssetManager is AssetManager {
    constructor(
        IERC20 _assetToken,
        IERC721 _kgh,
        address _securityCouncil,
        address _validatorRewardVault,
        IValidatorManager _validatorManager,
        uint128 _minDelegationPeriod,
        uint128 _bondAmount
    )
        AssetManager(
            _assetToken,
            _kgh,
            _securityCouncil,
            _validatorRewardVault,
            _validatorManager,
            _minDelegationPeriod,
            _bondAmount
        )
    {}

    function increaseRewardPerKgh(address validator, uint128 amount) external {
        _vaults[validator].asset.rewardPerKghStored += amount;
    }

    function increaseBaseReward(address validator, uint128 amount) external {
        _vaults[validator].asset.totalKro += amount;
    }
}

contract MockValidatorManager is ValidatorManager {
    constructor(ConstructorParams memory _constructorParams) ValidatorManager(_constructorParams) {}

    function sendToJail(address validator) external {
        _jail[validator] = uint128(block.timestamp) + HARD_JAIL_PERIOD_SECONDS;
    }

    function calculateBoostedReward(address validator) external view returns (uint128) {
        return _getBoostedReward(validator);
    }
}

// Tests the implementations of the AssetManager
contract AssetManagerTest is ValidatorSystemUpgrade_Initializer {
    MockAssetManager public mockAssetMgr;
    MockValidatorManager public mockValMgr;
    MockL2OutputOracle public mockOracle;
    address public validator = trusted;

    function setUp() public override {
        super.setUp();

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
        Proxy(payable(address(oracle))).upgradeTo(address(mockOracleImpl));
        mockOracle = MockL2OutputOracle(address(oracle));

        MockAssetManager assetManagerImpl = new MockAssetManager(
            IERC20(assetToken),
            IERC721(kgh),
            guardian,
            validatorRewardVault,
            valMgr,
            minDelegationPeriod,
            bondAmount
        );
        vm.prank(multisig);
        Proxy(payable(address(assetMgr))).upgradeTo(address(assetManagerImpl));
        mockAssetMgr = MockAssetManager(address(assetMgr));

        MockValidatorManager mockValMgrImpl = new MockValidatorManager(constructorParams);
        vm.prank(multisig);
        Proxy(payable(address(valMgr))).upgradeTo(address(mockValMgrImpl));
        mockValMgr = MockValidatorManager(address(valMgr));

        // Submit until terminateOutputIndex and set next output index to be finalized after it
        for (uint256 i = oracle.nextOutputIndex(); i <= terminateOutputIndex; i++) {
            _submitOutputRoot(pool.nextValidator());
        }
        vm.warp(oracle.finalizedAt(terminateOutputIndex));
        mockOracle.mockSetNextFinalizeOutputIndex(terminateOutputIndex + 1);
    }

    function _submitOutputRoot(address _validator) internal {
        uint256 nextBlockNumber = oracle.nextBlockNumber();
        warpToSubmitTime();
        vm.prank(_validator);
        mockOracle.addOutput(nextBlockNumber);
    }

    function _registerValidator(uint128 amount) internal {
        vm.startPrank(validator, validator);
        assetToken.approve(address(assetMgr), amount);
        valMgr.registerValidator(amount, 0, withdrawAcc);
        vm.stopPrank();
    }

    function _delegate(address _delegator, uint128 amount) internal {
        vm.startPrank(_delegator);
        assetToken.approve(address(assetMgr), amount);
        assetMgr.delegate(validator, amount);
        vm.stopPrank();
    }

    function _delegateKgh(address _delegator, uint256 tokenId) internal {
        kgh.mint(_delegator, tokenId);
        vm.startPrank(_delegator);
        kgh.approve(address(assetMgr), tokenId);
        assetMgr.delegateKgh(validator, tokenId);
        vm.stopPrank();
    }

    function _delegateKghBatch(uint256 kghCount) internal {
        uint256[] memory tokenIds = new uint256[](kghCount);
        for (uint256 i = 1; i < kghCount + 1; i++) {
            kgh.mint(delegator, i);
            vm.prank(delegator);
            kgh.approve(address(assetMgr), i);
            tokenIds[i - 1] = i;
        }

        vm.prank(delegator);
        assetMgr.delegateKghBatch(validator, tokenIds);
    }

    function _undelegate(address _delegator) internal returns (uint128) {
        uint128 delegatorKro = assetMgr.getKroAssets(validator, _delegator);
        vm.warp(assetMgr.canUndelegateKroAt(validator, _delegator));
        vm.prank(_delegator);
        assetMgr.undelegate(validator, delegatorKro);

        return delegatorKro;
    }

    function _undelegateKgh(address _delegator, uint256 tokenId) internal {
        vm.warp(assetMgr.canUndelegateKghAt(validator, _delegator, tokenId));
        vm.prank(_delegator);
        assetMgr.undelegateKgh(validator, tokenId);
    }

    function _undelegateKghBatch(uint256 kghCount) internal {
        uint256[] memory tokenIds = new uint256[](kghCount);
        for (uint256 i = 1; i < kghCount + 1; i++) {
            tokenIds[i - 1] = i;
        }

        vm.warp(assetMgr.canUndelegateKghAt(validator, delegator, tokenIds[0]));
        vm.prank(delegator);
        assetMgr.undelegateKghBatch(validator, tokenIds);
    }

    function test_constructor_succeeds() external {
        assertEq(address(assetMgr.ASSET_TOKEN()), address(assetToken));
        assertEq(address(assetMgr.KGH()), address(kgh));
        assertEq(assetMgr.SECURITY_COUNCIL(), guardian);
        assertEq(address(assetMgr.VALIDATOR_REWARD_VAULT()), validatorRewardVault);
        assertEq(address(assetMgr.VALIDATOR_MANAGER()), address(valMgr));
        assertEq(assetMgr.MIN_DELEGATION_PERIOD(), minDelegationPeriod);
        assertEq(assetMgr.BOND_AMOUNT(), bondAmount);
    }

    function test_depositToRegister_succeeds() external {
        uint256 beforeBalance = assetToken.balanceOf(validator);

        vm.prank(validator);
        assetToken.approve(address(assetMgr), minActivateAmount);
        vm.prank(address(valMgr));
        assetMgr.depositToRegister(validator, minActivateAmount, withdrawAcc);

        assertEq(assetMgr.getWithdrawAccount(validator), withdrawAcc);
        assertEq(assetMgr.totalValidatorKro(validator), minActivateAmount);
        assertEq(assetMgr.canWithdrawAt(validator), block.timestamp + minDelegationPeriod);
        assertEq(assetToken.balanceOf(validator), beforeBalance - minActivateAmount);
        assertEq(assetToken.balanceOf(address(assetMgr)), minActivateAmount);
    }

    function test_depositToRegister_callerNotValMgr_reverts() external {
        vm.expectRevert(IAssetManager.NotAllowedCaller.selector);
        assetMgr.depositToRegister(validator, minActivateAmount, withdrawAcc);
    }

    function test_depositToRegister_zeroWithdrawAcc_reverts() external {
        vm.prank(address(valMgr));
        vm.expectRevert(IAssetManager.ZeroAddress.selector);
        assetMgr.depositToRegister(validator, minActivateAmount, address(0));
    }

    function test_deposit_succeeds() external {
        _registerValidator(minActivateAmount);

        uint120 beforeWeight = valMgr.getWeight(validator);
        uint256 beforeBalance = assetToken.balanceOf(validator);
        uint128 beforeValidatorKro = assetMgr.totalValidatorKro(validator);

        vm.startPrank(validator);
        assetToken.approve(address(assetMgr), bondAmount);
        assetMgr.deposit(bondAmount);

        assertEq(assetMgr.totalValidatorKro(validator), beforeValidatorKro + bondAmount);
        assertEq(assetMgr.canWithdrawAt(validator), block.timestamp + minDelegationPeriod);
        assertEq(valMgr.getWeight(validator), beforeWeight + bondAmount);
        assertEq(assetToken.balanceOf(validator), beforeBalance - bondAmount);
    }

    function test_deposit_activate_succeeds() external {
        _registerValidator(minActivateAmount - 1);
        assertEq(valMgr.getWeight(validator), 0);

        vm.startPrank(validator);
        assetToken.approve(address(assetMgr), 1);
        assetMgr.deposit(1);

        assertEq(valMgr.getWeight(validator), minActivateAmount);
    }

    function test_deposit_notActivate_succeeds() external {
        _registerValidator(minActivateAmount - 2);
        assertEq(valMgr.getWeight(validator), 0);

        vm.startPrank(validator);
        assetToken.approve(address(assetMgr), 1);
        assetMgr.deposit(1);

        assertEq(valMgr.getWeight(validator), 0);
    }

    function test_deposit_inJailNotActivate_succeeds() external {
        _registerValidator(minActivateAmount - 1);
        assertEq(valMgr.getWeight(validator), 0);

        mockValMgr.sendToJail(validator);
        assertTrue(valMgr.inJail(validator));

        vm.startPrank(validator);
        assetToken.approve(address(assetMgr), 1);
        assetMgr.deposit(1);

        assertEq(valMgr.getWeight(validator), 0);
    }

    function test_deposit_zeroAsset_reverts() external {
        vm.expectRevert(IAssetManager.NotAllowedZeroInput.selector);
        assetMgr.deposit(0);
    }

    function test_deposit_validatorStatusNone_reverts() external {
        vm.prank(validator);
        vm.expectRevert(IAssetManager.ImproperValidatorStatus.selector);
        assetMgr.deposit(bondAmount);
    }

    function test_withdraw_succeeds() public {
        _registerValidator(minActivateAmount);
        assertEq(valMgr.getWeight(validator), minActivateAmount);
        uint256 beforeBalance = assetToken.balanceOf(validator);

        address withdrawAccount = assetMgr.getWithdrawAccount(validator);
        vm.warp(assetMgr.canWithdrawAt(validator));
        vm.prank(withdrawAccount);
        assetMgr.withdraw(validator, minActivateAmount);

        assertEq(assetMgr.totalValidatorKro(validator), 0);
        assertEq(valMgr.getWeight(validator), 0);
        assertEq(assetToken.balanceOf(validator), beforeBalance);
        assertEq(assetToken.balanceOf(withdrawAccount), minActivateAmount);
    }

    function test_withdraw_notWithdrawAcc_reverts() external {
        _registerValidator(minActivateAmount);

        vm.prank(validator);
        vm.expectRevert(IAssetManager.NotAllowedCaller.selector);
        assetMgr.withdraw(validator, minActivateAmount);
    }

    function test_withdraw_zeroAsset_reverts() external {
        _registerValidator(minActivateAmount);

        vm.prank(assetMgr.getWithdrawAccount(validator));
        vm.expectRevert(IAssetManager.NotAllowedZeroInput.selector);
        assetMgr.withdraw(validator, 0);
    }

    function test_withdraw_notElapsedMinDelegationPeriod_reverts() external {
        _registerValidator(minActivateAmount);

        vm.prank(assetMgr.getWithdrawAccount(validator));
        vm.expectRevert(IAssetManager.NotElapsedMinDelegationPeriod.selector);
        assetMgr.withdraw(validator, minActivateAmount);
    }

    function test_withdraw_notExpiredJailPeriod_reverts() external {
        _registerValidator(minActivateAmount);

        vm.warp(assetMgr.canWithdrawAt(validator));
        mockValMgr.sendToJail(validator);

        vm.prank(assetMgr.getWithdrawAccount(validator));
        vm.expectRevert(IAssetManager.ImproperValidatorStatus.selector);
        assetMgr.withdraw(validator, minActivateAmount);

        vm.warp(valMgr.jailExpiresAt(validator));
        vm.prank(assetMgr.getWithdrawAccount(validator));
        assetMgr.withdraw(validator, minActivateAmount);
    }

    function test_withdraw_insufficientValidatorKro_reverts() external {
        _registerValidator(minActivateAmount);

        vm.warp(assetMgr.canWithdrawAt(validator));
        vm.prank(assetMgr.getWithdrawAccount(validator));
        vm.expectRevert(IAssetManager.InsufficientAsset.selector);
        assetMgr.withdraw(validator, minActivateAmount + 1);
    }

    function test_withdraw_validatorKroBonded_reverts() external {
        _registerValidator(minActivateAmount);

        vm.prank(address(valMgr));
        assetMgr.bondValidatorKro(validator);

        address withdrawAccount = assetMgr.getWithdrawAccount(validator);
        vm.warp(assetMgr.canWithdrawAt(validator));
        vm.startPrank(withdrawAccount);
        vm.expectRevert(IAssetManager.InsufficientAsset.selector);
        assetMgr.withdraw(validator, minActivateAmount);

        uint128 validatorKroBonded = assetMgr.totalValidatorKroBonded(validator);
        assetMgr.withdraw(validator, minActivateAmount - validatorKroBonded);

        assertEq(assetMgr.totalValidatorKro(validator), validatorKroBonded);
        assertEq(valMgr.getWeight(validator), 0);
        assertEq(assetToken.balanceOf(withdrawAccount), minActivateAmount - validatorKroBonded);
    }

    function test_delegate_succeeds() external {
        _registerValidator(minActivateAmount);

        uint128 shares = assetMgr.previewDelegate(validator, bondAmount);
        uint256 beforeBalance = assetToken.balanceOf(delegator);
        uint120 beforeWeight = valMgr.getWeight(validator);

        _delegate(delegator, bondAmount);

        assertEq(assetToken.balanceOf(delegator), beforeBalance - bondAmount);
        assertEq(assetMgr.totalKroAssets(validator), bondAmount);
        assertEq(assetMgr.getKroTotalShareBalance(validator, delegator), shares);
        assertEq(
            assetMgr.canUndelegateKroAt(validator, delegator),
            block.timestamp + minDelegationPeriod
        );
        assertEq(valMgr.getWeight(validator), beforeWeight + bondAmount);
    }

    function test_delegate_validatorStatusNone_reverts() external {
        vm.expectRevert(IAssetManager.ImproperValidatorStatus.selector);
        assetMgr.delegate(validator, bondAmount);
    }

    function test_delegate_validatorStatusExited_reverts() external {
        test_withdraw_succeeds();
        assertTrue(valMgr.getStatus(validator) == IValidatorManager.ValidatorStatus.EXITED);

        vm.expectRevert(IAssetManager.ImproperValidatorStatus.selector);
        assetMgr.delegate(validator, bondAmount);
    }

    function test_delegate_validatorInJail_reverts() external {
        _registerValidator(minActivateAmount);
        mockValMgr.sendToJail(validator);
        assertTrue(valMgr.inJail(validator));

        vm.expectRevert(IAssetManager.ImproperValidatorStatus.selector);
        assetMgr.delegate(validator, bondAmount);
    }

    function test_delegate_zeroAsset_reverts() external {
        _registerValidator(minActivateAmount);

        vm.expectRevert(IAssetManager.NotAllowedZeroInput.selector);
        vm.prank(delegator);
        assetMgr.delegate(validator, 0);
    }

    function test_delegateKgh_succeeds() external {
        _registerValidator(minActivateAmount);
        uint256 beforeBalance = assetToken.balanceOf(delegator);

        uint256 tokenId = 0;
        _delegateKgh(delegator, tokenId);

        assertEq(assetToken.balanceOf(delegator), beforeBalance);
        assertEq(assetMgr.getKghReward(validator, delegator), 0);
        assertEq(kgh.ownerOf(tokenId), address(assetMgr));
        assertEq(assetMgr.totalKghNum(validator), 1);
        assertEq(assetMgr.getKghNum(validator, delegator), 1);
        assertEq(
            assetMgr.canUndelegateKghAt(validator, delegator, tokenId),
            block.timestamp + minDelegationPeriod
        );
    }

    function test_delegateKgh_claimBoostedReward_succeeds() external {
        _registerValidator(minActivateAmount);
        _delegateKgh(delegator, 0);

        uint256 beforeBalance = assetToken.balanceOf(delegator);

        // Increase boosted reward
        uint128 boostedReward = mockValMgr.calculateBoostedReward(validator);
        mockAssetMgr.increaseRewardPerKgh(validator, boostedReward);
        assertEq(assetMgr.getKghReward(validator, delegator), boostedReward);

        // Delegate one more KGH
        _delegateKgh(delegator, 1);

        assertEq(assetToken.balanceOf(delegator), beforeBalance + boostedReward);
    }

    function test_delegateKgh_validatorStatusNone_reverts() external {
        vm.expectRevert(IAssetManager.ImproperValidatorStatus.selector);
        assetMgr.delegateKgh(validator, 0);
    }

    function test_delegateKgh_validatorStatusExited_reverts() external {
        test_withdraw_succeeds();
        assertTrue(valMgr.getStatus(validator) == IValidatorManager.ValidatorStatus.EXITED);

        vm.expectRevert(IAssetManager.ImproperValidatorStatus.selector);
        assetMgr.delegateKgh(validator, 0);
    }

    function test_delegateKgh_validatorInJail_reverts() external {
        _registerValidator(minActivateAmount);
        mockValMgr.sendToJail(validator);
        assertTrue(valMgr.inJail(validator));

        vm.expectRevert(IAssetManager.ImproperValidatorStatus.selector);
        assetMgr.delegateKgh(validator, 0);
    }

    function test_delegateKghBatch_succeeds() external {
        _registerValidator(minActivateAmount);
        uint256 beforeBalance = assetToken.balanceOf(delegator);

        uint256 kghCount = 10;
        _delegateKghBatch(kghCount);

        assertEq(assetToken.balanceOf(delegator), beforeBalance);
        assertEq(assetMgr.getKghReward(validator, delegator), 0);
        assertEq(assetMgr.totalKghNum(validator), kghCount);
        assertEq(assetMgr.getKghNum(validator, delegator), kghCount);
        for (uint256 i = 1; i < kghCount + 1; i++) {
            assertEq(kgh.ownerOf(i), address(assetMgr));
            assertEq(
                assetMgr.canUndelegateKghAt(validator, delegator, i),
                block.timestamp + minDelegationPeriod
            );
        }
    }

    function test_delegateKghBatch_claimBoostedReward_succeeds() external {
        _registerValidator(minActivateAmount);
        _delegateKgh(delegator, 0);

        uint256 beforeBalance = assetToken.balanceOf(delegator);

        // Increase boosted reward
        uint128 boostedReward = mockValMgr.calculateBoostedReward(validator);
        mockAssetMgr.increaseRewardPerKgh(validator, boostedReward);
        assertEq(assetMgr.getKghReward(validator, delegator), boostedReward);

        // Delegate two more KGHs
        _delegateKghBatch(2);

        assertEq(assetToken.balanceOf(delegator), beforeBalance + boostedReward);
    }

    function test_delegateKghBatch_validatorStatusNone_reverts() external {
        uint256[] memory tokenIds = new uint256[](0);
        vm.expectRevert(IAssetManager.ImproperValidatorStatus.selector);
        assetMgr.delegateKghBatch(validator, tokenIds);
    }

    function test_delegateKghBatch_validatorStatusExited_reverts() external {
        test_withdraw_succeeds();
        assertTrue(valMgr.getStatus(validator) == IValidatorManager.ValidatorStatus.EXITED);

        uint256[] memory tokenIds = new uint256[](0);
        vm.expectRevert(IAssetManager.ImproperValidatorStatus.selector);
        assetMgr.delegateKghBatch(validator, tokenIds);
    }

    function test_delegateKghBatch_validatorInJail_reverts() external {
        _registerValidator(minActivateAmount);
        mockValMgr.sendToJail(validator);
        assertTrue(valMgr.inJail(validator));

        uint256[] memory tokenIds = new uint256[](0);
        vm.expectRevert(IAssetManager.ImproperValidatorStatus.selector);
        assetMgr.delegateKghBatch(validator, tokenIds);
    }

    function test_delegateKghBatch_zeroTokenIds_reverts() external {
        _registerValidator(minActivateAmount);

        uint256[] memory tokenIds = new uint256[](0);
        vm.expectRevert(IAssetManager.NotAllowedZeroInput.selector);
        assetMgr.delegateKghBatch(validator, tokenIds);
    }

    function test_undelegate_succeeds() external {
        _registerValidator(minActivateAmount);
        _delegate(delegator, bondAmount);
        mockAssetMgr.increaseBaseReward(validator, baseReward);
        uint256 beforeBalance = assetToken.balanceOf(delegator);

        uint128 delegatorKro = _undelegate(delegator);

        assertEq(assetMgr.totalKroAssets(validator), bondAmount + baseReward - delegatorKro);
        assertEq(
            valMgr.getWeight(validator),
            minActivateAmount + bondAmount + baseReward - delegatorKro
        );
        assertEq(assetToken.balanceOf(delegator), beforeBalance + delegatorKro);
    }

    function test_undelegate_severalDelegators_succeeds() external {
        _registerValidator(minActivateAmount);
        _delegate(delegator, bondAmount);
        mockAssetMgr.increaseBaseReward(validator, baseReward);

        address delegator2 = makeAddr("delegator2");
        assetToken.mint(delegator2, bondAmount);
        _delegate(delegator2, bondAmount);
        mockAssetMgr.increaseBaseReward(validator, baseReward);
        uint256 beforeBalance = assetToken.balanceOf(delegator);

        uint128 delegatorKro = _undelegate(delegator);

        assertEq(
            assetMgr.totalKroAssets(validator),
            bondAmount * 2 + baseReward * 2 - delegatorKro
        );
        assertEq(
            valMgr.getWeight(validator),
            minActivateAmount + bondAmount * 2 + baseReward * 2 - delegatorKro
        );
        assertEq(assetToken.balanceOf(delegator), beforeBalance + delegatorKro);
    }

    function test_undelegate_removedFromValidatorTree_succeeds() external {
        _registerValidator(minActivateAmount - 1);
        _delegate(delegator, bondAmount);

        vm.prank(validator);
        valMgr.activateValidator();
        assertEq(valMgr.getWeight(validator), minActivateAmount - 1 + bondAmount);

        _undelegate(delegator);

        assertEq(valMgr.getWeight(validator), 0);
    }

    function test_undelegate_zeroAsset_reverts() external {
        vm.expectRevert(IAssetManager.NotAllowedZeroInput.selector);
        assetMgr.undelegate(validator, 0);
    }

    function test_undelegate_largeAsset_reverts() external {
        vm.expectRevert(IAssetManager.InsufficientShare.selector);
        assetMgr.undelegate(validator, 1);
    }

    function test_undelegate_notElapsedMinDelegationPeriod_reverts() external {
        _registerValidator(minActivateAmount);
        _delegate(delegator, bondAmount);

        vm.expectRevert(IAssetManager.NotElapsedMinDelegationPeriod.selector);
        vm.prank(delegator);
        assetMgr.undelegate(validator, bondAmount);
    }

    function test_undelegateKgh_succeeds() external {
        _registerValidator(minActivateAmount);
        uint256 tokenId = 0;
        _delegateKgh(delegator, tokenId);

        uint128 boostedReward = mockValMgr.calculateBoostedReward(validator);
        mockAssetMgr.increaseRewardPerKgh(validator, boostedReward);

        uint256 beforeBalance = assetToken.balanceOf(delegator);
        boostedReward = assetMgr.getKghReward(validator, delegator);

        _undelegateKgh(delegator, tokenId);

        assertEq(assetMgr.getKghReward(validator, delegator), 0);
        assertEq(assetMgr.totalKghNum(validator), 0);
        assertEq(assetMgr.getKghNum(validator, delegator), 0);
        assertEq(assetMgr.canUndelegateKghAt(validator, delegator, tokenId), minDelegationPeriod);
        assertEq(kgh.ownerOf(tokenId), delegator);
        assertEq(assetToken.balanceOf(delegator), beforeBalance + boostedReward);
    }

    function test_undelegateKgh_noBoostedReward_succeeds() external {
        _registerValidator(minActivateAmount);
        uint256 tokenId = 0;
        _delegateKgh(delegator, tokenId);
        uint256 beforeBalance = assetToken.balanceOf(delegator);

        _undelegateKgh(delegator, tokenId);

        assertEq(assetToken.balanceOf(delegator), beforeBalance);
    }

    function test_undelegateKgh_severalDelegators_succeeds() external {
        _registerValidator(minActivateAmount);
        _delegateKgh(delegator, 0);
        uint128 boostedReward = mockValMgr.calculateBoostedReward(validator);
        mockAssetMgr.increaseRewardPerKgh(validator, boostedReward);

        address delegator2 = makeAddr("delegator2");
        uint256 tokenId = 1;
        _delegateKgh(delegator2, tokenId);
        boostedReward = mockValMgr.calculateBoostedReward(validator);
        mockAssetMgr.increaseRewardPerKgh(validator, boostedReward);

        uint256 beforeBalance = assetToken.balanceOf(delegator2);
        boostedReward = assetMgr.getKghReward(validator, delegator2);

        _undelegateKgh(delegator2, tokenId);

        assertEq(assetMgr.getKghReward(validator, delegator2), 0);
        assertEq(assetMgr.totalKghNum(validator), 1);
        assertEq(assetMgr.getKghNum(validator, delegator2), 0);
        assertEq(assetMgr.canUndelegateKghAt(validator, delegator2, tokenId), minDelegationPeriod);
        assertEq(kgh.ownerOf(tokenId), delegator2);
        assertEq(assetToken.balanceOf(delegator2), beforeBalance + boostedReward);
    }

    function test_undelegateKgh_invalidTokenIds_reverts() external {
        vm.expectRevert(IAssetManager.InvalidTokenIdsInput.selector);
        assetMgr.undelegateKgh(validator, 0);
    }

    function test_undelegateKgh_notElapsedMinDelegationPeriod_reverts() external {
        _registerValidator(minActivateAmount);
        _delegateKgh(delegator, 0);

        vm.expectRevert(IAssetManager.NotElapsedMinDelegationPeriod.selector);
        vm.prank(delegator);
        assetMgr.undelegateKgh(validator, 0);
    }

    function test_undelegateKghBatch_succeeds() external {
        _registerValidator(minActivateAmount);
        uint256 kghCount = 10;
        _delegateKghBatch(kghCount);

        uint128 boostedReward = mockValMgr.calculateBoostedReward(validator);
        mockAssetMgr.increaseRewardPerKgh(validator, boostedReward);

        uint256 beforeBalance = assetToken.balanceOf(delegator);
        boostedReward = assetMgr.getKghReward(validator, delegator);

        _undelegateKghBatch(kghCount);

        assertEq(assetMgr.getKghReward(validator, delegator), 0);
        assertEq(assetMgr.totalKghNum(validator), 0);
        assertEq(assetMgr.getKghNum(validator, delegator), 0);
        assertEq(assetToken.balanceOf(delegator), beforeBalance + boostedReward);
        for (uint256 i = 1; i < kghCount + 1; i++) {
            assertEq(kgh.ownerOf(i), delegator);
            assertEq(assetMgr.canUndelegateKghAt(validator, delegator, i), minDelegationPeriod);
        }
    }

    function test_undelegateKghBatch_noBoostedReward_succeeds() external {
        _registerValidator(minActivateAmount);
        uint256 kghCount = 10;
        _delegateKghBatch(kghCount);
        uint256 beforeBalance = assetToken.balanceOf(delegator);

        _undelegateKghBatch(kghCount);

        assertEq(assetToken.balanceOf(delegator), beforeBalance);
    }

    function test_undelegateKghBatch_zeroTokenIds_reverts() external {
        uint256[] memory tokenIds = new uint256[](0);
        vm.expectRevert(IAssetManager.NotAllowedZeroInput.selector);
        assetMgr.undelegateKghBatch(validator, tokenIds);
    }

    function test_undelegateKghBatch_invalidTokenIds_reverts() external {
        uint256[] memory tokenIds = new uint256[](1);
        tokenIds[0] = 0;
        vm.expectRevert(IAssetManager.InvalidTokenIdsInput.selector);
        assetMgr.undelegateKghBatch(validator, tokenIds);
    }

    function test_undelegateKghBatch_notElapsedMinDelegationPeriod_reverts() external {
        _registerValidator(minActivateAmount);
        _delegateKgh(delegator, 0);
        vm.warp(assetMgr.canUndelegateKghAt(validator, delegator, 0));
        _delegateKgh(delegator, 1);

        uint256[] memory tokenIds = new uint256[](2);
        tokenIds[0] = 0;
        tokenIds[1] = 1;
        vm.expectRevert(IAssetManager.NotElapsedMinDelegationPeriod.selector);
        vm.prank(delegator);
        assetMgr.undelegateKghBatch(validator, tokenIds);
    }

    function test_claimKghReward_succeeds() external {
        _registerValidator(minActivateAmount);
        _delegateKgh(delegator, 0);
        uint128 boostedReward = mockValMgr.calculateBoostedReward(validator);
        mockAssetMgr.increaseRewardPerKgh(validator, boostedReward);
        uint256 beforeBalance = assetToken.balanceOf(delegator);

        boostedReward = assetMgr.getKghReward(validator, delegator);

        vm.prank(delegator);
        assetMgr.claimKghReward(validator);

        assertEq(assetToken.balanceOf(delegator), beforeBalance + boostedReward);
    }

    function test_claimKghReward_zeroBoostedReward_reverts() external {
        _registerValidator(minActivateAmount);
        _delegateKgh(delegator, 0);

        vm.prank(delegator);
        vm.expectRevert(IAssetManager.InsufficientAsset.selector);
        assetMgr.claimKghReward(validator);
    }

    function test_bondValidatorKro_succeeds() public {
        _registerValidator(minActivateAmount);

        vm.prank(address(valMgr));
        assetMgr.bondValidatorKro(validator);

        assertEq(assetMgr.totalValidatorKro(validator), minActivateAmount);
        assertEq(assetMgr.totalValidatorKroBonded(validator), bondAmount);
    }

    function test_bondValidatorKro_callerNotValMgr_reverts() external {
        vm.expectRevert(IAssetManager.NotAllowedCaller.selector);
        assetMgr.bondValidatorKro(validator);
    }

    function test_bondValidatorKro_insufficientAsset_reverts() external {
        vm.prank(address(valMgr));
        vm.expectRevert(IAssetManager.InsufficientAsset.selector);
        assetMgr.bondValidatorKro(validator);
    }

    function test_unbondValidatorKro_succeeds() external {
        test_bondValidatorKro_succeeds();

        vm.prank(address(valMgr));
        assetMgr.unbondValidatorKro(validator);

        assertEq(assetMgr.totalValidatorKro(validator), minActivateAmount);
        assertEq(assetMgr.totalValidatorKroBonded(validator), 0);
    }

    function test_unbondValidatorKro_callerNotValMgr_reverts() external {
        vm.expectRevert(IAssetManager.NotAllowedCaller.selector);
        assetMgr.unbondValidatorKro(validator);
    }

    function test_increaseBalanceWithReward_succeeds() external {
        test_bondValidatorKro_succeeds();
        _delegateKgh(delegator, 0);

        uint256 beforeBalance = assetToken.balanceOf(address(assetMgr));
        uint256 beforeVaultBalance = assetToken.balanceOf(validatorRewardVault);

        uint128 boostedReward = 5e18;
        uint128 validatorReward = 10e18;
        uint128 totalReward = baseReward + boostedReward + validatorReward;

        vm.prank(address(valMgr));
        assetMgr.increaseBalanceWithReward(validator, baseReward, boostedReward, validatorReward);

        assertEq(assetToken.balanceOf(address(assetMgr)), beforeBalance + totalReward);
        assertEq(assetToken.balanceOf(validatorRewardVault), beforeVaultBalance - totalReward);
        assertEq(assetMgr.totalKroAssets(validator), baseReward);
        assertEq(assetMgr.totalValidatorKro(validator), minActivateAmount + validatorReward);
        assertEq(assetMgr.getKghReward(validator, delegator), boostedReward);
        assertEq(assetMgr.totalValidatorKroBonded(validator), 0);
    }

    function test_increaseBalanceWithReward_validatorIsSC_succeeds() external {
        uint256 beforeBalance = assetToken.balanceOf(address(assetMgr));
        uint256 beforeSCBalance = assetToken.balanceOf(guardian);
        uint256 beforeVaultBalance = assetToken.balanceOf(validatorRewardVault);

        uint128 boostedReward = 5e18;
        uint128 validatorReward = 10e18;
        uint128 totalReward = baseReward + boostedReward + validatorReward;

        vm.prank(address(valMgr));
        assetMgr.increaseBalanceWithReward(guardian, baseReward, boostedReward, validatorReward);

        assertEq(assetToken.balanceOf(address(assetMgr)), beforeBalance);
        assertEq(assetToken.balanceOf(guardian), beforeSCBalance + totalReward);
        assertEq(assetToken.balanceOf(validatorRewardVault), beforeVaultBalance - totalReward);
    }

    function test_increaseBalanceWithReward_callerNotValMgr_reverts() external {
        vm.expectRevert(IAssetManager.NotAllowedCaller.selector);
        assetMgr.increaseBalanceWithReward(validator, 0, 0, 0);
    }

    function test_increaseBalanceWithChallenge_succeeds() external {
        assetToken.mint(address(assetMgr), bondAmount);

        uint256 beforeSCBalance = assetToken.balanceOf(guardian);
        uint128 tax = (bondAmount * assetMgr.TAX_NUMERATOR()) / assetMgr.TAX_DENOMINATOR();

        vm.prank(address(valMgr));
        uint128 challengeReward = assetMgr.increaseBalanceWithChallenge(validator, bondAmount);

        assertEq(assetToken.balanceOf(guardian), beforeSCBalance + tax);
        assertEq(assetMgr.totalValidatorKro(validator), bondAmount - tax);
        assertEq(challengeReward, bondAmount - tax);
    }

    function test_increaseBalanceWithChallenge_winnerIsSC_succeeds() external {
        assetToken.mint(address(assetMgr), bondAmount);

        uint256 beforeSCBalance = assetToken.balanceOf(guardian);

        vm.prank(address(valMgr));
        uint128 challengeReward = assetMgr.increaseBalanceWithChallenge(guardian, bondAmount);

        assertEq(assetToken.balanceOf(guardian), beforeSCBalance + bondAmount);
        assertEq(challengeReward, bondAmount);
    }

    function test_increaseBalanceWithChallenge_callerNotValMgr_reverts() external {
        vm.expectRevert(IAssetManager.NotAllowedCaller.selector);
        assetMgr.increaseBalanceWithChallenge(validator, 0);
    }

    function test_decreaseBalanceWithChallenge_succeeds() external {
        test_bondValidatorKro_succeeds();

        vm.prank(address(valMgr));
        uint128 challengeReward = assetMgr.decreaseBalanceWithChallenge(validator);

        assertEq(assetMgr.totalValidatorKro(validator), minActivateAmount - bondAmount);
        assertEq(assetMgr.totalValidatorKroBonded(validator), 0);
        assertEq(challengeReward, bondAmount);
    }

    function test_decreaseBalanceWithChallenge_callerNotValMgr_reverts() external {
        vm.expectRevert(IAssetManager.NotAllowedCaller.selector);
        assetMgr.decreaseBalanceWithChallenge(validator);
    }

    function test_revertDecreaseBalanceWithChallenge_succeeds() external {
        vm.prank(address(valMgr));
        uint128 challengeReward = assetMgr.revertDecreaseBalanceWithChallenge(validator);

        assertEq(assetMgr.totalValidatorKro(validator), bondAmount);
        assertEq(assetMgr.totalValidatorKroBonded(validator), bondAmount);
        assertEq(challengeReward, bondAmount);
    }

    function test_revertDecreaseBalanceWithChallenge_callerNotValMgr_reverts() external {
        vm.expectRevert(IAssetManager.NotAllowedCaller.selector);
        assetMgr.revertDecreaseBalanceWithChallenge(validator);
    }
}
