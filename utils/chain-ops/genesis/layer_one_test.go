package genesis

import (
	"bytes"
	"encoding/json"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	"github.com/stretchr/testify/require"

	"github.com/kroma-network/kroma/bindings/bindings"
	"github.com/kroma-network/kroma/bindings/predeploys"
	"github.com/kroma-network/kroma/utils/chain-ops/deployer"
)

func TestBuildL1DeveloperGenesis(t *testing.T) {
	b, err := os.ReadFile("testdata/test-deploy-config-full.json")
	require.NoError(t, err)
	dec := json.NewDecoder(bytes.NewReader(b))
	config := new(DeployConfig)
	require.NoError(t, dec.Decode(config))
	config.L1GenesisBlockTimestamp = hexutil.Uint64(time.Now().Unix() - 100)

	genesis, err := BuildL1DeveloperGenesis(config)
	require.NoError(t, err)

	sim := backends.NewSimulatedBackend(
		genesis.Alloc,
		15000000,
	)
	callOpts := &bind.CallOpts{}

	oracle, err := bindings.NewL2OutputOracle(predeploys.DevL2OutputOracleAddr, sim)
	require.NoError(t, err)
	portal, err := bindings.NewKromaPortal(predeploys.DevKromaPortalAddr, sim)
	require.NoError(t, err)

	validator, err := oracle.VALIDATOR(callOpts)
	require.NoError(t, err)
	require.Equal(t, config.L2OutputOracleValidator, validator)

	// Same set of tests as exist in the deployment scripts
	interval, err := oracle.SUBMISSIONINTERVAL(callOpts)
	require.NoError(t, err)
	require.EqualValues(t, config.L2OutputOracleSubmissionInterval, interval.Uint64())

	startBlock, err := oracle.StartingBlockNumber(callOpts)
	require.NoError(t, err)
	require.EqualValues(t, 0, startBlock.Uint64())

	l2BlockTime, err := oracle.L2BLOCKTIME(callOpts)
	require.NoError(t, err)
	require.EqualValues(t, 2, l2BlockTime.Uint64())

	oracleAddr, err := portal.L2ORACLE(callOpts)
	require.NoError(t, err)
	require.EqualValues(t, predeploys.DevL2OutputOracleAddr, oracleAddr)

	msgr, err := bindings.NewL1CrossDomainMessenger(predeploys.DevL1CrossDomainMessengerAddr, sim)
	require.NoError(t, err)
	portalAddr, err := msgr.PORTAL(callOpts)
	require.NoError(t, err)
	require.Equal(t, predeploys.DevKromaPortalAddr, portalAddr)

	bridge, err := bindings.NewL1StandardBridge(predeploys.DevL1StandardBridgeAddr, sim)
	require.NoError(t, err)
	msgrAddr, err := bridge.MESSENGER(callOpts)
	require.NoError(t, err)
	require.Equal(t, predeploys.DevL1CrossDomainMessengerAddr, msgrAddr)
	otherBridge, err := bridge.OTHERBRIDGE(callOpts)
	require.NoError(t, err)
	require.Equal(t, predeploys.L2StandardBridgeAddr, otherBridge)

	factory, err := bindings.NewKromaMintableERC20(predeploys.DevKromaMintableERC20FactoryAddr, sim)
	require.NoError(t, err)
	bridgeAddr, err := factory.BRIDGE(callOpts)
	require.NoError(t, err)
	require.Equal(t, predeploys.DevL1StandardBridgeAddr, bridgeAddr)

	weth9, err := bindings.NewWETH9(predeploys.DevWETH9Addr, sim)
	require.NoError(t, err)
	decimals, err := weth9.Decimals(callOpts)
	require.NoError(t, err)
	require.Equal(t, uint8(18), decimals)
	symbol, err := weth9.Symbol(callOpts)
	require.NoError(t, err)
	require.Equal(t, "WETH", symbol)
	name, err := weth9.Name(callOpts)
	require.NoError(t, err)
	require.Equal(t, "Wrapped Ether", name)

	sysCfg, err := bindings.NewSystemConfig(predeploys.DevSystemConfigAddr, sim)
	require.NoError(t, err)
	cfg, err := sysCfg.ResourceConfig(&bind.CallOpts{})
	require.NoError(t, err)
	require.Equal(t, cfg, defaultResourceConfig)
	owner, err := sysCfg.Owner(&bind.CallOpts{})
	require.NoError(t, err)
	require.Equal(t, owner, config.FinalSystemOwner)
	overhead, err := sysCfg.Overhead(&bind.CallOpts{})
	require.NoError(t, err)
	require.Equal(t, overhead.Uint64(), config.GasPriceOracleOverhead)
	scalar, err := sysCfg.Scalar(&bind.CallOpts{})
	require.NoError(t, err)
	require.Equal(t, scalar.Uint64(), config.GasPriceOracleScalar)
	batcherHash, err := sysCfg.BatcherHash(&bind.CallOpts{})
	require.NoError(t, err)
	require.Equal(t, common.Hash(batcherHash), config.BatchSenderAddress.Hash())
	gasLimit, err := sysCfg.GasLimit(&bind.CallOpts{})
	require.NoError(t, err)
	require.Equal(t, gasLimit, uint64(config.L2GenesisBlockGasLimit))
	unsafeBlockSigner, err := sysCfg.UnsafeBlockSigner(&bind.CallOpts{})
	require.NoError(t, err)
	require.Equal(t, unsafeBlockSigner, config.P2PProposerAddress)

	// test that we can do deposits, etc.
	priv, err := crypto.HexToECDSA("ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")
	require.NoError(t, err)

	tOpts, err := bind.NewKeyedTransactorWithChainID(priv, deployer.ChainID)
	require.NoError(t, err)
	tOpts.Value = big.NewInt(0.001 * params.Ether)
	tOpts.GasLimit = 1_000_000
	_, err = bridge.BridgeETH(tOpts, 200000, nil)
	require.NoError(t, err)
}
