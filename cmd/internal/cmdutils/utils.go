package cmdutils

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/ethereum/go-ethereum/log"
	"gitlab.com/monetha/protocol-go-sdk/eth"
)

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

// CheckBalance checks if it's enough funds to spend specified gas limit
func CheckBalance(ctx context.Context, session *eth.Session, gasLimit int64) {
	enoughFunds, minBalance, err := session.IsEnoughFunds(ctx, gasLimit)
	CheckErr(err, "IsEnoughFunds")
	if !enoughFunds {
		log.Warn("not enough funds", "min_balance_required", minBalance.String())
		os.Exit(1)
	}
}
