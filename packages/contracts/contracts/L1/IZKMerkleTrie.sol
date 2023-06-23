// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

/**
 * @title IZKMerkleTrie
 */
interface IZKMerkleTrie {
    /**
     * @notice Verifies a proof that a given key/value pair is present in the trie.
     *
     * @param _key    Key of the node to search for, as a hex string.
     * @param _value  Value of the node to search for, as a hex string.
     * @param _proofs Merkle trie inclusion proof for the desired node.
     * @param _root   Known root of the Merkle trie. Used to verify that the included proof is
     *                correctly constructed.
     *
     * @return Whether or not the proof is valid.
     */
    function verifyInclusionProof(
        bytes32 _key,
        bytes memory _value,
        bytes[] memory _proofs,
        bytes32 _root
    ) external view returns (bool);
}
