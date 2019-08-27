package commands

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/monetha/go-verifiable-data/cmd/internal/cmdutils"
	"github.com/monetha/go-verifiable-data/cmd/internal/flag"
	"github.com/monetha/go-verifiable-data/deployer"
	"github.com/pkg/errors"
)

// DeployPassportCommand deploys Passport contract
type DeployPassportCommand struct {
	cmdutils.QuorumPrivateTxIOCommand
	PassportOwnerKey flag.ECDSAPrivateKeyFromFile `long:"ownerkey" required:"true"     description:"passport owner private key filename"`
	BackendURL       string                       `long:"backendurl"       required:"true" description:"Ethereum backend URL"`
	FactoryAddr      flag.EthereumAddress         `long:"factoryaddr"     required:"true" description:"Ethereum address of passport factory contract"`
	Verbosity        int                          `long:"verbosity"                        description:"log verbosity (0-9)" default:"2"`
	VModule          string                       `long:"vmodule"                          description:"log verbosity pattern"`
}

// Execute implements flags.Commander interface
func (c *DeployPassportCommand) Execute(args []string) error {
	cmdutils.InitLogging(log.Lvl(c.Verbosity), c.VModule)
	ctx := cmdutils.CreateCtrlCContext()

	ownerAddress := bind.NewKeyedTransactor(c.PassportOwnerKey.AsECDSAPrivateKey()).From
	log.Warn("Loaded configuration",
		"owner_address", ownerAddress.Hex(),
		"backend_url", c.BackendURL,
		"factory", c.FactoryAddr.AsCommonAddress().Hex())

	e, err := cmdutils.NewEth(ctx, c.BackendURL, c.QuorumEnclave, c.QuorumPrivateFor.AsStringArr())
	if err != nil {
		return err
	}

	e = e.NewHandleNonceBackend([]common.Address{ownerAddress})

	// Creating owner session
	ownerSession := e.NewSession(c.PassportOwnerKey.AsECDSAPrivateKey())
	err = cmdutils.CheckBalance(ctx, ownerSession, deployer.PassportGasLimit)
	if err != nil {
		return err
	}

	// Deploying passport contract
	_, err = deployer.New(ownerSession).
		DeployPassport(ctx, c.FactoryAddr.AsCommonAddress())
	if err != nil {
		return errors.Wrap(err, "failed to deploy passport")
	}

	log.Warn("Done.")
	return nil
}
