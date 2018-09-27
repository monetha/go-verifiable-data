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
	"gitlab.com/monetha/protocol-go-sdk/cmd/internal/cmdutils"
	"gitlab.com/monetha/protocol-go-sdk/deploy"
	"gitlab.com/monetha/protocol-go-sdk/eth"
	"gitlab.com/monetha/protocol-go-sdk/eth/backend"
)

func main() {
	var (
		backendURL   = flag.String("backendurl", "", "backend URL (simulated backend used if empty)")
		factoryAddr  = flag.String("factoryaddr", "", "Ethereum address of passport factory contract")
		ownerKeyFile = flag.String("ownerkey", "", "owner private key filename")
		ownerKeyHex  = flag.String("ownerkeyhex", "", "private key as hex (for testing)")
		verbosity    = flag.Int("verbosity", int(log.LvlWarn), "log verbosity (0-9)")
		vmodule      = flag.String("vmodule", "", "log verbosity pattern")

		ownerKey *ecdsa.PrivateKey
		err      error
	)
	flag.Parse()

	glogger := log.NewGlogHandler(log.StreamHandler(os.Stderr, log.TerminalFormat(false)))
	glogger.Verbosity(log.Lvl(*verbosity))
	glogger.Vmodule(*vmodule)
	log.Root().SetHandler(glogger)

	switch {
	case *factoryAddr == "" && *backendURL != "":
		utils.Fatalf("Use -factoryaddr to specify an address of passport factory contract")
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

	passportFactoryAddress := common.HexToAddress(*factoryAddr)
	ownerAddress := bind.NewKeyedTransactor(ownerKey).From
	log.Warn("Loaded configuration", "owner_address", ownerAddress.Hex(), "backend_url", *backendURL, "factory", passportFactoryAddress.Hex())

	ctx := cmdutils.CreateCtrlCContext()

	var (
		contractBackend backend.Backend
		gasPrice        *big.Int
	)
	if *backendURL == "" {
		monethaKey, err := crypto.GenerateKey()
		cmdutils.CheckErr(err, "generating key")
		monethaAddress := bind.NewKeyedTransactor(monethaKey).From

		alloc := core.GenesisAlloc{
			monethaAddress: {Balance: big.NewInt(deploy.PassportFactoryGasLimit)},
			ownerAddress:   {Balance: big.NewInt(deploy.PassportGasLimit)},
		}
		sim := backend.NewSimulatedBackendExtended(alloc, 10000000)
		sim.Commit()

		contractBackend = sim

		// retrieving suggested gas price
		gasPrice, err = contractBackend.SuggestGasPrice(ctx)
		cmdutils.CheckErr(err, "SuggestGasPrice")

		// creating owner session and checking balance
		monethaSession := eth.NewSession(contractBackend, monethaKey).SetGasPrice(gasPrice).SetLogFun(log.Warn)
		cmdutils.CheckBalance(ctx, monethaSession, deploy.PassportFactoryGasLimit)

		// deploying passport factory
		passportFactoryAddress, err = deploy.New(monethaSession).DeployPassportFactory(ctx)
		cmdutils.CheckErr(err, "create passport factory")

	} else {
		contractBackend, err = ethclient.Dial(*backendURL)
		if err != nil {
			utils.Fatalf("dial backend %v", err)
		}

		// retrieving suggested gas price
		gasPrice, err = contractBackend.SuggestGasPrice(ctx)
		cmdutils.CheckErr(err, "SuggestGasPrice")
	}

	// creating owner session and checking balance
	ownerSession := eth.NewSession(contractBackend, ownerKey).SetGasPrice(gasPrice).SetLogFun(log.Warn)
	cmdutils.CheckBalance(ctx, ownerSession, deploy.PassportGasLimit)

	// deploy passport
	_, err = deploy.New(ownerSession).
		DeployPassport(ctx, passportFactoryAddress)
	cmdutils.CheckErr(err, "deploying passport")

	log.Warn("Done.")
}
