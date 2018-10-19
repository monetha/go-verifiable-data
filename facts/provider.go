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
func (p *Provider) WriteTxData(ctx context.Context, passportAddress common.Address, key [32]byte, data []byte) (common.Hash, error) {
	backend := p.Backend
	factProviderAuth := &p.TransactOpts

	p.Log("Writing tx data to passport", "fact_provider", factProviderAuth.From.Hex(), "key", key)
	tx, err := contracts.InitPassportLogicContract(passportAddress, backend).SetTxDataBlockNumber(factProviderAuth, key, data)
	if err != nil {
		return common.Hash{}, fmt.Errorf("facts: SetTxDataBlockNumber: %v", err)
	}
	txHash := tx.Hash()
	_, err = p.WaitForTxReceipt(ctx, txHash)

	return txHash, err
}

// WriteBytes writes data for the specific key (uses Ethereum storage)
func (p *Provider) WriteBytes(ctx context.Context, passportAddress common.Address, key [32]byte, data []byte) (common.Hash, error) {
	backend := p.Backend
	factProviderAuth := &p.TransactOpts

	p.Log("Writing bytes to passport", "fact_provider", factProviderAuth.From.Hex(), "key", key)
	tx, err := contracts.InitPassportLogicContract(passportAddress, backend).SetBytes(factProviderAuth, key, data)
	if err != nil {
		return common.Hash{}, fmt.Errorf("facts: SetBytes: %v", err)
	}
	txHash := tx.Hash()
	_, err = p.WaitForTxReceipt(ctx, txHash)

	return txHash, err
}

// WriteString writes data for the specific key (uses Ethereum storage)
func (p *Provider) WriteString(ctx context.Context, passportAddress common.Address, key [32]byte, data string) (common.Hash, error) {
	backend := p.Backend
	factProviderAuth := &p.TransactOpts

	p.Log("Writing string to passport", "fact_provider", factProviderAuth.From.Hex(), "key", key)
	tx, err := contracts.InitPassportLogicContract(passportAddress, backend).SetString(factProviderAuth, key, data)
	if err != nil {
		return common.Hash{}, fmt.Errorf("facts: SetString: %v", err)
	}
	txHash := tx.Hash()
	_, err = p.WaitForTxReceipt(ctx, txHash)

	return txHash, err
}

// WriteAddress writes data for the specific key (uses Ethereum storage)
func (p *Provider) WriteAddress(ctx context.Context, passportAddress common.Address, key [32]byte, data common.Address) (common.Hash, error) {
	backend := p.Backend
	factProviderAuth := &p.TransactOpts

	p.Log("Writing address to passport", "fact_provider", factProviderAuth.From.Hex(), "key", key)
	tx, err := contracts.InitPassportLogicContract(passportAddress, backend).SetAddress(factProviderAuth, key, data)
	if err != nil {
		return common.Hash{}, fmt.Errorf("facts: SetAddress: %v", err)
	}
	txHash := tx.Hash()
	_, err = p.WaitForTxReceipt(ctx, txHash)

	return txHash, err
}

// WriteUint writes data for the specific key (uses Ethereum storage)
func (p *Provider) WriteUint(ctx context.Context, passportAddress common.Address, key [32]byte, data *big.Int) (common.Hash, error) {
	if data == nil {
		panic("big-integer cannot be nil")
	}

	if data.Sign() == -1 || data.BitLen() > 256 {
		return common.Hash{}, ErrOutOfUint256Range
	}

	data = new(big.Int).Set(data)

	backend := p.Backend
	factProviderAuth := &p.TransactOpts

	p.Log("Writing uint to passport", "fact_provider", factProviderAuth.From.Hex(), "key", key)
	tx, err := contracts.InitPassportLogicContract(passportAddress, backend).SetUint(factProviderAuth, key, data)
	if err != nil {
		return common.Hash{}, fmt.Errorf("facts: SetUint: %v", err)
	}
	txHash := tx.Hash()
	_, err = p.WaitForTxReceipt(ctx, txHash)

	return txHash, err
}

var (
	two255    = new(big.Int).Exp(big.NewInt(2), big.NewInt(255), nil) // = 2^255
	maxInt256 = new(big.Int).Sub(two255, big.NewInt(1))               // = 2^255-1
	minInt256 = new(big.Int).Neg(two255)                              // = -2^255
)

// WriteInt writes data for the specific key (uses Ethereum storage)
func (p *Provider) WriteInt(ctx context.Context, passportAddress common.Address, key [32]byte, data *big.Int) (common.Hash, error) {
	if data == nil {
		panic("big-integer cannot be nil")
	}
	if data.Cmp(maxInt256) == 1 || minInt256.Cmp(data) == 1 {
		return common.Hash{}, ErrOutOfInt256Range
	}

	data = new(big.Int).Set(data)

	backend := p.Backend
	factProviderAuth := &p.TransactOpts

	p.Log("Writing int to passport", "fact_provider", factProviderAuth.From.Hex(), "key", key)
	tx, err := contracts.InitPassportLogicContract(passportAddress, backend).SetInt(factProviderAuth, key, data)
	if err != nil {
		return common.Hash{}, fmt.Errorf("facts: SetInt: %v", err)
	}
	txHash := tx.Hash()
	_, err = p.WaitForTxReceipt(ctx, txHash)

	return txHash, err
}

// WriteBool writes data for the specific key (uses Ethereum storage)
func (p *Provider) WriteBool(ctx context.Context, passportAddress common.Address, key [32]byte, data bool) (common.Hash, error) {
	backend := p.Backend
	factProviderAuth := &p.TransactOpts

	p.Log("Writing bool to passport", "fact_provider", factProviderAuth.From.Hex(), "key", key)
	tx, err := contracts.InitPassportLogicContract(passportAddress, backend).SetBool(factProviderAuth, key, data)
	if err != nil {
		return common.Hash{}, fmt.Errorf("facts: SetBool: %v", err)
	}
	txHash := tx.Hash()
	_, err = p.WaitForTxReceipt(ctx, txHash)

	return txHash, err
}

// DeleteTxData deletes tx data for the specific key
func (p *Provider) DeleteTxData(ctx context.Context, passportAddress common.Address, key [32]byte) (common.Hash, error) {
	backend := p.Backend
	factProviderAuth := &p.TransactOpts

	p.Log("Deleting tx data from passport", "fact_provider", factProviderAuth.From.Hex(), "key", key)
	tx, err := contracts.InitPassportLogicContract(passportAddress, backend).DeleteTxDataBlockNumber(factProviderAuth, key)
	if err != nil {
		return common.Hash{}, fmt.Errorf("facts: DeleteTxDataBlockNumber: %v", err)
	}
	txHash := tx.Hash()
	_, err = p.WaitForTxReceipt(ctx, txHash)

	return txHash, err
}

// DeleteBytes deletes bytes data for the specific key
func (p *Provider) DeleteBytes(ctx context.Context, passportAddress common.Address, key [32]byte) (common.Hash, error) {
	backend := p.Backend
	factProviderAuth := &p.TransactOpts

	p.Log("Deleting bytes from passport", "fact_provider", factProviderAuth.From.Hex(), "key", key)
	tx, err := contracts.InitPassportLogicContract(passportAddress, backend).DeleteBytes(factProviderAuth, key)
	if err != nil {
		return common.Hash{}, fmt.Errorf("facts: DeleteBytes: %v", err)
	}
	txHash := tx.Hash()
	_, err = p.WaitForTxReceipt(ctx, txHash)

	return txHash, err
}

// DeleteString deletes string data for the specific key
func (p *Provider) DeleteString(ctx context.Context, passportAddress common.Address, key [32]byte) (common.Hash, error) {
	backend := p.Backend
	factProviderAuth := &p.TransactOpts

	p.Log("Deleting string from passport", "fact_provider", factProviderAuth.From.Hex(), "key", key)
	tx, err := contracts.InitPassportLogicContract(passportAddress, backend).DeleteString(factProviderAuth, key)
	if err != nil {
		return common.Hash{}, fmt.Errorf("facts: DeleteString: %v", err)
	}
	txHash := tx.Hash()
	_, err = p.WaitForTxReceipt(ctx, txHash)

	return txHash, err
}

// DeleteAddress deletes address data for the specific key
func (p *Provider) DeleteAddress(ctx context.Context, passportAddress common.Address, key [32]byte) (common.Hash, error) {
	backend := p.Backend
	factProviderAuth := &p.TransactOpts

	p.Log("Deleting address from passport", "fact_provider", factProviderAuth.From.Hex(), "key", key)
	tx, err := contracts.InitPassportLogicContract(passportAddress, backend).DeleteAddress(factProviderAuth, key)
	if err != nil {
		return common.Hash{}, fmt.Errorf("facts: DeleteAddress: %v", err)
	}
	txHash := tx.Hash()
	_, err = p.WaitForTxReceipt(ctx, txHash)

	return txHash, err
}

// DeleteUint deletes uint data for the specific key
func (p *Provider) DeleteUint(ctx context.Context, passportAddress common.Address, key [32]byte) (common.Hash, error) {
	backend := p.Backend
	factProviderAuth := &p.TransactOpts

	p.Log("Deleting uint from passport", "fact_provider", factProviderAuth.From.Hex(), "key", key)
	tx, err := contracts.InitPassportLogicContract(passportAddress, backend).DeleteUint(factProviderAuth, key)
	if err != nil {
		return common.Hash{}, fmt.Errorf("facts: DeleteUint: %v", err)
	}
	txHash := tx.Hash()
	_, err = p.WaitForTxReceipt(ctx, txHash)

	return txHash, err
}

// DeleteInt deletes int data for the specific key
func (p *Provider) DeleteInt(ctx context.Context, passportAddress common.Address, key [32]byte) (common.Hash, error) {
	backend := p.Backend
	factProviderAuth := &p.TransactOpts

	p.Log("Deleting int from passport", "fact_provider", factProviderAuth.From.Hex(), "key", key)
	tx, err := contracts.InitPassportLogicContract(passportAddress, backend).DeleteInt(factProviderAuth, key)
	if err != nil {
		return common.Hash{}, fmt.Errorf("facts: DeleteInt: %v", err)
	}
	txHash := tx.Hash()
	_, err = p.WaitForTxReceipt(ctx, txHash)

	return txHash, err
}

// DeleteBool deletes bool data for the specific key
func (p *Provider) DeleteBool(ctx context.Context, passportAddress common.Address, key [32]byte) (common.Hash, error) {
	backend := p.Backend
	factProviderAuth := &p.TransactOpts

	p.Log("Deleting bool from passport", "fact_provider", factProviderAuth.From.Hex(), "key", key)
	tx, err := contracts.InitPassportLogicContract(passportAddress, backend).DeleteBool(factProviderAuth, key)
	if err != nil {
		return common.Hash{}, fmt.Errorf("facts: DeleteBool: %v", err)
	}
	txHash := tx.Hash()
	_, err = p.WaitForTxReceipt(ctx, txHash)

	return txHash, err
}
