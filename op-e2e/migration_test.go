package op_e2e

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/eth/ethconfig"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/kroma-network/zktrie/trie"
	"github.com/stretchr/testify/require"

	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils/geth"
	rollupNode "github.com/ethereum-optimism/optimism/op-node/node"
	"github.com/ethereum-optimism/optimism/op-node/rollup/driver"
	"github.com/ethereum-optimism/optimism/op-node/rollup/sync"
	"github.com/ethereum-optimism/optimism/op-service/testlog"
	"github.com/kroma-network/kroma/kroma-bindings/predeploys"
	"github.com/kroma-network/kroma/op-e2e/e2eutils/wait"
)

func TestMigration(t *testing.T) {
	InitParallel(t)

	zero := hexutil.Uint64(0)
	one := hexutil.Uint64(1)
	mptTimeOffset := hexutil.Uint64(60)

	cfg := DefaultSystemConfig(t)
	cfg.DeployConfig.L2GenesisDeltaTimeOffset = &zero
	cfg.DeployConfig.L2GenesisEcotoneTimeOffset = &one
	cfg.DeployConfig.L2GenesisKromaMptTimeOffset = &mptTimeOffset

	// Setup historical rpc node.
	historicalRpcPort := 8045
	cfg.Nodes["historical"] = &rollupNode.Config{
		Driver: driver.Config{
			VerifierConfDepth:  0,
			SequencerConfDepth: 0,
			SequencerEnabled:   false,
		},
		L1EpochPollInterval:         time.Second * 4,
		RuntimeConfigReloadInterval: time.Minute * 10,
		ConfigPersistence:           &rollupNode.DisabledConfigPersistence{},
		Sync:                        sync.Config{SyncMode: sync.CLSync},
	}
	cfg.Loggers["historical"] = testlog.Logger(t, log.LevelInfo).New("role", "historical")
	cfg.GethOptions["historical"] = append(cfg.GethOptions["historical"], []geth.GethOption{
		func(ethCfg *ethconfig.Config, nodeCfg *node.Config) error {
			nodeCfg.HTTPPort = historicalRpcPort
			nodeCfg.HTTPModules = []string{"debug", "eth"}
			nodeCfg.HTTPHost = "127.0.0.1"
			return nil
		},
	}...)

	// Set historical rpc endpoint.
	for name := range cfg.Nodes {
		name := name
		cfg.GethOptions[name] = append(cfg.GethOptions[name], []geth.GethOption{
			func(ethCfg *ethconfig.Config, nodeCfg *node.Config) error {
				// Since the migration process requires preimages, enable storing preimage option.
				ethCfg.Preimages = true
				if name != "historical" {
					ethCfg.RollupHistoricalRPC = fmt.Sprintf("http://127.0.0.1:%d", historicalRpcPort)
				}
				// Deep copy the genesis
				dst := &core.Genesis{}
				b, _ := json.Marshal(ethCfg.Genesis)
				err := json.Unmarshal(b, dst)
				if err != nil {
					return err
				}
				ethCfg.Genesis = dst
				return nil
			},
		}...)
	}

	sys, err := cfg.Start(t)
	defer sys.Close()
	require.Nil(t, err, "Error starting up system")
	l1Cl := sys.Clients["l1"]
	l2Seq := sys.Clients["sequencer"]
	l2Verif := sys.Clients["verifier"]

	// Send L2 TX
	_, err = wait.ForTransferTxOnL2(sys.Cfg.L2ChainIDBig(), l2Seq, l2Verif, cfg.Secrets.Alice, common.Address{0xff, 0xff}, common.Big1)
	require.NoError(t, err)

	// Deploy contract on L2
	minimalCode := common.Hex2Bytes("601680600c6000396000f3fe3615600b576000803555005b60016000540160005500")
	chainId := cfg.L2ChainIDBig()
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	nonce, err := l2Seq.PendingNonceAt(ctx, cfg.Secrets.Addresses().Alice)
	require.NoError(t, err)
	tx := types.MustSignNewTx(cfg.Secrets.Alice, types.LatestSignerForChainID(chainId), &types.DynamicFeeTx{
		ChainID:   chainId,
		Nonce:     nonce,
		To:        nil,
		GasTipCap: big.NewInt(10),
		GasFeeCap: big.NewInt(200),
		Gas:       100_000,
		Data:      minimalCode,
	})
	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Duration(cfg.DeployConfig.L1BlockTime)*time.Second)
	defer cancel()
	err = l2Seq.SendTransaction(ctx, tx)
	require.NoError(t, err)
	deployReceipt, err := geth.WaitForTransaction(tx.Hash(), l2Verif, 4*time.Duration(cfg.DeployConfig.L1BlockTime)*time.Second)
	require.NoError(t, err)
	// Set storage
	tx = types.MustSignNewTx(cfg.Secrets.Alice, types.LatestSignerForChainID(chainId), &types.DynamicFeeTx{
		ChainID:   chainId,
		Nonce:     tx.Nonce() + 1,
		To:        &deployReceipt.ContractAddress,
		GasTipCap: big.NewInt(10),
		GasFeeCap: big.NewInt(200),
		Gas:       100_000,
	})
	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Duration(cfg.DeployConfig.L1BlockTime)*time.Second)
	defer cancel()
	err = l2Seq.SendTransaction(ctx, tx)
	require.NoError(t, err)
	_, err = geth.WaitForTransaction(tx.Hash(), l2Verif, 4*time.Duration(cfg.DeployConfig.L1BlockTime)*time.Second)
	require.NoError(t, err)
	// Check storage value
	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	slot := common.Hash{}
	value, err := l2Verif.StorageAt(ctx, *tx.To(), slot, nil)
	require.NoError(t, err)
	require.True(t, new(big.Int).SetBytes(value).Uint64() > 0)
	// Delete storage
	tx = types.MustSignNewTx(cfg.Secrets.Alice, types.LatestSignerForChainID(chainId), &types.DynamicFeeTx{
		ChainID:   chainId,
		Nonce:     tx.Nonce() + 1,
		To:        &deployReceipt.ContractAddress,
		GasTipCap: big.NewInt(10),
		GasFeeCap: big.NewInt(200),
		Gas:       100_000,
		Data:      slot.Bytes(),
	})
	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Duration(cfg.DeployConfig.L1BlockTime)*time.Second)
	defer cancel()
	err = l2Seq.SendTransaction(ctx, tx)
	require.NoError(t, err)
	_, err = geth.WaitForTransaction(tx.Hash(), l2Verif, 4*time.Duration(cfg.DeployConfig.L1BlockTime)*time.Second)
	require.NoError(t, err)
	// Check storage value
	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	value, err = l2Verif.StorageAt(ctx, *tx.To(), slot, nil)
	require.NoError(t, err)
	require.True(t, new(big.Int).SetBytes(value).Uint64() == 0)

	transitionBlockNumber := new(big.Int).SetUint64(uint64(mptTimeOffset) / cfg.DeployConfig.L2BlockTime)
	_, err = geth.WaitForBlock(transitionBlockNumber, l2Verif, 5*time.Minute)
	require.Nil(t, err)

	// Ensure that the transition block inserted into chain.
	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	transitionBlock, err := l2Verif.BlockByNumber(ctx, transitionBlockNumber)
	require.Nil(t, err)
	require.Equal(t, []byte("BEDROCK"), transitionBlock.Extra())

	// Ensure that the transition block has been finalized.
	l2Finalized, err := geth.WaitForBlockToBeFinalized(transitionBlockNumber, l2Verif, 1*time.Minute)
	require.NoError(t, err, "must be able to fetch a finalized L2 block")
	require.NotZerof(t, l2Finalized.NumberU64(), "must have finalized L2 block")

	validateL1BlockTxProof(t, l1Cl, l2Verif, transitionBlockNumber)

	// Check states after migration
	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	balance, err := l2Verif.BalanceAt(ctx, common.Address{0xff, 0xff}, nil)
	require.Nil(t, err)
	require.True(t, balance.Cmp(common.Big0) == 1)

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	code, err := l2Verif.CodeAt(ctx, deployReceipt.ContractAddress, nil)
	require.Nil(t, err)
	require.True(t, len(code) > 0)
}

func validateL1BlockTxProof(t *testing.T, l1Cl *ethclient.Client, l2Cl *ethclient.Client, number *big.Int) {
	l1BlockHashSlot := "0x2"
	l2GethCl := gethclient.New(l2Cl.Client())

	validateZktProof := func(hex string) {
		b := common.Hex2Bytes(strings.TrimPrefix(hex, "0x"))
		_, err := trie.DecodeSMTProof(b)
		require.Nil(t, err)
	}
	validateMptProof := func(hex string) {
		b := common.Hex2Bytes(strings.TrimPrefix(hex, "0x"))
		_, _, err := rlp.SplitList(b)
		require.Nil(t, err)
	}
	validateL1BlockHash := func(v *big.Int) {
		_, err := l1Cl.BlockByHash(context.Background(), common.BigToHash(v))
		require.Nil(t, err)
	}
	proof, err := l2GethCl.GetProof(context.Background(), predeploys.L1BlockAddr, []string{l1BlockHashSlot}, new(big.Int).Sub(number, common.Big1))
	require.Nil(t, err, "failed to validate state proof for pre-transition block")
	for _, accProof := range proof.AccountProof {
		validateZktProof(accProof)
	}
	for _, storageProof := range proof.StorageProof {
		for _, p := range storageProof.Proof {
			validateZktProof(p)
		}
		if storageProof.Key == l1BlockHashSlot {
			validateL1BlockHash(storageProof.Value)
		}
	}

	proof, err = l2GethCl.GetProof(context.Background(), predeploys.L1BlockAddr, []string{l1BlockHashSlot}, number)
	require.Nil(t, err, "failed to validate state proof for transition block")
	for _, accProof := range proof.AccountProof {
		validateMptProof(accProof)
	}
	for _, storageProof := range proof.StorageProof {
		for _, p := range storageProof.Proof {
			validateMptProof(p)
		}
		if storageProof.Key == l1BlockHashSlot {
			validateL1BlockHash(storageProof.Value)
		}
	}

	proof, err = l2GethCl.GetProof(context.Background(), predeploys.L1BlockAddr, []string{l1BlockHashSlot}, new(big.Int).Add(number, common.Big1))
	require.Nil(t, err)
	require.Nil(t, err, "failed to validate state proof for post-transition block")
	for _, accProof := range proof.AccountProof {
		validateMptProof(accProof)
	}
	for _, storageProof := range proof.StorageProof {
		for _, p := range storageProof.Proof {
			validateMptProof(p)
		}
		if storageProof.Key == l1BlockHashSlot {
			validateL1BlockHash(storageProof.Value)
		}
	}
}
