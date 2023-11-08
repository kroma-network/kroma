package flags

import (
	"time"

	kservice "github.com/ethereum-optimism/optimism/op-service"
	klog "github.com/ethereum-optimism/optimism/op-service/log"
	kmetrics "github.com/ethereum-optimism/optimism/op-service/metrics"
	kpprof "github.com/ethereum-optimism/optimism/op-service/pprof"
	krpc "github.com/ethereum-optimism/optimism/op-service/rpc"
	"github.com/ethereum-optimism/optimism/op-service/txmgr"
	"github.com/urfave/cli"
)

const envVarPrefix = "VALIDATOR"

var (
	// Required Flags

	L1EthRpcFlag = cli.StringFlag{
		Name:     "l1-eth-rpc",
		Usage:    "Websocket provider URL for L1",
		Required: true,
		EnvVar:   kservice.PrefixEnvVar(envVarPrefix, "L1_ETH_RPC"),
	}
	L2EthRpcFlag = cli.StringFlag{
		Name:     "l2-eth-rpc",
		Usage:    "HTTP provider URL for L2",
		Required: true,
		EnvVar:   kservice.PrefixEnvVar(envVarPrefix, "L2_ETH_RPC"),
	}
	RollupRpcFlag = cli.StringFlag{
		Name:     "rollup-rpc",
		Usage:    "HTTP provider URL for the rollup node",
		Required: true,
		EnvVar:   kservice.PrefixEnvVar(envVarPrefix, "ROLLUP_RPC"),
	}
	L2OOAddressFlag = cli.StringFlag{
		Name:     "l2oo-address",
		Usage:    "Address of the L2OutputOracle contract",
		Required: true,
		EnvVar:   kservice.PrefixEnvVar(envVarPrefix, "L2OO_ADDRESS"),
	}
	ColosseumAddressFlag = cli.StringFlag{
		Name:     "colosseum-address",
		Usage:    "Address of the Colosseum contract",
		Required: true,
		EnvVar:   kservice.PrefixEnvVar(envVarPrefix, "COLOSSEUM_ADDRESS"),
	}
	ValPoolAddressFlag = cli.StringFlag{
		Name:     "valpool-address",
		Usage:    "Address of the ValidatorPool contract",
		Required: true,
		EnvVar:   kservice.PrefixEnvVar(envVarPrefix, "VALPOOL_ADDRESS"),
	}
	OutputSubmitterEnabledFlag = cli.BoolFlag{
		Name:     "output-submitter.enabled",
		Usage:    "Enable l2 output submitter",
		EnvVar:   kservice.PrefixEnvVar(envVarPrefix, "OUTPUT_SUBMITTER_ENABLED"),
		Required: true,
	}
	ChallengerEnabledFlag = cli.BoolFlag{
		Name:     "challenger.enabled",
		Usage:    "Enable challenger",
		EnvVar:   kservice.PrefixEnvVar(envVarPrefix, "CHALLENGER_ENABLED"),
		Required: true,
	}
	ChallengerPollIntervalFlag = cli.DurationFlag{
		Name:     "challenger.poll-interval",
		Usage:    "Poll interval for challenge process",
		Required: true,
		EnvVar:   kservice.PrefixEnvVar(envVarPrefix, "CHALLENGER_POLL_INTERVAL"),
	}

	// Optional flags

	AllowNonFinalizedFlag = cli.BoolFlag{
		Name:   "allow-non-finalized",
		Usage:  "Allow the validator to submit outputs for L2 blocks derived from non-finalized L1 blocks.",
		EnvVar: kservice.PrefixEnvVar(envVarPrefix, "ALLOW_NON_FINALIZED"),
	}
	OutputSubmitterRetryIntervalFlag = cli.DurationFlag{
		Name:   "output-submitter.retry-interval",
		Usage:  "Retry interval for output submission process",
		EnvVar: kservice.PrefixEnvVar(envVarPrefix, "OUTPUT_SUBMITTER_RETRY_INTERVAL"),
		Value:  time.Second * 1,
	}
	OutputSubmitterRoundBufferFlag = cli.Uint64Flag{
		Name:   "output-submitter.round-buffer",
		Usage:  "Number of blocks before each round to start trying submission",
		EnvVar: kservice.PrefixEnvVar(envVarPrefix, "OUTPUT_SUBMITTER_ROUND_BUFFER"),
		Value:  30,
	}
	OutputSubmitterAllowPublicRoundFlag = cli.BoolFlag{
		Name:   "output-submitter.allow-public-round",
		Usage:  "Allows l2 output submitter in public round",
		EnvVar: kservice.PrefixEnvVar(envVarPrefix, "OUTPUT_SUBMITTER_ALLOW_PUBLIC_ROUND"),
	}
	ProverRPCFlag = cli.StringFlag{
		Name:   "prover-rpc-url",
		Usage:  "jsonRPC URL for kroma-prover.",
		EnvVar: kservice.PrefixEnvVar(envVarPrefix, "PROVER_RPC"),
	}
	SecurityCouncilAddressFlag = cli.StringFlag{
		Name:   "securitycouncil-address",
		Usage:  "Address of the SecurityCouncil contract",
		EnvVar: kservice.PrefixEnvVar(envVarPrefix, "SECURITYCOUNCIL_ADDRESS"),
	}
	GuardianEnabledFlag = cli.BoolFlag{
		Name:   "guardian.enabled",
		Usage:  "Enable guardian",
		EnvVar: kservice.PrefixEnvVar(envVarPrefix, "GUARDIAN_ENABLED"),
	}
	FetchingProofTimeoutFlag = cli.DurationFlag{
		Name:   "fetching-proof-timeout",
		Usage:  "Duration we will wait to fetching proof",
		EnvVar: kservice.PrefixEnvVar(envVarPrefix, "FETCHING_PROOF_TIMEOUT"),
		Value:  time.Hour * 2,
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
