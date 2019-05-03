package flag

import (
	"errors"
)

// ErrInvalidFactKeyLength error returned when fact key is not fit into 32 bytes
var ErrInvalidFactKeyLength = errors.New("the fact key string should fit into 32 bytes")

// FactKey is a fact key type
type FactKey [32]byte

// UnmarshalFlag implements flags.Unmarshaler interface
func (f *FactKey) UnmarshalFlag(value string) (err error) {
	if bs := []byte(value); len(bs) <= 32 {
		copy(f[:], bs)
	} else {
		err = ErrInvalidFactKeyLength
	}
	return
}
