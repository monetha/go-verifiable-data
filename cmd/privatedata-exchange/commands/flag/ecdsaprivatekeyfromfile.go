package flag

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/crypto"
)

// ECDSAPrivateKeyFromFile is a flag that reads ecdsa.PrivateKey from file
type ECDSAPrivateKeyFromFile ecdsa.PrivateKey

// UnmarshalFlag implements flags.Unmarshaler interface
func (k *ECDSAPrivateKeyFromFile) UnmarshalFlag(value string) error {
	key, err := crypto.LoadECDSA(value)
	if err != nil {
		return err
	}
	*k = (ECDSAPrivateKeyFromFile)(*key)
	return nil
}

// AsECDSAPrivateKey returns pointer to ecdsa.PrivateKey
func (k *ECDSAPrivateKeyFromFile) AsECDSAPrivateKey() *ecdsa.PrivateKey {
	return (*ecdsa.PrivateKey)(k)
}
