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
	"crypto/cipher"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/hmac"
	"crypto/subtle"
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

// PublicKey is a representation of an elliptic curve public key.
type PublicKey struct {
	X *big.Int
	Y *big.Int
	elliptic.Curve
	Params *Params
}

// ExportECDSA exports an ECIES public key as an ECDSA public key.
func (pub *PublicKey) ExportECDSA() *ecdsa.PublicKey {
	return &ecdsa.PublicKey{Curve: pub.Curve, X: pub.X, Y: pub.Y}
}

// ImportECDSAPublic imports an ECDSA public key as an ECIES public key.
func ImportECDSAPublic(pub *ecdsa.PublicKey) *PublicKey {
	return &PublicKey{
		X:      pub.X,
		Y:      pub.Y,
		Curve:  pub.Curve,
		Params: ParamsFromCurve(pub.Curve),
	}
}

// PrivateKey is a representation of an elliptic curve private key.
type PrivateKey struct {
	PublicKey
	D *big.Int
}

// ExportECDSA exports an ECIES private key as an ECDSA private key.
func (prv *PrivateKey) ExportECDSA() *ecdsa.PrivateKey {
	pub := &prv.PublicKey
	pubECDSA := pub.ExportECDSA()
	return &ecdsa.PrivateKey{PublicKey: *pubECDSA, D: prv.D}
}

// ImportECDSA imports an ECDSA private key as an ECIES private key.
func ImportECDSA(prv *ecdsa.PrivateKey) *PrivateKey {
	pub := ImportECDSAPublic(&prv.PublicKey)
	return &PrivateKey{*pub, prv.D}
}

// GenerateKey generates an elliptic curve public / private keypair. If params is nil,
// the recommended default parameters for the key will be chosen.
func GenerateKey(rand io.Reader, curve elliptic.Curve, params *Params) (prv *PrivateKey, err error) {
	pb, x, y, err := elliptic.GenerateKey(curve, rand)
	if err != nil {
		return
	}
	prv = new(PrivateKey)
	prv.PublicKey.X = x
	prv.PublicKey.Y = y
	prv.PublicKey.Curve = curve
	prv.D = new(big.Int).SetBytes(pb)
	if params == nil {
		params = ParamsFromCurve(curve)
	}
	prv.PublicKey.Params = params
	return
}

// MaxSharedKeyLength returns the maximum length of the shared key the
// public key can produce.
func MaxSharedKeyLength(pub *PublicKey) int {
	return (pub.Curve.Params().BitSize + 7) / 8
}

// GenerateShared generates shared secret keys for encryption using ECDH key agreement protocol.
func (prv *PrivateKey) GenerateShared(pub *PublicKey, skLen, macLen int) (sk []byte, err error) {
	if prv.PublicKey.Curve != pub.Curve {
		return nil, ErrInvalidCurve
	}
	if skLen+macLen > MaxSharedKeyLength(pub) {
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

// messageTag computes the MAC of a message (called the tag) as per
// SEC 1, 3.5.
func messageTag(hash func() hash.Hash, hmacKey, msg, shared []byte) []byte {
	mac := hmac.New(hash, hmacKey)
	mac.Write(msg)
	mac.Write(shared)
	tag := mac.Sum(nil)
	return tag
}

// Generate an initialisation vector for CTR mode.
func generateIV(blockSize int, rand io.Reader) (iv []byte, err error) {
	iv = make([]byte, blockSize)
	_, err = io.ReadFull(rand, iv)
	return
}

// symEncrypt carries out CTR encryption using the block cipher specified in the
// parameters.
func symEncrypt(rand io.Reader, params *Params, key, m []byte) (ct []byte, err error) {
	blockSize := params.BlockSize

	c, err := params.NewCipher(key)
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

// symDecrypt carries out CTR decryption using the block cipher specified in
// the parameters
func symDecrypt(params *Params, key, ct []byte) (m []byte, err error) {
	blockSize := params.BlockSize

	c, err := params.NewCipher(key)
	if err != nil {
		return
	}

	iv := ct[:blockSize]
	ctr := cipher.NewCTR(c, iv)

	m = make([]byte, len(ct)-blockSize)
	ctr.XORKeyStream(m, ct[blockSize:])
	return
}

// Encrypt encrypts a message using ECIES as specified in SEC 1, 5.1.
//
// s1 and s2 contain shared information that is not part of the resulting
// ciphertext. s1 is fed into key derivation, s2 is fed into the MAC. If the
// shared information parameters aren't being used, they should be nil.
func Encrypt(rand io.Reader, pub *PublicKey, msg, s1, s2 []byte) (ct []byte, err error) {
	params := pub.Params
	curve := pub.Curve

	if params == nil {
		if params = ParamsFromCurve(curve); params == nil {
			err = ErrUnsupportedECIESParameters
			return
		}
	}
	ephKey, err := GenerateKey(rand, curve, params)
	if err != nil {
		return
	}

	encKey, hmacKey, err := ephKey.generateEncryptionHMacKeys(params, pub, s1)
	if err != nil {
		return
	}

	encMsg, err := symEncrypt(rand, params, encKey, msg)
	if err != nil || len(encMsg) <= params.BlockSize {
		return
	}

	msgHmac := messageTag(params.NewHash, hmacKey, encMsg, s2)

	ephPub := ephKey.PublicKey
	ephPubBs := elliptic.Marshal(curve, ephPub.X, ephPub.Y)
	ct = make([]byte, len(ephPubBs)+len(encMsg)+len(msgHmac))
	copy(ct, ephPubBs)
	copy(ct[len(ephPubBs):], encMsg)
	copy(ct[len(ephPubBs)+len(encMsg):], msgHmac)
	return
}

// Decrypt decrypts an ECIES ciphertext.
func (prv *PrivateKey) Decrypt(ct, s1, s2 []byte) (msg []byte, err error) {
	if len(ct) == 0 {
		return nil, ErrInvalidMessage
	}
	params := prv.PublicKey.Params
	curve := prv.PublicKey.Curve

	if params == nil {
		if params = ParamsFromCurve(curve); params == nil {
			err = ErrUnsupportedECIESParameters
			return
		}
	}
	h := params.NewHash()

	var (
		rLen   int
		hLen   = h.Size()
		mStart int
		mEnd   int
	)

	switch ct[0] {
	case 2, 3, 4:
		rLen = (curve.Params().BitSize + 7) / 4
		if len(ct) < (rLen + hLen + 1) {
			err = ErrInvalidMessage
			return
		}
	default:
		err = ErrInvalidPublicKey
		return
	}

	mStart = rLen
	mEnd = len(ct) - hLen

	ephPubBs := ct[:rLen]

	ephPub := new(PublicKey)
	ephPub.Curve = curve
	ephPub.X, ephPub.Y = elliptic.Unmarshal(ephPub.Curve, ephPubBs)
	if ephPub.X == nil {
		err = ErrInvalidPublicKey
		return
	}
	if !curve.IsOnCurve(ephPub.X, ephPub.Y) {
		err = ErrInvalidCurve
		return
	}

	encKey, hmacKey, err := prv.generateEncryptionHMacKeys(params, ephPub, s1)
	if err != nil {
		return
	}

	encMsg := ct[mStart:mEnd]
	msgHmac := ct[mEnd:]

	calcMsgHmac := messageTag(params.NewHash, hmacKey, encMsg, s2)
	if subtle.ConstantTimeCompare(msgHmac, calcMsgHmac) != 1 {
		err = ErrInvalidMessage
		return
	}

	msg, err = symDecrypt(params, encKey, encMsg)
	return
}

func (prv *PrivateKey) generateEncryptionHMacKeys(params *Params, pub *PublicKey, s1 []byte) (encKey, hmacKey []byte, err error) {
	keyLen := params.KeyLen

	z, err := prv.GenerateShared(pub, keyLen, keyLen)
	if err != nil {
		return
	}

	hsh := params.NewHash()

	K, err := concatKDF(hsh, z, s1, keyLen+keyLen)
	if err != nil {
		return
	}

	encKey = K[:keyLen]
	hmacKey = K[keyLen:]
	hsh.Write(hmacKey)
	hmacKey = hsh.Sum(nil)
	hsh.Reset()

	return
}