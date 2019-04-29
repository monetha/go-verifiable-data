package cmdutils

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/monetha/reputation-go-sdk/eth"
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

// BoolFlag holds flag value and indicator if it was set from command line
type BoolFlag struct {
	set   bool
	value bool
}

// String returns string representation of the value
func (f *BoolFlag) String() string {
	return strconv.FormatBool(f.value)
}

// Set sets value from string representation
func (f *BoolFlag) Set(s string) error {
	v, err := strconv.ParseBool(s)
	f.value = v
	f.set = true
	return err
}

// IsSet returns true if value was set from command line
func (f *BoolFlag) IsSet() bool {
	return f.set
}

// GetValue retruns flag value
func (f *BoolFlag) GetValue() bool {
	return f.value
}

// BoolVar defines a bool flag with specified name, default value, and usage string.
// The argument value is the default value of the flag.
func BoolVar(name string, value bool, usage string) *BoolFlag {
	v := &BoolFlag{value: value}
	flag.Var(v, name, usage)
	return v
}

// AddressFlag holds flag value and indicator if it was set from command line
type AddressFlag struct {
	set   bool
	value common.Address
}

// String returns string representation of the value
func (f *AddressFlag) String() string {
	return f.value.Hex()
}

// Set sets value from string representation
func (f *AddressFlag) Set(s string) error {
	if !common.IsHexAddress(s) {
		return fmt.Errorf("invalid Ethereum address: %v", s)
	}
	f.value = common.HexToAddress(s)
	f.set = true
	return nil
}

// IsSet returns true if value was set from command line
func (f *AddressFlag) IsSet() bool {
	return f.set
}

// GetValue returns flag value
func (f *AddressFlag) GetValue() common.Address {
	return f.value
}

// AddressVar defines a bool flag with specified name, default value, and usage string.
// The argument value is the default value of the flag.
func AddressVar(name string, value common.Address, usage string) *AddressFlag {
	v := &AddressFlag{value: value}
	flag.Var(v, name, usage)
	return v
}

// HashFlag holds flag value and indicator if it was set from command line
type HashFlag struct {
	set   bool
	value common.Hash
}

// String returns string representation of the value
func (f *HashFlag) String() string {
	return f.value.Hex()
}

// Set sets value from string representation
func (f *HashFlag) Set(s string) error {
	f.value = common.HexToHash(s)
	f.set = true
	return nil
}

// IsSet returns true if value was set from command line
func (f *HashFlag) IsSet() bool {
	return f.set
}

// GetValue returns flag value
func (f *HashFlag) GetValue() common.Hash {
	return f.value
}

// HashVar defines a bool flag with specified name, default value, and usage string.
// The argument value is the default value of the flag.
func HashVar(name string, value common.Hash, usage string) *HashFlag {
	v := &HashFlag{value: value}
	flag.Var(v, name, usage)
	return v
}

// PrivateKeyFlag holds flag value and indicator if it was set from command line
type PrivateKeyFlag struct {
	set   bool
	value *ecdsa.PrivateKey
}

// String returns string representation of the value
func (f *PrivateKeyFlag) String() string {
	return hex.EncodeToString(crypto.FromECDSA(f.value))
}

// Set sets value from string representation
func (f *PrivateKeyFlag) Set(s string) (err error) {
	f.value, err = crypto.HexToECDSA(s)
	f.set = true
	return
}

// IsSet returns true if value was set from command line
func (f *PrivateKeyFlag) IsSet() bool {
	return f.set
}

// GetValue returns flag value
func (f *PrivateKeyFlag) GetValue() *ecdsa.PrivateKey {
	return f.value
}

// PrivateKeyFlagVar defines a private key flag with specified name, default value, and usage string.
// The argument value is the default value of the flag.
func PrivateKeyFlagVar(name string, value *ecdsa.PrivateKey, usage string) *PrivateKeyFlag {
	v := &PrivateKeyFlag{value: value}
	flag.Var(v, name, usage)
	return v
}
