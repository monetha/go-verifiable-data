package facts

import (
	"context"
	"io"

	"github.com/ethereum/go-ethereum/common"
	"github.com/monetha/reputation-go-sdk/eth"
	"github.com/monetha/reputation-go-sdk/ipfs"
	"github.com/pkg/errors"
)

// IPFSDataWriter writes data to IPFS and stores IPFS hash to Ethereum storage
type IPFSDataWriter struct {
	s  *eth.Session
	fs *ipfs.IPFS
}

// NewIPFSDataWriter creates new instance of IPFSDataWriter
func NewIPFSDataWriter(s *eth.Session, fs *ipfs.IPFS) *IPFSDataWriter {
	return &IPFSDataWriter{
		s:  s,
		fs: fs,
	}
}

// WriteIPFSDataResult holds result of invoking WriteIPFSData
type WriteIPFSDataResult struct {
	// IPFSHash is IPFS hash of data stored in IPFS
	IPFSHash string
	// TransactionHash is hash of storing IPFS hash transaction in Ethereum network
	TransactionHash common.Hash
}

// WriteIPFSData writes data to IPFS and stores IPFS hash to Ethereum network
func (p *IPFSDataWriter) WriteIPFSData(
	ctx context.Context,
	passportAddress common.Address,
	factKey [32]byte,
	r io.Reader,
) (*WriteIPFSDataResult, error) {
	addResult, err := p.fs.Add(ctx, r)
	if err != nil {
		return nil, errors.Wrap(err, "failed to add reader content to IPFS")
	}

	ipfsHash := addResult.Hash
	p.log("Data successfully uploaded to IPFS...", "txHash", ipfsHash)

	txHash, err := NewProvider(p.s).WriteIPFSHash(ctx, passportAddress, factKey, ipfsHash)
	if err != nil {
		return nil, errors.Wrap(err, "failed to write IPFS hash")
	}

	return &WriteIPFSDataResult{
		IPFSHash:        ipfsHash,
		TransactionHash: txHash,
	}, nil
}

func (p *IPFSDataWriter) log(msg string, ctx ...interface{}) {
	p.s.Log(msg, ctx...)
}
