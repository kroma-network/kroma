package metrics

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/kroma-network/kroma/components/node/eth"
	"github.com/kroma-network/kroma/components/node/rollup/derive"
	kmetrics "github.com/kroma-network/kroma/utils/service/metrics"
	txmetrics "github.com/kroma-network/kroma/utils/service/txmgr/metrics"
)

const Namespace = "kroma_batcher"

type Metricer interface {
	RecordInfo(version string)
	RecordUp()

	// Records all L1 and L2 block events
	kmetrics.RefMetricer

	// Record Tx metrics
	txmetrics.TxMetricer

	RecordLatestL1Block(l1ref eth.L1BlockRef)
	RecordL2BlocksLoaded(l2ref eth.L2BlockRef)
	RecordChannelOpened(id derive.ChannelID, numPendingBlocks int)
	RecordL2BlocksAdded(l2ref eth.L2BlockRef, numBlocksAdded, numPendingBlocks, inputBytes, outputComprBytes int)
	RecordChannelClosed(id derive.ChannelID, numPendingBlocks int, numFrames int, inputBytes int, outputComprBytes int, reason error)
	RecordChannelFullySubmitted(id derive.ChannelID)
	RecordChannelTimedOut(id derive.ChannelID)

	RecordBatchTxSubmitted()
	RecordBatchTxSuccess()
	RecordBatchTxFailed()

	Document() []kmetrics.DocumentedMetric
}

type Metrics struct {
	ns       string
	registry *prometheus.Registry
	factory  kmetrics.Factory

	kmetrics.RefMetrics
	txmetrics.TxMetrics

	Info prometheus.GaugeVec
	Up   prometheus.Gauge

	// label by opened, closed, fully_submitted, timed_out
	ChannelEvs kmetrics.EventVec

	PendingBlocksCount prometheus.GaugeVec
	BlocksAddedCount   prometheus.Gauge

	ChannelInputBytes   prometheus.GaugeVec
	ChannelReadyBytes   prometheus.Gauge
	ChannelOutputBytes  prometheus.Gauge
	ChannelClosedReason prometheus.Gauge
	ChannelNumFrames    prometheus.Gauge
	ChannelComprRatio   prometheus.Histogram

	BatcherTxEvs kmetrics.EventVec
}

var _ Metricer = (*Metrics)(nil)

func NewMetrics(procName string) *Metrics {
	if procName == "" {
		procName = "default"
	}
	ns := Namespace + "_" + procName

	registry := kmetrics.NewRegistry()
	factory := kmetrics.With(registry)

	return &Metrics{
		ns:       ns,
		registry: registry,
		factory:  factory,

		RefMetrics: kmetrics.MakeRefMetrics(ns, factory),
		TxMetrics:  txmetrics.MakeTxMetrics(ns, factory),

		Info: *factory.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: ns,
			Name:      "info",
			Help:      "Pseudo-metric tracking version and config info",
		}, []string{
			"version",
		}),
		Up: factory.NewGauge(prometheus.GaugeOpts{
			Namespace: ns,
			Name:      "up",
			Help:      "1 if the kroma-batcher has finished starting up",
		}),

		ChannelEvs: kmetrics.NewEventVec(factory, ns, "channel", "Channel", []string{"stage"}),

		PendingBlocksCount: *factory.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: ns,
			Name:      "pending_blocks_count",
			Help:      "Number of pending blocks, not added to a channel yet.",
		}, []string{"stage"}),
		BlocksAddedCount: factory.NewGauge(prometheus.GaugeOpts{
			Namespace: ns,
			Name:      "blocks_added_count",
			Help:      "Total number of blocks added to current channel.",
		}),

		ChannelInputBytes: *factory.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: ns,
			Name:      "input_bytes",
			Help:      "Number of input bytes to a channel.",
		}, []string{"stage"}),
		ChannelReadyBytes: factory.NewGauge(prometheus.GaugeOpts{
			Namespace: ns,
			Name:      "ready_bytes",
			Help:      "Number of bytes ready in the compression buffer.",
		}),
		ChannelOutputBytes: factory.NewGauge(prometheus.GaugeOpts{
			Namespace: ns,
			Name:      "output_bytes",
			Help:      "Number of compressed output bytes from a channel.",
		}),
		ChannelClosedReason: factory.NewGauge(prometheus.GaugeOpts{
			Namespace: ns,
			Name:      "channel_closed_reason",
			Help:      "Pseudo-metric to record the reason a channel got closed.",
		}),
		ChannelNumFrames: factory.NewGauge(prometheus.GaugeOpts{
			Namespace: ns,
			Name:      "channel_num_frames",
			Help:      "Total number of frames of closed channel.",
		}),
		ChannelComprRatio: factory.NewHistogram(prometheus.HistogramOpts{
			Namespace: ns,
			Name:      "channel_compr_ratio",
			Help:      "Compression ratios of closed channel.",
			Buckets:   append([]float64{0.1, 0.2}, prometheus.LinearBuckets(0.3, 0.05, 14)...),
		}),

		BatcherTxEvs: kmetrics.NewEventVec(factory, ns, "batcher_tx", "BatcherTx", []string{"stage"}),
	}
}

func (m *Metrics) Serve(ctx context.Context, host string, port int) error {
	return kmetrics.ListenAndServe(ctx, m.registry, host, port)
}

func (m *Metrics) Document() []kmetrics.DocumentedMetric {
	return m.factory.Document()
}

func (m *Metrics) StartBalanceMetrics(ctx context.Context,
	l log.Logger, client *ethclient.Client, account common.Address,
) {
	kmetrics.LaunchBalanceMetrics(ctx, l, m.registry, m.ns, client, account)
}

// RecordInfo sets a pseudo-metric that contains versioning and
// config info for the kroma-batcher.
func (m *Metrics) RecordInfo(version string) {
	m.Info.WithLabelValues(version).Set(1)
}

// RecordUp sets the up metric to 1.
func (m *Metrics) RecordUp() {
	prometheus.MustRegister()
	m.Up.Set(1)
}

const (
	StageLoaded         = "loaded"
	StageOpened         = "opened"
	StageAdded          = "added"
	StageClosed         = "closed"
	StageFullySubmitted = "fully_submitted"
	StageTimedOut       = "timed_out"

	TxStageSubmitted = "submitted"
	TxStageSuccess   = "success"
	TxStageFailed    = "failed"
)

func (m *Metrics) RecordLatestL1Block(l1ref eth.L1BlockRef) {
	m.RecordL1Ref("latest", l1ref)
}

// RecordL2BlocksLoaded should be called when a new L2 block was loaded into the
// channel manager (but not processed yet).
func (m *Metrics) RecordL2BlocksLoaded(l2ref eth.L2BlockRef) {
	m.RecordL2Ref(StageLoaded, l2ref)
}

func (m *Metrics) RecordChannelOpened(id derive.ChannelID, numPendingBlocks int) {
	m.ChannelEvs.Record(StageOpened)
	m.BlocksAddedCount.Set(0) // reset
	m.PendingBlocksCount.WithLabelValues(StageOpened).Set(float64(numPendingBlocks))
}

// RecordL2BlocksAdded should be called when L2 block were added to the channel
// builder, with the latest added block.
func (m *Metrics) RecordL2BlocksAdded(l2ref eth.L2BlockRef, numBlocksAdded, numPendingBlocks, inputBytes, outputComprBytes int) {
	m.RecordL2Ref(StageAdded, l2ref)
	m.BlocksAddedCount.Add(float64(numBlocksAdded))
	m.PendingBlocksCount.WithLabelValues(StageAdded).Set(float64(numPendingBlocks))
	m.ChannelInputBytes.WithLabelValues(StageAdded).Set(float64(inputBytes))
	m.ChannelReadyBytes.Set(float64(outputComprBytes))
}

func (m *Metrics) RecordChannelClosed(id derive.ChannelID, numPendingBlocks int, numFrames int, inputBytes int, outputComprBytes int, reason error) {
	m.ChannelEvs.Record(StageClosed)
	m.PendingBlocksCount.WithLabelValues(StageClosed).Set(float64(numPendingBlocks))
	m.ChannelNumFrames.Set(float64(numFrames))
	m.ChannelInputBytes.WithLabelValues(StageClosed).Set(float64(inputBytes))
	m.ChannelOutputBytes.Set(float64(outputComprBytes))

	var comprRatio float64
	if inputBytes > 0 {
		comprRatio = float64(outputComprBytes) / float64(inputBytes)
	}
	m.ChannelComprRatio.Observe(comprRatio)

	m.ChannelClosedReason.Set(float64(ClosedReasonToNum(reason)))
}

func ClosedReasonToNum(reason error) int {
	// CLI-3640
	return 0
}

func (m *Metrics) RecordChannelFullySubmitted(id derive.ChannelID) {
	m.ChannelEvs.Record(StageFullySubmitted)
}

func (m *Metrics) RecordChannelTimedOut(id derive.ChannelID) {
	m.ChannelEvs.Record(StageTimedOut)
}

func (m *Metrics) RecordBatchTxSubmitted() {
	m.BatcherTxEvs.Record(TxStageSubmitted)
}

func (m *Metrics) RecordBatchTxSuccess() {
	m.BatcherTxEvs.Record(TxStageSuccess)
}

func (m *Metrics) RecordBatchTxFailed() {
	m.BatcherTxEvs.Record(TxStageFailed)
}
