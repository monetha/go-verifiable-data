package backend

import (
	"context"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
)

// NewSimulatedBackendExtended creates a new binding backend using a simulated blockchain
// for testing purposes. It uses `backends.SimulatedBackend` under the hood, but extends it to support
// ethereum.TransactionReader interface.
func NewSimulatedBackendExtended(alloc core.GenesisAlloc, gasLimit uint64) *SimulatedBackendExt {
	return &SimulatedBackendExt{
		b: backends.NewSimulatedBackend(alloc, 10000000),
	}
}

type SimulatedBackendExt struct {
	b   *backends.SimulatedBackend
	txs sync.Map
}

func (b *SimulatedBackendExt) CodeAt(ctx context.Context, contract common.Address, blockNumber *big.Int) ([]byte, error) {
	return b.b.CodeAt(ctx, contract, blockNumber)
}

func (b *SimulatedBackendExt) CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	return b.b.CallContract(ctx, call, blockNumber)
}

func (b *SimulatedBackendExt) PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error) {
	return b.b.PendingCodeAt(ctx, account)
}

func (b *SimulatedBackendExt) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	return b.b.PendingNonceAt(ctx, account)
}

func (b *SimulatedBackendExt) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	return b.b.SuggestGasPrice(ctx)
}

func (b *SimulatedBackendExt) EstimateGas(ctx context.Context, call ethereum.CallMsg) (usedGas uint64, err error) {
	return b.b.EstimateGas(ctx, call)
}

func (b *SimulatedBackendExt) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	err := b.b.SendTransaction(ctx, tx)
	if err == nil {
		b.txs.Store(tx.Hash(), tx)
	}

	return err
}

func (b *SimulatedBackendExt) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	return b.b.TransactionReceipt(ctx, txHash)
}

func (b *SimulatedBackendExt) BalanceAt(ctx context.Context, address common.Address, blockNum *big.Int) (*big.Int, error) {
	return b.b.BalanceAt(ctx, address, blockNum)
}

func (b *SimulatedBackendExt) FilterLogs(ctx context.Context, query ethereum.FilterQuery) ([]types.Log, error) {
	return b.b.FilterLogs(ctx, query)
}

func (b *SimulatedBackendExt) SubscribeFilterLogs(ctx context.Context, query ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
	return b.b.SubscribeFilterLogs(ctx, query, ch)
}

func (b *SimulatedBackendExt) Commit() {
	b.b.Commit()
}

func (b *SimulatedBackendExt) Rollback() {
	b.b.Rollback()
}

func (b *SimulatedBackendExt) TransactionByHash(ctx context.Context, txHash common.Hash) (tx *types.Transaction, isPending bool, err error) {
	v, ok := b.txs.Load(txHash)
	if !ok {
		return nil, false, ethereum.NotFound
	}
	tx = v.(*types.Transaction)

	txr, _ := b.b.TransactionReceipt(ctx, txHash)
	isPending = txr == nil

	return
}
