package ipfs_test

import (
	"context"
	"net/http"
	"strings"
	"testing"

	"github.com/cheekybits/is"
	"github.com/dnaeon/go-vcr/recorder"
	"github.com/monetha/go-verifiable-data/ipfs"
)

func TestIPFS_Cat(t *testing.T) {
	// start http recorder
	r, err := recorder.New("fixtures/hello-world")
	if err != nil {
		panic(err)
	}
	defer func() { _ = r.Stop() }() // Make sure recorder is stopped once done with it

	is := is.New(t)
	expectedStr := "hello world!"

	fs, err := ipfs.NewWithClient("https://ipfs.infura.io:5001", &http.Client{Transport: r})
	if err != nil {
		panic(err)
	}

	ctx := context.TODO()
	addResult, err := fs.Add(ctx, strings.NewReader(expectedStr))
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Hash: %v", addResult.Hash)

	bs, err := fs.CatBytes(ctx, addResult.Hash)
	if err != nil {
		t.Fatal(err)
	}

	actualStr := string(bs)
	is.Equal(expectedStr, actualStr)
}
