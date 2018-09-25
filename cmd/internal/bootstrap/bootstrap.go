package bootstrap

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/contracts/chequebook"
	"github.com/ethereum/go-ethereum/log"
	"gitlab.com/monetha/protocol-go-sdk/cmd/internal/cmdutils"
	"gitlab.com/monetha/protocol-go-sdk/contracts"
)

// PassportFactory deploys PassportFactory contract and all contracts needed in order to deploy it
func PassportFactory(ctx context.Context, contractBackend chequebook.Backend, ownerAuth *bind.TransactOpts) common.Address {
	///////////////////////////////////////////////////////
	// PassportLogic
	///////////////////////////////////////////////////////

	log.Warn("Deploying PassportLogic", "owner_address", ownerAuth.From)
	passportLogicAddress, tx, passportLogicContract, err := contracts.DeployPassportLogicContract(ownerAuth, contractBackend)
	cmdutils.CheckErr(err, "deployment PassportLogic contract")
	cmdutils.CheckTx(ctx, contractBackend, tx.Hash())

	log.Warn("PassportLogic deployed", "contract_address", passportLogicAddress.Hex())
	_ = passportLogicContract

	///////////////////////////////////////////////////////
	// PassportLogicRegistry
	///////////////////////////////////////////////////////

	version := "0.1"
	log.Warn("Deploying PassportLogicRegistry", "owner_address", ownerAuth.From, "impl_version", version, "impl_address", passportLogicAddress)
	passportLogicRegistryAddress, tx, passportLogicRegistryContract, err := contracts.DeployPassportLogicRegistryContract(ownerAuth, contractBackend, version, passportLogicAddress)
	cmdutils.CheckErr(err, "deployment PassportLogicRegistry contract")
	cmdutils.CheckTx(ctx, contractBackend, tx.Hash())

	log.Warn("PassportLogicRegistry deployed", "contract_address", passportLogicRegistryAddress.Hex())
	_ = passportLogicRegistryContract

	///////////////////////////////////////////////////////
	// PassportFactory
	///////////////////////////////////////////////////////

	log.Warn("Deploying PassportFactory", "owner_address", ownerAuth.From, "registry", passportLogicRegistryAddress)
	passportFactoryAddress, tx, passportFactoryContract, err := contracts.DeployPassportFactoryContract(ownerAuth, contractBackend, passportLogicRegistryAddress)
	cmdutils.CheckErr(err, "deployment PassportFactory contract")
	cmdutils.CheckTx(ctx, contractBackend, tx.Hash())

	log.Warn("PassportFactory deployed", "contract_address", passportFactoryAddress.Hex())
	_ = passportFactoryContract

	return passportFactoryAddress
}
