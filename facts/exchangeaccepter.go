package facts

import (
	"context"
	"crypto/ecdsa"
	"crypto/subtle"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/monetha/reputation-go-sdk/contracts"
	"github.com/monetha/reputation-go-sdk/crypto/ecies"
	"github.com/monetha/reputation-go-sdk/eth"
	"github.com/monetha/reputation-go-sdk/ipfs"
	"github.com/monetha/reputation-go-sdk/types/exchange"
	"github.com/pkg/errors"
)

// ExchangeAcceptor allows passport owner to accept private data exchange proposition
type ExchangeAcceptor struct {
	passportOwnerKey *ecdsa.PrivateKey
	s                *eth.Session
	fs               *ipfs.IPFS
}

// NewExchangeAcceptor create new instance of ExchangeAcceptor
func NewExchangeAcceptor(e *eth.Eth, passportOwnerKey *ecdsa.PrivateKey, fs *ipfs.IPFS) *ExchangeAcceptor {
	return &ExchangeAcceptor{
		passportOwnerKey: passportOwnerKey,
		s:                e.NewSession(passportOwnerKey),
		fs:               fs,
	}
}

// AcceptPrivateDataExchange accepts private data exchange proposition (should be called only by the passport owner)
func (a *ExchangeAcceptor) AcceptPrivateDataExchange(ctx context.Context, passportAddress common.Address, exchangeIdx *big.Int) error {
	backend := a.s.Backend

	c := contracts.InitPassportLogicContract(passportAddress, backend)

	ex, err := c.PrivateDataExchanges(&bind.CallOpts{Context: ctx}, exchangeIdx)
	if err != nil {
		return errors.Wrap(err, "failed to get proposed private data exchange")
	}

	// it should be proposed
	exState := exchange.StateType(ex.State)
	if exState != exchange.Proposed {
		return ErrExchangeMustBeProposed
	}

	// TODO: add expiration check

	// decrypt and check hash of ExchangeKey
	publicKey := &ecdsa.PublicKey{}
	if err := ecies.UnmarshalPublicKey(a.passportOwnerKey.Curve, ex.EncryptedExchangeKey, publicKey); err != nil {
		return errors.Wrap(err, "failed to parse encrypted exchange key")
	}

	ec, err := ecies.New(a.passportOwnerKey)
	if err != nil {
		return errors.Wrap(err, "failed to create ECIES instance")
	}

	_, exchangeKey, exchangeKeyHash, err := deriveSecretKeyringMaterial(ec, publicKey, passportAddress, ex.FactProvider, ex.Key)
	if err != nil {
		return err
	}

	if subtle.ConstantTimeCompare(ex.ExchangeKeyHash[:], exchangeKeyHash[:]) != 1 {
		return errors.New("proposed exchange has invalid exchange key hash")
	}

	// decrypt data secret key
	dataSecretKey, err := NewPrivateDataReader(a.s.Eth, a.fs).
		DecryptSecretKey(
			ctx,
			a.passportOwnerKey,
			&PrivateDataHashes{
				DataIPFSHash: ex.DataIPFSHash,
				DataKeyHash:  ex.DataKeyHash,
			},
			passportAddress,
			ex.FactProvider,
			ex.Key,
		)
	if err != nil {
		return errors.Wrap(err, "failed to decrypt data encryption key")
	}

	// XOR data secret key with exchange key
	var encryptedDataKey [32]byte
	for i, b := range exchangeKey {
		encryptedDataKey[i] = dataSecretKey[i] ^ b
	}

	auth := a.s.TransactOpts
	auth.Context = ctx
	auth.Value = ex.DataRequesterValue // stake the same amount of ETH as data requester

	a.s.Log("Accepting private data exchange", "exchange_index", exchangeIdx.String(), "encrypted_key", encryptedDataKey)
	tx, err := c.AcceptPrivateDataExchange(&auth, exchangeIdx, encryptedDataKey)
	if err != nil {
		return errors.Wrap(err, "failed to accept proposed private data exchange")
	}

	_, err = a.s.WaitForTxReceipt(ctx, tx.Hash())
	return err
}
