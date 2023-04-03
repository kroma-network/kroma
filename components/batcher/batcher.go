package batcher

import (
	"context"
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/urfave/cli"

	"github.com/wemixkanvas/kanvas/components/batcher/metrics"
	"github.com/wemixkanvas/kanvas/utils"
	"github.com/wemixkanvas/kanvas/utils/monitoring"
	klog "github.com/wemixkanvas/kanvas/utils/service/log"
	krpc "github.com/wemixkanvas/kanvas/utils/service/rpc"
	"github.com/wemixkanvas/kanvas/utils/service/txmgr"
)

// Main is the entrypoint into the Batch Submitter.
func Main(version string, cliCtx *cli.Context) error {
	cliCfg := NewCLIConfig(cliCtx)
	if err := cliCfg.Check(); err != nil {
		return fmt.Errorf("invalid CLI flags: %w", err)
	}

	l := klog.NewLogger(cliCfg.LogConfig)
	m := metrics.NewMetrics("default")
	l.Info("Initializing Batch Submitter")

	batcherCfg, err := NewBatcherConfig(cliCfg, l, m)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	monitoring.MaybeStartPprof(ctx, cliCfg.PprofConfig, l)
	monitoring.MaybeStartMetrics(ctx, cliCfg.MetricsConfig, l, batcherCfg.L1Client, batcherCfg.From)
	server, err := monitoring.StartRPC(cliCfg.RPCConfig.ToServiceCLIConfig(), version, krpc.WithLogger(l))
	if err != nil {
		return err
	}
	defer server.Stop()

	m.RecordInfo(version)
	m.RecordUp()

	batcher, err := NewBatcher(ctx, *batcherCfg, l, m)
	if err != nil {
		return err
	}

	batcher.Start()
	<-utils.WaitInterrupt()
	batcher.Stop()

	return nil
}

type Batcher struct {
	ctx    context.Context
	cancel context.CancelFunc

	cfg            Config
	l              log.Logger
	batchSubmitter *BatchSubmitter
	txMgr          txmgr.TxManager

	wg sync.WaitGroup
}

func NewBatcher(parentCtx context.Context, cfg Config, l log.Logger, m metrics.Metricer) (*Batcher, error) {
	// Validate the batcher config
	if err := cfg.Check(); err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(parentCtx)

	batchSubmitter, err := NewBatchSubmitter(cfg, l, m)
	if err != nil {
		cancel()
		return nil, fmt.Errorf("failed to init batch submitter: %w", err)
	}

	balance, err := cfg.L1Client.BalanceAt(ctx, cfg.From, nil)
	if err != nil {
		cancel()
		return nil, err
	}

	l.Info("creating batcher", "batcher_addr", cfg.From, "batcher_bal", balance)

	return &Batcher{
		ctx:            ctx,
		cancel:         cancel,
		cfg:            cfg,
		l:              l,
		batchSubmitter: batchSubmitter,
		txMgr:          txmgr.NewSimpleTxManager("batcher", l, cfg.TxManagerConfig, cfg.L1Client),
	}, nil
}

func (b *Batcher) Start() {
	b.l.Info("starting Batch Submitter")
	b.wg.Add(1)
	go b.loop()
}

func (b *Batcher) Stop() {
	b.cancel()
	b.wg.Wait()
}

func (b *Batcher) loop() {
	defer b.wg.Done()

	ticker := time.NewTicker(b.cfg.PollInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if err := b.submitBatch(); err != nil {
				b.l.Error("failed to submit batch channel frame", "err", err)
			}
		case <-b.ctx.Done():
			return
		}
	}
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

func (b *Batcher) submitBatch() error {
	b.batchSubmitter.LoadBlocksIntoState(b.ctx)

blockLoop:
	for {
		l1tip, err := b.batchSubmitter.l1Tip(b.ctx)
		if err != nil {
			b.l.Error("Failed to query L1 tip", "error", err)
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

		tx, err := b.batchSubmitter.CreateSubmitTx(txdata.Bytes())
		if err != nil {
			// record it as a failed TX to resubmit the transaction.
			b.batchSubmitter.recordFailedTx(txdata.ID(), err)
			return fmt.Errorf("failed to create batch submit transaction: %w", err)
		}
		b.l.Info("creating batch submit tx", "to", tx.To, "from", b.cfg.From)
		// Record TX Status
		receipt, err := b.SendTransaction(b.ctx, tx)
		if err != nil {
			b.batchSubmitter.recordFailedTx(txdata.ID(), err)
			return fmt.Errorf("failed to send batch transaction: %w", err)
		}
		b.batchSubmitter.recordConfirmedTx(txdata.ID(), receipt)

		// hack to exit this loop. Proper fix is to do request another send tx or parallel tx sending
		// from the channel manager rather than sending the channel in a loop. This stalls b/c if the
		// context is cancelled while sending, it will never fully clearing the pending txns.
		select {
		case <-b.ctx.Done():
			break blockLoop
		default:
		}
	}

	return nil
}

// SendTransaction sends a transaction through the transaction manager which handles automatic
// price bumping, and returns transaction receipt.
// It also hardcodes a timeout of 100s.
func (b *Batcher) SendTransaction(ctx context.Context, tx *types.Transaction) (*types.Receipt, error) {
	// Wait until one of our submitted transactions confirms. If no
	// receipt is received it's likely our gas price was too low.
	cCtx, cancel := context.WithTimeout(ctx, 100*time.Second)
	defer cancel()

	b.l.Info("batcher sending transaction", "tx", tx.Hash())
	receipt, err := b.txMgr.Send(cCtx, tx)
	if err != nil {
		b.l.Error("batcher unable to publish tx", "err", err)
		return nil, err
	}

	// The transaction was successfully submitted
	b.l.Info("batcher tx successfully published", "tx_hash", receipt.TxHash)
	return receipt, nil
}
