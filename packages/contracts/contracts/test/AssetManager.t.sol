// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { ERC20 } from "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import { IERC721 } from "@openzeppelin/contracts/token/ERC721/IERC721.sol";

import { MockL2OutputOracle } from "./ValidatorPool.t.sol";
import { L2OutputOracle_Initializer } from "./CommonTest.t.sol";
import { BalancedWeightTree } from "../libraries/BalancedWeightTree.sol";
import { Constants } from "../libraries/Constants.sol";
import { Predeploys } from "../libraries/Predeploys.sol";
import { Types } from "../libraries/Types.sol";
import { Uint128Math } from "../libraries/Uint128Math.sol";
import { AssetManager } from "../L1/AssetManager.sol";
import { Colosseum } from "../L1/Colosseum.sol";
import { IKGHManager } from "../universal/IKGHManager.sol";
import { L2OutputOracle } from "../L1/L2OutputOracle.sol";
import { Proxy } from "../universal/Proxy.sol";
import { TestERC721 } from "./L1ERC721Bridge.t.sol";

contract MockKro is ERC20 {
    constructor() ERC20("Kroma", "KRO") {
        _mint(msg.sender, 1e27);
    }
}

contract MockKgh is TestERC721 {}

contract MockKghManager is IKGHManager {
    ERC20 public kro;

    constructor(ERC20 _kro) {
        kro = _kro;
    }

    function totalKroInKgh(uint256 /* tokenId */) external pure override returns (uint128) {
        return 100e18;
    }
}

contract MockAssetManager is AssetManager {
    using Uint128Math for uint128;
    using BalancedWeightTree for BalancedWeightTree.Tree;

    constructor(
        L2OutputOracle _l2OutputOracle,
        MockKro _token,
        IERC721 _kgh,
        IKGHManager _kghManager,
        address _securityCouncil,
        uint128 _maxOutputFinalizations,
        uint128 _baseReward,
        uint128 _slashingRateNumerator,
        uint128 _minSlashingAmount,
        uint128 _minRegisterAmount,
        uint128 _minStartAmount,
        uint256 _undelegationPeriod
    )
        AssetManager(
            _l2OutputOracle,
            _token,
            _kgh,
            _kghManager,
            _securityCouncil,
            _maxOutputFinalizations,
            _baseReward,
            _slashingRateNumerator,
            _minSlashingAmount,
            _minRegisterAmount,
            _minStartAmount,
            _undelegationPeriod
        )
    {}

    function submitOutput(uint256 _outputIndex, uint256 _expiresAt) external {
        _submittedOutputs[_outputIndex] = _expiresAt;
    }

    function modifyKghNum(address validator, uint128 amount) external {
        // We do not consider KROs in the KGH here, since this mock function
        // is only used for testing the boosted reward calculation.
        _vaults[validator].asset.totalKghShares += previewKghDelegate(validator) * amount;
        _vaults[validator].asset.totalKgh += amount;
    }

    function distributeReward() external {
        require(
            msg.sender == address(L2_ORACLE),
            "AssetManager: only oracle can distribute reward"
        );
        _distributeReward();
    }

    function getPendingKroReward(
        uint256 timestamp,
        address validator,
        address owner
    ) external view returns (uint128) {
        uint128 pendingShare = _vaults[validator].kroDelegators[owner].pendingKroShares[timestamp];
        uint128 pendingAsset = pendingShare.mulDiv(
            _vaults[validator].asset.totalPendingAssets,
            _vaults[validator].asset.totalPendingKroShares
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
            .kroShares;
        uint128 pendinKghShare = _vaults[validator]
            .kghDelegators[owner]
            .pendingShares[timestamp]
            .kghShares;
        uint128 pendingKroAsset = pendingKroShare.mulDiv(
            _vaults[validator].asset.totalPendingAssets,
            _vaults[validator].asset.totalPendingKroShares
        );
        uint128 pendingKghAsset = pendinKghShare.mulDiv(
            _vaults[validator].asset.totalPendingBoostedRewards,
            _vaults[validator].asset.totalPendingKghShares
        );
        return (pendingKroAsset, pendingKghAsset);
    }

    function totalKghAssets(address validator) public view virtual returns (uint128) {
        return _totalKghAssets(validator);
    }

    function insertToTree(address validator, uint128 kroAmount) external {
        _validatorTree.insert(validator, uint120(kroAmount));
    }

    function isValidatorRemoved() external view returns (bool) {
        return _validatorTree.removed == 1;
    }

    function setCommissionRate(address validator, uint8 _commissionRate) external {
        _vaults[validator].reward.commissionRate = _commissionRate;
    }
}

// Tests the implementations of the AssetManager
contract AssetManagerTest is L2OutputOracle_Initializer {
    using Types for *;
    using Constants for *;
    using Predeploys for *;

    MockAssetManager public assetManager;
    MockKro public kro;
    MockKgh public kgh;
    MockKghManager public kghManager;
    MockL2OutputOracle public mockOracle;
    uint128 public baseReward = 20e18;
    uint128 public maxOutputFinalizations = 10;
    address public validator = 0x000000000000000000000000000000000000AaaD;
    address public mockColosseum = 0x000000000000000000000000000000000000AAaE;
    address public delegator = 0x000000000000000000000000000000000000AAAF;

    function setUp() public override {
        super.setUp();

        MockL2OutputOracle mockOracleImpl = new MockL2OutputOracle(pool, address(mockColosseum));
        vm.prank(multisig);
        Proxy(payable(address(oracle))).upgradeTo(address(mockOracleImpl));
        mockOracle = MockL2OutputOracle(address(oracle));

        kro = new MockKro();
        kgh = new MockKgh();
        kghManager = new MockKghManager(kro);
        assetManager = new MockAssetManager({
            _l2OutputOracle: mockOracle,
            _token: kro,
            _kgh: kgh,
            _kghManager: kghManager,
            _securityCouncil: multisig,
            _maxOutputFinalizations: maxOutputFinalizations,
            _baseReward: baseReward,
            _slashingRateNumerator: 20, // 2%
            _minSlashingAmount: 1e18, // 1 KRO
            _minRegisterAmount: 10e18, // 10 KRO
            _minStartAmount: 100e18, // 100 KRO
            _undelegationPeriod: 7 days
        });

        // KRO bridged from L2 Validator Reward Vault
        kro.transfer(address(assetManager), 1e22);
    }

    function _submitOutputRoot(address _validator) internal {
        uint256 nextOutputIndex = mockOracle.nextOutputIndex();
        uint256 nextBlockNumber = mockOracle.nextBlockNumber();
        bytes32 outputRoot = keccak256(abi.encode(nextBlockNumber));

        warpToSubmitTime(nextBlockNumber);

        vm.prank(_validator);
        mockOracle.addOutput(outputRoot, nextBlockNumber);
        assetManager.submitOutput(nextOutputIndex, block.timestamp);
    }

    function _fillTokensForSlashing(
        uint128 kroAmount,
        uint256 asserterId,
        uint256 challengerId
    ) internal {
        kro.transfer(address(asserter), kroAmount);
        kro.transfer(address(challenger), kroAmount);

        vm.startPrank(asserter);
        kro.approve(address(assetManager), kroAmount);
        assetManager.delegate(asserter, kroAmount);
        if (asserterId != 0) {
            kgh.mint(asserter, asserterId);
            kgh.approve(address(assetManager), asserterId);
            assetManager.delegateKgh(asserter, asserterId);
        }
        vm.stopPrank();

        vm.startPrank(challenger);
        kro.approve(address(assetManager), kroAmount);
        assetManager.delegate(challenger, kroAmount);
        if (challengerId != 0) {
            kgh.mint(challenger, challengerId);
            kgh.approve(address(assetManager), challengerId);
            assetManager.delegateKgh(challenger, challengerId);
        }
        vm.stopPrank();
    }

    function _setUpKroDelegation(uint128 kroAmount) internal {
        kro.transfer(address(validator), kroAmount);
        kro.transfer(address(delegator), kroAmount);
        vm.startPrank(validator);
        kro.approve(address(assetManager), kroAmount);
        // self delegation
        assetManager.delegate(validator, kroAmount);
        vm.stopPrank();

        assetManager.insertToTree(validator, kroAmount);

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
        assetManager.delegate(validator, 100e18);
        vm.stopPrank();

        assetManager.insertToTree(validator, 100e18);

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
        assetManager.delegate(validator, 100e18);
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
        assertEq(address(assetManager.L2_ORACLE()), address(mockOracle));
        assertEq(address(assetManager.ASSET_TOKEN()), address(kro));
        assertEq(address(assetManager.KGH()), address(kgh));
        assertEq(assetManager.SECURITY_COUNCIL(), address(multisig));
        assertEq(assetManager.MAX_OUTPUT_FINALIZATIONS(), maxOutputFinalizations);
        assertEq(assetManager.BASE_REWARD(), baseReward);
        assertEq(assetManager.SLASHING_RATE_NUMERATOR(), 20);
        assertEq(assetManager.MIN_SLASHING_AMOUNT(), 1e18);
        assertEq(assetManager.MIN_REGISTER_AMOUNT(), 10e18);
        assertEq(assetManager.MIN_START_AMOUNT(), 100e18);
        assertEq(assetManager.UNDELEGATION_PERIOD(), 7 days);
    }

    function test_delegate_succeeds() external {
        _setUpKroDelegation(100e18);

        assertEq(assetManager.totalKroAssets(validator), 200e18);
    }

    function test_delegate_WithoutValidatorDelegation_reverts() external {
        vm.expectRevert("AssetManager: Vault is inactive");
        assetManager.delegate(validator, 100e18);
    }

    function test_delegateKgh_succeeds() external {
        (uint128 kroShares, uint128 kghShares) = _setUpKghDelegation(1);

        assertEq(kroShares, 1e26);
        assertEq(kghShares, 1e26);
        assertEq(assetManager.totalKghAssets(validator), 100e18);
    }

    function test_delegateKghBatch_succeeds() external {
        (uint128 kroShares, uint128 kghShares) = _setUpKghBatchDelegation(10);

        assertEq(kroShares, 1e27);
        assertEq(kghShares, 1e27);
        assertEq(assetManager.totalKghAssets(validator), 1_000e18);
    }

    function test_delegateKgh_WithoutValidatorDelegation_reverts() external {
        kgh.mint(delegator, 1);
        vm.expectRevert("AssetManager: Vault is inactive");
        vm.prank(delegator);
        assetManager.delegateKgh(validator, 1);
    }

    function test_initUndelegate_succeeds() external {
        assetManager.modifyKghNum(validator, 100);
        _setUpKroDelegation(100e18);
        _submitOutputRoot(validator);

        vm.prank(address(oracle));
        assetManager.distributeReward();

        vm.startPrank(delegator);
        // Fully undelegate
        assetManager.initUndelegate(validator, 1e26);
        vm.stopPrank();

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

        vm.prank(address(oracle));
        assetManager.distributeReward();

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

        vm.prank(address(oracle));
        assetManager.distributeReward();

        uint128 sharesToUndelegate = assetManager.getKroTotalShareBalance(validator, delegator);
        vm.startPrank(delegator);
        vm.expectRevert("AssetManager: Invalid amount of shares to undelegate");
        assetManager.initUndelegate(validator, sharesToUndelegate + 1);
    }

    function test_initUndelegate_removedFromWeightTree_succeeds() external {
        _setUpKroDelegation(100e18);

        uint128 kroShares = assetManager.getKroTotalShareBalance(validator, validator);
        vm.startPrank(validator);
        assetManager.initUndelegate(validator, kroShares);
        vm.stopPrank();

        assertEq(assetManager.totalKroAssets(validator), 100e18);
        assertEq(assetManager.isValidatorRemoved(), true);
    }

    function test_initUndelegateKgh_succeeds() external {
        assetManager.modifyKghNum(validator, 99);
        _setUpKghDelegation(100);
        _submitOutputRoot(validator);

        vm.prank(address(oracle));
        assetManager.distributeReward();

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

        vm.prank(address(oracle));
        assetManager.distributeReward();

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

    function test_initUndelegateKgh_ExactAmounts_succeeds() external {
        assetManager.modifyKghNum(validator, 99);
        _setUpKghDelegation(100);

        _submitOutputRoot(validator);

        vm.prank(address(oracle));
        assetManager.distributeReward();

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

    function test_initUndelegateKgh_NoShares_reverts() external {
        vm.expectRevert("AssetManager: No shares for the given tokenId");
        assetManager.initUndelegateKgh(validator, 1);
    }

    function test_initUndelegateKghBatch_NoShares_reverts() external {
        uint256[] memory tokenIds = new uint256[](1);
        tokenIds[0] = 1;
        vm.expectRevert("AssetManager: No shares for the given tokenId");
        assetManager.initUndelegateKghBatch(validator, tokenIds);
    }

    function test_initClaimValidatorReward_succeeds() external {
        _setUpKroDelegation(100e18);
        _submitOutputRoot(validator);
        // Set commission rate to 10%
        assetManager.setCommissionRate(validator, 10);

        vm.prank(address(oracle));
        assetManager.distributeReward();

        vm.startPrank(validator);
        assetManager.initClaimValidatorReward(2e18);
        vm.stopPrank();

        assertEq(assetManager.totalKroAssets(validator), 218e18);
    }

    function test_finalizeUndelegate_succeeds() external {
        assetManager.modifyKghNum(validator, 100);
        _setUpKroDelegation(100e18);
        _submitOutputRoot(validator);

        vm.prank(address(oracle));
        assetManager.distributeReward();

        uint128 sharesToUndelegate = assetManager.getKroTotalShareBalance(validator, delegator);
        vm.startPrank(delegator);
        assetManager.initUndelegate(validator, sharesToUndelegate);

        vm.warp(block.timestamp + 7 days);

        assetManager.finalizeUndelegate(validator);
        vm.stopPrank();

        assertEq(assetManager.totalKroAssets(validator), 110000000000000000001);
        assertEq(assetManager.ASSET_TOKEN().balanceOf(delegator), 109999999999999999999);
    }

    function test_finalizeUndelegateKgh_succeeds() external {
        _setUpKghDelegation(1);
        assertEq(assetManager.totalKghNum(validator), 1);

        vm.startPrank(delegator);
        assetManager.initUndelegateKgh(validator, 1);

        vm.warp(block.timestamp + 7 days);

        assetManager.finalizeUndelegateKgh(validator);
        vm.stopPrank();

        assertEq(assetManager.totalKghAssets(validator), 0);
        assertEq(kgh.balanceOf(delegator), 1);
    }

    function test_finalizeClaimValidatorReward_succeeds() external {
        _setUpKroDelegation(100e18);
        _submitOutputRoot(validator);
        // Set commission rate to 10%
        assetManager.setCommissionRate(validator, 10);

        vm.prank(address(oracle));
        assetManager.distributeReward();

        vm.startPrank(validator);
        assetManager.initClaimValidatorReward(2e18);
        vm.stopPrank();

        vm.warp(7 days + block.timestamp);

        vm.startPrank(validator);
        assetManager.finalizeClaimValidatorReward();
        vm.stopPrank();

        assertEq(assetManager.totalKroAssets(validator), 218e18);
        assertEq(assetManager.ASSET_TOKEN().balanceOf(validator), 2e18);
    }

    function test_finalizeUndelegate_WithNoPendingShares_reverts() external {
        vm.expectRevert("AssetManager: No pending shares to finalize");
        assetManager.finalizeUndelegate(validator);
    }

    function test_distributeReward_succeeds() external {
        assetManager.modifyKghNum(validator, 100);
        _submitOutputRoot(validator);

        vm.prank(address(oracle));
        assetManager.distributeReward();

        assertEq(assetManager.totalKroAssets(validator), baseReward);
        // 8 * arctan(0.01 * 100) * 1e18 will be calculated as 6283173600000736769
        assertEq(assetManager.totalKghAssets(validator) - 100 * 100e18, 6283173600000736769);
    }

    function test_slash_withSlashingRateNumerator_succeeds() external {
        _fillTokensForSlashing(100e18, 0, 0);
        _submitOutputRoot(asserter);

        vm.prank(mockColosseum);
        // Suppose that the challenge is successful, so the winner is challenger
        assetManager.slash(asserter, challenger, 0);
        // This will be done by the l2 output oracle contract in the real environment.
        vm.prank(address(challenger));
        mockOracle.replaceOutput(0, keccak256(abi.encode(1)));

        vm.warp(7 days + block.timestamp);

        vm.prank(address(oracle));
        assetManager.distributeReward();

        // slashing rate is 2%
        assertEq(assetManager.totalKroAssets(asserter), 98000000000000000000);
        assertEq(assetManager.totalKroAssets(challenger), 121600000000000000000);
        assertEq(assetManager.ASSET_TOKEN().balanceOf(multisig), 400000000000000000);
    }

    function test_slash_withMinSlashAmount_succeeds() external {
        _fillTokensForSlashing(20e18, 0, 0);
        _submitOutputRoot(asserter);

        vm.prank(mockColosseum);
        // Suppose that the challenge is successful, so the winner is challenger
        assetManager.slash(asserter, challenger, 0);
        // This will be done by the l2 output oracle contract in the real environment.
        vm.prank(address(challenger));
        mockOracle.replaceOutput(0, keccak256(abi.encode(1)));

        vm.roll(7 days + block.timestamp);

        vm.prank(address(oracle));
        assetManager.distributeReward();

        assertEq(assetManager.totalKroAssets(asserter), 19000000000000000000);
        assertEq(assetManager.totalKroAssets(challenger), 40800000000000000000);
        assertEq(assetManager.ASSET_TOKEN().balanceOf(multisig), 200000000000000000);
    }

    function test_slash_rewardSlashing_succeeds() external {
        _submitOutputRoot(asserter);
        _submitOutputRoot(challenger);

        assetManager.modifyKghNum(asserter, 100);
        assetManager.modifyKghNum(challenger, 100);

        vm.roll(7 days + block.timestamp);

        vm.prank(address(oracle));
        assetManager.distributeReward();

        // Assert that the reward is 6283173600000736769 except for the virtual KROs
        // generated by the KGHs.
        assertEq(assetManager.totalKghAssets(asserter) - 100 * 100e18, 6283173600000736769);
        assertEq(assetManager.totalKghAssets(challenger) - 100 * 100e18, 6283173600000736769);

        _submitOutputRoot(asserter);

        vm.prank(mockColosseum);
        // Suppose that the challenge is successful, so the winner is challenger.
        assetManager.slash(asserter, challenger, 2);
        // This will be done by the l2 output oracle contract in the real environment.
        vm.prank(address(challenger));
        mockOracle.replaceOutput(2, keccak256(abi.encode(1)));

        vm.roll(7 days + block.timestamp);

        vm.prank(address(oracle));
        // Slashed amount + base & boosted reward will go to the challenger.
        assetManager.distributeReward();

        // Total slashingAmount is 1e18.
        // KRO slashingAmount = slashingAmount * (baseReward / (baseReward + boostedReward))
        // = 760943115332139318
        // So asserter balance should be 19239056884667860682.
        assertEq(assetManager.totalKroAssets(asserter), 19239056884667860682);
        // Challenger balance should be 20608754492265711454 + 20000000000000000000 = 40608754492265711454,
        // with tax taken by security council.
        assertEq(assetManager.totalKroAssets(challenger), 40608754492265711454);

        // Total slashingAmount is 1e18.
        // KGH slashingAmount = slashingAmount * (boostedReward / (baseReward + boostedReward))
        // = 239056884667860682
        // And KGH reward amount is 6283173600000736769
        // So asserter balance should be 6044116715332876088.
        assertEq(assetManager.totalKghAssets(asserter) - 100 * 100e18, 6044116715332876088);
        // Challenger balance should be 6474419107735025314 + 6283173600000736769 = 12757592707735762083,
        // with tax taken by security council.
        assertEq(assetManager.totalKghAssets(challenger) - 100 * 100e18, 12757592707735762083);
        assertEq(assetManager.ASSET_TOKEN().balanceOf(multisig), 200000000000000000);
    }

    function test_slash_notColosseum_reverts() external {
        vm.prank(address(1));
        vm.expectRevert("AssetManager: Only Colosseum can call this function");
        assetManager.slash(asserter, challenger, 1);
    }
}
