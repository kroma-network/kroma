pragma solidity 0.8.15;

import { Test } from "forge-std/Test.sol";

import { Bytes } from "../libraries/Bytes.sol";

contract Bytes_equal_Test is Test {
    /**
     * @notice Manually checks equality of two dynamic `bytes` arrays in memory.
     *
     * @param _a The first `bytes` array to compare.
     * @param _b The second `bytes` array to compare.
     *
     * @return True if the two `bytes` arrays are equal in memory.
     */
    function manualEq(bytes memory _a, bytes memory _b) internal pure returns (bool) {
        bool _eq;
        assembly {
            _eq := and(
                // Check if the contents of the two bytes arrays are equal in memory.
                eq(keccak256(add(0x20, _a), mload(_a)), keccak256(add(0x20, _b), mload(_b))),
                // Check if the length of the two bytes arrays are equal in memory.
                // This is redundant given the above check, but included for completeness.
                eq(mload(_a), mload(_b))
            )
        }
        return _eq;
    }

    /**
     * @notice Tests that the `equal` function in the `Bytes` library returns `false` if given two
     *         non-equal byte arrays.
     */
    function testFuzz_equal_notEqual_works(bytes memory _a, bytes memory _b) public {
        vm.assume(!manualEq(_a, _b));
        assertFalse(Bytes.equal(_a, _b));
    }

    /**
     * @notice Test whether or not the `equal` function in the `Bytes` library is equivalent to
     *         manually checking equality of the two dynamic `bytes` arrays in memory.
     */
    function testDiff_equal_works(bytes memory _a, bytes memory _b) public {
        assertEq(Bytes.equal(_a, _b), manualEq(_a, _b));
    }
}
