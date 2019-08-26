package flag

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

// TXHash is a flag that parses transaction hash
type TXHash common.Hash

// UnmarshalFlag implements flags.Unmarshaler interface
func (a *TXHash) UnmarshalFlag(value string) error {
	vCheck := value

	if hasHexPrefix(vCheck) {
		vCheck = vCheck[2:]
	}

	if !isHex(vCheck) {
		return fmt.Errorf("invalid TX hash: %v", value)
	}

	*a = TXHash(common.HexToHash(value))
	return nil
}

// AsCommonHash returns common.Hash
func (a *TXHash) AsCommonHash() common.Hash {
	return common.Hash(*a)
}

// hasHexPrefix validates str begins with '0x' or '0X'.
func hasHexPrefix(str string) bool {
	return len(str) >= 2 && str[0] == '0' && (str[1] == 'x' || str[1] == 'X')
}

// isHexCharacter returns bool of c being a valid hexadecimal.
func isHexCharacter(c byte) bool {
	return ('0' <= c && c <= '9') || ('a' <= c && c <= 'f') || ('A' <= c && c <= 'F')
}

// isHex validates whether each byte is valid hexadecimal string.
func isHex(str string) bool {
	if len(str)%2 != 0 {
		return false
	}
	for _, c := range []byte(str) {
		if !isHexCharacter(c) {
			return false
		}
	}
	return true
}
