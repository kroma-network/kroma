package validator

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/ethereum-optimism/optimism/op-node/rollup"
	opservice "github.com/ethereum-optimism/optimism/op-service"
	"github.com/ethereum-optimism/optimism/op-service/client"
	"github.com/ethereum-optimism/optimism/op-service/dial"
	oplog "github.com/ethereum-optimism/optimism/op-service/log"
	opmetrics "github.com/ethereum-optimism/optimism/op-service/metrics"
	pprof "github.com/ethereum-optimism/optimism/op-service/pprof"
	oprpc "github.com/ethereum-optimism/optimism/op-service/rpc"
	"github.com/ethereum-optimism/optimism/op-service/sources"
	"github.com/ethereum-optimism/optimism/op-service/txmgr"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/urfave/cli/v2"

	chal "github.com/kroma-network/kroma/kroma-validator/challenge"
	"github.com/kroma-network/kroma/kroma-validator/flags"
	"github.com/kroma-network/kroma/kroma-validator/metrics"
)

// Config contains the well typed fields that are used to initialize the output submitter.
// It is intended for programmatic use.
type Config struct {
	L2OutputOracleAddr              common.Address
	ColosseumAddr                   common.Address
	SecurityCouncilAddr             common.Address
	ValidatorPoolAddr               common.Address
	ValidatorManagerAddr            common.Address
	AssetManagerAddr                common.Address
	NetworkTimeout                  time.Duration
	TxManager                       *txmgr.BufferedTxManager
	L1Client                        *ethclient.Client
	L2Client                        *ethclient.Client
	RollupClient                    *sources.RollupClient
	RollupConfig                    *rollup.Config
	ChallengePollInterval           time.Duration
	AllowNonFinalized               bool
	OutputSubmitterEnabled          bool
	OutputSubmitterAllowPublicRound bool
	OutputSubmitterRetryInterval    time.Duration
	OutputSubmitterRoundBuffer      uint64
	ChallengerEnabled               bool
	ZkEVMProofFetcher               *chal.ZkEVMProofFetcher
	ZkVMProofFetcher                *chal.ZkVMProofFetcher
	WitnessGenerator                *chal.WitnessGenerator
	GuardianEnabled                 bool
	GuardianPollInterval            time.Duration
}

// Check ensures that the [Config] is valid.
func (c *Config) Check() error {
	if err := c.RollupConfig.Check(); err != nil {
		return err
	}
	return nil
}

// CLIConfig is a well typed config that is parsed from the CLI params.
// This also contains config options for auxiliary services.
// It is transformed into a `Config` before the Validator is started.
type CLIConfig struct {
	// L1EthRpc is the Websocket provider URL for L1.
	L1EthRpc string

	// L2EthRpc is the HTTP provider URL for the L2 execution engine.
	L2EthRpc string

	// RollupRpc is the HTTP provider URL for the rollup node.
	RollupRpc string

	// L2OOAddress is the L2OutputOracle contract address.
	L2OOAddress string

	// ColosseumAddress is the Colosseum contract address.
	ColosseumAddress string

	// SecurityCouncilAddress is the SecurityCouncil contract address.
	SecurityCouncilAddress string

	// ValPoolAddress is the ValidatorPool contract address.
	ValPoolAddress string

	// ValMgrAddress is the ValidatorManager contract address.
	ValMgrAddress string

	// AssetManagerAddress is the AssetManager contract address.
	AssetManagerAddress string

	// ChallengePollInterval is how frequently to poll L1 to handle related challenges.
	ChallengePollInterval time.Duration

	// AllowNonFinalized can be set to true to submit outputs
	// for L2 blocks derived from non-finalized L1 data.
	AllowNonFinalized bool

	OutputSubmitterEnabled bool

	OutputSubmitterAllowPublicRound bool

	// OutputSubmitterRetryInterval is how frequently to retry output submission.
	OutputSubmitterRetryInterval time.Duration

	// OutputSubmitterRoundBuffer is how many blocks before each round to start trying submission.
	OutputSubmitterRoundBuffer uint64

	ChallengerEnabled bool

	// ZkEVMProverRPC is the URL of zkEVM prover JSON-RPC server.
	ZkEVMProverRPC string

	// ZkEVMNetworkTimeout is timeout to be connected with zkEVM prover.
	ZkEVMNetworkTimeout time.Duration

	// ZkVMProverRPC is the URL of zkVM prover JSON-RPC server.
	ZkVMProverRPC string

	// WitnessGeneratorRPC is the URL of zkVM witness generator JSON-RPC server.
	WitnessGeneratorRPC string

	GuardianEnabled bool

	// GuardianPollInterval is how frequently to poll L1 for inspection.
	GuardianPollInterval time.Duration

	TxMgrConfig   txmgr.CLIConfig
	RPCConfig     oprpc.CLIConfig
	LogConfig     oplog.CLIConfig
	MetricsConfig opmetrics.CLIConfig
	PprofConfig   pprof.CLIConfig
}

func (c CLIConfig) Check() error {
	if !(c.OutputSubmitterEnabled || c.ChallengerEnabled || c.GuardianEnabled) {
		return errors.New("one of output submitter, challenger, guardian should be enabled")
	}
	if err := c.RPCConfig.Check(); err != nil {
		return err
	}
	if err := c.MetricsConfig.Check(); err != nil {
		return err
	}
	if err := c.PprofConfig.Check(); err != nil {
		return err
	}
	if err := c.TxMgrConfig.Check(); err != nil {
		return err
	}
	return nil
}

// NewConfig parses the Config from the provided flags or environment variables.
func NewConfig(ctx *cli.Context) CLIConfig {
	return CLIConfig{
		// Required Flags
		L1EthRpc:               ctx.String(flags.L1EthRpcFlag.Name),
		L2EthRpc:               ctx.String(flags.L2EthRpcFlag.Name),
		RollupRpc:              ctx.String(flags.RollupRpcFlag.Name),
		L2OOAddress:            ctx.String(flags.L2OOAddressFlag.Name),
		ColosseumAddress:       ctx.String(flags.ColosseumAddressFlag.Name),
		ValPoolAddress:         ctx.String(flags.ValPoolAddressFlag.Name),
		ValMgrAddress:          ctx.String(flags.ValMgrAddressFlag.Name),
		AssetManagerAddress:    ctx.String(flags.AssetManagerAddressFlag.Name),
		OutputSubmitterEnabled: ctx.Bool(flags.OutputSubmitterEnabledFlag.Name),
		ChallengerEnabled:      ctx.Bool(flags.ChallengerEnabledFlag.Name),

		// Optional Flags
		ChallengePollInterval:           ctx.Duration(flags.ChallengePollIntervalFlag.Name),
		AllowNonFinalized:               ctx.Bool(flags.AllowNonFinalizedFlag.Name),
		OutputSubmitterRetryInterval:    ctx.Duration(flags.OutputSubmitterRetryIntervalFlag.Name),
		OutputSubmitterRoundBuffer:      ctx.Uint64(flags.OutputSubmitterRoundBufferFlag.Name),
		OutputSubmitterAllowPublicRound: ctx.Bool(flags.OutputSubmitterAllowPublicRoundFlag.Name),
		ZkEVMProverRPC:                  ctx.String(flags.ZkEVMProverRPCFlag.Name),
		ZkEVMNetworkTimeout:             ctx.Duration(flags.ZkEVMNetworkTimeoutFlag.Name),
		ZkVMProverRPC:                   ctx.String(flags.ZkVMProverRPCFlag.Name),
		WitnessGeneratorRPC:             ctx.String(flags.WitnessGeneratorRPCFlag.Name),
		GuardianEnabled:                 ctx.Bool(flags.GuardianEnabledFlag.Name),
		SecurityCouncilAddress:          ctx.String(flags.SecurityCouncilAddressFlag.Name),
		GuardianPollInterval:            ctx.Duration(flags.GuardianPollIntervalFlag.Name),

		TxMgrConfig:   txmgr.ReadCLIConfig(ctx),
		RPCConfig:     oprpc.ReadCLIConfig(ctx),
		LogConfig:     oplog.ReadCLIConfig(ctx),
		MetricsConfig: opmetrics.ReadCLIConfig(ctx),
		PprofConfig:   pprof.ReadCLIConfig(ctx),
	}
}

// NewValidatorConfig creates a validator config with given the CLIConfig
func NewValidatorConfig(cfg CLIConfig, l log.Logger, m metrics.Metricer) (*Config, error) {
	l2OOAddress, err := opservice.ParseAddress(cfg.L2OOAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to parse L2OOAddress: %w", err)
	}

	colosseumAddress, err := opservice.ParseAddress(cfg.ColosseumAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to parse ColosseumAddress: %w", err)
	}

	var securityCouncilAddress common.Address
	if cfg.GuardianEnabled {
		securityCouncilAddress, err = opservice.ParseAddress(cfg.SecurityCouncilAddress)
		if err != nil {
			return nil, fmt.Errorf("failed to parse SecurityCouncilAddress: %w", err)
		}
	}

	valPoolAddress, err := opservice.ParseAddress(cfg.ValPoolAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to parse ValPoolAddress: %w", err)
	}

	valMgrAddress, err := opservice.ParseAddress(cfg.ValMgrAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to parse ValMgrAddress: %w", err)
	}

	assetManagerAddress, err := opservice.ParseAddress(cfg.AssetManagerAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to parse AssetManagerAddress: %w", err)
	}

	txManager, err := txmgr.NewBufferedTxManager("validator", l, m, cfg.TxMgrConfig)
	if err != nil {
		return nil, err
	}

	// Connect to L1 and L2 providers. Perform these last since they are the most expensive.
	ctx := context.Background()
	l1Client, err := dial.DialEthClientWithTimeout(ctx, dial.DefaultDialTimeout, l, cfg.L1EthRpc)
	if err != nil {
		return nil, fmt.Errorf("failed to dial L1 RPC: %w", err)
	}

	l2Client, err := dial.DialEthClientWithTimeout(ctx, dial.DefaultDialTimeout, l, cfg.L2EthRpc)
	if err != nil {
		return nil, fmt.Errorf("failed to dial L2 RPC: %w", err)
	}

	rollupClient, err := dial.DialRollupClientWithTimeout(ctx, dial.DefaultDialTimeout, l, cfg.RollupRpc)
	if err != nil {
		return nil, fmt.Errorf("failed to dial rollup node RPC: %w", err)
	}

	rollupConfig, err := rollupClient.RollupConfig(ctx)
	if err != nil {
		return nil, err
	}

	// TODO(seolaoh): remove zkEVMProofFetcher after zkVM transition completed
	var zkEVMProofFetcher *chal.ZkEVMProofFetcher
	var zkVMProofFetcher *chal.ZkVMProofFetcher
	var witnessGenerator *chal.WitnessGenerator
	if cfg.ChallengerEnabled {
		if rollupConfig.IsKromaMPT(uint64(time.Now().Unix())) {
			pc, err := client.NewRPC(ctx, l.New("service", "prover"), cfg.ZkVMProverRPC)
			if err != nil {
				return nil, fmt.Errorf("failed to create zkVM prover rpc client: %w", err)
			}
			zkVMProofFetcher = chal.NewZkVMProofFetcher(pc)

			wc, err := client.NewRPC(ctx, l.New("service", "witness"), cfg.WitnessGeneratorRPC)
			if err != nil {
				return nil, fmt.Errorf("failed to create witness generator rpc client: %w", err)
			}
			witnessGenerator = chal.NewWitnessGenerator(wc)

			proverSpec, err := zkVMProofFetcher.Spec(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to request spec of zkVM prover: %w", err)
			}
			witnessGenSpec, err := witnessGenerator.Spec(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to reqeust spec of witness generator: %w", err)
			}
			if proverSpec.SP1Version != witnessGenSpec.SP1Version {
				return nil, errors.New("SP1 version of zkVM prover and witness generator mismatched")
			}
		} else {
			clientOpt := rpc.WithHTTPClient(&http.Client{
				Timeout: cfg.ZkEVMNetworkTimeout,
			})
			opts := []client.RPCOption{
				client.WithGethRPCOptions(clientOpt),
			}
			pc, err := client.NewRPC(ctx, l.New("service", "prover"), cfg.ZkEVMProverRPC, opts...)
			if err != nil {
				return nil, fmt.Errorf("failed to create zkEVM prover rpc client: %w", err)
			}
			zkEVMProofFetcher = chal.NewZkEVMProofFetcher(pc)
		}
	}

	return &Config{
		L2OutputOracleAddr:              l2OOAddress,
		ColosseumAddr:                   colosseumAddress,
		SecurityCouncilAddr:             securityCouncilAddress,
		ValidatorPoolAddr:               valPoolAddress,
		ValidatorManagerAddr:            valMgrAddress,
		AssetManagerAddr:                assetManagerAddress,
		NetworkTimeout:                  cfg.TxMgrConfig.NetworkTimeout,
		TxManager:                       txManager,
		L1Client:                        l1Client,
		L2Client:                        l2Client,
		RollupClient:                    rollupClient,
		RollupConfig:                    rollupConfig,
		ChallengePollInterval:           cfg.ChallengePollInterval,
		AllowNonFinalized:               cfg.AllowNonFinalized,
		OutputSubmitterEnabled:          cfg.OutputSubmitterEnabled,
		OutputSubmitterAllowPublicRound: cfg.OutputSubmitterAllowPublicRound,
		OutputSubmitterRetryInterval:    cfg.OutputSubmitterRetryInterval,
		OutputSubmitterRoundBuffer:      cfg.OutputSubmitterRoundBuffer,
		ChallengerEnabled:               cfg.ChallengerEnabled,
		ZkEVMProofFetcher:               zkEVMProofFetcher,
		ZkVMProofFetcher:                zkVMProofFetcher,
		WitnessGenerator:                witnessGenerator,
		GuardianEnabled:                 cfg.GuardianEnabled,
		GuardianPollInterval:            cfg.GuardianPollInterval,
	}, nil
}
