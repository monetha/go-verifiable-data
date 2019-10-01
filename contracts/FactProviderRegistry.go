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

// FactProviderRegistryContractABI is the input ABI used to generate the binding from.
const FactProviderRegistryContractABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"reclaimToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"factProviders\",\"outputs\":[{\"name\":\"initialized\",\"type\":\"bool\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"reputation_passport\",\"type\":\"address\"},{\"name\":\"website\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"reclaimEther\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"tokenFallback\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"factProvider\",\"type\":\"address\"}],\"name\":\"FactProviderAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"factProvider\",\"type\":\"address\"}],\"name\":\"FactProviderUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"factProvider\",\"type\":\"address\"}],\"name\":\"FactProviderDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_factProvider\",\"type\":\"address\"},{\"name\":\"_factProviderName\",\"type\":\"string\"},{\"name\":\"_factProviderReputationPassport\",\"type\":\"address\"},{\"name\":\"_factProviderWebsite\",\"type\":\"string\"}],\"name\":\"setFactProviderInfo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_factProvider\",\"type\":\"address\"}],\"name\":\"deleteFactProviderInfo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// FactProviderRegistryContractBin is the compiled bytecode used for deploying new contracts.
const FactProviderRegistryContractBin = `0x608060405260008054600160a060020a03191633179055341561002157600080fd5b610ae3806100306000396000f3006080604052600436106100985763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166317ffc32081146100a75780631df84d87146100ca57806353cb9293146100eb578063715018a61461020e57806385301a01146102235780638da5cb5b146102625780639f727c2714610293578063c0ee0b8a146102a8578063f2fde38b146102d9575b3480156100a457600080fd5b50005b3480156100b357600080fd5b506100c8600160a060020a03600435166102fa565b005b3480156100d657600080fd5b506100c8600160a060020a03600435166103c6565b3480156100f757600080fd5b5061010c600160a060020a0360043516610495565b60405180851515151581526020018060200184600160a060020a0316600160a060020a0316815260200180602001838103835286818151815260200191508051906020019080838360005b8381101561016f578181015183820152602001610157565b50505050905090810190601f16801561019c5780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b838110156101cf5781810151838201526020016101b7565b50505050905090810190601f1680156101fc5780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390f35b34801561021a57600080fd5b506100c86105e4565b34801561022f57600080fd5b506100c8600160a060020a0360048035821691602480358082019390810135926044351691606435908101910135610650565b34801561026e57600080fd5b5061027761081d565b60408051600160a060020a039092168252519081900360200190f35b34801561029f57600080fd5b506100c861082c565b3480156102b457600080fd5b506100c860048035600160a060020a031690602480359160443591820191013561087e565b3480156102e557600080fd5b506100c8600160a060020a0360043516610883565b60008054600160a060020a0316331461031257600080fd5b604080517f70a082310000000000000000000000000000000000000000000000000000000081523060048201529051600160a060020a038416916370a082319160248083019260209291908290030181600087803b15801561037357600080fd5b505af1158015610387573d6000803e3d6000fd5b505050506040513d602081101561039d57600080fd5b50516000549091506103c290600160a060020a0384811691168363ffffffff6108a316565b5050565b600054600160a060020a031633146103dd57600080fd5b600160a060020a03811660009081526001602052604090205460ff161561049257600160a060020a03811660009081526001602081905260408220805460ff19168155919061042e908301826109d8565b60028201805473ffffffffffffffffffffffffffffffffffffffff1916905561045b6003830160006109d8565b5050604051600160a060020a038216907fffcf9fb4992cc4ee4022ab3c2c5fa4914c18a25a2dc1614fc52c3ba8bbebf0aa90600090a25b50565b600160208181526000928352604092839020805481840180548651600261010097831615979097026000190190911695909504601f810185900485028601850190965285855260ff909116949193929091908301828280156105385780601f1061050d57610100808354040283529160200191610538565b820191906000526020600020905b81548152906001019060200180831161051b57829003601f168201915b505050506002838101546003850180546040805160206101006001851615026000190190931695909504601f81018390048302860183019091528085529596600160a060020a03909316959294509091908301828280156105da5780601f106105af576101008083540402835291602001916105da565b820191906000526020600020905b8154815290600101906020018083116105bd57829003601f168201915b5050505050905084565b600054600160a060020a031633146105fb57600080fd5b60008054604051600160a060020a03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a26000805473ffffffffffffffffffffffffffffffffffffffff19169055565b60008054600160a060020a0316331461066857600080fd5b50600160a060020a0386166000908152600160208181526040928390205483516080810185529283528351601f890183900483028101830190945287845260ff1692828201918990899081908401838280828437820191505050505050815260200185600160a060020a0316815260200184848080601f01602080910402602001604051908101604052809392919081815260200183838082843750505092909352505050600160a060020a03881660009081526001602081815260409092208351815460ff1916901515178155838301518051919361074d93850192910190610a1c565b50604082015160028201805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039092169190911790556060820151805161079c916003840191602090910190610a1c565b5090505080156107df57604051600160a060020a038816907ff5373d11a64f7dfd39311b28060d0cc68d73cde4dbb253a9cfd6b00f871932a490600090a2610814565b604051600160a060020a038816907f49fe72d660cc80a367041ae43bb1e4956396956549f81c0e6a0adfbc4a8698f790600090a25b50505050505050565b600054600160a060020a031681565b600054600160a060020a0316331461084357600080fd5b60008054604051600160a060020a0390911691303180156108fc02929091818181858888f19350505050158015610492573d6000803e3d6000fd5b600080fd5b600054600160a060020a0316331461089a57600080fd5b6104928161095b565b82600160a060020a031663a9059cbb83836040518363ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018083600160a060020a0316600160a060020a0316815260200182815260200192505050602060405180830381600087803b15801561091f57600080fd5b505af1158015610933573d6000803e3d6000fd5b505050506040513d602081101561094957600080fd5b5051151561095657600080fd5b505050565b600160a060020a038116151561097057600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b50805460018160011615610100020316600290046000825580601f106109fe5750610492565b601f0160209004906000526020600020908101906104929190610a9a565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610a5d57805160ff1916838001178555610a8a565b82800160010185558215610a8a579182015b82811115610a8a578251825591602001919060010190610a6f565b50610a96929150610a9a565b5090565b610ab491905b80821115610a965760008155600101610aa0565b905600a165627a7a723058205bd0c77a2c503097d0d053f99209c80b25c1108c5a0c8e77d66dccb5bdf9c1030029`

// DeployFactProviderRegistryContract deploys a new Ethereum contract, binding an instance of FactProviderRegistryContract to it.
func DeployFactProviderRegistryContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *FactProviderRegistryContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FactProviderRegistryContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FactProviderRegistryContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FactProviderRegistryContract{FactProviderRegistryContractCaller: FactProviderRegistryContractCaller{contract: contract}, FactProviderRegistryContractTransactor: FactProviderRegistryContractTransactor{contract: contract}, FactProviderRegistryContractFilterer: FactProviderRegistryContractFilterer{contract: contract}}, nil
}

// FactProviderRegistryContract is an auto generated Go binding around an Ethereum contract.
type FactProviderRegistryContract struct {
	FactProviderRegistryContractCaller     // Read-only binding to the contract
	FactProviderRegistryContractTransactor // Write-only binding to the contract
	FactProviderRegistryContractFilterer   // Log filterer for contract events
}

// FactProviderRegistryContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type FactProviderRegistryContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FactProviderRegistryContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FactProviderRegistryContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FactProviderRegistryContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FactProviderRegistryContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FactProviderRegistryContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FactProviderRegistryContractSession struct {
	Contract     *FactProviderRegistryContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                 // Call options to use throughout this session
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// FactProviderRegistryContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FactProviderRegistryContractCallerSession struct {
	Contract *FactProviderRegistryContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                       // Call options to use throughout this session
}

// FactProviderRegistryContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FactProviderRegistryContractTransactorSession struct {
	Contract     *FactProviderRegistryContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                       // Transaction auth options to use throughout this session
}

// FactProviderRegistryContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type FactProviderRegistryContractRaw struct {
	Contract *FactProviderRegistryContract // Generic contract binding to access the raw methods on
}

// FactProviderRegistryContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FactProviderRegistryContractCallerRaw struct {
	Contract *FactProviderRegistryContractCaller // Generic read-only contract binding to access the raw methods on
}

// FactProviderRegistryContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FactProviderRegistryContractTransactorRaw struct {
	Contract *FactProviderRegistryContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFactProviderRegistryContract creates a new instance of FactProviderRegistryContract, bound to a specific deployed contract.
func NewFactProviderRegistryContract(address common.Address, backend bind.ContractBackend) (*FactProviderRegistryContract, error) {
	contract, err := bindFactProviderRegistryContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FactProviderRegistryContract{FactProviderRegistryContractCaller: FactProviderRegistryContractCaller{contract: contract}, FactProviderRegistryContractTransactor: FactProviderRegistryContractTransactor{contract: contract}, FactProviderRegistryContractFilterer: FactProviderRegistryContractFilterer{contract: contract}}, nil
}

// NewFactProviderRegistryContractCaller creates a new read-only instance of FactProviderRegistryContract, bound to a specific deployed contract.
func NewFactProviderRegistryContractCaller(address common.Address, caller bind.ContractCaller) (*FactProviderRegistryContractCaller, error) {
	contract, err := bindFactProviderRegistryContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FactProviderRegistryContractCaller{contract: contract}, nil
}

// NewFactProviderRegistryContractTransactor creates a new write-only instance of FactProviderRegistryContract, bound to a specific deployed contract.
func NewFactProviderRegistryContractTransactor(address common.Address, transactor bind.ContractTransactor) (*FactProviderRegistryContractTransactor, error) {
	contract, err := bindFactProviderRegistryContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FactProviderRegistryContractTransactor{contract: contract}, nil
}

// NewFactProviderRegistryContractFilterer creates a new log filterer instance of FactProviderRegistryContract, bound to a specific deployed contract.
func NewFactProviderRegistryContractFilterer(address common.Address, filterer bind.ContractFilterer) (*FactProviderRegistryContractFilterer, error) {
	contract, err := bindFactProviderRegistryContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FactProviderRegistryContractFilterer{contract: contract}, nil
}

// bindFactProviderRegistryContract binds a generic wrapper to an already deployed contract.
func bindFactProviderRegistryContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FactProviderRegistryContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FactProviderRegistryContract *FactProviderRegistryContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FactProviderRegistryContract.Contract.FactProviderRegistryContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FactProviderRegistryContract *FactProviderRegistryContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FactProviderRegistryContract.Contract.FactProviderRegistryContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FactProviderRegistryContract *FactProviderRegistryContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FactProviderRegistryContract.Contract.FactProviderRegistryContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FactProviderRegistryContract *FactProviderRegistryContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FactProviderRegistryContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FactProviderRegistryContract *FactProviderRegistryContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FactProviderRegistryContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FactProviderRegistryContract *FactProviderRegistryContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FactProviderRegistryContract.Contract.contract.Transact(opts, method, params...)
}

// FactProviders is a free data retrieval call binding the contract method 0x53cb9293.
//
// Solidity: function factProviders(address ) constant returns(bool initialized, string name, address reputation_passport, string website)
func (_FactProviderRegistryContract *FactProviderRegistryContractCaller) FactProviders(opts *bind.CallOpts, arg0 common.Address) (struct {
	Initialized        bool
	Name               string
	ReputationPassport common.Address
	Website            string
}, error) {
	ret := new(struct {
		Initialized        bool
		Name               string
		ReputationPassport common.Address
		Website            string
	})
	out := ret
	err := _FactProviderRegistryContract.contract.Call(opts, out, "factProviders", arg0)
	return *ret, err
}

// FactProviders is a free data retrieval call binding the contract method 0x53cb9293.
//
// Solidity: function factProviders(address ) constant returns(bool initialized, string name, address reputation_passport, string website)
func (_FactProviderRegistryContract *FactProviderRegistryContractSession) FactProviders(arg0 common.Address) (struct {
	Initialized        bool
	Name               string
	ReputationPassport common.Address
	Website            string
}, error) {
	return _FactProviderRegistryContract.Contract.FactProviders(&_FactProviderRegistryContract.CallOpts, arg0)
}

// FactProviders is a free data retrieval call binding the contract method 0x53cb9293.
//
// Solidity: function factProviders(address ) constant returns(bool initialized, string name, address reputation_passport, string website)
func (_FactProviderRegistryContract *FactProviderRegistryContractCallerSession) FactProviders(arg0 common.Address) (struct {
	Initialized        bool
	Name               string
	ReputationPassport common.Address
	Website            string
}, error) {
	return _FactProviderRegistryContract.Contract.FactProviders(&_FactProviderRegistryContract.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_FactProviderRegistryContract *FactProviderRegistryContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FactProviderRegistryContract.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_FactProviderRegistryContract *FactProviderRegistryContractSession) Owner() (common.Address, error) {
	return _FactProviderRegistryContract.Contract.Owner(&_FactProviderRegistryContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_FactProviderRegistryContract *FactProviderRegistryContractCallerSession) Owner() (common.Address, error) {
	return _FactProviderRegistryContract.Contract.Owner(&_FactProviderRegistryContract.CallOpts)
}

// TokenFallback is a free data retrieval call binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(address _from, uint256 _value, bytes _data) constant returns()
func (_FactProviderRegistryContract *FactProviderRegistryContractCaller) TokenFallback(opts *bind.CallOpts, _from common.Address, _value *big.Int, _data []byte) error {
	var ()
	out := &[]interface{}{}
	err := _FactProviderRegistryContract.contract.Call(opts, out, "tokenFallback", _from, _value, _data)
	return err
}

// TokenFallback is a free data retrieval call binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(address _from, uint256 _value, bytes _data) constant returns()
func (_FactProviderRegistryContract *FactProviderRegistryContractSession) TokenFallback(_from common.Address, _value *big.Int, _data []byte) error {
	return _FactProviderRegistryContract.Contract.TokenFallback(&_FactProviderRegistryContract.CallOpts, _from, _value, _data)
}

// TokenFallback is a free data retrieval call binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(address _from, uint256 _value, bytes _data) constant returns()
func (_FactProviderRegistryContract *FactProviderRegistryContractCallerSession) TokenFallback(_from common.Address, _value *big.Int, _data []byte) error {
	return _FactProviderRegistryContract.Contract.TokenFallback(&_FactProviderRegistryContract.CallOpts, _from, _value, _data)
}

// DeleteFactProviderInfo is a paid mutator transaction binding the contract method 0x1df84d87.
//
// Solidity: function deleteFactProviderInfo(address _factProvider) returns()
func (_FactProviderRegistryContract *FactProviderRegistryContractTransactor) DeleteFactProviderInfo(opts *bind.TransactOpts, _factProvider common.Address) (*types.Transaction, error) {
	return _FactProviderRegistryContract.contract.Transact(opts, "deleteFactProviderInfo", _factProvider)
}

// DeleteFactProviderInfo is a paid mutator transaction binding the contract method 0x1df84d87.
//
// Solidity: function deleteFactProviderInfo(address _factProvider) returns()
func (_FactProviderRegistryContract *FactProviderRegistryContractSession) DeleteFactProviderInfo(_factProvider common.Address) (*types.Transaction, error) {
	return _FactProviderRegistryContract.Contract.DeleteFactProviderInfo(&_FactProviderRegistryContract.TransactOpts, _factProvider)
}

// DeleteFactProviderInfo is a paid mutator transaction binding the contract method 0x1df84d87.
//
// Solidity: function deleteFactProviderInfo(address _factProvider) returns()
func (_FactProviderRegistryContract *FactProviderRegistryContractTransactorSession) DeleteFactProviderInfo(_factProvider common.Address) (*types.Transaction, error) {
	return _FactProviderRegistryContract.Contract.DeleteFactProviderInfo(&_FactProviderRegistryContract.TransactOpts, _factProvider)
}

// ReclaimEther is a paid mutator transaction binding the contract method 0x9f727c27.
//
// Solidity: function reclaimEther() returns()
func (_FactProviderRegistryContract *FactProviderRegistryContractTransactor) ReclaimEther(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FactProviderRegistryContract.contract.Transact(opts, "reclaimEther")
}

// ReclaimEther is a paid mutator transaction binding the contract method 0x9f727c27.
//
// Solidity: function reclaimEther() returns()
func (_FactProviderRegistryContract *FactProviderRegistryContractSession) ReclaimEther() (*types.Transaction, error) {
	return _FactProviderRegistryContract.Contract.ReclaimEther(&_FactProviderRegistryContract.TransactOpts)
}

// ReclaimEther is a paid mutator transaction binding the contract method 0x9f727c27.
//
// Solidity: function reclaimEther() returns()
func (_FactProviderRegistryContract *FactProviderRegistryContractTransactorSession) ReclaimEther() (*types.Transaction, error) {
	return _FactProviderRegistryContract.Contract.ReclaimEther(&_FactProviderRegistryContract.TransactOpts)
}

// ReclaimToken is a paid mutator transaction binding the contract method 0x17ffc320.
//
// Solidity: function reclaimToken(address _token) returns()
func (_FactProviderRegistryContract *FactProviderRegistryContractTransactor) ReclaimToken(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _FactProviderRegistryContract.contract.Transact(opts, "reclaimToken", _token)
}

// ReclaimToken is a paid mutator transaction binding the contract method 0x17ffc320.
//
// Solidity: function reclaimToken(address _token) returns()
func (_FactProviderRegistryContract *FactProviderRegistryContractSession) ReclaimToken(_token common.Address) (*types.Transaction, error) {
	return _FactProviderRegistryContract.Contract.ReclaimToken(&_FactProviderRegistryContract.TransactOpts, _token)
}

// ReclaimToken is a paid mutator transaction binding the contract method 0x17ffc320.
//
// Solidity: function reclaimToken(address _token) returns()
func (_FactProviderRegistryContract *FactProviderRegistryContractTransactorSession) ReclaimToken(_token common.Address) (*types.Transaction, error) {
	return _FactProviderRegistryContract.Contract.ReclaimToken(&_FactProviderRegistryContract.TransactOpts, _token)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FactProviderRegistryContract *FactProviderRegistryContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FactProviderRegistryContract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FactProviderRegistryContract *FactProviderRegistryContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _FactProviderRegistryContract.Contract.RenounceOwnership(&_FactProviderRegistryContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FactProviderRegistryContract *FactProviderRegistryContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _FactProviderRegistryContract.Contract.RenounceOwnership(&_FactProviderRegistryContract.TransactOpts)
}

// SetFactProviderInfo is a paid mutator transaction binding the contract method 0x85301a01.
//
// Solidity: function setFactProviderInfo(address _factProvider, string _factProviderName, address _factProviderReputationPassport, string _factProviderWebsite) returns()
func (_FactProviderRegistryContract *FactProviderRegistryContractTransactor) SetFactProviderInfo(opts *bind.TransactOpts, _factProvider common.Address, _factProviderName string, _factProviderReputationPassport common.Address, _factProviderWebsite string) (*types.Transaction, error) {
	return _FactProviderRegistryContract.contract.Transact(opts, "setFactProviderInfo", _factProvider, _factProviderName, _factProviderReputationPassport, _factProviderWebsite)
}

// SetFactProviderInfo is a paid mutator transaction binding the contract method 0x85301a01.
//
// Solidity: function setFactProviderInfo(address _factProvider, string _factProviderName, address _factProviderReputationPassport, string _factProviderWebsite) returns()
func (_FactProviderRegistryContract *FactProviderRegistryContractSession) SetFactProviderInfo(_factProvider common.Address, _factProviderName string, _factProviderReputationPassport common.Address, _factProviderWebsite string) (*types.Transaction, error) {
	return _FactProviderRegistryContract.Contract.SetFactProviderInfo(&_FactProviderRegistryContract.TransactOpts, _factProvider, _factProviderName, _factProviderReputationPassport, _factProviderWebsite)
}

// SetFactProviderInfo is a paid mutator transaction binding the contract method 0x85301a01.
//
// Solidity: function setFactProviderInfo(address _factProvider, string _factProviderName, address _factProviderReputationPassport, string _factProviderWebsite) returns()
func (_FactProviderRegistryContract *FactProviderRegistryContractTransactorSession) SetFactProviderInfo(_factProvider common.Address, _factProviderName string, _factProviderReputationPassport common.Address, _factProviderWebsite string) (*types.Transaction, error) {
	return _FactProviderRegistryContract.Contract.SetFactProviderInfo(&_FactProviderRegistryContract.TransactOpts, _factProvider, _factProviderName, _factProviderReputationPassport, _factProviderWebsite)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_FactProviderRegistryContract *FactProviderRegistryContractTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _FactProviderRegistryContract.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_FactProviderRegistryContract *FactProviderRegistryContractSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _FactProviderRegistryContract.Contract.TransferOwnership(&_FactProviderRegistryContract.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_FactProviderRegistryContract *FactProviderRegistryContractTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _FactProviderRegistryContract.Contract.TransferOwnership(&_FactProviderRegistryContract.TransactOpts, _newOwner)
}

// FactProviderRegistryContractFactProviderAddedIterator is returned from FilterFactProviderAdded and is used to iterate over the raw logs and unpacked data for FactProviderAdded events raised by the FactProviderRegistryContract contract.
type FactProviderRegistryContractFactProviderAddedIterator struct {
	Event *FactProviderRegistryContractFactProviderAdded // Event containing the contract specifics and raw log

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
func (it *FactProviderRegistryContractFactProviderAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FactProviderRegistryContractFactProviderAdded)
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
		it.Event = new(FactProviderRegistryContractFactProviderAdded)
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
func (it *FactProviderRegistryContractFactProviderAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FactProviderRegistryContractFactProviderAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FactProviderRegistryContractFactProviderAdded represents a FactProviderAdded event raised by the FactProviderRegistryContract contract.
type FactProviderRegistryContractFactProviderAdded struct {
	FactProvider common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterFactProviderAdded is a free log retrieval operation binding the contract event 0x49fe72d660cc80a367041ae43bb1e4956396956549f81c0e6a0adfbc4a8698f7.
//
// Solidity: event FactProviderAdded(address indexed factProvider)
func (_FactProviderRegistryContract *FactProviderRegistryContractFilterer) FilterFactProviderAdded(opts *bind.FilterOpts, factProvider []common.Address) (*FactProviderRegistryContractFactProviderAddedIterator, error) {

	var factProviderRule []interface{}
	for _, factProviderItem := range factProvider {
		factProviderRule = append(factProviderRule, factProviderItem)
	}

	logs, sub, err := _FactProviderRegistryContract.contract.FilterLogs(opts, "FactProviderAdded", factProviderRule)
	if err != nil {
		return nil, err
	}
	return &FactProviderRegistryContractFactProviderAddedIterator{contract: _FactProviderRegistryContract.contract, event: "FactProviderAdded", logs: logs, sub: sub}, nil
}

// WatchFactProviderAdded is a free log subscription operation binding the contract event 0x49fe72d660cc80a367041ae43bb1e4956396956549f81c0e6a0adfbc4a8698f7.
//
// Solidity: event FactProviderAdded(address indexed factProvider)
func (_FactProviderRegistryContract *FactProviderRegistryContractFilterer) WatchFactProviderAdded(opts *bind.WatchOpts, sink chan<- *FactProviderRegistryContractFactProviderAdded, factProvider []common.Address) (event.Subscription, error) {

	var factProviderRule []interface{}
	for _, factProviderItem := range factProvider {
		factProviderRule = append(factProviderRule, factProviderItem)
	}

	logs, sub, err := _FactProviderRegistryContract.contract.WatchLogs(opts, "FactProviderAdded", factProviderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FactProviderRegistryContractFactProviderAdded)
				if err := _FactProviderRegistryContract.contract.UnpackLog(event, "FactProviderAdded", log); err != nil {
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

// FactProviderRegistryContractFactProviderDeletedIterator is returned from FilterFactProviderDeleted and is used to iterate over the raw logs and unpacked data for FactProviderDeleted events raised by the FactProviderRegistryContract contract.
type FactProviderRegistryContractFactProviderDeletedIterator struct {
	Event *FactProviderRegistryContractFactProviderDeleted // Event containing the contract specifics and raw log

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
func (it *FactProviderRegistryContractFactProviderDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FactProviderRegistryContractFactProviderDeleted)
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
		it.Event = new(FactProviderRegistryContractFactProviderDeleted)
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
func (it *FactProviderRegistryContractFactProviderDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FactProviderRegistryContractFactProviderDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FactProviderRegistryContractFactProviderDeleted represents a FactProviderDeleted event raised by the FactProviderRegistryContract contract.
type FactProviderRegistryContractFactProviderDeleted struct {
	FactProvider common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterFactProviderDeleted is a free log retrieval operation binding the contract event 0xffcf9fb4992cc4ee4022ab3c2c5fa4914c18a25a2dc1614fc52c3ba8bbebf0aa.
//
// Solidity: event FactProviderDeleted(address indexed factProvider)
func (_FactProviderRegistryContract *FactProviderRegistryContractFilterer) FilterFactProviderDeleted(opts *bind.FilterOpts, factProvider []common.Address) (*FactProviderRegistryContractFactProviderDeletedIterator, error) {

	var factProviderRule []interface{}
	for _, factProviderItem := range factProvider {
		factProviderRule = append(factProviderRule, factProviderItem)
	}

	logs, sub, err := _FactProviderRegistryContract.contract.FilterLogs(opts, "FactProviderDeleted", factProviderRule)
	if err != nil {
		return nil, err
	}
	return &FactProviderRegistryContractFactProviderDeletedIterator{contract: _FactProviderRegistryContract.contract, event: "FactProviderDeleted", logs: logs, sub: sub}, nil
}

// WatchFactProviderDeleted is a free log subscription operation binding the contract event 0xffcf9fb4992cc4ee4022ab3c2c5fa4914c18a25a2dc1614fc52c3ba8bbebf0aa.
//
// Solidity: event FactProviderDeleted(address indexed factProvider)
func (_FactProviderRegistryContract *FactProviderRegistryContractFilterer) WatchFactProviderDeleted(opts *bind.WatchOpts, sink chan<- *FactProviderRegistryContractFactProviderDeleted, factProvider []common.Address) (event.Subscription, error) {

	var factProviderRule []interface{}
	for _, factProviderItem := range factProvider {
		factProviderRule = append(factProviderRule, factProviderItem)
	}

	logs, sub, err := _FactProviderRegistryContract.contract.WatchLogs(opts, "FactProviderDeleted", factProviderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FactProviderRegistryContractFactProviderDeleted)
				if err := _FactProviderRegistryContract.contract.UnpackLog(event, "FactProviderDeleted", log); err != nil {
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

// FactProviderRegistryContractFactProviderUpdatedIterator is returned from FilterFactProviderUpdated and is used to iterate over the raw logs and unpacked data for FactProviderUpdated events raised by the FactProviderRegistryContract contract.
type FactProviderRegistryContractFactProviderUpdatedIterator struct {
	Event *FactProviderRegistryContractFactProviderUpdated // Event containing the contract specifics and raw log

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
func (it *FactProviderRegistryContractFactProviderUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FactProviderRegistryContractFactProviderUpdated)
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
		it.Event = new(FactProviderRegistryContractFactProviderUpdated)
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
func (it *FactProviderRegistryContractFactProviderUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FactProviderRegistryContractFactProviderUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FactProviderRegistryContractFactProviderUpdated represents a FactProviderUpdated event raised by the FactProviderRegistryContract contract.
type FactProviderRegistryContractFactProviderUpdated struct {
	FactProvider common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterFactProviderUpdated is a free log retrieval operation binding the contract event 0xf5373d11a64f7dfd39311b28060d0cc68d73cde4dbb253a9cfd6b00f871932a4.
//
// Solidity: event FactProviderUpdated(address indexed factProvider)
func (_FactProviderRegistryContract *FactProviderRegistryContractFilterer) FilterFactProviderUpdated(opts *bind.FilterOpts, factProvider []common.Address) (*FactProviderRegistryContractFactProviderUpdatedIterator, error) {

	var factProviderRule []interface{}
	for _, factProviderItem := range factProvider {
		factProviderRule = append(factProviderRule, factProviderItem)
	}

	logs, sub, err := _FactProviderRegistryContract.contract.FilterLogs(opts, "FactProviderUpdated", factProviderRule)
	if err != nil {
		return nil, err
	}
	return &FactProviderRegistryContractFactProviderUpdatedIterator{contract: _FactProviderRegistryContract.contract, event: "FactProviderUpdated", logs: logs, sub: sub}, nil
}

// WatchFactProviderUpdated is a free log subscription operation binding the contract event 0xf5373d11a64f7dfd39311b28060d0cc68d73cde4dbb253a9cfd6b00f871932a4.
//
// Solidity: event FactProviderUpdated(address indexed factProvider)
func (_FactProviderRegistryContract *FactProviderRegistryContractFilterer) WatchFactProviderUpdated(opts *bind.WatchOpts, sink chan<- *FactProviderRegistryContractFactProviderUpdated, factProvider []common.Address) (event.Subscription, error) {

	var factProviderRule []interface{}
	for _, factProviderItem := range factProvider {
		factProviderRule = append(factProviderRule, factProviderItem)
	}

	logs, sub, err := _FactProviderRegistryContract.contract.WatchLogs(opts, "FactProviderUpdated", factProviderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FactProviderRegistryContractFactProviderUpdated)
				if err := _FactProviderRegistryContract.contract.UnpackLog(event, "FactProviderUpdated", log); err != nil {
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

// FactProviderRegistryContractOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the FactProviderRegistryContract contract.
type FactProviderRegistryContractOwnershipRenouncedIterator struct {
	Event *FactProviderRegistryContractOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *FactProviderRegistryContractOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FactProviderRegistryContractOwnershipRenounced)
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
		it.Event = new(FactProviderRegistryContractOwnershipRenounced)
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
func (it *FactProviderRegistryContractOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FactProviderRegistryContractOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FactProviderRegistryContractOwnershipRenounced represents a OwnershipRenounced event raised by the FactProviderRegistryContract contract.
type FactProviderRegistryContractOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_FactProviderRegistryContract *FactProviderRegistryContractFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*FactProviderRegistryContractOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _FactProviderRegistryContract.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FactProviderRegistryContractOwnershipRenouncedIterator{contract: _FactProviderRegistryContract.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_FactProviderRegistryContract *FactProviderRegistryContractFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *FactProviderRegistryContractOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _FactProviderRegistryContract.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FactProviderRegistryContractOwnershipRenounced)
				if err := _FactProviderRegistryContract.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// FactProviderRegistryContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the FactProviderRegistryContract contract.
type FactProviderRegistryContractOwnershipTransferredIterator struct {
	Event *FactProviderRegistryContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FactProviderRegistryContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FactProviderRegistryContractOwnershipTransferred)
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
		it.Event = new(FactProviderRegistryContractOwnershipTransferred)
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
func (it *FactProviderRegistryContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FactProviderRegistryContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FactProviderRegistryContractOwnershipTransferred represents a OwnershipTransferred event raised by the FactProviderRegistryContract contract.
type FactProviderRegistryContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FactProviderRegistryContract *FactProviderRegistryContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FactProviderRegistryContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FactProviderRegistryContract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FactProviderRegistryContractOwnershipTransferredIterator{contract: _FactProviderRegistryContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FactProviderRegistryContract *FactProviderRegistryContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FactProviderRegistryContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FactProviderRegistryContract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FactProviderRegistryContractOwnershipTransferred)
				if err := _FactProviderRegistryContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
