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
	"gitlab.com/monetha/protocol-go-sdk/cmd/cmdutils"
	"gitlab.com/monetha/protocol-go-sdk/contracts"
)

var (
	oneEthInWei = big.NewInt(1000000000000000000)
)

func main() {
	var (
		backendURL   = flag.String("backendurl", "", "backend URL (simulated backend used if empty)")
		registryAddr = flag.String("registryaddr", "", "Ethereum address of passport logic registry contract")
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
	case *registryAddr == "":
		utils.Fatalf("Use -registryaddr to specify a private key")
	case *ownerKeyFile == "" && *ownerKeyHex == "":
		utils.Fatalf("Use -ownerkey or -ownerkeyhex to specify an address of passport logic registry contract")
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

	passportLogicRegistryAddress := common.HexToAddress(*registryAddr)
	ownerAuth := bind.NewKeyedTransactor(ownerKey)
	ownerAddress := cmdutils.KeyAddress(ownerKey, "invalid owner key")
	log.Warn("Loaded configuration", "owner_address", ownerAddress.Hex(), "backend_url", *backendURL, "registry", passportLogicRegistryAddress.Hex())

	var contractBackend chequebook.Backend
	if *backendURL == "" {
		alloc := core.GenesisAlloc{
			ownerAddress: {Balance: oneEthInWei},
		}
		sim := backends.NewSimulatedBackend(alloc, 10000000)
		sim.Commit()
		contractBackend = sim

	} else {
		contractBackend, err = ethclient.Dial(*backendURL)
		if err != nil {
			utils.Fatalf("dial backend %v", err)
		}
	}

	contractBackend = backend.NewHandleNonceBackend(contractBackend, []common.Address{ownerAddress})

	ctx := cmdutils.CreateCtrlCContext()

	///////////////////////////////////////////////////////
	// Check budgets
	///////////////////////////////////////////////////////

	cmdutils.CheckBalance(ctx, contractBackend, ownerAddress, oneEthInWei)

	///////////////////////////////////////////////////////
	// Deploying Passport
	///////////////////////////////////////////////////////

	log.Warn("Deploying Passport", "owner_address", ownerAddress, "registry", passportLogicRegistryAddress.Hex())
	passportAddress, tx, _, err := contracts.DeployPassportContract(ownerAuth, contractBackend, passportLogicRegistryAddress)
	cmdutils.CheckErr(err, "deployment Passport contract")
	cmdutils.CheckTx(ctx, contractBackend, tx.Hash())

	log.Warn("Passport deployed", "contract_address", passportAddress.Hex())

	log.Warn("Done.")
}
