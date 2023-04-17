package rollup

import (
	"errors"

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/wemixkanvas/kanvas/bindings/bindings"
	"github.com/wemixkanvas/kanvas/components/node/eth"
)

var ErrNilProof = errors.New("output root proof is nil")
var ErrUnknownOutputRootProofVersion = errors.New("unknown output root proof version")
var ErrVersionNotMatched = errors.New("output root version is not matched")

var V0 = [32]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
var V1 = [32]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}

func L2OutputRootVersion(cfg *Config, timestamp uint64) [32]byte {
	if cfg.IsBlue(timestamp) {
		return V1
	} else {
		return V0
	}
}

// ComputeL2OutputRootV0 computes the L2 output root by hashing an output root proof.
func ComputeL2OutputRoot(proofElements *bindings.TypesOutputRootProof) (eth.Bytes32, error) {
	if proofElements.Version == V0 {
		return ComputeL2OutputRootV0(proofElements)
	} else if proofElements.Version == V1 {
		return ComputeL2OutputRootV1(proofElements)
	} else {
		return eth.Bytes32{}, ErrUnknownOutputRootProofVersion
	}
}

// ComputeL2OutputRootV0 computes the L2 output root by hashing an output root proof ().
func ComputeL2OutputRootV0(proofElements *bindings.TypesOutputRootProof) (eth.Bytes32, error) {
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
	)
	return eth.Bytes32(digest), nil
}

// ComputeL2OutputRootV1 computes the L2 output root by hashing an output root proof.
func ComputeL2OutputRootV1(proofElements *bindings.TypesOutputRootProof) (eth.Bytes32, error) {
	if proofElements == nil {
		return eth.Bytes32{}, ErrNilProof
	}
	if proofElements.Version != V1 {
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
