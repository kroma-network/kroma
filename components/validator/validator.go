package validator

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/log"
	"github.com/urfave/cli"

	"github.com/kroma-network/kroma/components/validator/metrics"
	"github.com/kroma-network/kroma/utils"
	"github.com/kroma-network/kroma/utils/monitoring"
	klog "github.com/kroma-network/kroma/utils/service/log"
	krpc "github.com/kroma-network/kroma/utils/service/rpc"
	"github.com/kroma-network/kroma/utils/service/txmgr"
)

// Main is the entrypoint into the Validator. This method executes the
// service and blocks until the service exits.
func Main(version string, cliCtx *cli.Context) error {
	cliCfg := NewCLIConfig(cliCtx)
	if err := cliCfg.Check(); err != nil {
		return fmt.Errorf("invalid CLI flags: %w", err)
	}

	l := klog.NewLogger(cliCfg.LogConfig)
	m := metrics.NewMetrics("default")
	l.Info("initializing Validator")

	validatorCfg, err := NewValidatorConfig(cliCfg, l, m)
	if err != nil {
		l.Error("Unable to create validator config", "err", err)
		return err
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	monitoring.MaybeStartPprof(ctx, cliCfg.PprofConfig, l)
	monitoring.MaybeStartMetrics(ctx, cliCfg.MetricsConfig, l, m, validatorCfg.L1Client, validatorCfg.TxManager.From())
	server, err := monitoring.StartRPC(cliCfg.RPCConfig, version, krpc.WithLogger(l))
	if err != nil {
		return err
	}
	defer func() {
		if err = server.Stop(); err != nil {
			l.Error("Error shutting down http server: %w", err)
		}
	}()

	m.RecordInfo(version)
	m.RecordUp()

	validator, err := NewValidator(ctx, *validatorCfg, l, m)
	if err != nil {
		return err
	}

	validator.Start()
	<-utils.WaitInterrupt()
	validator.Stop()

	return nil
}

type Validator struct {
	ctx    context.Context
	cancel context.CancelFunc

	cfg        Config
	l          log.Logger
	metr       metrics.Metricer
	l2os       *L2OutputSubmitter
	challenger *Challenger

	wg sync.WaitGroup
}

func NewValidator(parentCtx context.Context, cfg Config, l log.Logger, m metrics.Metricer) (*Validator, error) {
	// Validate the validator config
	if err := cfg.Check(); err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(parentCtx)

	l2OutputSubmitter, err := NewL2OutputSubmitter(cfg, l)
	if err != nil {
		cancel()
		return nil, err
	}

	challenger, err := NewChallenger(ctx, cfg, l)
	if err != nil {
		cancel()
		return nil, err
	}

	return &Validator{
		ctx:        ctx,
		cancel:     cancel,
		cfg:        cfg,
		l:          l,
		l2os:       l2OutputSubmitter,
		challenger: challenger,
		metr:       m,
	}, nil
}

func (v *Validator) Start() {
	v.l.Info("starting Validator")
	v.wg.Add(1)
	go v.loop()
}

func (v *Validator) Stop() {
	v.l.Info("stopping Validator")
	if v.cfg.ProofFetcher != nil {
		if err := v.cfg.ProofFetcher.Close(); err != nil {
			v.l.Error("cannot close grpc connection: %w", err)
		}
	}
	v.cancel()
	v.wg.Wait()
}

func (v *Validator) loop() {
	defer v.wg.Done()

	ticker := time.NewTicker(v.cfg.PollInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if !v.cfg.OutputSubmitterDisabled {
				if err := v.submitL2Output(); err != nil {
					v.l.Error("failed to submit l2 output", "err", err)
				}
			}

			if err := v.submitChallengeTx(); err != nil {
				v.l.Error("failed to submit challenge tx", "err", err)
			}
		case <-v.ctx.Done():
			return
		}
	}
}

func (v *Validator) submitL2Output() error {
	ctx := v.ctx
	output, shouldSubmit, err := v.l2os.FetchNextOutputInfo(ctx)
	if err != nil {
		return fmt.Errorf("failed to fetch next output: %w", err)
	}
	if !shouldSubmit {
		return nil
	}

	data, err := v.l2os.SubmitL2OutputTxData(output)
	if err != nil {
		return fmt.Errorf("failed to create submit l2 output transaction data: %w", err)
	}

	txCandidate := txmgr.TxCandidate{
		TxData:   data,
		To:       &v.cfg.L2OutputOracleAddr,
		GasLimit: 0,
	}

	if err := v.sendTransaction(ctx, txCandidate); err != nil {
		return fmt.Errorf("failed to send submit l2 output transaction: %w", err)
	}
	v.metr.RecordL2OutputSubmitted(output.BlockRef)

	return nil
}

func (v *Validator) submitChallengeTx() error {
	tx, err := v.challenger.DetermineChallengeTx()
	if err != nil {
		return fmt.Errorf("failed to determine challenge transaction to submit: %w", err)
	}

	if tx == nil {
		return nil
	}

	txCandidate := txmgr.TxCandidate{
		TxData:   tx.Data(),
		To:       tx.To(),
		GasLimit: 0,
	}

	if err := v.sendTransaction(v.ctx, txCandidate); err != nil {
		return fmt.Errorf("failed to send challenge transaction: %w", err)
	}

	return nil
}

// sendTransaction creates & sends transactions through the underlying transaction manager.
func (v *Validator) sendTransaction(ctx context.Context, txCandidate txmgr.TxCandidate) error {
	receipt, err := v.cfg.TxManager.Send(ctx, txCandidate)
	if err != nil {
		return err
	}
	v.l.Info("validator tx successfully published", "tx_hash", receipt.TxHash)
	return nil
}
