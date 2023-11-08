package rpc

import (
	"github.com/urfave/cli"

	opservice "github.com/ethereum-optimism/optimism/op-service"
	oprpc "github.com/ethereum-optimism/optimism/op-service/rpc"
)

const (
	EnableAdminFlagName = "rpc.enable-admin"
)

func CLIFlags(envPrefix string) []cli.Flag {
	return []cli.Flag{
		cli.BoolFlag{
			Name:   EnableAdminFlagName,
			Usage:  "Enable the admin API (experimental)",
			EnvVar: opservice.PrefixEnvVar(envPrefix, "RPC_ENABLE_ADMIN"),
		},
	}
}

type CLIConfig struct {
	oprpc.CLIConfig
	EnableAdmin bool
}

func ReadCLIConfig(ctx *cli.Context) CLIConfig {
	return CLIConfig{
		CLIConfig:   oprpc.ReadCLIConfig(ctx),
		EnableAdmin: ctx.GlobalBool(EnableAdminFlagName),
	}
}

func (c *CLIConfig) ToServiceCLIConfig() oprpc.CLIConfig {
	return oprpc.CLIConfig{
		ListenAddr: c.ListenAddr,
		ListenPort: c.ListenPort,
	}
}
