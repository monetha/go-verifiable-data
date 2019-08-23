package commands

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/monetha/go-verifiable-data/cmd/internal/cmdutils"
	"github.com/monetha/go-verifiable-data/cmd/privatedata-exchange/commands/flag"
	"github.com/monetha/go-verifiable-data/deployer"
	"github.com/pkg/errors"
)

// DeployPassportFactoryCommand deploys PassportFactory contract
type DeployPassportFactoryCommand struct {
	cmdutils.QuorumPrivateTxIOCommand
	OwnerKey     flag.ECDSAPrivateKeyFromFile `long:"ownerkey" required:"true"     description:"factory contract owner private key filename"`
	BackendURL   string                       `long:"backendurl"       required:"true" description:"Ethereum backend URL"`
	RegistryAddr flag.EthereumAddress         `long:"registryaddr"     required:"true" description:"Ethereum address of passport logic registry contract"`
	Verbosity    int                          `long:"verbosity"                        description:"log verbosity (0-9)" default:"2"`
	VModule      string                       `long:"vmodule"                          description:"log verbosity pattern"`
}

// Execute implements flags.Commander interface
func (c *DeployPassportFactoryCommand) Execute(args []string) error {
	cmdutils.InitLogging(log.Lvl(c.Verbosity), c.VModule)
	ctx := cmdutils.CreateCtrlCContext()

	ownerAddress := bind.NewKeyedTransactor(c.OwnerKey.AsECDSAPrivateKey()).From
	log.Warn("Loaded configuration",
		"owner_address", ownerAddress.Hex(),
		"backend_url", c.BackendURL,
		"registry", c.RegistryAddr.AsCommonAddress().Hex())

	e, err := cmdutils.NewEth(ctx, c.BackendURL, c.QuorumEnclave, c.QuorumPrivateFor.AsStringArr())
	if err != nil {
		return err
	}

	e = e.NewHandleNonceBackend([]common.Address{ownerAddress})

	// Creating owner session
	ownerSession := e.NewSession(c.OwnerKey.AsECDSAPrivateKey())

	// Deploying passport factory contract
	_, err = deployer.New(ownerSession).DeployPassportFactory(ctx, c.RegistryAddr.AsCommonAddress())
	if err != nil {
		return errors.Wrap(err, "failed to deploy passport factory")
	}

	return nil
}
