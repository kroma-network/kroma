package actions

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/stretchr/testify/require"

	oppredeploys "github.com/ethereum-optimism/optimism/op-bindings/predeploys"
	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils"
	"github.com/ethereum-optimism/optimism/op-node/rollup/derive"
	"github.com/ethereum-optimism/optimism/op-service/testlog"
	"github.com/kroma-network/kroma/kroma-bindings/bindings"
	"github.com/kroma-network/kroma/kroma-bindings/predeploys"
	"github.com/kroma-network/kroma/kroma-chain-ops/genesis"
)

var (
	l1BlockMPTCodeHash           = common.HexToHash("0xc88a313aa75dc4fbf0b6850d9f9ae41e04243b7008cf3eadb29256d4a71c1dfd")
	baseFeeVaultMPTCodeHash      = common.HexToHash("0xc0ccbda46b89a834c65c871fb0ccb93f02c60268d5560a8e1b13722979bb38dd")
	sequencerFeeVaultMPTCodeHash = common.HexToHash("0xc0ccbda46b89a834c65c871fb0ccb93f02c60268d5560a8e1b13722979bb38dd")
	l1FeeVaultMPTCodeHash        = common.HexToHash("0xf96923a0115890fc344bc26402cbdc062ff2d6f63317b40dd7f284f2b37c6d64")
	gasPriceOracleMPTCodeHash    = common.HexToHash("0x3474b859129195c11a4b1aa14af951a2b3f83c7841aa5e73e67d55bd3379220c")
)

func TestKromaMPTNetworkUpgradeTransactions(gt *testing.T) {
	t := NewDefaultTesting(gt)
	dp := e2eutils.MakeDeployParams(t, defaultRollupTestParams)
	genesisBlock := hexutil.Uint64(0)
	ecotoneOffset := hexutil.Uint64(2)
	kromaMPTOffset := hexutil.Uint64(4)

	dp.DeployConfig.L1CancunTimeOffset = &genesisBlock // can be removed once Cancun on L1 is the default

	// Activate all forks at genesis, and schedule KromaMPT the block after
	dp.DeployConfig.L2GenesisRegolithTimeOffset = &genesisBlock
	dp.DeployConfig.L2GenesisCanyonTimeOffset = &genesisBlock
	dp.DeployConfig.L2GenesisDeltaTimeOffset = &genesisBlock
	dp.DeployConfig.L2GenesisEcotoneTimeOffset = &ecotoneOffset
	dp.DeployConfig.L2GenesisKromaMPTTimeOffset = &kromaMPTOffset
	require.NoError(t, dp.DeployConfig.Check(), "must have valid config")

	sd := e2eutils.Setup(t, dp, defaultAlloc)
	log := testlog.Logger(t, log.LevelDebug)
	_, _, _, sequencer, engine, verifier, _, _ := setupReorgTestActors(t, dp, sd, log)
	ethCl := engine.EthClient()

	// start op-nodes
	sequencer.ActL2PipelineFull(t)
	verifier.ActL2PipelineFull(t)

	// Get gas price from oracle
	gasPriceOracle, err := bindings.NewGasPriceOracleCaller(predeploys.GasPriceOracleAddr, ethCl)
	require.NoError(t, err)

	// Get fee vault recipients
	protocolFeeVault, err := bindings.NewProtocolVaultCaller(predeploys.ProtocolVaultAddr, ethCl)
	require.NoError(t, err)
	protocolFeeRecipient, err := protocolFeeVault.RECIPIENT(nil)
	require.NoError(t, err)
	require.NotEqual(t, protocolFeeRecipient, common.Address{})
	require.Equal(t, protocolFeeRecipient, dp.DeployConfig.ProtocolVaultRecipient)

	l1FeeVault, err := bindings.NewL1FeeVaultCaller(predeploys.L1FeeVaultAddr, ethCl)
	require.NoError(t, err)
	l1FeeRecipient, err := l1FeeVault.RECIPIENT(nil)
	require.NoError(t, err)
	require.NotEqual(t, l1FeeRecipient, common.Address{})
	require.Equal(t, l1FeeRecipient, dp.DeployConfig.L1FeeVaultRecipient)

	// Get current implementations addresses (by slot) for GasPriceOracle
	initialGasPriceOracleAddress, err := ethCl.StorageAt(context.Background(), predeploys.GasPriceOracleAddr, genesis.ImplementationSlot, nil)
	require.NoError(t, err)

	// Build to the pre-KromaMPT block
	sequencer.ActBuildL2ToPreKromaMPT(t)

	// get latest block
	latestBlock, err := ethCl.BlockByNumber(context.Background(), nil)
	require.NoError(t, err)
	require.Equal(t, sequencer.L2Unsafe().Number, latestBlock.Number().Uint64())

	transactions := latestBlock.Transactions()
	// L1Block: 1 set-L1-info + 5 deploys + 5 upgradeTo + 1 enable KromaMPT on GPO
	// See [derive.KromaMPTNetworkUpgradeTransactions]
	require.Equal(t, 12, len(transactions))

	_, err = derive.L1BlockInfoFromBytes(sd.RollupCfg, latestBlock.Time(), transactions[0].Data())
	require.NoError(t, err)
	require.Equal(t, derive.L1InfoEcotoneLen, len(transactions[0].Data()))

	// All transactions are successful
	for i := 1; i < 12; i++ {
		txn := transactions[i]
		receipt, err := ethCl.TransactionReceipt(context.Background(), txn.Hash())
		require.NoError(t, err)
		require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)
		require.NotEmpty(t, txn.Data(), "upgrade tx must provide input data")
	}

	expectedL1BlockAddress := crypto.CreateAddress(derive.L1BlockMPTDeployerAddress, 0)
	expectedBaseFeeVaultAddress := crypto.CreateAddress(derive.BaseFeeVaultDeployerAddress, 0)
	expectedL1FeeVaultAddress := crypto.CreateAddress(derive.L1FeeVaultDeployerAddress, 0)
	expectedSequencerFeeVaultAddress := crypto.CreateAddress(derive.SequencerFeeVaultDeployerAddress, 0)
	expectedGasPriceOracleAddress := crypto.CreateAddress(derive.GasPriceOracleMPTDeployerAddress, 0)

	// New Base Fee Vault impl is deployed
	updatedBaseFeeVaultAddress, err := ethCl.StorageAt(context.Background(), oppredeploys.BaseFeeVaultAddr, genesis.ImplementationSlot, latestBlock.Number())
	require.NoError(t, err)
	require.Equal(t, expectedBaseFeeVaultAddress, common.BytesToAddress(updatedBaseFeeVaultAddress))
	verifyCodeHashMatches(t, ethCl, expectedBaseFeeVaultAddress, baseFeeVaultMPTCodeHash)

	// New L1 Fee Vault impl is deployed
	updatedL1FeeVaultAddress, err := ethCl.StorageAt(context.Background(), oppredeploys.L1FeeVaultAddr, genesis.ImplementationSlot, latestBlock.Number())
	require.NoError(t, err)
	require.Equal(t, expectedL1FeeVaultAddress, common.BytesToAddress(updatedL1FeeVaultAddress))
	verifyCodeHashMatches(t, ethCl, expectedL1FeeVaultAddress, l1FeeVaultMPTCodeHash)

	// New Sequencer Fee Vault impl is deployed
	updatedSequencerFeeVaultAddress, err := ethCl.StorageAt(context.Background(), oppredeploys.SequencerFeeVaultAddr, genesis.ImplementationSlot, latestBlock.Number())
	require.NoError(t, err)
	require.Equal(t, expectedSequencerFeeVaultAddress, common.BytesToAddress(updatedSequencerFeeVaultAddress))
	verifyCodeHashMatches(t, ethCl, expectedSequencerFeeVaultAddress, sequencerFeeVaultMPTCodeHash)

	// Gas Price Oracle impl is updated
	updatedGasPriceOracleAddress, err := ethCl.StorageAt(context.Background(), predeploys.GasPriceOracleAddr, genesis.ImplementationSlot, latestBlock.Number())
	require.NoError(t, err)
	require.Equal(t, expectedGasPriceOracleAddress, common.BytesToAddress(updatedGasPriceOracleAddress))
	require.NotEqualf(t, initialGasPriceOracleAddress, updatedGasPriceOracleAddress, "Gas Price Oracle Proxy address should have changed")
	verifyCodeHashMatches(t, ethCl, expectedGasPriceOracleAddress, gasPriceOracleMPTCodeHash)

	// L1Block impl is deployed
	updatedL1BlockAddress, err := ethCl.StorageAt(context.Background(), oppredeploys.L1BlockAddr, genesis.ImplementationSlot, latestBlock.Number())
	require.NoError(t, err)
	require.Equal(t, expectedL1BlockAddress, common.BytesToAddress(updatedL1BlockAddress))
	verifyCodeHashMatches(t, ethCl, common.BytesToAddress(updatedL1BlockAddress), l1BlockMPTCodeHash)

	// Check that KromaMPT was activated
	isKromaMPT, err := gasPriceOracle.IsKromaMPT(nil)
	require.NoError(t, err)
	require.True(t, isKromaMPT)

	// Check if recipients are the same as before in the newly deployed in the updated fee vaults
	baseFeeVault, err := bindings.NewProtocolVaultCaller(oppredeploys.BaseFeeVaultAddr, ethCl)
	require.NoError(t, err)
	updatedBaseFeeRecipient, err := baseFeeVault.RECIPIENT(nil)
	require.NoError(t, err)
	require.Equal(t, protocolFeeRecipient, updatedBaseFeeRecipient)

	l1FeeVault, err = bindings.NewL1FeeVaultCaller(oppredeploys.L1FeeVaultAddr, ethCl)
	require.NoError(t, err)
	updatedL1FeeRecipient, err := l1FeeVault.RECIPIENT(nil)
	require.NoError(t, err)
	require.Equal(t, l1FeeRecipient, updatedL1FeeRecipient)

	sequencerFeeVault, err := bindings.NewProtocolVaultCaller(oppredeploys.SequencerFeeVaultAddr, ethCl)
	require.NoError(t, err)
	updatedSequencerFeeRecipient, err := sequencerFeeVault.RECIPIENT(nil)
	require.NoError(t, err)
	require.Equal(t, protocolFeeRecipient, updatedSequencerFeeRecipient)

	l1Block, err := bindings.NewL1BlockCaller(oppredeploys.L1BlockAddr, ethCl)
	require.NoError(t, err)

	// Get L1Block info, now set as 0
	initialL1BlockTimestamp, err := l1Block.Timestamp(nil)
	require.NoError(t, err)
	require.Equal(t, uint64(0), initialL1BlockTimestamp)
}
