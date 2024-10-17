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
)

type (
	ProofFetcher struct {
		rpcURL     string
		httpClient *http.Client
	}

	rpcRequest struct {
		JsonRPC string `json:"jsonrpc"`
		Method  string `json:"method"`
		Params  any    `json:"params"`
		Id      int    `json:"id"`
	}

	rpcResponse struct {
		JsonRPC string          `json:"jsonrpc"`
		Result  json.RawMessage `json:"result"`
		Error   *jsonRPCError   `json:"error"`
		Id      int             `json:"id"`
	}

	jsonRPCError struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    any    `json:"data"`
	}

	ZkEVMProofFetcher interface {
		FetchProofAndPair(ctx context.Context, trace string) (*ProofAndPair, error)
	}

	zkEVMProveResponse struct {
		FinalPair []byte `json:"final_pair"`
		Proof     []byte `json:"proof"`
	}

	ProofAndPair struct {
		Proof []*big.Int
		Pair  []*big.Int
	}
)

func NewProofFetcher(rpcURL string, networkTimeout time.Duration) (*ProofFetcher, error) {
	if rpcURL == "" {
		return nil, fmt.Errorf("empty RPC URL supplied")
	}

	return &ProofFetcher{
		rpcURL: rpcURL,
		httpClient: &http.Client{
			Timeout: networkTimeout,
		},
	}, nil
}

func (f *ProofFetcher) FetchProofAndPair(ctx context.Context, trace string) (*ProofAndPair, error) {
	resultBytes, err := f.fetch(ctx, "prove", []any{trace})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch proof and pair: %w", err)
	}

	var proveResult zkEVMProveResponse
	err = json.Unmarshal(resultBytes, &proveResult)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal as zkEVMProveResponse: %w", err)
	}

	proofAndPair := &ProofAndPair{
		Proof: Decode(proveResult.Proof),
		Pair:  Decode(proveResult.FinalPair),
	}

	return proofAndPair, nil
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

func (f *ProofFetcher) fetch(ctx context.Context, method string, params []any) (json.RawMessage, error) {
	reqBody := rpcRequest{
		JsonRPC: "2.0",
		Method:  method,
		Params:  params,
		Id:      0,
	}
	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, f.rpcURL, bytes.NewReader(reqBytes))
	if err != nil {
		return nil, err
	}

	res, err := f.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	resBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response rpcResponse
	if err = json.Unmarshal(resBytes, &response); err != nil {
		return nil, err
	}

	if response.Error != nil {
		return nil, fmt.Errorf("error occurred from zkEVM prover: %w", response.Error)
	}

	return response.Result, nil
}

func (j *jsonRPCError) Error() string { return fmt.Sprintf("[%d] %s", j.Code, j.Message) }
