// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

/**
 * @title Uint128Math
 * @notice A library for handling overflow-safe math on uint128, especially for mulDiv operations.
 *         This library is motivated from the open-source Openzeppelin's Math library.
 */
library Uint128Math {
    /**
     * @dev Returns the largest of two numbers.
     */
    function max(uint128 a, uint128 b) internal pure returns (uint128) {
        return a > b ? a : b;
    }

    /**
     * @notice Calculates floor(x * y / denominator) with full precision. Throws if result overflows a uint128 or denominator == 0
     * @dev Original credit to Remco Bloemen under MIT license (https://xn--2-umb.com/21/muldiv)
     * with further edits by Uniswap Labs and Openzeppelin also under MIT license.
     */
    function mulDiv(
        uint128 x,
        uint128 y,
        uint128 denominator
    ) internal pure returns (uint128 result) {
        unchecked {
            uint256 prod;
            assembly {
                prod := mul(x, y)
            }

            // Make sure the result is less than 2^128.
            require(denominator > (prod >> 128), "Uint128Math: mulDiv overflow");

            // Direct division as fallback since we can't guarantee not exceeding 128 bits without further checks.
            result = uint128(prod / uint256(denominator));
            return result;
        }
    }
}
