package deployer

import (
	"context"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/monetha/go-verifiable-data/contracts"
	"github.com/monetha/go-verifiable-data/eth"
)

const (
	// PassportLogicDeployGasLimit is a gas limit to deploy only PassportLogic contract
	PassportLogicDeployGasLimit = 3420000
	// PassportLogicRegistryDeployGasLimit is a gas limit to deploy only PassportLogicRegistry contract
	PassportLogicRegistryDeployGasLimit = 1120000
	// PassportFactoryDeployGasLimit is a gas limit to deploy only PassportFactory contract
	PassportFactoryDeployGasLimit = 1160000
	// AddPassportLogicGasLimit is a gas limit to call AddPassportLogic of PassportLogicRegistry contract
	AddPassportLogicGasLimit = 50000
	// SetCurrentPassportLogicGasLimit is a gas limit to call SetCurrentPassportLogic of PassportLogicRegistry contract
	SetCurrentPassportLogicGasLimit = 50000

	// PassportFactoryGasLimit is a minimum gas amount needed to fully deployer passport factory contract with all dependent contracts
	PassportFactoryGasLimit = PassportLogicDeployGasLimit + PassportLogicRegistryDeployGasLimit + PassportFactoryDeployGasLimit

	// PassportGasLimit is a minimum gas amount needed to fully deploy passport contract
	PassportGasLimit = 480000
)

// Deploy contains methods to deployer contracts
type Deploy eth.Session

// New converts session to Deploy
func New(s *eth.Session) *Deploy {
	return (*Deploy)(s)
}

// BootstrapResult hold result of Bootstrap execution
type BootstrapResult struct {
	PassportLogicAddress         common.Address
	PassportLogicRegistryAddress common.Address
	PassportFactoryAddress       common.Address
}

// Bootstrap deploys PassportFactory contract and all contracts needed in order to deploy it
func (d *Deploy) Bootstrap(ctx context.Context) (*BootstrapResult, error) {
	///////////////////////////////////////////////////////
	// deploying PassportLogic
	///////////////////////////////////////////////////////
	passportLogicAddress, err := d.DeployPassportLogic(ctx)
	if err != nil {
		return nil, err
	}

	///////////////////////////////////////////////////////
	// deploying PassportLogicRegistry
	///////////////////////////////////////////////////////
	version := "0.1"
	passportLogicRegistryAddress, err := d.DeployPassportLogicRegistry(ctx, version, passportLogicAddress)
	if err != nil {
		return nil, err
	}

	///////////////////////////////////////////////////////
	// deploying PassportFactory
	///////////////////////////////////////////////////////
	passportFactoryAddress, err := d.DeployPassportFactory(ctx, passportLogicRegistryAddress)
	if err != nil {
		return nil, err
	}

	return &BootstrapResult{
		PassportLogicAddress:         passportLogicAddress,
		PassportLogicRegistryAddress: passportLogicRegistryAddress,
		PassportFactoryAddress:       passportFactoryAddress,
	}, nil
}

// DeployPassportLogic deploys only PassportLogic contract
func (d *Deploy) DeployPassportLogic(ctx context.Context) (common.Address, error) {
	backend := d.Backend
	ownerAuth := d.TransactOpts
	ownerAuth.Context = ctx

	d.Log("Deploying PassportLogic", "owner_address", ownerAuth.From)
	passportLogicAddress, tx, _, err := contracts.DeployPassportLogicContract(&ownerAuth, backend)
	if err != nil {
		return common.Address{}, fmt.Errorf("deploying PassportLogic contract: %v", err)
	}
	txr, err := d.WaitForTxReceipt(ctx, tx.Hash())
	if err != nil {
		return common.Address{}, err
	}

	d.Log("PassportLogic deployed", "contract_address", passportLogicAddress.Hex(), "gas_used", txr.GasUsed)

	return passportLogicAddress, nil
}

// DeployPassportLogicRegistry deploys only PassportLogicRegistry contract
func (d *Deploy) DeployPassportLogicRegistry(ctx context.Context, passportLogicVersion string, passportLogicAddress common.Address) (common.Address, error) {
	backend := d.Backend
	ownerAuth := d.TransactOpts
	ownerAuth.Context = ctx

	d.Log("Deploying PassportLogicRegistry", "owner_address", ownerAuth.From, "impl_version", passportLogicVersion, "impl_address", passportLogicAddress)
	passportLogicRegistryAddress, tx, _, err := contracts.DeployPassportLogicRegistryContract(&ownerAuth, backend, passportLogicVersion, passportLogicAddress)
	if err != nil {
		return common.Address{}, fmt.Errorf("deploying PassportLogicRegistry contract: %v", err)
	}
	txr, err := d.WaitForTxReceipt(ctx, tx.Hash())
	if err != nil {
		return common.Address{}, err
	}

	d.Log("PassportLogicRegistry deployed", "contract_address", passportLogicRegistryAddress.Hex(), "gas_used", txr.GasUsed)

	return passportLogicRegistryAddress, nil
}

// DeployPassportFactory deploys only PassportFactory contract
func (d *Deploy) DeployPassportFactory(ctx context.Context, passportLogicRegistryAddress common.Address) (common.Address, error) {
	backend := d.Backend
	ownerAuth := d.TransactOpts
	ownerAuth.Context = ctx

	d.Log("Deploying PassportFactory", "owner_address", ownerAuth.From, "registry", passportLogicRegistryAddress)
	passportFactoryAddress, tx, _, err := contracts.DeployPassportFactoryContract(&ownerAuth, backend, passportLogicRegistryAddress)
	if err != nil {
		return common.Address{}, fmt.Errorf("deploying PassportFactory contract: %v", err)
	}
	txr, err := d.WaitForTxReceipt(ctx, tx.Hash())
	if err != nil {
		return common.Address{}, err
	}

	d.Log("PassportFactory deployed", "contract_address", passportFactoryAddress.Hex(), "gas_used", txr.GasUsed)

	return passportFactoryAddress, nil
}

// UpgradePassportLogic deploys PassportLogic contract and sets it as current in PassportLogicRegistry
// Returns address of deployed PassportLogic contract
func (d *Deploy) UpgradePassportLogic(ctx context.Context, passportLogicVersion string, passportLogicRegistryAddress common.Address) (common.Address, error) {
	backend := d.Backend
	ownerAuth := d.TransactOpts
	ownerAuth.Context = ctx

	///////////////////////////////////////////////////////
	// deploying PassportLogic
	///////////////////////////////////////////////////////

	d.Log("Deploying PassportLogic", "owner_address", ownerAuth.From)
	passportLogicAddress, tx, _, err := contracts.DeployPassportLogicContract(&ownerAuth, backend)
	if err != nil {
		return common.Address{}, fmt.Errorf("deploying PassportLogic contract: %v", err)
	}
	txr, err := d.WaitForTxReceipt(ctx, tx.Hash())
	if err != nil {
		return common.Address{}, err
	}

	d.Log("PassportLogic deployed", "contract_address", passportLogicAddress.Hex(), "gas_used", txr.GasUsed)

	///////////////////////////////////////////////////////
	// updating PassportLogicRegistry
	///////////////////////////////////////////////////////

	d.Log("Initializing PassportLogicRegistry", "address", passportLogicRegistryAddress)
	passportLogicRegistryContract, err := contracts.NewPassportLogicRegistryContract(passportLogicRegistryAddress, backend)
	if err != nil {
		return common.Address{}, fmt.Errorf("initializing PassportLogicRegistry contract: %v", err)
	}

	d.Log("Adding new PassportLogic to registry", "registry_address", passportLogicRegistryAddress, "logic_address", passportLogicAddress, "version", passportLogicVersion)
	tx, err = passportLogicRegistryContract.AddPassportLogic(&ownerAuth, passportLogicVersion, passportLogicAddress)
	if err != nil {
		return common.Address{}, fmt.Errorf("adding new PassportLogic to registry: %v", err)
	}
	txr, err = d.WaitForTxReceipt(ctx, tx.Hash())
	if err != nil {
		return common.Address{}, err
	}
	d.Log("New logic added to registry", "gas_used", txr.GasUsed)

	d.Log("Setting current PassportLogic in registry", "registry_address", passportLogicRegistryAddress, "version", passportLogicVersion)
	tx, err = passportLogicRegistryContract.SetCurrentPassportLogic(&ownerAuth, passportLogicVersion)
	if err != nil {
		return common.Address{}, fmt.Errorf("setting current PassportLogic in registry: %v", err)
	}
	txr, err = d.WaitForTxReceipt(ctx, tx.Hash())
	if err != nil {
		return common.Address{}, err
	}
	d.Log("New logic set as current in registry", "gas_used", txr.GasUsed)

	return passportLogicAddress, nil
}

// DeployPassport deploys only Passport contract using existing PassportFactory contract
func (d *Deploy) DeployPassport(ctx context.Context, passportFactoryAddress common.Address) (common.Address, error) {
	backend := d.Backend
	ownerAuth := d.TransactOpts
	ownerAuth.Context = ctx

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
	tx, err := passportFactoryContract.CreatePassport(&ownerAuth)
	if err != nil {
		return common.Address{}, fmt.Errorf("deploying Passport contract: %v", err)
	}
	txr, err := d.WaitForTxReceipt(ctx, tx.Hash())
	if err != nil {
		return common.Address{}, err
	}
	d.Log("New passport deployed", "gas_used", txr.GasUsed)

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

	d.Log("Initializing Passport contract", "passport", passportAddress.Hex())
	passportContract, err := contracts.NewPassportContract(passportAddress, backend)
	if err != nil {
		return common.Address{}, fmt.Errorf("initializing Passport contract: %v", err)
	}

	///////////////////////////////////////////////////////
	// Claiming ownership
	///////////////////////////////////////////////////////

	d.Log("Claiming ownership", "owner_address", ownerAuth.From)
	tx, err = passportContract.ClaimOwnership(&ownerAuth)
	if err != nil {
		return common.Address{}, fmt.Errorf("claiming ownership: %v", err)
	}
	txr, err = d.WaitForTxReceipt(ctx, tx.Hash())
	if err != nil {
		return common.Address{}, err
	}
	d.Log("Ownership claimed successfully", "gas_used", txr.GasUsed)

	return passportAddress, nil
}

// DeployFactProviderRegistry deploys only FactProviderRegistry contract
func (d *Deploy) DeployFactProviderRegistry(ctx context.Context) (common.Address, error) {
	backend := d.Backend
	ownerAuth := d.TransactOpts
	ownerAuth.Context = ctx

	d.Log("Deploying FactProviderRegistry", "owner_address", ownerAuth.From)
	factProviderRegistryAddress, tx, _, err := contracts.DeployFactProviderRegistryContract(&ownerAuth, backend)
	if err != nil {
		return common.Address{}, fmt.Errorf("deploying FactProviderRegistry contract: %v", err)
	}
	txr, err := d.WaitForTxReceipt(ctx, tx.Hash())
	if err != nil {
		return common.Address{}, err
	}

	d.Log("FactProviderRegistry deployed", "contract_address", factProviderRegistryAddress.Hex(), "gas_used", txr.GasUsed)

	return factProviderRegistryAddress, nil
}
