package eth

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"

	"github.com/kroma-network/kroma/bindings/bindings"
)

var ErrBlockIsEmpty = errors.New("block is empty")

type PublicInputProof struct {
	NextBlock                   *types.Header      `json:"nextBlock"`
	NextTransactions            types.Transactions `json:"nextTransactions"`
	L2ToL1MessagePasserBalance  *big.Int           `json:"l2ToL1MessagePasserBalance"`
	L2ToL1MessagePasserCodeHash common.Hash        `json:"l2ToL1MessagePasserCodeHash"`
	MerkleProof                 []hexutil.Bytes    `json:"merkleProof"`
}

type OutputResponse struct {
	Version               Bytes32           `json:"version"`
	OutputRoot            Bytes32           `json:"outputRoot"`
	BlockRef              L2BlockRef        `json:"blockRef"`
	NextBlockRef          L2BlockRef        `json:"nextBlockRef"`
	WithdrawalStorageRoot common.Hash       `json:"withdrawalStorageRoot"`
	StateRoot             common.Hash       `json:"stateRoot"`
	Status                *SyncStatus       `json:"syncStatus"`
	PublicInputProof      *PublicInputProof `json:"publicInputProof"`
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
	p := o.PublicInputProof
	if p.NextBlock == nil {
		return bindings.TypesPublicInput{}, ErrBlockIsEmpty
	}
	var withdrawalsRoot common.Hash
	if p.NextBlock.WithdrawalsHash != nil {
		withdrawalsRoot = *p.NextBlock.WithdrawalsHash
	}
	txHashes := make([][32]byte, len(p.NextTransactions))
	for i, tx := range p.NextTransactions {
		txHashes[i] = tx.Hash()
	}
	return bindings.TypesPublicInput{
		BlockHash:        o.NextBlockRef.Hash,
		ParentHash:       o.BlockRef.Hash,
		Timestamp:        p.NextBlock.Time,
		Number:           p.NextBlock.Number.Uint64(),
		GasLimit:         p.NextBlock.GasLimit,
		BaseFee:          p.NextBlock.BaseFee,
		TransactionsRoot: p.NextBlock.TxHash,
		StateRoot:        p.NextBlock.Root,
		WithdrawalsRoot:  withdrawalsRoot,
		TxHashes:         txHashes,
	}, nil
}

func (o *OutputResponse) ToBlockHeaderRLP() (bindings.TypesBlockHeaderRLP, error) {
	p := o.PublicInputProof
	if p.NextBlock == nil {
		return bindings.TypesBlockHeaderRLP{}, ErrBlockIsEmpty
	}
	uncleHash, err := rlp.EncodeToBytes(types.EmptyUncleHash)
	if err != nil {
		return bindings.TypesBlockHeaderRLP{}, err
	}
	coinbase, err := rlp.EncodeToBytes(p.NextBlock.Coinbase)
	if err != nil {
		return bindings.TypesBlockHeaderRLP{}, err
	}
	receiptsRoot, err := rlp.EncodeToBytes(p.NextBlock.ReceiptHash)
	if err != nil {
		return bindings.TypesBlockHeaderRLP{}, err
	}
	logsBloom, err := rlp.EncodeToBytes(p.NextBlock.Bloom)
	if err != nil {
		return bindings.TypesBlockHeaderRLP{}, err
	}
	difficulty, err := rlp.EncodeToBytes(p.NextBlock.Difficulty)
	if err != nil {
		return bindings.TypesBlockHeaderRLP{}, err
	}
	gasUsed, err := rlp.EncodeToBytes(p.NextBlock.GasUsed)
	if err != nil {
		return bindings.TypesBlockHeaderRLP{}, err
	}
	extraData, err := rlp.EncodeToBytes(p.NextBlock.Extra)
	if err != nil {
		return bindings.TypesBlockHeaderRLP{}, err
	}
	mixHash, err := rlp.EncodeToBytes(p.NextBlock.MixDigest)
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
