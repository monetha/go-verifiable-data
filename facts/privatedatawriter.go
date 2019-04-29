package facts

import (
	"bytes"
	"context"
	"io"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
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
	return newPrivateDataWriter(s, fs), nil
}

// NewPrivateDataWriterWithClient creates new instance of PrivateDataWriter using provided http client
func NewPrivateDataWriterWithClient(s *eth.Session, ipfsURL string, c *http.Client) (*PrivateDataWriter, error) {
	fs, err := ipfs.NewWithClient(ipfsURL, c)
	if err != nil {
		return nil, errors.Wrap(err, "creating instance of IPFS")
	}
	return newPrivateDataWriter(s, fs), nil
}

func newPrivateDataWriter(s *eth.Session, fs *ipfs.IPFS) *PrivateDataWriter {
	return &PrivateDataWriter{
		s:  s,
		fs: fs,
	}
}

// WritePrivateDataResult holds result of invoking WritePrivateData
type WritePrivateDataResult struct {
	// DataIPFSHash is IPFS hash of encrypted private data bundle
	DataIPFSHash string
	// DataKey is secret key that was used to encrypt the data
	DataKey []byte
	// DataKeyHash is hash of secret key
	DataKeyHash [32]byte
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

	skm, skmBytes, skmHash, err := deriveSecretKeyringMaterial(ec, ownerPubKey, passportAddress, w.s.TransactOpts.From, factKey)
	if err != nil {
		return nil, err
	}

	eam, err := ec.Params().EncryptAuth(rand, skm, data, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to encrypt message")
	}

	// data to store in IPFS
	ephemeralPublicKey := ecies.MarshalPublicKey(ec.PublicKey())
	encryptedMessage := eam.EncryptedMessage
	messageHMAC := eam.HMAC

	w.log("Writing ephemeral public key to IPFS...")
	ephemeralPublicKeyAddResult, err := w.fs.Add(ctx, bytes.NewReader(ephemeralPublicKey))
	if err != nil {
		return nil, errors.Wrap(err, "failed to add ephemeral public key to IPFS")
	}
	w.log("Ephemeral public key added to IPFS", "hash", ephemeralPublicKeyAddResult.Hash, "size", ephemeralPublicKeyAddResult.Size)

	w.log("Writing encrypted message to IPFS...")
	encryptedMessageAddResult, err := w.fs.Add(ctx, bytes.NewReader(encryptedMessage))
	if err != nil {
		return nil, errors.Wrap(err, "failed to add encrypted message to IPFS")
	}
	w.log("Encrypted message added to IPFS", "hash", encryptedMessageAddResult.Hash, "size", encryptedMessageAddResult.Size)

	w.log("Writing message HMAC to IPFS...")
	messageHMACAddResult, err := w.fs.Add(ctx, bytes.NewReader(messageHMAC))
	if err != nil {
		return nil, errors.Wrap(err, "failed to add HMAC to IPFS")
	}
	w.log("Message HMAC added to IPFS", "hash", messageHMACAddResult.Hash, "size", messageHMACAddResult.Size)

	w.log("Creating directory in IPFS...")
	// create directory in IPFS
	cid, err := w.fs.DagPutLinks(ctx, []ipfs.Link{
		ephemeralPublicKeyAddResult.ToLink(ipfsPublicKeyFileName),
		encryptedMessageAddResult.ToLink(ipfsEncryptedMessageFileName),
		messageHMACAddResult.ToLink(ipfsMessageHMACFileName),
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to create directory in IPFS")
	}
	dataIPFSHash := cid.String()
	w.log("Directory created in IPFS", "hash", dataIPFSHash)

	// write hashes to Ethereum
	privateDataHashes := &PrivateDataHashes{
		DataIPFSHash: dataIPFSHash,
		DataKeyHash:  skmHash,
	}

	w.log("Writing private data hashes to Ethereum", "passport", passportAddress, "fact_key", factKey, "ipfs_hash", dataIPFSHash, "data_key_hash", skmHash.String())
	txHash, err := NewProvider(w.s).WritePrivateDataHashes(ctx, passportAddress, factKey, privateDataHashes)
	if err != nil {
		return nil, errors.Wrap(err, "failed to write private data hashes to Ethereum network")
	}

	return &WritePrivateDataResult{
		DataIPFSHash:    dataIPFSHash,
		DataKey:         skmBytes,
		DataKeyHash:     skmHash,
		TransactionHash: txHash,
	}, nil
}

func (w *PrivateDataWriter) log(msg string, ctx ...interface{}) {
	w.s.Log(msg, ctx...)
}
