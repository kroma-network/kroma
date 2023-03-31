// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/**
 * @title CodeDeployer
 * @notice CodeDeployer is a library to deploy bytecode.
 */
library CodeDeployer {
    function deployCode(bytes memory _code) internal returns (address deployedAddress) {
        assembly {
            deployedAddress := create(0, add(_code, 0x20), mload(_code))
        }
    }
}
