package facts

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/subtle"
	"net/http"
	"path"

	"github.com/ethereum/go-ethereum/common"
	"github.com/monetha/reputation-go-sdk/crypto/ecies"
	"github.com/monetha/reputation-go-sdk/eth"
	"github.com/monetha/reputation-go-sdk/ipfs"
	"github.com/pkg/errors"
)

var (
	// ErrDerivedSecretKeyringMaterialIsInvalid returned when hash of derived secret keyring material does not match hash from fact provider
	ErrDerivedSecretKeyringMaterialIsInvalid = errors.New("facts: derived secret keyring material is invalid")
)

// PrivateDataReader allows to decrypt private data
type PrivateDataReader struct {
	e  *eth.Eth
	fs *ipfs.IPFS
}

// NewPrivateDataReader creates new instance of PrivateDataReader
func NewPrivateDataReader(e *eth.Eth, ipfsURL string) (*PrivateDataReader, error) {
	fs, err := ipfs.New(ipfsURL)
	if err != nil {
		return nil, errors.Wrap(err, "creating instance of IPFS")
	}
	return newPrivateDataReader(e, fs), nil
}

// NewPrivateDataReaderWithClient creates new instance of PrivateDataReader using provided http client
func NewPrivateDataReaderWithClient(e *eth.Eth, ipfsURL string, c *http.Client) (*PrivateDataReader, error) {
	fs, err := ipfs.NewWithClient(ipfsURL, c)
	if err != nil {
		return nil, errors.Wrap(err, "creating instance of IPFS")
	}
	return newPrivateDataReader(e, fs), nil
}

func newPrivateDataReader(e *eth.Eth, fs *ipfs.IPFS) *PrivateDataReader {
	return &PrivateDataReader{
		e:  e,
		fs: fs,
	}
}

// ReadPrivateData decrypts secret key and then decrypts private data using decrypted secret key
func (r *PrivateDataReader) ReadPrivateData(
	ctx context.Context,
	passportOwnerPrivateKey *ecdsa.PrivateKey,
	passportAddress common.Address,
	factProviderAddress common.Address,
	factKey [32]byte,
) ([]byte, error) {
	factProviderHashes, err := NewReader(r.e).ReadPrivateDataHashes(ctx, passportAddress, factProviderAddress, factKey)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read private data hashes from Ethereum network")
	}

	secretKey, err := r.DecryptSecretKey(ctx, passportOwnerPrivateKey, factProviderHashes, passportAddress, factProviderAddress, factKey)
	if err != nil {
		return nil, err
	}

	return r.DecryptPrivateData(ctx, factProviderHashes.DataIPFSHash, secretKey, passportOwnerPrivateKey.Curve)
}

// DecryptSecretKey reads ephemeral public key from IPFS and derives secret keyring material using passport owner private key.
// It returns ErrDerivedSecretKeyringMaterialIsInvalid error when hash of decrypted secret keyring material does not match data key hash from fact provider.
func (r *PrivateDataReader) DecryptSecretKey(
	ctx context.Context,
	passportOwnerPrivateKey *ecdsa.PrivateKey,
	factProviderHashes *PrivateDataHashes,
	passportAddress common.Address,
	factProviderAddress common.Address,
	factKey [32]byte,
) ([]byte, error) {
	pubKeyBytes, err := r.fs.CatBytes(ctx, path.Join(factProviderHashes.DataIPFSHash, ipfsPublicKeyFileName))
	if err != nil {
		return nil, errors.Wrap(err, "failed to get ephemeral public key from IPFS")
	}

	pubKey := &ecdsa.PublicKey{}
	if err := ecies.UnmarshalPublicKey(passportOwnerPrivateKey.Curve, pubKeyBytes, pubKey); err != nil {
		return nil, errors.Wrap(err, "failed to parse ephemeral public key")
	}

	ec, err := ecies.New(passportOwnerPrivateKey)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create ECIES instance")
	}

	_, skmBytes, skmHash, err := deriveSecretKeyringMaterial(ec, pubKey, passportAddress, factProviderAddress, factKey)
	if err != nil {
		return nil, err
	}

	if subtle.ConstantTimeCompare(factProviderHashes.DataKeyHash[:], skmHash[:]) != 1 {
		return nil, ErrDerivedSecretKeyringMaterialIsInvalid
	}

	return skmBytes, nil
}

// DecryptPrivateData reads encrypted data and HMAC and decrypts data using provided secret keyring material and elliptic curve.
// Default elliptic curve is used if it's nil.
func (r *PrivateDataReader) DecryptPrivateData(
	ctx context.Context,
	dataIPFSHash string,
	secretKey []byte,
	curve elliptic.Curve,
) ([]byte, error) {
	// using default curve if not provided
	if curve == nil {
		curve = ecies.DefaultCurve
	}

	skm := unmarshalSecretKeyringMaterial(secretKey)

	encryptedMessage, err := r.fs.CatBytes(ctx, path.Join(dataIPFSHash, ipfsEncryptedMessageFileName))
	if err != nil {
		return nil, errors.Wrap(err, "failed to get encrypted message from IPFS")
	}

	hmac, err := r.fs.CatBytes(ctx, path.Join(dataIPFSHash, ipfsMessageHMACFileName))
	if err != nil {
		return nil, errors.Wrap(err, "failed to get message HMAC from IPFS")
	}

	return ecies.ParamsFromCurve(curve).
		DecryptAuth(skm, &ecies.EncryptedAuthenticatedMessage{
			EncryptedMessage: encryptedMessage,
			HMAC:             hmac,
		}, nil)
}
