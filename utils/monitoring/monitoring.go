package monitoring

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"

	"github.com/kroma-network/kroma/utils/service/metrics"
	"github.com/kroma-network/kroma/utils/service/pprof"
	krpc "github.com/kroma-network/kroma/utils/service/rpc"
)

// NOTE(pangssu): MaybeStartPprof requires cancelable context to stop http server
func MaybeStartPprof(ctx context.Context, cfg pprof.CLIConfig, l log.Logger) {
	if cfg.Enabled {
		l.Info("starting pprof", "addr", cfg.ListenAddr, "port", cfg.ListenPort)
		go func() {
			if err := pprof.ListenAndServe(ctx, cfg.ListenAddr, cfg.ListenPort); err != nil {
				l.Error("failed to start pprof", "err", err)
			}
		}()
	}
}

// NOTE(pangssu): MaybeStartMetrics requires cancelable context to stop http server
func MaybeStartMetrics(ctx context.Context, cfg metrics.CLIConfig, l log.Logger, l1 *ethclient.Client, wallet common.Address) {
	if cfg.Enabled {
		registry := metrics.NewRegistry()
		l.Info("starting metrics server", "addr", cfg.ListenAddr, "port", cfg.ListenPort)
		go func() {
			if err := metrics.ListenAndServe(ctx, registry, cfg.ListenAddr, cfg.ListenPort); err != nil {
				l.Error("failed to start metrics server", err)
			}
		}()

		metrics.LaunchBalanceMetrics(ctx, l, registry, "", l1, wallet)
	}
}

func StartRPC(cfg krpc.CLIConfig, version string, opts ...krpc.ServerOption) (*krpc.Server, error) {
	server := krpc.NewServer(cfg.ListenAddr, cfg.ListenPort, version, opts...)
	if err := server.Start(); err != nil {
		return nil, fmt.Errorf("failed to start RPC server: %w", err)
	}

	return server, nil
}
