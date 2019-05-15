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

// DisputePrivateDataExchangeResult holds result of calling DisputePrivateDataExchange
type DisputePrivateDataExchangeResult struct {
	Successful bool
	Cheater    common.Address
}

// DisputePrivateDataExchange closes private data exchange after acceptance by dispute.
func (f *ExchangeDisputer) DisputePrivateDataExchange(ctx context.Context, passportAddress common.Address, exchangeIdx *big.Int, exchangeKey [32]byte) (*DisputePrivateDataExchangeResult, error) {
	backend := f.s.Backend

	c := contracts.InitPassportLogicContract(passportAddress, backend)

	// getting private data exchange by index
	ex, err := c.PrivateDataExchanges(&bind.CallOpts{Context: ctx}, exchangeIdx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get private data exchange")
	}

	// it should be either closed or accepted
	exState := exchange.StateType(ex.State)
	if exState != exchange.Accepted {
		return nil, ErrExchangeMustBeAccepted
	}

	auth := f.s.TransactOpts
	auth.Context = ctx

	// only data requester can call "dispute"
	if auth.From != ex.DataRequester {
		return nil, ErrOnlyDataRequesterAllowedToCall
	}

	// it should not be expired
	if isExpired(ex.StateExpired, f.clock.Now().Add(1*time.Hour)) {
		return nil, ErrExchangeIsExpiredOrWillBeExpiredVerySoon
	}

	// validating exchange key by comparing hashes
	exchangeKeyHash := crypto.Keccak256Hash(exchangeKey[:])
	if subtle.ConstantTimeCompare(ex.ExchangeKeyHash[:], exchangeKeyHash[:]) != 1 {
		return nil, ErrInvalidExchangeKey
	}

	f.s.Log("Dispute private data exchange", "exchange_index", exchangeIdx.String())
	tx, err := c.DisputePrivateDataExchange(&auth, exchangeIdx, exchangeKey)
	if err != nil {
		return nil, errors.Wrap(err, "failed to dispute accepted private data exchange")
	}

	txr, err := f.s.WaitForTxReceipt(ctx, tx.Hash())
	if err != nil {
		return nil, err
	}

	// filtering PrivateDataExchangeDisputed event to get cheater address and dispute result
	evs, err := contracts.FilterPrivateDataExchangeDisputed(ctx, txr.Logs, []*big.Int{exchangeIdx}, nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "filtering PrivateDataExchangeDisputed event")
	}
	if len(evs) != 1 {
		return nil, errors.New("expected exactly one PrivateDataExchangeDisputed event")
	}
	ev := evs[0]

	return &DisputePrivateDataExchangeResult{
		Successful: ev.Successful,
		Cheater:    ev.Cheater,
	}, nil
}
