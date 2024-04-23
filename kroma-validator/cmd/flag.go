package main

import (
	"github.com/urfave/cli/v2"
)

var TokenAmountFlag = &cli.StringFlag{
	Name:     "amount",
	Usage:    "Amount of governance token (in wei)",
	Required: true,
}

var EthAmountFlag = &cli.StringFlag{
	Name:     "amount",
	Usage:    "Amount of ETH (in wei)",
	Required: true,
}

var CommissionRateFlag = &cli.Uint64Flag{
	Name:     "commission-rate",
	Usage:    "The commission rate the validator sets (in percentage). Maximum 100.",
	Required: true,
}

var CommissionMaxChangeRateFlag = &cli.Uint64Flag{
	Name:     "commission-max-change-rate",
	Usage:    "The maximum changeable commission rate change (in percentage). Maximum 100.",
	Required: true,
}
