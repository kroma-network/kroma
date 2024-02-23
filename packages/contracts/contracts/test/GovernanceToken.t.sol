// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

// Testing utilities
import { CommonTest } from "./CommonTest.t.sol";

// Target contract
import { GovernanceToken } from "../governance/GovernanceToken.sol";

contract GovernanceToken_Test is CommonTest {
    address owner;
    address rando;
    GovernanceToken governanceToken;
    address bridge;
    address remoteToken;
    address mintManager;

    /// @dev Sets up the test suite.
    function setUp() public virtual override {
        super.setUp();

        owner = makeAddr("owner");
        rando = makeAddr("rando");
        bridge = makeAddr("bridge");
        remoteToken = makeAddr("remoteToken");
        mintManager = makeAddr("mintManager");
        governanceToken = new GovernanceToken(bridge, remoteToken, mintManager);
    }

    /// @dev Tests that the constructor sets the correct initial state.
    function test_constructor_succeeds() external {
        assertEq(governanceToken.MINT_MANAGER(), mintManager);
        assertEq(governanceToken.name(), "Kroma");
        assertEq(governanceToken.symbol(), "KRO");
        assertEq(governanceToken.decimals(), 18);
        assertEq(governanceToken.totalSupply(), 0);
    }

    /// @dev Tests the bridge contract can successfully call `mint`.
    function test_mint_fromBridge_succeeds() external {
        // Mint 100 tokens.
        vm.prank(bridge);
        governanceToken.mint(owner, 100);

        // Balances have updated correctly.
        assertEq(governanceToken.balanceOf(owner), 100);
        assertEq(governanceToken.totalSupply(), 100);
    }

    /// @dev Tests that the MintManager can successfully call `mint`.
    function test_mint_fromMintManager_succeeds() external {
        // Mint 100 tokens.
        vm.prank(mintManager);
        governanceToken.mint(owner, 100);

        // Balances have updated correctly.
        assertEq(governanceToken.balanceOf(owner), 100);
        assertEq(governanceToken.totalSupply(), 100);
    }

    /// @dev Tests that `mint` reverts when called by a non-owner.
    function test_mint_fromNotOwner_reverts() external {
        // Mint 100 tokens as rando.
        vm.prank(rando);
        vm.expectRevert("GovernanceToken: only minter can mint");
        governanceToken.mint(owner, 100);

        // Balance does not update.
        assertEq(governanceToken.balanceOf(owner), 0);
        assertEq(governanceToken.totalSupply(), 0);
    }

    function test_mint_maxCapExceeded_reverts() external {
        uint256 cap = governanceToken.cap();

        // Mint up to the maximum total supply.
        vm.prank(mintManager);
        governanceToken.mint(owner, cap);

        assertEq(governanceToken.balanceOf(owner), cap);
        assertEq(governanceToken.totalSupply(), cap);

        // Then mint one more to exceed the total supply.
        vm.prank(mintManager);
        vm.expectRevert("GovernanceToken: cap exceeded");
        governanceToken.mint(owner, 1);

        // Balance does not update.
        assertEq(governanceToken.balanceOf(owner), cap);
        assertEq(governanceToken.totalSupply(), cap);
    }

    /// @dev Tests that the owner can successfully call `burn`.
    function test_burn_succeeds() external {
        // Mint 100 tokens to rando.
        vm.prank(mintManager);
        governanceToken.mint(rando, 100);

        // Rando burns their tokens.
        vm.prank(rando);
        governanceToken.burn(50);

        // Balances have updated correctly.
        assertEq(governanceToken.balanceOf(rando), 50);
        assertEq(governanceToken.totalSupply(), 50);
    }

    /// @dev Tests that the bridge contract can successfully call `burn`.
    function test_burn_fromBridge_succeeds() external {
        // Mint 100 tokens to rando.
        vm.prank(mintManager);
        governanceToken.mint(rando, 100);

        // Bridge burns rando's tokens.
        vm.prank(bridge);
        governanceToken.burn(rando, 100);

        // Balances have updated correctly.
        assertEq(governanceToken.balanceOf(rando), 0);
        assertEq(governanceToken.totalSupply(), 0);
    }

    /// @dev Tests that the owner can successfully call `burnFrom`.
    function test_burnFrom_succeeds() external {
        // Mint 100 tokens to rando.
        vm.prank(mintManager);
        governanceToken.mint(rando, 100);

        // Rando approves owner to burn 50 tokens.
        vm.prank(rando);
        governanceToken.approve(owner, 50);

        // Owner burns 50 tokens from rando.
        vm.prank(owner);
        governanceToken.burnFrom(rando, 50);

        // Balances have updated correctly.
        assertEq(governanceToken.balanceOf(rando), 50);
        assertEq(governanceToken.totalSupply(), 50);
    }

    /// @dev Tests that `transfer` correctly transfers tokens.
    function test_transfer_succeeds() external {
        // Mint 100 tokens to rando.
        vm.prank(mintManager);
        governanceToken.mint(rando, 100);

        // Rando transfers 50 tokens to owner.
        vm.prank(rando);
        governanceToken.transfer(owner, 50);

        // Balances have updated correctly.
        assertEq(governanceToken.balanceOf(owner), 50);
        assertEq(governanceToken.balanceOf(rando), 50);
        assertEq(governanceToken.totalSupply(), 100);
    }

    /// @dev Tests that `approve` correctly sets allowances.
    function test_approve_succeeds() external {
        // Mint 100 tokens to rando.
        vm.prank(mintManager);
        governanceToken.mint(rando, 100);

        // Rando approves owner to spend 50 tokens.
        vm.prank(rando);
        governanceToken.approve(owner, 50);

        // Allowances have updated.
        assertEq(governanceToken.allowance(rando, owner), 50);
    }

    /// @dev Tests that `transferFrom` correctly transfers tokens.
    function test_transferFrom_succeeds() external {
        // Mint 100 tokens to rando.
        vm.prank(mintManager);
        governanceToken.mint(rando, 100);

        // Rando approves owner to spend 50 tokens.
        vm.prank(rando);
        governanceToken.approve(owner, 50);

        // Owner transfers 50 tokens from rando to owner.
        vm.prank(owner);
        governanceToken.transferFrom(rando, owner, 50);

        // Balances have updated correctly.
        assertEq(governanceToken.balanceOf(owner), 50);
        assertEq(governanceToken.balanceOf(rando), 50);
        assertEq(governanceToken.totalSupply(), 100);
    }

    /// @dev Tests that `increaseAllowance` correctly increases allowances.
    function test_increaseAllowance_succeeds() external {
        // Mint 100 tokens to rando.
        vm.prank(mintManager);
        governanceToken.mint(rando, 100);

        // Rando approves owner to spend 50 tokens.
        vm.prank(rando);
        governanceToken.approve(owner, 50);

        // Rando increases allowance by 50 tokens.
        vm.prank(rando);
        governanceToken.increaseAllowance(owner, 50);

        // Allowances have updated.
        assertEq(governanceToken.allowance(rando, owner), 100);
    }

    /// @dev Tests that `decreaseAllowance` correctly decreases allowances.
    function test_decreaseAllowance_succeeds() external {
        // Mint 100 tokens to rando.
        vm.prank(mintManager);
        governanceToken.mint(rando, 100);

        // Rando approves owner to spend 100 tokens.
        vm.prank(rando);
        governanceToken.approve(owner, 100);

        // Rando decreases allowance by 50 tokens.
        vm.prank(rando);
        governanceToken.decreaseAllowance(owner, 50);

        // Allowances have updated.
        assertEq(governanceToken.allowance(rando, owner), 50);
    }

    /// @dev Tests that `totalMinted` correctly returns the actual minted amount.
    function test_totalMinted_succeeds() external {
        // Mint 100 tokens to rando.
        vm.prank(mintManager);
        governanceToken.mint(rando, 100);

        assertEq(governanceToken.totalMinted(), 100);

        // Rando burns their tokens.
        vm.prank(rando);
        governanceToken.burn(50);

        // `totalMinted` does not changed.
        assertEq(governanceToken.totalMinted(), 100);
    }

    /// @dev Tests that `cap` correctly returns the maximum number of tokens that can be minted.
    function test_cap_succeeds() external {
        assertEq(governanceToken.cap(), 50_000_000 ether);
    }
}
