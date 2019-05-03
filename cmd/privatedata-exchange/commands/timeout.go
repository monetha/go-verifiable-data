package commands

import (
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/monetha/reputation-go-sdk/cmd/privatedata-exchange/commands/flag"
)

// TimeoutCommand handles timeout command
type TimeoutCommand struct {
	PassportAddress  flag.EthereumAddress         `long:"passportaddr" required:"true" description:"Ethereum address of passport contract"`
	ExchangeIndex    flag.ExchangeIndex           `long:"exchidx"      required:"true" description:"private data exchange index"`
	DataRequesterKey flag.ECDSAPrivateKeyFromFile `long:"requesterkey" required:"true" description:"data requester private key filename"`
	BackendURL       string                       `long:"backendurl"   required:"true" description:"Ethereum backend URL"`
}

// Execute implements flags.Commander interface
func (c *TimeoutCommand) Execute(args []string) error {
	fmt.Println("Timeout command execution")
	fmt.Println("Passport address:", c.PassportAddress.AsCommonAddress().String())
	fmt.Println("Private exchange idx:", c.ExchangeIndex.AsBigInt().String())
	fmt.Println("Data requester address:", crypto.PubkeyToAddress(c.DataRequesterKey.PublicKey).String())
	fmt.Println("Backend URL:", c.BackendURL)

	return nil
}
