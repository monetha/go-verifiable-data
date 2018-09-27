package factprovider

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"gitlab.com/monetha/protocol-go-sdk/contracts"
	"gitlab.com/monetha/protocol-go-sdk/eth"
)

// FactProvider contains methods to provide facts
type FactProvider eth.Session

// New converts session to FactProvider
func New(s *eth.Session) *FactProvider {
	return (*FactProvider)(s)
}

// WriteTxData writes data for the specific key
func (p *FactProvider) WriteTxData(ctx context.Context, passportAddress common.Address, key [32]byte, data []byte) error {
	backend := p.Backend
	factProviderAuth := &p.TransactOpts

	p.Log("Initialising passport", "passport", passportAddress)
	passportLogicContract, err := contracts.NewPassportLogicContract(passportAddress, backend)
	if err != nil {
		return fmt.Errorf("factprovider: NewPassportLogicContract: %v", err)
	}

	p.Log("Writing tx data to passport", "fact_provider", factProviderAuth.From.Hex(), "key", key)
	tx, err := passportLogicContract.SetTxDataBlockNumber(factProviderAuth, key, data)
	if err != nil {
		return fmt.Errorf("factprovider: SetTxDataBlockNumber: %v", err)
	}
	_, err = p.WaitForTxReceipt(ctx, tx.Hash())

	return err
}
