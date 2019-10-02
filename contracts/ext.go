package contracts

import (
	"context"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	methereum "github.com/monetha/go-ethereum"
)

var (
	// PassportLogicABI is prepared(parsed) ABI specification of PassportLogic contract
	PassportLogicABI abi.ABI
	// FactProviderRegistryABI is prepared(parsed) ABI specification of FactProviderRegistry contract
	FactProviderRegistryABI abi.ABI
)

func init() {
	var err error
	PassportLogicABI, err = abi.JSON(strings.NewReader(PassportLogicContractABI))
	if err != nil {
		panic("contracts: initializing PassportLogicContractABI: " + err.Error())
	}
	FactProviderRegistryABI, err = abi.JSON(strings.NewReader(FactProviderRegistryContractABI))
	if err != nil {
		panic("contracts: initializing FactProviderRegistryContractABI: " + err.Error())
	}
}

// InitPassportLogicContract creates a new instance of PassportLogicContract, bound to a specific deployed contract.
// This method is faster than NewPassportLogicContract and doesn't return an error, because it reuses parsed PassportLogicContractABI.
func InitPassportLogicContract(address common.Address, backend bind.ContractBackend) *PassportLogicContract {
	contract := bind.NewBoundContract(address, PassportLogicABI, backend, backend, backend)
	return &PassportLogicContract{
		PassportLogicContractCaller:     PassportLogicContractCaller{contract: contract},
		PassportLogicContractTransactor: PassportLogicContractTransactor{contract: contract},
		PassportLogicContractFilterer:   PassportLogicContractFilterer{contract: contract}}
}

// InitFactProviderRegistryContract creates a new instance of FactProviderRegistryContract, bound to a specific deployed contract.
// This method is faster than NewFactProviderRegistryContract and doesn't return an error, because it reuses parsed FactProviderRegistryABI.
func InitFactProviderRegistryContract(address common.Address, backend bind.ContractBackend) *FactProviderRegistryContract {
	contract := bind.NewBoundContract(address, FactProviderRegistryABI, backend, backend, backend)
	return &FactProviderRegistryContract{
		FactProviderRegistryContractCaller:     FactProviderRegistryContractCaller{contract: contract},
		FactProviderRegistryContractTransactor: FactProviderRegistryContractTransactor{contract: contract},
		FactProviderRegistryContractFilterer:   FactProviderRegistryContractFilterer{contract: contract}}
}

// PassportFactoryLogFilterer filters PassportFactory event logs
type PassportFactoryLogFilterer struct {
	abi abi.ABI
}

// NewPassportFactoryLogFilterer creates an instance of PassportFactoryLogFilterer
func NewPassportFactoryLogFilterer() (*PassportFactoryLogFilterer, error) {
	parsed, err := abi.JSON(strings.NewReader(PassportFactoryContractABI))
	if err != nil {
		return nil, err
	}
	return &PassportFactoryLogFilterer{parsed}, nil
}

// FilterPassportCreated parses event logs and returns PassportCreated events if any found.
func (f *PassportFactoryLogFilterer) FilterPassportCreated(ctx context.Context, logs []*types.Log, passport []common.Address, owner []common.Address) (events []PassportFactoryContractPassportCreated, err error) {
	cf := &PassportFactoryContractFilterer{
		contract: bind.NewBoundContract(common.Address{}, f.abi, nil, nil, methereum.SliceLogFilterer(logs)),
	}

	var it *PassportFactoryContractPassportCreatedIterator
	it, err = cf.FilterPassportCreated(&bind.FilterOpts{Context: ctx}, passport, owner)
	if err != nil {
		return
	}
	defer func() {
		if cErr := it.Close(); err == nil && cErr != nil {
			err = cErr
		}
	}()

	for it.Next() {
		if err = it.Error(); err != nil {
			return nil, err
		}

		ev := it.Event
		if ev == nil {
			continue
		}

		events = append(events, *ev)
	}

	return
}

// FilterPrivateDataExchangeProposed filters event
// PrivateDataExchangeProposed(uint256 indexed exchangeIdx, address indexed dataRequester, address indexed passportOwner)
// from logs.
func FilterPrivateDataExchangeProposed(
	ctx context.Context,
	logs []*types.Log,
	exchangeIdx []*big.Int,
	dataRequester []common.Address,
	passportOwner []common.Address,
) (events []PassportLogicContractPrivateDataExchangeProposed, err error) {
	cf := &PassportLogicContractFilterer{
		contract: bind.NewBoundContract(common.Address{}, PassportLogicABI, nil, nil, methereum.SliceLogFilterer(logs)),
	}

	it, err := cf.FilterPrivateDataExchangeProposed(&bind.FilterOpts{Context: ctx}, exchangeIdx, dataRequester, passportOwner)
	if err != nil {
		return
	}
	defer func() {
		if cErr := it.Close(); err == nil && cErr != nil {
			err = cErr
		}
	}()

	for it.Next() {
		if err = it.Error(); err != nil {
			return nil, err
		}

		ev := it.Event
		if ev == nil {
			continue
		}

		events = append(events, *ev)
	}

	return
}

// FilterPrivateDataExchangeDisputed filters event
// PrivateDataExchangeDisputed(uint256 indexed exchangeIdx, bool indexed successful, address indexed cheater)
// from logs.
func FilterPrivateDataExchangeDisputed(
	ctx context.Context,
	logs []*types.Log,
	exchangeIdx []*big.Int, successful []bool, cheater []common.Address,
) (events []PassportLogicContractPrivateDataExchangeDisputed, err error) {
	cf := &PassportLogicContractFilterer{
		contract: bind.NewBoundContract(common.Address{}, PassportLogicABI, nil, nil, methereum.SliceLogFilterer(logs)),
	}

	it, err := cf.FilterPrivateDataExchangeDisputed(&bind.FilterOpts{Context: ctx}, exchangeIdx, successful, cheater)
	if err != nil {
		return
	}
	defer func() {
		if cErr := it.Close(); err == nil && cErr != nil {
			err = cErr
		}
	}()

	for it.Next() {
		if err = it.Error(); err != nil {
			return nil, err
		}

		ev := it.Event
		if ev == nil {
			continue
		}

		events = append(events, *ev)
	}

	return
}
