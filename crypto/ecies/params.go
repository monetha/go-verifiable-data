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
	"crypto/sha256"
	"crypto/sha512"
	"hash"

	ethcrypto "github.com/ethereum/go-ethereum/crypto"
)

var (
	// DefaultCurve is an instance of the secp256k1 curve
	DefaultCurve = ethcrypto.S256()
)

// Params holds all the parameters of selected encryption scheme
type Params struct {
	NewHash   func() hash.Hash // hash function
	hashAlgo  crypto.Hash
	NewCipher func([]byte) (cipher.Block, error) // symmetric cipher
	BlockSize int                                // block size of symmetric cipher
	KeyLen    int                                // length of symmetric key
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
)

var paramsFromCurve = map[elliptic.Curve]*Params{
	ethcrypto.S256(): Aes128Sha256Params,
	elliptic.P256():  Aes128Sha256Params,
	elliptic.P384():  Aes256Sha384Params,
	elliptic.P521():  Aes256Sha512Params,
}

// AddParamsForCurve sets parameters for the specific elliptic curve.
func AddParamsForCurve(curve elliptic.Curve, params *Params) {
	paramsFromCurve[curve] = params
}

// ParamsFromCurve selects parameters optimal for the selected elliptic curve.
// Only the curves P256, P384, and P512 are supported.
func ParamsFromCurve(curve elliptic.Curve) (params *Params) {
	return paramsFromCurve[curve]
}
