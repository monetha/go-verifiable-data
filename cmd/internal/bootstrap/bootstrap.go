package bootstrap

import (
	"context"
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

// CreatePassportFactory deploys PassportFactory contract and all contracts needed in order to deploy it
func (b Bootstrap) CreatePassportFactory(ctx context.Context, contractBackend chequebook.Backend, ownerAuth *bind.TransactOpts) (common.Address, error) {
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

func (b Bootstrap) log(msg string, ctx ...interface{}) {
	l := b.Log
	if l != nil {
		l(msg, ctx...)
	}
}
