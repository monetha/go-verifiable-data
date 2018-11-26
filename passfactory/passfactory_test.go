package passfactory_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"gitlab.com/monetha/reputation-go-sdk/deployer"
	"gitlab.com/monetha/reputation-go-sdk/eth"
	"gitlab.com/monetha/reputation-go-sdk/eth/backend"
	"gitlab.com/monetha/reputation-go-sdk/passfactory"
)

func TestReader_FilterPassports(t *testing.T) {
	t.Run("nil filtering options returns all passes", func(t *testing.T) {
		passCreation, ps := prepareReaderFilterPassports(t, func(result passportCreationResult) *passfactory.PassportFilterOpts {
			return nil
		})

		if len(ps) != 1 {
			t.Errorf("Expected to get exactly one passport, got %v", len(ps))
		}

		rPass := ps[0]
		if rPass.ContractAddress != passCreation.PassportAddress {
			t.Errorf("expected to get passport address %v, got %v", passCreation.PassportAddress.Hex(), rPass.ContractAddress)
		}
	})

	t.Run("empty filtering options returns all passes", func(t *testing.T) {
		passCreation, ps := prepareReaderFilterPassports(t, func(result passportCreationResult) *passfactory.PassportFilterOpts {
			return &passfactory.PassportFilterOpts{}
		})

		if len(ps) != 1 {
			t.Errorf("Expected to get exactly one passport, got %v", len(ps))
		}

		rPass := ps[0]
		if rPass.ContractAddress != passCreation.PassportAddress {
			t.Errorf("expected to get passport address %v, got %v", passCreation.PassportAddress.Hex(), rPass.ContractAddress)
		}
	})

	t.Run("filtering by passport returns pass", func(t *testing.T) {
		passCreation, ps := prepareReaderFilterPassports(t, func(result passportCreationResult) *passfactory.PassportFilterOpts {
			return &passfactory.PassportFilterOpts{Passport: []common.Address{result.PassportAddress}}
		})

		if len(ps) != 1 {
			t.Errorf("Expected to get exactly one passport, got %v", len(ps))
		}

		rPass := ps[0]
		if rPass.ContractAddress != passCreation.PassportAddress {
			t.Errorf("expected to get passport address %v, got %v", passCreation.PassportAddress.Hex(), rPass.ContractAddress)
		}
	})

	t.Run("filtering by owner returns pass", func(t *testing.T) {
		passCreation, ps := prepareReaderFilterPassports(t, func(result passportCreationResult) *passfactory.PassportFilterOpts {
			return &passfactory.PassportFilterOpts{Owner: []common.Address{result.PassportOwnerAddress}}
		})

		if len(ps) != 1 {
			t.Errorf("Expected to get exactly one passport, got %v", len(ps))
		}

		rPass := ps[0]
		if rPass.ContractAddress != passCreation.PassportAddress {
			t.Errorf("expected to get passport address %v, got %v", passCreation.PassportAddress.Hex(), rPass.ContractAddress)
		}
	})

	t.Run("filtering by unknown passport address returns nothing", func(t *testing.T) {
		_, ps := prepareReaderFilterPassports(t, func(result passportCreationResult) *passfactory.PassportFilterOpts {
			return &passfactory.PassportFilterOpts{Passport: []common.Address{result.PassportOwnerAddress}}
		})

		if len(ps) != 0 {
			t.Errorf("Expected to get nothing, got %v", len(ps))
		}
	})
}

func prepareReaderFilterPassports(t *testing.T, optsFun func(passportCreationResult) *passfactory.PassportFilterOpts) (passportCreationResult, []*passfactory.Passport) {
	ctx := context.Background()
	passCreation, e := createPassport(ctx, t)

	opts := optsFun(passCreation)

	pf := passfactory.NewReader(e)
	if opts != nil {
		opts.Context = ctx
	}
	it, err := pf.FilterPassports(opts, passCreation.PassportFactoryAddress)
	if err != nil {
		t.Errorf("FilterPassports: %v", err)
	}

	ps, err := it.ToSlice()
	if err != nil {
		t.Errorf("ToSlice: %v", err)
	}

	return passCreation, ps
}

type passportCreationResult struct {
	PassportAddress        common.Address
	PassportOwnerAddress   common.Address
	PassportFactoryAddress common.Address
}

func createPassport(ctx context.Context, t *testing.T) (passportCreationResult, *eth.Eth) {
	monethaKey, err := crypto.GenerateKey()
	if err != nil {
		t.Errorf("crypto.GenerateKey() error = %v", err)
	}
	monethaAddress := bind.NewKeyedTransactor(monethaKey).From

	passportOwnerKey, err := crypto.GenerateKey()
	if err != nil {
		t.Errorf("crypto.GenerateKey() error = %v", err)
	}
	passportOwnerAddress := bind.NewKeyedTransactor(passportOwnerKey).From

	factProviderKey, err := crypto.GenerateKey()
	if err != nil {
		t.Errorf("crypto.GenerateKey() error = %v", err)
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
	e.UpdateSuggestedGasPrice(ctx)

	monethaSession := e.NewSession(monethaKey)
	// deploying passport factory with all dependencies: passport logic, passport logic registry
	passportFactoryAddress, err := deployer.New(monethaSession).DeployPassportFactory(ctx)
	if err != nil {
		t.Errorf("Deploy.DeployPassportFactory() error = %v", err)
	}

	passportOwnerSession := e.NewSession(passportOwnerKey)
	// deploying passport
	passportAddress, err := deployer.New(passportOwnerSession).DeployPassport(ctx, passportFactoryAddress)
	if err != nil {
		t.Errorf("Deploy.DeployPassport() error = %v", err)
	}

	return passportCreationResult{
		PassportAddress:        passportAddress,
		PassportOwnerAddress:   passportOwnerAddress,
		PassportFactoryAddress: passportFactoryAddress,
	}, e
}
