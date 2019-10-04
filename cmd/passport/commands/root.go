package commands

import (
	"os"

	"github.com/monetha/go-verifiable-data/cmd"
)

// RootCommand is a root command used to parse command line arguments
type RootCommand struct {
	Version               func()                       `short:"v" long:"version"  description:"Print the version of tool and exit"`
	DeployBootstrap       DeployBootstrapCommand       `command:"bootstrap"       description:"Deploy digital identity backbone contracts"`
	DeployPassportFactory DeployPassportFactoryCommand `command:"deploy-factory"  description:"Deploy digital identity factory contract"`
	DeployPassport        DeployPassportCommand        `command:"create"          description:"Deploy digital identity contract"`
	PassportList          PassportListCommand          `command:"list"            description:"Get list of digital identities from factory"`
	PassportPermission    PassportPermissionCommand    `command:"permission"      description:"Change permissions for digital identity"`
	UpgradePassportLogic  UpgradePassportLogicCommand  `command:"upgrade-logic"   description:"Upgrade digital identity logic"`
	ReadHistory           ReadHistoryCommand           `command:"history"         description:"Read history (changes) of digital identity"`
	ReadFactTx            ReadFactTxCommand            `command:"read-fact-tx"    description:"Read digital identity fact value using transaction hash"`
	ReadFact              ReadFactCommand              `command:"read-fact"       description:"Read latest fact value from digital identity"`
	WriteFact             WriteFactCommand             `command:"write-fact"      description:"Write fact value to digital identity"`
}

// Root is a prepared command to be used in command line arguments parsing
var Root RootCommand

func init() {
	Root.Version = func() {
		cmd.PrintVersion()
		os.Exit(0)
	}
}
