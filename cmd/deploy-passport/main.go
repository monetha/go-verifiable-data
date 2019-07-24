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
	"github.com/ethereum/go-ethereum/log"
	"github.com/monetha/go-verifiable-data/cmd"
	"github.com/monetha/go-verifiable-data/cmd/internal/cmdutils"
	internalFlag "github.com/monetha/go-verifiable-data/cmd/internal/flag"
	"github.com/monetha/go-verifiable-data/deployer"
	"github.com/monetha/go-verifiable-data/eth"
	"github.com/monetha/go-verifiable-data/eth/backend"
)

func main() {
	var (
		backendURL       = flag.String("backendurl", "", "backend URL (simulated backend used if empty)")
		factoryAddr      = flag.String("factoryaddr", "", "Ethereum address of passport factory contract")
		ownerKeyFile     = flag.String("ownerkey", "", "owner private key filename")
		ownerKeyHex      = flag.String("ownerkeyhex", "", "private key as hex (for testing)")
		verbosity        = flag.Int("verbosity", int(log.LvlWarn), "log verbosity (0-9)")
		vmodule          = flag.String("vmodule", "", "log verbosity pattern")
		quorumPrivateFor = flag.String("quorum_privatefor", "", "Quorum nodes public keys to make transaction private for, separated by commas")
		quorumEnclave    = flag.String("quorum_enclave", "", "Quorum enclave url for private transactions")

		ownerKey *ecdsa.PrivateKey
		err      error
	)
	flag.Parse()

	if cmd.HasPrintedVersion() {
		return
	}

	privateFor := &internalFlag.StringArray{}
	privateFor.UnmarshalFlag(*quorumPrivateFor)
	bf := cmdutils.NewBackendFactory(quorumEnclave, privateFor.AsStringArr())

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
		e *eth.Eth
	)
	if *backendURL == "" {
		monethaKey, err := crypto.GenerateKey()
		cmdutils.CheckErr(err, "generating key")
		monethaAddress := bind.NewKeyedTransactor(monethaKey).From

		alloc := core.GenesisAlloc{
			monethaAddress: {Balance: big.NewInt(deployer.PassportFactoryGasLimit)},
			ownerAddress:   {Balance: big.NewInt(deployer.PassportGasLimit)},
		}
		sim := backend.NewSimulatedBackendExtended(alloc, 10000000)
		sim.Commit()

		e = eth.New(sim, log.Warn)
		cmdutils.CheckErr(e.UpdateSuggestedGasPrice(ctx), "SuggestGasPrice")

		// creating owner session and checking balance
		monethaSession := e.NewSession(monethaKey)
		cmdutils.CheckBalance(ctx, monethaSession, deployer.PassportFactoryGasLimit)

		// deploying passport factory
		res, err := deployer.New(monethaSession).Bootstrap(ctx)
		cmdutils.CheckErr(err, "create passport factory")

		passportFactoryAddress = res.PassportFactoryAddress
	} else {
		client, err := bf.DialBackend(*backendURL)
		cmdutils.CheckErr(err, "bf.DialBackend")

		e = bf.NewEth(ctx, client)
		cmdutils.CheckErr(e.UpdateSuggestedGasPrice(ctx), "SuggestGasPrice")
	}

	e = e.NewHandleNonceBackend([]common.Address{ownerAddress})

	// creating owner session and checking balance
	ownerSession := e.NewSession(ownerKey)
	cmdutils.CheckBalance(ctx, ownerSession, deployer.PassportGasLimit)

	// deploying passport
	_, err = deployer.New(ownerSession).
		DeployPassport(ctx, passportFactoryAddress)
	cmdutils.CheckErr(err, "deploying passport")

	log.Warn("Done.")
}
