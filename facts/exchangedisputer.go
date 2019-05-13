package facts

import (
	"context"
	"crypto/subtle"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/monetha/reputation-go-sdk/contracts"
	"github.com/monetha/reputation-go-sdk/eth"
	"github.com/monetha/reputation-go-sdk/types/exchange"
	"github.com/pkg/errors"
)

// ExchangeDisputer allows to dispute data exchange
type ExchangeDisputer struct {
	s     *eth.Session
	clock Clock
}

// NewExchangeDisputer converts session to ExchangeDisputer
func NewExchangeDisputer(s *eth.Session, clock Clock) *ExchangeDisputer {
	if clock == nil {
		clock = realClock{}
	}

	return &ExchangeDisputer{s: s, clock: clock}
}

// DisputePrivateDataExchange closes private data exchange after acceptance by dispute.
func (f *ExchangeDisputer) DisputePrivateDataExchange(ctx context.Context, passportAddress common.Address, exchangeIdx *big.Int, exchangeKey [32]byte) error {
	backend := f.s.Backend

	c := contracts.InitPassportLogicContract(passportAddress, backend)

	// getting private data exchange by index
	ex, err := c.PrivateDataExchanges(&bind.CallOpts{Context: ctx}, exchangeIdx)
	if err != nil {
		return errors.Wrap(err, "failed to get private data exchange")
	}

	// it should be either closed or accepted
	exState := exchange.StateType(ex.State)
	if exState != exchange.Accepted {
		return ErrExchangeMustBeAccepted
	}

	auth := f.s.TransactOpts
	auth.Context = ctx

	// only data requester can call "dispute"
	if auth.From != ex.DataRequester {
		return ErrOnlyDataRequesterAllowedToCall
	}

	// it should not be expired
	if isExpired(ex.StateExpired, f.clock.Now().Add(1*time.Hour)) {
		return ErrExchangeIsExpiredOrWillBeExpiredVerySoon
	}

	// validating exchange key by comparing hashes
	exchangeKeyHash := crypto.Keccak256Hash(exchangeKey[:])
	if subtle.ConstantTimeCompare(ex.ExchangeKeyHash[:], exchangeKeyHash[:]) != 1 {
		return ErrInvalidExchangeKey
	}

	f.s.Log("Dispute private data exchange", "exchange_index", exchangeIdx.String())
	tx, err := c.DisputePrivateDataExchange(&auth, exchangeIdx, exchangeKey)
	if err != nil {
		return errors.Wrap(err, "failed to dispute accepted private data exchange")
	}

	_, err = f.s.WaitForTxReceipt(ctx, tx.Hash())
	return err
}
