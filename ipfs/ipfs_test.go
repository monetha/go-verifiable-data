package ipfs_test

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/cheekybits/is"
	"github.com/dnaeon/go-vcr/recorder"
	"github.com/monetha/reputation-go-sdk/ipfs"
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

	fs := ipfs.NewWithClient("https://ipfs.infura.io:5001", &http.Client{Transport: r})

	ctx := context.TODO()
	addResult, err := fs.Add(ctx, strings.NewReader(expectedStr))
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Hash: %v", addResult.Hash)

	rd, err := fs.Cat(ctx, addResult.Hash)
	if err != nil {
		t.Fatal(err)
	}

	actualStr, err := copyToString(rd)
	if err != nil {
		t.Fatal(err)
	}

	is.Equal(expectedStr, actualStr)
}

func copyToString(r io.Reader) (res string, err error) {
	if c, ok := r.(io.Closer); ok {
		defer func() { _ = c.Close() }()
	}

	var sb strings.Builder
	if _, err = io.Copy(&sb, r); err == nil {
		res = sb.String()
	}
	return
}
