package e2e

import (
	"context"
	"flag"
	"fmt"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/eth/ethconfig"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/stretchr/testify/require"

	"github.com/kroma-network/kroma/bindings/bindings"
	"github.com/kroma-network/kroma/bindings/predeploys"
	"github.com/kroma-network/kroma/components/node/client"
	"github.com/kroma-network/kroma/components/node/eth"
	rollupNode "github.com/kroma-network/kroma/components/node/node"
	"github.com/kroma-network/kroma/components/node/p2p"
	"github.com/kroma-network/kroma/components/node/rollup"
	"github.com/kroma-network/kroma/components/node/rollup/derive"
	"github.com/kroma-network/kroma/components/node/rollup/driver"
	"github.com/kroma-network/kroma/components/node/sources"
	"github.com/kroma-network/kroma/components/node/testlog"
	"github.com/kroma-network/kroma/components/node/withdrawals"
)

var enableParallelTesting bool = true

// Init testing to enable test flags
var _ = func() bool {
	testing.Init()
	return true
}()

var verboseGethNodes bool

func init() {
	flag.BoolVar(&verboseGethNodes, "gethlogs", true, "Enable logs on geth nodes")
	flag.Parse()
	if os.Getenv("E2E_DISABLE_PARALLEL") == "true" {
		enableParallelTesting = false
	}
}

func parallel(t *testing.T) {
	t.Helper()
	if enableParallelTesting {
		t.Parallel()
	}
}

func TestL2OutputSubmitter(t *testing.T) {
	parallel(t)
	if !verboseGethNodes {
		log.Root().SetHandler(log.DiscardHandler())
	}

	cfg := DefaultSystemConfig(t)
	cfg.NonFinalizedOutputs = true // speed up the time till we see checkpoint outputs

	sys, err := cfg.Start()
	require.Nil(t, err, "Error starting up system")
	defer sys.Close()

	l1Client := sys.Clients["l1"]

	rollupRPCClient, err := rpc.DialContext(context.Background(), sys.RollupNodes["proposer"].HTTPEndpoint())
	require.Nil(t, err)
	rollupClient := sources.NewRollupClient(client.NewBaseRPCClient(rollupRPCClient))

	//  OutputOracle is already deployed
	l2OutputOracle, err := bindings.NewL2OutputOracleCaller(predeploys.DevL2OutputOracleAddr, l1Client)
	require.Nil(t, err)

	initialOutputBlockNumber, err := l2OutputOracle.LatestBlockNumber(&bind.CallOpts{})
	require.Nil(t, err)

	// Wait until the second output submission from L2. The output submitter submits outputs from the
	// unsafe portion of the chain which gets reorged on startup. The proposer has an out of date view
	// when it creates it's first block and uses and old L1 Origin. It then does not submit a batch
	// for that block and subsequently reorgs to match what the syncer derives when running the
	// reconciliation process.
	l2Syncer := sys.Clients["syncer"]
	_, err = waitForL2Block(big.NewInt(6), l2Syncer, 10*time.Duration(cfg.DeployConfig.L2BlockTime)*time.Second)
	require.Nil(t, err)

	// Wait for batch submitter to update L2 output oracle.
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for {
		l2ooBlockNumber, err := l2OutputOracle.LatestBlockNumber(&bind.CallOpts{})
		require.Nil(t, err)

		// Wait for the L2 output oracle to have been changed from the initial
		// timestamp set in the contract constructor.
		if l2ooBlockNumber.Cmp(initialOutputBlockNumber) > 0 {
			// Retrieve the l2 output committed at this updated timestamp.
			committedL2Output, err := l2OutputOracle.GetL2OutputAfter(&bind.CallOpts{}, l2ooBlockNumber)
			require.NotEqual(t, [32]byte{}, committedL2Output.OutputRoot, "Empty L2 Output")
			require.Nil(t, err)

			// Fetch the corresponding L2 block and assert the committed L2
			// output matches the block's state root.
			//
			// NOTE: This assertion will change once the L2 output format is
			// finalized.
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			l2Output, err := rollupClient.OutputAtBlock(ctx, l2ooBlockNumber.Uint64(), false)
			require.Nil(t, err)
			require.Equal(t, l2Output.OutputRoot[:], committedL2Output.OutputRoot[:])
			break
		}

		select {
		case <-ctx.Done():
			t.Fatalf("State root oracle not updated")
		case <-ticker.C:
		}
	}
}

// TestSystemE2E sets up a L1 Geth node, a rollup node, and a L2 geth node and then confirms that L1 deposits are reflected on L2.
// All nodes are run in process (but are the full nodes, not mocked or stubbed).
func TestSystemE2E(t *testing.T) {
	parallel(t)
	if !verboseGethNodes {
		log.Root().SetHandler(log.DiscardHandler())
	}

	cfg := DefaultSystemConfig(t)

	sys, err := cfg.Start()
	require.Nil(t, err, "Error starting up system")
	defer sys.Close()

	log := testlog.Logger(t, log.LvlInfo)
	log.Info("genesis", "l2", sys.RollupConfig.Genesis.L2, "l1", sys.RollupConfig.Genesis.L1, "l2_time", sys.RollupConfig.Genesis.L2Time)

	l1Client := sys.Clients["l1"]
	l2Prop := sys.Clients["proposer"]
	l2Sync := sys.Clients["syncer"]

	// Transactor Account
	ethPrivKey := sys.cfg.Secrets.Alice

	// Send Transaction & wait for success
	fromAddr := sys.cfg.Secrets.Addresses().Alice

	// Find deposit contract
	depositContract, err := bindings.NewKromaPortal(predeploys.DevKromaPortalAddr, l1Client)
	require.Nil(t, err)

	// Create signer
	opts, err := bind.NewKeyedTransactorWithChainID(ethPrivKey, cfg.L1ChainIDBig())
	require.Nil(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	startBalance, err := l2Sync.BalanceAt(ctx, fromAddr, nil)
	require.Nil(t, err)

	// Finally send TX
	mintAmount := big.NewInt(1_000_000_000_000)
	opts.Value = mintAmount
	tx, err := depositContract.DepositTransaction(opts, fromAddr, common.Big0, 1_000_000, false, nil)
	require.Nil(t, err, "with deposit tx")

	receipt, err := waitForTransaction(tx.Hash(), l1Client, 3*time.Duration(cfg.DeployConfig.L1BlockTime)*time.Second)
	require.Nil(t, err, "Waiting for deposit tx on L1")

	reconstructedDep, err := derive.UnmarshalDepositLogEvent(receipt.Logs[0])
	require.NoError(t, err, "Could not reconstruct L2 Deposit")
	tx = types.NewTx(reconstructedDep)
	receipt, err = waitForL2Transaction(tx.Hash(), l2Sync, 6*time.Duration(cfg.DeployConfig.L1BlockTime)*time.Second)
	require.NoError(t, err)
	require.Equal(t, receipt.Status, types.ReceiptStatusSuccessful)

	// Confirm balance
	ctx, cancel = context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	endBalance, err := l2Sync.BalanceAt(ctx, fromAddr, nil)
	require.Nil(t, err)

	diff := new(big.Int)
	diff = diff.Sub(endBalance, startBalance)
	require.Equal(t, mintAmount, diff, "Did not get expected balance change")

	// Submit TX to L2 proposer node
	toAddr := common.Address{0xff, 0xff}
	tx = types.MustSignNewTx(ethPrivKey, types.LatestSignerForChainID(cfg.L2ChainIDBig()), &types.DynamicFeeTx{
		ChainID:   cfg.L2ChainIDBig(),
		Nonce:     1, // Already have deposit
		To:        &toAddr,
		Value:     big.NewInt(1_000_000_000),
		GasTipCap: big.NewInt(10),
		GasFeeCap: big.NewInt(200),
		Gas:       21000,
	})
	err = l2Prop.SendTransaction(context.Background(), tx)
	require.Nil(t, err, "Sending L2 tx to proposer")

	_, err = waitForL2Transaction(tx.Hash(), l2Prop, 3*time.Duration(cfg.DeployConfig.L1BlockTime)*time.Second)
	require.Nil(t, err, "Waiting for L2 tx on proposer")

	receipt, err = waitForL2Transaction(tx.Hash(), l2Sync, 10*time.Duration(cfg.DeployConfig.L1BlockTime)*time.Second)
	require.Nil(t, err, "Waiting for L2 tx on syncer")
	require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status, "TX should have succeeded")

	// Verify blocks match after batch submission on syncers and proposers
	syncBlock, err := l2Sync.BlockByNumber(context.Background(), receipt.BlockNumber)
	require.Nil(t, err)
	propBlock, err := l2Prop.BlockByNumber(context.Background(), receipt.BlockNumber)
	require.Nil(t, err)
	require.Equal(t, syncBlock.NumberU64(), propBlock.NumberU64(), "Syncer and proposer blocks not the same after including a batch tx")
	require.Equal(t, syncBlock.ParentHash(), propBlock.ParentHash(), "Syncer and proposer blocks parent hashes not the same after including a batch tx")
	require.Equal(t, syncBlock.Hash(), propBlock.Hash(), "Syncer and proposer blocks not the same after including a batch tx")

	rollupRPCClient, err := rpc.DialContext(context.Background(), sys.RollupNodes["proposer"].HTTPEndpoint())
	require.Nil(t, err)
	rollupClient := sources.NewRollupClient(client.NewBaseRPCClient(rollupRPCClient))
	// basic check that sync status works
	propStatus, err := rollupClient.SyncStatus(context.Background())
	require.Nil(t, err)
	require.LessOrEqual(t, propBlock.NumberU64(), propStatus.UnsafeL2.Number)
	// basic check that version endpoint works
	propVersion, err := rollupClient.Version(context.Background())
	require.Nil(t, err)
	require.NotEqual(t, "", propVersion)
}

// TestConfirmationDepth runs the rollup with both proposer and syncer not immediately processing the tip of the chain.
func TestConfirmationDepth(t *testing.T) {
	parallel(t)
	if !verboseGethNodes {
		log.Root().SetHandler(log.DiscardHandler())
	}

	cfg := DefaultSystemConfig(t)
	cfg.DeployConfig.ProposerWindowSize = 4
	cfg.DeployConfig.MaxProposerDrift = 10 * cfg.DeployConfig.L1BlockTime
	propConfDepth := uint64(2)
	syncConfDepth := uint64(5)
	cfg.Nodes["proposer"].Driver.ProposerConfDepth = propConfDepth
	cfg.Nodes["proposer"].Driver.SyncerConfDepth = 0
	cfg.Nodes["syncer"].Driver.SyncerConfDepth = syncConfDepth

	sys, err := cfg.Start()
	require.Nil(t, err, "Error starting up system")
	defer sys.Close()

	log := testlog.Logger(t, log.LvlInfo)
	log.Info("genesis", "l2", sys.RollupConfig.Genesis.L2, "l1", sys.RollupConfig.Genesis.L1, "l2_time", sys.RollupConfig.Genesis.L2Time)

	l1Client := sys.Clients["l1"]
	l2Prop := sys.Clients["proposer"]
	l2Sync := sys.Clients["syncer"]

	// Wait enough time for the proposer to submit a block with distance from L1 head, submit it,
	// and for the slower syncer to read a full proposer window and cover confirmation depth for reading and some margin
	<-time.After(time.Duration((cfg.DeployConfig.ProposerWindowSize+syncConfDepth+3)*cfg.DeployConfig.L1BlockTime) * time.Second)

	// within a second, get both L1 and L2 syncer and proposer block heads
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	l1Head, err := l1Client.BlockByNumber(ctx, nil)
	require.NoError(t, err)
	l2PropHead, err := l2Prop.BlockByNumber(ctx, nil)
	require.NoError(t, err)
	l2SyncHead, err := l2Sync.BlockByNumber(ctx, nil)
	require.NoError(t, err)

	propInfo, err := derive.L1InfoDepositTxData(l2PropHead.Transactions()[0].Data())
	require.NoError(t, err)
	require.LessOrEqual(t, propInfo.Number+propConfDepth, l1Head.NumberU64(), "the proposer L2 head block should have an origin older than the L1 head block by at least the proposer conf depth")

	syncInfo, err := derive.L1InfoDepositTxData(l2SyncHead.Transactions()[0].Data())
	require.NoError(t, err)
	require.LessOrEqual(t, syncInfo.Number+syncConfDepth, l1Head.NumberU64(), "the syncer L2 head block should have an origin older than the L1 head block by at least the syncer conf depth")
}

// TestPendingGasLimit tests the configuration of the gas limit of the pending block,
// and if it does not conflict with the regular gas limit on the syncer or proposer.
func TestPendingGasLimit(t *testing.T) {
	parallel(t)
	if !verboseGethNodes {
		log.Root().SetHandler(log.DiscardHandler())
	}

	cfg := DefaultSystemConfig(t)

	// configure the L2 gas limit to be high, and the pending gas limits to be lower for resource saving.
	cfg.DeployConfig.L2GenesisBlockGasLimit = 20_000_000
	cfg.GethOptions["proposer"] = []GethOption{
		func(ethCfg *ethconfig.Config, nodeCfg *node.Config) error {
			ethCfg.Miner.GasCeil = 10_000_000
			return nil
		},
	}
	cfg.GethOptions["syncer"] = []GethOption{
		func(ethCfg *ethconfig.Config, nodeCfg *node.Config) error {
			ethCfg.Miner.GasCeil = 9_000_000
			return nil
		},
	}

	sys, err := cfg.Start()
	require.Nil(t, err, "Error starting up system")
	defer sys.Close()

	log := testlog.Logger(t, log.LvlInfo)
	log.Info("genesis", "l2", sys.RollupConfig.Genesis.L2, "l1", sys.RollupConfig.Genesis.L1, "l2_time", sys.RollupConfig.Genesis.L2Time)

	l2Sync := sys.Clients["syncer"]
	l2Prop := sys.Clients["proposer"]

	checkGasLimit := func(client *ethclient.Client, number *big.Int, expected uint64) *types.Header {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		header, err := client.HeaderByNumber(ctx, number)
		cancel()
		require.NoError(t, err)
		require.Equal(t, expected, header.GasLimit)
		return header
	}

	// check if the gaslimits are matching the expected values,
	// and that the syncer/proposer can use their locally configured gas limit for the pending block.
	for {
		checkGasLimit(l2Prop, big.NewInt(-1), 10_000_000)
		checkGasLimit(l2Sync, big.NewInt(-1), 9_000_000)
		checkGasLimit(l2Prop, nil, 20_000_000)
		latestSyncHeader := checkGasLimit(l2Sync, nil, 20_000_000)

		// Stop once the syncer passes genesis:
		// this implies we checked a new block from the proposer, on both proposer and syncer nodes.
		if latestSyncHeader.Number.Uint64() > 0 {
			break
		}
		time.Sleep(500 * time.Millisecond)
	}
}

// TestFinalize tests if L2 finalizes after sufficient time after L1 finalizes
func TestFinalize(t *testing.T) {
	parallel(t)
	if !verboseGethNodes {
		log.Root().SetHandler(log.DiscardHandler())
	}

	cfg := DefaultSystemConfig(t)

	sys, err := cfg.Start()
	require.Nil(t, err, "Error starting up system")
	defer sys.Close()

	l2Prop := sys.Clients["proposer"]

	// as configured in the extra geth lifecycle in testing setup
	const finalizedDistance = 8
	// Wait enough time for L1 to finalize and L2 to confirm its data in finalized L1 blocks
	timeout := time.Duration((finalizedDistance+6)*cfg.DeployConfig.L1BlockTime) * time.Second * timeoutMultiplier
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// fetch the finalizes head of geth
	for {
		select {
		case <-ctx.Done():
			require.Fail(t, "timeout")
			return
		default:
			<-time.After(time.Second)
		}

		// poll until the finalized block number is greater than 0
		l2Finalized, err := waitForL2Block(big.NewInt(int64(rpc.FinalizedBlockNumber)), l2Prop, time.Second)
		require.NoError(t, err)
		if l2Finalized.NumberU64() > 0 {
			break
		}
	}
}

func TestMintOnRevertedDeposit(t *testing.T) {
	parallel(t)
	if !verboseGethNodes {
		log.Root().SetHandler(log.DiscardHandler())
	}
	cfg := DefaultSystemConfig(t)

	sys, err := cfg.Start()
	require.Nil(t, err, "Error starting up system")
	defer sys.Close()

	l1Client := sys.Clients["l1"]
	l2Syncer := sys.Clients["syncer"]

	// Find deposit contract
	depositContract, err := bindings.NewKromaPortal(predeploys.DevKromaPortalAddr, l1Client)
	require.Nil(t, err)
	l1Node := sys.Nodes["l1"]

	// create signer
	ks := l1Node.AccountManager().Backends(keystore.KeyStoreType)[0].(*keystore.KeyStore)
	opts, err := bind.NewKeyStoreTransactorWithChainID(ks, ks.Accounts()[0], cfg.L1ChainIDBig())
	require.Nil(t, err)
	fromAddr := opts.From

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	startBalance, err := l2Syncer.BalanceAt(ctx, fromAddr, nil)
	cancel()
	require.Nil(t, err)

	ctx, cancel = context.WithTimeout(context.Background(), 1*time.Second)
	startNonce, err := l2Syncer.NonceAt(ctx, fromAddr, nil)
	require.NoError(t, err)
	cancel()

	toAddr := common.Address{0xff, 0xff}
	mintAmount := big.NewInt(9_000_000)
	opts.Value = mintAmount
	value := new(big.Int).Mul(common.Big2, startBalance) // trigger a revert by transferring more than we have available
	tx, err := depositContract.DepositTransaction(opts, toAddr, value, 1_000_000, false, nil)
	require.Nil(t, err, "with deposit tx")

	receipt, err := waitForTransaction(tx.Hash(), l1Client, 3*time.Duration(cfg.DeployConfig.L1BlockTime)*time.Second)
	require.Nil(t, err, "Waiting for deposit tx on L1")

	reconstructedDep, err := derive.UnmarshalDepositLogEvent(receipt.Logs[0])
	require.NoError(t, err, "Could not reconstruct L2 Deposit")
	tx = types.NewTx(reconstructedDep)
	receipt, err = waitForL2Transaction(tx.Hash(), l2Syncer, 10*time.Duration(cfg.DeployConfig.L1BlockTime)*time.Second)
	require.NoError(t, err)
	require.Equal(t, receipt.Status, types.ReceiptStatusFailed)

	// Confirm balance
	ctx, cancel = context.WithTimeout(context.Background(), 1*time.Second)
	endBalance, err := l2Syncer.BalanceAt(ctx, fromAddr, nil)
	cancel()
	require.Nil(t, err)
	ctx, cancel = context.WithTimeout(context.Background(), 1*time.Second)
	toAddrBalance, err := l2Syncer.BalanceAt(ctx, toAddr, nil)
	require.NoError(t, err)
	cancel()

	diff := new(big.Int)
	diff = diff.Sub(endBalance, startBalance)
	require.Equal(t, mintAmount, diff, "Did not get expected balance change")
	require.Equal(t, common.Big0.Int64(), toAddrBalance.Int64(), "The recipient account balance should be zero")

	ctx, cancel = context.WithTimeout(context.Background(), 1*time.Second)
	endNonce, err := l2Syncer.NonceAt(ctx, fromAddr, nil)
	require.NoError(t, err)
	cancel()
	require.Equal(t, startNonce+1, endNonce, "Nonce of deposit sender should increment on L2, even if the deposit fails")
}

func TestMissingBatchE2E(t *testing.T) {
	parallel(t)
	if !verboseGethNodes {
		log.Root().SetHandler(log.DiscardHandler())
	}
	// Note this test zeroes the balance of the batch-submitter to make the batches unable to go into L1.
	// The test logs may look scary, but this is expected:
	// 'batcher unable to publish transaction    role=batcher   err="insufficient funds for gas * price + value"'

	cfg := DefaultSystemConfig(t)
	// small proposer window size so the test does not take as long
	cfg.DeployConfig.ProposerWindowSize = 4

	// Specifically set batch submitter balance to stop batches from being included
	cfg.Premine[cfg.Secrets.Addresses().Batcher] = big.NewInt(0)

	sys, err := cfg.Start()
	require.Nil(t, err, "Error starting up system")
	defer sys.Close()

	l2Prop := sys.Clients["proposer"]
	l2Sync := sys.Clients["syncer"]

	// Transactor Account
	ethPrivKey := cfg.Secrets.Alice

	// Submit TX to L2 proposer node
	toAddr := common.Address{0xff, 0xff}
	tx := types.MustSignNewTx(ethPrivKey, types.LatestSignerForChainID(cfg.L2ChainIDBig()), &types.DynamicFeeTx{
		ChainID:   cfg.L2ChainIDBig(),
		Nonce:     0,
		To:        &toAddr,
		Value:     big.NewInt(1_000_000_000),
		GasTipCap: big.NewInt(10),
		GasFeeCap: big.NewInt(200),
		Gas:       21000,
	})
	err = l2Prop.SendTransaction(context.Background(), tx)
	require.Nil(t, err, "Sending L2 tx to proposer")

	// Let it show up on the unsafe chain
	receipt, err := waitForL2Transaction(tx.Hash(), l2Prop, 3*time.Duration(cfg.DeployConfig.L1BlockTime)*time.Second)
	require.Nil(t, err, "Waiting for L2 tx on proposer")

	// Wait until the block it was first included in shows up in the safe chain on the syncer
	_, err = waitForL2Block(receipt.BlockNumber, l2Sync, time.Duration((sys.RollupConfig.ProposerWindowSize+4)*cfg.DeployConfig.L1BlockTime)*time.Second)
	require.Nil(t, err, "Waiting for block on syncer")

	// Assert that the transaction is not found on the syncer
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = l2Sync.TransactionReceipt(ctx, tx.Hash())
	require.Equal(t, ethereum.NotFound, err, "Found transaction in syncer when it should not have been included")

	// Wait a short time for the L2 reorg to occur on the proposer as well.
	// The proper thing to do is to wait until the proposer marks this block safe.
	<-time.After(2 * time.Second)

	// Assert that the reconciliation process did an L2 reorg on the proposer to remove the invalid block
	ctx2, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	block, err := l2Prop.BlockByNumber(ctx2, receipt.BlockNumber)
	require.Nil(t, err, "Get block from proposer")
	require.NotEqual(t, block.Hash(), receipt.BlockHash, "L2 Proposer did not reorg out transaction on it's safe chain")
}

func L1InfoFromState(ctx context.Context, contract *bindings.L1Block, l2Number *big.Int) (derive.L1BlockInfo, error) {
	var err error
	var out derive.L1BlockInfo
	opts := bind.CallOpts{
		BlockNumber: l2Number,
		Context:     ctx,
	}

	out.Number, err = contract.Number(&opts)
	if err != nil {
		return derive.L1BlockInfo{}, fmt.Errorf("failed to get number: %w", err)
	}

	out.Time, err = contract.Timestamp(&opts)
	if err != nil {
		return derive.L1BlockInfo{}, fmt.Errorf("failed to get timestamp: %w", err)
	}

	out.BaseFee, err = contract.Basefee(&opts)
	if err != nil {
		return derive.L1BlockInfo{}, fmt.Errorf("failed to get timestamp: %w", err)
	}

	blockHashBytes, err := contract.Hash(&opts)
	if err != nil {
		return derive.L1BlockInfo{}, fmt.Errorf("failed to get block hash: %w", err)
	}
	out.BlockHash = common.BytesToHash(blockHashBytes[:])

	out.SequenceNumber, err = contract.SequenceNumber(&opts)
	if err != nil {
		return derive.L1BlockInfo{}, fmt.Errorf("failed to get sequence number: %w", err)
	}

	overhead, err := contract.L1FeeOverhead(&opts)
	if err != nil {
		return derive.L1BlockInfo{}, fmt.Errorf("failed to get l1 fee overhead: %w", err)
	}
	out.L1FeeOverhead = eth.Bytes32(common.BigToHash(overhead))

	scalar, err := contract.L1FeeScalar(&opts)
	if err != nil {
		return derive.L1BlockInfo{}, fmt.Errorf("failed to get l1 fee scalar: %w", err)
	}
	out.L1FeeScalar = eth.Bytes32(common.BigToHash(scalar))

	batcherHash, err := contract.BatcherHash(&opts)
	if err != nil {
		return derive.L1BlockInfo{}, fmt.Errorf("failed to get batch sender: %w", err)
	}
	out.BatcherAddr = common.BytesToAddress(batcherHash[:])

	return out, nil
}

// TestSystemMockP2P sets up a L1 Geth node, a rollup node, and a L2 geth node and then confirms that
// the nodes can sync L2 blocks before they are confirmed on L1.
func TestSystemMockP2P(t *testing.T) {
	parallel(t)
	if !verboseGethNodes {
		log.Root().SetHandler(log.DiscardHandler())
	}

	cfg := DefaultSystemConfig(t)
	// slow down L1 blocks so we can see the L2 blocks arrive well before the L1 blocks do.
	// Keep the proposer window small so the L2 chain is started quick
	cfg.DeployConfig.L1BlockTime = 10

	// connect the nodes
	cfg.P2PTopology = map[string][]string{
		"syncer": {"proposer"},
	}

	var published, received []common.Hash
	propTracer, syncTracer := new(FnTracer), new(FnTracer)
	propTracer.OnPublishL2PayloadFn = func(ctx context.Context, payload *eth.ExecutionPayload) {
		published = append(published, payload.BlockHash)
	}
	syncTracer.OnUnsafeL2PayloadFn = func(ctx context.Context, from peer.ID, payload *eth.ExecutionPayload) {
		received = append(received, payload.BlockHash)
	}
	cfg.Nodes["proposer"].Tracer = propTracer
	cfg.Nodes["syncer"].Tracer = syncTracer

	sys, err := cfg.Start()
	require.Nil(t, err, "Error starting up system")
	defer sys.Close()

	l2Prop := sys.Clients["proposer"]
	l2Sync := sys.Clients["syncer"]

	// Transactor Account
	ethPrivKey := cfg.Secrets.Alice

	// Submit TX to L2 proposer node
	toAddr := common.Address{0xff, 0xff}
	tx := types.MustSignNewTx(ethPrivKey, types.LatestSignerForChainID(cfg.L2ChainIDBig()), &types.DynamicFeeTx{
		ChainID:   cfg.L2ChainIDBig(),
		Nonce:     0,
		To:        &toAddr,
		Value:     big.NewInt(1_000_000_000),
		GasTipCap: big.NewInt(10),
		GasFeeCap: big.NewInt(200),
		Gas:       21000,
	})
	err = l2Prop.SendTransaction(context.Background(), tx)
	require.Nil(t, err, "Sending L2 tx to proposer")

	// Wait for tx to be mined on the L2 proposer chain
	receiptProp, err := waitForL2Transaction(tx.Hash(), l2Prop, 6*time.Duration(sys.RollupConfig.BlockTime)*time.Second)
	require.Nil(t, err, "Waiting for L2 tx on proposer")

	// Wait until the block it was first included in shows up in the safe chain on the syncer
	receiptSync, err := waitForL2Transaction(tx.Hash(), l2Sync, 6*time.Duration(sys.RollupConfig.BlockTime)*time.Second)
	require.Nil(t, err, "Waiting for L2 tx on syncer")

	require.Equal(t, receiptProp, receiptSync)

	// Verify that everything that was received was published
	require.GreaterOrEqual(t, len(published), len(received))
	require.ElementsMatch(t, received, published[:len(received)])

	// Verify that the tx was received via p2p
	require.Contains(t, received, receiptSync.BlockHash)
}

// TestSystemMockP2P sets up a L1 Geth node, a rollup node, and a L2 geth node and then confirms that
// the nodes can sync L2 blocks before they are confirmed on L1.
//
// Test steps:
// 1. Spin up the nodes (P2P is disabled on the syncer)
// 2. Send a transaction to the proposer.
// 3. Wait for the TX to be mined on the proposer chain.
// 5. Wait for the syncer to detect a gap in the payload queue vs. the unsafe head
// 6. Wait for the RPC sync method to grab the block from the proposer over RPC and insert it into the syncer's unsafe chain.
// 7. Wait for the syncer to sync the unsafe chain into the safe chain.
// 8. Verify that the TX is included in the syncer's safe chain.
func TestSystemMockAltSync(t *testing.T) {
	parallel(t)
	if !verboseGethNodes {
		log.Root().SetHandler(log.DiscardHandler())
	}

	cfg := DefaultSystemConfig(t)
	// slow down L1 blocks so we can see the L2 blocks arrive well before the L1 blocks do.
	// Keep the proposer window small so the L2 chain is started quick
	cfg.DeployConfig.L1BlockTime = 10

	var published, received []common.Hash
	propTracer, syncTracer := new(FnTracer), new(FnTracer)
	propTracer.OnPublishL2PayloadFn = func(ctx context.Context, payload *eth.ExecutionPayload) {
		published = append(published, payload.BlockHash)
	}
	syncTracer.OnUnsafeL2PayloadFn = func(ctx context.Context, from peer.ID, payload *eth.ExecutionPayload) {
		received = append(received, payload.BlockHash)
	}
	cfg.Nodes["proposer"].Tracer = propTracer
	cfg.Nodes["syncer"].Tracer = syncTracer

	sys, err := cfg.Start(SystemConfigOption{
		key:  "afterRollupNodeStart",
		role: "proposer",
		action: func(sCfg *SystemConfig, system *System) {
			rpc, _ := system.Nodes["proposer"].Attach() // never errors
			cfg.Nodes["syncer"].L2Sync = &rollupNode.L2SyncRPCConfig{
				Rpc: client.NewBaseRPCClient(rpc),
			}
		},
	})
	require.Nil(t, err, "Error starting up system")
	defer sys.Close()

	l2Prop := sys.Clients["proposer"]
	l2Sync := sys.Clients["syncer"]

	// Transactor Account
	ethPrivKey := cfg.Secrets.Alice

	// Submit a TX to L2 proposer node
	toAddr := common.Address{0xff, 0xff}
	tx := types.MustSignNewTx(ethPrivKey, types.LatestSignerForChainID(cfg.L2ChainIDBig()), &types.DynamicFeeTx{
		ChainID:   cfg.L2ChainIDBig(),
		Nonce:     0,
		To:        &toAddr,
		Value:     big.NewInt(1_000_000_000),
		GasTipCap: big.NewInt(10),
		GasFeeCap: big.NewInt(200),
		Gas:       21000,
	})
	err = l2Prop.SendTransaction(context.Background(), tx)
	require.Nil(t, err, "Sending L2 tx to proposer")

	// Wait for tx to be mined on the L2 proposer chain
	receiptSeq, err := waitForTransaction(tx.Hash(), l2Prop, 6*time.Duration(sys.RollupConfig.BlockTime)*time.Second)
	require.Nil(t, err, "Waiting for L2 tx on proposer")

	// Wait for alt RPC sync to pick up the blocks on the proposer chain
	receiptVerif, err := waitForTransaction(tx.Hash(), l2Sync, 12*time.Duration(sys.RollupConfig.BlockTime)*time.Second)
	require.Nil(t, err, "Waiting for L2 tx on syncer")

	require.Equal(t, receiptSeq, receiptVerif)

	// Verify that the tx was received via RPC sync (P2P is disabled)
	require.Contains(t, received, receiptVerif.BlockHash)

	// Verify that everything that was received was published
	require.GreaterOrEqual(t, len(published), len(received))
	require.ElementsMatch(t, received, published[:len(received)])
}

// TestSystemDenseTopology sets up a dense p2p topology with 3 syncer nodes and 1 proposer node.
func TestSystemDenseTopology(t *testing.T) {
	t.Skip("Skipping dense topology test to avoid flakiness. @refcell address in p2p scoring pr.")

	parallel(t)
	if !verboseGethNodes {
		log.Root().SetHandler(log.DiscardHandler())
	}

	cfg := DefaultSystemConfig(t)
	// slow down L1 blocks so we can see the L2 blocks arrive well before the L1 blocks do.
	// Keep the proposer window small so the L2 chain is started quick
	cfg.DeployConfig.L1BlockTime = 10

	// Append additional nodes to the system to construct a dense p2p network
	cfg.Nodes["syncer2"] = &rollupNode.Config{
		Driver: driver.Config{
			SyncerConfDepth:   0,
			ProposerConfDepth: 0,
			ProposerEnabled:   false,
		},
		L1EpochPollInterval: time.Second * 4,
	}
	cfg.Nodes["syncer3"] = &rollupNode.Config{
		Driver: driver.Config{
			SyncerConfDepth:   0,
			ProposerConfDepth: 0,
			ProposerEnabled:   false,
		},
		L1EpochPollInterval: time.Second * 4,
	}
	cfg.Loggers["syncer2"] = testlog.Logger(t, log.LvlInfo).New("role", "syncer")
	cfg.Loggers["syncer3"] = testlog.Logger(t, log.LvlInfo).New("role", "syncer")

	// connect the nodes
	cfg.P2PTopology = map[string][]string{
		"syncer":  {"proposer", "syncer2", "syncer3"},
		"syncer2": {"proposer", "syncer", "syncer3"},
		"syncer3": {"proposer", "syncer", "syncer2"},
	}

	// Set peer scoring for each node, but without banning
	for _, node := range cfg.Nodes {
		params, err := p2p.GetPeerScoreParams("light", 2)
		require.NoError(t, err)
		node.P2P = &p2p.Config{
			PeerScoring:    params,
			BanningEnabled: false,
		}
	}

	var published, received1, received2, received3 []common.Hash
	propTracer, syncTracer, syncTracer2, syncTracer3 := new(FnTracer), new(FnTracer), new(FnTracer), new(FnTracer)
	propTracer.OnPublishL2PayloadFn = func(ctx context.Context, payload *eth.ExecutionPayload) {
		published = append(published, payload.BlockHash)
	}
	syncTracer.OnUnsafeL2PayloadFn = func(ctx context.Context, from peer.ID, payload *eth.ExecutionPayload) {
		received1 = append(received1, payload.BlockHash)
	}
	syncTracer2.OnUnsafeL2PayloadFn = func(ctx context.Context, from peer.ID, payload *eth.ExecutionPayload) {
		received2 = append(received2, payload.BlockHash)
	}
	syncTracer3.OnUnsafeL2PayloadFn = func(ctx context.Context, from peer.ID, payload *eth.ExecutionPayload) {
		received3 = append(received3, payload.BlockHash)
	}
	cfg.Nodes["proposer"].Tracer = propTracer
	cfg.Nodes["syncer"].Tracer = syncTracer
	cfg.Nodes["syncer2"].Tracer = syncTracer2
	cfg.Nodes["syncer3"].Tracer = syncTracer3

	sys, err := cfg.Start()
	require.Nil(t, err, "Error starting up system")
	defer sys.Close()

	l2Prop := sys.Clients["proposer"]
	l2Sync := sys.Clients["syncer"]
	l2Sync2 := sys.Clients["syncer2"]
	l2Sync3 := sys.Clients["syncer3"]

	// Transactor Account
	ethPrivKey := cfg.Secrets.Alice

	// Submit TX to L2 proposer node
	toAddr := common.Address{0xff, 0xff}
	tx := types.MustSignNewTx(ethPrivKey, types.LatestSignerForChainID(cfg.L2ChainIDBig()), &types.DynamicFeeTx{
		ChainID:   cfg.L2ChainIDBig(),
		Nonce:     0,
		To:        &toAddr,
		Value:     big.NewInt(1_000_000_000),
		GasTipCap: big.NewInt(10),
		GasFeeCap: big.NewInt(200),
		Gas:       21000,
	})
	err = l2Prop.SendTransaction(context.Background(), tx)
	require.NoError(t, err, "Sending L2 tx to proposer")

	// Wait for tx to be mined on the L2 proposer chain
	receiptSeq, err := waitForTransaction(tx.Hash(), l2Prop, 10*time.Duration(sys.RollupConfig.BlockTime)*time.Second)
	require.NoError(t, err, "Waiting for L2 tx on proposer")

	// Wait until the block it was first included in shows up in the safe chain on the syncer
	receiptVerif, err := waitForTransaction(tx.Hash(), l2Sync, 10*time.Duration(sys.RollupConfig.BlockTime)*time.Second)
	require.NoError(t, err, "Waiting for L2 tx on syncer")
	require.Equal(t, receiptSeq, receiptVerif)

	receiptVerif, err = waitForTransaction(tx.Hash(), l2Sync2, 10*time.Duration(sys.RollupConfig.BlockTime)*time.Second)
	require.NoError(t, err, "Waiting for L2 tx on syncer2")
	require.Equal(t, receiptSeq, receiptVerif)

	receiptVerif, err = waitForTransaction(tx.Hash(), l2Sync3, 10*time.Duration(sys.RollupConfig.BlockTime)*time.Second)
	require.NoError(t, err, "Waiting for L2 tx on syncer3")
	require.Equal(t, receiptSeq, receiptVerif)

	// Verify that everything that was received was published
	require.GreaterOrEqual(t, len(published), len(received1))
	require.GreaterOrEqual(t, len(published), len(received2))
	require.GreaterOrEqual(t, len(published), len(received3))
	require.ElementsMatch(t, published, received1[:len(published)])
	require.ElementsMatch(t, published, received2[:len(published)])
	require.ElementsMatch(t, published, received3[:len(published)])

	// Verify that the tx was received via p2p
	require.Contains(t, received1, receiptVerif.BlockHash)
	require.Contains(t, received2, receiptVerif.BlockHash)
	require.Contains(t, received3, receiptVerif.BlockHash)
}

func TestL1InfoContract(t *testing.T) {
	parallel(t)
	if !verboseGethNodes {
		log.Root().SetHandler(log.DiscardHandler())
	}

	cfg := DefaultSystemConfig(t)

	sys, err := cfg.Start()
	require.Nil(t, err, "Error starting up system")
	defer sys.Close()

	l1Client := sys.Clients["l1"]
	l2Prop := sys.Clients["proposer"]
	l2Sync := sys.Clients["syncer"]

	endSyncBlockNumber := big.NewInt(4)
	endPropBlockNumber := big.NewInt(6)
	endSyncBlock, err := waitForL2Block(endSyncBlockNumber, l2Sync, time.Minute)
	require.Nil(t, err)
	endPropBlock, err := waitForL2Block(endPropBlockNumber, l2Prop, time.Minute)
	require.Nil(t, err)

	propL1Info, err := bindings.NewL1Block(cfg.L1InfoPredeployAddress, l2Prop)
	require.Nil(t, err)

	syncL1Info, err := bindings.NewL1Block(cfg.L1InfoPredeployAddress, l2Sync)
	require.Nil(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	fillInfoLists := func(start *types.Block, contract *bindings.L1Block, client *ethclient.Client) ([]derive.L1BlockInfo, []derive.L1BlockInfo) {
		var txList, stateList []derive.L1BlockInfo
		for b := start; ; {
			var infoFromTx derive.L1BlockInfo
			require.NoError(t, infoFromTx.UnmarshalBinary(b.Transactions()[0].Data()))
			txList = append(txList, infoFromTx)

			infoFromState, err := L1InfoFromState(ctx, contract, b.Number())
			require.Nil(t, err)
			stateList = append(stateList, infoFromState)

			// Genesis L2 block contains no L1 Deposit TX
			if b.NumberU64() == 1 {
				return txList, stateList
			}
			b, err = client.BlockByHash(ctx, b.ParentHash())
			require.Nil(t, err)
		}
	}

	l1InfosFromProposerTransactions, l1InfosFromProposerState := fillInfoLists(endPropBlock, propL1Info, l2Prop)
	l1InfosFromSyncerTransactions, l1InfosFromSyncerState := fillInfoLists(endSyncBlock, syncL1Info, l2Sync)

	l1blocks := make(map[common.Hash]derive.L1BlockInfo)
	maxL1Hash := l1InfosFromProposerTransactions[0].BlockHash
	for h := maxL1Hash; ; {
		b, err := l1Client.BlockByHash(ctx, h)
		require.Nil(t, err)

		l1blocks[h] = derive.L1BlockInfo{
			Number:         b.NumberU64(),
			Time:           b.Time(),
			BaseFee:        b.BaseFee(),
			BlockHash:      h,
			SequenceNumber: 0, // ignored, will be overwritten
			BatcherAddr:    sys.RollupConfig.Genesis.SystemConfig.BatcherAddr,
			L1FeeOverhead:  sys.RollupConfig.Genesis.SystemConfig.Overhead,
			L1FeeScalar:    sys.RollupConfig.Genesis.SystemConfig.Scalar,
		}

		h = b.ParentHash()
		if b.NumberU64() == 0 {
			break
		}
	}

	checkInfoList := func(name string, list []derive.L1BlockInfo) {
		for _, info := range list {
			if expected, ok := l1blocks[info.BlockHash]; ok {
				expected.SequenceNumber = info.SequenceNumber // the seq nr is not part of the L1 info we know in advance, so we ignore it.
				require.Equal(t, expected, info)
			} else {
				t.Fatalf("Did not find block hash for L1 Info: %v in test %s", info, name)
			}
		}
	}

	checkInfoList("On proposer with tx", l1InfosFromProposerTransactions)
	checkInfoList("On proposer with state", l1InfosFromProposerState)
	checkInfoList("On syncer with tx", l1InfosFromSyncerTransactions)
	checkInfoList("On syncer with state", l1InfosFromSyncerState)

}

// calcGasFees determines the actual cost of the transaction given a specific basefee
// This does not include the L1 data fee charged from L2 transactions.
func calcGasFees(gasUsed uint64, gasTipCap *big.Int, gasFeeCap *big.Int, baseFee *big.Int) *big.Int {
	x := new(big.Int).Add(gasTipCap, baseFee)
	// If tip + basefee > gas fee cap, clamp it to the gas fee cap
	if x.Cmp(gasFeeCap) > 0 {
		x = gasFeeCap
	}
	return x.Mul(x, new(big.Int).SetUint64(gasUsed))
}

// calcL1GasUsed returns the gas used to include the transaction data in
// the calldata on L1
func calcL1GasUsed(data []byte, overhead *big.Int) *big.Int {
	var zeroes, ones uint64
	for _, byt := range data {
		if byt == 0 {
			zeroes++
		} else {
			ones++
		}
	}

	zeroesGas := zeroes * 4 // params.TxDataZeroGas
	onesGas := ones * 16    // params.TxDataNonZeroGasEIP2028
	l1Gas := new(big.Int).SetUint64(zeroesGas + onesGas)
	return new(big.Int).Add(l1Gas, overhead)
}

// TestWithdrawals checks that a deposit and then withdrawal execution succeeds. It verifies the
// balance changes on L1 and L2 and has to include gas fees in the balance checks.
// It does not check that the withdrawal can be executed prior to the end of the finality period.
func TestWithdrawals(t *testing.T) {
	parallel(t)
	if !verboseGethNodes {
		log.Root().SetHandler(log.DiscardHandler())
	}

	cfg := DefaultSystemConfig(t)
	cfg.DeployConfig.FinalizationPeriodSeconds = 2 // 2s finalization period

	sys, err := cfg.Start()
	require.Nil(t, err, "Error starting up system")
	defer sys.Close()

	l1Client := sys.Clients["l1"]
	l2Prop := sys.Clients["proposer"]
	l2Sync := sys.Clients["syncer"]

	// Transactor Account
	ethPrivKey := cfg.Secrets.Alice
	fromAddr := crypto.PubkeyToAddress(ethPrivKey.PublicKey)

	// Find deposit contract
	depositContract, err := bindings.NewKromaPortal(predeploys.DevKromaPortalAddr, l1Client)
	require.Nil(t, err)

	// Create L1 signer
	opts, err := bind.NewKeyedTransactorWithChainID(ethPrivKey, cfg.L1ChainIDBig())
	require.Nil(t, err)

	// Start L2 balance
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	startBalance, err := l2Sync.BalanceAt(ctx, fromAddr, nil)
	require.Nil(t, err)

	// Finally send TX
	mintAmount := big.NewInt(1_000_000_000_000)
	opts.Value = mintAmount
	tx, err := depositContract.DepositTransaction(opts, fromAddr, common.Big0, 1_000_000, false, nil)
	require.Nil(t, err, "with deposit tx")

	receipt, err := waitForTransaction(tx.Hash(), l1Client, 3*time.Duration(cfg.DeployConfig.L1BlockTime)*time.Second)
	require.Nil(t, err, "Waiting for deposit tx on L1")

	// Bind L2 Withdrawer Contract
	l2withdrawer, err := bindings.NewL2ToL1MessagePasser(predeploys.L2ToL1MessagePasserAddr, l2Prop)
	require.Nil(t, err, "binding withdrawer on L2")

	// Wait for deposit to arrive
	reconstructedDep, err := derive.UnmarshalDepositLogEvent(receipt.Logs[0])
	require.NoError(t, err, "Could not reconstruct L2 Deposit")
	tx = types.NewTx(reconstructedDep)
	receipt, err = waitForL2Transaction(tx.Hash(), l2Sync, 10*time.Duration(cfg.DeployConfig.L1BlockTime)*time.Second)
	require.NoError(t, err)
	require.Equal(t, receipt.Status, types.ReceiptStatusSuccessful)

	// Confirm L2 balance
	ctx, cancel = context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	endBalance, err := l2Sync.BalanceAt(ctx, fromAddr, nil)
	require.Nil(t, err)

	diff := new(big.Int)
	diff = diff.Sub(endBalance, startBalance)
	require.Equal(t, mintAmount, diff, "Did not get expected balance change after mint")

	// Start L2 balance for withdrawal
	ctx, cancel = context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	startBalance, err = l2Prop.BalanceAt(ctx, fromAddr, nil)
	require.Nil(t, err)

	// Initiate Withdrawal
	withdrawAmount := big.NewInt(500_000_000_000)
	l2opts, err := bind.NewKeyedTransactorWithChainID(ethPrivKey, cfg.L2ChainIDBig())
	require.Nil(t, err)
	l2opts.Value = withdrawAmount
	tx, err = l2withdrawer.InitiateWithdrawal(l2opts, fromAddr, big.NewInt(21000), nil)
	require.Nil(t, err, "sending initiate withdraw tx")

	receipt, err = waitForL2Transaction(tx.Hash(), l2Sync, 10*time.Duration(cfg.DeployConfig.L1BlockTime)*time.Second)
	require.Nil(t, err, "withdrawal initiated on L2 proposer")
	require.Equal(t, receipt.Status, types.ReceiptStatusSuccessful, "transaction failed")

	// Verify L2 balance after withdrawal
	ctx, cancel = context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	header, err := l2Sync.HeaderByNumber(ctx, receipt.BlockNumber)
	require.Nil(t, err)

	ctx, cancel = context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	endBalance, err = l2Sync.BalanceAt(ctx, fromAddr, nil)
	require.Nil(t, err)

	// Take fee into account
	diff = new(big.Int).Sub(startBalance, endBalance)
	fees := calcGasFees(receipt.GasUsed, tx.GasTipCap(), tx.GasFeeCap(), header.BaseFee)
	fees = fees.Add(fees, receipt.L1Fee)
	diff = diff.Sub(diff, fees)
	require.Equal(t, withdrawAmount, diff)

	// Take start balance on L1
	ctx, cancel = context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	startBalance, err = l1Client.BalanceAt(ctx, fromAddr, nil)
	require.Nil(t, err)

	// Get l2BlockNumber for proof generation
	ctx, cancel = context.WithTimeout(context.Background(), 40*time.Duration(cfg.DeployConfig.L1BlockTime)*time.Second*timeoutMultiplier)
	defer cancel()
	blockNumber, err := withdrawals.WaitForFinalizationPeriod(ctx, l1Client, predeploys.DevKromaPortalAddr, receipt.BlockNumber)
	require.Nil(t, err)

	ctx, cancel = context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	header, err = l2Sync.HeaderByNumber(ctx, new(big.Int).SetUint64(blockNumber))
	require.Nil(t, err)

	var nextHeader *types.Header = nil
	if sys.RollupConfig.IsBlue(header.Time) {
		ctx, cancel = context.WithTimeout(context.Background(), 1*time.Second)
		nextHeader, err = l2Sync.HeaderByNumber(ctx, new(big.Int).SetUint64(blockNumber+1))
		defer cancel()
		require.Nil(t, err)
	}

	rpcClient, err := rpc.Dial(sys.Nodes["syncer"].WSEndpoint())
	require.Nil(t, err)
	proofCl := gethclient.New(rpcClient)
	receiptCl := ethclient.NewClient(rpcClient)

	// Now create withdrawal
	oracle, err := bindings.NewL2OutputOracleCaller(predeploys.DevL2OutputOracleAddr, l1Client)
	require.Nil(t, err)

	params, err := withdrawals.ProveWithdrawalParameters(context.Background(), rollup.L2OutputRootVersion(sys.RollupConfig, header.Time), proofCl, receiptCl, tx.Hash(), header, nextHeader, oracle)
	require.Nil(t, err)

	portal, err := bindings.NewKromaPortal(predeploys.DevKromaPortalAddr, l1Client)
	require.Nil(t, err)

	opts.Value = nil

	// Prove withdrawal
	tx, err = portal.ProveWithdrawalTransaction(
		opts,
		bindings.TypesWithdrawalTransaction{
			Nonce:    params.Nonce,
			Sender:   params.Sender,
			Target:   params.Target,
			Value:    params.Value,
			GasLimit: params.GasLimit,
			Data:     params.Data,
		},
		params.L2OutputIndex,
		params.OutputRootProof,
		params.WithdrawalProof,
	)
	require.Nil(t, err)

	// Ensure that our withdrawal was proved successfully
	proveReceipt, err := waitForTransaction(tx.Hash(), l1Client, 3*time.Duration(cfg.DeployConfig.L1BlockTime)*time.Second)
	require.Nil(t, err, "prove withdrawal")
	require.Equal(t, types.ReceiptStatusSuccessful, proveReceipt.Status)

	// Wait for finalization and then create the Finalized Withdrawal Transaction
	ctx, cancel = context.WithTimeout(context.Background(), 30*time.Duration(cfg.DeployConfig.L1BlockTime)*time.Second)
	defer cancel()
	_, err = withdrawals.WaitForFinalizationPeriod(ctx, l1Client, predeploys.DevKromaPortalAddr, header.Number)
	require.Nil(t, err)

	// Finalize withdrawal
	tx, err = portal.FinalizeWithdrawalTransaction(
		opts,
		bindings.TypesWithdrawalTransaction{
			Nonce:    params.Nonce,
			Sender:   params.Sender,
			Target:   params.Target,
			Value:    params.Value,
			GasLimit: params.GasLimit,
			Data:     params.Data,
		},
	)
	require.Nil(t, err)

	// Ensure that our withdrawal was finalized successfully
	finalizeReceipt, err := waitForTransaction(tx.Hash(), l1Client, 3*time.Duration(cfg.DeployConfig.L1BlockTime)*time.Second)
	require.Nil(t, err, "finalize withdrawal")
	require.Equal(t, types.ReceiptStatusSuccessful, finalizeReceipt.Status)

	// Verify balance after withdrawal
	ctx, cancel = context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	header, err = l1Client.HeaderByNumber(ctx, finalizeReceipt.BlockNumber)
	require.Nil(t, err)

	ctx, cancel = context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	endBalance, err = l1Client.BalanceAt(ctx, fromAddr, nil)
	require.Nil(t, err)

	// Ensure that withdrawal - gas fees are added to the L1 balance
	// Fun fact, the fee is greater than the withdrawal amount
	// NOTE: The gas fees include *both* the ProveWithdrawalTransaction and FinalizeWithdrawalTransaction transactions.
	diff = new(big.Int).Sub(endBalance, startBalance)
	fees = calcGasFees(proveReceipt.GasUsed+finalizeReceipt.GasUsed, tx.GasTipCap(), tx.GasFeeCap(), header.BaseFee)
	withdrawAmount = withdrawAmount.Sub(withdrawAmount, fees)
	require.Equal(t, withdrawAmount, diff)
}

// TestFees checks that L1/L2 fees are handled.
func TestFees(t *testing.T) {
	parallel(t)
	if !verboseGethNodes {
		log.Root().SetHandler(log.DiscardHandler())
	}

	cfg := DefaultSystemConfig(t)
	// TODO: after we have the system config contract and new kroma-geth L1 cost utils,
	// we can pull in l1 costs into every e2e test and account for it in assertions easily etc.
	cfg.DeployConfig.GasPriceOracleOverhead = 2100
	cfg.DeployConfig.GasPriceOracleScalar = 1000_000

	sys, err := cfg.Start()
	require.Nil(t, err, "Error starting up system")
	defer sys.Close()

	l2Prop := sys.Clients["proposer"]
	l2Sync := sys.Clients["syncer"]

	// Transactor Account
	ethPrivKey := cfg.Secrets.Alice
	fromAddr := crypto.PubkeyToAddress(ethPrivKey.PublicKey)

	// Find gaspriceoracle contract
	gpoContract, err := bindings.NewGasPriceOracle(predeploys.GasPriceOracleAddr, l2Prop)
	require.Nil(t, err)

	overhead, err := gpoContract.Overhead(&bind.CallOpts{})
	require.Nil(t, err, "reading gpo overhead")
	decimals, err := gpoContract.DECIMALS(&bind.CallOpts{})
	require.Nil(t, err, "reading gpo decimals")
	scalar, err := gpoContract.Scalar(&bind.CallOpts{})
	require.Nil(t, err, "reading gpo scalar")

	require.Equal(t, overhead.Uint64(), uint64(2100), "wrong gpo overhead")
	require.Equal(t, decimals.Uint64(), uint64(6), "wrong gpo decimals")
	require.Equal(t, scalar.Uint64(), uint64(1_000_000), "wrong gpo scalar")

	// Check balances of ProtocolVault
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	protocolVaultStartBalance, err := l2Prop.BalanceAt(ctx, predeploys.ProtocolVaultAddr, nil)
	require.Nil(t, err)

	// Check balances of ProposerRewardVault
	ctx, cancel = context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	proposerRewardVaultStartBalance, err := l2Prop.BalanceAt(ctx, predeploys.ProposerRewardVaultAddr, nil)
	require.Nil(t, err)

	// Simple transfer from signer to random account
	ctx, cancel = context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	startBalance, err := l2Sync.BalanceAt(ctx, fromAddr, nil)
	require.Nil(t, err)

	toAddr := common.Address{0xff, 0xff}
	transferAmount := big.NewInt(1_000_000_000)
	gasTip := big.NewInt(10)
	tx := types.MustSignNewTx(ethPrivKey, types.LatestSignerForChainID(cfg.L2ChainIDBig()), &types.DynamicFeeTx{
		ChainID:   cfg.L2ChainIDBig(),
		Nonce:     0,
		To:        &toAddr,
		Value:     transferAmount,
		GasTipCap: gasTip,
		GasFeeCap: big.NewInt(200),
		Gas:       21000,
	})
	sender, err := types.LatestSignerForChainID(cfg.L2ChainIDBig()).Sender(tx)
	require.NoError(t, err)
	t.Logf("waiting for tx %s from %s to %s", tx.Hash(), sender, tx.To())
	err = l2Prop.SendTransaction(context.Background(), tx)
	require.Nil(t, err, "Sending L2 tx to proposer")

	_, err = waitForL2Transaction(tx.Hash(), l2Prop, 4*time.Duration(cfg.DeployConfig.L1BlockTime)*time.Second)
	require.Nil(t, err, "Waiting for L2 tx on proposer")

	receipt, err := waitForL2Transaction(tx.Hash(), l2Sync, 4*time.Duration(cfg.DeployConfig.L1BlockTime)*time.Second)
	require.Nil(t, err, "Waiting for L2 tx on syncer")
	require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status, "TX should have succeeded")

	ctx, cancel = context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	header, err := l2Prop.HeaderByNumber(ctx, receipt.BlockNumber)
	require.Nil(t, err)

	ctx, cancel = context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	validatorRewardVaultStartBalance, err := l2Prop.BalanceAt(ctx, header.Coinbase, safeAddBig(header.Number, big.NewInt(-1)))
	require.Nil(t, err)

	ctx, cancel = context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	validatorRewardVaultEndBalance, err := l2Prop.BalanceAt(ctx, header.Coinbase, header.Number)
	require.Nil(t, err)

	ctx, cancel = context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	endBalance, err := l2Prop.BalanceAt(ctx, fromAddr, header.Number)
	require.Nil(t, err)

	ctx, cancel = context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	protocolVaultEndBalance, err := l2Prop.BalanceAt(ctx, predeploys.ProtocolVaultAddr, header.Number)
	require.Nil(t, err)

	l1Header, err := sys.Clients["l1"].HeaderByNumber(ctx, nil)
	require.Nil(t, err)

	ctx, cancel = context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	proposerRewardVaultEndBalance, err := l2Prop.BalanceAt(ctx, predeploys.ProposerRewardVaultAddr, header.Number)
	require.Nil(t, err)

	// Diff fee recipients balances
	protocolVaultDiff := new(big.Int).Sub(protocolVaultEndBalance, protocolVaultStartBalance)
	proposerRewardVaultDiff := new(big.Int).Sub(proposerRewardVaultEndBalance, proposerRewardVaultStartBalance)
	validatorRewardVaultDiff := new(big.Int).Sub(validatorRewardVaultEndBalance, validatorRewardVaultStartBalance)

	// Tally L2 Fee
	l2Fee := gasTip.Mul(gasTip, new(big.Int).SetUint64(receipt.GasUsed))
	require.Equal(t, l2Fee, validatorRewardVaultDiff, "l2 fee mismatch")

	// Tally BaseFee
	baseFee := new(big.Int).Mul(header.BaseFee, new(big.Int).SetUint64(receipt.GasUsed))
	require.Equal(t, baseFee, protocolVaultDiff, "base fee fee mismatch")

	// Tally proposer reward
	bytes, err := tx.MarshalBinary()
	require.Nil(t, err)
	l1GasUsed := calcL1GasUsed(bytes, overhead)
	divisor := new(big.Int).Exp(big.NewInt(10), decimals, nil)
	l1Fee := new(big.Int).Mul(l1GasUsed, l1Header.BaseFee)
	l1Fee = l1Fee.Mul(l1Fee, scalar)
	l1Fee = l1Fee.Div(l1Fee, divisor)

	require.Equal(t, l1Fee, proposerRewardVaultDiff, "proposer reward mismatch")

	// Tally Proposer reward against GasPriceOracle
	gpoL1Fee, err := gpoContract.GetL1Fee(&bind.CallOpts{}, bytes)
	require.Nil(t, err)
	require.Equal(t, l1Fee, gpoL1Fee, "proposer reward mismatch")

	// Calculate total fee
	protocolVaultDiff.Add(protocolVaultDiff, validatorRewardVaultDiff)
	totalFee := new(big.Int).Add(protocolVaultDiff, proposerRewardVaultDiff)
	balanceDiff := new(big.Int).Sub(startBalance, endBalance)
	balanceDiff.Sub(balanceDiff, transferAmount)
	require.Equal(t, balanceDiff, totalFee, "balances should add up")
}

func TestStopStartProposer(t *testing.T) {
	parallel(t)
	if !verboseGethNodes {
		log.Root().SetHandler(log.DiscardHandler())
	}

	cfg := DefaultSystemConfig(t)
	sys, err := cfg.Start()
	require.Nil(t, err, "Error starting up system")
	defer sys.Close()

	l2Prop := sys.Clients["proposer"]
	rollupNode := sys.RollupNodes["proposer"]

	nodeRPC, err := rpc.DialContext(context.Background(), rollupNode.HTTPEndpoint())
	require.Nil(t, err, "Error dialing node")

	blockBefore := latestBlock(t, l2Prop)
	time.Sleep(time.Duration(cfg.DeployConfig.L2BlockTime+1) * time.Second)
	blockAfter := latestBlock(t, l2Prop)
	require.Greaterf(t, blockAfter, blockBefore, "Chain did not advance")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	blockHash := common.Hash{}
	err = nodeRPC.CallContext(ctx, &blockHash, "admin_stopProposer")
	require.Nil(t, err, "Error stopping proposer")

	blockBefore = latestBlock(t, l2Prop)
	time.Sleep(time.Duration(cfg.DeployConfig.L2BlockTime+1) * time.Second)
	blockAfter = latestBlock(t, l2Prop)
	require.Equal(t, blockAfter, blockBefore, "Chain advanced after stopping proposer")

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = nodeRPC.CallContext(ctx, nil, "admin_startProposer", blockHash)
	require.Nil(t, err, "Error starting proposer")

	blockBefore = latestBlock(t, l2Prop)
	time.Sleep(time.Duration(cfg.DeployConfig.L2BlockTime+1) * time.Second)
	blockAfter = latestBlock(t, l2Prop)
	require.Greater(t, blockAfter, blockBefore, "Chain did not advance after starting proposer")
}

func TestStopStartBatcher(t *testing.T) {
	parallel(t)
	if !verboseGethNodes {
		log.Root().SetHandler(log.DiscardHandler())
	}

	cfg := DefaultSystemConfig(t)
	sys, err := cfg.Start()
	require.Nil(t, err, "Error starting up system")
	defer sys.Close()

	rollupRPCClient, err := rpc.DialContext(context.Background(), sys.RollupNodes["syncer"].HTTPEndpoint())
	require.Nil(t, err)
	rollupClient := sources.NewRollupClient(client.NewBaseRPCClient(rollupRPCClient))

	l2Prop := sys.Clients["proposer"]
	l2Sync := sys.Clients["syncer"]

	// retrieve the initial sync status
	propStatus, err := rollupClient.SyncStatus(context.Background())
	require.Nil(t, err)

	nonce := uint64(0)
	sendTx := func() *types.Receipt {
		// Submit TX to L2 proposer node
		tx := types.MustSignNewTx(cfg.Secrets.Alice, types.LatestSignerForChainID(cfg.L2ChainIDBig()), &types.DynamicFeeTx{
			ChainID:   cfg.L2ChainIDBig(),
			Nonce:     nonce,
			To:        &common.Address{0xff, 0xff},
			Value:     big.NewInt(1_000_000_000),
			GasTipCap: big.NewInt(10),
			GasFeeCap: big.NewInt(200),
			Gas:       21000,
		})
		nonce++
		err = l2Prop.SendTransaction(context.Background(), tx)
		require.Nil(t, err, "Sending L2 tx to proposer")

		// Let it show up on the unsafe chain
		receipt, err := waitForTransaction(tx.Hash(), l2Prop, 3*time.Duration(cfg.DeployConfig.L1BlockTime)*time.Second)
		require.Nil(t, err, "Waiting for L2 tx on proposer")

		return receipt
	}
	// send a transaction
	receipt := sendTx()

	// wait until the block the tx was first included in shows up in the safe chain on the syncer
	safeBlockInclusionDuration := time.Duration(3*cfg.DeployConfig.L1BlockTime) * time.Second
	_, err = waitForL2Block(receipt.BlockNumber, l2Sync, safeBlockInclusionDuration)
	require.Nil(t, err, "Waiting for block on syncer")

	// ensure the safe chain advances
	newSeqStatus, err := rollupClient.SyncStatus(context.Background())
	require.Nil(t, err)
	require.Greater(t, newSeqStatus.SafeL2.Number, propStatus.SafeL2.Number, "Safe chain did not advance")

	// stop the batch submission
	sys.Batcher.Stop()

	// wait for any old safe blocks being submitted / derived
	time.Sleep(safeBlockInclusionDuration)

	// get the initial sync status
	propStatus, err = rollupClient.SyncStatus(context.Background())
	require.Nil(t, err)

	// send another tx
	sendTx()
	time.Sleep(safeBlockInclusionDuration)

	// ensure that the safe chain does not advance while the batcher is stopped
	newSeqStatus, err = rollupClient.SyncStatus(context.Background())
	require.Nil(t, err)
	require.Equal(t, newSeqStatus.SafeL2.Number, propStatus.SafeL2.Number, "Safe chain advanced while batcher was stopped")

	// start the batch submission
	sys.Batcher.Start()
	time.Sleep(safeBlockInclusionDuration)

	// send a third tx
	receipt = sendTx()

	// wait until the block the tx was first included in shows up in the safe chain on the syncer
	_, err = waitForL2Block(receipt.BlockNumber, l2Sync, safeBlockInclusionDuration)
	require.Nil(t, err, "Waiting for block on syncer")

	// ensure that the safe chain advances after restarting the batcher
	newSeqStatus, err = rollupClient.SyncStatus(context.Background())
	require.Nil(t, err)
	require.Greater(t, newSeqStatus.SafeL2.Number, propStatus.SafeL2.Number, "Safe chain did not advance after batcher was restarted")
}

func safeAddBig(a *big.Int, b *big.Int) *big.Int {
	return new(big.Int).Add(a, b)
}

func latestBlock(t *testing.T, client *ethclient.Client) uint64 {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	blockAfter, err := client.BlockNumber(ctx)
	require.Nil(t, err, "Error getting latest block")
	return blockAfter
}
