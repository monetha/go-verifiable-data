// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"math/big"
	"strings"
)

// PassportFactoryContractABI is the input ABI used to generate the binding from.
const PassportFactoryContractABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"reclaimToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"reclaimEther\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"tokenFallback\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_registry\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"passport\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"PassportCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"oldRegistry\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newRegistry\",\"type\":\"address\"}],\"name\":\"PassportLogicRegistryChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_registry\",\"type\":\"address\"}],\"name\":\"setRegistry\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRegistry\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"createPassport\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// PassportFactoryContractBin is the compiled bytecode used for deploying new contracts.
const PassportFactoryContractBin = `0x608060405234801561001057600080fd5b50604051602080610fa4833981016040525160008054600160a060020a03191633179055341561003f57600080fd5b61005181640100000000610057810204565b5061008e565b600160a060020a038116151561006c57600080fd5b60018054600160a060020a031916600160a060020a0392909216919091179055565b610f078061009d6000396000f3006080604052600436106100985763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166317ffc32081146100a75780632ec0faad146100ca5780635ab1bd53146100fb578063715018a6146101105780638da5cb5b146101255780639f727c271461013a578063a91ee0dc1461014f578063c0ee0b8a14610170578063f2fde38b146101a1575b3480156100a457600080fd5b50005b3480156100b357600080fd5b506100c8600160a060020a03600435166101c2565b005b3480156100d657600080fd5b506100df61028e565b60408051600160a060020a039092168252519081900360200190f35b34801561010757600080fd5b506100df61038c565b34801561011c57600080fd5b506100c861039b565b34801561013157600080fd5b506100df610407565b34801561014657600080fd5b506100c8610416565b34801561015b57600080fd5b506100c8600160a060020a036004351661046b565b34801561017c57600080fd5b506100c860048035600160a060020a03169060248035916044359182019101356104c5565b3480156101ad57600080fd5b506100c8600160a060020a03600435166104ca565b60008054600160a060020a031633146101da57600080fd5b604080517f70a082310000000000000000000000000000000000000000000000000000000081523060048201529051600160a060020a038416916370a082319160248083019260209291908290030181600087803b15801561023b57600080fd5b505af115801561024f573d6000803e3d6000fd5b505050506040513d602081101561026557600080fd5b505160005490915061028a90600160a060020a0384811691168363ffffffff6104ea16565b5050565b6001546000908190600160a060020a03166102a7610663565b600160a060020a03909116815260405190819003602001906000f0801580156102d4573d6000803e3d6000fd5b50604080517ff2fde38b0000000000000000000000000000000000000000000000000000000081523360048201529051919250600160a060020a0383169163f2fde38b9160248082019260009290919082900301818387803b15801561033957600080fd5b505af115801561034d573d6000803e3d6000fd5b5050604051339250600160a060020a03841691507f03f096f07a4d27c54645fa682640d44179c821150e16c48b27130ca928fa937c90600090a3919050565b600154600160a060020a031690565b600054600160a060020a031633146103b257600080fd5b60008054604051600160a060020a03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a26000805473ffffffffffffffffffffffffffffffffffffffff19169055565b600054600160a060020a031681565b600054600160a060020a0316331461042d57600080fd5b60008054604051600160a060020a0390911691303180156108fc02929091818181858888f19350505050158015610468573d6000803e3d6000fd5b50565b600054600160a060020a0316331461048257600080fd5b600154604051600160a060020a038084169216907f5c2abfd67230c0e47d6de28402bfe206c7a57283cba891416ed657fd70a714c290600090a3610468816105a2565b600080fd5b600054600160a060020a031633146104e157600080fd5b610468816105e6565b82600160a060020a031663a9059cbb83836040518363ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018083600160a060020a0316600160a060020a0316815260200182815260200192505050602060405180830381600087803b15801561056657600080fd5b505af115801561057a573d6000803e3d6000fd5b505050506040513d602081101561059057600080fd5b5051151561059d57600080fd5b505050565b600160a060020a03811615156105b757600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600160a060020a03811615156105fb57600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60405161086880610674833901905600608060405234801561001057600080fd5b5060405160208061086883398101604081815291517f6f72672e6d6f6e657468612e70726f78792e6f776e6572000000000000000000825291519081900360170190206000805160206108288339815191521461006957fe5b61007b3364010000000061015b810204565b604080517f6f72672e6d6f6e657468612e70726f78792e70656e64696e674f776e657200008152905190819003601e0190207fcfd0c6ea5352192d7d4c5d4e7a73c5da12c871730cb60ff57879cbe7b403bb52146100d557fe5b604080517f6f72672e6d6f6e657468612e70617373706f72742e70726f78792e726567697381527f7472790000000000000000000000000000000000000000000000000000000000602082015290519081900360230190206000805160206108488339815191521461014357fe5b6101558164010000000061016d810204565b5061021f565b60008051602061082883398151915255565b6000600160a060020a038216151561020c57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f43616e6e6f742073657420726567697374727920746f2061207a65726f20616460448201527f6472657373000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b5060008051602061084883398151915255565b6105fa8061022e6000396000f3006080604052600436106100985763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416634645a41881146100aa5780634e71e0c8146100cb578063715018a6146100e057806383197ef0146100f557806386d5c5f91461010a5780638da5cb5b1461013b578063e30c397814610150578063f2fde38b14610165578063f5074f4114610186575b6100a86100a36101a7565b610238565b005b3480156100b657600080fd5b506100a8600160a060020a036004351661025c565b3480156100d757600080fd5b506100a86102cb565b3480156100ec57600080fd5b506100a8610351565b34801561010157600080fd5b506100a86103b4565b34801561011657600080fd5b5061011f6103e3565b60408051600160a060020a039092168252519081900360200190f35b34801561014757600080fd5b5061011f6103f2565b34801561015c57600080fd5b5061011f6103fc565b34801561017157600080fd5b506100a8600160a060020a0360043516610406565b34801561019257600080fd5b506100a8600160a060020a036004351661042b565b60006101b1610453565b600160a060020a031663609725ef6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b15801561020757600080fd5b505af115801561021b573d6000803e3d6000fd5b505050506040513d602081101561023157600080fd5b5051905090565b3660008037600080366000845af43d6000803e808015610257573d6000f35b3d6000fd5b610264610478565b600160a060020a0316331461027857600080fd5b80600160a060020a031661028a610453565b600160a060020a03167f5c2abfd67230c0e47d6de28402bfe206c7a57283cba891416ed657fd70a714c260405160405180910390a36102c88161049d565b50565b6102d3610561565b600160a060020a031633146102e757600080fd5b6102ef610561565b600160a060020a0316610300610478565b600160a060020a03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3610345610340610561565b610586565b61034f60006105aa565b565b610359610478565b600160a060020a0316331461036d57600080fd5b610375610478565b600160a060020a03167ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482060405160405180910390a261034f6000610586565b6103bc610478565b600160a060020a031633146103d057600080fd5b6103d8610478565b600160a060020a0316ff5b60006103ed610453565b905090565b60006103ed610478565b60006103ed610561565b61040e610478565b600160a060020a0316331461042257600080fd5b6102c8816105aa565b610433610478565b600160a060020a0316331461044757600080fd5b80600160a060020a0316ff5b7fa04bab69e45aeb4c94a78ba5bc1be67ef28977c4fdf815a30b829a794eb67a4a5490565b7f3ca57e4b51fc2e18497b219410298879868edada7e6fe5132c8feceb0a080d225490565b6000600160a060020a038216151561053c57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f43616e6e6f742073657420726567697374727920746f2061207a65726f20616460448201527f6472657373000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b507fa04bab69e45aeb4c94a78ba5bc1be67ef28977c4fdf815a30b829a794eb67a4a55565b7fcfd0c6ea5352192d7d4c5d4e7a73c5da12c871730cb60ff57879cbe7b403bb525490565b7f3ca57e4b51fc2e18497b219410298879868edada7e6fe5132c8feceb0a080d2255565b7fcfd0c6ea5352192d7d4c5d4e7a73c5da12c871730cb60ff57879cbe7b403bb52555600a165627a7a7230582027bb57ac09535e72c3e85d0be1e55abee8f4ad72ffc717aabb25edc02408067400293ca57e4b51fc2e18497b219410298879868edada7e6fe5132c8feceb0a080d22a04bab69e45aeb4c94a78ba5bc1be67ef28977c4fdf815a30b829a794eb67a4aa165627a7a72305820c264c61a30abe464547d47bb79f282656e4638921c325f45dd01f08d380388160029`

// DeployPassportFactoryContract deploys a new Ethereum contract, binding an instance of PassportFactoryContract to it.
func DeployPassportFactoryContract(auth *bind.TransactOpts, backend bind.ContractBackend, _registry common.Address) (common.Address, *types.Transaction, *PassportFactoryContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PassportFactoryContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PassportFactoryContractBin), backend, _registry)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PassportFactoryContract{PassportFactoryContractCaller: PassportFactoryContractCaller{contract: contract}, PassportFactoryContractTransactor: PassportFactoryContractTransactor{contract: contract}, PassportFactoryContractFilterer: PassportFactoryContractFilterer{contract: contract}}, nil
}

// PassportFactoryContract is an auto generated Go binding around an Ethereum contract.
type PassportFactoryContract struct {
	PassportFactoryContractCaller     // Read-only binding to the contract
	PassportFactoryContractTransactor // Write-only binding to the contract
	PassportFactoryContractFilterer   // Log filterer for contract events
}

// PassportFactoryContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type PassportFactoryContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PassportFactoryContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PassportFactoryContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PassportFactoryContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PassportFactoryContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PassportFactoryContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PassportFactoryContractSession struct {
	Contract     *PassportFactoryContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// PassportFactoryContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PassportFactoryContractCallerSession struct {
	Contract *PassportFactoryContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// PassportFactoryContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PassportFactoryContractTransactorSession struct {
	Contract     *PassportFactoryContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// PassportFactoryContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type PassportFactoryContractRaw struct {
	Contract *PassportFactoryContract // Generic contract binding to access the raw methods on
}

// PassportFactoryContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PassportFactoryContractCallerRaw struct {
	Contract *PassportFactoryContractCaller // Generic read-only contract binding to access the raw methods on
}

// PassportFactoryContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PassportFactoryContractTransactorRaw struct {
	Contract *PassportFactoryContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPassportFactoryContract creates a new instance of PassportFactoryContract, bound to a specific deployed contract.
func NewPassportFactoryContract(address common.Address, backend bind.ContractBackend) (*PassportFactoryContract, error) {
	contract, err := bindPassportFactoryContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PassportFactoryContract{PassportFactoryContractCaller: PassportFactoryContractCaller{contract: contract}, PassportFactoryContractTransactor: PassportFactoryContractTransactor{contract: contract}, PassportFactoryContractFilterer: PassportFactoryContractFilterer{contract: contract}}, nil
}

// NewPassportFactoryContractCaller creates a new read-only instance of PassportFactoryContract, bound to a specific deployed contract.
func NewPassportFactoryContractCaller(address common.Address, caller bind.ContractCaller) (*PassportFactoryContractCaller, error) {
	contract, err := bindPassportFactoryContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PassportFactoryContractCaller{contract: contract}, nil
}

// NewPassportFactoryContractTransactor creates a new write-only instance of PassportFactoryContract, bound to a specific deployed contract.
func NewPassportFactoryContractTransactor(address common.Address, transactor bind.ContractTransactor) (*PassportFactoryContractTransactor, error) {
	contract, err := bindPassportFactoryContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PassportFactoryContractTransactor{contract: contract}, nil
}

// NewPassportFactoryContractFilterer creates a new log filterer instance of PassportFactoryContract, bound to a specific deployed contract.
func NewPassportFactoryContractFilterer(address common.Address, filterer bind.ContractFilterer) (*PassportFactoryContractFilterer, error) {
	contract, err := bindPassportFactoryContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PassportFactoryContractFilterer{contract: contract}, nil
}

// bindPassportFactoryContract binds a generic wrapper to an already deployed contract.
func bindPassportFactoryContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PassportFactoryContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PassportFactoryContract *PassportFactoryContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PassportFactoryContract.Contract.PassportFactoryContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PassportFactoryContract *PassportFactoryContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PassportFactoryContract.Contract.PassportFactoryContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PassportFactoryContract *PassportFactoryContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PassportFactoryContract.Contract.PassportFactoryContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PassportFactoryContract *PassportFactoryContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PassportFactoryContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PassportFactoryContract *PassportFactoryContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PassportFactoryContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PassportFactoryContract *PassportFactoryContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PassportFactoryContract.Contract.contract.Transact(opts, method, params...)
}

// GetRegistry is a free data retrieval call binding the contract method 0x5ab1bd53.
//
// Solidity: function getRegistry() constant returns(address)
func (_PassportFactoryContract *PassportFactoryContractCaller) GetRegistry(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PassportFactoryContract.contract.Call(opts, out, "getRegistry")
	return *ret0, err
}

// GetRegistry is a free data retrieval call binding the contract method 0x5ab1bd53.
//
// Solidity: function getRegistry() constant returns(address)
func (_PassportFactoryContract *PassportFactoryContractSession) GetRegistry() (common.Address, error) {
	return _PassportFactoryContract.Contract.GetRegistry(&_PassportFactoryContract.CallOpts)
}

// GetRegistry is a free data retrieval call binding the contract method 0x5ab1bd53.
//
// Solidity: function getRegistry() constant returns(address)
func (_PassportFactoryContract *PassportFactoryContractCallerSession) GetRegistry() (common.Address, error) {
	return _PassportFactoryContract.Contract.GetRegistry(&_PassportFactoryContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PassportFactoryContract *PassportFactoryContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PassportFactoryContract.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PassportFactoryContract *PassportFactoryContractSession) Owner() (common.Address, error) {
	return _PassportFactoryContract.Contract.Owner(&_PassportFactoryContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PassportFactoryContract *PassportFactoryContractCallerSession) Owner() (common.Address, error) {
	return _PassportFactoryContract.Contract.Owner(&_PassportFactoryContract.CallOpts)
}

// TokenFallback is a free data retrieval call binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(_from address, _value uint256, _data bytes) constant returns()
func (_PassportFactoryContract *PassportFactoryContractCaller) TokenFallback(opts *bind.CallOpts, _from common.Address, _value *big.Int, _data []byte) error {
	var ()
	out := &[]interface{}{}
	err := _PassportFactoryContract.contract.Call(opts, out, "tokenFallback", _from, _value, _data)
	return err
}

// TokenFallback is a free data retrieval call binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(_from address, _value uint256, _data bytes) constant returns()
func (_PassportFactoryContract *PassportFactoryContractSession) TokenFallback(_from common.Address, _value *big.Int, _data []byte) error {
	return _PassportFactoryContract.Contract.TokenFallback(&_PassportFactoryContract.CallOpts, _from, _value, _data)
}

// TokenFallback is a free data retrieval call binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(_from address, _value uint256, _data bytes) constant returns()
func (_PassportFactoryContract *PassportFactoryContractCallerSession) TokenFallback(_from common.Address, _value *big.Int, _data []byte) error {
	return _PassportFactoryContract.Contract.TokenFallback(&_PassportFactoryContract.CallOpts, _from, _value, _data)
}

// CreatePassport is a paid mutator transaction binding the contract method 0x2ec0faad.
//
// Solidity: function createPassport() returns(address)
func (_PassportFactoryContract *PassportFactoryContractTransactor) CreatePassport(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PassportFactoryContract.contract.Transact(opts, "createPassport")
}

// CreatePassport is a paid mutator transaction binding the contract method 0x2ec0faad.
//
// Solidity: function createPassport() returns(address)
func (_PassportFactoryContract *PassportFactoryContractSession) CreatePassport() (*types.Transaction, error) {
	return _PassportFactoryContract.Contract.CreatePassport(&_PassportFactoryContract.TransactOpts)
}

// CreatePassport is a paid mutator transaction binding the contract method 0x2ec0faad.
//
// Solidity: function createPassport() returns(address)
func (_PassportFactoryContract *PassportFactoryContractTransactorSession) CreatePassport() (*types.Transaction, error) {
	return _PassportFactoryContract.Contract.CreatePassport(&_PassportFactoryContract.TransactOpts)
}

// ReclaimEther is a paid mutator transaction binding the contract method 0x9f727c27.
//
// Solidity: function reclaimEther() returns()
func (_PassportFactoryContract *PassportFactoryContractTransactor) ReclaimEther(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PassportFactoryContract.contract.Transact(opts, "reclaimEther")
}

// ReclaimEther is a paid mutator transaction binding the contract method 0x9f727c27.
//
// Solidity: function reclaimEther() returns()
func (_PassportFactoryContract *PassportFactoryContractSession) ReclaimEther() (*types.Transaction, error) {
	return _PassportFactoryContract.Contract.ReclaimEther(&_PassportFactoryContract.TransactOpts)
}

// ReclaimEther is a paid mutator transaction binding the contract method 0x9f727c27.
//
// Solidity: function reclaimEther() returns()
func (_PassportFactoryContract *PassportFactoryContractTransactorSession) ReclaimEther() (*types.Transaction, error) {
	return _PassportFactoryContract.Contract.ReclaimEther(&_PassportFactoryContract.TransactOpts)
}

// ReclaimToken is a paid mutator transaction binding the contract method 0x17ffc320.
//
// Solidity: function reclaimToken(_token address) returns()
func (_PassportFactoryContract *PassportFactoryContractTransactor) ReclaimToken(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _PassportFactoryContract.contract.Transact(opts, "reclaimToken", _token)
}

// ReclaimToken is a paid mutator transaction binding the contract method 0x17ffc320.
//
// Solidity: function reclaimToken(_token address) returns()
func (_PassportFactoryContract *PassportFactoryContractSession) ReclaimToken(_token common.Address) (*types.Transaction, error) {
	return _PassportFactoryContract.Contract.ReclaimToken(&_PassportFactoryContract.TransactOpts, _token)
}

// ReclaimToken is a paid mutator transaction binding the contract method 0x17ffc320.
//
// Solidity: function reclaimToken(_token address) returns()
func (_PassportFactoryContract *PassportFactoryContractTransactorSession) ReclaimToken(_token common.Address) (*types.Transaction, error) {
	return _PassportFactoryContract.Contract.ReclaimToken(&_PassportFactoryContract.TransactOpts, _token)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PassportFactoryContract *PassportFactoryContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PassportFactoryContract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PassportFactoryContract *PassportFactoryContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _PassportFactoryContract.Contract.RenounceOwnership(&_PassportFactoryContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PassportFactoryContract *PassportFactoryContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PassportFactoryContract.Contract.RenounceOwnership(&_PassportFactoryContract.TransactOpts)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(_registry address) returns()
func (_PassportFactoryContract *PassportFactoryContractTransactor) SetRegistry(opts *bind.TransactOpts, _registry common.Address) (*types.Transaction, error) {
	return _PassportFactoryContract.contract.Transact(opts, "setRegistry", _registry)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(_registry address) returns()
func (_PassportFactoryContract *PassportFactoryContractSession) SetRegistry(_registry common.Address) (*types.Transaction, error) {
	return _PassportFactoryContract.Contract.SetRegistry(&_PassportFactoryContract.TransactOpts, _registry)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(_registry address) returns()
func (_PassportFactoryContract *PassportFactoryContractTransactorSession) SetRegistry(_registry common.Address) (*types.Transaction, error) {
	return _PassportFactoryContract.Contract.SetRegistry(&_PassportFactoryContract.TransactOpts, _registry)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_PassportFactoryContract *PassportFactoryContractTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _PassportFactoryContract.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_PassportFactoryContract *PassportFactoryContractSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _PassportFactoryContract.Contract.TransferOwnership(&_PassportFactoryContract.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_PassportFactoryContract *PassportFactoryContractTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _PassportFactoryContract.Contract.TransferOwnership(&_PassportFactoryContract.TransactOpts, _newOwner)
}

// PassportFactoryContractOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the PassportFactoryContract contract.
type PassportFactoryContractOwnershipRenouncedIterator struct {
	Event *PassportFactoryContractOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *PassportFactoryContractOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PassportFactoryContractOwnershipRenounced)
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
		it.Event = new(PassportFactoryContractOwnershipRenounced)
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
func (it *PassportFactoryContractOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PassportFactoryContractOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PassportFactoryContractOwnershipRenounced represents a OwnershipRenounced event raised by the PassportFactoryContract contract.
type PassportFactoryContractOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_PassportFactoryContract *PassportFactoryContractFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*PassportFactoryContractOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _PassportFactoryContract.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PassportFactoryContractOwnershipRenouncedIterator{contract: _PassportFactoryContract.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_PassportFactoryContract *PassportFactoryContractFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *PassportFactoryContractOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _PassportFactoryContract.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PassportFactoryContractOwnershipRenounced)
				if err := _PassportFactoryContract.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// PassportFactoryContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PassportFactoryContract contract.
type PassportFactoryContractOwnershipTransferredIterator struct {
	Event *PassportFactoryContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PassportFactoryContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PassportFactoryContractOwnershipTransferred)
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
		it.Event = new(PassportFactoryContractOwnershipTransferred)
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
func (it *PassportFactoryContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PassportFactoryContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PassportFactoryContractOwnershipTransferred represents a OwnershipTransferred event raised by the PassportFactoryContract contract.
type PassportFactoryContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_PassportFactoryContract *PassportFactoryContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PassportFactoryContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PassportFactoryContract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PassportFactoryContractOwnershipTransferredIterator{contract: _PassportFactoryContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_PassportFactoryContract *PassportFactoryContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PassportFactoryContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PassportFactoryContract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PassportFactoryContractOwnershipTransferred)
				if err := _PassportFactoryContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// PassportFactoryContractPassportCreatedIterator is returned from FilterPassportCreated and is used to iterate over the raw logs and unpacked data for PassportCreated events raised by the PassportFactoryContract contract.
type PassportFactoryContractPassportCreatedIterator struct {
	Event *PassportFactoryContractPassportCreated // Event containing the contract specifics and raw log

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
func (it *PassportFactoryContractPassportCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PassportFactoryContractPassportCreated)
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
		it.Event = new(PassportFactoryContractPassportCreated)
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
func (it *PassportFactoryContractPassportCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PassportFactoryContractPassportCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PassportFactoryContractPassportCreated represents a PassportCreated event raised by the PassportFactoryContract contract.
type PassportFactoryContractPassportCreated struct {
	Passport common.Address
	Owner    common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPassportCreated is a free log retrieval operation binding the contract event 0x03f096f07a4d27c54645fa682640d44179c821150e16c48b27130ca928fa937c.
//
// Solidity: e PassportCreated(passport indexed address, owner indexed address)
func (_PassportFactoryContract *PassportFactoryContractFilterer) FilterPassportCreated(opts *bind.FilterOpts, passport []common.Address, owner []common.Address) (*PassportFactoryContractPassportCreatedIterator, error) {

	var passportRule []interface{}
	for _, passportItem := range passport {
		passportRule = append(passportRule, passportItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _PassportFactoryContract.contract.FilterLogs(opts, "PassportCreated", passportRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &PassportFactoryContractPassportCreatedIterator{contract: _PassportFactoryContract.contract, event: "PassportCreated", logs: logs, sub: sub}, nil
}

// WatchPassportCreated is a free log subscription operation binding the contract event 0x03f096f07a4d27c54645fa682640d44179c821150e16c48b27130ca928fa937c.
//
// Solidity: e PassportCreated(passport indexed address, owner indexed address)
func (_PassportFactoryContract *PassportFactoryContractFilterer) WatchPassportCreated(opts *bind.WatchOpts, sink chan<- *PassportFactoryContractPassportCreated, passport []common.Address, owner []common.Address) (event.Subscription, error) {

	var passportRule []interface{}
	for _, passportItem := range passport {
		passportRule = append(passportRule, passportItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _PassportFactoryContract.contract.WatchLogs(opts, "PassportCreated", passportRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PassportFactoryContractPassportCreated)
				if err := _PassportFactoryContract.contract.UnpackLog(event, "PassportCreated", log); err != nil {
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

// PassportFactoryContractPassportLogicRegistryChangedIterator is returned from FilterPassportLogicRegistryChanged and is used to iterate over the raw logs and unpacked data for PassportLogicRegistryChanged events raised by the PassportFactoryContract contract.
type PassportFactoryContractPassportLogicRegistryChangedIterator struct {
	Event *PassportFactoryContractPassportLogicRegistryChanged // Event containing the contract specifics and raw log

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
func (it *PassportFactoryContractPassportLogicRegistryChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PassportFactoryContractPassportLogicRegistryChanged)
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
		it.Event = new(PassportFactoryContractPassportLogicRegistryChanged)
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
func (it *PassportFactoryContractPassportLogicRegistryChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PassportFactoryContractPassportLogicRegistryChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PassportFactoryContractPassportLogicRegistryChanged represents a PassportLogicRegistryChanged event raised by the PassportFactoryContract contract.
type PassportFactoryContractPassportLogicRegistryChanged struct {
	OldRegistry common.Address
	NewRegistry common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPassportLogicRegistryChanged is a free log retrieval operation binding the contract event 0x5c2abfd67230c0e47d6de28402bfe206c7a57283cba891416ed657fd70a714c2.
//
// Solidity: e PassportLogicRegistryChanged(oldRegistry indexed address, newRegistry indexed address)
func (_PassportFactoryContract *PassportFactoryContractFilterer) FilterPassportLogicRegistryChanged(opts *bind.FilterOpts, oldRegistry []common.Address, newRegistry []common.Address) (*PassportFactoryContractPassportLogicRegistryChangedIterator, error) {

	var oldRegistryRule []interface{}
	for _, oldRegistryItem := range oldRegistry {
		oldRegistryRule = append(oldRegistryRule, oldRegistryItem)
	}
	var newRegistryRule []interface{}
	for _, newRegistryItem := range newRegistry {
		newRegistryRule = append(newRegistryRule, newRegistryItem)
	}

	logs, sub, err := _PassportFactoryContract.contract.FilterLogs(opts, "PassportLogicRegistryChanged", oldRegistryRule, newRegistryRule)
	if err != nil {
		return nil, err
	}
	return &PassportFactoryContractPassportLogicRegistryChangedIterator{contract: _PassportFactoryContract.contract, event: "PassportLogicRegistryChanged", logs: logs, sub: sub}, nil
}

// WatchPassportLogicRegistryChanged is a free log subscription operation binding the contract event 0x5c2abfd67230c0e47d6de28402bfe206c7a57283cba891416ed657fd70a714c2.
//
// Solidity: e PassportLogicRegistryChanged(oldRegistry indexed address, newRegistry indexed address)
func (_PassportFactoryContract *PassportFactoryContractFilterer) WatchPassportLogicRegistryChanged(opts *bind.WatchOpts, sink chan<- *PassportFactoryContractPassportLogicRegistryChanged, oldRegistry []common.Address, newRegistry []common.Address) (event.Subscription, error) {

	var oldRegistryRule []interface{}
	for _, oldRegistryItem := range oldRegistry {
		oldRegistryRule = append(oldRegistryRule, oldRegistryItem)
	}
	var newRegistryRule []interface{}
	for _, newRegistryItem := range newRegistry {
		newRegistryRule = append(newRegistryRule, newRegistryItem)
	}

	logs, sub, err := _PassportFactoryContract.contract.WatchLogs(opts, "PassportLogicRegistryChanged", oldRegistryRule, newRegistryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PassportFactoryContractPassportLogicRegistryChanged)
				if err := _PassportFactoryContract.contract.UnpackLog(event, "PassportLogicRegistryChanged", log); err != nil {
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
