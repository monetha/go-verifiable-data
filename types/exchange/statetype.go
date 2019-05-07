//go:generate enumer -type=StateType

package exchange

// StateType is type for exchange state
type StateType uint8

// All exchange states
const (
	Closed   StateType = iota
	Proposed StateType = iota
	Accepted StateType = iota
)
