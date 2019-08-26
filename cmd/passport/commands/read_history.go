package commands

import (
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/log"
	"github.com/monetha/go-verifiable-data/cmd/internal/cmdutils"
	"github.com/monetha/go-verifiable-data/cmd/privatedata-exchange/commands/flag"
	"github.com/monetha/go-verifiable-data/facts"
	"github.com/pkg/errors"
)

// ReadHistoryCommand lists all passport fact changes
type ReadHistoryCommand struct {
	BackendURL      string               `long:"backendurl"       required:"true" description:"Ethereum backend URL"`
	PassportAddress flag.EthereumAddress `long:"passaddr" required:"true"     description:"Ethereum address of passport contract"`
	FileName        string               `long:"out"     required:"true" description:"Save retrieved passports to the specified file"`
	Verbosity       int                  `long:"verbosity"                        description:"log verbosity (0-9)" default:"2"`
	VModule         string               `long:"vmodule"                          description:"log verbosity pattern"`
}

// Execute implements flags.Commander interface
func (c *ReadHistoryCommand) Execute(args []string) error {
	cmdutils.InitLogging(log.Lvl(c.Verbosity), c.VModule)
	ctx := cmdutils.CreateCtrlCContext()

	log.Warn("Loaded configuration",
		"backend_url", c.BackendURL,
		"passport", c.PassportAddress.AsCommonAddress().Hex())

	e, err := cmdutils.NewEth(ctx, c.BackendURL, "", nil)
	if err != nil {
		return err
	}

	historian := facts.NewHistorian(e)

	f, err := os.Create(c.FileName)
	if err != nil {
		return errors.Wrap(err, "Could not create output file")
	}
	defer func() { _ = f.Close() }()

	it, err := historian.FilterChanges(&facts.ChangesFilterOpts{Context: ctx}, c.PassportAddress.AsCommonAddress())
	if err != nil {
		return errors.Wrap(err, "Could get digital identity changes")
	}

	// Iterate the whole history
	_, err = f.WriteString("fact_provider,key,data_type,change_type,block_number,tx_hash\n")
	if err != nil {
		return errors.Wrap(err, "Could not write to file")
	}

	for it.Next() {
		if err != nil {
			return errors.Wrap(it.Error(), "Could not get next history item")
		}

		ch := it.Change
		_, err = f.WriteString(fmt.Sprintf("%v,%v,%v,%v,%v,%v\n", ch.FactProvider.Hex(), string(ch.Key[:]), ch.DataType, ch.ChangeType, ch.Raw.BlockNumber, ch.Raw.TxHash.Hex()))
		if err != nil {
			return errors.Wrap(err, "Could not write to file")
		}
	}

	log.Warn("Done.")
	return nil
}
