package commands

import (
	"io/ioutil"

	"github.com/ethereum/go-ethereum/log"
	"github.com/monetha/reputation-go-sdk/cmd/internal/cmdutils"
	"github.com/monetha/reputation-go-sdk/cmd/privatedata-exchange/commands/flag"
	"github.com/monetha/reputation-go-sdk/facts"
	"github.com/monetha/reputation-go-sdk/ipfs"
	"github.com/pkg/errors"
)

var (
	// ErrDataFileExists returned when data file already exists
	ErrDataFileExists = errors.New("private data file already exists")
)

// ReadDataCommand handles `readdata` command
type ReadDataCommand struct {
	PassportAddress flag.EthereumAddress     `long:"passportaddr" required:"true" description:"Ethereum address of passport contract"`
	ExchangeIndex   flag.ExchangeIndex       `long:"exchidx"      required:"true" description:"private data exchange index"`
	ExchangeKey     flag.ExchangeKeyFromFile `long:"exchangekey"  required:"true" description:"exchange key filename"`
	DataFile        string                   `long:"datafile"     required:"true" description:"file name where to save the private data (output)"`
	BackendURL      string                   `long:"backendurl"   required:"true" description:"Ethereum backend URL"`
	IPFSURL         string                   `long:"ipfsurl"                      description:"IPFS node URL" default:"https://ipfs.infura.io:5001"`
	Verbosity       int                      `long:"verbosity"                    description:"log verbosity (0-9)" default:"2"`
	VModule         string                   `long:"vmodule"                      description:"log verbosity pattern"`
}

// Execute implements flags.Commander interface
func (c *ReadDataCommand) Execute(args []string) error {
	if fileExists(c.DataFile) {
		return ErrDataFileExists
	}

	initLogging(log.Lvl(c.Verbosity), c.VModule)
	ctx := cmdutils.CreateCtrlCContext()

	e, err := newEth(c.BackendURL)
	if err != nil {
		return err
	}

	fs, err := ipfs.New(c.IPFSURL)
	if err != nil {
		return errors.Wrap(err, "failed to create IPFS client")
	}

	data, err := facts.NewExchangeDataReader(
		e,
		fs,
	).ReadPrivateData(
		ctx,
		c.PassportAddress.AsCommonAddress(),
		c.ExchangeIndex.AsBigInt(),
		c.ExchangeKey,
	)
	if err != nil {
		return err
	}

	log.Warn("Writing private data to file", "file_name", c.DataFile)
	if err := ioutil.WriteFile(c.DataFile, data, 0600); err != nil {
		return errors.Wrap(err, "failed to write private data")
	}

	return nil
}
