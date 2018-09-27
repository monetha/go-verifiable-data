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
	"gitlab.com/monetha/protocol-go-sdk/facts"
)

func main() {
	var (
		backendURL   = flag.String("backendurl", "", "backend URL (simulated backend used if empty)")
		passportAddr = flag.String("passportaddr", "", "Ethereum address of passport contract")
		ownerKeyFile = flag.String("ownerkey", "", "fact provider private key filename")
		ownerKeyHex  = flag.String("ownerkeyhex", "", "fact provider private key as hex (for testing)")
		verbosity    = flag.Int("verbosity", int(log.LvlWarn), "log verbosity (0-9)")
		vmodule      = flag.String("vmodule", "", "log verbosity pattern")

		factProviderKey *ecdsa.PrivateKey
		err             error
	)
	flag.Parse()

	glogger := log.NewGlogHandler(log.StreamHandler(os.Stderr, log.TerminalFormat(false)))
	glogger.Verbosity(log.Lvl(*verbosity))
	glogger.Vmodule(*vmodule)
	log.Root().SetHandler(glogger)

	switch {
	case *passportAddr == "" && *backendURL != "":
		utils.Fatalf("Use -passportaddr to specify an address of passport contract")
	case *ownerKeyFile == "" && *ownerKeyHex == "":
		utils.Fatalf("Use -ownerkey or -ownerkeyhex to specify a private key of fact provider")
	case *ownerKeyFile != "" && *ownerKeyHex != "":
		utils.Fatalf("Options -ownerkey or -ownerkeyhex are mutually exclusive")
	case *ownerKeyFile != "":
		if factProviderKey, err = crypto.LoadECDSA(*ownerKeyFile); err != nil {
			utils.Fatalf("-ownerkey: %v", err)
		}
	case *ownerKeyHex != "":
		if factProviderKey, err = crypto.HexToECDSA(*ownerKeyHex); err != nil {
			utils.Fatalf("-ownerkeyhex: %v", err)
		}
	}

	passportAddress := common.HexToAddress(*passportAddr)
	factProviderAddress := bind.NewKeyedTransactor(factProviderKey).From
	log.Warn("Loaded configuration", "fact_provider", factProviderAddress.Hex(), "backend_url", *backendURL, "passport", passportAddress.Hex())

	ctx := cmdutils.CreateCtrlCContext()

	var (
		contractBackend backend.Backend
		gasPrice        *big.Int
	)
	if *backendURL == "" {
		monethaKey, err := crypto.GenerateKey()
		cmdutils.CheckErr(err, "generating key")
		monethaAddress := bind.NewKeyedTransactor(monethaKey).From

		passportOwnerKey, err := crypto.GenerateKey()
		cmdutils.CheckErr(err, "generating key")
		passportOwnerAddress := bind.NewKeyedTransactor(passportOwnerKey).From

		alloc := core.GenesisAlloc{
			monethaAddress:       {Balance: big.NewInt(deploy.PassportFactoryGasLimit)},
			passportOwnerAddress: {Balance: big.NewInt(deploy.PassportGasLimit)},
			factProviderAddress:  {Balance: big.NewInt(10000000000000)},
		}
		sim := backend.NewSimulatedBackendExtended(alloc, 10000000)
		sim.Commit()

		contractBackend = sim

		e := &eth.Eth{
			Backend: sim,
			LogFun:  log.Warn,
		}

		// retrieving suggested gas price
		gasPrice, err = contractBackend.SuggestGasPrice(ctx)
		cmdutils.CheckErr(err, "SuggestGasPrice")

		// creating owner session and checking balance
		monethaSession := eth.NewSession(e, monethaKey).SetGasPrice(gasPrice)
		cmdutils.CheckBalance(ctx, monethaSession, deploy.PassportFactoryGasLimit)

		// deploying passport factory
		passportFactoryAddress, err := deploy.New(monethaSession).DeployPassportFactory(ctx)
		cmdutils.CheckErr(err, "create passport factory")

		// creating passport owner session and checking balance
		passportOwnerSession := eth.NewSession(e, passportOwnerKey).SetGasPrice(gasPrice)
		cmdutils.CheckBalance(ctx, passportOwnerSession, deploy.PassportGasLimit)

		// deploying passport
		passportAddress, err = deploy.New(passportOwnerSession).DeployPassport(ctx, passportFactoryAddress)
	} else {
		contractBackend, err = ethclient.Dial(*backendURL)
		if err != nil {
			utils.Fatalf("dial backend %v", err)
		}

		// retrieving suggested gas price
		gasPrice, err = contractBackend.SuggestGasPrice(ctx)
		cmdutils.CheckErr(err, "SuggestGasPrice")
	}

	e := &eth.Eth{
		Backend: contractBackend,
		LogFun:  log.Warn,
	}

	factProviderSession := eth.NewSession(e, factProviderKey).SetGasPrice(gasPrice)

	// TODO: check balance

	// TODO: get key/data from command line
	var key [32]byte
	data := []byte("some text")

	err = facts.NewProvider(factProviderSession).
		WriteTxData(ctx, passportAddress, key, data)
	cmdutils.CheckErr(err, "WriteTxData")

	log.Warn("Done.")
}
