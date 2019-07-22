package pass

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/monetha/go-verifiable-data/contracts"
	"github.com/monetha/go-verifiable-data/eth"
)

// PermissionWriter modifies passport permissions. Only passport owner is allowed to modify permissions.
type PermissionWriter struct {
	*eth.Session
	passportLogicContract *contracts.PassportLogicContract
}

// NewPermissionWriter converts session to PermissionWriter
func NewPermissionWriter(s *eth.Session, passport common.Address) *PermissionWriter {
	return &PermissionWriter{Session: s, passportLogicContract: contracts.InitPassportLogicContract(passport, s.Backend)}
}

// SetWhitelistOnlyPermission enables or disables the use of a whitelist of fact providers.
func (w *PermissionWriter) SetWhitelistOnlyPermission(ctx context.Context, onlyWhitelist bool) (common.Hash, error) {
	tx, err := w.passportLogicContract.SetWhitelistOnlyPermission(w.transactOpts(ctx), onlyWhitelist)
	if err != nil {
		return common.Hash{}, fmt.Errorf("pass: SetWhitelistOnlyPermission: %v", err)
	}
	txHash := tx.Hash()
	_, err = w.WaitForTxReceipt(ctx, txHash)

	return txHash, err
}

// AddFactProviderToWhitelist allows owner to add fact provider to the whitelist.
func (w *PermissionWriter) AddFactProviderToWhitelist(ctx context.Context, factProvider common.Address) (common.Hash, error) {
	tx, err := w.passportLogicContract.AddFactProviderToWhitelist(w.transactOpts(ctx), factProvider)
	if err != nil {
		return common.Hash{}, fmt.Errorf("pass: AddFactProviderToWhitelist: %v", err)
	}
	txHash := tx.Hash()
	_, err = w.WaitForTxReceipt(ctx, txHash)

	return txHash, err
}

// RemoveFactProviderFromWhitelist allows owner to remove fact provider from the whitelist.
func (w *PermissionWriter) RemoveFactProviderFromWhitelist(ctx context.Context, factProvider common.Address) (common.Hash, error) {
	tx, err := w.passportLogicContract.RemoveFactProviderFromWhitelist(w.transactOpts(ctx), factProvider)
	if err != nil {
		return common.Hash{}, fmt.Errorf("pass: AddFactProviderToWhitelist: %v", err)
	}
	txHash := tx.Hash()
	_, err = w.WaitForTxReceipt(ctx, txHash)

	return txHash, err
}

func (w *PermissionWriter) transactOpts(ctx context.Context) *bind.TransactOpts {
	opts := w.TransactOpts
	opts.Context = ctx
	return &opts
}
