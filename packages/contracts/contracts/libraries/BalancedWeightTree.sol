// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

/**
 * @title BalancedWeightTree
 * @notice A self-balancing tree (balancing the weights) holds the keys and their weights and the
 *         weight sum of each child. A random integer smaller than the weight sum is taken and the
 *         tree is traversed to find the matching key.
 *         See https://github.com/yasharpm/Solidity-Weighted-Random-List.
 */
library BalancedWeightTree {
    /**
     * @notice Struct representing a node that constructs the tree.
     *
     * @custom:field addr        The address that owns the node.
     * @custom:field parent      The index of parent node.
     * @custom:field leftChild   The index of left child node.
     * @custom:field rightChild  The index of right child node.
     * @custom:field isLeftChild If the node is left child node of its parent node.
     * @custom:field weight      The weight of the node.
     * @custom:field weightSum   The weight sum of the node and its child nodes.
     */
    struct Node {
        address addr;
        uint32 parent;
        uint32 leftChild;
        uint32 rightChild;
        bool isLeftChild;
        uint120 weight;
        uint120 weightSum;
    }

    /**
     * @notice Struct representing a tree.
     *
     * @custom:field counter A counter used to assign a unique index to each node. The node index
     *                       starts with 1 and counter only increments.
     * @custom:field removed The cumulative number of removed nodes.
     * @custom:field root    The index of root node.
     * @custom:field nodes   A mapping of node index to node struct.
     * @custom:field nodeMap A mapping of owner address to node index.
     */
    struct Tree {
        uint32 counter;
        uint32 removed;
        uint32 root;
        mapping(uint32 => Node) nodes;
        mapping(address => uint32) nodeMap;
    }

    /**
     * @notice Inserts new node with the specified address and weight inside the tree.
     *
     * @param _tree   The tree to insert the new node.
     * @param _addr   The address that owns the new node.
     * @param _weight The weight of the new node.
     */
    function insert(Tree storage _tree, address _addr, uint120 _weight) internal {
        require(_addr != address(0), "BalancedWeightTree: zero address not allowed");
        require(_tree.nodeMap[_addr] == 0, "BalancedWeightTree: node already existing");

        Node memory newNode = Node({
            addr: _addr,
            weight: _weight,
            weightSum: _weight,
            parent: 0,
            leftChild: 0,
            rightChild: 0,
            isLeftChild: true
        });

        unchecked {
            _tree.counter++;
        }

        uint32 newNodeIndex = _tree.counter;
        _tree.nodes[newNodeIndex] = newNode;
        _tree.nodeMap[_addr] = newNodeIndex;

        if (_tree.root == 0) {
            _tree.root = newNodeIndex;
            return;
        }

        uint32 index = _tree.root;
        while (true) {
            Node storage node = _tree.nodes[index];

            unchecked {
                node.weightSum += _weight;
            }

            if (node.leftChild == 0) {
                _tree.nodes[newNodeIndex].parent = index;
                node.leftChild = newNodeIndex;

                _promote(_tree, newNodeIndex);

                return;
            } else if (node.rightChild == 0) {
                _tree.nodes[newNodeIndex].parent = index;
                _tree.nodes[newNodeIndex].isLeftChild = false;
                node.rightChild = newNodeIndex;

                _promote(_tree, newNodeIndex);

                return;
            } else if (
                _tree.nodes[node.leftChild].weightSum > _tree.nodes[node.rightChild].weightSum
            ) {
                index = node.rightChild;
            } else {
                index = node.leftChild;
            }
        }
    }

    /**
     * @notice Updates the weight of the node with the specified address. Returns true if the weight
     *         is updated, false if the node with specified address doesn't exist.
     *
     * @param _tree   The tree that includes the node to update.
     * @param _addr   The address that owns the node to update.
     * @param _weight The new weight to be assigned.
     *
     * @return If the weight is updated.
     */
    function update(Tree storage _tree, address _addr, uint120 _weight) internal returns (bool) {
        uint32 index = _tree.nodeMap[_addr];

        if (index == 0) {
            return false;
        }

        uint120 oldWeight = _tree.nodes[index].weight;
        _tree.nodes[index].weight = _weight;

        uint32 parentIndex = _tree.nodes[index].parent;
        if (_weight > oldWeight) {
            unchecked {
                uint120 weightDiff = _weight - oldWeight;
                _tree.nodes[index].weightSum += weightDiff;

                while (parentIndex != 0) {
                    _tree.nodes[parentIndex].weightSum += weightDiff;
                    parentIndex = _tree.nodes[parentIndex].parent;
                }
            }

            _promote(_tree, index);
        } else {
            unchecked {
                uint120 weightDiff = oldWeight - _weight;
                _tree.nodes[index].weightSum -= weightDiff;

                while (parentIndex != 0) {
                    _tree.nodes[parentIndex].weightSum -= weightDiff;
                    parentIndex = _tree.nodes[parentIndex].parent;
                }
            }

            _demote(_tree, index);
        }

        return true;
    }

    /**
     * @notice Removes the node with specified address from the tree. Returns true is the node is
     *         removed, false if it doesn't exist in the tree.
     *
     * @param _tree The tree that includes the node to remove.
     * @param _addr The address that owns the node to remove.
     *
     * @return If the node is removed.
     */
    function remove(Tree storage _tree, address _addr) internal returns (bool) {
        uint32 index = _tree.nodeMap[_addr];

        if (index == 0) {
            return false;
        }

        delete _tree.nodeMap[_addr];

        uint32 parentIndex = _tree.nodes[index].parent;
        uint120 weight = _tree.nodes[index].weight;
        while (parentIndex != 0) {
            unchecked {
                _tree.nodes[parentIndex].weightSum -= weight;
            }
            parentIndex = _tree.nodes[parentIndex].parent;
        }

        _pullUp(_tree, index);

        unchecked {
            ++_tree.removed;
        }

        return true;
    }

    /**
     * @notice Performs a weighted selection among the stored nodes. Returns the address of the
     *         selected node. If _weight is equal or greater than the weight sum of the tree, it
     *         returns zero address.
     *
     * @param _tree   The tree that includes the nodes to select.
     * @param _weight The random weight to be used for selection.
     *
     * @return The address of the selected node.
     */
    function select(Tree storage _tree, uint120 _weight) internal view returns (address) {
        uint32 index = _tree.root;
        while (true) {
            if (_tree.nodes[_tree.nodes[index].leftChild].weightSum > _weight) {
                index = _tree.nodes[index].leftChild;
                continue;
            }

            unchecked {
                _weight -= _tree.nodes[_tree.nodes[index].leftChild].weightSum;
            }

            if (_tree.nodes[index].weight > _weight) {
                return _tree.nodes[index].addr;
            }

            unchecked {
                _weight -= _tree.nodes[index].weight;
            }

            if (_tree.nodes[_tree.nodes[index].rightChild].weightSum > _weight) {
                index = _tree.nodes[index].rightChild;
            } else {
                return address(0);
            }
        }

        return address(0);
    }

    /**
     * @notice Promotes the node with higher weight to higher level of the tree. It is because to
     *         reduce the average number of traverses required since these nodes are more likely to
     *         be randomly selected.
     *
     * @param _tree  The tree that includes the node to promote.
     * @param _index The initial index of the target node to promote.
     */
    function _promote(Tree storage _tree, uint32 _index) private {
        Node storage node = _tree.nodes[_index];
        Node storage parentNode = _tree.nodes[node.parent];

        while (node.parent != 0 && node.weight > parentNode.weight) {
            address nodeAddr = node.addr;
            node.addr = parentNode.addr;
            parentNode.addr = nodeAddr;

            uint120 nodeWeight = node.weight;
            uint120 parentWeight = parentNode.weight;
            node.weight = parentWeight;
            parentNode.weight = nodeWeight;

            unchecked {
                node.weightSum -= nodeWeight - parentWeight;
            }

            _tree.nodeMap[node.addr] = _index;
            _tree.nodeMap[parentNode.addr] = node.parent;

            _index = node.parent;
            node = _tree.nodes[_index];
            parentNode = _tree.nodes[node.parent];
        }
    }

    /**
     * @notice Demotes the node with lower weight to lower level of the tree. It is because to
     *         reduce the average number of traverses required since these nodes are less likely to
     *         be randomly selected.
     *
     * @param _tree  The tree that includes the node to demote.
     * @param _index The initial index of the target node to demote.
     */
    function _demote(Tree storage _tree, uint32 _index) private {
        while (true) {
            Node storage node = _tree.nodes[_index];

            if (_tree.nodes[node.leftChild].weight > _tree.nodes[node.rightChild].weight) {
                if (_tree.nodes[node.leftChild].weight > node.weight) {
                    address nodeAddr = node.addr;
                    node.addr = _tree.nodes[node.leftChild].addr;
                    _tree.nodes[node.leftChild].addr = nodeAddr;

                    uint120 nodeWeight = node.weight;
                    uint120 leftChildWeight = _tree.nodes[node.leftChild].weight;
                    node.weight = leftChildWeight;
                    _tree.nodes[node.leftChild].weight = nodeWeight;

                    unchecked {
                        _tree.nodes[node.leftChild].weightSum -= leftChildWeight - nodeWeight;
                    }

                    _tree.nodeMap[node.addr] = _index;
                    _tree.nodeMap[_tree.nodes[node.leftChild].addr] = node.leftChild;

                    _index = node.leftChild;

                    continue;
                }

                return;
            } else if (_tree.nodes[node.rightChild].weight > node.weight) {
                address nodeAddr = node.addr;
                node.addr = _tree.nodes[node.rightChild].addr;
                _tree.nodes[node.rightChild].addr = nodeAddr;

                uint120 nodeWeight = node.weight;
                uint120 rightChildWeight = _tree.nodes[node.rightChild].weight;
                node.weight = rightChildWeight;
                _tree.nodes[node.rightChild].weight = nodeWeight;

                unchecked {
                    _tree.nodes[node.rightChild].weightSum -= rightChildWeight - nodeWeight;
                }

                _tree.nodeMap[node.addr] = _index;
                _tree.nodeMap[_tree.nodes[node.rightChild].addr] = node.rightChild;

                _index = node.rightChild;

                continue;
            }

            return;
        }
    }

    /**
     * @notice When removing a node, pulls up the remaining nodes with higher weight to higher level
     *         of the tree.
     *
     * @param _tree  The tree that includes the node to remove.
     * @param _index The initial index of the target node to remove.
     */
    function _pullUp(Tree storage _tree, uint32 _index) private {
        while (true) {
            Node storage node = _tree.nodes[_index];
            require(node.addr != address(0), "BalancedWeightTree: node not exists");

            if (node.leftChild == 0) {
                if (node.rightChild == 0) {
                    if (node.parent == 0) {
                        _tree.root = 0;
                    } else if (node.isLeftChild) {
                        _tree.nodes[node.parent].leftChild = 0;
                    } else {
                        _tree.nodes[node.parent].rightChild = 0;
                    }

                    delete _tree.nodes[_index];

                    return;
                } else {
                    node.addr = _tree.nodes[node.rightChild].addr;
                    node.weight = _tree.nodes[node.rightChild].weight;
                    node.weightSum = _tree.nodes[node.rightChild].weightSum;

                    _tree.nodeMap[_tree.nodes[node.rightChild].addr] = _index;

                    _index = node.rightChild;
                }
            } else if (node.rightChild == 0) {
                node.addr = _tree.nodes[node.leftChild].addr;
                node.weight = _tree.nodes[node.leftChild].weight;
                node.weightSum = _tree.nodes[node.leftChild].weightSum;

                _tree.nodeMap[_tree.nodes[node.leftChild].addr] = _index;

                _index = node.leftChild;
            } else if (_tree.nodes[node.leftChild].weight > _tree.nodes[node.rightChild].weight) {
                node.addr = _tree.nodes[node.leftChild].addr;
                node.weight = _tree.nodes[node.leftChild].weight;
                unchecked {
                    node.weightSum =
                        _tree.nodes[node.leftChild].weightSum +
                        _tree.nodes[node.rightChild].weightSum;
                }

                _tree.nodeMap[_tree.nodes[node.leftChild].addr] = _index;

                _index = node.leftChild;
            } else {
                node.addr = _tree.nodes[node.rightChild].addr;
                node.weight = _tree.nodes[node.rightChild].weight;
                unchecked {
                    node.weightSum =
                        _tree.nodes[node.leftChild].weightSum +
                        _tree.nodes[node.rightChild].weightSum;
                }

                _tree.nodeMap[_tree.nodes[node.rightChild].addr] = _index;

                _index = node.rightChild;
            }
        }
    }
}
