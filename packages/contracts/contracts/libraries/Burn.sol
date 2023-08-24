// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { SafeCall } from "./SafeCall.sol";

/**
 * @title Burn
 * @notice Utilities for burning stuff.
 */
library Burn {
    /**
     * Burns a given amount of ETH.
     * Note that execution engine of Kroma does not support SELFDESTRUCT opcode, so it sends ETH to zero address.
     *
     * @param _amount Amount of ETH to burn.
     */
    function eth(uint256 _amount) internal {
        SafeCall.call(address(0), gasleft(), _amount, "");
    }

    /**
     * Burns a given amount of gas.
     *
     * @param _amount Amount of gas to burn.
     */
    function gas(uint256 _amount) internal view {
        uint256 i = 0;
        uint256 initialGas = gasleft();
        while (initialGas - gasleft() < _amount) {
            ++i;
        }
    }
}
