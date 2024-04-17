// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { ERC20 } from "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import { IERC20 } from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import { IERC721 } from "@openzeppelin/contracts/token/ERC721/IERC721.sol";

import { Constants } from "../libraries/Constants.sol";
import { Predeploys } from "../libraries/Predeploys.sol";
import { Types } from "../libraries/Types.sol";
import { Uint128Math } from "../libraries/Uint128Math.sol";
import { AssetManager } from "../L1/AssetManager.sol";
import { IValidatorManager } from "../L1/IValidatorManager.sol";
import { IKGHManager } from "../universal/IKGHManager.sol";
import { Proxy } from "../universal/Proxy.sol";
import { L2OutputOracle_ValidatorSystemUpgrade_Initializer } from "./CommonTest.t.sol";
import { MockL2OutputOracle } from "./ValidatorManager.t.sol";

contract MockKro is ERC20 {
    constructor() ERC20("Kroma", "KRO") {
        _mint(msg.sender, 1e27);
    }
}

contract MockAssetManager is AssetManager {
    using Uint128Math for uint128;

    constructor(
        IERC20 _assetToken,
        IERC721 _kgh,
        IKGHManager _kghManager,
        address _securityCouncil,
        IValidatorManager _validatorManager,
        uint128 _undelegationPeriod,
        uint128 _slashingRate,
        uint128 _minSlashingAmount
    )
        AssetManager(
            _assetToken,
            _kgh,
            _kghManager,
            _securityCouncil,
            _validatorManager,
            _undelegationPeriod,
            _slashingRate,
            _minSlashingAmount
        )
    {}

    function modifyKghNum(address validator, uint128 amount) external {
        // We do not consider KROs in the KGH here, since this mock function
        // is only used for testing the boosted reward calculation.
        _vaults[validator].asset.totalKghShares += previewKghDelegate(validator) * amount;
        _vaults[validator].asset.totalKgh += amount;
    }

    function getPendingKroReward(
        uint256 timestamp,
        address validator,
        address owner
    ) external view returns (uint128) {
        uint128 pendingShare = _vaults[validator].kroDelegators[owner].pendingKroShares[timestamp];
        uint128 pendingAsset = pendingShare.mulDiv(
            _vaults[validator].pending.totalPendingAssets,
            _vaults[validator].pending.totalPendingKroShares
        );
        return pendingAsset;
    }

    function getPendingKghReward(
        uint256 timestamp,
        address validator,
        address owner
    ) external view returns (uint128, uint128) {
        uint128 pendingKroShare = _vaults[validator]
            .kghDelegators[owner]
            .pendingShares[timestamp]
            .kro;
        uint128 pendinKghShare = _vaults[validator]
            .kghDelegators[owner]
            .pendingShares[timestamp]
            .kgh;
        uint128 pendingKroAsset = pendingKroShare.mulDiv(
            _vaults[validator].pending.totalPendingAssets,
            _vaults[validator].pending.totalPendingKroShares
        );
        uint128 pendingKghAsset = pendinKghShare.mulDiv(
            _vaults[validator].pending.totalPendingBoostedRewards,
            _vaults[validator].pending.totalPendingKghShares
        );
        return (pendingKroAsset, pendingKghAsset);
    }

    function totalKghAssets(address validator) public view virtual returns (uint128) {
        return _totalKghAssets(validator);
    }
}

// Tests the implementations of the AssetManager
contract AssetManagerTest is L2OutputOracle_ValidatorSystemUpgrade_Initializer {
    using Types for *;
    using Constants for *;
    using Predeploys for *;

    MockAssetManager public assetManager;
    MockAssetManager public assetManagerImpl;
    MockKro public kro;
    MockL2OutputOracle public mockOracle;
    address public validator = trusted;
    address public delegator = 0x000000000000000000000000000000000000AAAF;
    uint128 public VKRO_PER_KGH;

    function setUp() public override {
        super.setUp();

        assetManager = MockAssetManager(address(new Proxy(multisig)));
        vm.label(address(assetManager), "AssetManager");

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
        Proxy(payable(address(oracle))).upgradeTo(address(mockOracleImpl));
        mockOracle = MockL2OutputOracle(address(oracle));

        kro = new MockKro();
        assetManagerImpl = new MockAssetManager(
            kro,
            kgh,
            kghManager,
            address(guardian),
            valMan,
            uint128(undelegationPeriod),
            slashingRate,
            minSlashingAmount
        );

        address assetManagerAddr = address(assetMan);

        vm.prank(multisig);
        Proxy(payable(assetManagerAddr)).upgradeTo(address(assetManagerImpl));
        assetManager = MockAssetManager(assetManagerAddr);

        VKRO_PER_KGH = assetManager.VKRO_PER_KGH();

        // KRO bridged from L2 Validator Reward Vault
        kro.transfer(address(assetManager), 1e22);

        // submit until terminateOutputIndex and set it latest finalized output
        for (uint256 i = mockOracle.nextOutputIndex(); i <= terminateOutputIndex; i++) {
            _submitOutputRoot(pool.nextValidator());
        }
        vm.warp(mockOracle.finalizedAt(terminateOutputIndex));
        mockOracle.mockSetLatestFinalizedOutputIndex(terminateOutputIndex);
    }

    function _submitOutputRoot(address _validator) internal {
        uint256 nextBlockNumber = mockOracle.nextBlockNumber();
        warpToSubmitTime();
        vm.prank(_validator);
        mockOracle.addOutput(nextBlockNumber);
    }

    function _setUpKroDelegation(uint128 kroAmount) internal {
        kro.transfer(address(validator), kroAmount);
        kro.transfer(address(delegator), kroAmount);
        vm.startPrank(validator);
        kro.approve(address(assetManager), kroAmount);
        // self delegation
        valMan.registerValidator(kroAmount, 0, 10);
        vm.stopPrank();

        vm.startPrank(delegator);
        kro.approve(address(assetManager), kroAmount);
        assetManager.delegate(validator, kroAmount);
        vm.stopPrank();
    }

    function _setUpKghDelegation(uint256 tokenId) internal returns (uint128, uint128) {
        kro.transfer(address(validator), 100e18);

        vm.startPrank(validator);
        kro.approve(address(assetManager), 100e18);
        // self delegation
        valMan.registerValidator(100e18, 0, 0);
        vm.stopPrank();

        kgh.mint(delegator, tokenId);
        vm.startPrank(delegator);
        kgh.approve(address(assetManager), tokenId);
        (uint128 kroShares, uint128 kghShares) = assetManager.delegateKgh(validator, tokenId);
        vm.stopPrank();
        return (kroShares, kghShares);
    }

    function _setUpKghBatchDelegation(uint256 kghCounts) internal returns (uint128, uint128) {
        kro.transfer(address(validator), 100e18);
        vm.startPrank(validator);
        kro.approve(address(assetManager), 100e18);
        valMan.registerValidator(100e18, 0, 10);
        vm.stopPrank();

        uint256[] memory tokenIds = new uint256[](kghCounts);
        for (uint256 i = 1; i < kghCounts + 1; i++) {
            kgh.mint(delegator, i);
            vm.prank(delegator);
            kgh.approve(address(assetManager), i);
            tokenIds[i - 1] = i;
        }

        vm.prank(delegator);
        (uint128 kroShares, uint128 kghShares) = assetManager.delegateKghBatch(validator, tokenIds);
        return (kroShares, kghShares);
    }

    function test_constructor_succeeds() external {
        assertEq(address(assetManager.ASSET_TOKEN()), address(kro));
        assertEq(address(assetManager.KGH()), address(kgh));
        assertEq(address(assetManager.KGH_MANAGER()), address(kghManager));
        assertEq(assetManager.SECURITY_COUNCIL(), address(guardian));
        assertEq(address(assetManager.VALIDATOR_MANAGER()), address(valMan));
        assertEq(assetManager.UNDELEGATION_PERIOD(), undelegationPeriod);
        assertEq(assetManager.SLASHING_RATE(), slashingRate);
        assertEq(assetManager.MIN_SLASHING_AMOUNT(), minSlashingAmount);
    }

    function test_constructor_largeSlashingRate_reverts() external {
        vm.expectRevert(AssetManager.InvalidConstructorParams.selector);
        new MockAssetManager(
            kro,
            kgh,
            kghManager,
            address(guardian),
            valMan,
            uint128(undelegationPeriod),
            1001,
            minSlashingAmount
        );
    }

    function test_delegate_succeeds() external {
        _setUpKroDelegation(100e18);

        assertEq(assetManager.totalKroAssets(validator), 200e18);
    }

    function test_delegate_withoutValidatorDelegation_reverts() external {
        vm.expectRevert(AssetManager.ImproperValidatorStatus.selector);
        assetManager.delegate(validator, 100e18);
    }

    function test_delegateKgh_succeeds() external {
        (uint128 kroShares, uint128 kghShares) = _setUpKghDelegation(1);

        assertEq(kroShares, 1e26);
        assertEq(kghShares, 1e26);
        assertEq(assetManager.totalKghAssets(validator), VKRO_PER_KGH);
    }

    function test_delegateKghBatch_succeeds() external {
        (uint128 kroShares, uint128 kghShares) = _setUpKghBatchDelegation(10);

        assertEq(kroShares, 1e27);
        assertEq(kghShares, 1e27);
        assertEq(assetManager.totalKghAssets(validator), VKRO_PER_KGH * 10);
    }

    function test_delegateKgh_withoutValidatorDelegation_reverts() external {
        kgh.mint(delegator, 1);
        vm.expectRevert(AssetManager.ImproperValidatorStatus.selector);
        vm.prank(delegator);
        assetManager.delegateKgh(validator, 1);
    }

    function test_initUndelegate_succeeds() public {
        assetManager.modifyKghNum(validator, 100);
        _setUpKroDelegation(100e18);
        _submitOutputRoot(validator);
        vm.warp(mockOracle.finalizedAt(mockOracle.latestOutputIndex()));

        vm.startPrank(address(mockOracle));
        valMan.afterSubmitL2Output(mockOracle.latestOutputIndex());
        vm.stopPrank();

        // Fully undelegate
        uint128 sharesToUndelegate = assetManager.getKroTotalShareBalance(validator, delegator);
        vm.prank(delegator);
        assetManager.initUndelegate(validator, sharesToUndelegate);

        uint128 pendingAssets = assetManager.getPendingKroReward(
            block.timestamp,
            validator,
            delegator
        );

        assertEq(assetManager.totalKroAssets(validator), 110000000000000000001);
        assertEq(pendingAssets, 109999999999999999999);
    }

    function test_initUndelegate_exactAmount_succeeds() external {
        assetManager.modifyKghNum(validator, 100);
        _setUpKroDelegation(9_990e18);

        address delegator3 = makeAddr("delegator3");
        kro.transfer(address(delegator3), 20e18);
        vm.startPrank(delegator3);
        kro.approve(address(assetManager), 20e18);
        assetManager.delegate(validator, 20e18);
        vm.stopPrank();

        _submitOutputRoot(validator);
        vm.warp(mockOracle.finalizedAt(mockOracle.latestOutputIndex()));

        vm.startPrank(address(mockOracle));
        valMan.afterSubmitL2Output(mockOracle.latestOutputIndex());
        vm.stopPrank();

        uint128 sharesToUndelegate = assetManager.getKroTotalShareBalance(validator, delegator3);
        vm.prank(delegator3);
        assetManager.initUndelegate(validator, sharesToUndelegate);

        uint128 pendingAssets = assetManager.getPendingKroReward(
            block.timestamp,
            validator,
            delegator3
        );

        // Total KRO was 20,000 and 20 KRO was undelegated. So the reward that delegator3 can take
        // is 20 * 20 / 20,000 = 0.02.
        assertEq(assetManager.totalKroAssets(validator), 19999980000000000000001);
        assertEq(pendingAssets, 20019999999999999999);
    }

    function test_initUndelegate_exceedsMaxAmount_reverts() external {
        assetManager.modifyKghNum(validator, 100);
        _setUpKroDelegation(100e18);
        _submitOutputRoot(validator);

        vm.startPrank(address(mockOracle));
        valMan.afterSubmitL2Output(mockOracle.latestOutputIndex());
        vm.stopPrank();

        uint128 sharesToUndelegate = assetManager.getKroTotalShareBalance(validator, delegator);
        vm.startPrank(delegator);
        vm.expectRevert(AssetManager.InsufficientShare.selector);
        assetManager.initUndelegate(validator, sharesToUndelegate + 1);
    }

    function test_initUndelegate_removedFromValidatorTree_succeeds() external {
        _setUpKroDelegation(minStartAmount);

        uint128 kroShares = assetManager.getKroTotalShareBalance(validator, delegator);
        vm.prank(delegator);
        assetManager.initUndelegate(validator, kroShares);

        uint128 minUndelegateShares = assetManager.previewDelegate(validator, 1);
        vm.prank(validator);
        assetManager.initUndelegate(validator, minUndelegateShares);

        assertEq(assetManager.totalKroAssets(validator), minStartAmount - 1);
    }

    function test_initUndelegateKgh_succeeds() external {
        assetManager.modifyKghNum(validator, 99);
        _setUpKghDelegation(100);
        _submitOutputRoot(validator);
        vm.warp(mockOracle.finalizedAt(mockOracle.latestOutputIndex()));

        vm.startPrank(address(mockOracle));
        valMan.afterSubmitL2Output(mockOracle.latestOutputIndex());
        vm.stopPrank();

        vm.startPrank(delegator);
        assetManager.initUndelegateKgh(validator, 100);
        vm.stopPrank();

        (, uint128 pendingAssets) = assetManager.getPendingKghReward(
            block.timestamp,
            validator,
            delegator
        );

        // Total boosted reward is 6283173600000736769.
        assertEq(pendingAssets, 62831736000007367);
    }

    function test_initUndelegateKghBatch_succeeds() external {
        _setUpKghBatchDelegation(100);
        _submitOutputRoot(validator);
        vm.warp(mockOracle.finalizedAt(mockOracle.latestOutputIndex()));

        vm.startPrank(address(mockOracle));
        valMan.afterSubmitL2Output(mockOracle.latestOutputIndex());
        vm.stopPrank();

        uint256[] memory tokenIds = new uint256[](100);
        for (uint256 i = 0; i < 100; i++) {
            tokenIds[i] = i + 1;
        }
        vm.prank(delegator);
        assetManager.initUndelegateKghBatch(validator, tokenIds);

        (, uint128 pendingAssets) = assetManager.getPendingKghReward(
            block.timestamp,
            validator,
            delegator
        );

        // Total boosted reward is 6283173600000736769.
        assertEq(pendingAssets, 6283173600000736700);
    }

    function test_initUndelegateKgh_exactAmounts_succeeds() external {
        assetManager.modifyKghNum(validator, 99);
        _setUpKghDelegation(100);

        _submitOutputRoot(validator);
        vm.warp(mockOracle.finalizedAt(mockOracle.latestOutputIndex()));

        vm.startPrank(address(mockOracle));
        valMan.afterSubmitL2Output(mockOracle.latestOutputIndex());
        vm.stopPrank();

        vm.startPrank(delegator);
        assetManager.initUndelegateKgh(validator, 100);
        vm.stopPrank();

        (uint128 pendingKroAssets, uint128 pendingKghAssets) = assetManager.getPendingKghReward(
            block.timestamp,
            validator,
            delegator
        );

        assertEq(pendingKroAssets, 9999999999999999999);
        // Total boosted reward is 6283173600000736769.
        assertEq(pendingKghAssets, 62831736000007367);
    }

    function test_initUndelegateKgh_noShares_reverts() external {
        vm.expectRevert(AssetManager.ShareNotExists.selector);
        assetManager.initUndelegateKgh(validator, 1);
    }

    function test_initUndelegateKghBatch_noShares_reverts() external {
        uint256[] memory tokenIds = new uint256[](1);
        tokenIds[0] = 1;
        vm.expectRevert(AssetManager.ShareNotExists.selector);
        assetManager.initUndelegateKghBatch(validator, tokenIds);
    }

    function test_initClaimValidatorReward_succeeds() public {
        _setUpKroDelegation(100e18);
        _submitOutputRoot(validator);
        vm.warp(mockOracle.finalizedAt(mockOracle.latestOutputIndex()));

        vm.warp(block.timestamp + commissionRateMinChangeSeconds);
        // Set commission rate to 10%
        vm.prank(validator);
        valMan.changeCommissionRate(10);

        vm.startPrank(address(mockOracle));
        valMan.afterSubmitL2Output(mockOracle.latestOutputIndex());
        vm.stopPrank();

        vm.startPrank(validator);
        assetManager.initClaimValidatorReward(2e18);
        vm.stopPrank();

        assertEq(assetManager.totalKroAssets(validator), 218e18);
    }

    function test_finalizeUndelegate_succeeds() external {
        test_initUndelegate_succeeds();

        vm.warp(block.timestamp + undelegationPeriod);

        vm.prank(delegator);
        assetManager.finalizeUndelegate(validator);

        assertEq(assetManager.totalKroAssets(validator), 110000000000000000001);
        assertEq(assetManager.ASSET_TOKEN().balanceOf(delegator), 109999999999999999999);
    }

    function test_finalizeUndelegate_zeroRequest_reverts() external {
        test_initUndelegate_succeeds();

        vm.warp(block.timestamp + undelegationPeriod);

        vm.prank(validator);
        vm.expectRevert(AssetManager.RequestNotExists.selector);
        assetManager.finalizeUndelegate(validator);
    }

    function test_finalizeUndelegate_undelegationPeriodNotElapsed_reverts() external {
        test_initUndelegate_succeeds();

        vm.prank(delegator);
        vm.expectRevert(AssetManager.FinalizedPendingNotExists.selector);
        assetManager.finalizeUndelegate(validator);
    }

    function test_finalizeUndelegateKgh_noReward_succeeds() external {
        _setUpKghDelegation(1);
        assertEq(assetManager.totalKghNum(validator), 1);

        vm.startPrank(delegator);
        assetManager.initUndelegateKgh(validator, 1);

        vm.warp(block.timestamp + undelegationPeriod);

        assetManager.finalizeUndelegateKgh(validator);
        vm.stopPrank();

        assertEq(assetManager.totalKghAssets(validator), 0);
        assertEq(kgh.balanceOf(delegator), 1);
        assertEq(assetManager.ASSET_TOKEN().balanceOf(delegator), 0);
    }

    function test_finalizeUndelegateKgh_rewardExists_succeeds() external {
        _setUpKghDelegation(1);
        assertEq(assetManager.totalKghNum(validator), 1);

        _submitOutputRoot(validator);
        vm.warp(mockOracle.finalizedAt(mockOracle.latestOutputIndex()));

        vm.startPrank(address(mockOracle));
        valMan.afterSubmitL2Output(mockOracle.latestOutputIndex());
        vm.stopPrank();

        vm.startPrank(delegator);
        assetManager.initUndelegateKgh(validator, 1);

        vm.warp(block.timestamp + undelegationPeriod);

        assetManager.finalizeUndelegateKgh(validator);
        vm.stopPrank();

        assertEq(assetManager.totalKghNum(validator), 0);
        assertEq(kgh.balanceOf(delegator), 1);
        assertTrue(assetManager.ASSET_TOKEN().balanceOf(delegator) > 0);
    }

    function test_finalizeUndelegateKgh_undelegationPeriodNotElapsed_reverts() external {
        _setUpKghDelegation(1);
        assertEq(assetManager.totalKghNum(validator), 1);

        vm.startPrank(delegator);
        assetManager.initUndelegateKgh(validator, 1);

        vm.expectRevert(AssetManager.FinalizedPendingNotExists.selector);
        assetManager.finalizeUndelegateKgh(validator);
    }

    function test_finalizeClaimValidatorReward_succeeds() external {
        test_initClaimValidatorReward_succeeds();

        vm.warp(undelegationPeriod + block.timestamp);

        vm.startPrank(validator);
        assetManager.finalizeClaimValidatorReward();
        vm.stopPrank();

        assertEq(assetManager.totalKroAssets(validator), 218e18);
        assertEq(assetManager.ASSET_TOKEN().balanceOf(validator), 2e18);
    }

    function test_finalizeUndelegate_withNoPendingShares_reverts() external {
        vm.expectRevert(AssetManager.PendingNotExists.selector);
        assetManager.finalizeUndelegate(validator);
    }
}
