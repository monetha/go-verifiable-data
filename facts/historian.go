package facts

import "gitlab.com/monetha/protocol-go-sdk/eth"

// Historian reads the fact history
type Historian eth.Eth

// NewHistorian converts eth to Historian
func NewHistorian(e *eth.Eth) *Historian {
	return (*Historian)(e)
}
