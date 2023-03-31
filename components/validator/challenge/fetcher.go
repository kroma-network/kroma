package challenge

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/wemixkanvas/kanvas/components/node/eth"
	pb "github.com/wemixkanvas/kanvas/components/validator/challenge/kanvas-grpc-proto"
)

type Fetcher struct {
	Client  pb.ProofClient
	logger  log.Logger
	conn    *grpc.ClientConn
	timeout time.Duration
}

func NewFetcher(grpcUrl string, timeout time.Duration, logger log.Logger) (*Fetcher, error) {
	if grpcUrl == "" {
		return nil, fmt.Errorf("no grpc url specified")
	}

	conn, err := grpc.Dial(grpcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to grpc server: %w", err)
	}

	return &Fetcher{
		Client:  pb.NewProofClient(conn),
		logger:  logger,
		conn:    conn,
		timeout: timeout,
	}, nil
}

type ProofAndPair struct {
	Proof []*big.Int
	Pair  []*big.Int
}

func (f *Fetcher) FetchProofAndPair(blockRef eth.L2BlockRef) (*ProofAndPair, error) {
	ctx, cancel := context.WithTimeout(context.Background(), f.timeout)
	defer cancel()

	blockNumberHex := fmt.Sprintf("0x%x", blockRef.Number)
	f.logger.Info("received block number hex", "hex", blockNumberHex)

	response, err := f.Client.Prove(ctx, &pb.ProofRequest{BlockNumberHex: blockNumberHex})
	if err != nil {
		f.logger.Warn("could not request", "err", err)
		return nil, err
	}

	result := &ProofAndPair{
		Proof: Decode(response.Proof),
		Pair:  Decode(response.FinalPair),
	}

	return result, nil
}

func (f *Fetcher) Close() error {
	f.logger.Info("Closing grpc connection")
	return f.conn.Close()
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
