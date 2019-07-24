package quorum

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"net/url"
	"path"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rpc"
)

// NewPrivateTransactor is a utility method to easily create a transaction signer
// from a single private key for Quorum private transactions.
func NewPrivateTransactor(ctx context.Context, key *ecdsa.PrivateKey, enclaveURL string) *bind.TransactOpts {
	keyAddr := crypto.PubkeyToAddress(key.PublicKey)

	opts := &bind.TransactOpts{
		From: keyAddr,
		Signer: func(signer types.Signer, address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			if address != keyAddr {
				return nil, errors.New("not authorized to sign this account")
			}

			// Send TX data to enclave for encryption
			rawData := tx.Data()
			dataHash, err := encryptTxData(ctx, rawData, enclaveURL)
			if err != nil {
				return nil, err
			}

			var encTx *types.Transaction
			if tx.To() == nil {
				encTx = types.NewContractCreation(tx.Nonce(), tx.Value(), tx.Gas(), tx.GasPrice(), dataHash)
			} else {
				encTx = types.NewTransaction(tx.Nonce(), *tx.To(), tx.Value(), tx.Gas(), tx.GasPrice(), dataHash)
			}

			privSigner := &privateTxSigner{
				Signer: signer,
			}

			signature, err := crypto.Sign(privSigner.Hash(encTx).Bytes(), key)
			if err != nil {
				return nil, err
			}

			return encTx.WithSignature(privSigner, signature)
		},
	}

	return opts
}

func encryptTxData(ctx context.Context, data []byte, enclaveURL string) ([]byte, error) {
	data64 := base64.StdEncoding.EncodeToString(data)

	bodyBytes, err := json.Marshal(map[string]string{
		"payload": data64,
	})

	if err != nil {
		return nil, err
	}

	body := bytes.NewBuffer(bodyBytes)

	u, err := url.Parse(enclaveURL)
	if err != nil {
		return nil, err
	}

	u.Path = path.Join(u.Path, "/storeraw")
	req, err := http.NewRequest("POST", u.String(), body)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/json")

	c := &http.Client{
		Transport: &http.Transport{
			Proxy:             http.ProxyFromEnvironment,
			DisableKeepAlives: true,
		},
	}

	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		buf := new(bytes.Buffer)
		buf.ReadFrom(res.Body)

		return nil, fmt.Errorf(buf.String())
	}

	var resBody enclaveResponse
	err = json.NewDecoder(res.Body).Decode(&resBody)
	if err != nil {
		return nil, err
	}

	keyBytes, err := base64.StdEncoding.DecodeString(resBody.Key)
	if err != nil {
		return nil, err
	}

	return keyBytes, err
}

type enclaveResponse struct {
	Key string `json:"key"`
}

type privateTxSigner struct {
	types.Signer
}

// SignatureValues returns the raw R, S, V values corresponding to the
// given signature and makes transaction Quorum's private tx.
func (sg *privateTxSigner) SignatureValues(tx *types.Transaction, sig []byte) (r, s, v *big.Int, err error) {
	r, s, v, err = sg.Signer.SignatureValues(tx, sig)
	if err != nil {
		return
	}

	if IsPrivateTx(tx) {
		return
	}

	if v.Int64() == 28 {
		v.SetUint64(38)
	} else {
		v.SetUint64(37)
	}

	return
}

type signedPrivateTxFormRestorer struct {
	types.Signer
}

// SignatureValues returns R, S, V values, restored from signed private transaction
// to the form where sender's public key can be recovered from transaction
func (sg *signedPrivateTxFormRestorer) SignatureValues(tx *types.Transaction, sig []byte) (r, s, v *big.Int, err error) {
	v, r, s = tx.RawSignatureValues()

	if !IsPrivateTx(tx) {
		return
	}

	v.SetUint64(v.Uint64() - 10)

	return
}

// IsPrivateTx checks whether given transaction is Quorum's private TX
// Tx is private when V value is 37 or 38
func IsPrivateTx(tx *types.Transaction) bool {
	v, _, _ := tx.RawSignatureValues()
	if v == nil {
		return false
	}

	return (*v).Uint64() == 37 || (*v).Uint64() == 38
}

// RestoreTxToSignedForm restores private transaction, retrieved from Quorum, to the form, which resulted
// after transaction signing. This includes modifying transaction's `v` field to its original value (because it is modified for private transactions)
// so that it would be possible to recover sender's public key from transaction data.
func RestoreTxToSignedForm(tx *types.Transaction) *types.Transaction {
	tx, _ = tx.WithSignature(&signedPrivateTxFormRestorer{}, nil)
	return tx
}

// DecryptTx decrypts Quorum's private transaction.
//
// It uses private data hash from "input" field and replaces it
// with decoded data that comes from calling `eth_getQuorumPayload` RPC method.
// Provided backend object must be connected to Quorum node which is a party to this private tx.
func DecryptTx(ctx context.Context, tx *types.Transaction, client *rpc.Client) (*types.Transaction, error) {
	if !IsPrivateTx(tx) {
		return tx, nil
	}

	dataHash := hexutil.Encode(tx.Data())

	var payloadStr string

	err := client.CallContext(ctx, &payloadStr, "eth_getQuorumPayload", dataHash)
	if err != nil {
		return nil, err
	}

	payloadBytes, err := hexutil.Decode(payloadStr)
	if err != nil {
		return nil, err
	}

	if tx.To() == nil {
		return types.NewContractCreation(tx.Nonce(), tx.Value(), tx.Gas(), tx.GasPrice(), payloadBytes), nil
	}

	return types.NewTransaction(tx.Nonce(), *tx.To(), tx.Value(), tx.Gas(), tx.GasPrice(), payloadBytes), nil
}
