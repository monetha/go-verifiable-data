package commands

import (
	"github.com/ethereum/go-ethereum/log"
	"github.com/monetha/reputation-go-sdk/cmd/internal/cmdutils"
	"github.com/monetha/reputation-go-sdk/cmd/privatedata-exchange/commands/flag"
	"github.com/monetha/reputation-go-sdk/facts"
)

// TimeoutCommand handles timeout command
type TimeoutCommand struct {
	PassportAddress  flag.EthereumAddress         `long:"passportaddr" required:"true" description:"Ethereum address of passport contract"`
	ExchangeIndex    flag.ExchangeIndex           `long:"exchidx"      required:"true" description:"private data exchange index"`
	DataRequesterKey flag.ECDSAPrivateKeyFromFile `long:"requesterkey" required:"true" description:"data requester private key filename"`
	BackendURL       string                       `long:"backendurl"   required:"true" description:"Ethereum backend URL"`
	Verbosity        int                          `long:"verbosity"                    description:"log verbosity (0-9)" default:"2"`
	VModule          string                       `long:"vmodule"                      description:"log verbosity pattern"`
}

// Execute implements flags.Commander interface
func (c *TimeoutCommand) Execute(args []string) error {
	initLogging(log.Lvl(c.Verbosity), c.VModule)
	ctx := cmdutils.CreateCtrlCContext()

	e, err := newEth(c.BackendURL)
	if err != nil {
		return err
	}

	err = facts.NewExchangeTimeouter(
		e.NewSession(c.DataRequesterKey.AsECDSAPrivateKey()),
	).TimeoutPrivateDataExchange(
		ctx,
		c.PassportAddress.AsCommonAddress(),
		c.ExchangeIndex.AsBigInt(),
	)
	if err != nil {
		return err
	}

	return nil
}
