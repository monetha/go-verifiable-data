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
	"github.com/monetha/reputation-go-sdk/cmd"
	"github.com/monetha/reputation-go-sdk/cmd/internal/cmdutils"
	"github.com/monetha/reputation-go-sdk/deployer"
	"github.com/monetha/reputation-go-sdk/eth"
	"github.com/monetha/reputation-go-sdk/eth/backend"
)

func main() {
	var (
		backendURL   = flag.String("backendurl", "", "backend URL (simulated backend used if empty)")
		versionStr   = flag.String("version", "", "The version of new passport logic contract (which will be deployed)")
		registryAddr = flag.String("registryaddr", "", "Ethereum address of passport logic registry contract")
		ownerKeyFile = flag.String("ownerkey", "", "Monetha owner private key filename")
		ownerKeyHex  = flag.String("ownerkeyhex", "", "Monetha owner private key as hex (for testing)")
		verbosity    = flag.Int("verbosity", int(log.LvlWarn), "log verbosity (0-9)")
		vmodule      = flag.String("vmodule", "", "log verbosity pattern")

		ownerKey *ecdsa.PrivateKey
		err      error
	)
	flag.Parse()

	if cmd.HasPrintedVersion() {
		return
	}

	glogger := log.NewGlogHandler(log.StreamHandler(os.Stderr, log.TerminalFormat(false)))
	glogger.Verbosity(log.Lvl(*verbosity))
	glogger.Vmodule(*vmodule)
	log.Root().SetHandler(glogger)

	switch {
	case *versionStr == "":
		utils.Fatalf("Use -version to specify the version of new passport logic contract")
	case *registryAddr == "" && *backendURL != "":
		utils.Fatalf("Use -registryaddr to specify an address of passport logic registry contract")
	case *ownerKeyFile == "" && *ownerKeyHex == "":
		utils.Fatalf("Use -ownerkey or -ownerkeyhex to specify a private key")
	case *ownerKeyFile != "" && *ownerKeyHex != "":
		utils.Fatalf("Options -ownerkey or -ownerkeyhex are mutually exclusive")
	case *ownerKeyFile != "":
		if ownerKey, err = crypto.LoadECDSA(*ownerKeyFile); err != nil {
			utils.Fatalf("-ownerkey: %v", err)
		}
	case *ownerKeyHex != "":
		if ownerKey, err = crypto.HexToECDSA(*ownerKeyHex); err != nil {
			utils.Fatalf("-ownerkeyhex: %v", err)
		}
	}

	passportRegistryAddress := common.HexToAddress(*registryAddr)
	ownerAddress := bind.NewKeyedTransactor(ownerKey).From
	log.Warn("Loaded configuration", "owner_address", ownerAddress.Hex(), "backend_url", *backendURL, "registry", passportRegistryAddress.Hex())

	ctx := cmdutils.CreateCtrlCContext()

	var (
		e *eth.Eth
	)
	if *backendURL == "" {
		alloc := core.GenesisAlloc{
			ownerAddress: {Balance: big.NewInt(deployer.PassportFactoryGasLimit +
				deployer.PassportLogicDeployGasLimit +
				deployer.AddPassportLogicGasLimit +
				deployer.SetCurrentPassportLogicGasLimit)},
		}
		sim := backend.NewSimulatedBackendExtended(alloc, 10000000)
		sim.Commit()

		e = eth.New(sim, log.Warn)
		cmdutils.CheckErr(e.UpdateSuggestedGasPrice(ctx), "SuggestGasPrice")

		// creating owner session and checking balance
		monethaSession := e.NewSession(ownerKey)
		cmdutils.CheckBalance(ctx, monethaSession, deployer.PassportFactoryGasLimit)

		// deploying passport factory
		res, err := deployer.New(monethaSession).Bootstrap(ctx)
		cmdutils.CheckErr(err, "create passport factory")
		passportRegistryAddress = res.PassportLogicRegistryAddress
	} else {
		client, err := ethclient.Dial(*backendURL)
		cmdutils.CheckErr(err, "ethclient.Dial")

		e = eth.New(client, log.Warn)
		cmdutils.CheckErr(e.UpdateSuggestedGasPrice(ctx), "SuggestGasPrice")
	}

	e = e.NewHandleNonceBackend([]common.Address{ownerAddress})

	// creating owner session and checking balance
	ownerSession := e.NewSession(ownerKey)

	// upgrading passport logic passport
	_, err = deployer.New(ownerSession).
		UpgradePassportLogic(ctx, *versionStr, passportRegistryAddress)
	cmdutils.CheckErr(err, "upgrading passport logic")

	log.Warn("Done.")
}
