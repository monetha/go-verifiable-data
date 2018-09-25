package bootstrap

import (
	"context"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/contracts/chequebook"
	"gitlab.com/monetha/protocol-go-sdk/cmd/internal/eth"
	"gitlab.com/monetha/protocol-go-sdk/cmd/internal/log"
	"gitlab.com/monetha/protocol-go-sdk/contracts"
)

// Bootstrap contains methods to deploy contracts
type Bootstrap struct {
	Log log.Fun
}

// DeployPassportFactory deploys PassportFactory contract and all contracts needed in order to deploy it
func (b Bootstrap) DeployPassportFactory(ctx context.Context, contractBackend chequebook.Backend, ownerAuth *bind.TransactOpts) (common.Address, error) {
	e := eth.Eth{Log: b.Log}

	///////////////////////////////////////////////////////
	// PassportLogic
	///////////////////////////////////////////////////////

	b.log("Deploying PassportLogic", "owner_address", ownerAuth.From)
	passportLogicAddress, tx, _, err := contracts.DeployPassportLogicContract(ownerAuth, contractBackend)
	if err != nil {
		return common.Address{}, fmt.Errorf("deploying PassportLogic contract: %v", err)
	}
	_, err = e.WaitForTxReceipt(ctx, contractBackend, tx.Hash())
	if err != nil {
		return common.Address{}, err
	}

	b.log("PassportLogic deployed", "contract_address", passportLogicAddress.Hex())

	///////////////////////////////////////////////////////
	// PassportLogicRegistry
	///////////////////////////////////////////////////////

	version := "0.1"
	b.log("Deploying PassportLogicRegistry", "owner_address", ownerAuth.From, "impl_version", version, "impl_address", passportLogicAddress)
	passportLogicRegistryAddress, tx, _, err := contracts.DeployPassportLogicRegistryContract(ownerAuth, contractBackend, version, passportLogicAddress)
	if err != nil {
		return common.Address{}, fmt.Errorf("deploying PassportLogicRegistry contract: %v", err)
	}
	_, err = e.WaitForTxReceipt(ctx, contractBackend, tx.Hash())
	if err != nil {
		return common.Address{}, err
	}

	b.log("PassportLogicRegistry deployed", "contract_address", passportLogicRegistryAddress.Hex())

	///////////////////////////////////////////////////////
	// PassportFactory
	///////////////////////////////////////////////////////

	b.log("Deploying PassportFactory", "owner_address", ownerAuth.From, "registry", passportLogicRegistryAddress)
	passportFactoryAddress, tx, _, err := contracts.DeployPassportFactoryContract(ownerAuth, contractBackend, passportLogicRegistryAddress)
	if err != nil {
		return common.Address{}, fmt.Errorf("deploying PassportFactory contract: %v", err)
	}
	_, err = e.WaitForTxReceipt(ctx, contractBackend, tx.Hash())
	if err != nil {
		return common.Address{}, err
	}

	b.log("PassportFactory deployed", "contract_address", passportFactoryAddress.Hex())

	return passportFactoryAddress, nil
}

// DeployPassport deploys only Passport contract using existing PassportFactory contract
func (b Bootstrap) DeployPassport(ctx context.Context, contractBackend chequebook.Backend, ownerAuth *bind.TransactOpts, passportFactoryAddress common.Address) (common.Address, error) {
	e := eth.Eth{Log: b.Log}

	///////////////////////////////////////////////////////
	// PassportFactory
	///////////////////////////////////////////////////////

	b.log("Initializing PassportFactory contract", "factory", passportFactoryAddress.Hex())
	passportFactoryContract, err := contracts.NewPassportFactoryContract(passportFactoryAddress, contractBackend)
	if err != nil {
		return common.Address{}, fmt.Errorf("initializing PassportFactory contract: %v", err)
	}

	///////////////////////////////////////////////////////
	// Passport
	///////////////////////////////////////////////////////

	b.log("Deploying Passport contract", "owner_address", ownerAuth.From.Hex())
	tx, err := passportFactoryContract.CreatePassport(ownerAuth)
	if err != nil {
		return common.Address{}, fmt.Errorf("deploying Passport contract: %v", err)
	}
	txr, err := e.WaitForTxReceipt(ctx, contractBackend, tx.Hash())
	if err != nil {
		return common.Address{}, err
	}

	///////////////////////////////////////////////////////
	// PassportCreated events
	///////////////////////////////////////////////////////

	lf, err := contracts.NewPassportFactoryLogFilterer()
	if err != nil {
		return common.Address{}, fmt.Errorf("initializing PassportFactoryLogFilterer: %v", err)
	}

	evs, err := lf.FilterPassportCreated(txr.Logs, nil, nil)
	if err != nil {
		return common.Address{}, fmt.Errorf("filtering PassportCreated events: %v", err)
	}
	if len(evs) != 1 {
		return common.Address{}, errors.New("expected exactly one PassportCreated event")
	}

	passportAddress := evs[0].Passport
	b.log("Passport deployed", "contract_address", passportAddress.Hex())

	///////////////////////////////////////////////////////
	// Passport
	///////////////////////////////////////////////////////

	b.log("Initializing Passport contract", "passport", passportFactoryAddress.Hex())
	passportContract, err := contracts.NewPassportContract(passportAddress, contractBackend)
	if err != nil {
		return common.Address{}, fmt.Errorf("initializing Passport contract: %v", err)
	}

	///////////////////////////////////////////////////////
	// ClaimOwnership
	///////////////////////////////////////////////////////

	b.log("Claiming ownership", "owner_address", ownerAuth.From)
	tx, err = passportContract.ClaimOwnership(ownerAuth)
	if err != nil {
		return common.Address{}, fmt.Errorf("claiming ownership: %v", err)
	}
	_, err = e.WaitForTxReceipt(ctx, contractBackend, tx.Hash())
	if err != nil {
		return common.Address{}, err
	}

	return passportAddress, nil
}

func (b Bootstrap) log(msg string, ctx ...interface{}) {
	l := b.Log
	if l != nil {
		l(msg, ctx...)
	}
}
