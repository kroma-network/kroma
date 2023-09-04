package flags

import (
	"fmt"
	"strings"
	"time"

	"github.com/urfave/cli/v2"

	"github.com/kroma-network/kroma/components/node/chaincfg"
	"github.com/kroma-network/kroma/components/node/sources"
	klog "github.com/kroma-network/kroma/utils/service/log"
)

// Flags

const EnvVarPrefix = "NODE"

func prefixEnvVars(name string) []string {
	return []string{EnvVarPrefix + "_" + name}
}

var (
	/* Required Flags */
	L1NodeAddr = &cli.StringFlag{
		Name:    "l1",
		Usage:   "Address of L1 User JSON-RPC endpoint to use (eth namespace required)",
		Value:   "http://127.0.0.1:8545",
		EnvVars: prefixEnvVars("L1_ETH_RPC"),
	}
	L2EngineAddr = &cli.StringFlag{
		Name:    "l2",
		Usage:   "Address of L2 Engine JSON-RPC endpoints to use (engine and eth namespace required)",
		EnvVars: prefixEnvVars("L2_ENGINE_RPC"),
	}
	RollupConfig = &cli.StringFlag{
		Name:    "rollup.config",
		Usage:   "Rollup chain parameters",
		EnvVars: prefixEnvVars("ROLLUP_CONFIG"),
	}
	Network = &cli.StringFlag{
		Name:    "network",
		Usage:   fmt.Sprintf("Predefined network selection. Available networks: %s", strings.Join(chaincfg.AvailableNetworks(), ", ")),
		EnvVars: prefixEnvVars("NETWORK"),
	}
	RPCListenAddr = &cli.StringFlag{
		Name:    "rpc.addr",
		Usage:   "RPC listening address",
		EnvVars: prefixEnvVars("RPC_ADDR"),
	}
	RPCListenPort = &cli.IntFlag{
		Name:    "rpc.port",
		Usage:   "RPC listening port",
		EnvVars: prefixEnvVars("RPC_PORT"),
	}
	RPCEnableAdmin = &cli.BoolFlag{
		Name:    "rpc.enable-admin",
		Usage:   "Enable the admin API (experimental)",
		EnvVars: prefixEnvVars("RPC_ENABLE_ADMIN"),
	}

	/* Optional Flags */
	L1TrustRPC = &cli.BoolFlag{
		Name:    "l1.trustrpc",
		Usage:   "Trust the L1 RPC, sync faster at risk of malicious/buggy RPC providing bad or inconsistent L1 data",
		EnvVars: prefixEnvVars("L1_TRUST_RPC"),
	}
	L1RPCProviderKind = &cli.GenericFlag{
		Name: "l1.rpckind",
		Usage: "The kind of RPC provider, used to inform optimal transactions receipts fetching, and thus reduce costs. Valid options: " +
			EnumString[sources.RPCProviderKind](sources.RPCProviderKinds),
		EnvVars: prefixEnvVars("L1_RPC_KIND"),
		Value: func() *sources.RPCProviderKind {
			out := sources.RPCKindBasic
			return &out
		}(),
	}
	L1RPCRateLimit = &cli.Float64Flag{
		Name:    "l1.rpc-rate-limit",
		Usage:   "Optional self-imposed global rate-limit on L1 RPC requests, specified in requests / second. Disabled if set to 0.",
		EnvVars: prefixEnvVars("L1_RPC_RATE_LIMIT"),
		Value:   0,
	}
	L1RPCMaxBatchSize = &cli.IntFlag{
		Name:    "l1.rpc-max-batch-size",
		Usage:   "Maximum number of RPC requests to bundle, e.g. during L1 blocks receipt fetching. The L1 RPC rate limit counts this as N items, but allows it to burst at once.",
		EnvVars: prefixEnvVars("L1_RPC_MAX_BATCH_SIZE"),
		Value:   20,
	}
	L1HTTPPollInterval = &cli.DurationFlag{
		Name:    "l1.http-poll-interval",
		Usage:   "Polling interval for latest-block subscription when using an HTTP RPC provider. Ignored for other types of RPC endpoints.",
		EnvVars: prefixEnvVars("L1_HTTP_POLL_INTERVAL"),
		Value:   time.Second * 12,
	}
	L2EngineJWTSecret = &cli.StringFlag{
		Name:        "l2.jwt-secret",
		Usage:       "Path to JWT secret key. Keys are 32 bytes, hex encoded in a file. A new key will be generated if left empty.",
		EnvVars:     prefixEnvVars("L2_ENGINE_AUTH"),
		Required:    false,
		Value:       "",
		Destination: new(string),
	}
	SyncerL1Confs = &cli.Uint64Flag{
		Name:     "syncer.l1-confs",
		Usage:    "Number of L1 blocks to keep distance from the L1 head before deriving L2 data from. Reorgs are supported, but may be slow to perform.",
		EnvVars:  prefixEnvVars("SYNCER_L1_CONFS"),
		Required: false,
		Value:    0,
	}
	ProposerEnabledFlag = &cli.BoolFlag{
		Name:    "proposer.enabled",
		Usage:   "Enable proposing of new L2 blocks. A separate batch submitter has to be deployed to publish the data for syncers.",
		EnvVars: prefixEnvVars("PROPOSER_ENABLED"),
	}
	ProposerStoppedFlag = &cli.BoolFlag{
		Name:    "proposer.stopped",
		Usage:   "Initialize the proposer in a stopped state. The proposer can be started using the admin_startProposer RPC",
		EnvVars: prefixEnvVars("PROPOSER_STOPPED"),
	}
	ProposerMaxSafeLagFlag = &cli.Uint64Flag{
		Name:     "proposer.max-safe-lag",
		Usage:    "Maximum number of L2 blocks for restricting the distance between L2 safe and unsafe. Disabled if 0.",
		EnvVars:  prefixEnvVars("PROPOSER_MAX_SAFE_LAG"),
		Required: false,
		Value:    0,
	}
	ProposerL1Confs = &cli.Uint64Flag{
		Name:     "proposer.l1-confs",
		Usage:    "Number of L1 blocks to keep distance from the L1 head as a proposer for picking an L1 origin.",
		EnvVars:  prefixEnvVars("PROPOSER_L1_CONFS"),
		Required: false,
		Value:    4,
	}
	L1EpochPollIntervalFlag = &cli.DurationFlag{
		Name:     "l1.epoch-poll-interval",
		Usage:    "Poll interval for retrieving new L1 epoch updates such as safe and finalized block changes. Disabled if 0 or negative.",
		EnvVars:  prefixEnvVars("L1_EPOCH_POLL_INTERVAL"),
		Required: false,
		Value:    time.Second * 12 * 32,
	}
	MetricsEnabledFlag = &cli.BoolFlag{
		Name:    "metrics.enabled",
		Usage:   "Enable the metrics server",
		EnvVars: prefixEnvVars("METRICS_ENABLED"),
	}
	MetricsAddrFlag = &cli.StringFlag{
		Name:    "metrics.addr",
		Usage:   "Metrics listening address",
		Value:   "0.0.0.0",
		EnvVars: prefixEnvVars("METRICS_ADDR"),
	}
	MetricsPortFlag = &cli.IntFlag{
		Name:    "metrics.port",
		Usage:   "Metrics listening port",
		Value:   7300,
		EnvVars: prefixEnvVars("METRICS_PORT"),
	}
	PprofEnabledFlag = &cli.BoolFlag{
		Name:    "pprof.enabled",
		Usage:   "Enable the pprof server",
		EnvVars: prefixEnvVars("PPROF_ENABLED"),
	}
	PprofAddrFlag = &cli.StringFlag{
		Name:    "pprof.addr",
		Usage:   "pprof listening address",
		Value:   "0.0.0.0",
		EnvVars: prefixEnvVars("PPROF_ADDR"),
	}
	PprofPortFlag = &cli.IntFlag{
		Name:    "pprof.port",
		Usage:   "pprof listening port",
		Value:   6060,
		EnvVars: prefixEnvVars("PPROF_PORT"),
	}
	SnapshotLog = &cli.StringFlag{
		Name:    "snapshotlog.file",
		Usage:   "Path to the snapshot log file",
		EnvVars: prefixEnvVars("SNAPSHOT_LOG"),
	}
	HeartbeatEnabledFlag = &cli.BoolFlag{
		Name:    "heartbeat.enabled",
		Usage:   "Enables or disables heartbeating",
		EnvVars: prefixEnvVars("HEARTBEAT_ENABLED"),
	}
	HeartbeatMonikerFlag = &cli.StringFlag{
		Name:    "heartbeat.moniker",
		Usage:   "Sets a moniker for this node",
		EnvVars: prefixEnvVars("HEARTBEAT_MONIKER"),
	}
	HeartbeatURLFlag = &cli.StringFlag{
		Name:    "heartbeat.url",
		Usage:   "Sets the URL to heartbeat to",
		EnvVars: prefixEnvVars("HEARTBEAT_URL"),
		Value:   "https://heartbeat.kroma-main.io",
	}
	BackupL2UnsafeSyncRPC = &cli.StringFlag{
		Name:     "l2.backup-unsafe-sync-rpc",
		Usage:    "Set the backup L2 unsafe sync RPC endpoint.",
		EnvVars:  prefixEnvVars("L2_BACKUP_UNSAFE_SYNC_RPC"),
		Required: false,
	}
	BackupL2UnsafeSyncRPCTrustRPC = &cli.StringFlag{
		Name: "l2.backup-unsafe-sync-rpc.trustrpc",
		Usage: "Like l1.trustrpc, configure if response data from the RPC needs to be verified, e.g. blockhash computation." +
			"This does not include checks if the blockhash is part of the canonical chain.",
		EnvVars:  prefixEnvVars("L2_BACKUP_UNSAFE_SYNC_RPC_TRUST_RPC"),
		Required: false,
	}
)

var requiredFlags = []cli.Flag{
	L1NodeAddr,
	L2EngineAddr,
	RPCListenAddr,
	RPCListenPort,
}

var optionalFlags = []cli.Flag{
	RollupConfig,
	Network,
	L1TrustRPC,
	L1RPCProviderKind,
	L1RPCRateLimit,
	L1RPCMaxBatchSize,
	L1HTTPPollInterval,
	L2EngineJWTSecret,
	SyncerL1Confs,
	ProposerEnabledFlag,
	ProposerStoppedFlag,
	ProposerMaxSafeLagFlag,
	ProposerL1Confs,
	L1EpochPollIntervalFlag,
	RPCEnableAdmin,
	MetricsEnabledFlag,
	MetricsAddrFlag,
	MetricsPortFlag,
	PprofEnabledFlag,
	PprofAddrFlag,
	PprofPortFlag,
	SnapshotLog,
	HeartbeatEnabledFlag,
	HeartbeatMonikerFlag,
	HeartbeatURLFlag,
	BackupL2UnsafeSyncRPC,
	BackupL2UnsafeSyncRPCTrustRPC,
}

// Flags contains the list of configuration options available to the binary.
var Flags []cli.Flag

func init() {
	optionalFlags = append(optionalFlags, p2pFlags...)
	optionalFlags = append(optionalFlags, klog.CLIFlags(EnvVarPrefix)...)
	Flags = append(requiredFlags, optionalFlags...)
}

func CheckRequired(ctx *cli.Context) error {
	for _, f := range requiredFlags {
		if !ctx.IsSet(f.Names()[0]) {
			return fmt.Errorf("flag %s is required", f.Names()[0])
		}
	}
	return nil
}
