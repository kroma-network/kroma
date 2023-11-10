package node

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"

	"github.com/ethereum-optimism/optimism/op-node/rollup"
	"github.com/ethereum-optimism/optimism/op-node/version"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum-optimism/optimism/op-service/metrics"
	"github.com/ethereum-optimism/optimism/op-service/rpc"
	"github.com/kroma-network/kroma/kroma-bindings/bindings"
	"github.com/kroma-network/kroma/kroma-bindings/predeploys"
)

type l2EthClient interface {
	InfoByHash(ctx context.Context, hash common.Hash) (eth.BlockInfo, error)
	// GetProof returns a proof of the account, it may return a nil result without error if the address was not found.
	// Optionally keys of the account storage trie can be specified to include with corresponding values in the proof.
	GetProof(ctx context.Context, address common.Address, storage []common.Hash, blockTag string) (*eth.AccountResult, error)
	OutputV0AtBlock(ctx context.Context, blockHash common.Hash) (*eth.OutputV0, error)

	// [Kroma: START]
	InfoAndTxsByHash(ctx context.Context, hash common.Hash) (eth.BlockInfo, types.Transactions, error)
	// [Kroma: END]
}

type driverClient interface {
	SyncStatus(ctx context.Context) (*eth.SyncStatus, error)
	BlockRefWithStatus(ctx context.Context, num uint64) (eth.L2BlockRef, *eth.SyncStatus, error)
	ResetDerivationPipeline(context.Context) error
	StartSequencer(ctx context.Context, blockHash common.Hash) error
	StopSequencer(context.Context) (common.Hash, error)
	SequencerActive(context.Context) (bool, error)

	// [Kroma: START]
	BlockRefsWithStatus(ctx context.Context, num uint64) (eth.L2BlockRef, eth.L2BlockRef, *eth.SyncStatus, error)
	// [Kroma: END]
}

type adminAPI struct {
	*rpc.CommonAdminAPI
	dr driverClient
}

func NewAdminAPI(dr driverClient, m metrics.RPCMetricer, log log.Logger) *adminAPI {
	return &adminAPI{
		CommonAdminAPI: rpc.NewCommonAdminAPI(m, log),
		dr:             dr,
	}
}

func (n *adminAPI) ResetDerivationPipeline(ctx context.Context) error {
	recordDur := n.M.RecordRPCServerRequest("admin_resetDerivationPipeline")
	defer recordDur()
	return n.dr.ResetDerivationPipeline(ctx)
}

func (n *adminAPI) StartSequencer(ctx context.Context, blockHash common.Hash) error {
	recordDur := n.M.RecordRPCServerRequest("admin_startSequencer")
	defer recordDur()
	return n.dr.StartSequencer(ctx, blockHash)
}

func (n *adminAPI) StopSequencer(ctx context.Context) (common.Hash, error) {
	recordDur := n.M.RecordRPCServerRequest("admin_stopSequencer")
	defer recordDur()
	return n.dr.StopSequencer(ctx)
}

func (n *adminAPI) SequencerActive(ctx context.Context) (bool, error) {
	recordDur := n.M.RecordRPCServerRequest("admin_sequencerActive")
	defer recordDur()
	return n.dr.SequencerActive(ctx)
}

type nodeAPI struct {
	config *rollup.Config
	client l2EthClient
	dr     driverClient
	log    log.Logger
	m      metrics.RPCMetricer
}

func NewNodeAPI(config *rollup.Config, l2Client l2EthClient, dr driverClient, log log.Logger, m metrics.RPCMetricer) *nodeAPI {
	return &nodeAPI{
		config: config,
		client: l2Client,
		dr:     dr,
		log:    log,
		m:      m,
	}
}

func (n *nodeAPI) OutputAtBlock(ctx context.Context, number hexutil.Uint64) (*eth.OutputResponse, error) {
	recordDur := n.m.RecordRPCServerRequest("optimism_outputAtBlock")
	defer recordDur()

	output, err := n.fetchOutputAtBlock(ctx, number)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func (n *nodeAPI) SyncStatus(ctx context.Context) (*eth.SyncStatus, error) {
	recordDur := n.m.RecordRPCServerRequest("optimism_syncStatus")
	defer recordDur()
	return n.dr.SyncStatus(ctx)
}

func (n *nodeAPI) RollupConfig(_ context.Context) (*rollup.Config, error) {
	recordDur := n.m.RecordRPCServerRequest("optimism_rollupConfig")
	defer recordDur()
	return n.config, nil
}

func (n *nodeAPI) Version(ctx context.Context) (string, error) {
	recordDur := n.m.RecordRPCServerRequest("optimism_version")
	defer recordDur()
	return version.Version + "-" + version.Meta, nil
}

func (n *nodeAPI) OutputWithProofAtBlock(ctx context.Context, number hexutil.Uint64) (*eth.OutputResponse, error) {
	recordDur := n.m.RecordRPCServerRequest("kroma_outputWithProofAtBlock")
	defer recordDur()

	output, err := n.fetchOutputAtBlock(ctx, number)
	if err != nil {
		return nil, err
	}

	nextHead, nextTxs, err := n.client.InfoAndTxsByHash(ctx, output.NextBlockRef.Hash)
	if err != nil {
		return nil, fmt.Errorf("failed to get L2 block by hash %s: %w", output.NextBlockRef, err)
	}
	nextBlock := nextHead.Header()

	// TODO(seolaoh): reuse the proof fetched in `fetchOutputAtBlock` function
	accountResult, err := n.client.GetProof(ctx, predeploys.L2ToL1MessagePasserAddr, []common.Hash{}, output.BlockRef.Hash.String())
	if err != nil {
		return nil, fmt.Errorf("failed to get proof of L2ToL1MessagePasser by hash %s: %w", output.BlockRef.Hash.String(), err)
	}
	l2ToL1MessagePasserBalance := accountResult.Balance.ToInt()
	l2ToL1MessagePasserCodeHash := accountResult.CodeHash
	merkleProof := accountResult.AccountProof

	output.PublicInputProof = &eth.PublicInputProof{
		NextBlock:                   nextBlock,
		NextTransactions:            nextTxs,
		L2ToL1MessagePasserBalance:  l2ToL1MessagePasserBalance,
		L2ToL1MessagePasserCodeHash: l2ToL1MessagePasserCodeHash,
		MerkleProof:                 merkleProof,
	}

	return output, nil
}

func (n *nodeAPI) fetchOutputAtBlock(ctx context.Context, number hexutil.Uint64) (*eth.OutputResponse, error) {
	ref, nextRef, status, err := n.dr.BlockRefsWithStatus(ctx, uint64(number))
	if err != nil {
		return nil, fmt.Errorf("failed to get L2 block ref with sync status: %w", err)
	}

	head, err := n.client.InfoByHash(ctx, ref.Hash)
	if err != nil {
		return nil, fmt.Errorf("failed to get L2 block by hash %s: %w", ref, err)
	}
	if head == nil {
		return nil, ethereum.NotFound
	}

	proof, err := n.client.GetProof(ctx, predeploys.L2ToL1MessagePasserAddr, []common.Hash{}, ref.Hash.String())
	if err != nil {
		return nil, fmt.Errorf("failed to get contract proof at block %s: %w", ref, err)
	}
	if proof == nil {
		return nil, fmt.Errorf("proof %w", ethereum.NotFound)
	}
	// make sure that the proof (including storage hash) that we retrieved is correct by verifying it against the state-root
	if err := proof.Verify(head.Root()); err != nil {
		n.log.Error("invalid withdrawal root detected in block", "stateRoot", head.Root(), "blocknum", number, "msg", err)
		return nil, fmt.Errorf("invalid withdrawal root hash, state root was %s: %w", head.Root(), err)
	}

	l2OutputRootVersion := eth.OutputVersionV0 // current version is 0
	l2OutputRoot, err := rollup.ComputeL2OutputRoot(&bindings.TypesOutputRootProof{
		Version:                  l2OutputRootVersion,
		StateRoot:                head.Root(),
		MessagePasserStorageRoot: proof.StorageHash,
		BlockHash:                head.Hash(),
		NextBlockHash:            nextRef.Hash,
	})
	if err != nil {
		n.log.Error("Error computing L2 output root, nil ptr passed to hashing function")
		return nil, err
	}

	return &eth.OutputResponse{
		Version:               l2OutputRootVersion,
		OutputRoot:            l2OutputRoot,
		BlockRef:              ref,
		NextBlockRef:          nextRef,
		WithdrawalStorageRoot: proof.StorageHash,
		StateRoot:             head.Root(),
		Status:                status,
	}, nil
}
