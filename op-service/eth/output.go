package eth

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"

	"github.com/kroma-network/kroma/kroma-bindings/bindings"
)

// [Kroma: START]
type PublicInputProof struct {
	NextBlock                   *types.Header      `json:"nextBlock"`
	NextTransactions            types.Transactions `json:"nextTransactions"`
	L2ToL1MessagePasserBalance  *big.Int           `json:"l2ToL1MessagePasserBalance"`
	L2ToL1MessagePasserCodeHash common.Hash        `json:"l2ToL1MessagePasserCodeHash"`
	MerkleProof                 []hexutil.Bytes    `json:"merkleProof"`
}

// [Kroma: END]

type OutputResponse struct {
	Version               Bytes32     `json:"version"`
	OutputRoot            Bytes32     `json:"outputRoot"`
	BlockRef              L2BlockRef  `json:"blockRef"`
	WithdrawalStorageRoot common.Hash `json:"withdrawalStorageRoot"`
	StateRoot             common.Hash `json:"stateRoot"`
	Status                *SyncStatus `json:"syncStatus"`
	// [Kroma: START]
	NextBlockRef     L2BlockRef        `json:"nextBlockRef"`
	PublicInputProof *PublicInputProof `json:"publicInputProof"`
	// [Kroma: END]
}

var (
	ErrBlockIsEmpty         = errors.New("block is empty")
	ErrInvalidOutput        = errors.New("invalid output")
	ErrInvalidOutputVersion = errors.New("invalid output version")

	OutputVersionV0 = Bytes32{}
)

type Output interface {
	// Version returns the version of the L2 output
	Version() Bytes32

	// Marshal a L2 output into a byte slice for hashing
	Marshal() []byte
}

type OutputV0 struct {
	StateRoot                Bytes32
	MessagePasserStorageRoot Bytes32
	BlockHash                common.Hash
	// [Kroma: START]
	NextBlockHash common.Hash
	// [Kroma: END]
}

func (o *OutputV0) Version() Bytes32 {
	return OutputVersionV0
}

func (o *OutputV0) Marshal() []byte {
	var buf [160]byte
	version := o.Version()
	copy(buf[:32], version[:])
	copy(buf[32:], o.StateRoot[:])
	copy(buf[64:], o.MessagePasserStorageRoot[:])
	copy(buf[96:], o.BlockHash[:])
	copy(buf[128:], o.NextBlockHash[:])
	return buf[:]
}

// OutputRoot returns the keccak256 hash of the marshaled L2 output
func OutputRoot(output Output) Bytes32 {
	marshaled := output.Marshal()
	return Bytes32(crypto.Keccak256Hash(marshaled))
}

func UnmarshalOutput(data []byte) (Output, error) {
	if len(data) < 32 {
		return nil, ErrInvalidOutput
	}
	var ver Bytes32
	copy(ver[:], data[:32])
	switch ver {
	case OutputVersionV0:
		return unmarshalOutputV0(data)
	default:
		return nil, ErrInvalidOutputVersion
	}
}

func unmarshalOutputV0(data []byte) (*OutputV0, error) {
	if len(data) != 160 {
		return nil, ErrInvalidOutput
	}
	var output OutputV0
	// data[:32] is the version
	copy(output.StateRoot[:], data[32:64])
	copy(output.MessagePasserStorageRoot[:], data[64:96])
	copy(output.BlockHash[:], data[96:128])
	copy(output.NextBlockHash[:], data[128:160])
	return &output, nil
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
