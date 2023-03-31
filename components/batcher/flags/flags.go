package flags

import (
	"github.com/urfave/cli"

	"github.com/wemixkanvas/kanvas/components/batcher/rpc"
	kservice "github.com/wemixkanvas/kanvas/utils/service"
	klog "github.com/wemixkanvas/kanvas/utils/service/log"
	kmetrics "github.com/wemixkanvas/kanvas/utils/service/metrics"
	kpprof "github.com/wemixkanvas/kanvas/utils/service/pprof"
	krpc "github.com/wemixkanvas/kanvas/utils/service/rpc"
	ksigner "github.com/wemixkanvas/kanvas/utils/signer/client"
)

const envVarPrefix = "BATCHER"

var (
	/* Required flags */

	L1EthRpcFlag = cli.StringFlag{
		Name:     "l1-eth-rpc",
		Usage:    "HTTP provider URL for L1",
		Required: true,
		EnvVar:   kservice.PrefixEnvVar(envVarPrefix, "L1_ETH_RPC"),
	}
	L2EthRpcFlag = cli.StringFlag{
		Name:     "l2-eth-rpc",
		Usage:    "HTTP provider URL for L2 execution engine",
		Required: true,
		EnvVar:   kservice.PrefixEnvVar(envVarPrefix, "L2_ETH_RPC"),
	}
	RollupRpcFlag = cli.StringFlag{
		Name:     "rollup-rpc",
		Usage:    "HTTP provider URL for Rollup node",
		Required: true,
		EnvVar:   kservice.PrefixEnvVar(envVarPrefix, "ROLLUP_RPC"),
	}
	SubSafetyMarginFlag = cli.Uint64Flag{
		Name: "sub-safety-margin",
		Usage: "The batcher tx submission safety margin (in #L1-blocks) to subtract " +
			"from a channel's timeout and proposing window, to guarantee safe inclusion " +
			"of a channel on L1.",
		Required: true,
		EnvVar:   kservice.PrefixEnvVar(envVarPrefix, "SUB_SAFETY_MARGIN"),
	}
	PollIntervalFlag = cli.DurationFlag{
		Name: "poll-interval",
		Usage: "Delay between querying L2 for more transactions and " +
			"creating a new batch",
		Required: true,
		EnvVar:   kservice.PrefixEnvVar(envVarPrefix, "POLL_INTERVAL"),
	}
	NumConfirmationsFlag = cli.Uint64Flag{
		Name: "num-confirmations",
		Usage: "Number of confirmations which we will wait after " +
			"appending a new batch",
		Required: true,
		EnvVar:   kservice.PrefixEnvVar(envVarPrefix, "NUM_CONFIRMATIONS"),
	}
	SafeAbortNonceTooLowCountFlag = cli.Uint64Flag{
		Name: "safe-abort-nonce-too-low-count",
		Usage: "Number of ErrNonceTooLow observations required to " +
			"give up on a tx at a particular nonce without receiving " +
			"confirmation",
		Required: true,
		EnvVar:   kservice.PrefixEnvVar(envVarPrefix, "SAFE_ABORT_NONCE_TOO_LOW_COUNT"),
	}
	ResubmissionTimeoutFlag = cli.DurationFlag{
		Name: "resubmission-timeout",
		Usage: "Duration we will wait before resubmitting a " +
			"transaction to L1",
		Required: true,
		EnvVar:   kservice.PrefixEnvVar(envVarPrefix, "RESUBMISSION_TIMEOUT"),
	}
	BatchInboxAddressFlag = cli.StringFlag{
		Name:     "proposer-batch-inbox-address",
		Usage:    "L1 Address to receive batch transactions",
		Required: true,
		EnvVar:   kservice.PrefixEnvVar(envVarPrefix, "BATCH_INBOX_ADDRESS"),
	}

	/* Optional flags */

	MaxChannelDurationFlag = cli.Uint64Flag{
		Name:   "max-channel-duration",
		Usage:  "The maximum duration of L1-blocks to keep a channel open. 0 to disable.",
		Value:  0,
		EnvVar: kservice.PrefixEnvVar(envVarPrefix, "MAX_CHANNEL_DURATION"),
	}
	MaxL1TxSizeBytesFlag = cli.Uint64Flag{
		Name:   "max-l1-tx-size-bytes",
		Usage:  "The maximum size of a batch tx submitted to L1.",
		Value:  120_000,
		EnvVar: kservice.PrefixEnvVar(envVarPrefix, "MAX_L1_TX_SIZE_BYTES"),
	}
	TargetL1TxSizeBytesFlag = cli.Uint64Flag{
		Name:   "target-l1-tx-size-bytes",
		Usage:  "The target size of a batch tx submitted to L1.",
		Value:  100_000,
		EnvVar: kservice.PrefixEnvVar(envVarPrefix, "TARGET_L1_TX_SIZE_BYTES"),
	}
	TargetNumFramesFlag = cli.IntFlag{
		Name:   "target-num-frames",
		Usage:  "The target number of frames to create per channel",
		Value:  1,
		EnvVar: kservice.PrefixEnvVar(envVarPrefix, "TARGET_NUM_FRAMES"),
	}
	ApproxComprRatioFlag = cli.Float64Flag{
		Name:   "approx-compr-ratio",
		Usage:  "The approximate compression ratio (<= 1.0)",
		Value:  1.0,
		EnvVar: kservice.PrefixEnvVar(envVarPrefix, "APPROX_COMPR_RATIO"),
	}
	MnemonicFlag = cli.StringFlag{
		Name:   "mnemonic",
		Usage:  "The mnemonic used to derive the wallets for the batcher",
		EnvVar: kservice.PrefixEnvVar(envVarPrefix, "MNEMONIC"),
	}
	HDPathFlag = cli.StringFlag{
		Name: "hd-path",
		Usage: "The HD path used to derive the batcher from the " +
			"mnemonic. The mnemonic flag must also be set.",
		EnvVar: kservice.PrefixEnvVar(envVarPrefix, "HD_PATH"),
	}
	PrivateKeyFlag = cli.StringFlag{
		Name:   "private-key",
		Usage:  "The private key to use with the batcher. Must not be used with mnemonic.",
		EnvVar: kservice.PrefixEnvVar(envVarPrefix, "PRIVATE_KEY"),
	}
)

var requiredFlags = []cli.Flag{
	L1EthRpcFlag,
	L2EthRpcFlag,
	RollupRpcFlag,
	SubSafetyMarginFlag,
	PollIntervalFlag,
	NumConfirmationsFlag,
	SafeAbortNonceTooLowCountFlag,
	ResubmissionTimeoutFlag,
	BatchInboxAddressFlag,
}

var optionalFlags = []cli.Flag{
	MaxChannelDurationFlag,
	MaxL1TxSizeBytesFlag,
	TargetL1TxSizeBytesFlag,
	TargetNumFramesFlag,
	ApproxComprRatioFlag,
	MnemonicFlag,
	HDPathFlag,
	PrivateKeyFlag,
}

func init() {
	requiredFlags = append(requiredFlags, krpc.CLIFlags(envVarPrefix)...)

	optionalFlags = append(optionalFlags, klog.CLIFlags(envVarPrefix)...)
	optionalFlags = append(optionalFlags, kmetrics.CLIFlags(envVarPrefix)...)
	optionalFlags = append(optionalFlags, kpprof.CLIFlags(envVarPrefix)...)
	optionalFlags = append(optionalFlags, ksigner.CLIFlags(envVarPrefix)...)
	optionalFlags = append(optionalFlags, rpc.CLIFlags(envVarPrefix)...)

	Flags = append(requiredFlags, optionalFlags...)
}

// Flags contains the list of configuration options available to the binary.
var Flags []cli.Flag
