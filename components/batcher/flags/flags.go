package flags

import (
	"github.com/urfave/cli"

	"github.com/kroma-network/kroma/components/batcher/rpc"
	kservice "github.com/kroma-network/kroma/utils/service"
	klog "github.com/kroma-network/kroma/utils/service/log"
	kmetrics "github.com/kroma-network/kroma/utils/service/metrics"
	kpprof "github.com/kroma-network/kroma/utils/service/pprof"
	krpc "github.com/kroma-network/kroma/utils/service/rpc"
	"github.com/kroma-network/kroma/utils/service/txmgr"
)

const envVarPrefix = "BATCHER"

var (
	// Required flags

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

	// Optional flags

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
	// Legacy Flags
	HDPathFlag = txmgr.BatcherHDPathFlag
)

var requiredFlags = []cli.Flag{
	L1EthRpcFlag,
	L2EthRpcFlag,
	RollupRpcFlag,
	SubSafetyMarginFlag,
	PollIntervalFlag,
}

var optionalFlags = []cli.Flag{
	MaxChannelDurationFlag,
	MaxL1TxSizeBytesFlag,
	TargetL1TxSizeBytesFlag,
	TargetNumFramesFlag,
	ApproxComprRatioFlag,
}

func init() {
	requiredFlags = append(requiredFlags, krpc.CLIFlags(envVarPrefix)...)

	optionalFlags = append(optionalFlags, klog.CLIFlags(envVarPrefix)...)
	optionalFlags = append(optionalFlags, kmetrics.CLIFlags(envVarPrefix)...)
	optionalFlags = append(optionalFlags, kpprof.CLIFlags(envVarPrefix)...)
	optionalFlags = append(optionalFlags, rpc.CLIFlags(envVarPrefix)...)
	optionalFlags = append(optionalFlags, txmgr.CLIFlags(envVarPrefix)...)

	Flags = append(requiredFlags, optionalFlags...)
}

// Flags contains the list of configuration options available to the binary.
var Flags []cli.Flag
