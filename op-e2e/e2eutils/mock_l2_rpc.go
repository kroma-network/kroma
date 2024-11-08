package e2eutils

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/ethereum-optimism/optimism/op-service/client"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum-optimism/optimism/op-service/testutils"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/kroma-network/kroma/op-e2e/testdata"
)

type MaliciousL2RPC struct {
	rpc client.RPC
	// targetBlockNumber is the block number for challenge
	targetBlockNumber *hexutil.Uint64
	proofType         testdata.ProofType
}

func NewMaliciousL2RPC(rpc client.RPC, proofType testdata.ProofType) (*MaliciousL2RPC, error) {
	if !testdata.ValidProofType(proofType) {
		return nil, fmt.Errorf("unexpected challenge proof type: %s", proofType)
	}
	return &MaliciousL2RPC{rpc: rpc, proofType: proofType}, nil
}

// SetTargetBlockNumber sets the first invalid block number for mocking malicious L2 RPC.
// After the m.targetBlockNumber, random output root will be returned for `optimism_outputAtBlock` CallContext
func (m *MaliciousL2RPC) SetTargetBlockNumber(lastValidBlockNumber uint64) {
	m.targetBlockNumber = new(hexutil.Uint64)
	*m.targetBlockNumber = hexutil.Uint64(lastValidBlockNumber)
}

func (m *MaliciousL2RPC) Close() {
	m.rpc.Close()
}

func (m *MaliciousL2RPC) CallContext(ctx context.Context, result interface{}, method string, args ...interface{}) error {
	if method == "optimism_outputAtBlock" || method == "kroma_outputWithProofAtBlock" {
		err := m.rpc.CallContext(ctx, result, method, args...)
		if err != nil {
			return err
		}

		blockNumber := args[0].(hexutil.Uint64)
		if m.targetBlockNumber != nil && *m.targetBlockNumber-1 == blockNumber {
			if method == "optimism_outputAtBlock" {
				if o, ok := result.(**eth.OutputResponse); ok {
					return testdata.SetPrevOutputResponse(*o, m.proofType)
				} else {
					return fmt.Errorf("invalid type for result: %T (method %s)", result, method)
				}
			}

			if o, ok := result.(**eth.OutputWithProofResponse); ok {
				return testdata.SetPrevOutputWithProofResponse(*o)
			} else {
				return fmt.Errorf("invalid type for result: %T (method %s)", result, method)
			}
		} else if m.targetBlockNumber != nil && *m.targetBlockNumber <= blockNumber {
			rng := rand.New(rand.NewSource(int64(blockNumber)))

			if o, ok := result.(**eth.OutputResponse); ok {
				(**o).OutputRoot = eth.Bytes32(testutils.RandomHash(rng))
				(**o).WithdrawalStorageRoot = testutils.RandomHash(rng)
				(**o).StateRoot = testutils.RandomHash(rng)

				return nil
			} else {
				return fmt.Errorf("invalid type for result: %T (method %s)", result, method)
			}
		}

		return nil
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
	proofType         testdata.ProofType
}

func NewHonestL2RPC(rpc client.RPC, proofType testdata.ProofType) (*HonestL2RPC, error) {
	if !testdata.ValidProofType(proofType) {
		return nil, fmt.Errorf("unexpected challenge proof type: %s", proofType)
	}
	return &HonestL2RPC{rpc: rpc, proofType: proofType}, nil
}

// SetTargetBlockNumber sets the target block number for challenge.
// At the m.targetBlockNumber, mocked output root will be returned for `optimism_outputAtBlock` CallContext
func (m *HonestL2RPC) SetTargetBlockNumber(lastValidBlockNumber uint64) {
	m.targetBlockNumber = new(hexutil.Uint64)
	*m.targetBlockNumber = hexutil.Uint64(lastValidBlockNumber)
}

func (m *HonestL2RPC) Close() {
	m.rpc.Close()
}

func (m *HonestL2RPC) CallContext(ctx context.Context, result interface{}, method string, args ...interface{}) error {
	if method == "optimism_outputAtBlock" || method == "kroma_outputWithProofAtBlock" {
		err := m.rpc.CallContext(ctx, result, method, args...)
		if err != nil {
			return err
		}

		blockNumber := args[0].(hexutil.Uint64)
		if m.targetBlockNumber != nil && *m.targetBlockNumber-1 == blockNumber {
			if method == "optimism_outputAtBlock" {
				if o, ok := result.(**eth.OutputResponse); ok {
					return testdata.SetPrevOutputResponse(*o, m.proofType)
				} else {
					return fmt.Errorf("invalid type for result: %T (method %s)", result, method)
				}
			}

			if o, ok := result.(**eth.OutputWithProofResponse); ok {
				return testdata.SetPrevOutputWithProofResponse(*o)
			} else {
				return fmt.Errorf("invalid type for result: %T (method %s)", result, method)
			}
		} else if m.targetBlockNumber != nil && *m.targetBlockNumber == blockNumber {
			if method == "optimism_outputAtBlock" {
				if o, ok := result.(**eth.OutputResponse); ok {
					return testdata.SetTargetOutputResponse(*o, m.proofType)
				} else {
					return fmt.Errorf("invalid type for result: %T (method %s)", result, method)
				}
			}

			if o, ok := result.(**eth.OutputWithProofResponse); ok {
				return testdata.SetTargetOutputWithProofResponse(*o)
			} else {
				return fmt.Errorf("invalid type for result: %T (method %s)", result, method)
			}
		}

		return nil
	}

	return m.rpc.CallContext(ctx, result, method, args...)
}

func (m *HonestL2RPC) BatchCallContext(ctx context.Context, b []rpc.BatchElem) error {
	return m.rpc.BatchCallContext(ctx, b)
}

func (m *HonestL2RPC) EthSubscribe(ctx context.Context, channel interface{}, args ...interface{}) (ethereum.Subscription, error) {
	return m.rpc.EthSubscribe(ctx, channel, args...)
}
