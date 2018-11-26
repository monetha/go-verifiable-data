package facts_test

import (
	"context"
	"math/big"
	"testing"

	"gitlab.com/monetha/reputation-go-sdk/facts"
)

func TestFactProvider_WriteUint(t *testing.T) {
	tests := []struct {
		name string
		key  [32]byte
		data *big.Int
		err  error
	}{
		{"0", [32]byte{9, 8, 7, 6}, bigInt(0), nil},
		{"2^256-1", [32]byte{9, 8, 7, 6}, sub(exp(bigInt(2), bigInt(256)), bigInt(1)), nil},
		{"-1 should fail", [32]byte{9, 8, 7, 6}, bigInt(-1), facts.ErrOutOfUint256Range},
		{"2^256 should fail", [32]byte{9, 8, 7, 6}, exp(bigInt(2), bigInt(256)), facts.ErrOutOfUint256Range},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()

			passportAddress, factProviderSession := createPassportAndFactProviderSession(ctx, t)

			_, err := facts.NewProvider(factProviderSession).WriteUint(ctx, passportAddress, tt.key, tt.data)
			if err != tt.err {
				t.Errorf("Provider.WriteUint() error = %v, want err = %v", err, tt.err)
			}
		})
	}
}

func TestFactProvider_WriteInt(t *testing.T) {
	tests := []struct {
		name string
		key  [32]byte
		data *big.Int
		err  error
	}{
		{"0", [32]byte{9, 8, 7, 6}, bigInt(0), nil},
		{"1", [32]byte{9, 8, 7, 6}, bigInt(1), nil},
		{"-1", [32]byte{9, 8, 7, 6}, bigInt(-1), nil},
		{"2^255-1", [32]byte{9, 8, 7, 6}, sub(exp(bigInt(2), bigInt(255)), bigInt(1)), nil},
		{"-2^255", [32]byte{9, 8, 7, 6}, neg(exp(bigInt(2), bigInt(255))), nil},
		{"2^255 should fail", [32]byte{9, 8, 7, 6}, exp(bigInt(2), bigInt(255)), facts.ErrOutOfInt256Range},
		{"-2^255-1 should fail", [32]byte{9, 8, 7, 6}, sub(neg(exp(bigInt(2), bigInt(255))), bigInt(1)), facts.ErrOutOfInt256Range},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()

			passportAddress, factProviderSession := createPassportAndFactProviderSession(ctx, t)

			_, err := facts.NewProvider(factProviderSession).WriteInt(ctx, passportAddress, tt.key, tt.data)
			if err != tt.err {
				t.Errorf("Provider.WriteInt() error = %v, want err = %v", err, tt.err)
			}
		})
	}
}
