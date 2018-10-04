package facts

import (
	"context"
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
