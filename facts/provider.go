package facts

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/monetha/go-verifiable-data/contracts"
	"github.com/monetha/go-verifiable-data/eth"
)

var (
	// ErrIsNotAllowedFactProvider error returned when fact provider is not allowed to modify (write or delete) passport facts.
	ErrIsNotAllowedFactProvider = errors.New("facts: not allowed fact provider")
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

func (p *Provider) initPassportLogicContractToModify(ctx context.Context, opts *bind.TransactOpts, passportAddress common.Address) (*contracts.PassportLogicContract, error) {
	c := contracts.InitPassportLogicContract(passportAddress, p.Backend)
	allowed, err := c.IsAllowedFactProvider(&bind.CallOpts{Context: ctx}, p.TransactOpts.From)
	if err != nil {
		return nil, fmt.Errorf("facts: IsAllowedFactProvider(): %v", err)
	}
	if !allowed {
		return nil, ErrIsNotAllowedFactProvider
	}

	return c, nil
}

// WriteTxData writes data for the specific key (uses transaction data)
func (p *Provider) WriteTxData(ctx context.Context, passportAddress common.Address, key [32]byte, data []byte) (common.Hash, error) {
	txHash, err := p.WriteTxDataNoWait(ctx, passportAddress, key, data)
	if err == nil {
		_, err = p.WaitForTxReceipt(ctx, txHash)
	}
	return txHash, err
}

// WriteTxDataNoWait writes data for the specific key (uses transaction data)
func (p *Provider) WriteTxDataNoWait(ctx context.Context, passportAddress common.Address, key [32]byte, data []byte) (common.Hash, error) {
	factProviderAuth := &p.TransactOpts

	c, err := p.initPassportLogicContractToModify(ctx, factProviderAuth, passportAddress)
	if err != nil {
		return common.Hash{}, err
	}

	p.Log("Writing tx data to passport", "fact_provider", factProviderAuth.From.Hex(), "key", key)
	tx, err := c.SetTxDataBlockNumber(factProviderAuth, key, data)
	if err != nil {
		return common.Hash{}, fmt.Errorf("facts: SetTxDataBlockNumber: %v", err)
	}
	return tx.Hash(), err
}

// WriteBytes writes data for the specific key (uses Ethereum storage)
func (p *Provider) WriteBytes(ctx context.Context, passportAddress common.Address, key [32]byte, data []byte) (common.Hash, error) {
	txHash, err := p.WriteBytesNoWait(ctx, passportAddress, key, data)
	if err == nil {
		_, err = p.WaitForTxReceipt(ctx, txHash)
	}
	return txHash, err
}

// WriteBytesNoWait writes data for the specific key (uses Ethereum storage).
// This method does not wait for the transaction to be mined. Use the method without the NoWait suffix if you need to make
// sure that the transaction was successfully mined.
func (p *Provider) WriteBytesNoWait(ctx context.Context, passportAddress common.Address, key [32]byte, data []byte) (common.Hash, error) {
	factProviderAuth := &p.TransactOpts

	c, err := p.initPassportLogicContractToModify(ctx, factProviderAuth, passportAddress)
	if err != nil {
		return common.Hash{}, err
	}

	p.Log("Writing bytes to passport", "fact_provider", factProviderAuth.From.Hex(), "key", key)
	tx, err := c.SetBytes(factProviderAuth, key, data)
	if err != nil {
		return common.Hash{}, fmt.Errorf("facts: SetBytes: %v", err)
	}
	return tx.Hash(), err
}

// WriteString writes data for the specific key (uses Ethereum storage)
func (p *Provider) WriteString(ctx context.Context, passportAddress common.Address, key [32]byte, data string) (common.Hash, error) {
	txHash, err := p.WriteStringNoWait(ctx, passportAddress, key, data)
	if err == nil {
		_, err = p.WaitForTxReceipt(ctx, txHash)
	}
	return txHash, err
}

// WriteStringNoWait writes data for the specific key (uses Ethereum storage).
// This method does not wait for the transaction to be mined. Use the method without the NoWait suffix if you need to make
// sure that the transaction was successfully mined.
func (p *Provider) WriteStringNoWait(ctx context.Context, passportAddress common.Address, key [32]byte, data string) (common.Hash, error) {
	factProviderAuth := &p.TransactOpts

	c, err := p.initPassportLogicContractToModify(ctx, factProviderAuth, passportAddress)
	if err != nil {
		return common.Hash{}, err
	}

	p.Log("Writing string to passport", "fact_provider", factProviderAuth.From.Hex(), "key", key)
	tx, err := c.SetString(factProviderAuth, key, data)
	if err != nil {
		return common.Hash{}, fmt.Errorf("facts: SetString: %v", err)
	}
	return tx.Hash(), err
}

// WriteAddress writes data for the specific key (uses Ethereum storage)
func (p *Provider) WriteAddress(ctx context.Context, passportAddress common.Address, key [32]byte, data common.Address) (common.Hash, error) {
	txHash, err := p.WriteAddressNoWait(ctx, passportAddress, key, data)
	if err == nil {
		_, err = p.WaitForTxReceipt(ctx, txHash)
	}
	return txHash, err
}

// WriteAddressNoWait writes data for the specific key (uses Ethereum storage).
// This method does not wait for the transaction to be mined. Use the method without the NoWait suffix if you need to make
// sure that the transaction was successfully mined.
func (p *Provider) WriteAddressNoWait(ctx context.Context, passportAddress common.Address, key [32]byte, data common.Address) (common.Hash, error) {
	factProviderAuth := &p.TransactOpts

	c, err := p.initPassportLogicContractToModify(ctx, factProviderAuth, passportAddress)
	if err != nil {
		return common.Hash{}, err
	}

	p.Log("Writing address to passport", "fact_provider", factProviderAuth.From.Hex(), "key", key)
	tx, err := c.SetAddress(factProviderAuth, key, data)
	if err != nil {
		return common.Hash{}, fmt.Errorf("facts: SetAddress: %v", err)
	}
	return tx.Hash(), err
}

// WriteUint writes data for the specific key (uses Ethereum storage)
func (p *Provider) WriteUint(ctx context.Context, passportAddress common.Address, key [32]byte, data *big.Int) (common.Hash, error) {
	txHash, err := p.WriteUintNoWait(ctx, passportAddress, key, data)
	if err == nil {
		_, err = p.WaitForTxReceipt(ctx, txHash)
	}
	return txHash, err
}

// WriteUintNoWait writes data for the specific key (uses Ethereum storage).
// This method does not wait for the transaction to be mined. Use the method without the NoWait suffix if you need to make
// sure that the transaction was successfully mined.
func (p *Provider) WriteUintNoWait(ctx context.Context, passportAddress common.Address, key [32]byte, data *big.Int) (common.Hash, error) {
	if data == nil {
		panic("big-integer cannot be nil")
	}

	if data.Sign() == -1 || data.BitLen() > 256 {
		return common.Hash{}, ErrOutOfUint256Range
	}

	data = new(big.Int).Set(data)

	factProviderAuth := &p.TransactOpts

	c, err := p.initPassportLogicContractToModify(ctx, factProviderAuth, passportAddress)
	if err != nil {
		return common.Hash{}, err
	}

	p.Log("Writing uint to passport", "fact_provider", factProviderAuth.From.Hex(), "key", key)
	tx, err := c.SetUint(factProviderAuth, key, data)
	if err != nil {
		return common.Hash{}, fmt.Errorf("facts: SetUint: %v", err)
	}
	return tx.Hash(), err
}

var (
	two255    = new(big.Int).Exp(big.NewInt(2), big.NewInt(255), nil) // = 2^255
	maxInt256 = new(big.Int).Sub(two255, big.NewInt(1))               // = 2^255-1
	minInt256 = new(big.Int).Neg(two255)                              // = -2^255
)

// WriteInt writes data for the specific key (uses Ethereum storage)
func (p *Provider) WriteInt(ctx context.Context, passportAddress common.Address, key [32]byte, data *big.Int) (common.Hash, error) {
	txHash, err := p.WriteIntNoWait(ctx, passportAddress, key, data)
	if err == nil {
		_, err = p.WaitForTxReceipt(ctx, txHash)
	}
	return txHash, err
}

// WriteIntNoWait writes data for the specific key (uses Ethereum storage).
// This method does not wait for the transaction to be mined. Use the method without the NoWait suffix if you need to make
// sure that the transaction was successfully mined.
func (p *Provider) WriteIntNoWait(ctx context.Context, passportAddress common.Address, key [32]byte, data *big.Int) (common.Hash, error) {
	if data == nil {
		panic("big-integer cannot be nil")
	}
	if data.Cmp(maxInt256) == 1 || minInt256.Cmp(data) == 1 {
		return common.Hash{}, ErrOutOfInt256Range
	}

	data = new(big.Int).Set(data)

	factProviderAuth := &p.TransactOpts

	c, err := p.initPassportLogicContractToModify(ctx, factProviderAuth, passportAddress)
	if err != nil {
		return common.Hash{}, err
	}

	p.Log("Writing int to passport", "fact_provider", factProviderAuth.From.Hex(), "key", key)
	tx, err := c.SetInt(factProviderAuth, key, data)
	if err != nil {
		return common.Hash{}, fmt.Errorf("facts: SetInt: %v", err)
	}
	return tx.Hash(), err
}

// WriteBool writes data for the specific key (uses Ethereum storage)
func (p *Provider) WriteBool(ctx context.Context, passportAddress common.Address, key [32]byte, data bool) (common.Hash, error) {
	txHash, err := p.WriteBoolNoWait(ctx, passportAddress, key, data)
	if err == nil {
		_, err = p.WaitForTxReceipt(ctx, txHash)
	}
	return txHash, err
}

// WriteBoolNoWait writes data for the specific key (uses Ethereum storage).
// This method does not wait for the transaction to be mined. Use the method without the NoWait suffix if you need to make
// sure that the transaction was successfully mined.
func (p *Provider) WriteBoolNoWait(ctx context.Context, passportAddress common.Address, key [32]byte, data bool) (common.Hash, error) {
	factProviderAuth := &p.TransactOpts

	c, err := p.initPassportLogicContractToModify(ctx, factProviderAuth, passportAddress)
	if err != nil {
		return common.Hash{}, err
	}

	p.Log("Writing bool to passport", "fact_provider", factProviderAuth.From.Hex(), "key", key)
	tx, err := c.SetBool(factProviderAuth, key, data)
	if err != nil {
		return common.Hash{}, fmt.Errorf("facts: SetBool: %v", err)
	}
	return tx.Hash(), err
}

// WriteIPFSHash writes IPFS hash for specific key (uses Ethereum storage to store the hash)
func (p *Provider) WriteIPFSHash(ctx context.Context, passportAddress common.Address, key [32]byte, hash string) (common.Hash, error) {
	txHash, err := p.WriteIPFSHashNoWait(ctx, passportAddress, key, hash)
	if err == nil {
		_, err = p.WaitForTxReceipt(ctx, txHash)
	}
	return txHash, err
}

// WriteIPFSHashNoWait writes IPFS hash for specific key (uses Ethereum storage to store the hash).
// This method does not wait for the transaction to be mined. Use the method without the NoWait suffix if you need to make
// sure that the transaction was successfully mined.
func (p *Provider) WriteIPFSHashNoWait(ctx context.Context, passportAddress common.Address, key [32]byte, hash string) (common.Hash, error) {
	factProviderAuth := &p.TransactOpts

	c, err := p.initPassportLogicContractToModify(ctx, factProviderAuth, passportAddress)
	if err != nil {
		return common.Hash{}, err
	}

	p.Log("Writing IPFS hash to passport", "fact_provider", factProviderAuth.From.Hex(), "key", key)
	tx, err := c.SetIPFSHash(factProviderAuth, key, hash)
	if err != nil {
		return common.Hash{}, fmt.Errorf("facts: SetIPFSHash: %v", err)
	}
	return tx.Hash(), err
}

// WritePrivateDataHashes writes IPFS hash of encrypted private data and hash of data encryption key
func (p *Provider) WritePrivateDataHashes(ctx context.Context, passportAddress common.Address, key [32]byte, privateData *PrivateDataHashes) (common.Hash, error) {
	txHash, err := p.WritePrivateDataHashesNoWait(ctx, passportAddress, key, privateData)
	if err == nil {
		_, err = p.WaitForTxReceipt(ctx, txHash)
	}
	return txHash, err
}

// WritePrivateDataHashesNoWait writes IPFS hash of encrypted private data and hash of data encryption key.
// This method does not wait for the transaction to be mined. Use the method without the NoWait suffix if you need to make
// sure that the transaction was successfully mined.
func (p *Provider) WritePrivateDataHashesNoWait(ctx context.Context, passportAddress common.Address, key [32]byte, privateData *PrivateDataHashes) (common.Hash, error) {
	factProviderAuth := &p.TransactOpts

	c, err := p.initPassportLogicContractToModify(ctx, factProviderAuth, passportAddress)
	if err != nil {
		return common.Hash{}, err
	}

	p.Log("Writing IPFS private data hashes to passport", "fact_provider", factProviderAuth.From.Hex(), "key", key)
	tx, err := c.SetPrivateDataHashes(factProviderAuth, key, privateData.DataIPFSHash, privateData.DataKeyHash)
	if err != nil {
		return common.Hash{}, fmt.Errorf("facts: SetPrivateDataHashes: %v", err)
	}
	return tx.Hash(), err
}

// DeleteTxData deletes tx data for the specific key
func (p *Provider) DeleteTxData(ctx context.Context, passportAddress common.Address, key [32]byte) (common.Hash, error) {
	txHash, err := p.DeleteTxDataNoWait(ctx, passportAddress, key)
	if err == nil {
		_, err = p.WaitForTxReceipt(ctx, txHash)
	}
	return txHash, err
}

// DeleteTxDataNoWait deletes tx data for the specific key.
// This method does not wait for the transaction to be mined. Use the method without the NoWait suffix if you need to make
// sure that the transaction was successfully mined.
func (p *Provider) DeleteTxDataNoWait(ctx context.Context, passportAddress common.Address, key [32]byte) (common.Hash, error) {
	factProviderAuth := &p.TransactOpts

	c, err := p.initPassportLogicContractToModify(ctx, factProviderAuth, passportAddress)
	if err != nil {
		return common.Hash{}, err
	}

	p.Log("Deleting tx data from passport", "fact_provider", factProviderAuth.From.Hex(), "key", key)
	tx, err := c.DeleteTxDataBlockNumber(factProviderAuth, key)
	if err != nil {
		return common.Hash{}, fmt.Errorf("facts: DeleteTxDataBlockNumber: %v", err)
	}
	return tx.Hash(), err
}

// DeleteBytes deletes bytes data for the specific key
func (p *Provider) DeleteBytes(ctx context.Context, passportAddress common.Address, key [32]byte) (common.Hash, error) {
	txHash, err := p.DeleteBytesNoWait(ctx, passportAddress, key)
	if err == nil {
		_, err = p.WaitForTxReceipt(ctx, txHash)
	}
	return txHash, err
}

// DeleteBytesNoWait deletes bytes data for the specific key.
// This method does not wait for the transaction to be mined. Use the method without the NoWait suffix if you need to make
// sure that the transaction was successfully mined.
func (p *Provider) DeleteBytesNoWait(ctx context.Context, passportAddress common.Address, key [32]byte) (common.Hash, error) {
	factProviderAuth := &p.TransactOpts

	c, err := p.initPassportLogicContractToModify(ctx, factProviderAuth, passportAddress)
	if err != nil {
		return common.Hash{}, err
	}

	p.Log("Deleting bytes from passport", "fact_provider", factProviderAuth.From.Hex(), "key", key)
	tx, err := c.DeleteBytes(factProviderAuth, key)
	if err != nil {
		return common.Hash{}, fmt.Errorf("facts: DeleteBytes: %v", err)
	}
	return tx.Hash(), err
}

// DeleteString deletes string data for the specific key
func (p *Provider) DeleteString(ctx context.Context, passportAddress common.Address, key [32]byte) (common.Hash, error) {
	txHash, err := p.DeleteStringNoWait(ctx, passportAddress, key)
	if err == nil {
		_, err = p.WaitForTxReceipt(ctx, txHash)
	}
	return txHash, err
}

// DeleteStringNoWait deletes string data for the specific key.
// This method does not wait for the transaction to be mined. Use the method without the NoWait suffix if you need to make
// sure that the transaction was successfully mined.
func (p *Provider) DeleteStringNoWait(ctx context.Context, passportAddress common.Address, key [32]byte) (common.Hash, error) {
	factProviderAuth := &p.TransactOpts

	c, err := p.initPassportLogicContractToModify(ctx, factProviderAuth, passportAddress)
	if err != nil {
		return common.Hash{}, err
	}

	p.Log("Deleting string from passport", "fact_provider", factProviderAuth.From.Hex(), "key", key)
	tx, err := c.DeleteString(factProviderAuth, key)
	if err != nil {
		return common.Hash{}, fmt.Errorf("facts: DeleteString: %v", err)
	}
	return tx.Hash(), err
}

// DeleteAddress deletes address data for the specific key
func (p *Provider) DeleteAddress(ctx context.Context, passportAddress common.Address, key [32]byte) (common.Hash, error) {
	txHash, err := p.DeleteAddressNoWait(ctx, passportAddress, key)
	if err == nil {
		_, err = p.WaitForTxReceipt(ctx, txHash)
	}
	return txHash, err
}

// DeleteAddressNoWait deletes address data for the specific key.
// This method does not wait for the transaction to be mined. Use the method without the NoWait suffix if you need to make
// sure that the transaction was successfully mined.
func (p *Provider) DeleteAddressNoWait(ctx context.Context, passportAddress common.Address, key [32]byte) (common.Hash, error) {
	factProviderAuth := &p.TransactOpts

	c, err := p.initPassportLogicContractToModify(ctx, factProviderAuth, passportAddress)
	if err != nil {
		return common.Hash{}, err
	}

	p.Log("Deleting address from passport", "fact_provider", factProviderAuth.From.Hex(), "key", key)
	tx, err := c.DeleteAddress(factProviderAuth, key)
	if err != nil {
		return common.Hash{}, fmt.Errorf("facts: DeleteAddress: %v", err)
	}

	return tx.Hash(), nil
}

// DeleteUint deletes uint data for the specific key
func (p *Provider) DeleteUint(ctx context.Context, passportAddress common.Address, key [32]byte) (common.Hash, error) {
	txHash, err := p.DeleteUintNoWait(ctx, passportAddress, key)
	if err == nil {
		_, err = p.WaitForTxReceipt(ctx, txHash)
	}
	return txHash, err
}

// DeleteUintNoWait deletes uint data for the specific key.
// This method does not wait for the transaction to be mined. Use the method without the NoWait suffix if you need to make
// sure that the transaction was successfully mined.
func (p *Provider) DeleteUintNoWait(ctx context.Context, passportAddress common.Address, key [32]byte) (common.Hash, error) {
	factProviderAuth := &p.TransactOpts
	c, err := p.initPassportLogicContractToModify(ctx, factProviderAuth, passportAddress)
	if err != nil {
		return common.Hash{}, err
	}
	p.Log("Deleting uint from passport", "fact_provider", factProviderAuth.From.Hex(), "key", key)
	tx, err := c.DeleteUint(factProviderAuth, key)
	if err != nil {
		return common.Hash{}, fmt.Errorf("facts: DeleteUint: %v", err)
	}
	return tx.Hash(), nil
}

// DeleteInt deletes int data for the specific key
func (p *Provider) DeleteInt(ctx context.Context, passportAddress common.Address, key [32]byte) (common.Hash, error) {
	txHash, err := p.DeleteIntNoWait(ctx, passportAddress, key)
	if err == nil {
		_, err = p.WaitForTxReceipt(ctx, txHash)
	}
	return txHash, err
}

// DeleteIntNoWait deletes int data for the specific key.
// This method does not wait for the transaction to be mined. Use the method without the NoWait suffix if you need to make
// sure that the transaction was successfully mined.
func (p *Provider) DeleteIntNoWait(ctx context.Context, passportAddress common.Address, key [32]byte) (common.Hash, error) {
	factProviderAuth := &p.TransactOpts

	c, err := p.initPassportLogicContractToModify(ctx, factProviderAuth, passportAddress)
	if err != nil {
		return common.Hash{}, err
	}

	p.Log("Deleting int from passport", "fact_provider", factProviderAuth.From.Hex(), "key", key)
	tx, err := c.DeleteInt(factProviderAuth, key)
	if err != nil {
		return common.Hash{}, fmt.Errorf("facts: DeleteInt: %v", err)
	}
	return tx.Hash(), err
}

// DeleteBool deletes bool data for the specific key
func (p *Provider) DeleteBool(ctx context.Context, passportAddress common.Address, key [32]byte) (common.Hash, error) {
	txHash, err := p.DeleteBoolNoWait(ctx, passportAddress, key)
	if err == nil {
		_, err = p.WaitForTxReceipt(ctx, txHash)
	}
	return txHash, err
}

// DeleteBoolNoWait deletes bool data for the specific key.
// This method does not wait for the transaction to be mined. Use the method without the NoWait suffix if you need to make
// sure that the transaction was successfully mined.
func (p *Provider) DeleteBoolNoWait(ctx context.Context, passportAddress common.Address, key [32]byte) (common.Hash, error) {
	factProviderAuth := &p.TransactOpts

	c, err := p.initPassportLogicContractToModify(ctx, factProviderAuth, passportAddress)
	if err != nil {
		return common.Hash{}, err
	}

	p.Log("Deleting bool from passport", "fact_provider", factProviderAuth.From.Hex(), "key", key)
	tx, err := c.DeleteBool(factProviderAuth, key)
	if err != nil {
		return common.Hash{}, fmt.Errorf("facts: DeleteBool: %v", err)
	}
	return tx.Hash(), err
}

// DeleteIPFSHash deletes IPFS hash for the specific key
func (p *Provider) DeleteIPFSHash(ctx context.Context, passportAddress common.Address, key [32]byte) (common.Hash, error) {
	txHash, err := p.DeleteIPFSHashNoWait(ctx, passportAddress, key)
	if err == nil {
		_, err = p.WaitForTxReceipt(ctx, txHash)
	}
	return txHash, err
}

// DeleteIPFSHashNoWait deletes IPFS hash for the specific key.
// This method does not wait for the transaction to be mined. Use the method without the NoWait suffix if you need to make
// sure that the transaction was successfully mined.
func (p *Provider) DeleteIPFSHashNoWait(ctx context.Context, passportAddress common.Address, key [32]byte) (common.Hash, error) {
	factProviderAuth := &p.TransactOpts

	c, err := p.initPassportLogicContractToModify(ctx, factProviderAuth, passportAddress)
	if err != nil {
		return common.Hash{}, err
	}

	p.Log("Deleting IPFS hash from passport", "fact_provider", factProviderAuth.From.Hex(), "key", key)
	tx, err := c.DeleteIPFSHash(factProviderAuth, key)
	if err != nil {
		return common.Hash{}, fmt.Errorf("facts: DeleteIPFSHash: %v", err)
	}
	return tx.Hash(), err
}

// DeletePrivateDataHashes deletes IPFS private data for the specific key
func (p *Provider) DeletePrivateDataHashes(ctx context.Context, passportAddress common.Address, key [32]byte) (common.Hash, error) {
	txHash, err := p.DeletePrivateDataHashesNoWait(ctx, passportAddress, key)
	if err == nil {
		_, err = p.WaitForTxReceipt(ctx, txHash)
	}
	return txHash, err
}

// DeletePrivateDataHashesNoWait deletes IPFS private data for the specific key
func (p *Provider) DeletePrivateDataHashesNoWait(ctx context.Context, passportAddress common.Address, key [32]byte) (common.Hash, error) {
	factProviderAuth := &p.TransactOpts

	c, err := p.initPassportLogicContractToModify(ctx, factProviderAuth, passportAddress)
	if err != nil {
		return common.Hash{}, err
	}

	p.Log("Deleting IPFS private data hashes from passport", "fact_provider", factProviderAuth.From.Hex(), "key", key)
	tx, err := c.DeletePrivateDataHashes(factProviderAuth, key)
	if err != nil {
		return common.Hash{}, fmt.Errorf("facts: DeletePrivateDataHashes: %v", err)
	}
	return tx.Hash(), err
}
