package e2eutils

import (
	"context"
	"math/big"
	"os"
	"path/filepath"

	"github.com/ethereum/go-ethereum/log"
	"golang.org/x/sync/errgroup"

	chal "github.com/kroma-network/kroma/components/validator/challenge"
)

type Fetcher struct {
	l        log.Logger
	mockPath string
}

func NewFetcher(logger log.Logger, path string) *Fetcher {
	return &Fetcher{
		l:        logger,
		mockPath: path,
	}
}

func read(path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (f *Fetcher) FetchProofAndPair(_ context.Context, _ string) (*chal.ProofAndPair, error) {
	decoded := make([][]*big.Int, 2)
	files := []string{"verify_circuit_proof.data", "verify_circuit_final_pair.data"}

	g, _ := errgroup.WithContext(context.Background())

	for i := 0; i < len(files); i++ {
		filePath := filepath.Join(f.mockPath, files[i])
		i := i

		g.Go(func() error {
			data, err := read(filePath)
			if err != nil {
				return err
			}

			decoded[i] = chal.Decode(data)
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}

	result := &chal.ProofAndPair{
		Proof: decoded[0],
		Pair:  decoded[1],
	}

	return result, nil
}
