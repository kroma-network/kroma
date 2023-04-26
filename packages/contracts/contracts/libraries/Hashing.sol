// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import { Encoding } from "./Encoding.sol";
import { RLPWriter } from "./rlp/RLPWriter.sol";
import { Types } from "./Types.sol";

/**
 * @title Hashing
 * @notice Hashing handles Kroma's various different hashing schemes.
 */
library Hashing {
    /**
     * @notice Computes the hash of the RLP encoded L2 transaction that would be generated when a
     *         given deposit is sent to the L2 system. Useful for searching for a deposit in the L2
     *         system.
     *
     * @param _tx User deposit transaction to hash.
     *
     * @return Hash of the RLP encoded L2 deposit transaction.
     */
    function hashDepositTransaction(Types.UserDepositTransaction memory _tx)
        internal
        pure
        returns (bytes32)
    {
        return keccak256(Encoding.encodeDepositTransaction(_tx));
    }

    /**
     * @notice Computes the deposit transaction's "source hash", a value that guarantees the hash
     *         of the L2 transaction that corresponds to a deposit is unique and is
     *         deterministically generated from L1 transaction data.
     *
     * @param _l1BlockHash Hash of the L1 block where the deposit was included.
     * @param _logIndex    The index of the log that created the deposit transaction.
     *
     * @return Hash of the deposit transaction's "source hash".
     */
    function hashDepositSource(bytes32 _l1BlockHash, uint64 _logIndex)
        internal
        pure
        returns (bytes32)
    {
        bytes32 depositId = keccak256(abi.encode(_l1BlockHash, _logIndex));
        return keccak256(abi.encode(bytes32(0), depositId));
    }

    /**
     * @notice Hashes the cross domain message based on the version that is encoded into the
     *         message nonce.
     *
     * @param _nonce    Message nonce with version encoded into the first two bytes.
     * @param _sender   Address of the sender of the message.
     * @param _target   Address of the target of the message.
     * @param _value    ETH value to send to the target.
     * @param _gasLimit Gas limit to use for the message.
     * @param _data     Data to send with the message.
     *
     * @return Hashed cross domain message.
     */
    function hashCrossDomainMessage(
        uint256 _nonce,
        address _sender,
        address _target,
        uint256 _value,
        uint256 _gasLimit,
        bytes memory _data
    ) internal pure returns (bytes32) {
        (, uint16 version) = Encoding.decodeVersionedNonce(_nonce);
        if (version == 0) {
            return hashCrossDomainMessageV0(_nonce, _sender, _target, _value, _gasLimit, _data);
        } else {
            revert("Hashing: unknown cross domain message version");
        }
    }

    /**
     * @notice Hashes a cross domain message based on the V0 (current) encoding.
     *
     * @param _nonce    Message nonce.
     * @param _sender   Address of the sender of the message.
     * @param _target   Address of the target of the message.
     * @param _value    ETH value to send to the target.
     * @param _gasLimit Gas limit to use for the message.
     * @param _data     Data to send with the message.
     *
     * @return Hashed cross domain message.
     */
    function hashCrossDomainMessageV0(
        uint256 _nonce,
        address _sender,
        address _target,
        uint256 _value,
        uint256 _gasLimit,
        bytes memory _data
    ) internal pure returns (bytes32) {
        return
            keccak256(
                Encoding.encodeCrossDomainMessageV0(
                    _nonce,
                    _sender,
                    _target,
                    _value,
                    _gasLimit,
                    _data
                )
            );
    }

    /**
     * @notice Derives the withdrawal hash according to the encoding in the L2 Withdrawer contract
     *
     * @param _tx Withdrawal transaction to hash.
     *
     * @return Hashed withdrawal transaction.
     */
    function hashWithdrawal(Types.WithdrawalTransaction memory _tx)
        internal
        pure
        returns (bytes32)
    {
        return
            keccak256(
                abi.encode(_tx.nonce, _tx.sender, _tx.target, _tx.value, _tx.gasLimit, _tx.data)
            );
    }

    /**
     * @notice Hashes the various elements of an output root proof into an output root hash which
     *         can be used to check if the proof is valid.
     *
     * @param _outputRootProof Output root proof which should hash to an output root.
     *
     * @return Hashed output root proof.
     */
    function hashOutputRootProof(Types.OutputRootProof memory _outputRootProof)
        internal
        pure
        returns (bytes32)
    {
        if (_outputRootProof.version == bytes32(uint256(0))) {
            return hashOutputRootProofV0(_outputRootProof);
        } else if (_outputRootProof.version == bytes32(uint256(1))) {
            return hashOutputRootProofV1(_outputRootProof);
        } else {
            revert("Hashing: unknown output root proof version");
        }
    }

    /**
     * @notice Hashes the various elements of an output root proof into an output root hash which
     *         can be used to check if the proof is valid. (version 0)
     *
     * @param _outputRootProof Output root proof which should hash to an output root.
     *
     * @return Hashed output root proof.
     */
    function hashOutputRootProofV0(Types.OutputRootProof memory _outputRootProof)
        internal
        pure
        returns (bytes32)
    {
        return
            keccak256(
                abi.encode(
                    _outputRootProof.version,
                    _outputRootProof.stateRoot,
                    _outputRootProof.messagePasserStorageRoot,
                    _outputRootProof.blockHash
                )
            );
    }

    /**
     * @notice Hashes the various elements of an output root proof into an output root hash which
     *         can be used to check if the proof is valid. (version 1)
     *
     * @param _outputRootProof Output root proof which should hash to an output root.
     *
     * @return Hashed output root proof.
     */
    function hashOutputRootProofV1(Types.OutputRootProof memory _outputRootProof)
        internal
        pure
        returns (bytes32)
    {
        return
            keccak256(
                abi.encode(
                    _outputRootProof.version,
                    _outputRootProof.stateRoot,
                    _outputRootProof.messagePasserStorageRoot,
                    _outputRootProof.blockHash,
                    _outputRootProof.nextBlockHash
                )
            );
    }

    /**
     * @notice Hashes the various elements of a block header into a block hash.
     *
     * @param _publicInput Public input which should hash to a block hash.
     * @param _rlps        Pre-RLP encoded data which should hash to a block hash.
     *
     * @return Hashed block header.
     */
    function hashBlockHeader(
        Types.PublicInput memory _publicInput,
        Types.BlockHeaderRLP memory _rlps
    ) internal pure returns (bytes32) {
        uint256 length = 16;
        if (_rlps.withdrawalsRoot.length > 0) {
            length += 1;
        }
        bytes[] memory raw = new bytes[](length);
        raw[0] = _rlps.parentHash;
        raw[1] = _rlps.uncleHash;
        raw[2] = RLPWriter.writeAddress(_publicInput.coinbase);
        raw[3] = RLPWriter.writeBytes(abi.encodePacked(_publicInput.stateRoot));
        raw[4] = RLPWriter.writeBytes(abi.encodePacked(_publicInput.transactionsRoot));
        raw[5] = _rlps.receiptsRoot;
        raw[6] = _rlps.logsBloom;
        raw[7] = RLPWriter.writeUint(_publicInput.difficulty);
        raw[8] = RLPWriter.writeUint(_publicInput.number);
        raw[9] = RLPWriter.writeUint(_publicInput.gasLimit);
        raw[10] = _rlps.gasUsed;
        raw[11] = RLPWriter.writeUint(_publicInput.timestamp);
        raw[12] = _rlps.extraData;
        raw[13] = _rlps.mixHash;
        raw[14] = _rlps.nonce;
        raw[15] = RLPWriter.writeUint(_publicInput.baseFee);
        if (length == 17) {
            raw[16] = _rlps.withdrawalsRoot;
        }
        return keccak256(RLPWriter.writeList(raw));
    }

    /**
     * @notice Hashes the various elements of a public input into a public input hash.
     *
     * @param _prevStateRoot Previous state root.
     * @param _publicInput   Public input which should hash to a public input hash.
     * @param _chainId       Chain ID.
     * @param _dummyHashes   Dummy hashes returned from generateDummyHashes().
     *
     * @return Hashed block header.
     */
    function hashPublicInput(
        bytes32 _prevStateRoot,
        Types.PublicInput memory _publicInput,
        uint256 _chainId,
        bytes32[] memory _dummyHashes
    ) internal pure returns (bytes32) {
        return
            keccak256(
                abi.encodePacked(
                    _publicInput.coinbase,
                    _publicInput.timestamp,
                    _publicInput.number,
                    _publicInput.difficulty,
                    _publicInput.gasLimit,
                    _publicInput.baseFee,
                    _chainId,
                    _publicInput.txHashes.length,
                    _prevStateRoot,
                    _publicInput.stateRoot,
                    _publicInput.txHashes,
                    _dummyHashes
                )
            );
    }

    /**
     * @notice Generates a bytes32 array filled with a dummy hash for the given length.
     *
     * @param _dummyHashes Dummy hash.
     * @param _length      A length of the array.
     *
     * @return Bytes32 array filled with dummy hash.
     */
    function generateDummyHashes(
        bytes32 _dummyHashes,
        uint256 _length
    ) internal pure returns (bytes32[] memory) {
        bytes32[] memory hashes = new bytes32[](_length);
        for (uint256 i = 0; i < _length; i++) {
            hashes[i] = _dummyHashes;
        }
        return hashes;
    }
}
