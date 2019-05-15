package facts

import (
	"errors"
	"math/big"
	"time"
)

var (
	// ErrExchangeIsExpiredOrWillBeExpiredVerySoon returned when exchange state is expired or will be expired very soon
	ErrExchangeIsExpiredOrWillBeExpiredVerySoon = errors.New("exchange is expired or will be expired very soon")
	// ErrExchangeHasNotExpired returned when exchange has not expired
	ErrExchangeHasNotExpired = errors.New("exchange has not expired")
	// ErrExchangeMustBeProposed returned when private data exchange is not in Proposed state
	ErrExchangeMustBeProposed = errors.New("private data exchange must be in Proposed state")
	// ErrExchangeMustBeClosedOrAccepted returned when private data exchange is not in Closed or Accepted state
	ErrExchangeMustBeClosedOrAccepted = errors.New("private data exchange must be in Closed or Accepted state")
	// ErrInvalidExchangeKey returned when hash of exchange key does not match exchange key hash stored in private data exchange
	ErrInvalidExchangeKey = errors.New("invalid exchange key")
	// ErrOnlyDataRequesterAllowedToCall returned when called by non-requester account
	ErrOnlyDataRequesterAllowedToCall = errors.New("only the data requester allowed to call")
)

// Clock interface is used to "mock" time.Now and time.After
type Clock interface {
	Now() time.Time
}

type realClock struct{}

func (realClock) Now() time.Time { return time.Now() }

func isExpired(expirationTimestamp *big.Int, now time.Time) bool {
	return expirationTimestamp != nil && expirationTimestamp.Cmp(big.NewInt(now.Unix())) == -1
}
