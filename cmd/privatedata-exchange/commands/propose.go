package commands

import (
	"crypto/rand"
	"io/ioutil"

	"github.com/ethereum/go-ethereum/log"
	"github.com/monetha/reputation-go-sdk/cmd/internal/cmdutils"
	"github.com/monetha/reputation-go-sdk/cmd/privatedata-exchange/commands/flag"
	"github.com/monetha/reputation-go-sdk/facts"
	"github.com/pkg/errors"
)

var (
	// ErrExchangeKeyFileExists error returned when exchange key file is already exists (to not overwrite it)
	ErrExchangeKeyFileExists = errors.New("exchange key file already exists")
)

// ProposeCommand handles propose command
type ProposeCommand struct {
	PassportAddress  flag.EthereumAddress         `long:"passportaddr"     required:"true" description:"Ethereum address of passport contract"`
	FactProvider     flag.EthereumAddress         `long:"factprovideraddr" required:"true" description:"Ethereum address of fact provider"`
	FactKey          flag.FactKey                 `long:"fkey"             required:"true" description:"the key of the fact (max. 32 bytes)"`
	DataRequesterKey flag.ECDSAPrivateKeyFromFile `long:"requesterkey"     required:"true" description:"data requester private key filename"`
	StakedValue      flag.EthereumWei             `long:"stake"            required:"true" description:"amount of ethers to stake (in wei)"`
	ExchangeKeyFile  string                       `long:"exchangekey"      required:"true" description:"file name where to save the exchange key (output)"`
	BackendURL       string                       `long:"backendurl"       required:"true" description:"Ethereum backend URL"`
	Verbosity        int                          `long:"verbosity"                        description:"log verbosity (0-9)" default:"2"`
	VModule          string                       `long:"vmodule"                          description:"log verbosity pattern"`
}

// Execute implements flags.Commander interface
func (c *ProposeCommand) Execute(args []string) error {
	if fileExistsAndNotTty(c.ExchangeKeyFile) {
		return ErrExchangeKeyFileExists
	}

	initLogging(log.Lvl(c.Verbosity), c.VModule)
	ctx := cmdutils.CreateCtrlCContext()

	e, err := newEth(ctx, c.BackendURL)
	if err != nil {
		return err
	}

	propResult, err := facts.NewExchangeProposer(
		e.NewSession(c.DataRequesterKey.AsECDSAPrivateKey()),
	).ProposePrivateDataExchange(
		ctx,
		c.PassportAddress.AsCommonAddress(),
		c.FactProvider.AsCommonAddress(),
		c.FactKey,
		c.StakedValue.AsBigInt(),
		rand.Reader,
	)
	if err != nil {
		return err
	}
	log.Warn("Private data exchange proposed", "exchange_index", propResult.ExchangeIdx.String())

	log.Warn("Writing exchange key to file", "file_name", c.ExchangeKeyFile)
	if err := ioutil.WriteFile(c.ExchangeKeyFile, propResult.ExchangeKey[:], 0400); err != nil {
		return errors.Wrap(err, "failed to write exchange key")
	}

	return nil
}
