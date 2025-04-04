package validator

import (
	"bytes"
	"context"
	"fmt"
	"math/big"
	_ "net/http/pprof"
	"sync"
	"time"

	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum-optimism/optimism/op-service/optsutils"
	"github.com/ethereum-optimism/optimism/op-service/watcher"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/log"

	"github.com/kroma-network/kroma/kroma-bindings/bindings"
	"github.com/kroma-network/kroma/kroma-validator/challenge"
)

// Guardian is responsible for validating outputs.
type Guardian struct {
	log    log.Logger
	cfg    Config
	ctx    context.Context
	cancel context.CancelFunc
	wg     sync.WaitGroup

	l2ooContract            *bindings.L2OutputOracle
	securityCouncilContract *bindings.SecurityCouncil
	colosseumContract       *bindings.Colosseum
	colosseumABI            *abi.ABI

	l1BlockTime             *big.Int
	l2BlockTime             *big.Int
	guardianPeriod          *big.Int
	maxClockDurationSeconds *big.Int
	challengeGracePeriod    *big.Int

	validationRequestedSub ethereum.Subscription
	deletionRequestedSub   ethereum.Subscription
	challengeCreatedSub    ethereum.Subscription

	validationRequestedChan chan *bindings.SecurityCouncilValidationRequested
	deletionRequestedChan   chan *bindings.SecurityCouncilDeletionRequested
	challengeCreatedChan    chan *bindings.ColosseumChallengeCreated

	checkpoint *big.Int
}

// NewGuardian creates a new Guardian.
func NewGuardian(cfg Config, l log.Logger) (*Guardian, error) {
	securityCouncilContract, err := bindings.NewSecurityCouncil(cfg.SecurityCouncilAddr, cfg.L1Client)
	if err != nil {
		return nil, err
	}

	l2ooContract, err := bindings.NewL2OutputOracle(cfg.L2OutputOracleAddr, cfg.L1Client)
	if err != nil {
		return nil, err
	}

	colosseumABI, err := bindings.ColosseumMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	colosseumContract, err := bindings.NewColosseum(cfg.ColosseumAddr, cfg.L1Client)
	if err != nil {
		return nil, err
	}

	return &Guardian{
		log:                     l.New("service", "guardian"),
		cfg:                     cfg,
		securityCouncilContract: securityCouncilContract,
		l2ooContract:            l2ooContract,
		colosseumContract:       colosseumContract,
		colosseumABI:            colosseumABI,
		l1BlockTime:             big.NewInt(12),
	}, nil
}

func (g *Guardian) InitConfig(ctx context.Context) error {
	contractWatcher := watcher.NewContractWatcher(ctx, g.cfg.L1Client, g.log)

	err := contractWatcher.WatchUpgraded(g.cfg.L2OutputOracleAddr, func() error {
		cCtx, cCancel := context.WithTimeout(ctx, g.cfg.NetworkTimeout)
		defer cCancel()
		l2BlockTime, err := g.l2ooContract.L2BLOCKTIME(optsutils.NewSimpleCallOpts(cCtx))
		if err != nil {
			return fmt.Errorf("failed to get l2 block time: %w", err)
		}
		g.l2BlockTime = l2BlockTime

		return nil
	})
	if err != nil {
		return fmt.Errorf("failed to initiate l2oo config: %w", err)
	}

	err = contractWatcher.WatchUpgraded(g.cfg.ColosseumAddr, func() error {
		cCtx, cCancel := context.WithTimeout(ctx, g.cfg.NetworkTimeout)
		defer cCancel()
		guardianPeriod, err := g.colosseumContract.GUARDIANPERIOD(optsutils.NewSimpleCallOpts(cCtx))
		if err != nil {
			return fmt.Errorf("failed to get creation period seconds: %w", err)
		}
		g.guardianPeriod = guardianPeriod

		cCtx, cCancel = context.WithTimeout(ctx, g.cfg.NetworkTimeout)
		defer cCancel()
		maxClockDurationSeconds, err := g.colosseumContract.MAXCLOCKDURATIONSECONDS(optsutils.NewSimpleCallOpts(cCtx))
		if err != nil {
			return fmt.Errorf("failed to get creation period seconds: %w", err)
		}
		g.maxClockDurationSeconds = maxClockDurationSeconds

		cCtx, cCancel = context.WithTimeout(ctx, g.cfg.NetworkTimeout)
		defer cCancel()
		challengeGracePeriod, err := g.colosseumContract.CHALLENGEGRACEPERIOD(optsutils.NewSimpleCallOpts(cCtx))
		if err != nil {
			return fmt.Errorf("failed to get creation period seconds: %w", err)
		}
		g.challengeGracePeriod = challengeGracePeriod

		return nil
	})
	if err != nil {
		return fmt.Errorf("failed to initiate colosseum config: %w", err)
	}

	return nil
}

func (g *Guardian) Start(ctx context.Context) error {
	g.ctx, g.cancel = context.WithCancel(ctx)

	if err := g.InitConfig(g.ctx); err != nil {
		return err
	}
	g.initSub()

	g.wg.Add(1)
	go g.scanPrevChallenges()

	g.wg.Add(1)
	go g.subscriptionLoop()

	g.wg.Add(1)
	go g.inspectorLoop()

	return nil
}

func (g *Guardian) Stop() error {
	g.validationRequestedSub.Unsubscribe()
	g.deletionRequestedSub.Unsubscribe()
	g.challengeCreatedSub.Unsubscribe()

	g.cancel()
	g.wg.Wait()

	close(g.validationRequestedChan)
	close(g.deletionRequestedChan)
	close(g.challengeCreatedChan)

	return nil
}

func (g *Guardian) initSub() {
	opts := optsutils.NewSimpleWatchOpts(g.ctx)

	g.validationRequestedChan = make(chan *bindings.SecurityCouncilValidationRequested)
	g.validationRequestedSub = event.ResubscribeErr(time.Second*10, func(ctx context.Context, err error) (event.Subscription, error) {
		if err != nil {
			g.log.Warn("resubscribing after failed ValidationRequested event", "err", err)
		}
		return g.securityCouncilContract.WatchValidationRequested(opts, g.validationRequestedChan, nil)
	})

	g.deletionRequestedChan = make(chan *bindings.SecurityCouncilDeletionRequested)
	g.deletionRequestedSub = event.ResubscribeErr(time.Second*10, func(ctx context.Context, err error) (event.Subscription, error) {
		if err != nil {
			g.log.Warn("resubscribing after failed DeletionRequested event", "err", err)
		}
		return g.securityCouncilContract.WatchDeletionRequested(opts, g.deletionRequestedChan, nil, nil)
	})

	g.challengeCreatedChan = make(chan *bindings.ColosseumChallengeCreated)
	g.challengeCreatedSub = event.ResubscribeErr(time.Second*10, func(ctx context.Context, err error) (event.Subscription, error) {
		if err != nil {
			g.log.Warn("resubscribing after failed ChallengeCreated event", "err", err)
		}
		return g.colosseumContract.WatchChallengeCreated(opts, g.challengeCreatedChan, nil, nil, nil)
	})
}

// scanPrevChallenges scans all the previous challenges before current L1 head within the finalization window to handle challenger timeout.
func (g *Guardian) scanPrevChallenges() {
	ticker := time.NewTicker(g.cfg.GuardianPollInterval)
	defer func() {
		ticker.Stop()
		g.wg.Done()
	}()

	for ; ; <-ticker.C {
		select {
		case <-g.ctx.Done():
			return
		default:
			status, err := g.cfg.RollupClient.SyncStatus(g.ctx)
			if err != nil {
				g.log.Error("failed to get sync status", "err", err)
				continue
			}

			toBlock := new(big.Int).SetUint64(status.HeadL1.Number)

			latestFinalizeOutput, err := g.l2ooContract.GetLatestFinalizeOutput(optsutils.NewSimpleCallOpts(g.ctx))
			if err != nil {
				g.log.Error("failed to get latest finalize output", "err", err)
				continue
			}

			output, err := g.cfg.RollupClient.OutputAtBlock(g.ctx, latestFinalizeOutput.L2BlockNumber.Uint64()+1)
			if err != nil {
				g.log.Error("failed to get local output", "err", err)
				continue
			}

			l1OriginNumber := output.BlockRef.L1Origin.Number
			fromBlock := new(big.Int).SetUint64(l1OriginNumber)

			challengeCreatedEvent := g.colosseumABI.Events[KeyEventChallengeCreated]

			addresses := []common.Address{g.cfg.ColosseumAddr}
			topics := []common.Hash{challengeCreatedEvent.ID}

			query := ethereum.FilterQuery{
				FromBlock: fromBlock,
				ToBlock:   toBlock,
				Addresses: addresses,
				Topics:    [][]common.Hash{topics},
			}

			logs, err := g.cfg.L1Client.FilterLogs(g.ctx, query)
			if err != nil {
				g.log.Error("failed to get event logs related to challenge creation", "err", err)
				continue
			}

			for _, vLog := range logs {
				ev, err := NewChallengeCreatedEvent(vLog)
				if err != nil {
					g.log.Error("failed to parse challenge created event", "err", err)
					continue
				}
				g.wg.Add(1)
				go g.processChallengerTimeout(ev)
			}

			return
		}
	}
}

// inspectorLoop finds and deletes outputs whose zk fault proving has failed due to an undeniable bug
// among whose creation period has passed but not finalized
func (g *Guardian) inspectorLoop() {
	defer g.wg.Done()

	ticker := time.NewTicker(g.cfg.GuardianPollInterval)
	defer ticker.Stop()

	for ; ; <-ticker.C {
		select {
		case <-g.ctx.Done():
			return
		default:
			if g.checkpoint == nil {
				func() {
					cCtx, cCancel := context.WithTimeout(g.ctx, g.cfg.NetworkTimeout)
					defer cCancel()
					startOutputIndex, err := g.l2ooContract.GetLatestFinalizeOutputIndex(optsutils.NewSimpleCallOpts(cCtx))
					if err != nil {
						g.log.Error("failed to get output index after", "err", err)
						return
					}

					cCtx, cCancel = context.WithTimeout(g.ctx, g.cfg.NetworkTimeout)
					defer cCancel()
					endOutputIndex, err := g.l2ooContract.LatestOutputIndex(optsutils.NewSimpleCallOpts(cCtx))
					if err != nil {
						g.log.Error("failed to get output index after", "err", err)
						return
					}

					for i := new(big.Int).Add(startOutputIndex, common.Big1); i.Cmp(endOutputIndex) <= 0; i.Add(i, common.Big1) {
						g.wg.Add(1)
						go g.inspectOutput(new(big.Int).Set(i))
					}
					g.checkpoint = endOutputIndex
				}()
			} else {
				func() {
					cCtx, cCancel := context.WithTimeout(g.ctx, g.cfg.NetworkTimeout)
					defer cCancel()
					outputIndex, err := g.l2ooContract.LatestOutputIndex(optsutils.NewSimpleCallOpts(cCtx))
					if err != nil {
						g.log.Error("failed to get output index after", "err", err)
						return
					}

					for i := new(big.Int).Add(g.checkpoint, common.Big1); i.Cmp(outputIndex) < 0; i.Add(i, common.Big1) {
						g.wg.Add(1)
						go g.inspectOutput(new(big.Int).Set(i))
					}
					g.checkpoint = outputIndex.Sub(outputIndex, common.Big1)
				}()
			}
		}
	}
}

// inspectOutput inspects if the output fails zk fault proof due to an undeniable bug.
func (g *Guardian) inspectOutput(outputIndex *big.Int) {
	g.log.Info("inspect output if there is an undeniable bug", "outputIndex", outputIndex)
	defer g.wg.Done()

	ticker := time.NewTicker(g.cfg.GuardianPollInterval)
	defer ticker.Stop()

	for ; ; <-ticker.C {
		select {
		case <-g.ctx.Done():
			return
		default:
			inGuardianPeriod, retry, err := g.isInGuardianPeriod(outputIndex)
			if err != nil {
				g.log.Error("unable to fetch if output is in guardian period", "err", err, "outputIndex", outputIndex)
				continue
			}
			if !inGuardianPeriod {
				if retry {
					continue
				} else {
					return
				}
			}

			shouldBeDeleted, err := g.shouldBeDeleted(outputIndex)
			if err != nil {
				g.log.Error("unable to inspect the output for force deletion", "err", err, "outputIndex", outputIndex)
				continue
			}
			if !shouldBeDeleted {
				g.log.Info("no need to delete output forcefully", "outputIndex", outputIndex)
				return
			}

			tx, err := g.RequestDeletion(g.ctx, outputIndex)
			if err != nil {
				g.log.Error("failed to create tx for output deletion", "err", err, "outputIndex", outputIndex)
				continue
			}

			if txResponse := g.cfg.TxManager.SendTransaction(g.ctx, tx); txResponse.Err != nil {
				g.log.Error("failed to send deletion request tx", "err", txResponse.Err, "outputIndex", outputIndex)
				continue
			}

			return
		}
	}
}

// subscriptionLoop handles event subscriptions.
func (g *Guardian) subscriptionLoop() {
	defer g.wg.Done()

	for {
		select {
		case ev := <-g.validationRequestedChan:
			g.wg.Add(1)
			go g.processOutputValidation(ev)
		case ev := <-g.deletionRequestedChan:
			g.wg.Add(1)
			go g.processOutputDeletion(ev)
		case ev := <-g.challengeCreatedChan:
			g.wg.Add(1)
			go g.processChallengerTimeout(ev)
		case <-g.ctx.Done():
			return
		}
	}
}

// processOutputValidation validates the deleted output and sends confirm tx when multi sig tx that requires confirmation is created.
func (g *Guardian) processOutputValidation(event *bindings.SecurityCouncilValidationRequested) {
	g.log.Info("processing validation of the deleted output", "l2BlockNumber", event.L2BlockNumber, "outputRoot", event.OutputRoot, "transactionId", event.TransactionId)

	ticker := time.NewTicker(10 * time.Second)
	defer func() {
		ticker.Stop()
		g.wg.Done()
	}()

	for ; ; <-ticker.C {
		select {
		case <-g.ctx.Done():
			return
		default:
			if err := g.tryConfirmRequestValidationTx(event); err != nil {
				g.log.Error("failed to create confirmation tx for output validation request", "err", err, "transactionId", event.TransactionId.String())
				continue
			}
			return
		}
	}
}

// processOutputDeletion validates the requested output deletion and sends confirm tx when multi sig tx that requires confirmation is created.
func (g *Guardian) processOutputDeletion(event *bindings.SecurityCouncilDeletionRequested) {
	g.log.Info("processing validation of the output to be deleted", "outputIndex", event.OutputIndex, "transactionId", event.TransactionId)

	ticker := time.NewTicker(10 * time.Second)
	defer func() {
		ticker.Stop()
		g.wg.Done()
	}()

	for ; ; <-ticker.C {
		select {
		case <-g.ctx.Done():
			return
		default:
			if err := g.tryConfirmRequestDeletionTx(event); err != nil {
				g.log.Error("failed to create confirmation tx for output deletion request", "err", err, "transactionId", event.TransactionId.String())
				continue
			}
			return
		}
	}
}

// processChallengerTimeout tracks created challenges and sends challenger timeout txs when challenger timed out.
func (g *Guardian) processChallengerTimeout(event *bindings.ColosseumChallengeCreated) {
	outputIndex, challenger := event.OutputIndex, event.Challenger
	g.log.Info("processing challenge to call challenger timeout", "outputIndex", outputIndex, "asserter", event.Asserter, "challenger", challenger)

	ticker := time.NewTicker(10 * time.Second)
	defer func() {
		ticker.Stop()
		g.wg.Done()
	}()

	for ; ; <-ticker.C {
		select {
		case <-g.ctx.Done():
			return
		default:
			inGuardianPeriod, retry, err := g.isInGuardianPeriod(outputIndex)
			if err != nil {
				g.log.Error("unable to fetch if output in guardian period", "err", err, "outputIndex", outputIndex)
				continue
			}
			if !inGuardianPeriod {
				if retry {
					continue
				} else {
					return
				}
			}

			retry, err = g.tryChallengerTimeoutTx(outputIndex, challenger)
			if err != nil {
				g.log.Error("failed to try challenge timeout tx", "err", err, "outputIndex", outputIndex, "challenger", challenger)
				continue
			}
			if retry {
				continue
			} else {
				return
			}
		}
	}
}

func (g *Guardian) tryConfirmRequestValidationTx(event *bindings.SecurityCouncilValidationRequested) error {
	outputIndex, err := g.getL2OutputIndexAfter(event.L2BlockNumber)
	if err != nil {
		return fmt.Errorf("failed to get output index after. (l2BlockNumber: %d): %w", event.L2BlockNumber.Int64(), err)
	}

	needConfirm, err := g.CheckConfirmCondition(g.ctx, event.TransactionId, outputIndex)
	if err != nil {
		return fmt.Errorf("failed to check confirm condition. (transactionId: %d): %w", event.TransactionId.Int64(), err)
	}
	if !needConfirm {
		return nil
	}

	isValid, err := g.ValidateL2Output(g.ctx, event.OutputRoot, event.L2BlockNumber.Uint64())
	if err != nil {
		return fmt.Errorf("failed to validate the deleted output. (transactionId: %d): %w", event.TransactionId.Int64(), err)
	}

	if isValid {
		g.log.Info("the deleted output is equal to guardian's output but deleted incorrectly, so confirm to dismiss challenge")

		tx, err := g.ConfirmTransaction(g.ctx, event.TransactionId)
		if err != nil {
			return fmt.Errorf("failed to create confirm tx. (transactionId: %d): %w", event.TransactionId.Int64(), err)
		}

		if txResponse := g.cfg.TxManager.SendTransaction(g.ctx, tx); txResponse.Err != nil {
			return fmt.Errorf("failed to send confirm tx. (transactionId: %d): %w", event.TransactionId.Int64(), txResponse.Err)
		}
	} else {
		g.log.Info("do nothing because the deleted output is not equal to guardian's output so deleted correctly")
	}

	return nil
}

func (g *Guardian) tryConfirmRequestDeletionTx(event *bindings.SecurityCouncilDeletionRequested) error {
	needConfirm, err := g.CheckConfirmCondition(g.ctx, event.TransactionId, event.OutputIndex)
	if err != nil {
		return fmt.Errorf("failed to check confirm condition. (transactionId: %d): %w", event.TransactionId.Int64(), err)
	}

	if !needConfirm {
		return nil
	}

	cCtx, cCancel := context.WithTimeout(g.ctx, g.cfg.NetworkTimeout)
	defer cCancel()
	output, err := g.l2ooContract.GetL2Output(optsutils.NewSimpleCallOpts(cCtx), event.OutputIndex)
	if err != nil {
		return fmt.Errorf("failed to get output from L2OutputOracle contract(outputIndex: %d): %w", event.OutputIndex.Uint64(), err)
	}

	isValid, err := g.ValidateL2Output(g.ctx, output.OutputRoot, output.L2BlockNumber.Uint64())
	if err != nil {
		return err
	}

	if isValid {
		g.log.Info("output deletion is requested, but it's valid", "transactionId", event.TransactionId, "outputIndex", event.OutputIndex)
		return nil
	}

	tx, err := g.ConfirmTransaction(g.ctx, event.TransactionId)
	if err != nil {
		return fmt.Errorf("failed to create confirm tx. (transactionId: %d): %w", event.TransactionId.Int64(), err)
	}

	if txResponse := g.cfg.TxManager.SendTransaction(g.ctx, tx); txResponse.Err != nil {
		return fmt.Errorf("failed to send confirm tx. (transactionId: %d): %w", event.TransactionId.Int64(), txResponse.Err)
	}

	return nil
}

func (g *Guardian) CheckConfirmCondition(ctx context.Context, transactionId *big.Int, outputIndex *big.Int) (bool, error) {
	outputFinalized, err := g.isOutputFinalized(ctx, outputIndex)
	if err != nil {
		return true, fmt.Errorf("failed to get if output is finalized. (outputIndex: %d): %w", outputIndex.Int64(), err)
	}
	if outputFinalized {
		g.log.Info("output is already finalized", "outputIndex", outputIndex)
		return false, nil
	}

	isConfirmed, err := g.isTransactionConfirmed(ctx, transactionId)
	if err != nil {
		return true, fmt.Errorf("failed to get confirmation. (transactionId: %d): %w", transactionId.Int64(), err)
	}
	if isConfirmed {
		g.log.Info("transaction is already confirmed", "transactionId", transactionId)
		return false, nil
	}

	cCtx, cCancel := context.WithTimeout(ctx, g.cfg.NetworkTimeout)
	defer cCancel()
	executionTx, err := g.securityCouncilContract.Transactions(optsutils.NewSimpleCallOpts(cCtx), transactionId)
	if err != nil {
		return true, fmt.Errorf("failed to get transaction with transactionId %d: %w", transactionId.Int64(), err)
	}
	if executionTx.Executed {
		g.log.Info("transaction is already executed", "transactionId", transactionId)
		return false, nil
	}

	return true, nil
}

func (g *Guardian) isInGuardianPeriod(outputIndex *big.Int) (inGuardianPeriod bool, retry bool, err error) {
	isFinalized, err := g.isOutputFinalized(g.ctx, outputIndex)
	if err != nil {
		return false, true, fmt.Errorf("unable to check if the output is finalized or not: %w", err)
	}
	// outputs that have been finalized are not target
	if isFinalized {
		g.log.Info("the output is finalized, no need to handle", "outputIndex", outputIndex)
		return false, false, nil
	}

	cCtx, cCancel := context.WithTimeout(g.ctx, g.cfg.NetworkTimeout)
	defer cCancel()

	assertion, err := g.colosseumContract.GetAssertion(optsutils.NewSimpleCallOpts(cCtx), outputIndex)

	if err != nil {
		return false, true, fmt.Errorf("unable to check if the output is in challenge creation period or not: %w", err)
	}

	if assertion.IsEnforced {
		g.log.Info("the output has enforced, no need to handle", "outputIndex", outputIndex)
		return false, false, nil
	}

	if assertion.AssertedAt.Cmp(common.Big0) == 0 {
		assertionStatus, err := g.colosseumContract.GetAssertionStatus(optsutils.NewSimpleCallOpts(cCtx), outputIndex)
		if err != nil {
			return false, true, fmt.Errorf("unable to check if the output is in challenge creation period or not: %w", err)
		}

		if assertionStatus != challenge.StatusChallengerTimeout {
			g.log.Info("guardian period has not started yet", "outputIndex", outputIndex)
			return false, true, nil
		}

	}

	return true, false, nil
}

func (g *Guardian) tryChallengerTimeoutTx(outputIndex *big.Int, challenger common.Address) (retry bool, err error) {
	status, err := g.getChallengeStatus(outputIndex, challenger)
	if err != nil {
		return true, fmt.Errorf("unable to fetch challenge status: %w", err)
	}

	if status == challenge.StatusNone {
		g.log.Info("challenge status is None, no need to call challenger timeout", "outputIndex", outputIndex, "challenger", challenger)
		return false, nil
	} else if status == challenge.StatusChallengerTimeout {
		g.log.Info("challenge status is ChallengerTimout, call challenger timeout", "outputIndex", outputIndex, "challenger", challenger)

		tx, err := g.challengerTimeout(g.ctx, outputIndex, challenger)
		if err != nil {
			return true, fmt.Errorf("failed to create challenger timeout tx: %w", err)
		}

		if txResponse := g.cfg.TxManager.SendTransaction(g.ctx, tx); txResponse.Err != nil {
			return true, fmt.Errorf("failed to send challenger timeout tx: %w", txResponse.Err)
		}

		return false, nil
	} else {
		return true, nil
	}
}

// shouldBeDeleted checks whether the output should have been deleted.
// This function should only be called during the guardian period,
// and it compares the target output to the local output of the guardian.
func (g *Guardian) shouldBeDeleted(outputIndex *big.Int) (bool, error) {
	cCtx, cCancel := context.WithTimeout(g.ctx, g.cfg.NetworkTimeout)
	defer cCancel()
	output, err := g.l2ooContract.GetL2Output(optsutils.NewSimpleCallOpts(cCtx), outputIndex)
	if err != nil {
		return false, fmt.Errorf("failed to get output from L2OutputOracle contract(outputIndex: %d): %w", outputIndex.Uint64(), err)
	}

	if IsOutputDeleted(output.OutputRoot) {
		g.log.Info("output has already been deleted", "outputIndex", outputIndex)
		return false, nil
	}

	isValid, err := g.ValidateL2Output(g.ctx, output.OutputRoot, output.L2BlockNumber.Uint64())
	if err != nil {
		return false, err
	}

	return !isValid, nil
}

func (g *Guardian) ValidateL2Output(ctx context.Context, outputRoot eth.Bytes32, l2BlockNumber uint64) (bool, error) {
	g.log.Info("validating deleted output as a result of challenge...", "l2BlockNumber", l2BlockNumber, "outputRoot", outputRoot)
	localOutputRoot, err := g.OutputRootAtBlock(ctx, l2BlockNumber)
	if err != nil {
		return false, fmt.Errorf("failed to get output root at block number %d: %w", l2BlockNumber, err)
	}
	isValid := bytes.Equal(outputRoot[:], localOutputRoot[:])
	return isValid, nil
}

func (g *Guardian) ConfirmTransaction(ctx context.Context, transactionId *big.Int) (*types.Transaction, error) {
	g.log.Info("crafting confirm tx", "transactionId", transactionId)
	txOpts := optsutils.NewSimpleTxOpts(ctx, g.cfg.TxManager.From(), g.cfg.TxManager.Signer)
	return g.securityCouncilContract.ConfirmTransaction(txOpts, transactionId)
}

func (g *Guardian) RequestDeletion(ctx context.Context, outputIndex *big.Int) (*types.Transaction, error) {
	g.log.Info("crafting requestDeletion tx", "outputIndex", outputIndex)
	txOpts := optsutils.NewSimpleTxOpts(ctx, g.cfg.TxManager.From(), g.cfg.TxManager.Signer)
	return g.securityCouncilContract.RequestDeletion(txOpts, outputIndex, false)
}

func (g *Guardian) challengerTimeout(ctx context.Context, outputIndex *big.Int, challenger common.Address) (*types.Transaction, error) {
	g.log.Info("crafting challenger timeout tx", "outputIndex", outputIndex, "challenger", challenger)
	txOpts := optsutils.NewSimpleTxOpts(ctx, g.cfg.TxManager.From(), g.cfg.TxManager.Signer)
	return g.colosseumContract.ChallengerTimeout(txOpts, outputIndex, challenger)
}

func (g *Guardian) OutputRootAtBlock(ctx context.Context, l2BlockNumber uint64) (eth.Bytes32, error) {
	cCtx, cCancel := context.WithTimeout(ctx, g.cfg.NetworkTimeout)
	defer cCancel()
	output, err := g.cfg.RollupClient.OutputAtBlock(cCtx, l2BlockNumber)
	if err != nil {
		return eth.Bytes32{}, err
	}
	return output.OutputRoot, nil
}

func (g *Guardian) getL2OutputIndexAfter(l2BlockNumber *big.Int) (*big.Int, error) {
	cCtx, cCancel := context.WithTimeout(g.ctx, g.cfg.NetworkTimeout)
	defer cCancel()
	return g.l2ooContract.GetL2OutputIndexAfter(optsutils.NewSimpleCallOpts(cCtx), l2BlockNumber)
}

func (g *Guardian) isOutputFinalized(ctx context.Context, outputIndex *big.Int) (bool, error) {
	cCtx, cCancel := context.WithTimeout(ctx, g.cfg.NetworkTimeout)
	defer cCancel()
	return g.l2ooContract.IsFinalized(optsutils.NewSimpleCallOpts(cCtx), outputIndex)
}

func (g *Guardian) isTransactionConfirmed(ctx context.Context, transactionId *big.Int) (bool, error) {
	cCtx, cCancel := context.WithTimeout(ctx, g.cfg.NetworkTimeout)
	defer cCancel()
	return g.securityCouncilContract.IsConfirmed(optsutils.NewSimpleCallOpts(cCtx), transactionId)
}

func (g *Guardian) getChallengeStatus(outputIndex *big.Int, challenger common.Address) (uint8, error) {
	cCtx, cCancel := context.WithTimeout(g.ctx, g.cfg.NetworkTimeout)
	defer cCancel()
	return g.colosseumContract.GetStatus(optsutils.NewSimpleCallOpts(cCtx), outputIndex, challenger)
}
