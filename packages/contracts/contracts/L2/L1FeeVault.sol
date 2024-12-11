// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { FeeVault } from "../universal/FeeVault.sol";
import { ISemver } from "../universal/ISemver.sol";

/**
 * @custom:proxied
 * @custom:predeploy 0x4200000000000000000000000000000000000007 (before Kroma MPT transition)
 * @custom:predeploy 0x420000000000000000000000000000000000001A (after Kroma MPT transition)
 * @title L1FeeVault
 * @notice The L1FeeVault accumulates the L1 portion of the transaction fees.
 */
contract L1FeeVault is FeeVault, ISemver {
    /**
     * @notice Semantic version.
     * @custom:semver 1.0.2
     */
    string public constant version = "1.0.2";

    /**
     * @notice Constructs the L1FeeVault contract.
     *
     * @param _recipient Address that will receive the accumulated fees.
     */
    constructor(address _recipient) FeeVault(_recipient, 0) {}
}
