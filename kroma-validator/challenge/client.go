package challenge

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	rpcURL     string
	httpClient *http.Client
}

type rpcRequest struct {
	JsonRPC string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  any    `json:"params"`
	Id      int    `json:"id"`
}

type rpcResponse struct {
	JsonRPC string          `json:"jsonrpc"`
	Result  json.RawMessage `json:"result"`
	Error   *jsonRPCError   `json:"error"`
	Id      int             `json:"id"`
}

type jsonRPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func NewClient(rpcURL string, networkTimeout time.Duration) *Client {
	return &Client{
		rpcURL: rpcURL,
		httpClient: &http.Client{
			Timeout: networkTimeout,
		},
	}
}

func (c *Client) fetch(ctx context.Context, method string, params []any) (json.RawMessage, error) {
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

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.rpcURL, bytes.NewReader(reqBytes))
	if err != nil {
		return nil, err
	}

	res, err := c.httpClient.Do(req)
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
		return nil, fmt.Errorf("error occurred from RPC provider: %w", response.Error)
	}

	return response.Result, nil
}

func (j *jsonRPCError) Error() string { return fmt.Sprintf("[%d] %s", j.Code, j.Message) }
