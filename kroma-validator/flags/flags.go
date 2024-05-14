package flags

import (
	"fmt"
	"time"

	opservice "github.com/ethereum-optimism/optimism/op-service"
	oplog "github.com/ethereum-optimism/optimism/op-service/log"
	opmetrics "github.com/ethereum-optimism/optimism/op-service/metrics"
	oppprof "github.com/ethereum-optimism/optimism/op-service/oppprof"
	oprpc "github.com/ethereum-optimism/optimism/op-service/rpc"
	"github.com/ethereum-optimism/optimism/op-service/txmgr"
	"github.com/urfave/cli/v2"
)

const EnvVarPrefix = "VALIDATOR"

func prefixEnvVars(name string) []string {
	return opservice.PrefixEnvVar(EnvVarPrefix, name)
}

var (
	// Required Flags

	L1EthRpcFlag = &cli.StringFlag{
		Name:     "l1-eth-rpc",
		Usage:    "Websocket provider URL for L1",
		Required: true,
		EnvVars:  prefixEnvVars("L1_ETH_RPC"),
	}
	L2EthRpcFlag = &cli.StringFlag{
		Name:     "l2-eth-rpc",
		Usage:    "HTTP provider URL for L2",
		Required: true,
		EnvVars:  prefixEnvVars("L2_ETH_RPC"),
	}
	RollupRpcFlag = &cli.StringFlag{
		Name:     "rollup-rpc",
		Usage:    "HTTP provider URL for the rollup node",
		Required: true,
		EnvVars:  prefixEnvVars("ROLLUP_RPC"),
	}
	L2OOAddressFlag = &cli.StringFlag{
		Name:     "l2oo-address",
		Usage:    "Address of the L2OutputOracle contract",
		Required: true,
		EnvVars:  prefixEnvVars("L2OO_ADDRESS"),
	}
	ColosseumAddressFlag = &cli.StringFlag{
		Name:     "colosseum-address",
		Usage:    "Address of the Colosseum contract",
		Required: true,
		EnvVars:  prefixEnvVars("COLOSSEUM_ADDRESS"),
	}
	ValPoolAddressFlag = &cli.StringFlag{
		Name:     "valpool-address",
		Usage:    "Address of the ValidatorPool contract",
		Required: true,
		EnvVars:  prefixEnvVars("VALPOOL_ADDRESS"),
	}
	ValMgrAddressFlag = &cli.StringFlag{
		Name:     "valmgr-address",
		Usage:    "Address of the ValidatorManager contract",
		Required: true,
		EnvVars:  prefixEnvVars("VALMGR_ADDRESS"),
	}
	AssetManagerAddressFlag = &cli.StringFlag{
		Name:     "assetmanager-address",
		Usage:    "Address of the AssetManager contract",
		Required: true,
		EnvVars:  prefixEnvVars("ASSETMANAGER_ADDRESS"),
	}
	OutputSubmitterEnabledFlag = &cli.BoolFlag{
		Name:     "output-submitter.enabled",
		Usage:    "Enable l2 output submitter",
		EnvVars:  prefixEnvVars("OUTPUT_SUBMITTER_ENABLED"),
		Required: true,
	}
	ChallengerEnabledFlag = &cli.BoolFlag{
		Name:     "challenger.enabled",
		Usage:    "Enable challenger",
		EnvVars:  prefixEnvVars("CHALLENGER_ENABLED"),
		Required: true,
	}
	ChallengerPollIntervalFlag = &cli.DurationFlag{
		Name:     "challenger.poll-interval",
		Usage:    "Poll interval for challenge process",
		Required: true,
		EnvVars:  prefixEnvVars("CHALLENGER_POLL_INTERVAL"),
	}

	// Optional flags

	AllowNonFinalizedFlag = &cli.BoolFlag{
		Name:    "allow-non-finalized",
		Usage:   "Allow the validator to submit outputs for L2 blocks derived from non-finalized L1 blocks.",
		EnvVars: prefixEnvVars("ALLOW_NON_FINALIZED"),
	}
	OutputSubmitterRetryIntervalFlag = &cli.DurationFlag{
		Name:    "output-submitter.retry-interval",
		Usage:   "Retry interval for output submission process",
		EnvVars: prefixEnvVars("OUTPUT_SUBMITTER_RETRY_INTERVAL"),
		Value:   time.Second * 1,
	}
	OutputSubmitterRoundBufferFlag = &cli.Uint64Flag{
		Name:    "output-submitter.round-buffer",
		Usage:   "Number of blocks before each round to start trying submission",
		EnvVars: prefixEnvVars("OUTPUT_SUBMITTER_ROUND_BUFFER"),
		Value:   30,
	}
	OutputSubmitterAllowPublicRoundFlag = &cli.BoolFlag{
		Name:    "output-submitter.allow-public-round",
		Usage:   "Allows l2 output submitter in public round",
		EnvVars: prefixEnvVars("OUTPUT_SUBMITTER_ALLOW_PUBLIC_ROUND"),
	}
	ProverRPCFlag = &cli.StringFlag{
		Name:    "prover-rpc-url",
		Usage:   "jsonRPC URL for kroma-prover.",
		EnvVars: prefixEnvVars("PROVER_RPC"),
	}
	SecurityCouncilAddressFlag = &cli.StringFlag{
		Name:    "securitycouncil-address",
		Usage:   "Address of the SecurityCouncil contract",
		EnvVars: prefixEnvVars("SECURITYCOUNCIL_ADDRESS"),
	}
	GuardianEnabledFlag = &cli.BoolFlag{
		Name:    "guardian.enabled",
		Usage:   "Enable guardian",
		EnvVars: prefixEnvVars("GUARDIAN_ENABLED"),
	}
	GuardianPollIntervalFlag = &cli.DurationFlag{
		Name:    "guardian.poll-interval",
		Usage:   "Poll interval for guardian inspection",
		EnvVars: prefixEnvVars("GUARDIAN_POLL_INTERVAL"),
		Value:   time.Minute,
	}
	FetchingProofTimeoutFlag = &cli.DurationFlag{
		Name:    "fetching-proof-timeout",
		Usage:   "Duration we will wait to fetching proof",
		EnvVars: prefixEnvVars("FETCHING_PROOF_TIMEOUT"),
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
	ValMgrAddressFlag,
	AssetManagerAddressFlag,
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
	GuardianPollIntervalFlag,
	FetchingProofTimeoutFlag,
}

func init() {
	optionalFlags = append(optionalFlags, oprpc.CLIFlags(EnvVarPrefix)...)
	optionalFlags = append(optionalFlags, oplog.CLIFlags(EnvVarPrefix)...)
	optionalFlags = append(optionalFlags, opmetrics.CLIFlags(EnvVarPrefix)...)
	optionalFlags = append(optionalFlags, oppprof.CLIFlags(EnvVarPrefix)...)
	optionalFlags = append(optionalFlags, txmgr.CLIFlags(EnvVarPrefix)...)

	Flags = append(requiredFlags, optionalFlags...)
}

// Flags contains the list of configuration options available to the binary.
var Flags []cli.Flag

func CheckRequired(ctx *cli.Context) error {
	for _, f := range requiredFlags {
		if !ctx.IsSet(f.Names()[0]) {
			return fmt.Errorf("flag %s is required", f.Names()[0])
		}
	}
	return nil
}
