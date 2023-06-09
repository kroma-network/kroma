package validator

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	_ "net/http/pprof"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"

	"github.com/kroma-network/kroma/bindings/bindings"
	"github.com/kroma-network/kroma/components/node/eth"
	"github.com/kroma-network/kroma/components/node/rollup"
	"github.com/kroma-network/kroma/components/validator/metrics"
	"github.com/kroma-network/kroma/utils"
	"github.com/kroma-network/kroma/utils/service/txmgr"
)

// L2OutputSubmitter is responsible for submitting outputs.
type L2OutputSubmitter struct {
	ctx    context.Context
	cancel context.CancelFunc

	cfg  Config
	log  log.Logger
	metr metrics.Metricer

	l2ooContract    *bindings.L2OutputOracleCaller
	l2ooABI         *abi.ABI
	valpoolContract *bindings.ValidatorPoolCaller

	submissionInterval *big.Int
	l2BlockTime        *big.Int

	txCandidatesChan chan<- txmgr.TxCandidate
	submitChan       chan struct{}

	wg sync.WaitGroup
}

// NewL2OutputSubmitter creates a new L2OutputSubmitter.
func NewL2OutputSubmitter(parentCtx context.Context, cfg Config, l log.Logger, m metrics.Metricer, txCandidatesChan chan<- txmgr.TxCandidate) (*L2OutputSubmitter, error) {
	ctx, cancel := context.WithCancel(parentCtx)

	submitChan := make(chan struct{}, 1)

	l2ooContract, err := bindings.NewL2OutputOracleCaller(cfg.L2OutputOracleAddr, cfg.L1Client)
	if err != nil {
		cancel()
		return nil, err
	}

	parsed, err := bindings.L2OutputOracleMetaData.GetAbi()
	if err != nil {
		cancel()
		return nil, err
	}

	valpoolContract, err := bindings.NewValidatorPoolCaller(cfg.ValidatorPoolAddr, cfg.L1Client)
	if err != nil {
		cancel()
		return nil, err
	}

	cCtx, cCancel := context.WithTimeout(ctx, cfg.NetworkTimeout)
	callOpts := utils.NewSimpleCallOpts(cCtx)
	l2BlockTime, err := l2ooContract.L2BLOCKTIME(callOpts)
	if err != nil {
		cancel()
		return nil, fmt.Errorf("failed to get l2 block time: %w", err)
	}
	cCancel()

	cCtx, cCancel = context.WithTimeout(ctx, cfg.NetworkTimeout)
	defer cCancel()
	callOpts = utils.NewSimpleCallOpts(cCtx)
	submissionInterval, err := l2ooContract.SUBMISSIONINTERVAL(callOpts)
	if err != nil {
		cancel()
		return nil, fmt.Errorf("failed to get submission interval: %w", err)
	}

	return &L2OutputSubmitter{
		ctx:                ctx,
		cancel:             cancel,
		cfg:                cfg,
		log:                l,
		metr:               m,
		l2ooContract:       l2ooContract,
		l2ooABI:            parsed,
		valpoolContract:    valpoolContract,
		submissionInterval: submissionInterval,
		l2BlockTime:        l2BlockTime,
		txCandidatesChan:   txCandidatesChan,
		submitChan:         submitChan,
	}, nil
}

func (l *L2OutputSubmitter) Start() error {
	l.log.Info("starting L2 Output Submitter")
	l.wg.Add(1)
	go l.loop()

	return nil
}

func (l *L2OutputSubmitter) Stop() error {
	l.log.Info("stopping L2 Output Submitter")
	l.cancel()
	l.wg.Wait()
	return nil
}

func (l *L2OutputSubmitter) loop() {
	defer l.wg.Done()

	l.submitChan <- struct{}{}

	for {
		select {
		case <-l.submitChan:
			if err := l.trySubmitL2Output(); err != nil {
				l.log.Error("failed to submit l2 output", "err", err)
				l.retryAfter(1)
			}
		case <-l.ctx.Done():
			return
		}
	}
}

func (l *L2OutputSubmitter) retryAfter(sec uint64) {
	time.AfterFunc(time.Duration(sec)*time.Second, func() {
		l.submitChan <- struct{}{}
	})
}

func (l *L2OutputSubmitter) trySubmitL2Output() error {
	nextBlockNumber, canSubmit, err := l.canSubmit(l.ctx)
	if err != nil {
		return fmt.Errorf("failed to check if it can submit: %w", err)
	}
	if !canSubmit {
		return nil
	}

	output, err := l.fetchOutput(l.ctx, nextBlockNumber)
	if err != nil {
		return fmt.Errorf("failed to fetch next output: %w", err)
	}

	data, err := submitL2OutputTxData(l.l2ooABI, output, l.cfg.OutputSubmitterBondAmount)
	if err != nil {
		return fmt.Errorf("failed to create submit l2 output transaction data: %w", err)
	}

	l.submitL2OutputTx(data)
	l.metr.RecordL2OutputSubmitted(output.BlockRef)
	l.retryAfter(1)

	return nil
}

// canSubmit checks if submission interval has elapsed and selected for next validator.
func (l *L2OutputSubmitter) canSubmit(ctx context.Context) (*big.Int, bool, error) {
	currentBlockNumber, nextBlockNumber, err := l.fetchBlockNumbers(ctx)
	if err != nil {
		return nil, false, err
	}

	var nextBlockNumberToWait *big.Int
	if l.cfg.RollupConfig.IsBlueBlock(nextBlockNumber.Uint64()) {
		nextBlockNumberToWait = new(big.Int).Add(nextBlockNumber, common.Big1)
	} else {
		nextBlockNumberToWait = nextBlockNumber
	}

	// Wait for L2 blocks proceeding when validator submission interval has not elapsed
	if currentBlockNumber.Cmp(nextBlockNumberToWait) < 0 {
		l.waitL2Blocks(currentBlockNumber, nextBlockNumberToWait)
		return nil, false, nil
	}

	// Check if selected for next validator or not
	isNextValidator, err := l.isNextValidator(ctx)
	if err != nil {
		return nil, false, err
	}
	// Wait for L2 blocks proceeding until next submission is available when not selected for next validator
	if !isNextValidator {
		nextBlockNumberToWait = new(big.Int).Add(nextBlockNumberToWait, l.submissionInterval)
		l.waitL2Blocks(currentBlockNumber, nextBlockNumberToWait)
		return nil, false, nil
	}

	return nextBlockNumber, true, nil
}

func (l *L2OutputSubmitter) fetchBlockNumbers(ctx context.Context) (*big.Int, *big.Int, error) {
	cCtx, cCancel := context.WithTimeout(ctx, l.cfg.NetworkTimeout)
	callOpts := utils.NewCallOptsWithSender(cCtx, l.cfg.TxManager.From())
	nextBlockNumber, err := l.l2ooContract.NextBlockNumber(callOpts)
	if err != nil {
		l.log.Error("validator unable to get next block number", "err", err)
		cCancel()
		return nil, nil, err
	}
	cCancel()

	// Fetch the current L2 heads
	cCtx, cCancel = context.WithTimeout(ctx, l.cfg.NetworkTimeout)
	defer cCancel()
	status, err := l.cfg.RollupClient.SyncStatus(cCtx)
	if err != nil {
		l.log.Error("validator unable to get sync status", "err", err)
		return nil, nil, err
	}

	// Use either the finalized or safe head depending on the config. Finalized head is default & safer.
	var currentBlockNumber *big.Int
	if l.cfg.AllowNonFinalized {
		currentBlockNumber = new(big.Int).SetUint64(status.SafeL2.Number)
	} else {
		currentBlockNumber = new(big.Int).SetUint64(status.FinalizedL2.Number)
	}

	return currentBlockNumber, nextBlockNumber, nil
}

func (l *L2OutputSubmitter) waitL2Blocks(currentBlockNumber *big.Int, targetBlockNumber *big.Int) {
	l.log.Info("validator submission interval has not elapsed", "currentBlockNumber", currentBlockNumber, "targetBlockNumber", targetBlockNumber)
	waitBlockNum := new(big.Int).Sub(targetBlockNumber, currentBlockNumber)
	waitSec := new(big.Int).Mul(waitBlockNum, l.l2BlockTime)
	if waitSec.Cmp(common.Big0) == -1 {
		waitSec = common.Big1
	}

	l.log.Info("wait for L2 blocks proceeding", "waitSec", waitSec)
	l.retryAfter(waitSec.Uint64())
}

func (l *L2OutputSubmitter) isNextValidator(ctx context.Context) (bool, error) {
	cCtx, cCancel := context.WithTimeout(ctx, l.cfg.NetworkTimeout)
	defer cCancel()
	callOpts := utils.NewCallOptsWithSender(cCtx, l.cfg.TxManager.From())
	nextValidator, err := l.valpoolContract.NextValidator(callOpts)
	if err != nil {
		l.log.Error("validator unable to get next validator address", "err", err)
		return false, err
	}

	if nextValidator != l.cfg.TxManager.From() {
		l.log.Info("not selected for next validator")
		return false, nil
	}

	return true, nil
}

// fetchOutput gets the output information to the corresponding block number.
// It returns the output info if the output can be made, otherwise error.
func (l *L2OutputSubmitter) fetchOutput(ctx context.Context, blockNumber *big.Int) (*eth.OutputResponse, error) {
	cCtx, cCancel := context.WithTimeout(ctx, l.cfg.NetworkTimeout)
	defer cCancel()
	output, err := l.cfg.RollupClient.OutputAtBlock(cCtx, blockNumber.Uint64(), false)
	if err != nil {
		l.log.Error("failed to fetch output at block number %d: %w", blockNumber, err)
		return nil, err
	}
	if output.Version != rollup.L2OutputRootVersion(l.cfg.RollupConfig, l.cfg.RollupConfig.ComputeTimestamp(blockNumber.Uint64())) {
		l.log.Error("l2 output version is not matched: %s", output.Version)
		return nil, errors.New("mismatched l2 output version")
	}
	if output.BlockRef.Number != blockNumber.Uint64() { // sanity check, e.g. in case of bad RPC caching
		l.log.Error("invalid block number", "next", blockNumber, "output", output.BlockRef.Number)
		return nil, errors.New("invalid block number")
	}

	return output, nil
}

// submitL2OutputTxData creates the transaction data for the submitL2OutputTx function.
func submitL2OutputTxData(abi *abi.ABI, output *eth.OutputResponse, bondAmount uint64) ([]byte, error) {
	return abi.Pack(
		"submitL2Output",
		output.OutputRoot,
		new(big.Int).SetUint64(output.BlockRef.Number),
		output.Status.CurrentL1.Hash,
		new(big.Int).SetUint64(output.Status.CurrentL1.Number),
		new(big.Int).SetUint64(bondAmount))
}

// submitL2OutputTx sends the l2 output submit tx to txCandidates channel to process validator's tx candidates in order.
func (l *L2OutputSubmitter) submitL2OutputTx(data []byte) {
	accessList := types.AccessList{
		types.AccessTuple{
			Address: l.cfg.ValidatorPoolAddr,
			StorageKeys: []common.Hash{
				common.HexToHash("0000000000000000000000000000000000000000000000000000000000000036"),
			},
		},
	}

	l.txCandidatesChan <- txmgr.TxCandidate{
		TxData:     data,
		To:         &l.cfg.L2OutputOracleAddr,
		GasLimit:   0,
		AccessList: accessList,
	}
}
