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

	ownerAuth := bind.NewKeyedTransactor(ownerKey)
	ownerAddress := cmdutils.KeyAddress(ownerKey, "invalid owner key")
	log.Warn("Loaded configuration", "owner_address", ownerAddress.Hex(), "backend_url", *backendURL)

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
	// PassportLogic
	///////////////////////////////////////////////////////

	log.Warn("Deploying PassportLogic", "owner_address", ownerAddress)
	passportLogicAddress, tx, passportLogicContract, err := contracts.DeployPassportLogicContract(ownerAuth, contractBackend)
	cmdutils.CheckErr(err, "deployment PassportLogic contract")
	cmdutils.CheckTx(ctx, contractBackend, tx.Hash())

	log.Warn("PassportLogic deployed", "contract_address", passportLogicAddress.Hex())
	_ = passportLogicContract

	///////////////////////////////////////////////////////
	// PassportLogicRegistry
	///////////////////////////////////////////////////////

	version := "0.1"
	log.Warn("Deploying PassportLogicRegistry", "owner_address", ownerAddress, "impl_version", version, "impl_address", passportLogicAddress)
	passportLogicRegistryAddress, tx, passportLogicRegistryContract, err := contracts.DeployPassportLogicRegistryContract(ownerAuth, contractBackend, version, passportLogicAddress)
	cmdutils.CheckErr(err, "deployment PassportLogicRegistry contract")
	cmdutils.CheckTx(ctx, contractBackend, tx.Hash())

	log.Warn("PassportLogicRegistry deployed", "contract_address", passportLogicRegistryAddress.Hex())
	_ = passportLogicRegistryContract

	///////////////////////////////////////////////////////
	// PassportFactory
	///////////////////////////////////////////////////////

	log.Warn("Deploying PassportFactory", "owner_address", ownerAddress, "registry", passportLogicRegistryAddress)
	passportFactoryAddress, tx, passportFactoryContract, err := contracts.DeployPassportFactoryContract(ownerAuth, contractBackend, passportLogicRegistryAddress)
	cmdutils.CheckErr(err, "deployment PassportFactory contract")
	cmdutils.CheckTx(ctx, contractBackend, tx.Hash())

	log.Warn("PassportFactory deployed", "contract_address", passportFactoryAddress.Hex())
	_ = passportFactoryContract

	log.Warn("Done.")
}
