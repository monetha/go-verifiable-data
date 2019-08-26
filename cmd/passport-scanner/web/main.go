// +build js,wasm

package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
	"syscall/js"

	"github.com/monetha/go-verifiable-data/cmd/passport-scanner/web/app"
	"github.com/monetha/go-verifiable-data/cmd/passport-scanner/web/dom"
	"github.com/monetha/go-verifiable-data/cmd/passport-scanner/web/logging"
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
		StartFromBlock         int
	}{
		BackendURL:             "https://ropsten.infura.io",
		PassportFactoryAddress: "0x35Cb95Db8E6d56D1CF8D5877EB13e9EE74e457F2",
		StartFromBlock:         5233000,
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
		PassListStartFromBlockInp:       dom.Document.GetElementById("passListStartFromBlockInp").AsInput(),
		GetPassportListButton:           dom.Document.GetElementById("getPassportListBtn").AsButton(),
		PassportListOutputDiv:           dom.Document.GetElementById("passportListOutput"),

		PassChangesPassAddressInput:  dom.Document.GetElementById("passChangesPassportAddressInp").AsInput(),
		PassChangesStartFromBlockInp: dom.Document.GetElementById("passChangesStartFromBlockInp").AsInput(),
		GetPassportChangesButton:     dom.Document.GetElementById("getPassportChangesBtn").AsButton(),
		PassportChangesOutputDiv:     dom.Document.GetElementById("passportChangesOutput"),
	}).
		RegisterCallBacks()

	log("Passport scanner is initialized.")

	<-done
	log("Shutting down passport scanner...")

	_ = a.Close()

	log("Passport scanner is shut down.")
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
