// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { Initializable } from "@openzeppelin/contracts/proxy/utils/Initializable.sol";

import { Predeploys } from "../libraries/Predeploys.sol";
import { ISemver } from "../universal/ISemver.sol";
import { GovernanceToken } from "./GovernanceToken.sol";

/**
 * @custom:proxied
 * @custom:predeploy 0x4200000000000000000000000000000000000070
 * @title MintManager
 * @notice MintManager issues a set amount of GovernanceToken for every block generation
 *         and distributes the token to a specified vaults.
 *         The amount of tokens minted per block decreases by a predetermined rate each epoch.
 *         Although the decimal of token is 18, but the mint amount is floored to 8 decimals.
 */
contract MintManager is Initializable, ISemver {
    /**
     * @notice The denominator of the decaying factor.
     */
    uint256 public constant DECAYING_DENOMINATOR = 10 ** 5;

    /**
     * @notice The value for truncating decimal point of the minting amount.
     */
    uint256 public constant FLOOR_UNIT = 10 ** 10;

    /**
     * @notice The denominator of each recipient's share.
     */
    uint256 public constant SHARE_DENOMINATOR = 10 ** 5;

    /**
     * @notice The GovernanceToken that the MintManager can mint tokens.
     */
    GovernanceToken public immutable GOVERNANCE_TOKEN;

    /**
     * @notice The number of the L2 block where the mint function is activated.
     */
    uint256 public immutable MINT_ACTIVATED_BLOCK;

    /**
     * @notice The amount of minting in the first epoch.
     */
    uint256 public immutable INIT_MINT_PER_BLOCK;

    /**
     * @notice The number of blocks per epoch. Can be updated via upgrade.
     */
    uint256 public immutable SLIDING_WINDOW_BLOCKS;

    /**
     * @notice The decaying factor that reduces the amount of minting. Can be updated via upgrade.
     */
    uint256 public immutable DECAYING_FACTOR;

    /**
     * @notice A list of recipient addresses that will receive tokens to be distributed.
     */
    address[] internal recipients;

    /**
     * @notice A mapping of the recipient's share.
     */
    mapping(address => uint256) public shareOf;

    /**
     * @notice The number of the last block that minted tokens.
     */
    uint256 internal lastMintedBlock;

    /**
     * @notice Reverts if the caller is not a L1Block contract.
     */
    modifier onlyL1Block() {
        require(msg.sender == Predeploys.L1_BLOCK_ATTRIBUTES, "MintManager: only the L1Block can call this function");
        _;
    }

    /**
     * @notice Semantic version.
     * @custom:semver 1.0.0
     */
    string public constant version = "1.0.0";

    /**
     * @notice Constructs the MintManager contract.
     *
     * @param _mintActivatedBlock  The number of L2 block number which the mint function is activated.
     * @param _initMintPerBlock    The amount of the initial minting per block.
     * @param _slidingWindowBlocks The number of blocks per epoch.
     * @param _decayingFactor      The decaying factor that reduces the amount of minting.
     */
    constructor(
        uint256 _mintActivatedBlock,
        uint256 _initMintPerBlock,
        uint256 _slidingWindowBlocks,
        uint256 _decayingFactor
    ) {
        MINT_ACTIVATED_BLOCK = _mintActivatedBlock;
        INIT_MINT_PER_BLOCK = _initMintPerBlock;
        SLIDING_WINDOW_BLOCKS = _slidingWindowBlocks;
        DECAYING_FACTOR = _decayingFactor;

        GOVERNANCE_TOKEN = GovernanceToken(Predeploys.GOVERNANCE_TOKEN);
    }

    /**
     * @notice Initializer.
     *
     * @param _recipients List of the recipient.
     * @param _shares     List of token distribution ratios for each recipient.
     */
    function initialize(
        address[] calldata _recipients,
        uint256[] calldata _shares
    ) external initializer {
        require(_recipients.length == _shares.length, "MintManager: invalid length of array");

        uint256 totalShares = 0;
        for (uint256 i = 0; i < _recipients.length; i++) {
            address recipient = _recipients[i];
            require(recipient != address(0), "MintManager: recipient address cannot be 0");
            uint256 share = _shares[i];
            require(share != 0, "MintManager: share cannot be 0");
            totalShares += share;

            recipients.push(recipient);
            shareOf[recipient] = share;
        }
        require(totalShares == SHARE_DENOMINATOR, "MintManager: invalid shares");
    }

    /**
     * @notice Mints and distributes tokens.
     *         It must be called every block and cannot be called more than once in the same block.
     *         If it is the first time tokens are being minted, mint the amount of tokens that
     *         should be minted up to the current block number.
     */
    function mint() external onlyL1Block {
        if (MINT_ACTIVATED_BLOCK > block.number) {
            return;
        }

        require(
            lastMintedBlock != block.number,
            "MintManager: tokens have already been minted in this block"
        );

        uint256 mintAmount;
        if (lastMintedBlock == 0) {
            mintAmount = _initialMintAmount(block.number);
        } else {
            mintAmount = mintAmountPerBlock(block.number);
        }

        if (mintAmount > 0) {
            for (uint256 i = 0; i < recipients.length; i++) {
                address recipient = recipients[i];
                uint256 share = shareOf[recipient];
                uint256 amount = (mintAmount * share) / SHARE_DENOMINATOR;

                GOVERNANCE_TOKEN.mint(recipient, amount);
            }

            lastMintedBlock = block.number;
        }
    }

    /**
     * @notice Returns the amount of tokens that should be minted at the given block number.
     *
     * @param _blockNumber The block number at the time of minting.
     *
     * @return The mint amount at the given block number.
     */
    function mintAmountPerBlock(uint256 _blockNumber) public view returns (uint256) {
        uint256 amount = INIT_MINT_PER_BLOCK;

        (uint256 epoch, ) = _getEpochAndOffset(_blockNumber);
        for (uint256 i = 1; i < epoch; i++) {
            amount = (amount * DECAYING_FACTOR) / DECAYING_DENOMINATOR;
            amount = (amount / FLOOR_UNIT) * FLOOR_UNIT;
        }

        return amount;
    }

    /**
     * @notice Returns the epoch number and the offset within that epoch for a given block number.
     *
     * @param _blockNumber The block number.
     *
     * @return The epoch number and the offset.
     */
    function _getEpochAndOffset(uint256 _blockNumber) private view returns (uint256, uint256) {
        uint256 epoch = (_blockNumber - 1) / SLIDING_WINDOW_BLOCKS + 1;
        uint256 offset = (_blockNumber - 1) % SLIDING_WINDOW_BLOCKS + 1;
        return (epoch, offset);
    }

    /**
     * @notice Returns the initial minting amount.
     *         It computes and returns the amount that should be minted up to the given block number.
     *
     * @param _blockNumber The block number at the time of minting.
     *
     * @return The amount of initial minting.
     */
    function _initialMintAmount(uint256 _blockNumber) private view returns (uint256) {
        uint256 amount = 0;
        uint256 mintPerBlock = INIT_MINT_PER_BLOCK;

        (uint256 epoch, uint256 offset) = _getEpochAndOffset(_blockNumber);
        for (uint256 i = 1; i < epoch; i++) {
            amount = amount + mintPerBlock * SLIDING_WINDOW_BLOCKS;
            mintPerBlock = (mintPerBlock * DECAYING_FACTOR) / DECAYING_DENOMINATOR;
            mintPerBlock = (mintPerBlock / FLOOR_UNIT) * FLOOR_UNIT;
        }

        if (offset > 0) {
            amount = amount + mintPerBlock * offset;
        }

        return amount;
    }
}
