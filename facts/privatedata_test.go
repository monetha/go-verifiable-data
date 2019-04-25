package facts_test

import (
	"bytes"
	"context"
	"math/rand"
	"net/http"
	"testing"

	"github.com/dnaeon/go-vcr/recorder"
	"github.com/monetha/reputation-go-sdk/facts"
)

func TestNewPrivateData_WritePrivateData_ReadPrivateData(t *testing.T) {
	// start http recorder
	r, err := recorder.New("fixtures/private-data-write-read")
	if err != nil {
		panic(err)
	}
	defer func() { _ = r.Stop() }() // Make sure recorder is stopped once done with it
	c := &http.Client{Transport: r}

	ipfsURL := "https://ipfs.infura.io:5001"

	ctx := context.Background()
	rnd := rand.New(rand.NewSource(1))

	pa := newPassportWithActors(ctx, t, rnd)

	passportAddress := pa.PassportAddress
	factProviderAddress := pa.FactProviderAddress
	factKey := [32]byte{1, 2, 3, 4, 5, 6, 7, 8, 9}
	data := []byte("this is a secret message")

	wr, err := facts.NewPrivateDataWriterWithClient(pa.FactProviderSession(), ipfsURL, c)
	if err != nil {
		t.Fatalf("facts.NewPrivateDataWriterWithClient: %v", err)
	}

	rd, err := facts.NewPrivateDataReaderWithClient(pa.Eth, ipfsURL, c)
	if err != nil {
		t.Fatalf("facts.NewPrivateDataReaderWithClient: %v", err)
	}

	_, err = wr.WritePrivateData(ctx, passportAddress, factKey, data, rnd)
	if err != nil {
		t.Fatalf("WritePrivateData: %v", err)
	}

	decryptedData, err := rd.ReadPrivateData(ctx, pa.PassportOwnerKey, passportAddress, factProviderAddress, factKey)
	if err != nil {
		t.Fatalf("ReadPrivateData: %v", err)
	}

	if bytes.Compare(data, decryptedData) != 0 {
		t.Errorf("wanted data %v, but got %v", data, decryptedData)
	}
}
