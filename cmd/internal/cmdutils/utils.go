package cmdutils

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/ethereum/go-ethereum/common"
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

type boolFlag struct {
	set   bool
	value bool
}

func (f *boolFlag) String() string {
	return strconv.FormatBool(f.value)
}

func (f *boolFlag) Set(s string) error {
	v, err := strconv.ParseBool(s)
	f.value = v
	f.set = true
	return err
}

func (f *boolFlag) IsSet() bool {
	return f.set
}

func (f *boolFlag) GetValue() bool {
	return f.value
}

// BoolVar defines a bool flag with specified name, default value, and usage string.
// The argument value is the default value of the flag.
func BoolVar(name string, value bool, usage string) *boolFlag {
	v := &boolFlag{value: value}
	flag.Var(v, name, usage)
	return v
}

type addressFlag struct {
	set   bool
	value common.Address
}

func (f *addressFlag) String() string {
	return f.value.Hex()
}

func (f *addressFlag) Set(s string) error {
	if !common.IsHexAddress(s) {
		return fmt.Errorf("invalid Ethereum address: %v", s)
	}
	f.value = common.HexToAddress(s)
	f.set = true
	return nil
}

func (f *addressFlag) IsSet() bool {
	return f.set
}

func (f *addressFlag) GetValue() common.Address {
	return f.value
}

// AddressVar defines a bool flag with specified name, default value, and usage string.
// The argument value is the default value of the flag.
func AddressVar(name string, value common.Address, usage string) *addressFlag {
	v := &addressFlag{value: value}
	flag.Var(v, name, usage)
	return v
}
