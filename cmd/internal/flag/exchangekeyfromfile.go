package flag

import (
	"errors"
	"io/ioutil"
)

// ErrInvalidExchangeKeyLength returned when the length of the exchange key is not equal to 32 bytes
var ErrInvalidExchangeKeyLength = errors.New("exchange key length must be 32 bytes")

// ExchangeKeyFromFile parses exchange key from file
type ExchangeKeyFromFile [32]byte

// UnmarshalFlag implements flags.Unmarshaler interface
func (f *ExchangeKeyFromFile) UnmarshalFlag(value string) error {
	key, err := ioutil.ReadFile(value)
	if err != nil {
		return err
	}

	if len(key) != 32 {
		return ErrInvalidExchangeKeyLength
	}

	copy(f[:], key)
	return nil
}

// AsBytes32 returns exchange key as [32]byte
func (f *ExchangeKeyFromFile) AsBytes32() [32]byte {
	return *f
}
