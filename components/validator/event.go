package validator

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

const (
	KeyEventOutputSubmitted  = "OutputSubmitted"
	KeyEventChallengeCreated = "ChallengeCreated"
)

type ChallengeCreatedEvent struct {
	OutputIndex *big.Int
	Asserter    common.Address
	Challenger  common.Address
}

func NewChallengeCreatedEvent(log types.Log) ChallengeCreatedEvent {
	return ChallengeCreatedEvent{
		OutputIndex: new(big.Int).SetBytes(log.Topics[1][:]),
		Asserter:    common.BytesToAddress(log.Topics[2][:]),
		Challenger:  common.BytesToAddress(log.Topics[3][:]),
	}
}

type OutputSubmittedEvent struct {
	ExpectedOutputRoot string
	OutputIndex        *big.Int
	L2BlockNumber      *big.Int
}

func NewOutputSubmittedEvent(log types.Log) OutputSubmittedEvent {
	return OutputSubmittedEvent{
		ExpectedOutputRoot: log.Topics[1].Hex(),
		OutputIndex:        new(big.Int).SetBytes(log.Topics[2][:]),
		L2BlockNumber:      new(big.Int).SetBytes(log.Topics[3][:]),
	}
}
