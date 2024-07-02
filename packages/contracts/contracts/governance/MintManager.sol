// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { Ownable2Step } from "@openzeppelin/contracts/access/Ownable2Step.sol";

import { GovernanceToken } from "./GovernanceToken.sol";

/**
 * @title MintManager
 * @notice MintManager issues mint cap amount of GovernanceToken at once (TGE) and distributes the
 *         tokens to specified recipients.
 */
contract MintManager is Ownable2Step {
    /**
     * @notice The amount of tokens that can be minted.
     */
    uint256 public constant MINT_CAP = 1_000_000_000;

    /**
     * @notice The denominator of each recipient's share.
     */
    uint256 public constant SHARE_DENOMINATOR = 10 ** 5;

    /**
     * @notice The GovernanceToken that the MintManager can mint.
     */
    GovernanceToken public immutable GOVERNANCE_TOKEN;

    /**
     * @notice True when already minted on this chain. MintManager can mint only once on each chain.
     */
    bool public minted;

    /**
     * @notice A list of recipient addresses that will receive tokens to be distributed.
     */
    address[] public recipients;

    /**
     * @notice A mapping of the recipient's address to share.
     */
    mapping(address => uint256) public shareOf;

    /**
     * @notice Constructs the MintManager contract.
     *
     * @param _governanceToken The GovernanceToken this contract can mint tokens of.
     * @param _owner           The owner of this contract.
     * @param _recipients      List of the recipients.
     * @param _shares          List of token distribution ratios for each recipient.
     */
    constructor(
        address _governanceToken,
        address _owner,
        address[] memory _recipients,
        uint256[] memory _shares
    ) {
        GOVERNANCE_TOKEN = GovernanceToken(_governanceToken);

        transferOwnership(_owner);

        require(_recipients.length == _shares.length, "MintManager: invalid length of array");

        uint256 totalShares = 0;
        for (uint256 i = 0; i < _recipients.length; i++) {
            address recipient = _recipients[i];
            require(recipient != address(0), "MintManager: recipient address cannot be 0");

            uint256 share = _shares[i];
            require(share != 0, "MintManager: share cannot be 0");

            if (shareOf[recipient] == 0) {
                recipients.push(recipient);
            }
            shareOf[recipient] += share;
            totalShares += share;
        }

        require(
            totalShares <= SHARE_DENOMINATOR,
            "MintManager: max total share is equal or less than SHARE_DENOMINATOR"
        );
    }

    /**
     * @notice Only the owner is allowed to mint mint cap amount of the GovernanceToken at once.
     */
    function mint() external onlyOwner {
        require(!minted, "MintManager: already minted on this chain");

        uint256 mintCap = MINT_CAP * 10 ** GOVERNANCE_TOKEN.decimals();

        uint256 totalAmount;
        for (uint256 i = 0; i < recipients.length; i++) {
            address recipient = recipients[i];
            uint256 share = shareOf[recipient];
            uint256 amount = (mintCap * share) / SHARE_DENOMINATOR;
            totalAmount += amount;
        }

        GOVERNANCE_TOKEN.mint(address(this), totalAmount);

        minted = true;
    }

    /**
     * @notice Only the owner is allowed to distribute the GovernanceToken to specified recipients.
     */
    function distribute() external onlyOwner {
        uint256 mintCap = MINT_CAP * 10 ** GOVERNANCE_TOKEN.decimals();

        for (uint256 i = 0; i < recipients.length; i++) {
            address recipient = recipients[i];
            uint256 share = shareOf[recipient];
            uint256 amount = (mintCap * share) / SHARE_DENOMINATOR;
            GOVERNANCE_TOKEN.transfer(recipient, amount);
        }
    }

    /**
     * @notice Only the owner is allowed to renounce the ownership of the GovernanceToken.
     */
    function renounceOwnershipOfToken() external onlyOwner {
        require(minted, "MintManager: not minted before renounce ownership");

        GOVERNANCE_TOKEN.renounceOwnership();
    }

    /**
     * @notice Only the owner is allowed to transfer the ownership of the GovernanceToken.
     *
     * @param newMintManager The new MintManager to own the GovernanceToken.
     */
    function transferOwnershipOfToken(address newMintManager) external onlyOwner {
        GOVERNANCE_TOKEN.transferOwnership(newMintManager);
    }

    /**
     * @notice Only the owner is allowed to accept the ownership of the GovernanceToken.
     */
    function acceptOwnershipOfToken() external onlyOwner {
        GOVERNANCE_TOKEN.acceptOwnership();
    }
}
