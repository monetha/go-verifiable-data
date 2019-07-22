package facts_test

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"fmt"
	"io"
	"math/big"
	"sync"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/monetha/go-verifiable-data/deployer"
	"github.com/monetha/go-verifiable-data/eth"
	"github.com/monetha/go-verifiable-data/eth/backend"
	sdklog "github.com/monetha/go-verifiable-data/log"
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

	if sdklog.IsTestLogSet() {
		glogHandler := log.NewGlogHandler(log.StreamHandler(tWriter{}, log.LogfmtFormat()))
		glogHandler.Verbosity(log.LvlWarn)
		log.Root().SetHandler(glogHandler)
	}

	e := eth.New(sim, log.Warn)
	if err := e.UpdateSuggestedGasPrice(ctx); err != nil {
		t.Fatalf("Eth.UpdateSuggestedGasPrice() error = %v", err)
	}

	monethaSession := e.NewSession(monethaKey)
	// deploying passport factory with all dependencies: passport logic, passport logic registry
	res, err := deployer.New(monethaSession).Bootstrap(ctx)
	if err != nil {
		t.Fatalf("Deploy.Bootstrap() error = %v", err)
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

type clockMock struct {
	mu  sync.RWMutex
	now time.Time // current time
}

func newClockMock() *clockMock {
	return &clockMock{now: time.Unix(0, 0)}
}

func (c *clockMock) Now() time.Time {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.now
}

func (c *clockMock) Add(d time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.now = c.now.Add(d)
}

type tWriter struct{}

func (t tWriter) Write(p []byte) (n int, err error) {
	fmt.Print(string(p))
	return len(p), nil
}
