// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// TypesBlockHeaderRLP is an auto generated low-level Go binding around an user-defined struct.
type TypesBlockHeaderRLP struct {
	UncleHash    []byte
	Coinbase     []byte
	ReceiptsRoot []byte
	LogsBloom    []byte
	Difficulty   []byte
	GasUsed      []byte
	ExtraData    []byte
	MixHash      []byte
	Nonce        []byte
}

// TypesChallenge is an auto generated low-level Go binding around an user-defined struct.
type TypesChallenge struct {
	Turn       uint8
	TimeoutAt  uint64
	Asserter   common.Address
	Challenger common.Address
	Segments   [][32]byte
	SegSize    *big.Int
	SegStart   *big.Int
	L1Head     [32]byte
}

// TypesOutputRootProof is an auto generated low-level Go binding around an user-defined struct.
type TypesOutputRootProof struct {
	Version                  [32]byte
	StateRoot                [32]byte
	MessagePasserStorageRoot [32]byte
	LatestBlockhash          [32]byte
	NextBlockHash            [32]byte
}

// TypesPublicInput is an auto generated low-level Go binding around an user-defined struct.
type TypesPublicInput struct {
	BlockHash        [32]byte
	ParentHash       [32]byte
	Timestamp        uint64
	Number           uint64
	GasLimit         uint64
	BaseFee          *big.Int
	TransactionsRoot [32]byte
	StateRoot        [32]byte
	WithdrawalsRoot  [32]byte
	TxHashes         [][32]byte
	BlobGasUsed      uint64
	ExcessBlobGas    uint64
	ParentBeaconRoot [32]byte
}

// TypesPublicInputProof is an auto generated low-level Go binding around an user-defined struct.
type TypesPublicInputProof struct {
	SrcOutputRootProof          TypesOutputRootProof
	DstOutputRootProof          TypesOutputRootProof
	PublicInput                 TypesPublicInput
	Rlps                        TypesBlockHeaderRLP
	L2ToL1MessagePasserBalance  [32]byte
	L2ToL1MessagePasserCodeHash [32]byte
	MerkleProof                 [][]byte
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

// TypesZkEvmProof is an auto generated low-level Go binding around an user-defined struct.
type TypesZkEvmProof struct {
	PublicInputProof TypesPublicInputProof
	Proof            []*big.Int
	Pair             []*big.Int
}

// TypesZkVmProof is an auto generated low-level Go binding around an user-defined struct.
type TypesZkVmProof struct {
	ZkVmProgramVKey [32]byte
	PublicValues    []byte
	ProofBytes      []byte
}
