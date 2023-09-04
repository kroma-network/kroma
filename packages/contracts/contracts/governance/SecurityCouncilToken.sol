// SPDX-License-Identifier: MIT
pragma solidity ^0.8.15;

import { Semver } from "../universal/Semver.sol";
import "../universal/KromaSoulBoundERC721.sol";

/**
 * @custom:proxied
 * @title SecurityCouncilToken
 * @notice The SecurityCouncilToken is a basic token based on KromaSoulBoundERC721.
 */
contract SecurityCouncilToken is KromaSoulBoundERC721, Semver {
    /**
     * @custom:semver 1.0.0
     */
    constructor() Semver(1, 0, 0) {}

    /**
     * @notice Initializer.
     *
     * @param _owner Owner of this token contract.
     */
    function initialize(address _owner) public initializer {
        __KromaSoulBoundERC721_init("KromaSecurityCouncil", "KSC", _owner);
    }

    function _baseURI() internal pure override returns (string memory) {
        return "https://nft.kroma.network/sc/";
    }
}
