package commands

import (
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/monetha/reputation-go-sdk/cmd/privatedata-exchange/commands/flag"
)

// FinishCommand handles finish command
type FinishCommand struct {
	PassportAddress flag.EthereumAddress         `long:"passportaddr" required:"true" description:"Ethereum address of passport contract"`
	ExchangeIndex   flag.ExchangeIndex           `long:"exchidx"      required:"true" description:"private data exchange index"`
	EthereumKey     flag.ECDSAPrivateKeyFromFile `long:"requesterkey" required:"true" description:"data requester (or passport owner, when expired) private key filename"`
	BackendURL      string                       `long:"backendurl"   required:"true" description:"Ethereum backend URL"`
}

// Execute implements flags.Commander interface
func (c *FinishCommand) Execute(args []string) error {
	fmt.Println("Finish command execution")
	fmt.Println("Passport address:", c.PassportAddress.AsCommonAddress().String())
	fmt.Println("Private exchange idx:", c.ExchangeIndex.AsBigInt().String())
	fmt.Println("Caller address:", crypto.PubkeyToAddress(c.EthereumKey.PublicKey).String())
	fmt.Println("Backend URL:", c.BackendURL)

	return nil
}
