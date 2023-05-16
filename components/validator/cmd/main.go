package main

import (
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/log"
	"github.com/urfave/cli"

	validator "github.com/kroma-network/kroma/components/validator"
	"github.com/kroma-network/kroma/components/validator/flags"
	klog "github.com/kroma-network/kroma/utils/service/log"
)

var (
	Version = ""
	Meta    = ""
)

func main() {
	klog.SetupDefaults()

	app := cli.NewApp()
	app.Flags = flags.Flags
	app.Version = fmt.Sprintf("%s-%s", Version, Meta)
	app.Name = "kroma-validator"
	app.Usage = "L2 Output Submitter and Challenger Service"
	app.Description = "Service for generating and submitting L2 output checkpoints to the L2OutputOracle contract as an L2 Output Submitter, " + "detecting and correcting invalid L2 outputs as a Challenger to ensure the integrity of the L2 state."

	app.Action = curryMain(Version)
	err := app.Run(os.Args)
	if err != nil {
		log.Crit("Application failed", "message", err)
	}
}

// curryMain transforms the validator.Main function into an app.Action
// This is done to capture the Version of the validator.
func curryMain(version string) func(ctx *cli.Context) error {
	return func(ctx *cli.Context) error {
		return validator.Main(version, ctx)
	}
}
