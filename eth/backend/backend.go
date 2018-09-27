package backend

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// Backend contains all methods required for the backend operations.
type Backend interface {
	bind.ContractBackend
	ethereum.TransactionReader
	BalanceAt(ctx context.Context, address common.Address, blockNum *big.Int) (*big.Int, error)
}
