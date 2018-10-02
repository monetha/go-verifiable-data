package txdata

import (
	"bytes"
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"gitlab.com/monetha/protocol-go-sdk/contracts"
)

func TestPassportLogicInputParser_ParseSetTxDataBlockNumberCallData(t *testing.T) {
	parser, err := NewPassportLogicInputParser()
	if err != nil {
		panic(err)
	}

	address := common.Address{}
	transactor := &transactorMock{}
	contractTransactor, err := contracts.NewPassportLogicContractTransactor(address, transactor)
	if err != nil {
		panic(err)
	}

	transactOpts := &bind.TransactOpts{Signer: bind.SignerFn(noSign)}

	tests := []struct {
		name  string
		_key  [32]byte
		_data []byte
	}{
		{"test 1", [32]byte{}, nil},
		{"test 2", [32]byte{}, ([]byte)("")},
		{"test 3", [32]byte{0x1, 0x2}, ([]byte)("test string")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx, err := contractTransactor.SetTxDataBlockNumber(transactOpts, tt._key, tt._data)
			if err != nil {
				t.Errorf("contractTransactor.SetTxDataBlockNumber: %v", err)
			}

			params, err := parser.ParseSetTxDataBlockNumberCallData(tx.Data())
			if err != nil {
				t.Errorf("parser.ParseSetTxDataBlockNumberCallData: %v", err)
			}

			if !bytes.Equal(tt._key[:], params.Key[:]) {
				was := common.Bytes2Hex(tt._key[:])
				exp := common.Bytes2Hex(params.Key[:])
				t.Errorf("expected key %v, got %v", exp, was)
			}

			if !bytes.Equal(tt._data, params.Data) {
				was := common.Bytes2Hex(tt._data)
				exp := common.Bytes2Hex(params.Data)
				t.Errorf("expected data %v, got %v", exp, was)
			}
		})
	}
}

func noSign(signer types.Signer, from common.Address, tx *types.Transaction) (*types.Transaction, error) {
	return tx, nil
}

type transactorMock struct {
}

func (*transactorMock) PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error) {
	return []byte{1}, nil
}

func (*transactorMock) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	return 1, nil
}

func (*transactorMock) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	return big.NewInt(1), nil
}

func (*transactorMock) EstimateGas(ctx context.Context, call ethereum.CallMsg) (gas uint64, err error) {
	return 1, nil
}

func (*transactorMock) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	return nil
}
