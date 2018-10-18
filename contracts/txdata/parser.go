package txdata

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
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
		Key  [32]byte
		Data []byte
	}

	// SetBytesParameters holds the parameters of the setBytes method call
	SetBytesParameters struct {
		Key  [32]byte
		Data []byte
	}

	// SetStringParameters holds the parameters of the setString method call
	SetStringParameters struct {
		Key  [32]byte
		Data []byte
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

// ParseSetStringCallData parses setBytes method call parameters from transaction input data.
func (p *PassportLogicInputParser) ParseSetStringCallData(input []byte) (parms *SetStringParameters, err error) {
	v := &SetStringParameters{}
	err = unpackInput(p.abi, v, "setString", input)
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
