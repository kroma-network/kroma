// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/**
 * @title IKGHManager
 * @notice Interface for contracts that are compatible with the KromaMintableERC721 standard.
 *         Tokens that follow this standard can be easily transferred across the ERC721 bridge.
 */
interface IKGHManager {
    function totalKroInKgh(uint256 tokenId) external view returns (uint128);
}
