package facts

import (
	"bytes"
	"context"
	"io"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/monetha/reputation-go-sdk/crypto/ecies"
	"github.com/monetha/reputation-go-sdk/eth"
	"github.com/monetha/reputation-go-sdk/ipfs"
	"github.com/pkg/errors"
)

// PrivateDataWriter allows to encrypt private data, add encrypted content to IPFS and write hashes of ecnrypted private data to Ethereum network
type PrivateDataWriter struct {
	s  *eth.Session
	fs *ipfs.IPFS
}

// NewPrivateDataWriter creates new instance of PrivateDataWriter
func NewPrivateDataWriter(s *eth.Session, ipfsURL string) (*PrivateDataWriter, error) {
	fs, err := ipfs.New(ipfsURL)
	if err != nil {
		return nil, errors.Wrap(err, "creating instance of IPFS")
	}
	return newPrivateDatatWriter(s, fs), nil
}

// NewPrivateDataWriterWithClient creates new instance of PrivateDataWriter using provided http client
func NewPrivateDataWriterWithClient(s *eth.Session, ipfsURL string, c *http.Client) (*PrivateDataWriter, error) {
	fs, err := ipfs.NewWithClient(ipfsURL, c)
	if err != nil {
		return nil, errors.Wrap(err, "creating instance of IPFS")
	}
	return newPrivateDatatWriter(s, fs), nil
}

func newPrivateDatatWriter(s *eth.Session, fs *ipfs.IPFS) *PrivateDataWriter {
	return &PrivateDataWriter{
		s:  s,
		fs: fs,
	}
}

// WritePrivateDataResult holds result of invoking WritePrivateData
type WritePrivateDataResult struct {
	// SecretKeyringMaterial are bytes of secret keyring material
	SecretKeyringMaterial []byte
	// DataIPFSHash is IPFS hash of encrypted private data bundle
	DataIPFSHash string
	// TransactionHash is hash of storing IPFS hashes transaction in Ethereum network
	TransactionHash common.Hash
}

// WritePrivateData encrypts private data, adds encrypted content to IPFS and then writes hashes of encrypted data to passport in Ethereum network.
// `rand` - used to generate random encryption key (use rand.Reader from "crypto/rand" package in real application)
func (w *PrivateDataWriter) WritePrivateData(
	ctx context.Context,
	passportAddress common.Address,
	factKey [32]byte,
	data []byte,
	rand io.Reader,
) (*WritePrivateDataResult, error) {
	ownerPubKey, err := NewReader(w.s.Eth).ReadOwnerPublicKey(ctx, passportAddress)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read owner public key")
	}

	ec, err := ecies.NewGenerate(ownerPubKey.Curve, rand)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create ECIES instance")
	}

	// using [Provider Address + Passport Address + factKey] as seed to derive secret keyring material
	skmSeed := concatSlices(
		w.s.TransactOpts.From.Bytes(),
		passportAddress.Bytes(),
		factKey[:])

	// deriving secret keyring material
	skm, err := ec.DeriveSecretKeyringMaterial(ownerPubKey, skmSeed)
	if err != nil {
		return nil, errors.Wrap(err, "failed to derive secret keyring material")
	}

	skmBytes, err := skm.Marshal()
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal secret keyring material")
	}
	skmHash := crypto.Keccak256Hash(skmBytes)

	encSeed := skmSeed // re-using the same seed for HMAC
	eam, err := ec.Params().EncryptAuth(rand, skm, data, encSeed)
	if err != nil {
		return nil, errors.Wrap(err, "failed to encrypt message")
	}

	// data to store in IPFS
	ephemeralPublicKey := ecies.MarshalPublicKey(ec.PublicKey())
	encryptedMessage := eam.EncryptedMessage
	messageHMAC := eam.HMAC

	ephemeralPublicKeyAddResult, err := w.fs.Add(ctx, bytes.NewReader(ephemeralPublicKey))
	if err != nil {
		return nil, errors.Wrap(err, "failed to add ephemeral public key to IPFS")
	}

	encryptedMessageAddResult, err := w.fs.Add(ctx, bytes.NewReader(encryptedMessage))
	if err != nil {
		return nil, errors.Wrap(err, "failed to add encrypted message to IPFS")
	}

	messageHMACAddResult, err := w.fs.Add(ctx, bytes.NewReader(messageHMAC))
	if err != nil {
		return nil, errors.Wrap(err, "failed to add HMAC to IPFS")
	}

	// create directory in IPFS
	cid, err := w.fs.DagPutLinks(ctx, []ipfs.Link{
		ephemeralPublicKeyAddResult.ToLink("public_key"),
		encryptedMessageAddResult.ToLink("encrypted_message"),
		messageHMACAddResult.ToLink("hmac"),
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to create directory in IPFS")
	}

	// write hashes to Ethereum
	privateDataHashes := &PrivateDataHashes{
		DataIPFSHash: cid.String(),
		DataKeyHash:  skmHash,
	}

	txHash, err := NewProvider(w.s).WritePrivateDataHashes(ctx, passportAddress, factKey, privateDataHashes)
	if err != nil {
		return nil, errors.Wrap(err, "failed to write private data hashes to Ethereum network")
	}

	return &WritePrivateDataResult{
		SecretKeyringMaterial: skmBytes,
		DataIPFSHash:          cid.String(),
		TransactionHash:       txHash,
	}, nil
}

func concatSlices(slices ...[]byte) []byte {
	var tmp []byte
	for _, s := range slices {
		tmp = append(tmp, s...)
	}
	return tmp
}
