package e2e

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"os"
	"path"
	"sort"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core"
	geth_eth "github.com/ethereum/go-ethereum/eth"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/rpc"
	mocknet "github.com/libp2p/go-libp2p/p2p/net/mock"
	"github.com/stretchr/testify/require"

	"github.com/kroma-network/kroma/bindings/predeploys"
	batcher "github.com/kroma-network/kroma/components/batcher"
	batchermetrics "github.com/kroma-network/kroma/components/batcher/metrics"
	"github.com/kroma-network/kroma/components/node/chaincfg"
	"github.com/kroma-network/kroma/components/node/eth"
	"github.com/kroma-network/kroma/components/node/metrics"
	rollupNode "github.com/kroma-network/kroma/components/node/node"
	"github.com/kroma-network/kroma/components/node/p2p"
	"github.com/kroma-network/kroma/components/node/rollup"
	"github.com/kroma-network/kroma/components/node/rollup/driver"
	"github.com/kroma-network/kroma/components/node/sources"
	"github.com/kroma-network/kroma/components/node/testlog"
	validator "github.com/kroma-network/kroma/components/validator"
	validatormetrics "github.com/kroma-network/kroma/components/validator/metrics"
	"github.com/kroma-network/kroma/e2e/e2eutils"
	"github.com/kroma-network/kroma/utils/chain-ops/genesis"
	klog "github.com/kroma-network/kroma/utils/service/log"
	"github.com/kroma-network/kroma/utils/service/txmgr"
)

var testingJWTSecret = [32]byte{123}

func newTxMgrConfig(l1Addr string, privKey *ecdsa.PrivateKey) txmgr.CLIConfig {
	return txmgr.CLIConfig{
		L1RPCURL:                  l1Addr,
		PrivateKey:                hexPriv(privKey),
		NumConfirmations:          1,
		SafeAbortNonceTooLowCount: 3,
		ResubmissionTimeout:       3 * time.Second,
		ReceiptQueryInterval:      50 * time.Millisecond,
		NetworkTimeout:            2 * time.Second,
		TxSendTimeout:             10 * time.Minute,
		TxNotInMempoolTimeout:     2 * time.Minute,
	}
}

func DefaultSystemConfig(t *testing.T) SystemConfig {
	secrets, err := e2eutils.DefaultMnemonicConfig.Secrets()
	require.NoError(t, err)
	addresses := secrets.Addresses()

	deployConfig := &genesis.DeployConfig{
		L1ChainID:   900,
		L2ChainID:   901,
		L2BlockTime: 1,

		FinalizationPeriodSeconds: 60 * 60 * 24,
		MaxProposerDrift:          10,
		ProposerWindowSize:        30,
		ChannelTimeout:            10,
		P2PProposerAddress:        addresses.ProposerP2P,
		BatchInboxAddress:         common.Address{0: 0x52, 19: 0xff}, // tbd
		BatchSenderAddress:        addresses.Batcher,

		DummyHash: common.HexToHash("0x8e556cf0e9ed5d6b6ad79247cddc30112cfee4a207fb13903eb834b447aebae9"),
		MaxTxs:    25,

		L2OutputOracleSubmissionInterval: 4,
		L2OutputOracleStartingTimestamp:  -1,
		L2OutputOracleValidator:          addresses.Validator,

		FinalSystemOwner: addresses.SysCfgOwner,

		L1BlockTime:                 2,
		L1GenesisBlockNonce:         4660,
		CliqueSignerAddress:         common.Address{}, // e2e used to run Clique, but now uses fake Proof of Stake.
		L1GenesisBlockTimestamp:     hexutil.Uint64(time.Now().Unix()),
		L1GenesisBlockGasLimit:      30_000_000,
		L1GenesisBlockDifficulty:    uint642big(1),
		L1GenesisBlockMixHash:       common.Hash{},
		L1GenesisBlockCoinbase:      common.Address{},
		L1GenesisBlockNumber:        0,
		L1GenesisBlockGasUsed:       0,
		L1GenesisBlockParentHash:    common.Hash{},
		L1GenesisBlockBaseFeePerGas: uint642big(7),

		L2GenesisBlockNonce:         0,
		L2GenesisBlockGasLimit:      30_000_000,
		L2GenesisBlockDifficulty:    uint642big(1),
		L2GenesisBlockMixHash:       common.Hash{},
		L2GenesisBlockNumber:        0,
		L2GenesisBlockGasUsed:       0,
		L2GenesisBlockParentHash:    common.Hash{},
		L2GenesisBlockBaseFeePerGas: uint642big(7),

		L2GenesisBlueTimeOffset: new(hexutil.Uint64),

		ColosseumChallengeTimeout: 120,
		ColosseumSegmentsLengths:  "3,3",

		GasPriceOracleOverhead: 2100,
		GasPriceOracleScalar:   1_000_000,

		ValidatorRewardVaultRecipient: common.Address{19: 1},
		ProtocolVaultRecipient:        common.Address{19: 2},
		ProposerRewardVaultRecipient:  common.Address{19: 3},

		DeploymentWaitConfirmations: 1,

		EIP1559Elasticity:  2,
		EIP1559Denominator: 8,

		FundDevAccounts: true,
	}

	if err := deployConfig.InitDeveloperDeployedAddresses(); err != nil {
		panic(err)
	}

	return SystemConfig{
		Secrets: secrets,

		Premine: make(map[common.Address]*big.Int),

		DeployConfig:           deployConfig,
		L1InfoPredeployAddress: predeploys.L1BlockAddr,
		JWTFilePath:            writeDefaultJWT(t),
		JWTSecret:              testingJWTSecret,
		Nodes: map[string]*rollupNode.Config{
			"syncer": {
				Driver: driver.Config{
					SyncerConfDepth:   0,
					ProposerConfDepth: 0,
					ProposerEnabled:   false,
				},
				L1EpochPollInterval: time.Second * 4,
			},
			"proposer": {
				Driver: driver.Config{
					SyncerConfDepth:   0,
					ProposerConfDepth: 0,
					ProposerEnabled:   true,
				},
				// Submitter PrivKey is set in system start for rollup nodes where proposer = true
				RPC: rollupNode.RPCConfig{
					ListenAddr:  "127.0.0.1",
					ListenPort:  0,
					EnableAdmin: true,
				},
				L1EpochPollInterval: time.Second * 4,
			},
		},
		Loggers: map[string]log.Logger{
			"syncer":    testlog.Logger(t, log.LvlInfo).New("role", "syncer"),
			"proposer":  testlog.Logger(t, log.LvlInfo).New("role", "proposer"),
			"batcher":   testlog.Logger(t, log.LvlInfo).New("role", "batcher"),
			"validator": testlog.Logger(t, log.LvlCrit).New("role", "validator"),
		},
		GethOptions:         map[string][]GethOption{},
		P2PTopology:         nil, // no P2P connectivity by default
		NonFinalizedOutputs: false,
	}
}

func writeDefaultJWT(t *testing.T) string {
	// Sadly the geth node config cannot load JWT secret from memory, it has to be a file
	jwtPath := path.Join(t.TempDir(), "jwt_secret")
	if err := os.WriteFile(jwtPath, []byte(hexutil.Encode(testingJWTSecret[:])), 0o600); err != nil {
		t.Fatalf("failed to prepare jwt file for geth: %v", err)
	}
	return jwtPath
}

type DepositContractConfig struct {
	L2Oracle           common.Address
	FinalizationPeriod *big.Int
}

type SystemConfig struct {
	Secrets                *e2eutils.Secrets
	L1InfoPredeployAddress common.Address

	DeployConfig *genesis.DeployConfig

	JWTFilePath string
	JWTSecret   [32]byte

	Premine         map[common.Address]*big.Int
	Nodes           map[string]*rollupNode.Config // Per node config. Don't use populate rollup.Config
	Loggers         map[string]log.Logger
	GethOptions     map[string][]GethOption
	ValidatorLogger log.Logger
	BatcherLogger   log.Logger

	// map of outbound connections to other nodes. Node names prefixed with "~" are unconnected but linked.
	// A nil map disables P2P completely.
	// Any node name not in the topology will not have p2p enabled.
	P2PTopology map[string][]string

	// Enables req-resp sync in the P2P nodes
	P2PReqRespSync bool

	// If the validator can make outputs for L2 blocks derived from L1 blocks which are not finalized on L1 yet.
	NonFinalizedOutputs bool

	// Explicitly disable batcher, for tests that rely on unsafe L2 payloads
	DisableBatcher bool
}

type System struct {
	cfg SystemConfig

	RollupConfig *rollup.Config

	L2GenesisCfg *core.Genesis

	// Connections to running nodes
	Nodes       map[string]*node.Node
	Backends    map[string]*geth_eth.Ethereum
	Clients     map[string]*ethclient.Client
	RollupNodes map[string]*rollupNode.KromaNode
	Validator   *validator.Validator
	Batcher     *batcher.Batcher
	Mocknet     mocknet.Mocknet
}

func (sys *System) Close() {
	if sys.Validator != nil {
		sys.Validator.Stop()
	}
	if sys.Batcher != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		sys.Batcher.Stop(ctx)
	}

	for _, node := range sys.RollupNodes {
		node.Close()
	}
	for _, node := range sys.Nodes {
		node.Close()
	}
	sys.Mocknet.Close()
}

type systemConfigHook func(sCfg *SystemConfig, s *System)

type SystemConfigOption struct {
	key    string
	role   string
	action systemConfigHook
}

type SystemConfigOptions struct {
	opts map[string]systemConfigHook
}

func NewSystemConfigOptions(_opts []SystemConfigOption) (SystemConfigOptions, error) {
	opts := make(map[string]systemConfigHook)
	for _, opt := range _opts {
		if _, ok := opts[opt.key+":"+opt.role]; ok {
			return SystemConfigOptions{}, fmt.Errorf("duplicate option for key %s and role %s", opt.key, opt.role)
		}
		opts[opt.key+":"+opt.role] = opt.action
	}

	return SystemConfigOptions{
		opts: opts,
	}, nil
}

func (s *SystemConfigOptions) Get(key, role string) (systemConfigHook, bool) {
	v, ok := s.opts[key+":"+role]
	return v, ok
}

func (cfg SystemConfig) Start(_opts ...SystemConfigOption) (*System, error) {
	opts, err := NewSystemConfigOptions(_opts)
	if err != nil {
		return nil, err
	}

	sys := &System{
		cfg:         cfg,
		Nodes:       make(map[string]*node.Node),
		Backends:    make(map[string]*geth_eth.Ethereum),
		Clients:     make(map[string]*ethclient.Client),
		RollupNodes: make(map[string]*rollupNode.KromaNode),
	}
	didErrAfterStart := false
	defer func() {
		if didErrAfterStart {
			for _, node := range sys.RollupNodes {
				node.Close()
			}
			for _, node := range sys.Nodes {
				node.Close()
			}
		}
	}()

	l1Genesis, err := genesis.BuildL1DeveloperGenesis(cfg.DeployConfig)
	if err != nil {
		return nil, err
	}

	for addr, amount := range cfg.Premine {
		if existing, ok := l1Genesis.Alloc[addr]; ok {
			l1Genesis.Alloc[addr] = core.GenesisAccount{
				Code:    existing.Code,
				Storage: existing.Storage,
				Balance: amount,
				Nonce:   existing.Nonce,
			}
		} else {
			l1Genesis.Alloc[addr] = core.GenesisAccount{
				Balance: amount,
				Nonce:   0,
			}
		}
	}

	l1Block := l1Genesis.ToBlock()
	l2Genesis, err := genesis.BuildL2DeveloperGenesis(cfg.DeployConfig, l1Block, true)
	if err != nil {
		return nil, err
	}
	sys.L2GenesisCfg = l2Genesis
	for addr, amount := range cfg.Premine {
		if existing, ok := l2Genesis.Alloc[addr]; ok {
			l2Genesis.Alloc[addr] = core.GenesisAccount{
				Code:    existing.Code,
				Storage: existing.Storage,
				Balance: amount,
				Nonce:   existing.Nonce,
			}
		} else {
			l2Genesis.Alloc[addr] = core.GenesisAccount{
				Balance: amount,
				Nonce:   0,
			}
		}
	}

	makeRollupConfig := func() rollup.Config {
		return rollup.Config{
			Genesis: rollup.Genesis{
				L1: eth.BlockID{
					Hash:   l1Block.Hash(),
					Number: 0,
				},
				L2: eth.BlockID{
					Hash:   l2Genesis.ToBlock().Hash(),
					Number: 0,
				},
				L2Time:       uint64(cfg.DeployConfig.L1GenesisBlockTimestamp),
				SystemConfig: e2eutils.SystemConfigFromDeployConfig(cfg.DeployConfig),
			},
			BlockTime:              cfg.DeployConfig.L2BlockTime,
			MaxProposerDrift:       cfg.DeployConfig.MaxProposerDrift,
			ProposerWindowSize:     cfg.DeployConfig.ProposerWindowSize,
			ChannelTimeout:         cfg.DeployConfig.ChannelTimeout,
			L1ChainID:              cfg.L1ChainIDBig(),
			L2ChainID:              cfg.L2ChainIDBig(),
			BatchInboxAddress:      cfg.DeployConfig.BatchInboxAddress,
			DepositContractAddress: predeploys.DevKromaPortalAddr,
			L1SystemConfigAddress:  predeploys.DevSystemConfigAddr,
			BlueTime:               cfg.DeployConfig.BlueTime(uint64(cfg.DeployConfig.L1GenesisBlockTimestamp)),
		}
	}
	defaultConfig := makeRollupConfig()
	sys.RollupConfig = &defaultConfig

	// Initialize nodes
	l1Node, l1Backend, err := initL1Geth(&cfg, l1Genesis, cfg.GethOptions["l1"]...)
	if err != nil {
		return nil, err
	}
	sys.Nodes["l1"] = l1Node
	sys.Backends["l1"] = l1Backend

	for name := range cfg.Nodes {
		node, backend, err := initL2Geth(name, big.NewInt(int64(cfg.DeployConfig.L2ChainID)), l2Genesis, cfg.JWTFilePath, cfg.GethOptions[name]...)
		if err != nil {
			return nil, err
		}
		sys.Nodes[name] = node
		sys.Backends[name] = backend
	}

	// Start
	err = l1Node.Start()
	if err != nil {
		didErrAfterStart = true
		return nil, err
	}
	for name, node := range sys.Nodes {
		if name == "l1" {
			continue
		}
		err = node.Start()
		if err != nil {
			didErrAfterStart = true
			return nil, err
		}
	}

	// Configure connections to L1 and L2 for rollup nodes.
	// TODO: refactor testing to use in-process rpc connections instead of websockets.

	for name, rollupCfg := range cfg.Nodes {
		configureL1(rollupCfg, l1Node)
		configureL2(rollupCfg, sys.Nodes[name], cfg.JWTSecret)

		rollupCfg.L2Sync = &rollupNode.PreparedL2SyncEndpoint{
			Client:   nil,
			TrustRPC: false,
		}
	}

	// Geth Clients
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	l1Srv, err := l1Node.RPCHandler()
	if err != nil {
		didErrAfterStart = true
		return nil, err
	}
	l1Client := ethclient.NewClient(rpc.DialInProc(l1Srv))
	sys.Clients["l1"] = l1Client
	for name, node := range sys.Nodes {
		client, err := ethclient.DialContext(ctx, node.WSEndpoint())
		if err != nil {
			didErrAfterStart = true
			return nil, err
		}
		sys.Clients[name] = client
	}

	_, err = waitForBlock(big.NewInt(2), l1Client, 6*time.Second*time.Duration(cfg.DeployConfig.L1BlockTime))
	if err != nil {
		return nil, fmt.Errorf("waiting for blocks: %w", err)
	}

	sys.Mocknet = mocknet.New()

	p2pNodes := make(map[string]*p2p.Prepared)
	if cfg.P2PTopology != nil {
		// create the peer if it doesn't exist yet.
		initHostMaybe := func(name string) (*p2p.Prepared, error) {
			if p, ok := p2pNodes[name]; ok {
				return p, nil
			}
			h, err := sys.Mocknet.GenPeer()
			if err != nil {
				return nil, fmt.Errorf("failed to init p2p host for node %s", name)
			}
			h.Network()
			_, ok := cfg.Nodes[name]
			if !ok {
				return nil, fmt.Errorf("node %s from p2p topology not found in actual nodes map", name)
			}
			// TODO we can enable discv5 in the testnodes to test discovery of new peers.
			// Would need to mock though, and the discv5 implementation does not provide nice mocks here.
			p := &p2p.Prepared{
				HostP2P:           h,
				LocalNode:         nil,
				UDPv5:             nil,
				EnableReqRespSync: cfg.P2PReqRespSync,
			}
			p2pNodes[name] = p
			return p, nil
		}
		for k, vs := range cfg.P2PTopology {
			peerA, err := initHostMaybe(k)
			if err != nil {
				return nil, fmt.Errorf("failed to setup mocknet peer %s", k)
			}
			for _, v := range vs {
				v = strings.TrimPrefix(v, "~")
				peerB, err := initHostMaybe(v)
				if err != nil {
					return nil, fmt.Errorf("failed to setup mocknet peer %s (peer of %s)", v, k)
				}
				if _, err := sys.Mocknet.LinkPeers(peerA.HostP2P.ID(), peerB.HostP2P.ID()); err != nil {
					return nil, fmt.Errorf("failed to setup mocknet link between %s and %s", k, v)
				}
				// connect the peers after starting the full rollup node
			}
		}
	}

	// Don't log state snapshots in test output
	snapLog := log.New()
	snapLog.SetHandler(log.DiscardHandler())

	// Rollup nodes

	// Ensure we are looping through the nodes in alphabetical order
	ks := make([]string, 0, len(cfg.Nodes))
	for k := range cfg.Nodes {
		ks = append(ks, k)
	}
	// Sort strings in ascending alphabetical order
	sort.Strings(ks)

	for _, name := range ks {
		nodeConfig := cfg.Nodes[name]
		c := *nodeConfig // copy
		c.Rollup = makeRollupConfig()

		if p, ok := p2pNodes[name]; ok {
			c.P2P = p

			if c.Driver.ProposerEnabled && c.P2PSigner == nil {
				c.P2PSigner = &p2p.PreparedSigner{Signer: p2p.NewLocalSigner(cfg.Secrets.ProposerP2P)}
			}
		}

		c.Rollup.LogDescription(cfg.Loggers[name], chaincfg.L2ChainIDToNetworkName)

		node, err := rollupNode.New(context.Background(), &c, cfg.Loggers[name], snapLog, "", metrics.NewMetrics(""))
		if err != nil {
			didErrAfterStart = true
			return nil, err
		}
		err = node.Start(context.Background())
		if err != nil {
			didErrAfterStart = true
			return nil, err
		}
		sys.RollupNodes[name] = node

		if action, ok := opts.Get("afterRollupNodeStart", name); ok {
			action(&cfg, sys)
		}
	}

	if cfg.P2PTopology != nil {
		// We only set up the connections after starting the actual nodes,
		// so GossipSub and other p2p protocols can be started before the connections go live.
		// This way protocol negotiation happens correctly.
		for k, vs := range cfg.P2PTopology {
			peerA := p2pNodes[k]
			for _, v := range vs {
				unconnected := strings.HasPrefix(v, "~")
				if unconnected {
					v = v[1:]
				}
				if !unconnected {
					peerB := p2pNodes[v]
					if _, err := sys.Mocknet.ConnectPeers(peerA.HostP2P.ID(), peerB.HostP2P.ID()); err != nil {
						return nil, fmt.Errorf("failed to setup mocknet connection between %s and %s", k, v)
					}
				}
			}
		}
	}

	// Validator (L2 Output Submitter + Challenger)
	validatorCliCfg := validator.CLIConfig{
		L1EthRpc:          sys.Nodes["l1"].WSEndpoint(),
		RollupRpc:         sys.RollupNodes["proposer"].HTTPEndpoint(),
		L2OOAddress:       predeploys.DevL2OutputOracleAddr.String(),
		ColosseumAddress:  predeploys.DevColosseumAddr.String(),
		PollInterval:      50 * time.Millisecond,
		AllowNonFinalized: cfg.NonFinalizedOutputs,
		TxMgrConfig:       newTxMgrConfig(sys.Nodes["l1"].WSEndpoint(), cfg.Secrets.Validator),
		LogConfig: klog.CLIConfig{
			Level:  "info",
			Format: "text",
		},
		ProverGrpc: "http://0.0.0.0:0",
	}

	validatorCfg, err := validator.NewValidatorConfig(validatorCliCfg, sys.cfg.Loggers["validator"], validatormetrics.NoopMetrics)
	if err != nil {
		return nil, fmt.Errorf("unable to init validator config: %w", err)
	}
	// replace to mock fetcher
	validatorCfg.ProofFetcher = e2eutils.NewFetcher(sys.cfg.Loggers["validator"])
	sys.Validator, err = validator.NewValidator(context.Background(), *validatorCfg, sys.cfg.Loggers["validator"], validatormetrics.NoopMetrics)
	if err != nil {
		return nil, fmt.Errorf("unable to setup validator: %w", err)
	}

	sys.Validator.Start()

	// Batcher (Batch Submitter)
	batcherCliCfg := batcher.CLIConfig{
		L1EthRpc:           sys.Nodes["l1"].WSEndpoint(),
		L2EthRpc:           sys.Nodes["proposer"].WSEndpoint(),
		RollupRpc:          sys.RollupNodes["proposer"].HTTPEndpoint(),
		MaxChannelDuration: 1,
		MaxL1TxSize:        120_000,
		TargetL1TxSize:     100_000,
		TargetNumFrames:    1,
		ApproxComprRatio:   0.4,
		SubSafetyMargin:    4,
		PollInterval:       50 * time.Millisecond,
		TxMgrConfig:        newTxMgrConfig(sys.Nodes["l1"].WSEndpoint(), cfg.Secrets.Batcher),
		LogConfig: klog.CLIConfig{
			Level:  "info",
			Format: "text",
		},
	}

	batcherCfg, err := batcher.NewBatcherConfig(batcherCliCfg, sys.cfg.Loggers["batcher"], batchermetrics.NoopMetrics)
	if err != nil {
		return nil, fmt.Errorf("unable to init batcher config: %w", err)
	}
	sys.Batcher, err = batcher.NewBatcher(context.Background(), *batcherCfg, sys.cfg.Loggers["batcher"], batchermetrics.NoopMetrics)
	if err != nil {
		return nil, fmt.Errorf("failed to setup batcher: %w", err)
	}

	// Batcher may be enabled later
	if !sys.cfg.DisableBatcher {
		sys.Batcher.Start()
	}

	return sys, nil
}

func configureL1(rollupNodeCfg *rollupNode.Config, l1Node *node.Node) {
	l1EndpointConfig := l1Node.WSEndpoint()
	useHTTP := os.Getenv("E2E_USE_HTTP") == "true"
	if useHTTP {
		log.Info("using HTTP client")
		l1EndpointConfig = l1Node.HTTPEndpoint()
	}
	rollupNodeCfg.L1 = &rollupNode.L1EndpointConfig{
		L1NodeAddr:       l1EndpointConfig,
		L1TrustRPC:       false,
		L1RPCKind:        sources.RPCKindBasic,
		RateLimit:        0,
		BatchSize:        20,
		HttpPollInterval: time.Millisecond * 100,
	}
}

func configureL2(rollupNodeCfg *rollupNode.Config, l2Node *node.Node, jwtSecret [32]byte) {
	useHTTP := os.Getenv("E2E_USE_HTTP") == "true"
	l2EndpointConfig := l2Node.WSAuthEndpoint()
	if useHTTP {
		l2EndpointConfig = l2Node.HTTPAuthEndpoint()
	}

	rollupNodeCfg.L2 = &rollupNode.L2EndpointConfig{
		L2EngineAddr:      l2EndpointConfig,
		L2EngineJWTSecret: jwtSecret,
	}
}

func (cfg SystemConfig) L1ChainIDBig() *big.Int {
	return new(big.Int).SetUint64(cfg.DeployConfig.L1ChainID)
}

func (cfg SystemConfig) L2ChainIDBig() *big.Int {
	return new(big.Int).SetUint64(cfg.DeployConfig.L2ChainID)
}

func uint642big(in uint64) *hexutil.Big {
	b := new(big.Int).SetUint64(in)
	hu := hexutil.Big(*b)
	return &hu
}

func hexPriv(in *ecdsa.PrivateKey) string {
	b := e2eutils.EncodePrivKey(in)
	return hexutil.Encode(b)
}
