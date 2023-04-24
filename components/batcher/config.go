package batcher

import (
	"context"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/urfave/cli"

	"github.com/kroma-network/kroma/components/batcher/flags"
	"github.com/kroma-network/kroma/components/batcher/metrics"
	"github.com/kroma-network/kroma/components/batcher/rpc"
	"github.com/kroma-network/kroma/components/node/rollup"
	"github.com/kroma-network/kroma/components/node/sources"
	"github.com/kroma-network/kroma/utils"
	klog "github.com/kroma-network/kroma/utils/service/log"
	kmetrics "github.com/kroma-network/kroma/utils/service/metrics"
	kpprof "github.com/kroma-network/kroma/utils/service/pprof"
	"github.com/kroma-network/kroma/utils/service/txmgr"
)

type Config struct {
	log          log.Logger
	metr         metrics.Metricer
	L1Client     *ethclient.Client
	L2Client     *ethclient.Client
	RollupClient *sources.RollupClient
	TxManager    txmgr.TxManager

	NetworkTimeout time.Duration
	PollInterval   time.Duration

	// Rollup config is queried at startup
	Rollup *rollup.Config

	// Channel builder parameters
	Channel ChannelConfig
}

// Check ensures that the [Config] is valid.
func (c *Config) Check() error {
	if err := c.Rollup.Check(); err != nil {
		return err
	}
	if err := c.Channel.Check(); err != nil {
		return err
	}
	return nil
}

type CLIConfig struct {
	// L1EthRpc is the HTTP provider URL for L1.
	L1EthRpc string

	// L2EthRpc is the HTTP provider URL for the L2 execution engine.
	L2EthRpc string

	// RollupRpc is the HTTP provider URL for the L2 rollup node.
	RollupRpc string

	// MaxChannelDuration is the maximum duration (in #L1-blocks) to keep a
	// channel open. This allows to more eagerly send batcher transactions
	// during times of low L2 transaction volume. Note that the effective
	// L1-block distance between batcher transactions is then MaxChannelDuration
	//
	// If 0, duration checks are disabled.
	MaxChannelDuration uint64

	// The batcher tx submission safety margin (in #L1-blocks) to subtract from
	// a channel's timeout and proposing window, to guarantee safe inclusion of
	// a channel on L1.
	SubSafetyMargin uint64

	// PollInterval is the delay between querying L2 for more transaction
	// and creating a new batch.
	PollInterval time.Duration

	// MaxL1TxSize is the maximum size of a batch tx submitted to L1.
	MaxL1TxSize uint64

	// TargetL1TxSize is the target size of a batch tx submitted to L1.
	TargetL1TxSize uint64

	// TargetNumFrames is the target number of frames per channel.
	TargetNumFrames int

	// ApproxComprRatio is the approximate compression ratio (<= 1.0) of the used
	// compression algorithm.
	ApproxComprRatio float64

	TxMgrConfig   txmgr.CLIConfig
	RPCConfig     rpc.CLIConfig
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
		L1EthRpc:        ctx.GlobalString(flags.L1EthRpcFlag.Name),
		L2EthRpc:        ctx.GlobalString(flags.L2EthRpcFlag.Name),
		RollupRpc:       ctx.GlobalString(flags.RollupRpcFlag.Name),
		SubSafetyMargin: ctx.GlobalUint64(flags.SubSafetyMarginFlag.Name),
		PollInterval:    ctx.GlobalDuration(flags.PollIntervalFlag.Name),

		// Optional Flags
		MaxChannelDuration: ctx.GlobalUint64(flags.MaxChannelDurationFlag.Name),
		MaxL1TxSize:        ctx.GlobalUint64(flags.MaxL1TxSizeBytesFlag.Name),
		TargetL1TxSize:     ctx.GlobalUint64(flags.TargetL1TxSizeBytesFlag.Name),
		TargetNumFrames:    ctx.GlobalInt(flags.TargetNumFramesFlag.Name),
		ApproxComprRatio:   ctx.GlobalFloat64(flags.ApproxComprRatioFlag.Name),
		TxMgrConfig:        txmgr.ReadCLIConfig(ctx),
		RPCConfig:          rpc.ReadCLIConfig(ctx),
		LogConfig:          klog.ReadCLIConfig(ctx),
		MetricsConfig:      kmetrics.ReadCLIConfig(ctx),
		PprofConfig:        kpprof.ReadCLIConfig(ctx),
	}
}

// NewBatcherConfig creates a batcher config with given the CLIConfig
func NewBatcherConfig(cfg CLIConfig, l log.Logger, m metrics.Metricer) (*Config, error) {
	ctx := context.Background()

	// Connect to L1 and L2 providers. Perform these last since they are the most expensive.
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

	rcfg, err := rollupClient.RollupConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("querying rollup config: %w", err)
	}

	txManager, err := txmgr.NewSimpleTxManager("batcher", l, m, cfg.TxMgrConfig)
	if err != nil {
		return nil, err
	}

	return &Config{
		log:            l,
		metr:           m,
		L1Client:       l1Client,
		L2Client:       l2Client,
		RollupClient:   rollupClient,
		PollInterval:   cfg.PollInterval,
		NetworkTimeout: cfg.TxMgrConfig.NetworkTimeout,
		TxManager:      txManager,
		Rollup:         rcfg,
		Channel: ChannelConfig{
			ProposerWindowSize: rcfg.ProposerWindowSize,
			ChannelTimeout:     rcfg.ChannelTimeout,
			MaxChannelDuration: cfg.MaxChannelDuration,
			SubSafetyMargin:    cfg.SubSafetyMargin,
			MaxFrameSize:       cfg.MaxL1TxSize - 1,    // subtract 1 byte for version
			TargetFrameSize:    cfg.TargetL1TxSize - 1, // subtract 1 byte for version
			TargetNumFrames:    cfg.TargetNumFrames,
			ApproxComprRatio:   cfg.ApproxComprRatio,
		},
	}, nil
}
