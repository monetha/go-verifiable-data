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

// Reader reads the facts
type Reader eth.Eth

// NewReader converts eth to Reader
func NewReader(e *eth.Eth) *Reader {
	return (*Reader)(e)
}

// ReadTxData reads the data from the specific key of the given data provider.
// `ethereum.NotFound` error returned in case no value exists for the given key.
func (r *Reader) ReadTxData(ctx context.Context, passport common.Address, factProvider common.Address, key [32]byte) ([]byte, error) {
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

	params, err := txdata.NewPassportLogicInputParser().ParseSetTxDataBlockNumberCallData(tx.Data())
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
func (r *Reader) ReadBytes(ctx context.Context, passport common.Address, factProvider common.Address, key [32]byte) ([]byte, error) {
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
func (r *Reader) ReadString(ctx context.Context, passport common.Address, factProvider common.Address, key [32]byte) (string, error) {
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
func (r *Reader) ReadAddress(ctx context.Context, passport common.Address, factProvider common.Address, key [32]byte) (common.Address, error) {
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
func (r *Reader) ReadUint(ctx context.Context, passport common.Address, factProvider common.Address, key [32]byte) (*big.Int, error) {
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
func (r *Reader) ReadInt(ctx context.Context, passport common.Address, factProvider common.Address, key [32]byte) (*big.Int, error) {
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
func (r *Reader) ReadBool(ctx context.Context, passport common.Address, factProvider common.Address, key [32]byte) (bool, error) {
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
