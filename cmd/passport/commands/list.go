package commands

import (
	"fmt"
	"os"

	"github.com/monetha/go-verifiable-data/cmd/internal/flag"

	"github.com/ethereum/go-ethereum/log"
	"github.com/monetha/go-verifiable-data/cmd/internal/cmdutils"
	"github.com/monetha/go-verifiable-data/passfactory"
	"github.com/pkg/errors"
)

// PassportListCommand lists all passports from given passport factory
type PassportListCommand struct {
	BackendURL  string               `long:"backendurl"       required:"true" description:"Ethereum backend URL"`
	FactoryAddr flag.EthereumAddress `long:"factoryaddr"     required:"true" description:"Ethereum address of passport factory contract"`
	FileName    string               `long:"out"     required:"true" description:"Save retrieved passports to the specified file"`
	Verbosity   int                  `long:"verbosity"                        description:"log verbosity (0-9)" default:"2"`
	VModule     string               `long:"vmodule"                          description:"log verbosity pattern"`
}

// Execute implements flags.Commander interface
func (c *PassportListCommand) Execute(args []string) error {
	cmdutils.InitLogging(log.Lvl(c.Verbosity), c.VModule)
	ctx := cmdutils.CreateCtrlCContext()

	log.Warn("Loaded configuration",
		"backend_url", c.BackendURL,
		"factory", c.FactoryAddr.AsCommonAddress().Hex())

	e, err := cmdutils.NewEth(ctx, c.BackendURL, "", nil)
	if err != nil {
		return err
	}

	pfr := passfactory.NewReader(e)
	filterOpts := &passfactory.PassportFilterOpts{
		Context: ctx,
	}

	it, err := pfr.FilterPassports(filterOpts, c.FactoryAddr.AsCommonAddress())
	if err != nil {
		return errors.Wrap(err, "could not get passports")
	}

	defer func() { _ = it.Close() }()

	log.Warn("Writing collected passports to file")

	f, err := os.Create(c.FileName)
	if err != nil {
		return errors.Wrap(err, "could not create output file")
	}
	defer func() { _ = f.Close() }()

	_, err = f.WriteString("passport_address,first_owner,block_number,tx_hash\n")
	if err != nil {
		return errors.Wrap(err, "could not write string to file")
	}

	for it.Next() {
		if it.Error() != nil {
			return errors.Wrap(it.Error(), "could not get next passport")
		}

		p := it.Passport
		_, err = f.WriteString(fmt.Sprintf("%v,%v,%v,%v\n", p.ContractAddress.Hex(), p.FirstOwner.Hex(), p.Raw.BlockNumber, p.Raw.TxHash.Hex()))
		if err != nil {
			return errors.Wrap(err, "could not write string to file")
		}
	}

	log.Warn("Done.")
	return nil
}
