package txmgr

import (
	"context"
	"errors"
	"sync"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/kroma-network/kroma/utils/service/txmgr/metrics"
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
	txCandidate  TxCandidate
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

	manager := BufferedTxManager{
		SimpleTxManager: *simpleTxManager,
	}
	return &manager, nil
}

func (m *BufferedTxManager) Start(ctx context.Context) error {
	m.l.Info("starting BufferedTxManager")
	m.txRequestChan = make(chan *TxRequest, m.Config.TxBufferSize)
	m.ctx, m.cancel = context.WithCancel(ctx)
	m.wg.Add(1)
	go m.listen()
	return nil
}

func (m *BufferedTxManager) Stop() error {
	m.l.Info("stopping BufferedTxManager")

	if m.cancel != nil {
		m.cancel()
	}
	m.wg.Wait()
	close(m.txRequestChan)
	return nil
}

func (m *BufferedTxManager) SubmitTransaction(ctx context.Context, txCandidate TxCandidate) *TxResponse {
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

	txResponse := txRequest.waitForResponse()

	return &TxResponse{txResponse.Receipt, txResponse.Err}
}

func (m *BufferedTxManager) listen() {
	defer m.wg.Done()
	for {
		select {
		case txRequest := <-m.txRequestChan:
			txReceipt, err := m.Send(txRequest.ctx, txRequest.txCandidate)
			if err != nil {
				m.l.Error("failed to send transaction in buffered tx manager", "err", err)
			}
			txRequest.responseChan <- &TxResponse{txReceipt, err}
		}
	}
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
			txResponse := TxResponse{Err: errors.New("context cancelled in WaitForResponse")}
			return &txResponse
		}
	}
}
