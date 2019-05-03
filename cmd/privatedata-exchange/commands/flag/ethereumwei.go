package flag

import (
	"fmt"
	"math/big"
	"strings"
)

var etherInWei = big.NewInt(1000000000000000000)

// EthereumWei parses wei
type EthereumWei big.Int

// UnmarshalFlag implements flags.Unmarshaler interface
func (f *EthereumWei) UnmarshalFlag(value string) (err error) {
	if idx, ok := new(big.Int).SetString(value, 10); ok {
		if idx.Sign() != -1 {
			*f = EthereumWei(*idx)
		} else {
			err = fmt.Errorf("non negative integer expected, but got: %v", value)
		}
	} else {
		err = fmt.Errorf("failed to parse integer: %v", value)
	}
	return
}

// AsBigInt returns wei as big.Int
func (f *EthereumWei) AsBigInt() *big.Int {
	return (*big.Int)(f)
}

// EthString converts wei to ETH and returns result as string
func (f *EthereumWei) EthString() string {
	r := new(big.Rat).SetFrac(f.AsBigInt(), etherInWei)
	return strings.TrimRight(strings.TrimRight(r.FloatString(18), "0"), ".")
}
