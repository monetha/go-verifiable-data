package quorum

import (
	"context"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/rpc"
)

// PrivateTxClient is a client to work with Quorum's private transactions
type PrivateTxClient struct {
	*ethclient.Client
	c          *rpc.Client
	privateFor []string
}

type privateForParams struct {
	PrivateFor []string `json:"privateFor"`
}

// Dial connects a client to the given URL.
func Dial(rawurl string, privateFor []string) (*PrivateTxClient, error) {
	return DialContext(context.Background(), rawurl, privateFor)
}

func DialContext(ctx context.Context, rawurl string, privateFor []string) (*PrivateTxClient, error) {
	c, err := rpc.DialContext(ctx, rawurl)
	if err != nil {
		return nil, err
	}

	return NewClient(c, privateFor), nil
}

// NewClient creates a client that uses the given RPC client.
func NewClient(c *rpc.Client, privateFor []string) *PrivateTxClient {
	return &PrivateTxClient{
		c:          c,
		Client:     ethclient.NewClient(c),
		privateFor: privateFor,
	}
}

// GetRPCClient returns raw RPC client that this client is connected with
func (ec *PrivateTxClient) GetRPCClient() *rpc.Client {
	return ec.c
}

// SendTransaction injects a signed private transaction into the pending pool for execution.
//
// If the transaction was a contract creation use the TransactionReceipt method to get the
// contract address after the transaction has been mined.
func (ec *PrivateTxClient) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	data, err := rlp.EncodeToBytes(tx)
	if err != nil {
		return err
	}

	privateFor := &privateForParams{
		PrivateFor: ec.privateFor,
	}

	return ec.c.CallContext(ctx, nil, "eth_sendRawPrivateTransaction", hexutil.Encode(data), privateFor)
}
