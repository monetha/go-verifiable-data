package facts_test

import (
	"bytes"
	"context"
	"io/ioutil"
	"math/rand"
	"net/http"
	"testing"

	"github.com/monetha/go-verifiable-data/ipfs"

	"github.com/dnaeon/go-vcr/recorder"
	"github.com/monetha/go-verifiable-data/facts"
)

func TestIPFSData(t *testing.T) {
	var (
		ipfsURL  = "https://ipfs.infura.io:5001"
		randSeed = int64(1)
		factKey  = [32]byte{1, 2, 3, 4, 5, 6, 7, 8, 9}
		factData = []byte("this is a public message")
	)

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

		fs, err := ipfs.NewWithClient(ipfsURL, c)
		if err != nil {
			t.Fatalf("creating instance of IPFS: %v", err)
		}

		wr := facts.NewIPFSDataWriter(pa.FactProviderSession(), fs)
		rd := facts.NewIPFSDataReader(pa.Eth, fs)

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

			if !bytes.Equal(factData, data) {
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

			if !bytes.Equal(factData, data) {
				t.Errorf("wanted data %v, but got %v", factData, data)
			}
		})
	})
}
