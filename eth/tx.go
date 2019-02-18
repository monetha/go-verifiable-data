package eth

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

// Transaction add new methods to types.Transaction
type Transaction types.Transaction

// GetSenderPublicKey retrieves public key of sender from transaction.
func (t *Transaction) GetSenderPublicKey() (*ecdsa.PublicKey, error) {
	tx := (*types.Transaction)(t)

	V, R, S := tx.RawSignatureValues()
	r, s := R.Bytes(), S.Bytes()

	var (
		signedHash common.Hash
		v          byte
	)

	if tx.Protected() {
		chainID := tx.ChainId()
		signedHash = types.NewEIP155Signer(chainID).Hash(tx)
		v = byte(V.Uint64() - 2*chainID.Uint64() - 8 - 27)
	} else {
		signedHash = types.HomesteadSigner{}.Hash(tx)
		v = byte(V.Uint64() - 27)
	}

	if !crypto.ValidateSignatureValues(v, R, S, true) {
		return nil, types.ErrInvalidSig
	}

	sig := make([]byte, 65)
	copy(sig[32-len(r):32], r)
	copy(sig[64-len(s):64], s)
	sig[64] = v

	publicKey, err := crypto.SigToPub(signedHash[:], sig)
	if err != nil {
		return nil, err
	}

	return publicKey, err
}
