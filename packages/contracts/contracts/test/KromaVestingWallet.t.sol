// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { GovernanceToken } from "../governance/GovernanceToken.sol";
import { KromaVestingWallet } from "../universal/KromaVestingWallet.sol";
import { Proxy } from "../universal/Proxy.sol";
import { CommonTest } from "./CommonTest.t.sol";

contract KromaVestingWalletTest is CommonTest {
    uint64 immutable SECONDS_PER_3_MONTHS = (60 * 60 * 24 * 365) / 4;

    KromaVestingWallet vestingWallet;
    address beneficiary;
    uint64 startTime;
    uint64 durationNum;
    uint64 durationSec;
    uint64 cliffDivider;
    uint64 vestingCycle;
    address vestingWalletOwner;

    GovernanceToken token;
    address tokenOwner;
    uint256 totalAllocation;

    function setUp() public override {
        super.setUp();

        beneficiary = makeAddr("beneficiary");
        vestingWalletOwner = makeAddr("vestingWalletOwner");
        vestingCycle = SECONDS_PER_3_MONTHS;
        startTime = vestingCycle;
        durationNum = 12;
        durationSec = vestingCycle * durationNum;
        cliffDivider = 4;

        tokenOwner = makeAddr("tokenOwner");
        totalAllocation = 15 * 10_000_000 * (10 ** 18);

        vestingWallet = KromaVestingWallet(payable(address(new Proxy(multisig))));
        KromaVestingWallet vestingWalletImpl = new KromaVestingWallet();
        vm.prank(multisig);
        toProxy(address(vestingWallet)).upgradeToAndCall(
            address(vestingWalletImpl),
            abi.encodeCall(
                vestingWallet.initialize,
                (
                    beneficiary,
                    startTime,
                    durationSec,
                    cliffDivider,
                    vestingCycle,
                    vestingWalletOwner
                )
            )
        );

        token = GovernanceToken(address(new Proxy(multisig)));
        GovernanceToken tokenImpl = new GovernanceToken(ZERO_ADDRESS, ZERO_ADDRESS);
        vm.prank(multisig);
        toProxy(address(token)).upgradeToAndCall(
            address(tokenImpl),
            abi.encodeCall(token.initialize, tokenOwner)
        );

        vm.prank(tokenOwner);
        token.mint(address(vestingWallet), totalAllocation);
        assertEq(token.balanceOf(address(vestingWallet)), totalAllocation);
    }

    function test_initialize_succeeds() external {
        assertEq(vestingWallet.cliffDivider(), cliffDivider);
        assertEq(vestingWallet.vestingCycle(), vestingCycle);
        assertEq(vestingWallet.start(), startTime);
        assertEq(vestingWallet.duration(), durationSec);
        assertEq(vestingWallet.beneficiary(), beneficiary);
        assertEq(vestingWallet.owner(), vestingWalletOwner);
    }

    function test_setBeneficiary_owner_succeeds() external {
        address newBeneficiary = makeAddr("newBeneficary");
        vm.prank(vestingWalletOwner);
        vestingWallet.setBeneficiary(newBeneficiary);
    }

    function test_setBeneficiary_beneficiary_succeeds() external {
        address newBeneficiary = makeAddr("newBeneficary");
        vm.prank(beneficiary);
        vestingWallet.setBeneficiary(newBeneficiary);
    }

    function test_setBeneficiary_randomAddress_reverts() external {
        address newBeneficiary = makeAddr("newBeneficary");
        vm.prank(tokenOwner);
        vm.expectRevert("KromaVestingWallet: caller is not beneficiary or owner");
        vestingWallet.setBeneficiary(newBeneficiary);
    }

    function test_setBeneficiary_zeroBeneficiary_reverts() external {
        vm.prank(vestingWalletOwner);
        vm.expectRevert("KromaVestingWallet: beneficiary is zero address");
        vestingWallet.setBeneficiary(ZERO_ADDRESS);
    }

    function test_release_token_succeeds() external {
        // Ensure test env is set properly
        assertEq(token.balanceOf(beneficiary), 0);
        assertTrue(block.timestamp < startTime);

        uint256 cliffAmount = totalAllocation / cliffDivider;
        uint256 vestingAmountPer3Months = (totalAllocation - cliffAmount) / durationNum;

        vm.startPrank(beneficiary);
        vestingWallet.release(address(token));
        assertEq(token.balanceOf(beneficiary), 0);

        vm.warp(startTime);
        vestingWallet.release(address(token));
        assertEq(token.balanceOf(beneficiary), cliffAmount);

        vm.warp(startTime + vestingCycle / 2);
        vestingWallet.release(address(token));
        assertEq(token.balanceOf(beneficiary), cliffAmount);

        vm.warp(startTime + vestingCycle);
        vestingWallet.release(address(token));
        assertEq(token.balanceOf(beneficiary), cliffAmount + vestingAmountPer3Months);

        vm.warp(startTime + durationSec);
        vestingWallet.release(address(token));
        assertEq(token.balanceOf(beneficiary), totalAllocation);
    }

    function test_release_tokenAfterFullyVested_succeeds() external {
        // Ensure test env is set properly
        assertEq(token.balanceOf(beneficiary), 0);

        vm.warp(startTime + durationSec + 1);

        vm.prank(beneficiary);
        vestingWallet.release(address(token));
        assertEq(token.balanceOf(beneficiary), totalAllocation);
    }

    function test_release_tokenWhenBeneficiaryChanged_succeeds() external {
        // Ensure test env is set properly
        assertEq(token.balanceOf(beneficiary), 0);

        uint256 cliffAmount = totalAllocation / cliffDivider;
        uint256 vestingAmountPer3Months = (totalAllocation - cliffAmount) / durationNum;

        vm.warp(startTime + vestingCycle);

        vm.startPrank(beneficiary);
        vestingWallet.release(address(token));
        assertEq(token.balanceOf(beneficiary), cliffAmount + vestingAmountPer3Months);

        address newBeneficary = makeAddr("newBeneficary");
        vestingWallet.setBeneficiary(newBeneficary);
        vm.stopPrank();

        vm.warp(startTime + vestingCycle * 2);
        vm.prank(newBeneficary);
        vestingWallet.release(address(token));
        assertEq(token.balanceOf(beneficiary), cliffAmount + vestingAmountPer3Months);
        assertEq(token.balanceOf(newBeneficary), vestingAmountPer3Months);
    }

    function test_release_succeeds() external {
        uint256 totalEthAllocation = 1 ether;
        vm.deal(address(vestingWallet), totalEthAllocation);

        // Ensure test env is set properly
        assertEq(beneficiary.balance, 0);
        assertTrue(block.timestamp < startTime);

        uint256 cliffAmount = totalEthAllocation / cliffDivider;
        uint256 vestingAmountPer3Months = (totalEthAllocation - cliffAmount) / durationNum;

        vm.startPrank(beneficiary);
        vestingWallet.release();
        assertEq(beneficiary.balance, 0);

        vm.warp(startTime);
        vestingWallet.release();
        assertEq(beneficiary.balance, cliffAmount);

        vm.warp(startTime + vestingCycle / 2);
        vestingWallet.release();
        assertEq(beneficiary.balance, cliffAmount);

        vm.warp(startTime + vestingCycle);
        vestingWallet.release();
        assertEq(beneficiary.balance, cliffAmount + vestingAmountPer3Months);

        vm.warp(startTime + durationSec);
        vestingWallet.release();
        assertEq(beneficiary.balance, totalEthAllocation);
    }

    function test_release_afterFullyVested_succeeds() external {
        uint256 totalEthAllocation = 1 ether;
        vm.deal(address(vestingWallet), totalEthAllocation);

        // Ensure test env is set properly
        assertEq(beneficiary.balance, 0);

        vm.warp(startTime + durationSec + 1);

        vm.prank(beneficiary);
        vestingWallet.release();
        assertEq(beneficiary.balance, totalEthAllocation);
    }

    function test_release_whenBeneficiaryChanged_succeeds() external {
        uint256 totalEthAllocation = 1 ether;
        vm.deal(address(vestingWallet), totalEthAllocation);

        // Ensure test env is set properly
        assertEq(beneficiary.balance, 0);

        uint256 cliffAmount = totalEthAllocation / cliffDivider;
        uint256 vestingAmountPer3Months = (totalEthAllocation - cliffAmount) / durationNum;

        vm.warp(startTime + vestingCycle);

        vm.startPrank(beneficiary);
        vestingWallet.release();
        assertEq(beneficiary.balance, cliffAmount + vestingAmountPer3Months);

        address newBeneficary = makeAddr("newBeneficary");
        vestingWallet.setBeneficiary(newBeneficary);
        vm.stopPrank();

        vm.warp(startTime + vestingCycle * 2);
        vm.prank(newBeneficary);
        vestingWallet.release();
        assertEq(beneficiary.balance, cliffAmount + vestingAmountPer3Months);
        assertEq(newBeneficary.balance, vestingAmountPer3Months);
    }

    function test_release_notBeneficiary_reverts() external {
        vm.startPrank(vestingWalletOwner);

        vm.expectRevert("KromaVestingWallet: caller is not beneficiary");
        vestingWallet.release(address(token));

        vm.expectRevert("KromaVestingWallet: caller is not beneficiary");
        vestingWallet.release();
    }

    function test_migrateTokensToNewWallet_succeeds() external {
        KromaVestingWallet newVestingWallet = new KromaVestingWallet();

        vm.prank(vestingWalletOwner);
        vestingWallet.migrateTokensToNewWallet(address(token), address(newVestingWallet));

        assertEq(token.balanceOf(address(vestingWallet)), 0);
        assertEq(token.balanceOf(address(newVestingWallet)), totalAllocation);
    }

    function test_migrateEthToNewWallet_succeeds() external {
        uint256 totalEthAllocation = 1 ether;
        vm.deal(address(vestingWallet), totalEthAllocation);

        KromaVestingWallet newVestingWallet = new KromaVestingWallet();

        vm.prank(vestingWalletOwner);
        vestingWallet.migrateEthToNewWallet(address(newVestingWallet));

        assertEq(address(vestingWallet).balance, 0);
        assertEq(address(newVestingWallet).balance, totalEthAllocation);
    }

    function test_migrateTokensToNewWallet_notOwner_reverts() external {
        vm.startPrank(beneficiary);

        vm.expectRevert("Ownable: caller is not the owner");
        vestingWallet.migrateTokensToNewWallet(address(token), beneficiary);

        vm.expectRevert("Ownable: caller is not the owner");
        vestingWallet.migrateEthToNewWallet(beneficiary);
    }

    function test_migrateTokensToNewWallet_targetNotContract_reverts() external {
        vm.startPrank(vestingWalletOwner);

        vm.expectRevert("KromaVestingWallet: new wallet must be a contract");
        vestingWallet.migrateTokensToNewWallet(address(token), beneficiary);

        vm.expectRevert("KromaVestingWallet: new wallet must be a contract");
        vestingWallet.migrateEthToNewWallet(beneficiary);
    }
}
