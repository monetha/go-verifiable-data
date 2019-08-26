package flag

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

// EthereumAddress is a flag that parses Ethereum address
type EthereumAddress common.Address

// UnmarshalFlag implements flags.Unmarshaler interface
func (a *EthereumAddress) UnmarshalFlag(value string) error {
	if !common.IsHexAddress(value) {
		return fmt.Errorf("invalid Ethereum address: %v", value)
	}
	*a = EthereumAddress(common.HexToAddress(value))
	return nil
}

// AsCommonAddress returns common.Address
func (a *EthereumAddress) AsCommonAddress() common.Address {
	return common.Address(*a)
}
