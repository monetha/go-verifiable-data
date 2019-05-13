package facts

import (
	"errors"
	"math/big"
	"time"
)

var (
	// ErrExchangeIsExpired returned when exchange state is expired
	ErrExchangeIsExpired = errors.New("exchange is expired or will be expired very soon")
	// ErrExchangeMustBeProposed returned when private data exchange is not in Proposed state
	ErrExchangeMustBeProposed = errors.New("private data exchange must be in Proposed state")
	// ErrExchangeMustBeClosedOrAccepted returned when private data exchange is not in Closed or Accepted state
	ErrExchangeMustBeClosedOrAccepted = errors.New("private data exchange must be in Closed or Accepted state")
	// ErrInvalidExchangeKey returned when hash of exchange key does not match exchange key hash stored in private data exchange
	ErrInvalidExchangeKey = errors.New("invalid exchange key")
)

var oneHourSecondsInt = big.NewInt(60 * 60)

func oneHourBeforeExpiration(expirationTimestamp *big.Int) bool {
	now := big.NewInt(time.Now().Unix())
	nowPlusHour := now.Add(now, oneHourSecondsInt)

	return expirationTimestamp != nil && expirationTimestamp.Cmp(nowPlusHour) == -1
}
