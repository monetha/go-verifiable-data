package main

import (
	"crypto/ecdsa"
	"flag"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/contracts/chequebook"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/monetha/go-ethereum/backend"
	"gitlab.com/monetha/protocol-go-sdk/cmd/internal/cmdutils"
	"gitlab.com/monetha/protocol-go-sdk/cmd/internal/deploy"
)

var (
	oneEthInWei = big.NewInt(1000000000000000000)
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
	ownerAuth := bind.NewKeyedTransactor(ownerKey)
	log.Warn("Loaded configuration", "owner_address", ownerAuth.From.Hex(), "backend_url", *backendURL, "factory", passportFactoryAddress.Hex())

	ctx := cmdutils.CreateCtrlCContext()

	d := deploy.Deploy{Log: log.Warn}
	var contractBackend chequebook.Backend
	if *backendURL == "" {
		alloc := core.GenesisAlloc{
			ownerAuth.From: {Balance: new(big.Int).Add(oneEthInWei, oneEthInWei)},
		}
		sim := backends.NewSimulatedBackend(alloc, 10000000)
		sim.Commit()

		contractBackend = sim

		passportFactoryAddress, err = d.DeployPassportFactory(ctx, contractBackend, ownerAuth)
		cmdutils.CheckErr(err, "create passport factory")

	} else {
		contractBackend, err = ethclient.Dial(*backendURL)
		if err != nil {
			utils.Fatalf("dial backend %v", err)
		}
	}

	contractBackend = backend.NewHandleNonceBackend(contractBackend, []common.Address{ownerAuth.From})

	cmdutils.CheckBalance(ctx, contractBackend, ownerAuth.From, oneEthInWei)

	_, err = d.DeployPassport(ctx, contractBackend, ownerAuth, passportFactoryAddress)
	cmdutils.CheckErr(err, "deploying passport")

	log.Warn("Done.")
}
