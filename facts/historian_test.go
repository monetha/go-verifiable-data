package facts_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/go-test/deep"
	"gitlab.com/monetha/protocol-go-sdk/facts"
	"gitlab.com/monetha/protocol-go-sdk/types/change"
	"gitlab.com/monetha/protocol-go-sdk/types/data"
)

func TestHistorian_FilterChanges(t *testing.T) {

	ctx := context.Background()

	passportAddress, factProviderSession := createPassportAndFactProviderSession(ctx, t)
	factProviderAddress := factProviderSession.TransactOpts.From

	provider := facts.NewProvider(factProviderSession)

	key := [32]byte{99, 88, 77, 66, 55, 44, 33, 22, 11}

	var allChanges []changeDetails
	// make some updates
	if err := provider.WriteTxData(ctx, passportAddress, key, []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}); err != nil {
		t.Errorf("Provider.WriteTxData() error = %v", err)
	}
	allChanges = append(allChanges, changeDetails{ChangeType: change.Updated, DataType: data.TxData, FactProvider: factProviderAddress, Key: key})
	if err := provider.WriteBytes(ctx, passportAddress, key, []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}); err != nil {
		t.Errorf("Provider.WriteBytes() error = %v", err)
	}
	allChanges = append(allChanges, changeDetails{ChangeType: change.Updated, DataType: data.Bytes, FactProvider: factProviderAddress, Key: key})
	if err := provider.WriteString(ctx, passportAddress, key, "test only string"); err != nil {
		t.Errorf("Provider.WriteString() error = %v", err)
	}
	allChanges = append(allChanges, changeDetails{ChangeType: change.Updated, DataType: data.String, FactProvider: factProviderAddress, Key: key})
	if err := provider.WriteAddress(ctx, passportAddress, key, common.HexToAddress("0xaF4DcE16Da2877f8c9e00544c93B62Ac40631F16")); err != nil {
		t.Errorf("Provider.WriteAddress() error = %v", err)
	}
	allChanges = append(allChanges, changeDetails{ChangeType: change.Updated, DataType: data.Address, FactProvider: factProviderAddress, Key: key})
	if err := provider.WriteInt(ctx, passportAddress, key, big.NewInt(-123456)); err != nil {
		t.Errorf("Provider.WriteInt() error = %v", err)
	}
	allChanges = append(allChanges, changeDetails{ChangeType: change.Updated, DataType: data.Int, FactProvider: factProviderAddress, Key: key})
	if err := provider.WriteUint(ctx, passportAddress, key, big.NewInt(123456)); err != nil {
		t.Errorf("Provider.WriteUint() error = %v", err)
	}
	allChanges = append(allChanges, changeDetails{ChangeType: change.Updated, DataType: data.Uint, FactProvider: factProviderAddress, Key: key})
	if err := provider.WriteBool(ctx, passportAddress, key, true); err != nil {
		t.Errorf("Provider.WriteBool() error = %v", err)
	}
	allChanges = append(allChanges, changeDetails{ChangeType: change.Updated, DataType: data.Bool, FactProvider: factProviderAddress, Key: key})

	// make some deletes
	if err := provider.DeleteTxData(ctx, passportAddress, key); err != nil {
		t.Errorf("Provider.DeleteTxData() error = %v", err)
	}
	allChanges = append(allChanges, changeDetails{ChangeType: change.Deleted, DataType: data.TxData, FactProvider: factProviderAddress, Key: key})
	if err := provider.DeleteBytes(ctx, passportAddress, key); err != nil {
		t.Errorf("Provider.DeleteBytes() error = %v", err)
	}
	allChanges = append(allChanges, changeDetails{ChangeType: change.Deleted, DataType: data.Bytes, FactProvider: factProviderAddress, Key: key})
	if err := provider.DeleteString(ctx, passportAddress, key); err != nil {
		t.Errorf("Provider.DeleteString() error = %v", err)
	}
	allChanges = append(allChanges, changeDetails{ChangeType: change.Deleted, DataType: data.String, FactProvider: factProviderAddress, Key: key})
	if err := provider.DeleteAddress(ctx, passportAddress, key); err != nil {
		t.Errorf("Provider.DeleteAddress() error = %v", err)
	}
	allChanges = append(allChanges, changeDetails{ChangeType: change.Deleted, DataType: data.Address, FactProvider: factProviderAddress, Key: key})
	if err := provider.DeleteInt(ctx, passportAddress, key); err != nil {
		t.Errorf("Provider.DeleteInt() error = %v", err)
	}
	allChanges = append(allChanges, changeDetails{ChangeType: change.Deleted, DataType: data.Int, FactProvider: factProviderAddress, Key: key})
	if err := provider.DeleteUint(ctx, passportAddress, key); err != nil {
		t.Errorf("Provider.DeleteUint() error = %v", err)
	}
	allChanges = append(allChanges, changeDetails{ChangeType: change.Deleted, DataType: data.Uint, FactProvider: factProviderAddress, Key: key})
	if err := provider.DeleteBool(ctx, passportAddress, key); err != nil {
		t.Errorf("Provider.DeleteBool() error = %v", err)
	}
	allChanges = append(allChanges, changeDetails{ChangeType: change.Deleted, DataType: data.Bool, FactProvider: factProviderAddress, Key: key})

	eth := factProviderSession.Eth
	hist := facts.NewHistorian(eth)

	tests := []struct {
		name       string
		changeType change.Type
		dataType   data.Type
	}{
		{"filter txdata update", change.Updated, data.TxData},
		{"filter bytes update", change.Updated, data.Bytes},
		{"filter string update", change.Updated, data.String},
		{"filter address update", change.Updated, data.Address},
		{"filter int update", change.Updated, data.Int},
		{"filter bool update", change.Updated, data.Bool},
		{"filter txdata delete", change.Updated, data.TxData},
		{"filter bytes delete", change.Deleted, data.Bytes},
		{"filter string delete", change.Deleted, data.String},
		{"filter address delete", change.Deleted, data.Address},
		{"filter int delete", change.Deleted, data.Int},
		{"filter bool delete", change.Deleted, data.Bool},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testFilterChangesExactOne(t, hist, tt.changeType, tt.dataType, passportAddress, factProviderAddress, key)
		})
	}

	t.Run("filter all order", func(t *testing.T) {
		testFilterChangesAll(t, hist, passportAddress, allChanges)
	})

	t.Run("filter all updates order", func(t *testing.T) {
		testFilterChangesAllUpdates(t, hist, passportAddress, allChanges)
	})

	t.Run("filter all deletes order", func(t *testing.T) {
		testFilterChangesAllDeletes(t, hist, passportAddress, allChanges)
	})
}

type changeDetails struct {
	ChangeType   change.Type
	DataType     data.Type
	FactProvider common.Address
	Key          [32]byte
}

func testFilterChangesExactOne(t *testing.T, hist *facts.Historian, changeType change.Type, dataType data.Type, passportAddress common.Address, factProviderAddress common.Address, key [32]byte) {
	opts := &facts.ChangesFilterOpts{
		ChangeType:   []change.Type{changeType},
		DataType:     []data.Type{dataType},
		FactProvider: []common.Address{factProviderAddress},
		Key:          [][32]byte{key},
	}

	it, err := hist.FilterChanges(opts, passportAddress)
	if err != nil {
		t.Errorf("Historian.FilterChanges() error = %v", err)
	}

	changes, err := it.ToSlice()
	if err != nil {
		t.Errorf("ChangeIterator.ToSlice() error = %v", err)
	}

	if len(changes) != 1 {
		t.Errorf("Expected to get exactly one change, got %v", len(changes))
	}

	ch := changes[0]

	expected := changeDetails{
		ChangeType:   changeType,
		DataType:     dataType,
		FactProvider: factProviderAddress,
		Key:          key,
	}

	actual := changeDetails{
		ChangeType:   ch.ChangeType,
		DataType:     ch.DataType,
		FactProvider: ch.FactProvider,
		Key:          ch.Key,
	}

	if diff := deep.Equal(expected, actual); diff != nil {
		t.Error(diff)
	}
}

func testFilterChangesAll(t *testing.T, hist *facts.Historian, passportAddress common.Address, allChanges []changeDetails) {
	it, err := hist.FilterChanges(nil, passportAddress)
	if err != nil {
		t.Errorf("Historian.FilterChanges() error = %v", err)
	}

	changes, err := it.ToSlice()
	if err != nil {
		t.Errorf("ChangeIterator.ToSlice() error = %v", err)
	}

	var actualChanges []changeDetails
	for _, ch := range changes {
		actualChanges = append(actualChanges, changeDetails{
			ChangeType:   ch.ChangeType,
			DataType:     ch.DataType,
			FactProvider: ch.FactProvider,
			Key:          ch.Key,
		})
	}

	if diff := deep.Equal(allChanges, actualChanges); diff != nil {
		t.Error(diff)
	}
}

func testFilterChangesAllUpdates(t *testing.T, hist *facts.Historian, passportAddress common.Address, allChanges []changeDetails) {
	opts := &facts.ChangesFilterOpts{ChangeType: []change.Type{change.Updated}}
	it, err := hist.FilterChanges(opts, passportAddress)
	if err != nil {
		t.Errorf("Historian.FilterChanges() error = %v", err)
	}

	changes, err := it.ToSlice()
	if err != nil {
		t.Errorf("ChangeIterator.ToSlice() error = %v", err)
	}

	var actualChanges []changeDetails
	for _, ch := range changes {
		actualChanges = append(actualChanges, changeDetails{
			ChangeType:   ch.ChangeType,
			DataType:     ch.DataType,
			FactProvider: ch.FactProvider,
			Key:          ch.Key,
		})
	}

	var expectedChanges []changeDetails
	for _, ch := range allChanges {
		if ch.ChangeType != change.Updated {
			continue
		}
		expectedChanges = append(expectedChanges, ch)
	}

	if diff := deep.Equal(expectedChanges, actualChanges); diff != nil {
		t.Error(diff)
	}
}

func testFilterChangesAllDeletes(t *testing.T, hist *facts.Historian, passportAddress common.Address, allChanges []changeDetails) {
	opts := &facts.ChangesFilterOpts{ChangeType: []change.Type{change.Deleted}}
	it, err := hist.FilterChanges(opts, passportAddress)
	if err != nil {
		t.Errorf("Historian.FilterChanges() error = %v", err)
	}

	changes, err := it.ToSlice()
	if err != nil {
		t.Errorf("ChangeIterator.ToSlice() error = %v", err)
	}

	var actualChanges []changeDetails
	for _, ch := range changes {
		actualChanges = append(actualChanges, changeDetails{
			ChangeType:   ch.ChangeType,
			DataType:     ch.DataType,
			FactProvider: ch.FactProvider,
			Key:          ch.Key,
		})
	}

	var expectedChanges []changeDetails
	for _, ch := range allChanges {
		if ch.ChangeType != change.Deleted {
			continue
		}
		expectedChanges = append(expectedChanges, ch)
	}

	if diff := deep.Equal(expectedChanges, actualChanges); diff != nil {
		t.Error(diff)
	}
}
