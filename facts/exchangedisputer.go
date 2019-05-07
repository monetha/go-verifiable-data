package facts

import (
	"context"
	"math/big"

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
	tx, err := c.DisputePrivateDataExchange(&f.TransactOpts, exchangeIdx, exchangeKey)
	if err != nil {
		return errors.Wrap(err, "failed to dispute accepted private data exchange")
	}

	_, err = f.WaitForTxReceipt(ctx, tx.Hash())
	return err
}
