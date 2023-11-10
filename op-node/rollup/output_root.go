package rollup

import (
	"errors"

	"github.com/ethereum-optimism/optimism/op-service/eth"

	"github.com/kroma-network/kroma/kroma-bindings/bindings"
)

var (
	ErrNilProof                      = errors.New("output root proof is nil")
	ErrUnknownOutputRootProofVersion = errors.New("unknown output root proof version")
	ErrVersionNotMatched             = errors.New("output root version is not matched")
)

// ComputeL2OutputRoot computes the L2 output root by hashing an output root proof.
func ComputeL2OutputRoot(proofElements *bindings.TypesOutputRootProof) (eth.Bytes32, error) {
	if proofElements == nil {
		return eth.Bytes32{}, ErrNilProof
	}

	if eth.Bytes32(proofElements.Version) != eth.OutputVersionV0 {
		return eth.Bytes32{}, errors.New("unsupported output root version")
	}
	l2Output := eth.OutputV0{
		StateRoot:                eth.Bytes32(proofElements.StateRoot),
		MessagePasserStorageRoot: proofElements.MessagePasserStorageRoot,
		BlockHash:                proofElements.BlockHash,
		NextBlockHash:            proofElements.NextBlockHash,
	}
	return eth.OutputRoot(&l2Output), nil
}

// computeL2OutputRootV0 computes the L2 output root by hashing an output root proof.
func computeL2OutputRootV0(proofElements *bindings.TypesOutputRootProof) (eth.Bytes32, error) {
	if proofElements == nil {
		return eth.Bytes32{}, ErrNilProof
	}
	if proofElements.Version != eth.OutputVersionV0 {
		return eth.Bytes32{}, ErrVersionNotMatched
	}

	l2Output := eth.OutputV0{
		StateRoot:                eth.Bytes32(proofElements.StateRoot),
		MessagePasserStorageRoot: proofElements.MessagePasserStorageRoot,
		BlockHash:                proofElements.BlockHash,
		NextBlockHash:            proofElements.NextBlockHash,
	}
	return eth.OutputRoot(&l2Output), nil
}
