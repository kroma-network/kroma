package batcher

import (
	"context"
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/urfave/cli"

	"github.com/kroma-network/kroma/components/batcher/metrics"
	"github.com/kroma-network/kroma/utils"
	"github.com/kroma-network/kroma/utils/monitoring"
	klog "github.com/kroma-network/kroma/utils/service/log"
	krpc "github.com/kroma-network/kroma/utils/service/rpc"
	"github.com/kroma-network/kroma/utils/service/txmgr"
)

// Main is the entrypoint into the Batcher.
func Main(version string, cliCtx *cli.Context) error {
	cliCfg := NewCLIConfig(cliCtx)
	if err := cliCfg.Check(); err != nil {
		return fmt.Errorf("invalid CLI flags: %w", err)
	}

	l := klog.NewLogger(cliCfg.LogConfig)
	m := metrics.NewMetrics("default")
	l.Info("Initializing Batcher")

	batcherCfg, err := NewBatcherConfig(cliCfg, l, m)
	if err != nil {
		l.Error("Unable to create batcher config", "err", err)
		return err
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	monitoring.MaybeStartPprof(ctx, cliCfg.PprofConfig, l)
	monitoring.MaybeStartMetrics(ctx, cliCfg.MetricsConfig, l, batcherCfg.L1Client, batcherCfg.TxManager.From())
	server, err := monitoring.StartRPC(cliCfg.RPCConfig.ToServiceCLIConfig(), version, krpc.WithLogger(l))
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

	batcher, err := NewBatcher(ctx, *batcherCfg, l, m)
	if err != nil {
		return err
	}

	batcher.Start()
	<-utils.WaitInterrupt()
	batcher.Stop(context.Background())

	return nil
}

type Batcher struct {
	shutdownCtx       context.Context
	cancelShutdownCtx context.CancelFunc
	killCtx           context.Context
	cancelKillCtx     context.CancelFunc

	cfg            Config
	l              log.Logger
	batchSubmitter *BatchSubmitter

	wg sync.WaitGroup
}

func NewBatcher(parentCtx context.Context, cfg Config, l log.Logger, m metrics.Metricer) (*Batcher, error) {
	// Validate the batcher config
	if err := cfg.Check(); err != nil {
		return nil, err
	}

	batchSubmitter, err := NewBatchSubmitter(cfg, l, m)
	if err != nil {
		return nil, fmt.Errorf("failed to init batch submitter: %w", err)
	}

	balance, err := cfg.L1Client.BalanceAt(parentCtx, cfg.TxManager.From(), nil)
	if err != nil {
		return nil, err
	}

	l.Info("creating batcher", "batcher_addr", cfg.TxManager.From(), "batcher_bal", balance)

	return &Batcher{
		cfg:            cfg,
		l:              l,
		batchSubmitter: batchSubmitter,
	}, nil
}

func (b *Batcher) Start() {
	b.l.Info("starting Batcher")

	b.shutdownCtx, b.cancelShutdownCtx = context.WithCancel(context.Background())
	b.killCtx, b.cancelKillCtx = context.WithCancel(context.Background())

	b.wg.Add(1)
	go b.loop()

	b.l.Info("Batcher started")
}

func (b *Batcher) Stop(ctx context.Context) {
	b.l.Info("stopping Batcher")

	// go routine will call cancelKillCtx() if the passed in ctx is ever Done
	cancelKill := b.cancelKillCtx
	wrapped, cancel := context.WithCancel(ctx)
	defer cancel()
	go func() {
		<-wrapped.Done()
		cancelKill()
	}()

	b.cancelShutdownCtx()
	b.wg.Wait()
	b.cancelKillCtx()

	b.l.Info("Batcher stopped")
}

// The following things occur:
// New L2 block (reorg or not)
// L1 transaction is confirmed
//
// What the batcher does:
// Ensure that channels are created & submitted as frames for an L2 range
//
// Error conditions:
// Submitted batch, but it is not valid
// Missed L2 block somehow.

func (b *Batcher) loop() {
	defer b.wg.Done()

	ticker := time.NewTicker(b.cfg.PollInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			b.batchSubmitter.LoadBlocksIntoState(b.shutdownCtx)
			if err := b.submitBatch(b.killCtx); err != nil {
				b.l.Error("failed to submit batch channel frame", "err", err)
			}
		case <-b.shutdownCtx.Done():
			if err := b.submitBatch(b.killCtx); err != nil {
				b.l.Error("failed to submit batch channel frame", "err", err)
			}
			return
		}
	}
}

// submitBatch loops through the block data loaded into `state` and
// submits the associated data to the L1 in the form of channel frames.
func (b *Batcher) submitBatch(ctx context.Context) error {
	for {
		// Attempt to gracefully terminate the current channel, ensuring that no new frames will be
		// produced. Any remaining frames must still be published to the L1 to prevent stalling.
		select {
		case <-ctx.Done():
			err := b.batchSubmitter.state.Close()
			if err != nil {
				b.l.Error("failed to close the channel manager", "err", err)
			}
		case <-b.shutdownCtx.Done():
			err := b.batchSubmitter.state.Close()
			if err != nil {
				b.l.Error("failed to close the channel manager", "err", err)
			}
		default:
		}

		l1tip, err := b.batchSubmitter.l1Tip(ctx)
		if err != nil {
			b.l.Error("failed to query L1 tip", "err", err)
			break
		}
		b.batchSubmitter.recordL1Tip(l1tip)

		// Collect next transaction data
		txdata, err := b.batchSubmitter.state.TxData(l1tip.ID())
		if err == io.EOF {
			b.l.Trace("no transaction data available")
			break
		} else if err != nil {
			b.l.Error("unable to get tx data", "err", err)
			break
		}

		// Record TX Status
		receipt, err := b.sendTransaction(ctx, txdata.Bytes())
		if err != nil {
			b.batchSubmitter.recordFailedTx(txdata.ID(), err)
			return fmt.Errorf("failed to send batch submit transaction: %w", err)
		}
		b.batchSubmitter.recordConfirmedTx(txdata.ID(), receipt)
	}

	return nil
}

// sendTransaction creates & submits a transaction to the batch inbox address with the given `data`.
// It currently uses the underlying `txmgr` to handle transaction sending & price management.
// This is a blocking method. It should not be called concurrently.
func (b *Batcher) sendTransaction(ctx context.Context, data []byte) (*types.Receipt, error) {
	// Do the gas estimation offline. A value of 0 will cause the [txmgr] to estimate the gas limit.
	intrinsicGas, err := core.IntrinsicGas(data, nil, false, true, true, false)
	if err != nil {
		return nil, fmt.Errorf("failed to calculate intrinsic gas: %w", err)
	}

	// Send the transaction through the txmgr
	receipt, err := b.cfg.TxManager.Send(ctx, txmgr.TxCandidate{
		To:       &b.batchSubmitter.Rollup.BatchInboxAddress,
		TxData:   data,
		GasLimit: intrinsicGas,
	})
	if err != nil {
		b.l.Error("batcher unable to publish tx", "err", err)
		return nil, err
	}

	// The transaction was successfully submitted
	b.l.Info("batcher tx successfully published", "tx_hash", receipt.TxHash)
	return receipt, nil
}
