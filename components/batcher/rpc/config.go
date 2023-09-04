package rpc

import (
	"github.com/urfave/cli/v2"

	kservice "github.com/kroma-network/kroma/utils/service"
	krpc "github.com/kroma-network/kroma/utils/service/rpc"
)

const (
	EnableAdminFlagName = "rpc.enable-admin"
)

func CLIFlags(envPrefix string) []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:    EnableAdminFlagName,
			Usage:   "Enable the admin API (experimental)",
			EnvVars: kservice.PrefixEnvVar(envPrefix, "RPC_ENABLE_ADMIN"),
		},
	}
}

type CLIConfig struct {
	krpc.CLIConfig
	EnableAdmin bool
}

func ReadCLIConfig(ctx *cli.Context) CLIConfig {
	return CLIConfig{
		CLIConfig:   krpc.ReadCLIConfig(ctx),
		EnableAdmin: ctx.Bool(EnableAdminFlagName),
	}
}

func (c *CLIConfig) ToServiceCLIConfig() krpc.CLIConfig {
	return krpc.CLIConfig{
		ListenAddr: c.ListenAddr,
		ListenPort: c.ListenPort,
	}
}
