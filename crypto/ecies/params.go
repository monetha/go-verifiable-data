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

// This file contains parameters for ECIES encryption, specifying the
// symmetric encryption and HMAC parameters.

import (
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/elliptic"
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/subtle"
	"errors"
	"hash"
	"io"

	ethcrypto "github.com/ethereum/go-ethereum/crypto"
)

var (
	// DefaultCurve is an instance of the secp256k1 curve
	DefaultCurve = ethcrypto.S256()

	// ErrEncryptionKeyMACKeyLengthsMustBeTheSame returned when EncryptionKey and MACKey lengths don't match
	ErrEncryptionKeyMACKeyLengthsMustBeTheSame = errors.New("EncryptionKey and MACKey lengths must match")
)

// SecretKeyringMaterial hold the encryption key and MAC key
type SecretKeyringMaterial struct {
	EncryptionKey []byte
	MACKey        []byte
}

// Marshal converts secret keyring material into byte slice
func (skm *SecretKeyringMaterial) Marshal() ([]byte, error) {
	if len(skm.EncryptionKey) != len(skm.MACKey) {
		return nil, ErrEncryptionKeyMACKeyLengthsMustBeTheSame
	}

	bs := make([]byte, len(skm.EncryptionKey)+len(skm.MACKey))
	copy(bs, skm.EncryptionKey)
	copy(bs[len(skm.EncryptionKey):], skm.MACKey)

	return bs, nil
}

// Unmarshal restores secret keyring material from byte slice
func (skm *SecretKeyringMaterial) Unmarshal(bs []byte) error {
	if len(bs)%2 != 0 {
		return ErrEncryptionKeyMACKeyLengthsMustBeTheSame
	}

	keyLen := len(bs) / 2

	encryptionKey := make([]byte, keyLen)
	copy(encryptionKey, bs)

	macKey := make([]byte, keyLen)
	copy(macKey, bs[keyLen:])

	skm.EncryptionKey = encryptionKey
	skm.MACKey = macKey

	return nil
}

// EncryptedAuthenticatedMessage holds the encrypted message and HMAC
type EncryptedAuthenticatedMessage struct {
	EncryptedMessage []byte
	HMAC             []byte
}

// NewHashFun is a function type that returns a new hash.Hash computing the checksum.
type NewHashFun func() hash.Hash

// NewCipherFun is a function type that creates and returns a new cipher.Block
// The key argument should be the AES key,
// either 16, 24, or 32 bytes to select
// AES-128, AES-192, or AES-256.
type NewCipherFun func([]byte) (cipher.Block, error)

// Params holds all the parameters of selected encryption scheme
type Params struct {
	NewHash   NewHashFun // hash function
	hashAlgo  crypto.Hash
	NewCipher NewCipherFun // symmetric cipher
	BlockSize int          // block size of symmetric cipher
	KeyLen    int          // length of symmetric key
}

// Standard ECIES parameters:
// * ECIES using AES128 and HMAC-SHA-256-16
// * ECIES using AES256 and HMAC-SHA-256-32
// * ECIES using AES256 and HMAC-SHA-384-48
// * ECIES using AES256 and HMAC-SHA-512-64

var (
	// Aes128Sha256Params using AES128 and HMAC-SHA-256-16
	Aes128Sha256Params = &Params{
		NewHash:   sha256.New,
		hashAlgo:  crypto.SHA256,
		NewCipher: aes.NewCipher,
		BlockSize: aes.BlockSize,
		KeyLen:    16,
	}

	// Aes256Sha256Params using AES256 and HMAC-SHA-256-32
	Aes256Sha256Params = &Params{
		NewHash:   sha256.New,
		hashAlgo:  crypto.SHA256,
		NewCipher: aes.NewCipher,
		BlockSize: aes.BlockSize,
		KeyLen:    32,
	}

	// Aes256Sha384Params using AES256 and HMAC-SHA-384-48
	Aes256Sha384Params = &Params{
		NewHash:   sha512.New384,
		hashAlgo:  crypto.SHA384,
		NewCipher: aes.NewCipher,
		BlockSize: aes.BlockSize,
		KeyLen:    32,
	}

	// Aes256Sha512Params using AES256 and HMAC-SHA-512-64
	Aes256Sha512Params = &Params{
		NewHash:   sha512.New,
		hashAlgo:  crypto.SHA512,
		NewCipher: aes.NewCipher,
		BlockSize: aes.BlockSize,
		KeyLen:    32,
	}

	paramsFromCurve = map[elliptic.Curve]*Params{
		ethcrypto.S256(): Aes128Sha256Params,
		elliptic.P256():  Aes128Sha256Params,
		elliptic.P384():  Aes256Sha384Params,
		elliptic.P521():  Aes256Sha512Params,
	}

	// DefaultParams holds default parameters for the default curve
	DefaultParams = ParamsFromCurve(DefaultCurve)
)

// ParamsFromCurve selects parameters optimal for the selected elliptic curve.
// Only the curves P256, P384, and P512 are supported.
func ParamsFromCurve(curve elliptic.Curve) (params *Params) {
	return paramsFromCurve[curve]
}

// EncryptAuth encrypts message using provided secret keyring material and returns encrypted message with the HMAC
// s2 contains shared information that is not part of the resulting ciphertext, it's fed into the MAC. If the
// shared information parameters aren't being used, they should be nil.
func (p *Params) EncryptAuth(rand io.Reader, skm *SecretKeyringMaterial, msg, s2 []byte) (eam *EncryptedAuthenticatedMessage, err error) {
	blockSize := p.BlockSize
	encMsg, err := symEncrypt(rand, p.NewCipher, blockSize, skm.EncryptionKey, msg)
	if err != nil {
		return
	}

	msgHmac := messageTag(p.NewHash, skm.MACKey, encMsg, s2)
	eam = &EncryptedAuthenticatedMessage{
		EncryptedMessage: encMsg,
		HMAC:             msgHmac,
	}
	return
}

// DecryptAuth checks encrypted message HMAC and if it's valid decrypts the message
// s2 contains shared information that is not part of the ciphertext, it's fed into the MAC. If the
// shared information parameters aren't being used, they should be nil.
func (p *Params) DecryptAuth(skm *SecretKeyringMaterial, eam *EncryptedAuthenticatedMessage, s2 []byte) (msg []byte, err error) {
	encMsg := eam.EncryptedMessage
	calcMsgHmac := messageTag(p.NewHash, skm.MACKey, encMsg, s2)
	if subtle.ConstantTimeCompare(eam.HMAC, calcMsgHmac) != 1 {
		err = ErrInvalidMessage
		return
	}

	msg, err = symDecrypt(p.NewCipher, p.BlockSize, skm.EncryptionKey, encMsg)
	return
}

// Generate an initialisation vector for CTR mode.
func generateIV(blockSize int, rand io.Reader) (iv []byte, err error) {
	iv = make([]byte, blockSize)
	_, err = io.ReadFull(rand, iv)
	return
}

// symEncrypt carries out CTR encryption using the block cipher specified in the
// parameters.
func symEncrypt(rand io.Reader, newCipher NewCipherFun, blockSize int, key, m []byte) (ct []byte, err error) {
	c, err := newCipher(key)
	if err != nil {
		return
	}

	iv, err := generateIV(blockSize, rand)
	if err != nil {
		return
	}
	ctr := cipher.NewCTR(c, iv)

	ct = make([]byte, blockSize+len(m))
	copy(ct, iv)
	ctr.XORKeyStream(ct[blockSize:], m)
	return
}

// messageTag computes the MAC of a message (called the tag) as per
// SEC 1, 3.5.
func messageTag(newHash NewHashFun, hmacKey, msg, shared []byte) []byte {
	mac := hmac.New(newHash, hmacKey)
	mac.Write(msg)
	mac.Write(shared)
	tag := mac.Sum(nil)
	return tag
}

// DecryptAuth carries out CTR decryption using the block cipher specified in
// the parameters
func symDecrypt(newCipher NewCipherFun, blockSize int, key, ct []byte) (m []byte, err error) {
	c, err := newCipher(key)
	if err != nil {
		return
	}

	iv := ct[:blockSize]
	ctr := cipher.NewCTR(c, iv)

	m = make([]byte, len(ct)-blockSize)
	ctr.XORKeyStream(m, ct[blockSize:])
	return
}
