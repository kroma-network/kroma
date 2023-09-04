package metrics

import (
	"errors"
	"math"

	"github.com/urfave/cli/v2"

	kservice "github.com/kroma-network/kroma/utils/service"
)

const (
	EnabledFlagName    = "metrics.enabled"
	ListenAddrFlagName = "metrics.addr"
	PortFlagName       = "metrics.port"
)

func CLIFlags(envPrefix string) []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:    EnabledFlagName,
			Usage:   "Enable the metrics server",
			EnvVars: kservice.PrefixEnvVar(envPrefix, "METRICS_ENABLED"),
		},
		&cli.StringFlag{
			Name:    ListenAddrFlagName,
			Usage:   "Metrics listening address",
			Value:   "0.0.0.0",
			EnvVars: kservice.PrefixEnvVar(envPrefix, "METRICS_ADDR"),
		},
		&cli.IntFlag{
			Name:    PortFlagName,
			Usage:   "Metrics listening port",
			Value:   7300,
			EnvVars: kservice.PrefixEnvVar(envPrefix, "METRICS_PORT"),
		},
	}
}

type CLIConfig struct {
	Enabled    bool
	ListenAddr string
	ListenPort int
}

func (m CLIConfig) Check() error {
	if !m.Enabled {
		return nil
	}

	if m.ListenPort < 0 || m.ListenPort > math.MaxUint16 {
		return errors.New("invalid metrics port")
	}

	return nil
}

func ReadCLIConfig(ctx *cli.Context) CLIConfig {
	return CLIConfig{
		Enabled:    ctx.Bool(EnabledFlagName),
		ListenAddr: ctx.String(ListenAddrFlagName),
		ListenPort: ctx.Int(PortFlagName),
	}
}

func ReadLocalCLIConfig(ctx *cli.Context) CLIConfig {
	return CLIConfig{
		Enabled:    ctx.Bool(EnabledFlagName),
		ListenAddr: ctx.String(ListenAddrFlagName),
		ListenPort: ctx.Int(PortFlagName),
	}
}
