package facts

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/monetha/reputation-go-sdk/contracts"
	"github.com/monetha/reputation-go-sdk/eth"
	"github.com/monetha/reputation-go-sdk/types/exchange"
	"github.com/pkg/errors"
)

// ExchangeTimeouter allows to timeout data exchange/close it after proposition has expired.
type ExchangeTimeouter eth.Session

// NewExchangeTimeouter converts session to ExchangeTimeouter
func NewExchangeTimeouter(s *eth.Session) *ExchangeTimeouter {
	return (*ExchangeTimeouter)(s)
}

// TimeoutPrivateDataExchange closes private data exchange after proposition has expired.
func (f *ExchangeTimeouter) TimeoutPrivateDataExchange(ctx context.Context, passportAddress common.Address, exchangeIdx *big.Int) error {
	backend := f.Backend

	c := contracts.InitPassportLogicContract(passportAddress, backend)

	ex, err := c.PrivateDataExchanges(nil, exchangeIdx)
	if err != nil {
		return errors.Wrap(err, "failed to get proposed private data exchange")
	}

	// it should be proposed
	exState := exchange.StateType(ex.State)
	if exState != exchange.Proposed {
		return ErrExchangeMustBeProposed
	}

	// TODO: add expiration check

	f.Log("Timeout private data exchange", "exchange_index", exchangeIdx.String())
	tx, err := c.TimeoutPrivateDataExchange(&f.TransactOpts, exchangeIdx)
	if err != nil {
		return errors.Wrap(err, "failed to close proposed private data exchange")
	}

	_, err = f.WaitForTxReceipt(ctx, tx.Hash())
	return err
}
