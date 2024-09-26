package rollup

import (
	"errors"

	"github.com/ethereum-optimism/optimism/op-service/eth"

	"github.com/kroma-network/kroma/kroma-bindings/bindings"
)

// ComputeKromaL2OutputRoot computes the L2 output root by hashing an output root proof.
func ComputeKromaL2OutputRoot(proofElements *bindings.TypesOutputRootProof) (eth.Bytes32, error) {
	if proofElements == nil {
		return eth.Bytes32{}, ErrNilProof
	}

	if eth.Bytes32(proofElements.Version) != eth.OutputVersionV0 {
		return eth.Bytes32{}, errors.New("unsupported output root version")
	}
	l2Output := eth.KromaOutputV0{
		OutputV0: eth.OutputV0{
			StateRoot:                eth.Bytes32(proofElements.StateRoot),
			MessagePasserStorageRoot: proofElements.MessagePasserStorageRoot,
			BlockHash:                proofElements.LatestBlockhash,
		},
		NextBlockHash: proofElements.NextBlockHash,
	}
	return eth.OutputRoot(&l2Output), nil
}
