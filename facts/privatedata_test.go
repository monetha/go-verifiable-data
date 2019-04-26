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

func TestPrivateData(t *testing.T) {
	arrangeActAssert := func(actsAssert func(
		ctx context.Context,
		pa *passportWithActors,
		rd *facts.PrivateDataReader,
		wpdRes *facts.WritePrivateDataResult,
	)) {
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

		wr, err := facts.NewPrivateDataWriterWithClient(pa.FactProviderSession(), ipfsURL, c)
		if err != nil {
			t.Fatalf("facts.NewPrivateDataWriterWithClient: %v", err)
		}

		rd, err := facts.NewPrivateDataReaderWithClient(pa.Eth, ipfsURL, c)
		if err != nil {
			t.Fatalf("facts.NewPrivateDataReaderWithClient: %v", err)
		}

		wpdRes, err := wr.WritePrivateData(ctx, pa.PassportAddress, factKey, factData, rnd)
		if err != nil {
			t.Fatalf("WritePrivateData: %v", err)
		}

		actsAssert(ctx, pa, rd, wpdRes)
	}

	t.Run("ReadPrivateData", func(t *testing.T) {
		arrangeActAssert(func(
			ctx context.Context,
			pa *passportWithActors,
			rd *facts.PrivateDataReader,
			wpdRes *facts.WritePrivateDataResult,
		) {
			decryptedData, err := rd.ReadPrivateData(ctx, pa.PassportOwnerKey, pa.PassportAddress, pa.FactProviderAddress, factKey)
			if err != nil {
				t.Fatalf("ReadPrivateData: %v", err)
			}

			if bytes.Compare(factData, decryptedData) != 0 {
				t.Errorf("wanted data %v, but got %v", factData, decryptedData)
			}
		})
	})

	t.Run("ReadSecretKey", func(t *testing.T) {
		arrangeActAssert(func(
			ctx context.Context,
			pa *passportWithActors,
			rd *facts.PrivateDataReader,
			wpdRes *facts.WritePrivateDataResult,
		) {
			secretKey, err := rd.ReadSecretKey(ctx, pa.PassportOwnerKey, pa.PassportAddress, pa.FactProviderAddress, factKey)
			if err != nil {
				t.Fatalf("ReadPrivateData: %v", err)
			}

			if bytes.Compare(wpdRes.DataKey, secretKey) != 0 {
				t.Errorf("wanted secret key %v, but got %v", wpdRes.DataKey, secretKey)
			}
		})
	})

	t.Run("DecryptPrivateData", func(t *testing.T) {
		arrangeActAssert(func(
			ctx context.Context,
			pa *passportWithActors,
			rd *facts.PrivateDataReader,
			wpdRes *facts.WritePrivateDataResult,
		) {
			decryptedData, err := rd.DecryptPrivateData(ctx, wpdRes.DataIPFSHash, wpdRes.DataKey, nil)
			if err != nil {
				t.Fatalf("ReadPrivateData: %v", err)
			}

			if bytes.Compare(factData, decryptedData) != 0 {
				t.Errorf("wanted data %v, but got %v", factData, decryptedData)
			}
		})
	})

	t.Run("ReadHistoryPrivateData", func(t *testing.T) {
		arrangeActAssert(func(
			ctx context.Context,
			pa *passportWithActors,
			rd *facts.PrivateDataReader,
			wpdRes *facts.WritePrivateDataResult,
		) {
			decryptedData, err := rd.ReadHistoryPrivateData(ctx, pa.PassportOwnerKey, pa.PassportAddress, wpdRes.TransactionHash)
			if err != nil {
				t.Fatalf("ReadPrivateData: %v", err)
			}

			if bytes.Compare(factData, decryptedData) != 0 {
				t.Errorf("wanted data %v, but got %v", factData, decryptedData)
			}
		})
	})
}
