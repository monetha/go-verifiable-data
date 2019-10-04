package commands

import (
	"fmt"

	"github.com/ethereum/go-ethereum/log"
	"github.com/monetha/go-verifiable-data/cmd/internal/cmdutils"
	"github.com/monetha/go-verifiable-data/cmd/internal/flag"
	"github.com/monetha/go-verifiable-data/factprovider"
	"github.com/pkg/errors"
)

// GetFactProviderInfoCommand gets fact provider information from FactProviderRegistry contract
type GetFactProviderInfoCommand struct {
	cmdutils.QuorumPrivateTxIOCommand
	BackendURL                  string               `long:"backendurl"       required:"true" description:"Ethereum backend URL"`
	FactProviderRegistryAddress flag.EthereumAddress `long:"registryaddr"     required:"true" description:"Ethereum address of fact provider registry contract"`
	FactProviderAddress         flag.EthereumAddress `long:"provideraddr"     required:"true" description:"Ethereum address of fact provider"`
	Verbosity                   int                  `long:"verbosity"                        description:"log verbosity (0-9)" default:"2"`
	VModule                     string               `long:"vmodule"                          description:"log verbosity pattern"`
}

// Execute implements flags.Commander interface
func (c *GetFactProviderInfoCommand) Execute(args []string) error {
	cmdutils.InitLogging(log.Lvl(c.Verbosity), c.VModule)
	ctx := cmdutils.CreateCtrlCContext()

	log.Warn("Loaded configuration",
		"backend_url", c.BackendURL)

	e, err := cmdutils.NewEth(ctx, c.BackendURL, c.QuorumEnclave, c.QuorumPrivateFor.AsStringArr())
	if err != nil {
		return err
	}

	// Reading information about fact provider
	rd := factprovider.NewRegistryReader(e, c.FactProviderRegistryAddress.AsCommonAddress())

	info, err := rd.GetFactProviderInfo(ctx, c.FactProviderAddress.AsCommonAddress())
	if err != nil {
		return errors.Wrap(err, "failed to get fact provider information from fact provider registry contract")
	}
	if info == nil {
		fmt.Println("No information found about fact provider")
	} else {
		fmt.Println("Fact provider name:", info.Name)
		fmt.Println("Fact provider passport:", info.ReputationPassport.String())
		fmt.Println("Fact provider website:", info.Website)
	}

	return nil
}
