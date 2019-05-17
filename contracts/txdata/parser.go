package txdata

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/monetha/reputation-go-sdk/contracts"
)

type (
	// SetTxDataBlockNumberParameters holds the parameters of the setTxDataBlockNumber method call
	SetTxDataBlockNumberParameters struct {
		Key  [32]byte `abi:"_key"`
		Data []byte   `abi:"_data"`
	}

	// SetBytesParameters holds the parameters of the setBytes method call
	SetBytesParameters struct {
		Key  [32]byte `abi:"_key"`
		Data []byte   `abi:"_value"`
	}

	// SetStringParameters holds the parameters of the setString method call
	SetStringParameters struct {
		Key  [32]byte `abi:"_key"`
		Data string   `abi:"_value"`
	}

	// SetAddressParameters holds the parameters of the setAddress method call
	SetAddressParameters struct {
		Key  [32]byte       `abi:"_key"`
		Data common.Address `abi:"_value"`
	}

	// SetUintParameters holds the parameters of the setUint method call
	SetUintParameters struct {
		Key  [32]byte `abi:"_key"`
		Data *big.Int `abi:"_value"`
	}

	// SetIntParameters holds the parameters of the setInt method call
	SetIntParameters struct {
		Key  [32]byte `abi:"_key"`
		Data *big.Int `abi:"_value"`
	}

	// SetBoolParameters holds the parameters of the setBool method call
	SetBoolParameters struct {
		Key  [32]byte `abi:"_key"`
		Data bool     `abi:"_value"`
	}

	// SetIPFSHashParameters holds the parameters of the setIPFSHash method call
	SetIPFSHashParameters struct {
		Key  [32]byte `abi:"_key"`
		Hash string   `abi:"_value"`
	}

	// SetPrivateDataParameters holds parameters of the setPrivateData method call
	SetPrivateDataParameters struct {
		Key          [32]byte `abi:"_key"`
		DataIPFSHash string   `abi:"_dataIPFSHash"`
		DataKeyHash  [32]byte `abi:"_dataKeyHash"`
	}
)

// ParseSetTxDataBlockNumberCallData parses setTxDataBlockNumber method call parameters from transaction input data.
func ParseSetTxDataBlockNumberCallData(input []byte) (parms *SetTxDataBlockNumberParameters, err error) {
	v := &SetTxDataBlockNumberParameters{}
	err = unpackInput(contracts.PassportLogicABI, v, "setTxDataBlockNumber", input)
	if err == nil {
		parms = v
	}
	return
}

// ParseSetBytesCallData parses setBytes method call parameters from transaction input data.
func ParseSetBytesCallData(input []byte) (parms *SetBytesParameters, err error) {
	v := &SetBytesParameters{}
	err = unpackInput(contracts.PassportLogicABI, v, "setBytes", input)
	if err == nil {
		parms = v
	}
	return
}

// ParseSetStringCallData parses setString method call parameters from transaction input data.
func ParseSetStringCallData(input []byte) (parms *SetStringParameters, err error) {
	v := &SetStringParameters{}
	err = unpackInput(contracts.PassportLogicABI, v, "setString", input)
	if err == nil {
		parms = v
	}
	return
}

// ParseSetAddressCallData parses setAddress method call parameters from transaction input data.
func ParseSetAddressCallData(input []byte) (parms *SetAddressParameters, err error) {
	v := &SetAddressParameters{}
	err = unpackInput(contracts.PassportLogicABI, v, "setAddress", input)
	if err == nil {
		parms = v
	}
	return
}

// ParseSetUintCallData parses setUint method call parameters from transaction input data.
func ParseSetUintCallData(input []byte) (parms *SetUintParameters, err error) {
	v := &SetUintParameters{}
	err = unpackInput(contracts.PassportLogicABI, v, "setUint", input)
	if err == nil {
		parms = v
	}
	return
}

// ParseSetIntCallData parses setInt method call parameters from transaction input data.
func ParseSetIntCallData(input []byte) (parms *SetIntParameters, err error) {
	v := &SetIntParameters{}
	err = unpackInput(contracts.PassportLogicABI, v, "setInt", input)
	if err == nil {
		parms = v
	}
	return
}

// ParseSetBoolCallData parses setBool method call parameters from transaction input data.
func ParseSetBoolCallData(input []byte) (parms *SetBoolParameters, err error) {
	v := &SetBoolParameters{}
	err = unpackInput(contracts.PassportLogicABI, v, "setBool", input)
	if err == nil {
		parms = v
	}
	return
}

// ParseSetIPFSHashCallData parses setIPFSHash method call parameters from transaction input data.
func ParseSetIPFSHashCallData(input []byte) (parms *SetIPFSHashParameters, err error) {
	v := &SetIPFSHashParameters{}
	err = unpackInput(contracts.PassportLogicABI, v, "setIPFSHash", input)
	if err == nil {
		parms = v
	}
	return
}

// ParseSetPrivateDataHashesCallData parses setPrivateData method call parameters from transaction input data
func ParseSetPrivateDataHashesCallData(input []byte) (parms *SetPrivateDataParameters, err error) {
	v := &SetPrivateDataParameters{}
	err = unpackInput(contracts.PassportLogicABI, v, "setPrivateDataHashes", input)
	if err == nil {
		parms = v
	}
	return
}

func unpackInput(abi abi.ABI, v interface{}, name string, input []byte) (err error) {
	if len(input) < 4 {
		return fmt.Errorf("invalid ABI-data, incomplete method signature of (%d bytes)", len(input))
	}

	sigData, argData := input[:4], input[4:]
	if len(argData)%32 != 0 {
		return fmt.Errorf("not ABI-encoded data; length should be a multiple of 32 (was %d)", len(argData))
	}

	parsedMethod, err := abi.MethodById(sigData)
	if err != nil {
		return err
	}

	if expectedMethod, ok := abi.Methods[name]; ok {
		if expectedMethod.Name != parsedMethod.Name {
			return fmt.Errorf("parsed method signature %v, expected method signature: %v", parsedMethod.Sig(), expectedMethod.Sig())
		}

		return parsedMethod.Inputs.Unpack(v, argData)
	}

	return fmt.Errorf("could not locate named method %v", name)
}
