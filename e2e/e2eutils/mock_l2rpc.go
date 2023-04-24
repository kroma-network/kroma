package e2eutils

import (
	"context"
	"math/rand"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/kroma-network/kroma/components/node/client"
	"github.com/kroma-network/kroma/components/node/eth"
	"github.com/kroma-network/kroma/components/node/testutils"
)

type MockL2RPC struct {
	rpc                  client.RPC
	lastValidBlockNumber *hexutil.Uint64
}

func NewRPC(rpc client.RPC) *MockL2RPC {
	return &MockL2RPC{rpc: rpc}
}

func (m *MockL2RPC) SetLastValidBlockNumber(lastValidBlockNumber uint64) {
	m.lastValidBlockNumber = new(hexutil.Uint64)
	*m.lastValidBlockNumber = hexutil.Uint64(lastValidBlockNumber)
}

func (m *MockL2RPC) Close() {
	m.rpc.Close()
}

func (m *MockL2RPC) CallContext(ctx context.Context, result interface{}, method string, args ...interface{}) error {
	if method == "kroma_outputAtBlock" {
		blockNumber := args[0].(hexutil.Uint64)
		includeNextBlock := args[1].(bool)

		err := m.rpc.CallContext(ctx, &result, "kroma_outputAtBlock", blockNumber, includeNextBlock)
		if err != nil {
			return err
		}
		if m.lastValidBlockNumber != nil && *m.lastValidBlockNumber < blockNumber {
			rng := rand.New(rand.NewSource(int64(blockNumber)))

			s := result.(**eth.OutputResponse)
			(*s).OutputRoot = eth.Bytes32(testutils.RandomHash(rng))
			(*s).WithdrawalStorageRoot = testutils.RandomHash(rng)
			(*s).StateRoot = testutils.RandomHash(rng)

			return nil
		}
	}

	return m.rpc.CallContext(ctx, result, method, args...)
}

func (m *MockL2RPC) BatchCallContext(ctx context.Context, b []rpc.BatchElem) error {
	return m.rpc.BatchCallContext(ctx, b)
}

func (m *MockL2RPC) EthSubscribe(ctx context.Context, channel interface{}, args ...interface{}) (ethereum.Subscription, error) {
	return m.rpc.EthSubscribe(ctx, channel, args...)
}
