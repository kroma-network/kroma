package validator

import (
	"bytes"
	"context"
	"fmt"
	"math/big"
	_ "net/http/pprof"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/log"

	"github.com/kroma-network/kroma/bindings/bindings"
	"github.com/kroma-network/kroma/components/node/eth"
	"github.com/kroma-network/kroma/utils"
	"github.com/kroma-network/kroma/utils/service/txmgr"
)

const (
	GasLimitMultiplier  = 150
	GasLimitDenominator = 100
)

// Guardian is responsible for validating outputs
type Guardian struct {
	log    log.Logger
	cfg    Config
	ctx    context.Context
	cancel context.CancelFunc
	wg     sync.WaitGroup

	securityCouncilContract *bindings.SecurityCouncil
	securityCouncilSub      ethereum.Subscription

	validationRequestedChan chan *bindings.SecurityCouncilValidationRequested

	txCandidatesChan chan<- txmgr.TxCandidate
}

// NewGuardian creates a new Guardian
func NewGuardian(cfg Config, l log.Logger) (*Guardian, error) {
	securityCouncilContract, err := bindings.NewSecurityCouncil(cfg.SecurityCouncilAddr, cfg.L1Client)
	if err != nil {
		return nil, err
	}

	return &Guardian{
		log:                     l,
		cfg:                     cfg,
		securityCouncilContract: securityCouncilContract,
		validationRequestedChan: make(chan *bindings.SecurityCouncilValidationRequested),
	}, nil
}

func (g *Guardian) Start(ctx context.Context, txCandidatesChan chan<- txmgr.TxCandidate) error {
	g.ctx, g.cancel = context.WithCancel(ctx)
	g.log.Info("starting guardian")

	watchOpts := &bind.WatchOpts{Context: g.ctx, Start: nil}

	g.securityCouncilSub = event.ResubscribeErr(time.Second*10, func(ctx context.Context, err error) (event.Subscription, error) {
		if err != nil {
			g.log.Warn("resubscribing after failed SecurityCouncilValidationRequested event", "err", err)
		}
		return g.securityCouncilContract.WatchValidationRequested(watchOpts, g.validationRequestedChan, nil)
	})

	g.txCandidatesChan = txCandidatesChan
	g.wg.Add(1)
	go g.handleValidationRequested(g.ctx)

	return nil
}

func (g *Guardian) Stop() error {
	g.log.Info("stopping guardian")

	if g.securityCouncilSub != nil {
		g.securityCouncilSub.Unsubscribe()
	}

	g.cancel()
	g.wg.Wait()

	close(g.validationRequestedChan)

	return nil
}

func (g *Guardian) ValidateL2Output(ctx context.Context, outputRoot eth.Bytes32, l2BlockNumber uint64) (bool, error) {
	g.log.Info("validating output...", "blockNumber", l2BlockNumber, "outputRoot", outputRoot)
	localOutputRoot, err := g.outputRootAtBlock(ctx, l2BlockNumber)
	if err != nil {
		return false, fmt.Errorf("failed to get output root at block %d: %w", l2BlockNumber, err)
	}
	isValid := bytes.Equal(outputRoot[:], localOutputRoot[:])
	return isValid, nil
}

func (g *Guardian) ConfirmTransaction(ctx context.Context, transactionId *big.Int) (*types.Transaction, error) {
	cCtx, cCancel := context.WithTimeout(ctx, g.cfg.NetworkTimeout)
	executionTx, err := g.securityCouncilContract.Transactions(&bind.CallOpts{Context: cCtx}, transactionId)
	cCancel()
	if err != nil {
		return nil, err
	}
	if executionTx.Executed {
		log.Warn("transaction %d already executed.", transactionId)
		return nil, nil
	}

	cCtx, cCancel = context.WithTimeout(ctx, g.cfg.NetworkTimeout)
	executionGas, err := g.cfg.L1Client.EstimateGas(cCtx, ethereum.CallMsg{
		From:  g.cfg.SecurityCouncilAddr,
		To:    &executionTx.Destination,
		Value: executionTx.Value,
		Data:  executionTx.Data,
	})
	cCancel()
	if err != nil {
		return nil, fmt.Errorf("failed to estimate confirmation tx gas: %w", err)
	}

	cCtx, cCancel = context.WithTimeout(ctx, g.cfg.NetworkTimeout)
	txOpts := utils.NewSimpleTxOpts(cCtx, g.cfg.TxManager.From(), g.cfg.TxManager.Signer)
	confirmTx, err := g.securityCouncilContract.ConfirmTransaction(txOpts, transactionId)
	cCancel()
	if err != nil {
		return nil, err
	}

	cCtx, cCancel = context.WithTimeout(ctx, g.cfg.NetworkTimeout)
	confirmationGas, err := g.cfg.L1Client.EstimateGas(cCtx, ethereum.CallMsg{
		From:  txOpts.From,
		To:    confirmTx.To(),
		Data:  confirmTx.Data(),
		Value: confirmTx.Value(),
	})
	cCancel()
	if err != nil {
		return nil, fmt.Errorf("failed to estimate confirmation tx gas: %w", err)
	}

	txOpts.Context = ctx
	txOpts.GasLimit = (confirmationGas * GasLimitMultiplier / GasLimitDenominator) + executionGas
	return g.securityCouncilContract.ConfirmTransaction(txOpts, transactionId)
}

func (g *Guardian) handleValidationRequested(ctx context.Context) {
	defer g.wg.Done()
	for {
		select {
		case ev := <-g.validationRequestedChan:
			g.wg.Add(1)
			go g.processOutputValidation(ctx, ev)
		case <-ctx.Done():
			return
		}
	}
}

// TODO(pangssu): Retry a failed or missed transaction.
func (g *Guardian) processOutputValidation(ctx context.Context, event *bindings.SecurityCouncilValidationRequested) {
	ticker := time.NewTicker(10 * time.Second)
	defer func() {
		ticker.Stop()
		g.wg.Done()
	}()

	for {
	Loop:
		select {
		case <-ticker.C:
			cCtx, cCancel := context.WithTimeout(ctx, g.cfg.NetworkTimeout)
			callOpts := utils.NewCallOptsWithSender(cCtx, g.cfg.TxManager.From())
			isConfirmed, err := g.securityCouncilContract.IsConfirmed(callOpts, event.TransactionId)
			cCancel()
			if err != nil {
				g.log.Error("failed to get confirmation", "transactionId", event.TransactionId, "err", err)
				break Loop
			}

			if isConfirmed {
				g.log.Info("this transaction has already been confirmed", "transactionId", event.TransactionId)
				return
			}

			isValid, err := g.ValidateL2Output(ctx, event.OutputRoot, event.L2BlockNumber.Uint64())
			if err != nil {
				g.log.Error("failed to verify the output to be replaced", "transactionId", event.TransactionId, "err", err)
				break Loop
			}

			if isValid {
				g.log.Info("the output to be replaced is valid")

				cCtx, cCancel := context.WithTimeout(ctx, g.cfg.NetworkTimeout)
				tx, err := g.ConfirmTransaction(cCtx, event.TransactionId)
				cCancel()
				if err != nil {
					g.log.Error("failed to create confirm tx", "transactionId", event.TransactionId, "err", err)
					break Loop
				}
				if tx == nil {
					g.log.Error("confirm tx is nil", "transactionId", event.TransactionId, "err", err)
					return
				}

				g.sendTransaction(tx)
			} else {
				g.log.Info("ignore this challenge because the output to be replaced is invalid")
			}
			return
		case <-ctx.Done():
			return
		}
	}
}

func (g *Guardian) outputRootAtBlock(ctx context.Context, blockNumber uint64) (eth.Bytes32, error) {
	cCtx, cCancel := context.WithTimeout(ctx, g.cfg.NetworkTimeout)
	defer cCancel()
	output, err := g.cfg.RollupClient.OutputAtBlock(cCtx, blockNumber)
	if err != nil {
		return eth.Bytes32{}, err
	}
	return output.OutputRoot, nil
}

func (g *Guardian) sendTransaction(tx *types.Transaction) {
	g.txCandidatesChan <- txmgr.TxCandidate{
		TxData:     tx.Data(),
		To:         tx.To(),
		GasLimit:   tx.Gas(),
		AccessList: nil,
	}
}
