package validator

import (
	"context"
	"errors"
	"math/big"
	_ "net/http/pprof"
	"sync"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"

	"github.com/wemixkanvas/kanvas/bindings/bindings"
	"github.com/wemixkanvas/kanvas/components/node/eth"
	"github.com/wemixkanvas/kanvas/utils"
)

var supportedL2OutputVersion = eth.Bytes32{}

// L2OutputSubmitter is responsible for submitting outputs
type L2OutputSubmitter struct {
	wg   sync.WaitGroup
	done chan struct{}
	log  log.Logger
	cfg  Config

	ctx    context.Context
	cancel context.CancelFunc

	l2ooContract *bindings.L2OutputOracle
}

// NewL2OutputSubmitter creates a new L2 Output Submitter
func NewL2OutputSubmitter(ctx context.Context, cfg Config, l log.Logger) (*L2OutputSubmitter, error) {
	l2ooContract, err := bindings.NewL2OutputOracle(cfg.L2OutputOracleAddr, cfg.L1Client)
	if err != nil {
		return nil, err
	}

	return &L2OutputSubmitter{
		done:         make(chan struct{}),
		log:          l,
		cfg:          cfg,
		ctx:          ctx,
		l2ooContract: l2ooContract,
	}, nil
}

// FetchNextOutputInfo gets the block number of the next output.
// It returns: the next block number, if the output should be made, error
func (l *L2OutputSubmitter) FetchNextOutputInfo(ctx context.Context) (*eth.OutputResponse, bool, error) {
	callOpts := utils.NewCallOptsWithSender(ctx, l.cfg.From)
	nextCheckpointBlock, err := l.l2ooContract.NextBlockNumber(callOpts)
	if err != nil {
		l.log.Error("validator unable to get next block number", "err", err)
		return nil, false, err
	}
	// Fetch the current L2 heads
	status, err := l.cfg.RollupClient.SyncStatus(ctx)
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
	// Ensure that we do not submit a block in the future
	if currentBlockNumber.Cmp(nextCheckpointBlock) < 0 {
		l.log.Info("validator submission interval has not elapsed", "currentBlockNumber", currentBlockNumber, "nextBlockNumber", nextCheckpointBlock)
		return nil, false, nil
	}

	output, err := l.cfg.RollupClient.OutputAtBlock(ctx, nextCheckpointBlock.Uint64())
	if err != nil {
		l.log.Error("failed to fetch output at block %d: %w", nextCheckpointBlock, err)
		return nil, false, err
	}
	if output.Version != supportedL2OutputVersion {
		l.log.Error("unsupported l2 output version: %s", output.Version)
		return nil, false, errors.New("unsupported l2 output version")
	}
	if output.BlockRef.Number != nextCheckpointBlock.Uint64() { // sanity check, e.g. in case of bad RPC caching
		l.log.Error("invalid blockNumber", "next", nextCheckpointBlock, "output", output.BlockRef.Number)
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

// CreateSubmitL2OutputTx transforms an output response into a signed submit l2 output transaction.
// It does not send the transaction to the transaction pool.
func (l *L2OutputSubmitter) CreateSubmitL2OutputTx(ctx context.Context, output *eth.OutputResponse) (*types.Transaction, error) {
	opts := utils.NewSimpleTxOpts(ctx, l.cfg.From, l.cfg.SignerFn)

	tx, err := l.l2ooContract.SubmitL2Output(
		opts,
		output.OutputRoot,
		new(big.Int).SetUint64(output.BlockRef.Number),
		output.Status.CurrentL1.Hash,
		new(big.Int).SetUint64(output.Status.CurrentL1.Number))
	if err != nil {
		l.log.Error("failed to create the CreateOutputTx transaction", "err", err)
		return nil, err
	}
	return tx, nil
}
