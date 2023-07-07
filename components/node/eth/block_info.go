package eth

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

var _ BlockInfo = (&types.Block{})

type BlockInfo interface {
	Hash() common.Hash
	ParentHash() common.Hash
	Coinbase() common.Address
	Root() common.Hash // state-root
	NumberU64() uint64
	Time() uint64
	// MixDigest field, reused for randomness after The Merge (Bellatrix hardfork)
	MixDigest() common.Hash
	BaseFee() *big.Int
	TxHash() common.Hash
	ReceiptHash() common.Hash
	GasUsed() uint64
	GasLimit() uint64
	Bloom() types.Bloom
	Extra() []byte

	// NOTE(chokobole): I would like to add a WithdrawalsRoot() or WithdrawalsHash()
	// method, but it is not feasible because this interface must remain compatible
	// with the types.Block constraint.
	Header() *types.Header
}

func InfoToL1BlockRef(info BlockInfo) L1BlockRef {
	return L1BlockRef{
		Hash:       info.Hash(),
		Number:     info.NumberU64(),
		ParentHash: info.ParentHash(),
		Time:       info.Time(),
	}
}

type NumberAndHash interface {
	Hash() common.Hash
	NumberU64() uint64
}

func ToBlockID(b NumberAndHash) BlockID {
	return BlockID{
		Hash:   b.Hash(),
		Number: b.NumberU64(),
	}
}

// headerBlockInfo is a conversion type of types.Header turning it into a
// BlockInfo.
type headerBlockInfo struct{ header *types.Header }

func (h headerBlockInfo) Hash() common.Hash {
	return h.header.Hash()
}

func (h headerBlockInfo) ParentHash() common.Hash {
	return h.header.ParentHash
}

func (h headerBlockInfo) Coinbase() common.Address {
	return h.header.Coinbase
}

func (h headerBlockInfo) Root() common.Hash {
	return h.header.Root
}

func (h headerBlockInfo) NumberU64() uint64 {
	return h.header.Number.Uint64()
}

func (h headerBlockInfo) Time() uint64 {
	return h.header.Time
}

func (h headerBlockInfo) MixDigest() common.Hash {
	return h.header.MixDigest
}

func (h headerBlockInfo) BaseFee() *big.Int {
	return h.header.BaseFee
}

func (h headerBlockInfo) TxHash() common.Hash {
	return h.header.TxHash
}

func (h headerBlockInfo) ReceiptHash() common.Hash {
	return h.header.ReceiptHash
}

func (h headerBlockInfo) WithdrawalsHash() *common.Hash {
	return h.header.WithdrawalsHash
}

func (h headerBlockInfo) GasUsed() uint64 {
	return h.header.GasUsed
}

func (h headerBlockInfo) GasLimit() uint64 {
	return h.header.GasLimit
}

func (h headerBlockInfo) Bloom() types.Bloom {
	return h.header.Bloom
}

func (h headerBlockInfo) Extra() []byte {
	return h.header.Extra
}

func (h headerBlockInfo) Header() *types.Header {
	return h.header
}

// HeaderBlockInfo returns h as a BlockInfo implementation.
func HeaderBlockInfo(h *types.Header) BlockInfo {
	return headerBlockInfo{h}
}
