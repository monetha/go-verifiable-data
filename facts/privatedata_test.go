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

func TestNewPrivateData_ReadPrivateData(t *testing.T) {
	// start http recorder
	r, err := recorder.New("fixtures/private-data-write-read")
	if err != nil {
		panic(err)
	}
	defer func() { _ = r.Stop() }() // Make sure recorder is stopped once done with it
	c := &http.Client{Transport: r}

	ctx := context.Background()
	rnd := rand.New(rand.NewSource(randSeed))

	pa := newPassportWithActors(ctx, t, rnd)

	passportAddress := pa.PassportAddress
	factProviderAddress := pa.FactProviderAddress

	wr, err := facts.NewPrivateDataWriterWithClient(pa.FactProviderSession(), ipfsURL, c)
	if err != nil {
		t.Fatalf("facts.NewPrivateDataWriterWithClient: %v", err)
	}

	rd, err := facts.NewPrivateDataReaderWithClient(pa.Eth, ipfsURL, c)
	if err != nil {
		t.Fatalf("facts.NewPrivateDataReaderWithClient: %v", err)
	}

	_, err = wr.WritePrivateData(ctx, passportAddress, factKey, factData, rnd)
	if err != nil {
		t.Fatalf("WritePrivateData: %v", err)
	}

	decryptedData, err := rd.ReadPrivateData(ctx, pa.PassportOwnerKey, passportAddress, factProviderAddress, factKey)
	if err != nil {
		t.Fatalf("ReadPrivateData: %v", err)
	}

	if bytes.Compare(factData, decryptedData) != 0 {
		t.Errorf("wanted data %v, but got %v", factData, decryptedData)
	}
}

func TestNewPrivateData_ReadSecretKey(t *testing.T) {
	// start http recorder
	r, err := recorder.New("fixtures/private-data-write-read")
	if err != nil {
		panic(err)
	}
	defer func() { _ = r.Stop() }() // Make sure recorder is stopped once done with it
	c := &http.Client{Transport: r}

	ctx := context.Background()
	rnd := rand.New(rand.NewSource(randSeed))

	pa := newPassportWithActors(ctx, t, rnd)

	passportAddress := pa.PassportAddress
	factProviderAddress := pa.FactProviderAddress

	wr, err := facts.NewPrivateDataWriterWithClient(pa.FactProviderSession(), ipfsURL, c)
	if err != nil {
		t.Fatalf("facts.NewPrivateDataWriterWithClient: %v", err)
	}

	rd, err := facts.NewPrivateDataReaderWithClient(pa.Eth, ipfsURL, c)
	if err != nil {
		t.Fatalf("facts.NewPrivateDataReaderWithClient: %v", err)
	}

	wpdRes, err := wr.WritePrivateData(ctx, passportAddress, factKey, factData, rnd)
	if err != nil {
		t.Fatalf("WritePrivateData: %v", err)
	}

	secretKey, err := rd.ReadSecretKey(ctx, pa.PassportOwnerKey, passportAddress, factProviderAddress, factKey)
	if err != nil {
		t.Fatalf("ReadPrivateData: %v", err)
	}

	if bytes.Compare(wpdRes.DataKey, secretKey) != 0 {
		t.Errorf("wanted secret key %v, but got %v", wpdRes.DataKey, secretKey)
	}
}

func TestNewPrivateData_DecryptPrivateData(t *testing.T) {
	// start http recorder
	r, err := recorder.New("fixtures/private-data-write-read")
	if err != nil {
		panic(err)
	}
	defer func() { _ = r.Stop() }() // Make sure recorder is stopped once done with it
	c := &http.Client{Transport: r}

	ctx := context.Background()
	rnd := rand.New(rand.NewSource(randSeed))

	pa := newPassportWithActors(ctx, t, rnd)

	passportAddress := pa.PassportAddress

	wr, err := facts.NewPrivateDataWriterWithClient(pa.FactProviderSession(), ipfsURL, c)
	if err != nil {
		t.Fatalf("facts.NewPrivateDataWriterWithClient: %v", err)
	}

	rd, err := facts.NewPrivateDataReaderWithClient(pa.Eth, ipfsURL, c)
	if err != nil {
		t.Fatalf("facts.NewPrivateDataReaderWithClient: %v", err)
	}

	wpdRes, err := wr.WritePrivateData(ctx, passportAddress, factKey, factData, rnd)
	if err != nil {
		t.Fatalf("WritePrivateData: %v", err)
	}

	decryptedData, err := rd.DecryptPrivateData(ctx, wpdRes.DataIPFSHash, wpdRes.DataKey, nil)
	if err != nil {
		t.Fatalf("ReadPrivateData: %v", err)
	}

	if bytes.Compare(factData, decryptedData) != 0 {
		t.Errorf("wanted data %v, but got %v", factData, decryptedData)
	}
}

func TestNewPrivateData_ReadHistoryPrivateData(t *testing.T) {
	// start http recorder
	r, err := recorder.New("fixtures/private-data-write-read")
	if err != nil {
		panic(err)
	}
	defer func() { _ = r.Stop() }() // Make sure recorder is stopped once done with it
	c := &http.Client{Transport: r}

	ctx := context.Background()
	rnd := rand.New(rand.NewSource(randSeed))

	pa := newPassportWithActors(ctx, t, rnd)

	passportAddress := pa.PassportAddress

	wr, err := facts.NewPrivateDataWriterWithClient(pa.FactProviderSession(), ipfsURL, c)
	if err != nil {
		t.Fatalf("facts.NewPrivateDataWriterWithClient: %v", err)
	}

	rd, err := facts.NewPrivateDataReaderWithClient(pa.Eth, ipfsURL, c)
	if err != nil {
		t.Fatalf("facts.NewPrivateDataReaderWithClient: %v", err)
	}

	wpdRes, err := wr.WritePrivateData(ctx, passportAddress, factKey, factData, rnd)
	if err != nil {
		t.Fatalf("WritePrivateData: %v", err)
	}

	decryptedData, err := rd.ReadHistoryPrivateData(ctx, pa.PassportOwnerKey, passportAddress, wpdRes.TransactionHash)
	if err != nil {
		t.Fatalf("ReadPrivateData: %v", err)
	}

	if bytes.Compare(factData, decryptedData) != 0 {
		t.Errorf("wanted data %v, but got %v", factData, decryptedData)
	}
}
