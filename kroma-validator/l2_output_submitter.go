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

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"

	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum-optimism/optimism/op-service/optsutils"
	"github.com/ethereum-optimism/optimism/op-service/txmgr"
	"github.com/ethereum-optimism/optimism/op-service/watcher"
	"github.com/kroma-network/kroma/kroma-bindings/bindings"
	"github.com/kroma-network/kroma/kroma-validator/metrics"
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
	requiredBondAmount  *big.Int

	submitChan chan struct{}

	wg sync.WaitGroup
}

// NewL2OutputSubmitter creates a new L2OutputSubmitter.
func NewL2OutputSubmitter(cfg Config, l log.Logger, m metrics.Metricer) (*L2OutputSubmitter, error) {
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

	return &L2OutputSubmitter{
		cfg:             cfg,
		log:             l.New("service", "submitter"),
		metr:            m,
		l2ooContract:    l2ooContract,
		l2ooABI:         parsed,
		valpoolContract: valpoolContract,
	}, nil
}

func (l *L2OutputSubmitter) InitConfig(ctx context.Context) error {
	contractWatcher := watcher.NewContractWatcher(ctx, l.cfg.L1Client, l.log)

	err := contractWatcher.WatchUpgraded(l.cfg.L2OutputOracleAddr, func() error {
		cCtx, cCancel := context.WithTimeout(ctx, l.cfg.NetworkTimeout)
		defer cCancel()
		l2BlockTime, err := l.l2ooContract.L2BLOCKTIME(optsutils.NewSimpleCallOpts(cCtx))
		if err != nil {
			return fmt.Errorf("failed to get l2 block time: %w", err)
		}
		l.l2BlockTime = l2BlockTime

		cCtx, cCancel = context.WithTimeout(ctx, l.cfg.NetworkTimeout)
		defer cCancel()
		submissionInterval, err := l.l2ooContract.SUBMISSIONINTERVAL(optsutils.NewSimpleCallOpts(cCtx))
		if err != nil {
			return fmt.Errorf("failed to get submission interval: %w", err)
		}
		singleRoundInterval := new(big.Int).Div(submissionInterval, new(big.Int).SetUint64(roundNums))
		l.singleRoundInterval = singleRoundInterval

		return nil
	})
	if err != nil {
		return fmt.Errorf("failed to initiate l2oo config: %w", err)
	}

	err = contractWatcher.WatchUpgraded(l.cfg.ValidatorPoolAddr, func() error {
		cCtx, cCancel := context.WithTimeout(ctx, l.cfg.NetworkTimeout)
		defer cCancel()
		requiredBondAmount, err := l.valpoolContract.REQUIREDBONDAMOUNT(optsutils.NewSimpleCallOpts(cCtx))
		if err != nil {
			return fmt.Errorf("failed to get required bond amount: %w", err)
		}
		l.requiredBondAmount = requiredBondAmount

		return nil
	})
	if err != nil {
		return fmt.Errorf("failed to initiate valpool config: %w", err)
	}

	return nil
}

func (l *L2OutputSubmitter) Start(ctx context.Context) error {
	l.ctx, l.cancel = context.WithCancel(ctx)
	l.submitChan = make(chan struct{}, 1)

	if err := l.InitConfig(l.ctx); err != nil {
		return err
	}

	l.wg.Add(1)
	go l.loop()

	return nil
}

func (l *L2OutputSubmitter) Stop() error {
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

	data, err := SubmitL2OutputTxData(l.l2ooABI, output)
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

	currentBlockNumber, err := l.FetchCurrentBlockNumber(ctx)
	if err != nil {
		return defaultWaitTime
	}

	hasEnoughDeposit, err := l.HasEnoughDeposit(ctx)
	if err != nil {
		return defaultWaitTime
	}
	if !hasEnoughDeposit {
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

// HasEnoughDeposit checks if validator has enough deposit to bond when trying output submission.
func (l *L2OutputSubmitter) HasEnoughDeposit(ctx context.Context) (bool, error) {
	cCtx, cCancel := context.WithTimeout(ctx, l.cfg.NetworkTimeout)
	defer cCancel()
	from := l.cfg.TxManager.From()
	balance, err := l.valpoolContract.BalanceOf(optsutils.NewSimpleCallOpts(cCtx), from)
	if err != nil {
		return false, fmt.Errorf("failed to fetch deposit amount: %w", err)
	}

	if balance.Cmp(l.requiredBondAmount) == -1 {
		l.log.Warn(
			"deposit is less than bond attempt amount",
			"requiredBondAmount", l.requiredBondAmount,
			"deposit", balance,
		)
		return false, nil
	}
	l.log.Info("deposit amount", "deposit", balance)
	l.metr.RecordDepositAmount(balance)

	return true, nil
}

func (l *L2OutputSubmitter) FetchNextBlockNumber(ctx context.Context) (*big.Int, error) {
	cCtx, cCancel := context.WithTimeout(ctx, l.cfg.NetworkTimeout)
	defer cCancel()
	nextBlockNumber, err := l.l2ooContract.NextBlockNumber(optsutils.NewSimpleCallOpts(cCtx))
	if err != nil {
		l.log.Error("unable to get next block number", "err", err)
		return nil, err
	}

	return nextBlockNumber, nil
}

func (l *L2OutputSubmitter) FetchCurrentBlockNumber(ctx context.Context) (*big.Int, error) {
	// fetch the current L2 heads
	cCtx, cCancel := context.WithTimeout(ctx, l.cfg.NetworkTimeout)
	defer cCancel()
	status, err := l.cfg.RollupClient.SyncStatus(cCtx)
	if err != nil {
		l.log.Error("unable to get sync status", "err", err)
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
	l.log.Info("submission interval has not elapsed", "currentBlockNumber", currentBlockNumber, "targetBlockNumber", targetBlockNumber)
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
	canJoinPublicRound  bool
}

func (r *roundInfo) canJoinRound() bool {
	joinPriority := !r.isPublicRound && r.isPriorityValidator
	joinPublic := r.isPublicRound && r.canJoinPublicRound
	return joinPriority || joinPublic
}

// fetchCurrentRound fetches next validator address from ValidatorPool contract.
// It returns if current round is public round, and if selected for priority validator if it's a priority round.
func (l *L2OutputSubmitter) fetchCurrentRound(ctx context.Context) (roundInfo, error) {
	cCtx, cCancel := context.WithTimeout(ctx, l.cfg.NetworkTimeout)
	defer cCancel()
	ri := roundInfo{canJoinPublicRound: l.cfg.OutputSubmitterAllowPublicRound}
	nextValidator, err := l.valpoolContract.NextValidator(optsutils.NewSimpleCallOpts(cCtx))
	if err != nil {
		l.log.Error("unable to get next validator address", "err", err)
		ri.isPublicRound = false
		ri.isPriorityValidator = false
		return ri, err
	}

	l.metr.RecordNextValidator(nextValidator)

	if bytes.Equal(nextValidator[:], PublicRoundAddress[:]) {
		l.log.Info("current round is public round")
		ri.isPublicRound = true
		ri.isPriorityValidator = false
		return ri, nil
	}

	if nextValidator == l.cfg.TxManager.From() {
		l.log.Info("current round is priority round, and selected for priority validator")
		ri.isPublicRound = false
		ri.isPriorityValidator = true
		return ri, nil
	}

	l.log.Info("current round is priority round, and not selected for priority validator")
	ri.isPublicRound = false
	ri.isPriorityValidator = false
	return ri, nil
}

// FetchOutput gets the output information to the corresponding block number.
// It returns the output info if the output can be made, otherwise error.
func (l *L2OutputSubmitter) FetchOutput(ctx context.Context, blockNumber *big.Int) (*eth.OutputResponse, error) {
	cCtx, cCancel := context.WithTimeout(ctx, l.cfg.NetworkTimeout)
	defer cCancel()
	output, err := l.cfg.RollupClient.OutputAtBlock(cCtx, blockNumber.Uint64())
	if err != nil {
		l.log.Error("failed to fetch output at ", "block number", blockNumber.Uint64(), " err", err)
		return nil, err
	}
	if output.Version != eth.OutputVersionV0 {
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
func SubmitL2OutputTxData(abi *abi.ABI, output *eth.OutputResponse) ([]byte, error) {
	return abi.Pack(
		"submitL2Output",
		output.OutputRoot,
		new(big.Int).SetUint64(output.BlockRef.Number),
		output.Status.CurrentL1.Hash,
		new(big.Int).SetUint64(output.Status.CurrentL1.Number),
	)
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

	// Do the gas estimation and set 150% of it to gas limit to prevent tx failed because of dynamic gas usage in unbond and priority validator selection
	gasTipCap, basefee, err := l.cfg.TxManager.SuggestGasPriceCaps(l.ctx)
	if err != nil {
		return &txmgr.TxResponse{
			Receipt: nil,
			Err:     fmt.Errorf("failed to get gas price info: %w", err),
		}
	}
	gasFeeCap := txmgr.CalcGasFeeCap(basefee, gasTipCap)

	to := &l.cfg.L2OutputOracleAddr
	estimatedGas, err := l.cfg.L1Client.EstimateGas(l.ctx, ethereum.CallMsg{
		From:      l.cfg.TxManager.From(),
		To:        to,
		GasFeeCap: gasFeeCap,
		GasTipCap: gasTipCap,
		Data:      data,
	})
	if err != nil {
		return &txmgr.TxResponse{
			Receipt: nil,
			Err:     fmt.Errorf("failed to estimate gas: %w", err),
		}
	}

	return l.cfg.TxManager.SendTxCandidate(l.ctx, &txmgr.TxCandidate{
		TxData:     data,
		To:         to,
		GasLimit:   estimatedGas * 3 / 2,
		AccessList: accessList,
	})
}

func (l *L2OutputSubmitter) L2ooAbi() *abi.ABI {
	return l.l2ooABI
}
