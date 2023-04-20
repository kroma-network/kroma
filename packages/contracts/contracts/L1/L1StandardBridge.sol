// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { Predeploys } from "../libraries/Predeploys.sol";
import { Semver } from "../universal/Semver.sol";
import { StandardBridge } from "../universal/StandardBridge.sol";

/**
 * @custom:proxied
 * @title L1StandardBridge
 * @notice The L1StandardBridge is responsible for transfering ETH and ERC20 tokens between L1 and
 *         L2. In the case that an ERC20 token is native to L1, it will be escrowed within this
 *         contract. If the ERC20 token is native to L2, it will be burnt. ETH is instead stored
 *         inside the KromaPortal contract.
 *         NOTE: this contract is not intended to support all variations of ERC20 tokens. Examples
 *         of some token types that may not be properly supported by this contract include, but are
 *         not limited to: tokens with transfer fees, rebasing tokens, and tokens with blocklists.
 */
contract L1StandardBridge is StandardBridge, Semver {
    /**
     * @custom:semver 0.1.0
     *
     * @param _messenger Address of the L1CrossDomainMessenger.
     */
    constructor(address payable _messenger)
        Semver(0, 1, 0)
        StandardBridge(_messenger, payable(Predeploys.L2_STANDARD_BRIDGE))
    {}

    /**
     * @notice Allows EOAs to bridge ETH by sending directly to the bridge.
     */
    receive() external payable override onlyEOA {
        _initiateBridgeETH(msg.sender, msg.sender, msg.value, RECEIVE_DEFAULT_GAS_LIMIT, bytes(""));
    }
}
