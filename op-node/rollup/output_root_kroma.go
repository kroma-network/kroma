package rollup

import (
	"github.com/ethereum-optimism/optimism/op-service/eth"

	"github.com/kroma-network/kroma/kroma-bindings/bindings"
)

func ComputeKromaL2Output(proofElements *bindings.TypesOutputRootProof) (eth.Bytes32, error) {
	if proofElements == nil {
		return eth.Bytes32{}, ErrNilProof
	}
	if proofElements.Version != eth.OutputVersionV0 {
		return eth.Bytes32{}, ErrVersionNotMatched
	}

	l2Output := eth.KromaOutputV0{
		OutputV0: eth.OutputV0{
			StateRoot:                eth.Bytes32(proofElements.StateRoot),
			MessagePasserStorageRoot: proofElements.MessagePasserStorageRoot,
			BlockHash:                proofElements.BlockHash,
		},
		NextBlockHash:            proofElements.NextBlockHash,
	}
	return eth.OutputRoot(&l2Output), nil
}

// computeKromaL2OutputRootV0 computes the L2 output root by hashing an output root proof.
func computeKromaL2OutputRootV0(proofElements *bindings.TypesOutputRootProof) (eth.Bytes32, error) {
	if proofElements == nil {
		return eth.Bytes32{}, ErrNilProof
	}
	if proofElements.Version != eth.OutputVersionV0 {
		return eth.Bytes32{}, ErrVersionNotMatched
	}

	l2Output := eth.KromaOutputV0{
		OutputV0: eth.OutputV0{
			StateRoot:                eth.Bytes32(proofElements.StateRoot),
			MessagePasserStorageRoot: proofElements.MessagePasserStorageRoot,
			BlockHash:                proofElements.BlockHash,
		},
		NextBlockHash:            proofElements.NextBlockHash,
	}
	return eth.OutputRoot(&l2Output), nil
}
