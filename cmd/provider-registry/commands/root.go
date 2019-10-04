package commands

import (
	"os"

	"github.com/monetha/go-verifiable-data/cmd"
)

// RootCommand is a root command used to parse command line arguments
type RootCommand struct {
	Version                           func()                            `short:"v" long:"version"  description:"Print the version of tool and exit"`
	DeployFactProviderRegistryCommand DeployFactProviderRegistryCommand `command:"deploy" description:"Deploy fact provider registry"`
	SetFactProviderInfoCommand        SetFactProviderInfoCommand        `command:"set"    description:"Set fact provider information"`
	GetFactProviderInfoCommand        GetFactProviderInfoCommand        `command:"get"    description:"Get fact provider information"`
	DeleteFactProviderInfoCommand     DeleteFactProviderInfoCommand     `command:"delete" description:"Delete fact provider information"`
}

// Root is a prepared command to be used in command line arguments parsing
var Root RootCommand

func init() {
	Root.Version = func() {
		cmd.PrintVersion()
		os.Exit(0)
	}
}
