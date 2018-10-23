package main

import (
	"crypto/ecdsa"
	"flag"
	"fmt"
	"math/big"
	"os"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"gitlab.com/monetha/protocol-go-sdk/cmd/internal/cmdutils"
	"gitlab.com/monetha/protocol-go-sdk/deployer"
	"gitlab.com/monetha/protocol-go-sdk/eth"
	"gitlab.com/monetha/protocol-go-sdk/eth/backend"
	"gitlab.com/monetha/protocol-go-sdk/pass"
)

type boolFlag struct {
	set   bool
	value bool
}

func (f *boolFlag) String() string {
	return strconv.FormatBool(f.value)
}

func (f *boolFlag) Set(s string) error {
	v, err := strconv.ParseBool(s)
	f.value = v
	f.set = true
	return err
}

type addressFlag struct {
	set   bool
	value common.Address
}

func (f *addressFlag) String() string {
	return f.value.Hex()
}

func (f *addressFlag) Set(s string) error {
	if !common.IsHexAddress(s) {
		return fmt.Errorf("invalid Ethereum address: %v", s)
	}
	f.value = common.HexToAddress(s)
	f.set = true
	return nil
}

func flagBool(name string, value bool, usage string) *boolFlag {
	v := &boolFlag{value: value}
	flag.Var(v, name, usage)
	return v
}

func flagAddress(name string, value common.Address, usage string) *addressFlag {
	v := &addressFlag{value: value}
	flag.Var(v, name, usage)
	return v
}

func main() {
	var (
		backendURL         = flag.String("backendurl", "", "backend URL (simulated backend used if empty)")
		passAddr           = flag.String("passaddr", "", "Ethereum address of passport contract")
		ownerKeyFile       = flag.String("ownerkey", "", "owner private key filename")
		ownerKeyHex        = flag.String("ownerkeyhex", "", "private key as hex (for testing)")
		addFactProvider    = flagAddress("add", common.Address{}, "add fact provider to the whitelist")
		removeFactProvider = flagAddress("remove", common.Address{}, "remove fact provider from the whitelist")
		onlyWhitelist      = flagBool("onlywhitelist", false, "enables or disables the use of a whitelist of fact providers")
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
	case *passAddr == "" && *backendURL != "":
		utils.Fatalf("Use -passaddr to specify an address of passport contract")
	case *ownerKeyFile == "" && *ownerKeyHex == "":
		utils.Fatalf("Use -ownerkey or -ownerkeyhex to specify a private key")
	case *ownerKeyFile != "" && *ownerKeyHex != "":
		utils.Fatalf("Options -ownerkey or -ownerkeyhex are mutually exclusive")
	case !addFactProvider.set && !removeFactProvider.set && !onlyWhitelist.set:
		utils.Fatalf("Use -add, -remove or -onlywhitelist to specify what to do")
	case (addFactProvider.set && removeFactProvider.set) ||
		(addFactProvider.set && onlyWhitelist.set) ||
		(removeFactProvider.set && onlyWhitelist.set):
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

	passportAddress := common.HexToAddress(*passAddr)
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
	case addFactProvider.set:
		log.Warn("adding provider to whitelist", "fact_provider", addFactProvider.value)
		cmdutils.CheckErr(ignoreHash(w.AddFactProviderToWhitelist(ctx, addFactProvider.value)), "AddFactProviderToWhitelist")
	case removeFactProvider.set:
		log.Warn("removing provider from whitelist", "fact_provider", removeFactProvider.value)
		cmdutils.CheckErr(ignoreHash(w.RemoveFactProviderFromWhitelist(ctx, removeFactProvider.value)), "RemoveFactProviderFromWhitelist")
	case onlyWhitelist.set:
		log.Warn("set whitelist only permission", "value", onlyWhitelist.value)
		cmdutils.CheckErr(ignoreHash(w.SetWhitelistOnlyPermission(ctx, onlyWhitelist.value)), "SetWhitelistOnlyPermission")
	}

	log.Warn("Done.")
}

func ignoreHash(_ common.Hash, err error) error {
	return err
}
