// +build js,wasm

package main

import (
	"syscall/js"
)

func main() {
	js.Global().Get("console").Call("log", "Hello world Go/wasm!")
}
