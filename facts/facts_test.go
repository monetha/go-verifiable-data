package facts_test

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"io"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/monetha/reputation-go-sdk/deployer"
	"github.com/monetha/reputation-go-sdk/eth"
	"github.com/monetha/reputation-go-sdk/eth/backend"
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

func generatePrivateKey(t *testing.T, rnd io.Reader) (*ecdsa.PrivateKey, common.Address) {
	key, err := ecdsa.GenerateKey(crypto.S256(), rnd)
	if err != nil {
		t.Fatalf("ecdsa.GenerateKey() error = %v", err)
	}
	addr := crypto.PubkeyToAddress(key.PublicKey)
	return key, addr
}

type account struct {
	Address common.Address
	Balance int64
}

func newPassportWithActors(
	ctx context.Context,
	t *testing.T,
	rnd io.Reader,
	accounts ...account,
) *passportWithActors {
	monethaKey, monethaAddress := generatePrivateKey(t, rnd)
	passportOwnerKey, passportOwnerAddress := generatePrivateKey(t, rnd)
	factProviderKey, factProviderAddress := generatePrivateKey(t, rnd)

	alloc := core.GenesisAlloc{
		monethaAddress:       {Balance: big.NewInt(deployer.PassportFactoryGasLimit)},
		passportOwnerAddress: {Balance: big.NewInt(math.MaxInt64)},
		factProviderAddress:  {Balance: big.NewInt(math.MaxInt64)},
	}
	for _, v := range accounts {
		alloc[v.Address] = core.GenesisAccount{Balance: big.NewInt(v.Balance)}
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
