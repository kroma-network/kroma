package txmgr

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/txpool"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"

	"github.com/kroma-network/kroma/utils/service/txmgr/metrics"
)

// Geth defaults the priceBump to 10
// Set it to 15% to be more aggressive about including transactions
const priceBump int64 = 15

// new = old * (100 + priceBump) / 100
var (
	priceBumpPercent = big.NewInt(100 + priceBump)
	oneHundred       = big.NewInt(100)
)

// ErrTxReceiptNotSucceed is the error returned when tx confirmed but the status is not success.
var ErrTxReceiptNotSucceed = errors.New("transaction confirmed but the status is not success")

// TxManager is an interface that allows callers to reliably publish txs,
// bumping the gas price if needed, and obtain the receipt of the resulting tx.
//
//go:generate mockery --name TxManager --output ./mocks
type TxManager interface {
	// Send is used to create & send a transaction. It will handle increasing
	// the gas price & ensuring that the transaction remains in the transaction pool.
	// It can be stopped by cancelling the provided context; however, the transaction
	// may be included on L1 even if the context is cancelled.
	//
	// NOTE: Send should be called by AT MOST one caller at a time.
	Send(ctx context.Context, candidate TxCandidate) (*types.Receipt, error)

	// From returns the sending address associated with the instance of the transaction manager.
	// It is static for a single instance of a TxManager.
	From() common.Address
}

// SimpleTxManager is an implementation of TxManager that performs linear fee
// bumping of a tx until it confirms.
type SimpleTxManager struct {
	Config  // directly embed the config
	name    string
	chainID *big.Int

	backend ETHBackend
	l       log.Logger
	metr    metrics.TxMetricer
}

// NewSimpleTxManager initializes a new SimpleTxManager with the passed Config.
func NewSimpleTxManager(name string, l log.Logger, m metrics.TxMetricer, cfg CLIConfig) (*SimpleTxManager, error) {
	conf, err := NewConfig(cfg, l)
	if err != nil {
		return nil, err
	}

	return &SimpleTxManager{
		chainID: conf.ChainID,
		name:    name,
		Config:  conf,
		backend: conf.Backend,
		l:       l.New("service", name),
		metr:    m,
	}, nil
}

func (m *SimpleTxManager) From() common.Address {
	return m.Config.From
}

// TxCandidate is a transaction candidate that can be submitted to ask the
// [TxManager] to construct a transaction with gas price bounds.
type TxCandidate struct {
	// TxData is the transaction data to be used in the constructed tx.
	TxData []byte
	// To is the recipient of the constructed tx. Nil means contract creation.
	To *common.Address
	// GasLimit is the gas limit to be used in the constructed tx.
	GasLimit uint64
	// AccessList is an EIP-2930 access list.
	AccessList types.AccessList
	// Value is the value that is passed to the constructed tx.
	Value *big.Int
}

// Send is used to publish a transaction with incrementally higher gas prices
// until the transaction eventually confirms. This method blocks until an
// invocation of sendTx returns (called with differing gas prices). The method
// may be canceled using the passed context.
//
// The transaction manager handles all signing. If and only if the gas limit is 0, the
// transaction manager will do a gas estimation.
//
// NOTE: Send should be called by AT MOST one caller at a time.
func (m *SimpleTxManager) Send(ctx context.Context, candidate TxCandidate) (*types.Receipt, error) {
	if m.TxSendTimeout != 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, m.TxSendTimeout)
		defer cancel()
	}
	tx, err := m.craftTx(ctx, candidate)
	if err != nil {
		return nil, fmt.Errorf("failed to create the tx: %w", err)
	}
	return m.send(ctx, tx)
}

// craftTx creates the signed transaction
// It queries L1 for the current fee market conditions as well as for the nonce.
// NOTE: This method SHOULD NOT publish the resulting transaction.
// NOTE: If the [TxCandidate.GasLimit] is non-zero, it will be used as the transaction's gas.
// NOTE: Otherwise, the [SimpleTxManager] will query the specified backend for an estimate.
func (m *SimpleTxManager) craftTx(ctx context.Context, candidate TxCandidate) (*types.Transaction, error) {
	gasTipCap, basefee, err := m.suggestGasPriceCaps(ctx)
	if err != nil {
		m.metr.RPCError()
		return nil, fmt.Errorf("failed to get gas price info: %w", err)
	}
	gasFeeCap := calcGasFeeCap(basefee, gasTipCap)

	// Fetch the sender's nonce from the latest known block (nil `blockNumber`)
	childCtx, cancel := context.WithTimeout(ctx, m.NetworkTimeout)
	defer cancel()
	nonce, err := m.backend.NonceAt(childCtx, m.From(), nil)
	if err != nil {
		m.metr.RPCError()
		return nil, fmt.Errorf("failed to get nonce: %w", err)
	}
	m.metr.RecordNonce(nonce)

	msg := ethereum.CallMsg{
		From:       m.From(),
		To:         candidate.To,
		Gas:        params.MaxGasLimit,
		Data:       candidate.TxData,
		Value:      candidate.Value,
		AccessList: candidate.AccessList,
	}

	gasLimit := candidate.GasLimit
	if gasLimit == 0 {
		gas, err := m.backend.EstimateGas(ctx, msg)
		if err != nil {
			return nil, fmt.Errorf("failed to estimate gas: %w", err)
		}
		fmt.Printf("gas: %d\n", gas)
		gasLimit = gas
	}

	if candidate.AccessList == nil {
		accessList, gasUsed, vmErr, err := m.backend.CreateAccessList(ctx, msg)
		if err != nil {
			return nil, fmt.Errorf("failed to create access list: %w", err)
		}
		if vmErr != "" {
			return nil, fmt.Errorf("failed to create access list due to vm err: %s", vmErr)
		}

		// If the estimated gas without applying an access list is lower than when it is applied,
		// rebuild the access list without the tuple for the target contract.
		if gasUsed > gasLimit {
			msg.AccessList = types.AccessList{}
			for _, tuple := range *accessList {
				if tuple.Address.String() != candidate.To.String() {
					msg.AccessList = append(msg.AccessList, tuple)
				}
			}
		} else {
			msg.AccessList = *accessList
		}
	}

	m.l.Info("creating tx", "to", msg.To, "from", m.From())

	rawTx := &types.DynamicFeeTx{
		ChainID:    m.chainID,
		Nonce:      nonce,
		To:         msg.To,
		Gas:        gasLimit,
		GasTipCap:  gasTipCap,
		GasFeeCap:  gasFeeCap,
		Value:      msg.Value,
		Data:       msg.Data,
		AccessList: msg.AccessList,
	}

	ctx, cancel = context.WithTimeout(ctx, m.NetworkTimeout)
	defer cancel()
	return m.Signer(ctx, m.From(), types.NewTx(rawTx))
}

// send submits the same transaction several times with increasing gas prices as necessary.
// It waits for the transaction to be confirmed on chain.
func (m *SimpleTxManager) send(ctx context.Context, tx *types.Transaction) (*types.Receipt, error) {
	var wg sync.WaitGroup
	defer wg.Wait()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	sendState := NewSendState(m.SafeAbortNonceTooLowCount, m.TxNotInMempoolTimeout)
	receiptChan := make(chan *types.Receipt, 1)
	sendTxAsync := func(tx *types.Transaction) {
		defer wg.Done()
		m.publishAndWaitForTx(ctx, tx, sendState, receiptChan)
	}

	// Immediately publish a transaction before starting the resubmission loop
	wg.Add(1)
	go sendTxAsync(tx)

	ticker := time.NewTicker(m.ResubmissionTimeout)
	defer ticker.Stop()

	bumpCounter := 0
	for {
		select {
		case <-ticker.C:
			// Don't resubmit a transaction if it has been mined, but we are waiting for the conf depth.
			if sendState.IsWaitingForConfirmation() {
				continue
			}
			// If we see lots of unrecoverable errors (and no pending transactions) abort sending the transaction.
			if sendState.ShouldAbortImmediately() {
				m.l.Warn("Aborting transaction submission")
				return nil, errors.New("aborted transaction sending")
			}
			// Increase the gas price & submit the new transaction
			tx = m.increaseGasPrice(ctx, tx)
			wg.Add(1)
			bumpCounter += 1
			go sendTxAsync(tx)

		case <-ctx.Done():
			return nil, ctx.Err()

		case receipt := <-receiptChan:
			m.metr.RecordGasBumpCount(bumpCounter)
			m.metr.TxConfirmed(receipt)
			// If transaction confirmed but the status is not success, return ErrTxReceiptNotSucceed
			if receipt.Status != types.ReceiptStatusSuccessful {
				return receipt, ErrTxReceiptNotSucceed
			}
			m.l.Info("Transaction receipt status successful", "hash", receipt.TxHash)
			return receipt, nil
		}
	}
}

// publishAndWaitForTx publishes the transaction to the transaction pool and then waits for it with [waitMined].
// It should be called in a new go-routine. It will send the receipt to receiptChan in a non-blocking way if a receipt is found
// for the transaction.
func (m *SimpleTxManager) publishAndWaitForTx(ctx context.Context, tx *types.Transaction, sendState *SendState, receiptChan chan *types.Receipt) {
	l := m.l.New("hash", tx.Hash(), "nonce", tx.Nonce(), "gasTipCap", tx.GasTipCap(), "gasFeeCap", tx.GasFeeCap())
	l.Info("publishing transaction")

	cCtx, cancel := context.WithTimeout(ctx, m.NetworkTimeout)
	defer cancel()
	t := time.Now()
	err := m.backend.SendTransaction(cCtx, tx)
	sendState.ProcessSendError(err)

	// Properly log & exit if there is an error
	if err != nil {
		switch {
		case errStringMatch(err, core.ErrNonceTooLow):
			l.Warn("nonce too low", "err", err)
			m.metr.TxPublished("nonce_to_low")
		case errStringMatch(err, context.Canceled):
			m.metr.RPCError()
			l.Warn("transaction send cancelled", "err", err)
			m.metr.TxPublished("context_cancelled")
		case errStringMatch(err, txpool.ErrAlreadyKnown):
			l.Warn("resubmitted already known transaction", "err", err)
			m.metr.TxPublished("tx_already_known")
		case errStringMatch(err, txpool.ErrReplaceUnderpriced):
			l.Warn("transaction replacement is underpriced", "err", err)
			m.metr.TxPublished("tx_replacement_underpriced")
		case errStringMatch(err, txpool.ErrUnderpriced):
			l.Warn("transaction is underpriced", "err", err)
			m.metr.TxPublished("tx_underpriced")
		default:
			m.metr.RPCError()
			l.Error("unable to publish transaction", "err", err)
			m.metr.TxPublished("unknown_error")
		}
		return
	}
	m.metr.TxPublished("")

	l.Info("Transaction successfully published")
	// Poll for the transaction to be ready & then send the result to receiptChan
	receipt, err := m.waitMined(ctx, tx, sendState)
	if err != nil {
		l.Warn("Transaction receipt not found", "err", err)
		return
	}
	select {
	case receiptChan <- receipt:
		m.metr.RecordTxConfirmationLatency(time.Since(t).Milliseconds())
	default:
	}
}

// waitMined waits for the transaction to be mined or for the context to be cancelled.
func (m *SimpleTxManager) waitMined(ctx context.Context, tx *types.Transaction, sendState *SendState) (*types.Receipt, error) {
	txHash := tx.Hash()
	queryTicker := time.NewTicker(m.ReceiptQueryInterval)
	defer queryTicker.Stop()
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-queryTicker.C:
			if receipt := m.queryReceipt(ctx, txHash, sendState); receipt != nil {
				return receipt, nil
			}
		}
	}
}

// queryReceipt queries for the receipt and returns the receipt if it has passed the confirmation depth
func (m *SimpleTxManager) queryReceipt(ctx context.Context, txHash common.Hash, sendState *SendState) *types.Receipt {
	ctx, cancel := context.WithTimeout(ctx, m.NetworkTimeout)
	defer cancel()
	receipt, err := m.backend.TransactionReceipt(ctx, txHash)
	if errors.Is(err, ethereum.NotFound) {
		sendState.TxNotMined(txHash)
		m.l.Trace("Transaction not yet mined", "hash", txHash)
		return nil
	} else if err != nil {
		m.metr.RPCError()
		m.l.Info("Receipt retrieval failed", "hash", txHash, "err", err)
		return nil
	} else if receipt == nil {
		m.metr.RPCError()
		m.l.Warn("Receipt and error are both nil", "hash", txHash)
		return nil
	}

	// Receipt is confirmed to be valid from this point on
	sendState.TxMined(txHash)

	txHeight := receipt.BlockNumber.Uint64()
	tipHeight, err := m.backend.BlockNumber(ctx)
	if err != nil {
		m.l.Error("Unable to fetch block number", "err", err)
		return nil
	}

	m.l.Debug("Transaction mined, checking confirmations", "hash", txHash, "txHeight", txHeight,
		"tipHeight", tipHeight, "numConfirmations", m.NumConfirmations)

	// The transaction is considered confirmed when
	// txHeight+numConfirmations-1 <= tipHeight. Note that the -1 is
	// needed to account for the fact that confirmations have an
	// inherent off-by-one, i.e. when using 1 confirmation the
	// transaction should be confirmed when txHeight is equal to
	// tipHeight. The equation is rewritten in this form to avoid
	// underflows.
	if txHeight+m.NumConfirmations <= tipHeight+1 {
		m.l.Info("Transaction confirmed", "hash", txHash)
		return receipt
	}

	// Safe to subtract since we know the LHS above is greater.
	confsRemaining := (txHeight + m.NumConfirmations) - (tipHeight + 1)
	m.l.Debug("Transaction not yet confirmed", "hash", txHash, "confsRemaining", confsRemaining)
	return nil
}

// increaseGasPrice takes the previous transaction & potentially clones then signs it with a higher tip.
// If the tip + basefee suggested by the network are not greater than the previous values, the same transaction
// will be returned. If they are greater, this function will ensure that they are at least greater by 15% than
// the previous transaction's value to ensure that the price bump is large enough.
//
// We do not re-estimate the amount of gas used because for some stateful transactions (like output proposals) the
// act of including the transaction renders the repeat of the transaction invalid.
//
// If it encounters an error with creating the new transaction, it will return the old transaction.
func (m *SimpleTxManager) increaseGasPrice(ctx context.Context, tx *types.Transaction) *types.Transaction {
	tip, basefee, err := m.suggestGasPriceCaps(ctx)
	if err != nil {
		m.l.Warn("failed to get suggested gas tip and basefee", "err", err)
		return tx
	}
	gasTipCap, gasFeeCap := updateFees(tx.GasTipCap(), tx.GasFeeCap(), tip, basefee, m.l)

	if tx.GasTipCapIntCmp(gasTipCap) == 0 && tx.GasFeeCapIntCmp(gasFeeCap) == 0 {
		return tx
	}

	rawTx := &types.DynamicFeeTx{
		ChainID:    tx.ChainId(),
		Nonce:      tx.Nonce(),
		GasTipCap:  gasTipCap,
		GasFeeCap:  gasFeeCap,
		Gas:        tx.Gas(),
		To:         tx.To(),
		Value:      tx.Value(),
		Data:       tx.Data(),
		AccessList: tx.AccessList(),
	}
	ctx, cancel := context.WithTimeout(ctx, m.NetworkTimeout)
	defer cancel()
	newTx, err := m.Signer(ctx, m.From(), types.NewTx(rawTx))
	if err != nil {
		m.l.Warn("failed to sign new transaction", "err", err)
		return tx
	}
	return newTx
}

// suggestGasPriceCaps suggests what the new tip & new basefee should be based on the current L1 conditions
func (m *SimpleTxManager) suggestGasPriceCaps(ctx context.Context) (*big.Int, *big.Int, error) {
	cCtx, cancel := context.WithTimeout(ctx, m.NetworkTimeout)
	defer cancel()
	tip, err := m.backend.SuggestGasTipCap(cCtx)
	if err != nil {
		m.metr.RPCError()
		return nil, nil, fmt.Errorf("failed to fetch the suggested gas tip cap: %w", err)
	} else if tip == nil {
		return nil, nil, errors.New("the suggested tip was nil")
	}
	cCtx, cancel = context.WithTimeout(ctx, m.NetworkTimeout)
	defer cancel()
	head, err := m.backend.HeaderByNumber(cCtx, nil)
	if err != nil {
		m.metr.RPCError()
		return nil, nil, fmt.Errorf("failed to fetch the suggested basefee: %w", err)
	} else if head.BaseFee == nil {
		return nil, nil, errors.New("txmgr does not support pre-london blocks that do not have a basefee")
	}
	return tip, head.BaseFee, nil
}

// calcThresholdValue returns x * priceBumpPercent / 100
func calcThresholdValue(x *big.Int) *big.Int {
	threshold := new(big.Int).Mul(priceBumpPercent, x)
	threshold = threshold.Div(threshold, oneHundred)
	return threshold
}

// updateFees takes the old tip/basefee & the new tip/basefee and then suggests
// a gasTipCap and gasFeeCap that satisfies geth's required fee bumps
// Geth: FC and Tip must be bumped if any increase
func updateFees(oldTip, oldFeeCap, newTip, newBaseFee *big.Int, lgr log.Logger) (*big.Int, *big.Int) {
	newFeeCap := calcGasFeeCap(newBaseFee, newTip)
	lgr = lgr.New("old_tip", oldTip, "old_feecap", oldFeeCap, "new_tip", newTip, "new_feecap", newFeeCap)
	// If the new prices are less than the old price, reuse the old prices
	if oldTip.Cmp(newTip) >= 0 && oldFeeCap.Cmp(newFeeCap) >= 0 {
		lgr.Debug("Reusing old tip and feecap")
		return oldTip, oldFeeCap
	}
	// Determine if we need to increase the suggested values
	thresholdTip := calcThresholdValue(oldTip)
	thresholdFeeCap := calcThresholdValue(oldFeeCap)
	if newTip.Cmp(thresholdTip) >= 0 && newFeeCap.Cmp(thresholdFeeCap) >= 0 {
		lgr.Debug("Using new tip and feecap")
		return newTip, newFeeCap
	} else if newTip.Cmp(thresholdTip) >= 0 && newFeeCap.Cmp(thresholdFeeCap) < 0 {
		// Tip has gone up, but basefee is flat or down.
		// TODO(CLI-3714): Do we need to recalculate the FC here?
		lgr.Debug("Using new tip and threshold feecap")
		return newTip, thresholdFeeCap
	} else if newTip.Cmp(thresholdTip) < 0 && newFeeCap.Cmp(thresholdFeeCap) >= 0 {
		// Basefee has gone up, but the tip hasn't. Recalculate the feecap because if the tip went up a lot
		// not enough of the feecap may be dedicated to paying the basefee.
		lgr.Debug("Using threshold tip and recalculated feecap")
		return thresholdTip, calcGasFeeCap(newBaseFee, thresholdTip)

	} else {
		// TODO(CLI-3713): Should we skip the bump in this case?
		lgr.Debug("Using threshold tip and threshold feecap")
		return thresholdTip, thresholdFeeCap
	}
}

// calcGasFeeCap deterministically computes the recommended gas fee cap given
// the base fee and gasTipCap. The resulting gasFeeCap is equal to:
//
//	gasTipCap + 2*baseFee.
func calcGasFeeCap(baseFee, gasTipCap *big.Int) *big.Int {
	return new(big.Int).Add(
		gasTipCap,
		new(big.Int).Mul(baseFee, big.NewInt(2)),
	)
}

// errStringMatch returns true if err.Error() is a substring in target.Error() or if both are nil.
// It can accept nil errors without issue.
func errStringMatch(err, target error) bool {
	if err == nil && target == nil {
		return true
	} else if err == nil || target == nil {
		return false
	}
	return strings.Contains(err.Error(), target.Error())
}
