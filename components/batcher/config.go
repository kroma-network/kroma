package batcher

import (
	"context"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/urfave/cli"

	"github.com/wemixkanvas/kanvas/components/batcher/flags"
	"github.com/wemixkanvas/kanvas/components/batcher/metrics"
	"github.com/wemixkanvas/kanvas/components/batcher/rpc"
	"github.com/wemixkanvas/kanvas/components/node/rollup"
	"github.com/wemixkanvas/kanvas/components/node/sources"
	"github.com/wemixkanvas/kanvas/utils"
	kcrypto "github.com/wemixkanvas/kanvas/utils/service/crypto"
	klog "github.com/wemixkanvas/kanvas/utils/service/log"
	kmetrics "github.com/wemixkanvas/kanvas/utils/service/metrics"
	kpprof "github.com/wemixkanvas/kanvas/utils/service/pprof"
	"github.com/wemixkanvas/kanvas/utils/service/txmgr"
	ksigner "github.com/wemixkanvas/kanvas/utils/signer/client"
)

type Config struct {
	log          log.Logger
	metr         metrics.Metricer
	L1Client     *ethclient.Client
	L2Client     *ethclient.Client
	RollupClient *sources.RollupClient

	PollInterval time.Duration
	From         common.Address

	TxManagerConfig txmgr.Config

	// RollupConfig is queried at startup
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
	/* Required Params */

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
	// + NumConfirmations because the batcher waits for NumConfirmations blocks
	// after sending a batcher tx and only then starts a new channel.
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

	// NumConfirmations is the number of confirmations which we will wait after
	// appending new batches.
	NumConfirmations uint64

	// SafeAbortNonceTooLowCount is the number of ErrNonceTooLowObservations
	// required to give up on a tx at a particular nonce without receiving
	// confirmation.
	SafeAbortNonceTooLowCount uint64

	// ResubmissionTimeout is time we will wait before resubmitting a
	// transaction.
	ResubmissionTimeout time.Duration

	// Mnemonic is the HD seed used to derive the wallet private keys for
	// the batcher.
	Mnemonic string

	// HDPath is the derivation path used to obtain the private key for
	// the batcher.
	HDPath string

	// PrivateKey is the private key used for the batcher.
	PrivateKey string

	RPCConfig rpc.CLIConfig

	/* Optional Params */

	// MaxL1TxSize is the maximum size of a batch tx submitted to L1.
	MaxL1TxSize uint64

	// TargetL1TxSize is the target size of a batch tx submitted to L1.
	TargetL1TxSize uint64

	// TargetNumFrames is the target number of frames per channel.
	TargetNumFrames int

	// ApproxComprRatio is the approximate compression ratio (<= 1.0) of the used
	// compression algorithm.
	ApproxComprRatio float64

	LogConfig klog.CLIConfig

	MetricsConfig kmetrics.CLIConfig

	PprofConfig kpprof.CLIConfig

	// SignerConfig contains the client config for signer service
	SignerConfig ksigner.CLIConfig
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
	if err := c.SignerConfig.Check(); err != nil {
		return err
	}
	return nil
}

// NewCLIConfig parses the CLIConfig from the provided flags or environment variables.
func NewCLIConfig(ctx *cli.Context) CLIConfig {
	return CLIConfig{
		/* Required Flags */
		L1EthRpc:                  ctx.GlobalString(flags.L1EthRpcFlag.Name),
		L2EthRpc:                  ctx.GlobalString(flags.L2EthRpcFlag.Name),
		RollupRpc:                 ctx.GlobalString(flags.RollupRpcFlag.Name),
		SubSafetyMargin:           ctx.GlobalUint64(flags.SubSafetyMarginFlag.Name),
		PollInterval:              ctx.GlobalDuration(flags.PollIntervalFlag.Name),
		NumConfirmations:          ctx.GlobalUint64(flags.NumConfirmationsFlag.Name),
		SafeAbortNonceTooLowCount: ctx.GlobalUint64(flags.SafeAbortNonceTooLowCountFlag.Name),
		ResubmissionTimeout:       ctx.GlobalDuration(flags.ResubmissionTimeoutFlag.Name),

		/* Optional Flags */
		MaxChannelDuration: ctx.GlobalUint64(flags.MaxChannelDurationFlag.Name),
		MaxL1TxSize:        ctx.GlobalUint64(flags.MaxL1TxSizeBytesFlag.Name),
		TargetL1TxSize:     ctx.GlobalUint64(flags.TargetL1TxSizeBytesFlag.Name),
		TargetNumFrames:    ctx.GlobalInt(flags.TargetNumFramesFlag.Name),
		ApproxComprRatio:   ctx.GlobalFloat64(flags.ApproxComprRatioFlag.Name),
		Mnemonic:           ctx.GlobalString(flags.MnemonicFlag.Name),
		HDPath:             ctx.GlobalString(flags.HDPathFlag.Name),
		PrivateKey:         ctx.GlobalString(flags.PrivateKeyFlag.Name),
		RPCConfig:          rpc.ReadCLIConfig(ctx),
		LogConfig:          klog.ReadCLIConfig(ctx),
		MetricsConfig:      kmetrics.ReadCLIConfig(ctx),
		PprofConfig:        kpprof.ReadCLIConfig(ctx),
		SignerConfig:       ksigner.ReadCLIConfig(ctx),
	}
}

// NewBatcherConfig creates a batcher config with given the CLIConfig
func NewBatcherConfig(cfg CLIConfig, l log.Logger, m metrics.Metricer) (*Config, error) {
	signer, fromAddress, err := kcrypto.SignerFactoryFromConfig(l, cfg.PrivateKey, cfg.Mnemonic, cfg.HDPath, cfg.SignerConfig)
	if err != nil {
		return nil, err
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

	rcfg, err := rollupClient.RollupConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("querying rollup config: %w", err)
	}

	txMgrCfg := txmgr.Config{
		ResubmissionTimeout:       cfg.ResubmissionTimeout,
		ReceiptQueryInterval:      time.Second,
		NumConfirmations:          cfg.NumConfirmations,
		SafeAbortNonceTooLowCount: cfg.SafeAbortNonceTooLowCount,
		From:                      fromAddress,
		Signer:                    signer(rcfg.L1ChainID),
	}

	return &Config{
		log:             l,
		metr:            m,
		L1Client:        l1Client,
		L2Client:        l2Client,
		RollupClient:    rollupClient,
		PollInterval:    cfg.PollInterval,
		TxManagerConfig: txMgrCfg,
		From:            fromAddress,
		Rollup:          rcfg,
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
