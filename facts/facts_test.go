package facts_test

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"io"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/monetha/reputation-go-sdk/deployer"
	"github.com/monetha/reputation-go-sdk/eth"
	"github.com/monetha/reputation-go-sdk/eth/backend"
)

var (
	ipfsURL  = "https://ipfs.infura.io:5001"
	randSeed = int64(1)
	factKey  = [32]byte{1, 2, 3, 4, 5, 6, 7, 8, 9}
	factData = []byte("this is a secret message")
)

func bigInt(x int64) *big.Int             { return big.NewInt(x) }
func exp(x *big.Int, y *big.Int) *big.Int { return new(big.Int).Exp(x, y, nil) }
func sub(x *big.Int, y *big.Int) *big.Int { return new(big.Int).Sub(x, y) }
func neg(x *big.Int) *big.Int             { return new(big.Int).Neg(x) }

type passportWithActors struct {
	*eth.Eth
	PassportOwnerKey     *ecdsa.PrivateKey
	PassportOwnerAddress common.Address
	FactProviderKey      *ecdsa.PrivateKey
	FactProviderAddress  common.Address
	PassportAddress      common.Address
}

func (pwa *passportWithActors) PassportOwnerSession() *eth.Session {
	return pwa.NewSession(pwa.PassportOwnerKey)
}

func (pwa *passportWithActors) FactProviderSession() *eth.Session {
	return pwa.NewSession(pwa.FactProviderKey)
}

func newPassportWithActors(ctx context.Context, t *testing.T, rnd io.Reader) *passportWithActors {
	monethaKey, err := ecdsa.GenerateKey(crypto.S256(), rnd)
	if err != nil {
		t.Fatalf("ecdsa.GenerateKey() error = %v", err)
	}
	monethaAddress := bind.NewKeyedTransactor(monethaKey).From

	passportOwnerKey, err := ecdsa.GenerateKey(crypto.S256(), rnd)
	if err != nil {
		t.Fatalf("ecdsa.GenerateKey() error = %v", err)
	}
	passportOwnerAddress := bind.NewKeyedTransactor(passportOwnerKey).From

	factProviderKey, err := ecdsa.GenerateKey(crypto.S256(), rnd)
	if err != nil {
		t.Fatalf("ecdsa.GenerateKey() error = %v", err)
	}
	factProviderAddress := bind.NewKeyedTransactor(factProviderKey).From
	alloc := core.GenesisAlloc{
		monethaAddress:       {Balance: big.NewInt(deployer.PassportFactoryGasLimit)},
		passportOwnerAddress: {Balance: big.NewInt(math.MaxInt64)},
		factProviderAddress:  {Balance: big.NewInt(math.MaxInt64)},
	}
	sim := backend.NewSimulatedBackendExtended(alloc, 10000000)
	sim.Commit()

	e := eth.New(sim, nil)
	if err := e.UpdateSuggestedGasPrice(ctx); err != nil {
		t.Fatalf("Eth.UpdateSuggestedGasPrice() error = %v", err)
	}

	monethaSession := e.NewSession(monethaKey)
	// deploying passport factory with all dependencies: passport logic, passport logic registry
	res, err := deployer.New(monethaSession).DeployPassportFactory(ctx)
	if err != nil {
		t.Fatalf("Deploy.DeployPassportFactory() error = %v", err)
	}
	passportFactoryAddress := res.PassportFactoryAddress

	passportOwnerSession := e.NewSession(passportOwnerKey)
	// deploying passport
	passportAddress, err := deployer.New(passportOwnerSession).DeployPassport(ctx, passportFactoryAddress)
	if err != nil {
		t.Fatalf("Deploy.DeployPassport() error = %v", err)
	}

	return &passportWithActors{
		Eth:                  e,
		PassportOwnerKey:     passportOwnerKey,
		PassportOwnerAddress: passportOwnerAddress,
		FactProviderKey:      factProviderKey,
		FactProviderAddress:  factProviderAddress,
		PassportAddress:      passportAddress,
	}
}

func createPassportAndFactProviderSession(ctx context.Context, t *testing.T) (common.Address, *eth.Session) {
	pa := newPassportWithActors(ctx, t, rand.Reader)

	return pa.PassportAddress, pa.FactProviderSession()
}
