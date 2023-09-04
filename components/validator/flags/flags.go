package flags

import (
	"time"

	"github.com/urfave/cli/v2"

	kservice "github.com/kroma-network/kroma/utils/service"
	klog "github.com/kroma-network/kroma/utils/service/log"
	kmetrics "github.com/kroma-network/kroma/utils/service/metrics"
	kpprof "github.com/kroma-network/kroma/utils/service/pprof"
	krpc "github.com/kroma-network/kroma/utils/service/rpc"
	"github.com/kroma-network/kroma/utils/service/txmgr"
)

const envVarPrefix = "VALIDATOR"

func prefixEnvVar(name string) []string {
	return kservice.PrefixEnvVar(envVarPrefix, name)
}

var (
	// Required Flags

	L1EthRpcFlag = &cli.StringFlag{
		Name:     "l1-eth-rpc",
		Usage:    "Websocket provider URL for L1",
		Required: true,
		EnvVars:  prefixEnvVar("L1_ETH_RPC"),
	}
	L2EthRpcFlag = &cli.StringFlag{
		Name:     "l2-eth-rpc",
		Usage:    "HTTP provider URL for L2",
		Required: true,
		EnvVars:  prefixEnvVar("L2_ETH_RPC"),
	}
	RollupRpcFlag = &cli.StringFlag{
		Name:     "rollup-rpc",
		Usage:    "HTTP provider URL for the rollup node",
		Required: true,
		EnvVars:  prefixEnvVar("ROLLUP_RPC"),
	}
	L2OOAddressFlag = &cli.StringFlag{
		Name:     "l2oo-address",
		Usage:    "Address of the L2OutputOracle contract",
		Required: true,
		EnvVars:  prefixEnvVar("L2OO_ADDRESS"),
	}
	ColosseumAddressFlag = &cli.StringFlag{
		Name:     "colosseum-address",
		Usage:    "Address of the Colosseum contract",
		Required: true,
		EnvVars:  prefixEnvVar("COLOSSEUM_ADDRESS"),
	}
	ValPoolAddressFlag = &cli.StringFlag{
		Name:     "valpool-address",
		Usage:    "Address of the ValidatorPool contract",
		Required: true,
		EnvVars:  prefixEnvVar("VALPOOL_ADDRESS"),
	}
	OutputSubmitterEnabledFlag = &cli.BoolFlag{
		Name:     "output-submitter.enabled",
		Usage:    "Enable l2 output submitter",
		EnvVars:  prefixEnvVar("OUTPUT_SUBMITTER_ENABLED"),
		Required: true,
	}
	ChallengerEnabledFlag = &cli.BoolFlag{
		Name:     "challenger.enabled",
		Usage:    "Enable challenger",
		EnvVars:  prefixEnvVar("CHALLENGER_ENABLED"),
		Required: true,
	}
	ChallengerPollIntervalFlag = &cli.DurationFlag{
		Name:     "challenger.poll-interval",
		Usage:    "Poll interval for challenge process",
		Required: true,
		EnvVars:  prefixEnvVar("CHALLENGER_POLL_INTERVAL"),
	}

	// Optional flags

	AllowNonFinalizedFlag = &cli.BoolFlag{
		Name:    "allow-non-finalized",
		Usage:   "Allow the validator to submit outputs for L2 blocks derived from non-finalized L1 blocks.",
		EnvVars: prefixEnvVar("ALLOW_NON_FINALIZED"),
	}
	OutputSubmitterRetryIntervalFlag = &cli.DurationFlag{
		Name:    "output-submitter.retry-interval",
		Usage:   "Retry interval for output submission process",
		EnvVars: prefixEnvVar("OUTPUT_SUBMITTER_RETRY_INTERVAL"),
		Value:   time.Second * 1,
	}
	OutputSubmitterRoundBufferFlag = &cli.Uint64Flag{
		Name:    "output-submitter.round-buffer",
		Usage:   "Number of blocks before each round to start trying submission",
		EnvVars: prefixEnvVar("OUTPUT_SUBMITTER_ROUND_BUFFER"),
		Value:   30,
	}
	OutputSubmitterAllowPublicRoundFlag = &cli.BoolFlag{
		Name:    "output-submitter.allow-public-round",
		Usage:   "Allows l2 output submitter in public round",
		EnvVars: prefixEnvVar("OUTPUT_SUBMITTER_ALLOW_PUBLIC_ROUND"),
	}
	ProverRPCFlag = &cli.StringFlag{
		Name:    "prover-rpc-url",
		Usage:   "jsonRPC URL for kroma-prover.",
		EnvVars: prefixEnvVar("PROVER_RPC"),
	}
	SecurityCouncilAddressFlag = &cli.StringFlag{
		Name:    "securitycouncil-address",
		Usage:   "Address of the SecurityCouncil contract",
		EnvVars: prefixEnvVar("SECURITYCOUNCIL_ADDRESS"),
	}
	GuardianEnabledFlag = &cli.BoolFlag{
		Name:    "guardian.enabled",
		Usage:   "Enable guardian",
		EnvVars: prefixEnvVar("GUARDIAN_ENABLED"),
	}
	FetchingProofTimeoutFlag = &cli.DurationFlag{
		Name:    "fetching-proof-timeout",
		Usage:   "Duration we will wait to fetching proof",
		EnvVars: prefixEnvVar("FETCHING_PROOF_TIMEOUT"),
		Value:   time.Hour * 4,
	}
)

var requiredFlags = []cli.Flag{
	L1EthRpcFlag,
	L2EthRpcFlag,
	RollupRpcFlag,
	L2OOAddressFlag,
	ColosseumAddressFlag,
	ValPoolAddressFlag,
	OutputSubmitterEnabledFlag,
	ChallengerEnabledFlag,
	ChallengerPollIntervalFlag,
}

var optionalFlags = []cli.Flag{
	AllowNonFinalizedFlag,
	OutputSubmitterRetryIntervalFlag,
	OutputSubmitterRoundBufferFlag,
	OutputSubmitterAllowPublicRoundFlag,
	ProverRPCFlag,
	SecurityCouncilAddressFlag,
	GuardianEnabledFlag,
	FetchingProofTimeoutFlag,
}

func init() {
	requiredFlags = append(requiredFlags, krpc.CLIFlags(envVarPrefix)...)

	optionalFlags = append(optionalFlags, klog.CLIFlags(envVarPrefix)...)
	optionalFlags = append(optionalFlags, kmetrics.CLIFlags(envVarPrefix)...)
	optionalFlags = append(optionalFlags, kpprof.CLIFlags(envVarPrefix)...)
	optionalFlags = append(optionalFlags, txmgr.CLIFlags(envVarPrefix)...)

	Flags = append(requiredFlags, optionalFlags...)
}

// Flags contains the list of configuration options available to the binary.
var Flags []cli.Flag
