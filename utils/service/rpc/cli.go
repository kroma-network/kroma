package rpc

import (
	"errors"
	"math"

	"github.com/urfave/cli/v2"

	kservice "github.com/kroma-network/kroma/utils/service"
)

const (
	ListenAddrFlagName = "rpc.addr"
	PortFlagName       = "rpc.port"
)

func CLIFlags(envPrefix string) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    ListenAddrFlagName,
			Usage:   "rpc listening address",
			Value:   "0.0.0.0",
			EnvVars: kservice.PrefixEnvVar(envPrefix, "RPC_ADDR"),
		},
		&cli.IntFlag{
			Name:    PortFlagName,
			Usage:   "rpc listening port",
			Value:   8545,
			EnvVars: kservice.PrefixEnvVar(envPrefix, "RPC_PORT"),
		},
	}
}

type CLIConfig struct {
	ListenAddr string
	ListenPort int
}

func (c CLIConfig) Check() error {
	if c.ListenPort < 0 || c.ListenPort > math.MaxUint16 {
		return errors.New("invalid RPC port")
	}

	return nil
}

func ReadCLIConfig(ctx *cli.Context) CLIConfig {
	return CLIConfig{
		ListenAddr: ctx.String(ListenAddrFlagName),
		ListenPort: ctx.Int(PortFlagName),
	}
}
