package validator

import (
	"bytes"
	"context"
	"fmt"
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

	txCandidatesChan chan<- txmgr.TxCandidate
}

// NewGuardian creates a new Guardian
func NewGuardian(parentCtx context.Context, cfg Config, l log.Logger, txCandidateChan chan<- txmgr.TxCandidate) (*Guardian, error) {
	securityCouncilContract, err := bindings.NewSecurityCouncil(cfg.SecurityCouncilAddr, cfg.L1Client)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(parentCtx)

	return &Guardian{
		log:                     l,
		cfg:                     cfg,
		ctx:                     ctx,
		cancel:                  cancel,
		securityCouncilContract: securityCouncilContract,
		validationRequestedChan: make(chan *bindings.SecurityCouncilValidationRequested),
		txCandidatesChan:        txCandidateChan,
	}, nil
}

func (g *Guardian) Start() error {
	g.log.Info("start Guardian")

	watchOpts := &bind.WatchOpts{Context: g.ctx, Start: nil}

	g.securityCouncilSub = event.ResubscribeErr(g.cfg.ResubscribeBackoffMax, func(ctx context.Context, err error) (event.Subscription, error) {
		if err != nil {
			g.log.Error("resubscribing after failed SecurityCouncilValidationRequested event", "err", err)
		}
		return g.securityCouncilContract.WatchValidationRequested(watchOpts, g.validationRequestedChan, nil)
	})

	g.wg.Add(1)
	go g.handleValidationRequested()

	return nil
}

func (g *Guardian) Stop() error {
	g.log.Info("stop Guardian")

	if g.securityCouncilSub != nil {
		g.securityCouncilSub.Unsubscribe()
	}
	close(g.validationRequestedChan)

	g.cancel()
	g.wg.Wait()

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
	ticker := time.NewTicker(time.Minute)
	defer func() {
		ticker.Stop()
		g.wg.Done()
	}()

	for {
	Loop:
		select {
		case <-ticker.C:
			cCtx, cCancel := context.WithTimeout(g.ctx, g.cfg.NetworkTimeout)
			callOpts := utils.NewCallOptsWithSender(cCtx, g.cfg.TxManager.From())
			isConfirmed, err := g.securityCouncilContract.IsConfirmed(callOpts, event.TransactionId)
			cCancel()
			if err != nil {
				g.log.Error("IsConfirmed failed", "err", err, "transactionId", event.TransactionId)
				break Loop
			}

			if isConfirmed {
				g.log.Info(fmt.Sprintf("Skip validate L2Output. Current tx[%+v] status(confirmed) : (%+v)", event.TransactionId, isConfirmed))
				return
			}

			isValid, err := g.validateL2Output(event.OutputRoot, event.L2BlockNumber.Uint64())
			if err != nil {
				g.log.Error("validateL2Output failed", "err", err, "l2BlockNumber", event.L2BlockNumber.Uint64())
				break Loop
			}
			if isValid {
				cCtx, cCancel := context.WithTimeout(g.ctx, g.cfg.NetworkTimeout)
				txOpts := utils.NewSimpleTxOpts(cCtx, g.cfg.TxManager.From(), g.cfg.TxManager.Signer)
				tx, err := g.securityCouncilContract.ConfirmTransaction(txOpts, event.TransactionId)
				cCancel()
				if err != nil {
					g.log.Error("tx call ConfirmTransaction failed", "err", err, "transactionId", event.TransactionId)
					break Loop
				}
				g.sendTransaction(tx)
			}
			return
		case <-g.ctx.Done():
			return
		}
	}
}

func (g *Guardian) validateL2Output(outputRoot eth.Bytes32, l2BlockNumber uint64) (bool, error) {
	localOutputRoot, err := g.outputRootAtBlock(l2BlockNumber)
	if err != nil {
		return false, fmt.Errorf("failed to get outputRootAtBlock: %w", err)
	}
	isValid := bytes.Equal(outputRoot[:], localOutputRoot[:])
	return isValid, nil
}

func (g *Guardian) outputRootAtBlock(blockNumber uint64) (eth.Bytes32, error) {
	cCtx, cCancel := context.WithTimeout(g.ctx, g.cfg.NetworkTimeout)
	defer cCancel()
	output, err := g.cfg.RollupClient.OutputAtBlock(cCtx, blockNumber, false)
	if err != nil {
		return eth.Bytes32{}, err
	}
	return output.OutputRoot, nil
}

func (g *Guardian) sendTransaction(tx *types.Transaction) {
	g.txCandidatesChan <- txmgr.TxCandidate{
		TxData:     tx.Data(),
		To:         tx.To(),
		GasLimit:   0,
		AccessList: nil,
	}
}
