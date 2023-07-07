// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { FeeVault } from "../universal/FeeVault.sol";
import { Semver } from "../universal/Semver.sol";

/**
 * @custom:proxied
 * @custom:predeploy 0x4200000000000000000000000000000000000006
 * @title ProtocolVault
 * @notice The ProtocolVault accumulates transaction fees to fund network operation.
 */
contract ProtocolVault is FeeVault, Semver {
    /**
     * @custom:semver 0.1.0
     *
     * @param _recipient Address that will receive the accumulated fees.
     */
    constructor(address _recipient) FeeVault(_recipient, 10 ether) Semver(0, 1, 0) {}
}
