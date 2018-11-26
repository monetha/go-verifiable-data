// +build js,wasm

package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
	"syscall/js"

	"gitlab.com/monetha/reputation-go-sdk/cmd/passport-scanner/web/app"
	"gitlab.com/monetha/reputation-go-sdk/cmd/passport-scanner/web/dom"
	"gitlab.com/monetha/reputation-go-sdk/cmd/passport-scanner/web/logging"
)

func main() {
	log := logging.Fun
	log("Initializing passport scanner...")

	htmlMarkupTemplate := getTMPL("main.template")

	done := make(chan struct{})

	var sb strings.Builder
	if err := htmlMarkupTemplate.Execute(&sb, struct {
		BackendURL             string
		PassportFactoryAddress string
	}{
		BackendURL:             "https://ropsten.infura.io",
		PassportFactoryAddress: "0x87b7Ec2602Da6C9e4D563d788e1e29C064A364a2",
	}); err != nil {
		panic(err)
	}

	dom.Document.
		GetElementById("box").
		SetInnerHTML(sb.String())

	// $('body').scrollspy({ target: '#main-nav' })
	opts := js.Global().Get("Object").Invoke() // New JS Object
	opts.Set("target", "#main-nav")
	js.Global().
		Get("$").Invoke("body").
		Call("scrollspy", opts)

	a := (&app.App{
		Log:             log,
		BackendURLInput: dom.Document.GetElementById("backendURLInp").AsInput(),

		PassListPassFactoryAddressInput: dom.Document.GetElementById("passListPassportFactoryAddressInp").AsInput(),
		GetPassportListButton:           dom.Document.GetElementById("getPassportListBtn").AsButton(),
		PassportListOutputDiv:           dom.Document.GetElementById("passportListOutput"),

		PassChangesPassAddressInput: dom.Document.GetElementById("passChangesPassportAddressInp").AsInput(),
		GetPassportChangesButton:    dom.Document.GetElementById("getPassportChangesBtn").AsButton(),
		PassportChangesOutputDiv:    dom.Document.GetElementById("passportChangesOutput"),
	}).
		RegisterCallBacks()

	log("Protocol scanner is initialized.")

	<-done
	log("Shutting down passport scanner...")

	_ = a.Close()

	log("Protocol scanner is shut down.")
}

// getTMPL parses template from assets/tmpl folder
func getTMPL(name string) *template.Template {
	resp, err := http.Get("tmpl/" + name)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return template.Must(template.New(name).Parse(string(b)))
}
