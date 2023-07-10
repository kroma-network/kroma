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

func (g *Guardian) Start(ctx context.Context) error {
	g.ctx, g.cancel = context.WithCancel(ctx)
	g.log.Info("starting guardian")

	watchOpts := &bind.WatchOpts{Context: g.ctx, Start: nil}

	g.securityCouncilSub = event.ResubscribeErr(time.Second*10, func(ctx context.Context, err error) (event.Subscription, error) {
		if err != nil {
			g.log.Warn("resubscribing after failed SecurityCouncilValidationRequested event", "err", err)
		}
		return g.securityCouncilContract.WatchValidationRequested(watchOpts, g.validationRequestedChan, nil)
	})

	g.wg.Add(1)
	go g.handleValidationRequested(g.ctx)

	return nil
}

func (g *Guardian) Stop() error {
	g.log.Info("stopping guardian")

	if g.securityCouncilSub != nil {
		g.securityCouncilSub.Unsubscribe()
	}
	if g.cancel != nil {
		g.cancel()
	}
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
	txOpts := utils.NewSimpleTxOpts(ctx, g.cfg.TxManager.From(), g.cfg.TxManager.Signer)
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

func (g *Guardian) tryConfirmTransaction(ctx context.Context, event *bindings.SecurityCouncilValidationRequested) error {
	cCtx, cCancel := context.WithTimeout(ctx, g.cfg.NetworkTimeout)
	callOpts := utils.NewCallOptsWithSender(cCtx, g.cfg.TxManager.From())
	isConfirmed, err := g.securityCouncilContract.IsConfirmed(callOpts, event.TransactionId)
	cCancel()
	if err != nil {
		g.log.Error("failed to get confirmation", "transactionId", event.TransactionId, "err", err)
		return err
	}

	if isConfirmed {
		g.log.Info("this transaction has already been confirmed", "transactionId", event.TransactionId)
		return nil
	}

	isValid, err := g.ValidateL2Output(ctx, event.OutputRoot, event.L2BlockNumber.Uint64())
	if err != nil {
		g.log.Error("failed to verify the output to be replaced", "transactionId", event.TransactionId, "err", err)
		return err
	}

	if isValid {
		g.log.Info("the output to be replaced is valid")

		cCtx, cCancel := context.WithTimeout(ctx, g.cfg.NetworkTimeout)
		tx, err := g.ConfirmTransaction(cCtx, event.TransactionId)
		cCancel()
		if err != nil {
			g.log.Error("failed to create confirm tx", "transactionId", event.TransactionId, "err", err)
			return err
		}
		if tx == nil {
			return nil
		}

		if err := g.sendTransaction(ctx, tx); err != nil {
			return err
		}
	}

	return nil
}

func (g *Guardian) processOutputValidation(ctx context.Context, event *bindings.SecurityCouncilValidationRequested) {
	ticker := time.NewTicker(10 * time.Second)
	defer func() {
		ticker.Stop()
		g.wg.Done()
	}()

	for ; ; <-ticker.C {
		select {
		case <-ctx.Done():
			return
		default:
			err := g.tryConfirmTransaction(ctx, event)
			if err != nil {
				continue
			}
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

func (g *Guardian) sendTransaction(ctx context.Context, tx *types.Transaction) error {
	txResponse := g.cfg.TxManager.SubmitTransaction(ctx, txmgr.TxCandidate{
		TxData:   tx.Data(),
		To:       tx.To(),
		GasLimit: 0,
	})
	if txResponse.Err != nil {
		g.log.Error("failed to send tx with response", "err", txResponse.Err)
		return txResponse.Err
	}

	return nil
}
