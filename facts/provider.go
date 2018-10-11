package facts

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"gitlab.com/monetha/protocol-go-sdk/contracts"
	"gitlab.com/monetha/protocol-go-sdk/eth"
)

var (
	// ErrOutOfUint256Range error returned when value is out of uint256 range [0; 2^256-1]
	ErrOutOfUint256Range = errors.New("facts: out of uint256 range")
	// ErrOutOfInt256Range error returned when value is out of int256 range [-2^255; 2^255-1]
	ErrOutOfInt256Range = errors.New("facts: out of int256 range")
)

// Provider provides the facts
type Provider eth.Session

// NewProvider converts session to Provider
func NewProvider(s *eth.Session) *Provider {
	return (*Provider)(s)
}

// WriteTxData writes data for the specific key (uses transaction data)
func (p *Provider) WriteTxData(ctx context.Context, passportAddress common.Address, key [32]byte, data []byte) error {
	backend := p.Backend
	factProviderAuth := &p.TransactOpts

	p.Log("Initialising passport", "passport", passportAddress)
	passportLogicContract, err := contracts.NewPassportLogicContract(passportAddress, backend)
	if err != nil {
		return fmt.Errorf("facts: NewPassportLogicContract: %v", err)
	}

	p.Log("Writing tx data to passport", "fact_provider", factProviderAuth.From.Hex(), "key", key)
	tx, err := passportLogicContract.SetTxDataBlockNumber(factProviderAuth, key, data)
	if err != nil {
		return fmt.Errorf("facts: SetTxDataBlockNumber: %v", err)
	}
	_, err = p.WaitForTxReceipt(ctx, tx.Hash())

	return err
}

// WriteBytes writes data for the specific key (uses Ethereum storage)
func (p *Provider) WriteBytes(ctx context.Context, passportAddress common.Address, key [32]byte, data []byte) error {
	backend := p.Backend
	factProviderAuth := &p.TransactOpts

	p.Log("Initialising passport", "passport", passportAddress)
	passportLogicContract, err := contracts.NewPassportLogicContract(passportAddress, backend)
	if err != nil {
		return fmt.Errorf("facts: NewPassportLogicContract: %v", err)
	}

	p.Log("Writing bytes to passport", "fact_provider", factProviderAuth.From.Hex(), "key", key)
	tx, err := passportLogicContract.SetBytes(factProviderAuth, key, data)
	if err != nil {
		return fmt.Errorf("facts: SetBytes: %v", err)
	}
	_, err = p.WaitForTxReceipt(ctx, tx.Hash())

	return err
}

// WriteString writes data for the specific key (uses Ethereum storage)
func (p *Provider) WriteString(ctx context.Context, passportAddress common.Address, key [32]byte, data string) error {
	backend := p.Backend
	factProviderAuth := &p.TransactOpts

	p.Log("Initialising passport", "passport", passportAddress)
	passportLogicContract, err := contracts.NewPassportLogicContract(passportAddress, backend)
	if err != nil {
		return fmt.Errorf("facts: NewPassportLogicContract: %v", err)
	}

	p.Log("Writing string to passport", "fact_provider", factProviderAuth.From.Hex(), "key", key)
	tx, err := passportLogicContract.SetString(factProviderAuth, key, data)
	if err != nil {
		return fmt.Errorf("facts: SetString: %v", err)
	}
	_, err = p.WaitForTxReceipt(ctx, tx.Hash())

	return err
}

// WriteAddress writes data for the specific key (uses Ethereum storage)
func (p *Provider) WriteAddress(ctx context.Context, passportAddress common.Address, key [32]byte, data common.Address) error {
	backend := p.Backend
	factProviderAuth := &p.TransactOpts

	p.Log("Initialising passport", "passport", passportAddress)
	passportLogicContract, err := contracts.NewPassportLogicContract(passportAddress, backend)
	if err != nil {
		return fmt.Errorf("facts: NewPassportLogicContract: %v", err)
	}

	p.Log("Writing address to passport", "fact_provider", factProviderAuth.From.Hex(), "key", key)
	tx, err := passportLogicContract.SetAddress(factProviderAuth, key, data)
	if err != nil {
		return fmt.Errorf("facts: SetAddress: %v", err)
	}
	_, err = p.WaitForTxReceipt(ctx, tx.Hash())

	return err
}

// WriteUint writes data for the specific key (uses Ethereum storage)
func (p *Provider) WriteUint(ctx context.Context, passportAddress common.Address, key [32]byte, data *big.Int) error {
	if data == nil {
		panic("big-integer cannot be nil")
	}

	if data.Sign() == -1 || data.BitLen() > 256 {
		return ErrOutOfUint256Range
	}

	data = new(big.Int).Set(data)

	backend := p.Backend
	factProviderAuth := &p.TransactOpts

	p.Log("Initialising passport", "passport", passportAddress)
	passportLogicContract, err := contracts.NewPassportLogicContract(passportAddress, backend)
	if err != nil {
		return fmt.Errorf("facts: NewPassportLogicContract: %v", err)
	}

	p.Log("Writing uint to passport", "fact_provider", factProviderAuth.From.Hex(), "key", key)
	tx, err := passportLogicContract.SetUint(factProviderAuth, key, data)
	if err != nil {
		return fmt.Errorf("facts: SetUint: %v", err)
	}
	_, err = p.WaitForTxReceipt(ctx, tx.Hash())

	return err
}

var (
	two255    = new(big.Int).Exp(big.NewInt(2), big.NewInt(255), nil) // = 2^255
	maxInt256 = new(big.Int).Sub(two255, big.NewInt(1))               // = 2^255-1
	minInt256 = new(big.Int).Neg(two255)                              // = -2^255
)

// WriteInt writes data for the specific key (uses Ethereum storage)
func (p *Provider) WriteInt(ctx context.Context, passportAddress common.Address, key [32]byte, data *big.Int) error {
	if data == nil {
		panic("big-integer cannot be nil")
	}
	if data.Cmp(maxInt256) == 1 || minInt256.Cmp(data) == 1 {
		return ErrOutOfInt256Range
	}

	data = new(big.Int).Set(data)

	backend := p.Backend
	factProviderAuth := &p.TransactOpts

	p.Log("Initialising passport", "passport", passportAddress)
	passportLogicContract, err := contracts.NewPassportLogicContract(passportAddress, backend)
	if err != nil {
		return fmt.Errorf("facts: NewPassportLogicContract: %v", err)
	}

	p.Log("Writing int to passport", "fact_provider", factProviderAuth.From.Hex(), "key", key)
	tx, err := passportLogicContract.SetInt(factProviderAuth, key, data)
	if err != nil {
		return fmt.Errorf("facts: SetInt: %v", err)
	}
	_, err = p.WaitForTxReceipt(ctx, tx.Hash())

	return err
}

// WriteBool writes data for the specific key (uses Ethereum storage)
func (p *Provider) WriteBool(ctx context.Context, passportAddress common.Address, key [32]byte, data bool) error {
	backend := p.Backend
	factProviderAuth := &p.TransactOpts

	p.Log("Initialising passport", "passport", passportAddress)
	passportLogicContract, err := contracts.NewPassportLogicContract(passportAddress, backend)
	if err != nil {
		return fmt.Errorf("facts: NewPassportLogicContract: %v", err)
	}

	p.Log("Writing bool to passport", "fact_provider", factProviderAuth.From.Hex(), "key", key)
	tx, err := passportLogicContract.SetBool(factProviderAuth, key, data)
	if err != nil {
		return fmt.Errorf("facts: SetBool: %v", err)
	}
	_, err = p.WaitForTxReceipt(ctx, tx.Hash())

	return err
}
