package facts

import (
	"context"
	"crypto/subtle"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/monetha/go-verifiable-data/contracts"
	"github.com/monetha/go-verifiable-data/eth"
	"github.com/monetha/go-verifiable-data/types/exchange"
	"github.com/pkg/errors"
)

const disputeGasLimit = 60000

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
	noWaitResult, err := f.DisputePrivateDataExchangeNoWait(ctx, passportAddress, exchangeIdx, exchangeKey)
	if err != nil {
		return nil, err
	}
	return f.GetDisputePrivateDataExchangeResult(ctx, noWaitResult)
}

// DisputePrivateDataExchangeResultNoWait holds result of calling DisputePrivateDataExchangeNoWait
type DisputePrivateDataExchangeResultNoWait struct {
	ExchangeIdx *big.Int
	TxHash      common.Hash
}

// DisputePrivateDataExchangeNoWait closes private data exchange after acceptance by dispute.
// This method does not wait for the transaction to be mined. Use the method without the NoWait suffix if you need to make
// sure that the transaction was successfully mined.
func (f *ExchangeDisputer) DisputePrivateDataExchangeNoWait(ctx context.Context, passportAddress common.Address, exchangeIdx *big.Int, exchangeKey [32]byte) (*DisputePrivateDataExchangeResultNoWait, error) {
	backend := f.s.Backend

	c := contracts.InitPassportLogicContract(passportAddress, backend)

	// getting private data exchange by index
	ex, err := c.PrivateDataExchanges(&bind.CallOpts{Context: ctx}, exchangeIdx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get private data exchange")
	}

	// it should be accepted
	exState := exchange.StateType(ex.State)
	if exState != exchange.Accepted {
		return nil, ErrExchangeMustBeAccepted
	}

	auth := f.s.TransactOpts
	auth.Context = ctx
	auth.GasLimit = disputeGasLimit

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

	return &DisputePrivateDataExchangeResultNoWait{
		ExchangeIdx: exchangeIdx,
		TxHash:      tx.Hash(),
	}, nil
}

// GetDisputePrivateDataExchangeResult waits for transaction to be mined if it's not mined yet and retrieves the dispute result for the provided private data exchange from tx logs.
func (f *ExchangeDisputer) GetDisputePrivateDataExchangeResult(ctx context.Context, noWaitResult *DisputePrivateDataExchangeResultNoWait) (*DisputePrivateDataExchangeResult, error) {
	txr, err := f.s.WaitForTxReceipt(ctx, noWaitResult.TxHash)
	if err != nil {
		return nil, err
	}

	// filtering PrivateDataExchangeDisputed event to get cheater address and dispute result
	evs, err := contracts.FilterPrivateDataExchangeDisputed(ctx, txr.Logs, []*big.Int{noWaitResult.ExchangeIdx}, nil, nil)
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
