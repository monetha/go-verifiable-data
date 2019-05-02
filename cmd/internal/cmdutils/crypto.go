package cmdutils

import (
	"crypto/ecdsa"
	"math/rand"

	"github.com/ethereum/go-ethereum/crypto"
)

var (
	// TestMonethaAdminKey is private key of Monetha admin (only used in simulated block-chain for test purposes)
	TestMonethaAdminKey = mustPrivateKey("163f5f0f9a621d72fedd85ffca3d08d131ab4e812181e0d30ffd1c885d20aac7")
	// TestPassportOwnerKey is private key of passport owner (only used in simulated block-chain for test purposes)
	TestPassportOwnerKey = mustPrivateKey("6694d2c422acd209aac6a3d9841b92f396e82373934b9baacbbcdf376b5a492a")
	// TestFactProviderKey is private key of fact provider (only used in simulated block-chain for test purposes)
	TestFactProviderKey = mustPrivateKey("a2ff6cd471c483f21df01a203e0f81de8093f10bb1d822eb0f4644a3c225919c")
)

// NewMathRandFixedSeed creates a new source of pseudo-random numbers that uses fixed seed, so it produces the same set of
// pseudo-random numbers.
func NewMathRandFixedSeed() *rand.Rand { return rand.New(rand.NewSource(1)) }

func mustPrivateKey(s string) *ecdsa.PrivateKey {
	k, err := crypto.HexToECDSA(s)
	if err != nil {
		panic(err)
	}
	return k
}
