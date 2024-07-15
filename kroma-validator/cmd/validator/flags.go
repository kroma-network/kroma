package validator

import (
	"github.com/urfave/cli/v2"
)

var TokenAmountFlag = &cli.StringFlag{
	Name:     "amount",
	Usage:    "Amount of asset token (in wei)",
	Required: true,
}

var EthAmountFlag = &cli.StringFlag{
	Name:     "amount",
	Usage:    "Amount of ETH (in wei)",
	Required: true,
}

var AddressFlag = &cli.StringFlag{
	Name:     "address",
	Usage:    "Address to receive ETH",
	Required: true,
}

var CommissionRateFlag = &cli.Uint64Flag{
	Name:     "commission-rate",
	Usage:    "The commission rate earned by the validator (in percentage). Maximum 100.",
	Required: true,
}

var WithdrawAccountFlag = &cli.StringFlag{
	Name:     "withdraw-account",
	Usage:    "Address to withdraw deposited asset token",
	Required: true,
}
