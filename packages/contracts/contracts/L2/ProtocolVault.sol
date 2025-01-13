// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { FeeVault } from "../universal/FeeVault.sol";
import { ISemver } from "../universal/ISemver.sol";

/**
 * @custom:proxied
 * @custom:predeploy 0x4200000000000000000000000000000000000006 (before Kroma MPT transition)
 * @custom:predeploy 0x4200000000000000000000000000000000000011 (for SequencerFeeVault after Kroma MPT transition)
 * @custom:predeploy 0x4200000000000000000000000000000000000019 (for BaseFeeVault after Kroma MPT transition)
 * @title ProtocolVault
 * @notice The ProtocolVault accumulates transaction fees to fund network operation.
 */
contract ProtocolVault is FeeVault, ISemver {
    /**
     * @notice Semantic version.
     * @custom:semver 1.0.1
     */
    string public constant version = "1.0.1";

    /**
     * @notice Constructs the ProtocolVault contract.
     *
     * @param _recipient Address that will receive the accumulated fees.
     */
    constructor(address _recipient) FeeVault(_recipient, 0) {}
}
