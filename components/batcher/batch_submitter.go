package batcher

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	_ "net/http/pprof"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"

	"github.com/kroma-network/kroma/components/batcher/metrics"
	"github.com/kroma-network/kroma/components/node/eth"
)

// BatchSubmitter encapsulates a service responsible for submitting L2 tx
// batches to L1 for availability.
type BatchSubmitter struct {
	Config // directly embed the config

	// lastStoredBlock is the last block loaded into `state`. If it is empty it should be set to the l2 safe head.
	lastStoredBlock eth.BlockID
	lastL1Tip       eth.L1BlockRef

	state *channelManager
}

// NewBatchSubmitter initializes the BatchSubmitter, gathering any resources
// that will be needed during operation.
func NewBatchSubmitter(cfg Config, l log.Logger, m metrics.Metricer) (*BatchSubmitter, error) {
	return &BatchSubmitter{
		Config: cfg,
		state:  NewChannelManager(l, m, cfg.Channel),
	}, nil
}

// LoadBlocksIntoState loads all blocks since the previous stored block
// It does the following:
// 1. Fetch the sync status of the proposer
// 2. Check if the sync status is valid or if we are all the way up to date
// 3. Check if it needs to initialize state OR it is lagging (todo: lagging just means race condition?)
// 4. Load all new blocks into the local state.
func (b *BatchSubmitter) LoadBlocksIntoState(ctx context.Context) {
	start, end, err := b.calculateL2BlockRangeToStore(ctx)
	if err != nil {
		b.log.Trace("unable to calculate L2 block range", "err", err)
		return
	}

	// Add all blocks to "state"
	for i := start.Number + 1; i < end.Number+1; i++ {
		id, err := b.loadBlockIntoState(ctx, i)
		if errors.Is(err, ErrReorg) {
			b.log.Warn("found L2 reorg", "block_number", i)
			b.state.Clear()
			b.lastStoredBlock = eth.BlockID{}
			return
		} else if err != nil {
			b.log.Warn("failed to load block into state", "err", err)
			return
		}
		b.lastStoredBlock = id
	}
}

// loadBlockIntoState fetches & stores a single block into `state`. It returns the block it loaded.
func (b *BatchSubmitter) loadBlockIntoState(ctx context.Context, blockNumber uint64) (eth.BlockID, error) {
	ctx, cancel := context.WithTimeout(ctx, b.NetworkTimeout)
	block, err := b.L2Client.BlockByNumber(ctx, new(big.Int).SetUint64(blockNumber))
	cancel()
	if err != nil {
		return eth.BlockID{}, err
	}
	if err := b.state.AddL2Block(block); err != nil {
		return eth.BlockID{}, err
	}
	id := eth.ToBlockID(block)
	b.log.Info("added L2 block to local state", "block", id, "tx_count", len(block.Transactions()), "time", block.Time())
	return id, nil
}

// calculateL2BlockRangeToStore determines the range (start,end) that should be loaded into the local state.
// It also takes care of initializing some local state (i.e. will modify b.lastStoredBlock in certain conditions)
func (b *BatchSubmitter) calculateL2BlockRangeToStore(ctx context.Context) (eth.BlockID, eth.BlockID, error) {
	ctx, cancel := context.WithTimeout(ctx, b.NetworkTimeout)
	defer cancel()
	syncStatus, err := b.RollupClient.SyncStatus(ctx)
	// Ensure that we have the sync status
	if err != nil {
		return eth.BlockID{}, eth.BlockID{}, fmt.Errorf("failed to get sync status: %w", err)
	}
	if syncStatus.HeadL1 == (eth.L1BlockRef{}) {
		return eth.BlockID{}, eth.BlockID{}, errors.New("empty sync status")
	}

	// Check last stored to see if it needs to be set on startup OR set if is lagged behind.
	// It lagging implies that the kroma-node processed some batches that where submitted prior to the current instance of the kroma-batcher being alive.
	if b.lastStoredBlock == (eth.BlockID{}) {
		b.log.Info("Starting batch-submitter work at safe-head", "safe", syncStatus.SafeL2)
		b.lastStoredBlock = syncStatus.SafeL2.ID()
	} else if b.lastStoredBlock.Number < syncStatus.SafeL2.Number {
		b.log.Warn("last submitted block lagged behind L2 safe head: batch submission will continue from the safe head now", "last", b.lastStoredBlock, "safe", syncStatus.SafeL2)
		b.lastStoredBlock = syncStatus.SafeL2.ID()
	}

	// Check if we should even attempt to load any blocks. TODO: May not need this check
	if syncStatus.SafeL2.Number >= syncStatus.UnsafeL2.Number {
		return eth.BlockID{}, eth.BlockID{}, errors.New("L2 safe head ahead of L2 unsafe head")
	}

	return b.lastStoredBlock, syncStatus.UnsafeL2.ID(), nil
}

func (b *BatchSubmitter) recordL1Tip(l1tip eth.L1BlockRef) {
	if b.lastL1Tip == l1tip {
		return
	}
	b.lastL1Tip = l1tip
	b.metr.RecordLatestL1Block(l1tip)
}

func (b *BatchSubmitter) recordFailedTx(id txID, err error) {
	b.log.Warn("Failed to send transaction", "err", err)
	b.state.TxFailed(id)
}

func (b *BatchSubmitter) recordConfirmedTx(id txID, receipt *types.Receipt) {
	b.log.Info("Transaction confirmed", "tx_hash", receipt.TxHash, "status", receipt.Status, "block_hash", receipt.BlockHash, "block_number", receipt.BlockNumber)
	l1block := eth.BlockID{Number: receipt.BlockNumber.Uint64(), Hash: receipt.BlockHash}
	b.state.TxConfirmed(id, l1block)
}

// l1Tip gets the current L1 tip as a L1BlockRef. The passed context is assumed
// to be a lifetime context, so it is internally wrapped with a network timeout.
func (b *BatchSubmitter) l1Tip(ctx context.Context) (eth.L1BlockRef, error) {
	tctx, cancel := context.WithTimeout(ctx, b.NetworkTimeout)
	defer cancel()
	head, err := b.L1Client.HeaderByNumber(tctx, nil)
	if err != nil {
		return eth.L1BlockRef{}, fmt.Errorf("getting latest L1 block: %w", err)
	}
	return eth.InfoToL1BlockRef(eth.HeaderBlockInfo(head)), nil
}
