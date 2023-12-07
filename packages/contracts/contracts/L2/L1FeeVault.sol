// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { FeeVault } from "../universal/FeeVault.sol";
import { Semver } from "../universal/Semver.sol";

/**
 * @custom:proxied
 * @custom:predeploy 0x4200000000000000000000000000000000000007
 * @title L1FeeVault
 * @notice The L1FeeVault accumulates the L1 portion of the transaction fees.
 */
contract L1FeeVault is FeeVault, Semver {
    /**
     * @custom:semver 1.0.2
     *
     * @param _recipient Address that will receive the accumulated fees.
     */
    constructor(address _recipient) FeeVault(_recipient, 0) Semver(1, 0, 2) {}
}
