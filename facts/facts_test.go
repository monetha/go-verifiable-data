package facts_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
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

func createPassportAndFactProviderSession(ctx context.Context, t *testing.T) (common.Address, *eth.Session) {
	monethaKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatalf("crypto.GenerateKey() error = %v", err)
	}
	monethaAddress := bind.NewKeyedTransactor(monethaKey).From

	passportOwnerKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatalf("crypto.GenerateKey() error = %v", err)
	}
	passportOwnerAddress := bind.NewKeyedTransactor(passportOwnerKey).From

	factProviderKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatalf("crypto.GenerateKey() error = %v", err)
	}
	factProviderAddress := bind.NewKeyedTransactor(factProviderKey).From
	alloc := core.GenesisAlloc{
		monethaAddress:       {Balance: big.NewInt(deployer.PassportFactoryGasLimit)},
		passportOwnerAddress: {Balance: big.NewInt(deployer.PassportGasLimit)},
		factProviderAddress:  {Balance: big.NewInt(10000000000000)},
	}
	sim := backend.NewSimulatedBackendExtended(alloc, 10000000)
	sim.Commit()

	e := eth.New(sim, nil)
	if err := e.UpdateSuggestedGasPrice(ctx); err != nil {
		t.Fatalf("Eth.UpdateSuggestedGasPrice() error = %v", err)
	}

	monethaSession := e.NewSession(monethaKey)
	// deploying passport factory with all dependencies: passport logic, passport logic registry
	passportFactoryAddress, err := deployer.New(monethaSession).DeployPassportFactory(ctx)
	if err != nil {
		t.Fatalf("Deploy.DeployPassportFactory() error = %v", err)
	}

	passportOwnerSession := e.NewSession(passportOwnerKey)
	// deploying passport
	passportAddress, err := deployer.New(passportOwnerSession).DeployPassport(ctx, passportFactoryAddress)
	if err != nil {
		t.Fatalf("Deploy.DeployPassport() error = %v", err)
	}

	factProviderSession := e.NewSession(factProviderKey)

	return passportAddress, factProviderSession
}
