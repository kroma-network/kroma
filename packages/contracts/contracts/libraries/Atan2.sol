// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

/**
 * @title Atan2
 * @notice A library for calculating the arctangent of a fraction y / x. Based on fixed-point
 *         math library, it provides 1E-12 precision with 40 fractional bits.
 *         Originally from https://github.com/NovakDistributed/macroverse.
 */
library Atan2 {
    /**
     * @notice The value of pi/2 in radians, represented as a fixed-point number.
     */
    uint256 internal constant REAL_HALF_PI = 1727108826179;

    /**
     * @notice Calculate atan(y / x).
     * @dev Uses the Chebyshev polynomial approach to approximate arctan(x) where x is [0, 1].
     * @dev 0.999974x-0.332568x^3+0.193235x^5-0.115729x^7+0.0519505x^9-0.0114658x^11
     *
     * @param real_y The numerator of the fraction y / x.
     * @param real_x The denominator of the fraction y / x.
     *
     * @return result The angle in radians of the fraction y / x.
     */
    function atan2(uint256 real_y, uint256 real_x) internal pure returns (uint256 result) {
        assembly {
            let frac

            switch lt(real_x, real_y)
            case 0 {
                frac := div(mul(real_y, shl(40, 1)), real_x)
            }
            case 1 {
                frac := div(mul(real_x, shl(40, 1)), real_y)
            }

            // Initialize variables to be used in the polynomial.
            let frac_squared := shr(40, mul(frac, frac))
            let frac_cubed := shr(40, mul(frac_squared, frac))
            let frac_five_squared := shr(40, mul(frac_squared, frac_cubed))
            let frac_seven_squared := shr(40, mul(frac_squared, frac_five_squared))
            let frac_nine_squared := shr(40, mul(frac_squared, frac_seven_squared))
            let frac_eleven_squared := shr(40, mul(frac_squared, frac_nine_squared))

            // Calculate the polynomial using unsigned integers.
            // Start with the x^1 term, and then subtract or add the other terms based on the coefficient signs.
            result := shr(40, mul(1099483040474, frac)) // x^1 term

            // x^5 term
            result := add(result, shr(40, mul(212464129393, frac_five_squared)))

            // x^9 term
            result := add(result, shr(40, mul(57120178819, frac_nine_squared)))

            // x^3 term, subtract because original coefficient is negative
            result := sub(result, shr(40, mul(365662383026, frac_cubed)))

            // x^7 term, subtract because original coefficient is negative
            result := sub(result, shr(40, mul(127245381171, frac_seven_squared)))

            // x^11 term, subtract because original coefficient is negative
            result := sub(result, shr(40, mul(12606780422, frac_eleven_squared)))

            if gt(real_y, real_x) {
                result := sub(REAL_HALF_PI, result)
            }
        }
    }
}
