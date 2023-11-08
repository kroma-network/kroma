package validator

import (
	"context"
	"errors"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/urfave/cli"

	"github.com/ethereum-optimism/optimism/op-node/rollup"
	"github.com/ethereum-optimism/optimism/op-node/sources"
	klog "github.com/ethereum-optimism/optimism/op-service/log"
	kmetrics "github.com/ethereum-optimism/optimism/op-service/metrics"
	kpprof "github.com/ethereum-optimism/optimism/op-service/pprof"
	krpc "github.com/ethereum-optimism/optimism/op-service/rpc"
	"github.com/ethereum-optimism/optimism/op-service/txmgr"
	chal "github.com/kroma-network/kroma/components/validator/challenge"
	"github.com/kroma-network/kroma/components/validator/flags"
	"github.com/kroma-network/kroma/components/validator/metrics"
	"github.com/kroma-network/kroma/utils"
)

// Config contains the well typed fields that are used to initialize the output submitter.
// It is intended for programmatic use.
type Config struct {
	L2OutputOracleAddr              common.Address
	ColosseumAddr                   common.Address
	SecurityCouncilAddr             common.Address
	ValidatorPoolAddr               common.Address
	ChallengerPollInterval          time.Duration
	NetworkTimeout                  time.Duration
	TxManager                       *txmgr.BufferedTxManager
	L1Client                        *ethclient.Client
	L2Client                        *ethclient.Client
	RollupClient                    *sources.RollupClient
	RollupConfig                    *rollup.Config
	AllowNonFinalized               bool
	OutputSubmitterEnabled          bool
	OutputSubmitterAllowPublicRound bool
	OutputSubmitterRetryInterval    time.Duration
	OutputSubmitterRoundBuffer      uint64
	ChallengerEnabled               bool
	GuardianEnabled                 bool
	ProofFetcher                    ProofFetcher
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

	// ChallengerPollInterval is how frequently to poll L2 for new finalized outputs.
	ChallengerPollInterval time.Duration

	// ProverRPC is the URL of prover jsonRPC server.
	ProverRPC string

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

	GuardianEnabled bool

	FetchingProofTimeout time.Duration

	TxMgrConfig   txmgr.CLIConfig
	RPCConfig     krpc.CLIConfig
	LogConfig     klog.CLIConfig
	MetricsConfig kmetrics.CLIConfig
	PprofConfig   kpprof.CLIConfig
}

func (c CLIConfig) Check() error {
	if !(c.OutputSubmitterEnabled || c.ChallengerEnabled || c.GuardianEnabled) {
		return errors.New("one of output submitter, challenger, guardian should be enabled")
	}
	if c.OutputSubmitterAllowPublicRound && !c.OutputSubmitterEnabled {
		return errors.New("OutputSubmitterAllowPublicRound is meaningful when OutputSubmitterEnabled enabled")
	}
	if err := c.RPCConfig.Check(); err != nil {
		return err
	}
	if err := c.LogConfig.Check(); err != nil {
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

// NewCLIConfig parses the CLIConfig from the provided flags or environment variables.
func NewCLIConfig(ctx *cli.Context) CLIConfig {
	return CLIConfig{
		// Required Flags
		L1EthRpc:               ctx.GlobalString(flags.L1EthRpcFlag.Name),
		L2EthRpc:               ctx.GlobalString(flags.L2EthRpcFlag.Name),
		RollupRpc:              ctx.GlobalString(flags.RollupRpcFlag.Name),
		L2OOAddress:            ctx.GlobalString(flags.L2OOAddressFlag.Name),
		ColosseumAddress:       ctx.GlobalString(flags.ColosseumAddressFlag.Name),
		ValPoolAddress:         ctx.GlobalString(flags.ValPoolAddressFlag.Name),
		OutputSubmitterEnabled: ctx.GlobalBool(flags.OutputSubmitterEnabledFlag.Name),
		ChallengerEnabled:      ctx.GlobalBool(flags.ChallengerEnabledFlag.Name),
		ChallengerPollInterval: ctx.GlobalDuration(flags.ChallengerPollIntervalFlag.Name),
		TxMgrConfig:            txmgr.ReadCLIConfig(ctx),

		// Optional Flags
		AllowNonFinalized:               ctx.GlobalBool(flags.AllowNonFinalizedFlag.Name),
		OutputSubmitterRetryInterval:    ctx.GlobalDuration(flags.OutputSubmitterRetryIntervalFlag.Name),
		OutputSubmitterRoundBuffer:      ctx.GlobalUint64(flags.OutputSubmitterRoundBufferFlag.Name),
		OutputSubmitterAllowPublicRound: ctx.GlobalBool(flags.OutputSubmitterAllowPublicRoundFlag.Name),
		SecurityCouncilAddress:          ctx.GlobalString(flags.SecurityCouncilAddressFlag.Name),
		ProverRPC:                       ctx.GlobalString(flags.ProverRPCFlag.Name),
		GuardianEnabled:                 ctx.GlobalBool(flags.GuardianEnabledFlag.Name),
		FetchingProofTimeout:            ctx.GlobalDuration(flags.FetchingProofTimeoutFlag.Name),
		RPCConfig:                       krpc.ReadCLIConfig(ctx),
		LogConfig:                       klog.ReadCLIConfig(ctx),
		MetricsConfig:                   kmetrics.ReadCLIConfig(ctx),
		PprofConfig:                     kpprof.ReadCLIConfig(ctx),
	}
}

// NewValidatorConfig creates a validator config with given the CLIConfig
func NewValidatorConfig(cfg CLIConfig, l log.Logger, m metrics.Metricer) (*Config, error) {
	l2ooAddress, err := utils.ParseAddress(cfg.L2OOAddress)
	if err != nil {
		return nil, err
	}

	colosseumAddress, err := utils.ParseAddress(cfg.ColosseumAddress)
	if err != nil {
		return nil, err
	}

	var securityCouncilAddress common.Address
	if cfg.GuardianEnabled {
		securityCouncilAddress, err = utils.ParseAddress(cfg.SecurityCouncilAddress)
		if err != nil {
			return nil, err
		}
	}

	valPoolAddress, err := utils.ParseAddress(cfg.ValPoolAddress)
	if err != nil {
		return nil, err
	}

	txManager, err := txmgr.NewBufferedTxManager("validator", l, m, cfg.TxMgrConfig)
	if err != nil {
		return nil, err
	}

	if cfg.ChallengerEnabled && len(cfg.ProverRPC) == 0 {
		return nil, errors.New("ProverRPC is required when challenger enabled, but given empty")
	}

	var fetcher ProofFetcher
	if len(cfg.ProverRPC) > 0 {
		fetcher, err = chal.NewFetcher(cfg.ProverRPC, cfg.FetchingProofTimeout, l)
		if err != nil {
			return nil, err
		}
	}

	// Connect to L1 and L2 providers. Perform these last since they are the most expensive.
	ctx := context.Background()
	l1Client, err := utils.DialEthClientWithTimeout(ctx, cfg.L1EthRpc)
	if err != nil {
		return nil, err
	}

	l2Client, err := utils.DialEthClientWithTimeout(ctx, cfg.L2EthRpc)
	if err != nil {
		return nil, err
	}

	rollupClient, err := utils.DialRollupClientWithTimeout(ctx, cfg.RollupRpc)
	if err != nil {
		return nil, err
	}

	rollupConfig, err := rollupClient.RollupConfig(ctx)
	if err != nil {
		return nil, err
	}

	return &Config{
		L2OutputOracleAddr:              l2ooAddress,
		ColosseumAddr:                   colosseumAddress,
		SecurityCouncilAddr:             securityCouncilAddress,
		ValidatorPoolAddr:               valPoolAddress,
		ChallengerPollInterval:          cfg.ChallengerPollInterval,
		NetworkTimeout:                  cfg.TxMgrConfig.NetworkTimeout,
		TxManager:                       txManager,
		L1Client:                        l1Client,
		L2Client:                        l2Client,
		RollupClient:                    rollupClient,
		RollupConfig:                    rollupConfig,
		AllowNonFinalized:               cfg.AllowNonFinalized,
		OutputSubmitterEnabled:          cfg.OutputSubmitterEnabled,
		OutputSubmitterAllowPublicRound: cfg.OutputSubmitterAllowPublicRound,
		OutputSubmitterRetryInterval:    cfg.OutputSubmitterRetryInterval,
		OutputSubmitterRoundBuffer:      cfg.OutputSubmitterRoundBuffer,
		ChallengerEnabled:               cfg.ChallengerEnabled,
		GuardianEnabled:                 cfg.GuardianEnabled,
		ProofFetcher:                    fetcher,
	}, nil
}
