package commands

import (
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/monetha/reputation-go-sdk/cmd/privatedata-exchange/commands/flag"
)

// AcceptCommand handles accept command
type AcceptCommand struct {
	PassportAddress  flag.EthereumAddress         `long:"passportaddr" required:"true" description:"Ethereum address of passport contract"`
	ExchangeIndex    flag.ExchangeIndex           `long:"exchidx"      required:"true" description:"private data exchange index"`
	PassportOwnerKey flag.ECDSAPrivateKeyFromFile `long:"ownerkey"     required:"true" description:"passport owner private key filename"`
	BackendURL       string                       `long:"backendurl"   required:"true" description:"Ethereum backend URL"`
	IPFSURL          string                       `long:"ipfsurl"                      description:"IPFS node URL" default:"https://ipfs.infura.io:5001"`
}

// Execute implements flags.Commander interface
func (c *AcceptCommand) Execute(args []string) error {
	fmt.Println("Accept command execution")
	fmt.Println("Passport address:", c.PassportAddress.AsCommonAddress().String())
	fmt.Println("Private exchange idx:", c.ExchangeIndex.AsBigInt().String())
	fmt.Println("Passport owner address:", crypto.PubkeyToAddress(c.PassportOwnerKey.PublicKey).String())
	fmt.Println("Backend URL:", c.BackendURL)
	fmt.Println("IPFS URL:", c.IPFSURL)

	return nil
}
