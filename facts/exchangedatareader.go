package facts

import (
	"context"
	"crypto/subtle"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/monetha/go-verifiable-data/contracts"
	"github.com/monetha/go-verifiable-data/eth"
	"github.com/monetha/go-verifiable-data/ipfs"
	"github.com/monetha/go-verifiable-data/types/exchange"
	"github.com/pkg/errors"
)

var (
	// ErrDecryptedSecretKeyIsInvalid returned when secret key hash does not match secret key hash stored in private data exchange
	ErrDecryptedSecretKeyIsInvalid = errors.New("decrypted data secret key is invalid, it's safe to dispute")
)

// ExchangeDataReader allows to read and decrypt private data from private data exchange
type ExchangeDataReader struct {
	e  *eth.Eth
	fs *ipfs.IPFS
}

// NewExchangeDataReader creates new instance of ExchangeDataReader
func NewExchangeDataReader(e *eth.Eth, fs *ipfs.IPFS) *ExchangeDataReader {
	return &ExchangeDataReader{e: e, fs: fs}
}

// ReadPrivateData decrypts data secret key from private data exchange, then it reads and decrypts private data.
func (r *ExchangeDataReader) ReadPrivateData(ctx context.Context, passportAddress common.Address, exchangeIdx *big.Int, exchangeKey [32]byte) ([]byte, error) {
	backend := r.e.Backend

	c := contracts.InitPassportLogicContract(passportAddress, backend)

	// getting private data exchange by index
	ex, err := c.PrivateDataExchanges(&bind.CallOpts{Context: ctx}, exchangeIdx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get private data exchange")
	}

	// it should be either closed or accepted
	exState := exchange.StateType(ex.State)
	if exState != exchange.Closed && exState != exchange.Accepted {
		return nil, ErrExchangeMustBeClosedOrAccepted
	}

	// validating exchange key by comparing hashes
	exchangeKeyHash := crypto.Keccak256Hash(exchangeKey[:])
	if subtle.ConstantTimeCompare(ex.ExchangeKeyHash[:], exchangeKeyHash[:]) != 1 {
		return nil, ErrInvalidExchangeKey
	}

	// decrypting secret key from private data exchange using exchange key
	var secretDataKey [32]byte
	for i, b := range exchangeKey {
		secretDataKey[i] = ex.EncryptedDataKey[i] ^ b
	}

	// validating secret key by comparing hashes
	secretDataKeyHash := crypto.Keccak256Hash(secretDataKey[:])
	if subtle.ConstantTimeCompare(ex.DataKeyHash[:], secretDataKeyHash[:]) != 1 {
		return nil, ErrDecryptedSecretKeyIsInvalid
	}

	return NewPrivateDataReader(r.e, r.fs).
		DecryptPrivateData(ctx, ex.DataIPFSHash, secretDataKey[:], nil)
}
