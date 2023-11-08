package rollup

import (
	"errors"

	"github.com/ethereum-optimism/optimism/op-bindings/bindings"
	"github.com/ethereum-optimism/optimism/op-node/eth"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	ErrNilProof                      = errors.New("output root proof is nil")
	ErrUnknownOutputRootProofVersion = errors.New("unknown output root proof version")
	ErrVersionNotMatched             = errors.New("output root version is not matched")
)

var V0 = [32]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

func L2OutputRootVersion(cfg *Config, timestamp uint64) [32]byte {
	return V0
}

// ComputeL2OutputRoot computes the L2 output root by hashing an output root proof.
func ComputeL2OutputRoot(proofElements *bindings.TypesOutputRootProof) (eth.Bytes32, error) {
	if proofElements.Version == V0 {
		return computeL2OutputRootV0(proofElements)
	} else {
		return eth.Bytes32{}, ErrUnknownOutputRootProofVersion
	}
}

// computeL2OutputRootV0 computes the L2 output root by hashing an output root proof.
func computeL2OutputRootV0(proofElements *bindings.TypesOutputRootProof) (eth.Bytes32, error) {
	if proofElements == nil {
		return eth.Bytes32{}, ErrNilProof
	}
	if proofElements.Version != V0 {
		return eth.Bytes32{}, ErrVersionNotMatched
	}

	digest := crypto.Keccak256Hash(
		proofElements.Version[:],
		proofElements.StateRoot[:],
		proofElements.MessagePasserStorageRoot[:],
		proofElements.BlockHash[:],
		proofElements.NextBlockHash[:],
	)
	return eth.Bytes32(digest), nil
}
