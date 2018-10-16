package facts

import (
	"context"
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
	// Block number in which the change has happened
	BlockNumber uint64
	// Hash of the transaction in which he change has happened
	TxHash common.Hash
	// Index of the transaction in the block
	TxIndex uint
}

type eventMetaInfo struct {
	EventName  string
	ChangeType change.Type
	DataType   data.Type
}

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
		/*
			case log := <-it.logs:
				var topics []string
				for _, value := range log.Topics {
					topics = append(topics, value.Hex())
				}
				fmt.Printf("Block number: %v Topics: %v Data: %v\n", log.BlockNumber, topics, common.Bytes2Hex(log.Data))
				return true
		*/

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
	/*
		case log := <-it.logs:
			var topics []string
			for _, value := range log.Topics {
				topics = append(topics, value.Hex())
			}
			fmt.Printf("Block number: %v Topics: %v Data: %v\n", log.BlockNumber, topics, common.Bytes2Hex(log.Data))
			return true
	*/
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
		ChangeType:  emi.ChangeType,
		DataType:    emi.DataType,
		BlockNumber: log.BlockNumber,
		TxHash:      log.TxHash,
		TxIndex:     log.TxIndex,
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

					eventId := passportLogicContractABI.Events[eventName].Id()
					eventMetaInfos[eventId] = eventMetaInfo{
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
