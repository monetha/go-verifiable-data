package contracts

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	methereum "github.com/monetha/go-ethereum"
)

// PassportFactoryLogFilterer filters PassportFactory event logs
type PassportFactoryLogFilterer struct {
	abi abi.ABI
}

// NewPassportFactoryLogFilterer creates an instance of PassportFactoryLogFilterer
func NewPassportFactoryLogFilterer() (*PassportFactoryLogFilterer, error) {
	parsed, err := abi.JSON(strings.NewReader(PassportFactoryContractABI))
	if err != nil {
		return nil, err
	}
	return &PassportFactoryLogFilterer{parsed}, nil
}

// FilterPassportCreated parses event logs and returns PassportCreated events if any found.
func (f *PassportFactoryLogFilterer) FilterPassportCreated(logs []*types.Log, passport []common.Address, owner []common.Address) (events []PassportFactoryContractPassportCreated, err error) {
	cf := &PassportFactoryContractFilterer{
		contract: bind.NewBoundContract(common.Address{}, f.abi, nil, nil, methereum.SliceLogFilterer(logs)),
	}

	var it *PassportFactoryContractPassportCreatedIterator
	it, err = cf.FilterPassportCreated(nil, passport, owner)
	if err != nil {
		return
	}
	defer func() {
		if cErr := it.Close(); err == nil && cErr != nil {
			err = cErr
		}
	}()

	for it.Next() {
		if err = it.Error(); err != nil {
			return nil, err
		}

		ev := it.Event
		if ev == nil {
			continue
		}

		events = append(events, *ev)
	}

	return
}
