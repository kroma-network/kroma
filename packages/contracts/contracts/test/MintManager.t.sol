// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

// Testing utilities
import { CommonTest } from "./CommonTest.t.sol";

// Target contract
import { GovernanceToken } from "../governance/GovernanceToken.sol";
import { MintManager } from "../governance/MintManager.sol";
import { Predeploys } from "../libraries/Predeploys.sol";

contract MintManagerTest is CommonTest {
    address owner;
    address rando;
    GovernanceToken governanceToken;
    MintManager mintManager;

    uint256 constant SHARE_DENOMINATOR = 10 ** 5;
    uint256 constant MINT_ACTIVATED_BLOCK = 0;
    uint256 constant INIT_MINT_PER_BLOCK = 1 ether;
    uint256 constant SLIDING_WINDOW_BLOCKS = 3_888_000;
    uint256 constant DECAYING_FACTOR = 92224;
    uint256 constant MINT_EPOCH_1 = 100000000 * 10 ** 10;
    uint256 constant MINT_EPOCH_2 = 92224000 * 10 ** 10;
    uint256 constant MINT_EPOCH_3 = 85052661 * 10 ** 10;
    uint256 constant MINT_EPOCH_4 = 78438966 * 10 ** 10;

    address[] recipients = new address[](10);
    uint256[] shares = new uint256[](10);

    function setUp() public virtual override {
        super.setUp();

        owner = makeAddr("owner");
        rando = makeAddr("rando");

        mintManager = MintManager(Predeploys.MINT_MANAGER);
        governanceToken = GovernanceToken(Predeploys.GOVERNANCE_TOKEN);

        MintManager mintManagerImpl = new MintManager(
            MINT_ACTIVATED_BLOCK,
            INIT_MINT_PER_BLOCK,
            SLIDING_WINDOW_BLOCKS,
            DECAYING_FACTOR
        );
        GovernanceToken govTokenImpl = new GovernanceToken(
            address(0),
            address(0),
            address(mintManager)
        );

        vm.etch(address(mintManager), address(mintManagerImpl).code);
        vm.etch(address(governanceToken), address(govTokenImpl).code);

        // Initialize
        for (uint256 i = 0; i < recipients.length; i++) {
            string memory name = string(abi.encodePacked("recipient", i));
            recipients[i] = makeAddr(name);
            shares[i] = SHARE_DENOMINATOR / recipients.length;
        }

        mintManager.initialize(recipients, shares);

        for (uint256 i = 0; i < recipients.length; i++) {
            assertEq(mintManager.shareOf(recipients[i]), shares[i]);
        }
    }

    function test_constructor_succeeds() external {
        assertEq(address(mintManager.GOVERNANCE_TOKEN()), address(governanceToken));
        assertEq(mintManager.SLIDING_WINDOW_BLOCKS(), SLIDING_WINDOW_BLOCKS);
        assertEq(mintManager.DECAYING_FACTOR(), DECAYING_FACTOR);
    }

    function test_initializer_invalidShares_reverts() external {
        mintManager = new MintManager(
            MINT_ACTIVATED_BLOCK,
            INIT_MINT_PER_BLOCK,
            SLIDING_WINDOW_BLOCKS,
            DECAYING_FACTOR
        );

        recipients = new address[](10);
        shares = new uint256[](10);

        for (uint256 i = 0; i < recipients.length; i++) {
            string memory name = string(abi.encodePacked("recipient", i));
            recipients[i] = makeAddr(name);
            shares[i] = 1;
        }

        vm.expectRevert("MintManager: invalid shares");
        mintManager.initialize(recipients, shares);
    }

    function test_initializer_zeroShares_reverts() external {
        mintManager = new MintManager(
            MINT_ACTIVATED_BLOCK,
            INIT_MINT_PER_BLOCK,
            SLIDING_WINDOW_BLOCKS,
            DECAYING_FACTOR
        );

        recipients = new address[](1);
        shares = new uint256[](1);

        recipients[0] = makeAddr("recipient");
        shares[0] = 0;

        vm.expectRevert("MintManager: share cannot be 0");
        mintManager.initialize(recipients, shares);
    }

    function test_initializer_zeroRecipient_reverts() external {
        mintManager = new MintManager(
            MINT_ACTIVATED_BLOCK,
            INIT_MINT_PER_BLOCK,
            SLIDING_WINDOW_BLOCKS,
            DECAYING_FACTOR
        );

        recipients = new address[](1);
        shares = new uint256[](1);

        recipients[0] = address(0);
        shares[0] = SHARE_DENOMINATOR;

        vm.expectRevert("MintManager: recipient address cannot be 0");
        mintManager.initialize(recipients, shares);
    }

    function test_mint_default_succeeds() external {
        uint256 minted;

        minted = mintManager.mintAmountPerBlock(block.number);
        vm.prank(Predeploys.L1_BLOCK_ATTRIBUTES);
        mintManager.mint();

        uint256[] memory prevBalances = new uint256[](recipients.length);
        for (uint256 i = 0; i < recipients.length; i++) {
            uint256 balance = governanceToken.balanceOf(recipients[i]);
            prevBalances[i] = balance;
        }

        uint256 supply = governanceToken.totalSupply();
        assertEq(supply, minted);

        vm.roll(block.number + 1);
        vm.prank(Predeploys.L1_BLOCK_ATTRIBUTES);
        mintManager.mint();

        minted = mintManager.mintAmountPerBlock(block.number);
        assertEq(governanceToken.totalSupply(), supply + minted);

        for (uint256 i = 0; i < recipients.length; i++) {
            uint256 balance = governanceToken.balanceOf(recipients[i]);
            uint256 expected = prevBalances[i] + (minted * shares[i]) / SHARE_DENOMINATOR;
            assertEq(balance, expected);
        }
    }

    function test_mint_initial_succeeds() external {
        uint256 blockNumber = 12345678;

        vm.roll(blockNumber);
        assertEq(governanceToken.totalSupply(), 0);

        // epoch     is      4 = ceil(12345678/3888000)
        // remainder is 681678 = 12345678%3888000
        uint256 expected = 0;
        // epoch 1 - 1 token per block
        expected += MINT_EPOCH_1 * SLIDING_WINDOW_BLOCKS;
        // epoch 2 - 0.92224000 token per block
        expected += MINT_EPOCH_2 * SLIDING_WINDOW_BLOCKS;
        // epoch 3 - 0.85052661 token per block
        expected += MINT_EPOCH_3 * SLIDING_WINDOW_BLOCKS;
        // epoch 4 - 0.78438966 token per block
        expected += MINT_EPOCH_4 * (blockNumber % SLIDING_WINDOW_BLOCKS);

        vm.prank(Predeploys.L1_BLOCK_ATTRIBUTES);
        mintManager.mint();
        assertEq(governanceToken.totalSupply(), expected);
    }

    function test_mint_notActivated_succeeds() external {
        mintManager = new MintManager(
            type(uint256).max,
            INIT_MINT_PER_BLOCK,
            SLIDING_WINDOW_BLOCKS,
            DECAYING_FACTOR
        );

        vm.prank(Predeploys.L1_BLOCK_ATTRIBUTES);
        mintManager.mint();

        assertEq(governanceToken.totalSupply(), 0);
    }

    function test_mint_notMintCaller_reverts() external {
        vm.expectRevert("MintManager: only the L1Block can call this function");
        vm.prank(makeAddr("someone"));
        mintManager.mint();
    }

    function test_mintAmountPerBlock_firstEpoch_succeeds() external {
        uint256 mintAmount = mintManager.mintAmountPerBlock(block.number);
        assertEq(mintAmount, MINT_EPOCH_1);
    }

    function test_mintAmountPerBlock_whenEpochIncreased_succeeds() external {
        uint256 prevMintAmount = mintManager.mintAmountPerBlock(block.number);
        assertEq(prevMintAmount, MINT_EPOCH_1);

        vm.roll(block.number + SLIDING_WINDOW_BLOCKS);

        uint256 mintAmount = mintManager.mintAmountPerBlock(block.number);
        assertEq(mintAmount, MINT_EPOCH_2);
        assertLt(mintAmount, prevMintAmount);
        prevMintAmount = mintAmount;

        vm.roll(block.number + SLIDING_WINDOW_BLOCKS);

        mintAmount = mintManager.mintAmountPerBlock(block.number);
        assertEq(mintAmount, MINT_EPOCH_3);
        assertLt(mintAmount, prevMintAmount);
        prevMintAmount = mintAmount;

        vm.roll(block.number + SLIDING_WINDOW_BLOCKS);

        mintAmount = mintManager.mintAmountPerBlock(block.number);
        assertEq(mintAmount, MINT_EPOCH_4);
        assertLt(mintAmount, prevMintAmount);
    }
}
