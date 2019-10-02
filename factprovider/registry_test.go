package factprovider_test

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"io"
	"math"
	"math/big"
	"math/rand"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/go-test/deep"
	"github.com/monetha/go-verifiable-data/deployer"
	"github.com/monetha/go-verifiable-data/eth"
	"github.com/monetha/go-verifiable-data/eth/backend"
	"github.com/monetha/go-verifiable-data/factprovider"
	sdklog "github.com/monetha/go-verifiable-data/log"
)

func TestRegistry(t *testing.T) {
	const (
		randSeed = int64(1)
	)

	var (
		factProviderInfo = &factprovider.Info{
			Name:               "Company Ltd",
			ReputationPassport: common.HexToAddress("0x1"),
			Website:            "https://company.io",
		}
		factProviderAddress1 = common.HexToAddress("0x2")
		factProviderAddress2 = common.HexToAddress("0x3")
	)

	type testContext struct {
		context.Context
		*testing.T
		*registryWithActors
	}

	type actAssertFun func(
		tc *testContext,
	)

	arrangeActAssert := func(t *testing.T, actAssert actAssertFun) {
		ctx := context.Background()
		rnd := rand.New(rand.NewSource(randSeed))

		rwa := createRegistryWithActors(ctx, t, rnd)
		actAssert(&testContext{
			Context:            ctx,
			T:                  t,
			registryWithActors: rwa,
		})
	}

	t.Run("Monetha account is allowed to set fact provider info", func(t *testing.T) {
		arrangeActAssert(t, func(tc *testContext) {
			rw := factprovider.NewRegistryWriter(tc.MonethaSession, tc.FactProviderRegistryAddress)

			_, err := rw.SetFactProviderInfo(tc.Context, factProviderAddress1, factProviderInfo)
			if err != nil {
				t.Fatalf(err.Error())
			}
		})
	})

	t.Run("anyone is allowed to read fact provider info", func(t *testing.T) {
		arrangeActAssert(t, func(tc *testContext) {
			rw := factprovider.NewRegistryWriter(tc.MonethaSession, tc.FactProviderRegistryAddress)

			_, err := rw.SetFactProviderInfo(tc.Context, factProviderAddress1, factProviderInfo)
			if err != nil {
				t.Fatalf(err.Error())
			}

			rd := factprovider.NewRegistryReader(tc.Eth, tc.FactProviderRegistryAddress)
			actual, err := rd.GetFactProviderInfo(tc.Context, factProviderAddress1)
			if err != nil {
				t.Fatalf(err.Error())
			}
			if diff := deep.Equal(factProviderInfo, actual); diff != nil {
				t.Error(diff)
			}
		})
	})

	t.Run("reading non-registered fact provider info results in nil", func(t *testing.T) {
		arrangeActAssert(t, func(tc *testContext) {
			rw := factprovider.NewRegistryWriter(tc.MonethaSession, tc.FactProviderRegistryAddress)

			_, err := rw.SetFactProviderInfo(tc.Context, factProviderAddress1, factProviderInfo)
			if err != nil {
				t.Fatalf(err.Error())
			}

			rd := factprovider.NewRegistryReader(tc.Eth, tc.FactProviderRegistryAddress)
			actual, err := rd.GetFactProviderInfo(tc.Context, factProviderAddress2)
			if err != nil {
				t.Fatalf(err.Error())
			}
			if actual != nil {
				t.Fatalf("expecting nil, got %v", *actual)
			}
		})
	})

	t.Run("non-Monetha account is not allowed to set fact provider info", func(t *testing.T) {
		arrangeActAssert(t, func(tc *testContext) {
			rw := factprovider.NewRegistryWriter(tc.NonMonethaSession, tc.FactProviderRegistryAddress)

			_, err := rw.SetFactProviderInfo(tc.Context, factProviderAddress1, factProviderInfo)
			if err == nil {
				t.Fatal("expecting error")
			}
		})
	})

	t.Run("Monetha account is allowed to delete fact provider info", func(t *testing.T) {
		arrangeActAssert(t, func(tc *testContext) {
			rw := factprovider.NewRegistryWriter(tc.MonethaSession, tc.FactProviderRegistryAddress)

			_, err := rw.SetFactProviderInfo(tc.Context, factProviderAddress1, factProviderInfo)
			if err != nil {
				t.Fatalf(err.Error())
			}

			_, err = rw.DeleteFactProviderInfo(tc.Context, factProviderAddress1)
			if err != nil {
				t.Fatalf(err.Error())
			}
		})
	})

	t.Run("non-Monetha account is not allowed to delete fact provider info", func(t *testing.T) {
		arrangeActAssert(t, func(tc *testContext) {
			rw := factprovider.NewRegistryWriter(tc.MonethaSession, tc.FactProviderRegistryAddress)

			_, err := rw.SetFactProviderInfo(tc.Context, factProviderAddress1, factProviderInfo)
			if err != nil {
				t.Fatalf(err.Error())
			}

			rw2 := factprovider.NewRegistryWriter(tc.NonMonethaSession, tc.FactProviderRegistryAddress)
			_, err = rw2.DeleteFactProviderInfo(tc.Context, factProviderAddress1)
			if err == nil {
				t.Fatal("expecting error")
			}
		})
	})

	t.Run("reading deleted fact provider info results in nil", func(t *testing.T) {
		arrangeActAssert(t, func(tc *testContext) {
			rw := factprovider.NewRegistryWriter(tc.MonethaSession, tc.FactProviderRegistryAddress)

			_, err := rw.SetFactProviderInfo(tc.Context, factProviderAddress1, factProviderInfo)
			if err != nil {
				t.Fatalf(err.Error())
			}

			_, err = rw.DeleteFactProviderInfo(tc.Context, factProviderAddress1)
			if err != nil {
				t.Fatalf(err.Error())
			}

			rd := factprovider.NewRegistryReader(tc.Eth, tc.FactProviderRegistryAddress)
			actual, err := rd.GetFactProviderInfo(tc.Context, factProviderAddress1)
			if err != nil {
				t.Fatalf(err.Error())
			}
			if actual != nil {
				t.Fatalf("expecting nil, got %v", *actual)
			}
		})
	})
}

type account struct {
	Address common.Address
	Balance int64
}

type registryWithActors struct {
	*eth.Eth
	FactProviderRegistryAddress common.Address
	MonethaSession              *eth.Session
	NonMonethaSession           *eth.Session
}

func createRegistryWithActors(ctx context.Context, t *testing.T, rnd io.Reader, accounts ...account) *registryWithActors {
	monethaKey, monethaAddress := generatePrivateKey(t, rnd)
	nonMonethaKey, nonMonethaAddress := generatePrivateKey(t, rnd)

	alloc := core.GenesisAlloc{
		monethaAddress:    {Balance: big.NewInt(math.MaxInt64)},
		nonMonethaAddress: {Balance: big.NewInt(math.MaxInt64)},
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

	mthSession := e.NewSession(monethaKey)

	factProviderRegistryAddress, err := deployer.New(mthSession).DeployFactProviderRegistry(ctx)
	if err != nil {
		t.Fatalf("DeployFactProviderRegistry() error = %v", err)
	}

	return &registryWithActors{
		Eth:                         e,
		FactProviderRegistryAddress: factProviderRegistryAddress,
		MonethaSession:              mthSession,
		NonMonethaSession:           e.NewSession(nonMonethaKey),
	}
}

func generatePrivateKey(t *testing.T, rnd io.Reader) (*ecdsa.PrivateKey, common.Address) {
	key, err := ecdsa.GenerateKey(crypto.S256(), rnd)
	if err != nil {
		t.Fatalf("ecdsa.GenerateKey() error = %v", err)
	}
	addr := crypto.PubkeyToAddress(key.PublicKey)
	return key, addr
}

type tWriter struct{}

func (t tWriter) Write(p []byte) (n int, err error) {
	fmt.Print(string(p))
	return len(p), nil
}
