package commands

import (
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/monetha/reputation-go-sdk/cmd/internal/cmdutils"
	"github.com/monetha/reputation-go-sdk/cmd/privatedata-exchange/commands/flag"
	"github.com/monetha/reputation-go-sdk/facts"
	"github.com/monetha/reputation-go-sdk/ipfs"
	"github.com/pkg/errors"
)

// AcceptCommand handles accept command
type AcceptCommand struct {
	PassportAddress  flag.EthereumAddress         `long:"passportaddr" required:"true" description:"Ethereum address of passport contract"`
	ExchangeIndex    flag.ExchangeIndex           `long:"exchidx"      required:"true" description:"private data exchange index"`
	PassportOwnerKey flag.ECDSAPrivateKeyFromFile `long:"ownerkey"     required:"true" description:"passport owner private key filename"`
	BackendURL       string                       `long:"backendurl"   required:"true" description:"Ethereum backend URL"`
	IPFSURL          string                       `long:"ipfsurl"                      description:"IPFS node URL" default:"https://ipfs.infura.io:5001"`
	Verbosity        int                          `long:"verbosity"                    description:"log verbosity (0-9)" default:"2"`
	VModule          string                       `long:"vmodule"                      description:"log verbosity pattern"`
}

// Execute implements flags.Commander interface
func (c *AcceptCommand) Execute(args []string) error {
	fmt.Println("Accept command execution")
	fmt.Println("Passport address:", c.PassportAddress.AsCommonAddress().String())
	fmt.Println("Private exchange idx:", c.ExchangeIndex.AsBigInt().String())
	fmt.Println("Passport owner address:", crypto.PubkeyToAddress(c.PassportOwnerKey.PublicKey).String())
	fmt.Println("Backend URL:", c.BackendURL)
	fmt.Println("IPFS URL:", c.IPFSURL)

	initLogging(log.Lvl(c.Verbosity), c.VModule)
	ctx := cmdutils.CreateCtrlCContext()

	e, err := newEth(c.BackendURL)
	if err != nil {
		return err
	}

	fs, err := ipfs.New(c.IPFSURL)
	if err != nil {
		return errors.Wrap(err, "failed to create IPFS client")
	}

	err = facts.NewExchangeAcceptor(
		e,
		c.PassportOwnerKey.AsECDSAPrivateKey(),
		fs,
	).AcceptPrivateDataExchange(
		ctx,
		c.PassportAddress.AsCommonAddress(),
		c.ExchangeIndex.AsBigInt(),
	)
	if err != nil {
		return err
	}

	return nil
}
