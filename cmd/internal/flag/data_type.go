package flag

import (
	"fmt"
	"strings"

	"github.com/monetha/go-verifiable-data/types/data"
)

// DataType is a flag that parses data type of fact
type DataType data.Type

var (
	dataTypes      = make(map[string]data.Type)
	dataTypeSetStr string
)

func init() {
	keys := make([]string, 0, len(data.TypeValues()))

	for _, key := range data.TypeValues() {
		keyStr := strings.ToLower(key.String())
		dataTypes[keyStr] = key
		keys = append(keys, keyStr)
	}

	dataTypeSetStr = strings.Join(keys, ", ")
}

// DataTypeSet returns comma separated string of allowed data types
func DataTypeSet() string {
	return dataTypeSetStr
}

// UnmarshalFlag implements flags.Unmarshaler interface
func (a *DataType) UnmarshalFlag(value string) error {
	dataType, ok := dataTypes[strings.ToLower(value)]
	if !ok {
		return fmt.Errorf("unsupported data type of fact '%v', use one of: %v", value, dataTypeSetStr)
	}

	*a = DataType(dataType)

	return nil
}

// AsType returns data.Type
func (a *DataType) AsType() data.Type {
	return data.Type(*a)
}
