package challenge

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum-optimism/optimism/op-service/client"
	"github.com/ethereum/go-ethereum/common"
)

type ZkEVMProofFetcher struct {
	rpc client.RPC
}

type ZkEVMProveResponse struct {
	FinalPair []byte `json:"final_pair"`
	Proof     []byte `json:"proof"`
}

type ProofAndPair struct {
	Proof []*big.Int
	Pair  []*big.Int
}

func NewZkEVMProofFetcher(rpc client.RPC) *ZkEVMProofFetcher {
	return &ZkEVMProofFetcher{rpc}
}

func (z *ZkEVMProofFetcher) FetchProofAndPair(ctx context.Context, trace string) (*ProofAndPair, error) {
	var output *ZkEVMProveResponse
	err := z.rpc.CallContext(ctx, &output, "prove", trace)
	if err != nil {
		return nil, fmt.Errorf("failed to request prove: %w", err)
	}

	proofAndPair := &ProofAndPair{
		Proof: decode(output.Proof),
		Pair:  decode(output.FinalPair),
	}

	return proofAndPair, nil
}

func decode(data []byte) []*big.Int {
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

type ZkVMProofFetcher struct {
	rpc client.RPC
}

type HexBytes []byte

type ZkVMProofResponse struct {
	RequestStatus RequestStatusType `json:"request_status"`
	VKeyHash      common.Hash       `json:"vkey_hash"`
	RequestID     string            `json:"request_id"`
	PublicValues  HexBytes          `json:"public_values"`
	Proof         HexBytes          `json:"proof"`
}

func NewZkVMProofFetcher(rpc client.RPC) *ZkVMProofFetcher {
	return &ZkVMProofFetcher{rpc}
}

func (z *ZkVMProofFetcher) Spec(ctx context.Context) (*SpecResponse, error) {
	var output *SpecResponse
	err := z.rpc.CallContext(ctx, &output, "spec")
	return output, err
}

func (z *ZkVMProofFetcher) RequestProve(ctx context.Context, blockHash string, l1Head string, witness string) (*RequestStatusType, error) {
	var output *RequestStatusType
	err := z.rpc.CallContext(ctx, &output, "requestProve", blockHash, l1Head, witness)
	return output, err
}

func (z *ZkVMProofFetcher) GetProof(ctx context.Context, blockHash string, l1Head string) (*ZkVMProofResponse, error) {
	var output *ZkVMProofResponse
	err := z.rpc.CallContext(ctx, &output, "getProof", blockHash, l1Head)
	return output, err
}

// UnmarshalJSON handles the conversion from a hex string to a byte array.
func (h *HexBytes) UnmarshalJSON(data []byte) error {
	// Remove quotes around the hex string
	str := string(data)
	if len(str) >= 2 && str[0] == '"' && str[len(str)-1] == '"' {
		str = str[1 : len(str)-1]
	}

	// Decode the hex string to byte array
	str = strings.TrimPrefix(str, "0x")
	decoded, err := hex.DecodeString(str)
	if err != nil {
		return err
	}

	*h = decoded
	return nil
}
