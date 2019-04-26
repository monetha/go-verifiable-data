package facts_test

import (
	"bytes"
	"context"
	"io/ioutil"
	"math/rand"
	"net/http"
	"testing"

	"github.com/dnaeon/go-vcr/recorder"
	"github.com/monetha/reputation-go-sdk/facts"
)

func TestIPFSData(t *testing.T) {
	arrangeActAssert := func(actAssert func(
		ctx context.Context,
		pa *passportWithActors,
		rd *facts.IPFSDataReader,
		wRes *facts.WriteIPFSDataResult,
	)) {
		// start http recorder
		r, err := recorder.New("fixtures/ipfs-data-write-read")
		if err != nil {
			panic(err)
		}
		defer func() { _ = r.Stop() }() // Make sure recorder is stopped once done with it

		c := &http.Client{Transport: r}

		ctx := context.Background()
		rnd := rand.New(rand.NewSource(randSeed))

		pa := newPassportWithActors(ctx, t, rnd)

		wr, err := facts.NewIPFSDataWriterWithClient(pa.FactProviderSession(), ipfsURL, c)
		if err != nil {
			t.Fatalf("facts.NewIPFSDataWriterWithClient: %v", err)
		}

		rd, err := facts.NewIPFSDataReaderWithClient(pa.Eth, ipfsURL, c)
		if err != nil {
			t.Fatalf("facts.NewIPFSDataReaderWithClient: %v", err)
		}

		wRes, err := wr.WriteIPFSData(ctx, pa.PassportAddress, factKey, bytes.NewReader(factData))
		if err != nil {
			t.Fatalf("WriteIPFSData: %v", err)
		}

		actAssert(ctx, pa, rd, wRes)
	}

	t.Run("ReadIPFSData", func(t *testing.T) {
		arrangeActAssert(func(
			ctx context.Context,
			pa *passportWithActors,
			rd *facts.IPFSDataReader,
			wRes *facts.WriteIPFSDataResult,
		) {
			dr, err := rd.ReadIPFSData(ctx, pa.PassportAddress, pa.FactProviderAddress, factKey)
			if err != nil {
				t.Fatalf("ReadIPFSData: %v", err)
			}
			defer func() {
				if cErr := dr.Close(); cErr != nil {
					t.Fatalf("Close data reader: %v", cErr)
				}
			}()

			data, err := ioutil.ReadAll(dr)
			if err != nil {
				t.Fatalf("Data read: %v", err)
			}

			if bytes.Compare(factData, data) != 0 {
				t.Errorf("wanted data %v, but got %v", factData, data)
			}
		})
	})

	t.Run("ReadHistoryIPFSData", func(t *testing.T) {
		arrangeActAssert(func(
			ctx context.Context,
			pa *passportWithActors,
			rd *facts.IPFSDataReader,
			wRes *facts.WriteIPFSDataResult,
		) {
			dr, err := rd.ReadHistoryIPFSData(ctx, pa.PassportAddress, wRes.TransactionHash)
			if err != nil {
				t.Fatalf("ReadIPFSData: %v", err)
			}
			defer func() {
				if cErr := dr.Close(); cErr != nil {
					t.Fatalf("Close data reader: %v", cErr)
				}
			}()

			data, err := ioutil.ReadAll(dr)
			if err != nil {
				t.Fatalf("Data read: %v", err)
			}

			if bytes.Compare(factData, data) != 0 {
				t.Errorf("wanted data %v, but got %v", factData, data)
			}
		})
	})
}
