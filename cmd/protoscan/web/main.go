// +build js,wasm

package main

import (
	"gitlab.com/monetha/protocol-go-sdk/cmd/protoscan/web/app"
	"gitlab.com/monetha/protocol-go-sdk/cmd/protoscan/web/dom"
	"gitlab.com/monetha/protocol-go-sdk/cmd/protoscan/web/logging"
	"gitlab.com/monetha/protocol-go-sdk/log"
)

func main() {
	run("box", logging.Fun)
}

func run(elementId string, lf log.Fun) {
	lf("Initializing protocol scanner...")

	done := make(chan struct{})

	const (
		backendURLStr             = "Backend URL"
		passportFactoryAddressStr = "Passport factory address"
	)

	backendURLInput := dom.TextInput().WithClass("form-control").WithPlaceholder(backendURLStr).WithValue("https://ropsten.infura.io")
	passFactoryAddressInput := dom.TextInput().WithClass("form-control").WithPlaceholder(passportFactoryAddressStr).WithValue("0x87b7Ec2602Da6C9e4D563d788e1e29C064A364a2")
	getPassportListButton := dom.Button("Get passport list").WithClass("btn btn-primary btn-block")
	passportListOutputDiv := dom.Div()

	dom.Document.
		GetElementById(elementId).
		AppendChild(
			dom.Div().WithClass("row").WithChildren(
				dom.Div().WithClass("col-3").WithChildren(
					dom.Form().WithChildren(
						dom.Div().WithClass("form-group").WithChildren(
							dom.Label(backendURLStr),
							backendURLInput,
						),
						dom.Div().WithClass("form-group").WithChildren(
							dom.Label(passportFactoryAddressStr),
							passFactoryAddressInput,
						),
						getPassportListButton,
					),
				),
				dom.Div().WithClass("col-9").WithChildren(
					dom.Div().WithClass("row").WithChildren(
						passportListOutputDiv.WithClass("col-12"),
					),
				),
			),
		)

	a := (&app.App{
		Log:                     lf,
		BackendURLInput:         backendURLInput,
		PassFactoryAddressInput: passFactoryAddressInput,
		GetPassportListButton:   getPassportListButton,
		PassportListOutputDiv:   passportListOutputDiv,
	}).
		SetupOnClickGetPassportList()

	lf("Protocol scanner is initialized.")

	<-done
	lf("Shutting down protocol scanner...")

	_ = a.Close()

	lf("Protocol scanner is shut down.")
}
