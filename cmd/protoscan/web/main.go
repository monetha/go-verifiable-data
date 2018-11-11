// +build js,wasm

package main

import (
	"html/template"
	"strings"
	"syscall/js"

	"gitlab.com/monetha/protocol-go-sdk/cmd/protoscan/web/app"
	"gitlab.com/monetha/protocol-go-sdk/cmd/protoscan/web/dom"
	"gitlab.com/monetha/protocol-go-sdk/cmd/protoscan/web/logging"
)

var htmlMarkupTmpl = template.Must(template.New("htmlMarkup").Parse(
	`<nav id="main-nav" class="navbar navbar-expand-md navbar-dark fixed-top bg-dark">
    <a class="navbar-brand" href="#">Monetha</a>
    <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarCollapse"
            aria-controls="navbarCollapse" aria-expanded="false" aria-label="Toggle navigation">
        <span class="navbar-toggler-icon"></span>
    </button>
    <div class="collapse navbar-collapse" id="navbarCollapse">
        <ul class="navbar-nav mr-auto">
            <li class="nav-item">
                <a class="nav-link" href="#passport-list">Passport list</a>
            </li>
            <li class="nav-item">
                <a class="nav-link" href="#passport-changes">Passport changes</a>
            </li>
        </ul>
        <form class="form-inline mt-3 mt-md-0" id="navBarBackendURLForm">
            <div class="input-group">
                <div class="input-group-prepend">
                    <span class="input-group-text">Backend URL</span>
                </div>
                <input class="form-control mr-sm-3" type="text" placeholder="Backend URL"
                       aria-label="Backend URL"
                       value="{{.BackendURL}}" id="backendURLInp">
            </div>
        </form>
    </div>
</nav>

<div id="passport-list" class="container-fluid content-section">
    <div class="row justify-content-center">
        <div class="col-auto"><h1>Passport list</h1></div>
    </div>
    <div class="row">
        <div class="col-3">
            <form>
                <div class="form-group">
                    <label for="passListPassportFactoryAddressInp">Passport factory address</label>
                    <input type="text" class="form-control" placeholder="Passport factory address"
                           value="{{.PassportFactoryAddress}}" id="passListPassportFactoryAddressInp">
                </div>
                <button class="btn btn-primary btn-block" id="getPassportListBtn">Get passport list &raquo;</button>
            </form>
        </div>
        <div class="col-9">
            <div class="row">
                <div class="col-12" id="passportListOutput">
                </div>
            </div>
        </div>
    </div>
</div>

<div id="passport-changes" class="container-fluid content-section">
    <div class="row justify-content-center">
        <div class="col-auto"><h1>Passport changes</h1></div>
    </div>
    <div class="row">
        <div class="col-3">
            <form>
                <div class="form-group">
                    <label for="passChangesPassportAddressInp">Passport address</label>
                    <input type="text" class="form-control" placeholder="Passport address"
                           id="passChangesPassportAddressInp">
                </div>
                <button class="btn btn-primary btn-block" id="getPassportChangesBtn">Get passport changes &raquo;</button>
            </form>
        </div>
        <div class="col-9">
            <div class="row">
                <div class="col-12" id="passportChangesOutput">
                </div>
            </div>
        </div>
    </div>
</div>`))

func main() {
	log := logging.Fun
	log("Initializing protocol scanner...")

	done := make(chan struct{})

	var sb strings.Builder
	if err := htmlMarkupTmpl.Execute(&sb, struct {
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
		SetupOnClickGetPassportList().
		SetupOnClickGetPassportChanges()

	log("Protocol scanner is initialized.")

	<-done
	log("Shutting down protocol scanner...")

	_ = a.Close()

	log("Protocol scanner is shut down.")
}
