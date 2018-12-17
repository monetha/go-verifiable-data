package main

import (
	"crypto/ecdsa"
	"flag"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/monetha/reputation-go-sdk/cmd/internal/cmdutils"
	"github.com/monetha/reputation-go-sdk/deployer"
	"github.com/monetha/reputation-go-sdk/eth"
	"github.com/monetha/reputation-go-sdk/eth/backend"
	"github.com/monetha/reputation-go-sdk/pass"
)

func main() {
	var (
		backendURL         = flag.String("backendurl", "", "backend URL (simulated backend used if empty)")
		passAddr           = cmdutils.AddressVar("passaddr", common.Address{}, "Ethereum address of passport contract")
		ownerKeyFile       = flag.String("ownerkey", "", "owner private key filename")
		ownerKeyHex        = flag.String("ownerkeyhex", "", "private key as hex (for testing)")
		addFactProvider    = cmdutils.AddressVar("add", common.Address{}, "add fact provider to the whitelist")
		removeFactProvider = cmdutils.AddressVar("remove", common.Address{}, "remove fact provider from the whitelist")
		onlyWhitelist      = cmdutils.BoolVar("onlywhitelist", false, "enables or disables the use of a whitelist of fact providers")
		verbosity          = flag.Int("verbosity", int(log.LvlWarn), "log verbosity (0-9)")
		vmodule            = flag.String("vmodule", "", "log verbosity pattern")

		ownerKey *ecdsa.PrivateKey
		err      error
	)
	flag.Parse()

	glogger := log.NewGlogHandler(log.StreamHandler(os.Stderr, log.TerminalFormat(false)))
	glogger.Verbosity(log.Lvl(*verbosity))
	glogger.Vmodule(*vmodule)
	log.Root().SetHandler(glogger)

	switch {
	case !passAddr.IsSet() && *backendURL != "":
		utils.Fatalf("Use -passaddr to specify an address of passport contract")
	case *ownerKeyFile == "" && *ownerKeyHex == "":
		utils.Fatalf("Use -ownerkey or -ownerkeyhex to specify a private key")
	case *ownerKeyFile != "" && *ownerKeyHex != "":
		utils.Fatalf("Options -ownerkey or -ownerkeyhex are mutually exclusive")
	case !addFactProvider.IsSet() && !removeFactProvider.IsSet() && !onlyWhitelist.IsSet():
		utils.Fatalf("Use -add, -remove or -onlywhitelist to specify what to do")
	case (addFactProvider.IsSet() && removeFactProvider.IsSet()) ||
		(addFactProvider.IsSet() && onlyWhitelist.IsSet()) ||
		(removeFactProvider.IsSet() && onlyWhitelist.IsSet()):
		utils.Fatalf("Options -add, -remove or -onlywhitelist are mutually exclusive")
	case *ownerKeyFile != "":
		if ownerKey, err = crypto.LoadECDSA(*ownerKeyFile); err != nil {
			utils.Fatalf("-ownerkey: %v", err)
		}
	case *ownerKeyHex != "":
		if ownerKey, err = crypto.HexToECDSA(*ownerKeyHex); err != nil {
			utils.Fatalf("-ownerkeyhex: %v", err)
		}
	}

	passportAddress := passAddr.GetValue()
	ownerAddress := bind.NewKeyedTransactor(ownerKey).From
	log.Warn("Loaded configuration", "owner_address", ownerAddress.Hex(), "backend_url", *backendURL, "passport", passportAddress.Hex())

	ctx := cmdutils.CreateCtrlCContext()

	var (
		e *eth.Eth
	)
	if *backendURL == "" {
		monethaKey, err := crypto.GenerateKey()
		cmdutils.CheckErr(err, "generating key")
		monethaAddress := bind.NewKeyedTransactor(monethaKey).From

		alloc := core.GenesisAlloc{
			monethaAddress: {Balance: big.NewInt(deployer.PassportFactoryGasLimit)},
			ownerAddress:   {Balance: big.NewInt(deployer.PassportGasLimit + 10000000)},
		}
		sim := backend.NewSimulatedBackendExtended(alloc, 10000000)
		sim.Commit()

		e = eth.New(sim, log.Warn)
		cmdutils.CheckErr(e.UpdateSuggestedGasPrice(ctx), "SuggestGasPrice")

		// creating owner session and checking balance
		monethaSession := e.NewSession(monethaKey)
		cmdutils.CheckBalance(ctx, monethaSession, deployer.PassportFactoryGasLimit)

		// deploying passport factory
		passportFactoryAddress, err := deployer.New(monethaSession).DeployPassportFactory(ctx)
		cmdutils.CheckErr(err, "create passport factory")

		// deploying passport
		ownerSession := e.NewSession(ownerKey)
		passportAddress, err = deployer.New(ownerSession).DeployPassport(ctx, passportFactoryAddress)
		cmdutils.CheckErr(err, "deploying passport")
	} else {
		client, err := ethclient.Dial(*backendURL)
		cmdutils.CheckErr(err, "ethclient.Dial")

		e = eth.New(client, log.Warn)
		cmdutils.CheckErr(e.UpdateSuggestedGasPrice(ctx), "SuggestGasPrice")
	}

	e = e.NewHandleNonceBackend([]common.Address{ownerAddress})

	// creating owner session and checking balance
	ownerSession := e.NewSession(ownerKey)

	w := pass.NewPermissionWriter(ownerSession, passportAddress)

	switch {
	case addFactProvider.IsSet():
		log.Warn("adding provider to whitelist", "fact_provider", addFactProvider.GetValue())
		cmdutils.CheckErr(ignoreHash(w.AddFactProviderToWhitelist(ctx, addFactProvider.GetValue())), "AddFactProviderToWhitelist")
	case removeFactProvider.IsSet():
		log.Warn("removing provider from whitelist", "fact_provider", removeFactProvider.GetValue())
		cmdutils.CheckErr(ignoreHash(w.RemoveFactProviderFromWhitelist(ctx, removeFactProvider.GetValue())), "RemoveFactProviderFromWhitelist")
	case onlyWhitelist.IsSet():
		log.Warn("set whitelist only permission", "value", onlyWhitelist.GetValue())
		cmdutils.CheckErr(ignoreHash(w.SetWhitelistOnlyPermission(ctx, onlyWhitelist.GetValue())), "SetWhitelistOnlyPermission")
	}

	log.Warn("Done.")
}

func ignoreHash(_ common.Hash, err error) error {
	return err
}
