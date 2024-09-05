// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { Test } from "forge-std/Test.sol";
import { Strings } from "@openzeppelin/contracts/utils/Strings.sol";

import { BalancedWeightTree } from "../libraries/BalancedWeightTree.sol";

contract BalancedWeightTree_Test is Test {
    BalancedWeightTree.Tree internal tree;

    function setUp() public {
        for (uint256 i; i < 100; i++) {
            address addr = makeAddr(Strings.toString(i));
            BalancedWeightTree.insert(tree, addr, uint120(rand(i, 1000)));
        }
    }

    function rand(uint256 _mixer, uint256 _range) private view returns (uint256) {
        uint256 seed = uint256(
            keccak256(abi.encodePacked(block.difficulty, block.timestamp, _mixer))
        );
        return seed % _range;
    }

    function testFuzz_insert_succeeds(address _addr, uint120 _weight) external {
        vm.assume(_addr != address(0));
        uint32 index = tree.nodeMap[_addr];
        vm.assume(index == 0);

        uint32 counter = tree.counter;
        BalancedWeightTree.insert(tree, _addr, _weight);

        index = tree.nodeMap[_addr];
        assertEq(tree.counter, counter + 1);
        assertEq(tree.nodes[index].addr, _addr);
        assertEq(tree.nodes[index].weight, _weight);
    }

    function testFuzz_update_succeeds(uint32 _index, uint120 _weight) external {
        uint32 index = (_index % tree.counter) + 1;
        address addr = tree.nodes[index].addr;

        bool succeed = BalancedWeightTree.update(tree, addr, _weight);

        uint32 newIndex = tree.nodeMap[addr];
        assertTrue(succeed);
        assertEq(tree.nodes[newIndex].weight, _weight);
    }

    function testFuzz_remove_succeeds(uint32 _index) external {
        uint32 removed = tree.removed;
        uint32 index = (_index % tree.counter) + 1;
        address addr = tree.nodes[index].addr;

        bool succeed = BalancedWeightTree.remove(tree, addr);

        assertTrue(succeed);
        assertEq(tree.removed, removed + 1);
        assertEq(tree.nodeMap[addr], 0);
        assertTrue(tree.nodes[index].addr != addr);
    }

    function testFuzz_select_succeeds(uint120 _weight) external {
        address selected = BalancedWeightTree.select(
            tree,
            _weight % tree.nodes[tree.root].weightSum
        );

        assertTrue(selected != address(0));
    }
}
