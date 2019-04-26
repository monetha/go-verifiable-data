package facts

import (
	"context"
	"io"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/monetha/reputation-go-sdk/eth"
	"github.com/monetha/reputation-go-sdk/ipfs"
	"github.com/pkg/errors"
)

// IPFSDataReader reads IPFS hash from Ethereum network and then retrieves data from IPFS
type IPFSDataReader struct {
	e  *eth.Eth
	fs *ipfs.IPFS
}

// NewIPFSDataReader creates new instance of IPFSDataReader
func NewIPFSDataReader(e *eth.Eth, ipfsURL string) (*IPFSDataReader, error) {
	fs, err := ipfs.New(ipfsURL)
	if err != nil {
		return nil, errors.Wrap(err, "creating instance of IPFS")
	}
	return newIPFSDataReader(e, fs), nil
}

// NewIPFSDataReaderWithClient creates new instance of IPFSDataReader using provided http client
func NewIPFSDataReaderWithClient(e *eth.Eth, ipfsURL string, c *http.Client) (*IPFSDataReader, error) {
	fs, err := ipfs.NewWithClient(ipfsURL, c)
	if err != nil {
		return nil, errors.Wrap(err, "creating instance of IPFS")
	}
	return newIPFSDataReader(e, fs), nil
}

func newIPFSDataReader(e *eth.Eth, fs *ipfs.IPFS) *IPFSDataReader {
	return &IPFSDataReader{
		e:  e,
		fs: fs,
	}
}

// ReadIPFSData reads IPFS hash from Ethereum network and then retrieves data from IPFS
func (p *IPFSDataReader) ReadIPFSData(
	ctx context.Context,
	passportAddress common.Address,
	factProviderAddress common.Address,
	factKey [32]byte,
) (io.ReadCloser, error) {
	hash, err := NewReader(p.e).ReadIPFSHash(ctx, passportAddress, factProviderAddress, factKey)
	if err != nil {
		return nil, errors.Wrap(err, "reading IPFS hash from Ethereum network")
	}

	p.log("Reading from IPFS...", "hash", hash)

	rc, err := p.fs.Cat(ctx, hash)
	if err != nil {
		return nil, errors.Wrap(err, "cat the content from IPFS failed")
	}

	return rc, nil
}

// ReadHistoryIPFSData reads IPFS hash from transaction and then retrieves data from IPFS
func (p *IPFSDataReader) ReadHistoryIPFSData(
	ctx context.Context,
	passportAddress common.Address,
	txHash common.Hash,
) (io.ReadCloser, error) {
	hi, err := NewHistorian(p.e).GetHistoryItemOfWriteIPFSHash(ctx, passportAddress, txHash)
	if err != nil {
		return nil, errors.Wrap(err, "reading IPFS hash from Ethereum transaction parameters")
	}
	hash := hi.Hash

	p.log("Reading from IPFS...", "hash", hash)

	rc, err := p.fs.Cat(ctx, hash)
	if err != nil {
		return nil, errors.Wrap(err, "cat the content from IPFS failed")
	}

	return rc, nil
}

func (p *IPFSDataReader) log(msg string, ctx ...interface{}) {
	lf := p.e.LogFun
	if lf != nil {
		lf(msg, ctx...)
	}
}
