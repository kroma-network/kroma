package challenge

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
)

type RequestStatusType string

const (
	Requested  RequestStatusType = "Requested"
	Processing RequestStatusType = "Processing"
	Completed  RequestStatusType = "Completed"
)

type WitnessGenerator interface {
	RequestWitness(ctx context.Context, blockHash common.Hash, l1Head common.Hash) (RequestStatusType, error)
}
