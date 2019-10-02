package factprovider

import (
	"github.com/ethereum/go-ethereum/common"
)

// FactProviderInfo holds fact provider information
type FactProviderInfo struct {
	Name               string
	ReputationPassport common.Address
	Website            string
}
