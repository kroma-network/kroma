// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { Bytes } from "../libraries/Bytes.sol";
import { NodeReader } from "../libraries/NodeReader.sol";
import { IZKMerkleTrie } from "./interfaces/IZKMerkleTrie.sol";
import { ZKTrieHasher } from "./ZKTrieHasher.sol";

/**
 * @custom:proxied
 * @title ZKMerkleTrie
 * @notice The ZKMerkleTrie is contract which can produce a hash according to ZKTrie.
 *         This owns an interface of Poseidon2 that is required to compute hash used by ZKTrie.
 */
contract ZKMerkleTrie is IZKMerkleTrie, ZKTrieHasher {
    /**
     * @notice Struct representing a node in the trie.
     */
    struct TrieNode {
        bytes encoded;
        NodeReader.Node decoded;
    }

    /**
     * @notice Magic hash which indicates
     *         See https://github.com/kroma-network/zktrie/blob/main/trie/zk_trie_proof.go.
     */
    bytes32 private constant MAGIC_SMT_BYTES_HASH =
        keccak256(
            hex"5448495320495320534f4d45204d4147494320425954455320464f5220534d54206d3172525867503278704449"
        );

    /**
     * @param _poseidon2 The address of poseidon2 contract.
     */
    constructor(address _poseidon2) ZKTrieHasher(_poseidon2) {}

    /**
     * @notice Checks if a given bytes is MAGIC_SMT_BYTES_HASH.
     *
     * @param _value Bytes to be compared.
     */
    function isMagicSmtBytesHash(bytes memory _value) private pure returns (bool) {
        return keccak256(_value) == MAGIC_SMT_BYTES_HASH;
    }

    /**
     * @inheritdoc IZKMerkleTrie
     */
    function verifyInclusionProof(
        bytes32 _key,
        bytes memory _value,
        bytes[] memory _proofs,
        bytes32 _root
    ) external view returns (bool) {
        (bool exists, bytes memory value) = this.get(_key, _proofs, _root);
        return (exists && Bytes.equal(_value, value));
    }

    /**
     * @notice Retrieves the value associated with a given key.
     *
     * @param _key    Key to search for, as hex bytes.
     * @param _proofs Merkle trie inclusion proof for the key.
     * @param _root   Known root of the Merkle trie.
     *
     * @return Whether or not the key exists.
     * @return Value of the key if it exists.
     */
    function get(
        bytes32 _key,
        bytes[] memory _proofs,
        bytes32 _root
    ) external view returns (bool, bytes memory) {
        require(_proofs.length >= 2, "ZKMerkleTrie: provided proof is too short");
        require(
            isMagicSmtBytesHash(_proofs[_proofs.length - 1]),
            "ZKMerkleTrie: the last item is not magic hash"
        );
        bytes32 key = _hashElem(_key);
        TrieNode[] memory nodes = _parseProofs(_proofs);
        NodeReader.Node memory currentNode;
        bytes32 computedKey = bytes32(0);
        bool exists = false;
        bool empty = false;
        bytes memory value = bytes("");
        for (uint256 i = nodes.length - 2; i >= 0; ) {
            currentNode = nodes[i].decoded;
            if (currentNode.nodeType == NodeReader.NodeType.MIDDLE) {
                bool isLeft = _isLeft(key, i);
                if (isLeft) {
                    require(computedKey == currentNode.childL, "ZKMerkleTrie: invalid key L");
                } else {
                    require(computedKey == currentNode.childR, "ZKMerkleTrie: invalid key R");
                }
                computedKey = _hashFixed2Elems(currentNode.childL, currentNode.childR);
            } else if (currentNode.nodeType == NodeReader.NodeType.LEAF) {
                require(!exists && !empty, "ZKMerkleTrie: duplicated terminal node");
                exists = currentNode.nodeKey == key;
                if (!exists) {
                    break;
                }
                computedKey = _hashFixed3Elems(
                    bytes32(uint256(1)),
                    currentNode.nodeKey,
                    _valueHash(currentNode.compressedFlags, currentNode.valuePreimage)
                );
                bytes32[] memory valuePreimage = currentNode.valuePreimage;
                uint256 len = valuePreimage.length;
                assembly {
                    value := valuePreimage
                    mstore(value, mul(len, 32))
                }
                if (currentNode.keyPreimage != bytes32(0)) {
                    // NOTE(chokobole): The comparison order is important, because in this setting,
                    // first condition is mostly evaluted to be true. When we're sure about
                    // database preimage, then we need to enable just one of check below!
                    require(
                        currentNode.keyPreimage == _key || currentNode.keyPreimage == key,
                        "ZKMerkleTrie: invalid key preimage"
                    );
                }
            } else if (currentNode.nodeType == NodeReader.NodeType.EMPTY) {
                require(!exists && !empty, "ZKMerkleTrie: duplicated terminal node");
                empty = true;
            }
            if (i == 0) {
                require(computedKey == _root, "ZKMerkeTrie: invalid root");
                break;
            }
            unchecked {
                --i;
            }
        }
        return (exists, value);
    }

    /**
     * @notice Parses an array of proof elements into a new array that contains both the original
     *         encoded element and the decoded element.
     *
     * @param _proofs Array of proof elements to parse.
     *
     * @return TrieNode parsed into easily accessible structs.
     */
    function _parseProofs(bytes[] memory _proofs) private pure returns (TrieNode[] memory) {
        uint256 length = _proofs.length;
        TrieNode[] memory nodes = new TrieNode[](length);
        // NOTE(chokobole): Last proof is MAGIC_SMT_BYTES_HASH!
        for (uint256 i = 0; i < length - 1; ) {
            NodeReader.Node memory node = NodeReader.readNode(_proofs[i]);
            nodes[i] = TrieNode({ encoded: _proofs[i], decoded: node });
            unchecked {
                ++i;
            }
        }
        return nodes;
    }

    /**
     * @notice Computes merkle path at index n based on a given keyPreimage.
     *
     * @param _keyPreimage Keypreimage.
     * @param _n           Bit to mask.
     *
     * @return Whether merkle path is left or not.
     */
    function _isLeft(bytes32 _keyPreimage, uint256 _n) private pure returns (bool) {
        require(_n < 256, "ZKMerkleTrie: too long depth");
        return _keyPreimage & bytes32(1 << _n) == 0;
    }
}
