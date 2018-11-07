// +build js,wasm

package main

import (
	"gitlab.com/monetha/protocol-go-sdk/cmd/protoscan/web/app"
	"gitlab.com/monetha/protocol-go-sdk/cmd/protoscan/web/logging"
)

func main() {
	app.Run("box", logging.Fun)
}
