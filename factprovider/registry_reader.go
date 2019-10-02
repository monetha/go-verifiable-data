package factprovider

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/monetha/go-verifiable-data/contracts"
	"github.com/monetha/go-verifiable-data/eth"
)

// RegistryReader reads public information of registered fact providers.
type RegistryReader struct {
	*eth.Eth
	registryContract *contracts.FactProviderRegistryContract
}

// NewRegistryReader converts Ethereum backend to RegistryReader
func NewRegistryReader(e *eth.Eth, factProviderRegistry common.Address) *RegistryReader {
	return &RegistryReader{Eth: e, registryContract: contracts.InitFactProviderRegistryContract(factProviderRegistry, e.Backend)}
}

// GetFactProviderInfo retrieves fact provider information given it's address. It returns nil when registry contains no information about fact provider.
func (r *RegistryReader) GetFactProviderInfo(ctx context.Context, factProviderAddress common.Address) (*FactProviderInfo, error) {
	fp, err := r.registryContract.FactProviders(&bind.CallOpts{Context: ctx}, factProviderAddress)
	if err != nil {
		return nil, err
	}
	if !fp.Initialized {
		return nil, nil
	}
	return &FactProviderInfo{
		Name:               fp.Name,
		ReputationPassport: fp.ReputationPassport,
		Website:            fp.Website,
	}, nil
}
