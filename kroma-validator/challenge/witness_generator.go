package challenge

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
)

const (
	RequestNone       RequestStatusType = "None"
	RequestProcessing RequestStatusType = "Processing"
	RequestCompleted  RequestStatusType = "Completed"
	RequestFailed     RequestStatusType = "Failed"
)

type SpecResponse struct {
	Version    string      `json:"version"`
	SP1Version string      `json:"sp1_version"`
	VKeyHash   common.Hash `json:"vkey_hash"`
}

type RequestStatusType string

type WitnessResponse struct {
	RequestStatus RequestStatusType `json:"status"`
	VKeyHash      common.Hash       `json:"vkey_hash"`
	Witness       string            `json:"witness"`
}

type WitnessGenerator interface {
	Spec(ctx context.Context) (*SpecResponse, error)
	RequestWitness(ctx context.Context, blockHash string, l1Head string) (*RequestStatusType, error)
	GetWitness(ctx context.Context, blockHash string, l1Head string) (*WitnessResponse, error)
}

func (c *Client) Spec(ctx context.Context) (*SpecResponse, error) {
	return send[SpecResponse](ctx, c, "spec", []any{})
}

func (c *Client) RequestWitness(ctx context.Context, blockHash string, l1Head string) (*RequestStatusType, error) {
	return send[RequestStatusType](ctx, c, "requestWitness", []any{blockHash, l1Head})
}

func (c *Client) GetWitness(ctx context.Context, blockHash string, l1Head string) (*WitnessResponse, error) {
	return send[WitnessResponse](ctx, c, "getWitness", []any{blockHash, l1Head})
}
