package eth

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"

	"github.com/kroma-network/kroma/kroma-bindings/bindings"
)

type PublicInputProof struct {
	NextBlock                   *types.Header      `json:"nextBlock"`
	NextTransactions            types.Transactions `json:"nextTransactions"`
	L2ToL1MessagePasserBalance  *big.Int           `json:"l2ToL1MessagePasserBalance"`
	L2ToL1MessagePasserCodeHash common.Hash        `json:"l2ToL1MessagePasserCodeHash"`
	MerkleProof                 []hexutil.Bytes    `json:"merkleProof"`
}

type OutputWithProofResponse struct {
	OutputResponse

	PublicInputProof *PublicInputProof `json:"publicInputProof"`
}

var ErrBlockIsEmpty = errors.New("block is empty")

type KromaOutputV0 struct {
	OutputV0

	NextBlockHash common.Hash
}

func (o *KromaOutputV0) Version() Bytes32 {
	return OutputVersionV0
}

func (o *KromaOutputV0) Marshal() []byte {
	var buf [160]byte
	version := o.Version()
	copy(buf[:32], version[:])
	copy(buf[32:], o.StateRoot[:])
	copy(buf[64:], o.MessagePasserStorageRoot[:])
	copy(buf[96:], o.BlockHash[:])
	copy(buf[128:], o.NextBlockHash[:])
	return buf[:]
}

func unmarshalKromaOutputV0(data []byte) (*KromaOutputV0, error) {
	if len(data) != 160 {
		return nil, ErrInvalidOutput
	}
	var output KromaOutputV0
	// data[:32] is the version
	copy(output.StateRoot[:], data[32:64])
	copy(output.MessagePasserStorageRoot[:], data[64:96])
	copy(output.BlockHash[:], data[96:128])
	copy(output.NextBlockHash[:], data[128:160])
	return &output, nil
}

func (o *OutputWithProofResponse) ToOutputRootProof() bindings.TypesOutputRootProof {
	return bindings.TypesOutputRootProof{
		Version:                  o.Version,
		StateRoot:                o.StateRoot,
		MessagePasserStorageRoot: o.WithdrawalStorageRoot,
		LatestBlockhash:          o.BlockRef.Hash,
		NextBlockHash:            o.NextBlockRef.Hash,
	}
}

func (o *OutputWithProofResponse) ToPublicInput() (bindings.TypesPublicInput, error) {
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
	var blobGasUsed, excessBlobGas uint64
	if p.NextBlock.BlobGasUsed != nil {
		blobGasUsed = *p.NextBlock.BlobGasUsed
	}
	if p.NextBlock.ExcessBlobGas != nil {
		excessBlobGas = *p.NextBlock.ExcessBlobGas
	}
	var parentBeaconRoot common.Hash
	if o.PublicInputProof.NextBlock.ParentBeaconRoot != nil {
		parentBeaconRoot = *o.PublicInputProof.NextBlock.ParentBeaconRoot
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
		BlobGasUsed:      blobGasUsed,
		ExcessBlobGas:    excessBlobGas,
		ParentBeaconRoot: parentBeaconRoot,
	}, nil
}

func (o *OutputWithProofResponse) ToBlockHeaderRLP() (bindings.TypesBlockHeaderRLP, error) {
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
