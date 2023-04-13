package actions

import (
	"context"
	"errors"

	"github.com/ethereum/go-ethereum/log"
	"github.com/stretchr/testify/require"

	"github.com/wemixkanvas/kanvas/components/node/eth"
	"github.com/wemixkanvas/kanvas/components/node/metrics"
	"github.com/wemixkanvas/kanvas/components/node/rollup"
	"github.com/wemixkanvas/kanvas/components/node/rollup/derive"
	"github.com/wemixkanvas/kanvas/components/node/rollup/driver"
)

// MockL1OriginSelector is a shim to override the origin as proposer, so we can force it to stay on an older origin.
type MockL1OriginSelector struct {
	actual         *driver.L1OriginSelector
	originOverride eth.L1BlockRef // override which origin gets picked
}

func (m *MockL1OriginSelector) FindL1Origin(ctx context.Context, l2Head eth.L2BlockRef) (eth.L1BlockRef, error) {
	if m.originOverride != (eth.L1BlockRef{}) {
		return m.originOverride, nil
	}
	return m.actual.FindL1Origin(ctx, l2Head)
}

// L2Proposer is an actor that functions like a rollup node,
// without the full P2P/API/Node stack, but just the derivation state, and simplified driver with sequencing ability.
type L2Proposer struct {
	L2Syncer

	proposer *driver.Proposer

	failL2GossipUnsafeBlock error // mock error

	mockL1OriginSelector *MockL1OriginSelector
}

func NewL2Proposer(t Testing, log log.Logger, l1 derive.L1Fetcher, eng L2API, cfg *rollup.Config, propConfDepth uint64) *L2Proposer {
	syncer := NewL2Syncer(t, log, l1, eng, cfg)
	attrBuilder := derive.NewFetchingAttributesBuilder(cfg, l1, eng)
	propConfDepthL1 := driver.NewConfDepth(propConfDepth, syncer.l1State.L1Head, l1)
	l1OriginSelector := &MockL1OriginSelector{
		actual: driver.NewL1OriginSelector(log, cfg, propConfDepthL1),
	}
	return &L2Proposer{
		L2Syncer:                *syncer,
		proposer:                driver.NewProposer(log, cfg, syncer.derivation, attrBuilder, l1OriginSelector, metrics.NoopMetrics),
		mockL1OriginSelector:    l1OriginSelector,
		failL2GossipUnsafeBlock: nil,
	}
}

// ActL2StartBlock starts building of a new L2 block on top of the head
func (p *L2Proposer) ActL2StartBlock(t Testing) {
	p.ActL2StartBlockCheckErr(t, nil)
}

func (p *L2Proposer) ActL2StartBlockCheckErr(t Testing, checkErr error) {
	if !p.l2PipelineIdle {
		t.InvalidAction("cannot start L2 build when derivation is not idle")
		return
	}
	if p.l2Building {
		t.InvalidAction("already started building L2 block")
		return
	}

	err := p.proposer.StartBuildingBlock(t.Ctx())
	if checkErr == nil {
		require.NoError(t, err, "failed to start block building")
	} else {
		require.ErrorIs(t, err, checkErr, "expected typed error")
	}

	if errors.Is(err, derive.ErrReset) {
		p.derivation.Reset()
	}

	if err == nil {
		p.l2Building = true
	}
}

// ActL2EndBlock completes a new L2 block and applies it to the L2 chain as new canonical unsafe head
func (p *L2Proposer) ActL2EndBlock(t Testing) {
	if !p.l2Building {
		t.InvalidAction("cannot end L2 block building when no block is being built")
		return
	}
	p.l2Building = false

	_, err := p.proposer.CompleteBuildingBlock(t.Ctx())
	// TODO: there may be legitimate temporary errors here, if we mock engine API RPC-failure.
	// For advanced tests we can catch those and print a warning instead.
	require.NoError(t, err)

	// TODO: action-test publishing of payload on p2p
}

// ActL2KeepL1Origin makes the proposer use the current L1 origin, even if the next origin is available.
func (p *L2Proposer) ActL2KeepL1Origin(t Testing) {
	parent := p.derivation.UnsafeL2Head()
	// force old origin, for testing purposes
	oldOrigin, err := p.l1.L1BlockRefByHash(t.Ctx(), parent.L1Origin.Hash)
	require.NoError(t, err, "failed to get current origin: %s", parent.L1Origin)
	p.mockL1OriginSelector.originOverride = oldOrigin
}

// ActBuildToL1Head builds empty blocks until (incl.) the L1 head becomes the L2 origin
func (p *L2Proposer) ActBuildToL1Head(t Testing) {
	for p.derivation.UnsafeL2Head().L1Origin.Number < p.l1State.L1Head().Number {
		p.ActL2PipelineFull(t)
		p.ActL2StartBlock(t)
		p.ActL2EndBlock(t)
	}
}

// ActBuildToL1HeadUnsafe builds empty blocks until (incl.) the L1 head becomes the L1 origin of the L2 head
func (p *L2Proposer) ActBuildToL1HeadUnsafe(t Testing) {
	for p.derivation.UnsafeL2Head().L1Origin.Number < p.l1State.L1Head().Number {
		// Note: the derivation pipeline does not run, we are just sequencing a block on top of the existing L2 chain.
		p.ActL2StartBlock(t)
		p.ActL2EndBlock(t)
	}
}

// ActBuildToL1HeadExcl builds empty blocks until (excl.) the L1 head becomes the L1 origin of the L2 head
func (p *L2Proposer) ActBuildToL1HeadExcl(t Testing) {
	for {
		p.ActL2PipelineFull(t)
		nextOrigin, err := p.mockL1OriginSelector.FindL1Origin(t.Ctx(), p.derivation.UnsafeL2Head())
		require.NoError(t, err)
		if nextOrigin.Number >= p.l1State.L1Head().Number {
			break
		}
		p.ActL2StartBlock(t)
		p.ActL2EndBlock(t)
	}
}

// ActBuildToL1HeadExclUnsafe builds empty blocks until (excl.) the L1 head becomes the L1 origin of the L2 head, without safe-head progression.
func (p *L2Proposer) ActBuildToL1HeadExclUnsafe(t Testing) {
	for {
		// Note: the derivation pipeline does not run, we are just sequencing a block on top of the existing L2 chain.
		nextOrigin, err := p.mockL1OriginSelector.FindL1Origin(t.Ctx(), p.derivation.UnsafeL2Head())
		require.NoError(t, err)
		if nextOrigin.Number >= p.l1State.L1Head().Number {
			break
		}
		p.ActL2StartBlock(t)
		p.ActL2EndBlock(t)
	}
}

func (p *L2Proposer) ActBuildL2ToBlue(t Testing) {
	require.NotNil(t, p.rollupCfg.BlueTime, "cannot activate Blue when it is not scheduled")
	for p.L2Unsafe().Time < *p.rollupCfg.BlueTime {
		p.ActL2StartBlock(t)
		p.ActL2EndBlock(t)
	}
}
