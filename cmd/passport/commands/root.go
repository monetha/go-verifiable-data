package commands

import (
	"os"

	"github.com/monetha/go-verifiable-data/cmd"
)

// RootCommand is a root command used to parse command line arguments
type RootCommand struct {
	Version               func()                       `short:"v" long:"version" description:"Print the version of tool and exit"`
	DeployBootstrap       DeployBootstrapCommand       `command:"deploy-bootstrap"        description:"Deploy digital identity backbone contracts"`
	DeployPassportFactory DeployPassportFactoryCommand `command:"deploy-passport-factory" description:"Deploy digital identity factory contract"`
	DeployPassport        DeployPassportCommand        `command:"deploy-passport" description:"Deploy digital identity contract"`
	PassportList          PassportListCommand          `command:"passport-list" description:"Get list of digital identities from factory"`
	PassportPermission    PassportPermissionCommand    `command:"passport-permission" description:"Change permissions for digital identity"`
	UpgradePassportLogic  UpgradePassportLogicCommand  `command:"upgrade-passport-logic" description:"Upgrade digital identity logic"`
}

// Root is a prepared command to be used in command line arguments parsing
var Root RootCommand

func init() {
	Root.Version = func() {
		cmd.PrintVersion()
		os.Exit(0)
	}
}
