package cmdutils

import (
	"crypto/ecdsa"
	"io"
	"math/rand"

	"github.com/ethereum/go-ethereum/crypto"
)

// NewMathRand creates a new source of pseudo-random numbers that uses provided seed
func NewMathRand(seed int64) *rand.Rand {
	return rand.New(rand.NewSource(seed))
}

// GenerateKey generates a public and private key pair. It uses the secp256k1 curve.
func GenerateKey(rand io.Reader) (*ecdsa.PrivateKey, error) {
	return ecdsa.GenerateKey(crypto.S256(), rand)
}
