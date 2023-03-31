// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/**
 * @title Bytes32
 * @notice Bytes32 is a library for manipulating byte32.
 */
library Bytes32 {
    /**
     * @notice Splits bytes32 to high and low parts.
     *
     * @param _bytes Bytes32 to split.
     *
     * @return High part of bytes32.
     * @return Low part of bytes32.
     */
    function split(bytes32 _bytes) internal pure returns (bytes32, bytes32) {
        bytes16 high = bytes16(_bytes);
        bytes16 low = bytes16(uint128(uint256(_bytes)));
        return (fromBytes16(high), fromBytes16(low));
    }

    /**
     * @notice Converts bytes16 to bytes32.
     *
     * @param _bytes Bytes to constrcut to bytes32.
     *
     * @return Bytes32 constructed from bytes16.
     */
    function fromBytes16(bytes16 _bytes) internal pure returns (bytes32) {
        return bytes32(uint256(uint128(_bytes)));
    }
}
