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

// IPassportLogicContractABI is the input ABI used to generate the binding from.
const IPassportLogicContractABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes32\"},{\"name\":\"_value\",\"type\":\"address\"}],\"name\":\"setAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes32\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"setUint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes32\"},{\"name\":\"_value\",\"type\":\"int256\"}],\"name\":\"setInt\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes32\"},{\"name\":\"_value\",\"type\":\"bool\"}],\"name\":\"setBool\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes32\"},{\"name\":\"_value\",\"type\":\"string\"}],\"name\":\"setString\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes32\"},{\"name\":\"_value\",\"type\":\"bytes\"}],\"name\":\"setBytes\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes32\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"setTxDataBlockNumber\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes32\"}],\"name\":\"deleteAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes32\"}],\"name\":\"deleteUint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes32\"}],\"name\":\"deleteInt\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes32\"}],\"name\":\"deleteBool\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes32\"}],\"name\":\"deleteString\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes32\"}],\"name\":\"deleteBytes\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes32\"}],\"name\":\"deleteTxDataBlockNumber\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_factProvider\",\"type\":\"address\"},{\"name\":\"_key\",\"type\":\"bytes32\"}],\"name\":\"getAddress\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"},{\"name\":\"value\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_factProvider\",\"type\":\"address\"},{\"name\":\"_key\",\"type\":\"bytes32\"}],\"name\":\"getUint\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_factProvider\",\"type\":\"address\"},{\"name\":\"_key\",\"type\":\"bytes32\"}],\"name\":\"getInt\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"},{\"name\":\"value\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_factProvider\",\"type\":\"address\"},{\"name\":\"_key\",\"type\":\"bytes32\"}],\"name\":\"getBool\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"},{\"name\":\"value\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_factProvider\",\"type\":\"address\"},{\"name\":\"_key\",\"type\":\"bytes32\"}],\"name\":\"getString\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"},{\"name\":\"value\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_factProvider\",\"type\":\"address\"},{\"name\":\"_key\",\"type\":\"bytes32\"}],\"name\":\"getBytes\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"},{\"name\":\"value\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_factProvider\",\"type\":\"address\"},{\"name\":\"_key\",\"type\":\"bytes32\"}],\"name\":\"getTxDataBlockNumber\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"},{\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IPassportLogicContract is an auto generated Go binding around an Ethereum contract.
type IPassportLogicContract struct {
	IPassportLogicContractCaller     // Read-only binding to the contract
	IPassportLogicContractTransactor // Write-only binding to the contract
	IPassportLogicContractFilterer   // Log filterer for contract events
}

// IPassportLogicContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type IPassportLogicContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPassportLogicContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IPassportLogicContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPassportLogicContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPassportLogicContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPassportLogicContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPassportLogicContractSession struct {
	Contract     *IPassportLogicContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IPassportLogicContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPassportLogicContractCallerSession struct {
	Contract *IPassportLogicContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// IPassportLogicContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPassportLogicContractTransactorSession struct {
	Contract     *IPassportLogicContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// IPassportLogicContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type IPassportLogicContractRaw struct {
	Contract *IPassportLogicContract // Generic contract binding to access the raw methods on
}

// IPassportLogicContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPassportLogicContractCallerRaw struct {
	Contract *IPassportLogicContractCaller // Generic read-only contract binding to access the raw methods on
}

// IPassportLogicContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPassportLogicContractTransactorRaw struct {
	Contract *IPassportLogicContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIPassportLogicContract creates a new instance of IPassportLogicContract, bound to a specific deployed contract.
func NewIPassportLogicContract(address common.Address, backend bind.ContractBackend) (*IPassportLogicContract, error) {
	contract, err := bindIPassportLogicContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPassportLogicContract{IPassportLogicContractCaller: IPassportLogicContractCaller{contract: contract}, IPassportLogicContractTransactor: IPassportLogicContractTransactor{contract: contract}, IPassportLogicContractFilterer: IPassportLogicContractFilterer{contract: contract}}, nil
}

// NewIPassportLogicContractCaller creates a new read-only instance of IPassportLogicContract, bound to a specific deployed contract.
func NewIPassportLogicContractCaller(address common.Address, caller bind.ContractCaller) (*IPassportLogicContractCaller, error) {
	contract, err := bindIPassportLogicContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPassportLogicContractCaller{contract: contract}, nil
}

// NewIPassportLogicContractTransactor creates a new write-only instance of IPassportLogicContract, bound to a specific deployed contract.
func NewIPassportLogicContractTransactor(address common.Address, transactor bind.ContractTransactor) (*IPassportLogicContractTransactor, error) {
	contract, err := bindIPassportLogicContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPassportLogicContractTransactor{contract: contract}, nil
}

// NewIPassportLogicContractFilterer creates a new log filterer instance of IPassportLogicContract, bound to a specific deployed contract.
func NewIPassportLogicContractFilterer(address common.Address, filterer bind.ContractFilterer) (*IPassportLogicContractFilterer, error) {
	contract, err := bindIPassportLogicContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPassportLogicContractFilterer{contract: contract}, nil
}

// bindIPassportLogicContract binds a generic wrapper to an already deployed contract.
func bindIPassportLogicContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IPassportLogicContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPassportLogicContract *IPassportLogicContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IPassportLogicContract.Contract.IPassportLogicContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPassportLogicContract *IPassportLogicContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPassportLogicContract.Contract.IPassportLogicContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPassportLogicContract *IPassportLogicContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPassportLogicContract.Contract.IPassportLogicContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPassportLogicContract *IPassportLogicContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IPassportLogicContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPassportLogicContract *IPassportLogicContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPassportLogicContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPassportLogicContract *IPassportLogicContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPassportLogicContract.Contract.contract.Transact(opts, method, params...)
}

// GetAddress is a free data retrieval call binding the contract method 0x7ac4ed64.
//
// Solidity: function getAddress(_factProvider address, _key bytes32) constant returns(success bool, value address)
func (_IPassportLogicContract *IPassportLogicContractCaller) GetAddress(opts *bind.CallOpts, _factProvider common.Address, _key [32]byte) (struct {
	Success bool
	Value   common.Address
}, error) {
	ret := new(struct {
		Success bool
		Value   common.Address
	})
	out := ret
	err := _IPassportLogicContract.contract.Call(opts, out, "getAddress", _factProvider, _key)
	return *ret, err
}

// GetAddress is a free data retrieval call binding the contract method 0x7ac4ed64.
//
// Solidity: function getAddress(_factProvider address, _key bytes32) constant returns(success bool, value address)
func (_IPassportLogicContract *IPassportLogicContractSession) GetAddress(_factProvider common.Address, _key [32]byte) (struct {
	Success bool
	Value   common.Address
}, error) {
	return _IPassportLogicContract.Contract.GetAddress(&_IPassportLogicContract.CallOpts, _factProvider, _key)
}

// GetAddress is a free data retrieval call binding the contract method 0x7ac4ed64.
//
// Solidity: function getAddress(_factProvider address, _key bytes32) constant returns(success bool, value address)
func (_IPassportLogicContract *IPassportLogicContractCallerSession) GetAddress(_factProvider common.Address, _key [32]byte) (struct {
	Success bool
	Value   common.Address
}, error) {
	return _IPassportLogicContract.Contract.GetAddress(&_IPassportLogicContract.CallOpts, _factProvider, _key)
}

// GetBool is a free data retrieval call binding the contract method 0x9d74b37d.
//
// Solidity: function getBool(_factProvider address, _key bytes32) constant returns(success bool, value bool)
func (_IPassportLogicContract *IPassportLogicContractCaller) GetBool(opts *bind.CallOpts, _factProvider common.Address, _key [32]byte) (struct {
	Success bool
	Value   bool
}, error) {
	ret := new(struct {
		Success bool
		Value   bool
	})
	out := ret
	err := _IPassportLogicContract.contract.Call(opts, out, "getBool", _factProvider, _key)
	return *ret, err
}

// GetBool is a free data retrieval call binding the contract method 0x9d74b37d.
//
// Solidity: function getBool(_factProvider address, _key bytes32) constant returns(success bool, value bool)
func (_IPassportLogicContract *IPassportLogicContractSession) GetBool(_factProvider common.Address, _key [32]byte) (struct {
	Success bool
	Value   bool
}, error) {
	return _IPassportLogicContract.Contract.GetBool(&_IPassportLogicContract.CallOpts, _factProvider, _key)
}

// GetBool is a free data retrieval call binding the contract method 0x9d74b37d.
//
// Solidity: function getBool(_factProvider address, _key bytes32) constant returns(success bool, value bool)
func (_IPassportLogicContract *IPassportLogicContractCallerSession) GetBool(_factProvider common.Address, _key [32]byte) (struct {
	Success bool
	Value   bool
}, error) {
	return _IPassportLogicContract.Contract.GetBool(&_IPassportLogicContract.CallOpts, _factProvider, _key)
}

// GetBytes is a free data retrieval call binding the contract method 0x6556f767.
//
// Solidity: function getBytes(_factProvider address, _key bytes32) constant returns(success bool, value bytes)
func (_IPassportLogicContract *IPassportLogicContractCaller) GetBytes(opts *bind.CallOpts, _factProvider common.Address, _key [32]byte) (struct {
	Success bool
	Value   []byte
}, error) {
	ret := new(struct {
		Success bool
		Value   []byte
	})
	out := ret
	err := _IPassportLogicContract.contract.Call(opts, out, "getBytes", _factProvider, _key)
	return *ret, err
}

// GetBytes is a free data retrieval call binding the contract method 0x6556f767.
//
// Solidity: function getBytes(_factProvider address, _key bytes32) constant returns(success bool, value bytes)
func (_IPassportLogicContract *IPassportLogicContractSession) GetBytes(_factProvider common.Address, _key [32]byte) (struct {
	Success bool
	Value   []byte
}, error) {
	return _IPassportLogicContract.Contract.GetBytes(&_IPassportLogicContract.CallOpts, _factProvider, _key)
}

// GetBytes is a free data retrieval call binding the contract method 0x6556f767.
//
// Solidity: function getBytes(_factProvider address, _key bytes32) constant returns(success bool, value bytes)
func (_IPassportLogicContract *IPassportLogicContractCallerSession) GetBytes(_factProvider common.Address, _key [32]byte) (struct {
	Success bool
	Value   []byte
}, error) {
	return _IPassportLogicContract.Contract.GetBytes(&_IPassportLogicContract.CallOpts, _factProvider, _key)
}

// GetInt is a free data retrieval call binding the contract method 0x95ee8bae.
//
// Solidity: function getInt(_factProvider address, _key bytes32) constant returns(success bool, value int256)
func (_IPassportLogicContract *IPassportLogicContractCaller) GetInt(opts *bind.CallOpts, _factProvider common.Address, _key [32]byte) (struct {
	Success bool
	Value   *big.Int
}, error) {
	ret := new(struct {
		Success bool
		Value   *big.Int
	})
	out := ret
	err := _IPassportLogicContract.contract.Call(opts, out, "getInt", _factProvider, _key)
	return *ret, err
}

// GetInt is a free data retrieval call binding the contract method 0x95ee8bae.
//
// Solidity: function getInt(_factProvider address, _key bytes32) constant returns(success bool, value int256)
func (_IPassportLogicContract *IPassportLogicContractSession) GetInt(_factProvider common.Address, _key [32]byte) (struct {
	Success bool
	Value   *big.Int
}, error) {
	return _IPassportLogicContract.Contract.GetInt(&_IPassportLogicContract.CallOpts, _factProvider, _key)
}

// GetInt is a free data retrieval call binding the contract method 0x95ee8bae.
//
// Solidity: function getInt(_factProvider address, _key bytes32) constant returns(success bool, value int256)
func (_IPassportLogicContract *IPassportLogicContractCallerSession) GetInt(_factProvider common.Address, _key [32]byte) (struct {
	Success bool
	Value   *big.Int
}, error) {
	return _IPassportLogicContract.Contract.GetInt(&_IPassportLogicContract.CallOpts, _factProvider, _key)
}

// GetString is a free data retrieval call binding the contract method 0xe318de73.
//
// Solidity: function getString(_factProvider address, _key bytes32) constant returns(success bool, value string)
func (_IPassportLogicContract *IPassportLogicContractCaller) GetString(opts *bind.CallOpts, _factProvider common.Address, _key [32]byte) (struct {
	Success bool
	Value   string
}, error) {
	ret := new(struct {
		Success bool
		Value   string
	})
	out := ret
	err := _IPassportLogicContract.contract.Call(opts, out, "getString", _factProvider, _key)
	return *ret, err
}

// GetString is a free data retrieval call binding the contract method 0xe318de73.
//
// Solidity: function getString(_factProvider address, _key bytes32) constant returns(success bool, value string)
func (_IPassportLogicContract *IPassportLogicContractSession) GetString(_factProvider common.Address, _key [32]byte) (struct {
	Success bool
	Value   string
}, error) {
	return _IPassportLogicContract.Contract.GetString(&_IPassportLogicContract.CallOpts, _factProvider, _key)
}

// GetString is a free data retrieval call binding the contract method 0xe318de73.
//
// Solidity: function getString(_factProvider address, _key bytes32) constant returns(success bool, value string)
func (_IPassportLogicContract *IPassportLogicContractCallerSession) GetString(_factProvider common.Address, _key [32]byte) (struct {
	Success bool
	Value   string
}, error) {
	return _IPassportLogicContract.Contract.GetString(&_IPassportLogicContract.CallOpts, _factProvider, _key)
}

// GetTxDataBlockNumber is a free data retrieval call binding the contract method 0x174a6277.
//
// Solidity: function getTxDataBlockNumber(_factProvider address, _key bytes32) constant returns(success bool, blockNumber uint256)
func (_IPassportLogicContract *IPassportLogicContractCaller) GetTxDataBlockNumber(opts *bind.CallOpts, _factProvider common.Address, _key [32]byte) (struct {
	Success     bool
	BlockNumber *big.Int
}, error) {
	ret := new(struct {
		Success     bool
		BlockNumber *big.Int
	})
	out := ret
	err := _IPassportLogicContract.contract.Call(opts, out, "getTxDataBlockNumber", _factProvider, _key)
	return *ret, err
}

// GetTxDataBlockNumber is a free data retrieval call binding the contract method 0x174a6277.
//
// Solidity: function getTxDataBlockNumber(_factProvider address, _key bytes32) constant returns(success bool, blockNumber uint256)
func (_IPassportLogicContract *IPassportLogicContractSession) GetTxDataBlockNumber(_factProvider common.Address, _key [32]byte) (struct {
	Success     bool
	BlockNumber *big.Int
}, error) {
	return _IPassportLogicContract.Contract.GetTxDataBlockNumber(&_IPassportLogicContract.CallOpts, _factProvider, _key)
}

// GetTxDataBlockNumber is a free data retrieval call binding the contract method 0x174a6277.
//
// Solidity: function getTxDataBlockNumber(_factProvider address, _key bytes32) constant returns(success bool, blockNumber uint256)
func (_IPassportLogicContract *IPassportLogicContractCallerSession) GetTxDataBlockNumber(_factProvider common.Address, _key [32]byte) (struct {
	Success     bool
	BlockNumber *big.Int
}, error) {
	return _IPassportLogicContract.Contract.GetTxDataBlockNumber(&_IPassportLogicContract.CallOpts, _factProvider, _key)
}

// GetUint is a free data retrieval call binding the contract method 0x71658552.
//
// Solidity: function getUint(_factProvider address, _key bytes32) constant returns(success bool, value uint256)
func (_IPassportLogicContract *IPassportLogicContractCaller) GetUint(opts *bind.CallOpts, _factProvider common.Address, _key [32]byte) (struct {
	Success bool
	Value   *big.Int
}, error) {
	ret := new(struct {
		Success bool
		Value   *big.Int
	})
	out := ret
	err := _IPassportLogicContract.contract.Call(opts, out, "getUint", _factProvider, _key)
	return *ret, err
}

// GetUint is a free data retrieval call binding the contract method 0x71658552.
//
// Solidity: function getUint(_factProvider address, _key bytes32) constant returns(success bool, value uint256)
func (_IPassportLogicContract *IPassportLogicContractSession) GetUint(_factProvider common.Address, _key [32]byte) (struct {
	Success bool
	Value   *big.Int
}, error) {
	return _IPassportLogicContract.Contract.GetUint(&_IPassportLogicContract.CallOpts, _factProvider, _key)
}

// GetUint is a free data retrieval call binding the contract method 0x71658552.
//
// Solidity: function getUint(_factProvider address, _key bytes32) constant returns(success bool, value uint256)
func (_IPassportLogicContract *IPassportLogicContractCallerSession) GetUint(_factProvider common.Address, _key [32]byte) (struct {
	Success bool
	Value   *big.Int
}, error) {
	return _IPassportLogicContract.Contract.GetUint(&_IPassportLogicContract.CallOpts, _factProvider, _key)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_IPassportLogicContract *IPassportLogicContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IPassportLogicContract.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_IPassportLogicContract *IPassportLogicContractSession) Owner() (common.Address, error) {
	return _IPassportLogicContract.Contract.Owner(&_IPassportLogicContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_IPassportLogicContract *IPassportLogicContractCallerSession) Owner() (common.Address, error) {
	return _IPassportLogicContract.Contract.Owner(&_IPassportLogicContract.CallOpts)
}

// DeleteAddress is a paid mutator transaction binding the contract method 0x0e14a376.
//
// Solidity: function deleteAddress(_key bytes32) returns()
func (_IPassportLogicContract *IPassportLogicContractTransactor) DeleteAddress(opts *bind.TransactOpts, _key [32]byte) (*types.Transaction, error) {
	return _IPassportLogicContract.contract.Transact(opts, "deleteAddress", _key)
}

// DeleteAddress is a paid mutator transaction binding the contract method 0x0e14a376.
//
// Solidity: function deleteAddress(_key bytes32) returns()
func (_IPassportLogicContract *IPassportLogicContractSession) DeleteAddress(_key [32]byte) (*types.Transaction, error) {
	return _IPassportLogicContract.Contract.DeleteAddress(&_IPassportLogicContract.TransactOpts, _key)
}

// DeleteAddress is a paid mutator transaction binding the contract method 0x0e14a376.
//
// Solidity: function deleteAddress(_key bytes32) returns()
func (_IPassportLogicContract *IPassportLogicContractTransactorSession) DeleteAddress(_key [32]byte) (*types.Transaction, error) {
	return _IPassportLogicContract.Contract.DeleteAddress(&_IPassportLogicContract.TransactOpts, _key)
}

// DeleteBool is a paid mutator transaction binding the contract method 0x2c62ff2d.
//
// Solidity: function deleteBool(_key bytes32) returns()
func (_IPassportLogicContract *IPassportLogicContractTransactor) DeleteBool(opts *bind.TransactOpts, _key [32]byte) (*types.Transaction, error) {
	return _IPassportLogicContract.contract.Transact(opts, "deleteBool", _key)
}

// DeleteBool is a paid mutator transaction binding the contract method 0x2c62ff2d.
//
// Solidity: function deleteBool(_key bytes32) returns()
func (_IPassportLogicContract *IPassportLogicContractSession) DeleteBool(_key [32]byte) (*types.Transaction, error) {
	return _IPassportLogicContract.Contract.DeleteBool(&_IPassportLogicContract.TransactOpts, _key)
}

// DeleteBool is a paid mutator transaction binding the contract method 0x2c62ff2d.
//
// Solidity: function deleteBool(_key bytes32) returns()
func (_IPassportLogicContract *IPassportLogicContractTransactorSession) DeleteBool(_key [32]byte) (*types.Transaction, error) {
	return _IPassportLogicContract.Contract.DeleteBool(&_IPassportLogicContract.TransactOpts, _key)
}

// DeleteBytes is a paid mutator transaction binding the contract method 0x616b59f6.
//
// Solidity: function deleteBytes(_key bytes32) returns()
func (_IPassportLogicContract *IPassportLogicContractTransactor) DeleteBytes(opts *bind.TransactOpts, _key [32]byte) (*types.Transaction, error) {
	return _IPassportLogicContract.contract.Transact(opts, "deleteBytes", _key)
}

// DeleteBytes is a paid mutator transaction binding the contract method 0x616b59f6.
//
// Solidity: function deleteBytes(_key bytes32) returns()
func (_IPassportLogicContract *IPassportLogicContractSession) DeleteBytes(_key [32]byte) (*types.Transaction, error) {
	return _IPassportLogicContract.Contract.DeleteBytes(&_IPassportLogicContract.TransactOpts, _key)
}

// DeleteBytes is a paid mutator transaction binding the contract method 0x616b59f6.
//
// Solidity: function deleteBytes(_key bytes32) returns()
func (_IPassportLogicContract *IPassportLogicContractTransactorSession) DeleteBytes(_key [32]byte) (*types.Transaction, error) {
	return _IPassportLogicContract.Contract.DeleteBytes(&_IPassportLogicContract.TransactOpts, _key)
}

// DeleteInt is a paid mutator transaction binding the contract method 0x8c160095.
//
// Solidity: function deleteInt(_key bytes32) returns()
func (_IPassportLogicContract *IPassportLogicContractTransactor) DeleteInt(opts *bind.TransactOpts, _key [32]byte) (*types.Transaction, error) {
	return _IPassportLogicContract.contract.Transact(opts, "deleteInt", _key)
}

// DeleteInt is a paid mutator transaction binding the contract method 0x8c160095.
//
// Solidity: function deleteInt(_key bytes32) returns()
func (_IPassportLogicContract *IPassportLogicContractSession) DeleteInt(_key [32]byte) (*types.Transaction, error) {
	return _IPassportLogicContract.Contract.DeleteInt(&_IPassportLogicContract.TransactOpts, _key)
}

// DeleteInt is a paid mutator transaction binding the contract method 0x8c160095.
//
// Solidity: function deleteInt(_key bytes32) returns()
func (_IPassportLogicContract *IPassportLogicContractTransactorSession) DeleteInt(_key [32]byte) (*types.Transaction, error) {
	return _IPassportLogicContract.Contract.DeleteInt(&_IPassportLogicContract.TransactOpts, _key)
}

// DeleteString is a paid mutator transaction binding the contract method 0xf6bb3cc4.
//
// Solidity: function deleteString(_key bytes32) returns()
func (_IPassportLogicContract *IPassportLogicContractTransactor) DeleteString(opts *bind.TransactOpts, _key [32]byte) (*types.Transaction, error) {
	return _IPassportLogicContract.contract.Transact(opts, "deleteString", _key)
}

// DeleteString is a paid mutator transaction binding the contract method 0xf6bb3cc4.
//
// Solidity: function deleteString(_key bytes32) returns()
func (_IPassportLogicContract *IPassportLogicContractSession) DeleteString(_key [32]byte) (*types.Transaction, error) {
	return _IPassportLogicContract.Contract.DeleteString(&_IPassportLogicContract.TransactOpts, _key)
}

// DeleteString is a paid mutator transaction binding the contract method 0xf6bb3cc4.
//
// Solidity: function deleteString(_key bytes32) returns()
func (_IPassportLogicContract *IPassportLogicContractTransactorSession) DeleteString(_key [32]byte) (*types.Transaction, error) {
	return _IPassportLogicContract.Contract.DeleteString(&_IPassportLogicContract.TransactOpts, _key)
}

// DeleteTxDataBlockNumber is a paid mutator transaction binding the contract method 0xa2b6cbe1.
//
// Solidity: function deleteTxDataBlockNumber(_key bytes32) returns()
func (_IPassportLogicContract *IPassportLogicContractTransactor) DeleteTxDataBlockNumber(opts *bind.TransactOpts, _key [32]byte) (*types.Transaction, error) {
	return _IPassportLogicContract.contract.Transact(opts, "deleteTxDataBlockNumber", _key)
}

// DeleteTxDataBlockNumber is a paid mutator transaction binding the contract method 0xa2b6cbe1.
//
// Solidity: function deleteTxDataBlockNumber(_key bytes32) returns()
func (_IPassportLogicContract *IPassportLogicContractSession) DeleteTxDataBlockNumber(_key [32]byte) (*types.Transaction, error) {
	return _IPassportLogicContract.Contract.DeleteTxDataBlockNumber(&_IPassportLogicContract.TransactOpts, _key)
}

// DeleteTxDataBlockNumber is a paid mutator transaction binding the contract method 0xa2b6cbe1.
//
// Solidity: function deleteTxDataBlockNumber(_key bytes32) returns()
func (_IPassportLogicContract *IPassportLogicContractTransactorSession) DeleteTxDataBlockNumber(_key [32]byte) (*types.Transaction, error) {
	return _IPassportLogicContract.Contract.DeleteTxDataBlockNumber(&_IPassportLogicContract.TransactOpts, _key)
}

// DeleteUint is a paid mutator transaction binding the contract method 0xe2b202bf.
//
// Solidity: function deleteUint(_key bytes32) returns()
func (_IPassportLogicContract *IPassportLogicContractTransactor) DeleteUint(opts *bind.TransactOpts, _key [32]byte) (*types.Transaction, error) {
	return _IPassportLogicContract.contract.Transact(opts, "deleteUint", _key)
}

// DeleteUint is a paid mutator transaction binding the contract method 0xe2b202bf.
//
// Solidity: function deleteUint(_key bytes32) returns()
func (_IPassportLogicContract *IPassportLogicContractSession) DeleteUint(_key [32]byte) (*types.Transaction, error) {
	return _IPassportLogicContract.Contract.DeleteUint(&_IPassportLogicContract.TransactOpts, _key)
}

// DeleteUint is a paid mutator transaction binding the contract method 0xe2b202bf.
//
// Solidity: function deleteUint(_key bytes32) returns()
func (_IPassportLogicContract *IPassportLogicContractTransactorSession) DeleteUint(_key [32]byte) (*types.Transaction, error) {
	return _IPassportLogicContract.Contract.DeleteUint(&_IPassportLogicContract.TransactOpts, _key)
}

// SetAddress is a paid mutator transaction binding the contract method 0xca446dd9.
//
// Solidity: function setAddress(_key bytes32, _value address) returns()
func (_IPassportLogicContract *IPassportLogicContractTransactor) SetAddress(opts *bind.TransactOpts, _key [32]byte, _value common.Address) (*types.Transaction, error) {
	return _IPassportLogicContract.contract.Transact(opts, "setAddress", _key, _value)
}

// SetAddress is a paid mutator transaction binding the contract method 0xca446dd9.
//
// Solidity: function setAddress(_key bytes32, _value address) returns()
func (_IPassportLogicContract *IPassportLogicContractSession) SetAddress(_key [32]byte, _value common.Address) (*types.Transaction, error) {
	return _IPassportLogicContract.Contract.SetAddress(&_IPassportLogicContract.TransactOpts, _key, _value)
}

// SetAddress is a paid mutator transaction binding the contract method 0xca446dd9.
//
// Solidity: function setAddress(_key bytes32, _value address) returns()
func (_IPassportLogicContract *IPassportLogicContractTransactorSession) SetAddress(_key [32]byte, _value common.Address) (*types.Transaction, error) {
	return _IPassportLogicContract.Contract.SetAddress(&_IPassportLogicContract.TransactOpts, _key, _value)
}

// SetBool is a paid mutator transaction binding the contract method 0xabfdcced.
//
// Solidity: function setBool(_key bytes32, _value bool) returns()
func (_IPassportLogicContract *IPassportLogicContractTransactor) SetBool(opts *bind.TransactOpts, _key [32]byte, _value bool) (*types.Transaction, error) {
	return _IPassportLogicContract.contract.Transact(opts, "setBool", _key, _value)
}

// SetBool is a paid mutator transaction binding the contract method 0xabfdcced.
//
// Solidity: function setBool(_key bytes32, _value bool) returns()
func (_IPassportLogicContract *IPassportLogicContractSession) SetBool(_key [32]byte, _value bool) (*types.Transaction, error) {
	return _IPassportLogicContract.Contract.SetBool(&_IPassportLogicContract.TransactOpts, _key, _value)
}

// SetBool is a paid mutator transaction binding the contract method 0xabfdcced.
//
// Solidity: function setBool(_key bytes32, _value bool) returns()
func (_IPassportLogicContract *IPassportLogicContractTransactorSession) SetBool(_key [32]byte, _value bool) (*types.Transaction, error) {
	return _IPassportLogicContract.Contract.SetBool(&_IPassportLogicContract.TransactOpts, _key, _value)
}

// SetBytes is a paid mutator transaction binding the contract method 0x2e28d084.
//
// Solidity: function setBytes(_key bytes32, _value bytes) returns()
func (_IPassportLogicContract *IPassportLogicContractTransactor) SetBytes(opts *bind.TransactOpts, _key [32]byte, _value []byte) (*types.Transaction, error) {
	return _IPassportLogicContract.contract.Transact(opts, "setBytes", _key, _value)
}

// SetBytes is a paid mutator transaction binding the contract method 0x2e28d084.
//
// Solidity: function setBytes(_key bytes32, _value bytes) returns()
func (_IPassportLogicContract *IPassportLogicContractSession) SetBytes(_key [32]byte, _value []byte) (*types.Transaction, error) {
	return _IPassportLogicContract.Contract.SetBytes(&_IPassportLogicContract.TransactOpts, _key, _value)
}

// SetBytes is a paid mutator transaction binding the contract method 0x2e28d084.
//
// Solidity: function setBytes(_key bytes32, _value bytes) returns()
func (_IPassportLogicContract *IPassportLogicContractTransactorSession) SetBytes(_key [32]byte, _value []byte) (*types.Transaction, error) {
	return _IPassportLogicContract.Contract.SetBytes(&_IPassportLogicContract.TransactOpts, _key, _value)
}

// SetInt is a paid mutator transaction binding the contract method 0x3e49bed0.
//
// Solidity: function setInt(_key bytes32, _value int256) returns()
func (_IPassportLogicContract *IPassportLogicContractTransactor) SetInt(opts *bind.TransactOpts, _key [32]byte, _value *big.Int) (*types.Transaction, error) {
	return _IPassportLogicContract.contract.Transact(opts, "setInt", _key, _value)
}

// SetInt is a paid mutator transaction binding the contract method 0x3e49bed0.
//
// Solidity: function setInt(_key bytes32, _value int256) returns()
func (_IPassportLogicContract *IPassportLogicContractSession) SetInt(_key [32]byte, _value *big.Int) (*types.Transaction, error) {
	return _IPassportLogicContract.Contract.SetInt(&_IPassportLogicContract.TransactOpts, _key, _value)
}

// SetInt is a paid mutator transaction binding the contract method 0x3e49bed0.
//
// Solidity: function setInt(_key bytes32, _value int256) returns()
func (_IPassportLogicContract *IPassportLogicContractTransactorSession) SetInt(_key [32]byte, _value *big.Int) (*types.Transaction, error) {
	return _IPassportLogicContract.Contract.SetInt(&_IPassportLogicContract.TransactOpts, _key, _value)
}

// SetString is a paid mutator transaction binding the contract method 0x6e899550.
//
// Solidity: function setString(_key bytes32, _value string) returns()
func (_IPassportLogicContract *IPassportLogicContractTransactor) SetString(opts *bind.TransactOpts, _key [32]byte, _value string) (*types.Transaction, error) {
	return _IPassportLogicContract.contract.Transact(opts, "setString", _key, _value)
}

// SetString is a paid mutator transaction binding the contract method 0x6e899550.
//
// Solidity: function setString(_key bytes32, _value string) returns()
func (_IPassportLogicContract *IPassportLogicContractSession) SetString(_key [32]byte, _value string) (*types.Transaction, error) {
	return _IPassportLogicContract.Contract.SetString(&_IPassportLogicContract.TransactOpts, _key, _value)
}

// SetString is a paid mutator transaction binding the contract method 0x6e899550.
//
// Solidity: function setString(_key bytes32, _value string) returns()
func (_IPassportLogicContract *IPassportLogicContractTransactorSession) SetString(_key [32]byte, _value string) (*types.Transaction, error) {
	return _IPassportLogicContract.Contract.SetString(&_IPassportLogicContract.TransactOpts, _key, _value)
}

// SetTxDataBlockNumber is a paid mutator transaction binding the contract method 0x5b2a372d.
//
// Solidity: function setTxDataBlockNumber(_key bytes32, _data bytes) returns()
func (_IPassportLogicContract *IPassportLogicContractTransactor) SetTxDataBlockNumber(opts *bind.TransactOpts, _key [32]byte, _data []byte) (*types.Transaction, error) {
	return _IPassportLogicContract.contract.Transact(opts, "setTxDataBlockNumber", _key, _data)
}

// SetTxDataBlockNumber is a paid mutator transaction binding the contract method 0x5b2a372d.
//
// Solidity: function setTxDataBlockNumber(_key bytes32, _data bytes) returns()
func (_IPassportLogicContract *IPassportLogicContractSession) SetTxDataBlockNumber(_key [32]byte, _data []byte) (*types.Transaction, error) {
	return _IPassportLogicContract.Contract.SetTxDataBlockNumber(&_IPassportLogicContract.TransactOpts, _key, _data)
}

// SetTxDataBlockNumber is a paid mutator transaction binding the contract method 0x5b2a372d.
//
// Solidity: function setTxDataBlockNumber(_key bytes32, _data bytes) returns()
func (_IPassportLogicContract *IPassportLogicContractTransactorSession) SetTxDataBlockNumber(_key [32]byte, _data []byte) (*types.Transaction, error) {
	return _IPassportLogicContract.Contract.SetTxDataBlockNumber(&_IPassportLogicContract.TransactOpts, _key, _data)
}

// SetUint is a paid mutator transaction binding the contract method 0xe2a4853a.
//
// Solidity: function setUint(_key bytes32, _value uint256) returns()
func (_IPassportLogicContract *IPassportLogicContractTransactor) SetUint(opts *bind.TransactOpts, _key [32]byte, _value *big.Int) (*types.Transaction, error) {
	return _IPassportLogicContract.contract.Transact(opts, "setUint", _key, _value)
}

// SetUint is a paid mutator transaction binding the contract method 0xe2a4853a.
//
// Solidity: function setUint(_key bytes32, _value uint256) returns()
func (_IPassportLogicContract *IPassportLogicContractSession) SetUint(_key [32]byte, _value *big.Int) (*types.Transaction, error) {
	return _IPassportLogicContract.Contract.SetUint(&_IPassportLogicContract.TransactOpts, _key, _value)
}

// SetUint is a paid mutator transaction binding the contract method 0xe2a4853a.
//
// Solidity: function setUint(_key bytes32, _value uint256) returns()
func (_IPassportLogicContract *IPassportLogicContractTransactorSession) SetUint(_key [32]byte, _value *big.Int) (*types.Transaction, error) {
	return _IPassportLogicContract.Contract.SetUint(&_IPassportLogicContract.TransactOpts, _key, _value)
}
