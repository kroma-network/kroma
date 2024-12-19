package op_e2e

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"math/big"
	"net"
	"os"
	"path"
	"sort"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	bss "github.com/ethereum-optimism/optimism/op-batcher/batcher"
	batcherFlags "github.com/ethereum-optimism/optimism/op-batcher/flags"
	"github.com/ethereum-optimism/optimism/op-e2e/config"
	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils"
	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils/batcher"
	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils/fakebeacon"
	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils/geth"
	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils/wait"
	"github.com/ethereum-optimism/optimism/op-node/chaincfg"
	"github.com/ethereum-optimism/optimism/op-node/metrics"
	rollupNode "github.com/ethereum-optimism/optimism/op-node/node"
	"github.com/ethereum-optimism/optimism/op-node/p2p"
	"github.com/ethereum-optimism/optimism/op-node/p2p/store"
	"github.com/ethereum-optimism/optimism/op-node/rollup"
	"github.com/ethereum-optimism/optimism/op-node/rollup/derive"
	"github.com/ethereum-optimism/optimism/op-node/rollup/driver"
	"github.com/ethereum-optimism/optimism/op-node/rollup/sync"
	"github.com/ethereum-optimism/optimism/op-service/cliapp"
	"github.com/ethereum-optimism/optimism/op-service/clock"
	"github.com/ethereum-optimism/optimism/op-service/dial"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	oplog "github.com/ethereum-optimism/optimism/op-service/log"
	"github.com/ethereum-optimism/optimism/op-service/sources"
	"github.com/ethereum-optimism/optimism/op-service/testlog"
	"github.com/ethereum-optimism/optimism/op-service/txmgr"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core"
	geth_eth "github.com/ethereum/go-ethereum/eth"
	"github.com/ethereum/go-ethereum/eth/ethconfig"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
	ds "github.com/ipfs/go-datastore"
	dsSync "github.com/ipfs/go-datastore/sync"
	ic "github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/peerstore"
	"github.com/libp2p/go-libp2p/p2p/host/peerstore/pstoremem"
	mocknet "github.com/libp2p/go-libp2p/p2p/net/mock"
	ma "github.com/multiformats/go-multiaddr"
	"github.com/stretchr/testify/require"

	"github.com/kroma-network/kroma/kroma-bindings/predeploys"
	"github.com/kroma-network/kroma/kroma-chain-ops/genesis"
	validator "github.com/kroma-network/kroma/kroma-validator"
	chal "github.com/kroma-network/kroma/kroma-validator/challenge"
	validatormetrics "github.com/kroma-network/kroma/kroma-validator/metrics"
	valhelper "github.com/kroma-network/kroma/op-e2e/e2eutils/validator"
	"github.com/kroma-network/kroma/op-e2e/testdata"
	"github.com/kroma-network/kroma/op-service/client"
)

var testingJWTSecret = [32]byte{123}

func newTxMgrConfig(l1Addr string, privKey *ecdsa.PrivateKey) txmgr.CLIConfig {
	return txmgr.CLIConfig{
		L1RPCURL:                  l1Addr,
		PrivateKey:                hexPriv(privKey),
		NumConfirmations:          1,
		SafeAbortNonceTooLowCount: 3,
		FeeLimitMultiplier:        5,
		ResubmissionTimeout:       3 * time.Second,
		ReceiptQueryInterval:      50 * time.Millisecond,
		NetworkTimeout:            2 * time.Second,
		TxNotInMempoolTimeout:     2 * time.Minute,
		// [Kroma: START]
		TxSendTimeout: 10 * time.Minute,
		TxBufferSize:  10,
		// [Kroma: END]
	}
}

func DefaultSystemConfig(t *testing.T) SystemConfig {
	config.ExternalL2TestParms.SkipIfNecessary(t)

	secrets, err := e2eutils.DefaultMnemonicConfig.Secrets()
	require.NoError(t, err)
	deployConfig := config.DeployConfig.Copy()
	deployConfig.L1GenesisBlockTimestamp = hexutil.Uint64(time.Now().Unix())
	e2eutils.ApplyDeployConfigForks(deployConfig)
	// [Kroma: START]
	deployConfig.L2GenesisDeltaTimeOffset = nil
	deployConfig.L2GenesisEcotoneTimeOffset = nil
	deployConfig.L2GenesisKromaMPTTimeOffset = nil
	// [Kroma: END]
	require.NoError(t, deployConfig.Check(), "Deploy config is invalid, do you need to run make devnet-allocs?")
	l1Deployments := config.L1Deployments.Copy()
	require.NoError(t, l1Deployments.Check(deployConfig))

	require.Equal(t, secrets.Addresses().Batcher, deployConfig.BatchSenderAddress)
	require.Equal(t, secrets.Addresses().SequencerP2P, deployConfig.P2PSequencerAddress)
	// [Kroma: START]
	require.Equal(t, secrets.Addresses().TrustedValidator, deployConfig.ValidatorPoolTrustedValidator)
	require.Equal(t, secrets.Addresses().TrustedValidator, deployConfig.ValidatorManagerTrustedValidator)
	// [Kroma: END]

	// Tests depend on premine being filled with secrets addresses
	premine := make(map[common.Address]*big.Int)
	for _, addr := range secrets.Addresses().All() {
		premine[addr] = new(big.Int).Mul(big.NewInt(1000), big.NewInt(params.Ether))
	}

	return SystemConfig{
		Secrets:                secrets,
		Premine:                premine,
		DeployConfig:           deployConfig,
		L1Deployments:          config.L1Deployments,
		L1InfoPredeployAddress: predeploys.KromaL1BlockAddr,
		JWTFilePath:            writeDefaultJWT(t),
		JWTSecret:              testingJWTSecret,
		BlobsPath:              t.TempDir(),
		Nodes: map[string]*rollupNode.Config{
			"sequencer": {
				Driver: driver.Config{
					VerifierConfDepth:  0,
					SequencerConfDepth: 0,
					SequencerEnabled:   true,
				},
				// Submitter PrivKey is set in system start for rollup nodes where sequencer = true
				RPC: rollupNode.RPCConfig{
					ListenAddr:  "127.0.0.1",
					ListenPort:  0,
					EnableAdmin: true,
				},
				L1EpochPollInterval:         time.Second * 2,
				RuntimeConfigReloadInterval: time.Minute * 10,
				ConfigPersistence:           &rollupNode.DisabledConfigPersistence{},
				Sync:                        sync.Config{SyncMode: sync.CLSync},
			},
			"verifier": {
				Driver: driver.Config{
					VerifierConfDepth:  0,
					SequencerConfDepth: 0,
					SequencerEnabled:   false,
				},
				L1EpochPollInterval:         time.Second * 4,
				RuntimeConfigReloadInterval: time.Minute * 10,
				ConfigPersistence:           &rollupNode.DisabledConfigPersistence{},
				Sync:                        sync.Config{SyncMode: sync.CLSync},
			},
		},
		Loggers: map[string]log.Logger{
			"verifier":  testlog.Logger(t, log.LevelInfo).New("role", "verifier"),
			"sequencer": testlog.Logger(t, log.LevelInfo).New("role", "sequencer"),
			"batcher":   testlog.Logger(t, log.LevelInfo).New("role", "batcher"),
			"validator": testlog.Logger(t, log.LevelCrit).New("role", "validator"),
			// [Kroma: START]
			"challenger": testlog.Logger(t, log.LevelCrit).New("role", "challenger"),
			"guardian":   testlog.Logger(t, log.LevelCrit).New("role", "guardian"),
			// [Kroma: END]
		},
		GethOptions:            map[string][]geth.GethOption{},
		P2PTopology:            nil, // no P2P connectivity by default
		NonFinalizedOutputs:    false,
		ExternalL2Shim:         config.ExternalL2Shim,
		DataAvailabilityType:   batcherFlags.CalldataType,
		MaxPendingTransactions: 1,
		BatcherTargetNumFrames: 1,
		// [Kroma: START]
		ChallengeProofType: testdata.DefaultProofType,
		// [Kroma: END]
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

	DeployConfig  *genesis.DeployConfig
	L1Deployments *genesis.L1Deployments

	JWTFilePath string
	JWTSecret   [32]byte

	BlobsPath string

	Premine         map[common.Address]*big.Int
	Nodes           map[string]*rollupNode.Config // Per node config. Don't use populate rollup.Config
	Loggers         map[string]log.Logger
	GethOptions     map[string][]geth.GethOption
	ValidatorLogger log.Logger
	BatcherLogger   log.Logger

	ExternalL2Shim string

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

	// Configure data-availability type that is used by the batcher.
	DataAvailabilityType batcherFlags.DataAvailabilityType

	// Max L1 tx size for the batcher transactions
	BatcherMaxL1TxSizeBytes uint64

	// Target number of frames to create per channel. Can be used to create
	// multi-blob transactions.
	// Default is 1 if unset.
	BatcherTargetNumFrames int

	// whether to actually use BatcherMaxL1TxSizeBytes for blobs, insteaf of max blob size
	BatcherUseMaxTxSizeForBlobs bool

	// SupportL1TimeTravel determines if the L1 node supports quickly skipping forward in time
	SupportL1TimeTravel bool

	// MaxPendingTransactions determines how many transactions the batcher will try to send
	// concurrently. 0 means unlimited.
	MaxPendingTransactions uint64

	// [Kroma: START]
	// EnableChallenge enables challenge setup, the validator will act as a malicious one and
	// the challenger and guardian will act as honest parties.
	EnableChallenge bool

	// ChallengeProofType configures challenge proof type that is used for proving fault.
	ChallengeProofType testdata.ProofType

	// ValidatorVersion makes the version of validator system other than V1.
	ValidatorVersion uint8

	// SetupMPTMigration sets up the environment to migrate to Kroma MPT.
	SetupMPTMigration bool
	// [Kroma: END]
}

type GethInstance struct {
	Backend *geth_eth.Ethereum
	Node    *node.Node
}

func (gi *GethInstance) HTTPEndpoint() string {
	return gi.Node.HTTPEndpoint()
}

func (gi *GethInstance) WSEndpoint() string {
	return gi.Node.WSEndpoint()
}

func (gi *GethInstance) WSAuthEndpoint() string {
	return gi.Node.WSAuthEndpoint()
}

func (gi *GethInstance) HTTPAuthEndpoint() string {
	return gi.Node.HTTPAuthEndpoint()
}

func (gi *GethInstance) Close() error {
	return gi.Node.Close()
}

// EthInstance is either an in process Geth or external process exposing its
// endpoints over the network
type EthInstance interface {
	HTTPEndpoint() string
	WSEndpoint() string
	HTTPAuthEndpoint() string
	WSAuthEndpoint() string
	Close() error
}

type System struct {
	Cfg SystemConfig

	RollupConfig *rollup.Config

	L2GenesisCfg *core.Genesis

	// Connections to running nodes
	EthInstances map[string]EthInstance
	Clients      map[string]*ethclient.Client
	RawClients   map[string]*rpc.Client
	RollupNodes  map[string]*rollupNode.OpNode
	Validator    *validator.Validator
	// [Kroma: START]
	Challenger *validator.Validator
	Guardian   *validator.Validator
	// [Kroma: END]
	BatchSubmitter *bss.BatcherService
	Mocknet        mocknet.Mocknet

	L1BeaconAPIAddr string

	// TimeTravelClock is nil unless SystemConfig.SupportL1TimeTravel was set to true
	// It provides access to the clock instance used by the L1 node. Calling TimeTravelClock.AdvanceBy
	// allows tests to quickly time travel L1 into the future.
	// Note that this time travel may occur in a single block, creating a very large difference in the Time
	// on sequential blocks.
	TimeTravelClock *clock.AdvancingClock

	t      *testing.T
	closed atomic.Bool

	// rollupClients caches the lazily created RollupClient instances so they can be reused and closed
	rollupClients map[string]*sources.RollupClient
}

// AdvanceTime advances the system clock by the given duration.
// If the [System.TimeTravelClock] is nil, this is a no-op.
func (sys *System) AdvanceTime(d time.Duration) {
	if sys.TimeTravelClock != nil {
		sys.TimeTravelClock.AdvanceTime(d)
	}
}

func (sys *System) L1BeaconEndpoint() string {
	return sys.L1BeaconAPIAddr
}

func (sys *System) L1BeaconHTTPClient() *sources.BeaconHTTPClient {
	logger := testlog.Logger(sys.t, log.LevelInfo).New("component", "beaconClient")
	return sources.NewBeaconHTTPClient(client.NewBasicHTTPClient(sys.L1BeaconEndpoint(), logger))
}

func (sys *System) NodeEndpoint(name string) string {
	return selectEndpoint(sys.EthInstances[name])
}

func (sys *System) NodeClient(name string) *ethclient.Client {
	return sys.Clients[name]
}

func (sys *System) RollupEndpoint(name string) string {
	return sys.RollupNodes[name].HTTPEndpoint()
}

func (sys *System) RollupClient(name string) *sources.RollupClient {
	client, ok := sys.rollupClients[name]
	if ok {
		return client
	}
	logger := testlog.Logger(sys.t, log.LevelInfo).New("rollupClient", name)
	endpoint := sys.RollupEndpoint(name)
	client, err := dial.DialRollupClientWithTimeout(context.Background(), 30*time.Second, logger, endpoint)
	require.NoErrorf(sys.t, err, "Failed to dial rollup client %v", name)
	sys.rollupClients[name] = client
	return client
}

func (sys *System) L1Deployments() *genesis.L1Deployments {
	return sys.Cfg.L1Deployments
}

func (sys *System) RollupCfg() *rollup.Config {
	return sys.RollupConfig
}

func (sys *System) L2Genesis() *core.Genesis {
	return sys.L2GenesisCfg
}

func (sys *System) L1Slot(l1Timestamp uint64) uint64 {
	return (l1Timestamp - uint64(sys.Cfg.DeployConfig.L1GenesisBlockTimestamp)) /
		sys.Cfg.DeployConfig.L1BlockTime
}

func (sys *System) Close() {
	if !sys.closed.CompareAndSwap(false, true) {
		// Already closed.
		return
	}
	postCtx, postCancel := context.WithCancel(context.Background())
	postCancel() // immediate shutdown, no allowance for idling

	var combinedErr error
	if sys.Validator != nil {
		if err := sys.Validator.Stop(); err != nil {
			combinedErr = errors.Join(combinedErr, fmt.Errorf("stop L2OutputSubmitter: %w", err))
		}
	}
	// [Kroma: START]
	if sys.Challenger != nil {
		if err := sys.Challenger.Stop(); err != nil {
			combinedErr = errors.Join(combinedErr, fmt.Errorf("stop Challenger: %w", err))
		}
	}
	if sys.Guardian != nil {
		if err := sys.Guardian.Stop(); err != nil {
			combinedErr = errors.Join(combinedErr, fmt.Errorf("stop Guardian: %w", err))
		}
	}
	// [Kroma: END]
	if sys.BatchSubmitter != nil {
		if err := sys.BatchSubmitter.Kill(); err != nil && !errors.Is(err, bss.ErrAlreadyStopped) {
			combinedErr = errors.Join(combinedErr, fmt.Errorf("stop BatchSubmitter: %w", err))
		}
	}

	for name, node := range sys.RollupNodes {
		if err := node.Stop(postCtx); err != nil && !errors.Is(err, rollupNode.ErrAlreadyClosed) {
			combinedErr = errors.Join(combinedErr, fmt.Errorf("stop rollup node %v: %w", name, err))
		}
	}
	for name, ei := range sys.EthInstances {
		if err := ei.Close(); err != nil && !errors.Is(err, node.ErrNodeStopped) {
			combinedErr = errors.Join(combinedErr, fmt.Errorf("stop EthInstance %v: %w", name, err))
		}
	}
	for _, client := range sys.rollupClients {
		client.Close()
	}
	if sys.Mocknet != nil {
		if err := sys.Mocknet.Close(); err != nil {
			combinedErr = errors.Join(combinedErr, fmt.Errorf("stop Mocknet: %w", err))
		}
	}
	require.NoError(sys.t, combinedErr, "Failed to stop system")
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

func (cfg SystemConfig) Start(t *testing.T, _opts ...SystemConfigOption) (*System, error) {
	// [Kroma: START]
	cfg.DeployConfig.ValidatorPoolRoundDuration = cfg.DeployConfig.L2OutputOracleSubmissionInterval * cfg.DeployConfig.L2BlockTime / 2
	cfg.DeployConfig.ValidatorManagerRoundDurationSeconds = cfg.DeployConfig.L2OutputOracleSubmissionInterval * cfg.DeployConfig.L2BlockTime / 2
	// [Kroma: END]
	opts, err := NewSystemConfigOptions(_opts)
	if err != nil {
		return nil, err
	}

	sys := &System{
		t:             t,
		Cfg:           cfg,
		EthInstances:  make(map[string]EthInstance),
		Clients:       make(map[string]*ethclient.Client),
		RawClients:    make(map[string]*rpc.Client),
		RollupNodes:   make(map[string]*rollupNode.OpNode),
		rollupClients: make(map[string]*sources.RollupClient),
	}
	// Automatically stop the system at the end of the test
	t.Cleanup(sys.Close)

	c := clock.SystemClock
	if cfg.SupportL1TimeTravel {
		sys.TimeTravelClock = clock.NewAdvancingClock(100 * time.Millisecond)
		c = sys.TimeTravelClock
	}

	if err := cfg.DeployConfig.Check(); err != nil {
		return nil, err
	}

	l1Genesis, err := genesis.BuildL1DeveloperGenesis(cfg.DeployConfig, config.L1Allocs, config.L1Deployments)
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
	l2Genesis, err := genesis.BuildL2Genesis(cfg.DeployConfig, l1Block)
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
			MaxSequencerDrift:      cfg.DeployConfig.MaxSequencerDrift,
			SeqWindowSize:          cfg.DeployConfig.SequencerWindowSize,
			ChannelTimeout:         cfg.DeployConfig.ChannelTimeout,
			L1ChainID:              cfg.L1ChainIDBig(),
			L2ChainID:              cfg.L2ChainIDBig(),
			BatchInboxAddress:      cfg.DeployConfig.BatchInboxAddress,
			DepositContractAddress: cfg.DeployConfig.KromaPortalProxy,
			L1SystemConfigAddress:  cfg.DeployConfig.SystemConfigProxy,
			RegolithTime:           cfg.DeployConfig.RegolithTime(uint64(cfg.DeployConfig.L1GenesisBlockTimestamp)),
			CanyonTime:             cfg.DeployConfig.CanyonTime(uint64(cfg.DeployConfig.L1GenesisBlockTimestamp)),
			DeltaTime:              cfg.DeployConfig.DeltaTime(uint64(cfg.DeployConfig.L1GenesisBlockTimestamp)),
			EcotoneTime:            cfg.DeployConfig.EcotoneTime(uint64(cfg.DeployConfig.L1GenesisBlockTimestamp)),
			KromaMPTTime:           cfg.DeployConfig.KromaMPTTime(uint64(cfg.DeployConfig.L1GenesisBlockTimestamp)),
			FjordTime:              cfg.DeployConfig.FjordTime(uint64(cfg.DeployConfig.L1GenesisBlockTimestamp)),
			InteropTime:            cfg.DeployConfig.InteropTime(uint64(cfg.DeployConfig.L1GenesisBlockTimestamp)),
			/* [Kroma: START]
			ProtocolVersionsAddress: cfg.L1Deployments.ProtocolVersionsProxy,
			[Kroma: END] */
		}
	}
	defaultConfig := makeRollupConfig()
	if err := defaultConfig.Check(); err != nil {
		return nil, err
	}
	sys.RollupConfig = &defaultConfig

	// Create a fake Beacon node to hold on to blobs created by the L1 miner, and to serve them to L2
	bcn := fakebeacon.NewBeacon(testlog.Logger(t, log.LevelInfo).New("role", "l1_cl"),
		path.Join(cfg.BlobsPath, "l1_cl"), l1Genesis.Timestamp, cfg.DeployConfig.L1BlockTime)
	t.Cleanup(func() {
		_ = bcn.Close()
	})
	require.NoError(t, bcn.Start("127.0.0.1:0"))
	beaconApiAddr := bcn.BeaconAddr()
	require.NotEmpty(t, beaconApiAddr, "beacon API listener must be up")
	sys.L1BeaconAPIAddr = beaconApiAddr

	// Initialize nodes
	l1Node, l1Backend, err := geth.InitL1(cfg.DeployConfig.L1ChainID, cfg.DeployConfig.L1BlockTime, l1Genesis, c,
		path.Join(cfg.BlobsPath, "l1_el"), bcn, cfg.GethOptions["l1"]...)
	if err != nil {
		return nil, err
	}
	sys.EthInstances["l1"] = &GethInstance{
		Backend: l1Backend,
		Node:    l1Node,
	}
	err = l1Node.Start()
	if err != nil {
		return nil, err
	}

	// [Kroma: START]
	initializeNode := func(name string) (EthInstance, error) {
		var ethClient EthInstance
		if cfg.ExternalL2Shim == "" {
			node, backend, err := geth.InitL2(name, big.NewInt(int64(cfg.DeployConfig.L2ChainID)), l2Genesis, cfg.JWTFilePath, cfg.GethOptions[name]...)
			if err != nil {
				return nil, err
			}
			gethInst := &GethInstance{
				Backend: backend,
				Node:    node,
			}
			err = gethInst.Node.Start()
			if err != nil {
				return nil, err
			}
			ethClient = gethInst
		} else {
			if len(cfg.GethOptions[name]) > 0 {
				t.Skip("External L2 nodes do not support configuration through GethOptions")
			}
			ethClient = (&ExternalRunner{
				Name:    name,
				BinPath: cfg.ExternalL2Shim,
				Genesis: l2Genesis,
				JWTPath: cfg.JWTFilePath,
			}).Run(t)
		}
		return ethClient, nil
	}

	if cfg.SetupMPTMigration {
		ethClient, err := setupNodesForMPT(t, &cfg, initializeNode)
		if err != nil {
			return nil, err
		}
		sys.EthInstances["historical"] = ethClient
	}

	for name := range cfg.Nodes {
		if name == "historical" {
			continue
		}
		ethClient, err := initializeNode(name)
		if err != nil {
			return nil, err
		}
		// [Kroma: END]
		sys.EthInstances[name] = ethClient
	}

	// Configure connections to L1 and L2 for rollup nodes.
	// TODO: refactor testing to allow use of in-process rpc connections instead
	// of only websockets (which are required for external eth client tests).
	for name, nodeCfg := range cfg.Nodes {
		configureL1(nodeCfg, sys.EthInstances["l1"])
		configureL2(nodeCfg, sys.EthInstances[name], cfg.JWTSecret)
		if sys.RollupConfig.EcotoneTime != nil {
			nodeCfg.Beacon = &rollupNode.L1BeaconEndpointConfig{BeaconAddr: sys.L1BeaconAPIAddr}
		}
	}

	// Geth Clients
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	l1Srv, err := l1Node.RPCHandler()
	if err != nil {
		return nil, err
	}
	rawL1Client := rpc.DialInProc(l1Srv)
	l1Client := ethclient.NewClient(rawL1Client)
	sys.Clients["l1"] = l1Client
	sys.RawClients["l1"] = rawL1Client
	for name, ethInst := range sys.EthInstances {
		rawClient, err := rpc.DialContext(ctx, ethInst.WSEndpoint())
		if err != nil {
			return nil, err
		}
		client := ethclient.NewClient(rawClient)
		sys.RawClients[name] = rawClient
		sys.Clients[name] = client
	}

	_, err = geth.WaitForBlock(big.NewInt(2), l1Client, 6*time.Second*time.Duration(cfg.DeployConfig.L1BlockTime))
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
			h, err := sys.newMockNetPeer()
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
	snapLog := log.NewLogger(log.DiscardHandler())

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
		if err := c.LoadPersisted(cfg.Loggers[name]); err != nil {
			return nil, err
		}

		if p, ok := p2pNodes[name]; ok {
			c.P2P = p

			if c.Driver.SequencerEnabled && c.P2PSigner == nil {
				c.P2PSigner = &p2p.PreparedSigner{Signer: p2p.NewLocalSigner(cfg.Secrets.SequencerP2P)}
			}
		}

		c.Rollup.LogDescription(cfg.Loggers[name], chaincfg.L2ChainIDToNetworkDisplayName)
		l := cfg.Loggers[name]
		var cycle cliapp.Lifecycle
		c.Cancel = func(errCause error) {
			l.Warn("node requested early shutdown!", "err", errCause)
			go func() {
				postCtx, postCancel := context.WithCancel(context.Background())
				postCancel() // don't allow the stopping to continue for longer than needed
				if err := cycle.Stop(postCtx); err != nil {
					t.Error(err)
				}
				l.Warn("closed op-node!")
			}()
		}
		node, err := rollupNode.New(context.Background(), &c, l, snapLog, "", metrics.NewMetrics(""))
		if err != nil {
			return nil, err
		}
		cycle = node
		err = node.Start(context.Background())
		if err != nil {
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

	// Don't start batch submitter and proposer if there's no sequencer.
	if sys.RollupNodes["sequencer"] == nil {
		return sys, nil
	}

	// [Kroma: START]
	// Run validator node (L2 Output Submitter, Asserter)
	validatorCliCfg := validator.CLIConfig{
		L1EthRpc:                        sys.EthInstances["l1"].WSEndpoint(),
		L2EthRpc:                        sys.EthInstances["sequencer"].HTTPEndpoint(),
		RollupRpc:                       sys.RollupNodes["sequencer"].HTTPEndpoint(),
		L2OOAddress:                     config.L1Deployments.L2OutputOracleProxy.Hex(),
		ColosseumAddress:                config.L1Deployments.ColosseumProxy.Hex(),
		ValPoolAddress:                  config.L1Deployments.ValidatorPoolProxy.Hex(),
		ValMgrAddress:                   config.L1Deployments.ValidatorManagerProxy.Hex(),
		AssetManagerAddress:             config.L1Deployments.AssetManagerProxy.Hex(),
		ChallengePollInterval:           500 * time.Millisecond,
		TxMgrConfig:                     newTxMgrConfig(sys.EthInstances["l1"].WSEndpoint(), cfg.Secrets.TrustedValidator),
		AllowNonFinalized:               cfg.NonFinalizedOutputs,
		OutputSubmitterRetryInterval:    500 * time.Millisecond,
		OutputSubmitterRoundBuffer:      cfg.DeployConfig.L2OutputOracleSubmissionInterval / 2,
		OutputSubmitterEnabled:          true,
		OutputSubmitterAllowPublicRound: true,
		LogConfig: oplog.CLIConfig{
			Level:  log.LevelInfo,
			Format: oplog.FormatText,
		},
	}

	// deposit to ValidatorPool to be a validator
	validatorHelper := sys.ValidatorHelper()
	validatorHelper.DepositToValPool(cfg.Secrets.TrustedValidator, big.NewInt(params.Ether))

	if cfg.ValidatorVersion == valhelper.ValidatorV2 {
		// register to ValidatorManager to be a validator
		depositAmount := new(big.Int).Mul(cfg.DeployConfig.ValidatorManagerMinActivateAmount.ToInt(), common.Big256)
		validatorHelper.RegisterToValMgr(cfg.Secrets.TrustedValidator, depositAmount, cfg.Secrets.Addresses().TrustedValidator)

		// set up ValidatorRewardVault(Mallory) to be able to provide asset tokens to AssetManager
		validatorHelper.ApproveAssetToken(cfg.Secrets.Mallory, cfg.L1Deployments.AssetManagerProxy, new(big.Int).SetUint64(math.MaxUint64))

		// Redeploy and upgrade ValidatorPool to set the termination index to a smaller value for ValidatorManager test
		deployTx, upgradeTx, err := e2eutils.RedeployValPoolToTerminate(
			cfg.DeployConfig.ValidatorPoolTerminateOutputIndex.ToInt(),
			l1Client,
			cfg.Secrets,
			cfg.L1ChainIDBig(),
			cfg.L1Deployments,
			cfg.DeployConfig,
		)
		if err != nil {
			return nil, fmt.Errorf("unable to redeploy ValidatorPool: %w", err)
		}

		// Check deploy tx and upgrade tx submission were successful
		err = waitTxs(l1Client, deployTx.Hash(), upgradeTx.Hash())
		if err != nil {
			return nil, err
		}
	}

	validatorCfg, err := validator.NewValidatorConfig(validatorCliCfg, sys.Cfg.Loggers["validator"], validatormetrics.NoopMetrics)
	if err != nil {
		return nil, fmt.Errorf("unable to init validator config: %w", err)
	}

	// Replace to mock RPC client
	cl, err := rpc.DialHTTP(validatorCliCfg.RollupRpc)
	if err != nil {
		return nil, fmt.Errorf("unable to init validator rollup rpc client: %w", err)
	}
	rpcCl := client.NewBaseRPCClient(cl)
	validatorMaliciousL2RPC, err := e2eutils.NewMaliciousL2RPC(rpcCl, cfg.ChallengeProofType)
	if err != nil {
		return nil, fmt.Errorf("unable to init validator malicious rpc: %w", err)
	}
	validatorCfg.RollupClient = sources.NewRollupClient(validatorMaliciousL2RPC)
	validatorCfg.L2Client = sys.Clients["sequencer"]

	// For challenge setup, set target block number for submitting invalid output
	if cfg.EnableChallenge {
		validatorMaliciousL2RPC.SetTargetBlockNumber(testdata.TargetBlockNumber)
	}

	sys.Validator, err = validator.NewValidator(*validatorCfg, sys.Cfg.Loggers["validator"], validatormetrics.NoopMetrics)
	if err != nil {
		return nil, fmt.Errorf("unable to setup validator: %w", err)
	}

	if err := sys.Validator.Start(); err != nil {
		return nil, fmt.Errorf("unable to start validator: %w", err)
	}

	if cfg.EnableChallenge {
		if err = startChallengeSystem(sys, &cfg); err != nil {
			return nil, fmt.Errorf("unable to start challenge system: %w", err)
		}
	}
	// [Kroma: END]

	var batchType uint = derive.SingularBatchType
	if cfg.DeployConfig.L2GenesisDeltaTimeOffset != nil && *cfg.DeployConfig.L2GenesisDeltaTimeOffset == hexutil.Uint64(0) {
		batchType = derive.SpanBatchType
	}
	// batcher defaults if unset
	batcherMaxL1TxSizeBytes := cfg.BatcherMaxL1TxSizeBytes
	if batcherMaxL1TxSizeBytes == 0 {
		batcherMaxL1TxSizeBytes = 120_000
	}
	batcherTargetNumFrames := cfg.BatcherTargetNumFrames
	if batcherTargetNumFrames == 0 {
		batcherTargetNumFrames = 1
	}
	batcherCLIConfig := &bss.CLIConfig{
		L1EthRpc:                 sys.EthInstances["l1"].WSEndpoint(),
		L2EthRpc:                 sys.EthInstances["sequencer"].WSEndpoint(),
		RollupRpc:                sys.RollupNodes["sequencer"].HTTPEndpoint(),
		MaxPendingTransactions:   cfg.MaxPendingTransactions,
		MaxChannelDuration:       1,
		MaxL1TxSize:              batcherMaxL1TxSizeBytes,
		TestUseMaxTxSizeForBlobs: cfg.BatcherUseMaxTxSizeForBlobs,
		TargetNumFrames:          int(batcherTargetNumFrames),
		ApproxComprRatio:         0.4,
		SubSafetyMargin:          4,
		PollInterval:             50 * time.Millisecond,
		TxMgrConfig:              newTxMgrConfig(sys.EthInstances["l1"].WSEndpoint(), cfg.Secrets.Batcher),
		LogConfig: oplog.CLIConfig{
			Level:  log.LevelInfo,
			Format: oplog.FormatText,
		},
		Stopped:              sys.Cfg.DisableBatcher, // Batch submitter may be enabled later
		BatchType:            batchType,
		DataAvailabilityType: sys.Cfg.DataAvailabilityType,
	}
	// Batch Submitter
	batcher, err := bss.BatcherServiceFromCLIConfig(context.Background(), "0.0.1", batcherCLIConfig, sys.Cfg.Loggers["batcher"])
	if err != nil {
		return nil, fmt.Errorf("failed to setup batch submitter: %w", err)
	}
	if err := batcher.Start(context.Background()); err != nil {
		return nil, errors.Join(fmt.Errorf("failed to start batch submitter: %w", err), batcher.Stop(context.Background()))
	}
	sys.BatchSubmitter = batcher

	return sys, nil
}

// [Kroma: START]

func setupNodesForMPT(
	t *testing.T, cfg *SystemConfig, initializeNode func(name string) (EthInstance, error),
) (EthInstance, error) {
	// Setup historical rpc node.
	historicalName := "historical"
	cfg.Nodes[historicalName] = &rollupNode.Config{
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
	cfg.Loggers[historicalName] = testlog.Logger(t, log.LevelInfo).New("role", historicalName)
	cfg.GethOptions[historicalName] = append(cfg.GethOptions[historicalName], []geth.GethOption{
		func(ethCfg *ethconfig.Config, nodeCfg *node.Config) error {
			nodeCfg.HTTPModules = []string{"debug", "eth"}
			ethCfg.RollupHistoricalRPC = ""
			ethCfg.DisableMPTMigration = true
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

	ethClient, err := initializeNode(historicalName)
	if err != nil {
		return nil, err
	}

	// Set historical rpc endpoint.
	historicalEndpoint := ethClient.HTTPEndpoint()
	for name := range cfg.Nodes {
		if name == historicalName {
			continue
		}
		name := name
		cfg.GethOptions[name] = append(cfg.GethOptions[name], []geth.GethOption{
			func(ethCfg *ethconfig.Config, nodeCfg *node.Config) error {
				// Since the migration process requires preimages, enable storing preimage option.
				ethCfg.Preimages = true
				ethCfg.RollupHistoricalRPC = historicalEndpoint
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

	return ethClient, nil
}

func startChallengeSystem(sys *System, cfg *SystemConfig) error {
	l1Client := sys.NodeClient("l1")

	// Deploy MockColosseum which has setL1Head function for challenge test
	deployTx, upgradeTx, err := e2eutils.ReplaceWithMockColosseum(
		l1Client,
		cfg.Secrets.SysCfgOwner,
		cfg.L1ChainIDBig(),
		cfg.L1Deployments,
		cfg.DeployConfig,
	)
	if err != nil {
		return fmt.Errorf("unable to replace Colosseum: %w", err)
	}

	// Check deploy tx and upgrade tx submission were successful
	err = waitTxs(l1Client, deployTx.Hash(), upgradeTx.Hash())
	if err != nil {
		return err
	}

	// Deploy SP1Verifier contract
	sp1Verifier, deployTx, err := e2eutils.DeploySP1Verifier(
		l1Client,
		cfg.Secrets.SysCfgOwner,
		cfg.L1ChainIDBig(),
	)
	if err != nil {
		return fmt.Errorf("unable to deploy SP1Verifier: %w", err)
	}

	err = waitTxs(l1Client, deployTx.Hash())
	if err != nil {
		return err
	}

	// Deploy new ZKProofVerifier impl and upgrade ZKProofVerifier proxy
	deployTx, upgradeTx, err = e2eutils.RedeployZKProofVerifier(
		l1Client,
		cfg.Secrets.SysCfgOwner,
		cfg.L1ChainIDBig(),
		cfg.L1Deployments,
		cfg.DeployConfig,
		sp1Verifier,
		cfg.DeployConfig.ZKProofVerifierVKey,
	)
	if err != nil {
		return fmt.Errorf("unable to redeploy ZKProofVerifier: %w", err)
	}

	// Check deploy tx and upgrade tx submission were successful
	err = waitTxs(l1Client, deployTx.Hash(), upgradeTx.Hash())
	if err != nil {
		return err
	}

	// Run validator node (Challenger)
	challengerCliCfg := validator.CLIConfig{
		L1EthRpc:              sys.EthInstances["l1"].WSEndpoint(),
		L2EthRpc:              sys.EthInstances["sequencer"].HTTPEndpoint(),
		RollupRpc:             sys.RollupNodes["sequencer"].HTTPEndpoint(),
		L2OOAddress:           config.L1Deployments.L2OutputOracleProxy.Hex(),
		ColosseumAddress:      config.L1Deployments.ColosseumProxy.Hex(),
		ValPoolAddress:        config.L1Deployments.ValidatorPoolProxy.Hex(),
		ValMgrAddress:         config.L1Deployments.ValidatorManagerProxy.Hex(),
		AssetManagerAddress:   config.L1Deployments.AssetManagerProxy.Hex(),
		ChallengePollInterval: 500 * time.Millisecond,
		TxMgrConfig:           newTxMgrConfig(sys.EthInstances["l1"].WSEndpoint(), cfg.Secrets.Challenger1),
		AllowNonFinalized:     cfg.NonFinalizedOutputs,
		ChallengerEnabled:     false, // to bypass prover connection check in NewValidatorConfig
		LogConfig: oplog.CLIConfig{
			Level:  log.LevelInfo,
			Format: oplog.FormatText,
		},
	}

	challengerCfg, err := validator.NewValidatorConfig(challengerCliCfg, sys.Cfg.Loggers["challenger"], validatormetrics.NoopMetrics)
	if err != nil {
		return fmt.Errorf("unable to init challenger config: %w", err)
	}

	// Replace to mock RPC client
	cl, err := rpc.DialHTTP(challengerCliCfg.RollupRpc)
	if err != nil {
		return fmt.Errorf("unable to init challenger rollup rpc client: %w", err)
	}
	rpcCl := client.NewBaseRPCClient(cl)
	challengerHonestL2RPC, err := e2eutils.NewHonestL2RPC(rpcCl, cfg.ChallengeProofType)
	if err != nil {
		return fmt.Errorf("unable to init challenger honest rpc: %w", err)
	}
	challengerCfg.RollupClient = sources.NewRollupClient(challengerHonestL2RPC)
	challengerCfg.L2Client = sys.Clients["sequencer"]

	// For challenge setup, set target block number for submitting invalid output
	challengerHonestL2RPC.SetTargetBlockNumber(testdata.TargetBlockNumber)

	// Replace to mock proof fetcher
	challengerCfg.ChallengerEnabled = true
	challengerCfg.ZkEVMProofFetcher = chal.NewZkEVMProofFetcher(e2eutils.NewMockRPCWithData("./testdata/proof"))
	challengerCfg.ZkVMProofFetcher = chal.NewZkVMProofFetcher(e2eutils.NewMockRPC())
	challengerCfg.WitnessGenerator = chal.NewWitnessGenerator(e2eutils.NewMockRPC())
	sys.Challenger, err = validator.NewValidator(*challengerCfg, sys.Cfg.Loggers["challenger"], validatormetrics.NoopMetrics)
	if err != nil {
		return fmt.Errorf("unable to setup challenger: %w", err)
	}

	if err := sys.Challenger.Start(); err != nil {
		return fmt.Errorf("unable to start challenger: %w", err)
	}

	// Run validator node (Guardian)
	guardianCliCfg := validator.CLIConfig{
		L1EthRpc:               sys.EthInstances["l1"].WSEndpoint(),
		L2EthRpc:               sys.EthInstances["sequencer"].HTTPEndpoint(),
		RollupRpc:              sys.RollupNodes["sequencer"].HTTPEndpoint(),
		L2OOAddress:            config.L1Deployments.L2OutputOracleProxy.Hex(),
		ColosseumAddress:       config.L1Deployments.ColosseumProxy.Hex(),
		ValPoolAddress:         config.L1Deployments.ValidatorPoolProxy.Hex(),
		ValMgrAddress:          config.L1Deployments.ValidatorManagerProxy.Hex(),
		AssetManagerAddress:    config.L1Deployments.AssetManagerProxy.Hex(),
		SecurityCouncilAddress: config.L1Deployments.SecurityCouncilProxy.Hex(),
		TxMgrConfig:            newTxMgrConfig(sys.EthInstances["l1"].WSEndpoint(), cfg.Secrets.Guardian),
		AllowNonFinalized:      cfg.NonFinalizedOutputs,
		GuardianEnabled:        true,
		GuardianPollInterval:   time.Second,
		LogConfig: oplog.CLIConfig{
			Level:  log.LevelInfo,
			Format: oplog.FormatText,
		},
	}

	guardianCfg, err := validator.NewValidatorConfig(guardianCliCfg, sys.Cfg.Loggers["guardian"], validatormetrics.NoopMetrics)
	if err != nil {
		return fmt.Errorf("unable to init guardian config: %w", err)
	}

	// Replace to mock RPC client
	cl, err = rpc.DialHTTP(guardianCliCfg.RollupRpc)
	if err != nil {
		return fmt.Errorf("unable to init guardian rollup rpc client: %w", err)
	}
	rpcCl = client.NewBaseRPCClient(cl)
	guardianHonestL2RPC, err := e2eutils.NewHonestL2RPC(rpcCl, cfg.ChallengeProofType)
	if err != nil {
		return fmt.Errorf("unable to init guardian honest rpc: %w", err)
	}
	guardianCfg.RollupClient = sources.NewRollupClient(guardianHonestL2RPC)
	guardianCfg.L2Client = sys.Clients["sequencer"]

	// For challenge setup, set target block number for submitting invalid output
	guardianHonestL2RPC.SetTargetBlockNumber(testdata.TargetBlockNumber)

	sys.Guardian, err = validator.NewValidator(*guardianCfg, sys.Cfg.Loggers["guardian"], validatormetrics.NoopMetrics)
	if err != nil {
		return fmt.Errorf("unable to setup guardian: %w", err)
	}

	if err := sys.Guardian.Start(); err != nil {
		return fmt.Errorf("unable to start guardian: %w", err)
	}

	return nil
}

func waitTxs(client *ethclient.Client, txs ...common.Hash) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	for _, tx := range txs {
		_, err := wait.ForReceiptOK(ctx, client, tx)
		if err != nil {
			return fmt.Errorf("failed to wait tx success: %w", err)
		}
	}
	return nil
}

// [Kroma: END]

// IP6 range that gets blackholed (in case our traffic ever makes it out onto
// the internet).
var blackholeIP6 = net.ParseIP("100::")

// mocknet doesn't allow us to add a peerstore without fully creating the peer ourselves
func (sys *System) newMockNetPeer() (host.Host, error) {
	sk, _, err := ic.GenerateECDSAKeyPair(rand.Reader)
	if err != nil {
		return nil, err
	}
	id, err := peer.IDFromPrivateKey(sk)
	if err != nil {
		return nil, err
	}
	suffix := id
	if len(id) > 8 {
		suffix = id[len(id)-8:]
	}
	ip := append(net.IP{}, blackholeIP6...)
	copy(ip[net.IPv6len-len(suffix):], suffix)
	a, err := ma.NewMultiaddr(fmt.Sprintf("/ip6/%s/tcp/4242", ip))
	if err != nil {
		return nil, fmt.Errorf("failed to create test multiaddr: %w", err)
	}
	p, err := peer.IDFromPublicKey(sk.GetPublic())
	if err != nil {
		return nil, err
	}

	ps, err := pstoremem.NewPeerstore()
	if err != nil {
		return nil, err
	}
	ps.AddAddr(p, a, peerstore.PermanentAddrTTL)
	_ = ps.AddPrivKey(p, sk)
	_ = ps.AddPubKey(p, sk.GetPublic())

	ds := dsSync.MutexWrap(ds.NewMapDatastore())
	eps, err := store.NewExtendedPeerstore(context.Background(), log.Root(), clock.SystemClock, ps, ds, 24*time.Hour)
	if err != nil {
		return nil, err
	}
	return sys.Mocknet.AddPeerWithPeerstore(p, eps)
}

func (sys *System) BatcherHelper() *batcher.Helper {
	return batcher.NewHelper(sys.t, sys.Cfg.Secrets.Batcher, sys.RollupConfig, sys.NodeClient("l1"))
}

// [Kroma: START]

func (sys *System) ValidatorHelper() *valhelper.Helper {
	return valhelper.NewHelper(sys.t, sys.NodeClient("l1"), sys.Cfg.L1ChainIDBig(), sys.Cfg.DeployConfig.L1BlockTime)
}

// [Kroma: END]

func UseHTTP() bool {
	return os.Getenv("OP_E2E_USE_HTTP") == "true"
}

func selectEndpoint(node EthInstance) string {
	if UseHTTP() {
		log.Info("using HTTP client")
		return node.HTTPEndpoint()
	}
	return node.WSEndpoint()
}

func configureL1(rollupNodeCfg *rollupNode.Config, l1Node EthInstance) {
	l1EndpointConfig := selectEndpoint(l1Node)
	rollupNodeCfg.L1 = &rollupNode.L1EndpointConfig{
		L1NodeAddr:       l1EndpointConfig,
		L1TrustRPC:       false,
		L1RPCKind:        sources.RPCKindStandard,
		RateLimit:        0,
		BatchSize:        20,
		HttpPollInterval: time.Millisecond * 100,
		MaxConcurrency:   10,
	}
}

type WSOrHTTPEndpoint interface {
	WSAuthEndpoint() string
	HTTPAuthEndpoint() string
}

func configureL2(rollupNodeCfg *rollupNode.Config, l2Node WSOrHTTPEndpoint, jwtSecret [32]byte) {
	l2EndpointConfig := l2Node.WSAuthEndpoint()
	if UseHTTP() {
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

func hexPriv(in *ecdsa.PrivateKey) string {
	b := e2eutils.EncodePrivKey(in)
	return hexutil.Encode(b)
}
