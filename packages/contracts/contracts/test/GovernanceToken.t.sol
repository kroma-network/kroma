// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

// Testing utilities
import { CommonTest } from "./CommonTest.t.sol";

// Target contract dependencies
import { Proxy } from "../universal/Proxy.sol";

// Target contract
import { GovernanceToken } from "../governance/GovernanceToken.sol";

contract GovernanceToken_Test is CommonTest {
    address rando;
    GovernanceToken governanceToken;
    address bridge;
    address remoteToken;
    address mintManager;

    event Mint(address indexed account, uint256 amount);

    event Burn(address indexed account, uint256 amount);

    /// @dev Sets up the test suite.
    function setUp() public virtual override {
        super.setUp();

        rando = makeAddr("rando");
        bridge = makeAddr("bridge");
        remoteToken = makeAddr("remoteToken");
        mintManager = makeAddr("mintManager");

        governanceToken = GovernanceToken(address(new Proxy(multisig)));
        GovernanceToken govTokenImpl = new GovernanceToken(bridge, remoteToken);
        vm.prank(multisig);
        toProxy(address(governanceToken)).upgradeToAndCall(
            address(govTokenImpl),
            abi.encodeCall(governanceToken.initialize, mintManager)
        );
        assertEq(governanceToken.pendingOwner(), mintManager);

        vm.prank(mintManager);
        governanceToken.acceptOwnership();
        assertEq(governanceToken.owner(), mintManager);
    }

    /// @dev Tests that the constructor sets the correct initial state.
    function test_constructor_succeeds() external {
        assertEq(governanceToken.BRIDGE(), bridge);
        assertEq(governanceToken.REMOTE_TOKEN(), remoteToken);
        assertEq(governanceToken.name(), "Kroma");
        assertEq(governanceToken.symbol(), "KRO");
        assertEq(governanceToken.decimals(), 18);
        assertEq(governanceToken.totalSupply(), 0);
    }

    /// @dev Tests that the owner can successfully call `mint`.
    function test_mint_fromOwner_succeeds() external {
        // Mint 100 tokens.
        vm.expectEmit(true, false, false, true, address(governanceToken));
        emit Mint(rando, 100);
        vm.prank(mintManager);
        governanceToken.mint(rando, 100);

        // Balances have updated correctly.
        assertEq(governanceToken.balanceOf(rando), 100);
        assertEq(governanceToken.totalSupply(), 100);
    }

    /// @dev Tests the bridge contract can successfully call `mint`.
    function test_mint_fromBridge_succeeds() external {
        // Mint 100 tokens.
        vm.expectEmit(true, false, false, true, address(governanceToken));
        emit Mint(rando, 100);
        vm.prank(bridge);
        governanceToken.mint(rando, 100);

        // Balances have updated correctly.
        assertEq(governanceToken.balanceOf(rando), 100);
        assertEq(governanceToken.totalSupply(), 100);
    }

    /// @dev Tests that `mint` reverts when called by a non-minter.
    function test_mint_fromNotMinter_reverts() external {
        // Mint 100 tokens as rando.
        vm.prank(rando);
        vm.expectRevert("GovernanceToken: only bridge or owner can mint");
        governanceToken.mint(rando, 100);

        // Balance does not update.
        assertEq(governanceToken.balanceOf(rando), 0);
        assertEq(governanceToken.totalSupply(), 0);
    }

    /// @dev Tests that the bridge contract can successfully call `burn`.
    function test_burn_fromBridge_succeeds() external {
        // Mint 100 tokens to rando.
        vm.prank(mintManager);
        governanceToken.mint(rando, 100);

        // Bridge burns rando's tokens.
        vm.expectEmit(true, false, false, true, address(governanceToken));
        emit Burn(rando, 100);
        vm.prank(bridge);
        governanceToken.burn(rando, 100);

        // Balances have updated correctly.
        assertEq(governanceToken.balanceOf(rando), 0);
        assertEq(governanceToken.totalSupply(), 0);
    }

    /// @dev Tests that other than bridge contract cannot call `burn`.
    function test_burn_fromNotBridge_reverts() external {
        // Mint 100 tokens to rando.
        vm.prank(mintManager);
        governanceToken.mint(rando, 100);

        vm.expectRevert("KromaMintableERC20: only bridge can mint and burn");
        vm.prank(rando);
        governanceToken.burn(rando, 100);
    }

    /// @dev Tests that `transfer` correctly transfers tokens.
    function test_transfer_succeeds() external {
        // Mint 100 tokens to rando.
        vm.prank(mintManager);
        governanceToken.mint(rando, 100);

        // Rando transfers 50 tokens to alice.
        vm.prank(rando);
        governanceToken.transfer(alice, 50);

        // Balances have updated correctly.
        assertEq(governanceToken.balanceOf(alice), 50);
        assertEq(governanceToken.balanceOf(rando), 50);
        assertEq(governanceToken.totalSupply(), 100);
    }

    /// @dev Tests that `approve` correctly sets allowances.
    function test_approve_succeeds() external {
        // Mint 100 tokens to rando.
        vm.prank(mintManager);
        governanceToken.mint(rando, 100);

        // Rando approves alice to spend 50 tokens.
        vm.prank(rando);
        governanceToken.approve(alice, 50);

        // Allowances have updated.
        assertEq(governanceToken.allowance(rando, alice), 50);
    }

    /// @dev Tests that `transferFrom` correctly transfers tokens.
    function test_transferFrom_succeeds() external {
        // Mint 100 tokens to rando.
        vm.prank(mintManager);
        governanceToken.mint(rando, 100);

        // Rando approves alice to spend 50 tokens.
        vm.prank(rando);
        governanceToken.approve(alice, 50);

        // Alice transfers 50 tokens from rando to alice.
        vm.prank(alice);
        governanceToken.transferFrom(rando, alice, 50);

        // Balances have updated correctly.
        assertEq(governanceToken.balanceOf(alice), 50);
        assertEq(governanceToken.balanceOf(rando), 50);
        assertEq(governanceToken.totalSupply(), 100);
    }

    /// @dev Tests that `increaseAllowance` correctly increases allowances.
    function test_increaseAllowance_succeeds() external {
        // Mint 100 tokens to rando.
        vm.prank(mintManager);
        governanceToken.mint(rando, 100);

        // Rando approves alice to spend 50 tokens.
        vm.prank(rando);
        governanceToken.approve(alice, 50);

        // Rando increases allowance by 50 tokens.
        vm.prank(rando);
        governanceToken.increaseAllowance(alice, 50);

        // Allowances have updated.
        assertEq(governanceToken.allowance(rando, alice), 100);
    }

    /// @dev Tests that `decreaseAllowance` correctly decreases allowances.
    function test_decreaseAllowance_succeeds() external {
        // Mint 100 tokens to rando.
        vm.prank(mintManager);
        governanceToken.mint(rando, 100);

        // Rando approves alice to spend 100 tokens.
        vm.prank(rando);
        governanceToken.approve(alice, 100);

        // Rando decreases allowance by 50 tokens.
        vm.prank(rando);
        governanceToken.decreaseAllowance(alice, 50);

        // Allowances have updated.
        assertEq(governanceToken.allowance(rando, alice), 50);
    }
}
