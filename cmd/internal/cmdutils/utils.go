package cmdutils

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/contracts/chequebook"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
)

func CheckBalance(ctx context.Context, backend chequebook.Backend, address common.Address, minWei *big.Int) {
	log.Warn("Checking balance", "address", address.Hex())

	balance, err := backend.BalanceAt(ctx, address, nil)
	hint := fmt.Sprintf("BalanceAt(%v)", address.Hex())
	CheckErr(err, hint)

	if balance.Cmp(minWei) == -1 {
		CheckErr(fmt.Errorf("balance too low: %v wei < %v wei", balance, minWei), hint)
	}
}

func CheckErr(err error, hint string) {
	if err != nil {
		log.Error(hint, "err", err)
		os.Exit(1)
	}
}

func CheckTx(ctx context.Context, backend chequebook.Backend, txHash common.Hash) {
	err := WaitForTx(ctx, backend, txHash)
	CheckErr(err, fmt.Sprintf("tx 0x%x", txHash))
}

func CheckTxReceipt(ctx context.Context, backend chequebook.Backend, txHash common.Hash) *types.Receipt {
	txr, err := WaitForTxReceipt(ctx, backend, txHash)
	CheckErr(err, fmt.Sprintf("tx 0x%x", txHash))
	return txr
}

func WaitForTx(ctx context.Context, backend chequebook.Backend, txHash common.Hash) error {
	_, err := WaitForTxReceipt(ctx, backend, txHash)
	if err != nil {
		return err
	}
	return nil
}

func WaitForTxReceipt(ctx context.Context, backend chequebook.Backend, txHash common.Hash) (tr *types.Receipt, err error) {
	log.Warn("Waiting for transaction", "hash", txHash.Hex())

	type commiter interface {
		Commit()
	}
	if sim, ok := backend.(commiter); ok {
		sim.Commit()
		tr, err = onlySuccessfulReceipt(backend.TransactionReceipt(ctx, txHash))
		return
	}

	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(4 * time.Second):
		}

		tr, err = onlySuccessfulReceipt(backend.TransactionReceipt(ctx, txHash))
		if err == ethereum.NotFound {
			continue
		}
		return
	}
}

func onlySuccessfulReceipt(tr *types.Receipt, err error) (*types.Receipt, error) {
	if err != nil {
		return nil, err
	}
	if tr.Status != types.ReceiptStatusSuccessful {
		return nil, fmt.Errorf("tx failed: %+v", tr)
	}
	log.Warn("Transaction successfully mined", "tx_hash", tr.TxHash.Hex(), "cumulative_gas_used", tr.CumulativeGasUsed)
	return tr, nil
}

func CreateCtrlCContext() context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		defer cancel()

		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT)
		<-sigChan
		log.Warn("got interrupt signal")
	}()

	return ctx
}
