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

	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum-optimism/optimism/op-service/optsutils"
	"github.com/ethereum-optimism/optimism/op-service/txmgr"
	"github.com/ethereum-optimism/optimism/op-service/watcher"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"

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

	l2OOContract     *bindings.L2OutputOracleCaller
	l2OOABI          *abi.ABI
	valPoolContract  *bindings.ValidatorPoolCaller
	valMgrContract   *bindings.ValidatorManagerCaller
	assetMgrContract *bindings.AssetManagerCaller

	singleRoundInterval     *big.Int
	l2BlockTime             *big.Int
	requiredBondAmountV1    *big.Int
	requiredBondAmountV2    *big.Int
	valPoolTerminationIndex *big.Int

	submitChan chan struct{}

	wg sync.WaitGroup
}

// NewL2OutputSubmitter creates a new L2OutputSubmitter.
func NewL2OutputSubmitter(cfg Config, l log.Logger, m metrics.Metricer) (*L2OutputSubmitter, error) {
	l2OOContract, err := bindings.NewL2OutputOracleCaller(cfg.L2OutputOracleAddr, cfg.L1Client)
	if err != nil {
		return nil, err
	}

	parsedL2OOAbi, err := bindings.L2OutputOracleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	valPoolContract, err := bindings.NewValidatorPoolCaller(cfg.ValidatorPoolAddr, cfg.L1Client)
	if err != nil {
		return nil, err
	}

	valMgrContract, err := bindings.NewValidatorManagerCaller(cfg.ValidatorManagerAddr, cfg.L1Client)
	if err != nil {
		return nil, err
	}

	assetMgrContract, err := bindings.NewAssetManagerCaller(cfg.AssetManagerAddr, cfg.L1Client)
	if err != nil {
		return nil, err
	}

	return &L2OutputSubmitter{
		cfg:              cfg,
		log:              l.New("service", "submitter"),
		metr:             m,
		l2OOContract:     l2OOContract,
		l2OOABI:          parsedL2OOAbi,
		valPoolContract:  valPoolContract,
		valMgrContract:   valMgrContract,
		assetMgrContract: assetMgrContract,
	}, nil
}

func (l *L2OutputSubmitter) InitConfig(ctx context.Context) error {
	contractWatcher := watcher.NewContractWatcher(ctx, l.cfg.L1Client, l.log)

	err := contractWatcher.WatchUpgraded(l.cfg.L2OutputOracleAddr, func() error {
		cCtx, cCancel := context.WithTimeout(ctx, l.cfg.NetworkTimeout)
		defer cCancel()
		l2BlockTime, err := l.l2OOContract.L2BLOCKTIME(optsutils.NewSimpleCallOpts(cCtx))
		if err != nil {
			return fmt.Errorf("failed to get l2 block time: %w", err)
		}
		l.l2BlockTime = l2BlockTime

		cCtx, cCancel = context.WithTimeout(ctx, l.cfg.NetworkTimeout)
		defer cCancel()
		submissionInterval, err := l.l2OOContract.SUBMISSIONINTERVAL(optsutils.NewSimpleCallOpts(cCtx))
		if err != nil {
			return fmt.Errorf("failed to get submission interval: %w", err)
		}
		singleRoundInterval := new(big.Int).Div(submissionInterval, new(big.Int).SetUint64(roundNums))
		l.singleRoundInterval = singleRoundInterval

		return nil
	})
	if err != nil {
		return fmt.Errorf("failed to initiate l2OO config: %w", err)
	}

	err = contractWatcher.WatchUpgraded(l.cfg.ValidatorPoolAddr, func() error {
		cCtx, cCancel := context.WithTimeout(ctx, l.cfg.NetworkTimeout)
		defer cCancel()
		requiredBondAmountV1, err := l.valPoolContract.REQUIREDBONDAMOUNT(optsutils.NewSimpleCallOpts(cCtx))
		if err != nil {
			return fmt.Errorf("failed to get required bond amount: %w", err)
		}
		l.requiredBondAmountV1 = requiredBondAmountV1

		cCtx, cCancel = context.WithTimeout(ctx, l.cfg.NetworkTimeout)
		defer cCancel()
		valPoolTerminationIndex, err := l.valPoolContract.TERMINATEOUTPUTINDEX(optsutils.NewSimpleCallOpts(cCtx))
		if err != nil {
			return fmt.Errorf("failed to get valPool termination index: %w", err)
		}
		l.valPoolTerminationIndex = valPoolTerminationIndex

		return nil
	})
	if err != nil {
		return fmt.Errorf("failed to initiate valPool config: %w", err)
	}

	err = contractWatcher.WatchUpgraded(l.cfg.AssetManagerAddr, func() error {
		cCtx, cCancel := context.WithTimeout(ctx, l.cfg.NetworkTimeout)
		defer cCancel()
		requiredBondAmountV2, err := l.assetMgrContract.BONDAMOUNT(optsutils.NewSimpleCallOpts(cCtx))
		if err != nil {
			return fmt.Errorf("failed to get required bond amount of assetMgr: %w", err)
		}
		l.requiredBondAmountV2 = requiredBondAmountV2

		return nil
	})
	if err != nil {
		return fmt.Errorf("failed to initiate assetMgr config: %w", err)
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
	defaultWaitTime := l.cfg.OutputSubmitterRetryInterval

	nextBlockNumber, err := l.FetchNextBlockNumber(ctx)
	if err != nil {
		return defaultWaitTime, err
	}

	outputIndex, err := l.FetchNextOutputIndex(ctx)
	if err != nil {
		return defaultWaitTime, err
	}

	calculatedWaitTime := l.CalculateWaitTime(ctx, nextBlockNumber, outputIndex)
	if calculatedWaitTime > 0 {
		return calculatedWaitTime, nil
	}

	canSubmitOutput, err := l.CanSubmitOutput(ctx, outputIndex)
	if err != nil || !canSubmitOutput {
		return defaultWaitTime, err
	}

	if err = l.doSubmitL2Output(ctx, nextBlockNumber, outputIndex); err != nil {
		return defaultWaitTime, err
	}

	// successfully submitted. start next loop immediately.
	return 0, nil
}

// doSubmitL2Output submits l2 Output submission transaction.
func (l *L2OutputSubmitter) doSubmitL2Output(ctx context.Context, nextBlockNumber *big.Int, outputIndex *big.Int) error {
	output, err := l.FetchOutput(ctx, nextBlockNumber)
	if err != nil {
		return err
	}

	data, err := SubmitL2OutputTxData(l.l2OOABI, output)
	if err != nil {
		return fmt.Errorf("failed to create submit l2 output transaction data: %w", err)
	}

	if txResponse := l.submitL2OutputTx(data, outputIndex); txResponse.Err != nil {
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
func (l *L2OutputSubmitter) CalculateWaitTime(ctx context.Context, nextBlockNumber *big.Int, outputIndex *big.Int) time.Duration {
	defaultWaitTime := l.cfg.OutputSubmitterRetryInterval

	currentBlockNumber, err := l.FetchCurrentBlockNumber(ctx)
	if err != nil {
		return defaultWaitTime
	}

	if _, err = l.CanSubmitOutput(ctx, outputIndex); err != nil {
		l.log.Error("failed to check the validator can submit output", "err", err)
		return defaultWaitTime
	}

	l.log.Info("current status before submit", "currentBlockNumber", currentBlockNumber, "nextBlockNumberToSubmit", nextBlockNumber)

	// Wait for L2 blocks proceeding when validator submission interval has not elapsed
	// Need to wait next block number to submit plus 1 because of next block hash inclusion
	nextBlockNumberToWait := new(big.Int).Add(nextBlockNumber, common.Big1)
	roundBuffer := new(big.Int).SetUint64(l.cfg.OutputSubmitterRoundBuffer)
	if currentBlockNumber.Cmp(nextBlockNumberToWait) < 0 {
		nextBlockNumberToWait = new(big.Int).Sub(nextBlockNumber, roundBuffer)
		l.log.Info("submission interval has not elapsed", "currentBlockNumber", currentBlockNumber, "nextBlockNumberToWait", nextBlockNumberToWait)
		return l.getLeftTimeForL2Blocks(currentBlockNumber, nextBlockNumberToWait)
	}

	// Check if it's a public round, or selected for priority validator
	roundInfo, err := l.fetchCurrentRound(ctx, outputIndex)
	if err != nil {
		return defaultWaitTime
	}

	if !roundInfo.canJoinRound() {
		// wait for L2 blocks proceeding until public round when not selected for priority validator
		roundIntervalToWait := new(big.Int).Sub(l.singleRoundInterval, roundBuffer)
		nextBlockNumberToWait = new(big.Int).Add(nextBlockNumber, roundIntervalToWait)
		l.log.Info("wait for next submission chance", "currentBlockNumber", currentBlockNumber, "nextBlockNumberToWait", nextBlockNumberToWait)
		return l.getLeftTimeForL2Blocks(currentBlockNumber, nextBlockNumberToWait)
	}

	// no need to wait
	return 0
}

// CanSubmitOutput checks that the validator satisfies the condition to submit L2Output.
func (l *L2OutputSubmitter) CanSubmitOutput(ctx context.Context, outputIndex *big.Int) (bool, error) {
	cCtx, cCancel := context.WithTimeout(ctx, l.cfg.NetworkTimeout)
	defer cCancel()
	from := l.cfg.TxManager.From()

	var balance, requiredBondAmount *big.Int
	if l.IsValPoolTerminated(outputIndex) {
		if isInJail, err := l.IsInJail(ctx); err != nil {
			return false, err
		} else if isInJail {
			l.log.Warn("validator is in jail")
			return false, nil
		}

		validatorStatus, err := l.GetValidatorStatus(ctx)
		if err != nil {
			return false, err
		}
		l.metr.RecordValidatorStatus(validatorStatus)

		if validatorStatus != StatusActive {
			l.log.Warn("validator is not in the status to submit output", "currentStatus", validatorStatus)
			return false, nil
		}

		balance, err = l.assetMgrContract.TotalValidatorKroNotBonded(optsutils.NewSimpleCallOpts(cCtx), from)
		if err != nil {
			return false, fmt.Errorf("failed to fetch balance: %w", err)
		}
		requiredBondAmount = l.requiredBondAmountV2
	} else {
		var err error
		balance, err = l.valPoolContract.BalanceOf(optsutils.NewSimpleCallOpts(cCtx), from)
		if err != nil {
			return false, fmt.Errorf("failed to fetch deposit amount: %w", err)
		}
		requiredBondAmount = l.requiredBondAmountV1
	}
	l.metr.RecordUnbondedDepositAmount(balance)

	// Check if the unbonded deposit amount is less than the required bond amount
	if balance.Cmp(requiredBondAmount) == -1 {
		l.log.Warn(
			"unbonded deposit is less than bond attempt amount",
			"requiredBondAmount", requiredBondAmount,
			"unbonded_deposit", balance,
		)
		return false, nil
	}

	l.log.Info("unbonded deposit amount and required bond amount",
		"unbonded_deposit", balance, "required_bond", requiredBondAmount)

	return true, nil
}

func (l *L2OutputSubmitter) FetchNextOutputIndex(ctx context.Context) (*big.Int, error) {
	cCtx, cCancel := context.WithTimeout(ctx, l.cfg.NetworkTimeout)
	defer cCancel()

	outputIndex, err := l.l2OOContract.NextOutputIndex(optsutils.NewSimpleCallOpts(cCtx))
	if err != nil {
		return nil, fmt.Errorf("failed to fetch next output index: %w", err)
	}

	return outputIndex, nil
}

func (l *L2OutputSubmitter) FetchNextBlockNumber(ctx context.Context) (*big.Int, error) {
	cCtx, cCancel := context.WithTimeout(ctx, l.cfg.NetworkTimeout)
	defer cCancel()
	nextBlockNumber, err := l.l2OOContract.NextBlockNumber(optsutils.NewSimpleCallOpts(cCtx))
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

func (l *L2OutputSubmitter) IsValPoolTerminated(outputIndex *big.Int) bool {
	return l.valPoolTerminationIndex.Cmp(outputIndex) < 0
}

func (l *L2OutputSubmitter) GetValidatorStatus(ctx context.Context) (uint8, error) {
	cCtx, cCancel := context.WithTimeout(ctx, l.cfg.NetworkTimeout)
	defer cCancel()
	validatorStatus, err := l.valMgrContract.GetStatus(optsutils.NewSimpleCallOpts(cCtx), l.cfg.TxManager.From())
	if err != nil {
		return 0, fmt.Errorf("failed to fetch the validator status: %w", err)
	}
	return validatorStatus, nil
}

func (l *L2OutputSubmitter) IsInJail(ctx context.Context) (bool, error) {
	cCtx, cCancel := context.WithTimeout(ctx, l.cfg.NetworkTimeout)
	defer cCancel()
	isInJail, err := l.valMgrContract.InJail(optsutils.NewSimpleCallOpts(cCtx), l.cfg.TxManager.From())
	if err != nil {
		return false, fmt.Errorf("failed to fetch the jail status: %w", err)
	}

	return isInJail, nil
}

func (l *L2OutputSubmitter) getLeftTimeForL2Blocks(currentBlockNumber *big.Int, targetBlockNumber *big.Int) time.Duration {
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

// fetchCurrentRound fetches next validator address from ValidatorPool or ValidatorManager contract.
// It returns if current round is public round, and if selected for priority validator if it's a priority round.
func (l *L2OutputSubmitter) fetchCurrentRound(ctx context.Context, outputIndex *big.Int) (roundInfo, error) {
	cCtx, cCancel := context.WithTimeout(ctx, l.cfg.NetworkTimeout)
	defer cCancel()
	ri := roundInfo{canJoinPublicRound: l.cfg.OutputSubmitterAllowPublicRound}

	nextValidator, err := l.getNextValidatorAddress(cCtx, outputIndex)
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

// getNextValidatorAddress selects the appropriate contract and retrieves the next validator address.
func (l *L2OutputSubmitter) getNextValidatorAddress(ctx context.Context, outputIndex *big.Int) (common.Address, error) {
	opts := optsutils.NewSimpleCallOpts(ctx)
	if l.IsValPoolTerminated(outputIndex) {
		return l.valMgrContract.NextValidator(opts)
	}
	return l.valPoolContract.NextValidator(opts)
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
func (l *L2OutputSubmitter) submitL2OutputTx(data []byte, outputIndex *big.Int) *txmgr.TxResponse {
	var name string
	var accessListAddr common.Address
	if l.IsValPoolTerminated(outputIndex) {
		name = "ValidatorManager"
		accessListAddr = l.cfg.ValidatorManagerAddr
	} else {
		name = "ValidatorPool"
		accessListAddr = l.cfg.ValidatorPoolAddr
	}

	layout, err := bindings.GetStorageLayout(name)
	if err != nil {
		return &txmgr.TxResponse{
			Receipt: nil,
			Err:     fmt.Errorf("failed to get storage layout: %w", err),
		}
	}

	var outputIndexSlot, priorityValidatorSlot common.Hash
	for _, entry := range layout.Storage {
		switch entry.Label {
		// ValidatorPool
		case "nextUnbondOutputIndex":
			outputIndexSlot = common.BigToHash(big.NewInt(int64(entry.Slot)))
		case "nextPriorityValidator":
			priorityValidatorSlot = common.BigToHash(big.NewInt(int64(entry.Slot)))
		// ValidatorManager
		case "_nextPriorityValidator":
			priorityValidatorSlot = common.BigToHash(big.NewInt(int64(entry.Slot)))
		}
	}

	var storageKeys []common.Hash
	if l.IsValPoolTerminated(outputIndex) {
		storageKeys = []common.Hash{priorityValidatorSlot}
	} else {
		storageKeys = []common.Hash{outputIndexSlot, priorityValidatorSlot}
	}

	// If provide accessList that is not actually accessed, the transaction may not be executed due to exceeding the estimated gas limit
	accessList := types.AccessList{
		types.AccessTuple{
			Address:     accessListAddr,
			StorageKeys: storageKeys,
		},
	}

	// Do the gas estimation and set 150% of it to gas limit to prevent tx failed because of dynamic gas usage in unbond and priority validator selection
	gasTipCap, baseFee, _, err := l.cfg.TxManager.SuggestGasPriceCaps(l.ctx)
	if err != nil {
		return &txmgr.TxResponse{
			Receipt: nil,
			Err:     fmt.Errorf("failed to get gas price info: %w", err),
		}
	}
	gasFeeCap := txmgr.CalcGasFeeCap(baseFee, gasTipCap)

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

func (l *L2OutputSubmitter) L2OOAbi() *abi.ABI {
	return l.l2OOABI
}
