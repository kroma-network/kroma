package actions

import (
	"context"
	"errors"
	"io"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	gnode "github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/stretchr/testify/require"

	"github.com/kroma-network/kroma/components/node/client"
	"github.com/kroma-network/kroma/components/node/eth"
	"github.com/kroma-network/kroma/components/node/node"
	"github.com/kroma-network/kroma/components/node/rollup"
	"github.com/kroma-network/kroma/components/node/rollup/derive"
	"github.com/kroma-network/kroma/components/node/rollup/driver"
	"github.com/kroma-network/kroma/components/node/sources"
	"github.com/kroma-network/kroma/components/node/testutils"
)

// L2Syncer is an actor that functions like a rollup node,
// without the full P2P/API/Node stack, but just the derivation state, and simplified driver.
type L2Syncer struct {
	log log.Logger

	eng interface {
		derive.Engine
		L2BlockRefByNumber(ctx context.Context, num uint64) (eth.L2BlockRef, error)
	}

	// L2 rollup
	derivation *derive.DerivationPipeline

	l1      derive.L1Fetcher
	l1State *driver.L1State

	l2PipelineIdle bool
	l2Building     bool

	rollupCfg *rollup.Config

	rpc *rpc.Server

	failRPC error // mock error
}

type L2API interface {
	derive.Engine
	L2BlockRefByNumber(ctx context.Context, num uint64) (eth.L2BlockRef, error)
	InfoByHash(ctx context.Context, hash common.Hash) (eth.BlockInfo, error)
	InfoAndTxsByHash(ctx context.Context, hash common.Hash) (eth.BlockInfo, types.Transactions, error)
	// GetProof returns a proof of the account, it may return a nil result without error if the address was not found.
	GetProof(ctx context.Context, address common.Address, storage []common.Hash, blockTag string) (*eth.AccountResult, error)
}

func NewL2Syncer(t Testing, log log.Logger, l1 derive.L1Fetcher, eng L2API, cfg *rollup.Config) *L2Syncer {
	metrics := &testutils.TestDerivationMetrics{}
	pipeline := derive.NewDerivationPipeline(log, cfg, l1, eng, metrics)
	pipeline.Reset()

	rollupNode := &L2Syncer{
		log:            log,
		eng:            eng,
		derivation:     pipeline,
		l1:             l1,
		l1State:        driver.NewL1State(log, metrics),
		l2PipelineIdle: true,
		l2Building:     false,
		rollupCfg:      cfg,
		rpc:            rpc.NewServer(),
	}
	t.Cleanup(rollupNode.rpc.Stop)

	// setup RPC server for rollup node, hooked to the actor as backend
	m := &testutils.TestRPCMetrics{}
	backend := &l2SyncerBackend{syncer: rollupNode}
	apis := []rpc.API{
		{
			Namespace:     "kroma",
			Service:       node.NewNodeAPI(cfg, eng, backend, log, m),
			Public:        true,
			Authenticated: false,
		},
		{
			Namespace:     "admin",
			Version:       "",
			Service:       node.NewAdminAPI(backend, m),
			Public:        true, // TODO: this field is deprecated. Do we even need this anymore?
			Authenticated: false,
		},
	}
	require.NoError(t, gnode.RegisterApis(apis, nil, rollupNode.rpc), "failed to set up APIs")
	return rollupNode
}

type l2SyncerBackend struct {
	syncer *L2Syncer
}

func (s *l2SyncerBackend) BlockRefsWithStatus(ctx context.Context, num uint64) (eth.L2BlockRef, eth.L2BlockRef, *eth.SyncStatus, error) {
	ref, err := s.syncer.eng.L2BlockRefByNumber(ctx, num)
	nextRef := eth.L2BlockRef{}
	if err == nil && s.syncer.rollupCfg.IsBlueBlock(num) {
		nextRef, err = s.syncer.eng.L2BlockRefByNumber(ctx, num+1)
	}
	return ref, nextRef, s.syncer.SyncStatus(), err
}

func (s *l2SyncerBackend) SyncStatus(ctx context.Context) (*eth.SyncStatus, error) {
	return s.syncer.SyncStatus(), nil
}

func (s *l2SyncerBackend) ResetDerivationPipeline(ctx context.Context) error {
	s.syncer.derivation.Reset()
	return nil
}

func (s *l2SyncerBackend) StartProposer(ctx context.Context, blockHash common.Hash) error {
	return nil
}

func (s *l2SyncerBackend) StopProposer(ctx context.Context) (common.Hash, error) {
	return common.Hash{}, errors.New("stopping the L2Syncer proposer is not supported")
}

func (s *L2Syncer) L2Finalized() eth.L2BlockRef {
	return s.derivation.Finalized()
}

func (s *L2Syncer) L2Safe() eth.L2BlockRef {
	return s.derivation.SafeL2Head()
}

func (s *L2Syncer) L2Unsafe() eth.L2BlockRef {
	return s.derivation.UnsafeL2Head()
}

func (s *L2Syncer) SyncStatus() *eth.SyncStatus {
	return &eth.SyncStatus{
		CurrentL1:          s.derivation.Origin(),
		CurrentL1Finalized: s.derivation.FinalizedL1(),
		HeadL1:             s.l1State.L1Head(),
		SafeL1:             s.l1State.L1Safe(),
		FinalizedL1:        s.l1State.L1Finalized(),
		UnsafeL2:           s.L2Unsafe(),
		SafeL2:             s.L2Safe(),
		FinalizedL2:        s.L2Finalized(),
		UnsafeL2SyncTarget: s.derivation.UnsafeL2SyncTarget(),
	}
}

func (s *L2Syncer) RollupClient() *sources.RollupClient {
	return sources.NewRollupClient(s.RPCClient())
}

func (s *L2Syncer) RPCClient() client.RPC {
	cl := rpc.DialInProc(s.rpc)
	return testutils.RPCErrFaker{
		RPC: client.NewBaseRPCClient(cl),
		ErrFn: func() error {
			err := s.failRPC
			s.failRPC = nil // reset back, only error once.
			return err
		},
	}
}

// ActRPCFail makes the next L2 RPC request fail
func (s *L2Syncer) ActRPCFail(t Testing) {
	if s.failRPC != nil { // already set to fail?
		t.InvalidAction("already set a mock rpc error")
		return
	}
	s.failRPC = errors.New("mock RPC error")
}

func (s *L2Syncer) ActL1HeadSignal(t Testing) {
	head, err := s.l1.L1BlockRefByLabel(t.Ctx(), eth.Unsafe)
	require.NoError(t, err)
	s.l1State.HandleNewL1HeadBlock(head)
}

func (s *L2Syncer) ActL1SafeSignal(t Testing) {
	safe, err := s.l1.L1BlockRefByLabel(t.Ctx(), eth.Safe)
	require.NoError(t, err)
	s.l1State.HandleNewL1SafeBlock(safe)
}

func (s *L2Syncer) ActL1FinalizedSignal(t Testing) {
	finalized, err := s.l1.L1BlockRefByLabel(t.Ctx(), eth.Finalized)
	require.NoError(t, err)
	s.l1State.HandleNewL1FinalizedBlock(finalized)
	s.derivation.Finalize(finalized)
}

// ActL2PipelineStep runs one iteration of the L2 derivation pipeline
func (s *L2Syncer) ActL2PipelineStep(t Testing) {
	if s.l2Building {
		t.InvalidAction("cannot derive new data while building L2 block")
		return
	}

	s.l2PipelineIdle = false
	err := s.derivation.Step(t.Ctx())
	if err == io.EOF {
		s.l2PipelineIdle = true
		return
	} else if err != nil && errors.Is(err, derive.NotEnoughData) {
		return
	} else if err != nil && errors.Is(err, derive.ErrReset) {
		s.log.Warn("Derivation pipeline is reset", "err", err)
		s.derivation.Reset()
		return
	} else if err != nil && errors.Is(err, derive.ErrTemporary) {
		s.log.Warn("Derivation process temporary error", "err", err)
		return
	} else if err != nil && errors.Is(err, derive.ErrCritical) {
		t.Fatalf("derivation failed critically: %v", err)
	} else {
		return
	}
}

func (s *L2Syncer) ActL2PipelineFull(t Testing) {
	s.l2PipelineIdle = false
	for !s.l2PipelineIdle {
		s.ActL2PipelineStep(t)
	}
}

// ActL2UnsafeGossipReceive creates an action that can receive an unsafe execution payload, like gossipsub
func (s *L2Syncer) ActL2UnsafeGossipReceive(payload *eth.ExecutionPayload) Action {
	return func(t Testing) {
		s.derivation.AddUnsafePayload(payload)
	}
}
