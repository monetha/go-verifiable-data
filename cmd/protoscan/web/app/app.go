// +build js,wasm

package app

import (
	"context"
	"fmt"
	"github.com/dennwc/dom/js"
	"io"
	"strconv"

	"github.com/dennwc/dom"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/monetha/protocol-go-sdk/cmd/protoscan/web/rx"
	"gitlab.com/monetha/protocol-go-sdk/eth"
	"gitlab.com/monetha/protocol-go-sdk/log"
	"gitlab.com/monetha/protocol-go-sdk/passfactory"
)

type box dom.Element

func (b *box) appendInput(label string, typ string, value string) *dom.Input {
	inp := dom.Doc.NewInput("text")
	inp.JSRef().Set("placeholder", label)
	inp.SetValue(value)
	b.AppendChild(inp)

	return inp
}

func (b *box) appendButton(s string) *dom.Button {
	bt := dom.Doc.NewButton(s)
	b.AppendChild(bt)
	return bt
}

func (b *box) appendDiv() *box {
	d := dom.Doc.CreateElement("div")
	b.AppendChild(d)
	return (*box)(d)
}

func (b *box) appendTable() *table {
	t := dom.Doc.CreateElement("table")
	b.AppendChild(t)
	return (*table)(t)
}

func (b *box) appendTextNode(s string) *dom.Element {
	n := createTextNode(s)
	b.AppendChild(n)
	return n
}

type table dom.Element

func (t *table) appendHeader(els ...*dom.Element) {
	header := t.JSRef().Call("createTHead")
	newRow := header.Call("insertRow", -1)
	for _, el := range els {
		newCell := dom.AsElement(js.ValueOf(newRow.Call("insertCell", -1)))
		newCell.AppendChild(el)
	}
}


func (t *table) appendRow(els ...*dom.Element) {
	newRow := t.JSRef().Call("insertRow", -1)
	for _, el := range els {
		newCell := dom.AsElement(js.ValueOf(newRow.Call("insertCell", -1)))
		newCell.AppendChild(el)
	}
}

func createTextNode(s string) *dom.Element {
	return dom.AsElement(js.ValueOf(dom.Doc.JSRef().Call("createTextNode", s)))
}

func Run(elementId string, lf log.Fun) {
	lf("Initializing protocol scanner...")

	done := make(chan struct{})

	box := (*box)(dom.Doc.GetElementById(elementId))

	backendURLInput := box.appendInput("Backend URL", "text", "https://ropsten.infura.io")
	passFactoryAddressInput := box.appendInput("Passport factory address", "text", "0x87b7Ec2602Da6C9e4D563d788e1e29C064A364a2")

	passFactoryAddressInput.AddEventListener("keyup", dom.EventHandler(func(e dom.Event) {
		lf("Factory address changed", "passport_factory_address", e.Target().JSRef().Get("value"))
	}))

	getPassportListButton := box.appendButton("Get passport list")
	resultDiv := box.appendDiv()

	getPassportListButton.OnClick(dom.EventHandler(func(e dom.Event) {
		passportFactoryAddressStr := passFactoryAddressInput.Value()
		passportFactoryAddress := common.HexToAddress(passportFactoryAddressStr)

		backendURL := backendURLInput.Value()

		resultDiv.Remove()
		resultDiv = box.appendDiv()
		resultStatus := resultDiv.appendTextNode("Filtering passports...")
		resultTable := resultDiv.appendTable()
		resultTable.appendHeader(
			createTextNode("Passport address"),
			createTextNode("First owner address"),
			createTextNode("Block number"),
			createTextNode("Transaction hash"),
			)

		lf("Getting passport list from passport factory...", "backend_url", backendURL, "passport_factory_address", passportFactoryAddress.Hex())

		(&passportFilterer{
			Log:        lf,
			BackendURL: backendURL,
		}).FilterAsync(passportFactoryAddress, &passportObserverFun{
			OnErrorFun: func(err error) {
				lf("passport filtering error", "error", err.Error())
				resultStatus.SetInnerHTML(err.Error())
			},
			OnCompletedFun: func() {
				lf("passport filtering completed")
				resultStatus.Remove()
			},
			OnNextFun: func(p *passfactory.Passport) {
				lf("next passport", "contract_address", p.ContractAddress.Hex(), "first_owner_address", p.FirstOwner.Hex(), "block_number", p.Raw.BlockNumber, "tx_hash", p.Raw.TxHash.Hex())
				resultTable.appendRow(
					createTextNode(p.ContractAddress.Hex()),
					createTextNode(p.FirstOwner.Hex()),
					createTextNode(strconv.FormatUint(p.Raw.BlockNumber, 10)),
					createTextNode(p.Raw.TxHash.Hex()),
					)
			},
		})
	}))

	lf("Protocol scanner is initialized.")

	<-done
	lf("Shutting down protocol scanner...")

	lf("Protocol scanner is shut down.")
}

type passportObserverFun struct {
	OnErrorFun     func(err error)
	OnCompletedFun func()
	OnNextFun      func(passport *passfactory.Passport)
}

func (o *passportObserverFun) OnError(err error) {
	o.OnErrorFun(err)
}

func (o *passportObserverFun) OnCompleted() {
	o.OnCompletedFun()
}

func (o *passportObserverFun) OnNext(passport *passfactory.Passport) {
	o.OnNextFun(passport)
}

type passportObserver interface {
	rx.Observer
	OnNext(passport *passfactory.Passport)
}

type passportFilterer struct {
	Context    context.Context
	Log        log.Fun
	BackendURL string
}

func (f *passportFilterer) FilterAsync(passportFactoryAddress common.Address, o passportObserver) io.Closer {
	backendURL := f.BackendURL
	lf := f.Log
	onNext := o.OnNext

	return rx.RunAsyncObserver(f.Context, o, func(ctx context.Context) (err error) {
		client, err := ethclient.Dial(backendURL)
		if err != nil {
			return fmt.Errorf("ethclient.Dial: %v", err)
		}
		defer client.Close()

		e := eth.New(client, lf)

		pfr := passfactory.NewReader(e)
		filterOpts := &passfactory.PassportFilterOpts{
			Context: ctx,
		}

		var it *passfactory.PassportIterator
		it, err = pfr.FilterPassports(filterOpts, passportFactoryAddress)
		if err != nil {
			err = fmt.Errorf("Reader.FilterPassports: %v", err)
			return
		}
		defer func() {
			if cErr := it.Close(); err == nil && cErr != nil {
				err = cErr
			}
		}()

		for it.Next() {
			if err = it.Error(); err != nil {
				err = fmt.Errorf("PassportIterator.Next: %v", err)
				return
			}

			onNext(it.Passport)
		}

		return nil
	})
}
