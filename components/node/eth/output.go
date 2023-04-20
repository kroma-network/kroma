package eth

import (
	"github.com/ethereum/go-ethereum/common"
)

type OutputResponse struct {
	Version               Bytes32     `json:"version"`
	OutputRoot            Bytes32     `json:"outputRoot"`
	BlockRef              L2BlockRef  `json:"blockRef"`
	NextBlockRef          L2BlockRef  `json:"nextBlockRef"`
	WithdrawalStorageRoot common.Hash `json:"withdrawalStorageRoot"`
	StateRoot             common.Hash `json:"stateRoot"`
	Status                *SyncStatus `json:"syncStatus"`
}
