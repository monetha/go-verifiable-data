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

// ExchangeStatusGetter allows to get status of private data exchange
type ExchangeStatusGetter eth.Eth

// NewExchangeStatusGetter converts eth.Eth to ExchangeStatusGetter
func NewExchangeStatusGetter(e *eth.Eth) *ExchangeStatusGetter {
	return (*ExchangeStatusGetter)(e)
}

// PrivateDataExchangeStatus holds details of private data exchange
type PrivateDataExchangeStatus struct {
	DataRequester        common.Address
	DataRequesterStaked  *big.Int
	PassportOwner        common.Address
	PassportOwnerStaked  *big.Int
	FactProvider         common.Address
	FactKey              [32]byte
	DataIPFSHash         string
	DataKeyHash          common.Hash
	EncryptedExchangeKey []byte
	ExchangeKeyHash      common.Hash
	EncryptedDataKey     [32]byte
	State                exchange.StateType
	StateExpirationTime  time.Time
}

// GetPrivateDataExchangeStatus returns the status of private data exchange
func (g *ExchangeStatusGetter) GetPrivateDataExchangeStatus(ctx context.Context, passportAddress common.Address, exchangeIdx *big.Int) (*PrivateDataExchangeStatus, error) {
	backend := g.Backend

	c := contracts.InitPassportLogicContract(passportAddress, backend)

	ex, err := c.PrivateDataExchanges(&bind.CallOpts{Context: ctx}, exchangeIdx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get private data exchange")
	}

	return &PrivateDataExchangeStatus{
		DataRequester:        ex.DataRequester,
		DataRequesterStaked:  ex.DataRequesterValue,
		PassportOwner:        ex.PassportOwner,
		PassportOwnerStaked:  ex.PassportOwnerValue,
		FactProvider:         ex.FactProvider,
		FactKey:              ex.Key,
		DataIPFSHash:         ex.DataIPFSHash,
		DataKeyHash:          ex.DataKeyHash,
		EncryptedExchangeKey: ex.EncryptedExchangeKey,
		ExchangeKeyHash:      ex.ExchangeKeyHash,
		EncryptedDataKey:     ex.EncryptedDataKey,
		State:                exchange.StateType(ex.State),
		StateExpirationTime:  time.Unix(ex.StateExpired.Int64(), 0),
	}, nil
}
