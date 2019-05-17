package flag

import (
	"fmt"
	"math/big"
)

// ExchangeIndex parses exchange index
type ExchangeIndex big.Int

// UnmarshalFlag implements flags.Unmarshaler interface
func (f *ExchangeIndex) UnmarshalFlag(value string) (err error) {
	if idx, ok := new(big.Int).SetString(value, 10); ok {
		if idx.Sign() != -1 {
			*f = ExchangeIndex(*idx)
		} else {
			err = fmt.Errorf("non negative integer expected, but got: %v", value)
		}
	} else {
		err = fmt.Errorf("failed to parse integer: %v", value)
	}
	return
}

// AsBigInt returns exchange index as big.Int
func (f *ExchangeIndex) AsBigInt() *big.Int {
	return (*big.Int)(f)
}
