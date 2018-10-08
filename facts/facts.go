package facts

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"gitlab.com/monetha/protocol-go-sdk/contracts"
	"gitlab.com/monetha/protocol-go-sdk/contracts/txdata"
	"gitlab.com/monetha/protocol-go-sdk/eth"
)

var (
	passportLogicInputParser *txdata.PassportLogicInputParser

	// ErrOutOfUint256Range error returned when value is out of uint256 range [0; 2^256-1]
	ErrOutOfUint256Range = errors.New("facts: out of uint256 range")
	// ErrOutOfInt256Range error returned when value is out of int256 range [-2^255; 2^255-1]
	ErrOutOfInt256Range = errors.New("facts: out of int256 range")
)

func init() {
	var err error
	passportLogicInputParser, err = txdata.NewPassportLogicInputParser()
	if err != nil {
		panic("facts: initializing PassportLogicInputParser: " + err.Error())
	}
}

// FactProvider provides the facts
type FactProvider eth.Session

// FactReader reads the facts
type FactReader eth.Eth

// NewProvider converts session to FactProvider
func NewProvider(s *eth.Session) *FactProvider {
	return (*FactProvider)(s)
}

// NewReader converts eth to FactReader
func NewReader(e *eth.Eth) *FactReader {
	return (*FactReader)(e)
}

// WriteTxData writes data for the specific key (uses transaction data)
func (p *FactProvider) WriteTxData(ctx context.Context, passportAddress common.Address, key [32]byte, data []byte) error {
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
func (p *FactProvider) WriteBytes(ctx context.Context, passportAddress common.Address, key [32]byte, data []byte) error {
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
func (p *FactProvider) WriteString(ctx context.Context, passportAddress common.Address, key [32]byte, data string) error {
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
func (p *FactProvider) WriteAddress(ctx context.Context, passportAddress common.Address, key [32]byte, data common.Address) error {
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
func (p *FactProvider) WriteUint(ctx context.Context, passportAddress common.Address, key [32]byte, data *big.Int) error {
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
func (p *FactProvider) WriteInt(ctx context.Context, passportAddress common.Address, key [32]byte, data *big.Int) error {
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
func (p *FactProvider) WriteBool(ctx context.Context, passportAddress common.Address, key [32]byte, data bool) error {
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

// ReadTxData reads the data from the specific key of the given data provider.
// `ethereum.NotFound` error returned in case no value exists for the given key.
func (r *FactReader) ReadTxData(ctx context.Context, passport common.Address, factProvider common.Address, key [32]byte) ([]byte, error) {
	backend := r.Backend

	(*eth.Eth)(r).Log("Initialising passport", "passport", passport)
	passportLogicContract, err := contracts.NewPassportLogicContract(passport, backend)
	if err != nil {
		return nil, fmt.Errorf("facts: NewPassportLogicContract: %v", err)
	}

	var bn struct {
		Success     bool
		BlockNumber *big.Int
	}

	(*eth.Eth)(r).Log("Getting block number for tx data", "fact_provider", factProvider, "key", key)
	bn, err = passportLogicContract.GetTxDataBlockNumber(nil, factProvider, key)
	if err != nil {
		return nil, fmt.Errorf("facts: GetTxDataBlockNumber: %v", err)
	}
	// check if block number exists for the key
	if !bn.Success {
		// no data
		return nil, ethereum.NotFound
	}

	blockNumber := bn.BlockNumber.Uint64()
	filterOpts := &bind.FilterOpts{
		Start:   blockNumber,
		End:     &blockNumber,
		Context: ctx,
	}
	(*eth.Eth)(r).Log("Filtering TxDataUpdated events", "fact_provider", factProvider, "key", key, "block_number", blockNumber)
	it, err := passportLogicContract.FilterTxDataUpdated(filterOpts, []common.Address{factProvider}, [][32]byte{key})
	if err != nil {
		return nil, fmt.Errorf("facts: FilterTxDataUpdated: %v", err)
	}
	defer it.Close()

	// reading the latest TxDataUpdated event
	var (
		txDataUpdatedEvent      contracts.PassportLogicContractTxDataUpdated
		txDataUpdatedEventFound bool
	)
	for it.Next() {
		if err = it.Error(); err != nil {
			return nil, err
		}

		ev := it.Event
		if ev == nil || ev.Raw.Removed {
			continue
		}

		txDataUpdatedEvent = *ev
		txDataUpdatedEventFound = true
	}

	if !txDataUpdatedEventFound {
		// no event found but it must exists
		return nil, fmt.Errorf("facts: no TxDataUpdated event found for fact provider (%v) and key (%v) in block %v", factProvider.Hex(), key, blockNumber)
	}

	txHash := txDataUpdatedEvent.Raw.TxHash
	(*eth.Eth)(r).Log("Getting transaction by hash", "tx_hash", txHash.Hex())
	tx, _, err := backend.TransactionByHash(ctx, txHash)
	if err != nil {
		return nil, fmt.Errorf("facts: TransactionByHash(%v): %v", txHash, err)
	}

	params, err := passportLogicInputParser.ParseSetTxDataBlockNumberCallData(tx.Data())
	if err != nil {
		return nil, err
	}
	if params.Key != key {
		return nil, fmt.Errorf("facts: parsed key parameter %v does not match provided key %v", params.Key, key)
	}

	return params.Data, nil
}

// ReadBytes reads the data from the specific key of the given data provider.
// `ethereum.NotFound` error returned in case no value exists for the given key.
func (r *FactReader) ReadBytes(ctx context.Context, passport common.Address, factProvider common.Address, key [32]byte) ([]byte, error) {
	backend := r.Backend

	(*eth.Eth)(r).Log("Initialising passport", "passport", passport)
	passportLogicContract, err := contracts.NewPassportLogicContract(passport, backend)
	if err != nil {
		return nil, fmt.Errorf("facts: NewPassportLogicContract: %v", err)
	}

	var res struct {
		Success bool
		Value   []byte
	}

	(*eth.Eth)(r).Log("Getting bytes", "fact_provider", factProvider, "key", key)
	res, err = passportLogicContract.GetBytes(nil, factProvider, key)
	if err != nil {
		return nil, fmt.Errorf("facts: GetBytes: %v", err)
	}
	// check if block number exists for the key
	if !res.Success {
		// no data
		return nil, ethereum.NotFound
	}

	return res.Value, nil
}

// ReadString reads the data from the specific key of the given data provider.
// `ethereum.NotFound` error returned in case no value exists for the given key.
func (r *FactReader) ReadString(ctx context.Context, passport common.Address, factProvider common.Address, key [32]byte) (string, error) {
	backend := r.Backend

	(*eth.Eth)(r).Log("Initialising passport", "passport", passport)
	passportLogicContract, err := contracts.NewPassportLogicContract(passport, backend)
	if err != nil {
		return "", fmt.Errorf("facts: NewPassportLogicContract: %v", err)
	}

	var res struct {
		Success bool
		Value   string
	}

	(*eth.Eth)(r).Log("Getting string", "fact_provider", factProvider, "key", key)
	res, err = passportLogicContract.GetString(nil, factProvider, key)
	if err != nil {
		return "", fmt.Errorf("facts: GetBytes: %v", err)
	}
	// check if block number exists for the key
	if !res.Success {
		// no data
		return "", ethereum.NotFound
	}

	return res.Value, nil
}

// ReadAddress reads the data from the specific key of the given data provider.
// `ethereum.NotFound` error returned in case no value exists for the given key.
func (r *FactReader) ReadAddress(ctx context.Context, passport common.Address, factProvider common.Address, key [32]byte) (common.Address, error) {
	backend := r.Backend

	(*eth.Eth)(r).Log("Initialising passport", "passport", passport)
	passportLogicContract, err := contracts.NewPassportLogicContract(passport, backend)
	if err != nil {
		return common.Address{}, fmt.Errorf("facts: NewPassportLogicContract: %v", err)
	}

	var res struct {
		Success bool
		Value   common.Address
	}

	(*eth.Eth)(r).Log("Getting address", "fact_provider", factProvider, "key", key)
	res, err = passportLogicContract.GetAddress(nil, factProvider, key)
	if err != nil {
		return common.Address{}, fmt.Errorf("facts: GetAddress: %v", err)
	}
	// check if block number exists for the key
	if !res.Success {
		// no data
		return common.Address{}, ethereum.NotFound
	}

	return res.Value, nil
}

// ReadUint reads the data from the specific key of the given data provider.
// `ethereum.NotFound` error returned in case no value exists for the given key.
func (r *FactReader) ReadUint(ctx context.Context, passport common.Address, factProvider common.Address, key [32]byte) (*big.Int, error) {
	backend := r.Backend

	(*eth.Eth)(r).Log("Initialising passport", "passport", passport)
	passportLogicContract, err := contracts.NewPassportLogicContract(passport, backend)
	if err != nil {
		return nil, fmt.Errorf("facts: NewPassportLogicContract: %v", err)
	}

	var res struct {
		Success bool
		Value   *big.Int
	}

	(*eth.Eth)(r).Log("Getting uint", "fact_provider", factProvider, "key", key)
	res, err = passportLogicContract.GetUint(nil, factProvider, key)
	if err != nil {
		return nil, fmt.Errorf("facts: GetUint: %v", err)
	}
	// check if block number exists for the key
	if !res.Success {
		// no data
		return nil, ethereum.NotFound
	}

	return res.Value, nil
}

// ReadInt reads the data from the specific key of the given data provider.
// `ethereum.NotFound` error returned in case no value exists for the given key.
func (r *FactReader) ReadInt(ctx context.Context, passport common.Address, factProvider common.Address, key [32]byte) (*big.Int, error) {
	backend := r.Backend

	(*eth.Eth)(r).Log("Initialising passport", "passport", passport)
	passportLogicContract, err := contracts.NewPassportLogicContract(passport, backend)
	if err != nil {
		return nil, fmt.Errorf("facts: NewPassportLogicContract: %v", err)
	}

	var res struct {
		Success bool
		Value   *big.Int
	}

	(*eth.Eth)(r).Log("Getting int", "fact_provider", factProvider, "key", key)
	res, err = passportLogicContract.GetInt(nil, factProvider, key)
	if err != nil {
		return nil, fmt.Errorf("facts: GetUint: %v", err)
	}
	// check if block number exists for the key
	if !res.Success {
		// no data
		return nil, ethereum.NotFound
	}

	return res.Value, nil
}

// ReadBool reads the data from the specific key of the given data provider.
// `ethereum.NotFound` error returned in case no value exists for the given key.
func (r *FactReader) ReadBool(ctx context.Context, passport common.Address, factProvider common.Address, key [32]byte) (bool, error) {
	backend := r.Backend

	(*eth.Eth)(r).Log("Initialising passport", "passport", passport)
	passportLogicContract, err := contracts.NewPassportLogicContract(passport, backend)
	if err != nil {
		return false, fmt.Errorf("facts: NewPassportLogicContract: %v", err)
	}

	var res struct {
		Success bool
		Value   bool
	}

	(*eth.Eth)(r).Log("Getting bool", "fact_provider", factProvider, "key", key)
	res, err = passportLogicContract.GetBool(nil, factProvider, key)
	if err != nil {
		return false, fmt.Errorf("facts: GetBool: %v", err)
	}
	// check if block number exists for the key
	if !res.Success {
		// no data
		return false, ethereum.NotFound
	}

	return res.Value, nil
}
