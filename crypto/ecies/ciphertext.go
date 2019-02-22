package ecies

import (
	"crypto/elliptic"
)

// CipherText holds parts of encrypted message
type CipherText struct {
	EphemeralPublicKey *PublicKey
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

	ct.EphemeralPublicKey = &PublicKey{
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
