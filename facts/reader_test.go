package facts_test

import (
	"bytes"
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"gitlab.com/monetha/protocol-go-sdk/facts"
)

func TestFactReader_ReadTxData(t *testing.T) {
	tests := []struct {
		name string
		key  [32]byte
		data []byte
	}{
		{"test 1", [32]byte{}, nil},
		{"test 2", [32]byte{}, []byte{}},
		{"test 3", [32]byte{}, []byte{1, 2, 3, 4}},
		{"test 4", [32]byte{9, 8, 7, 6}, []byte{1, 2, 3, 4}},
		{"test 5", [32]byte{9, 8, 7, 6}, make([]byte, 10000)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()

			passportAddress, factProviderSession := createPassportAndFactProviderSession(ctx, t)

			e := factProviderSession.Eth
			factProviderAddress := factProviderSession.TransactOpts.From

			err := facts.NewProvider(factProviderSession).WriteTxData(ctx, passportAddress, tt.key, tt.data)
			if err != nil {
				t.Errorf("FactProvider.WriteTxData() error = %v", err)
			}

			readData, err := facts.NewReader(e).ReadTxData(ctx, passportAddress, factProviderAddress, tt.key)
			if err != nil {
				t.Errorf("FactReader.ReadTxData() error = %v", err)
			}

			if !bytes.Equal(tt.data, readData) {
				t.Errorf("Expected data = %v, read data = %v", tt.data, readData)
			}
		})
	}

	t.Run("reading non existing key value", func(t *testing.T) {
		ctx := context.Background()

		passportAddress, factProviderSession := createPassportAndFactProviderSession(ctx, t)

		e := factProviderSession.Eth
		factProviderAddress := factProviderSession.TransactOpts.From

		key := [32]byte{1, 2, 3}

		_, err := facts.NewReader(e).ReadTxData(ctx, passportAddress, factProviderAddress, key)
		if err != ethereum.NotFound {
			t.Errorf("FactReader.ReadTxData() expecting error = %v, got error = %v", ethereum.NotFound, err)
		}
	})
}

func TestFactReader_ReadBytes(t *testing.T) {
	tests := []struct {
		name string
		key  [32]byte
		data []byte
	}{
		{"test 1", [32]byte{}, nil},
		{"test 2", [32]byte{}, []byte{}},
		{"test 3", [32]byte{}, []byte{1, 2, 3, 4}},
		{"test 4", [32]byte{9, 8, 7, 6}, []byte{1, 2, 3, 4}},
		{"test 5", [32]byte{9, 8, 7, 6}, make([]byte, 10000)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()

			passportAddress, factProviderSession := createPassportAndFactProviderSession(ctx, t)

			e := factProviderSession.Eth
			factProviderAddress := factProviderSession.TransactOpts.From

			err := facts.NewProvider(factProviderSession).WriteBytes(ctx, passportAddress, tt.key, tt.data)
			if err != nil {
				t.Errorf("FactProvider.WriteBytes() error = %v", err)
			}

			readData, err := facts.NewReader(e).ReadBytes(ctx, passportAddress, factProviderAddress, tt.key)
			if err != nil {
				t.Errorf("FactReader.ReadBytes() error = %v", err)
			}

			if !bytes.Equal(tt.data, readData) {
				t.Errorf("Expected data = %v, read data = %v", tt.data, readData)
			}
		})
	}

	t.Run("reading non existing key value", func(t *testing.T) {
		ctx := context.Background()

		passportAddress, factProviderSession := createPassportAndFactProviderSession(ctx, t)

		e := factProviderSession.Eth
		factProviderAddress := factProviderSession.TransactOpts.From

		key := [32]byte{1, 2, 3}

		_, err := facts.NewReader(e).ReadBytes(ctx, passportAddress, factProviderAddress, key)
		if err != ethereum.NotFound {
			t.Errorf("FactReader.ReadBytes() expecting error = %v, got error = %v", ethereum.NotFound, err)
		}
	})
}

func TestFactReader_ReadString(t *testing.T) {
	tests := []struct {
		name string
		key  [32]byte
		data string
	}{
		{"test 1", [32]byte{}, ""},
		{"test 2", [32]byte{}, "1234"},
		{"test 3", [32]byte{}, "Hello, 中国!"},
		{"test 4", [32]byte{9, 8, 7, 6}, "Hello, 中国!"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()

			passportAddress, factProviderSession := createPassportAndFactProviderSession(ctx, t)

			e := factProviderSession.Eth
			factProviderAddress := factProviderSession.TransactOpts.From

			err := facts.NewProvider(factProviderSession).WriteString(ctx, passportAddress, tt.key, tt.data)
			if err != nil {
				t.Errorf("FactProvider.WriteString() error = %v", err)
			}

			readData, err := facts.NewReader(e).ReadString(ctx, passportAddress, factProviderAddress, tt.key)
			if err != nil {
				t.Errorf("FactReader.ReadString() error = %v", err)
			}

			if tt.data != readData {
				t.Errorf("Expected data = %v, read data = %v", tt.data, readData)
			}
		})
	}

	t.Run("reading non existing key value", func(t *testing.T) {
		ctx := context.Background()

		passportAddress, factProviderSession := createPassportAndFactProviderSession(ctx, t)

		e := factProviderSession.Eth
		factProviderAddress := factProviderSession.TransactOpts.From

		key := [32]byte{1, 2, 3}

		_, err := facts.NewReader(e).ReadString(ctx, passportAddress, factProviderAddress, key)
		if err != ethereum.NotFound {
			t.Errorf("FactReader.ReadString() expecting error = %v, got error = %v", ethereum.NotFound, err)
		}
	})
}

func TestFactReader_ReadAddress(t *testing.T) {
	tests := []struct {
		name string
		key  [32]byte
		data common.Address
	}{
		{"test 1", [32]byte{}, common.Address{}},
		{"test 2", [32]byte{}, common.HexToAddress("0xaF4DcE16Da2877f8c9e00544c93B62Ac40631F16")},
		{"test 3", [32]byte{9, 8, 7, 6}, common.HexToAddress("0xaF4DcE16Da2877f8c9e00544c93B62Ac40631F16")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()

			passportAddress, factProviderSession := createPassportAndFactProviderSession(ctx, t)

			e := factProviderSession.Eth
			factProviderAddress := factProviderSession.TransactOpts.From

			err := facts.NewProvider(factProviderSession).WriteAddress(ctx, passportAddress, tt.key, tt.data)
			if err != nil {
				t.Errorf("FactProvider.WriteAddress() error = %v", err)
			}

			readData, err := facts.NewReader(e).ReadAddress(ctx, passportAddress, factProviderAddress, tt.key)
			if err != nil {
				t.Errorf("FactReader.ReadAddress() error = %v", err)
			}

			if tt.data != readData {
				t.Errorf("Expected data = %v, read data = %v", tt.data, readData)
			}
		})
	}

	t.Run("reading non existing key value", func(t *testing.T) {
		ctx := context.Background()

		passportAddress, factProviderSession := createPassportAndFactProviderSession(ctx, t)

		e := factProviderSession.Eth
		factProviderAddress := factProviderSession.TransactOpts.From

		key := [32]byte{1, 2, 3}

		_, err := facts.NewReader(e).ReadAddress(ctx, passportAddress, factProviderAddress, key)
		if err != ethereum.NotFound {
			t.Errorf("FactReader.ReadAddress() expecting error = %v, got error = %v", ethereum.NotFound, err)
		}
	})
}

func TestFactReader_ReadUint(t *testing.T) {
	tests := []struct {
		name string
		key  [32]byte
		data *big.Int
	}{
		{"test 1", [32]byte{}, new(big.Int)},
		{"test 2", [32]byte{}, big.NewInt(1234567890)},
		{"test 3", [32]byte{9, 8, 7, 6}, sub(exp(bigInt(2), bigInt(256)), bigInt(1))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()

			passportAddress, factProviderSession := createPassportAndFactProviderSession(ctx, t)

			e := factProviderSession.Eth
			factProviderAddress := factProviderSession.TransactOpts.From

			err := facts.NewProvider(factProviderSession).WriteUint(ctx, passportAddress, tt.key, tt.data)
			if err != nil {
				t.Errorf("FactProvider.WriteUint() error = %v", err)
			}

			readData, err := facts.NewReader(e).ReadUint(ctx, passportAddress, factProviderAddress, tt.key)
			if err != nil {
				t.Errorf("FactReader.ReadUint() error = %v", err)
			}

			if tt.data.Cmp(readData) != 0 {
				t.Errorf("Expected data = %v, read data = %v", tt.data.String(), readData.String())
			}
		})
	}

	t.Run("reading non existing key value", func(t *testing.T) {
		ctx := context.Background()

		passportAddress, factProviderSession := createPassportAndFactProviderSession(ctx, t)

		e := factProviderSession.Eth
		factProviderAddress := factProviderSession.TransactOpts.From

		key := [32]byte{1, 2, 3}

		_, err := facts.NewReader(e).ReadUint(ctx, passportAddress, factProviderAddress, key)
		if err != ethereum.NotFound {
			t.Errorf("FactReader.ReadUint() expecting error = %v, got error = %v", ethereum.NotFound, err)
		}
	})
}

func TestFactReader_ReadInt(t *testing.T) {
	tests := []struct {
		name string
		key  [32]byte
		data *big.Int
	}{
		{"test 1", [32]byte{}, new(big.Int)},
		{"test 2", [32]byte{}, big.NewInt(1)},
		{"test 3", [32]byte{}, big.NewInt(-1)},
		{"test 4", [32]byte{}, big.NewInt(1234567890)},
		{"test 5", [32]byte{}, big.NewInt(-1234567890)},
		{"test 6", [32]byte{9, 8, 7, 6}, sub(exp(bigInt(2), bigInt(255)), bigInt(1))},
		{"test 7", [32]byte{9, 8, 7, 6}, neg(exp(bigInt(2), bigInt(255)))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()

			passportAddress, factProviderSession := createPassportAndFactProviderSession(ctx, t)

			e := factProviderSession.Eth
			factProviderAddress := factProviderSession.TransactOpts.From

			err := facts.NewProvider(factProviderSession).WriteInt(ctx, passportAddress, tt.key, tt.data)
			if err != nil {
				t.Errorf("FactProvider.WriteInt() error = %v", err)
			}

			readData, err := facts.NewReader(e).ReadInt(ctx, passportAddress, factProviderAddress, tt.key)
			if err != nil {
				t.Errorf("FactReader.ReadInt() error = %v", err)
			}

			if tt.data.Cmp(readData) != 0 {
				t.Errorf("Expected data = %v, read data = %v", tt.data.String(), readData.String())
			}
		})
	}

	t.Run("reading non existing key value", func(t *testing.T) {
		ctx := context.Background()

		passportAddress, factProviderSession := createPassportAndFactProviderSession(ctx, t)

		e := factProviderSession.Eth
		factProviderAddress := factProviderSession.TransactOpts.From

		key := [32]byte{1, 2, 3}

		_, err := facts.NewReader(e).ReadInt(ctx, passportAddress, factProviderAddress, key)
		if err != ethereum.NotFound {
			t.Errorf("FactReader.ReadInt() expecting error = %v, got error = %v", ethereum.NotFound, err)
		}
	})
}

func TestFactReader_ReadBool(t *testing.T) {
	tests := []struct {
		name string
		key  [32]byte
		data bool
	}{
		{"test 1", [32]byte{}, false},
		{"test 2", [32]byte{}, true},
		{"test 3", [32]byte{9, 8, 7, 6}, false},
		{"test 4", [32]byte{9, 8, 7, 6}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()

			passportAddress, factProviderSession := createPassportAndFactProviderSession(ctx, t)

			e := factProviderSession.Eth
			factProviderAddress := factProviderSession.TransactOpts.From

			err := facts.NewProvider(factProviderSession).WriteBool(ctx, passportAddress, tt.key, tt.data)
			if err != nil {
				t.Errorf("FactProvider.WriteBool() error = %v", err)
			}

			readData, err := facts.NewReader(e).ReadBool(ctx, passportAddress, factProviderAddress, tt.key)
			if err != nil {
				t.Errorf("FactReader.ReadBool() error = %v", err)
			}

			if tt.data != readData {
				t.Errorf("Expected data = %v, read data = %v", tt.data, readData)
			}
		})
	}

	t.Run("reading non existing key value", func(t *testing.T) {
		ctx := context.Background()

		passportAddress, factProviderSession := createPassportAndFactProviderSession(ctx, t)

		e := factProviderSession.Eth
		factProviderAddress := factProviderSession.TransactOpts.From

		key := [32]byte{1, 2, 3}

		_, err := facts.NewReader(e).ReadBool(ctx, passportAddress, factProviderAddress, key)
		if err != ethereum.NotFound {
			t.Errorf("FactReader.ReadBool() expecting error = %v, got error = %v", ethereum.NotFound, err)
		}
	})
}
