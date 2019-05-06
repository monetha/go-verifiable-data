package facts

import (
	"context"
	"math/big"

	"github.com/monetha/reputation-go-sdk/contracts"
	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/monetha/reputation-go-sdk/eth"
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
	tx, err := c.FinishPrivateDataExchange(&f.TransactOpts, exchangeIdx)
	if err != nil {
		return errors.Wrap(err, "failed to finish accepted private data exchange")
	}

	_, err = f.WaitForTxReceipt(ctx, tx.Hash())
	return err
}
