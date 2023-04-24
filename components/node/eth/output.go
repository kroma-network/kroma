package eth

import (
	"errors"
	"math/big"

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

func (o *OutputResponse) ToPublicInput(ChainId *big.Int) (bindings.TypesPublicInput, error) {
	if o.NextBlock == nil {
		return bindings.TypesPublicInput{}, ErrBlockIsEmpty
	}
	txHashes := make([][32]byte, len(o.NextTransactions))
	for i, tx := range o.NextTransactions {
		txHashes[i] = tx.Hash()
	}
	return bindings.TypesPublicInput{
		Coinbase:         o.NextBlock.Coinbase,
		Timestamp:        o.NextBlock.Time,
		Number:           o.NextBlock.Number.Uint64(),
		Difficulty:       common.Big0,
		GasLimit:         new(big.Int).SetUint64(o.NextBlock.GasLimit),
		BaseFee:          o.NextBlock.BaseFee,
		ChainId:          ChainId,
		TransactionsRoot: o.NextBlock.TxHash,
		StateRoot:        o.NextBlock.Root,
		TxHashes:         txHashes,
	}, nil
}

func (o *OutputResponse) ToBlockHeaderRLP() (bindings.TypesBlockHeaderRLP, error) {
	if o.NextBlock == nil {
		return bindings.TypesBlockHeaderRLP{}, ErrBlockIsEmpty
	}
	parentHash, err := rlp.EncodeToBytes(o.NextBlock.ParentHash)
	if err != nil {
		return bindings.TypesBlockHeaderRLP{}, err
	}
	uncleHash, err := rlp.EncodeToBytes(types.EmptyUncleHash)
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
	var withdrawalsRoot []byte
	if o.NextBlock.WithdrawalsHash != nil {
		withdrawalsRoot, err = rlp.EncodeToBytes(*o.NextBlock.WithdrawalsHash)
		if err != nil {
			return bindings.TypesBlockHeaderRLP{}, err
		}
	}

	return bindings.TypesBlockHeaderRLP{
		ParentHash:      parentHash,
		UncleHash:       uncleHash,
		ReceiptsRoot:    receiptsRoot,
		LogsBloom:       logsBloom,
		GasUsed:         gasUsed,
		ExtraData:       extraData,
		MixHash:         mixHash,
		Nonce:           nonce,
		WithdrawalsRoot: withdrawalsRoot,
	}, nil
}
