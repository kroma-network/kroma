package batcher

import (
	"errors"
	"fmt"
	"io"
	"math"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"

	"github.com/kroma-network/kroma/components/batcher/metrics"
	"github.com/kroma-network/kroma/components/node/eth"
	"github.com/kroma-network/kroma/components/node/rollup/derive"
)

var ErrReorg = errors.New("block does not extend existing chain")

// channelManager stores a contiguous set of blocks & turns them into channels.
// Upon receiving tx confirmation (or a tx failure), it does channel error handling.
//
// For simplicity, it only creates a single pending channel at a time & waits for
// the channel to either successfully be submitted or timeout before creating a new
// channel.
// Functions on channelManager are not safe for concurrent access.
type channelManager struct {
	log  log.Logger
	metr metrics.Metricer
	cfg  ChannelConfig

	// All blocks since the last request for new tx data.
	blocks []*types.Block
	// last block hash - for reorg detection
	tip common.Hash

	// Pending data returned by TxData waiting on Tx Confirmed/Failed

	// pending channel builder
	pendingChannel *channelBuilder
	// Set of unconfirmed txID -> frame data. For tx resubmission
	pendingTransactions map[txID]txData
	// Set of confirmed txID -> inclusion block. For determining if the channel is timed out
	confirmedTransactions map[txID]eth.BlockID

	// if set to true, prevents production of any new channel frames
	closed bool
}

func NewChannelManager(log log.Logger, metr metrics.Metricer, cfg ChannelConfig) *channelManager {
	return &channelManager{
		log:  log,
		metr: metr,
		cfg:  cfg,

		pendingTransactions:   make(map[txID]txData),
		confirmedTransactions: make(map[txID]eth.BlockID),
	}
}

// Clear clears the entire state of the channel manager.
// It is intended to be used after an L2 reorg.
func (c *channelManager) Clear() {
	c.log.Trace("clearing channel manager state")
	c.blocks = c.blocks[:0]
	c.tip = common.Hash{}
	c.closed = false
	c.clearPendingChannel()
}

// TxFailed records a transaction as failed. It will attempt to resubmit the data
// in the failed transaction.
func (c *channelManager) TxFailed(id txID) {
	if data, ok := c.pendingTransactions[id]; ok {
		c.log.Trace("marked transaction as failed", "id", id)
		// Note: when the batcher is changed to send multiple frames per tx,
		// this needs to be changed to iterate over all frames of the tx data
		// and re-queue them.
		c.pendingChannel.PushFrame(data.Frame())
		delete(c.pendingTransactions, id)
	} else {
		c.log.Warn("unknown transaction marked as failed", "id", id)
	}

	c.metr.RecordBatchTxFailed()
	if c.closed && len(c.confirmedTransactions) == 0 && len(c.pendingTransactions) == 0 {
		c.log.Info("Channel has no submitted transactions, clearing for shutdown", "chID", c.pendingChannel.ID())
		c.clearPendingChannel()
	}
}

// TxConfirmed marks a transaction as confirmed on L1. Unfortunately even if all frames in
// a channel have been marked as confirmed on L1 the channel may be invalid & need to be
// resubmitted.
// This function may reset the pending channel if the pending channel has timed out.
func (c *channelManager) TxConfirmed(id txID, inclusionBlock eth.BlockID) {
	c.metr.RecordBatchTxSubmitted()
	c.log.Debug("marked transaction as confirmed", "id", id, "block", inclusionBlock)
	if _, ok := c.pendingTransactions[id]; !ok {
		c.log.Warn("unknown transaction marked as confirmed", "id", id, "block", inclusionBlock)
		// TODO: This can occur if we clear the channel while there are still pending transactions
		// We need to keep track of stale transactions instead
		return
	}
	delete(c.pendingTransactions, id)
	c.confirmedTransactions[id] = inclusionBlock
	c.pendingChannel.FramePublished(inclusionBlock.Number)

	// If this channel timed out, put the pending blocks back into the local saved blocks
	// and then reset this state so it can try to build a new channel.
	if c.pendingChannelIsTimedOut() {
		c.metr.RecordChannelTimedOut(c.pendingChannel.ID())
		c.log.Warn("Channel timed out", "id", c.pendingChannel.ID())
		c.blocks = append(c.pendingChannel.Blocks(), c.blocks...)
		c.clearPendingChannel()
	}
	// If we are done with this channel, record that.
	if c.pendingChannelIsFullySubmitted() {
		c.metr.RecordChannelFullySubmitted(c.pendingChannel.ID())
		c.log.Info("Channel is fully submitted", "id", c.pendingChannel.ID())
		c.clearPendingChannel()
	}
}

// clearPendingChannel resets all pending state back to an initialized but empty state.
// TODO: Create separate "pending" state
func (c *channelManager) clearPendingChannel() {
	c.pendingChannel = nil
	c.pendingTransactions = make(map[txID]txData)
	c.confirmedTransactions = make(map[txID]eth.BlockID)
}

// pendingChannelIsTimedOut returns true if submitted channel has timed out.
// A channel has timed out if the difference in L1 Inclusion blocks between
// the first & last included block is greater than or equal to the channel timeout.
func (c *channelManager) pendingChannelIsTimedOut() bool {
	if c.pendingChannel == nil {
		return false // no channel to be timed out
	}
	// No confirmed transactions => not timed out
	if len(c.confirmedTransactions) == 0 {
		return false
	}
	// If there are confirmed transactions, find the first + last confirmed block numbers
	min := uint64(math.MaxUint64)
	max := uint64(0)
	for _, inclusionBlock := range c.confirmedTransactions {
		if inclusionBlock.Number < min {
			min = inclusionBlock.Number
		}
		if inclusionBlock.Number > max {
			max = inclusionBlock.Number
		}
	}
	return max-min >= c.cfg.ChannelTimeout
}

// pendingChannelIsFullySubmitted returns true if the channel has been fully submitted.
func (c *channelManager) pendingChannelIsFullySubmitted() bool {
	if c.pendingChannel == nil {
		return false // todo: can decide either way here. Nonsensical answer though
	}
	return c.pendingChannel.IsFull() && len(c.pendingTransactions)+c.pendingChannel.NumFrames() == 0
}

// nextTxData pops off c.datas & handles updating the internal state
func (c *channelManager) nextTxData() (txData, error) {
	if c.pendingChannel == nil || !c.pendingChannel.HasFrame() {
		c.log.Trace("no next tx data")
		return txData{}, io.EOF // TODO: not enough data error instead
	}

	frame := c.pendingChannel.NextFrame()
	txdata := txData{frame}
	id := txdata.ID()

	c.log.Trace("returning next tx data", "id", id)
	c.pendingTransactions[id] = txdata
	return txdata, nil
}

// TxData returns the next tx data that should be submitted to L1.
//
// It currently only uses one frame per transaction. If the pending channel is
// full, it only returns the remaining frames of this channel until it got
// successfully fully sent to L1. It returns io.EOF if there's no pending frame.
func (c *channelManager) TxData(l1Head eth.BlockID) (txData, error) {
	dataPending := c.pendingChannel != nil && c.pendingChannel.HasFrame()
	c.log.Debug("Requested tx data", "l1Head", l1Head, "data_pending", dataPending, "blocks_pending", len(c.blocks))

	// Short circuit if there is a pending frame or the channel manager is closed.
	if dataPending || c.closed {
		return c.nextTxData()
	}

	// No pending frame, so we have to add new blocks to the channel

	// If we have no saved blocks, we will not be able to create valid frames
	if len(c.blocks) == 0 {
		return txData{}, io.EOF
	}

	if err := c.ensurePendingChannel(l1Head); err != nil {
		return txData{}, err
	}

	if err := c.processBlocks(); err != nil {
		return txData{}, err
	}

	// Register current L1 head only after all pending blocks have been
	// processed. Even if a timeout will be triggered now, it is better to have
	// all pending blocks be included in this channel for submission.
	c.registerL1Block(l1Head)

	if err := c.outputFrames(); err != nil {
		return txData{}, err
	}

	return c.nextTxData()
}

func (c *channelManager) ensurePendingChannel(l1Head eth.BlockID) error {
	if c.pendingChannel != nil {
		return nil
	}

	cb, err := newChannelBuilder(c.cfg)
	if err != nil {
		return fmt.Errorf("creating new channel: %w", err)
	}
	c.pendingChannel = cb
	c.log.Info("Created channel",
		"id", cb.ID(),
		"l1Head", l1Head,
		"blocks_pending", len(c.blocks))
	c.metr.RecordChannelOpened(cb.ID(), len(c.blocks))

	return nil
}

// registerL1Block registers the given block at the pending channel.
func (c *channelManager) registerL1Block(l1Head eth.BlockID) {
	c.pendingChannel.RegisterL1Block(l1Head.Number)
	c.log.Debug("new L1-block registered at channel builder",
		"l1Head", l1Head,
		"channel_full", c.pendingChannel.IsFull(),
		"full_reason", c.pendingChannel.FullErr(),
	)
}

// processBlocks adds blocks from the blocks queue to the pending channel until
// either the queue got exhausted or the channel is full.
func (c *channelManager) processBlocks() error {
	var (
		blocksAdded int
		_chFullErr  *ChannelFullError // throw away, just for type checking
		latestL2ref eth.L2BlockRef
	)
	for i, block := range c.blocks {
		l1info, err := c.pendingChannel.AddBlock(block)
		if errors.As(err, &_chFullErr) {
			// current block didn't get added because channel is already full
			break
		} else if err != nil {
			return fmt.Errorf("adding block[%d] to channel builder: %w", i, err)
		}
		blocksAdded += 1
		latestL2ref = l2BlockRefFromBlockAndL1Info(block, l1info)
		// current block got added but channel is now full
		if c.pendingChannel.IsFull() {
			break
		}
	}

	if blocksAdded == len(c.blocks) {
		// all blocks processed, reuse slice
		c.blocks = c.blocks[:0]
	} else {
		// remove processed blocks
		c.blocks = c.blocks[blocksAdded:]
	}

	c.metr.RecordL2BlocksAdded(latestL2ref,
		blocksAdded,
		len(c.blocks),
		c.pendingChannel.InputBytes(),
		c.pendingChannel.ReadyBytes())
	c.log.Debug("Added blocks to channel",
		"blocks_added", blocksAdded,
		"blocks_pending", len(c.blocks),
		"channel_full", c.pendingChannel.IsFull(),
		"input_bytes", c.pendingChannel.InputBytes(),
		"ready_bytes", c.pendingChannel.ReadyBytes(),
	)
	return nil
}

func (c *channelManager) outputFrames() error {
	if err := c.pendingChannel.OutputFrames(); err != nil {
		return fmt.Errorf("creating frames with channel builder: %w", err)
	}
	if !c.pendingChannel.IsFull() {
		return nil
	}

	inBytes, outBytes := c.pendingChannel.InputBytes(), c.pendingChannel.OutputBytes()
	c.metr.RecordChannelClosed(
		c.pendingChannel.ID(),
		len(c.blocks),
		c.pendingChannel.NumFrames(),
		inBytes,
		outBytes,
		c.pendingChannel.FullErr(),
	)

	var comprRatio float64
	if inBytes > 0 {
		comprRatio = float64(outBytes) / float64(inBytes)
	}
	c.log.Info("Channel closed",
		"id", c.pendingChannel.ID(),
		"blocks_pending", len(c.blocks),
		"num_frames", c.pendingChannel.NumFrames(),
		"input_bytes", inBytes,
		"output_bytes", outBytes,
		"full_reason", c.pendingChannel.FullErr(),
		"compr_ratio", comprRatio,
	)
	return nil
}

// AddL2Block adds an L2 block to the internal blocks queue. It returns ErrReorg
// if the block does not extend the last block loaded into the state. If no
// blocks were added yet, the parent hash check is skipped.
func (c *channelManager) AddL2Block(block *types.Block) error {
	if c.tip != (common.Hash{}) && c.tip != block.ParentHash() {
		return ErrReorg
	}
	c.blocks = append(c.blocks, block)
	c.tip = block.Hash()

	return nil
}

func l2BlockRefFromBlockAndL1Info(block *types.Block, l1info derive.L1BlockInfo) eth.L2BlockRef {
	return eth.L2BlockRef{
		Hash:           block.Hash(),
		Number:         block.NumberU64(),
		ParentHash:     block.ParentHash(),
		Time:           block.Time(),
		L1Origin:       eth.BlockID{Hash: l1info.BlockHash, Number: l1info.Number},
		SequenceNumber: l1info.SequenceNumber,
	}
}

// Close closes the current pending channel, if one exists, outputs any remaining frames,
// and prevents the creation of any new channels.
// Any outputted frames still need to be published.
func (c *channelManager) Close() error {
	if c.closed {
		return nil
	}

	c.closed = true

	// Any pending state can be proactively cleared if there are no submitted transactions
	if len(c.confirmedTransactions) == 0 && len(c.pendingTransactions) == 0 {
		c.clearPendingChannel()
	}

	if c.pendingChannel == nil {
		return nil
	}

	c.pendingChannel.Close()

	return c.outputFrames()
}
