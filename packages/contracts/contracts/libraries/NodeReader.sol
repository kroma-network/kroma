// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/**
 * @title NodeReader
 * @notice NodeReader is a library for reading ZKTrie Node.
 */
library NodeReader {
    /**
     * @notice Node types.
     *         See https://github.com/wemixkanvas/zktrie/blob/main/types/README.md.
     *
     * @custom:value MIDDLE Represents a middle node.
     * @custom:value LEAF   Represents a leaf node.
     * @custom:value EMPTY  Represents a empty node.
     * @custom:value ROOT   Represents a middle node.
     */
    enum NodeType {
        MIDDLE,
        LEAF,
        EMPTY,
        ROOT
    }

    /**
     * @notice Struct representing a Node.
     *         See https://github.com/wemixkanvas/zktrie/blob/main/types/README.md.
     */
    struct Node {
        NodeType nodeType;
        bytes32 childL;
        bytes32 childR;
        bytes32 nodeKey;
        bytes32[] valuePreimage;
        uint32 compressedFlags;
        bytes32 valueHash;
        bytes32 keyPreimage;
    }

    /**
     * @notice Struct representing an Item.
     */
    struct Item {
        bytes ptr;
        uint256 len;
    }

    /**
     * @notice Converts bytes to Item.
     *
     * @param _bytes bytes to convert.
     *
     * @return Item referencing _bytes.
     */
    function toItem(bytes memory _bytes) internal pure returns (Item memory) {
        bytes memory ptr;
        assembly {
            ptr := add(_bytes, 32)
        }
        return Item({ ptr: ptr, len: _bytes.length });
    }

    /**
     * @notice Reads an Item into an uint8.
     *         Internal ptr and length is updated automatically.
     *
     * @param _item Item to read.
     *
     * @return An uint8 value.
     */
    function readUint8(Item memory _item) internal pure returns (uint8) {
        require(_item.len >= 1, "NodeReader: too short for uint8");
        bytes memory newPtr;
        bytes memory ptr = _item.ptr;
        uint8 ret;
        assembly {
            ret := shr(248, mload(ptr))
            newPtr := add(ptr, 1)
        }
        _item.ptr = newPtr;
        _item.len -= 1;
        return ret;
    }

    /**
     * @notice Reads an Item into compressed flags and length of values.
     *         Internal ptr and length is updated automatically.
     *
     * @param _item Item to read.
     *
     * @return Compressed flags.
     * @return Length of values.
     */
    function readCompressedFlags(Item memory _item) internal pure returns (uint32, uint8) {
        require(_item.len >= 4, "NodeReader: too short for uint32");
        bytes memory newPtr;
        bytes memory ptr = _item.ptr;
        uint32 temp;
        uint8 flag;
        uint8 len;
        assembly {
            temp := mload(ptr)
            len := shr(248, temp)
            flag := shr(240, temp)
            newPtr := add(ptr, 4)
        }
        _item.ptr = newPtr;
        _item.len -= 4;
        return (flag, len);
    }

    /**
     * @notice Reads an Item into a bytes32.
     *         Internal ptr and length is updated automatically.
     *
     * @param _item Item to read.
     *
     * @return A bytes32 value.
     */
    function readBytes32(Item memory _item) internal pure returns (bytes32) {
        require(_item.len >= 32, "NodeReader: too short for bytes32");
        bytes memory newPtr;
        bytes memory ptr = _item.ptr;
        bytes32 ret;
        assembly {
            ret := mload(ptr)
            newPtr := add(ptr, 32)
        }
        _item.ptr = newPtr;
        _item.len -= 32;
        return ret;
    }

    /**
     * @notice Reads an Item by n bytes into a bytes32.
     *         Internal ptr and length is updated automatically.
     *
     * @param _item Item to read.
     *
     * @return A bytes32 value.
     */
    function readBytesN(Item memory _item, uint256 _length) internal pure returns (bytes32) {
        require(_item.len >= _length, "NodeReader: too short for n bytes");
        bytes memory newPtr;
        bytes memory ptr = _item.ptr;
        bytes32 ret;
        uint256 to = 256 - _length * 8;
        assembly {
            newPtr := add(ptr, _length)
            ret := shr(to, mload(ptr))
        }
        _item.ptr = newPtr;
        _item.len -= _length;
        return ret;
    }

    /**
     * @notice Reads bytes into a Node.
     *
     * @param _proof Bytes to read.
     *
     * @return A decoded Node.
     */
    function readNode(bytes memory _proof) internal pure returns (Node memory) {
        Node memory node;
        Item memory item = toItem(_proof);
        uint256 nodeType = readUint8(item);
        if (nodeType == uint256(NodeType.MIDDLE)) {
            // TODO(chokobole): Do the length check as much as possible at once and read the bytes.
            node.childL = readBytes32(item);
            node.childR = readBytes32(item);
        } else if (nodeType == uint256(NodeType.LEAF)) {
            // TODO(chokobole): Do the length check as much as possible at once and read the bytes.
            node.nodeKey = readBytes32(item);
            (uint32 compressedFlags, uint256 valuePreimageLen) = readCompressedFlags(item);
            node.compressedFlags = compressedFlags;
            node.valuePreimage = new bytes32[](valuePreimageLen);
            for (uint256 i = 0; i < valuePreimageLen; ) {
                node.valuePreimage[i] = readBytes32(item);
                unchecked {
                    ++i;
                }
            }
            uint256 keyPreimageLen = readUint8(item);
            if (keyPreimageLen > 0) {
                node.keyPreimage = readBytesN(item, keyPreimageLen);
            }
        } else if (nodeType == uint256(NodeType.EMPTY)) {
            // Do nothing.
        } else if (nodeType == uint256(NodeType.ROOT)) {
            revert("NodeReader: unexpected root node type");
        } else {
            revert("NodeReader: invalid node type");
        }
        node.nodeType = NodeType(nodeType);
        return node;
    }
}
