package challenge

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/log"
)

type (
	ProofType int32

	JsonRpcError struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    any    `json:"data"`
	}

	response struct {
		Jsonrpc string         `json:"jsonrpc"`
		Result  *ProveResponse `json:"result"`
		Error   *JsonRpcError  `json:"error"`
		Id      string         `json:"id"`
	}

	ProveResponse struct {
		FinalPair []byte `json:"final_pair,omitempty"`
		Proof     []byte `json:"proof,omitempty"`
	}

	ProverClient interface {
		Prove(ctx context.Context, traceString string, proofType ProofType) (*ProveResponse, error)
	}

	Fetcher struct {
		Client  JsonRPCProverClient
		logger  log.Logger
		timeout time.Duration
	}
)

func NewFetcher(rpcURL string, timeout time.Duration, logger log.Logger) (*Fetcher, error) {
	if rpcURL == "" {
		return nil, fmt.Errorf("no RPC URL specified")
	}

	return &Fetcher{
		Client:  JsonRPCProverClient{rpcURL},
		logger:  logger,
		timeout: timeout,
	}, nil
}

type ProofAndPair struct {
	Proof []*big.Int
	Pair  []*big.Int
}

func (f *Fetcher) FetchProofAndPair(ctx context.Context, trace string) (*ProofAndPair, error) {
	cCtx, cCancel := context.WithTimeout(ctx, f.timeout)
	defer cCancel()

	// NOTE(0xHansLee): only ProofType_AGG(4) is used for proof.
	// https://github.com/kroma-network/kroma-prover/blob/dev/prover-server/src/spec.rs#L10-L16
	resp, err := f.Client.Prove(cCtx, trace, 4)
	if err != nil {
		f.logger.Error("could not request fault proof", "err", err)
		return nil, err
	}

	result := &ProofAndPair{
		Proof: Decode(resp.Proof),
		Pair:  Decode(resp.FinalPair),
	}

	return result, nil
}

func Decode(data []byte) []*big.Int {
	result := make([]*big.Int, len(data)/32)

	for i := 0; i < len(data)/32; i++ {
		// The best is data is given in Big Endian.
		for j := 0; j < 16; j++ {
			data[i*32+j], data[i*32+31-j] = data[i*32+31-j], data[i*32+j]
		}
		result[i] = new(big.Int).SetBytes(data[i*32 : (i+1)*32])
	}

	return result
}

var _ ProverClient = (*JsonRPCProverClient)(nil)

type JsonRPCProverClient struct {
	address string
}

func (c JsonRPCProverClient) Prove(ctx context.Context, traceString string, proofType ProofType) (*ProveResponse, error) {
	reqBody := struct {
		Jsonrpc string `json:"jsonrpc"`
		Method  string `json:"method"`
		Params  any    `json:"params"`
		Id      string `json:"id"`
	}{
		Jsonrpc: "2.0",
		Method:  "prove",
		Params:  []any{traceString, proofType},
		Id:      "0",
	}

	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to json.Marshal %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.address, bytes.NewReader(reqBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to create new reqBody for proof: %w", err)
	}

	cli := http.Client{}
	res, err := cli.Do(req)
	if err != nil {
		return nil, err
	}

	respBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var resp response
	if err := json.Unmarshal(respBytes, &resp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if resp.Error != nil {
		return nil, fmt.Errorf("error occurs from zk prover: %s", resp.Error)
	}

	return resp.Result, nil
}

func (j *JsonRpcError) Error() string { return fmt.Sprintf("[%d] %s", j.Code, j.Message) }
