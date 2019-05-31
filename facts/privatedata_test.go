package facts_test

import (
	"bytes"
	"context"
	"math/rand"
	"net/http"
	"testing"

	"github.com/monetha/reputation-go-sdk/ipfs"

	"github.com/dnaeon/go-vcr/recorder"
	"github.com/monetha/reputation-go-sdk/facts"
)

func TestPrivateData(t *testing.T) {
	var (
		ipfsURL  = "https://ipfs.infura.io:5001"
		randSeed = int64(1)
		factKey  = [32]byte{1, 2, 3, 4, 5, 6, 7, 8, 9}
		factData = []byte("this is a secret message")
	)

	type testContext struct {
		context.Context
		*testing.T
		*passportWithActors
		*facts.PrivateDataReader
		*facts.WritePrivateDataResult
	}

	type actAssertFun func(
		tc *testContext,
	)

	arrangeActAssert := func(t *testing.T, actAssert actAssertFun) {
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

		fs, err := ipfs.NewWithClient(ipfsURL, c)
		if err != nil {
			t.Fatalf("creating instance of IPFS: %v", err)
		}

		wr := facts.NewPrivateDataWriter(pa.FactProviderSession(), fs)
		rd := facts.NewPrivateDataReader(pa.Eth, fs)

		wpdRes, err := wr.WritePrivateData(ctx, pa.PassportAddress, factKey, factData, rnd)
		if err != nil {
			t.Fatalf("WritePrivateData: %v", err)
		}

		actAssert(&testContext{ctx, t, pa, rd, wpdRes})
	}

	t.Run("ReadPrivateData", func(t *testing.T) {
		arrangeActAssert(t, func(tc *testContext) {
			decryptedData, err := tc.ReadPrivateData(tc, tc.PassportOwnerKey, tc.PassportAddress, tc.FactProviderAddress, factKey)
			if err != nil {
				tc.Fatalf("ReadPrivateData: %v", err)
			}

			if !bytes.Equal(factData, decryptedData) {
				tc.Errorf("wanted data %v, but got %v", factData, decryptedData)
			}
		})
	})

	t.Run("ReadPrivateData with invalid passport owner key", func(t *testing.T) {
		arrangeActAssert(t, func(tc *testContext) {
			_, err := tc.ReadPrivateData(tc, tc.FactProviderKey, tc.PassportAddress, tc.FactProviderAddress, factKey)
			if err != facts.ErrInvalidPassportOwnerKey {
				tc.Errorf("ReadPrivateData: expected error %v, but got %v", facts.ErrInvalidPassportOwnerKey, err)
			}
		})
	})

	t.Run("ReadPrivateDataUsingSecretKey", func(t *testing.T) {
		arrangeActAssert(t, func(tc *testContext) {
			decryptedData, err := tc.ReadPrivateDataUsingSecretKey(tc, tc.DataKey, tc.PassportAddress, tc.FactProviderAddress, factKey)
			if err != nil {
				tc.Fatalf("ReadPrivateDataUsingSecretKey: %v", err)
			}

			if !bytes.Equal(factData, decryptedData) {
				tc.Errorf("wanted data %v, but got %v", factData, decryptedData)
			}
		})
	})

	t.Run("ReadPrivateDataUsingSecretKey with invalid secret data key", func(t *testing.T) {
		arrangeActAssert(t, func(tc *testContext) {
			invalidDataKey := make([]byte, len(tc.DataKey))
			for idx, value := range tc.DataKey {
				invalidDataKey[idx] = ^value
			}
			_, err := tc.ReadPrivateDataUsingSecretKey(tc, invalidDataKey, tc.PassportAddress, tc.FactProviderAddress, factKey)
			if err != facts.ErrInvalidSecretKey {
				tc.Errorf("ReadPrivateData: expected error %v, but got %v", facts.ErrInvalidSecretKey, err)
			}
		})
	})

	t.Run("ReadSecretKey", func(t *testing.T) {
		arrangeActAssert(t, func(tc *testContext) {
			secretKey, err := tc.ReadSecretKey(tc, tc.PassportOwnerKey, tc.PassportAddress, tc.FactProviderAddress, factKey)
			if err != nil {
				tc.Fatalf("ReadPrivateData: %v", err)
			}

			if !bytes.Equal(tc.DataKey, secretKey) {
				tc.Errorf("wanted secret key %v, but got %v", tc.DataKey, secretKey)
			}
		})
	})

	t.Run("ReadSecretKey with invalid passport owner key", func(t *testing.T) {
		arrangeActAssert(t, func(tc *testContext) {
			_, err := tc.ReadSecretKey(tc, tc.FactProviderKey, tc.PassportAddress, tc.FactProviderAddress, factKey)
			if err != facts.ErrInvalidPassportOwnerKey {
				tc.Errorf("ReadPrivateData: expected error %v, but got %v", facts.ErrInvalidPassportOwnerKey, err)
			}
		})
	})

	t.Run("DecryptPrivateData", func(t *testing.T) {
		arrangeActAssert(t, func(tc *testContext) {
			decryptedData, err := tc.DecryptPrivateData(tc, tc.DataIPFSHash, tc.DataKey, nil)
			if err != nil {
				tc.Fatalf("ReadPrivateData: %v", err)
			}

			if !bytes.Equal(factData, decryptedData) {
				tc.Errorf("wanted data %v, but got %v", factData, decryptedData)
			}
		})
	})

	t.Run("DecryptPrivateData with invalid secret data key", func(t *testing.T) {
		arrangeActAssert(t, func(tc *testContext) {
			invalidDataKey := make([]byte, len(tc.DataKey))
			for idx, value := range tc.DataKey {
				invalidDataKey[idx] = ^value
			}
			_, err := tc.DecryptPrivateData(tc, tc.DataIPFSHash, invalidDataKey, nil)
			if err != facts.ErrInvalidSecretKey {
				tc.Errorf("ReadPrivateData: expected error %v, but got %v", facts.ErrInvalidSecretKey, err)
			}
		})
	})

	t.Run("ReadHistoryPrivateData", func(t *testing.T) {
		arrangeActAssert(t, func(tc *testContext) {
			decryptedData, err := tc.ReadHistoryPrivateData(tc, tc.PassportOwnerKey, tc.PassportAddress, tc.TransactionHash)
			if err != nil {
				tc.Fatalf("ReadPrivateData: %v", err)
			}

			if !bytes.Equal(factData, decryptedData) {
				tc.Errorf("wanted data %v, but got %v", factData, decryptedData)
			}
		})
	})

	t.Run("ReadHistoryPrivateData with invalid passport owner key", func(t *testing.T) {
		arrangeActAssert(t, func(tc *testContext) {
			_, err := tc.ReadHistoryPrivateData(tc, tc.FactProviderKey, tc.PassportAddress, tc.TransactionHash)
			if err != facts.ErrInvalidPassportOwnerKey {
				tc.Errorf("ReadPrivateData: expected error %v, but got %v", facts.ErrInvalidPassportOwnerKey, err)
			}
		})
	})

	t.Run("ReadHistoryPrivateDataUsingSecretKey", func(t *testing.T) {
		arrangeActAssert(t, func(tc *testContext) {
			decryptedData, err := tc.ReadHistoryPrivateDataUsingSecretKey(tc, tc.DataKey, tc.PassportAddress, tc.TransactionHash)
			if err != nil {
				tc.Fatalf("ReadHistoryPrivateDataUsingSecretKey: %v", err)
			}

			if !bytes.Equal(factData, decryptedData) {
				tc.Errorf("wanted data %v, but got %v", factData, decryptedData)
			}
		})
	})

	t.Run("ReadHistoryPrivateDataUsingSecretKey with invalid secret data key", func(t *testing.T) {
		arrangeActAssert(t, func(tc *testContext) {
			invalidDataKey := make([]byte, len(tc.DataKey))
			for idx, value := range tc.DataKey {
				invalidDataKey[idx] = ^value
			}
			_, err := tc.ReadHistoryPrivateDataUsingSecretKey(tc, invalidDataKey, tc.PassportAddress, tc.TransactionHash)
			if err != facts.ErrInvalidSecretKey {
				tc.Errorf("ReadHistoryPrivateDataUsingSecretKey: expected error %v, but got %v", facts.ErrInvalidSecretKey, err)
			}
		})
	})
}
