package testutils

import (
	"math/big"
	"math/rand"

	"github.com/ethereum-optimism/optimism/op-node/eth"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type MockBlockInfo struct {
	// Prefixed all fields with "Info" to avoid collisions with the interface method names.

	InfoHash       common.Hash
	InfoParentHash common.Hash
	InfoCoinbase   common.Address
	InfoRoot       common.Hash
	InfoNum        uint64
	InfoTime       uint64
	InfoMixDigest  [32]byte
	InfoBaseFee    *big.Int
	// NOTE: kroma add
	InfoTransactionsRoot common.Hash
	InfoReceiptRoot      common.Hash
	// NOTE: kroma add
	InfoWithdrawalsRoot *common.Hash
	InfoGasUsed         uint64
	// NOTE: kroma add
	InfoGasLimit uint64
	// NOTE: kroma add
	InfoBloom types.Bloom
	// NOTE: kroma add
	InfoExtra []byte
}

func (l *MockBlockInfo) Hash() common.Hash {
	return l.InfoHash
}

func (l *MockBlockInfo) ParentHash() common.Hash {
	return l.InfoParentHash
}

func (l *MockBlockInfo) Coinbase() common.Address {
	return l.InfoCoinbase
}

func (l *MockBlockInfo) Root() common.Hash {
	return l.InfoRoot
}

func (l *MockBlockInfo) NumberU64() uint64 {
	return l.InfoNum
}

func (l *MockBlockInfo) Time() uint64 {
	return l.InfoTime
}

func (l *MockBlockInfo) MixDigest() common.Hash {
	return l.InfoMixDigest
}

func (l *MockBlockInfo) BaseFee() *big.Int {
	return l.InfoBaseFee
}

func (l *MockBlockInfo) TxHash() common.Hash {
	return l.InfoTransactionsRoot
}

func (l *MockBlockInfo) ReceiptHash() common.Hash {
	return l.InfoReceiptRoot
}

func (l *MockBlockInfo) GasUsed() uint64 {
	return l.InfoGasUsed
}

func (l *MockBlockInfo) GasLimit() uint64 {
	return l.InfoGasLimit
}

func (l *MockBlockInfo) Bloom() types.Bloom {
	return l.InfoBloom
}

func (l *MockBlockInfo) Extra() []byte {
	return l.InfoExtra
}

func (l *MockBlockInfo) Header() *types.Header {
	return nil
}

func (l *MockBlockInfo) ID() eth.BlockID {
	return eth.BlockID{Hash: l.InfoHash, Number: l.InfoNum}
}

func (l *MockBlockInfo) BlockRef() eth.L1BlockRef {
	return eth.L1BlockRef{
		Hash:       l.InfoHash,
		Number:     l.InfoNum,
		ParentHash: l.InfoParentHash,
		Time:       l.InfoTime,
	}
}

func RandomBlockInfo(rng *rand.Rand) *MockBlockInfo {
	return &MockBlockInfo{
		InfoParentHash:  RandomHash(rng),
		InfoNum:         rng.Uint64(),
		InfoTime:        rng.Uint64(),
		InfoHash:        RandomHash(rng),
		InfoBaseFee:     big.NewInt(rng.Int63n(1000_000 * 1e9)), // a million GWEI
		InfoReceiptRoot: types.EmptyMPTRootHash,
		InfoRoot:        RandomHash(rng),
		InfoGasUsed:     rng.Uint64(),
	}
}

func MakeBlockInfo(fn func(l *MockBlockInfo)) func(rng *rand.Rand) *MockBlockInfo {
	return func(rng *rand.Rand) *MockBlockInfo {
		b := RandomBlockInfo(rng)
		if fn != nil {
			fn(b)
		}
		return b
	}
}

func NewMockBlockInfoWithHeader(header *types.Header) MockBlockInfo {
	return MockBlockInfo{
		InfoHash:             header.Hash(),
		InfoParentHash:       header.ParentHash,
		InfoCoinbase:         header.Coinbase,
		InfoRoot:             header.Root,
		InfoNum:              header.Number.Uint64(),
		InfoTime:             header.Time,
		InfoMixDigest:        header.MixDigest,
		InfoBaseFee:          header.BaseFee,
		InfoTransactionsRoot: header.TxHash,
		InfoReceiptRoot:      header.ReceiptHash,
		InfoWithdrawalsRoot:  header.WithdrawalsHash,
		InfoGasUsed:          header.GasUsed,
		InfoGasLimit:         header.GasLimit,
		InfoBloom:            header.Bloom,
		InfoExtra:            header.Extra,
	}
}
