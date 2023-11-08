package e2eutils

import (
	"context"
	"math/rand"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/ethereum-optimism/optimism/op-node/client"
	"github.com/ethereum-optimism/optimism/op-node/eth"
	"github.com/ethereum-optimism/optimism/op-node/testutils"
	"github.com/ethereum-optimism/optimism/op-e2e/testdata"
)

type MaliciousL2RPC struct {
	rpc client.RPC
	// targetBlockNumber is the block number for challenge
	targetBlockNumber *hexutil.Uint64
}

func NewMaliciousL2RPC(rpc client.RPC) *MaliciousL2RPC {
	return &MaliciousL2RPC{rpc: rpc}
}

// SetTargetBlockNumber sets the first invalid block number for mocking malicious L2 RPC.
// After the m.targetBlockNumber, random output root will be returned for `kroma_outputAtBlock` CallContext
func (m *MaliciousL2RPC) SetTargetBlockNumber(lastValidBlockNumber uint64) {
	m.targetBlockNumber = new(hexutil.Uint64)
	*m.targetBlockNumber = hexutil.Uint64(lastValidBlockNumber)
}

func (m *MaliciousL2RPC) Close() {
	m.rpc.Close()
}

func (m *MaliciousL2RPC) CallContext(ctx context.Context, result interface{}, method string, args ...interface{}) error {
	if method == "kroma_outputAtBlock" || method == "kroma_outputWithProofAtBlock" {
		blockNumber := args[0].(hexutil.Uint64)

		err := m.rpc.CallContext(ctx, &result, method, blockNumber)
		if err != nil {
			return err
		}
		if m.targetBlockNumber != nil && *m.targetBlockNumber-1 == blockNumber {
			return testdata.SetPrevOutputResponse(result.(**eth.OutputResponse))
		} else if m.targetBlockNumber != nil && *m.targetBlockNumber <= blockNumber {
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

func (m *MaliciousL2RPC) BatchCallContext(ctx context.Context, b []rpc.BatchElem) error {
	return m.rpc.BatchCallContext(ctx, b)
}

func (m *MaliciousL2RPC) EthSubscribe(ctx context.Context, channel interface{}, args ...interface{}) (ethereum.Subscription, error) {
	return m.rpc.EthSubscribe(ctx, channel, args...)
}

type HonestL2RPC struct {
	rpc client.RPC
	// targetBlockNumber is the block number for challenge
	targetBlockNumber *hexutil.Uint64
}

func NewHonestL2RPC(rpc client.RPC) *HonestL2RPC {
	return &HonestL2RPC{rpc: rpc}
}

// SetTargetBlockNumber sets the target block number for challenge.
// At the m.targetBlockNumber, mocked output root will be returned for `kroma_outputAtBlock` CallContext
func (m *HonestL2RPC) SetTargetBlockNumber(lastValidBlockNumber uint64) {
	m.targetBlockNumber = new(hexutil.Uint64)
	*m.targetBlockNumber = hexutil.Uint64(lastValidBlockNumber)
}

func (m *HonestL2RPC) Close() {
	m.rpc.Close()
}

func (m *HonestL2RPC) CallContext(ctx context.Context, result interface{}, method string, args ...interface{}) error {
	if method == "kroma_outputAtBlock" || method == "kroma_outputWithProofAtBlock" {
		blockNumber := args[0].(hexutil.Uint64)

		err := m.rpc.CallContext(ctx, &result, method, blockNumber)
		if err != nil {
			return err
		}
		if m.targetBlockNumber != nil && *m.targetBlockNumber-1 == blockNumber {
			return testdata.SetPrevOutputResponse(result.(**eth.OutputResponse))
		} else if m.targetBlockNumber != nil && *m.targetBlockNumber == blockNumber {
			return testdata.SetTargetOutputResponse(result.(**eth.OutputResponse))
		}
	}

	return m.rpc.CallContext(ctx, result, method, args...)
}

func (m *HonestL2RPC) BatchCallContext(ctx context.Context, b []rpc.BatchElem) error {
	return m.rpc.BatchCallContext(ctx, b)
}

func (m *HonestL2RPC) EthSubscribe(ctx context.Context, channel interface{}, args ...interface{}) (ethereum.Subscription, error) {
	return m.rpc.EthSubscribe(ctx, channel, args...)
}
