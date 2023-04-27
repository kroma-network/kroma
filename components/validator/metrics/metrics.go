package metrics

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/kroma-network/kroma/components/node/eth"
	kmetrics "github.com/kroma-network/kroma/utils/service/metrics"
	txmetrics "github.com/kroma-network/kroma/utils/service/txmgr/metrics"
)

const (
	Namespace         = "kroma_validator"
	L2OutputSubmitted = "submitted"
)

type Metricer interface {
	RecordInfo(version string)
	RecordUp()

	// Records all L1 and L2 block events
	kmetrics.RefMetricer

	// Record Tx metrics
	txmetrics.TxMetricer

	RecordL2OutputSubmitted(l2ref eth.L2BlockRef)
}

type Metrics struct {
	ns       string
	registry *prometheus.Registry
	factory  kmetrics.Factory

	kmetrics.RefMetrics
	txmetrics.TxMetrics

	Info prometheus.GaugeVec
	Up   prometheus.Gauge
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
			Help:      "1 if the kroma-validator has finished starting up",
		}),
	}
}

func (m *Metrics) Serve(ctx context.Context, host string, port int) error {
	return kmetrics.ListenAndServe(ctx, m.registry, host, port)
}

func (m *Metrics) StartBalanceMetrics(ctx context.Context,
	l log.Logger, client *ethclient.Client, account common.Address,
) {
	kmetrics.LaunchBalanceMetrics(ctx, l, m.registry, m.ns, client, account)
}

// RecordInfo sets a pseudo-metric that contains versioning and
// config info for the kroma-validator.
func (m *Metrics) RecordInfo(version string) {
	m.Info.WithLabelValues(version).Set(1)
}

// RecordUp sets the up metric to 1.
func (m *Metrics) RecordUp() {
	prometheus.MustRegister()
	m.Up.Set(1)
}

// RecordL2OutputSubmitted should be called when new L2 output is submitted
func (m *Metrics) RecordL2OutputSubmitted(l2ref eth.L2BlockRef) {
	m.RecordL2Ref(L2OutputSubmitted, l2ref)
}
