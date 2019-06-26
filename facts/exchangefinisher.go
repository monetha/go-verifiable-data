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

var (
	// ErrExchangeMustBeAccepted returned when private data exchange is not in Accepted state
	ErrExchangeMustBeAccepted = errors.New("private data exchange must be in Accepted state")
	// ErrOnlyExchangeParticipantsAreAllowedToCall returned when called not by data requester or passport owner
	ErrOnlyExchangeParticipantsAreAllowedToCall = errors.New("only data requester or passport owner are allowed to call")
)

// ExchangeFinisher allows to finish data exchange
type ExchangeFinisher struct {
	s     *eth.Session
	clock Clock
}

// NewExchangeFinisher converts session to ExchangeFinisher
func NewExchangeFinisher(s *eth.Session, clock Clock) *ExchangeFinisher {
	if clock == nil {
		clock = realClock{}
	}

	return &ExchangeFinisher{s: s, clock: clock}
}

// FinishPrivateDataExchange closes private data exchange after acceptance. It's supposed to be called by the data requester,
// but passport owner also can call it after data exchange is expired.
func (f *ExchangeFinisher) FinishPrivateDataExchange(ctx context.Context, passportAddress common.Address, exchangeIdx *big.Int) error {
	backend := f.s.Backend

	c := contracts.InitPassportLogicContract(passportAddress, backend)

	ex, err := c.PrivateDataExchanges(&bind.CallOpts{Context: ctx}, exchangeIdx)
	if err != nil {
		return errors.Wrap(err, "failed to get proposed private data exchange")
	}

	// it should be accepted
	exState := exchange.StateType(ex.State)
	if exState != exchange.Accepted {
		return ErrExchangeMustBeAccepted
	}

	auth := f.s.TransactOpts
	auth.Context = ctx

	if auth.From == ex.PassportOwner {
		// now should be 1 minute after expiration
		if !isExpired(ex.StateExpired, f.clock.Now().Add(-1*time.Minute)) {
			return ErrExchangeHasNotExpired
		}
	} else if auth.From != ex.DataRequester {
		return ErrOnlyExchangeParticipantsAreAllowedToCall
	}

	tx, err := c.FinishPrivateDataExchange(&auth, exchangeIdx)
	if err != nil {
		return errors.Wrap(err, "failed to finish accepted private data exchange")
	}

	_, err = f.s.WaitForTxReceipt(ctx, tx.Hash())
	return err
}
