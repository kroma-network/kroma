package flags

import (
	"time"

	"github.com/urfave/cli"

	kservice "github.com/wemixkanvas/kanvas/utils/service"
	klog "github.com/wemixkanvas/kanvas/utils/service/log"
	kmetrics "github.com/wemixkanvas/kanvas/utils/service/metrics"
	kpprof "github.com/wemixkanvas/kanvas/utils/service/pprof"
	krpc "github.com/wemixkanvas/kanvas/utils/service/rpc"
	ksigner "github.com/wemixkanvas/kanvas/utils/signer/client"
)

const envVarPrefix = "VALIDATOR"

var (
	/* Required Flags */

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
	ProverGrpcFlag = cli.StringFlag{
		Name:     "prover-grpc-url",
		Usage:    "gRPC URL for kanvas-prover.",
		Required: true,
		EnvVar:   kservice.PrefixEnvVar(envVarPrefix, "PROVER_GRPC"),
	}

	/* Optional flags */

	MnemonicFlag = cli.StringFlag{
		Name:   "mnemonic",
		Usage:  "The mnemonic used to derive the wallets for the validator",
		EnvVar: kservice.PrefixEnvVar(envVarPrefix, "MNEMONIC"),
	}
	HDPathFlag = cli.StringFlag{
		Name: "hd-path",
		Usage: "The HD path used to derive the validator from the " +
			"mnemonic. The mnemonic flag must also be set.",
		EnvVar: kservice.PrefixEnvVar(envVarPrefix, "HD_PATH"),
	}
	PrivateKeyFlag = cli.StringFlag{
		Name:   "private-key",
		Usage:  "The private key to use with the validator. Must not be used with mnemonic.",
		EnvVar: kservice.PrefixEnvVar(envVarPrefix, "PRIVATE_KEY"),
	}
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
	PollIntervalFlag,
	NumConfirmationsFlag,
	SafeAbortNonceTooLowCountFlag,
	ResubmissionTimeoutFlag,
	ProverGrpcFlag,
}

var optionalFlags = []cli.Flag{
	MnemonicFlag,
	HDPathFlag,
	PrivateKeyFlag,
	AllowNonFinalizedFlag,
	OutputSubmitterDisabledFlag,
	ChallengerDisabledFlag,
	FetchingProofTimeoutFlag,
}

func init() {
	requiredFlags = append(requiredFlags, krpc.CLIFlags(envVarPrefix)...)

	optionalFlags = append(optionalFlags, klog.CLIFlags(envVarPrefix)...)
	optionalFlags = append(optionalFlags, kmetrics.CLIFlags(envVarPrefix)...)
	optionalFlags = append(optionalFlags, kpprof.CLIFlags(envVarPrefix)...)
	optionalFlags = append(optionalFlags, ksigner.CLIFlags(envVarPrefix)...)

	Flags = append(requiredFlags, optionalFlags...)
}

// Flags contains the list of configuration options available to the binary.
var Flags []cli.Flag
