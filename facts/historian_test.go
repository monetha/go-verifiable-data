package facts_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/go-test/deep"
	"github.com/monetha/go-verifiable-data/facts"
	"github.com/monetha/go-verifiable-data/types/change"
	"github.com/monetha/go-verifiable-data/types/data"
)

func TestHistorian_FilterChanges(t *testing.T) {

	ctx := context.Background()

	passportAddress, factProviderSession := createPassportAndFactProviderSession(ctx, t)
	factProviderAddress := factProviderSession.TransactOpts.From

	provider := &providerChanges{t: t, p: facts.NewProvider(factProviderSession)}

	key := [32]byte{99, 88, 77, 66, 55, 44, 33, 22, 11}

	// make some updates
	provider.WriteTxData(ctx, passportAddress, key, []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0})
	provider.WriteBytes(ctx, passportAddress, key, []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0})
	provider.WriteString(ctx, passportAddress, key, "test only string")
	provider.WriteAddress(ctx, passportAddress, key, common.HexToAddress("0xaF4DcE16Da2877f8c9e00544c93B62Ac40631F16"))
	provider.WriteInt(ctx, passportAddress, key, big.NewInt(-123456))
	provider.WriteUint(ctx, passportAddress, key, big.NewInt(123456))
	provider.WriteBool(ctx, passportAddress, key, true)
	provider.WriteIPFSHash(ctx, passportAddress, key, "QmTp2hEo8eXRp6wg7jXv1BLCMh5a4F3B7buAUZNZUu772j")
	provider.WritePrivateDataHashes(ctx, passportAddress, key, &facts.PrivateDataHashes{DataIPFSHash: "QmTp2hEo8eXRp6wg7jXv1BLCMh5a4F3B7buAUZNZUu772j", DataKeyHash: [32]byte{1, 2, 3}})

	// make some deletes
	provider.DeleteTxData(ctx, passportAddress, key)
	provider.DeleteBytes(ctx, passportAddress, key)
	provider.DeleteString(ctx, passportAddress, key)
	provider.DeleteAddress(ctx, passportAddress, key)
	provider.DeleteInt(ctx, passportAddress, key)
	provider.DeleteUint(ctx, passportAddress, key)
	provider.DeleteBool(ctx, passportAddress, key)
	provider.DeleteIPFSHash(ctx, passportAddress, key)
	provider.DeletePrivateDataHashes(ctx, passportAddress, key)

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
		{"filter ipfs hash update", change.Updated, data.IPFS},
		{"filter txdata delete", change.Updated, data.TxData},
		{"filter bytes delete", change.Deleted, data.Bytes},
		{"filter string delete", change.Deleted, data.String},
		{"filter address delete", change.Deleted, data.Address},
		{"filter int delete", change.Deleted, data.Int},
		{"filter bool delete", change.Deleted, data.Bool},
		{"filter ipfs hash delete", change.Deleted, data.IPFS},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testFilterChangesExactOne(t, hist, tt.changeType, tt.dataType, passportAddress, factProviderAddress, key, provider.AllChanges)
		})
	}

	t.Run("filter all order", func(t *testing.T) {
		testFilterChangesAll(t, hist, passportAddress, provider.AllChanges)
	})

	t.Run("filter all updates order", func(t *testing.T) {
		testFilterChangesAllUpdates(t, hist, passportAddress, provider.AllChanges)
	})

	t.Run("filter all deletes order", func(t *testing.T) {
		testFilterChangesAllDeletes(t, hist, passportAddress, provider.AllChanges)
	})
}

type changeDetails struct {
	ChangeType   change.Type
	DataType     data.Type
	FactProvider common.Address
	Key          [32]byte
	TxHash       common.Hash
}

type providerChanges struct {
	t          *testing.T
	p          *facts.Provider
	AllChanges []changeDetails
}

func (c *providerChanges) WriteTxData(ctx context.Context, passportAddress common.Address, key [32]byte, value []byte) {
	txHash, err := c.p.WriteTxData(ctx, passportAddress, key, value)
	if err != nil {
		c.t.Fatalf("Provider.WriteTxData() error = %v", err)
	}
	c.AllChanges = append(c.AllChanges, changeDetails{ChangeType: change.Updated, DataType: data.TxData, FactProvider: c.p.TransactOpts.From, Key: key, TxHash: txHash})
}

func (c *providerChanges) WriteBytes(ctx context.Context, passportAddress common.Address, key [32]byte, value []byte) {
	txHash, err := c.p.WriteBytes(ctx, passportAddress, key, value)
	if err != nil {
		c.t.Fatalf("Provider.WriteBytes() error = %v", err)
	}
	c.AllChanges = append(c.AllChanges, changeDetails{ChangeType: change.Updated, DataType: data.Bytes, FactProvider: c.p.TransactOpts.From, Key: key, TxHash: txHash})
}

func (c *providerChanges) WriteString(ctx context.Context, passportAddress common.Address, key [32]byte, value string) {
	txHash, err := c.p.WriteString(ctx, passportAddress, key, value)
	if err != nil {
		c.t.Fatalf("Provider.WriteString() error = %v", err)
	}
	c.AllChanges = append(c.AllChanges, changeDetails{ChangeType: change.Updated, DataType: data.String, FactProvider: c.p.TransactOpts.From, Key: key, TxHash: txHash})
}

func (c *providerChanges) WriteAddress(ctx context.Context, passportAddress common.Address, key [32]byte, value common.Address) {
	txHash, err := c.p.WriteAddress(ctx, passportAddress, key, value)
	if err != nil {
		c.t.Fatalf("Provider.WriteAddress() error = %v", err)
	}
	c.AllChanges = append(c.AllChanges, changeDetails{ChangeType: change.Updated, DataType: data.Address, FactProvider: c.p.TransactOpts.From, Key: key, TxHash: txHash})
}

func (c *providerChanges) WriteUint(ctx context.Context, passportAddress common.Address, key [32]byte, value *big.Int) {
	txHash, err := c.p.WriteUint(ctx, passportAddress, key, value)
	if err != nil {
		c.t.Fatalf("Provider.WriteUint() error = %v", err)
	}
	c.AllChanges = append(c.AllChanges, changeDetails{ChangeType: change.Updated, DataType: data.Uint, FactProvider: c.p.TransactOpts.From, Key: key, TxHash: txHash})
}

func (c *providerChanges) WriteInt(ctx context.Context, passportAddress common.Address, key [32]byte, value *big.Int) {
	txHash, err := c.p.WriteInt(ctx, passportAddress, key, value)
	if err != nil {
		c.t.Fatalf("Provider.WriteInt() error = %v", err)
	}
	c.AllChanges = append(c.AllChanges, changeDetails{ChangeType: change.Updated, DataType: data.Int, FactProvider: c.p.TransactOpts.From, Key: key, TxHash: txHash})
}

func (c *providerChanges) WriteBool(ctx context.Context, passportAddress common.Address, key [32]byte, value bool) {
	txHash, err := c.p.WriteBool(ctx, passportAddress, key, value)
	if err != nil {
		c.t.Fatalf("Provider.WriteBool() error = %v", err)
	}
	c.AllChanges = append(c.AllChanges, changeDetails{ChangeType: change.Updated, DataType: data.Bool, FactProvider: c.p.TransactOpts.From, Key: key, TxHash: txHash})
}

func (c *providerChanges) WriteIPFSHash(ctx context.Context, passportAddress common.Address, key [32]byte, value string) {
	txHash, err := c.p.WriteIPFSHash(ctx, passportAddress, key, value)
	if err != nil {
		c.t.Fatalf("Provider.WriteIPFSHash() error = %v", err)
	}
	c.AllChanges = append(c.AllChanges, changeDetails{ChangeType: change.Updated, DataType: data.IPFS, FactProvider: c.p.TransactOpts.From, Key: key, TxHash: txHash})
}

func (c *providerChanges) WritePrivateDataHashes(ctx context.Context, passportAddress common.Address, key [32]byte, privateData *facts.PrivateDataHashes) {
	txHash, err := c.p.WritePrivateDataHashes(ctx, passportAddress, key, privateData)
	if err != nil {
		c.t.Fatalf("Provider.WritePrivateDataHashes() error = %v", err)
	}
	c.AllChanges = append(c.AllChanges, changeDetails{ChangeType: change.Updated, DataType: data.PrivateData, FactProvider: c.p.TransactOpts.From, Key: key, TxHash: txHash})
}

func (c *providerChanges) DeleteTxData(ctx context.Context, passportAddress common.Address, key [32]byte) {
	txHash, err := c.p.DeleteTxData(ctx, passportAddress, key)
	if err != nil {
		c.t.Fatalf("Provider.DeleteTxData() error = %v", err)
	}
	c.AllChanges = append(c.AllChanges, changeDetails{ChangeType: change.Deleted, DataType: data.TxData, FactProvider: c.p.TransactOpts.From, Key: key, TxHash: txHash})
}

func (c *providerChanges) DeleteBytes(ctx context.Context, passportAddress common.Address, key [32]byte) {
	txHash, err := c.p.DeleteBytes(ctx, passportAddress, key)
	if err != nil {
		c.t.Fatalf("Provider.DeleteBytes() error = %v", err)
	}
	c.AllChanges = append(c.AllChanges, changeDetails{ChangeType: change.Deleted, DataType: data.Bytes, FactProvider: c.p.TransactOpts.From, Key: key, TxHash: txHash})
}

func (c *providerChanges) DeleteString(ctx context.Context, passportAddress common.Address, key [32]byte) {
	txHash, err := c.p.DeleteString(ctx, passportAddress, key)
	if err != nil {
		c.t.Fatalf("Provider.DeleteString() error = %v", err)
	}
	c.AllChanges = append(c.AllChanges, changeDetails{ChangeType: change.Deleted, DataType: data.String, FactProvider: c.p.TransactOpts.From, Key: key, TxHash: txHash})
}

func (c *providerChanges) DeleteAddress(ctx context.Context, passportAddress common.Address, key [32]byte) {
	txHash, err := c.p.DeleteAddress(ctx, passportAddress, key)
	if err != nil {
		c.t.Fatalf("Provider.DeleteAddress() error = %v", err)
	}
	c.AllChanges = append(c.AllChanges, changeDetails{ChangeType: change.Deleted, DataType: data.Address, FactProvider: c.p.TransactOpts.From, Key: key, TxHash: txHash})
}

func (c *providerChanges) DeleteUint(ctx context.Context, passportAddress common.Address, key [32]byte) {
	txHash, err := c.p.DeleteUint(ctx, passportAddress, key)
	if err != nil {
		c.t.Fatalf("Provider.DeleteUint() error = %v", err)
	}
	c.AllChanges = append(c.AllChanges, changeDetails{ChangeType: change.Deleted, DataType: data.Uint, FactProvider: c.p.TransactOpts.From, Key: key, TxHash: txHash})
}

func (c *providerChanges) DeleteInt(ctx context.Context, passportAddress common.Address, key [32]byte) {
	txHash, err := c.p.DeleteInt(ctx, passportAddress, key)
	if err != nil {
		c.t.Fatalf("Provider.DeleteInt() error = %v", err)
	}
	c.AllChanges = append(c.AllChanges, changeDetails{ChangeType: change.Deleted, DataType: data.Int, FactProvider: c.p.TransactOpts.From, Key: key, TxHash: txHash})
}

func (c *providerChanges) DeleteBool(ctx context.Context, passportAddress common.Address, key [32]byte) {
	txHash, err := c.p.DeleteBool(ctx, passportAddress, key)
	if err != nil {
		c.t.Fatalf("Provider.DeleteBool() error = %v", err)
	}
	c.AllChanges = append(c.AllChanges, changeDetails{ChangeType: change.Deleted, DataType: data.Bool, FactProvider: c.p.TransactOpts.From, Key: key, TxHash: txHash})
}

func (c *providerChanges) DeleteIPFSHash(ctx context.Context, passportAddress common.Address, key [32]byte) {
	txHash, err := c.p.DeleteIPFSHash(ctx, passportAddress, key)
	if err != nil {
		c.t.Fatalf("Provider.DeleteIPFSHash() error = %v", err)
	}
	c.AllChanges = append(c.AllChanges, changeDetails{ChangeType: change.Deleted, DataType: data.IPFS, FactProvider: c.p.TransactOpts.From, Key: key, TxHash: txHash})
}

func (c *providerChanges) DeletePrivateDataHashes(ctx context.Context, passportAddress common.Address, key [32]byte) {
	txHash, err := c.p.DeletePrivateDataHashes(ctx, passportAddress, key)
	if err != nil {
		c.t.Fatalf("Provider.DeletePrivateDataHashes() error = %v", err)
	}
	c.AllChanges = append(c.AllChanges, changeDetails{ChangeType: change.Deleted, DataType: data.PrivateData, FactProvider: c.p.TransactOpts.From, Key: key, TxHash: txHash})
}

func testFilterChangesExactOne(t *testing.T, hist *facts.Historian, changeType change.Type, dataType data.Type, passportAddress common.Address, factProviderAddress common.Address, key [32]byte, allChanges []changeDetails) {
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

	var expected changeDetails
	for _, v := range allChanges {
		if v.ChangeType == changeType &&
			v.DataType == dataType &&
			v.FactProvider == factProviderAddress &&
			v.Key == key {
			expected = v
			break
		}
	}

	actual := changeToDetails(ch)

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
		actualChanges = append(actualChanges, changeToDetails(ch))
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
		actualChanges = append(actualChanges, changeToDetails(ch))
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
		actualChanges = append(actualChanges, changeToDetails(ch))
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

func changeToDetails(ch *facts.Change) changeDetails {
	return changeDetails{
		ChangeType:   ch.ChangeType,
		DataType:     ch.DataType,
		FactProvider: ch.FactProvider,
		Key:          ch.Key,
		TxHash:       ch.Raw.TxHash,
	}
}

func TestHistorian_GetHistoryItemOfWriteTxData(t *testing.T) {
	ctx := context.Background()

	passportAddress, factProviderSession := createPassportAndFactProviderSession(ctx, t)
	ehi := &facts.WriteTxDataHistoryItem{
		FactProvider: factProviderSession.TransactOpts.From,
		Key:          [32]byte{1, 2, 3, 4, 5, 6, 7, 8, 9},
		Data:         ([]byte)("this is test only message"),
	}

	provider := facts.NewProvider(factProviderSession)

	txHash, err := provider.WriteTxData(ctx, passportAddress, ehi.Key, ehi.Data)
	if err != nil {
		t.Errorf("Provider.WriteTxData() error = %v", err)
	}

	hi, err := facts.NewHistorian(factProviderSession.Eth).GetHistoryItemOfWriteTxData(ctx, passportAddress, txHash)
	if err != nil {
		t.Errorf("Historian.GetHistoryItemOfWriteTxData() error = %v", err)
	}

	if diff := deep.Equal(ehi, hi); diff != nil {
		t.Error(diff)
	}
}

func TestHistorian_GetHistoryItemOfWriteBytes(t *testing.T) {
	ctx := context.Background()

	passportAddress, factProviderSession := createPassportAndFactProviderSession(ctx, t)
	ehi := &facts.WriteBytesHistoryItem{
		FactProvider: factProviderSession.TransactOpts.From,
		Key:          [32]byte{1, 2, 3, 4, 5, 6, 7, 8, 9},
		Data:         ([]byte)("this is test only message"),
	}

	provider := facts.NewProvider(factProviderSession)

	txHash, err := provider.WriteBytes(ctx, passportAddress, ehi.Key, ehi.Data)
	if err != nil {
		t.Errorf("Provider.WriteBytes() error = %v", err)
	}

	hi, err := facts.NewHistorian(factProviderSession.Eth).GetHistoryItemOfWriteBytes(ctx, passportAddress, txHash)
	if err != nil {
		t.Errorf("Historian.GetHistoryItemOfWriteBytes() error = %v", err)
	}

	if diff := deep.Equal(ehi, hi); diff != nil {
		t.Error(diff)
	}
}

func TestHistorian_GetHistoryItemOfWriteString(t *testing.T) {
	ctx := context.Background()

	passportAddress, factProviderSession := createPassportAndFactProviderSession(ctx, t)
	ehi := &facts.WriteStringHistoryItem{
		FactProvider: factProviderSession.TransactOpts.From,
		Key:          [32]byte{1, 2, 3, 4, 5, 6, 7, 8, 9},
		Data:         "this is test only message",
	}

	provider := facts.NewProvider(factProviderSession)

	txHash, err := provider.WriteString(ctx, passportAddress, ehi.Key, ehi.Data)
	if err != nil {
		t.Errorf("Provider.WriteString() error = %v", err)
	}

	hi, err := facts.NewHistorian(factProviderSession.Eth).GetHistoryItemOfWriteString(ctx, passportAddress, txHash)
	if err != nil {
		t.Errorf("Historian.GetHistoryItemOfWriteString error = %v", err)
	}

	if diff := deep.Equal(ehi, hi); diff != nil {
		t.Error(diff)
	}
}

func TestHistorian_GetHistoryItemOfWriteAddress(t *testing.T) {
	ctx := context.Background()

	passportAddress, factProviderSession := createPassportAndFactProviderSession(ctx, t)
	ehi := &facts.WriteAddressHistoryItem{
		FactProvider: factProviderSession.TransactOpts.From,
		Key:          [32]byte{1, 2, 3, 4, 5, 6, 7, 8, 9},
		Data:         common.HexToAddress("0xaF4DcE16Da2877f8c9e00544c93B62Ac40631F16"),
	}

	provider := facts.NewProvider(factProviderSession)

	txHash, err := provider.WriteAddress(ctx, passportAddress, ehi.Key, ehi.Data)
	if err != nil {
		t.Errorf("Provider.WriteAddress() error = %v", err)
	}

	hi, err := facts.NewHistorian(factProviderSession.Eth).GetHistoryItemOfWriteAddress(ctx, passportAddress, txHash)
	if err != nil {
		t.Errorf("Historian.GetHistoryItemOfWriteAddress() error = %v", err)
	}

	if diff := deep.Equal(ehi, hi); diff != nil {
		t.Error(diff)
	}
}

func TestHistorian_GetHistoryItemOfWriteUint(t *testing.T) {
	ctx := context.Background()

	passportAddress, factProviderSession := createPassportAndFactProviderSession(ctx, t)
	ehi := &facts.WriteUintHistoryItem{
		FactProvider: factProviderSession.TransactOpts.From,
		Key:          [32]byte{1, 2, 3, 4, 5, 6, 7, 8, 9},
		Data:         big.NewInt(1234567890),
	}

	provider := facts.NewProvider(factProviderSession)

	txHash, err := provider.WriteUint(ctx, passportAddress, ehi.Key, ehi.Data)
	if err != nil {
		t.Errorf("Provider.WriteUint() error = %v", err)
	}

	hi, err := facts.NewHistorian(factProviderSession.Eth).GetHistoryItemOfWriteUint(ctx, passportAddress, txHash)
	if err != nil {
		t.Errorf("Historian.GetHistoryItemOfWriteUint() error = %v", err)
	}

	if diff := deep.Equal(ehi, hi); diff != nil {
		t.Error(diff)
	}
}

func TestHistorian_GetHistoryItemOfWriteInt(t *testing.T) {
	ctx := context.Background()

	passportAddress, factProviderSession := createPassportAndFactProviderSession(ctx, t)
	ehi := &facts.WriteIntHistoryItem{
		FactProvider: factProviderSession.TransactOpts.From,
		Key:          [32]byte{1, 2, 3, 4, 5, 6, 7, 8, 9},
		Data:         big.NewInt(-1234567890),
	}

	provider := facts.NewProvider(factProviderSession)

	txHash, err := provider.WriteInt(ctx, passportAddress, ehi.Key, ehi.Data)
	if err != nil {
		t.Errorf("Provider.WriteInt() error = %v", err)
	}

	hi, err := facts.NewHistorian(factProviderSession.Eth).GetHistoryItemOfWriteInt(ctx, passportAddress, txHash)
	if err != nil {
		t.Errorf("Historian.GetHistoryItemOfWriteInt() error = %v", err)
	}

	if diff := deep.Equal(ehi, hi); diff != nil {
		t.Error(diff)
	}
}

func TestHistorian_GetHistoryItemOfWriteBool(t *testing.T) {
	ctx := context.Background()

	passportAddress, factProviderSession := createPassportAndFactProviderSession(ctx, t)
	ehi := &facts.WriteBoolHistoryItem{
		FactProvider: factProviderSession.TransactOpts.From,
		Key:          [32]byte{1, 2, 3, 4, 5, 6, 7, 8, 9},
		Data:         true,
	}

	provider := facts.NewProvider(factProviderSession)

	txHash, err := provider.WriteBool(ctx, passportAddress, ehi.Key, ehi.Data)
	if err != nil {
		t.Errorf("Provider.WriteBool() error = %v", err)
	}

	hi, err := facts.NewHistorian(factProviderSession.Eth).GetHistoryItemOfWriteBool(ctx, passportAddress, txHash)
	if err != nil {
		t.Errorf("Historian.GetHistoryItemOfWriteBool() error = %v", err)
	}

	if diff := deep.Equal(ehi, hi); diff != nil {
		t.Error(diff)
	}
}

func TestHistorian_GetHistoryItemOfWriteIPFSHash(t *testing.T) {
	ctx := context.Background()

	passportAddress, factProviderSession := createPassportAndFactProviderSession(ctx, t)
	ehi := &facts.WriteIPFSHashHistoryItem{
		FactProvider: factProviderSession.TransactOpts.From,
		Key:          [32]byte{1, 2, 3, 4, 5, 6, 7, 8, 9},
		Hash:         "QmTp2hEo8eXRp6wg7jXv1BLCMh5a4F3B7buAUZNZUu772j",
	}

	provider := facts.NewProvider(factProviderSession)

	txHash, err := provider.WriteIPFSHash(ctx, passportAddress, ehi.Key, ehi.Hash)
	if err != nil {
		t.Errorf("Provider.WriteIPFSHash() error = %v", err)
	}

	hi, err := facts.NewHistorian(factProviderSession.Eth).GetHistoryItemOfWriteIPFSHash(ctx, passportAddress, txHash)
	if err != nil {
		t.Errorf("Historian.GetHistoryItemOfWriteIPFSHash() error = %v", err)
	}

	if diff := deep.Equal(ehi, hi); diff != nil {
		t.Error(diff)
	}
}
