package main

import (
	"os"

	"github.com/jessevdk/go-flags"
	"github.com/monetha/reputation-go-sdk/cmd/privatedata-exchange/commands"
)

func main() {
	p := flags.NewParser(&commands.Exchange, flags.Default)

	_, err := p.Parse()
	handleError(err)
}

func handleError(err error) {
	if err != nil {
		if flagsErr, ok := err.(*flags.Error); ok {
			if flagsErr.Type == flags.ErrHelp {
				os.Exit(0)
			}
		}

		os.Exit(1)
	}
}
