package main

import (
	"fmt"

	"os"

	"github.com/ethereum/go-ethereum/log"
	"github.com/urfave/cli/v2"

	"github.com/ethereum-optimism/optimism/op-service/cliapp"
	oplog "github.com/ethereum-optimism/optimism/op-service/log"

	val "github.com/kroma-network/kroma/kroma-validator"
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
	app.Name = "validator"
	app.Usage = "L2 Output Submitter and Challenger Service"
	app.Description = "Service for generating and submitting L2 output checkpoints to the L2OutputOracle contract as an L2 Output Submitter, " + "detecting and correcting invalid L2 outputs as a Challenger to ensure the integrity of the L2 state."
	app.Action = curryMain(Version)
	app.Commands = cli.Commands{
		{
			Name:        "deposit",
			Usage:       "Deposit ETH into ValidatorPool to be used as bond",
			Description: "This command will be deprecated in a future release of validator system V2. Please use the 'register' command to register as a validator.",
			Flags:       []cli.Flag{cmd.EthAmountFlag},
			Action:      cmd.Deposit,
		},
		{
			Name:        "withdraw",
			Usage:       "Withdraw ETH from ValidatorPool",
			Description: "This command will be deprecated in a future release of validator system V2. You can still use this command to withdraw your asset from the ValidatorPool.",
			Flags:       []cli.Flag{cmd.EthAmountFlag},
			Action:      cmd.Withdraw,
		},
		{
			Name:        "unbond",
			Usage:       "Attempt to unbond in ValidatorPool",
			Description: "This command will be deprecated in a future release of validator system V2. You can still use this command to unbond your asset from the ValidatorPool.",
			Action:      cmd.Unbond,
		},
		{
			Name:  "register",
			Usage: "(EXPERIMENTAL) Register the validator to ValidatorManager",
			Flags: []cli.Flag{
				cmd.TokenAmountFlag,
				cmd.CommissionRateFlag,
				cmd.CommissionMaxChangeRateFlag,
			},
			Action: cmd.RegisterValidator,
		},
		{
			Name:   "activate",
			Usage:  "(EXPERIMENTAL) Activate the validator",
			Action: cmd.Activate,
		},
		{
			Name:   "unjail",
			Usage:  "(EXPERIMENTAL) Attempt to unjail the validator",
			Action: cmd.Unjail,
		},
		{
			Name:   "changeCommissionRate",
			Usage:  "(EXPERIMENTAL) Change the commission rate of the validator",
			Flags:  []cli.Flag{cmd.CommissionRateFlag},
			Action: cmd.ChangeCommissionRate,
		},
		{
			Name:   "delegate",
			Usage:  "(EXPERIMENTAL) Attempt to self-delegate governance tokens",
			Flags:  []cli.Flag{cmd.TokenAmountFlag},
			Action: cmd.Delegate,
		},
		{
			Name:  "undelegate",
			Usage: "(EXPERIMENTAL) Undelegate governance tokens",
			Subcommands: []*cli.Command{
				{
					Name:   "init",
					Usage:  "Initiate an undelegation of governance tokens",
					Flags:  []cli.Flag{cmd.TokenAmountFlag},
					Action: cmd.InitUndelegate,
				},
				{
					Name:   "finalize",
					Usage:  "Finalize an undelegation of governance tokens",
					Action: cmd.FinalizeUndelegate,
				},
			},
		},
		{
			Name:  "claim",
			Usage: "(EXPERIMENTAL) Claim validator rewards",
			Subcommands: []*cli.Command{
				{
					Name:   "init",
					Usage:  "Initiate a claim of validator rewards",
					Flags:  []cli.Flag{cmd.TokenAmountFlag},
					Action: cmd.InitClaimValidatorReward,
				},
				{
					Name:   "finalize",
					Usage:  "Finalize a claim of validator rewards",
					Action: cmd.FinalizeClaimValidatorReward,
				},
			},
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
		return val.Main(version, ctx)
	}
}
