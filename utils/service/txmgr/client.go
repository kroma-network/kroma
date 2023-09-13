package txmgr

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

// ETHBackend is the set of methods that the transaction manager uses to resubmit gas & determine
// when transactions are included on L1.
type ETHBackend interface {
	// BlockNumber returns the most recent block number.
	BlockNumber(ctx context.Context) (uint64, error)

	// TransactionReceipt queries the backend for a receipt associated with
	// txHash. If lookup does not fail, but the transaction is not found,
	// nil should be returned for both values.
	TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error)

	// SendTransaction submits a signed transaction to L1.
	SendTransaction(ctx context.Context, tx *types.Transaction) error

	// These functions are used to estimate what the basefee & priority fee should be set to.
	// TODO(CLI-3318): Maybe need a generic interface to support different RPC providers
	HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error)
	SuggestGasTipCap(ctx context.Context) (*big.Int, error)
	// NonceAt returns the account nonce of the given account.
	// The block number can be nil, in which case the nonce is taken from the latest known block.
	NonceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (uint64, error)
	// PendingNonceAt returns the pending nonce.
	PendingNonceAt(ctx context.Context, account common.Address) (uint64, error)
	// EstimateGas returns an estimate of the amount of gas needed to execute the given
	// transaction against the current pending block.
	EstimateGas(ctx context.Context, msg ethereum.CallMsg) (uint64, error)
	// CreateAccessList tries to create an access list for a specific transaction based on the
	// current pending state of the blockchain.
	CreateAccessList(ctx context.Context, msg ethereum.CallMsg) (*types.AccessList, uint64, string, error)
}

type CombinedClient struct {
	ec *ethclient.Client
	gc *gethclient.Client
}

func NewCombinedClient(c *rpc.Client) *CombinedClient {
	return &CombinedClient{
		ec: ethclient.NewClient(c),
		gc: gethclient.New(c),
	}
}

func (c *CombinedClient) BlockNumber(ctx context.Context) (uint64, error) {
	return c.ec.BlockNumber(ctx)
}

func (c *CombinedClient) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	return c.ec.TransactionReceipt(ctx, txHash)
}

func (c *CombinedClient) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	return c.ec.SendTransaction(ctx, tx)
}

func (c *CombinedClient) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	return c.ec.HeaderByNumber(ctx, number)
}

func (c *CombinedClient) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	return c.ec.SuggestGasTipCap(ctx)
}

func (c *CombinedClient) NonceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (uint64, error) {
	return c.ec.NonceAt(ctx, account, blockNumber)
}

func (c *CombinedClient) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	return c.ec.PendingNonceAt(ctx, account)
}

func (c *CombinedClient) EstimateGas(ctx context.Context, msg ethereum.CallMsg) (uint64, error) {
	return c.ec.EstimateGas(ctx, msg)
}

func (c *CombinedClient) CreateAccessList(ctx context.Context, msg ethereum.CallMsg) (*types.AccessList, uint64, string, error) {
	return c.gc.CreateAccessList(ctx, msg)
}
