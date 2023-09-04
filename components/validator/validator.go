package validator

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/log"
	"github.com/urfave/cli/v2"

	"github.com/kroma-network/kroma/bindings/bindings"
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

	l2ooContract *bindings.L2OutputOracleCaller
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
		guardian, err = NewGuardian(ctx, cfg, l)
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

	if err := v.cfg.TxManager.Start(v.ctx); err != nil {
		return fmt.Errorf("cannot start TxManager: %w", err)
	}

	if v.cfg.OutputSubmitterEnabled {
		if err := v.l2os.Start(v.ctx); err != nil {
			return fmt.Errorf("cannot start l2 output submitter: %w", err)
		}
	}

	if v.cfg.OutputSubmitterEnabled || v.cfg.ChallengerEnabled {
		if err := v.challenger.Start(v.ctx); err != nil {
			return fmt.Errorf("cannot start challenger: %w", err)
		}
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

	if v.cfg.OutputSubmitterEnabled {
		if err := v.l2os.Stop(); err != nil {
			return fmt.Errorf("failed to stop l2 output submitter: %w", err)
		}
	}

	if v.cfg.OutputSubmitterEnabled || v.cfg.ChallengerEnabled {
		if err := v.challenger.Stop(); err != nil {
			return fmt.Errorf("failed to stop challenger: %w", err)
		}
	}

	if v.cfg.GuardianEnabled {
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
	latestBlockNumber, err := v.l2ooContract.LatestBlockNumber(utils.NewSimpleCallOpts(cCtx))
	if err != nil {
		return nil, fmt.Errorf("unable to get latest block number of L2OutputOracle contract: %w", err)
	}

	return latestBlockNumber, nil
}
