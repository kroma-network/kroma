package node

import (
	"context"

	"github.com/libp2p/go-libp2p/core/peer"

	"github.com/ethereum-optimism/optimism/op-service/eth"
)

// Tracer configures the KromaNode to share events
type Tracer interface {
	OnNewL1Head(ctx context.Context, sig eth.L1BlockRef)
	OnUnsafeL2Payload(ctx context.Context, from peer.ID, payload *eth.ExecutionPayload)
	OnPublishL2Payload(ctx context.Context, payload *eth.ExecutionPayload)
}

type noKromaTracer struct{}

func (n noKromaTracer) OnNewL1Head(ctx context.Context, sig eth.L1BlockRef) {}

func (n noKromaTracer) OnUnsafeL2Payload(ctx context.Context, from peer.ID, payload *eth.ExecutionPayload) {
}

func (n noKromaTracer) OnPublishL2Payload(ctx context.Context, payload *eth.ExecutionPayload) {}

var _ Tracer = (*noKromaTracer)(nil)
