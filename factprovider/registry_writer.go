package factprovider

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/monetha/go-verifiable-data/contracts"
	"github.com/monetha/go-verifiable-data/eth"
	"github.com/pkg/errors"
)

// RegistryWriter allows to write to fact provider registry
type RegistryWriter struct {
	*eth.Session
	c *contracts.FactProviderRegistryContract
}

// NewRegistryWriter converts session to RegistryWriter
func NewRegistryWriter(s *eth.Session, factProviderRegistry common.Address) *RegistryWriter {
	c := contracts.InitFactProviderRegistryContract(factProviderRegistry, s.Backend)
	return &RegistryWriter{
		Session: s,
		c:       c,
	}
}

// SetFactProviderInfo writes information about fact provider into registry.
func (r *RegistryWriter) SetFactProviderInfo(ctx context.Context, factProvider common.Address, info *FactProviderInfo) (common.Hash, error) {
	txHash, err := r.SetFactProviderInfoNoWait(ctx, factProvider, info)
	if err == nil {
		_, err = r.WaitForTxReceipt(ctx, txHash)
	}
	return txHash, err
}

// SetFactProviderInfoNoWait writes information about fact provider into registry.
// This method does not wait for the transaction to be mined. Use the method without the NoWait suffix if you need to make
// sure that the transaction was successfully mined.
func (r *RegistryWriter) SetFactProviderInfoNoWait(ctx context.Context, factProvider common.Address, info *FactProviderInfo) (common.Hash, error) {
	tx, err := r.c.SetFactProviderInfo(r.txOpts(ctx), factProvider, info.Name, info.ReputationPassport, info.Website)
	if err != nil {
		return common.Hash{}, errors.Wrap(err, "failed to set fact provider info")
	}
	return tx.Hash(), nil
}

// DeleteFactProviderInfo deletes information about fact provider from registry.
func (r *RegistryWriter) DeleteFactProviderInfo(ctx context.Context, factProvider common.Address) (common.Hash, error) {
	txHash, err := r.DeleteFactProviderInfoNoWait(ctx, factProvider)
	if err == nil {
		_, err = r.WaitForTxReceipt(ctx, txHash)
	}
	return txHash, err
}

// DeleteFactProviderInfoNoWait deletes information about fact provider from registry.
// This method does not wait for the transaction to be mined. Use the method without the NoWait suffix if you need to make
// sure that the transaction was successfully mined.
func (r *RegistryWriter) DeleteFactProviderInfoNoWait(ctx context.Context, factProvider common.Address) (common.Hash, error) {
	tx, err := r.c.DeleteFactProviderInfo(r.txOpts(ctx), factProvider)
	if err != nil {
		return common.Hash{}, errors.Wrap(err, "failed to delete fact provider info")
	}
	return tx.Hash(), nil
}

func (r *RegistryWriter) txOpts(ctx context.Context) *bind.TransactOpts {
	opts := r.TransactOpts
	opts.Context = ctx
	return &opts
}
