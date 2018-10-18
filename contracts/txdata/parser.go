package txdata

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"gitlab.com/monetha/protocol-go-sdk/contracts"
)

// PassportLogicInputParser allows to parse transaction input data
type PassportLogicInputParser struct {
	abi abi.ABI
}

// NewPassportLogicInputParser creates an instance of PassportLogicInputParser
func NewPassportLogicInputParser() (*PassportLogicInputParser, error) {
	parsedABI, err := abi.JSON(strings.NewReader(contracts.PassportLogicContractABI))
	if err != nil {
		return nil, err
	}
	return &PassportLogicInputParser{parsedABI}, nil
}

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
)

// ParseSetTxDataBlockNumberCallData parses setTxDataBlockNumber method call parameters from transaction input data.
func (p *PassportLogicInputParser) ParseSetTxDataBlockNumberCallData(input []byte) (parms *SetTxDataBlockNumberParameters, err error) {
	v := &SetTxDataBlockNumberParameters{}
	err = unpackInput(p.abi, v, "setTxDataBlockNumber", input)
	if err == nil {
		parms = v
	}
	return
}

// ParseSetBytesCallData parses setBytes method call parameters from transaction input data.
func (p *PassportLogicInputParser) ParseSetBytesCallData(input []byte) (parms *SetBytesParameters, err error) {
	v := &SetBytesParameters{}
	err = unpackInput(p.abi, v, "setBytes", input)
	if err == nil {
		parms = v
	}
	return
}

// ParseSetStringCallData parses setString method call parameters from transaction input data.
func (p *PassportLogicInputParser) ParseSetStringCallData(input []byte) (parms *SetStringParameters, err error) {
	v := &SetStringParameters{}
	err = unpackInput(p.abi, v, "setString", input)
	if err == nil {
		parms = v
	}
	return
}

// ParseSetAddressCallData parses setAddress method call parameters from transaction input data.
func (p *PassportLogicInputParser) ParseSetAddressCallData(input []byte) (parms *SetAddressParameters, err error) {
	v := &SetAddressParameters{}
	err = unpackInput(p.abi, v, "setAddress", input)
	if err == nil {
		parms = v
	}
	return
}

// ParseSetUintCallData parses setUint method call parameters from transaction input data.
func (p *PassportLogicInputParser) ParseSetUintCallData(input []byte) (parms *SetUintParameters, err error) {
	v := &SetUintParameters{}
	err = unpackInput(p.abi, v, "setUint", input)
	if err == nil {
		parms = v
	}
	return
}

// ParseSetIntCallData parses setInt method call parameters from transaction input data.
func (p *PassportLogicInputParser) ParseSetIntCallData(input []byte) (parms *SetIntParameters, err error) {
	v := &SetIntParameters{}
	err = unpackInput(p.abi, v, "setInt", input)
	if err == nil {
		parms = v
	}
	return
}

// ParseSetBoolCallData parses setBool method call parameters from transaction input data.
func (p *PassportLogicInputParser) ParseSetBoolCallData(input []byte) (parms *SetBoolParameters, err error) {
	v := &SetBoolParameters{}
	err = unpackInput(p.abi, v, "setBool", input)
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
