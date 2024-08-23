// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { GovernanceToken } from "../governance/GovernanceToken.sol";
import { KromaVestingWallet } from "../universal/KromaVestingWallet.sol";
import { Proxy } from "../universal/Proxy.sol";
import { CommonTest } from "./CommonTest.t.sol";

contract KromaVestingWalletTest is CommonTest {
    uint64 immutable SECONDS_PER_3_MONTHS = 365 days / 4;

    KromaVestingWallet vestingWallet;
    address beneficiary;
    uint64 startTime;
    uint64 durationNum;
    uint64 durationSec;
    uint64 cliffDivider;
    uint64 vestingCycle;

    GovernanceToken token;
    address tokenOwner;
    uint256 totalAllocation;

    function setUp() public override {
        super.setUp();

        beneficiary = makeAddr("beneficiary");
        vestingCycle = SECONDS_PER_3_MONTHS;
        startTime = vestingCycle;
        durationNum = 12;
        durationSec = vestingCycle * durationNum;
        cliffDivider = 4;

        tokenOwner = makeAddr("tokenOwner");
        totalAllocation = 15 * 10_000_000 * (10 ** 18);

        vestingWallet = KromaVestingWallet(payable(address(new Proxy(multisig))));
        KromaVestingWallet vestingWalletImpl = new KromaVestingWallet(cliffDivider, vestingCycle);
        vm.prank(multisig);
        toProxy(address(vestingWallet)).upgradeToAndCall(
            address(vestingWalletImpl),
            abi.encodeCall(vestingWallet.initialize, (beneficiary, startTime, durationSec))
        );

        token = GovernanceToken(address(new Proxy(multisig)));
        GovernanceToken tokenImpl = new GovernanceToken(ZERO_ADDRESS, ZERO_ADDRESS);
        vm.prank(multisig);
        toProxy(address(token)).upgradeToAndCall(
            address(tokenImpl),
            abi.encodeCall(token.initialize, tokenOwner)
        );
        vm.prank(tokenOwner);
        token.acceptOwnership();

        vm.prank(tokenOwner);
        token.mint(address(vestingWallet), totalAllocation);
        assertEq(token.balanceOf(address(vestingWallet)), totalAllocation);
    }

    function test_constructor_succeeds() external {
        assertEq(vestingWallet.CLIFF_DIVIDER(), cliffDivider);
        assertEq(vestingWallet.VESTING_CYCLE(), vestingCycle);
    }

    function test_constructor_zeroValues_reverts() external {
        vm.expectRevert("KromaVestingWallet: cliff divider is zero");
        new KromaVestingWallet(0, vestingCycle);

        vm.expectRevert("KromaVestingWallet: vesting cycle is zero");
        new KromaVestingWallet(cliffDivider, 0);
    }

    function test_initialize_succeeds() external {
        assertEq(vestingWallet.beneficiary(), beneficiary);
        assertEq(vestingWallet.start(), startTime);
        assertEq(vestingWallet.duration(), durationSec);
    }

    function test_initialize_durationNotMultiple_reverts() external {
        vestingWallet = KromaVestingWallet(payable(address(new Proxy(multisig))));
        KromaVestingWallet vestingWalletImpl = new KromaVestingWallet(cliffDivider, vestingCycle);

        vm.prank(multisig);
        vm.expectRevert("Proxy: delegatecall to new implementation contract failed");
        toProxy(address(vestingWallet)).upgradeToAndCall(
            address(vestingWalletImpl),
            abi.encodeCall(vestingWallet.initialize, (beneficiary, startTime, durationSec + 1))
        );
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

    function test_release_notBeneficiary_reverts() external {
        vm.startPrank(tokenOwner);

        vm.expectRevert("KromaVestingWallet: caller is not beneficiary");
        vestingWallet.release(address(token));

        vm.expectRevert("KromaVestingWallet: caller is not beneficiary");
        vestingWallet.release();
    }
}
