package challenge

import (
	"context"

	"github.com/ethereum-optimism/optimism/op-service/client"
	"github.com/ethereum/go-ethereum/common"
)

const (
	RequestNone       RequestStatusType = "None"
	RequestProcessing RequestStatusType = "Processing"
	RequestCompleted  RequestStatusType = "Completed"
	RequestFailed     RequestStatusType = "Failed"
)

type WitnessGenerator struct {
	rpc client.RPC
}

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

func NewWitnessGenerator(rpc client.RPC) *WitnessGenerator {
	return &WitnessGenerator{rpc}
}

func (w *WitnessGenerator) Spec(ctx context.Context) (*SpecResponse, error) {
	var output *SpecResponse
	err := w.rpc.CallContext(ctx, &output, "spec")
	return output, err
}

func (w *WitnessGenerator) RequestWitness(ctx context.Context, blockHash string, l1Head string) (*RequestStatusType, error) {
	var output *RequestStatusType
	err := w.rpc.CallContext(ctx, &output, "requestWitness", blockHash, l1Head)
	return output, err
}

func (w *WitnessGenerator) GetWitness(ctx context.Context, blockHash string, l1Head string) (*WitnessResponse, error) {
	var output *WitnessResponse
	err := w.rpc.CallContext(ctx, &output, "getWitness", blockHash, l1Head)
	return output, err
}
