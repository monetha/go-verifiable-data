package bootstrap

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/contracts/chequebook"
	"gitlab.com/monetha/protocol-go-sdk/cmd/internal/cmdutils"
	"gitlab.com/monetha/protocol-go-sdk/cmd/internal/log"
	"gitlab.com/monetha/protocol-go-sdk/contracts"
)

type Bootstrap struct {
	log log.Fun
}

// New returns an instance of Bootstrap
func New(lf log.Fun) *Bootstrap {
	return &Bootstrap{log: lf}
}

// CreatePassportFactory deploys PassportFactory contract and all contracts needed in order to deploy it
func (b *Bootstrap) CreatePassportFactory(ctx context.Context, contractBackend chequebook.Backend, ownerAuth *bind.TransactOpts) common.Address {
	///////////////////////////////////////////////////////
	// PassportLogic
	///////////////////////////////////////////////////////

	b.log("Deploying PassportLogic", "owner_address", ownerAuth.From)
	passportLogicAddress, tx, _, err := contracts.DeployPassportLogicContract(ownerAuth, contractBackend)
	cmdutils.CheckErr(err, "deployment PassportLogic contract")
	cmdutils.CheckTx(ctx, contractBackend, tx.Hash())

	b.log("PassportLogic deployed", "contract_address", passportLogicAddress.Hex())

	///////////////////////////////////////////////////////
	// PassportLogicRegistry
	///////////////////////////////////////////////////////

	version := "0.1"
	b.log("Deploying PassportLogicRegistry", "owner_address", ownerAuth.From, "impl_version", version, "impl_address", passportLogicAddress)
	passportLogicRegistryAddress, tx, _, err := contracts.DeployPassportLogicRegistryContract(ownerAuth, contractBackend, version, passportLogicAddress)
	cmdutils.CheckErr(err, "deployment PassportLogicRegistry contract")
	cmdutils.CheckTx(ctx, contractBackend, tx.Hash())

	b.log("PassportLogicRegistry deployed", "contract_address", passportLogicRegistryAddress.Hex())

	///////////////////////////////////////////////////////
	// PassportFactory
	///////////////////////////////////////////////////////

	b.log("Deploying PassportFactory", "owner_address", ownerAuth.From, "registry", passportLogicRegistryAddress)
	passportFactoryAddress, tx, _, err := contracts.DeployPassportFactoryContract(ownerAuth, contractBackend, passportLogicRegistryAddress)
	cmdutils.CheckErr(err, "deployment PassportFactory contract")
	cmdutils.CheckTx(ctx, contractBackend, tx.Hash())

	b.log("PassportFactory deployed", "contract_address", passportFactoryAddress.Hex())

	return passportFactoryAddress
}
