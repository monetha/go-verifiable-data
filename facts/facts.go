package facts

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"gitlab.com/monetha/protocol-go-sdk/contracts"
	"gitlab.com/monetha/protocol-go-sdk/eth"
)

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

// WriteTxData writes data for the specific key
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

// ReadTxData reads the data from the specific key of the given data provider
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
		return nil, nil
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

	txData := tx.Data()

	// TODO: optimize parsing

	parsed, err := abi.JSON(strings.NewReader(contracts.PassportLogicContractABI))
	if err != nil {
		return nil, err
	}
	var v struct {
		Key  [32]byte
		Data []byte
	}
	err = customABI(parsed).UnpackInput(&v, "setTxDataBlockNumber", txData)
	if err != nil {
		return nil, err
	}
	if v.Key != key {
		return nil, fmt.Errorf("facts: parsed key parameter %v does not match provided key %v", v.Key, key)
	}

	return v.Data, nil
}

type customABI abi.ABI

// UnpackInput parses input according to the abi specification
func (abi customABI) UnpackInput(v interface{}, name string, input []byte) (err error) {
	if len(input) <= 4 {
		return fmt.Errorf("abi: insufficient input data")
	}

	input = input[4:] // remove signature

	// since there can't be naming collisions with contracts and events,
	// we need to decide whether we're calling a method or an event
	if method, ok := abi.Methods[name]; ok {
		if len(input)%32 != 0 {
			return fmt.Errorf("abi: improperly formatted input")
		}
		return method.Inputs.Unpack(v, input)
	} else if event, ok := abi.Events[name]; ok {
		return event.Inputs.Unpack(v, input)
	}
	return fmt.Errorf("abi: could not locate named method or event")
}
