package facts_test

import (
	"context"
	"math/big"
	"testing"

	"gitlab.com/monetha/protocol-go-sdk/facts"
)

func TestHistorian_FilterChanges(t *testing.T) {
	ctx := context.Background()

	passportAddress, factProviderSession := createPassportAndFactProviderSession(ctx, t)

	t.Logf("fact provider address: %v", factProviderSession.TransactOpts.From.Hex())

	err := facts.NewProvider(factProviderSession).WriteUint(ctx, passportAddress, [32]byte{99, 88, 77}, big.NewInt(123456))
	if err != nil {
		t.Errorf("Provider.WriteUint() error = %v", err)
	}

	hist := facts.NewHistorian(factProviderSession.Eth)
	it, err := hist.FilterChanges(nil, passportAddress)
	if err != nil {
		t.Errorf("Historian.FilterChanges() error = %v", err)
	}
	defer func() {
		if err := it.Close(); err != nil {
			t.Errorf("ChangeIterator.Close() error = %v", err)
		}
	}()

	for it.Next() {
		if err := it.Error(); err != nil {
			t.Errorf("ChangeIterator.Next() error = %v", err)
			return
		}

		t.Logf("Change: %+v", it.Change)
	}
}
