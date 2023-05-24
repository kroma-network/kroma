package flags

import (
	"time"

	"github.com/urfave/cli"

	kservice "github.com/kroma-network/kroma/utils/service"
	klog "github.com/kroma-network/kroma/utils/service/log"
	kmetrics "github.com/kroma-network/kroma/utils/service/metrics"
	kpprof "github.com/kroma-network/kroma/utils/service/pprof"
	krpc "github.com/kroma-network/kroma/utils/service/rpc"
	"github.com/kroma-network/kroma/utils/service/txmgr"
)

const envVarPrefix = "VALIDATOR"

var (
	// Required Flags

	L1EthRpcFlag = cli.StringFlag{
		Name:     "l1-eth-rpc",
		Usage:    "HTTP provider URL for L1",
		Required: true,
		EnvVar:   kservice.PrefixEnvVar(envVarPrefix, "L1_ETH_RPC"),
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
	ChallengerPollIntervalFlag = cli.DurationFlag{
		Name:     "challenger.poll-interval",
		Usage:    "Poll interval for challenge process",
		Required: true,
		EnvVar:   kservice.PrefixEnvVar(envVarPrefix, "CHALLENGER_POLL_INTERVAL"),
	}
	ProverGrpcFlag = cli.StringFlag{
		Name:     "prover-grpc-url",
		Usage:    "gRPC URL for kroma-prover.",
		Required: false,
		EnvVar:   kservice.PrefixEnvVar(envVarPrefix, "PROVER_GRPC"),
	}

	// Optional flags

	AllowNonFinalizedFlag = cli.BoolFlag{
		Name:   "allow-non-finalized",
		Usage:  "Allow the validator to submit outputs for L2 blocks derived from non-finalized L1 blocks.",
		EnvVar: kservice.PrefixEnvVar(envVarPrefix, "ALLOW_NON_FINALIZED"),
	}
	OutputSubmitterDisabledFlag = cli.BoolFlag{
		Name:   "output-submitter.disabled",
		Usage:  "Disable l2 output submitter",
		EnvVar: kservice.PrefixEnvVar(envVarPrefix, "OUTPUT_SUBMITTER_DISABLED"),
	}
	OutputSubmitterBondAmountFlag = cli.Uint64Flag{
		Name:   "output-submitter.bond-amount",
		Usage:  "Amount to bond when submitting each output (in wei)",
		EnvVar: kservice.PrefixEnvVar(envVarPrefix, "OUTPUT_SUBMITTER_BOND_AMOUNT"),
		Value:  1,
	}
	ChallengerDisabledFlag = cli.BoolFlag{
		Name:   "challenger.disabled",
		Usage:  "Disable challenger",
		EnvVar: kservice.PrefixEnvVar(envVarPrefix, "CHALLENGER_DISABLED"),
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
	RollupRpcFlag,
	L2OOAddressFlag,
	ColosseumAddressFlag,
	ValPoolAddressFlag,
	ChallengerPollIntervalFlag,
	ProverGrpcFlag,
}

var optionalFlags = []cli.Flag{
	AllowNonFinalizedFlag,
	OutputSubmitterDisabledFlag,
	OutputSubmitterBondAmountFlag,
	ChallengerDisabledFlag,
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
