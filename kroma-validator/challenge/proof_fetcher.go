package challenge

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
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

func (c *Client) FetchProofAndPair(ctx context.Context, trace string) (*ProofAndPair, error) {
	proveResult, err := send[zkEVMProveResponse](ctx, c, "prove", []any{trace})
	if err != nil {
		return nil, fmt.Errorf("failed to request prove: %w", err)
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

type HexBytes []byte

type ZkVMProofResponse struct {
	RequestStatus RequestStatusType `json:"request_status"`
	VKeyHash      common.Hash       `json:"vkey_hash"`
	RequestID     string            `json:"request_id"`
	PublicValues  HexBytes          `json:"public_values"`
	Proof         HexBytes          `json:"proof"`
}

type ZkVMProofFetcher interface {
	Spec(ctx context.Context) (*SpecResponse, error)
	RequestProve(ctx context.Context, blockHash string, l1Head string, witness string) (*RequestStatusType, error)
	GetProof(ctx context.Context, blockHash string, l1Head string) (*ZkVMProofResponse, error)
}

func (c *Client) RequestProve(ctx context.Context, blockHash string, l1Head string, witness string) (*RequestStatusType, error) {
	return send[RequestStatusType](ctx, c, "requestProve", []any{blockHash, l1Head, witness})
}

func (c *Client) GetProof(ctx context.Context, blockHash string, l1Head string) (*ZkVMProofResponse, error) {
	return send[ZkVMProofResponse](ctx, c, "getProof", []any{blockHash, l1Head})
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
