package flag

import (
	"strings"
)

// StringArray parses comma separated string
type StringArray []string

// UnmarshalFlag implements flags.Unmarshaler interface
func (k *StringArray) UnmarshalFlag(value string) error {
	if value == "" {
		*k = []string{}
	} else {
		*k = strings.Split(value, ",")
	}

	return nil
}

// AsStringArr returns value as string array
func (k *StringArray) AsStringArr() []string {
	return *k
}
