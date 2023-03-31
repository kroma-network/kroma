package validator

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/urfave/cli"

	"github.com/wemixkanvas/kanvas/utils"
	"github.com/wemixkanvas/kanvas/utils/monitoring"
	klog "github.com/wemixkanvas/kanvas/utils/service/log"
	krpc "github.com/wemixkanvas/kanvas/utils/service/rpc"
	"github.com/wemixkanvas/kanvas/utils/service/txmgr"
)

// Main is the entrypoint into the Validator. This method executes the
// service and blocks until the service exits.
func Main(version string, cliCtx *cli.Context) error {
	cliCfg := NewCLIConfig(cliCtx)
	if err := cliCfg.Check(); err != nil {
		return fmt.Errorf("invalid CLI flags: %w", err)
	}

	l := klog.NewLogger(cliCfg.LogConfig)
	l.Info("initializing Validator")

	validatorCfg, err := NewValidatorConfig(cliCfg, l)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	monitoring.MaybeStartPprof(ctx, cliCfg.PprofConfig, l)
	monitoring.MaybeStartMetrics(ctx, cliCfg.MetricsConfig, l, validatorCfg.L1Client, validatorCfg.From)
	server, err := monitoring.StartRPC(cliCfg.RPCConfig, version, krpc.WithLogger(l))
	if err != nil {
		return err
	}
	defer server.Stop()

	validator, err := NewValidator(ctx, *validatorCfg, l)
	if err != nil {
		return err
	}

	validator.Start()
	<-utils.WaitInterrupt()
	validator.Stop()

	return nil
}

type Validator struct {
	ctx        context.Context
	cancel     context.CancelFunc
	cfg        Config
	l          log.Logger
	l2os       *L2OutputSubmitter
	challenger *Challenger
	txMgr      txmgr.TxManager

	wg sync.WaitGroup
}

func NewValidator(parentCtx context.Context, cfg Config, l log.Logger) (*Validator, error) {
	ctx, cancel := context.WithCancel(parentCtx)

	l2OutputSubmitter, err := NewL2OutputSubmitter(ctx, cfg, l)
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
		txMgr:      txmgr.NewSimpleTxManager("validator", l, cfg.TxManagerConfig, cfg.L1Client),
	}, nil
}

func (v *Validator) Start() {
	v.l.Info("starting Validator")
	v.wg.Add(1)
	go v.loop()
}

func (v *Validator) Stop() {
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
		default:
		}
	}
}

func (v *Validator) submitL2Output() error {
	cCtx, cancel := context.WithTimeout(v.ctx, 3*time.Minute)
	defer cancel()

	output, shouldSubmit, err := v.l2os.FetchNextOutputInfo(cCtx)
	if err != nil {
		return fmt.Errorf("failed to fetch next output: %w", err)
	}
	if !shouldSubmit {
		return nil
	}

	tx, err := v.l2os.CreateSubmitL2OutputTx(cCtx, output)
	if err != nil {
		return fmt.Errorf("failed to create submit l2 output transaction: %w", err)
	}
	if err := v.SendTransaction(cCtx, tx); err != nil {
		return fmt.Errorf("failed to send submit l2 output transaction: %w", err)
	}

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

	if err := v.SendTransaction(v.ctx, tx); err != nil {
		return fmt.Errorf("failed to send challenge transaction: %w", err)
	}

	return nil
}

// SendTransaction sends a transaction through the transaction manager which handles automatic
// price bumping.
// It also hardcodes a timeout of 100s.
func (v *Validator) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	// Wait until one of our submitted transactions confirms. If no
	// receipt is received it's likely our gas price was too low.
	cCtx, cancel := context.WithTimeout(ctx, 100*time.Second)
	defer cancel()
	v.l.Info("validator sending transaction", "tx", tx.Hash())
	receipt, err := v.txMgr.Send(cCtx, tx)
	if err != nil {
		v.l.Error("validator unable to publish tx", "err", err)
		return err
	}

	// The transaction was successfully submitted
	v.l.Info("validator tx successfully published", "tx_hash", receipt.TxHash)
	return nil
}
