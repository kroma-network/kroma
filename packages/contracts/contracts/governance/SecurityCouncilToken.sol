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
     * @custom:semver 0.1.0
     */
    constructor() Semver(0, 1, 0) {}

    /**
     * @notice Initializer.
     *
     * @param _owner Owner of this token contract.
     */
    function initialize(address _owner) public initializer {
        __KromaSoulBoundERC721_init("KromaSecurityCouncil", "KSC", _owner);
    }

    // TODO(ayaan): set base URI for security council SBT
    function _baseURI() internal pure override returns (string memory) {
        return "";
    }
}
