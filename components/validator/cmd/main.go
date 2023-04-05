package main

import (
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/log"
	"github.com/urfave/cli"

	validator "github.com/wemixkanvas/kanvas/components/validator"
	"github.com/wemixkanvas/kanvas/components/validator/flags"
	klog "github.com/wemixkanvas/kanvas/utils/service/log"
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
	app.Name = "validator"
	app.Usage = "L2Output Submitter"
	app.Description = "Service for generating and submitting L2 Output checkpoints to the L2OutputOracle contract"

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
