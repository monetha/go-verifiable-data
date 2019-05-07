package facts

import (
	"context"
	"crypto/subtle"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/monetha/reputation-go-sdk/types/exchange"

	"github.com/ethereum/go-ethereum/common"
	"github.com/monetha/reputation-go-sdk/contracts"
	"github.com/monetha/reputation-go-sdk/eth"
	"github.com/pkg/errors"
)

// ExchangeDisputer allows to dispute data exchange
type ExchangeDisputer eth.Session

// NewExchangeDisputer converts session to ExchangeDisputer
func NewExchangeDisputer(s *eth.Session) *ExchangeDisputer {
	return (*ExchangeDisputer)(s)
}

// DisputePrivateDataExchange closes private data exchange after acceptance by dispute.
func (f *ExchangeDisputer) DisputePrivateDataExchange(ctx context.Context, passportAddress common.Address, exchangeIdx *big.Int, exchangeKey [32]byte) error {
	backend := f.Backend

	c := contracts.InitPassportLogicContract(passportAddress, backend)

	// getting private data exchange by index
	ex, err := c.PrivateDataExchanges(nil, exchangeIdx)
	if err != nil {
		return errors.Wrap(err, "failed to get private data exchange")
	}

	// it should be either closed or accepted
	exState := exchange.StateType(ex.State)
	if exState != exchange.Closed && exState != exchange.Accepted {
		return ErrExchangeMustBeClosedOrAccepted
	}

	// validating exchange key by comparing hashes
	exchangeKeyHash := crypto.Keccak256Hash(exchangeKey[:])
	if subtle.ConstantTimeCompare(ex.ExchangeKeyHash[:], exchangeKeyHash[:]) != 1 {
		return ErrInvalidExchangeKey
	}

	// TODO: add expiration check

	f.Log("Dispute private datat exchange", "exchange_index", exchangeIdx.String())
	tx, err := c.DisputePrivateDataExchange(&f.TransactOpts, exchangeIdx, exchangeKey)
	if err != nil {
		return errors.Wrap(err, "failed to dispute accepted private data exchange")
	}

	_, err = f.WaitForTxReceipt(ctx, tx.Hash())
	return err
}
