package eth

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/contracts/chequebook"
	"github.com/ethereum/go-ethereum/core/types"
	"gitlab.com/monetha/protocol-go-sdk/log"
)

// Eth simplifies some operations with the Ethereum network
type Eth struct {
	Backend chequebook.Backend
	LogFun  log.Fun
}

// WaitForTxReceipt waits until the transaction is successfully mined. It returns error if receipt status is not equal to `types.ReceiptStatusSuccessful`.
func (e *Eth) WaitForTxReceipt(ctx context.Context, txHash common.Hash) (tr *types.Receipt, err error) {
	backend := e.Backend

	txHashStr := txHash.Hex()
	e.Log("Waiting for transaction", "hash", txHashStr)

	defer func() {
		if err != nil {
			err = fmt.Errorf("waiting for tx(%v): %v", txHashStr, err)
		}
	}()

	type commiter interface {
		Commit()
	}
	if sim, ok := backend.(commiter); ok {
		sim.Commit()
		tr, err = e.onlySuccessfulReceipt(backend.TransactionReceipt(ctx, txHash))
		return
	}

	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(4 * time.Second):
		}

		tr, err = e.onlySuccessfulReceipt(backend.TransactionReceipt(ctx, txHash))
		if err == ethereum.NotFound {
			continue
		}
		return
	}
}

func (e *Eth) onlySuccessfulReceipt(tr *types.Receipt, err error) (*types.Receipt, error) {
	if err != nil {
		return nil, err
	}
	if tr.Status != types.ReceiptStatusSuccessful {
		return nil, fmt.Errorf("tx failed: %+v", tr)
	}
	e.Log("Transaction successfully mined", "tx_hash", tr.TxHash.Hex(), "cumulative_gas_used", tr.CumulativeGasUsed)
	return tr, nil
}

// Log writes message to the log with the context data
func (e *Eth) Log(msg string, ctx ...interface{}) {
	lf := e.LogFun
	if lf != nil {
		lf(msg, ctx...)
	}
}

// Session provides holds basic pre-configured parameters like backend, authorization, logging
type Session struct {
	*Eth
	TransactOpts bind.TransactOpts
}

// NewSession creates an instance of Session
func NewSession(backend chequebook.Backend, key *ecdsa.PrivateKey) *Session {
	return &Session{
		Eth: &Eth{
			Backend: backend,
		},
		TransactOpts: *bind.NewKeyedTransactor(key),
	}
}

// SetLogFun sets Log function
func (s *Session) SetLogFun(lf log.Fun) *Session {
	s.LogFun = lf
	return s
}

// SetGasPrice sets gas price
func (s *Session) SetGasPrice(gasPrice *big.Int) *Session {
	s.TransactOpts.GasPrice = gasPrice
	return s
}

// IsEnoughFunds retrieves current account balance and checks if it's enough funds given gas limit.
// SetGasPrice needs to be called with non-nil parameter before calling this method.
func (s *Session) IsEnoughFunds(ctx context.Context, gasLimit int64) (enough bool, minBalance *big.Int, err error) {
	gasPrice := s.TransactOpts.GasPrice
	if gasPrice == nil {
		panic("gas price must be non nil")
	}

	minBalance = new(big.Int).Mul(big.NewInt(gasLimit), gasPrice)

	s.Log("Getting balance", "address", s.TransactOpts.From.Hex())

	var balance *big.Int
	balance, err = s.Backend.BalanceAt(ctx, s.TransactOpts.From, nil)
	if err != nil {
		err = fmt.Errorf("backend BalanceAt(%v): %v", s.TransactOpts.From.Hex(), err)
		return
	}

	if balance.Cmp(minBalance) == -1 {
		return
	}

	enough = true
	return
}
