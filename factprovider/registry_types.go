package factprovider

import (
	"github.com/ethereum/go-ethereum/common"
)

// Info holds fact provider information
type Info struct {
	Name               string
	ReputationPassport common.Address
	Website            string
}
