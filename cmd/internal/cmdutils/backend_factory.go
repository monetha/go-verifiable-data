package cmdutils

import (
	"context"

	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/monetha/go-verifiable-data/eth"
	"github.com/monetha/go-verifiable-data/eth/backend"
	"github.com/monetha/go-verifiable-data/eth/backend/ethclient"
	"github.com/monetha/go-verifiable-data/quorum"
)

// BackendFactory provides functions for creating clients, which are able to communicate
// with Ethereum backend based on configuration
type BackendFactory struct {
	quorumEnclave    *string
	quorumPrivateFor []string
}

// NewBackendFactory creates factory for creating Ethereum client
func NewBackendFactory(quorumEnclave *string, quorumPrivateFor []string) *BackendFactory {
	if len(quorumPrivateFor) > 0 && (quorumEnclave == nil || *quorumEnclave == "") {
		utils.Fatalf("Cannot use Quorum's privateFor when enclave URL is not specified")
	}

	f := &BackendFactory{
		quorumEnclave:    quorumEnclave,
		quorumPrivateFor: quorumPrivateFor,
	}

	return f
}

// DialBackendContext connects an Ethereum client to the given URL.
func (f *BackendFactory) DialBackendContext(ctx context.Context, backendURL string) (backend.Backend, error) {
	if f.isPrivateQuorum() {
		return quorum.DialContext(ctx, backendURL, f.quorumPrivateFor, *f.quorumEnclave)
	}

	return ethclient.DialContext(ctx, backendURL)
}

// DialBackend connects an Ethereum client to the given URL.
func (f *BackendFactory) DialBackend(backendURL string) (backend.Backend, error) {
	if f.isPrivateQuorum() {
		return quorum.Dial(backendURL, f.quorumPrivateFor, *f.quorumEnclave)
	}

	return ethclient.Dial(backendURL)
}

// NewEth creates new instance of Eth
func (f *BackendFactory) NewEth(ctx context.Context, b backend.Backend) *eth.Eth {
	e := eth.New(b, log.Warn)

	if f.isPrivateQuorum() {
		if pb, ok := b.(*quorum.PrivateTxClient); ok {
			e.SetTxDecrypter(func(ctx context.Context, tx *types.Transaction) (*types.Transaction, error) {
				return quorum.DecryptTx(ctx, tx, pb.GetRPCClient())
			})
		}
	}

	return e
}

func (f *BackendFactory) isPrivateQuorum() bool {
	return f.quorumEnclave != nil && *f.quorumEnclave != ""
}
