package facts_test

import (
	"bytes"
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"gitlab.com/monetha/protocol-go-sdk/deployer"
	"gitlab.com/monetha/protocol-go-sdk/eth"
	"gitlab.com/monetha/protocol-go-sdk/eth/backend"
	"gitlab.com/monetha/protocol-go-sdk/facts"
)

func TestFactReader_ReadTxData(t *testing.T) {

	tests := []struct {
		name string
		key  [32]byte
		data []byte
	}{
		{"test 1", [32]byte{}, nil},
		{"test 2", [32]byte{}, []byte{}},
		{"test 3", [32]byte{}, []byte{1, 2, 3, 4}},
		{"test 4", [32]byte{9, 8, 7, 6}, []byte{1, 2, 3, 4}},
		{"test 5", [32]byte{9, 8, 7, 6}, make([]byte, 10000)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			passportAddress, factProviderSession := createPassportAndFactProviderSession(t)

			e := factProviderSession.Eth
			factProviderAddress := factProviderSession.TransactOpts.From

			err := facts.NewProvider(factProviderSession).WriteTxData(context.TODO(), passportAddress, tt.key, tt.data)
			if err != nil {
				t.Errorf("FactProvider.WriteTxData() error = %v", err)
			}

			readData, err := facts.NewReader(e).ReadTxData(context.TODO(), passportAddress, factProviderAddress, tt.key)
			if err != nil {
				t.Errorf("FactReader.ReadTxData() error = %v", err)
			}

			if !bytes.Equal(tt.data, readData) {
				t.Errorf("Expected data = %v, read data = %v", tt.data, readData)
			}
		})
	}

	t.Run("reading non existing key value", func(t *testing.T) {
		passportAddress, factProviderSession := createPassportAndFactProviderSession(t)

		e := factProviderSession.Eth
		factProviderAddress := factProviderSession.TransactOpts.From

		key := [32]byte{1, 2, 3}

		_, err := facts.NewReader(e).ReadTxData(context.TODO(), passportAddress, factProviderAddress, key)
		if err != ethereum.NotFound {
			t.Errorf("FactReader.ReadTxData() expecting error = %v, got error = %v", ethereum.NotFound, err)
		}
	})
}

func createPassportAndFactProviderSession(t *testing.T) (common.Address, *eth.Session) {
	ctx := context.TODO()

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

	factProviderSession := e.NewSession(factProviderKey)

	return passportAddress, factProviderSession
}
