package validator

import (
	"context"
	"fmt"
	"sync"

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

	if err := validator.Start(); err != nil {
		l.Error("failed to start validator", "err", err)
		return err
	}
	<-utils.WaitInterrupt()
	if err := validator.Stop(); err != nil {
		l.Error("failed to stop validator", "err", err)
		return err
	}

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
	guardian   *Guardian

	txCandidatesChan chan txmgr.TxCandidate

	wg sync.WaitGroup
}

func NewValidator(parentCtx context.Context, cfg Config, l log.Logger, m metrics.Metricer) (*Validator, error) {
	// Validate the validator config
	if err := cfg.Check(); err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(parentCtx)

	txCandidatesChan := make(chan txmgr.TxCandidate, 10)

	l2OutputSubmitter, err := NewL2OutputSubmitter(ctx, cfg, l, m, txCandidatesChan)
	if err != nil {
		cancel()
		return nil, err
	}

	challenger, err := NewChallenger(ctx, cfg, l, txCandidatesChan)
	if err != nil {
		cancel()
		return nil, err
	}

	guardian, err := NewGuardian(ctx, cfg, l, txCandidatesChan)
	if err != nil {
		cancel()
		return nil, err
	}

	return &Validator{
		ctx:              ctx,
		cancel:           cancel,
		cfg:              cfg,
		l:                l,
		metr:             m,
		l2os:             l2OutputSubmitter,
		challenger:       challenger,
		guardian:         guardian,
		txCandidatesChan: txCandidatesChan,
	}, nil
}

func (v *Validator) Start() error {
	v.l.Info("starting Validator")

	if !v.cfg.OutputSubmitterDisabled {
		if err := v.l2os.Start(); err != nil {
			return fmt.Errorf("cannot start l2 output submitter: %w", err)
		}
	}

	if err := v.challenger.Start(); err != nil {
		return fmt.Errorf("cannot start challenger: %w", err)
	}

	if v.cfg.GuardianEnabled {
		if err := v.guardian.Start(); err != nil {
			return fmt.Errorf("cannot start guardian: %w", err)
		}
	}

	v.wg.Add(1)
	go v.loop()

	return nil
}

func (v *Validator) Stop() error {
	v.l.Info("stopping Validator")
	if v.cfg.ProofFetcher != nil {
		if err := v.cfg.ProofFetcher.Close(); err != nil {
			return fmt.Errorf("cannot close gRPC connection: %w", err)
		}
	}

	if !v.cfg.OutputSubmitterDisabled {
		if err := v.l2os.Stop(); err != nil {
			return fmt.Errorf("failed to stop l2 output submitter: %w", err)
		}
	}

	if err := v.challenger.Stop(); err != nil {
		return fmt.Errorf("failed to stop challenger: %w", err)
	}

	if v.cfg.GuardianEnabled {
		if err := v.guardian.Stop(); err != nil {
			return fmt.Errorf("failed to stop guardian: %w", err)
		}
	}

	v.cancel()
	v.wg.Wait()
	return nil
}

func (v *Validator) loop() {
	defer v.wg.Done()

	for {
		select {
		case txCandidate := <-v.txCandidatesChan:
			if err := v.sendTransaction(v.ctx, txCandidate); err != nil {
				v.l.Error("failed to submit transaction of validator", "err", err)
			}
		case <-v.ctx.Done():
			return
		}
	}
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
