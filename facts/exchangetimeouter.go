package facts

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/monetha/reputation-go-sdk/contracts"
	"github.com/monetha/reputation-go-sdk/eth"
	"github.com/monetha/reputation-go-sdk/types/exchange"
	"github.com/pkg/errors"
)

// ExchangeTimeouter allows to timeout data exchange/close it after proposition has expired.
type ExchangeTimeouter struct {
	s     *eth.Session
	clock Clock
}

// NewExchangeTimeouter converts session to ExchangeTimeouter
func NewExchangeTimeouter(s *eth.Session, clock Clock) *ExchangeTimeouter {
	if clock == nil {
		clock = realClock{}
	}

	return &ExchangeTimeouter{
		s:     s,
		clock: clock,
	}
}

// TimeoutPrivateDataExchange closes private data exchange after proposition has expired.
func (f *ExchangeTimeouter) TimeoutPrivateDataExchange(ctx context.Context, passportAddress common.Address, exchangeIdx *big.Int) error {
	backend := f.s.Backend

	c := contracts.InitPassportLogicContract(passportAddress, backend)

	ex, err := c.PrivateDataExchanges(&bind.CallOpts{Context: ctx}, exchangeIdx)
	if err != nil {
		return errors.Wrap(err, "failed to get proposed private data exchange")
	}

	// it should be proposed
	exState := exchange.StateType(ex.State)
	if exState != exchange.Proposed {
		return ErrExchangeMustBeProposed
	}

	auth := f.s.TransactOpts
	auth.Context = ctx

	// only data requester can call "timeout"
	if auth.From != ex.DataRequester {
		return ErrOnlyDataRequesterAllowedToCall
	}

	// now should be 1 minute after expiration
	if isExpired(ex.StateExpired, f.clock.Now().Add(-1*time.Minute)) {
		return ErrExchangeHasNotExpired
	}

	f.s.Log("Timeout private data exchange", "exchange_index", exchangeIdx.String())
	tx, err := c.TimeoutPrivateDataExchange(&auth, exchangeIdx)
	if err != nil {
		return errors.Wrap(err, "failed to close proposed private data exchange")
	}

	_, err = f.s.WaitForTxReceipt(ctx, tx.Hash())
	return err
}
