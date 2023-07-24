package validator

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/log"
	"github.com/urfave/cli"

	"github.com/kroma-network/kroma/components/validator/metrics"
	"github.com/kroma-network/kroma/utils"
	"github.com/kroma-network/kroma/utils/monitoring"
	klog "github.com/kroma-network/kroma/utils/service/log"
	krpc "github.com/kroma-network/kroma/utils/service/rpc"
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
}

func NewValidator(ctx context.Context, cfg Config, l log.Logger, m metrics.Metricer) (*Validator, error) {
	// Validate the validator config
	if err := cfg.Check(); err != nil {
		return nil, err
	}

	var err error
	var l2os *L2OutputSubmitter
	if cfg.OutputSubmitterEnabled {
		l2os, err = NewL2OutputSubmitter(ctx, cfg, l, m)
		if err != nil {
			return nil, err
		}
	}

	challenger, err := NewChallenger(ctx, cfg, l, m)
	if err != nil {
		return nil, err
	}

	var guardian *Guardian
	if cfg.GuardianEnabled {
		guardian, err = NewGuardian(cfg, l)
		if err != nil {
			return nil, err
		}
	}

	return &Validator{
		cfg:        cfg,
		l:          l,
		metr:       m,
		l2os:       l2os,
		challenger: challenger,
		guardian:   guardian,
	}, nil
}

func (v *Validator) Start() error {
	v.ctx, v.cancel = context.WithCancel(context.Background())
	v.l.Info("starting Validator", "outputSubmitter", v.cfg.OutputSubmitterEnabled, "challenger", v.cfg.ChallengerEnabled, "guardian", v.cfg.GuardianEnabled)

	if err := v.cfg.TxManager.Start(v.ctx); err != nil {
		return fmt.Errorf("cannot start TxManager: %w", err)
	}

	if v.cfg.OutputSubmitterEnabled {
		if err := v.l2os.Start(v.ctx); err != nil {
			return fmt.Errorf("cannot start l2 output submitter: %w", err)
		}
	}

	if err := v.challenger.Start(v.ctx); err != nil {
		return fmt.Errorf("cannot start challenger: %w", err)
	}

	if v.cfg.GuardianEnabled {
		if err := v.guardian.Start(v.ctx); err != nil {
			return fmt.Errorf("cannot start guardian: %w", err)
		}
	}

	return nil
}

func (v *Validator) Stop() error {
	v.l.Info("stopping Validator")
	if err := v.cfg.TxManager.Stop(); err != nil {
		return fmt.Errorf("failed to stop TxManager: %w", err)
	}

	if v.cfg.ProofFetcher != nil {
		if err := v.cfg.ProofFetcher.Close(); err != nil {
			return fmt.Errorf("cannot close gRPC connection: %w", err)
		}
	}

	if v.cfg.OutputSubmitterEnabled {
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

	return nil
}
