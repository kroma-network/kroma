package validator

import (
	"bytes"
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

const (
	roundNums      = 2
	publicRoundHex = "0xffffffffffffffffffffffffffffffffffffffff"
)

var PublicRoundAddress = common.HexToAddress(publicRoundHex)

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

	singleRoundInterval *big.Int
	l2BlockTime         *big.Int

	txCandidatesChan chan<- txmgr.TxCandidate
	submitChan       chan struct{}

	wg sync.WaitGroup
}

// NewL2OutputSubmitter creates a new L2OutputSubmitter.
func NewL2OutputSubmitter(ctx context.Context, cfg Config, l log.Logger, m metrics.Metricer) (*L2OutputSubmitter, error) {
	l2ooContract, err := bindings.NewL2OutputOracleCaller(cfg.L2OutputOracleAddr, cfg.L1Client)
	if err != nil {
		return nil, err
	}

	parsed, err := bindings.L2OutputOracleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	valpoolContract, err := bindings.NewValidatorPoolCaller(cfg.ValidatorPoolAddr, cfg.L1Client)
	if err != nil {
		return nil, err
	}

	cCtx, cCancel := context.WithTimeout(ctx, cfg.NetworkTimeout)
	callOpts := utils.NewSimpleCallOpts(cCtx)
	l2BlockTime, err := l2ooContract.L2BLOCKTIME(callOpts)
	if err != nil {
		cCancel()
		return nil, fmt.Errorf("failed to get l2 block time: %w", err)
	}
	cCancel()

	cCtx, cCancel = context.WithTimeout(ctx, cfg.NetworkTimeout)
	defer cCancel()
	callOpts = utils.NewSimpleCallOpts(cCtx)
	submissionInterval, err := l2ooContract.SUBMISSIONINTERVAL(callOpts)
	if err != nil {
		return nil, fmt.Errorf("failed to get submission interval: %w", err)
	}
	singleRoundInterval := new(big.Int).Div(submissionInterval, new(big.Int).SetUint64(roundNums))

	return &L2OutputSubmitter{
		cfg:                 cfg,
		log:                 l,
		metr:                m,
		l2ooContract:        l2ooContract,
		l2ooABI:             parsed,
		valpoolContract:     valpoolContract,
		singleRoundInterval: singleRoundInterval,
		l2BlockTime:         l2BlockTime,
	}, nil
}

func (l *L2OutputSubmitter) Start(ctx context.Context, txCandidatesChan chan<- txmgr.TxCandidate) error {
	l.ctx, l.cancel = context.WithCancel(ctx)
	l.log.Info("starting L2 Output Submitter")

	l.submitChan = make(chan struct{}, 1)
	l.txCandidatesChan = txCandidatesChan
	l.wg.Add(1)
	go l.loop()

	return nil
}

func (l *L2OutputSubmitter) Stop() error {
	l.log.Info("stopping L2 Output Submitter")

	l.cancel()
	l.wg.Wait()

	close(l.submitChan)

	return nil
}

func (l *L2OutputSubmitter) loop() {
	defer l.wg.Done()

	l.submitChan <- struct{}{}

	for {
		select {
		case <-l.submitChan:
			if err := l.trySubmitL2Output(l.ctx); err != nil {
				l.log.Error("failed to submit l2 output", "err", err)
				l.retryAfter(l.cfg.OutputSubmitterRetryInterval)
			}
		case <-l.ctx.Done():
			return
		}
	}
}

func (l *L2OutputSubmitter) retryAfter(d time.Duration) {
	l.wg.Add(1)

	time.AfterFunc(d, func() {
		l.submitChan <- struct{}{}
		l.wg.Done()
	})
}

// TODO(seolaoh): return wait duration explicitly, and handle `retryAfter` function calls at once.
func (l *L2OutputSubmitter) trySubmitL2Output(ctx context.Context) error {
	nextBlockNumber, canSubmit, err := l.CanSubmit(ctx)
	if err != nil {
		return fmt.Errorf("failed to check if it can submit: %w", err)
	}
	if !canSubmit {
		return nil
	}

	output, err := l.FetchOutput(ctx, nextBlockNumber)
	if err != nil {
		return fmt.Errorf("failed to fetch next output: %w", err)
	}

	data, err := SubmitL2OutputTxData(l.l2ooABI, output, l.cfg.OutputSubmitterBondAmount)
	if err != nil {
		return fmt.Errorf("failed to create submit l2 output transaction data: %w", err)
	}

	if err := l.submitL2OutputTx(data, nextBlockNumber); err != nil {
		return fmt.Errorf("failed to submit l2 output transaction: %w", err)
	}
	l.metr.RecordL2OutputSubmitted(output.BlockRef)
	l.retryAfter(l.cfg.OutputSubmitterRetryInterval)

	return nil
}

// CanSubmit checks if submission interval has elapsed and current round conditions.
func (l *L2OutputSubmitter) CanSubmit(ctx context.Context) (*big.Int, bool, error) {
	hasEnoughDeposit, err := l.checkDeposit(ctx)
	if err != nil {
		return nil, false, err
	}
	if !hasEnoughDeposit {
		l.retryAfter(l.cfg.OutputSubmitterRetryInterval)
		return nil, false, nil
	}

	currentBlockNumber, nextBlockNumber, err := l.fetchBlockNumbers(ctx)
	if err != nil {
		return nil, false, err
	}
	l.log.Info("current status before submit", "currentBlockNumber", currentBlockNumber, "nextBlockNumberToSubmit", nextBlockNumber)

	var nextBlockNumberToWait *big.Int
	if l.cfg.RollupConfig.IsBlueBlock(nextBlockNumber.Uint64()) {
		nextBlockNumberToWait = new(big.Int).Add(nextBlockNumber, common.Big1)
	} else {
		nextBlockNumberToWait = nextBlockNumber
	}

	// Wait for L2 blocks proceeding when validator submission interval has not elapsed
	roundBuffer := new(big.Int).SetUint64(l.cfg.OutputSubmitterRoundBuffer)
	if currentBlockNumber.Cmp(nextBlockNumberToWait) < 0 {
		nextBlockNumberToWait = new(big.Int).Sub(nextBlockNumber, roundBuffer)
		l.waitL2Blocks(currentBlockNumber, nextBlockNumberToWait)
		return nil, false, nil
	}

	// Check if it's a public round, or selected for priority validator
	roundInfo, err := l.fetchCurrentRound(ctx)
	if err != nil {
		return nil, false, err
	}
	// if it's a public round, try to submit right now
	if roundInfo.isPublicRound {
		return nextBlockNumber, true, nil
	}
	// if it's a priority round, wait for L2 blocks proceeding until public round when not selected for priority validator
	if !roundInfo.isPriorityValidator {
		roundIntervalToWait := new(big.Int).Sub(l.singleRoundInterval, roundBuffer)
		nextBlockNumberToWait = new(big.Int).Add(nextBlockNumber, roundIntervalToWait)
		l.waitL2Blocks(currentBlockNumber, nextBlockNumberToWait)
		return nil, false, nil
	}

	return nextBlockNumber, true, nil
}

func (l *L2OutputSubmitter) checkDeposit(ctx context.Context) (bool, error) {
	cCtx, cCancel := context.WithTimeout(ctx, l.cfg.NetworkTimeout)
	defer cCancel()
	from := l.cfg.TxManager.From()
	callOpts := utils.NewCallOptsWithSender(cCtx, from)
	balance, err := l.valpoolContract.BalanceOf(callOpts, from)
	if err != nil {
		return false, fmt.Errorf("failed to fetch validator deposit amount: %w", err)
	}

	if balance.Cmp(new(big.Int).SetUint64(l.cfg.OutputSubmitterBondAmount)) == -1 {
		l.log.Warn("validator deposit is less than bond attempt amount", "bondAttemptAmount", l.cfg.OutputSubmitterBondAmount, "deposit", balance)
		return false, nil
	}
	l.log.Info("validator deposit amount", "deposit", balance)

	return true, nil
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

	var waitDuration time.Duration
	if waitBlockNum.Cmp(common.Big0) == -1 {
		waitDuration = l.cfg.OutputSubmitterRetryInterval
	} else {
		waitDuration = time.Duration(new(big.Int).Mul(waitBlockNum, l.l2BlockTime).Uint64()) * time.Second
	}

	l.log.Info("wait for L2 blocks proceeding", "waitDuration", waitDuration)
	l.retryAfter(waitDuration)
}

type roundInfo struct {
	isPublicRound       bool
	isPriorityValidator bool
}

// fetchCurrentRound fetches next validator address from ValidatorPool contract.
// It returns if current round is public round, and if selected for priority validator if it's a priority round.
func (l *L2OutputSubmitter) fetchCurrentRound(ctx context.Context) (roundInfo, error) {
	cCtx, cCancel := context.WithTimeout(ctx, l.cfg.NetworkTimeout)
	defer cCancel()
	callOpts := utils.NewCallOptsWithSender(cCtx, l.cfg.TxManager.From())
	nextValidator, err := l.valpoolContract.NextValidator(callOpts)
	if err != nil {
		l.log.Error("validator unable to get next validator address", "err", err)
		return roundInfo{
			isPublicRound:       false,
			isPriorityValidator: false,
		}, err
	}

	if bytes.Equal(nextValidator[:], PublicRoundAddress[:]) {
		l.log.Info("current round is public round")
		return roundInfo{
			isPublicRound:       true,
			isPriorityValidator: false,
		}, nil
	}

	if nextValidator == l.cfg.TxManager.From() {
		l.log.Info("current round is priority round, and selected for priority validator")
		return roundInfo{
			isPublicRound:       false,
			isPriorityValidator: true,
		}, nil
	}

	l.log.Info("current round is priority round, and not selected for priority validator")
	return roundInfo{
		isPublicRound:       false,
		isPriorityValidator: false,
	}, nil
}

// FetchOutput gets the output information to the corresponding block number.
// It returns the output info if the output can be made, otherwise error.
func (l *L2OutputSubmitter) FetchOutput(ctx context.Context, blockNumber *big.Int) (*eth.OutputResponse, error) {
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

// SubmitL2OutputTxData creates the transaction data for the submitL2OutputTx function.
func SubmitL2OutputTxData(abi *abi.ABI, output *eth.OutputResponse, bondAmount uint64) ([]byte, error) {
	return abi.Pack(
		"submitL2Output",
		output.OutputRoot,
		new(big.Int).SetUint64(output.BlockRef.Number),
		output.Status.CurrentL1.Hash,
		new(big.Int).SetUint64(output.Status.CurrentL1.Number),
		new(big.Int).SetUint64(bondAmount))
}

// submitL2OutputTx creates l2 output submit tx candidate and sends it to txCandidates channel to process validator's tx candidates in order.
func (l *L2OutputSubmitter) submitL2OutputTx(data []byte, nextBlockNumber *big.Int) error {
	layout, err := bindings.GetStorageLayout("ValidatorPool")
	if err != nil {
		return fmt.Errorf("failed to get storage layout: %w", err)
	}

	var outputIndexSlot, priorityValidatorSlot common.Hash
	for _, entry := range layout.Storage {
		switch entry.Label {
		case "nextUnbondOutputIndex":
			outputIndexSlot = common.BigToHash(big.NewInt(int64(entry.Slot)))
		case "nextPriorityValidator":
			priorityValidatorSlot = common.BigToHash(big.NewInt(int64(entry.Slot)))
		}
	}

	storageKeys := []common.Hash{outputIndexSlot, priorityValidatorSlot}

	// If provide accessList that is not actually accessed, the transaction may not be executed due to exceeding the estimated gas limit
	accessList := types.AccessList{
		types.AccessTuple{
			Address:     l.cfg.ValidatorPoolAddr,
			StorageKeys: storageKeys,
		},
	}

	l.txCandidatesChan <- txmgr.TxCandidate{
		TxData:     data,
		To:         &l.cfg.L2OutputOracleAddr,
		GasLimit:   0,
		AccessList: accessList,
	}

	return nil
}

func (l *L2OutputSubmitter) L2ooAbi() *abi.ABI {
	return l.l2ooABI
}
