package commands

import (
	"fmt"
	"strings"
	"time"

	"github.com/monetha/go-verifiable-data/cmd/internal/flag"

	"github.com/ethereum/go-ethereum/log"
	"github.com/hako/durafmt"
	"github.com/monetha/go-verifiable-data/cmd/internal/cmdutils"
	"github.com/monetha/go-verifiable-data/facts"
	"github.com/monetha/go-verifiable-data/types/exchange"
)

// StatusCommand handles status command
type StatusCommand struct {
	PassportAddress flag.EthereumAddress `long:"passportaddr" required:"true" description:"Ethereum address of passport contract"`
	ExchangeIndex   flag.ExchangeIndex   `long:"exchidx"      required:"true" description:"private data exchange index"`
	BackendURL      string               `long:"backendurl"   required:"true" description:"Ethereum backend URL"`
	Verbosity       int                  `long:"verbosity"                    description:"log verbosity (0-9)" default:"2"`
	VModule         string               `long:"vmodule"                      description:"log verbosity pattern"`
}

// Execute implements flags.Commander interface
func (c *StatusCommand) Execute(args []string) error {
	cmdutils.InitLogging(log.Lvl(c.Verbosity), c.VModule)
	ctx := cmdutils.CreateCtrlCContext()

	e, err := cmdutils.NewEth(ctx, c.BackendURL, "", nil)
	if err != nil {
		return err
	}

	exchangeIdx := c.ExchangeIndex.AsBigInt()
	status, err := facts.NewExchangeStatusGetter(
		e,
	).GetPrivateDataExchangeStatus(
		ctx,
		c.PassportAddress.AsCommonAddress(),
		exchangeIdx,
	)
	if err != nil {
		return err
	}

	stateDetails := []string{status.State.String()}
	if status.State != exchange.Closed {
		agoDur := time.Since(status.StateExpirationTime)
		fmtStr := "expired %v ago"
		if agoDur < 0 {
			fmtStr = "expires in %v"
			agoDur = -agoDur
		}
		stateDetails = append(stateDetails, fmt.Sprintf(fmtStr, durafmt.Parse(agoDur.Round(time.Minute)).String()))
	}

	fmt.Printf("Private data exchange:               %v (%v)\n", exchangeIdx.String(), strings.Join(stateDetails, ", "))
	fmt.Printf("Data requester address:              %v\n", status.DataRequester.String())
	fmt.Printf("Data requester staked:               %v ETH\n", (*flag.EthereumWei)(status.DataRequesterStaked).EthString())
	fmt.Printf("Passport owner address:              %v\n", status.PassportOwner.String())
	fmt.Printf("Passport owner staked:               %v ETH\n", (*flag.EthereumWei)(status.PassportOwnerStaked).EthString())
	fmt.Printf("Private data fact provider:          %v\n", status.FactProvider.String())
	fmt.Printf("Private data fact key:               %v\n", string(status.FactKey[:]))
	fmt.Printf("Private data IPFS hash:              %v\n", status.DataIPFSHash)

	return nil
}
