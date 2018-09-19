package main

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"flag"
	"fmt"
	"math/big"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/contracts/chequebook"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/monetha/go-ethereum/backend"
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
	ownerAddress := keyAddress(ownerKey, "invalid owner key")
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

	ctx := createCtrlCContext()

	///////////////////////////////////////////////////////
	// Check budgets
	///////////////////////////////////////////////////////

	checkBalance(ctx, contractBackend, ownerAddress, oneEthInWei)

	///////////////////////////////////////////////////////
	// PassportLogic
	///////////////////////////////////////////////////////

	log.Warn("Deploying PassportLogic", "owner_address", ownerAddress)
	passportLogicAddress, tx, passportLogicContract, err := contracts.DeployPassportLogicContract(ownerAuth, contractBackend)
	checkErr(err, "deployment PassportLogic contract")
	checkTx(ctx, contractBackend, tx.Hash())

	log.Warn("PassportLogic deployed", "contract_address", passportLogicAddress.Hex())
	_ = passportLogicContract

	///////////////////////////////////////////////////////
	// PassportLogicRegistry
	///////////////////////////////////////////////////////

	version := "0.1"
	log.Warn("Deploying PassportLogicRegistry", "owner_address", ownerAddress, "impl_version", version, "impl_address", passportLogicAddress)
	passportLogicRegistryAddress, tx, passportLogicRegistryContract, err := contracts.DeployPassportLogicRegistryContract(ownerAuth, contractBackend, version, passportLogicAddress)
	checkErr(err, "deployment PassportLogicRegistry contract")
	checkTx(ctx, contractBackend, tx.Hash())

	log.Warn("PassportLogicRegistry deployed", "contract_address", passportLogicRegistryAddress.Hex())
	_ = passportLogicRegistryContract

	///////////////////////////////////////////////////////
	// PassportFactory
	///////////////////////////////////////////////////////

	log.Warn("Deploying PassportFactory", "owner_address", ownerAddress, "registry", passportLogicRegistryAddress)
	passportFactoryAddress, tx, passportFactoryContract, err := contracts.DeployPassportFactoryContract(ownerAuth, contractBackend, passportLogicRegistryAddress)
	checkErr(err, "deployment PassportFactory contract")
	checkTx(ctx, contractBackend, tx.Hash())

	log.Warn("PassportFactory deployed", "contract_address", passportFactoryAddress.Hex())
	_ = passportFactoryContract

	log.Warn("Done.")
}

func checkErr(err error, hint string) {
	if err != nil {
		log.Error(hint, "err", err)
		os.Exit(1)
	}
}

func checkBalance(ctx context.Context, backend chequebook.Backend, address common.Address, minWei *big.Int) {
	log.Warn("Checking balance", "address", address.Hex())

	balance, err := backend.BalanceAt(ctx, address, nil)
	hint := fmt.Sprintf("BalanceAt(%v)", address.Hex())
	checkErr(err, hint)

	if balance.Cmp(minWei) == -1 {
		checkErr(fmt.Errorf("balance too low: %v wei < %v wei", balance, minWei), hint)
	}
}

func checkTx(ctx context.Context, backend chequebook.Backend, txHash common.Hash) {
	err := waitForTx(ctx, backend, txHash)
	checkErr(err, fmt.Sprintf("tx 0x%x", txHash))
}

func waitForTx(ctx context.Context, backend chequebook.Backend, txHash common.Hash) error {
	log.Warn("Waiting for transaction", "hash", txHash.Hex())

	type commiter interface {
		Commit()
	}
	if sim, ok := backend.(commiter); ok {
		sim.Commit()
		tr, err := backend.TransactionReceipt(ctx, txHash)
		if err != nil {
			return err
		}
		if tr.Status != types.ReceiptStatusSuccessful {
			return fmt.Errorf("tx failed: %+v", tr)
		}
		return nil
	}

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(4 * time.Second):
		}

		tr, err := backend.TransactionReceipt(ctx, txHash)
		if err != nil {
			if err == ethereum.NotFound {
				continue
			} else {
				return err
			}
		} else {
			if tr.Status != types.ReceiptStatusSuccessful {
				return fmt.Errorf("tx failed: %+v", tr)
			}
			return nil
		}
	}
}

func keyAddress(privateKeyECDSA *ecdsa.PrivateKey, errorMsg string) common.Address {
	addr, err := pubkeyToAddress(privateKeyECDSA.PublicKey)
	if err != nil {
		utils.Fatalf(errorMsg, err)
	}
	return addr
}

func pubkeyToAddress(p ecdsa.PublicKey) (common.Address, error) {
	pubBytes := crypto.FromECDSAPub(&p)
	if pubBytes == nil {
		return common.Address{}, errors.New("invalid key")
	}
	return common.Address(common.BytesToAddress(crypto.Keccak256(pubBytes[1:])[12:])), nil
}

func createCtrlCContext() context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		defer cancel()

		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT)
		<-sigChan
		log.Warn("got interrupt signal")
	}()

	return ctx
}
