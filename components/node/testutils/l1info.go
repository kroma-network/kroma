package testutils

import (
	"math/big"
	"math/rand"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/kroma-network/kroma/components/node/eth"
)

type MockBlockInfo struct {
	// Prefixed all fields with "Info" to avoid collisions with the interface method names.

	InfoHash             common.Hash
	InfoParentHash       common.Hash
	InfoCoinbase         common.Address
	InfoRoot             common.Hash
	InfoNum              uint64
	InfoTime             uint64
	InfoMixDigest        [32]byte
	InfoBaseFee          *big.Int
	InfoTransactionsRoot common.Hash
	InfoReceiptRoot      common.Hash
	InfoWithdrawalsRoot  *common.Hash
	InfoGasUsed          uint64
	InfoGasLimit         uint64
	InfoBloom            types.Bloom
	InfoExtra            []byte
}

func (m *MockBlockInfo) Hash() common.Hash {
	return m.InfoHash
}

func (m *MockBlockInfo) ParentHash() common.Hash {
	return m.InfoParentHash
}

func (m *MockBlockInfo) Coinbase() common.Address {
	return m.InfoCoinbase
}

func (m *MockBlockInfo) Root() common.Hash {
	return m.InfoRoot
}

func (m *MockBlockInfo) NumberU64() uint64 {
	return m.InfoNum
}

func (m *MockBlockInfo) Time() uint64 {
	return m.InfoTime
}

func (m *MockBlockInfo) MixDigest() common.Hash {
	return m.InfoMixDigest
}

func (m *MockBlockInfo) BaseFee() *big.Int {
	return m.InfoBaseFee
}

func (m *MockBlockInfo) TxHash() common.Hash {
	return m.InfoTransactionsRoot
}

func (m *MockBlockInfo) ReceiptHash() common.Hash {
	return m.InfoReceiptRoot
}

func (m *MockBlockInfo) GasUsed() uint64 {
	return m.InfoGasUsed
}

func (m *MockBlockInfo) GasLimit() uint64 {
	return m.InfoGasLimit
}

func (m *MockBlockInfo) Bloom() types.Bloom {
	return m.InfoBloom
}

func (m *MockBlockInfo) Extra() []byte {
	return m.InfoExtra
}

func (m *MockBlockInfo) Header() *types.Header {
	return nil
}

func (m *MockBlockInfo) ID() eth.BlockID {
	return eth.BlockID{Hash: m.InfoHash, Number: m.InfoNum}
}

func (m *MockBlockInfo) BlockRef() eth.L1BlockRef {
	return eth.L1BlockRef{
		Hash:       m.InfoHash,
		Number:     m.InfoNum,
		ParentHash: m.InfoParentHash,
		Time:       m.InfoTime,
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

func MakeBlockInfo(fn func(m *MockBlockInfo)) func(rng *rand.Rand) *MockBlockInfo {
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
