// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { NodeReader } from "../libraries/NodeReader.sol";
import { CommonTest } from "./CommonTest.t.sol";

contract NodeReader_Test is CommonTest {
    function test_readUint8_bytestring00() external {
        NodeReader.Item memory item = NodeReader.toItem(hex"00");
        assertEq(NodeReader.readUint8(item), uint8(0x00));
    }

    function test_readUint8_bytestring01() external {
        NodeReader.Item memory item = NodeReader.toItem(hex"01");
        assertEq(NodeReader.readUint8(item), uint8(0x01));
    }

    function test_readUint8_bytestring7f() external {
        NodeReader.Item memory item = NodeReader.toItem(hex"7f");
        assertEq(NodeReader.readUint8(item), uint8(0x7f));
    }

    function test_readUint8_too_short_bytestring() external {
        vm.expectRevert("NodeReader: too short for uint8");
        NodeReader.Item memory item = NodeReader.toItem(hex"");
        NodeReader.readUint8(item);
    }

    function test_readCompressedFlags_length01_and_flag000003() external {
        NodeReader.Item memory item = NodeReader.toItem(hex"01030000");
        (uint32 compressedFlags, uint8 len) = NodeReader.readCompressedFlags(item);
        assertEq(compressedFlags, uint32(3));
        assertEq(len, uint8(1));
    }

    function test_readCompressedFlags_too_short_byte() external {
        vm.expectRevert("NodeReader: too short for uint32");
        NodeReader.Item memory item = NodeReader.toItem(hex"");
        NodeReader.readCompressedFlags(item);
    }

    function test_readBytes32_32bytesting() external {
        NodeReader.Item memory item = NodeReader.toItem(
            hex"000102030405060708090a0b0c0d0e0f101112131415161718192a2b2c2d2e2f"
        );
        assertEq(
            NodeReader.readBytes32(item),
            hex"000102030405060708090a0b0c0d0e0f101112131415161718192a2b2c2d2e2f"
        );
    }

    function test_readBytes32_too_short_byte() external {
        vm.expectRevert("NodeReader: too short for bytes32");
        NodeReader.Item memory item = NodeReader.toItem(hex"");
        NodeReader.readBytes32(item);
    }

    function test_readBytesN_4bytesting() external {
        NodeReader.Item memory item = NodeReader.toItem(hex"0001020304");
        assertEq(NodeReader.readBytesN(item, 4), bytes32(uint256(0x010203)));
    }

    function test_readBytesN_too_short_byte() external {
        vm.expectRevert("NodeReader: too short for n bytes");
        NodeReader.Item memory item = NodeReader.toItem(hex"0001020304");
        NodeReader.readBytesN(item, 6);
    }

    function test_readNode_middle_node() external {
        bytes memory middleHex = new bytes(65);
        bytes32 childL = hex"0000000000000000000000000000000000000000000000000000000000000000";
        bytes32 childR = hex"04470b58d80eeb26da85b2c2db5c254900656fb459c07729f556ff02534ab32a";
        assembly {
            mstore8(add(middleHex, 32), 0)
            mstore(add(middleHex, 33), childL)
            mstore(add(middleHex, 65), childR)
        }
        NodeReader.Node memory node = NodeReader.readNode(middleHex);
        assertEq(uint256(node.nodeType), uint256(NodeReader.NodeType.MIDDLE));
        assertEq(node.childL, childL);
        assertEq(node.childR, childR);
    }

    function test_readNode_leaf_node() external {
        bytes memory leafHex = new bytes(102);
        bytes32 nodeKey = hex"7f9d3bbc51d12566ecc6049ca6bf76e32828c22b197405f63a833b566fe7da0a";
        bytes32 value = hex"0000000000000000000000000000000000000000000000000000000000000001";
        assembly {
            mstore8(add(leafHex, 32), 1)
            mstore(add(leafHex, 33), nodeKey)
            mstore8(add(leafHex, 65), 1)
            mstore8(add(leafHex, 66), 1)
            mstore8(add(leafHex, 67), 0)
            mstore8(add(leafHex, 68), 0)
            mstore(add(leafHex, 69), value)
            mstore8(add(leafHex, 101), 0)
        }
        NodeReader.Node memory node = NodeReader.readNode(leafHex);
        assertEq(uint256(node.nodeType), uint256(NodeReader.NodeType.LEAF));
        assertEq(node.nodeKey, nodeKey);
        assertEq(node.compressedFlags, 1);
        assertEq(node.valuePreimage.length, 1);
        assertEq(node.valuePreimage[0], value);
    }

    function test_readNode_empty_node() external {
        bytes memory emptyHex = new bytes(1);
        assembly {
            mstore8(add(emptyHex, 32), 2)
        }
        NodeReader.Node memory node = NodeReader.readNode(emptyHex);
        assertEq(uint256(node.nodeType), uint256(NodeReader.NodeType.EMPTY));
    }

    function test_readNode_root_node() external {
        vm.expectRevert("NodeReader: unexpected root node type");
        bytes memory rootHex = new bytes(1);
        assembly {
            mstore8(add(rootHex, 32), 3)
        }
        NodeReader.readNode(rootHex);
    }

    function test_readNode_invalid_node() external {
        vm.expectRevert("NodeReader: invalid node type");
        bytes memory invalidHex = new bytes(1);
        assembly {
            mstore8(add(invalidHex, 32), 4)
        }
        NodeReader.readNode(invalidHex);
    }
}
