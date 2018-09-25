package eth

import (
	"context"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/contracts/chequebook"
	"github.com/ethereum/go-ethereum/core/types"
	"gitlab.com/monetha/protocol-go-sdk/cmd/internal/log"
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
