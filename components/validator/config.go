package validator

import (
	"context"
	"errors"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/urfave/cli"

	"github.com/kroma-network/kroma/components/node/rollup"
	"github.com/kroma-network/kroma/components/node/sources"
	chal "github.com/kroma-network/kroma/components/validator/challenge"
	"github.com/kroma-network/kroma/components/validator/flags"
	"github.com/kroma-network/kroma/components/validator/metrics"
	"github.com/kroma-network/kroma/utils"
	klog "github.com/kroma-network/kroma/utils/service/log"
	kmetrics "github.com/kroma-network/kroma/utils/service/metrics"
	kpprof "github.com/kroma-network/kroma/utils/service/pprof"
	krpc "github.com/kroma-network/kroma/utils/service/rpc"
	"github.com/kroma-network/kroma/utils/service/txmgr"
)

// Config contains the well typed fields that are used to initialize the output submitter.
// It is intended for programmatic use.
type Config struct {
	L2OutputOracleAddr      common.Address
	ColosseumAddr           common.Address
	PollInterval            time.Duration
	NetworkTimeout          time.Duration
	TxManager               *txmgr.SimpleTxManager
	L1Client                *ethclient.Client
	RollupClient            *sources.RollupClient
	RollupConfig            *rollup.Config
	AllowNonFinalized       bool
	OutputSubmitterDisabled bool
	ChallengerDisabled      bool
	ProofFetcher            ProofFetcher
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
	// L1EthRpc is the HTTP provider URL for L1.
	L1EthRpc string

	// RollupRpc is the HTTP provider URL for the rollup node.
	RollupRpc string

	// L2OOAddress is the L2OutputOracle contract address.
	L2OOAddress string

	// ColosseumAddress is the Colosseum contract address.
	ColosseumAddress string

	// PollInterval is how frequently to poll L2 for new finalized outputs.
	PollInterval time.Duration

	// ProverGrpc is the URL of prover grpc server.
	ProverGrpc string

	// AllowNonFinalized can be set to true to submit outputs
	// for L2 blocks derived from non-finalized L1 data.
	AllowNonFinalized bool

	OutputSubmitterDisabled bool

	ChallengerDisabled bool

	FetchingProofTimeout time.Duration

	TxMgrConfig   txmgr.CLIConfig
	RPCConfig     krpc.CLIConfig
	LogConfig     klog.CLIConfig
	MetricsConfig kmetrics.CLIConfig
	PprofConfig   kpprof.CLIConfig
}

func (c CLIConfig) Check() error {
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
		L1EthRpc:         ctx.GlobalString(flags.L1EthRpcFlag.Name),
		RollupRpc:        ctx.GlobalString(flags.RollupRpcFlag.Name),
		L2OOAddress:      ctx.GlobalString(flags.L2OOAddressFlag.Name),
		ColosseumAddress: ctx.GlobalString(flags.ColosseumAddressFlag.Name),
		PollInterval:     ctx.GlobalDuration(flags.PollIntervalFlag.Name),
		ProverGrpc:       ctx.GlobalString(flags.ProverGrpcFlag.Name),
		TxMgrConfig:      txmgr.ReadCLIConfig(ctx),

		// Optional Flags
		AllowNonFinalized:       ctx.GlobalBool(flags.AllowNonFinalizedFlag.Name),
		OutputSubmitterDisabled: ctx.GlobalBool(flags.OutputSubmitterDisabledFlag.Name),
		ChallengerDisabled:      ctx.GlobalBool(flags.ChallengerDisabledFlag.Name),
		FetchingProofTimeout:    ctx.GlobalDuration(flags.FetchingProofTimeoutFlag.Name),
		RPCConfig:               krpc.ReadCLIConfig(ctx),
		LogConfig:               klog.ReadCLIConfig(ctx),
		MetricsConfig:           kmetrics.ReadCLIConfig(ctx),
		PprofConfig:             kpprof.ReadCLIConfig(ctx),
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

	txManager, err := txmgr.NewSimpleTxManager("validator", l, m, cfg.TxMgrConfig)
	if err != nil {
		return nil, err
	}

	if !cfg.ChallengerDisabled && len(cfg.ProverGrpc) == 0 {
		return nil, errors.New("ProverGrpc is required but given empty")
	}

	var fetcher ProofFetcher
	if len(cfg.ProverGrpc) > 0 {
		fetcher, err = chal.NewFetcher(cfg.ProverGrpc, cfg.FetchingProofTimeout, l)
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

	rollupClient, err := utils.DialRollupClientWithTimeout(ctx, cfg.RollupRpc)
	if err != nil {
		return nil, err
	}

	rollupConfig, err := rollupClient.RollupConfig(ctx)
	if err != nil {
		return nil, err
	}

	return &Config{
		L2OutputOracleAddr:      l2ooAddress,
		ColosseumAddr:           colosseumAddress,
		PollInterval:            cfg.PollInterval,
		NetworkTimeout:          cfg.TxMgrConfig.NetworkTimeout,
		TxManager:               txManager,
		L1Client:                l1Client,
		RollupClient:            rollupClient,
		RollupConfig:            rollupConfig,
		AllowNonFinalized:       cfg.AllowNonFinalized,
		OutputSubmitterDisabled: cfg.OutputSubmitterDisabled,
		ChallengerDisabled:      cfg.ChallengerDisabled,
		ProofFetcher:            fetcher,
	}, nil
}
