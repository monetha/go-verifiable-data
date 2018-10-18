package facts

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"gitlab.com/monetha/protocol-go-sdk/contracts"
	"gitlab.com/monetha/protocol-go-sdk/eth"
	"gitlab.com/monetha/protocol-go-sdk/types/change"
	"gitlab.com/monetha/protocol-go-sdk/types/data"
)

var (
	passportLogicContractABI abi.ABI
)

func init() {
	var err error
	passportLogicContractABI, err = abi.JSON(strings.NewReader(contracts.PassportLogicContractABI))
	if err != nil {
		panic("facts: initializing PassportLogicContractABI: " + err.Error())
	}
}

// Historian reads the fact history
type Historian eth.Eth

// NewHistorian converts eth to Historian
func NewHistorian(e *eth.Eth) *Historian {
	return (*Historian)(e)
}

// Change contains information about data change
type Change struct {
	// ChangeType represents type of data change
	ChangeType change.Type
	// DataType represents type of data
	DataType data.Type
	// Address of fact provider
	FactProvider common.Address
	// Key of the value
	Key [32]byte
	// Blockchain specific contextual infos
	Raw types.Log
}

type eventMetaInfo struct {
	EventName  string
	ChangeType change.Type
	DataType   data.Type
}

// ChangeIterator is returned from FilterChanges and is used to iterate over the changes and unpacked data for
// Updated and Deleted events raised by the PassportLogic contract.
type ChangeIterator struct {
	Change *Change // Change containing the info of the last retrieved change

	contract *bind.BoundContract // Generic contract to use for unpacking event data

	eventMetaInfos map[common.Hash]eventMetaInfo // event meta information by topic id

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}

	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			if err := it.parseEvent(log); err != nil {
				it.fail = err
				return false
			}
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		if err := it.parseEvent(log); err != nil {
			it.fail = err
			return false
		}
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *ChangeIterator) parseEvent(log types.Log) error {
	emi := it.eventMetaInfos[log.Topics[0]]
	it.Change = &Change{
		ChangeType: emi.ChangeType,
		DataType:   emi.DataType,
		Raw:        log,
	}
	event := &struct {
		FactProvider common.Address
		Key          [32]byte
	}{}
	if err := it.contract.UnpackLog(event, emi.EventName, log); err != nil {
		return err
	}
	it.Change.FactProvider = event.FactProvider
	it.Change.Key = event.Key
	return nil
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ToSlice retrieves all changes and saves them into slice.
func (it *ChangeIterator) ToSlice() (ps []*Change, err error) {
	defer func() {
		if cErr := it.Close(); err == nil && cErr != nil {
			err = cErr
		}
	}()

	for it.Next() {
		if err := it.Error(); err != nil {
			return nil, err
		}
		ps = append(ps, it.Change)
	}

	return
}

// ChangesFilterOpts is the collection of options to fine tune filtering for changes.
type ChangesFilterOpts struct {
	Start        uint64           // Start of the queried range
	End          *uint64          // End of the range (nil = latest)
	ChangeType   []change.Type    // ChangeType is a slice of change types to filter (nil = all change types)
	DataType     []data.Type      // DataType is a slice of data types to filter (nil = all data types)
	FactProvider []common.Address // FactProvider is a slice of fact providers to filter (nil = all fact providers)
	Key          [][32]byte       // Key is a slice of keys to filter (nil = all keys)
	Context      context.Context  // Network context to support cancellation and timeouts (nil = no timeout)
}

var (
	eventNameByChangeTypeDataType = map[change.Type]map[data.Type]string{
		change.Deleted: {
			data.Address: "AddressDeleted",
			data.Bool:    "BoolDeleted",
			data.Bytes:   "BytesDeleted",
			data.Int:     "IntDeleted",
			data.String:  "StringDeleted",
			data.TxData:  "TxDataDeleted",
			data.Uint:    "UintDeleted",
		},
		change.Updated: {
			data.Address: "AddressUpdated",
			data.Bool:    "BoolUpdated",
			data.Bytes:   "BytesUpdated",
			data.Int:     "IntUpdated",
			data.String:  "StringUpdated",
			data.TxData:  "TxDataUpdated",
			data.Uint:    "UintUpdated",
		},
	}
)

type changeTypeSet map[change.Type]struct{}

func (s *changeTypeSet) Add(t change.Type) {
	if *s == nil {
		*s = changeTypeSet{t: struct{}{}}
		return
	}
	(*s)[t] = struct{}{}
}

func (s *changeTypeSet) IsEmpty() bool {
	return len(*s) == 0
}

func (s *changeTypeSet) Contains(t change.Type) (ok bool) {
	if *s != nil {
		_, ok = (*s)[t]
	}
	return
}

type dataTypeSet map[data.Type]struct{}

func (s *dataTypeSet) Add(t data.Type) {
	if *s == nil {
		*s = dataTypeSet{t: struct{}{}}
		return
	}
	(*s)[t] = struct{}{}
}

func (s *dataTypeSet) IsEmpty() bool {
	return len(*s) == 0
}

func (s *dataTypeSet) Contains(t data.Type) (ok bool) {
	if *s != nil {
		_, ok = (*s)[t]
	}
	return
}

// FilterChanges retrieves changes from event log of passport.
func (h *Historian) FilterChanges(opts *ChangesFilterOpts, passportAddress common.Address) (*ChangeIterator, error) {
	if opts == nil {
		opts = &ChangesFilterOpts{}
	}

	// prepare filter sets
	var (
		changeTypes changeTypeSet
		dataTypes   dataTypeSet
	)
	for _, value := range opts.ChangeType {
		changeTypes.Add(value)
	}
	for _, value := range opts.DataType {
		dataTypes.Add(value)
	}

	// prepare event names
	var eventNames []string
	eventMetaInfos := make(map[common.Hash]eventMetaInfo)
	for changeType, eventNamesByDataType := range eventNameByChangeTypeDataType {

		if changeTypes.IsEmpty() || changeTypes.Contains(changeType) {
			for dataType, eventName := range eventNamesByDataType {

				if dataTypes.IsEmpty() || dataTypes.Contains(dataType) {
					eventNames = append(eventNames, eventName)

					eventID := passportLogicContractABI.Events[eventName].Id()
					eventMetaInfos[eventID] = eventMetaInfo{
						EventName:  eventName,
						ChangeType: changeType,
						DataType:   dataType,
					}
				}
			}
		}
	}

	// create log filterer
	logFilterer := eth.NewContractLogFilterer(passportAddress, passportLogicContractABI, h.Backend)

	var factProviderRule []interface{}
	for _, factProviderItem := range opts.FactProvider {
		factProviderRule = append(factProviderRule, factProviderItem)
	}
	var keyRule []interface{}
	for _, keyItem := range opts.Key {
		keyRule = append(keyRule, keyItem)
	}

	filterOpts := &bind.FilterOpts{
		Start:   opts.Start,
		End:     opts.End,
		Context: opts.Context,
	}

	logs, sub, err := logFilterer.FilterLogs(filterOpts, eventNames, factProviderRule, keyRule)
	if err != nil {
		return nil, err
	}

	boundContract := bind.NewBoundContract(passportAddress, passportLogicContractABI, h.Backend, h.Backend, h.Backend)
	return &ChangeIterator{logs: logs, sub: sub, contract: boundContract, eventMetaInfos: eventMetaInfos}, nil
}

type (
	// WriteTxDataHistoryItem holds parameters of WriteTxData call
	WriteTxDataHistoryItem struct {
		FactProvider common.Address
		Key          [32]byte
		Data         []byte
	}

	// WriteBytesHistoryItem holds parameters of WriteBytes call
	WriteBytesHistoryItem struct {
		FactProvider common.Address
		Key          [32]byte
		Data         []byte
	}

	// WriteStringHistoryItem holds parameters of WriteString call
	WriteStringHistoryItem struct {
		FactProvider common.Address
		Key          [32]byte
		Data         string
	}

	// WriteAddressHistoryItem holds parameters of WriteAddress call
	WriteAddressHistoryItem struct {
		FactProvider common.Address
		Key          [32]byte
		Data         common.Address
	}

	// WriteUintHistoryItem holds parameters of WriteUint call
	WriteUintHistoryItem struct {
		FactProvider common.Address
		Key          [32]byte
		Data         *big.Int
	}

	// WriteIntHistoryItem holds parameters of WriteInt call
	WriteIntHistoryItem struct {
		FactProvider common.Address
		Key          [32]byte
		Data         *big.Int
	}

	// WriteBoolHistoryItem holds parameters of WriteBool call
	WriteBoolHistoryItem struct {
		FactProvider common.Address
		Key          [32]byte
		Data         bool
	}
)

// GetHistoryItemOfWriteTxData returns the data value that was set in the given transaction.
func (h *Historian) GetHistoryItemOfWriteTxData(ctx context.Context, passport common.Address, txHash common.Hash) (*WriteTxDataHistoryItem, error) {
	from, txData, err := h.getTransactionSenderData(ctx, txHash)
	if err != nil {
		return nil, err
	}

	params, err := passportLogicInputParser.ParseSetTxDataBlockNumberCallData(txData)
	if err != nil {
		return nil, err
	}

	return &WriteTxDataHistoryItem{
		FactProvider: from,
		Key:          params.Key,
		Data:         params.Data,
	}, nil
}

// GetHistoryItemOfWriteBytes returns the data value that was set in the given transaction.
func (h *Historian) GetHistoryItemOfWriteBytes(ctx context.Context, passport common.Address, txHash common.Hash) (*WriteBytesHistoryItem, error) {
	from, txData, err := h.getTransactionSenderData(ctx, txHash)
	if err != nil {
		return nil, err
	}

	params, err := passportLogicInputParser.ParseSetBytesCallData(txData)
	if err != nil {
		return nil, err
	}

	return &WriteBytesHistoryItem{
		FactProvider: from,
		Key:          params.Key,
		Data:         params.Data,
	}, nil
}

// GetHistoryItemOfWriteString returns the data value that was set in the given transaction.
func (h *Historian) GetHistoryItemOfWriteString(ctx context.Context, passport common.Address, txHash common.Hash) (*WriteStringHistoryItem, error) {
	from, txData, err := h.getTransactionSenderData(ctx, txHash)
	if err != nil {
		return nil, err
	}

	params, err := passportLogicInputParser.ParseSetStringCallData(txData)
	if err != nil {
		return nil, err
	}

	return &WriteStringHistoryItem{
		FactProvider: from,
		Key:          params.Key,
		Data:         params.Data,
	}, nil
}

// GetHistoryItemOfWriteAddress returns the data value that was set in the given transaction.
func (h *Historian) GetHistoryItemOfWriteAddress(ctx context.Context, passport common.Address, txHash common.Hash) (*WriteAddressHistoryItem, error) {
	from, txData, err := h.getTransactionSenderData(ctx, txHash)
	if err != nil {
		return nil, err
	}

	params, err := passportLogicInputParser.ParseSetAddressCallData(txData)
	if err != nil {
		return nil, err
	}

	return &WriteAddressHistoryItem{
		FactProvider: from,
		Key:          params.Key,
		Data:         params.Data,
	}, nil
}

// GetHistoryItemOfWriteUint returns the data value that was set in the given transaction.
func (h *Historian) GetHistoryItemOfWriteUint(ctx context.Context, passport common.Address, txHash common.Hash) (*WriteUintHistoryItem, error) {
	from, txData, err := h.getTransactionSenderData(ctx, txHash)
	if err != nil {
		return nil, err
	}

	params, err := passportLogicInputParser.ParseSetUintCallData(txData)
	if err != nil {
		return nil, err
	}

	return &WriteUintHistoryItem{
		FactProvider: from,
		Key:          params.Key,
		Data:         params.Data,
	}, nil
}

// GetHistoryItemOfWriteInt returns the data value that was set in the given transaction.
func (h *Historian) GetHistoryItemOfWriteInt(ctx context.Context, passport common.Address, txHash common.Hash) (*WriteIntHistoryItem, error) {
	from, txData, err := h.getTransactionSenderData(ctx, txHash)
	if err != nil {
		return nil, err
	}

	params, err := passportLogicInputParser.ParseSetIntCallData(txData)
	if err != nil {
		return nil, err
	}

	return &WriteIntHistoryItem{
		FactProvider: from,
		Key:          params.Key,
		Data:         params.Data,
	}, nil
}

// GetHistoryItemOfWriteBool returns the data value that was set in the given transaction.
func (h *Historian) GetHistoryItemOfWriteBool(ctx context.Context, passport common.Address, txHash common.Hash) (*WriteBoolHistoryItem, error) {
	from, txData, err := h.getTransactionSenderData(ctx, txHash)
	if err != nil {
		return nil, err
	}

	params, err := passportLogicInputParser.ParseSetBoolCallData(txData)
	if err != nil {
		return nil, err
	}

	return &WriteBoolHistoryItem{
		FactProvider: from,
		Key:          params.Key,
		Data:         params.Data,
	}, nil
}

func (h *Historian) getTransactionSenderData(ctx context.Context, txHash common.Hash) (common.Address, []byte, error) {
	(*eth.Eth)(h).Log("Getting transaction by hash", "tx_hash", txHash.Hex())
	tx, _, err := h.Backend.TransactionByHash(ctx, txHash)
	if err != nil {
		return common.Address{}, nil, fmt.Errorf("facts: TransactionByHash(%v): %v", txHash, err)
	}

	from, err := types.Sender(types.HomesteadSigner{}, tx)
	if err != nil {
		return common.Address{}, nil, fmt.Errorf("facts: types.Sender(): %v", err)
	}

	return from, tx.Data(), nil
}
