package validator

import (
	"context"
	"fmt"
	"math/big"
	"time"

	opservice "github.com/ethereum-optimism/optimism/op-service"
	oplog "github.com/ethereum-optimism/optimism/op-service/log"
	"github.com/ethereum-optimism/optimism/op-service/monitoring"
	"github.com/ethereum-optimism/optimism/op-service/opio"
	"github.com/ethereum-optimism/optimism/op-service/optsutils"
	oprpc "github.com/ethereum-optimism/optimism/op-service/rpc"
	"github.com/ethereum/go-ethereum/log"
	"github.com/urfave/cli/v2"

	"github.com/kroma-network/kroma/kroma-bindings/bindings"
	"github.com/kroma-network/kroma/kroma-validator/flags"
	"github.com/kroma-network/kroma/kroma-validator/metrics"
)

// Main is the entrypoint into the Validator. This method executes the
// service and blocks until the service exits.
func Main(version string, cliCtx *cli.Context) error {
	if err := flags.CheckRequired(cliCtx); err != nil {
		return err
	}
	cfg := NewConfig(cliCtx)
	if err := cfg.Check(); err != nil {
		return fmt.Errorf("invalid CLI flags: %w", err)
	}

	l := oplog.NewLogger(oplog.AppOut(cliCtx), cfg.LogConfig)
	oplog.SetGlobalLogHandler(l.Handler())
	opservice.ValidateEnvVars(flags.EnvVarPrefix, flags.Flags, l)
	m := metrics.NewMetrics("default")
	l.Info("initializing Validator")

	validatorCfg, err := NewValidatorConfig(cfg, l, m)
	if err != nil {
		l.Error("Unable to create validator config", "err", err)
		return err
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	monitoring.MaybeStartPprof(ctx, cfg.PprofConfig, l)
	monitoring.MaybeStartMetrics(ctx, cfg.MetricsConfig, l, m, validatorCfg.L1Client, validatorCfg.TxManager.From())
	server, err := monitoring.StartRPC(cfg.RPCConfig, version, oprpc.WithLogger(l))
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

	validator, err := NewValidator(*validatorCfg, l, m)
	if err != nil {
		return err
	}

	if err := validator.Start(); err != nil {
		l.Error("failed to start validator", "err", err)
		return err
	}
	opio.BlockOnInterrupts()
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

	l2ooContract *bindings.L2OutputOracleCaller
}

func NewValidator(cfg Config, l log.Logger, m metrics.Metricer) (*Validator, error) {
	// Validate the validator config
	if err := cfg.Check(); err != nil {
		return nil, err
	}

	var err error
	var l2os *L2OutputSubmitter
	if cfg.OutputSubmitterEnabled {
		l2os, err = NewL2OutputSubmitter(cfg, l, m)
		if err != nil {
			return nil, err
		}
	}

	var challenger *Challenger
	if cfg.OutputSubmitterEnabled || cfg.ChallengerEnabled {
		challenger, err = NewChallenger(cfg, l, m)
		if err != nil {
			return nil, err
		}
	}

	var guardian *Guardian
	if cfg.GuardianEnabled {
		guardian, err = NewGuardian(cfg, l)
		if err != nil {
			return nil, err
		}
	}

	l2ooContract, err := bindings.NewL2OutputOracleCaller(cfg.L2OutputOracleAddr, cfg.L1Client)
	if err != nil {
		return nil, err
	}

	return &Validator{
		cfg:          cfg,
		l:            l,
		metr:         m,
		l2os:         l2os,
		challenger:   challenger,
		guardian:     guardian,
		l2ooContract: l2ooContract,
	}, nil
}

func (v *Validator) Start() error {
	v.ctx, v.cancel = context.WithCancel(context.Background())
	v.l.Info("starting Validator", "outputSubmitter", v.cfg.OutputSubmitterEnabled, "challenger", v.cfg.ChallengerEnabled, "guardian", v.cfg.GuardianEnabled)

	// wait for kroma node to sync completed
	v.waitSyncCompleted()

	if v.cfg.TxManager != nil {
		if err := v.cfg.TxManager.Start(v.ctx); err != nil {
			return fmt.Errorf("cannot start TxManager: %w", err)
		}
	}

	if v.l2os != nil {
		if err := v.l2os.Start(v.ctx); err != nil {
			return fmt.Errorf("cannot start l2 output submitter: %w", err)
		}
	}

	if v.challenger != nil {
		if err := v.challenger.Start(v.ctx); err != nil {
			return fmt.Errorf("cannot start challenger: %w", err)
		}
	}

	if v.guardian != nil {
		if err := v.guardian.Start(v.ctx); err != nil {
			return fmt.Errorf("cannot start guardian: %w", err)
		}
	}

	return nil
}

func (v *Validator) Stop() error {
	v.l.Info("stopping Validator")

	if v.cfg.TxManager != nil {
		if err := v.cfg.TxManager.Stop(); err != nil {
			return fmt.Errorf("failed to stop TxManager: %w", err)
		}
	}

	if v.l2os != nil {
		if err := v.l2os.Stop(); err != nil {
			return fmt.Errorf("failed to stop l2 output submitter: %w", err)
		}
	}

	if v.challenger != nil {
		if err := v.challenger.Stop(); err != nil {
			return fmt.Errorf("failed to stop challenger: %w", err)
		}
	}

	if v.guardian != nil {
		if err := v.guardian.Stop(); err != nil {
			return fmt.Errorf("failed to stop guardian: %w", err)
		}
	}

	v.cancel()

	return nil
}

func (v *Validator) waitSyncCompleted() {
	v.l.Info("start waiting for kroma node to sync")

	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	for ; ; <-ticker.C {
		select {
		case <-v.ctx.Done():
			return
		default:
			currentBlockNumber, err := v.fetchCurrentBlockNumber()
			if err != nil {
				v.l.Debug(err.Error())
				continue
			}
			latestBlockNumber, err := v.fetchLatestBlockNumber()
			if err != nil {
				v.l.Error(err.Error())
				continue
			}
			if latestBlockNumber.Cmp(currentBlockNumber) == 1 {
				v.l.Info("wait for kroma node to sync", "currentBlockNumber", currentBlockNumber, "latestSubmittedBlockNumber", latestBlockNumber)
				continue
			}
			v.l.Info("kroma node sync completed")
			return
		}
	}
}

func (v *Validator) fetchCurrentBlockNumber() (*big.Int, error) {
	// fetch the current L2 heads
	cCtx, cCancel := context.WithTimeout(v.ctx, v.cfg.NetworkTimeout)
	defer cCancel()
	status, err := v.cfg.RollupClient.SyncStatus(cCtx)
	if err != nil {
		return nil, fmt.Errorf("unable to get sync status: %w", err)
	}

	// Use either the finalized or safe head depending on the config. Finalized head is default & safer.
	var currentBlockNumber *big.Int
	if v.cfg.AllowNonFinalized {
		currentBlockNumber = new(big.Int).SetUint64(status.SafeL2.Number)
	} else {
		currentBlockNumber = new(big.Int).SetUint64(status.FinalizedL2.Number)
	}

	return currentBlockNumber, nil
}

func (v *Validator) fetchLatestBlockNumber() (*big.Int, error) {
	cCtx, cCancel := context.WithTimeout(v.ctx, v.cfg.NetworkTimeout)
	defer cCancel()
	latestBlockNumber, err := v.l2ooContract.LatestBlockNumber(optsutils.NewSimpleCallOpts(cCtx))
	if err != nil {
		return nil, fmt.Errorf("unable to get latest block number of L2OutputOracle contract: %w", err)
	}

	return latestBlockNumber, nil
}
