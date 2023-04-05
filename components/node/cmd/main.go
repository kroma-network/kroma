package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/ethereum/go-ethereum/log"
	"github.com/urfave/cli"

	knode "github.com/wemixkanvas/kanvas/components/node"
	"github.com/wemixkanvas/kanvas/components/node/chaincfg"
	"github.com/wemixkanvas/kanvas/components/node/cmd/doc"
	"github.com/wemixkanvas/kanvas/components/node/cmd/genesis"
	"github.com/wemixkanvas/kanvas/components/node/cmd/p2p"
	"github.com/wemixkanvas/kanvas/components/node/flags"
	"github.com/wemixkanvas/kanvas/components/node/heartbeat"
	"github.com/wemixkanvas/kanvas/components/node/metrics"
	"github.com/wemixkanvas/kanvas/components/node/node"
	"github.com/wemixkanvas/kanvas/components/node/version"
	klog "github.com/wemixkanvas/kanvas/utils/service/log"
	kpprof "github.com/wemixkanvas/kanvas/utils/service/pprof"
)

var (
	Version = ""
	Meta    = ""
)

// VersionWithMeta holds the textual version string including the metadata.
var VersionWithMeta = func() string {
	if Version != "" {
		version.Version = Version
	}
	if Meta != "" {
		version.Meta = Meta
	}
	return fmt.Sprintf("%s-%s", version.Version, version.Meta)
}()

func main() {
	// Set up logger with a default INFO level in case we fail to parse flags,
	// otherwise the final critical log won't show what the parsing error was.
	klog.SetupDefaults()

	app := cli.NewApp()
	app.Version = VersionWithMeta
	app.Flags = flags.Flags
	app.Name = "kanvas-node"
	app.Usage = "Kanvas Rollup Node"
	app.Description = "The Kanvas Rollup Node derives L2 block inputs from L1 data and drives an external L2 Execution Engine to build a L2 chain."
	app.Action = RollupNodeMain
	app.Commands = []cli.Command{
		{
			Name:        "p2p",
			Subcommands: p2p.Subcommands,
		},
		{
			Name:        "genesis",
			Subcommands: genesis.Subcommands,
		},
		{
			Name:        "doc",
			Subcommands: doc.Subcommands,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Crit("Application failed", "message", err)
	}
}

func RollupNodeMain(ctx *cli.Context) error {
	log.Info("Initializing Rollup Node")
	logCfg := klog.ReadCLIConfig(ctx)
	if err := logCfg.Check(); err != nil {
		log.Error("Unable to create the log config", "error", err)
		return err
	}
	log := klog.NewLogger(logCfg)
	m := metrics.NewMetrics("default")

	cfg, err := knode.NewConfig(ctx, log)
	if err != nil {
		log.Error("Unable to create the rollup node config", "error", err)
		return err
	}
	snapshotLog, err := knode.NewSnapshotLogger(ctx)
	if err != nil {
		log.Error("Unable to create snapshot root logger", "error", err)
		return err
	}

	// Only pretty-print the banner if it is a terminal log. Other log it as key-value pairs.
	if logCfg.Format == "terminal" {
		log.Info("rollup config:\n" + cfg.Rollup.Description(chaincfg.L2ChainIDToNetworkName))
	} else {
		cfg.Rollup.LogDescription(log, chaincfg.L2ChainIDToNetworkName)
	}

	n, err := node.New(context.Background(), cfg, log, snapshotLog, VersionWithMeta, m)
	if err != nil {
		log.Error("Unable to create the rollup node", "error", err)
		return err
	}
	log.Info("Starting rollup node", "version", VersionWithMeta)

	if err := n.Start(context.Background()); err != nil {
		log.Error("Unable to start rollup node", "error", err)
		return err
	}
	defer n.Close()

	m.RecordInfo(VersionWithMeta)
	m.RecordUp()
	log.Info("Rollup node started")

	if cfg.Heartbeat.Enabled {
		var peerID string
		if cfg.P2P.Disabled() {
			peerID = "disabled"
		} else {
			peerID = n.P2P().Host().ID().String()
		}

		beatCtx, beatCtxCancel := context.WithCancel(context.Background())
		payload := &heartbeat.Payload{
			Version: version.Version,
			Meta:    version.Meta,
			Moniker: cfg.Heartbeat.Moniker,
			PeerID:  peerID,
			ChainID: cfg.Rollup.L2ChainID.Uint64(),
		}
		go func() {
			if err := heartbeat.Beat(beatCtx, log, cfg.Heartbeat.URL, payload); err != nil {
				log.Error("heartbeat goroutine crashed", "err", err)
			}
		}()
		defer beatCtxCancel()
	}

	if cfg.Pprof.Enabled {
		pprofCtx, pprofCancel := context.WithCancel(context.Background())
		go func() {
			log.Info("pprof server started", "addr", net.JoinHostPort(cfg.Pprof.ListenAddr, strconv.Itoa(cfg.Pprof.ListenPort)))
			if err := kpprof.ListenAndServe(pprofCtx, cfg.Pprof.ListenAddr, cfg.Pprof.ListenPort); err != nil {
				log.Error("error starting pprof", "err", err)
			}
		}()
		defer pprofCancel()
	}

	interruptChannel := make(chan os.Signal, 1)
	signal.Notify(interruptChannel, []os.Signal{
		os.Interrupt,
		os.Kill,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	}...)
	<-interruptChannel

	return nil

}
