package txmgr

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"

	"github.com/ethereum-optimism/optimism/op-service/txmgr/metrics"
)

type BufferedTxManager struct {
	SimpleTxManager // directly embed
	wg              sync.WaitGroup
	txRequestChan   chan *TxRequest
	ctx             context.Context
	cancel          context.CancelFunc
}

type TxRequest struct {
	ctx          context.Context
	txCandidate  *TxCandidate
	responseChan chan *TxResponse
}

type TxResponse struct {
	Receipt *types.Receipt
	Err     error
}

func NewBufferedTxManager(name string, l log.Logger, m metrics.TxMetricer, cfg CLIConfig) (*BufferedTxManager, error) {
	simpleTxManager, err := NewSimpleTxManager(name, l, m, cfg)
	if err != nil {
		return nil, err
	}

	return &BufferedTxManager{
		SimpleTxManager: *simpleTxManager,
	}, nil
}

func (m *BufferedTxManager) Start(ctx context.Context) error {
	m.txRequestChan = make(chan *TxRequest, m.Config.TxBufferSize)
	m.ctx, m.cancel = context.WithCancel(ctx)
	m.wg.Add(1)
	go m.listen(m.ctx)
	return nil
}

func (m *BufferedTxManager) Stop() error {
	m.cancel()
	m.wg.Wait()
	close(m.txRequestChan)
	return nil
}

func (m *BufferedTxManager) listen(ctx context.Context) {
	defer m.wg.Done()
	for {
		select {
		case txRequest := <-m.txRequestChan:
			txReceipt, err := m.Send(txRequest.ctx, *txRequest.txCandidate)
			if err != nil {
				m.l.Error("failed to send transaction in buffered tx manager", "err", err)
			}
			txRequest.responseChan <- &TxResponse{txReceipt, err}
		case <-ctx.Done():
			return
		}
	}
}

func (m *BufferedTxManager) submitTransaction(ctx context.Context, txCandidate *TxCandidate) *TxResponse {
	responseChan := make(chan *TxResponse)
	defer close(responseChan)

	txRequest := &TxRequest{
		ctx:          ctx,
		txCandidate:  txCandidate,
		responseChan: responseChan,
	}
	if !m.tryEnqueue(txRequest) {
		return &TxResponse{
			nil, errors.New("submit transaction failed in tryEnqueue"),
		}
	}
	return txRequest.waitForResponse()
}

func (m *BufferedTxManager) SendTxCandidate(ctx context.Context, txCandidate *TxCandidate) *TxResponse {
	return m.submitTransaction(ctx, txCandidate)
}

func (m *BufferedTxManager) SendTransaction(ctx context.Context, tx *types.Transaction) *TxResponse {
	return m.SendTxCandidate(ctx, &TxCandidate{
		TxData:   tx.Data(),
		To:       tx.To(),
		GasLimit: 0,
	})
}

func (m *BufferedTxManager) tryEnqueue(txRequest *TxRequest) bool {
	select {
	case m.txRequestChan <- txRequest:
		return true
	default:
		return false
	}
}

func (r *TxRequest) waitForResponse() *TxResponse {
	for {
		select {
		case response := <-r.responseChan:
			return response
		case <-r.ctx.Done():
			return &TxResponse{Err: fmt.Errorf("context cancelled in WaitForResponse: %w", r.ctx.Err())}
		}
	}
}
