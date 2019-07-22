package commands

import (
	"context"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/monetha/go-verifiable-data/eth"
	"github.com/pkg/errors"
)

func initLogging(level log.Lvl, vmodule string) {
	glogger := log.NewGlogHandler(log.StreamHandler(os.Stderr, log.TerminalFormat(false)))
	glogger.Verbosity(level)
	_ = glogger.Vmodule(vmodule)
	log.Root().SetHandler(glogger)
}

func newEth(ctx context.Context, backendURL string) (e *eth.Eth, err error) {
	client, err := ethclient.DialContext(ctx, backendURL)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect ETH client")
	}

	e = eth.New(client, log.Warn)
	if err = e.UpdateSuggestedGasPrice(ctx); err != nil {
		e = nil
	}
	return
}

func fileExistsAndNotTty(name string) bool {
	fInfo, err := os.Stat(name)
	if os.IsNotExist(err) {
		return false
	}
	return (fInfo.Mode() & os.ModeCharDevice) == 0 // ignore tty
}
