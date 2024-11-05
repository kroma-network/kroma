package e2eutils

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/rpc"
	"golang.org/x/sync/errgroup"

	chal "github.com/kroma-network/kroma/kroma-validator/challenge"
)

type MockRPC struct {
	dataDir string
}

func NewMockRPC(dataDir string) *MockRPC {
	return &MockRPC{dataDir}
}

func (m *MockRPC) Close() {}

func (m *MockRPC) CallContext(ctx context.Context, result any, method string, _ ...any) error {
	switch method {
	case "prove":
		proveRes, err := m.prove(ctx)
		if err != nil {
			return err
		}

		if r, ok := result.(**chal.ZkEVMProveResponse); ok {
			*r = proveRes
		} else {
			return fmt.Errorf("invalid type for result: %T (method %s)", result, method)
		}
	default:
		return fmt.Errorf("CallContext invalid method %s", method)
	}

	return nil
}

func read(path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (m *MockRPC) prove(ctx context.Context) (*chal.ZkEVMProveResponse, error) {
	buf := make([][]byte, 2)
	files := []string{"verify_circuit_final_pair.data", "verify_circuit_proof.data"}

	g, _ := errgroup.WithContext(ctx)

	for i := 0; i < len(files); i++ {
		filePath := filepath.Join(m.dataDir, files[i])
		i := i

		g.Go(func() error {
			data, err := read(filePath)
			if err != nil {
				return err
			}

			buf[i] = data
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}

	result := &chal.ZkEVMProveResponse{
		FinalPair: buf[0],
		Proof:     buf[1],
	}

	return result, nil
}

func (m *MockRPC) BatchCallContext(_ context.Context, _ []rpc.BatchElem) error {
	return errors.New("BatchCallContext should not be called")
}

func (m *MockRPC) EthSubscribe(_ context.Context, _ any, _ ...any) (ethereum.Subscription, error) {
	return nil, errors.New("EthSubscribe should not be called")
}
