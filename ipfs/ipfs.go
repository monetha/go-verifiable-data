package ipfs

import (
	"context"
	"io"
	"io/ioutil"
	"net/http"

	"gitlab.com/monetha/protocol-go-sdk/ipfs/files"
)

// IPFS provides limited functionality to interact with the IPFS (https://ipfs.io)
type IPFS struct {
	url     string
	httpcli *http.Client
}

// New creates new instance of IPFS from the provided url
func New(url string) *IPFS {
	c := &http.Client{
		Transport: &http.Transport{
			Proxy:             http.ProxyFromEnvironment,
			DisableKeepAlives: true,
		},
	}

	return NewWithClient(url, c)
}

// NewWithClient creates new instance of IPFS from the provided url and http.Client
func NewWithClient(url string, c *http.Client) *IPFS {
	return &IPFS{
		url:     url,
		httpcli: c,
	}
}

type object struct {
	Hash string
}

// Add a file to ipfs from the given reader, returns the hash of the added file
func (f *IPFS) Add(ctx context.Context, r io.Reader) (string, error) {
	var rc io.ReadCloser
	if rclose, ok := r.(io.ReadCloser); ok {
		rc = rclose
	} else {
		rc = ioutil.NopCloser(r)
	}

	// handler expects an array of files
	fr := files.NewReaderFile("", "", rc, nil)
	slf := files.NewSliceFile("", "", []files.File{fr})
	fileReader := files.NewMultiFileReader(slf, true)

	var out object
	return out.Hash, f.request("add").
		Option("progress", false).
		Option("pin", true).
		Body(fileReader).
		Exec(ctx, &out)
}

// Cat the content at the given path. Callers need to drain and close the returned reader after usage.
func (f *IPFS) Cat(ctx context.Context, path string) (io.ReadCloser, error) {
	resp, err := f.request("cat", path).Send(ctx)
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, resp.Error
	}

	return resp.Output, nil
}

func (f *IPFS) request(command string, args ...string) *requestBuilder {
	return &requestBuilder{
		command: command,
		args:    args,
		ipfs:    f,
	}
}
