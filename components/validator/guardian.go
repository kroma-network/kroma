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
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/log"

	"github.com/kroma-network/kroma/bindings/bindings"
	"github.com/kroma-network/kroma/components/node/eth"
	"github.com/kroma-network/kroma/utils"
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
	securityCouncilSub      ethereum.Subscription

	validationRequestedChan chan *bindings.SecurityCouncilValidationRequested
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

	return &Guardian{
		log:                     l,
		cfg:                     cfg,
		securityCouncilContract: securityCouncilContract,
		l2ooContract:            l2ooContract,
		validationRequestedChan: make(chan *bindings.SecurityCouncilValidationRequested),
	}, nil
}

func (g *Guardian) Start(ctx context.Context) error {
	g.ctx, g.cancel = context.WithCancel(ctx)

	watchOpts := utils.NewSimpleWatchOpts(g.ctx)
	g.securityCouncilSub = event.ResubscribeErr(time.Second*10, func(ctx context.Context, err error) (event.Subscription, error) {
		if err != nil {
			g.log.Warn("resubscribing after failed ValidationRequested event", "err", err)
		}
		return g.securityCouncilContract.WatchValidationRequested(watchOpts, g.validationRequestedChan, nil)
	})

	g.wg.Add(1)
	go g.handleValidationRequested()

	return nil
}

func (g *Guardian) Stop() error {
	if g.securityCouncilSub != nil {
		g.securityCouncilSub.Unsubscribe()
	}

	g.cancel()
	g.wg.Wait()
	close(g.validationRequestedChan)

	return nil
}

func (g *Guardian) handleValidationRequested() {
	defer g.wg.Done()

	for {
		select {
		case ev := <-g.validationRequestedChan:
			g.wg.Add(1)
			go g.processOutputValidation(ev)
		case <-g.ctx.Done():
			return
		}
	}
}

func (g *Guardian) processOutputValidation(event *bindings.SecurityCouncilValidationRequested) {
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
			err := g.tryConfirmTransaction(event)
			if err != nil {
				g.log.Error(err.Error())
				continue
			}
			return
		}
	}
}

func (g *Guardian) tryConfirmTransaction(event *bindings.SecurityCouncilValidationRequested) error {
	needConfirm, err := g.checkConfirmCondition(event.TransactionId, event.L2BlockNumber)
	if err != nil {
		return fmt.Errorf("failed to check confirm condition. (transactionId: %d): %w", event.TransactionId.Int64(), err)
	}
	if !needConfirm {
		return nil
	}

	isEqual, err := g.ValidateL2Output(g.ctx, event.OutputRoot, event.L2BlockNumber.Uint64())
	if err != nil {
		return fmt.Errorf("failed to validate the deleted output. (transactionId: %d): %w", event.TransactionId.Int64(), err)
	}
	if isEqual {
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

func (g *Guardian) checkConfirmCondition(transactionId *big.Int, l2BlockNumber *big.Int) (bool, error) {
	outputIndex, err := g.getL2OutputIndexAfter(l2BlockNumber)
	if err != nil {
		return true, fmt.Errorf("failed to get output index after. (l2BlockNumber: %d): %w", l2BlockNumber.Int64(), err)
	}

	outputFinalized, err := g.IsOutputFinalized(g.ctx, outputIndex)
	if err != nil {
		return true, fmt.Errorf("failed to get if output is finalized. (outputIndex: %d): %w", outputIndex.Int64(), err)
	}
	if outputFinalized {
		g.log.Info("output is already finalized", "outputIndex", outputIndex)
		return false, nil
	}

	isConfirmed, err := g.isTransactionConfirmed(transactionId)
	if err != nil {
		return true, fmt.Errorf("failed to get confirmation. (transactionId: %d): %w", transactionId.Int64(), err)
	}
	if isConfirmed {
		g.log.Info("transaction is already confirmed", "transactionId", transactionId)
		return false, nil
	}

	cCtx, cCancel := context.WithTimeout(g.ctx, g.cfg.NetworkTimeout)
	executionTx, err := g.securityCouncilContract.Transactions(utils.NewSimpleCallOpts(cCtx), transactionId)
	cCancel()
	if err != nil {
		return true, fmt.Errorf("failed to get transaction with transactionId %d: %w", transactionId.Int64(), err)
	}
	if executionTx.Executed {
		g.log.Info("transaction is already executed", "transactionId", transactionId)
		return false, nil
	}

	return true, nil
}

func (g *Guardian) ValidateL2Output(ctx context.Context, outputRoot eth.Bytes32, l2BlockNumber uint64) (bool, error) {
	g.log.Info("validating deleted output as a result of challenge...", "l2BlockNumber", l2BlockNumber, "outputRoot", outputRoot)
	localOutputRoot, err := g.OutputRootAtBlock(ctx, l2BlockNumber)
	if err != nil {
		return false, fmt.Errorf("failed to get output root at block number %d: %w", l2BlockNumber, err)
	}
	isEqual := bytes.Equal(outputRoot[:], localOutputRoot[:])
	return isEqual, nil
}

func (g *Guardian) ConfirmTransaction(ctx context.Context, transactionId *big.Int) (*types.Transaction, error) {
	txOpts := utils.NewSimpleTxOpts(ctx, g.cfg.TxManager.From(), g.cfg.TxManager.Signer)
	return g.securityCouncilContract.ConfirmTransaction(txOpts, transactionId)
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
	return g.l2ooContract.GetL2OutputIndexAfter(utils.NewSimpleCallOpts(cCtx), l2BlockNumber)
}

func (g *Guardian) IsOutputFinalized(ctx context.Context, outputIndex *big.Int) (bool, error) {
	cCtx, cCancel := context.WithTimeout(ctx, g.cfg.NetworkTimeout)
	defer cCancel()
	return g.l2ooContract.IsFinalized(utils.NewSimpleCallOpts(cCtx), outputIndex)
}

func (g *Guardian) isTransactionConfirmed(transactionId *big.Int) (bool, error) {
	cCtx, cCancel := context.WithTimeout(g.ctx, g.cfg.NetworkTimeout)
	defer cCancel()
	return g.securityCouncilContract.IsConfirmed(utils.NewSimpleCallOpts(cCtx), transactionId)
}
