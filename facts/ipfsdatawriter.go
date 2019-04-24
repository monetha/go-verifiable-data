package facts

import (
	"context"
	"io"

	"github.com/ethereum/go-ethereum/common"
	"github.com/monetha/reputation-go-sdk/ipfs"
	"github.com/pkg/errors"
)

// Logger is an interface for logging
type Logger interface {
	Log(msg string, ctx ...interface{})
}

// IPFSHashWriter is an interface to write IPFS hash to Ethereum storage
type IPFSHashWriter interface {
	WriteIPFSHash(ctx context.Context, passportAddress common.Address, key [32]byte, hash string) (common.Hash, error)
}

// IPFSContentAdder is an interface to add data to IPFS
type IPFSContentAdder interface {
	Add(ctx context.Context, r io.Reader) (*ipfs.AddResult, error)
}

// IPFSDataWriter writes data to IPFS and stores IPFS hash to Ethereum storage
type IPFSDataWriter struct {
	l Logger
	w IPFSHashWriter
	a IPFSContentAdder
}

// NewIPFSDataWriter creates new instance of IPFSDataWriter
func NewIPFSDataWriter(l Logger, w IPFSHashWriter, a IPFSContentAdder) *IPFSDataWriter {
	return &IPFSDataWriter{l: l, w: w, a: a}
}

// WriteIPFSDataResult holds result of invoking WriteIPFSData
type WriteIPFSDataResult struct {
	// IPFSHash is IPFS hash of data stored in IPFS
	IPFSHash string
	// TransactionHash is hash of storing IPFS hash transaction in Ethereum network
	TransactionHash common.Hash
}

// WriteIPFSData writes data to IPFS and stores IPFS hash to Ethereum network
func (p *IPFSDataWriter) WriteIPFSData(ctx context.Context, passportAddress common.Address, key [32]byte, r io.Reader) (*WriteIPFSDataResult, error) {
	addResult, err := p.a.Add(ctx, r)
	if err != nil {
		return nil, errors.Wrap(err, "failed to add reader content to IPFS")
	}

	ipfsHash := addResult.Hash
	p.l.Log("Data successfully uploaded to IPFS...", "txHash", ipfsHash)

	txHash, err := p.w.WriteIPFSHash(ctx, passportAddress, key, ipfsHash)
	if err != nil {
		return nil, errors.Wrap(err, "failed to write IPFS hash")
	}

	return &WriteIPFSDataResult{
		IPFSHash:        ipfsHash,
		TransactionHash: txHash,
	}, nil
}
