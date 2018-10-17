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

	// make some updates
	if err := provider.WriteTxData(ctx, passportAddress, key, []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}); err != nil {
		t.Errorf("Provider.WriteTxData() error = %v", err)
	}
	if err := provider.WriteBytes(ctx, passportAddress, key, []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}); err != nil {
		t.Errorf("Provider.WriteBytes() error = %v", err)
	}
	if err := provider.WriteString(ctx, passportAddress, key, "test only string"); err != nil {
		t.Errorf("Provider.WriteString() error = %v", err)
	}
	if err := provider.WriteAddress(ctx, passportAddress, key, common.HexToAddress("0xaF4DcE16Da2877f8c9e00544c93B62Ac40631F16")); err != nil {
		t.Errorf("Provider.WriteAddress() error = %v", err)
	}
	if err := provider.WriteInt(ctx, passportAddress, key, big.NewInt(-123456)); err != nil {
		t.Errorf("Provider.WriteInt() error = %v", err)
	}
	if err := provider.WriteUint(ctx, passportAddress, key, big.NewInt(123456)); err != nil {
		t.Errorf("Provider.WriteUint() error = %v", err)
	}
	if err := provider.WriteBool(ctx, passportAddress, key, true); err != nil {
		t.Errorf("Provider.WriteBool() error = %v", err)
	}

	// make some deletes
	if err := provider.DeleteTxData(ctx, passportAddress, key); err != nil {
		t.Errorf("Provider.DeleteTxData() error = %v", err)
	}
	if err := provider.DeleteBytes(ctx, passportAddress, key); err != nil {
		t.Errorf("Provider.DeleteBytes() error = %v", err)
	}
	if err := provider.DeleteString(ctx, passportAddress, key); err != nil {
		t.Errorf("Provider.DeleteString() error = %v", err)
	}
	if err := provider.DeleteAddress(ctx, passportAddress, key); err != nil {
		t.Errorf("Provider.DeleteAddress() error = %v", err)
	}
	if err := provider.DeleteInt(ctx, passportAddress, key); err != nil {
		t.Errorf("Provider.DeleteInt() error = %v", err)
	}
	if err := provider.DeleteUint(ctx, passportAddress, key); err != nil {
		t.Errorf("Provider.DeleteUint() error = %v", err)
	}
	if err := provider.DeleteBool(ctx, passportAddress, key); err != nil {
		t.Errorf("Provider.DeleteBool() error = %v", err)
	}

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
		t.Errorf("ChnageIterator.ToSlice() error = %v", err)
	}

	if len(changes) != 1 {
		t.Errorf("Expected to get exactly one change, got %v", len(changes))
	}

	ch := changes[0]
	type x struct {
		ChangeType   change.Type
		DataType     data.Type
		FactProvider common.Address
		Key          [32]byte
	}

	expected := x{
		ChangeType:   changeType,
		DataType:     dataType,
		FactProvider: factProviderAddress,
		Key:          key,
	}

	actual := x{
		ChangeType:   ch.ChangeType,
		DataType:     ch.DataType,
		FactProvider: ch.FactProvider,
		Key:          ch.Key,
	}

	if diff := deep.Equal(expected, actual); diff != nil {
		t.Error(diff)
	}
}
