package txdata

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/go-test/deep"
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
		name   string
		params *SetTxDataBlockNumberParameters
	}{
		{"test 1", &SetTxDataBlockNumberParameters{[32]byte{}, []byte{}}},
		{"test 2", &SetTxDataBlockNumberParameters{[32]byte{0x1, 0x2}, ([]byte)("test string")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx, err := contractTransactor.SetTxDataBlockNumber(transactOpts, tt.params.Key, tt.params.Data)
			if err != nil {
				t.Errorf("contractTransactor.SetTxDataBlockNumber: %v", err)
			}

			params, err := parser.ParseSetTxDataBlockNumberCallData(tx.Data())
			if err != nil {
				t.Errorf("parser.ParseSetTxDataBlockNumberCallData: %v", err)
			}

			if diff := deep.Equal(tt.params, params); diff != nil {
				t.Error(diff)
			}
		})
	}
}

func TestPassportLogicInputParser_ParseSetBytesCallData(t *testing.T) {
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
		name   string
		params *SetBytesParameters
	}{
		{"test 1", &SetBytesParameters{[32]byte{}, []byte{}}},
		{"test 2", &SetBytesParameters{[32]byte{0x1, 0x2}, ([]byte)("test string")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx, err := contractTransactor.SetBytes(transactOpts, tt.params.Key, tt.params.Data)
			if err != nil {
				t.Errorf("contractTransactor.SetBytes: %v", err)
			}

			params, err := parser.ParseSetBytesCallData(tx.Data())
			if err != nil {
				t.Errorf("parser.ParseSetBytesCallData: %v", err)
			}

			if diff := deep.Equal(tt.params, params); diff != nil {
				t.Error(diff)
			}
		})
	}
}

func TestPassportLogicInputParser_ParseSetStringCallData(t *testing.T) {
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
		name   string
		params *SetStringParameters
	}{
		{"test 1", &SetStringParameters{[32]byte{}, ""}},
		{"test 2", &SetStringParameters{[32]byte{0x1, 0x2}, "test string"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx, err := contractTransactor.SetString(transactOpts, tt.params.Key, tt.params.Data)
			if err != nil {
				t.Errorf("contractTransactor.SetString: %v", err)
			}

			params, err := parser.ParseSetStringCallData(tx.Data())
			if err != nil {
				t.Errorf("parser.ParseSetStringCallData: %v", err)
			}

			if diff := deep.Equal(tt.params, params); diff != nil {
				t.Error(diff)
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
