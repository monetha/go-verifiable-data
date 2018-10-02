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

// SetTxDataBlockNumberParameters holds the parameters of the setTxDataBlockNumber method call
type SetTxDataBlockNumberParameters struct {
	Key  [32]byte
	Data []byte
}

// ParseSetTxDataBlockNumberCallData parses setTxDataBlockNumber method call parameters from transaction input data.
func (p *PassportLogicInputParser) ParseSetTxDataBlockNumberCallData(input []byte) (parms *SetTxDataBlockNumberParameters, err error) {
	err = abiExt(p.abi).UnpackInput(parms, "setTxDataBlockNumber", input)
	return
}

type abiExt abi.ABI

// UnpackInput parses input according to the abi specification
func (abi abiExt) UnpackInput(v interface{}, name string, input []byte) (err error) {
	if len(input) <= 4 {
		return fmt.Errorf("abi: insufficient input data")
	}

	input = input[4:] // remove signature

	// since there can't be naming collisions with contracts and events,
	// we need to decide whether we're calling a method or an event
	if method, ok := abi.Methods[name]; ok {
		if len(input)%32 != 0 {
			return fmt.Errorf("abi: improperly formatted input")
		}
		return method.Inputs.Unpack(v, input)
	} else if event, ok := abi.Events[name]; ok {
		return event.Inputs.Unpack(v, input)
	}
	return fmt.Errorf("abi: could not locate named method or event")
}
