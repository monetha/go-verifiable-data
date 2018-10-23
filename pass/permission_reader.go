package pass

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"gitlab.com/monetha/protocol-go-sdk/contracts"
	"gitlab.com/monetha/protocol-go-sdk/eth"
)

// PermissionReader reads passport permissions. Anyone is allowed to read passport permissions.
type PermissionReader struct {
	*eth.Eth
	passportLogicContract *contracts.PassportLogicContract
}

// NewPermissionReader converts session to PermissionReader
func NewPermissionReader(e *eth.Eth, passport common.Address) *PermissionReader {
	return &PermissionReader{Eth: e, passportLogicContract: contracts.InitPassportLogicContract(passport, e.Backend)}
}

// IsWhitelistOnlyPermissionSet returns true when a whitelist of fact providers is enabled.
func (r *PermissionReader) IsWhitelistOnlyPermissionSet(ctx context.Context) (bool, error) {
	return r.passportLogicContract.IsWhitelistOnlyPermissionSet(&bind.CallOpts{Context: ctx})
}

// IsFactProviderInWhitelist returns true if fact provider is added to the whitelist.
func (r *PermissionReader) IsFactProviderInWhitelist(ctx context.Context, factProvider common.Address) (bool, error) {
	return r.passportLogicContract.IsFactProviderInWhitelist(&bind.CallOpts{Context: ctx}, factProvider)
}

// IsAllowedFactProvider returns true when the given address is an allowed fact provider.
func (r *PermissionReader) IsAllowedFactProvider(ctx context.Context, factProvider common.Address) (bool, error) {
	return r.passportLogicContract.IsAllowedFactProvider(&bind.CallOpts{Context: ctx}, factProvider)
}
