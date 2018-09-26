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

// Eth contains code to wait for transaction
type Eth struct {
	Log log.Fun
}

// WaitForTxReceipt waits until the transaction is successfully mined. It returns error if receipt status is not equal to `types.ReceiptStatusSuccessful`.
func (t Eth) WaitForTxReceipt(ctx context.Context, backend chequebook.Backend, txHash common.Hash) (tr *types.Receipt, err error) {
	txHashStr := txHash.Hex()
	t.log("Waiting for transaction", "hash", txHashStr)

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
		tr, err = t.onlySuccessfulReceipt(backend.TransactionReceipt(ctx, txHash))
		return
	}

	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(4 * time.Second):
		}

		tr, err = t.onlySuccessfulReceipt(backend.TransactionReceipt(ctx, txHash))
		if err == ethereum.NotFound {
			continue
		}
		return
	}
}

func (t Eth) onlySuccessfulReceipt(tr *types.Receipt, err error) (*types.Receipt, error) {
	if err != nil {
		return nil, err
	}
	if tr.Status != types.ReceiptStatusSuccessful {
		return nil, fmt.Errorf("tx failed: %+v", tr)
	}
	t.log("Transaction successfully mined", "tx_hash", tr.TxHash.Hex(), "cumulative_gas_used", tr.CumulativeGasUsed)
	return tr, nil
}

func (t Eth) log(msg string, ctx ...interface{}) {
	l := t.Log
	if l != nil {
		l(msg, ctx...)
	}
}

// PrepareAuth prepares authorization data required to create a valid Ethereum transaction.
// It retrieves the currently suggested gas price to allow a timely	execution of a transaction and sets it in authorization data.
// It also checks if the balance of the given account has enough weis to execute transaction(s) for the given `minGasLimit`.
func PrepareAuth(ctx context.Context, contractBackend chequebook.Backend, key *ecdsa.PrivateKey, minGasLimit int64) (*bind.TransactOpts, error) {
	auth := bind.NewKeyedTransactor(key)

	gasPrice, err := contractBackend.SuggestGasPrice(ctx)
	if err != nil {
		return nil, fmt.Errorf("backend SuggestGasPrice: %v", err)
	}
	auth.GasPrice = gasPrice

	minBalance := new(big.Int).Mul(big.NewInt(minGasLimit), gasPrice)

	balance, err := contractBackend.BalanceAt(ctx, auth.From, nil)
	if err != nil {
		return nil, fmt.Errorf("backend BalanceAt(%v): %v", auth.From.Hex(), err)
	}
	if balance.Cmp(minBalance) == -1 {
		return nil, fmt.Errorf("balance too low: %v wei < %v wei", balance, minBalance)
	}

	return auth, nil
}
