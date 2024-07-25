package main

import (
	"fmt"
	"os"

	"github.com/ethereum-optimism/optimism/op-service/cliapp"
	oplog "github.com/ethereum-optimism/optimism/op-service/log"
	"github.com/ethereum/go-ethereum/log"
	"github.com/urfave/cli/v2"

	"github.com/kroma-network/kroma/kroma-validator"
	cmd "github.com/kroma-network/kroma/kroma-validator/cmd/validator"
	"github.com/kroma-network/kroma/kroma-validator/flags"
)

var (
	Version = ""
	Meta    = ""
)

func main() {
	oplog.SetupDefaults()

	app := cli.NewApp()
	app.Flags = cliapp.ProtectFlags(flags.Flags)
	app.Version = fmt.Sprintf("%s-%s", Version, Meta)
	app.Name = "kroma-validator"
	app.Usage = "L2 Output Submitter and Challenger Service"
	app.Description = "Service for generating and submitting L2 output checkpoints to the L2OutputOracle contract as an L2 Output Submitter, " + "detecting and correcting invalid L2 outputs as a Challenger to ensure the integrity of the L2 state."
	app.Action = curryMain(Version)
	app.Commands = cli.Commands{
		{
			Name:  "register",
			Usage: "Register as new validator to ValidatorManager",
			Flags: []cli.Flag{
				cmd.TokenAmountFlag,
				cmd.CommissionRateFlag,
				cmd.WithdrawAccountFlag,
			},
			Action: cmd.Register,
		},
		{
			Name:   "activate",
			Usage:  "Activate the validator in ValidatorManager",
			Action: cmd.Activate,
		},
		{
			Name:   "unjail",
			Usage:  "Attempt to unjail the validator in ValidatorManager",
			Action: cmd.Unjail,
		},
		{
			Name:  "changeCommission",
			Usage: "Change the commission rate of the validator in ValidatorManager",
			Subcommands: []*cli.Command{
				{
					Name:   "init",
					Usage:  "Initiate the commission rate change",
					Flags:  []cli.Flag{cmd.CommissionRateFlag},
					Action: cmd.InitCommissionChange,
				},
				{
					Name:   "finalize",
					Usage:  "Finalize the commission rate change",
					Action: cmd.FinalizeCommissionChange,
				},
			},
		},
		{
			Name:   "depositKro",
			Usage:  "Attempt to deposit asset tokens to AssetManager to be used as bond",
			Flags:  []cli.Flag{cmd.TokenAmountFlag},
			Action: cmd.DepositKro,
		},
		{
			Name:        "deposit",
			Usage:       "(DEPRECATED) Deposit ETH into ValidatorPool to be used as bond",
			Description: "This command is deprecated since the release of validator system V2. Please use the 'register' command to register as a validator.",
			Flags:       []cli.Flag{cmd.EthAmountFlag},
			Action:      cmd.Deposit,
		},
		{
			Name:        "withdraw",
			Usage:       "(DEPRECATED) Withdraw ETH from ValidatorPool",
			Description: "This command is deprecated since the release of validator system V2. You can still use this command to withdraw your asset from the ValidatorPool.",
			Flags:       []cli.Flag{cmd.EthAmountFlag},
			Action:      cmd.Withdraw,
		},
		{
			Name:        "withdrawTo",
			Usage:       "(DEPRECATED) Withdraw ETH from ValidatorPool to specific address",
			Description: "This command is deprecated since the release of validator system V2. You can still use this command to withdraw your asset from the ValidatorPool to specific address.",
			Flags:       []cli.Flag{cmd.AddressFlag, cmd.EthAmountFlag},
			Action:      cmd.WithdrawTo,
		},
		{
			Name:        "unbond",
			Usage:       "(DEPRECATED) Attempt to unbond in ValidatorPool",
			Description: "This command is deprecated since the release of validator system V2. You can still use this command to unbond your asset from the ValidatorPool.",
			Action:      cmd.Unbond,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Crit("Application failed", "message", err)
	}
}

// curryMain transforms the kroma-validator.Main function into an app.Action
// This is done to capture the Version of the validator.
func curryMain(version string) func(ctx *cli.Context) error {
	return func(ctx *cli.Context) error {
		return validator.Main(version, ctx)
	}
}
