package facts

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/monetha/go-verifiable-data/crypto/ecies"
	"github.com/pkg/errors"
)

const (
	ipfsPublicKeyFileName        = "public_key"
	ipfsEncryptedMessageFileName = "encrypted_message"
	ipfsMessageHMACFileName      = "hmac"
)

func deriveSecretKeyringMaterial(
	ec *ecies.ECIES,
	pubKey *ecdsa.PublicKey,
	passportAddress common.Address,
	factProviderAddress common.Address,
	factKey [32]byte,
) (skm *ecies.SecretKeyringMaterial, skmBytes []byte, skmHash common.Hash, err error) {
	// using [Provider Address + Passport Address + factKey] as seed to derive secret keyring material and HMAC
	var seed []byte
	for _, s := range [][]byte{
		factProviderAddress.Bytes(),
		passportAddress.Bytes(),
		factKey[:],
	} {
		seed = append(seed, s...)
	}

	skm, err = ec.DeriveSecretKeyringMaterial(pubKey, seed)
	if err != nil {
		err = errors.Wrap(err, "failed to derive secret keyring material")
		return
	}
	// take only part of MAC otherwise EncryptionKey + MACKey won't fit into 32 bytes array
	skm.MACKey = skm.MACKey[:len(skm.EncryptionKey)]

	skmBytes = make([]byte, len(skm.MACKey)+len(skm.EncryptionKey))
	copy(skmBytes, skm.EncryptionKey)
	copy(skmBytes[len(skm.EncryptionKey):], skm.MACKey)

	skmHash = crypto.Keccak256Hash(skmBytes)

	return
}

func unmarshalSecretKeyringMaterial(skmBytes []byte) *ecies.SecretKeyringMaterial {
	keyLen := len(skmBytes) / 2
	return &ecies.SecretKeyringMaterial{
		EncryptionKey: skmBytes[:keyLen],
		MACKey:        skmBytes[keyLen:],
	}
}
