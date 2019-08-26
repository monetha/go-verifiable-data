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
	"strings"
	"syscall"

	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/monetha/go-verifiable-data/eth"
	"github.com/monetha/go-verifiable-data/types/data"
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

var (
	dataTypes      = make(map[string]data.Type)
	dataTypeSetStr string
)

func init() {
	keys := make([]string, 0, len(data.TypeValues()))

	for _, key := range data.TypeValues() {
		keyStr := strings.ToLower(key.String())
		dataTypes[keyStr] = key
		keys = append(keys, keyStr)
	}

	dataTypeSetStr = strings.Join(keys, ", ")
}

// DataTypeSetStr returns comma separated string of allowed data types
func DataTypeSetStr() string {
	return dataTypeSetStr
}

// DataTypeFlag holds flag value and indicator if it was set from command line
type DataTypeFlag struct {
	set   bool
	value data.Type
}

// String returns string representation of the value
func (f *DataTypeFlag) String() string {
	return strings.ToLower(f.value.String())
}

// Set sets value from string representation
func (f *DataTypeFlag) Set(s string) (err error) {
	var ok bool
	f.value, ok = dataTypes[s]
	f.set = true
	if !ok {
		err = fmt.Errorf("unsupported data type of fact '%v', use one of: %v", s, dataTypeSetStr)
	}
	return
}

// IsSet returns true if value was set from command line
func (f *DataTypeFlag) IsSet() bool {
	return f.set
}

// GetValue returns flag value
func (f *DataTypeFlag) GetValue() data.Type {
	return f.value
}

// DataTypeFlagVar defines a data type flag with specified name, default value, and usage string.
// The argument value is the default value of the flag.
func DataTypeFlagVar(name string, value data.Type, usage string) *DataTypeFlag {
	v := &DataTypeFlag{value: value}
	flag.Var(v, name, usage)
	return v
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
