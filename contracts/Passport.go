// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// PassportContractABI is the input ABI used to generate the binding from.
const PassportContractABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"claimOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"destroy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"destroyAndSend\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_registry\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousRegistry\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newRegistry\",\"type\":\"address\"}],\"name\":\"PassportLogicRegistryChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_registry\",\"type\":\"address\"}],\"name\":\"changePassportLogicRegistry\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getPassportLogicRegistry\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// PassportContractBin is the compiled bytecode used for deploying new contracts.
const PassportContractBin = `0x608060405234801561001057600080fd5b5060405160208061086883398101604081815291517f6f72672e6d6f6e657468612e70726f78792e6f776e6572000000000000000000825291519081900360170190206000805160206108288339815191521461006957fe5b61007b3364010000000061015b810204565b604080517f6f72672e6d6f6e657468612e70726f78792e70656e64696e674f776e657200008152905190819003601e0190207fcfd0c6ea5352192d7d4c5d4e7a73c5da12c871730cb60ff57879cbe7b403bb52146100d557fe5b604080517f6f72672e6d6f6e657468612e70617373706f72742e70726f78792e726567697381527f7472790000000000000000000000000000000000000000000000000000000000602082015290519081900360230190206000805160206108488339815191521461014357fe5b6101558164010000000061016d810204565b5061021f565b60008051602061082883398151915255565b6000600160a060020a038216151561020c57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f43616e6e6f742073657420726567697374727920746f2061207a65726f20616460448201527f6472657373000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b5060008051602061084883398151915255565b6105fa8061022e6000396000f3006080604052600436106100985763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416634645a41881146100aa5780634e71e0c8146100cb578063715018a6146100e057806383197ef0146100f557806386d5c5f91461010a5780638da5cb5b1461013b578063e30c397814610150578063f2fde38b14610165578063f5074f4114610186575b6100a86100a36101a7565b610238565b005b3480156100b657600080fd5b506100a8600160a060020a036004351661025c565b3480156100d757600080fd5b506100a86102cb565b3480156100ec57600080fd5b506100a8610351565b34801561010157600080fd5b506100a86103b4565b34801561011657600080fd5b5061011f6103e3565b60408051600160a060020a039092168252519081900360200190f35b34801561014757600080fd5b5061011f6103f2565b34801561015c57600080fd5b5061011f6103fc565b34801561017157600080fd5b506100a8600160a060020a0360043516610406565b34801561019257600080fd5b506100a8600160a060020a036004351661042b565b60006101b1610453565b600160a060020a031663609725ef6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b15801561020757600080fd5b505af115801561021b573d6000803e3d6000fd5b505050506040513d602081101561023157600080fd5b5051905090565b3660008037600080366000845af43d6000803e808015610257573d6000f35b3d6000fd5b610264610478565b600160a060020a0316331461027857600080fd5b80600160a060020a031661028a610453565b600160a060020a03167f5c2abfd67230c0e47d6de28402bfe206c7a57283cba891416ed657fd70a714c260405160405180910390a36102c88161049d565b50565b6102d3610561565b600160a060020a031633146102e757600080fd5b6102ef610561565b600160a060020a0316610300610478565b600160a060020a03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3610345610340610561565b610586565b61034f60006105aa565b565b610359610478565b600160a060020a0316331461036d57600080fd5b610375610478565b600160a060020a03167ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482060405160405180910390a261034f6000610586565b6103bc610478565b600160a060020a031633146103d057600080fd5b6103d8610478565b600160a060020a0316ff5b60006103ed610453565b905090565b60006103ed610478565b60006103ed610561565b61040e610478565b600160a060020a0316331461042257600080fd5b6102c8816105aa565b610433610478565b600160a060020a0316331461044757600080fd5b80600160a060020a0316ff5b7fa04bab69e45aeb4c94a78ba5bc1be67ef28977c4fdf815a30b829a794eb67a4a5490565b7f3ca57e4b51fc2e18497b219410298879868edada7e6fe5132c8feceb0a080d225490565b6000600160a060020a038216151561053c57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f43616e6e6f742073657420726567697374727920746f2061207a65726f20616460448201527f6472657373000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b507fa04bab69e45aeb4c94a78ba5bc1be67ef28977c4fdf815a30b829a794eb67a4a55565b7fcfd0c6ea5352192d7d4c5d4e7a73c5da12c871730cb60ff57879cbe7b403bb525490565b7f3ca57e4b51fc2e18497b219410298879868edada7e6fe5132c8feceb0a080d2255565b7fcfd0c6ea5352192d7d4c5d4e7a73c5da12c871730cb60ff57879cbe7b403bb52555600a165627a7a7230582027bb57ac09535e72c3e85d0be1e55abee8f4ad72ffc717aabb25edc02408067400293ca57e4b51fc2e18497b219410298879868edada7e6fe5132c8feceb0a080d22a04bab69e45aeb4c94a78ba5bc1be67ef28977c4fdf815a30b829a794eb67a4a`

// DeployPassportContract deploys a new Ethereum contract, binding an instance of PassportContract to it.
func DeployPassportContract(auth *bind.TransactOpts, backend bind.ContractBackend, _registry common.Address) (common.Address, *types.Transaction, *PassportContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PassportContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PassportContractBin), backend, _registry)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PassportContract{PassportContractCaller: PassportContractCaller{contract: contract}, PassportContractTransactor: PassportContractTransactor{contract: contract}, PassportContractFilterer: PassportContractFilterer{contract: contract}}, nil
}

// PassportContract is an auto generated Go binding around an Ethereum contract.
type PassportContract struct {
	PassportContractCaller     // Read-only binding to the contract
	PassportContractTransactor // Write-only binding to the contract
	PassportContractFilterer   // Log filterer for contract events
}

// PassportContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type PassportContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PassportContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PassportContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PassportContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PassportContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PassportContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PassportContractSession struct {
	Contract     *PassportContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PassportContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PassportContractCallerSession struct {
	Contract *PassportContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// PassportContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PassportContractTransactorSession struct {
	Contract     *PassportContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// PassportContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type PassportContractRaw struct {
	Contract *PassportContract // Generic contract binding to access the raw methods on
}

// PassportContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PassportContractCallerRaw struct {
	Contract *PassportContractCaller // Generic read-only contract binding to access the raw methods on
}

// PassportContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PassportContractTransactorRaw struct {
	Contract *PassportContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPassportContract creates a new instance of PassportContract, bound to a specific deployed contract.
func NewPassportContract(address common.Address, backend bind.ContractBackend) (*PassportContract, error) {
	contract, err := bindPassportContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PassportContract{PassportContractCaller: PassportContractCaller{contract: contract}, PassportContractTransactor: PassportContractTransactor{contract: contract}, PassportContractFilterer: PassportContractFilterer{contract: contract}}, nil
}

// NewPassportContractCaller creates a new read-only instance of PassportContract, bound to a specific deployed contract.
func NewPassportContractCaller(address common.Address, caller bind.ContractCaller) (*PassportContractCaller, error) {
	contract, err := bindPassportContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PassportContractCaller{contract: contract}, nil
}

// NewPassportContractTransactor creates a new write-only instance of PassportContract, bound to a specific deployed contract.
func NewPassportContractTransactor(address common.Address, transactor bind.ContractTransactor) (*PassportContractTransactor, error) {
	contract, err := bindPassportContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PassportContractTransactor{contract: contract}, nil
}

// NewPassportContractFilterer creates a new log filterer instance of PassportContract, bound to a specific deployed contract.
func NewPassportContractFilterer(address common.Address, filterer bind.ContractFilterer) (*PassportContractFilterer, error) {
	contract, err := bindPassportContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PassportContractFilterer{contract: contract}, nil
}

// bindPassportContract binds a generic wrapper to an already deployed contract.
func bindPassportContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PassportContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PassportContract *PassportContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PassportContract.Contract.PassportContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PassportContract *PassportContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PassportContract.Contract.PassportContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PassportContract *PassportContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PassportContract.Contract.PassportContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PassportContract *PassportContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PassportContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PassportContract *PassportContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PassportContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PassportContract *PassportContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PassportContract.Contract.contract.Transact(opts, method, params...)
}

// GetPassportLogicRegistry is a free data retrieval call binding the contract method 0x86d5c5f9.
//
// Solidity: function getPassportLogicRegistry() constant returns(address)
func (_PassportContract *PassportContractCaller) GetPassportLogicRegistry(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PassportContract.contract.Call(opts, out, "getPassportLogicRegistry")
	return *ret0, err
}

// GetPassportLogicRegistry is a free data retrieval call binding the contract method 0x86d5c5f9.
//
// Solidity: function getPassportLogicRegistry() constant returns(address)
func (_PassportContract *PassportContractSession) GetPassportLogicRegistry() (common.Address, error) {
	return _PassportContract.Contract.GetPassportLogicRegistry(&_PassportContract.CallOpts)
}

// GetPassportLogicRegistry is a free data retrieval call binding the contract method 0x86d5c5f9.
//
// Solidity: function getPassportLogicRegistry() constant returns(address)
func (_PassportContract *PassportContractCallerSession) GetPassportLogicRegistry() (common.Address, error) {
	return _PassportContract.Contract.GetPassportLogicRegistry(&_PassportContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PassportContract *PassportContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PassportContract.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PassportContract *PassportContractSession) Owner() (common.Address, error) {
	return _PassportContract.Contract.Owner(&_PassportContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PassportContract *PassportContractCallerSession) Owner() (common.Address, error) {
	return _PassportContract.Contract.Owner(&_PassportContract.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() constant returns(address)
func (_PassportContract *PassportContractCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PassportContract.contract.Call(opts, out, "pendingOwner")
	return *ret0, err
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() constant returns(address)
func (_PassportContract *PassportContractSession) PendingOwner() (common.Address, error) {
	return _PassportContract.Contract.PendingOwner(&_PassportContract.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() constant returns(address)
func (_PassportContract *PassportContractCallerSession) PendingOwner() (common.Address, error) {
	return _PassportContract.Contract.PendingOwner(&_PassportContract.CallOpts)
}

// ChangePassportLogicRegistry is a paid mutator transaction binding the contract method 0x4645a418.
//
// Solidity: function changePassportLogicRegistry(_registry address) returns()
func (_PassportContract *PassportContractTransactor) ChangePassportLogicRegistry(opts *bind.TransactOpts, _registry common.Address) (*types.Transaction, error) {
	return _PassportContract.contract.Transact(opts, "changePassportLogicRegistry", _registry)
}

// ChangePassportLogicRegistry is a paid mutator transaction binding the contract method 0x4645a418.
//
// Solidity: function changePassportLogicRegistry(_registry address) returns()
func (_PassportContract *PassportContractSession) ChangePassportLogicRegistry(_registry common.Address) (*types.Transaction, error) {
	return _PassportContract.Contract.ChangePassportLogicRegistry(&_PassportContract.TransactOpts, _registry)
}

// ChangePassportLogicRegistry is a paid mutator transaction binding the contract method 0x4645a418.
//
// Solidity: function changePassportLogicRegistry(_registry address) returns()
func (_PassportContract *PassportContractTransactorSession) ChangePassportLogicRegistry(_registry common.Address) (*types.Transaction, error) {
	return _PassportContract.Contract.ChangePassportLogicRegistry(&_PassportContract.TransactOpts, _registry)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_PassportContract *PassportContractTransactor) ClaimOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PassportContract.contract.Transact(opts, "claimOwnership")
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_PassportContract *PassportContractSession) ClaimOwnership() (*types.Transaction, error) {
	return _PassportContract.Contract.ClaimOwnership(&_PassportContract.TransactOpts)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_PassportContract *PassportContractTransactorSession) ClaimOwnership() (*types.Transaction, error) {
	return _PassportContract.Contract.ClaimOwnership(&_PassportContract.TransactOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_PassportContract *PassportContractTransactor) Destroy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PassportContract.contract.Transact(opts, "destroy")
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_PassportContract *PassportContractSession) Destroy() (*types.Transaction, error) {
	return _PassportContract.Contract.Destroy(&_PassportContract.TransactOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_PassportContract *PassportContractTransactorSession) Destroy() (*types.Transaction, error) {
	return _PassportContract.Contract.Destroy(&_PassportContract.TransactOpts)
}

// DestroyAndSend is a paid mutator transaction binding the contract method 0xf5074f41.
//
// Solidity: function destroyAndSend(_recipient address) returns()
func (_PassportContract *PassportContractTransactor) DestroyAndSend(opts *bind.TransactOpts, _recipient common.Address) (*types.Transaction, error) {
	return _PassportContract.contract.Transact(opts, "destroyAndSend", _recipient)
}

// DestroyAndSend is a paid mutator transaction binding the contract method 0xf5074f41.
//
// Solidity: function destroyAndSend(_recipient address) returns()
func (_PassportContract *PassportContractSession) DestroyAndSend(_recipient common.Address) (*types.Transaction, error) {
	return _PassportContract.Contract.DestroyAndSend(&_PassportContract.TransactOpts, _recipient)
}

// DestroyAndSend is a paid mutator transaction binding the contract method 0xf5074f41.
//
// Solidity: function destroyAndSend(_recipient address) returns()
func (_PassportContract *PassportContractTransactorSession) DestroyAndSend(_recipient common.Address) (*types.Transaction, error) {
	return _PassportContract.Contract.DestroyAndSend(&_PassportContract.TransactOpts, _recipient)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PassportContract *PassportContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PassportContract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PassportContract *PassportContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _PassportContract.Contract.RenounceOwnership(&_PassportContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PassportContract *PassportContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PassportContract.Contract.RenounceOwnership(&_PassportContract.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_PassportContract *PassportContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PassportContract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_PassportContract *PassportContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PassportContract.Contract.TransferOwnership(&_PassportContract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_PassportContract *PassportContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PassportContract.Contract.TransferOwnership(&_PassportContract.TransactOpts, newOwner)
}

// PassportContractOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the PassportContract contract.
type PassportContractOwnershipRenouncedIterator struct {
	Event *PassportContractOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *PassportContractOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PassportContractOwnershipRenounced)
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
		it.Event = new(PassportContractOwnershipRenounced)
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
func (it *PassportContractOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PassportContractOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PassportContractOwnershipRenounced represents a OwnershipRenounced event raised by the PassportContract contract.
type PassportContractOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_PassportContract *PassportContractFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*PassportContractOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _PassportContract.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PassportContractOwnershipRenouncedIterator{contract: _PassportContract.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_PassportContract *PassportContractFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *PassportContractOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _PassportContract.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PassportContractOwnershipRenounced)
				if err := _PassportContract.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// PassportContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PassportContract contract.
type PassportContractOwnershipTransferredIterator struct {
	Event *PassportContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PassportContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PassportContractOwnershipTransferred)
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
		it.Event = new(PassportContractOwnershipTransferred)
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
func (it *PassportContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PassportContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PassportContractOwnershipTransferred represents a OwnershipTransferred event raised by the PassportContract contract.
type PassportContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_PassportContract *PassportContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PassportContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PassportContract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PassportContractOwnershipTransferredIterator{contract: _PassportContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_PassportContract *PassportContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PassportContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PassportContract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PassportContractOwnershipTransferred)
				if err := _PassportContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// PassportContractPassportLogicRegistryChangedIterator is returned from FilterPassportLogicRegistryChanged and is used to iterate over the raw logs and unpacked data for PassportLogicRegistryChanged events raised by the PassportContract contract.
type PassportContractPassportLogicRegistryChangedIterator struct {
	Event *PassportContractPassportLogicRegistryChanged // Event containing the contract specifics and raw log

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
func (it *PassportContractPassportLogicRegistryChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PassportContractPassportLogicRegistryChanged)
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
		it.Event = new(PassportContractPassportLogicRegistryChanged)
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
func (it *PassportContractPassportLogicRegistryChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PassportContractPassportLogicRegistryChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PassportContractPassportLogicRegistryChanged represents a PassportLogicRegistryChanged event raised by the PassportContract contract.
type PassportContractPassportLogicRegistryChanged struct {
	PreviousRegistry common.Address
	NewRegistry      common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterPassportLogicRegistryChanged is a free log retrieval operation binding the contract event 0x5c2abfd67230c0e47d6de28402bfe206c7a57283cba891416ed657fd70a714c2.
//
// Solidity: e PassportLogicRegistryChanged(previousRegistry indexed address, newRegistry indexed address)
func (_PassportContract *PassportContractFilterer) FilterPassportLogicRegistryChanged(opts *bind.FilterOpts, previousRegistry []common.Address, newRegistry []common.Address) (*PassportContractPassportLogicRegistryChangedIterator, error) {

	var previousRegistryRule []interface{}
	for _, previousRegistryItem := range previousRegistry {
		previousRegistryRule = append(previousRegistryRule, previousRegistryItem)
	}
	var newRegistryRule []interface{}
	for _, newRegistryItem := range newRegistry {
		newRegistryRule = append(newRegistryRule, newRegistryItem)
	}

	logs, sub, err := _PassportContract.contract.FilterLogs(opts, "PassportLogicRegistryChanged", previousRegistryRule, newRegistryRule)
	if err != nil {
		return nil, err
	}
	return &PassportContractPassportLogicRegistryChangedIterator{contract: _PassportContract.contract, event: "PassportLogicRegistryChanged", logs: logs, sub: sub}, nil
}

// WatchPassportLogicRegistryChanged is a free log subscription operation binding the contract event 0x5c2abfd67230c0e47d6de28402bfe206c7a57283cba891416ed657fd70a714c2.
//
// Solidity: e PassportLogicRegistryChanged(previousRegistry indexed address, newRegistry indexed address)
func (_PassportContract *PassportContractFilterer) WatchPassportLogicRegistryChanged(opts *bind.WatchOpts, sink chan<- *PassportContractPassportLogicRegistryChanged, previousRegistry []common.Address, newRegistry []common.Address) (event.Subscription, error) {

	var previousRegistryRule []interface{}
	for _, previousRegistryItem := range previousRegistry {
		previousRegistryRule = append(previousRegistryRule, previousRegistryItem)
	}
	var newRegistryRule []interface{}
	for _, newRegistryItem := range newRegistry {
		newRegistryRule = append(newRegistryRule, newRegistryItem)
	}

	logs, sub, err := _PassportContract.contract.WatchLogs(opts, "PassportLogicRegistryChanged", previousRegistryRule, newRegistryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PassportContractPassportLogicRegistryChanged)
				if err := _PassportContract.contract.UnpackLog(event, "PassportLogicRegistryChanged", log); err != nil {
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
