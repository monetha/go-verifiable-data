package facts

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/monetha/go-verifiable-data/contracts"
	"github.com/monetha/go-verifiable-data/eth"
	"github.com/monetha/go-verifiable-data/types/exchange"
	"github.com/pkg/errors"
)

const timeoutGasLimit = 60000

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
	txHash, err := f.TimeoutPrivateDataExchangeNoWait(ctx, passportAddress, exchangeIdx)
	if err == nil {
		_, err = f.s.WaitForTxReceipt(ctx, txHash)
	}
	return err
}

// TimeoutPrivateDataExchangeNoWait closes private data exchange after proposition has expired.
// This method does not wait for the transaction to be mined. Use the method without the NoWait suffix if you need to make
// sure that the transaction was successfully mined.
func (f *ExchangeTimeouter) TimeoutPrivateDataExchangeNoWait(ctx context.Context, passportAddress common.Address, exchangeIdx *big.Int) (common.Hash, error) {
	backend := f.s.Backend

	c := contracts.InitPassportLogicContract(passportAddress, backend)

	ex, err := c.PrivateDataExchanges(&bind.CallOpts{Context: ctx}, exchangeIdx)
	if err != nil {
		return common.Hash{}, errors.Wrap(err, "failed to get proposed private data exchange")
	}

	// it should be proposed
	exState := exchange.StateType(ex.State)
	if exState != exchange.Proposed {
		return common.Hash{}, ErrExchangeMustBeProposed
	}

	auth := f.s.TransactOpts
	auth.Context = ctx
	auth.GasLimit = timeoutGasLimit

	// only data requester can call "timeout"
	if auth.From != ex.DataRequester {
		return common.Hash{}, ErrOnlyDataRequesterAllowedToCall
	}

	// now should be 1 minute after expiration
	if !isExpired(ex.StateExpired, f.clock.Now().Add(-1*time.Minute)) {
		return common.Hash{}, ErrExchangeHasNotExpired
	}

	f.s.Log("Timeout private data exchange", "exchange_index", exchangeIdx.String())
	tx, err := c.TimeoutPrivateDataExchange(&auth, exchangeIdx)
	if err != nil {
		return common.Hash{}, errors.Wrap(err, "failed to close proposed private data exchange")
	}

	return tx.Hash(), nil
}
