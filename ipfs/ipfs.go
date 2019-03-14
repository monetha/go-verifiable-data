package ipfs

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/monetha/reputation-go-sdk/ipfs/files"
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

// AddResult contains result of Add command
type AddResult struct {
	Hash string `json:"Hash"`
	Size uint64 `json:"Size,string"`
}

// Add a file to ipfs from the given reader, returns the hash of the added file
func (f *IPFS) Add(ctx context.Context, r io.Reader) (*AddResult, error) {
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

	var out AddResult
	return &out, f.request("add").
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

type dagNode struct {
	Data  string `json:"data"`
	Links []Link `json:"links"`
}

// Link represents an IPFS Merkle DAG Link between Nodes.
type Link struct {
	// multihash of the target object
	Cid Cid `json:"Cid"`
	// utf string name. should be unique per object
	Name string `json:"Name"`
	// cumulative size of target object
	Size uint64 `json:"Size"`
}

// DagPutLinks puts directory containing links and returns Cid of directory.
func (f *IPFS) DagPutLinks(ctx context.Context, links []Link) (Cid, error) {
	dagJSONBytes, err := json.Marshal(dagNode{
		Data:  "CAE=",
		Links: links,
	})
	if err != nil {
		return "", err
	}

	fr := files.NewReaderFile("", "", ioutil.NopCloser(bytes.NewBuffer(dagJSONBytes)), nil)
	slf := files.NewSliceFile("", "", []files.File{fr})
	fileReader := files.NewMultiFileReader(slf, true)

	var out struct {
		Cid Cid `json:"Cid"`
	}

	return out.Cid, f.request("dag/put").
		Option("format", "protobuf").
		Option("input-enc", "json").
		Option("pin", true).
		Body(fileReader).
		Exec(ctx, &out)
}

// DagGetLinks gets directory links.
func (f *IPFS) DagGetLinks(ctx context.Context, cid Cid) ([]Link, error) {
	var out dagNode
	return out.Links, f.request("dag/get", cid.String()).Exec(ctx, &out)
}

func (f *IPFS) request(command string, args ...string) *requestBuilder {
	return &requestBuilder{
		command: command,
		args:    args,
		ipfs:    f,
	}
}
