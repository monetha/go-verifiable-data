package main

import (
	"crypto/ecdsa"
	"flag"
	"io/ioutil"
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
	"gitlab.com/monetha/protocol-go-sdk/deployer"
	"gitlab.com/monetha/protocol-go-sdk/eth"
	"gitlab.com/monetha/protocol-go-sdk/eth/backend"
	"gitlab.com/monetha/protocol-go-sdk/facts"
)

func main() {
	var (
		backendURL   = flag.String("backendurl", "", "backend URL (simulated backend used if empty)")
		passportAddr = flag.String("passportaddr", "", "Ethereum address of passport contract")
		factKeyStr   = flag.String("factkey", "", "the key of the fact (max. 32 bytes)")
		fileName     = flag.String("filename", "", "filename of the file with fact data")
		ownerKeyFile = flag.String("ownerkey", "", "fact provider private key filename")
		ownerKeyHex  = flag.String("ownerkeyhex", "", "fact provider private key as hex (for testing)")
		verbosity    = flag.Int("verbosity", int(log.LvlWarn), "log verbosity (0-9)")
		vmodule      = flag.String("vmodule", "", "log verbosity pattern")

		factProviderKey *ecdsa.PrivateKey
		factKey         [32]byte
		factData        []byte
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
	case *fileName == "":
		utils.Fatalf("Use -filename to specify the file name of the file with fact data")
	case *factKeyStr == "":
		utils.Fatalf("Use -factkey to specify the key of the fact")
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

	if factKeyBytes := []byte(*factKeyStr); len(factKeyBytes) < 32 {
		copy(factKey[:], factKeyBytes)
	} else {
		utils.Fatalf("The key string should fit into 32 bytes")
	}

	if factData, err = ioutil.ReadFile(*fileName); err != nil {
		utils.Fatalf("failed to read fact data file: %v", err)
	}

	passportAddress := common.HexToAddress(*passportAddr)
	factProviderAddress := bind.NewKeyedTransactor(factProviderKey).From
	log.Warn("Loaded configuration", "fact_provider", factProviderAddress.Hex(), "backend_url", *backendURL, "passport", passportAddress.Hex())

	ctx := cmdutils.CreateCtrlCContext()

	var (
		e *eth.Eth
	)
	if *backendURL == "" {
		monethaKey, err := crypto.GenerateKey()
		cmdutils.CheckErr(err, "generating key")
		monethaAddress := bind.NewKeyedTransactor(monethaKey).From

		passportOwnerKey, err := crypto.GenerateKey()
		cmdutils.CheckErr(err, "generating key")
		passportOwnerAddress := bind.NewKeyedTransactor(passportOwnerKey).From

		alloc := core.GenesisAlloc{
			monethaAddress:       {Balance: big.NewInt(deployer.PassportFactoryGasLimit)},
			passportOwnerAddress: {Balance: big.NewInt(deployer.PassportGasLimit)},
			factProviderAddress:  {Balance: big.NewInt(10000000000000)},
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

		// creating passport owner session and checking balance
		passportOwnerSession := e.NewSession(passportOwnerKey)
		cmdutils.CheckBalance(ctx, passportOwnerSession, deployer.PassportGasLimit)

		// deploying passport
		passportAddress, err = deployer.New(passportOwnerSession).DeployPassport(ctx, passportFactoryAddress)
		cmdutils.CheckErr(err, "create passport")
	} else {
		client, err := ethclient.Dial(*backendURL)
		cmdutils.CheckErr(err, "ethclient.Dial")

		e = eth.New(client, log.Warn)
		cmdutils.CheckErr(e.UpdateSuggestedGasPrice(ctx), "SuggestGasPrice")
	}

	factProviderSession := e.NewSession(factProviderKey)

	// TODO: check balance

	err = facts.NewProvider(factProviderSession).
		WriteTxData(ctx, passportAddress, factKey, factData)
	cmdutils.CheckErr(err, "WriteTxData")

	log.Warn("Done.")
}
