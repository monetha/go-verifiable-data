// Copyright (c) 2019 Dmitrij Koniajev @ Monetha
// Copyright (c) 2013 Kyle Isom <kyle@tyrfingr.is>
// Copyright (c) 2012 The Go Authors. All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//    * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//    * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//    * Neither the name of Google Inc. nor the names of its
// contributors may be used to endorse or promote products derived from
// this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package ecies

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"fmt"
	"hash"
	"io"
	"math/big"
)

// Errors returned by the package
var (
	ErrImport                     = fmt.Errorf("ecies: failed to import key")
	ErrInvalidCurve               = fmt.Errorf("ecies: invalid elliptic curve")
	ErrInvalidParams              = fmt.Errorf("ecies: invalid ECIES parameters")
	ErrInvalidPublicKey           = fmt.Errorf("ecies: invalid public key")
	ErrSharedKeyIsPointAtInfinity = fmt.Errorf("ecies: shared key is point at infinity")
	ErrSharedKeyTooBig            = fmt.Errorf("ecies: shared key params are too big")
	ErrKeyDataTooLong             = fmt.Errorf("ecies: can't supply requested key data")
	ErrSharedTooLong              = fmt.Errorf("ecies: shared secret is too long")
	ErrInvalidMessage             = fmt.Errorf("ecies: invalid message")
	ErrUnsupportedECIESParameters = fmt.Errorf("ecies: unsupported ECIES parameters")
)

// ECIES implements Elliptic Curve Integrated Encryption Scheme
type ECIES struct {
	params *Params           // the parameters of selected encryption scheme
	prv    *ecdsa.PrivateKey // private key used for encryption/decryption
}

// NewGenerate creates an instance of ECIES by generating a public and private key pair for specified elliptic curve,
// and selecting default parameters for encryption scheme.
func NewGenerate(c elliptic.Curve, rand io.Reader) (*ECIES, error) {
	prv, err := ecdsa.GenerateKey(c, rand)
	if err != nil {
		return nil, err
	}
	return NewWithParams(prv, nil)
}

// New creates an instance of ECIES from specified private key, it tries automatically select appropriate parameters for encryption scheme.
func New(prv *ecdsa.PrivateKey) (*ECIES, error) {
	return NewWithParams(prv, nil)
}

// Must is a helper that wraps a call to a function returning (*ECIES, error)
// and panics if the error is non-nil. It is intended for use in variable initializations
// such as
//	var e = ecies.Must(ecies.NewGenerate(ecies.DefaultCurve, rand.Reader))
func Must(e *ECIES, err error) *ECIES {
	if err != nil {
		panic(err)
	}
	return e
}

// NewWithParams creates an instance of ECIES from specified private key and params.
// If params is nil, the recommended default parameters for the encryption scheme will be chosen.
func NewWithParams(prv *ecdsa.PrivateKey, params *Params) (*ECIES, error) {
	if params == nil {
		if params = ParamsFromCurve(prv.PublicKey.Curve); params == nil {
			return nil, ErrUnsupportedECIESParameters
		}
	}

	return &ECIES{
		params: params,
		prv:    prv,
	}, nil
}

// MaxSharedKeyLength returns the maximum length of the shared key the internal public key can produce.
func (ci *ECIES) MaxSharedKeyLength() int {
	return maxSharedKeyLength(&ci.prv.PublicKey)
}

// PublicKey returns pointer to the internal public key.
func (ci *ECIES) PublicKey() *ecdsa.PublicKey {
	return &ci.prv.PublicKey
}

// PrivateKey returns pointer to the internal private key.
func (ci *ECIES) PrivateKey() *ecdsa.PrivateKey {
	return ci.prv
}

// Params returns the parameters of selected encryption scheme.
func (ci *ECIES) Params() *Params {
	return ci.params
}

// DeriveSecretKeyringMaterial derives secret keyring material by computing shared secret from private and public keys and
// passing it as a parameter to the KDF.
func (ci *ECIES) DeriveSecretKeyringMaterial(pub *ecdsa.PublicKey, s1 []byte) (skm *SecretKeyringMaterial, err error) {
	prv := ci.prv
	if prv.PublicKey.Curve != pub.Curve {
		return nil, ErrInvalidCurve
	}

	params := ci.params
	keyLen := params.KeyLen
	z, err := ci.GenerateShared(pub, keyLen, keyLen)
	if err != nil {
		return
	}

	newHash := params.NewHash
	hsh := newHash()

	K, err := concatKDF(hsh, z, s1, keyLen+keyLen)
	if err != nil {
		return
	}

	encKey := K[:keyLen]
	macKey := K[keyLen:]
	hsh.Write(macKey)
	macKey = hsh.Sum(nil)

	skm = &SecretKeyringMaterial{
		EncryptionKey: encKey,
		MACKey:        macKey,
	}

	return
}

// GenerateShared generates shared secret keys for encryption using ECDH key agreement protocol.
func (ci *ECIES) GenerateShared(pub *ecdsa.PublicKey, skLen, macLen int) (sk []byte, err error) {
	prv := ci.prv
	if prv.PublicKey.Curve != pub.Curve {
		return nil, ErrInvalidCurve
	}

	if skLen+macLen > maxSharedKeyLength(pub) {
		return nil, ErrSharedKeyTooBig
	}

	x, _ := pub.Curve.ScalarMult(pub.X, pub.Y, prv.D.Bytes())
	if x == nil {
		return nil, ErrSharedKeyIsPointAtInfinity
	}

	sk = make([]byte, skLen+macLen)
	skBytes := x.Bytes()
	copy(sk[len(sk)-len(skBytes):], skBytes)
	return sk, nil
}

// maxSharedKeyLength returns the maximum length of the shared key the
// public key can produce.
func maxSharedKeyLength(pub *ecdsa.PublicKey) int {
	return (pub.Curve.Params().BitSize + 7) / 8
}

// Encrypt encrypts a message using ECIES as specified in SEC 1, 5.1.
//
// s1 and s2 contain shared information that is not part of the resulting
// ciphertext. s1 is fed into key derivation, s2 is fed into the MAC. If the
// shared information parameters aren't being used, they should be nil.
func (ci *ECIES) Encrypt(rand io.Reader, pub *ecdsa.PublicKey, msg, s1, s2 []byte) (ct []byte, err error) {
	c, err := ci.EncryptToCipherText(rand, pub, msg, s1, s2)
	if err != nil {
		return
	}

	return c.Marshal()
}

// EncryptToCipherText encrypts a message using ECIES as specified in SEC 1, 5.1.
// Instead of bytes it returns all the parts of encrypted message.
//
// s1 and s2 contain shared information that is not part of the resulting
// ciphertext. s1 is fed into key derivation, s2 is fed into the MAC. If the
// shared information parameters aren't being used, they should be nil.
func (ci *ECIES) EncryptToCipherText(rand io.Reader, pub *ecdsa.PublicKey, msg, s1, s2 []byte) (ct *CipherText, err error) {
	ciPubKey := ci.prv.PublicKey
	if ciPubKey.Curve != pub.Curve {
		return nil, ErrInvalidCurve
	}
	params := ci.params

	skm, err := ci.DeriveSecretKeyringMaterial(pub, s1)
	if err != nil {
		return
	}

	eam, err := params.EncryptAuth(rand, skm, msg, s2)
	if err != nil {
		return
	}

	ct = &CipherText{
		EphemeralPublicKey:            &ciPubKey,
		EncryptedAuthenticatedMessage: eam,
	}
	return
}

// Decrypt decrypts an ECIES ciphertext.
func (ci *ECIES) Decrypt(ct, s1, s2 []byte) (msg []byte, err error) {
	params := ci.params
	curve := ci.prv.PublicKey.Curve

	newHash := params.NewHash
	c := &CipherText{}
	err = c.Unmarshal(ct, curve, newHash().Size())
	if err != nil {
		return
	}

	skm, err := ci.DeriveSecretKeyringMaterial(c.EphemeralPublicKey, s1)
	if err != nil {
		return
	}

	return params.DecryptAuth(skm, c.EncryptedAuthenticatedMessage, s2)
}

// Encrypt encrypts a message using ECIES as specified in SEC 1, 5.1.
//
// s1 and s2 contain shared information that is not part of the resulting
// ciphertext. s1 is fed into key derivation, s2 is fed into the MAC. If the
// shared information parameters aren't being used, they should be nil.
func Encrypt(rand io.Reader, pub *ecdsa.PublicKey, msg, s1, s2 []byte) (ct []byte, err error) {
	e, err := NewGenerate(pub.Curve, rand)
	if err != nil {
		return
	}

	return e.Encrypt(rand, pub, msg, s1, s2)
}

// CipherText holds parts of encrypted message
type CipherText struct {
	EphemeralPublicKey *ecdsa.PublicKey
	*EncryptedAuthenticatedMessage
}

// Marshal converts parts of encrypted message into byte slice.
func (ct *CipherText) Marshal() ([]byte, error) {
	ephPub := ct.EphemeralPublicKey
	curve := ephPub.Curve
	encMsg := ct.EncryptedMessage
	msgHmac := ct.HMAC

	ephPubBs := elliptic.Marshal(curve, ephPub.X, ephPub.Y)
	ctBs := make([]byte, len(ephPubBs)+len(encMsg)+len(msgHmac))
	copy(ctBs, ephPubBs)
	copy(ctBs[len(ephPubBs):], encMsg)
	copy(ctBs[len(ephPubBs)+len(encMsg):], msgHmac)

	return ctBs, nil
}

// Unmarshal splits the bytes of encrypted message, serialized by Marshal, into the parts of encrypted message.
func (ct *CipherText) Unmarshal(b []byte, curve elliptic.Curve, hashSize int) error {
	if len(b) == 0 {
		return ErrInvalidMessage
	}

	var (
		rLen   int
		hLen   = hashSize
		mStart int
		mEnd   int
	)

	switch b[0] {
	case 2, 3, 4:
		rLen = (curve.Params().BitSize + 7) / 4
		if len(b) < (rLen + hLen + 1) {
			return ErrInvalidMessage
		}
	default:
		return ErrInvalidPublicKey
	}

	mStart = rLen
	mEnd = len(b) - hLen

	ephPubBs := b[:rLen]

	x, y := elliptic.Unmarshal(curve, ephPubBs)
	if x == nil {
		return ErrInvalidPublicKey
	}
	if !curve.IsOnCurve(x, y) {
		return ErrInvalidCurve
	}

	ct.EphemeralPublicKey = &ecdsa.PublicKey{
		Curve: curve,
		X:     x,
		Y:     y,
	}
	ct.EncryptedAuthenticatedMessage = &EncryptedAuthenticatedMessage{
		EncryptedMessage: b[mStart:mEnd],
		HMAC:             b[mEnd:],
	}

	return nil
}

var (
	big2To32   = new(big.Int).Exp(big.NewInt(2), big.NewInt(32), nil)
	big2To32M1 = new(big.Int).Sub(big2To32, big.NewInt(1))
)

func incCounter(ctr []byte) {
	if ctr[3]++; ctr[3] != 0 {
		return
	}
	if ctr[2]++; ctr[2] != 0 {
		return
	}
	if ctr[1]++; ctr[1] != 0 {
		return
	}
	if ctr[0]++; ctr[0] != 0 {
		return
	}
}

// NIST SP 800-56 Concatenation Key Derivation Function (see section 5.8.1).
func concatKDF(hash hash.Hash, z, s1 []byte, kdLen int) (k []byte, err error) {
	if s1 == nil {
		s1 = make([]byte, 0)
	}

	reps := ((kdLen + 7) * 8) / (hash.BlockSize() * 8)
	if big.NewInt(int64(reps)).Cmp(big2To32M1) > 0 {
		fmt.Println(big2To32M1)
		return nil, ErrKeyDataTooLong
	}

	counter := []byte{0, 0, 0, 1}
	k = make([]byte, 0)

	for i := 0; i <= reps; i++ {
		hash.Write(counter)
		hash.Write(z)
		hash.Write(s1)
		k = append(k, hash.Sum(nil)...)
		hash.Reset()
		incCounter(counter)
	}

	k = k[:kdLen]
	return
}
