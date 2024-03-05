// SPDX-License-Identifier: MIT
pragma solidity ^0.8.15;

import "@openzeppelin/contracts-upgradeable/governance/TimelockControllerUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

import { ISemver } from "../universal/ISemver.sol";

/**
 * @custom:proxied
 * @title TimeLock
 * @notice The TimeLock is a timelock controller based on OpenZeppelin TimelockController.
 */
contract TimeLock is Initializable, TimelockControllerUpgradeable, ISemver {
    /**
     * @notice Semantic version.
     * @custom:semver 1.0.0
     */
    string public constant version = "1.0.0";

    /**
     * @notice Initializer.
     *
     * @param _minDelay  Initial minimum delay for operations.
     * @param _proposers Accounts to be granted proposer and canceller roles.
     * @param _executors Accounts to be granted executor role.
     * @param _admin     Optional account to be granted admin role; disable with zero address.
     */
    function initialize(
        uint256 _minDelay,
        address[] memory _proposers,
        address[] memory _executors,
        address _admin
    ) public initializer {
        __TimelockController_init(_minDelay, _proposers, _executors, _admin);
    }
}
