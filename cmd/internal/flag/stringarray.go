package flag

import (
	"strings"
)

// StringArray parses comma separated string
type StringArray []string

// UnmarshalFlag implements flags.Unmarshaler interface
func (k *StringArray) UnmarshalFlag(value string) error {
	if value == "" {
		*k = (StringArray)([]string{})
		return nil
	}

	a := strings.Split(value, ",")

	*k = (StringArray)(a)
	return nil
}

// AsStringArr returns value as string array
func (k *StringArray) AsStringArr() []string {
	return ([]string)(*k)
}
