package test

import (
	"math/rand"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/trie"

	"github.com/kroma-network/kroma/components/node/eth"
	"github.com/kroma-network/kroma/components/node/rollup/derive"
	"github.com/kroma-network/kroma/components/node/testutils"
)

// RandomL2Block returns a random block whose first transaction is a random
// L1 Info Deposit transaction.
func RandomL2Block(rng *rand.Rand, txCount int) (*types.Block, []*types.Receipt) {
	l1Block := types.NewBlock(testutils.RandomHeader(rng),
		nil, nil, nil, trie.NewStackTrie(nil))
	l1InfoTx, err := derive.L1InfoDeposit(0, l1Block, eth.SystemConfig{})
	if err != nil {
		panic("L1InfoDeposit: " + err.Error())
	}
	return testutils.RandomBlockPrependTxs(rng, txCount, types.NewTx(l1InfoTx))
}
