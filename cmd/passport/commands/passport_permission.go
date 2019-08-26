package commands

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/monetha/go-verifiable-data/cmd/internal/cmdutils"
	"github.com/monetha/go-verifiable-data/cmd/privatedata-exchange/commands/flag"
	"github.com/monetha/go-verifiable-data/pass"
	"github.com/pkg/errors"
)

// PassportPermissionCommand changes permissions for passport
type PassportPermissionCommand struct {
	cmdutils.QuorumPrivateTxIOCommand
	PassportAddress    flag.EthereumAddress         `long:"passaddr" required:"true"     description:"Ethereum address of passport contract"`
	PassportOwnerKey   flag.ECDSAPrivateKeyFromFile `long:"ownerkey" required:"true"     description:"passport owner private key filename"`
	BackendURL         string                       `long:"backendurl"       required:"true" description:"Ethereum backend URL"`
	AddFactProvider    *flag.EthereumAddress        `long:"add"  description:"add data source to the whitelist"`
	RemoveFactProvider *flag.EthereumAddress        `long:"remove"  description:"remove data source from the whitelist"`
	EnableWhitelist    *bool                        `long:"enablewhitelist"                        description:"enables the use of a whitelist of data sources"`
	DisableWhitelist   *bool                        `long:"disablewhitelist"                        description:"disables the use of a whitelist of data sources"`
	Verbosity          int                          `long:"verbosity"                        description:"log verbosity (0-9)" default:"2"`
	VModule            string                       `long:"vmodule"                          description:"log verbosity pattern"`
}

// Execute implements flags.Commander interface
func (c *PassportPermissionCommand) Execute(args []string) error {
	cmdutils.InitLogging(log.Lvl(c.Verbosity), c.VModule)
	ctx := cmdutils.CreateCtrlCContext()

	enabledMutExclFlags := []bool{
		c.AddFactProvider != nil,
		c.RemoveFactProvider != nil,
		c.EnableWhitelist != nil,
		c.DisableWhitelist != nil,
	}

	enabledFlagCount := 0
	for _, v := range enabledMutExclFlags {
		if v {
			enabledFlagCount++
		}
	}

	switch {
	case enabledFlagCount == 0:
		return fmt.Errorf("use --add, --remove, --enablewhitelist or --disablewhitelist to specify what to do")
	case enabledFlagCount > 1:
		return fmt.Errorf("options --add, --remove, --enablewhitelist or --disablewhitelist are mutually exclusive")
	}

	ownerAddress := bind.NewKeyedTransactor(c.PassportOwnerKey.AsECDSAPrivateKey()).From
	log.Warn("Loaded configuration",
		"owner_address", ownerAddress.Hex(),
		"backend_url", c.BackendURL,
		"passport", c.PassportAddress.AsCommonAddress().Hex())

	e, err := cmdutils.NewEth(ctx, c.BackendURL, c.QuorumEnclave, c.QuorumPrivateFor.AsStringArr())
	if err != nil {
		return err
	}

	e = e.NewHandleNonceBackend([]common.Address{ownerAddress})

	// Creating owner session
	ownerSession := e.NewSession(c.PassportOwnerKey.AsECDSAPrivateKey())

	w := pass.NewPermissionWriter(ownerSession, c.PassportAddress.AsCommonAddress())

	switch {
	case c.AddFactProvider != nil:
		log.Warn("adding provider to whitelist", "fact_provider", c.AddFactProvider.AsCommonAddress())
		err = ignoreHash(w.AddFactProviderToWhitelist(ctx, c.AddFactProvider.AsCommonAddress()))
		if err != nil {
			return errors.Wrap(err, "could not add data source to whitelist")
		}
	case c.RemoveFactProvider != nil:
		log.Warn("removing provider from whitelist", "fact_provider", c.RemoveFactProvider.AsCommonAddress())
		err = ignoreHash(w.RemoveFactProviderFromWhitelist(ctx, c.RemoveFactProvider.AsCommonAddress()))
		if err != nil {
			return errors.Wrap(err, "could not remove data source from whitelist")
		}
	case c.EnableWhitelist != nil:
		log.Warn("enable whitelist only permission")
		err = ignoreHash(w.SetWhitelistOnlyPermission(ctx, true))
		if err != nil {
			return errors.Wrap(err, "could enable whitelist only permission")
		}
	case c.DisableWhitelist != nil:
		log.Warn("disable whitelist only permission")
		err = ignoreHash(w.SetWhitelistOnlyPermission(ctx, false))
		if err != nil {
			return errors.Wrap(err, "could disable whitelist only permission")
		}
	}

	log.Warn("Done.")
	return nil
}

func ignoreHash(_ common.Hash, err error) error {
	return err
}
