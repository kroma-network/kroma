// SPDX-License-Identifier: MIT
pragma solidity ^0.8.15;

import { ISemver } from "../universal/ISemver.sol";
import "../universal/KromaSoulBoundERC721.sol";

/**
 * @custom:proxied
 * @title SecurityCouncilToken
 * @notice The SecurityCouncilToken is a basic token based on KromaSoulBoundERC721.
 */
contract SecurityCouncilToken is KromaSoulBoundERC721, ISemver {
    /**
     * @notice Semantic version.
     * @custom:semver 1.0.1
     */
    string public constant version = "1.0.1";

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
