package commands

import (
	"github.com/monetha/go-verifiable-data/deployer"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/monetha/go-verifiable-data/cmd/internal/cmdutils"
	"github.com/monetha/go-verifiable-data/cmd/privatedata-exchange/commands/flag"
	"github.com/pkg/errors"
)

// UpgradePassportLogicCommand upgrades passport logic
type UpgradePassportLogicCommand struct {
	cmdutils.QuorumPrivateTxIOCommand
	PassportOwnerKey flag.ECDSAPrivateKeyFromFile `long:"ownerkey" required:"true"     description:"passport owner private key filename"`
	RegistryAddr     flag.EthereumAddress         `long:"registryaddr"     required:"true" description:"Ethereum address of passport logic registry contract"`
	Version          string                       `long:"newversion"       required:"true" description:"the version of new passport logic contract (which will be deployed)"`
	BackendURL       string                       `long:"backendurl"       required:"true" description:"Ethereum backend URL"`
	Verbosity        int                          `long:"verbosity"                        description:"log verbosity (0-9)" default:"2"`
	VModule          string                       `long:"vmodule"                          description:"log verbosity pattern"`
}

// Execute implements flags.Commander interface
func (c *UpgradePassportLogicCommand) Execute(args []string) error {
	cmdutils.InitLogging(log.Lvl(c.Verbosity), c.VModule)
	ctx := cmdutils.CreateCtrlCContext()

	ownerAddress := bind.NewKeyedTransactor(c.PassportOwnerKey.AsECDSAPrivateKey()).From
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
	ownerSession := e.NewSession(c.PassportOwnerKey.AsECDSAPrivateKey())

	// Upgrading passport logic
	_, err = deployer.New(ownerSession).
		UpgradePassportLogic(ctx, c.Version, c.RegistryAddr.AsCommonAddress())
	if err != nil {
		return errors.Wrap(err, "failed to upgrade passport logic")
	}

	log.Warn("Done.")
	return nil
}
