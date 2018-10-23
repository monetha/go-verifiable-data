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

// PassportLogicContractABI is the input ABI used to generate the binding from.
const PassportLogicContractABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes32\"}],\"name\":\"deleteAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_factProvider\",\"type\":\"address\"},{\"name\":\"_key\",\"type\":\"bytes32\"}],\"name\":\"getTxDataBlockNumber\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"},{\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isAllowedFactProvider\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_onlyWhitelist\",\"type\":\"bool\"}],\"name\":\"setWhitelistOnlyPermission\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes32\"}],\"name\":\"deleteBool\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes32\"},{\"name\":\"_value\",\"type\":\"bytes\"}],\"name\":\"setBytes\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes32\"},{\"name\":\"_value\",\"type\":\"int256\"}],\"name\":\"setInt\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claimOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes32\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"setTxDataBlockNumber\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes32\"}],\"name\":\"deleteBytes\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isWhitelistOnlyPermissionSet\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_factProvider\",\"type\":\"address\"},{\"name\":\"_key\",\"type\":\"bytes32\"}],\"name\":\"getBytes\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"},{\"name\":\"value\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"removeFactProviderFromWhitelist\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes32\"},{\"name\":\"_value\",\"type\":\"string\"}],\"name\":\"setString\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_factProvider\",\"type\":\"address\"},{\"name\":\"_key\",\"type\":\"bytes32\"}],\"name\":\"getUint\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_factProvider\",\"type\":\"address\"},{\"name\":\"_key\",\"type\":\"bytes32\"}],\"name\":\"getAddress\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"},{\"name\":\"value\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isFactProviderInWhitelist\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes32\"}],\"name\":\"deleteInt\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_factProvider\",\"type\":\"address\"},{\"name\":\"_key\",\"type\":\"bytes32\"}],\"name\":\"getInt\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"},{\"name\":\"value\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_factProvider\",\"type\":\"address\"},{\"name\":\"_key\",\"type\":\"bytes32\"}],\"name\":\"getBool\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"},{\"name\":\"value\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes32\"}],\"name\":\"deleteTxDataBlockNumber\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes32\"},{\"name\":\"_value\",\"type\":\"bool\"}],\"name\":\"setBool\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"addFactProviderToWhitelist\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes32\"},{\"name\":\"_value\",\"type\":\"address\"}],\"name\":\"setAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes32\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"setUint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes32\"}],\"name\":\"deleteUint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_factProvider\",\"type\":\"address\"},{\"name\":\"_key\",\"type\":\"bytes32\"}],\"name\":\"getString\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"},{\"name\":\"value\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes32\"}],\"name\":\"deleteString\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"factProvider\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"TxDataUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"factProvider\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"TxDataDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"factProvider\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"BytesUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"factProvider\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"BytesDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"factProvider\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"StringUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"factProvider\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"StringDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"factProvider\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"BoolUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"factProvider\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"BoolDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"factProvider\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"IntUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"factProvider\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"IntDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"factProvider\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"UintUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"factProvider\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"UintDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"factProvider\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"AddressUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"factProvider\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"AddressDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"onlyWhitelist\",\"type\":\"bool\"}],\"name\":\"WhitelistOnlyPermissionSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"factProvider\",\"type\":\"address\"}],\"name\":\"WhitelistFactProviderAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"factProvider\",\"type\":\"address\"}],\"name\":\"WhitelistFactProviderRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// PassportLogicContractBin is the compiled bytecode used for deploying new contracts.
const PassportLogicContractBin = `0x608060408190527f6f72672e6d6f6e657468612e70726f78792e6f776e6572000000000000000000905261003b3364010000000061009a810204565b604080517f6f72672e6d6f6e657468612e70726f78792e70656e64696e674f776e657200008152905190819003601e0190207fcfd0c6ea5352192d7d4c5d4e7a73c5da12c871730cb60ff57879cbe7b403bb521461009557fe5b6100be565b7f3ca57e4b51fc2e18497b219410298879868edada7e6fe5132c8feceb0a080d2255565b61154e806100cd6000396000f3006080604052600436106101955763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630e14a376811461019a578063174a6277146101b457806320888158146101f35780632b42a89a146102285780632c62ff2d146102425780632e28d0841461025a5780633e49bed01461027e5780634e71e0c8146102995780635b2a372d146102ae578063616b59f6146102d257806361a2d076146102ea5780636556f767146102ff5780636d4e7802146103a65780636e899550146103c7578063715018a6146103eb57806371658552146104005780637ac4ed64146104245780638ae44a081461046b5780638c1600951461048c5780638da5cb5b146104a457806395ee8bae146104d55780639d74b37d146104f9578063a2b6cbe114610538578063abfdcced14610550578063ac0aeb1a1461056d578063ca446dd91461058e578063e2a4853a146105b2578063e2b202bf146105cd578063e30c3978146105e5578063e318de73146105fa578063f2fde38b1461061e578063f6bb3cc41461063f575b600080fd5b3480156101a657600080fd5b506101b2600435610657565b005b3480156101c057600080fd5b506101d8600160a060020a0360043516602435610663565b60408051921515835260208301919091528051918290030190f35b3480156101ff57600080fd5b50610214600160a060020a036004351661067b565b604080519115158252519081900360200190f35b34801561023457600080fd5b506101b260043515156106d1565b34801561024e57600080fd5b506101b260043561072a565b34801561026657600080fd5b506101b2600480359060248035908101910135610733565b34801561028a57600080fd5b506101b2600435602435610772565b3480156102a557600080fd5b506101b2610780565b3480156102ba57600080fd5b506101b2600480359060248035908101910135610806565b3480156102de57600080fd5b506101b260043561088b565b3480156102f657600080fd5b50610214610894565b34801561030b57600080fd5b50610323600160a060020a036004351660243561089e565b604051808315151515815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561036a578181015183820152602001610352565b50505050905090810190601f1680156103975780820380516001836020036101000a031916815260200191505b50935050505060405180910390f35b3480156103b257600080fd5b506101b2600160a060020a03600435166108ac565b3480156103d357600080fd5b506101b2600480359060248035908101910135610911565b3480156103f757600080fd5b506101b261094b565b34801561040c57600080fd5b506101d8600160a060020a03600435166024356109ae565b34801561043057600080fd5b50610448600160a060020a03600435166024356109bb565b604080519215158352600160a060020a0390911660208301528051918290030190f35b34801561047757600080fd5b50610214600160a060020a03600435166109c8565b34801561049857600080fd5b506101b26004356109e6565b3480156104b057600080fd5b506104b96109ef565b60408051600160a060020a039092168252519081900360200190f35b3480156104e157600080fd5b506101d8600160a060020a03600435166024356109fe565b34801561050557600080fd5b5061051d600160a060020a0360043516602435610a0b565b60408051921515835290151560208301528051918290030190f35b34801561054457600080fd5b506101b2600435610a18565b34801561055c57600080fd5b506101b26004356024351515610a7f565b34801561057957600080fd5b506101b2600160a060020a0360043516610a89565b34801561059a57600080fd5b506101b2600435600160a060020a0360243516610af1565b3480156105be57600080fd5b506101b2600435602435610afb565b3480156105d957600080fd5b506101b2600435610b05565b3480156105f157600080fd5b506104b9610b0e565b34801561060657600080fd5b50610323600160a060020a0360043516602435610b18565b34801561062a57600080fd5b506101b2600160a060020a0360043516610b26565b34801561064b57600080fd5b506101b2600435610b4b565b61066081610b54565b50565b6000806106708484610bc7565b915091509250929050565b60075460009060ff1615806106a85750600160a060020a03821660009081526008602052604090205460ff165b806106cb57506106b6610bfb565b600160a060020a031682600160a060020a0316145b92915050565b6106d9610bfb565b600160a060020a031633146106ed57600080fd5b6007805460ff19168215159081179091556040517f6e1eb3fa845047cf9d5fd1deb1b7770b4e26163dd4d88cefb14fe7964a70059590600090a250565b61066081610c20565b61076d8383838080601f01602080910402602001604051908101604052809392919081815260200183838082843750610c82945050505050565b505050565b61077c8282610d1c565b5050565b610788610da0565b600160a060020a0316331461079c57600080fd5b6107a4610da0565b600160a060020a03166107b5610bfb565b600160a060020a03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36107fa6107f5610da0565b610dc5565b6108046000610de9565b565b61080f3361067b565b151561081a57600080fd5b6040805180820182526001808252436020808401918252336000818152600683528681208a82529092528582209451855460ff191690151517855591519390920192909255915185927fcbde9cd310365ff3de3ad9cb44ed688d8880c7b1d15df733a73880fb84d032de91a3505050565b61066081610e0d565b60075460ff165b90565b600060606106708484610e83565b6108b4610bfb565b600160a060020a031633146108c857600080fd5b600160a060020a038116600081815260086020526040808220805460ff19169055517f5bf47f016603dfff957e6086bd9a350212686e8f1d9208b3db47c5a3107c78789190a250565b61076d8383838080601f01602080910402602001604051908101604052809392919081815260200183838082843750610f4d945050505050565b610953610bfb565b600160a060020a0316331461096757600080fd5b61096f610bfb565b600160a060020a03167ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482060405160405180910390a26108046000610dc5565b6000806106708484610fe7565b600080610670848461101a565b600160a060020a031660009081526008602052604090205460ff1690565b6106608161104d565b60006109f9610bfb565b905090565b60008061067084846110b4565b60008061067084846110e8565b610a213361067b565b1515610a2c57600080fd5b336000818152600660209081526040808320858452909152808220805460ff19168155600101829055518392917f9b9532c8ba5879554dee7c3e9374bd26dbfb2f79cf7f247bcf90648fc81fc16491a350565b61077c828261111d565b610a91610bfb565b600160a060020a03163314610aa557600080fd5b600160a060020a038116600081815260086020526040808220805460ff19166001179055517fff8ac5f0e9cdc8f7ffbbc29ab1301ee649cf01cf7a6a9265a2d90587612c41759190a250565b61077c82826111b0565b61077c828261125e565b610660816112e1565b60006109f9610da0565b600060606106708484611348565b610b2e610bfb565b600160a060020a03163314610b4257600080fd5b61066081610de9565b610660816113d4565b610b5d3361067b565b1515610b6857600080fd5b33600081815260208181526040808320858452909152808220805474ffffffffffffffffffffffffffffffffffffffffff19169055518392917fed9474c4702ee9da4d0d2085840fafed5912b8e13daf1d6ea341ebe98803aec291a350565b600160a060020a0391909116600090815260066020908152604080832093835292905220805460019091015460ff90911691565b7f3ca57e4b51fc2e18497b219410298879868edada7e6fe5132c8feceb0a080d225490565b610c293361067b565b1515610c3457600080fd5b336000818152600360209081526040808320858452909152808220805461ffff19169055518392917f191bd7fc1f22ce6b74b7d7e42126f0ed225b109728607c99971de9e78b0db6eb91a350565b610c8b3361067b565b1515610c9657600080fd5b60408051808201825260018082526020808301858152336000908152600583528581208882528352949094208351815460ff191690151517815593518051939493610ce893850192919091019061144a565b505060405183915033907fee56af66250a5cbba377132a356dc403d520776e30969594927024c44af8440b90600090a35050565b610d253361067b565b1515610d3057600080fd5b60408051808201825260018082526020808301858152336000818152600284528681208982529093528583209451855460ff19169015151785559051939092019290925591518492917fdb4ad6a3ed2bdacbae02da45a5bb54192cdb7d72648674d13e08d2fc4c45ac6691a35050565b7fcfd0c6ea5352192d7d4c5d4e7a73c5da12c871730cb60ff57879cbe7b403bb525490565b7f3ca57e4b51fc2e18497b219410298879868edada7e6fe5132c8feceb0a080d2255565b7fcfd0c6ea5352192d7d4c5d4e7a73c5da12c871730cb60ff57879cbe7b403bb5255565b610e163361067b565b1515610e2157600080fd5b3360009081526005602090815260408083208484529091528120805460ff1916815590610e5160018301826114c8565b5050604051819033907f0eaf95ab41777ff88e28181b1c43418524616bcef6d37c3b216015b5548cb10890600090a350565b600160a060020a038216600090815260056020908152604080832084845282528083208054600180830180548551601f600294831615610100026000190190921693909304908101879004870283018701909552848252606095939460ff9093169390928391830182828015610f3a5780601f10610f0f57610100808354040283529160200191610f3a565b820191906000526020600020905b815481529060010190602001808311610f1d57829003601f168201915b5050505050905092509250509250929050565b610f563361067b565b1515610f6157600080fd5b60408051808201825260018082526020808301858152336000908152600483528581208882528352949094208351815460ff191690151517815593518051939493610fb393850192919091019061144a565b505060405183915033907f43e6b7e3323b4598401023341c086c07c3ff5577f594b5aab9c065f2c3c9d59090600090a35050565b600160a060020a03919091166000908152600160208181526040808420948452939052919020805491015460ff90911691565b600160a060020a03918216600090815260208181526040808320938352929052205460ff81169261010090910490911690565b6110563361067b565b151561106157600080fd5b336000818152600260209081526040808320858452909152808220805460ff19168155600101829055518392917f3c3ff48e02e407eb1e78310d11b5e3f9e735263a9cafc2bcf4aa981b8ecb32a591a350565b600160a060020a0391909116600090815260026020908152604080832093835292905220805460019091015460ff90911691565b600160a060020a03919091166000908152600360209081526040808320938352929052205460ff808216926101009092041690565b6111263361067b565b151561113157600080fd5b6040805180820182526001815282151560208083019182523360008181526003835285812088825290925284822093518454935115156101000261ff001991151560ff1990951694909417169290921790925591518492917f68cc3496efaac4c1f2c0cd52da916138f6c5fc541992f05d97423a89b6914ae591a35050565b6111b93361067b565b15156111c457600080fd5b60408051808201825260018152600160a060020a03838116602080840191825233600081815280835286812089825290925285822094518554935160ff199094169015151774ffffffffffffffffffffffffffffffffffffffff001916610100939094169290920292909217909255915184927f8e7e6ab6c4613205e833e1faf5415d78dd10130a8828f729c08036e2a7a6277091a35050565b6112673361067b565b151561127257600080fd5b604080518082018252600180825260208083018581523360008181528484528681208982529093528583209451855460ff19169015151785559051939092019290925591518492917fff76cb7634629f8d05011ab2a58380dfc0743157a70ba4173f866e113eca75c091a35050565b6112ea3361067b565b15156112f557600080fd5b336000818152600160208181526040808420868552909152808320805460ff19168155909101829055518392917fd5f853f7aaba549b5a46be1de24cac4c20e716b4c603d24b2b0b7b5d97ca1c4d91a350565b600160a060020a038216600090815260046020908152604080832084845282528083208054600180830180548551601f600294831615610100026000190190921693909304908101879004870283018701909552848252606095939460ff9093169390928391830182828015610f3a5780601f10610f0f57610100808354040283529160200191610f3a565b6113dd3361067b565b15156113e857600080fd5b3360009081526004602090815260408083208484529091528120805460ff191681559061141860018301826114c8565b5050604051819033907f491a0e0281af9fc88ff024cffc24db22eddd4f826cb30d84936592967dc9237c90600090a350565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061148b57805160ff19168380011785556114b8565b828001600101855582156114b8579182015b828111156114b857825182559160200191906001019061149d565b506114c4929150611508565b5090565b50805460018160011615610100020316600290046000825580601f106114ee5750610660565b601f01602090049060005260206000209081019061066091905b61089b91905b808211156114c4576000815560010161150e5600a165627a7a72305820dd374e3d3a76c60aec14d1129c09c04bcc5d078abce023bc54ebd19a16c6b8600029`

// DeployPassportLogicContract deploys a new Ethereum contract, binding an instance of PassportLogicContract to it.
func DeployPassportLogicContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PassportLogicContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PassportLogicContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PassportLogicContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PassportLogicContract{PassportLogicContractCaller: PassportLogicContractCaller{contract: contract}, PassportLogicContractTransactor: PassportLogicContractTransactor{contract: contract}, PassportLogicContractFilterer: PassportLogicContractFilterer{contract: contract}}, nil
}

// PassportLogicContract is an auto generated Go binding around an Ethereum contract.
type PassportLogicContract struct {
	PassportLogicContractCaller     // Read-only binding to the contract
	PassportLogicContractTransactor // Write-only binding to the contract
	PassportLogicContractFilterer   // Log filterer for contract events
}

// PassportLogicContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type PassportLogicContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PassportLogicContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PassportLogicContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PassportLogicContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PassportLogicContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PassportLogicContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PassportLogicContractSession struct {
	Contract     *PassportLogicContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// PassportLogicContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PassportLogicContractCallerSession struct {
	Contract *PassportLogicContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// PassportLogicContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PassportLogicContractTransactorSession struct {
	Contract     *PassportLogicContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// PassportLogicContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type PassportLogicContractRaw struct {
	Contract *PassportLogicContract // Generic contract binding to access the raw methods on
}

// PassportLogicContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PassportLogicContractCallerRaw struct {
	Contract *PassportLogicContractCaller // Generic read-only contract binding to access the raw methods on
}

// PassportLogicContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PassportLogicContractTransactorRaw struct {
	Contract *PassportLogicContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPassportLogicContract creates a new instance of PassportLogicContract, bound to a specific deployed contract.
func NewPassportLogicContract(address common.Address, backend bind.ContractBackend) (*PassportLogicContract, error) {
	contract, err := bindPassportLogicContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PassportLogicContract{PassportLogicContractCaller: PassportLogicContractCaller{contract: contract}, PassportLogicContractTransactor: PassportLogicContractTransactor{contract: contract}, PassportLogicContractFilterer: PassportLogicContractFilterer{contract: contract}}, nil
}

// NewPassportLogicContractCaller creates a new read-only instance of PassportLogicContract, bound to a specific deployed contract.
func NewPassportLogicContractCaller(address common.Address, caller bind.ContractCaller) (*PassportLogicContractCaller, error) {
	contract, err := bindPassportLogicContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PassportLogicContractCaller{contract: contract}, nil
}

// NewPassportLogicContractTransactor creates a new write-only instance of PassportLogicContract, bound to a specific deployed contract.
func NewPassportLogicContractTransactor(address common.Address, transactor bind.ContractTransactor) (*PassportLogicContractTransactor, error) {
	contract, err := bindPassportLogicContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PassportLogicContractTransactor{contract: contract}, nil
}

// NewPassportLogicContractFilterer creates a new log filterer instance of PassportLogicContract, bound to a specific deployed contract.
func NewPassportLogicContractFilterer(address common.Address, filterer bind.ContractFilterer) (*PassportLogicContractFilterer, error) {
	contract, err := bindPassportLogicContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PassportLogicContractFilterer{contract: contract}, nil
}

// bindPassportLogicContract binds a generic wrapper to an already deployed contract.
func bindPassportLogicContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PassportLogicContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PassportLogicContract *PassportLogicContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PassportLogicContract.Contract.PassportLogicContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PassportLogicContract *PassportLogicContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PassportLogicContract.Contract.PassportLogicContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PassportLogicContract *PassportLogicContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PassportLogicContract.Contract.PassportLogicContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PassportLogicContract *PassportLogicContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PassportLogicContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PassportLogicContract *PassportLogicContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PassportLogicContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PassportLogicContract *PassportLogicContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PassportLogicContract.Contract.contract.Transact(opts, method, params...)
}

// GetAddress is a free data retrieval call binding the contract method 0x7ac4ed64.
//
// Solidity: function getAddress(_factProvider address, _key bytes32) constant returns(success bool, value address)
func (_PassportLogicContract *PassportLogicContractCaller) GetAddress(opts *bind.CallOpts, _factProvider common.Address, _key [32]byte) (struct {
	Success bool
	Value   common.Address
}, error) {
	ret := new(struct {
		Success bool
		Value   common.Address
	})
	out := ret
	err := _PassportLogicContract.contract.Call(opts, out, "getAddress", _factProvider, _key)
	return *ret, err
}

// GetAddress is a free data retrieval call binding the contract method 0x7ac4ed64.
//
// Solidity: function getAddress(_factProvider address, _key bytes32) constant returns(success bool, value address)
func (_PassportLogicContract *PassportLogicContractSession) GetAddress(_factProvider common.Address, _key [32]byte) (struct {
	Success bool
	Value   common.Address
}, error) {
	return _PassportLogicContract.Contract.GetAddress(&_PassportLogicContract.CallOpts, _factProvider, _key)
}

// GetAddress is a free data retrieval call binding the contract method 0x7ac4ed64.
//
// Solidity: function getAddress(_factProvider address, _key bytes32) constant returns(success bool, value address)
func (_PassportLogicContract *PassportLogicContractCallerSession) GetAddress(_factProvider common.Address, _key [32]byte) (struct {
	Success bool
	Value   common.Address
}, error) {
	return _PassportLogicContract.Contract.GetAddress(&_PassportLogicContract.CallOpts, _factProvider, _key)
}

// GetBool is a free data retrieval call binding the contract method 0x9d74b37d.
//
// Solidity: function getBool(_factProvider address, _key bytes32) constant returns(success bool, value bool)
func (_PassportLogicContract *PassportLogicContractCaller) GetBool(opts *bind.CallOpts, _factProvider common.Address, _key [32]byte) (struct {
	Success bool
	Value   bool
}, error) {
	ret := new(struct {
		Success bool
		Value   bool
	})
	out := ret
	err := _PassportLogicContract.contract.Call(opts, out, "getBool", _factProvider, _key)
	return *ret, err
}

// GetBool is a free data retrieval call binding the contract method 0x9d74b37d.
//
// Solidity: function getBool(_factProvider address, _key bytes32) constant returns(success bool, value bool)
func (_PassportLogicContract *PassportLogicContractSession) GetBool(_factProvider common.Address, _key [32]byte) (struct {
	Success bool
	Value   bool
}, error) {
	return _PassportLogicContract.Contract.GetBool(&_PassportLogicContract.CallOpts, _factProvider, _key)
}

// GetBool is a free data retrieval call binding the contract method 0x9d74b37d.
//
// Solidity: function getBool(_factProvider address, _key bytes32) constant returns(success bool, value bool)
func (_PassportLogicContract *PassportLogicContractCallerSession) GetBool(_factProvider common.Address, _key [32]byte) (struct {
	Success bool
	Value   bool
}, error) {
	return _PassportLogicContract.Contract.GetBool(&_PassportLogicContract.CallOpts, _factProvider, _key)
}

// GetBytes is a free data retrieval call binding the contract method 0x6556f767.
//
// Solidity: function getBytes(_factProvider address, _key bytes32) constant returns(success bool, value bytes)
func (_PassportLogicContract *PassportLogicContractCaller) GetBytes(opts *bind.CallOpts, _factProvider common.Address, _key [32]byte) (struct {
	Success bool
	Value   []byte
}, error) {
	ret := new(struct {
		Success bool
		Value   []byte
	})
	out := ret
	err := _PassportLogicContract.contract.Call(opts, out, "getBytes", _factProvider, _key)
	return *ret, err
}

// GetBytes is a free data retrieval call binding the contract method 0x6556f767.
//
// Solidity: function getBytes(_factProvider address, _key bytes32) constant returns(success bool, value bytes)
func (_PassportLogicContract *PassportLogicContractSession) GetBytes(_factProvider common.Address, _key [32]byte) (struct {
	Success bool
	Value   []byte
}, error) {
	return _PassportLogicContract.Contract.GetBytes(&_PassportLogicContract.CallOpts, _factProvider, _key)
}

// GetBytes is a free data retrieval call binding the contract method 0x6556f767.
//
// Solidity: function getBytes(_factProvider address, _key bytes32) constant returns(success bool, value bytes)
func (_PassportLogicContract *PassportLogicContractCallerSession) GetBytes(_factProvider common.Address, _key [32]byte) (struct {
	Success bool
	Value   []byte
}, error) {
	return _PassportLogicContract.Contract.GetBytes(&_PassportLogicContract.CallOpts, _factProvider, _key)
}

// GetInt is a free data retrieval call binding the contract method 0x95ee8bae.
//
// Solidity: function getInt(_factProvider address, _key bytes32) constant returns(success bool, value int256)
func (_PassportLogicContract *PassportLogicContractCaller) GetInt(opts *bind.CallOpts, _factProvider common.Address, _key [32]byte) (struct {
	Success bool
	Value   *big.Int
}, error) {
	ret := new(struct {
		Success bool
		Value   *big.Int
	})
	out := ret
	err := _PassportLogicContract.contract.Call(opts, out, "getInt", _factProvider, _key)
	return *ret, err
}

// GetInt is a free data retrieval call binding the contract method 0x95ee8bae.
//
// Solidity: function getInt(_factProvider address, _key bytes32) constant returns(success bool, value int256)
func (_PassportLogicContract *PassportLogicContractSession) GetInt(_factProvider common.Address, _key [32]byte) (struct {
	Success bool
	Value   *big.Int
}, error) {
	return _PassportLogicContract.Contract.GetInt(&_PassportLogicContract.CallOpts, _factProvider, _key)
}

// GetInt is a free data retrieval call binding the contract method 0x95ee8bae.
//
// Solidity: function getInt(_factProvider address, _key bytes32) constant returns(success bool, value int256)
func (_PassportLogicContract *PassportLogicContractCallerSession) GetInt(_factProvider common.Address, _key [32]byte) (struct {
	Success bool
	Value   *big.Int
}, error) {
	return _PassportLogicContract.Contract.GetInt(&_PassportLogicContract.CallOpts, _factProvider, _key)
}

// GetString is a free data retrieval call binding the contract method 0xe318de73.
//
// Solidity: function getString(_factProvider address, _key bytes32) constant returns(success bool, value string)
func (_PassportLogicContract *PassportLogicContractCaller) GetString(opts *bind.CallOpts, _factProvider common.Address, _key [32]byte) (struct {
	Success bool
	Value   string
}, error) {
	ret := new(struct {
		Success bool
		Value   string
	})
	out := ret
	err := _PassportLogicContract.contract.Call(opts, out, "getString", _factProvider, _key)
	return *ret, err
}

// GetString is a free data retrieval call binding the contract method 0xe318de73.
//
// Solidity: function getString(_factProvider address, _key bytes32) constant returns(success bool, value string)
func (_PassportLogicContract *PassportLogicContractSession) GetString(_factProvider common.Address, _key [32]byte) (struct {
	Success bool
	Value   string
}, error) {
	return _PassportLogicContract.Contract.GetString(&_PassportLogicContract.CallOpts, _factProvider, _key)
}

// GetString is a free data retrieval call binding the contract method 0xe318de73.
//
// Solidity: function getString(_factProvider address, _key bytes32) constant returns(success bool, value string)
func (_PassportLogicContract *PassportLogicContractCallerSession) GetString(_factProvider common.Address, _key [32]byte) (struct {
	Success bool
	Value   string
}, error) {
	return _PassportLogicContract.Contract.GetString(&_PassportLogicContract.CallOpts, _factProvider, _key)
}

// GetTxDataBlockNumber is a free data retrieval call binding the contract method 0x174a6277.
//
// Solidity: function getTxDataBlockNumber(_factProvider address, _key bytes32) constant returns(success bool, blockNumber uint256)
func (_PassportLogicContract *PassportLogicContractCaller) GetTxDataBlockNumber(opts *bind.CallOpts, _factProvider common.Address, _key [32]byte) (struct {
	Success     bool
	BlockNumber *big.Int
}, error) {
	ret := new(struct {
		Success     bool
		BlockNumber *big.Int
	})
	out := ret
	err := _PassportLogicContract.contract.Call(opts, out, "getTxDataBlockNumber", _factProvider, _key)
	return *ret, err
}

// GetTxDataBlockNumber is a free data retrieval call binding the contract method 0x174a6277.
//
// Solidity: function getTxDataBlockNumber(_factProvider address, _key bytes32) constant returns(success bool, blockNumber uint256)
func (_PassportLogicContract *PassportLogicContractSession) GetTxDataBlockNumber(_factProvider common.Address, _key [32]byte) (struct {
	Success     bool
	BlockNumber *big.Int
}, error) {
	return _PassportLogicContract.Contract.GetTxDataBlockNumber(&_PassportLogicContract.CallOpts, _factProvider, _key)
}

// GetTxDataBlockNumber is a free data retrieval call binding the contract method 0x174a6277.
//
// Solidity: function getTxDataBlockNumber(_factProvider address, _key bytes32) constant returns(success bool, blockNumber uint256)
func (_PassportLogicContract *PassportLogicContractCallerSession) GetTxDataBlockNumber(_factProvider common.Address, _key [32]byte) (struct {
	Success     bool
	BlockNumber *big.Int
}, error) {
	return _PassportLogicContract.Contract.GetTxDataBlockNumber(&_PassportLogicContract.CallOpts, _factProvider, _key)
}

// GetUint is a free data retrieval call binding the contract method 0x71658552.
//
// Solidity: function getUint(_factProvider address, _key bytes32) constant returns(success bool, value uint256)
func (_PassportLogicContract *PassportLogicContractCaller) GetUint(opts *bind.CallOpts, _factProvider common.Address, _key [32]byte) (struct {
	Success bool
	Value   *big.Int
}, error) {
	ret := new(struct {
		Success bool
		Value   *big.Int
	})
	out := ret
	err := _PassportLogicContract.contract.Call(opts, out, "getUint", _factProvider, _key)
	return *ret, err
}

// GetUint is a free data retrieval call binding the contract method 0x71658552.
//
// Solidity: function getUint(_factProvider address, _key bytes32) constant returns(success bool, value uint256)
func (_PassportLogicContract *PassportLogicContractSession) GetUint(_factProvider common.Address, _key [32]byte) (struct {
	Success bool
	Value   *big.Int
}, error) {
	return _PassportLogicContract.Contract.GetUint(&_PassportLogicContract.CallOpts, _factProvider, _key)
}

// GetUint is a free data retrieval call binding the contract method 0x71658552.
//
// Solidity: function getUint(_factProvider address, _key bytes32) constant returns(success bool, value uint256)
func (_PassportLogicContract *PassportLogicContractCallerSession) GetUint(_factProvider common.Address, _key [32]byte) (struct {
	Success bool
	Value   *big.Int
}, error) {
	return _PassportLogicContract.Contract.GetUint(&_PassportLogicContract.CallOpts, _factProvider, _key)
}

// IsAllowedFactProvider is a free data retrieval call binding the contract method 0x20888158.
//
// Solidity: function isAllowedFactProvider(_address address) constant returns(bool)
func (_PassportLogicContract *PassportLogicContractCaller) IsAllowedFactProvider(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PassportLogicContract.contract.Call(opts, out, "isAllowedFactProvider", _address)
	return *ret0, err
}

// IsAllowedFactProvider is a free data retrieval call binding the contract method 0x20888158.
//
// Solidity: function isAllowedFactProvider(_address address) constant returns(bool)
func (_PassportLogicContract *PassportLogicContractSession) IsAllowedFactProvider(_address common.Address) (bool, error) {
	return _PassportLogicContract.Contract.IsAllowedFactProvider(&_PassportLogicContract.CallOpts, _address)
}

// IsAllowedFactProvider is a free data retrieval call binding the contract method 0x20888158.
//
// Solidity: function isAllowedFactProvider(_address address) constant returns(bool)
func (_PassportLogicContract *PassportLogicContractCallerSession) IsAllowedFactProvider(_address common.Address) (bool, error) {
	return _PassportLogicContract.Contract.IsAllowedFactProvider(&_PassportLogicContract.CallOpts, _address)
}

// IsFactProviderInWhitelist is a free data retrieval call binding the contract method 0x8ae44a08.
//
// Solidity: function isFactProviderInWhitelist(_address address) constant returns(bool)
func (_PassportLogicContract *PassportLogicContractCaller) IsFactProviderInWhitelist(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PassportLogicContract.contract.Call(opts, out, "isFactProviderInWhitelist", _address)
	return *ret0, err
}

// IsFactProviderInWhitelist is a free data retrieval call binding the contract method 0x8ae44a08.
//
// Solidity: function isFactProviderInWhitelist(_address address) constant returns(bool)
func (_PassportLogicContract *PassportLogicContractSession) IsFactProviderInWhitelist(_address common.Address) (bool, error) {
	return _PassportLogicContract.Contract.IsFactProviderInWhitelist(&_PassportLogicContract.CallOpts, _address)
}

// IsFactProviderInWhitelist is a free data retrieval call binding the contract method 0x8ae44a08.
//
// Solidity: function isFactProviderInWhitelist(_address address) constant returns(bool)
func (_PassportLogicContract *PassportLogicContractCallerSession) IsFactProviderInWhitelist(_address common.Address) (bool, error) {
	return _PassportLogicContract.Contract.IsFactProviderInWhitelist(&_PassportLogicContract.CallOpts, _address)
}

// IsWhitelistOnlyPermissionSet is a free data retrieval call binding the contract method 0x61a2d076.
//
// Solidity: function isWhitelistOnlyPermissionSet() constant returns(bool)
func (_PassportLogicContract *PassportLogicContractCaller) IsWhitelistOnlyPermissionSet(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PassportLogicContract.contract.Call(opts, out, "isWhitelistOnlyPermissionSet")
	return *ret0, err
}

// IsWhitelistOnlyPermissionSet is a free data retrieval call binding the contract method 0x61a2d076.
//
// Solidity: function isWhitelistOnlyPermissionSet() constant returns(bool)
func (_PassportLogicContract *PassportLogicContractSession) IsWhitelistOnlyPermissionSet() (bool, error) {
	return _PassportLogicContract.Contract.IsWhitelistOnlyPermissionSet(&_PassportLogicContract.CallOpts)
}

// IsWhitelistOnlyPermissionSet is a free data retrieval call binding the contract method 0x61a2d076.
//
// Solidity: function isWhitelistOnlyPermissionSet() constant returns(bool)
func (_PassportLogicContract *PassportLogicContractCallerSession) IsWhitelistOnlyPermissionSet() (bool, error) {
	return _PassportLogicContract.Contract.IsWhitelistOnlyPermissionSet(&_PassportLogicContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PassportLogicContract *PassportLogicContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PassportLogicContract.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PassportLogicContract *PassportLogicContractSession) Owner() (common.Address, error) {
	return _PassportLogicContract.Contract.Owner(&_PassportLogicContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PassportLogicContract *PassportLogicContractCallerSession) Owner() (common.Address, error) {
	return _PassportLogicContract.Contract.Owner(&_PassportLogicContract.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() constant returns(address)
func (_PassportLogicContract *PassportLogicContractCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PassportLogicContract.contract.Call(opts, out, "pendingOwner")
	return *ret0, err
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() constant returns(address)
func (_PassportLogicContract *PassportLogicContractSession) PendingOwner() (common.Address, error) {
	return _PassportLogicContract.Contract.PendingOwner(&_PassportLogicContract.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() constant returns(address)
func (_PassportLogicContract *PassportLogicContractCallerSession) PendingOwner() (common.Address, error) {
	return _PassportLogicContract.Contract.PendingOwner(&_PassportLogicContract.CallOpts)
}

// AddFactProviderToWhitelist is a paid mutator transaction binding the contract method 0xac0aeb1a.
//
// Solidity: function addFactProviderToWhitelist(_address address) returns()
func (_PassportLogicContract *PassportLogicContractTransactor) AddFactProviderToWhitelist(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _PassportLogicContract.contract.Transact(opts, "addFactProviderToWhitelist", _address)
}

// AddFactProviderToWhitelist is a paid mutator transaction binding the contract method 0xac0aeb1a.
//
// Solidity: function addFactProviderToWhitelist(_address address) returns()
func (_PassportLogicContract *PassportLogicContractSession) AddFactProviderToWhitelist(_address common.Address) (*types.Transaction, error) {
	return _PassportLogicContract.Contract.AddFactProviderToWhitelist(&_PassportLogicContract.TransactOpts, _address)
}

// AddFactProviderToWhitelist is a paid mutator transaction binding the contract method 0xac0aeb1a.
//
// Solidity: function addFactProviderToWhitelist(_address address) returns()
func (_PassportLogicContract *PassportLogicContractTransactorSession) AddFactProviderToWhitelist(_address common.Address) (*types.Transaction, error) {
	return _PassportLogicContract.Contract.AddFactProviderToWhitelist(&_PassportLogicContract.TransactOpts, _address)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_PassportLogicContract *PassportLogicContractTransactor) ClaimOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PassportLogicContract.contract.Transact(opts, "claimOwnership")
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_PassportLogicContract *PassportLogicContractSession) ClaimOwnership() (*types.Transaction, error) {
	return _PassportLogicContract.Contract.ClaimOwnership(&_PassportLogicContract.TransactOpts)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_PassportLogicContract *PassportLogicContractTransactorSession) ClaimOwnership() (*types.Transaction, error) {
	return _PassportLogicContract.Contract.ClaimOwnership(&_PassportLogicContract.TransactOpts)
}

// DeleteAddress is a paid mutator transaction binding the contract method 0x0e14a376.
//
// Solidity: function deleteAddress(_key bytes32) returns()
func (_PassportLogicContract *PassportLogicContractTransactor) DeleteAddress(opts *bind.TransactOpts, _key [32]byte) (*types.Transaction, error) {
	return _PassportLogicContract.contract.Transact(opts, "deleteAddress", _key)
}

// DeleteAddress is a paid mutator transaction binding the contract method 0x0e14a376.
//
// Solidity: function deleteAddress(_key bytes32) returns()
func (_PassportLogicContract *PassportLogicContractSession) DeleteAddress(_key [32]byte) (*types.Transaction, error) {
	return _PassportLogicContract.Contract.DeleteAddress(&_PassportLogicContract.TransactOpts, _key)
}

// DeleteAddress is a paid mutator transaction binding the contract method 0x0e14a376.
//
// Solidity: function deleteAddress(_key bytes32) returns()
func (_PassportLogicContract *PassportLogicContractTransactorSession) DeleteAddress(_key [32]byte) (*types.Transaction, error) {
	return _PassportLogicContract.Contract.DeleteAddress(&_PassportLogicContract.TransactOpts, _key)
}

// DeleteBool is a paid mutator transaction binding the contract method 0x2c62ff2d.
//
// Solidity: function deleteBool(_key bytes32) returns()
func (_PassportLogicContract *PassportLogicContractTransactor) DeleteBool(opts *bind.TransactOpts, _key [32]byte) (*types.Transaction, error) {
	return _PassportLogicContract.contract.Transact(opts, "deleteBool", _key)
}

// DeleteBool is a paid mutator transaction binding the contract method 0x2c62ff2d.
//
// Solidity: function deleteBool(_key bytes32) returns()
func (_PassportLogicContract *PassportLogicContractSession) DeleteBool(_key [32]byte) (*types.Transaction, error) {
	return _PassportLogicContract.Contract.DeleteBool(&_PassportLogicContract.TransactOpts, _key)
}

// DeleteBool is a paid mutator transaction binding the contract method 0x2c62ff2d.
//
// Solidity: function deleteBool(_key bytes32) returns()
func (_PassportLogicContract *PassportLogicContractTransactorSession) DeleteBool(_key [32]byte) (*types.Transaction, error) {
	return _PassportLogicContract.Contract.DeleteBool(&_PassportLogicContract.TransactOpts, _key)
}

// DeleteBytes is a paid mutator transaction binding the contract method 0x616b59f6.
//
// Solidity: function deleteBytes(_key bytes32) returns()
func (_PassportLogicContract *PassportLogicContractTransactor) DeleteBytes(opts *bind.TransactOpts, _key [32]byte) (*types.Transaction, error) {
	return _PassportLogicContract.contract.Transact(opts, "deleteBytes", _key)
}

// DeleteBytes is a paid mutator transaction binding the contract method 0x616b59f6.
//
// Solidity: function deleteBytes(_key bytes32) returns()
func (_PassportLogicContract *PassportLogicContractSession) DeleteBytes(_key [32]byte) (*types.Transaction, error) {
	return _PassportLogicContract.Contract.DeleteBytes(&_PassportLogicContract.TransactOpts, _key)
}

// DeleteBytes is a paid mutator transaction binding the contract method 0x616b59f6.
//
// Solidity: function deleteBytes(_key bytes32) returns()
func (_PassportLogicContract *PassportLogicContractTransactorSession) DeleteBytes(_key [32]byte) (*types.Transaction, error) {
	return _PassportLogicContract.Contract.DeleteBytes(&_PassportLogicContract.TransactOpts, _key)
}

// DeleteInt is a paid mutator transaction binding the contract method 0x8c160095.
//
// Solidity: function deleteInt(_key bytes32) returns()
func (_PassportLogicContract *PassportLogicContractTransactor) DeleteInt(opts *bind.TransactOpts, _key [32]byte) (*types.Transaction, error) {
	return _PassportLogicContract.contract.Transact(opts, "deleteInt", _key)
}

// DeleteInt is a paid mutator transaction binding the contract method 0x8c160095.
//
// Solidity: function deleteInt(_key bytes32) returns()
func (_PassportLogicContract *PassportLogicContractSession) DeleteInt(_key [32]byte) (*types.Transaction, error) {
	return _PassportLogicContract.Contract.DeleteInt(&_PassportLogicContract.TransactOpts, _key)
}

// DeleteInt is a paid mutator transaction binding the contract method 0x8c160095.
//
// Solidity: function deleteInt(_key bytes32) returns()
func (_PassportLogicContract *PassportLogicContractTransactorSession) DeleteInt(_key [32]byte) (*types.Transaction, error) {
	return _PassportLogicContract.Contract.DeleteInt(&_PassportLogicContract.TransactOpts, _key)
}

// DeleteString is a paid mutator transaction binding the contract method 0xf6bb3cc4.
//
// Solidity: function deleteString(_key bytes32) returns()
func (_PassportLogicContract *PassportLogicContractTransactor) DeleteString(opts *bind.TransactOpts, _key [32]byte) (*types.Transaction, error) {
	return _PassportLogicContract.contract.Transact(opts, "deleteString", _key)
}

// DeleteString is a paid mutator transaction binding the contract method 0xf6bb3cc4.
//
// Solidity: function deleteString(_key bytes32) returns()
func (_PassportLogicContract *PassportLogicContractSession) DeleteString(_key [32]byte) (*types.Transaction, error) {
	return _PassportLogicContract.Contract.DeleteString(&_PassportLogicContract.TransactOpts, _key)
}

// DeleteString is a paid mutator transaction binding the contract method 0xf6bb3cc4.
//
// Solidity: function deleteString(_key bytes32) returns()
func (_PassportLogicContract *PassportLogicContractTransactorSession) DeleteString(_key [32]byte) (*types.Transaction, error) {
	return _PassportLogicContract.Contract.DeleteString(&_PassportLogicContract.TransactOpts, _key)
}

// DeleteTxDataBlockNumber is a paid mutator transaction binding the contract method 0xa2b6cbe1.
//
// Solidity: function deleteTxDataBlockNumber(_key bytes32) returns()
func (_PassportLogicContract *PassportLogicContractTransactor) DeleteTxDataBlockNumber(opts *bind.TransactOpts, _key [32]byte) (*types.Transaction, error) {
	return _PassportLogicContract.contract.Transact(opts, "deleteTxDataBlockNumber", _key)
}

// DeleteTxDataBlockNumber is a paid mutator transaction binding the contract method 0xa2b6cbe1.
//
// Solidity: function deleteTxDataBlockNumber(_key bytes32) returns()
func (_PassportLogicContract *PassportLogicContractSession) DeleteTxDataBlockNumber(_key [32]byte) (*types.Transaction, error) {
	return _PassportLogicContract.Contract.DeleteTxDataBlockNumber(&_PassportLogicContract.TransactOpts, _key)
}

// DeleteTxDataBlockNumber is a paid mutator transaction binding the contract method 0xa2b6cbe1.
//
// Solidity: function deleteTxDataBlockNumber(_key bytes32) returns()
func (_PassportLogicContract *PassportLogicContractTransactorSession) DeleteTxDataBlockNumber(_key [32]byte) (*types.Transaction, error) {
	return _PassportLogicContract.Contract.DeleteTxDataBlockNumber(&_PassportLogicContract.TransactOpts, _key)
}

// DeleteUint is a paid mutator transaction binding the contract method 0xe2b202bf.
//
// Solidity: function deleteUint(_key bytes32) returns()
func (_PassportLogicContract *PassportLogicContractTransactor) DeleteUint(opts *bind.TransactOpts, _key [32]byte) (*types.Transaction, error) {
	return _PassportLogicContract.contract.Transact(opts, "deleteUint", _key)
}

// DeleteUint is a paid mutator transaction binding the contract method 0xe2b202bf.
//
// Solidity: function deleteUint(_key bytes32) returns()
func (_PassportLogicContract *PassportLogicContractSession) DeleteUint(_key [32]byte) (*types.Transaction, error) {
	return _PassportLogicContract.Contract.DeleteUint(&_PassportLogicContract.TransactOpts, _key)
}

// DeleteUint is a paid mutator transaction binding the contract method 0xe2b202bf.
//
// Solidity: function deleteUint(_key bytes32) returns()
func (_PassportLogicContract *PassportLogicContractTransactorSession) DeleteUint(_key [32]byte) (*types.Transaction, error) {
	return _PassportLogicContract.Contract.DeleteUint(&_PassportLogicContract.TransactOpts, _key)
}

// RemoveFactProviderFromWhitelist is a paid mutator transaction binding the contract method 0x6d4e7802.
//
// Solidity: function removeFactProviderFromWhitelist(_address address) returns()
func (_PassportLogicContract *PassportLogicContractTransactor) RemoveFactProviderFromWhitelist(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _PassportLogicContract.contract.Transact(opts, "removeFactProviderFromWhitelist", _address)
}

// RemoveFactProviderFromWhitelist is a paid mutator transaction binding the contract method 0x6d4e7802.
//
// Solidity: function removeFactProviderFromWhitelist(_address address) returns()
func (_PassportLogicContract *PassportLogicContractSession) RemoveFactProviderFromWhitelist(_address common.Address) (*types.Transaction, error) {
	return _PassportLogicContract.Contract.RemoveFactProviderFromWhitelist(&_PassportLogicContract.TransactOpts, _address)
}

// RemoveFactProviderFromWhitelist is a paid mutator transaction binding the contract method 0x6d4e7802.
//
// Solidity: function removeFactProviderFromWhitelist(_address address) returns()
func (_PassportLogicContract *PassportLogicContractTransactorSession) RemoveFactProviderFromWhitelist(_address common.Address) (*types.Transaction, error) {
	return _PassportLogicContract.Contract.RemoveFactProviderFromWhitelist(&_PassportLogicContract.TransactOpts, _address)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PassportLogicContract *PassportLogicContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PassportLogicContract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PassportLogicContract *PassportLogicContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _PassportLogicContract.Contract.RenounceOwnership(&_PassportLogicContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PassportLogicContract *PassportLogicContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PassportLogicContract.Contract.RenounceOwnership(&_PassportLogicContract.TransactOpts)
}

// SetAddress is a paid mutator transaction binding the contract method 0xca446dd9.
//
// Solidity: function setAddress(_key bytes32, _value address) returns()
func (_PassportLogicContract *PassportLogicContractTransactor) SetAddress(opts *bind.TransactOpts, _key [32]byte, _value common.Address) (*types.Transaction, error) {
	return _PassportLogicContract.contract.Transact(opts, "setAddress", _key, _value)
}

// SetAddress is a paid mutator transaction binding the contract method 0xca446dd9.
//
// Solidity: function setAddress(_key bytes32, _value address) returns()
func (_PassportLogicContract *PassportLogicContractSession) SetAddress(_key [32]byte, _value common.Address) (*types.Transaction, error) {
	return _PassportLogicContract.Contract.SetAddress(&_PassportLogicContract.TransactOpts, _key, _value)
}

// SetAddress is a paid mutator transaction binding the contract method 0xca446dd9.
//
// Solidity: function setAddress(_key bytes32, _value address) returns()
func (_PassportLogicContract *PassportLogicContractTransactorSession) SetAddress(_key [32]byte, _value common.Address) (*types.Transaction, error) {
	return _PassportLogicContract.Contract.SetAddress(&_PassportLogicContract.TransactOpts, _key, _value)
}

// SetBool is a paid mutator transaction binding the contract method 0xabfdcced.
//
// Solidity: function setBool(_key bytes32, _value bool) returns()
func (_PassportLogicContract *PassportLogicContractTransactor) SetBool(opts *bind.TransactOpts, _key [32]byte, _value bool) (*types.Transaction, error) {
	return _PassportLogicContract.contract.Transact(opts, "setBool", _key, _value)
}

// SetBool is a paid mutator transaction binding the contract method 0xabfdcced.
//
// Solidity: function setBool(_key bytes32, _value bool) returns()
func (_PassportLogicContract *PassportLogicContractSession) SetBool(_key [32]byte, _value bool) (*types.Transaction, error) {
	return _PassportLogicContract.Contract.SetBool(&_PassportLogicContract.TransactOpts, _key, _value)
}

// SetBool is a paid mutator transaction binding the contract method 0xabfdcced.
//
// Solidity: function setBool(_key bytes32, _value bool) returns()
func (_PassportLogicContract *PassportLogicContractTransactorSession) SetBool(_key [32]byte, _value bool) (*types.Transaction, error) {
	return _PassportLogicContract.Contract.SetBool(&_PassportLogicContract.TransactOpts, _key, _value)
}

// SetBytes is a paid mutator transaction binding the contract method 0x2e28d084.
//
// Solidity: function setBytes(_key bytes32, _value bytes) returns()
func (_PassportLogicContract *PassportLogicContractTransactor) SetBytes(opts *bind.TransactOpts, _key [32]byte, _value []byte) (*types.Transaction, error) {
	return _PassportLogicContract.contract.Transact(opts, "setBytes", _key, _value)
}

// SetBytes is a paid mutator transaction binding the contract method 0x2e28d084.
//
// Solidity: function setBytes(_key bytes32, _value bytes) returns()
func (_PassportLogicContract *PassportLogicContractSession) SetBytes(_key [32]byte, _value []byte) (*types.Transaction, error) {
	return _PassportLogicContract.Contract.SetBytes(&_PassportLogicContract.TransactOpts, _key, _value)
}

// SetBytes is a paid mutator transaction binding the contract method 0x2e28d084.
//
// Solidity: function setBytes(_key bytes32, _value bytes) returns()
func (_PassportLogicContract *PassportLogicContractTransactorSession) SetBytes(_key [32]byte, _value []byte) (*types.Transaction, error) {
	return _PassportLogicContract.Contract.SetBytes(&_PassportLogicContract.TransactOpts, _key, _value)
}

// SetInt is a paid mutator transaction binding the contract method 0x3e49bed0.
//
// Solidity: function setInt(_key bytes32, _value int256) returns()
func (_PassportLogicContract *PassportLogicContractTransactor) SetInt(opts *bind.TransactOpts, _key [32]byte, _value *big.Int) (*types.Transaction, error) {
	return _PassportLogicContract.contract.Transact(opts, "setInt", _key, _value)
}

// SetInt is a paid mutator transaction binding the contract method 0x3e49bed0.
//
// Solidity: function setInt(_key bytes32, _value int256) returns()
func (_PassportLogicContract *PassportLogicContractSession) SetInt(_key [32]byte, _value *big.Int) (*types.Transaction, error) {
	return _PassportLogicContract.Contract.SetInt(&_PassportLogicContract.TransactOpts, _key, _value)
}

// SetInt is a paid mutator transaction binding the contract method 0x3e49bed0.
//
// Solidity: function setInt(_key bytes32, _value int256) returns()
func (_PassportLogicContract *PassportLogicContractTransactorSession) SetInt(_key [32]byte, _value *big.Int) (*types.Transaction, error) {
	return _PassportLogicContract.Contract.SetInt(&_PassportLogicContract.TransactOpts, _key, _value)
}

// SetString is a paid mutator transaction binding the contract method 0x6e899550.
//
// Solidity: function setString(_key bytes32, _value string) returns()
func (_PassportLogicContract *PassportLogicContractTransactor) SetString(opts *bind.TransactOpts, _key [32]byte, _value string) (*types.Transaction, error) {
	return _PassportLogicContract.contract.Transact(opts, "setString", _key, _value)
}

// SetString is a paid mutator transaction binding the contract method 0x6e899550.
//
// Solidity: function setString(_key bytes32, _value string) returns()
func (_PassportLogicContract *PassportLogicContractSession) SetString(_key [32]byte, _value string) (*types.Transaction, error) {
	return _PassportLogicContract.Contract.SetString(&_PassportLogicContract.TransactOpts, _key, _value)
}

// SetString is a paid mutator transaction binding the contract method 0x6e899550.
//
// Solidity: function setString(_key bytes32, _value string) returns()
func (_PassportLogicContract *PassportLogicContractTransactorSession) SetString(_key [32]byte, _value string) (*types.Transaction, error) {
	return _PassportLogicContract.Contract.SetString(&_PassportLogicContract.TransactOpts, _key, _value)
}

// SetTxDataBlockNumber is a paid mutator transaction binding the contract method 0x5b2a372d.
//
// Solidity: function setTxDataBlockNumber(_key bytes32, _data bytes) returns()
func (_PassportLogicContract *PassportLogicContractTransactor) SetTxDataBlockNumber(opts *bind.TransactOpts, _key [32]byte, _data []byte) (*types.Transaction, error) {
	return _PassportLogicContract.contract.Transact(opts, "setTxDataBlockNumber", _key, _data)
}

// SetTxDataBlockNumber is a paid mutator transaction binding the contract method 0x5b2a372d.
//
// Solidity: function setTxDataBlockNumber(_key bytes32, _data bytes) returns()
func (_PassportLogicContract *PassportLogicContractSession) SetTxDataBlockNumber(_key [32]byte, _data []byte) (*types.Transaction, error) {
	return _PassportLogicContract.Contract.SetTxDataBlockNumber(&_PassportLogicContract.TransactOpts, _key, _data)
}

// SetTxDataBlockNumber is a paid mutator transaction binding the contract method 0x5b2a372d.
//
// Solidity: function setTxDataBlockNumber(_key bytes32, _data bytes) returns()
func (_PassportLogicContract *PassportLogicContractTransactorSession) SetTxDataBlockNumber(_key [32]byte, _data []byte) (*types.Transaction, error) {
	return _PassportLogicContract.Contract.SetTxDataBlockNumber(&_PassportLogicContract.TransactOpts, _key, _data)
}

// SetUint is a paid mutator transaction binding the contract method 0xe2a4853a.
//
// Solidity: function setUint(_key bytes32, _value uint256) returns()
func (_PassportLogicContract *PassportLogicContractTransactor) SetUint(opts *bind.TransactOpts, _key [32]byte, _value *big.Int) (*types.Transaction, error) {
	return _PassportLogicContract.contract.Transact(opts, "setUint", _key, _value)
}

// SetUint is a paid mutator transaction binding the contract method 0xe2a4853a.
//
// Solidity: function setUint(_key bytes32, _value uint256) returns()
func (_PassportLogicContract *PassportLogicContractSession) SetUint(_key [32]byte, _value *big.Int) (*types.Transaction, error) {
	return _PassportLogicContract.Contract.SetUint(&_PassportLogicContract.TransactOpts, _key, _value)
}

// SetUint is a paid mutator transaction binding the contract method 0xe2a4853a.
//
// Solidity: function setUint(_key bytes32, _value uint256) returns()
func (_PassportLogicContract *PassportLogicContractTransactorSession) SetUint(_key [32]byte, _value *big.Int) (*types.Transaction, error) {
	return _PassportLogicContract.Contract.SetUint(&_PassportLogicContract.TransactOpts, _key, _value)
}

// SetWhitelistOnlyPermission is a paid mutator transaction binding the contract method 0x2b42a89a.
//
// Solidity: function setWhitelistOnlyPermission(_onlyWhitelist bool) returns()
func (_PassportLogicContract *PassportLogicContractTransactor) SetWhitelistOnlyPermission(opts *bind.TransactOpts, _onlyWhitelist bool) (*types.Transaction, error) {
	return _PassportLogicContract.contract.Transact(opts, "setWhitelistOnlyPermission", _onlyWhitelist)
}

// SetWhitelistOnlyPermission is a paid mutator transaction binding the contract method 0x2b42a89a.
//
// Solidity: function setWhitelistOnlyPermission(_onlyWhitelist bool) returns()
func (_PassportLogicContract *PassportLogicContractSession) SetWhitelistOnlyPermission(_onlyWhitelist bool) (*types.Transaction, error) {
	return _PassportLogicContract.Contract.SetWhitelistOnlyPermission(&_PassportLogicContract.TransactOpts, _onlyWhitelist)
}

// SetWhitelistOnlyPermission is a paid mutator transaction binding the contract method 0x2b42a89a.
//
// Solidity: function setWhitelistOnlyPermission(_onlyWhitelist bool) returns()
func (_PassportLogicContract *PassportLogicContractTransactorSession) SetWhitelistOnlyPermission(_onlyWhitelist bool) (*types.Transaction, error) {
	return _PassportLogicContract.Contract.SetWhitelistOnlyPermission(&_PassportLogicContract.TransactOpts, _onlyWhitelist)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_PassportLogicContract *PassportLogicContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PassportLogicContract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_PassportLogicContract *PassportLogicContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PassportLogicContract.Contract.TransferOwnership(&_PassportLogicContract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_PassportLogicContract *PassportLogicContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PassportLogicContract.Contract.TransferOwnership(&_PassportLogicContract.TransactOpts, newOwner)
}

// PassportLogicContractAddressDeletedIterator is returned from FilterAddressDeleted and is used to iterate over the raw logs and unpacked data for AddressDeleted events raised by the PassportLogicContract contract.
type PassportLogicContractAddressDeletedIterator struct {
	Event *PassportLogicContractAddressDeleted // Event containing the contract specifics and raw log

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
func (it *PassportLogicContractAddressDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PassportLogicContractAddressDeleted)
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
		it.Event = new(PassportLogicContractAddressDeleted)
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
func (it *PassportLogicContractAddressDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PassportLogicContractAddressDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PassportLogicContractAddressDeleted represents a AddressDeleted event raised by the PassportLogicContract contract.
type PassportLogicContractAddressDeleted struct {
	FactProvider common.Address
	Key          [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAddressDeleted is a free log retrieval operation binding the contract event 0xed9474c4702ee9da4d0d2085840fafed5912b8e13daf1d6ea341ebe98803aec2.
//
// Solidity: e AddressDeleted(factProvider indexed address, key indexed bytes32)
func (_PassportLogicContract *PassportLogicContractFilterer) FilterAddressDeleted(opts *bind.FilterOpts, factProvider []common.Address, key [][32]byte) (*PassportLogicContractAddressDeletedIterator, error) {

	var factProviderRule []interface{}
	for _, factProviderItem := range factProvider {
		factProviderRule = append(factProviderRule, factProviderItem)
	}
	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _PassportLogicContract.contract.FilterLogs(opts, "AddressDeleted", factProviderRule, keyRule)
	if err != nil {
		return nil, err
	}
	return &PassportLogicContractAddressDeletedIterator{contract: _PassportLogicContract.contract, event: "AddressDeleted", logs: logs, sub: sub}, nil
}

// WatchAddressDeleted is a free log subscription operation binding the contract event 0xed9474c4702ee9da4d0d2085840fafed5912b8e13daf1d6ea341ebe98803aec2.
//
// Solidity: e AddressDeleted(factProvider indexed address, key indexed bytes32)
func (_PassportLogicContract *PassportLogicContractFilterer) WatchAddressDeleted(opts *bind.WatchOpts, sink chan<- *PassportLogicContractAddressDeleted, factProvider []common.Address, key [][32]byte) (event.Subscription, error) {

	var factProviderRule []interface{}
	for _, factProviderItem := range factProvider {
		factProviderRule = append(factProviderRule, factProviderItem)
	}
	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _PassportLogicContract.contract.WatchLogs(opts, "AddressDeleted", factProviderRule, keyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PassportLogicContractAddressDeleted)
				if err := _PassportLogicContract.contract.UnpackLog(event, "AddressDeleted", log); err != nil {
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

// PassportLogicContractAddressUpdatedIterator is returned from FilterAddressUpdated and is used to iterate over the raw logs and unpacked data for AddressUpdated events raised by the PassportLogicContract contract.
type PassportLogicContractAddressUpdatedIterator struct {
	Event *PassportLogicContractAddressUpdated // Event containing the contract specifics and raw log

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
func (it *PassportLogicContractAddressUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PassportLogicContractAddressUpdated)
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
		it.Event = new(PassportLogicContractAddressUpdated)
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
func (it *PassportLogicContractAddressUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PassportLogicContractAddressUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PassportLogicContractAddressUpdated represents a AddressUpdated event raised by the PassportLogicContract contract.
type PassportLogicContractAddressUpdated struct {
	FactProvider common.Address
	Key          [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAddressUpdated is a free log retrieval operation binding the contract event 0x8e7e6ab6c4613205e833e1faf5415d78dd10130a8828f729c08036e2a7a62770.
//
// Solidity: e AddressUpdated(factProvider indexed address, key indexed bytes32)
func (_PassportLogicContract *PassportLogicContractFilterer) FilterAddressUpdated(opts *bind.FilterOpts, factProvider []common.Address, key [][32]byte) (*PassportLogicContractAddressUpdatedIterator, error) {

	var factProviderRule []interface{}
	for _, factProviderItem := range factProvider {
		factProviderRule = append(factProviderRule, factProviderItem)
	}
	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _PassportLogicContract.contract.FilterLogs(opts, "AddressUpdated", factProviderRule, keyRule)
	if err != nil {
		return nil, err
	}
	return &PassportLogicContractAddressUpdatedIterator{contract: _PassportLogicContract.contract, event: "AddressUpdated", logs: logs, sub: sub}, nil
}

// WatchAddressUpdated is a free log subscription operation binding the contract event 0x8e7e6ab6c4613205e833e1faf5415d78dd10130a8828f729c08036e2a7a62770.
//
// Solidity: e AddressUpdated(factProvider indexed address, key indexed bytes32)
func (_PassportLogicContract *PassportLogicContractFilterer) WatchAddressUpdated(opts *bind.WatchOpts, sink chan<- *PassportLogicContractAddressUpdated, factProvider []common.Address, key [][32]byte) (event.Subscription, error) {

	var factProviderRule []interface{}
	for _, factProviderItem := range factProvider {
		factProviderRule = append(factProviderRule, factProviderItem)
	}
	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _PassportLogicContract.contract.WatchLogs(opts, "AddressUpdated", factProviderRule, keyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PassportLogicContractAddressUpdated)
				if err := _PassportLogicContract.contract.UnpackLog(event, "AddressUpdated", log); err != nil {
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

// PassportLogicContractBoolDeletedIterator is returned from FilterBoolDeleted and is used to iterate over the raw logs and unpacked data for BoolDeleted events raised by the PassportLogicContract contract.
type PassportLogicContractBoolDeletedIterator struct {
	Event *PassportLogicContractBoolDeleted // Event containing the contract specifics and raw log

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
func (it *PassportLogicContractBoolDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PassportLogicContractBoolDeleted)
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
		it.Event = new(PassportLogicContractBoolDeleted)
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
func (it *PassportLogicContractBoolDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PassportLogicContractBoolDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PassportLogicContractBoolDeleted represents a BoolDeleted event raised by the PassportLogicContract contract.
type PassportLogicContractBoolDeleted struct {
	FactProvider common.Address
	Key          [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBoolDeleted is a free log retrieval operation binding the contract event 0x191bd7fc1f22ce6b74b7d7e42126f0ed225b109728607c99971de9e78b0db6eb.
//
// Solidity: e BoolDeleted(factProvider indexed address, key indexed bytes32)
func (_PassportLogicContract *PassportLogicContractFilterer) FilterBoolDeleted(opts *bind.FilterOpts, factProvider []common.Address, key [][32]byte) (*PassportLogicContractBoolDeletedIterator, error) {

	var factProviderRule []interface{}
	for _, factProviderItem := range factProvider {
		factProviderRule = append(factProviderRule, factProviderItem)
	}
	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _PassportLogicContract.contract.FilterLogs(opts, "BoolDeleted", factProviderRule, keyRule)
	if err != nil {
		return nil, err
	}
	return &PassportLogicContractBoolDeletedIterator{contract: _PassportLogicContract.contract, event: "BoolDeleted", logs: logs, sub: sub}, nil
}

// WatchBoolDeleted is a free log subscription operation binding the contract event 0x191bd7fc1f22ce6b74b7d7e42126f0ed225b109728607c99971de9e78b0db6eb.
//
// Solidity: e BoolDeleted(factProvider indexed address, key indexed bytes32)
func (_PassportLogicContract *PassportLogicContractFilterer) WatchBoolDeleted(opts *bind.WatchOpts, sink chan<- *PassportLogicContractBoolDeleted, factProvider []common.Address, key [][32]byte) (event.Subscription, error) {

	var factProviderRule []interface{}
	for _, factProviderItem := range factProvider {
		factProviderRule = append(factProviderRule, factProviderItem)
	}
	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _PassportLogicContract.contract.WatchLogs(opts, "BoolDeleted", factProviderRule, keyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PassportLogicContractBoolDeleted)
				if err := _PassportLogicContract.contract.UnpackLog(event, "BoolDeleted", log); err != nil {
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

// PassportLogicContractBoolUpdatedIterator is returned from FilterBoolUpdated and is used to iterate over the raw logs and unpacked data for BoolUpdated events raised by the PassportLogicContract contract.
type PassportLogicContractBoolUpdatedIterator struct {
	Event *PassportLogicContractBoolUpdated // Event containing the contract specifics and raw log

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
func (it *PassportLogicContractBoolUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PassportLogicContractBoolUpdated)
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
		it.Event = new(PassportLogicContractBoolUpdated)
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
func (it *PassportLogicContractBoolUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PassportLogicContractBoolUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PassportLogicContractBoolUpdated represents a BoolUpdated event raised by the PassportLogicContract contract.
type PassportLogicContractBoolUpdated struct {
	FactProvider common.Address
	Key          [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBoolUpdated is a free log retrieval operation binding the contract event 0x68cc3496efaac4c1f2c0cd52da916138f6c5fc541992f05d97423a89b6914ae5.
//
// Solidity: e BoolUpdated(factProvider indexed address, key indexed bytes32)
func (_PassportLogicContract *PassportLogicContractFilterer) FilterBoolUpdated(opts *bind.FilterOpts, factProvider []common.Address, key [][32]byte) (*PassportLogicContractBoolUpdatedIterator, error) {

	var factProviderRule []interface{}
	for _, factProviderItem := range factProvider {
		factProviderRule = append(factProviderRule, factProviderItem)
	}
	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _PassportLogicContract.contract.FilterLogs(opts, "BoolUpdated", factProviderRule, keyRule)
	if err != nil {
		return nil, err
	}
	return &PassportLogicContractBoolUpdatedIterator{contract: _PassportLogicContract.contract, event: "BoolUpdated", logs: logs, sub: sub}, nil
}

// WatchBoolUpdated is a free log subscription operation binding the contract event 0x68cc3496efaac4c1f2c0cd52da916138f6c5fc541992f05d97423a89b6914ae5.
//
// Solidity: e BoolUpdated(factProvider indexed address, key indexed bytes32)
func (_PassportLogicContract *PassportLogicContractFilterer) WatchBoolUpdated(opts *bind.WatchOpts, sink chan<- *PassportLogicContractBoolUpdated, factProvider []common.Address, key [][32]byte) (event.Subscription, error) {

	var factProviderRule []interface{}
	for _, factProviderItem := range factProvider {
		factProviderRule = append(factProviderRule, factProviderItem)
	}
	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _PassportLogicContract.contract.WatchLogs(opts, "BoolUpdated", factProviderRule, keyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PassportLogicContractBoolUpdated)
				if err := _PassportLogicContract.contract.UnpackLog(event, "BoolUpdated", log); err != nil {
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

// PassportLogicContractBytesDeletedIterator is returned from FilterBytesDeleted and is used to iterate over the raw logs and unpacked data for BytesDeleted events raised by the PassportLogicContract contract.
type PassportLogicContractBytesDeletedIterator struct {
	Event *PassportLogicContractBytesDeleted // Event containing the contract specifics and raw log

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
func (it *PassportLogicContractBytesDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PassportLogicContractBytesDeleted)
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
		it.Event = new(PassportLogicContractBytesDeleted)
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
func (it *PassportLogicContractBytesDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PassportLogicContractBytesDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PassportLogicContractBytesDeleted represents a BytesDeleted event raised by the PassportLogicContract contract.
type PassportLogicContractBytesDeleted struct {
	FactProvider common.Address
	Key          [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBytesDeleted is a free log retrieval operation binding the contract event 0x0eaf95ab41777ff88e28181b1c43418524616bcef6d37c3b216015b5548cb108.
//
// Solidity: e BytesDeleted(factProvider indexed address, key indexed bytes32)
func (_PassportLogicContract *PassportLogicContractFilterer) FilterBytesDeleted(opts *bind.FilterOpts, factProvider []common.Address, key [][32]byte) (*PassportLogicContractBytesDeletedIterator, error) {

	var factProviderRule []interface{}
	for _, factProviderItem := range factProvider {
		factProviderRule = append(factProviderRule, factProviderItem)
	}
	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _PassportLogicContract.contract.FilterLogs(opts, "BytesDeleted", factProviderRule, keyRule)
	if err != nil {
		return nil, err
	}
	return &PassportLogicContractBytesDeletedIterator{contract: _PassportLogicContract.contract, event: "BytesDeleted", logs: logs, sub: sub}, nil
}

// WatchBytesDeleted is a free log subscription operation binding the contract event 0x0eaf95ab41777ff88e28181b1c43418524616bcef6d37c3b216015b5548cb108.
//
// Solidity: e BytesDeleted(factProvider indexed address, key indexed bytes32)
func (_PassportLogicContract *PassportLogicContractFilterer) WatchBytesDeleted(opts *bind.WatchOpts, sink chan<- *PassportLogicContractBytesDeleted, factProvider []common.Address, key [][32]byte) (event.Subscription, error) {

	var factProviderRule []interface{}
	for _, factProviderItem := range factProvider {
		factProviderRule = append(factProviderRule, factProviderItem)
	}
	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _PassportLogicContract.contract.WatchLogs(opts, "BytesDeleted", factProviderRule, keyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PassportLogicContractBytesDeleted)
				if err := _PassportLogicContract.contract.UnpackLog(event, "BytesDeleted", log); err != nil {
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

// PassportLogicContractBytesUpdatedIterator is returned from FilterBytesUpdated and is used to iterate over the raw logs and unpacked data for BytesUpdated events raised by the PassportLogicContract contract.
type PassportLogicContractBytesUpdatedIterator struct {
	Event *PassportLogicContractBytesUpdated // Event containing the contract specifics and raw log

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
func (it *PassportLogicContractBytesUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PassportLogicContractBytesUpdated)
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
		it.Event = new(PassportLogicContractBytesUpdated)
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
func (it *PassportLogicContractBytesUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PassportLogicContractBytesUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PassportLogicContractBytesUpdated represents a BytesUpdated event raised by the PassportLogicContract contract.
type PassportLogicContractBytesUpdated struct {
	FactProvider common.Address
	Key          [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBytesUpdated is a free log retrieval operation binding the contract event 0xee56af66250a5cbba377132a356dc403d520776e30969594927024c44af8440b.
//
// Solidity: e BytesUpdated(factProvider indexed address, key indexed bytes32)
func (_PassportLogicContract *PassportLogicContractFilterer) FilterBytesUpdated(opts *bind.FilterOpts, factProvider []common.Address, key [][32]byte) (*PassportLogicContractBytesUpdatedIterator, error) {

	var factProviderRule []interface{}
	for _, factProviderItem := range factProvider {
		factProviderRule = append(factProviderRule, factProviderItem)
	}
	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _PassportLogicContract.contract.FilterLogs(opts, "BytesUpdated", factProviderRule, keyRule)
	if err != nil {
		return nil, err
	}
	return &PassportLogicContractBytesUpdatedIterator{contract: _PassportLogicContract.contract, event: "BytesUpdated", logs: logs, sub: sub}, nil
}

// WatchBytesUpdated is a free log subscription operation binding the contract event 0xee56af66250a5cbba377132a356dc403d520776e30969594927024c44af8440b.
//
// Solidity: e BytesUpdated(factProvider indexed address, key indexed bytes32)
func (_PassportLogicContract *PassportLogicContractFilterer) WatchBytesUpdated(opts *bind.WatchOpts, sink chan<- *PassportLogicContractBytesUpdated, factProvider []common.Address, key [][32]byte) (event.Subscription, error) {

	var factProviderRule []interface{}
	for _, factProviderItem := range factProvider {
		factProviderRule = append(factProviderRule, factProviderItem)
	}
	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _PassportLogicContract.contract.WatchLogs(opts, "BytesUpdated", factProviderRule, keyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PassportLogicContractBytesUpdated)
				if err := _PassportLogicContract.contract.UnpackLog(event, "BytesUpdated", log); err != nil {
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

// PassportLogicContractIntDeletedIterator is returned from FilterIntDeleted and is used to iterate over the raw logs and unpacked data for IntDeleted events raised by the PassportLogicContract contract.
type PassportLogicContractIntDeletedIterator struct {
	Event *PassportLogicContractIntDeleted // Event containing the contract specifics and raw log

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
func (it *PassportLogicContractIntDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PassportLogicContractIntDeleted)
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
		it.Event = new(PassportLogicContractIntDeleted)
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
func (it *PassportLogicContractIntDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PassportLogicContractIntDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PassportLogicContractIntDeleted represents a IntDeleted event raised by the PassportLogicContract contract.
type PassportLogicContractIntDeleted struct {
	FactProvider common.Address
	Key          [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterIntDeleted is a free log retrieval operation binding the contract event 0x3c3ff48e02e407eb1e78310d11b5e3f9e735263a9cafc2bcf4aa981b8ecb32a5.
//
// Solidity: e IntDeleted(factProvider indexed address, key indexed bytes32)
func (_PassportLogicContract *PassportLogicContractFilterer) FilterIntDeleted(opts *bind.FilterOpts, factProvider []common.Address, key [][32]byte) (*PassportLogicContractIntDeletedIterator, error) {

	var factProviderRule []interface{}
	for _, factProviderItem := range factProvider {
		factProviderRule = append(factProviderRule, factProviderItem)
	}
	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _PassportLogicContract.contract.FilterLogs(opts, "IntDeleted", factProviderRule, keyRule)
	if err != nil {
		return nil, err
	}
	return &PassportLogicContractIntDeletedIterator{contract: _PassportLogicContract.contract, event: "IntDeleted", logs: logs, sub: sub}, nil
}

// WatchIntDeleted is a free log subscription operation binding the contract event 0x3c3ff48e02e407eb1e78310d11b5e3f9e735263a9cafc2bcf4aa981b8ecb32a5.
//
// Solidity: e IntDeleted(factProvider indexed address, key indexed bytes32)
func (_PassportLogicContract *PassportLogicContractFilterer) WatchIntDeleted(opts *bind.WatchOpts, sink chan<- *PassportLogicContractIntDeleted, factProvider []common.Address, key [][32]byte) (event.Subscription, error) {

	var factProviderRule []interface{}
	for _, factProviderItem := range factProvider {
		factProviderRule = append(factProviderRule, factProviderItem)
	}
	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _PassportLogicContract.contract.WatchLogs(opts, "IntDeleted", factProviderRule, keyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PassportLogicContractIntDeleted)
				if err := _PassportLogicContract.contract.UnpackLog(event, "IntDeleted", log); err != nil {
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

// PassportLogicContractIntUpdatedIterator is returned from FilterIntUpdated and is used to iterate over the raw logs and unpacked data for IntUpdated events raised by the PassportLogicContract contract.
type PassportLogicContractIntUpdatedIterator struct {
	Event *PassportLogicContractIntUpdated // Event containing the contract specifics and raw log

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
func (it *PassportLogicContractIntUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PassportLogicContractIntUpdated)
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
		it.Event = new(PassportLogicContractIntUpdated)
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
func (it *PassportLogicContractIntUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PassportLogicContractIntUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PassportLogicContractIntUpdated represents a IntUpdated event raised by the PassportLogicContract contract.
type PassportLogicContractIntUpdated struct {
	FactProvider common.Address
	Key          [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterIntUpdated is a free log retrieval operation binding the contract event 0xdb4ad6a3ed2bdacbae02da45a5bb54192cdb7d72648674d13e08d2fc4c45ac66.
//
// Solidity: e IntUpdated(factProvider indexed address, key indexed bytes32)
func (_PassportLogicContract *PassportLogicContractFilterer) FilterIntUpdated(opts *bind.FilterOpts, factProvider []common.Address, key [][32]byte) (*PassportLogicContractIntUpdatedIterator, error) {

	var factProviderRule []interface{}
	for _, factProviderItem := range factProvider {
		factProviderRule = append(factProviderRule, factProviderItem)
	}
	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _PassportLogicContract.contract.FilterLogs(opts, "IntUpdated", factProviderRule, keyRule)
	if err != nil {
		return nil, err
	}
	return &PassportLogicContractIntUpdatedIterator{contract: _PassportLogicContract.contract, event: "IntUpdated", logs: logs, sub: sub}, nil
}

// WatchIntUpdated is a free log subscription operation binding the contract event 0xdb4ad6a3ed2bdacbae02da45a5bb54192cdb7d72648674d13e08d2fc4c45ac66.
//
// Solidity: e IntUpdated(factProvider indexed address, key indexed bytes32)
func (_PassportLogicContract *PassportLogicContractFilterer) WatchIntUpdated(opts *bind.WatchOpts, sink chan<- *PassportLogicContractIntUpdated, factProvider []common.Address, key [][32]byte) (event.Subscription, error) {

	var factProviderRule []interface{}
	for _, factProviderItem := range factProvider {
		factProviderRule = append(factProviderRule, factProviderItem)
	}
	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _PassportLogicContract.contract.WatchLogs(opts, "IntUpdated", factProviderRule, keyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PassportLogicContractIntUpdated)
				if err := _PassportLogicContract.contract.UnpackLog(event, "IntUpdated", log); err != nil {
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

// PassportLogicContractOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the PassportLogicContract contract.
type PassportLogicContractOwnershipRenouncedIterator struct {
	Event *PassportLogicContractOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *PassportLogicContractOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PassportLogicContractOwnershipRenounced)
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
		it.Event = new(PassportLogicContractOwnershipRenounced)
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
func (it *PassportLogicContractOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PassportLogicContractOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PassportLogicContractOwnershipRenounced represents a OwnershipRenounced event raised by the PassportLogicContract contract.
type PassportLogicContractOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_PassportLogicContract *PassportLogicContractFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*PassportLogicContractOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _PassportLogicContract.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PassportLogicContractOwnershipRenouncedIterator{contract: _PassportLogicContract.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_PassportLogicContract *PassportLogicContractFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *PassportLogicContractOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _PassportLogicContract.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PassportLogicContractOwnershipRenounced)
				if err := _PassportLogicContract.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// PassportLogicContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PassportLogicContract contract.
type PassportLogicContractOwnershipTransferredIterator struct {
	Event *PassportLogicContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PassportLogicContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PassportLogicContractOwnershipTransferred)
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
		it.Event = new(PassportLogicContractOwnershipTransferred)
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
func (it *PassportLogicContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PassportLogicContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PassportLogicContractOwnershipTransferred represents a OwnershipTransferred event raised by the PassportLogicContract contract.
type PassportLogicContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_PassportLogicContract *PassportLogicContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PassportLogicContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PassportLogicContract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PassportLogicContractOwnershipTransferredIterator{contract: _PassportLogicContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_PassportLogicContract *PassportLogicContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PassportLogicContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PassportLogicContract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PassportLogicContractOwnershipTransferred)
				if err := _PassportLogicContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// PassportLogicContractStringDeletedIterator is returned from FilterStringDeleted and is used to iterate over the raw logs and unpacked data for StringDeleted events raised by the PassportLogicContract contract.
type PassportLogicContractStringDeletedIterator struct {
	Event *PassportLogicContractStringDeleted // Event containing the contract specifics and raw log

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
func (it *PassportLogicContractStringDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PassportLogicContractStringDeleted)
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
		it.Event = new(PassportLogicContractStringDeleted)
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
func (it *PassportLogicContractStringDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PassportLogicContractStringDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PassportLogicContractStringDeleted represents a StringDeleted event raised by the PassportLogicContract contract.
type PassportLogicContractStringDeleted struct {
	FactProvider common.Address
	Key          [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterStringDeleted is a free log retrieval operation binding the contract event 0x491a0e0281af9fc88ff024cffc24db22eddd4f826cb30d84936592967dc9237c.
//
// Solidity: e StringDeleted(factProvider indexed address, key indexed bytes32)
func (_PassportLogicContract *PassportLogicContractFilterer) FilterStringDeleted(opts *bind.FilterOpts, factProvider []common.Address, key [][32]byte) (*PassportLogicContractStringDeletedIterator, error) {

	var factProviderRule []interface{}
	for _, factProviderItem := range factProvider {
		factProviderRule = append(factProviderRule, factProviderItem)
	}
	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _PassportLogicContract.contract.FilterLogs(opts, "StringDeleted", factProviderRule, keyRule)
	if err != nil {
		return nil, err
	}
	return &PassportLogicContractStringDeletedIterator{contract: _PassportLogicContract.contract, event: "StringDeleted", logs: logs, sub: sub}, nil
}

// WatchStringDeleted is a free log subscription operation binding the contract event 0x491a0e0281af9fc88ff024cffc24db22eddd4f826cb30d84936592967dc9237c.
//
// Solidity: e StringDeleted(factProvider indexed address, key indexed bytes32)
func (_PassportLogicContract *PassportLogicContractFilterer) WatchStringDeleted(opts *bind.WatchOpts, sink chan<- *PassportLogicContractStringDeleted, factProvider []common.Address, key [][32]byte) (event.Subscription, error) {

	var factProviderRule []interface{}
	for _, factProviderItem := range factProvider {
		factProviderRule = append(factProviderRule, factProviderItem)
	}
	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _PassportLogicContract.contract.WatchLogs(opts, "StringDeleted", factProviderRule, keyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PassportLogicContractStringDeleted)
				if err := _PassportLogicContract.contract.UnpackLog(event, "StringDeleted", log); err != nil {
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

// PassportLogicContractStringUpdatedIterator is returned from FilterStringUpdated and is used to iterate over the raw logs and unpacked data for StringUpdated events raised by the PassportLogicContract contract.
type PassportLogicContractStringUpdatedIterator struct {
	Event *PassportLogicContractStringUpdated // Event containing the contract specifics and raw log

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
func (it *PassportLogicContractStringUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PassportLogicContractStringUpdated)
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
		it.Event = new(PassportLogicContractStringUpdated)
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
func (it *PassportLogicContractStringUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PassportLogicContractStringUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PassportLogicContractStringUpdated represents a StringUpdated event raised by the PassportLogicContract contract.
type PassportLogicContractStringUpdated struct {
	FactProvider common.Address
	Key          [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterStringUpdated is a free log retrieval operation binding the contract event 0x43e6b7e3323b4598401023341c086c07c3ff5577f594b5aab9c065f2c3c9d590.
//
// Solidity: e StringUpdated(factProvider indexed address, key indexed bytes32)
func (_PassportLogicContract *PassportLogicContractFilterer) FilterStringUpdated(opts *bind.FilterOpts, factProvider []common.Address, key [][32]byte) (*PassportLogicContractStringUpdatedIterator, error) {

	var factProviderRule []interface{}
	for _, factProviderItem := range factProvider {
		factProviderRule = append(factProviderRule, factProviderItem)
	}
	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _PassportLogicContract.contract.FilterLogs(opts, "StringUpdated", factProviderRule, keyRule)
	if err != nil {
		return nil, err
	}
	return &PassportLogicContractStringUpdatedIterator{contract: _PassportLogicContract.contract, event: "StringUpdated", logs: logs, sub: sub}, nil
}

// WatchStringUpdated is a free log subscription operation binding the contract event 0x43e6b7e3323b4598401023341c086c07c3ff5577f594b5aab9c065f2c3c9d590.
//
// Solidity: e StringUpdated(factProvider indexed address, key indexed bytes32)
func (_PassportLogicContract *PassportLogicContractFilterer) WatchStringUpdated(opts *bind.WatchOpts, sink chan<- *PassportLogicContractStringUpdated, factProvider []common.Address, key [][32]byte) (event.Subscription, error) {

	var factProviderRule []interface{}
	for _, factProviderItem := range factProvider {
		factProviderRule = append(factProviderRule, factProviderItem)
	}
	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _PassportLogicContract.contract.WatchLogs(opts, "StringUpdated", factProviderRule, keyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PassportLogicContractStringUpdated)
				if err := _PassportLogicContract.contract.UnpackLog(event, "StringUpdated", log); err != nil {
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

// PassportLogicContractTxDataDeletedIterator is returned from FilterTxDataDeleted and is used to iterate over the raw logs and unpacked data for TxDataDeleted events raised by the PassportLogicContract contract.
type PassportLogicContractTxDataDeletedIterator struct {
	Event *PassportLogicContractTxDataDeleted // Event containing the contract specifics and raw log

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
func (it *PassportLogicContractTxDataDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PassportLogicContractTxDataDeleted)
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
		it.Event = new(PassportLogicContractTxDataDeleted)
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
func (it *PassportLogicContractTxDataDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PassportLogicContractTxDataDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PassportLogicContractTxDataDeleted represents a TxDataDeleted event raised by the PassportLogicContract contract.
type PassportLogicContractTxDataDeleted struct {
	FactProvider common.Address
	Key          [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTxDataDeleted is a free log retrieval operation binding the contract event 0x9b9532c8ba5879554dee7c3e9374bd26dbfb2f79cf7f247bcf90648fc81fc164.
//
// Solidity: e TxDataDeleted(factProvider indexed address, key indexed bytes32)
func (_PassportLogicContract *PassportLogicContractFilterer) FilterTxDataDeleted(opts *bind.FilterOpts, factProvider []common.Address, key [][32]byte) (*PassportLogicContractTxDataDeletedIterator, error) {

	var factProviderRule []interface{}
	for _, factProviderItem := range factProvider {
		factProviderRule = append(factProviderRule, factProviderItem)
	}
	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _PassportLogicContract.contract.FilterLogs(opts, "TxDataDeleted", factProviderRule, keyRule)
	if err != nil {
		return nil, err
	}
	return &PassportLogicContractTxDataDeletedIterator{contract: _PassportLogicContract.contract, event: "TxDataDeleted", logs: logs, sub: sub}, nil
}

// WatchTxDataDeleted is a free log subscription operation binding the contract event 0x9b9532c8ba5879554dee7c3e9374bd26dbfb2f79cf7f247bcf90648fc81fc164.
//
// Solidity: e TxDataDeleted(factProvider indexed address, key indexed bytes32)
func (_PassportLogicContract *PassportLogicContractFilterer) WatchTxDataDeleted(opts *bind.WatchOpts, sink chan<- *PassportLogicContractTxDataDeleted, factProvider []common.Address, key [][32]byte) (event.Subscription, error) {

	var factProviderRule []interface{}
	for _, factProviderItem := range factProvider {
		factProviderRule = append(factProviderRule, factProviderItem)
	}
	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _PassportLogicContract.contract.WatchLogs(opts, "TxDataDeleted", factProviderRule, keyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PassportLogicContractTxDataDeleted)
				if err := _PassportLogicContract.contract.UnpackLog(event, "TxDataDeleted", log); err != nil {
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

// PassportLogicContractTxDataUpdatedIterator is returned from FilterTxDataUpdated and is used to iterate over the raw logs and unpacked data for TxDataUpdated events raised by the PassportLogicContract contract.
type PassportLogicContractTxDataUpdatedIterator struct {
	Event *PassportLogicContractTxDataUpdated // Event containing the contract specifics and raw log

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
func (it *PassportLogicContractTxDataUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PassportLogicContractTxDataUpdated)
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
		it.Event = new(PassportLogicContractTxDataUpdated)
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
func (it *PassportLogicContractTxDataUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PassportLogicContractTxDataUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PassportLogicContractTxDataUpdated represents a TxDataUpdated event raised by the PassportLogicContract contract.
type PassportLogicContractTxDataUpdated struct {
	FactProvider common.Address
	Key          [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTxDataUpdated is a free log retrieval operation binding the contract event 0xcbde9cd310365ff3de3ad9cb44ed688d8880c7b1d15df733a73880fb84d032de.
//
// Solidity: e TxDataUpdated(factProvider indexed address, key indexed bytes32)
func (_PassportLogicContract *PassportLogicContractFilterer) FilterTxDataUpdated(opts *bind.FilterOpts, factProvider []common.Address, key [][32]byte) (*PassportLogicContractTxDataUpdatedIterator, error) {

	var factProviderRule []interface{}
	for _, factProviderItem := range factProvider {
		factProviderRule = append(factProviderRule, factProviderItem)
	}
	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _PassportLogicContract.contract.FilterLogs(opts, "TxDataUpdated", factProviderRule, keyRule)
	if err != nil {
		return nil, err
	}
	return &PassportLogicContractTxDataUpdatedIterator{contract: _PassportLogicContract.contract, event: "TxDataUpdated", logs: logs, sub: sub}, nil
}

// WatchTxDataUpdated is a free log subscription operation binding the contract event 0xcbde9cd310365ff3de3ad9cb44ed688d8880c7b1d15df733a73880fb84d032de.
//
// Solidity: e TxDataUpdated(factProvider indexed address, key indexed bytes32)
func (_PassportLogicContract *PassportLogicContractFilterer) WatchTxDataUpdated(opts *bind.WatchOpts, sink chan<- *PassportLogicContractTxDataUpdated, factProvider []common.Address, key [][32]byte) (event.Subscription, error) {

	var factProviderRule []interface{}
	for _, factProviderItem := range factProvider {
		factProviderRule = append(factProviderRule, factProviderItem)
	}
	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _PassportLogicContract.contract.WatchLogs(opts, "TxDataUpdated", factProviderRule, keyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PassportLogicContractTxDataUpdated)
				if err := _PassportLogicContract.contract.UnpackLog(event, "TxDataUpdated", log); err != nil {
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

// PassportLogicContractUintDeletedIterator is returned from FilterUintDeleted and is used to iterate over the raw logs and unpacked data for UintDeleted events raised by the PassportLogicContract contract.
type PassportLogicContractUintDeletedIterator struct {
	Event *PassportLogicContractUintDeleted // Event containing the contract specifics and raw log

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
func (it *PassportLogicContractUintDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PassportLogicContractUintDeleted)
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
		it.Event = new(PassportLogicContractUintDeleted)
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
func (it *PassportLogicContractUintDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PassportLogicContractUintDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PassportLogicContractUintDeleted represents a UintDeleted event raised by the PassportLogicContract contract.
type PassportLogicContractUintDeleted struct {
	FactProvider common.Address
	Key          [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUintDeleted is a free log retrieval operation binding the contract event 0xd5f853f7aaba549b5a46be1de24cac4c20e716b4c603d24b2b0b7b5d97ca1c4d.
//
// Solidity: e UintDeleted(factProvider indexed address, key indexed bytes32)
func (_PassportLogicContract *PassportLogicContractFilterer) FilterUintDeleted(opts *bind.FilterOpts, factProvider []common.Address, key [][32]byte) (*PassportLogicContractUintDeletedIterator, error) {

	var factProviderRule []interface{}
	for _, factProviderItem := range factProvider {
		factProviderRule = append(factProviderRule, factProviderItem)
	}
	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _PassportLogicContract.contract.FilterLogs(opts, "UintDeleted", factProviderRule, keyRule)
	if err != nil {
		return nil, err
	}
	return &PassportLogicContractUintDeletedIterator{contract: _PassportLogicContract.contract, event: "UintDeleted", logs: logs, sub: sub}, nil
}

// WatchUintDeleted is a free log subscription operation binding the contract event 0xd5f853f7aaba549b5a46be1de24cac4c20e716b4c603d24b2b0b7b5d97ca1c4d.
//
// Solidity: e UintDeleted(factProvider indexed address, key indexed bytes32)
func (_PassportLogicContract *PassportLogicContractFilterer) WatchUintDeleted(opts *bind.WatchOpts, sink chan<- *PassportLogicContractUintDeleted, factProvider []common.Address, key [][32]byte) (event.Subscription, error) {

	var factProviderRule []interface{}
	for _, factProviderItem := range factProvider {
		factProviderRule = append(factProviderRule, factProviderItem)
	}
	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _PassportLogicContract.contract.WatchLogs(opts, "UintDeleted", factProviderRule, keyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PassportLogicContractUintDeleted)
				if err := _PassportLogicContract.contract.UnpackLog(event, "UintDeleted", log); err != nil {
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

// PassportLogicContractUintUpdatedIterator is returned from FilterUintUpdated and is used to iterate over the raw logs and unpacked data for UintUpdated events raised by the PassportLogicContract contract.
type PassportLogicContractUintUpdatedIterator struct {
	Event *PassportLogicContractUintUpdated // Event containing the contract specifics and raw log

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
func (it *PassportLogicContractUintUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PassportLogicContractUintUpdated)
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
		it.Event = new(PassportLogicContractUintUpdated)
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
func (it *PassportLogicContractUintUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PassportLogicContractUintUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PassportLogicContractUintUpdated represents a UintUpdated event raised by the PassportLogicContract contract.
type PassportLogicContractUintUpdated struct {
	FactProvider common.Address
	Key          [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUintUpdated is a free log retrieval operation binding the contract event 0xff76cb7634629f8d05011ab2a58380dfc0743157a70ba4173f866e113eca75c0.
//
// Solidity: e UintUpdated(factProvider indexed address, key indexed bytes32)
func (_PassportLogicContract *PassportLogicContractFilterer) FilterUintUpdated(opts *bind.FilterOpts, factProvider []common.Address, key [][32]byte) (*PassportLogicContractUintUpdatedIterator, error) {

	var factProviderRule []interface{}
	for _, factProviderItem := range factProvider {
		factProviderRule = append(factProviderRule, factProviderItem)
	}
	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _PassportLogicContract.contract.FilterLogs(opts, "UintUpdated", factProviderRule, keyRule)
	if err != nil {
		return nil, err
	}
	return &PassportLogicContractUintUpdatedIterator{contract: _PassportLogicContract.contract, event: "UintUpdated", logs: logs, sub: sub}, nil
}

// WatchUintUpdated is a free log subscription operation binding the contract event 0xff76cb7634629f8d05011ab2a58380dfc0743157a70ba4173f866e113eca75c0.
//
// Solidity: e UintUpdated(factProvider indexed address, key indexed bytes32)
func (_PassportLogicContract *PassportLogicContractFilterer) WatchUintUpdated(opts *bind.WatchOpts, sink chan<- *PassportLogicContractUintUpdated, factProvider []common.Address, key [][32]byte) (event.Subscription, error) {

	var factProviderRule []interface{}
	for _, factProviderItem := range factProvider {
		factProviderRule = append(factProviderRule, factProviderItem)
	}
	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _PassportLogicContract.contract.WatchLogs(opts, "UintUpdated", factProviderRule, keyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PassportLogicContractUintUpdated)
				if err := _PassportLogicContract.contract.UnpackLog(event, "UintUpdated", log); err != nil {
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

// PassportLogicContractWhitelistFactProviderAddedIterator is returned from FilterWhitelistFactProviderAdded and is used to iterate over the raw logs and unpacked data for WhitelistFactProviderAdded events raised by the PassportLogicContract contract.
type PassportLogicContractWhitelistFactProviderAddedIterator struct {
	Event *PassportLogicContractWhitelistFactProviderAdded // Event containing the contract specifics and raw log

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
func (it *PassportLogicContractWhitelistFactProviderAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PassportLogicContractWhitelistFactProviderAdded)
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
		it.Event = new(PassportLogicContractWhitelistFactProviderAdded)
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
func (it *PassportLogicContractWhitelistFactProviderAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PassportLogicContractWhitelistFactProviderAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PassportLogicContractWhitelistFactProviderAdded represents a WhitelistFactProviderAdded event raised by the PassportLogicContract contract.
type PassportLogicContractWhitelistFactProviderAdded struct {
	FactProvider common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterWhitelistFactProviderAdded is a free log retrieval operation binding the contract event 0xff8ac5f0e9cdc8f7ffbbc29ab1301ee649cf01cf7a6a9265a2d90587612c4175.
//
// Solidity: e WhitelistFactProviderAdded(factProvider indexed address)
func (_PassportLogicContract *PassportLogicContractFilterer) FilterWhitelistFactProviderAdded(opts *bind.FilterOpts, factProvider []common.Address) (*PassportLogicContractWhitelistFactProviderAddedIterator, error) {

	var factProviderRule []interface{}
	for _, factProviderItem := range factProvider {
		factProviderRule = append(factProviderRule, factProviderItem)
	}

	logs, sub, err := _PassportLogicContract.contract.FilterLogs(opts, "WhitelistFactProviderAdded", factProviderRule)
	if err != nil {
		return nil, err
	}
	return &PassportLogicContractWhitelistFactProviderAddedIterator{contract: _PassportLogicContract.contract, event: "WhitelistFactProviderAdded", logs: logs, sub: sub}, nil
}

// WatchWhitelistFactProviderAdded is a free log subscription operation binding the contract event 0xff8ac5f0e9cdc8f7ffbbc29ab1301ee649cf01cf7a6a9265a2d90587612c4175.
//
// Solidity: e WhitelistFactProviderAdded(factProvider indexed address)
func (_PassportLogicContract *PassportLogicContractFilterer) WatchWhitelistFactProviderAdded(opts *bind.WatchOpts, sink chan<- *PassportLogicContractWhitelistFactProviderAdded, factProvider []common.Address) (event.Subscription, error) {

	var factProviderRule []interface{}
	for _, factProviderItem := range factProvider {
		factProviderRule = append(factProviderRule, factProviderItem)
	}

	logs, sub, err := _PassportLogicContract.contract.WatchLogs(opts, "WhitelistFactProviderAdded", factProviderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PassportLogicContractWhitelistFactProviderAdded)
				if err := _PassportLogicContract.contract.UnpackLog(event, "WhitelistFactProviderAdded", log); err != nil {
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

// PassportLogicContractWhitelistFactProviderRemovedIterator is returned from FilterWhitelistFactProviderRemoved and is used to iterate over the raw logs and unpacked data for WhitelistFactProviderRemoved events raised by the PassportLogicContract contract.
type PassportLogicContractWhitelistFactProviderRemovedIterator struct {
	Event *PassportLogicContractWhitelistFactProviderRemoved // Event containing the contract specifics and raw log

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
func (it *PassportLogicContractWhitelistFactProviderRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PassportLogicContractWhitelistFactProviderRemoved)
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
		it.Event = new(PassportLogicContractWhitelistFactProviderRemoved)
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
func (it *PassportLogicContractWhitelistFactProviderRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PassportLogicContractWhitelistFactProviderRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PassportLogicContractWhitelistFactProviderRemoved represents a WhitelistFactProviderRemoved event raised by the PassportLogicContract contract.
type PassportLogicContractWhitelistFactProviderRemoved struct {
	FactProvider common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterWhitelistFactProviderRemoved is a free log retrieval operation binding the contract event 0x5bf47f016603dfff957e6086bd9a350212686e8f1d9208b3db47c5a3107c7878.
//
// Solidity: e WhitelistFactProviderRemoved(factProvider indexed address)
func (_PassportLogicContract *PassportLogicContractFilterer) FilterWhitelistFactProviderRemoved(opts *bind.FilterOpts, factProvider []common.Address) (*PassportLogicContractWhitelistFactProviderRemovedIterator, error) {

	var factProviderRule []interface{}
	for _, factProviderItem := range factProvider {
		factProviderRule = append(factProviderRule, factProviderItem)
	}

	logs, sub, err := _PassportLogicContract.contract.FilterLogs(opts, "WhitelistFactProviderRemoved", factProviderRule)
	if err != nil {
		return nil, err
	}
	return &PassportLogicContractWhitelistFactProviderRemovedIterator{contract: _PassportLogicContract.contract, event: "WhitelistFactProviderRemoved", logs: logs, sub: sub}, nil
}

// WatchWhitelistFactProviderRemoved is a free log subscription operation binding the contract event 0x5bf47f016603dfff957e6086bd9a350212686e8f1d9208b3db47c5a3107c7878.
//
// Solidity: e WhitelistFactProviderRemoved(factProvider indexed address)
func (_PassportLogicContract *PassportLogicContractFilterer) WatchWhitelistFactProviderRemoved(opts *bind.WatchOpts, sink chan<- *PassportLogicContractWhitelistFactProviderRemoved, factProvider []common.Address) (event.Subscription, error) {

	var factProviderRule []interface{}
	for _, factProviderItem := range factProvider {
		factProviderRule = append(factProviderRule, factProviderItem)
	}

	logs, sub, err := _PassportLogicContract.contract.WatchLogs(opts, "WhitelistFactProviderRemoved", factProviderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PassportLogicContractWhitelistFactProviderRemoved)
				if err := _PassportLogicContract.contract.UnpackLog(event, "WhitelistFactProviderRemoved", log); err != nil {
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

// PassportLogicContractWhitelistOnlyPermissionSetIterator is returned from FilterWhitelistOnlyPermissionSet and is used to iterate over the raw logs and unpacked data for WhitelistOnlyPermissionSet events raised by the PassportLogicContract contract.
type PassportLogicContractWhitelistOnlyPermissionSetIterator struct {
	Event *PassportLogicContractWhitelistOnlyPermissionSet // Event containing the contract specifics and raw log

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
func (it *PassportLogicContractWhitelistOnlyPermissionSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PassportLogicContractWhitelistOnlyPermissionSet)
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
		it.Event = new(PassportLogicContractWhitelistOnlyPermissionSet)
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
func (it *PassportLogicContractWhitelistOnlyPermissionSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PassportLogicContractWhitelistOnlyPermissionSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PassportLogicContractWhitelistOnlyPermissionSet represents a WhitelistOnlyPermissionSet event raised by the PassportLogicContract contract.
type PassportLogicContractWhitelistOnlyPermissionSet struct {
	OnlyWhitelist bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterWhitelistOnlyPermissionSet is a free log retrieval operation binding the contract event 0x6e1eb3fa845047cf9d5fd1deb1b7770b4e26163dd4d88cefb14fe7964a700595.
//
// Solidity: e WhitelistOnlyPermissionSet(onlyWhitelist indexed bool)
func (_PassportLogicContract *PassportLogicContractFilterer) FilterWhitelistOnlyPermissionSet(opts *bind.FilterOpts, onlyWhitelist []bool) (*PassportLogicContractWhitelistOnlyPermissionSetIterator, error) {

	var onlyWhitelistRule []interface{}
	for _, onlyWhitelistItem := range onlyWhitelist {
		onlyWhitelistRule = append(onlyWhitelistRule, onlyWhitelistItem)
	}

	logs, sub, err := _PassportLogicContract.contract.FilterLogs(opts, "WhitelistOnlyPermissionSet", onlyWhitelistRule)
	if err != nil {
		return nil, err
	}
	return &PassportLogicContractWhitelistOnlyPermissionSetIterator{contract: _PassportLogicContract.contract, event: "WhitelistOnlyPermissionSet", logs: logs, sub: sub}, nil
}

// WatchWhitelistOnlyPermissionSet is a free log subscription operation binding the contract event 0x6e1eb3fa845047cf9d5fd1deb1b7770b4e26163dd4d88cefb14fe7964a700595.
//
// Solidity: e WhitelistOnlyPermissionSet(onlyWhitelist indexed bool)
func (_PassportLogicContract *PassportLogicContractFilterer) WatchWhitelistOnlyPermissionSet(opts *bind.WatchOpts, sink chan<- *PassportLogicContractWhitelistOnlyPermissionSet, onlyWhitelist []bool) (event.Subscription, error) {

	var onlyWhitelistRule []interface{}
	for _, onlyWhitelistItem := range onlyWhitelist {
		onlyWhitelistRule = append(onlyWhitelistRule, onlyWhitelistItem)
	}

	logs, sub, err := _PassportLogicContract.contract.WatchLogs(opts, "WhitelistOnlyPermissionSet", onlyWhitelistRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PassportLogicContractWhitelistOnlyPermissionSet)
				if err := _PassportLogicContract.contract.UnpackLog(event, "WhitelistOnlyPermissionSet", log); err != nil {
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
