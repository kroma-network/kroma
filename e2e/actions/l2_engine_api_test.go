package actions

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"github.com/stretchr/testify/require"

	"github.com/kroma-network/kroma/components/node/eth"
	"github.com/kroma-network/kroma/components/node/rollup/derive"
	"github.com/kroma-network/kroma/components/node/testlog"
)

var (
	gasLimit     = eth.Uint64Quantity(30_000_000)
	feeRecipient = common.Address{}
)

func RunEngineAPITests(t *testing.T, createBackend func() EngineBackend) {
	t.Run("CreateBlock", func(t *testing.T) {
		api := newTestHelper(t, createBackend)

		block := api.addBlock()
		api.assert.Equal(block.BlockHash, api.headHash(), "should create and import new block")
	})

	t.Run("IncludeRequiredTransactions", func(t *testing.T) {
		api := newTestHelper(t, createBackend)
		genesis := api.backend.CurrentHeader()

		txData, err := derive.L1InfoDeposit(1, eth.HeaderBlockInfo(genesis), eth.SystemConfig{})
		api.assert.NoError(err)
		tx := types.NewTx(txData)
		block := api.addBlock(tx)
		api.assert.Equal(block.BlockHash, api.headHash(), "should create and import new block")
		imported := api.backend.GetBlockByHash(block.BlockHash)
		api.assert.Len(imported.Transactions(), 1, "should include transaction")

		api.assert.NotEqual(genesis.Root, block.StateRoot)
		newState, err := api.backend.StateAt(common.Hash(block.StateRoot))
		require.NoError(t, err, "imported block state should be available")
		require.NotNil(t, newState)
	})

	t.Run("RejectCreatingBlockWithInvalidRequiredTransaction", func(t *testing.T) {
		api := newTestHelper(t, createBackend)
		genesis := api.backend.CurrentHeader()

		txData, err := derive.L1InfoDeposit(1, eth.HeaderBlockInfo(genesis), eth.SystemConfig{})
		api.assert.NoError(err)
		txData.Gas = uint64(gasLimit + 1)
		tx := types.NewTx(txData)
		txRlp, err := tx.MarshalBinary()
		api.assert.NoError(err)

		result, err := api.engine.ForkchoiceUpdatedV1(api.ctx, &eth.ForkchoiceState{
			HeadBlockHash:      genesis.Hash(),
			SafeBlockHash:      genesis.Hash(),
			FinalizedBlockHash: genesis.Hash(),
		}, &eth.PayloadAttributes{
			Timestamp:             eth.Uint64Quantity(genesis.Time + 1),
			PrevRandao:            eth.Bytes32(genesis.MixDigest),
			SuggestedFeeRecipient: feeRecipient,
			Transactions:          []eth.Data{txRlp},
			NoTxPool:              true,
			GasLimit:              &gasLimit,
		})
		api.assert.Error(err)
		api.assert.Equal(eth.ExecutionInvalid, result.PayloadStatus.Status)
	})

	t.Run("IgnoreUpdateHeadToOlderBlock", func(t *testing.T) {
		api := newTestHelper(t, createBackend)
		genesisHash := api.headHash()
		api.addBlock()
		block := api.addBlock()
		api.assert.Equal(block.BlockHash, api.headHash(), "should have extended chain")

		api.forkChoiceUpdated(genesisHash, genesisHash, genesisHash)
		api.assert.Equal(block.BlockHash, api.headHash(), "should not have reset chain head")
	})

	t.Run("AllowBuildingOnOlderBlock", func(t *testing.T) {
		api := newTestHelper(t, createBackend)
		genesis := api.backend.CurrentHeader()
		api.addBlock()
		block := api.addBlock()
		api.assert.Equal(block.BlockHash, api.headHash(), "should have extended chain")

		payloadID := api.startBlockBuilding(genesis, eth.Uint64Quantity(genesis.Time+3))
		api.assert.Equal(block.BlockHash, api.headHash(), "should not reset chain head when building starts")

		payload := api.getPayload(payloadID)
		api.assert.Equal(genesis.Hash(), payload.ParentHash, "should have old block as parent")

		api.newPayload(payload)
		api.forkChoiceUpdated(payload.BlockHash, genesis.Hash(), genesis.Hash())
		api.assert.Equal(payload.BlockHash, api.headHash(), "should reorg to block built on old parent")
	})

	t.Run("RejectInvalidBlockHash", func(t *testing.T) {
		api := newTestHelper(t, createBackend)

		// Invalid because BlockHash won't be correct (among many other reasons)
		block := &eth.ExecutionPayload{}
		r, err := api.engine.NewPayloadV1(api.ctx, block)
		api.assert.NoError(err)
		api.assert.Equal(eth.ExecutionInvalidBlockHash, r.Status)
	})

	t.Run("RejectBlockWithInvalidStateTransition", func(t *testing.T) {
		api := newTestHelper(t, createBackend)
		genesis := api.backend.CurrentHeader()

		// Build a valid block
		payloadID := api.startBlockBuilding(genesis, eth.Uint64Quantity(genesis.Time+2))
		newBlock := api.getPayload(payloadID)

		// But then make it invalid by changing the state root
		newBlock.StateRoot = eth.Bytes32(genesis.TxHash)
		updateBlockHash(newBlock)

		r, err := api.engine.NewPayloadV1(api.ctx, newBlock)
		api.assert.NoError(err)
		api.assert.Equal(eth.ExecutionInvalid, r.Status)
	})

	t.Run("RejectBlockWithSameTimeAsParent", func(t *testing.T) {
		api := newTestHelper(t, createBackend)
		genesis := api.backend.CurrentHeader()

		// Start with a valid time
		payloadID := api.startBlockBuilding(genesis, eth.Uint64Quantity(genesis.Time+1))
		newBlock := api.getPayload(payloadID)

		// Then make it invalid to check NewPayload rejects it
		newBlock.Timestamp = eth.Uint64Quantity(genesis.Time)
		updateBlockHash(newBlock)

		r, err := api.engine.NewPayloadV1(api.ctx, newBlock)
		api.assert.NoError(err)
		api.assert.Equal(eth.ExecutionInvalid, r.Status)
	})

	t.Run("RejectBlockWithTimeBeforeParent", func(t *testing.T) {
		api := newTestHelper(t, createBackend)
		genesis := api.backend.CurrentHeader()

		// Start with a valid time
		payloadID := api.startBlockBuilding(genesis, eth.Uint64Quantity(genesis.Time+1))
		newBlock := api.getPayload(payloadID)

		// Then make it invalid to check NewPayload rejects it
		newBlock.Timestamp = eth.Uint64Quantity(genesis.Time - 1)
		updateBlockHash(newBlock)

		r, err := api.engine.NewPayloadV1(api.ctx, newBlock)
		api.assert.NoError(err)
		api.assert.Equal(eth.ExecutionInvalid, r.Status)
	})

	t.Run("RejectCreateBlockWithSameTimeAsParent", func(t *testing.T) {
		api := newTestHelper(t, createBackend)
		genesis := api.backend.CurrentHeader()

		result, err := api.engine.ForkchoiceUpdatedV1(api.ctx, &eth.ForkchoiceState{
			HeadBlockHash:      genesis.Hash(),
			SafeBlockHash:      genesis.Hash(),
			FinalizedBlockHash: genesis.Hash(),
		}, &eth.PayloadAttributes{
			Timestamp:             eth.Uint64Quantity(genesis.Time),
			PrevRandao:            eth.Bytes32(genesis.MixDigest),
			SuggestedFeeRecipient: feeRecipient,
			Transactions:          nil,
			NoTxPool:              true,
			GasLimit:              &gasLimit,
		})
		api.assert.Error(err)
		api.assert.Equal(eth.ExecutionInvalid, result.PayloadStatus.Status)
	})

	t.Run("RejectCreateBlockWithTimeBeforeParent", func(t *testing.T) {
		api := newTestHelper(t, createBackend)
		genesis := api.backend.CurrentHeader()

		result, err := api.engine.ForkchoiceUpdatedV1(api.ctx, &eth.ForkchoiceState{
			HeadBlockHash:      genesis.Hash(),
			SafeBlockHash:      genesis.Hash(),
			FinalizedBlockHash: genesis.Hash(),
		}, &eth.PayloadAttributes{
			Timestamp:             eth.Uint64Quantity(genesis.Time - 1),
			PrevRandao:            eth.Bytes32(genesis.MixDigest),
			SuggestedFeeRecipient: feeRecipient,
			Transactions:          nil,
			NoTxPool:              true,
			GasLimit:              &gasLimit,
		})
		api.assert.Error(err)
		api.assert.Equal(eth.ExecutionInvalid, result.PayloadStatus.Status)
	})

	t.Run("RejectCreateBlockWithGasLimitAboveMax", func(t *testing.T) {
		api := newTestHelper(t, createBackend)
		genesis := api.backend.CurrentHeader()

		gasLimit := eth.Uint64Quantity(params.MaxGasLimit + 1)

		result, err := api.engine.ForkchoiceUpdatedV1(api.ctx, &eth.ForkchoiceState{
			HeadBlockHash:      genesis.Hash(),
			SafeBlockHash:      genesis.Hash(),
			FinalizedBlockHash: genesis.Hash(),
		}, &eth.PayloadAttributes{
			Timestamp:             eth.Uint64Quantity(genesis.Time + 1),
			PrevRandao:            eth.Bytes32(genesis.MixDigest),
			SuggestedFeeRecipient: feeRecipient,
			Transactions:          nil,
			NoTxPool:              true,
			GasLimit:              &gasLimit,
		})
		api.assert.Error(err)
		api.assert.Equal(eth.ExecutionInvalid, result.PayloadStatus.Status)
	})

	t.Run("UpdateSafeAndFinalizedHead", func(t *testing.T) {
		api := newTestHelper(t, createBackend)

		finalized := api.addBlock()
		safe := api.addBlock()
		head := api.addBlock()

		api.forkChoiceUpdated(head.BlockHash, safe.BlockHash, finalized.BlockHash)
		api.assert.Equal(head.BlockHash, api.headHash(), "should update head block")
		api.assert.Equal(safe.BlockHash, api.safeHash(), "should update safe block")
		api.assert.Equal(finalized.BlockHash, api.finalHash(), "should update finalized block")
	})

	t.Run("RejectSafeHeadWhenNotAncestor", func(t *testing.T) {
		api := newTestHelper(t, createBackend)
		genesis := api.backend.CurrentHeader()

		api.addBlock()
		chainA2 := api.addBlock()
		chainA3 := api.addBlock()

		chainB1 := api.addBlockWithParent(genesis, eth.Uint64Quantity(genesis.Time+3))

		result, err := api.engine.ForkchoiceUpdatedV1(api.ctx, &eth.ForkchoiceState{
			HeadBlockHash:      chainA3.BlockHash,
			SafeBlockHash:      chainB1.BlockHash,
			FinalizedBlockHash: chainA2.BlockHash,
		}, nil)
		api.assert.ErrorContains(err, "Invalid forkchoice state", "should return error from forkChoiceUpdated")
		api.assert.Equal(eth.ExecutionInvalid, result.PayloadStatus.Status, "forkChoiceUpdated should return invalid")
		api.assert.Nil(result.PayloadID, "should not provide payload ID when invalid")
	})

	t.Run("RejectFinalizedHeadWhenNotAncestor", func(t *testing.T) {
		api := newTestHelper(t, createBackend)
		genesis := api.backend.CurrentHeader()

		api.addBlock()
		chainA2 := api.addBlock()
		chainA3 := api.addBlock()

		chainB1 := api.addBlockWithParent(genesis, eth.Uint64Quantity(genesis.Time+3))

		result, err := api.engine.ForkchoiceUpdatedV1(api.ctx, &eth.ForkchoiceState{
			HeadBlockHash:      chainA3.BlockHash,
			SafeBlockHash:      chainA2.BlockHash,
			FinalizedBlockHash: chainB1.BlockHash,
		}, nil)
		api.assert.ErrorContains(err, "Invalid forkchoice state", "should return error from forkChoiceUpdated")
		api.assert.Equal(eth.ExecutionInvalid, result.PayloadStatus.Status, "forkChoiceUpdated should return invalid")
		api.assert.Nil(result.PayloadID, "should not provide payload ID when invalid")
	})
}

// Updates the block hash to the expected value based on the other fields in the payload
func updateBlockHash(newBlock *eth.ExecutionPayload) {
	// And fix up the block hash
	newHash, _ := newBlock.CheckBlockHash()
	newBlock.BlockHash = newHash
}

type testHelper struct {
	t       *testing.T
	ctx     context.Context
	engine  *L2EngineAPI
	backend EngineBackend
	assert  *require.Assertions
}

func newTestHelper(t *testing.T, createBackend func() EngineBackend) *testHelper {
	logger := testlog.Logger(t, log.LvlDebug)
	ctx := context.Background()
	backend := createBackend()
	api := NewL2EngineAPI(logger, backend)
	test := &testHelper{
		t:       t,
		ctx:     ctx,
		engine:  api,
		backend: backend,
		assert:  require.New(t),
	}
	return test
}

func (h *testHelper) headHash() common.Hash {
	return h.backend.CurrentHeader().Hash()
}

func (h *testHelper) safeHash() common.Hash {
	return h.backend.CurrentSafeBlock().Hash()
}

func (h *testHelper) finalHash() common.Hash {
	return h.backend.CurrentFinalBlock().Hash()
}

func (h *testHelper) Log(args ...any) {
	h.t.Log(args...)
}

func (h *testHelper) addBlock(txs ...*types.Transaction) *eth.ExecutionPayload {
	head := h.backend.CurrentHeader()
	return h.addBlockWithParent(head, eth.Uint64Quantity(head.Time+2), txs...)
}

func (h *testHelper) addBlockWithParent(head *types.Header, timestamp eth.Uint64Quantity, txs ...*types.Transaction) *eth.ExecutionPayload {
	prevHead := h.backend.CurrentHeader()
	id := h.startBlockBuilding(head, timestamp, txs...)

	block := h.getPayload(id)
	h.assert.Equal(timestamp, block.Timestamp, "should create block with correct timestamp")
	h.assert.Equal(head.Hash(), block.ParentHash, "should have correct parent")
	h.assert.Len(block.Transactions, len(txs))

	h.newPayload(block)

	// Should not have changed the chain head yet
	h.assert.Equal(prevHead, h.backend.CurrentHeader())

	h.forkChoiceUpdated(block.BlockHash, head.Hash(), head.Hash())
	h.assert.Equal(block.BlockHash, h.backend.CurrentHeader().Hash())
	return block
}

func (h *testHelper) forkChoiceUpdated(head common.Hash, safe common.Hash, finalized common.Hash) {
	h.Log("forkChoiceUpdated", "head", head, "safe", safe, "finalized", finalized)
	result, err := h.engine.ForkchoiceUpdatedV1(h.ctx, &eth.ForkchoiceState{
		HeadBlockHash:      head,
		SafeBlockHash:      safe,
		FinalizedBlockHash: finalized,
	}, nil)
	h.assert.NoError(err)
	h.assert.Equal(eth.ExecutionValid, result.PayloadStatus.Status, "forkChoiceUpdated should return valid")
	h.assert.Nil(result.PayloadStatus.ValidationError, "should not have validation error when valid")
	h.assert.Nil(result.PayloadID, "should not provide payload ID when block building not requested")
}

func (h *testHelper) startBlockBuilding(head *types.Header, newBlockTimestamp eth.Uint64Quantity, txs ...*types.Transaction) *eth.PayloadID {
	h.Log("Start block building", "head", head.Hash(), "timestamp", newBlockTimestamp)
	var txData []eth.Data
	for _, tx := range txs {
		rlp, err := tx.MarshalBinary()
		h.assert.NoError(err, "Failed to marshall tx %v", tx)
		txData = append(txData, rlp)
	}
	result, err := h.engine.ForkchoiceUpdatedV1(h.ctx, &eth.ForkchoiceState{
		HeadBlockHash:      head.Hash(),
		SafeBlockHash:      head.Hash(),
		FinalizedBlockHash: head.Hash(),
	}, &eth.PayloadAttributes{
		Timestamp:             newBlockTimestamp,
		PrevRandao:            eth.Bytes32(head.MixDigest),
		SuggestedFeeRecipient: feeRecipient,
		Transactions:          txData,
		NoTxPool:              true,
		GasLimit:              &gasLimit,
	})
	h.assert.NoError(err)
	h.assert.Equal(eth.ExecutionValid, result.PayloadStatus.Status)
	id := result.PayloadID
	h.assert.NotNil(id)
	return id
}

func (h *testHelper) getPayload(id *eth.PayloadID) *eth.ExecutionPayload {
	h.Log("getPayload", "id", id)
	block, err := h.engine.GetPayloadV1(h.ctx, *id)
	h.assert.NoError(err)
	h.assert.NotNil(block)
	return block
}

func (h *testHelper) newPayload(block *eth.ExecutionPayload) {
	h.Log("newPayload", "hash", block.BlockHash)
	r, err := h.engine.NewPayloadV1(h.ctx, block)
	h.assert.NoError(err)
	h.assert.Equal(eth.ExecutionValid, r.Status)
	h.assert.Nil(r.ValidationError)
}
