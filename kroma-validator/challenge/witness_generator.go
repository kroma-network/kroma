package challenge

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

const (
	WitnessNone       RequestStatusType = "None"
	WitnessRequested  RequestStatusType = "Requested"
	WitnessProcessing RequestStatusType = "Processing"
	WitnessCompleted  RequestStatusType = "Completed"
)

type RequestStatusType string

type WitnessResponse struct {
	RequestStatus RequestStatusType `json:"status"`
	VKey          common.Hash       `json:"vkey_hash"`
	PublicValues  []byte            `json:"witness"`
}

type VKeyAndPublicValues struct {
	VKey         common.Hash
	PublicValues []byte
}

type WitnessGenerator interface {
	RequestWitness(ctx context.Context, blockHash common.Hash, l1Head common.Hash) (RequestStatusType, error)
	GetWitness(ctx context.Context, blockHash common.Hash, l1Head common.Hash) (*VKeyAndPublicValues, error)
}

func (c *Client) RequestWitness(ctx context.Context, blockHash common.Hash, l1Head common.Hash) (RequestStatusType, error) {
	resultBytes, err := c.fetch(ctx, "requestWitness", []any{blockHash, l1Head})
	if err != nil {
		return WitnessNone, fmt.Errorf("failed to request witness: %w", err)
	}

	var requestStatus RequestStatusType
	err = json.Unmarshal(resultBytes, &requestStatus)
	if err != nil {
		return WitnessNone, fmt.Errorf("failed to unmarshal as RequestStatusType: %w", err)
	}

	return requestStatus, nil
}

func (c *Client) GetWitness(ctx context.Context, blockHash common.Hash, l1Head common.Hash) (*VKeyAndPublicValues, error) {
	resultBytes, err := c.fetch(ctx, "getWitness", []any{blockHash, l1Head})
	if err != nil {
		return nil, fmt.Errorf("failed to get witness: %w", err)
	}

	var witnessResult WitnessResponse
	err = json.Unmarshal(resultBytes, &witnessResult)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal as WitnessResponse: %w", err)
	}

	if witnessResult.RequestStatus != WitnessCompleted {
		return nil, nil
	}

	return &VKeyAndPublicValues{
		VKey:         witnessResult.VKey,
		PublicValues: witnessResult.PublicValues,
	}, nil
}

//// UnmarshalJSON handles the conversion from a hex string to a byte array
//func (h *HexBytes) UnmarshalJSON(data []byte) error {
//	// Remove quotes around the hex string
//	str := string(data)
//	if len(str) >= 2 && str[0] == '"' && str[len(str)-1] == '"' {
//		str = str[1 : len(str)-1]
//	}
//
//	// Decode the hex string to byte array
//	decoded, err := hex.DecodeString(str)
//	if err != nil {
//		return err
//	}
//
//	*h = decoded
//	return nil
//}
