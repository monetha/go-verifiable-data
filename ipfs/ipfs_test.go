package ipfs_test

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/cheekybits/is"
	"github.com/monetha/reputation-go-sdk/ipfs"
)

func TestIPFS_Cat(t *testing.T) {
	is := is.New(t)
	expectedStr := "hello world!"

	fs := ipfs.NewWithClient("https://ipfs.infura.io:5001",
		newTestClient(func(req *http.Request) *http.Response {
			defer cleanUpRequest(req)

			if req.Method != "POST" {
				return httpResponse(http.StatusMethodNotAllowed, `method not allowed`)
			}

			switch req.URL.String() {
			case "https://ipfs.infura.io:5001/api/v0/add?pin=true&progress=false":
				return httpResponse(http.StatusOK, `{"Name":"QmTp2hEo8eXRp6wg7jXv1BLCMh5a4F3B7buAUZNZUu772j","Hash":"QmTp2hEo8eXRp6wg7jXv1BLCMh5a4F3B7buAUZNZUu772j","Size":"20"}`)
			case "https://ipfs.infura.io:5001/api/v0/cat?arg=QmTp2hEo8eXRp6wg7jXv1BLCMh5a4F3B7buAUZNZUu772j":
				return httpResponse(http.StatusOK, expectedStr)
			default:
				return httpResponse(http.StatusNotFound, `not found`)
			}
		}))

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

func cleanUpRequest(req *http.Request) {
	if req != nil {
		if rb := req.Body; rb != nil {
			_, _ = io.Copy(ioutil.Discard, rb)
			_ = rb.Close()
		}
	}
}

func copyToString(r io.Reader) (res string, err error) {
	if c, ok := r.(io.Closer); ok {
		defer c.Close()
	}

	var sb strings.Builder
	if _, err = io.Copy(&sb, r); err == nil {
		res = sb.String()
	}
	return
}

func httpResponse(statusCode int, body string) *http.Response {
	return &http.Response{
		StatusCode: statusCode,
		Body:       ioutil.NopCloser(bytes.NewBufferString(body)),
		Header:     make(http.Header),
	}
}

type roundTripFunc func(req *http.Request) *http.Response

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func newTestClient(fn roundTripFunc) *http.Client {
	return &http.Client{
		Transport: roundTripFunc(fn),
	}
}
