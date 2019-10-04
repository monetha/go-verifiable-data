package commands

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/monetha/go-verifiable-data/cmd/internal/cmdutils"
	"github.com/monetha/go-verifiable-data/cmd/internal/flag"
	"github.com/monetha/go-verifiable-data/factprovider"
	"github.com/pkg/errors"
)

// SetFactProviderInfoCommand sets fact provider information in FactProviderRegistry contract
type SetFactProviderInfoCommand struct {
	cmdutils.QuorumPrivateTxIOCommand
	OwnerKey                    flag.ECDSAPrivateKeyFromFile `long:"ownerkey"         required:"true" description:"fact provider registry owner private key filename"`
	BackendURL                  string                       `long:"backendurl"       required:"true" description:"Ethereum backend URL"`
	FactProviderRegistryAddress flag.EthereumAddress         `long:"registryaddr"     required:"true" description:"Ethereum address of fact provider registry contract"`
	FactProviderAddress         flag.EthereumAddress         `long:"provideraddr"     required:"true" description:"Ethereum address of fact provider"`
	FactProviderName            string                       `long:"providername"     required:"true" description:"Name of fact provider"`
	FactProviderPassportAddress flag.EthereumAddress         `long:"providerpassaddr"                 description:"Ethereum address of passport of fact provider"`
	FactProviderWebsite         string                       `long:"providerwebsite"                  description:"Website of fact provider"`
	Verbosity                   int                          `long:"verbosity"                        description:"log verbosity (0-9)" default:"2"`
	VModule                     string                       `long:"vmodule"                          description:"log verbosity pattern"`
}

// Execute implements flags.Commander interface
func (c *SetFactProviderInfoCommand) Execute(args []string) error {
	cmdutils.InitLogging(log.Lvl(c.Verbosity), c.VModule)
	ctx := cmdutils.CreateCtrlCContext()

	ownerAddress := bind.NewKeyedTransactor(c.OwnerKey.AsECDSAPrivateKey()).From
	log.Warn("Loaded configuration",
		"owner_address", ownerAddress.Hex(),
		"backend_url", c.BackendURL)

	e, err := cmdutils.NewEth(ctx, c.BackendURL, c.QuorumEnclave, c.QuorumPrivateFor.AsStringArr())
	if err != nil {
		return err
	}

	e = e.NewHandleNonceBackend([]common.Address{ownerAddress})

	// Creating owner session
	ownerSession := e.NewSession(c.OwnerKey.AsECDSAPrivateKey())

	// Writing information about fact provider
	_, err = factprovider.
		NewRegistryWriter(ownerSession, c.FactProviderRegistryAddress.AsCommonAddress()).
		SetFactProviderInfo(ctx, c.FactProviderAddress.AsCommonAddress(), &factprovider.Info{
			Name:               c.FactProviderName,
			ReputationPassport: c.FactProviderPassportAddress.AsCommonAddress(),
			Website:            c.FactProviderWebsite,
		})
	if err != nil {
		return errors.Wrap(err, "failed to set fact provider information in fact provider registry contract")
	}

	log.Warn("Done.")
	return nil
}
