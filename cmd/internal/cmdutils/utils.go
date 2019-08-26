package cmdutils

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/log"
	"github.com/monetha/go-verifiable-data/eth"
)

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
func CheckBalance(ctx context.Context, session *eth.Session, gasLimit int64) error {
	enoughFunds, minBalance, err := session.IsEnoughFunds(ctx, gasLimit)
	if err != nil {
		return err
	}

	if !enoughFunds {
		return fmt.Errorf("not enough funds. Minimum balance required is %v", minBalance.String())
	}

	return nil
}

// InitLogging inits logging for root logger
func InitLogging(level log.Lvl, vmodule string) {
	glogger := log.NewGlogHandler(log.StreamHandler(os.Stderr, log.TerminalFormat(false)))
	glogger.Verbosity(level)
	_ = glogger.Vmodule(vmodule)
	log.Root().SetHandler(glogger)
}

// NewEth creates Ethereum helper, preconfigured to work with specific blockchain from given parameters
func NewEth(ctx context.Context, backendURL string, enclaveURL string, privateFor []string) (e *eth.Eth, err error) {
	bf := NewBackendFactory(&enclaveURL, privateFor)

	client, err := bf.DialBackendContext(ctx, backendURL)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect ETH client")
	}

	e = bf.NewEth(ctx, client)
	if err = e.UpdateSuggestedGasPrice(ctx); err != nil {
		e = nil
	}
	return
}
