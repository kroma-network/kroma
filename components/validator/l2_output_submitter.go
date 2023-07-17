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

	submitChan chan struct{}

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

func (l *L2OutputSubmitter) Start(ctx context.Context) error {
	l.ctx, l.cancel = context.WithCancel(ctx)
	l.log.Info("starting L2 Output Submitter")

	l.submitChan = make(chan struct{}, 1)
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

	for ; ; <-l.submitChan {
		select {
		case <-l.ctx.Done():
			return
		default:
			l.repeatSubmitL2Output(l.ctx)
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

func (l *L2OutputSubmitter) repeatSubmitL2Output(ctx context.Context) {
	waitTime, err := l.trySubmitL2Output(ctx)
	if err != nil {
		l.log.Error("failed to submit L2Output", "err", err)
	}
	l.retryAfter(waitTime)
}

// trySubmitL2Output checks if the validator can submit l2 output and tries to submit it.
// If it needs to wait, it will calculate how long the validator should wait and
// try again after the delay.
func (l *L2OutputSubmitter) trySubmitL2Output(ctx context.Context) (time.Duration, error) {
	nextBlockNumber, err := l.FetchNextBlockNumber(ctx)
	if err != nil {
		return l.cfg.OutputSubmitterRetryInterval, err
	}

	calculatedWaitTime := l.CalculateWaitTime(ctx, nextBlockNumber)
	if calculatedWaitTime > 0 {
		return calculatedWaitTime, nil
	}

	if err = l.doSubmitL2Output(ctx, nextBlockNumber); err != nil {
		return l.cfg.OutputSubmitterRetryInterval, err
	}

	// successfully submitted. start next loop immediately.
	return 0, nil
}

// doSubmitL2Output submits l2 Output submission transaction.
func (l *L2OutputSubmitter) doSubmitL2Output(ctx context.Context, nextBlockNumber *big.Int) error {
	output, err := l.FetchOutput(ctx, nextBlockNumber)
	if err != nil {
		return err
	}

	data, err := SubmitL2OutputTxData(l.l2ooABI, output, l.cfg.OutputSubmitterBondAmount)
	if err != nil {
		return fmt.Errorf("failed to create submit l2 output transaction data: %w", err)
	}

	if txResponse := l.submitL2OutputTx(data); txResponse.Err != nil {
		return txResponse.Err
	}

	// Successfully submitted
	l.log.Info("L2output successfully submitted", "blockNumber", output.BlockRef.Number)
	l.metr.RecordL2OutputSubmitted(output.BlockRef)
	// go to try next submission immediately
	return nil
}

// CalculateWaitTime checks the conditions for submitting L2Output and calculates the required latency.
// Returns time 0 if the conditions are such that submission is possible immediately.
func (l *L2OutputSubmitter) CalculateWaitTime(ctx context.Context, nextBlockNumber *big.Int) time.Duration {
	defaultWaitTime := l.cfg.OutputSubmitterRetryInterval
	hasEnoughDeposit, err := l.checkDeposit(ctx)
	if err != nil {
		return defaultWaitTime
	}
	if !hasEnoughDeposit {
		return defaultWaitTime
	}

	currentBlockNumber, err := l.fetchCurrentBlockNumber(ctx)
	if err != nil {
		return defaultWaitTime
	}
	l.log.Info("current status before submit", "currentBlockNumber", currentBlockNumber, "nextBlockNumberToSubmit", nextBlockNumber)

	// Wait for L2 blocks proceeding when validator submission interval has not elapsed
	// Need to wait next block number to submit plus 1 because of next block hash inclusion
	nextBlockNumberToWait := new(big.Int).Add(nextBlockNumber, common.Big1)
	roundBuffer := new(big.Int).SetUint64(l.cfg.OutputSubmitterRoundBuffer)
	if currentBlockNumber.Cmp(nextBlockNumberToWait) < 0 {
		nextBlockNumberToWait = new(big.Int).Sub(nextBlockNumber, roundBuffer)
		return l.getLeftTimeForL2Blocks(currentBlockNumber, nextBlockNumberToWait)
	}

	// Check if it's a public round, or selected for priority validator
	roundInfo, err := l.fetchCurrentRound(ctx)
	if err != nil {
		return defaultWaitTime
	}

	if !roundInfo.canJoinRound() {
		// wait for L2 blocks proceeding until public round when not selected for priority validator
		roundIntervalToWait := new(big.Int).Sub(l.singleRoundInterval, roundBuffer)
		nextBlockNumberToWait = new(big.Int).Add(nextBlockNumber, roundIntervalToWait)
		return l.getLeftTimeForL2Blocks(currentBlockNumber, nextBlockNumberToWait)
	}

	// no need to wait
	return 0
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
	l.metr.RecordDepositAmount(balance)

	return true, nil
}

func (l *L2OutputSubmitter) FetchNextBlockNumber(ctx context.Context) (*big.Int, error) {
	cCtx, cCancel := context.WithTimeout(ctx, l.cfg.NetworkTimeout)
	defer cCancel()
	callOpts := utils.NewCallOptsWithSender(cCtx, l.cfg.TxManager.From())
	return l.l2ooContract.NextBlockNumber(callOpts)
}

func (l *L2OutputSubmitter) fetchCurrentBlockNumber(ctx context.Context) (*big.Int, error) {
	// Fetch the current L2 heads
	cCtx, cCancel := context.WithTimeout(ctx, l.cfg.NetworkTimeout)
	defer cCancel()
	status, err := l.cfg.RollupClient.SyncStatus(cCtx)
	if err != nil {
		l.log.Error("validator unable to get sync status", "err", err)
		return nil, err
	}

	// Use either the finalized or safe head depending on the config. Finalized head is default & safer.
	var currentBlockNumber *big.Int
	if l.cfg.AllowNonFinalized {
		currentBlockNumber = new(big.Int).SetUint64(status.SafeL2.Number)
	} else {
		currentBlockNumber = new(big.Int).SetUint64(status.FinalizedL2.Number)
	}

	return currentBlockNumber, nil
}

func (l *L2OutputSubmitter) getLeftTimeForL2Blocks(currentBlockNumber *big.Int, targetBlockNumber *big.Int) time.Duration {
	l.log.Info("validator submission interval has not elapsed", "currentBlockNumber", currentBlockNumber, "targetBlockNumber", targetBlockNumber)
	waitBlockNum := new(big.Int).Sub(targetBlockNumber, currentBlockNumber)

	var waitDuration time.Duration
	if waitBlockNum.Cmp(common.Big0) == -1 {
		waitDuration = l.cfg.OutputSubmitterRetryInterval
	} else {
		waitDuration = time.Duration(new(big.Int).Mul(waitBlockNum, l.l2BlockTime).Uint64()) * time.Second
	}

	l.log.Info("wait for L2 blocks proceeding", "waitDuration", waitDuration)
	return waitDuration
}

type roundInfo struct {
	isPublicRound       bool
	isPriorityValidator bool
}

func (r *roundInfo) canJoinRound() bool {
	joinPriority := !r.isPublicRound && r.isPriorityValidator
	joinPublic := r.isPublicRound
	return joinPriority || joinPublic
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

	l.metr.RecordNextValidator(nextValidator)

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
	output, err := l.cfg.RollupClient.OutputAtBlock(cCtx, blockNumber.Uint64())
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
func (l *L2OutputSubmitter) submitL2OutputTx(data []byte) *txmgr.TxResponse {
	layout, err := bindings.GetStorageLayout("ValidatorPool")
	if err != nil {
		return &txmgr.TxResponse{
			Receipt: nil,
			Err:     fmt.Errorf("failed to get storage layout: %w", err),
		}
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

	return l.cfg.TxManager.SendTxCandidate(l.ctx, &txmgr.TxCandidate{
		TxData:     data,
		To:         &l.cfg.L2OutputOracleAddr,
		GasLimit:   0,
		AccessList: accessList,
	})
}

func (l *L2OutputSubmitter) L2ooAbi() *abi.ABI {
	return l.l2ooABI
}
