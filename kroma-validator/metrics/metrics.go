package metrics

import (
	"context"
	"math/big"

	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum-optimism/optimism/op-service/httputil"
	opmetrics "github.com/ethereum-optimism/optimism/op-service/metrics"
	txmetrics "github.com/ethereum-optimism/optimism/op-service/txmgr/metrics"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	Namespace         = "kroma_validator"
	L2OutputSubmitted = "submitted"
)

type Metricer interface {
	RecordInfo(version string)
	RecordUp()

	// Records all L1 and L2 block events
	opmetrics.RefMetricer

	// Record Tx metrics
	txmetrics.TxMetricer

	RecordL2OutputSubmitted(l2ref eth.L2BlockRef)
	RecordUnbondedDepositAmount(amount *big.Int)
	RecordValidatorStatus(status uint8)
	RecordNextValidator(address common.Address)
	RecordChallengeCheckpoint(outputIndex *big.Int)
}

type Metrics struct {
	ns       string
	registry *prometheus.Registry
	factory  opmetrics.Factory

	opmetrics.RefMetrics
	txmetrics.TxMetrics
	opmetrics.RPCMetrics

	Info                  prometheus.GaugeVec
	Up                    prometheus.Gauge
	UnbondedDepositAmount prometheus.Gauge
	ValidatorStatus       prometheus.Gauge
	NextValidator         prometheus.GaugeVec
	ChallengeCheckpoint   prometheus.Gauge
}

var _ Metricer = (*Metrics)(nil)

func NewMetrics(procName string) *Metrics {
	if procName == "" {
		procName = "default"
	}
	ns := Namespace + "_" + procName

	registry := opmetrics.NewRegistry()
	factory := opmetrics.With(registry)

	return &Metrics{
		ns:       ns,
		registry: registry,
		factory:  factory,

		RefMetrics: opmetrics.MakeRefMetrics(ns, factory),
		TxMetrics:  txmetrics.MakeTxMetrics(ns, factory),
		RPCMetrics: opmetrics.MakeRPCMetrics(ns, factory),

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
		UnbondedDepositAmount: factory.NewGauge(prometheus.GaugeOpts{
			Namespace: ns,
			Name:      "unbonded_deposit_amount",
			Help:      "The amount of Validator balance excluding the bonded amount",
		}),
		ValidatorStatus: factory.NewGauge(prometheus.GaugeOpts{
			Namespace: ns,
			Name:      "validator_status",
			Help:      "The status of validator in the ValidatorManager contract",
		}),
		NextValidator: *factory.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: ns,
			Name:      "next_validator",
			Help:      "The address of the next validator",
		}, []string{
			"address",
		}),
		ChallengeCheckpoint: factory.NewGauge(prometheus.GaugeOpts{
			Namespace: ns,
			Name:      "challenge_checkpoint",
			Help:      "The output index that the challenge function last checked",
		}),
	}
}

func (m *Metrics) Start(host string, port int) (*httputil.HTTPServer, error) {
	return opmetrics.StartServer(m.registry, host, port)
}

func (m *Metrics) StartBalanceMetrics(ctx context.Context,
	l log.Logger, client *ethclient.Client, account common.Address,
) {
	// TODO(7684): util was refactored to close, but ctx is still being used by caller for shutdown
	balanceMetric := opmetrics.LaunchBalanceMetrics(l, m.registry, m.ns, client, account)
	go func() {
		<-ctx.Done()
		_ = balanceMetric.Close()
	}()
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

// RecordUnbondedDepositAmount sets the amount deposited into the ValidatorPool contract.
func (m *Metrics) RecordUnbondedDepositAmount(amount *big.Int) {
	m.UnbondedDepositAmount.Set(opmetrics.WeiToEther(amount))
}

// RecordValidatorStatus sets the status of validator in the ValidatorManager contract.
func (m *Metrics) RecordValidatorStatus(status uint8) {
	m.ValidatorStatus.Set(float64(status))
}

// RecordNextValidator sets the address of the next validator.
func (m *Metrics) RecordNextValidator(address common.Address) {
	m.NextValidator.WithLabelValues(address.String()).Set(1)
}

// RecordChallengeCheckpoint sets the output index that the challenge function last checked.
func (m *Metrics) RecordChallengeCheckpoint(outputIndex *big.Int) {
	m.ChallengeCheckpoint.Set(float64(outputIndex.Uint64()))
}
