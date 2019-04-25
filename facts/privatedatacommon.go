package facts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	ipfsPublicKeyFileName        = "public_key"
	ipfsEncryptedMessageFileName = "encrypted_message"
	ipfsMessageHMACFileName      = "hmac"
)

type skmSeedParams struct {
	PassportAddress     common.Address
	FactProviderAddress common.Address
	FactKey             [32]byte
}

// using [Provider Address + Passport Address + factKey] as seed to derive secret keyring material and HMAC
func createSecretKeyringMaterialSeed(p *skmSeedParams) []byte {
	var tmp []byte
	for _, s := range [][]byte{
		p.FactProviderAddress.Bytes(),
		p.PassportAddress.Bytes(),
		p.FactKey[:],
	} {
		tmp = append(tmp, s...)
	}
	return tmp
}
