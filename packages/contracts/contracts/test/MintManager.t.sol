// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

// Testing utilities
import { CommonTest } from "./CommonTest.t.sol";

// Target contract dependencies
import { GovernanceToken } from "../governance/GovernanceToken.sol";
import { Proxy } from "../universal/Proxy.sol";

// Target contract
import { MintManager } from "../governance/MintManager.sol";

contract MintManagerTest is CommonTest {
    address owner;
    address rando;
    GovernanceToken governanceToken;
    MintManager mintManager;

    address[] recipients = new address[](10);
    uint256[] shares = new uint256[](10);
    uint256 SHARE_DENOMINATOR;
    uint256 MINT_CAP;

    /// @dev Sets up the test suite.
    function setUp() public virtual override {
        super.setUp();

        owner = makeAddr("owner");
        rando = makeAddr("rando");

        governanceToken = GovernanceToken(address(new Proxy(multisig)));

        SHARE_DENOMINATOR = 10 ** 5;
        for (uint256 i = 0; i < recipients.length; i++) {
            string memory name = string(abi.encodePacked("recipient", i));
            recipients[i] = makeAddr(name);
            shares[i] = SHARE_DENOMINATOR / recipients.length;
        }
        mintManager = new MintManager(address(governanceToken), owner, recipients, shares);
        assertEq(mintManager.pendingOwner(), owner);

        vm.prank(owner);
        mintManager.acceptOwnership();
        assertEq(mintManager.owner(), owner);

        GovernanceToken govTokenImpl = new GovernanceToken(address(0), address(0));
        vm.prank(multisig);
        toProxy(address(governanceToken)).upgradeToAndCall(
            address(govTokenImpl),
            abi.encodeCall(governanceToken.initialize, address(mintManager))
        );
        assertEq(governanceToken.pendingOwner(), address(mintManager));

        vm.prank(owner);
        mintManager.acceptOwnershipOfToken();
        assertEq(governanceToken.owner(), address(mintManager));

        MINT_CAP = mintManager.MINT_CAP() * 10 ** governanceToken.decimals();
    }

    /// @dev Tests that the constructor properly configures the contract.
    function test_constructor_succeeds() external {
        assertEq(address(mintManager.GOVERNANCE_TOKEN()), address(governanceToken));

        assertEq(mintManager.owner(), owner);

        for (uint256 i = 0; i < recipients.length; i++) {
            assertEq(mintManager.recipients(i), recipients[i]);
            assertEq(mintManager.shareOf(recipients[i]), shares[i]);
        }
    }

    function test_constructor_sameRecipient_succeeds() external {
        recipients = new address[](3);
        shares = new uint256[](3);

        recipients[0] = address(1);
        recipients[1] = address(2);
        recipients[2] = address(2);

        shares[0] = 3;
        shares[1] = 1;
        shares[2] = 2;

        mintManager = new MintManager(address(governanceToken), owner, recipients, shares);

        vm.expectRevert(bytes(""));
        mintManager.recipients(2);
        assertEq(mintManager.recipients(0), recipients[0]);
        assertEq(mintManager.recipients(1), recipients[1]);
        assertEq(mintManager.shareOf(recipients[0]), shares[0]);
        assertEq(mintManager.shareOf(recipients[1]), shares[1] + shares[2]);
    }

    function test_constructor_invalidLengthArray_reverts() external {
        recipients = new address[](2);
        shares = new uint256[](1);

        vm.expectRevert("MintManager: invalid length of array");
        mintManager = new MintManager(address(governanceToken), owner, recipients, shares);
    }

    function test_constructor_zeroRecipient_reverts() external {
        recipients = new address[](1);
        shares = new uint256[](1);

        recipients[0] = address(0);
        shares[0] = SHARE_DENOMINATOR;

        vm.expectRevert("MintManager: recipient address cannot be 0");
        mintManager = new MintManager(address(governanceToken), owner, recipients, shares);
    }

    function test_constructor_zeroShares_reverts() external {
        recipients = new address[](1);
        shares = new uint256[](1);

        recipients[0] = makeAddr("recipient");
        shares[0] = 0;

        vm.expectRevert("MintManager: share cannot be 0");
        mintManager = new MintManager(address(governanceToken), owner, recipients, shares);
    }

    function test_constructor_tooManyShares_reverts() external {
        recipients = new address[](10);
        shares = new uint256[](10);

        for (uint256 i = 0; i < recipients.length; i++) {
            string memory name = string(abi.encodePacked("recipient", i));
            recipients[i] = makeAddr(name);
            shares[i] = SHARE_DENOMINATOR / (recipients.length - 1);
        }

        vm.expectRevert("MintManager: max total share is equal or less than SHARE_DENOMINATOR");
        mintManager = new MintManager(address(governanceToken), owner, recipients, shares);
    }

    function test_mint_succeeds() public {
        assertFalse(mintManager.minted());

        // Mint once.
        vm.prank(owner);
        mintManager.mint();

        assertTrue(mintManager.minted());

        uint256 totalAmount;
        for (uint256 i = 0; i < recipients.length; i++) {
            address recipient = recipients[i];
            uint256 share = mintManager.shareOf(recipient);
            uint256 amount = (MINT_CAP * share) / SHARE_DENOMINATOR;
            totalAmount += amount;
        }

        // Token balance increases.
        assertEq(governanceToken.balanceOf(address(mintManager)), totalAmount);
    }

    function test_mint_fromNotOwner_reverts() external {
        // Mint from rando fails.
        vm.prank(rando);
        vm.expectRevert("Ownable: caller is not the owner");
        mintManager.mint();
    }

    function test_mint_alreadyMinted_reverts() external {
        test_mint_succeeds();

        // Mint again.
        vm.prank(owner);
        vm.expectRevert("MintManager: already minted on this chain");
        mintManager.mint();
    }

    function test_distribute_succeeds() public {
        test_mint_succeeds();

        vm.prank(owner);
        mintManager.distribute();

        for (uint256 i = 0; i < recipients.length; i++) {
            address recipient = recipients[i];
            uint256 share = mintManager.shareOf(recipient);
            uint256 amount = (MINT_CAP * share) / SHARE_DENOMINATOR;
            assertEq(governanceToken.balanceOf(recipient), amount);
        }
    }

    function test_distribute_fromNotOwner_reverts() external {
        vm.prank(rando);
        vm.expectRevert("Ownable: caller is not the owner");
        mintManager.distribute();
    }

    function test_renounceOwnershipOfToken_succeeds() external {
        test_mint_succeeds();

        assertEq(governanceToken.owner(), address(mintManager));

        vm.prank(owner);
        mintManager.renounceOwnershipOfToken();

        assertEq(governanceToken.owner(), address(0));
    }

    function test_renounceOwnershipOfToken_fromNotOwner_reverts() external {
        vm.prank(rando);
        vm.expectRevert("Ownable: caller is not the owner");
        mintManager.renounceOwnershipOfToken();
    }

    function test_renounceOwnershipOfToken_beforeMinted_reverts() external {
        vm.prank(owner);
        vm.expectRevert("MintManager: not minted before renounce ownership");
        mintManager.renounceOwnershipOfToken();
    }

    function test_transferAndAcceptOwnershipOfToken_succeeds() external {
        assertEq(governanceToken.owner(), address(mintManager));

        address newOwner = makeAddr("newOwner");
        MintManager newMintManager = new MintManager(
            address(governanceToken),
            newOwner,
            recipients,
            shares
        );
        vm.prank(newOwner);
        newMintManager.acceptOwnership();

        vm.prank(owner);
        mintManager.transferOwnershipOfToken(address(newMintManager));
        assertEq(governanceToken.pendingOwner(), address(newMintManager));
        assertEq(governanceToken.owner(), address(mintManager));

        vm.prank(newOwner);
        newMintManager.acceptOwnershipOfToken();
        assertEq(governanceToken.pendingOwner(), ZERO_ADDRESS);
        assertEq(governanceToken.owner(), address(newMintManager));
    }

    function test_transferOwnershipOfToken_fromNotOwner_reverts() external {
        vm.prank(rando);
        vm.expectRevert("Ownable: caller is not the owner");
        mintManager.transferOwnershipOfToken(rando);
    }

    function test_acceptOwnershipOfToken_fromNotOwner_reverts() external {
        vm.prank(rando);
        vm.expectRevert("Ownable: caller is not the owner");
        mintManager.acceptOwnershipOfToken();
    }
}
