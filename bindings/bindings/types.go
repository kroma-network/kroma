// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// TypesBlockHeaderRLP is an auto generated low-level Go binding around an user-defined struct.
type TypesBlockHeaderRLP struct {
	ParentHash      []byte
	UncleHash       []byte
	ReceiptsRoot    []byte
	LogsBloom       []byte
	GasUsed         []byte
	ExtraData       []byte
	MixHash         []byte
	Nonce           []byte
	WithdrawalsRoot []byte
}

// TypesChallenge is an auto generated low-level Go binding around an user-defined struct.
type TypesChallenge struct {
	Turn       *big.Int
	Asserter   common.Address
	Challenger common.Address
	Segments   [][32]byte
	SegStart   *big.Int
	SegSize    *big.Int
	TimeoutAt  *big.Int
	OutputRoot [32]byte
	Approved   bool
}

// TypesOutputRootProof is an auto generated low-level Go binding around an user-defined struct.
type TypesOutputRootProof struct {
	Version                  [32]byte
	StateRoot                [32]byte
	MessagePasserStorageRoot [32]byte
	BlockHash                [32]byte
	NextBlockHash            [32]byte
}

// TypesPublicInput is an auto generated low-level Go binding around an user-defined struct.
type TypesPublicInput struct {
	Coinbase         common.Address
	Timestamp        uint64
	Number           uint64
	Difficulty       *big.Int
	GasLimit         *big.Int
	BaseFee          *big.Int
	ChainId          *big.Int
	TransactionsRoot [32]byte
	StateRoot        [32]byte
	TxHashes         [][32]byte
}

// TypesWithdrawalTransaction is an auto generated low-level Go binding around an user-defined struct.
type TypesWithdrawalTransaction struct {
	Nonce    *big.Int
	Sender   common.Address
	Target   common.Address
	Value    *big.Int
	GasLimit *big.Int
	Data     []byte
}

// TypesCheckpointOutput is an auto generated low-level Go binding around an user-defined struct.
type TypesCheckpointOutput struct {
	Submitter     common.Address
	OutputRoot    [32]byte
	Timestamp     *big.Int
	L2BlockNumber *big.Int
}

// TypesBond is an auto generated low-level Go binding around an user-defined struct.
type TypesBond struct {
	Amount    *big.Int
	ExpiresAt *big.Int
}

