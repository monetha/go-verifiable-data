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

// PassportLogicRegistryContractABI is the input ABI used to generate the binding from.
const PassportLogicRegistryContractABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"reclaimToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"reclaimEther\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"tokenFallback\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_version\",\"type\":\"string\"},{\"name\":\"_implementation\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"version\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"PassportLogicAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"version\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"CurrentPassportLogicSet\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_version\",\"type\":\"string\"},{\"name\":\"_implementation\",\"type\":\"address\"}],\"name\":\"addPassportLogic\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_version\",\"type\":\"string\"}],\"name\":\"getPassportLogic\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_version\",\"type\":\"string\"}],\"name\":\"setCurrentPassportLogic\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurrentPassportLogicVersion\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurrentPassportLogic\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// PassportLogicRegistryContractBin is the compiled bytecode used for deploying new contracts.
const PassportLogicRegistryContractBin = `0x60806040523480156200001157600080fd5b506040516200135738038062001357833981016040528051602082015160008054600160a060020a0319163317905591019034156200004f57600080fd5b62000064828264010000000062000080810204565b620000788264010000000062000342810204565b505062000642565b600160a060020a03811615156200011e57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f43616e6e6f742073657420696d706c656d656e746174696f6e20746f2061207a60448201527f65726f2061646472657373000000000000000000000000000000000000000000606482015290519081900360840190fd5b6003826040518082805190602001908083835b60208310620001525780518252601f19909201916020918201910162000131565b51815160209384036101000a6000190180199092169116179052920194855250604051938490030190922054600160a060020a03161591506200021e905057604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f43616e6e6f74207265706c616365206578697374696e672076657273696f6e2060448201527f696d706c656d656e746174696f6e000000000000000000000000000000000000606482015290519081900360840190fd5b806003836040518082805190602001908083835b60208310620002535780518252601f19909201916020918201910162000232565b51815160209384036101000a60001901801990921691161790529201948552506040805194859003820185208054600160a060020a031916600160a060020a0397881617905594861684820152848452865194840194909452505083517f7471eb04045ae72adb2fb73deb1e873113901110dd66dbde715232f2495a0cd8928592859290918291606083019186019080838360005b8381101562000302578181015183820152602001620002e8565b50505050905090810190601f168015620003305780820380516001836020036101000a031916815260200191505b50935050505060405180910390a15050565b6003816040518082805190602001908083835b60208310620003765780518252601f19909201916020918201910162000355565b51815160209384036101000a6000190180199092169116179052920194855250604051938490030190922054600160a060020a03161515915062000443905057604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602481018290527f43616e6e6f7420736574206e6f6e2d6578697374696e672070617373706f727460448201527f206c6f6769632061732063757272656e7420696d706c656d656e746174696f6e606482015290519081900360840190fd5b8051620004589060019060208401906200059d565b506003816040518082805190602001908083835b602083106200048d5780518252601f1990920191602091820191016200046c565b518151600019602094850361010090810a820192831692199390931691909117909252949092019687526040805197889003820188205460028054600160a060020a031916600160a060020a03928316178082559091169289018390528189526001805480821615909702909401909516949094049387018490527f4e366bf178b123bb29442cddeceedf743c23cbb40cffa5f577217fe2c54a0b19969195509350918291506060820190859080156200058b5780601f106200055f576101008083540402835291602001916200058b565b820191906000526020600020905b8154815290600101906020018083116200056d57829003601f168201915b5050935050505060405180910390a150565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620005e057805160ff191683800117855562000610565b8280016001018555821562000610579182015b8281111562000610578251825591602001919060010190620005f3565b506200061e92915062000622565b5090565b6200063f91905b808211156200061e576000815560010162000629565b90565b610d0580620006526000396000f3006080604052600436106100ae5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166317ffc32081146100bd578063593efdf1146100e0578063609725ef14610139578063715018a61461016a5780638da5cb5b1461017f5780639a1295d9146101945780639f727c27146101b4578063ba612493146101c9578063c0ee0b8a14610253578063c7c40fbb14610284578063f2fde38b146102e8575b3480156100ba57600080fd5b50005b3480156100c957600080fd5b506100de600160a060020a0360043516610309565b005b3480156100ec57600080fd5b506040805160206004803580820135601f81018490048402850184019095528484526100de9436949293602493928401919081908401838280828437509497506103d59650505050505050565b34801561014557600080fd5b5061014e6103f8565b60408051600160a060020a039092168252519081900360200190f35b34801561017657600080fd5b506100de610408565b34801561018b57600080fd5b5061014e610474565b3480156101a057600080fd5b5061014e6004803560248101910135610483565b3480156101c057600080fd5b506100de6104bc565b3480156101d557600080fd5b506101de61050e565b6040805160208082528351818301528351919283929083019185019080838360005b83811015610218578181015183820152602001610200565b50505050905090810190601f1680156102455780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561025f57600080fd5b506100de60048035600160a060020a03169060248035916044359182019101356105a3565b34801561029057600080fd5b506040805160206004803580820135601f81018490048402850184019095528484526100de94369492936024939284019190819084018382808284375094975050509235600160a060020a031693506105a892505050565b3480156102f457600080fd5b506100de600160a060020a03600435166105c9565b60008054600160a060020a0316331461032157600080fd5b604080517f70a082310000000000000000000000000000000000000000000000000000000081523060048201529051600160a060020a038416916370a082319160248083019260209291908290030181600087803b15801561038257600080fd5b505af1158015610396573d6000803e3d6000fd5b505050506040513d60208110156103ac57600080fd5b50516000549091506103d190600160a060020a0384811691168363ffffffff6105e916565b5050565b600054600160a060020a031633146103ec57600080fd5b6103f5816106a1565b50565b600254600160a060020a03165b90565b600054600160a060020a0316331461041f57600080fd5b60008054604051600160a060020a03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a26000805473ffffffffffffffffffffffffffffffffffffffff19169055565b600054600160a060020a031681565b6000600383836040518083838082843790910194855250506040519283900360200190922054600160a060020a03169250505092915050565b600054600160a060020a031633146104d357600080fd5b60008054604051600160a060020a0390911691303180156108fc02929091818181858888f193505050501580156103f5573d6000803e3d6000fd5b60018054604080516020601f600260001961010087891615020190951694909404938401819004810282018101909252828152606093909290918301828280156105995780601f1061056e57610100808354040283529160200191610599565b820191906000526020600020905b81548152906001019060200180831161057c57829003601f168201915b5050505050905090565b600080fd5b600054600160a060020a031633146105bf57600080fd5b6103d182826108fe565b600054600160a060020a031633146105e057600080fd5b6103f581610bc4565b82600160a060020a031663a9059cbb83836040518363ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018083600160a060020a0316600160a060020a0316815260200182815260200192505050602060405180830381600087803b15801561066557600080fd5b505af1158015610679573d6000803e3d6000fd5b505050506040513d602081101561068f57600080fd5b5051151561069c57600080fd5b505050565b6003816040518082805190602001908083835b602083106106d35780518252601f1990920191602091820191016106b4565b51815160209384036101000a6000190180199092169116179052920194855250604051938490030190922054600160a060020a03161515915061079f905057604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602481018290527f43616e6e6f7420736574206e6f6e2d6578697374696e672070617373706f727460448201527f206c6f6769632061732063757272656e7420696d706c656d656e746174696f6e606482015290519081900360840190fd5b80516107b2906001906020840190610c41565b506003816040518082805190602001908083835b602083106107e55780518252601f1990920191602091820191016107c6565b518151600019602094850361010090810a82019283169219939093169190911790925294909201968752604080519788900382018820546002805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03928316178082559091169289018390528189526001805480821615909702909401909516949094049387018490527f4e366bf178b123bb29442cddeceedf743c23cbb40cffa5f577217fe2c54a0b19969195509350918291506060820190859080156108ec5780601f106108c1576101008083540402835291602001916108ec565b820191906000526020600020905b8154815290600101906020018083116108cf57829003601f168201915b5050935050505060405180910390a150565b600160a060020a038116151561099b57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f43616e6e6f742073657420696d706c656d656e746174696f6e20746f2061207a60448201527f65726f2061646472657373000000000000000000000000000000000000000000606482015290519081900360840190fd5b6003826040518082805190602001908083835b602083106109cd5780518252601f1990920191602091820191016109ae565b51815160209384036101000a6000190180199092169116179052920194855250604051938490030190922054600160a060020a0316159150610a98905057604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f43616e6e6f74207265706c616365206578697374696e672076657273696f6e2060448201527f696d706c656d656e746174696f6e000000000000000000000000000000000000606482015290519081900360840190fd5b806003836040518082805190602001908083835b60208310610acb5780518252601f199092019160209182019101610aac565b51815160209384036101000a6000190180199092169116179052920194855250604080519485900382018520805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0397881617905594861684820152848452865194840194909452505083517f7471eb04045ae72adb2fb73deb1e873113901110dd66dbde715232f2495a0cd8928592859290918291606083019186019080838360005b83811015610b85578181015183820152602001610b6d565b50505050905090810190601f168015610bb25780820380516001836020036101000a031916815260200191505b50935050505060405180910390a15050565b600160a060020a0381161515610bd957600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610c8257805160ff1916838001178555610caf565b82800160010185558215610caf579182015b82811115610caf578251825591602001919060010190610c94565b50610cbb929150610cbf565b5090565b61040591905b80821115610cbb5760008155600101610cc55600a165627a7a7230582030be6f6b69b04014d32f60ed59003aff34c27c83b36b507d48ae67ba320048f50029`

// DeployPassportLogicRegistryContract deploys a new Ethereum contract, binding an instance of PassportLogicRegistryContract to it.
func DeployPassportLogicRegistryContract(auth *bind.TransactOpts, backend bind.ContractBackend, _version string, _implementation common.Address) (common.Address, *types.Transaction, *PassportLogicRegistryContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PassportLogicRegistryContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PassportLogicRegistryContractBin), backend, _version, _implementation)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PassportLogicRegistryContract{PassportLogicRegistryContractCaller: PassportLogicRegistryContractCaller{contract: contract}, PassportLogicRegistryContractTransactor: PassportLogicRegistryContractTransactor{contract: contract}, PassportLogicRegistryContractFilterer: PassportLogicRegistryContractFilterer{contract: contract}}, nil
}

// PassportLogicRegistryContract is an auto generated Go binding around an Ethereum contract.
type PassportLogicRegistryContract struct {
	PassportLogicRegistryContractCaller     // Read-only binding to the contract
	PassportLogicRegistryContractTransactor // Write-only binding to the contract
	PassportLogicRegistryContractFilterer   // Log filterer for contract events
}

// PassportLogicRegistryContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type PassportLogicRegistryContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PassportLogicRegistryContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PassportLogicRegistryContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PassportLogicRegistryContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PassportLogicRegistryContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PassportLogicRegistryContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PassportLogicRegistryContractSession struct {
	Contract     *PassportLogicRegistryContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                  // Call options to use throughout this session
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// PassportLogicRegistryContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PassportLogicRegistryContractCallerSession struct {
	Contract *PassportLogicRegistryContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                        // Call options to use throughout this session
}

// PassportLogicRegistryContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PassportLogicRegistryContractTransactorSession struct {
	Contract     *PassportLogicRegistryContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                        // Transaction auth options to use throughout this session
}

// PassportLogicRegistryContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type PassportLogicRegistryContractRaw struct {
	Contract *PassportLogicRegistryContract // Generic contract binding to access the raw methods on
}

// PassportLogicRegistryContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PassportLogicRegistryContractCallerRaw struct {
	Contract *PassportLogicRegistryContractCaller // Generic read-only contract binding to access the raw methods on
}

// PassportLogicRegistryContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PassportLogicRegistryContractTransactorRaw struct {
	Contract *PassportLogicRegistryContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPassportLogicRegistryContract creates a new instance of PassportLogicRegistryContract, bound to a specific deployed contract.
func NewPassportLogicRegistryContract(address common.Address, backend bind.ContractBackend) (*PassportLogicRegistryContract, error) {
	contract, err := bindPassportLogicRegistryContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PassportLogicRegistryContract{PassportLogicRegistryContractCaller: PassportLogicRegistryContractCaller{contract: contract}, PassportLogicRegistryContractTransactor: PassportLogicRegistryContractTransactor{contract: contract}, PassportLogicRegistryContractFilterer: PassportLogicRegistryContractFilterer{contract: contract}}, nil
}

// NewPassportLogicRegistryContractCaller creates a new read-only instance of PassportLogicRegistryContract, bound to a specific deployed contract.
func NewPassportLogicRegistryContractCaller(address common.Address, caller bind.ContractCaller) (*PassportLogicRegistryContractCaller, error) {
	contract, err := bindPassportLogicRegistryContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PassportLogicRegistryContractCaller{contract: contract}, nil
}

// NewPassportLogicRegistryContractTransactor creates a new write-only instance of PassportLogicRegistryContract, bound to a specific deployed contract.
func NewPassportLogicRegistryContractTransactor(address common.Address, transactor bind.ContractTransactor) (*PassportLogicRegistryContractTransactor, error) {
	contract, err := bindPassportLogicRegistryContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PassportLogicRegistryContractTransactor{contract: contract}, nil
}

// NewPassportLogicRegistryContractFilterer creates a new log filterer instance of PassportLogicRegistryContract, bound to a specific deployed contract.
func NewPassportLogicRegistryContractFilterer(address common.Address, filterer bind.ContractFilterer) (*PassportLogicRegistryContractFilterer, error) {
	contract, err := bindPassportLogicRegistryContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PassportLogicRegistryContractFilterer{contract: contract}, nil
}

// bindPassportLogicRegistryContract binds a generic wrapper to an already deployed contract.
func bindPassportLogicRegistryContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PassportLogicRegistryContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PassportLogicRegistryContract *PassportLogicRegistryContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PassportLogicRegistryContract.Contract.PassportLogicRegistryContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PassportLogicRegistryContract *PassportLogicRegistryContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PassportLogicRegistryContract.Contract.PassportLogicRegistryContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PassportLogicRegistryContract *PassportLogicRegistryContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PassportLogicRegistryContract.Contract.PassportLogicRegistryContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PassportLogicRegistryContract *PassportLogicRegistryContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PassportLogicRegistryContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PassportLogicRegistryContract *PassportLogicRegistryContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PassportLogicRegistryContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PassportLogicRegistryContract *PassportLogicRegistryContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PassportLogicRegistryContract.Contract.contract.Transact(opts, method, params...)
}

// GetCurrentPassportLogic is a free data retrieval call binding the contract method 0x609725ef.
//
// Solidity: function getCurrentPassportLogic() constant returns(address)
func (_PassportLogicRegistryContract *PassportLogicRegistryContractCaller) GetCurrentPassportLogic(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PassportLogicRegistryContract.contract.Call(opts, out, "getCurrentPassportLogic")
	return *ret0, err
}

// GetCurrentPassportLogic is a free data retrieval call binding the contract method 0x609725ef.
//
// Solidity: function getCurrentPassportLogic() constant returns(address)
func (_PassportLogicRegistryContract *PassportLogicRegistryContractSession) GetCurrentPassportLogic() (common.Address, error) {
	return _PassportLogicRegistryContract.Contract.GetCurrentPassportLogic(&_PassportLogicRegistryContract.CallOpts)
}

// GetCurrentPassportLogic is a free data retrieval call binding the contract method 0x609725ef.
//
// Solidity: function getCurrentPassportLogic() constant returns(address)
func (_PassportLogicRegistryContract *PassportLogicRegistryContractCallerSession) GetCurrentPassportLogic() (common.Address, error) {
	return _PassportLogicRegistryContract.Contract.GetCurrentPassportLogic(&_PassportLogicRegistryContract.CallOpts)
}

// GetCurrentPassportLogicVersion is a free data retrieval call binding the contract method 0xba612493.
//
// Solidity: function getCurrentPassportLogicVersion() constant returns(string)
func (_PassportLogicRegistryContract *PassportLogicRegistryContractCaller) GetCurrentPassportLogicVersion(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _PassportLogicRegistryContract.contract.Call(opts, out, "getCurrentPassportLogicVersion")
	return *ret0, err
}

// GetCurrentPassportLogicVersion is a free data retrieval call binding the contract method 0xba612493.
//
// Solidity: function getCurrentPassportLogicVersion() constant returns(string)
func (_PassportLogicRegistryContract *PassportLogicRegistryContractSession) GetCurrentPassportLogicVersion() (string, error) {
	return _PassportLogicRegistryContract.Contract.GetCurrentPassportLogicVersion(&_PassportLogicRegistryContract.CallOpts)
}

// GetCurrentPassportLogicVersion is a free data retrieval call binding the contract method 0xba612493.
//
// Solidity: function getCurrentPassportLogicVersion() constant returns(string)
func (_PassportLogicRegistryContract *PassportLogicRegistryContractCallerSession) GetCurrentPassportLogicVersion() (string, error) {
	return _PassportLogicRegistryContract.Contract.GetCurrentPassportLogicVersion(&_PassportLogicRegistryContract.CallOpts)
}

// GetPassportLogic is a free data retrieval call binding the contract method 0x9a1295d9.
//
// Solidity: function getPassportLogic(string _version) constant returns(address)
func (_PassportLogicRegistryContract *PassportLogicRegistryContractCaller) GetPassportLogic(opts *bind.CallOpts, _version string) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PassportLogicRegistryContract.contract.Call(opts, out, "getPassportLogic", _version)
	return *ret0, err
}

// GetPassportLogic is a free data retrieval call binding the contract method 0x9a1295d9.
//
// Solidity: function getPassportLogic(string _version) constant returns(address)
func (_PassportLogicRegistryContract *PassportLogicRegistryContractSession) GetPassportLogic(_version string) (common.Address, error) {
	return _PassportLogicRegistryContract.Contract.GetPassportLogic(&_PassportLogicRegistryContract.CallOpts, _version)
}

// GetPassportLogic is a free data retrieval call binding the contract method 0x9a1295d9.
//
// Solidity: function getPassportLogic(string _version) constant returns(address)
func (_PassportLogicRegistryContract *PassportLogicRegistryContractCallerSession) GetPassportLogic(_version string) (common.Address, error) {
	return _PassportLogicRegistryContract.Contract.GetPassportLogic(&_PassportLogicRegistryContract.CallOpts, _version)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PassportLogicRegistryContract *PassportLogicRegistryContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PassportLogicRegistryContract.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PassportLogicRegistryContract *PassportLogicRegistryContractSession) Owner() (common.Address, error) {
	return _PassportLogicRegistryContract.Contract.Owner(&_PassportLogicRegistryContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PassportLogicRegistryContract *PassportLogicRegistryContractCallerSession) Owner() (common.Address, error) {
	return _PassportLogicRegistryContract.Contract.Owner(&_PassportLogicRegistryContract.CallOpts)
}

// TokenFallback is a free data retrieval call binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(address _from, uint256 _value, bytes _data) constant returns()
func (_PassportLogicRegistryContract *PassportLogicRegistryContractCaller) TokenFallback(opts *bind.CallOpts, _from common.Address, _value *big.Int, _data []byte) error {
	var ()
	out := &[]interface{}{}
	err := _PassportLogicRegistryContract.contract.Call(opts, out, "tokenFallback", _from, _value, _data)
	return err
}

// TokenFallback is a free data retrieval call binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(address _from, uint256 _value, bytes _data) constant returns()
func (_PassportLogicRegistryContract *PassportLogicRegistryContractSession) TokenFallback(_from common.Address, _value *big.Int, _data []byte) error {
	return _PassportLogicRegistryContract.Contract.TokenFallback(&_PassportLogicRegistryContract.CallOpts, _from, _value, _data)
}

// TokenFallback is a free data retrieval call binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(address _from, uint256 _value, bytes _data) constant returns()
func (_PassportLogicRegistryContract *PassportLogicRegistryContractCallerSession) TokenFallback(_from common.Address, _value *big.Int, _data []byte) error {
	return _PassportLogicRegistryContract.Contract.TokenFallback(&_PassportLogicRegistryContract.CallOpts, _from, _value, _data)
}

// AddPassportLogic is a paid mutator transaction binding the contract method 0xc7c40fbb.
//
// Solidity: function addPassportLogic(string _version, address _implementation) returns()
func (_PassportLogicRegistryContract *PassportLogicRegistryContractTransactor) AddPassportLogic(opts *bind.TransactOpts, _version string, _implementation common.Address) (*types.Transaction, error) {
	return _PassportLogicRegistryContract.contract.Transact(opts, "addPassportLogic", _version, _implementation)
}

// AddPassportLogic is a paid mutator transaction binding the contract method 0xc7c40fbb.
//
// Solidity: function addPassportLogic(string _version, address _implementation) returns()
func (_PassportLogicRegistryContract *PassportLogicRegistryContractSession) AddPassportLogic(_version string, _implementation common.Address) (*types.Transaction, error) {
	return _PassportLogicRegistryContract.Contract.AddPassportLogic(&_PassportLogicRegistryContract.TransactOpts, _version, _implementation)
}

// AddPassportLogic is a paid mutator transaction binding the contract method 0xc7c40fbb.
//
// Solidity: function addPassportLogic(string _version, address _implementation) returns()
func (_PassportLogicRegistryContract *PassportLogicRegistryContractTransactorSession) AddPassportLogic(_version string, _implementation common.Address) (*types.Transaction, error) {
	return _PassportLogicRegistryContract.Contract.AddPassportLogic(&_PassportLogicRegistryContract.TransactOpts, _version, _implementation)
}

// ReclaimEther is a paid mutator transaction binding the contract method 0x9f727c27.
//
// Solidity: function reclaimEther() returns()
func (_PassportLogicRegistryContract *PassportLogicRegistryContractTransactor) ReclaimEther(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PassportLogicRegistryContract.contract.Transact(opts, "reclaimEther")
}

// ReclaimEther is a paid mutator transaction binding the contract method 0x9f727c27.
//
// Solidity: function reclaimEther() returns()
func (_PassportLogicRegistryContract *PassportLogicRegistryContractSession) ReclaimEther() (*types.Transaction, error) {
	return _PassportLogicRegistryContract.Contract.ReclaimEther(&_PassportLogicRegistryContract.TransactOpts)
}

// ReclaimEther is a paid mutator transaction binding the contract method 0x9f727c27.
//
// Solidity: function reclaimEther() returns()
func (_PassportLogicRegistryContract *PassportLogicRegistryContractTransactorSession) ReclaimEther() (*types.Transaction, error) {
	return _PassportLogicRegistryContract.Contract.ReclaimEther(&_PassportLogicRegistryContract.TransactOpts)
}

// ReclaimToken is a paid mutator transaction binding the contract method 0x17ffc320.
//
// Solidity: function reclaimToken(address _token) returns()
func (_PassportLogicRegistryContract *PassportLogicRegistryContractTransactor) ReclaimToken(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _PassportLogicRegistryContract.contract.Transact(opts, "reclaimToken", _token)
}

// ReclaimToken is a paid mutator transaction binding the contract method 0x17ffc320.
//
// Solidity: function reclaimToken(address _token) returns()
func (_PassportLogicRegistryContract *PassportLogicRegistryContractSession) ReclaimToken(_token common.Address) (*types.Transaction, error) {
	return _PassportLogicRegistryContract.Contract.ReclaimToken(&_PassportLogicRegistryContract.TransactOpts, _token)
}

// ReclaimToken is a paid mutator transaction binding the contract method 0x17ffc320.
//
// Solidity: function reclaimToken(address _token) returns()
func (_PassportLogicRegistryContract *PassportLogicRegistryContractTransactorSession) ReclaimToken(_token common.Address) (*types.Transaction, error) {
	return _PassportLogicRegistryContract.Contract.ReclaimToken(&_PassportLogicRegistryContract.TransactOpts, _token)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PassportLogicRegistryContract *PassportLogicRegistryContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PassportLogicRegistryContract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PassportLogicRegistryContract *PassportLogicRegistryContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _PassportLogicRegistryContract.Contract.RenounceOwnership(&_PassportLogicRegistryContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PassportLogicRegistryContract *PassportLogicRegistryContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PassportLogicRegistryContract.Contract.RenounceOwnership(&_PassportLogicRegistryContract.TransactOpts)
}

// SetCurrentPassportLogic is a paid mutator transaction binding the contract method 0x593efdf1.
//
// Solidity: function setCurrentPassportLogic(string _version) returns()
func (_PassportLogicRegistryContract *PassportLogicRegistryContractTransactor) SetCurrentPassportLogic(opts *bind.TransactOpts, _version string) (*types.Transaction, error) {
	return _PassportLogicRegistryContract.contract.Transact(opts, "setCurrentPassportLogic", _version)
}

// SetCurrentPassportLogic is a paid mutator transaction binding the contract method 0x593efdf1.
//
// Solidity: function setCurrentPassportLogic(string _version) returns()
func (_PassportLogicRegistryContract *PassportLogicRegistryContractSession) SetCurrentPassportLogic(_version string) (*types.Transaction, error) {
	return _PassportLogicRegistryContract.Contract.SetCurrentPassportLogic(&_PassportLogicRegistryContract.TransactOpts, _version)
}

// SetCurrentPassportLogic is a paid mutator transaction binding the contract method 0x593efdf1.
//
// Solidity: function setCurrentPassportLogic(string _version) returns()
func (_PassportLogicRegistryContract *PassportLogicRegistryContractTransactorSession) SetCurrentPassportLogic(_version string) (*types.Transaction, error) {
	return _PassportLogicRegistryContract.Contract.SetCurrentPassportLogic(&_PassportLogicRegistryContract.TransactOpts, _version)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_PassportLogicRegistryContract *PassportLogicRegistryContractTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _PassportLogicRegistryContract.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_PassportLogicRegistryContract *PassportLogicRegistryContractSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _PassportLogicRegistryContract.Contract.TransferOwnership(&_PassportLogicRegistryContract.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_PassportLogicRegistryContract *PassportLogicRegistryContractTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _PassportLogicRegistryContract.Contract.TransferOwnership(&_PassportLogicRegistryContract.TransactOpts, _newOwner)
}

// PassportLogicRegistryContractCurrentPassportLogicSetIterator is returned from FilterCurrentPassportLogicSet and is used to iterate over the raw logs and unpacked data for CurrentPassportLogicSet events raised by the PassportLogicRegistryContract contract.
type PassportLogicRegistryContractCurrentPassportLogicSetIterator struct {
	Event *PassportLogicRegistryContractCurrentPassportLogicSet // Event containing the contract specifics and raw log

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
func (it *PassportLogicRegistryContractCurrentPassportLogicSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PassportLogicRegistryContractCurrentPassportLogicSet)
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
		it.Event = new(PassportLogicRegistryContractCurrentPassportLogicSet)
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
func (it *PassportLogicRegistryContractCurrentPassportLogicSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PassportLogicRegistryContractCurrentPassportLogicSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PassportLogicRegistryContractCurrentPassportLogicSet represents a CurrentPassportLogicSet event raised by the PassportLogicRegistryContract contract.
type PassportLogicRegistryContractCurrentPassportLogicSet struct {
	Version        string
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterCurrentPassportLogicSet is a free log retrieval operation binding the contract event 0x4e366bf178b123bb29442cddeceedf743c23cbb40cffa5f577217fe2c54a0b19.
//
// Solidity: event CurrentPassportLogicSet(string version, address implementation)
func (_PassportLogicRegistryContract *PassportLogicRegistryContractFilterer) FilterCurrentPassportLogicSet(opts *bind.FilterOpts) (*PassportLogicRegistryContractCurrentPassportLogicSetIterator, error) {

	logs, sub, err := _PassportLogicRegistryContract.contract.FilterLogs(opts, "CurrentPassportLogicSet")
	if err != nil {
		return nil, err
	}
	return &PassportLogicRegistryContractCurrentPassportLogicSetIterator{contract: _PassportLogicRegistryContract.contract, event: "CurrentPassportLogicSet", logs: logs, sub: sub}, nil
}

// WatchCurrentPassportLogicSet is a free log subscription operation binding the contract event 0x4e366bf178b123bb29442cddeceedf743c23cbb40cffa5f577217fe2c54a0b19.
//
// Solidity: event CurrentPassportLogicSet(string version, address implementation)
func (_PassportLogicRegistryContract *PassportLogicRegistryContractFilterer) WatchCurrentPassportLogicSet(opts *bind.WatchOpts, sink chan<- *PassportLogicRegistryContractCurrentPassportLogicSet) (event.Subscription, error) {

	logs, sub, err := _PassportLogicRegistryContract.contract.WatchLogs(opts, "CurrentPassportLogicSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PassportLogicRegistryContractCurrentPassportLogicSet)
				if err := _PassportLogicRegistryContract.contract.UnpackLog(event, "CurrentPassportLogicSet", log); err != nil {
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

// PassportLogicRegistryContractOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the PassportLogicRegistryContract contract.
type PassportLogicRegistryContractOwnershipRenouncedIterator struct {
	Event *PassportLogicRegistryContractOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *PassportLogicRegistryContractOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PassportLogicRegistryContractOwnershipRenounced)
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
		it.Event = new(PassportLogicRegistryContractOwnershipRenounced)
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
func (it *PassportLogicRegistryContractOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PassportLogicRegistryContractOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PassportLogicRegistryContractOwnershipRenounced represents a OwnershipRenounced event raised by the PassportLogicRegistryContract contract.
type PassportLogicRegistryContractOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_PassportLogicRegistryContract *PassportLogicRegistryContractFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*PassportLogicRegistryContractOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _PassportLogicRegistryContract.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PassportLogicRegistryContractOwnershipRenouncedIterator{contract: _PassportLogicRegistryContract.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_PassportLogicRegistryContract *PassportLogicRegistryContractFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *PassportLogicRegistryContractOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _PassportLogicRegistryContract.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PassportLogicRegistryContractOwnershipRenounced)
				if err := _PassportLogicRegistryContract.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// PassportLogicRegistryContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PassportLogicRegistryContract contract.
type PassportLogicRegistryContractOwnershipTransferredIterator struct {
	Event *PassportLogicRegistryContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PassportLogicRegistryContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PassportLogicRegistryContractOwnershipTransferred)
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
		it.Event = new(PassportLogicRegistryContractOwnershipTransferred)
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
func (it *PassportLogicRegistryContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PassportLogicRegistryContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PassportLogicRegistryContractOwnershipTransferred represents a OwnershipTransferred event raised by the PassportLogicRegistryContract contract.
type PassportLogicRegistryContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PassportLogicRegistryContract *PassportLogicRegistryContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PassportLogicRegistryContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PassportLogicRegistryContract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PassportLogicRegistryContractOwnershipTransferredIterator{contract: _PassportLogicRegistryContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PassportLogicRegistryContract *PassportLogicRegistryContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PassportLogicRegistryContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PassportLogicRegistryContract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PassportLogicRegistryContractOwnershipTransferred)
				if err := _PassportLogicRegistryContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// PassportLogicRegistryContractPassportLogicAddedIterator is returned from FilterPassportLogicAdded and is used to iterate over the raw logs and unpacked data for PassportLogicAdded events raised by the PassportLogicRegistryContract contract.
type PassportLogicRegistryContractPassportLogicAddedIterator struct {
	Event *PassportLogicRegistryContractPassportLogicAdded // Event containing the contract specifics and raw log

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
func (it *PassportLogicRegistryContractPassportLogicAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PassportLogicRegistryContractPassportLogicAdded)
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
		it.Event = new(PassportLogicRegistryContractPassportLogicAdded)
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
func (it *PassportLogicRegistryContractPassportLogicAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PassportLogicRegistryContractPassportLogicAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PassportLogicRegistryContractPassportLogicAdded represents a PassportLogicAdded event raised by the PassportLogicRegistryContract contract.
type PassportLogicRegistryContractPassportLogicAdded struct {
	Version        string
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterPassportLogicAdded is a free log retrieval operation binding the contract event 0x7471eb04045ae72adb2fb73deb1e873113901110dd66dbde715232f2495a0cd8.
//
// Solidity: event PassportLogicAdded(string version, address implementation)
func (_PassportLogicRegistryContract *PassportLogicRegistryContractFilterer) FilterPassportLogicAdded(opts *bind.FilterOpts) (*PassportLogicRegistryContractPassportLogicAddedIterator, error) {

	logs, sub, err := _PassportLogicRegistryContract.contract.FilterLogs(opts, "PassportLogicAdded")
	if err != nil {
		return nil, err
	}
	return &PassportLogicRegistryContractPassportLogicAddedIterator{contract: _PassportLogicRegistryContract.contract, event: "PassportLogicAdded", logs: logs, sub: sub}, nil
}

// WatchPassportLogicAdded is a free log subscription operation binding the contract event 0x7471eb04045ae72adb2fb73deb1e873113901110dd66dbde715232f2495a0cd8.
//
// Solidity: event PassportLogicAdded(string version, address implementation)
func (_PassportLogicRegistryContract *PassportLogicRegistryContractFilterer) WatchPassportLogicAdded(opts *bind.WatchOpts, sink chan<- *PassportLogicRegistryContractPassportLogicAdded) (event.Subscription, error) {

	logs, sub, err := _PassportLogicRegistryContract.contract.WatchLogs(opts, "PassportLogicAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PassportLogicRegistryContractPassportLogicAdded)
				if err := _PassportLogicRegistryContract.contract.UnpackLog(event, "PassportLogicAdded", log); err != nil {
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
