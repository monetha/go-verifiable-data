package cmdutils

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"os/signal"
	"syscall"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/contracts/chequebook"
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

// CheckErr prints error if it's not nil and causes the current program to exit with the status code equal to 1.
// The program terminates immediately; deferred functions are not run.
func CheckErr(err error, hint string) {
	if err != nil {
		log.Error(hint, "err", err)
		os.Exit(1)
	}
}

// CreateCtrlCContext returns context which's Done channel is closed when application should be terminated (on SIGINT, SIGTERM, SIGHUP, SIGQUIT signal)
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
