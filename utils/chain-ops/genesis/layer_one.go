package genesis

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/trie"

	"github.com/kroma-network/kroma/bindings/bindings"
	"github.com/kroma-network/kroma/bindings/predeploys"
	"github.com/kroma-network/kroma/utils/chain-ops/deployer"
	"github.com/kroma-network/kroma/utils/chain-ops/state"
)

var proxies = []string{
	"SystemConfigProxy",
	"L2OutputOracleProxy",
	"L1CrossDomainMessengerProxy",
	"L1StandardBridgeProxy",
	"KromaPortalProxy",
	"KromaMintableERC20FactoryProxy",
	"ColosseumProxy",
}

var portalMeteringSlot = common.Hash{31: 0x01}

var zeroHash common.Hash

func BuildL1DeveloperGenesis(config *DeployConfig) (*core.Genesis, error) {
	if config.L2OutputOracleStartingTimestamp != -1 {
		return nil, errors.New("l2oo starting timestamp must be -1")
	}

	if config.L1GenesisBlockTimestamp == 0 {
		return nil, errors.New("must specify l1 genesis block timestamp")
	}

	genesis, err := NewL1Genesis(config)
	if err != nil {
		return nil, err
	}

	backend := deployer.NewBackendWithGenesisTimestamp(uint64(config.L1GenesisBlockTimestamp), false)

	deployments, err := deployL1Contracts(config, backend)
	if err != nil {
		return nil, err
	}

	depsByName := make(map[string]deployer.Deployment)
	depsByAddr := make(map[common.Address]deployer.Deployment)
	for _, dep := range deployments {
		depsByName[dep.Name] = dep
		depsByAddr[dep.Address] = dep
	}

	opts, err := bind.NewKeyedTransactorWithChainID(deployer.TestKey, deployer.ChainID)
	if err != nil {
		return nil, err
	}
	sysCfgABI, err := bindings.SystemConfigMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	gasLimit := uint64(config.L2GenesisBlockGasLimit)
	if gasLimit == 0 {
		gasLimit = defaultL2GasLimit
	}
	data, err := sysCfgABI.Pack(
		"initialize",
		config.FinalSystemOwner,
		uint642Big(config.GasPriceOracleOverhead),
		uint642Big(config.GasPriceOracleScalar),
		config.BatchSenderAddress.Hash(),
		gasLimit,
		config.P2PProposerAddress,
	)
	if err != nil {
		return nil, fmt.Errorf("cannot abi encode initialize for SystemConfig: %w", err)
	}
	if _, err := upgradeProxy(
		backend,
		opts,
		depsByName["SystemConfigProxy"].Address,
		depsByName["SystemConfig"].Address,
		data,
	); err != nil {
		return nil, err
	}

	l2ooABI, err := bindings.L2OutputOracleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	data, err = l2ooABI.Pack(
		"initialize",
		big.NewInt(0),
		uint642Big(uint64(config.L1GenesisBlockTimestamp)),
	)
	if err != nil {
		return nil, fmt.Errorf("cannot abi encode initialize for L2OutputOracle: %w", err)
	}
	if _, err := upgradeProxy(
		backend,
		opts,
		depsByName["L2OutputOracleProxy"].Address,
		depsByName["L2OutputOracle"].Address,
		data,
	); err != nil {
		return nil, err
	}

	colosseumABI, err := bindings.ColosseumMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	data, err = colosseumABI.Pack("initialize", parseSegsLengthsConfig(config.ColosseumSegmentsLengths))
	if err != nil {
		return nil, err
	}
	if _, err := upgradeProxy(
		backend,
		opts,
		depsByName["ColosseumProxy"].Address,
		depsByName["Colosseum"].Address,
		data,
	); err != nil {
		return nil, err
	}

	portalABI, err := bindings.KromaPortalMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	data, err = portalABI.Pack("initialize", false)
	if err != nil {
		return nil, fmt.Errorf("cannot abi encode initialize for KromaPortal: %w", err)
	}
	if _, err := upgradeProxy(
		backend,
		opts,
		depsByName["KromaPortalProxy"].Address,
		depsByName["KromaPortal"].Address,
		data,
	); err != nil {
		return nil, err
	}
	l1XDMABI, err := bindings.L1CrossDomainMessengerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	data, err = l1XDMABI.Pack("initialize")
	if err != nil {
		return nil, fmt.Errorf("cannot abi encode initialize for L1CrossDomainMessenger: %w", err)
	}
	if _, err := upgradeProxy(
		backend,
		opts,
		depsByName["L1CrossDomainMessengerProxy"].Address,
		depsByName["L1CrossDomainMessenger"].Address,
		data,
	); err != nil {
		return nil, err
	}

	if _, err := upgradeProxy(
		backend,
		opts,
		depsByName["L1StandardBridgeProxy"].Address,
		depsByName["L1StandardBridge"].Address,
		nil,
	); err != nil {
		return nil, err
	}

	var lastUpgradeTx *types.Transaction
	if lastUpgradeTx, err = upgradeProxy(
		backend,
		opts,
		depsByName["KromaMintableERC20FactoryProxy"].Address,
		depsByName["KromaMintableERC20Factory"].Address,
		nil,
	); err != nil {
		return nil, err
	}

	// Commit all the upgrades at once, then wait for the last
	// transaction to be mined. The simulator performs async
	// processing, and as such we need to wait for the transaction
	// receipt to appear before considering the above transactions
	// committed to the chain.

	backend.Commit()
	if _, err := bind.WaitMined(context.Background(), backend, lastUpgradeTx); err != nil {
		return nil, err
	}

	memDB := state.NewMemoryStateDB(genesis)
	if err := SetL1Proxies(memDB, predeploys.DevProxyAdminAddr); err != nil {
		return nil, err
	}
	FundDevAccounts(memDB)
	SetPrecompileBalances(memDB)

	for name, proxyAddr := range predeploys.DevPredeploys {
		memDB.SetState(*proxyAddr, ImplementationSlot, depsByName[name].Address.Hash())

		// Special case for WETH since it was not designed to be behind a proxy
		if name == "WETH9" {
			name, _ := state.EncodeStringValue("Wrapped Ether", 0)
			symbol, _ := state.EncodeStringValue("WETH", 0)
			decimals, _ := state.EncodeUintValue(18, 0)
			memDB.SetState(*proxyAddr, common.Hash{}, name)
			memDB.SetState(*proxyAddr, common.Hash{31: 0x01}, symbol)
			memDB.SetState(*proxyAddr, common.Hash{31: 0x02}, decimals)
		}

		if name == "ZKMerkleTrie" {
			poseidon2Addr, _ := state.EncodeAddressValue(predeploys.DevPoseidon2Addr, 0)
			memDB.SetState(*proxyAddr, common.Hash{}, poseidon2Addr)
		}
	}

	stateDB, err := backend.Blockchain().State()
	if err != nil {
		return nil, err
	}

	for _, dep := range deployments {
		st, err := stateDB.StorageTrie(dep.Address)
		if err != nil {
			return nil, fmt.Errorf("failed to open storage trie of %s: %w", dep.Address, err)
		}
		if st == nil {
			return nil, fmt.Errorf("missing account %s in state, address: %s", dep.Name, dep.Address)
		}
		iter := trie.NewIterator(st.NodeIterator(nil))

		depAddr := dep.Address
		if strings.HasSuffix(dep.Name, "Proxy") {
			depAddr = *predeploys.DevPredeploys[strings.TrimSuffix(dep.Name, "Proxy")]
		}

		memDB.CreateAccount(depAddr)
		memDB.SetCode(depAddr, dep.Bytecode)

		for iter.Next() {
			_, data, _, err := rlp.Split(iter.Value)
			if err != nil {
				return nil, err
			}

			key := common.BytesToHash(st.GetKey(iter.Key))
			value := common.BytesToHash(data)

			if depAddr == predeploys.DevKromaPortalAddr && key == portalMeteringSlot {
				// We need to manually set the block number in the resource
				// metering storage slot to zero. Otherwise, deposits will
				// revert.
				copy(value[:24], zeroHash[:])
			}

			memDB.SetState(depAddr, key, value)
		}
	}
	return memDB.Genesis(), nil
}

func deployL1Contracts(config *DeployConfig, backend *backends.SimulatedBackend) ([]deployer.Deployment, error) {
	constructors := make([]deployer.Constructor, 0)
	for _, proxy := range proxies {
		constructors = append(constructors, deployer.Constructor{
			Name: proxy,
		})
	}
	gasLimit := uint64(config.L2GenesisBlockGasLimit)
	if gasLimit == 0 {
		gasLimit = defaultL2GasLimit
	}
	constructors = append(constructors, []deployer.Constructor{
		{
			Name: "SystemConfig",
			Args: []interface{}{
				config.FinalSystemOwner,
				uint642Big(config.GasPriceOracleOverhead),
				uint642Big(config.GasPriceOracleScalar),
				config.BatchSenderAddress.Hash(), // left-padded 32 bytes value, version is zero anyway
				gasLimit,
				config.P2PProposerAddress,
			},
		},
		{
			Name: "L2OutputOracle",
			Args: []interface{}{
				uint642Big(config.L2OutputOracleSubmissionInterval),
				uint642Big(config.L2BlockTime),
				big.NewInt(0),
				uint642Big(uint64(config.L1GenesisBlockTimestamp)),
				config.L2OutputOracleValidator,
				predeploys.DevColosseumAddr,
				uint642Big(config.FinalizationPeriodSeconds),
			},
		},
		{
			Name: "Poseidon2",
		},
		{
			Name: "ZKMerkleTrie",
		},
		{
			Name: "ZKVerifier",
		},
		{
			Name: "KromaPortal",
			Args: []interface{}{
				config.PortalGuardian,
				true, // _paused
			},
		},
		{
			Name: "Colosseum",
			Args: []interface{}{
				uint642Big(config.L2OutputOracleSubmissionInterval),
				uint642Big(config.ColosseumChallengeTimeout),
				parseSegsLengthsConfig(config.ColosseumSegmentsLengths),
			},
		},
		{
			Name: "L1CrossDomainMessenger",
		},
		{
			Name: "L1StandardBridge",
		},
		{
			Name: "L1ERC721Bridge",
		},
		{
			Name: "KromaMintableERC20Factory",
		},
		{
			Name: "ProxyAdmin",
			Args: []interface{}{
				common.Address{19: 0x01},
			},
		},
		{
			Name: "WETH9",
		},
	}...)
	return deployer.Deploy(backend, constructors, l1Deployer)
}

func l1Deployer(backend *backends.SimulatedBackend, opts *bind.TransactOpts, deployment deployer.Constructor) (*types.Transaction, error) {
	var tx *types.Transaction
	var err error

	switch deployment.Name {
	case "SystemConfig":
		_, tx, _, err = bindings.DeploySystemConfig(
			opts,
			backend,
			/* owner= */ deployment.Args[0].(common.Address),
			/* overhead= */ deployment.Args[1].(*big.Int),
			/* scalar= */ deployment.Args[2].(*big.Int),
			/* batcherHash= */ deployment.Args[3].(common.Hash),
			/* gasLimit= */ deployment.Args[4].(uint64),
			/* unsafeBlockSigner= */ deployment.Args[5].(common.Address),
		)
	case "L2OutputOracle":
		_, tx, _, err = bindings.DeployL2OutputOracle(
			opts,
			backend,
			/* submissionInterval= */ deployment.Args[0].(*big.Int),
			/* l2BlockTime= */ deployment.Args[1].(*big.Int),
			/* startingBlockNumber= */ deployment.Args[2].(*big.Int),
			/* startingTimestamp= */ deployment.Args[3].(*big.Int),
			/* validator= */ deployment.Args[4].(common.Address),
			/* challenger= */ deployment.Args[5].(common.Address),
			/* finalizationPeriodSeconds= */ deployment.Args[6].(*big.Int),
		)
	case "KromaPortal":
		_, tx, _, err = bindings.DeployKromaPortal(
			opts,
			backend,
			predeploys.DevL2OutputOracleAddr,
			/* guardian= */ deployment.Args[0].(common.Address),
			/* paused= */ deployment.Args[1].(bool),
			predeploys.DevZKMerkleTrieAddr,
		)
	case "Colosseum":
		_, tx, _, err = bindings.DeployColosseum(
			opts,
			backend,
			predeploys.DevL2OutputOracleAddr,
			predeploys.DevZKVerifierAddr,
			/* submissionInterval= */ deployment.Args[0].(*big.Int),
			/* challengeTimeout= */ deployment.Args[1].(*big.Int),
			/* segmentsLengths= */ deployment.Args[2].([]*big.Int),
		)
	case "L1CrossDomainMessenger":
		_, tx, _, err = bindings.DeployL1CrossDomainMessenger(
			opts,
			backend,
			predeploys.DevKromaPortalAddr,
		)
	case "L1StandardBridge":
		_, tx, _, err = bindings.DeployL1StandardBridge(
			opts,
			backend,
			predeploys.DevL1CrossDomainMessengerAddr,
		)
	case "KromaMintableERC20Factory":
		_, tx, _, err = bindings.DeployKromaMintableERC20Factory(
			opts,
			backend,
			predeploys.DevL1StandardBridgeAddr,
		)
	case "ProxyAdmin":
		_, tx, _, err = bindings.DeployProxyAdmin(
			opts,
			backend,
			/* owner= */ common.Address{},
		)
	case "WETH9":
		_, tx, _, err = bindings.DeployWETH9(
			opts,
			backend,
		)
	case "L1ERC721Bridge":
		_, tx, _, err = bindings.DeployL1ERC721Bridge(
			opts,
			backend,
			predeploys.DevL1CrossDomainMessengerAddr,
			predeploys.L2ERC721BridgeAddr,
		)
	case "Poseidon2":
		_, tx, _, err = bind.DeployContract(
			opts,
			abi.ABI{},
			PoseidonByteCode,
			backend,
		)
	case "ZKMerkleTrie":
		_, tx, _, err = bindings.DeployZKMerkleTrie(
			opts,
			backend,
			predeploys.DevPoseidon2Addr,
		)
	case "ZKVerifier":
		_, tx, _, err = bindings.DeployZKVerifier(
			opts,
			backend,
		)
	default:
		if strings.HasSuffix(deployment.Name, "Proxy") {
			_, tx, _, err = bindings.DeployProxy(opts, backend, deployer.TestAddress)
		} else {
			err = fmt.Errorf("unknown contract %s", deployment.Name)
		}
	}

	if err != nil {
		err = fmt.Errorf("cannot deploy %s: %w", deployment.Name, err)
	}

	return tx, err
}

func upgradeProxy(backend *backends.SimulatedBackend, opts *bind.TransactOpts, proxyAddr common.Address, implAddr common.Address, callData []byte) (*types.Transaction, error) {
	var tx *types.Transaction
	proxy, err := bindings.NewProxy(proxyAddr, backend)
	if err != nil {
		return nil, err
	}
	if callData == nil {
		tx, err = proxy.UpgradeTo(opts, implAddr)
	} else {
		tx, err = proxy.UpgradeToAndCall(
			opts,
			implAddr,
			callData,
		)
	}
	return tx, err
}
