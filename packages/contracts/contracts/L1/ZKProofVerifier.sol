// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

// Contracts
import { ZKVerifier } from "contracts/L1/ZKVerifier.sol";

// Libraries
import { Hashing } from "contracts/libraries/Hashing.sol";
import { Predeploys } from "contracts/libraries/Predeploys.sol";
import { Types } from "contracts/libraries/Types.sol";

// Interfaces
import { ISemver } from "contracts/universal/ISemver.sol";
import { ISP1Verifier } from "contracts/vendor/ISP1Verifier.sol";
import { IZKMerkleTrie } from "contracts/L1/interfaces/IZKMerkleTrie.sol";

/// @custom:proxied true
/// @title ZKProofVerifier
/// @notice The ZKProofVerifier contract verifies public inputs and corresponding ZK proofs.
///         Currently it verifies zkEVM proofs using ZKVerifier contract, and zkVM proofs using
///         SP1Verifier contract.
contract ZKProofVerifier is ISemver {
    /// @notice Address of the ZKVerifier contract.
    ZKVerifier internal immutable ZK_VERIFIER;

    /// @notice The dummy transaction hash for zkEVM proofs. This is used to pad if the
    ///         number of transactions is less than MAX_TXS. This is same as:
    ///         unsignedTx = {
    ///           nonce: 0,
    ///           gasLimit: 0,
    ///           gasPrice: 0,
    ///           to: address(0),
    ///           value: 0,
    ///           data: '0x',
    ///           chainId: CHAIN_ID,
    ///         }
    ///         signature = sign(unsignedTx, 0x1)
    ///         dummyHash = keccak256(rlp({
    ///           ...unsignedTx,
    ///           signature,
    ///         }))
    bytes32 internal immutable DUMMY_HASH;

    /// @notice The maximum number of transactions for zkEVM proofs.
    uint256 internal immutable MAX_TXS;

    /// @notice Address that has the ability to verify the merkle proof in ZKTrie.
    address internal immutable ZK_MERKLE_TRIE;

    /// @notice Address of the SP1VerifierGateway contract.
    ISP1Verifier internal immutable SP1_VERIFIER;

    /// @notice The verification key for the zkVM program.
    bytes32 internal immutable ZKVM_PROGRAM_V_KEY;

    /// @notice Reverts when the zkVM program verification key is invalid.
    error InvalidZkVmVKey();

    /// @notice Reverts when the public input is invalid.
    error InvalidPublicInput();

    /// @notice Reverts when the ZK proof is invalid.
    error InvalidZkProof();

    /// @notice Reverts when the inclusion proof is invalid.
    error InvalidInclusionProof();

    /// @notice Reverts when the block hash is mismatched between source and destination output root
    ///         proof. (only for zkEVM proof)
    error BlockHashMismatchedBtwSrcAndDst();

    /// @notice Reverts when the source output root is mismatched.
    error SrcOutputMismatched();

    /// @notice Reverts when the destination output root is matched. (only for fault proof)
    error DstOutputMatched();

    /// @notice Reverts when the block hash is mismatched.
    error BlockHashMismatched();

    /// @notice Reverts when the state root is mismatched.
    error StateRootMismatched();

    /// @notice Semantic version.
    /// @custom:semver 1.0.0
    string public constant version = "1.0.0";

    /// @notice Constructs the ZKProofVerifier contract.
    /// @param _zkVerifier Address of the ZKVerifier contract.
    /// @param _dummyHash Dummy hash for zkEVM proofs.
    /// @param _maxTxs Number of max transactions per block for zkEVM proofs.
    /// @param _zkMerkleTrie Address of the ZKMerkleTrie contract.
    /// @param _sp1Verifier Address of the SP1VerifierGateway contract.
    /// @param _zkVmProgramVKey The verification key for the zkVM program.
    constructor(
        ZKVerifier _zkVerifier,
        bytes32 _dummyHash,
        uint256 _maxTxs,
        address _zkMerkleTrie,
        ISP1Verifier _sp1Verifier,
        bytes32 _zkVmProgramVKey
    ) {
        ZK_VERIFIER = _zkVerifier;
        DUMMY_HASH = _dummyHash;
        MAX_TXS = _maxTxs;
        ZK_MERKLE_TRIE = _zkMerkleTrie;
        SP1_VERIFIER = _sp1Verifier;
        ZKVM_PROGRAM_V_KEY = _zkVmProgramVKey;
    }

    /// @notice Getter for the address of ZKVerifier contract.
    function zkVerifier() external view returns (ZKVerifier) {
        return ZK_VERIFIER;
    }

    /// @notice Getter for the dummy transaction hash for zkEVM proofs.
    function dummyHash() external view returns (bytes32) {
        return DUMMY_HASH;
    }

    /// @notice Getter for the maximum number of transactions for zkEVM proofs.
    function maxTxs() external view returns (uint256) {
        return MAX_TXS;
    }

    /// @notice Getter for the address of ZKMerkleTrie contract.
    function zkMerkleTrie() external view returns (address) {
        return ZK_MERKLE_TRIE;
    }

    /// @notice Getter for the address of SP1VerifierGateway contract.
    function sp1Verifier() external view returns (ISP1Verifier) {
        return SP1_VERIFIER;
    }

    /// @notice Getter for the verification key for the zkVM program.
    function zkVmProgramVKey() external view returns (bytes32) {
        return ZKVM_PROGRAM_V_KEY;
    }

    /// @notice Verifies zkEVM public inputs and proof.
    /// @param _zkEvmProof The public input and proof using zkEVM.
    /// @param _storedSrcOutput The stored source output root.
    /// @param _storedDstOutput The stored destination output root. It will only be used for fault proving.
    /// @return publicInputHash_ Hash of public input.
    function verifyZkEvmProof(
        Types.ZkEvmProof calldata _zkEvmProof,
        bytes32 _storedSrcOutput,
        bytes32 _storedDstOutput
    ) external view returns (bytes32 publicInputHash_) {
        Types.PublicInputProof calldata publicInputProof = _zkEvmProof.publicInputProof;

        if (
            publicInputProof.srcOutputRootProof.nextBlockHash !=
            publicInputProof.dstOutputRootProof.latestBlockhash
        ) revert BlockHashMismatchedBtwSrcAndDst();

        _validatePublicInputOutput(
            _storedSrcOutput,
            _storedDstOutput,
            Hashing.hashOutputRootProof(publicInputProof.srcOutputRootProof),
            Hashing.hashOutputRootProof(publicInputProof.dstOutputRootProof)
        );
        _validateZkEvmPublicInput(
            publicInputProof.dstOutputRootProof,
            publicInputProof.publicInput,
            publicInputProof.rlps
        );
        _validateWithdrawalStorageRoot(
            publicInputProof.merkleProof,
            publicInputProof.l2ToL1MessagePasserBalance,
            publicInputProof.l2ToL1MessagePasserCodeHash,
            publicInputProof.dstOutputRootProof.messagePasserStorageRoot,
            publicInputProof.dstOutputRootProof.stateRoot
        );

        publicInputHash_ = _hashZkEvmPublicInput(
            publicInputProof.srcOutputRootProof.stateRoot,
            publicInputProof.publicInput
        );

        if (!ZK_VERIFIER.verify(_zkEvmProof.proof, _zkEvmProof.pair, publicInputHash_))
            revert InvalidZkProof();
    }

    /// @notice Verifies zkVM public inputs and proof.
    /// @param _zkVmProof The public input and proof using zkVM.
    /// @param _storedSrcOutput The stored source output root.
    /// @param _storedDstOutput The stored destination output root. It will only be used for fault proving.
    /// @param _storedL1Head The stored L1 block hash.
    /// @return publicInputHash_ Hash of public input.
    function verifyZkVmProof(
        Types.ZkVmProof calldata _zkVmProof,
        bytes32 _storedSrcOutput,
        bytes32 _storedDstOutput,
        bytes32 _storedL1Head
    ) external view returns (bytes32 publicInputHash_) {
        if (_zkVmProof.zkVmProgramVKey != ZKVM_PROGRAM_V_KEY) revert InvalidZkVmVKey();

        _validatePublicInputOutput(
            _storedSrcOutput,
            _storedDstOutput,
            bytes32(_zkVmProof.publicValues[8:40]), // skip ABI-encoding prefix at publicValues[0:8].
            bytes32(_zkVmProof.publicValues[48:80]) // skip ABI-encoding prefix at publicValues[40:48].
        );

        // Check if the L1 block hash is correct.
        // Skip ABI-encoding prefix at publicValues[80:88].
        if (bytes32(_zkVmProof.publicValues[88:120]) != _storedL1Head) revert InvalidPublicInput();

        SP1_VERIFIER.verifyProof(
            ZKVM_PROGRAM_V_KEY,
            _zkVmProof.publicValues,
            _zkVmProof.proofBytes
        );

        publicInputHash_ = keccak256(_zkVmProof.publicValues);
    }

    /// @notice Checks if the public input outputs are valid. Reverts if they are invalid.
    /// @param _storedSrcOutput The stored source output root.
    /// @param _storedDstOutput The stored destination output root.
    /// @param _publicInputSrcOutput The source output root of public input.
    /// @param _publicInputDstOutput The destination output root of public input.
    function _validatePublicInputOutput(
        bytes32 _storedSrcOutput,
        bytes32 _storedDstOutput,
        bytes32 _publicInputSrcOutput,
        bytes32 _publicInputDstOutput
    ) internal pure {
        if (_storedSrcOutput != _publicInputSrcOutput) revert SrcOutputMismatched();
        // If _storedDstOutput is non-zero, it is fault proving case, not validity proving.
        // Then assert _publicInputDstOutput is different with on-chain stored destination output.
        if (_storedDstOutput != bytes32(0)) {
            if (_storedDstOutput == _publicInputDstOutput) revert DstOutputMatched();
        }
    }

    /// @notice Checks if the public input of zkEVM proof is valid. Reverts if it is invalid.
    /// @param _dstOutputRootProof Proof of the destination output root.
    /// @param _publicInput Ingredients to compute the public input used by zkEVM proof verification.
    /// @param _rlps Pre-encoded RLPs to compute the latest block hash of the destination output
    ///              root proof.
    function _validateZkEvmPublicInput(
        Types.OutputRootProof calldata _dstOutputRootProof,
        Types.PublicInput calldata _publicInput,
        Types.BlockHeaderRLP calldata _rlps
    ) internal pure {
        if (_publicInput.stateRoot != _dstOutputRootProof.stateRoot) revert StateRootMismatched();

        // parentBeaconRoot is non-zero for Cancun block
        bytes32 blockHash = _publicInput.parentBeaconRoot != bytes32(0)
            ? Hashing.hashBlockHeaderCancun(_publicInput, _rlps)
            : Hashing.hashBlockHeaderShanghai(_publicInput, _rlps);

        if (_dstOutputRootProof.latestBlockhash != blockHash) revert BlockHashMismatched();
    }

    /// @notice Checks if the L2ToL1MesagePasser account is included in the given state root.
    /// @param _merkleProof Merkle proof of L2ToL1MessagePasser account against the state root.
    /// @param _l2ToL1MessagePasserBalance Balance of the L2ToL1MessagePasser account.
    /// @param _l2ToL1MessagePasserCodeHash Codehash of the L2ToL1MessagePasser account.
    /// @param _messagePasserStorageRoot Storage root of the L2ToL1MessagePasser account.
    /// @param _stateRoot State root.
    function _validateWithdrawalStorageRoot(
        bytes[] calldata _merkleProof,
        bytes32 _l2ToL1MessagePasserBalance,
        bytes32 _l2ToL1MessagePasserCodeHash,
        bytes32 _messagePasserStorageRoot,
        bytes32 _stateRoot
    ) internal view {
        // TODO(chokobole): Can we fix the codeHash?
        bytes memory l2ToL1MessagePasserAccount = abi.encodePacked(
            uint256(0), // nonce
            _l2ToL1MessagePasserBalance, // balance,
            _l2ToL1MessagePasserCodeHash, // codeHash,
            _messagePasserStorageRoot // storage root
        );

        if (
            !IZKMerkleTrie(ZK_MERKLE_TRIE).verifyInclusionProof(
                bytes32(bytes20(Predeploys.L2_TO_L1_MESSAGE_PASSER)),
                l2ToL1MessagePasserAccount,
                _merkleProof,
                _stateRoot
            )
        ) revert InvalidInclusionProof();
    }

    /// @notice Hashes the public input for zkEVM proof with padding dummy transactions.
    /// @param _prevStateRoot Previous state root.
    /// @param _publicInput Ingredients to compute the public input used by zkEVM proof verification.
    /// @return Hash of public input for zkEVM proof.
    function _hashZkEvmPublicInput(
        bytes32 _prevStateRoot,
        Types.PublicInput calldata _publicInput
    ) internal view returns (bytes32) {
        bytes32[] memory dummyHashes;
        if (_publicInput.txHashes.length < MAX_TXS) {
            dummyHashes = Hashing.generateDummyHashes(
                DUMMY_HASH,
                MAX_TXS - _publicInput.txHashes.length
            );
        }

        // NOTE(chokobole): We cannot calculate the Ethereum transaction root solely
        // based on transaction hashes. It is necessary to have access to the original
        // transactions. Considering the imposed constraints and the difficulty
        // of providing a preimage that would generate the desired public input hash
        // from an attacker's perspective, we have decided to omit the verification
        // using the transaction root.
        return Hashing.hashZkEvmPublicInput(_prevStateRoot, _publicInput, dummyHashes);
    }
}
