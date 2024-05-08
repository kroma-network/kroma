package validator

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/kroma-network/kroma/kroma-bindings/bindings"
)

const (
	KeyEventOutputSubmitted  = "OutputSubmitted"
	KeyEventChallengeCreated = "ChallengeCreated"
	KeyEventReadyToProve     = "ReadyToProve"
)

func NewChallengeCreatedEvent(log types.Log) *bindings.ColosseumChallengeCreated {
	return &bindings.ColosseumChallengeCreated{
		OutputIndex: new(big.Int).SetBytes(log.Topics[1][:]),
		Asserter:    common.BytesToAddress(log.Topics[2][:]),
		Challenger:  common.BytesToAddress(log.Topics[3][:]),
	}
}

func NewOutputSubmittedEvent(log types.Log) *bindings.L2OutputOracleOutputSubmitted {
	var outputRoot [32]byte
	copy(outputRoot[:], log.Topics[1][:])

	return &bindings.L2OutputOracleOutputSubmitted{
		OutputRoot:    outputRoot,
		L2OutputIndex: new(big.Int).SetBytes(log.Topics[2][:]),
		L2BlockNumber: new(big.Int).SetBytes(log.Topics[3][:]),
	}
}
