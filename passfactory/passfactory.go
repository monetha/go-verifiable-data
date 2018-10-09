package passfactory

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"gitlab.com/monetha/protocol-go-sdk/contracts"
	"gitlab.com/monetha/protocol-go-sdk/eth"
)

// Reader retrieves data from passport factory
type Reader eth.Eth

// NewReader converts session to Reader
func NewReader(e *eth.Eth) *Reader {
	return (*Reader)(e)
}

// Passport contains basic information about created passport
type Passport struct {
	// ContractAddress contains address of passport contract
	ContractAddress common.Address
	// FirstOwner contains address of first passport owner
	FirstOwner common.Address
	// Block number in which the passport was created
	BlockNumber uint64
	// Hash of the transaction in which the passport was created
	TxHash common.Hash
	// Index of the transaction in the block
	TxIndex uint
}

// PassportIterator is returned from FilterPassports and is used to iterate over the passports and unpacked data for
// PassportCreated events raised by the PassportFactoryContract contract.
type PassportIterator struct {
	it       *contracts.PassportFactoryContractPassportCreatedIterator
	Passport *Passport // Passport containing the info of the last retrieved passport
}

// Next advances the iterator to the subsequent passport, returning whether there
// are any more passports found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (pit *PassportIterator) Next() (next bool) {
	it := pit.it
	next = it.Next()

	ev := it.Event
	if !next || ev == nil || ev.Raw.Removed {
		pit.Passport = nil
		return
	}

	rawLog := ev.Raw
	pit.Passport = &Passport{
		ContractAddress: ev.Passport,
		FirstOwner:      ev.Owner,
		BlockNumber:     rawLog.BlockNumber,
		TxHash:          rawLog.TxHash,
		TxIndex:         rawLog.TxIndex,
	}
	return
}

// Error returns any retrieval or parsing error occurred during filtering.
func (pit *PassportIterator) Error() error {
	return pit.it.Error()
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (pit *PassportIterator) Close() error {
	pit.it.Close()
	return nil
}

// ToSlice retrieves all passports and saves them into slice.
func (pit *PassportIterator) ToSlice() ([]*Passport, error) {
	defer pit.Close()

	var ps []*Passport
	for pit.Next() {
		if err := pit.Error(); err != nil {
			return nil, err
		}
		ps = append(ps, pit.Passport)
	}

	return ps, nil
}

// PassportFilterOpts is the collection of options to fine tune filtering for passports.
type PassportFilterOpts struct {
	Start    uint64           // Start of the queried range
	End      *uint64          // End of the range (nil = latest)
	Passport []common.Address // Passport is a slice of passports to filter (nil = all passports)
	Owner    []common.Address // Owner is a slice of first owners to filter (nil = all owners)
	Context  context.Context  // Network context to support cancellation and timeouts (nil = no timeout)
}

// FilterPassports retrieves passports from event log of passport factory.
func (r *Reader) FilterPassports(opts *PassportFilterOpts, passportFactoryAddress common.Address) (*PassportIterator, error) {
	if opts == nil {
		opts = &PassportFilterOpts{}
	}

	backend := r.Backend

	(*eth.Eth)(r).Log("Initialising passport factory contract", "passport_factory", passportFactoryAddress.String())
	contract, err := contracts.NewPassportFactoryContract(passportFactoryAddress, backend)
	if err != nil {
		return nil, fmt.Errorf("passfactory: NewPassportFactoryContract: %v", err)
	}

	filterOpts := &bind.FilterOpts{
		Start:   opts.Start,
		End:     opts.End,
		Context: opts.Context,
	}
	it, err := contract.FilterPassportCreated(filterOpts, opts.Passport, opts.Owner)
	if err != nil {
		return nil, fmt.Errorf("passfactory: FilterPassportCreated: %v", err)
	}

	return &PassportIterator{it: it}, nil
}
