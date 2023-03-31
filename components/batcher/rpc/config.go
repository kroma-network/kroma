package rpc

import (
	"github.com/urfave/cli"

	kservice "github.com/wemixkanvas/kanvas/utils/service"
	krpc "github.com/wemixkanvas/kanvas/utils/service/rpc"
)

const (
	EnableAdminFlagName = "rpc.enable-admin"
)

func CLIFlags(envPrefix string) []cli.Flag {
	return []cli.Flag{
		cli.BoolFlag{
			Name:   EnableAdminFlagName,
			Usage:  "Enable the admin API (experimental)",
			EnvVar: kservice.PrefixEnvVar(envPrefix, "RPC_ENABLE_ADMIN"),
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
		EnableAdmin: ctx.GlobalBool(EnableAdminFlagName),
	}
}

func (c *CLIConfig) ToServiceCLIConfig() krpc.CLIConfig {
	return krpc.CLIConfig{
		ListenAddr: c.ListenAddr,
		ListenPort: c.ListenPort,
	}
}
