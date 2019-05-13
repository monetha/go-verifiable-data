package facts

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/monetha/reputation-go-sdk/contracts"
	"github.com/monetha/reputation-go-sdk/eth"
	"github.com/monetha/reputation-go-sdk/types/exchange"
	"github.com/pkg/errors"
)

var (
	// ErrExchangeMustBeAccepted returned when private data exchange is not in Accepted state
	ErrExchangeMustBeAccepted = errors.New("private data exchange must be in Accepted state")
)

// ExchangeFinisher allows to finish data exchange
type ExchangeFinisher eth.Session

// NewExchangeFinisher converts session to ExchangeFinisher
func NewExchangeFinisher(s *eth.Session) *ExchangeFinisher {
	return (*ExchangeFinisher)(s)
}

// FinishPrivateDataExchange closes private data exchange after acceptance. It's supposed to be called by the data requester,
// but passport owner also can call it after data exchange is expired.
func (f *ExchangeFinisher) FinishPrivateDataExchange(ctx context.Context, passportAddress common.Address, exchangeIdx *big.Int) error {
	backend := f.Backend

	c := contracts.InitPassportLogicContract(passportAddress, backend)

	ex, err := c.PrivateDataExchanges(&bind.CallOpts{Context: ctx}, exchangeIdx)
	if err != nil {
		return errors.Wrap(err, "failed to get proposed private data exchange")
	}

	// it should be proposed
	exState := exchange.StateType(ex.State)
	if exState != exchange.Accepted {
		return ErrExchangeMustBeAccepted
	}

	// TODO: add expiration check

	tx, err := c.FinishPrivateDataExchange(&f.TransactOpts, exchangeIdx)
	if err != nil {
		return errors.Wrap(err, "failed to finish accepted private data exchange")
	}

	_, err = f.WaitForTxReceipt(ctx, tx.Hash())
	return err
}
