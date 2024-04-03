package eth

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/misc/eip4844"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
)

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
	// BlobBaseFee returns the result of computing the blob fee from excessDataGas, or nil if the
	// block isn't a Dencun (4844 capable) block
	BlobBaseFee() *big.Int
	ReceiptHash() common.Hash
	GasUsed() uint64
	GasLimit() uint64
	ParentBeaconRoot() *common.Hash // Dencun extension

	// HeaderRLP returns the RLP of the block header as per consensus rules
	// Returns an error if the header RLP could not be written
	HeaderRLP() ([]byte, error)

	// [Kroma: START]
	TxHash() common.Hash
	Bloom() types.Bloom
	Extra() []byte

	// NOTE(chokobole): I would like to add a WithdrawalsRoot() or WithdrawalsHash()
	// method, but it is not feasible because this interface must remain compatible
	// with the types.Block constraint.
	Header() *types.Header
	// [Kroma: END]
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

// blockInfo is a conversion type of types.Block turning it into a BlockInfo
type blockInfo struct{ *types.Block }

func (b blockInfo) BlobBaseFee() *big.Int {
	ebg := b.ExcessBlobGas()
	if ebg == nil {
		return nil
	}
	return eip4844.CalcBlobFee(*ebg)
}

func (b blockInfo) HeaderRLP() ([]byte, error) {
	return rlp.EncodeToBytes(b.Header())
}

func (b blockInfo) ParentBeaconRoot() *common.Hash {
	return b.Block.BeaconRoot()
}

func BlockToInfo(b *types.Block) BlockInfo {
	return blockInfo{b}
}

var _ BlockInfo = (*blockInfo)(nil)

// headerBlockInfo is a conversion type of types.Header turning it into a
// BlockInfo.
type headerBlockInfo struct{ header *types.Header }

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

func (h headerBlockInfo) BlobBaseFee() *big.Int {
	if h.header.ExcessBlobGas == nil {
		return nil
	}
	return eip4844.CalcBlobFee(*h.header.ExcessBlobGas)
}

func (h headerBlockInfo) ReceiptHash() common.Hash {
	return h.header.ReceiptHash
}

func (h headerBlockInfo) GasUsed() uint64 {
	return h.header.GasUsed
}

func (h headerBlockInfo) GasLimit() uint64 {
	return h.header.GasLimit
}

func (h headerBlockInfo) ParentBeaconRoot() *common.Hash {
	return h.header.ParentBeaconRoot
}

func (h headerBlockInfo) HeaderRLP() ([]byte, error) {
	return rlp.EncodeToBytes(h.header)
}

// [Kroma: START]
func (h headerBlockInfo) Hash() common.Hash {
	return h.header.Hash()
}

func (h headerBlockInfo) Header() *types.Header {
	return h.header
}

func (h headerBlockInfo) TxHash() common.Hash {
	return h.header.TxHash
}

func (h headerBlockInfo) Bloom() types.Bloom {
	return h.header.Bloom
}

func (h headerBlockInfo) Extra() []byte {
	return h.header.Extra
}

// [Kroma: END]

// HeaderBlockInfo returns h as a BlockInfo implementation.
func HeaderBlockInfo(h *types.Header) BlockInfo {
	return headerBlockInfo{h}
}
