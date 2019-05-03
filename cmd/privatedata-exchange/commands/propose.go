package commands

import (
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/monetha/reputation-go-sdk/cmd/privatedata-exchange/commands/flag"
)

// ProposeCommand handles propose command
type ProposeCommand struct {
	PassportAddress  flag.EthereumAddress         `long:"passportaddr"     required:"true" description:"Ethereum address of passport contract"`
	FactProvider     flag.EthereumAddress         `long:"factprovideraddr" required:"true" description:"Ethereum address of fact provider"`
	FactKey          flag.FactKey                 `long:"fkey"             required:"true" description:"the key of the fact (max. 32 bytes)"`
	DataRequesterKey flag.ECDSAPrivateKeyFromFile `long:"requesterkey"     required:"true" description:"data requester private key filename"`
	StakedValue      flag.EthereumWei             `long:"stake"            required:"true" description:"amount of ethers to stake (in wei)"`
	ExchangeKeyFile  string                       `long:"exchangekey"      required:"true" description:"file name where to save the exchange key (output)"`
	BackendURL       string                       `long:"backendurl"       required:"true" description:"Ethereum backend URL"`
}

// Execute implements flags.Commander interface
func (c *ProposeCommand) Execute(args []string) error {
	fmt.Println("Propose command execution")
	fmt.Println("Passport address:", c.PassportAddress.AsCommonAddress().String())
	fmt.Println("Fact provider address:", c.FactProvider.AsCommonAddress().String())
	fmt.Println("Fact key:", c.FactKey)
	fmt.Println("Data requester address:", crypto.PubkeyToAddress(c.DataRequesterKey.PublicKey).String())
	fmt.Println("Staked amount:", c.StakedValue.EthString(), "ETH")
	fmt.Println("Backend URL:", c.BackendURL)

	return nil
}
