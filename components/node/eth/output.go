package eth

import (
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"

	"github.com/kroma-network/kroma/bindings/bindings"
)

var ErrBlockIsEmpty = errors.New("block is empty")

type OutputResponse struct {
	Version               Bytes32            `json:"version"`
	OutputRoot            Bytes32            `json:"outputRoot"`
	BlockRef              L2BlockRef         `json:"blockRef"`
	NextBlockRef          L2BlockRef         `json:"nextBlockRef"`
	WithdrawalStorageRoot common.Hash        `json:"withdrawalStorageRoot"`
	StateRoot             common.Hash        `json:"stateRoot"`
	Status                *SyncStatus        `json:"syncStatus"`
	NextBlock             *types.Header      `json:"nextBlock"`
	NextTransactions      types.Transactions `json:"nextTransactions"`
}

func (o *OutputResponse) ToOutputRootProof() bindings.TypesOutputRootProof {
	return bindings.TypesOutputRootProof{
		Version:                  o.Version,
		StateRoot:                o.StateRoot,
		MessagePasserStorageRoot: o.WithdrawalStorageRoot,
		BlockHash:                o.BlockRef.Hash,
		NextBlockHash:            o.NextBlockRef.Hash,
	}
}

func (o *OutputResponse) ToPublicInput() (bindings.TypesPublicInput, error) {
	if o.NextBlock == nil {
		return bindings.TypesPublicInput{}, ErrBlockIsEmpty
	}
	var withdrawalsRoot common.Hash
	if o.NextBlock.WithdrawalsHash != nil {
		withdrawalsRoot = *o.NextBlock.WithdrawalsHash
	}
	txHashes := make([][32]byte, len(o.NextTransactions))
	for i, tx := range o.NextTransactions {
		txHashes[i] = tx.Hash()
	}
	return bindings.TypesPublicInput{
		BlockHash:        o.NextBlockRef.Hash,
		ParentHash:       o.BlockRef.Hash,
		Timestamp:        o.NextBlock.Time,
		Number:           o.NextBlock.Number.Uint64(),
		GasLimit:         o.NextBlock.GasLimit,
		BaseFee:          o.NextBlock.BaseFee,
		TransactionsRoot: o.NextBlock.TxHash,
		StateRoot:        o.NextBlock.Root,
		WithdrawalsRoot:  withdrawalsRoot,
		TxHashes:         txHashes,
	}, nil
}

func (o *OutputResponse) ToBlockHeaderRLP() (bindings.TypesBlockHeaderRLP, error) {
	if o.NextBlock == nil {
		return bindings.TypesBlockHeaderRLP{}, ErrBlockIsEmpty
	}
	uncleHash, err := rlp.EncodeToBytes(types.EmptyUncleHash)
	if err != nil {
		return bindings.TypesBlockHeaderRLP{}, err
	}
	coinbase, err := rlp.EncodeToBytes(o.NextBlock.Coinbase)
	if err != nil {
		return bindings.TypesBlockHeaderRLP{}, err
	}
	receiptsRoot, err := rlp.EncodeToBytes(o.NextBlock.ReceiptHash)
	if err != nil {
		return bindings.TypesBlockHeaderRLP{}, err
	}
	logsBloom, err := rlp.EncodeToBytes(o.NextBlock.Bloom)
	if err != nil {
		return bindings.TypesBlockHeaderRLP{}, err
	}
	difficulty, err := rlp.EncodeToBytes(o.NextBlock.Difficulty)
	if err != nil {
		return bindings.TypesBlockHeaderRLP{}, err
	}
	gasUsed, err := rlp.EncodeToBytes(o.NextBlock.GasUsed)
	if err != nil {
		return bindings.TypesBlockHeaderRLP{}, err
	}
	extraData, err := rlp.EncodeToBytes(o.NextBlock.Extra)
	if err != nil {
		return bindings.TypesBlockHeaderRLP{}, err
	}
	mixHash, err := rlp.EncodeToBytes(o.NextBlock.MixDigest)
	if err != nil {
		return bindings.TypesBlockHeaderRLP{}, err
	}
	nonce, err := rlp.EncodeToBytes(types.BlockNonce{})
	if err != nil {
		return bindings.TypesBlockHeaderRLP{}, err
	}

	return bindings.TypesBlockHeaderRLP{
		UncleHash:    uncleHash,
		Coinbase:     coinbase,
		ReceiptsRoot: receiptsRoot,
		LogsBloom:    logsBloom,
		Difficulty:   difficulty,
		GasUsed:      gasUsed,
		ExtraData:    extraData,
		MixHash:      mixHash,
		Nonce:        nonce,
	}, nil
}
