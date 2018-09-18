// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// IPassportLogicRegistryContractABI is the input ABI used to generate the binding from.
const IPassportLogicRegistryContractABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"version\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"PassportLogicAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"version\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"CurrentPassportLogicSet\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"name\":\"_version\",\"type\":\"string\"}],\"name\":\"getPassportLogic\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurrentPassportLogicVersion\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurrentPassportLogic\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IPassportLogicRegistryContract is an auto generated Go binding around an Ethereum contract.
type IPassportLogicRegistryContract struct {
	IPassportLogicRegistryContractCaller     // Read-only binding to the contract
	IPassportLogicRegistryContractTransactor // Write-only binding to the contract
	IPassportLogicRegistryContractFilterer   // Log filterer for contract events
}

// IPassportLogicRegistryContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type IPassportLogicRegistryContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPassportLogicRegistryContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IPassportLogicRegistryContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPassportLogicRegistryContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPassportLogicRegistryContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPassportLogicRegistryContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPassportLogicRegistryContractSession struct {
	Contract     *IPassportLogicRegistryContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                   // Call options to use throughout this session
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// IPassportLogicRegistryContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPassportLogicRegistryContractCallerSession struct {
	Contract *IPassportLogicRegistryContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                         // Call options to use throughout this session
}

// IPassportLogicRegistryContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPassportLogicRegistryContractTransactorSession struct {
	Contract     *IPassportLogicRegistryContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                         // Transaction auth options to use throughout this session
}

// IPassportLogicRegistryContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type IPassportLogicRegistryContractRaw struct {
	Contract *IPassportLogicRegistryContract // Generic contract binding to access the raw methods on
}

// IPassportLogicRegistryContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPassportLogicRegistryContractCallerRaw struct {
	Contract *IPassportLogicRegistryContractCaller // Generic read-only contract binding to access the raw methods on
}

// IPassportLogicRegistryContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPassportLogicRegistryContractTransactorRaw struct {
	Contract *IPassportLogicRegistryContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIPassportLogicRegistryContract creates a new instance of IPassportLogicRegistryContract, bound to a specific deployed contract.
func NewIPassportLogicRegistryContract(address common.Address, backend bind.ContractBackend) (*IPassportLogicRegistryContract, error) {
	contract, err := bindIPassportLogicRegistryContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPassportLogicRegistryContract{IPassportLogicRegistryContractCaller: IPassportLogicRegistryContractCaller{contract: contract}, IPassportLogicRegistryContractTransactor: IPassportLogicRegistryContractTransactor{contract: contract}, IPassportLogicRegistryContractFilterer: IPassportLogicRegistryContractFilterer{contract: contract}}, nil
}

// NewIPassportLogicRegistryContractCaller creates a new read-only instance of IPassportLogicRegistryContract, bound to a specific deployed contract.
func NewIPassportLogicRegistryContractCaller(address common.Address, caller bind.ContractCaller) (*IPassportLogicRegistryContractCaller, error) {
	contract, err := bindIPassportLogicRegistryContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPassportLogicRegistryContractCaller{contract: contract}, nil
}

// NewIPassportLogicRegistryContractTransactor creates a new write-only instance of IPassportLogicRegistryContract, bound to a specific deployed contract.
func NewIPassportLogicRegistryContractTransactor(address common.Address, transactor bind.ContractTransactor) (*IPassportLogicRegistryContractTransactor, error) {
	contract, err := bindIPassportLogicRegistryContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPassportLogicRegistryContractTransactor{contract: contract}, nil
}

// NewIPassportLogicRegistryContractFilterer creates a new log filterer instance of IPassportLogicRegistryContract, bound to a specific deployed contract.
func NewIPassportLogicRegistryContractFilterer(address common.Address, filterer bind.ContractFilterer) (*IPassportLogicRegistryContractFilterer, error) {
	contract, err := bindIPassportLogicRegistryContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPassportLogicRegistryContractFilterer{contract: contract}, nil
}

// bindIPassportLogicRegistryContract binds a generic wrapper to an already deployed contract.
func bindIPassportLogicRegistryContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IPassportLogicRegistryContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPassportLogicRegistryContract *IPassportLogicRegistryContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IPassportLogicRegistryContract.Contract.IPassportLogicRegistryContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPassportLogicRegistryContract *IPassportLogicRegistryContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPassportLogicRegistryContract.Contract.IPassportLogicRegistryContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPassportLogicRegistryContract *IPassportLogicRegistryContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPassportLogicRegistryContract.Contract.IPassportLogicRegistryContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPassportLogicRegistryContract *IPassportLogicRegistryContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IPassportLogicRegistryContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPassportLogicRegistryContract *IPassportLogicRegistryContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPassportLogicRegistryContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPassportLogicRegistryContract *IPassportLogicRegistryContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPassportLogicRegistryContract.Contract.contract.Transact(opts, method, params...)
}

// GetCurrentPassportLogic is a free data retrieval call binding the contract method 0x609725ef.
//
// Solidity: function getCurrentPassportLogic() constant returns(address)
func (_IPassportLogicRegistryContract *IPassportLogicRegistryContractCaller) GetCurrentPassportLogic(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IPassportLogicRegistryContract.contract.Call(opts, out, "getCurrentPassportLogic")
	return *ret0, err
}

// GetCurrentPassportLogic is a free data retrieval call binding the contract method 0x609725ef.
//
// Solidity: function getCurrentPassportLogic() constant returns(address)
func (_IPassportLogicRegistryContract *IPassportLogicRegistryContractSession) GetCurrentPassportLogic() (common.Address, error) {
	return _IPassportLogicRegistryContract.Contract.GetCurrentPassportLogic(&_IPassportLogicRegistryContract.CallOpts)
}

// GetCurrentPassportLogic is a free data retrieval call binding the contract method 0x609725ef.
//
// Solidity: function getCurrentPassportLogic() constant returns(address)
func (_IPassportLogicRegistryContract *IPassportLogicRegistryContractCallerSession) GetCurrentPassportLogic() (common.Address, error) {
	return _IPassportLogicRegistryContract.Contract.GetCurrentPassportLogic(&_IPassportLogicRegistryContract.CallOpts)
}

// GetCurrentPassportLogicVersion is a free data retrieval call binding the contract method 0xba612493.
//
// Solidity: function getCurrentPassportLogicVersion() constant returns(string)
func (_IPassportLogicRegistryContract *IPassportLogicRegistryContractCaller) GetCurrentPassportLogicVersion(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _IPassportLogicRegistryContract.contract.Call(opts, out, "getCurrentPassportLogicVersion")
	return *ret0, err
}

// GetCurrentPassportLogicVersion is a free data retrieval call binding the contract method 0xba612493.
//
// Solidity: function getCurrentPassportLogicVersion() constant returns(string)
func (_IPassportLogicRegistryContract *IPassportLogicRegistryContractSession) GetCurrentPassportLogicVersion() (string, error) {
	return _IPassportLogicRegistryContract.Contract.GetCurrentPassportLogicVersion(&_IPassportLogicRegistryContract.CallOpts)
}

// GetCurrentPassportLogicVersion is a free data retrieval call binding the contract method 0xba612493.
//
// Solidity: function getCurrentPassportLogicVersion() constant returns(string)
func (_IPassportLogicRegistryContract *IPassportLogicRegistryContractCallerSession) GetCurrentPassportLogicVersion() (string, error) {
	return _IPassportLogicRegistryContract.Contract.GetCurrentPassportLogicVersion(&_IPassportLogicRegistryContract.CallOpts)
}

// GetPassportLogic is a free data retrieval call binding the contract method 0x9a1295d9.
//
// Solidity: function getPassportLogic(_version string) constant returns(address)
func (_IPassportLogicRegistryContract *IPassportLogicRegistryContractCaller) GetPassportLogic(opts *bind.CallOpts, _version string) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IPassportLogicRegistryContract.contract.Call(opts, out, "getPassportLogic", _version)
	return *ret0, err
}

// GetPassportLogic is a free data retrieval call binding the contract method 0x9a1295d9.
//
// Solidity: function getPassportLogic(_version string) constant returns(address)
func (_IPassportLogicRegistryContract *IPassportLogicRegistryContractSession) GetPassportLogic(_version string) (common.Address, error) {
	return _IPassportLogicRegistryContract.Contract.GetPassportLogic(&_IPassportLogicRegistryContract.CallOpts, _version)
}

// GetPassportLogic is a free data retrieval call binding the contract method 0x9a1295d9.
//
// Solidity: function getPassportLogic(_version string) constant returns(address)
func (_IPassportLogicRegistryContract *IPassportLogicRegistryContractCallerSession) GetPassportLogic(_version string) (common.Address, error) {
	return _IPassportLogicRegistryContract.Contract.GetPassportLogic(&_IPassportLogicRegistryContract.CallOpts, _version)
}

// IPassportLogicRegistryContractCurrentPassportLogicSetIterator is returned from FilterCurrentPassportLogicSet and is used to iterate over the raw logs and unpacked data for CurrentPassportLogicSet events raised by the IPassportLogicRegistryContract contract.
type IPassportLogicRegistryContractCurrentPassportLogicSetIterator struct {
	Event *IPassportLogicRegistryContractCurrentPassportLogicSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IPassportLogicRegistryContractCurrentPassportLogicSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPassportLogicRegistryContractCurrentPassportLogicSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IPassportLogicRegistryContractCurrentPassportLogicSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IPassportLogicRegistryContractCurrentPassportLogicSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPassportLogicRegistryContractCurrentPassportLogicSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPassportLogicRegistryContractCurrentPassportLogicSet represents a CurrentPassportLogicSet event raised by the IPassportLogicRegistryContract contract.
type IPassportLogicRegistryContractCurrentPassportLogicSet struct {
	Version        string
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterCurrentPassportLogicSet is a free log retrieval operation binding the contract event 0x4e366bf178b123bb29442cddeceedf743c23cbb40cffa5f577217fe2c54a0b19.
//
// Solidity: e CurrentPassportLogicSet(version string, implementation address)
func (_IPassportLogicRegistryContract *IPassportLogicRegistryContractFilterer) FilterCurrentPassportLogicSet(opts *bind.FilterOpts) (*IPassportLogicRegistryContractCurrentPassportLogicSetIterator, error) {

	logs, sub, err := _IPassportLogicRegistryContract.contract.FilterLogs(opts, "CurrentPassportLogicSet")
	if err != nil {
		return nil, err
	}
	return &IPassportLogicRegistryContractCurrentPassportLogicSetIterator{contract: _IPassportLogicRegistryContract.contract, event: "CurrentPassportLogicSet", logs: logs, sub: sub}, nil
}

// WatchCurrentPassportLogicSet is a free log subscription operation binding the contract event 0x4e366bf178b123bb29442cddeceedf743c23cbb40cffa5f577217fe2c54a0b19.
//
// Solidity: e CurrentPassportLogicSet(version string, implementation address)
func (_IPassportLogicRegistryContract *IPassportLogicRegistryContractFilterer) WatchCurrentPassportLogicSet(opts *bind.WatchOpts, sink chan<- *IPassportLogicRegistryContractCurrentPassportLogicSet) (event.Subscription, error) {

	logs, sub, err := _IPassportLogicRegistryContract.contract.WatchLogs(opts, "CurrentPassportLogicSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPassportLogicRegistryContractCurrentPassportLogicSet)
				if err := _IPassportLogicRegistryContract.contract.UnpackLog(event, "CurrentPassportLogicSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// IPassportLogicRegistryContractPassportLogicAddedIterator is returned from FilterPassportLogicAdded and is used to iterate over the raw logs and unpacked data for PassportLogicAdded events raised by the IPassportLogicRegistryContract contract.
type IPassportLogicRegistryContractPassportLogicAddedIterator struct {
	Event *IPassportLogicRegistryContractPassportLogicAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IPassportLogicRegistryContractPassportLogicAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPassportLogicRegistryContractPassportLogicAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IPassportLogicRegistryContractPassportLogicAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IPassportLogicRegistryContractPassportLogicAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPassportLogicRegistryContractPassportLogicAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPassportLogicRegistryContractPassportLogicAdded represents a PassportLogicAdded event raised by the IPassportLogicRegistryContract contract.
type IPassportLogicRegistryContractPassportLogicAdded struct {
	Version        string
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterPassportLogicAdded is a free log retrieval operation binding the contract event 0x7471eb04045ae72adb2fb73deb1e873113901110dd66dbde715232f2495a0cd8.
//
// Solidity: e PassportLogicAdded(version string, implementation address)
func (_IPassportLogicRegistryContract *IPassportLogicRegistryContractFilterer) FilterPassportLogicAdded(opts *bind.FilterOpts) (*IPassportLogicRegistryContractPassportLogicAddedIterator, error) {

	logs, sub, err := _IPassportLogicRegistryContract.contract.FilterLogs(opts, "PassportLogicAdded")
	if err != nil {
		return nil, err
	}
	return &IPassportLogicRegistryContractPassportLogicAddedIterator{contract: _IPassportLogicRegistryContract.contract, event: "PassportLogicAdded", logs: logs, sub: sub}, nil
}

// WatchPassportLogicAdded is a free log subscription operation binding the contract event 0x7471eb04045ae72adb2fb73deb1e873113901110dd66dbde715232f2495a0cd8.
//
// Solidity: e PassportLogicAdded(version string, implementation address)
func (_IPassportLogicRegistryContract *IPassportLogicRegistryContractFilterer) WatchPassportLogicAdded(opts *bind.WatchOpts, sink chan<- *IPassportLogicRegistryContractPassportLogicAdded) (event.Subscription, error) {

	logs, sub, err := _IPassportLogicRegistryContract.contract.WatchLogs(opts, "PassportLogicAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPassportLogicRegistryContractPassportLogicAdded)
				if err := _IPassportLogicRegistryContract.contract.UnpackLog(event, "PassportLogicAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}
