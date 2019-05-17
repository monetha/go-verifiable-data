package deployer

import (
	"context"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/monetha/reputation-go-sdk/contracts"
	"github.com/monetha/reputation-go-sdk/eth"
)

const (
	// PassportLogicDeployGasLimit is a gas limit to deploy only PassportLogic contract
	PassportLogicDeployGasLimit = 3300000
	// PassportLogicRegistryDeployGasLimit is a gas limit to deploy only PassportLogicRegistry contract
	PassportLogicRegistryDeployGasLimit = 1120000
	// PassportFactoryDeployGasLimit is a gas limit to deploy only PassportFactory contract
	PassportFactoryDeployGasLimit = 1020000
	// AddPassportLogicGasLimit is a gas limit to call AddPassportLogic of PassportLogicRegistry contract
	AddPassportLogicGasLimit = 50000
	// SetCurrentPassportLogicGasLimit is a gas limit to call SetCurrentPassportLogic of PassportLogicRegistry contract
	SetCurrentPassportLogicGasLimit = 50000

	// PassportFactoryGasLimit is a minimum gas amount needed to fully deployer passport factory contract with all dependent contracts
	PassportFactoryGasLimit = PassportLogicDeployGasLimit + PassportLogicRegistryDeployGasLimit + PassportFactoryDeployGasLimit

	// PassportGasLimit is a minimum gas amount needed to fully deployer passport contract
	PassportGasLimit = 460000
)

// Deploy contains methods to deployer contracts
type Deploy eth.Session

// New converts session to Deploy
func New(s *eth.Session) *Deploy {
	return (*Deploy)(s)
}

// DeployPassportFactoryResult hold result of DeployPassportFactory execution
type DeployPassportFactoryResult struct {
	PassportLogicAddress         common.Address
	PassportLogicRegistryAddress common.Address
	PassportFactoryAddress       common.Address
}

// DeployPassportFactory deploys PassportFactory contract and all contracts needed in order to deployer it
func (d *Deploy) DeployPassportFactory(ctx context.Context) (*DeployPassportFactoryResult, error) {
	backend := d.Backend
	ownerAuth := &d.TransactOpts

	///////////////////////////////////////////////////////
	// deploying PassportLogic
	///////////////////////////////////////////////////////

	d.Log("Deploying PassportLogic", "owner_address", ownerAuth.From)
	passportLogicAddress, tx, _, err := contracts.DeployPassportLogicContract(ownerAuth, backend)
	if err != nil {
		return nil, fmt.Errorf("deploying PassportLogic contract: %v", err)
	}
	_, err = d.WaitForTxReceipt(ctx, tx.Hash())
	if err != nil {
		return nil, err
	}

	d.Log("PassportLogic deployed", "contract_address", passportLogicAddress.Hex())

	///////////////////////////////////////////////////////
	// deploying PassportLogicRegistry
	///////////////////////////////////////////////////////

	version := "0.1"
	d.Log("Deploying PassportLogicRegistry", "owner_address", ownerAuth.From, "impl_version", version, "impl_address", passportLogicAddress)
	passportLogicRegistryAddress, tx, _, err := contracts.DeployPassportLogicRegistryContract(ownerAuth, backend, version, passportLogicAddress)
	if err != nil {
		return nil, fmt.Errorf("deploying PassportLogicRegistry contract: %v", err)
	}
	_, err = d.WaitForTxReceipt(ctx, tx.Hash())
	if err != nil {
		return nil, err
	}

	d.Log("PassportLogicRegistry deployed", "contract_address", passportLogicRegistryAddress.Hex())

	///////////////////////////////////////////////////////
	// deploying PassportFactory
	///////////////////////////////////////////////////////

	d.Log("Deploying PassportFactory", "owner_address", ownerAuth.From, "registry", passportLogicRegistryAddress)
	passportFactoryAddress, tx, _, err := contracts.DeployPassportFactoryContract(ownerAuth, backend, passportLogicRegistryAddress)
	if err != nil {
		return nil, fmt.Errorf("deploying PassportFactory contract: %v", err)
	}
	_, err = d.WaitForTxReceipt(ctx, tx.Hash())
	if err != nil {
		return nil, err
	}

	d.Log("PassportFactory deployed", "contract_address", passportFactoryAddress.Hex())

	return &DeployPassportFactoryResult{
		PassportLogicAddress:         passportLogicAddress,
		PassportLogicRegistryAddress: passportLogicRegistryAddress,
		PassportFactoryAddress:       passportFactoryAddress,
	}, nil
}

// UpgradePassportLogic deploys PassportLogic contract and sets it as current in PassportLogicRegistry
// Returns address of deployed PassportLogic contract
func (d *Deploy) UpgradePassportLogic(ctx context.Context, passportLogicVersion string, passportLogicRegistryAddress common.Address) (common.Address, error) {
	backend := d.Backend
	ownerAuth := &d.TransactOpts

	///////////////////////////////////////////////////////
	// deploying PassportLogic
	///////////////////////////////////////////////////////

	d.Log("Deploying PassportLogic", "owner_address", ownerAuth.From)
	passportLogicAddress, tx, _, err := contracts.DeployPassportLogicContract(ownerAuth, backend)
	if err != nil {
		return common.Address{}, fmt.Errorf("deploying PassportLogic contract: %v", err)
	}
	_, err = d.WaitForTxReceipt(ctx, tx.Hash())
	if err != nil {
		return common.Address{}, err
	}

	d.Log("PassportLogic deployed", "contract_address", passportLogicAddress.Hex())

	///////////////////////////////////////////////////////
	// updating PassportLogicRegistry
	///////////////////////////////////////////////////////

	d.Log("Initializing PassportLogicRegistry", "address", passportLogicRegistryAddress)
	passportLogicRegistryContract, err := contracts.NewPassportLogicRegistryContract(passportLogicRegistryAddress, backend)
	if err != nil {
		return common.Address{}, fmt.Errorf("initializing PassportLogicRegistry contract: %v", err)
	}

	d.Log("Adding new PassportLogic to registry", "registry_address", passportLogicRegistryAddress, "logic_address", passportLogicAddress, "version", passportLogicVersion)
	tx, err = passportLogicRegistryContract.AddPassportLogic(ownerAuth, passportLogicVersion, passportLogicAddress)
	if err != nil {
		return common.Address{}, fmt.Errorf("adding new PassportLogic to registry: %v", err)
	}
	_, err = d.WaitForTxReceipt(ctx, tx.Hash())
	if err != nil {
		return common.Address{}, err
	}

	d.Log("Setting current PassportLogic in registry", "registry_address", passportLogicRegistryAddress, "version", passportLogicVersion)
	tx, err = passportLogicRegistryContract.SetCurrentPassportLogic(ownerAuth, passportLogicVersion)
	if err != nil {
		return common.Address{}, fmt.Errorf("setting current PassportLogic in registry: %v", err)
	}
	_, err = d.WaitForTxReceipt(ctx, tx.Hash())
	if err != nil {
		return common.Address{}, err
	}

	return passportLogicAddress, nil
}

// DeployPassport deploys only Passport contract using existing PassportFactory contract
func (d *Deploy) DeployPassport(ctx context.Context, passportFactoryAddress common.Address) (common.Address, error) {
	backend := d.Backend
	ownerAuth := &d.TransactOpts

	///////////////////////////////////////////////////////
	// initializing PassportFactory
	///////////////////////////////////////////////////////

	d.Log("Initializing PassportFactory contract", "factory", passportFactoryAddress.Hex())
	passportFactoryContract, err := contracts.NewPassportFactoryContract(passportFactoryAddress, backend)
	if err != nil {
		return common.Address{}, fmt.Errorf("initializing PassportFactory contract: %v", err)
	}

	///////////////////////////////////////////////////////
	// deploying Passport
	///////////////////////////////////////////////////////

	d.Log("Deploying Passport contract", "owner_address", ownerAuth.From.Hex())
	tx, err := passportFactoryContract.CreatePassport(ownerAuth)
	if err != nil {
		return common.Address{}, fmt.Errorf("deploying Passport contract: %v", err)
	}
	txr, err := d.WaitForTxReceipt(ctx, tx.Hash())
	if err != nil {
		return common.Address{}, err
	}

	///////////////////////////////////////////////////////
	// filtering PassportCreated events
	///////////////////////////////////////////////////////

	lf, err := contracts.NewPassportFactoryLogFilterer()
	if err != nil {
		return common.Address{}, fmt.Errorf("initializing PassportFactoryLogFilterer: %v", err)
	}

	evs, err := lf.FilterPassportCreated(ctx, txr.Logs, nil, []common.Address{ownerAuth.From})
	if err != nil {
		return common.Address{}, fmt.Errorf("filtering PassportCreated events: %v", err)
	}
	if len(evs) != 1 {
		return common.Address{}, errors.New("expected exactly one PassportCreated event")
	}

	passportAddress := evs[0].Passport
	d.Log("Passport deployed", "contract_address", passportAddress.Hex())

	///////////////////////////////////////////////////////
	// initializing Passport
	///////////////////////////////////////////////////////

	d.Log("Initializing Passport contract", "passport", passportFactoryAddress.Hex())
	passportContract, err := contracts.NewPassportContract(passportAddress, backend)
	if err != nil {
		return common.Address{}, fmt.Errorf("initializing Passport contract: %v", err)
	}

	///////////////////////////////////////////////////////
	// Claiming ownership
	///////////////////////////////////////////////////////

	d.Log("Claiming ownership", "owner_address", ownerAuth.From)
	tx, err = passportContract.ClaimOwnership(ownerAuth)
	if err != nil {
		return common.Address{}, fmt.Errorf("claiming ownership: %v", err)
	}
	_, err = d.WaitForTxReceipt(ctx, tx.Hash())
	if err != nil {
		return common.Address{}, err
	}

	return passportAddress, nil
}
