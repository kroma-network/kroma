package validator

import (
	"context"
	"errors"
	"math/big"
	_ "net/http/pprof"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"

	"github.com/kroma-network/kroma/bindings/bindings"
	"github.com/kroma-network/kroma/components/node/eth"
	"github.com/kroma-network/kroma/components/node/rollup"
	"github.com/kroma-network/kroma/utils"
)

// L2OutputSubmitter is responsible for submitting outputs
type L2OutputSubmitter struct {
	log log.Logger
	cfg Config

	l2ooContract *bindings.L2OutputOracleCaller
	l2ooABI      *abi.ABI
}

// NewL2OutputSubmitter creates a new L2OutputSubmitter
func NewL2OutputSubmitter(cfg Config, l log.Logger) (*L2OutputSubmitter, error) {
	l2ooContract, err := bindings.NewL2OutputOracleCaller(cfg.L2OutputOracleAddr, cfg.L1Client)
	if err != nil {
		return nil, err
	}

	parsed, err := bindings.L2OutputOracleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	return &L2OutputSubmitter{
		log:          l,
		cfg:          cfg,
		l2ooContract: l2ooContract,
		l2ooABI:      parsed,
	}, nil
}

// FetchNextOutputInfo gets the block number of the next output.
// It returns: the next block number, if the output should be made, error
func (l *L2OutputSubmitter) FetchNextOutputInfo(ctx context.Context) (*eth.OutputResponse, bool, error) {
	nextBlockCtx, cancelNextBlockCtx := context.WithTimeout(ctx, l.cfg.NetworkTimeout)
	defer cancelNextBlockCtx()
	callOpts := utils.NewCallOptsWithSender(nextBlockCtx, l.cfg.TxManager.From())
	nextCheckpointBlock, err := l.l2ooContract.NextBlockNumber(callOpts)
	if err != nil {
		l.log.Error("validator unable to get next block number", "err", err)
		return nil, false, err
	}

	// Fetch the current L2 heads
	syncStatusCtx, cancelSyncStatusCtx := context.WithTimeout(ctx, l.cfg.NetworkTimeout)
	defer cancelSyncStatusCtx()
	status, err := l.cfg.RollupClient.SyncStatus(syncStatusCtx)
	if err != nil {
		l.log.Error("validator unable to get sync status", "err", err)
		return nil, false, err
	}

	// Use either the finalized or safe head depending on the config. Finalized head is default & safer.
	var currentBlockNumber *big.Int
	if l.cfg.AllowNonFinalized {
		currentBlockNumber = new(big.Int).SetUint64(status.SafeL2.Number)
	} else {
		currentBlockNumber = new(big.Int).SetUint64(status.FinalizedL2.Number)
	}
	var nextBlockNumber *big.Int
	if l.cfg.RollupConfig.IsBlueBlock(nextCheckpointBlock.Uint64()) {
		nextBlockNumber = new(big.Int).Add(nextCheckpointBlock, common.Big1)
	} else {
		nextBlockNumber = nextCheckpointBlock
	}
	// Ensure that we do not submit a block in the future
	if currentBlockNumber.Cmp(nextBlockNumber) < 0 {
		l.log.Info("validator submission interval has not elapsed", "currentBlockNumber", currentBlockNumber, "nextBlockNumber", nextCheckpointBlock)
		return nil, false, nil
	}

	return l.fetchOutput(ctx, nextCheckpointBlock)
}

func (l *L2OutputSubmitter) fetchOutput(ctx context.Context, block *big.Int) (*eth.OutputResponse, bool, error) {
	ctx, cancel := context.WithTimeout(ctx, l.cfg.NetworkTimeout)
	defer cancel()
	output, err := l.cfg.RollupClient.OutputAtBlock(ctx, block.Uint64(), false)
	if err != nil {
		l.log.Error("failed to fetch output at block %d: %w", block, err)
		return nil, false, err
	}
	if output.Version != rollup.L2OutputRootVersion(l.cfg.RollupConfig, l.cfg.RollupConfig.ComputeTimestamp(block.Uint64())) {
		l.log.Error("l2 output version is not matched: %s", output.Version)
		return nil, false, errors.New("mismatched l2 output version")
	}
	if output.BlockRef.Number != block.Uint64() { // sanity check, e.g. in case of bad RPC caching
		l.log.Error("invalid blockNumber", "next", block, "output", output.BlockRef.Number)
		return nil, false, errors.New("invalid blockNumber")
	}

	// Always submit if it's part of the Finalized L2 chain. Or if allowed, if it's part of the safe L2 chain.
	if !(output.BlockRef.Number <= output.Status.FinalizedL2.Number || (l.cfg.AllowNonFinalized && output.BlockRef.Number <= output.Status.SafeL2.Number)) {
		l.log.Debug("not submitting yet, L2 block is not ready for submission",
			"l2_output", output.BlockRef,
			"l2_safe", output.Status.SafeL2,
			"l2_finalized", output.Status.FinalizedL2,
			"allow_non_finalized", l.cfg.AllowNonFinalized)
		return nil, false, nil
	}
	return output, true, nil
}

// SubmitL2OutputTxData creates the transaction data for the submitL2Output function
func (l *L2OutputSubmitter) SubmitL2OutputTxData(output *eth.OutputResponse) ([]byte, error) {
	return submitL2OutputTxData(l.l2ooABI, output)
}

func submitL2OutputTxData(abi *abi.ABI, output *eth.OutputResponse) ([]byte, error) {
	return abi.Pack(
		"submitL2Output",
		output.OutputRoot,
		new(big.Int).SetUint64(output.BlockRef.Number),
		output.Status.CurrentL1.Hash,
		new(big.Int).SetUint64(output.Status.CurrentL1.Number))
}
