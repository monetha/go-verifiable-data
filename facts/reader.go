package facts

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/monetha/go-verifiable-data/contracts"
	"github.com/monetha/go-verifiable-data/contracts/txdata"
	"github.com/monetha/go-verifiable-data/eth"
)

var (
	// ErrOwnerMustClaimOwnership error returned when the owner hasn't claimed ownership of a passport, but an attempt is
	// made to retrieve the public address of the owner
	ErrOwnerMustClaimOwnership = errors.New("facts: owner must claim ownership of the passport")
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

	passportLogicContract := contracts.InitPassportLogicContract(passport, backend)
	var bn struct {
		Success     bool
		BlockNumber *big.Int
	}

	(*eth.Eth)(r).Log("Getting block number for tx data", "fact_provider", factProvider, "key", key)
	bn, err := passportLogicContract.GetTxDataBlockNumber(&bind.CallOpts{Context: ctx}, factProvider, key)
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
		return nil, fmt.Errorf("facts: TransactionByHash(%v): %v", txHash.String(), err)
	}

	if r.TxDecrypter != nil {
		tx, err = r.TxDecrypter(ctx, tx)
		if err != nil {
			return nil, fmt.Errorf("facts: TxDecrypter(%v): %v", txHash.String(), err)
		}
	}

	params, err := txdata.ParseSetTxDataBlockNumberCallData(tx.Data())
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

	var res struct {
		Success bool
		Value   []byte
	}

	(*eth.Eth)(r).Log("Getting bytes", "fact_provider", factProvider, "key", key)
	res, err := contracts.InitPassportLogicContract(passport, backend).GetBytes(&bind.CallOpts{Context: ctx}, factProvider, key)
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

	var res struct {
		Success bool
		Value   string
	}

	(*eth.Eth)(r).Log("Getting string", "fact_provider", factProvider, "key", key)
	res, err := contracts.InitPassportLogicContract(passport, backend).GetString(&bind.CallOpts{Context: ctx}, factProvider, key)
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

	var res struct {
		Success bool
		Value   common.Address
	}

	(*eth.Eth)(r).Log("Getting address", "fact_provider", factProvider, "key", key)
	res, err := contracts.InitPassportLogicContract(passport, backend).GetAddress(&bind.CallOpts{Context: ctx}, factProvider, key)
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

	var res struct {
		Success bool
		Value   *big.Int
	}

	(*eth.Eth)(r).Log("Getting uint", "fact_provider", factProvider, "key", key)
	res, err := contracts.InitPassportLogicContract(passport, backend).GetUint(&bind.CallOpts{Context: ctx}, factProvider, key)
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

	var res struct {
		Success bool
		Value   *big.Int
	}

	(*eth.Eth)(r).Log("Getting int", "fact_provider", factProvider, "key", key)
	res, err := contracts.InitPassportLogicContract(passport, backend).GetInt(&bind.CallOpts{Context: ctx}, factProvider, key)
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

	var res struct {
		Success bool
		Value   bool
	}

	(*eth.Eth)(r).Log("Getting bool", "fact_provider", factProvider, "key", key)
	res, err := contracts.InitPassportLogicContract(passport, backend).GetBool(&bind.CallOpts{Context: ctx}, factProvider, key)
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

// ReadIPFSHash reads the IPFS hash from the specific key of the given data provider.
// `ethereum.NotFound` error returned in case no value exists for the given key.
func (r *Reader) ReadIPFSHash(ctx context.Context, passport common.Address, factProvider common.Address, key [32]byte) (string, error) {
	backend := r.Backend

	var res struct {
		Success bool
		Value   string
	}

	(*eth.Eth)(r).Log("Getting IPFS hash", "fact_provider", factProvider, "key", key)
	res, err := contracts.InitPassportLogicContract(passport, backend).GetIPFSHash(&bind.CallOpts{Context: ctx}, factProvider, key)
	if err != nil {
		return "", fmt.Errorf("facts: GetBool: %v", err)
	}
	// check if block number exists for the key
	if !res.Success {
		// no data
		return "", ethereum.NotFound
	}

	return res.Value, nil
}

// ReadPrivateDataHashes reads the private data from the specific key of the given data provider.
// `ethereum.NotFound` error returned in case no value exists for the given key.
func (r *Reader) ReadPrivateDataHashes(ctx context.Context, passport common.Address, factProvider common.Address, key [32]byte) (*PrivateDataHashes, error) {
	backend := r.Backend

	var res struct {
		Success      bool
		DataIPFSHash string
		DataKeyHash  [32]byte
	}

	(*eth.Eth)(r).Log("Getting IPFS private data hashes", "fact_provider", factProvider, "key", key)
	res, err := contracts.InitPassportLogicContract(passport, backend).GetPrivateDataHashes(&bind.CallOpts{Context: ctx}, factProvider, key)
	if err != nil {
		return nil, fmt.Errorf("facts: GetPrivateDataHashes: %v", err)
	}
	// check if block number exists for the key
	if !res.Success {
		// no data
		return nil, ethereum.NotFound
	}

	return &PrivateDataHashes{DataIPFSHash: res.DataIPFSHash, DataKeyHash: res.DataKeyHash}, nil
}

// ReadOwnerPublicKey reads the Ethereum public key of the passport owner. Owner must claim ownership of the passport,
// before this method can be invoked.
func (r *Reader) ReadOwnerPublicKey(ctx context.Context, passport common.Address) (pk *ecdsa.PublicKey, err error) {
	backend := r.Backend

	plc := contracts.InitPassportLogicContract(passport, backend)

	newOwner, err := plc.Owner(nil)
	if err != nil {
		err = fmt.Errorf("facts: getting Owner: %v", err)
		return
	}

	(*eth.Eth)(r).Log("Filtering OwnershipTransferred", "newOwner", newOwner)
	ownershipTransferredEvent, err := (*extPassportLogicContract)(plc).GetFirstOwnershipTransferredEvent(ctx, newOwner)
	if err != nil {
		err = fmt.Errorf("facts: FilterOwnershipTransferred: %v", err)
		return
	}

	txHash := ownershipTransferredEvent.Raw.TxHash
	(*eth.Eth)(r).Log("Getting transaction by hash", "tx_hash", txHash.Hex())
	tx, _, err := backend.TransactionByHash(ctx, txHash)
	if err != nil {
		err = fmt.Errorf("facts: TransactionByHash(%v): %v", txHash.String(), err)
		return
	}

	if r.SignedTxRestorer != nil {
		tx, err = r.SignedTxRestorer(ctx, tx)
		if err != nil {
			err = fmt.Errorf("facts: SignedTxRestorer(%v): %v", txHash.String(), err)
			return
		}
	}

	return (*eth.Transaction)(tx).GetSenderPublicKey()
}

type extPassportLogicContract contracts.PassportLogicContract

func (c *extPassportLogicContract) GetFirstOwnershipTransferredEvent(ctx context.Context, newOwner common.Address) (ev *contracts.PassportLogicContractOwnershipTransferred, err error) {
	filterOpts := &bind.FilterOpts{
		Context: ctx,
	}
	it, err := c.FilterOwnershipTransferred(filterOpts, nil, []common.Address{newOwner})
	if err != nil {
		return
	}
	defer it.Close()

	for it.Next() {
		if err = it.Error(); err != nil {
			return
		}

		ev := it.Event
		if ev == nil || ev.Raw.Removed {
			continue
		}

		return ev, nil
	}

	return nil, ErrOwnerMustClaimOwnership
}
