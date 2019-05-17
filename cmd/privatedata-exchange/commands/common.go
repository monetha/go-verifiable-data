package commands

import (
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/monetha/reputation-go-sdk/eth"
	"github.com/pkg/errors"
)

func initLogging(level log.Lvl, vmodule string) {
	glogger := log.NewGlogHandler(log.StreamHandler(os.Stderr, log.TerminalFormat(false)))
	glogger.Verbosity(level)
	_ = glogger.Vmodule(vmodule)
	log.Root().SetHandler(glogger)
}

func newEth(backendURL string) (*eth.Eth, error) {
	client, err := ethclient.Dial(backendURL)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect ETH client")
	}

	return eth.New(client, log.Warn), nil
}

func fileExistsAndNotTty(name string) bool {
	fInfo, err := os.Stat(name)
	if os.IsNotExist(err) {
		return false
	}
	return (fInfo.Mode() & os.ModeCharDevice) == 0 // ignore tty
}
