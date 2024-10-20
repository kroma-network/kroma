package challenge

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
)

type zkEVMProveResponse struct {
	FinalPair []byte `json:"final_pair"`
	Proof     []byte `json:"proof"`
}

type ProofAndPair struct {
	Proof []*big.Int
	Pair  []*big.Int
}

type ZkEVMProofFetcher interface {
	FetchProofAndPair(ctx context.Context, trace string) (*ProofAndPair, error)
}

type ZkVMProofFetcher interface {
}

func (c *Client) FetchProofAndPair(ctx context.Context, trace string) (*ProofAndPair, error) {
	resultBytes, err := c.fetch(ctx, "prove", []any{trace})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch proof and pair: %w", err)
	}

	var proveResult zkEVMProveResponse
	err = json.Unmarshal(resultBytes, &proveResult)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal as zkEVMProveResponse: %w", err)
	}

	proofAndPair := &ProofAndPair{
		Proof: Decode(proveResult.Proof),
		Pair:  Decode(proveResult.FinalPair),
	}

	return proofAndPair, nil
}

func Decode(data []byte) []*big.Int {
	result := make([]*big.Int, len(data)/32)

	for i := 0; i < len(data)/32; i++ {
		// The best is data is given in Big Endian.
		for j := 0; j < 16; j++ {
			data[i*32+j], data[i*32+31-j] = data[i*32+31-j], data[i*32+j]
		}
		result[i] = new(big.Int).SetBytes(data[i*32 : (i+1)*32])
	}

	return result
}
