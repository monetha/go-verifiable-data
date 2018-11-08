// +build js,wasm

package main

import (
	"context"
	"fmt"
	"io"
	"strconv"
	"syscall/js"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/monetha/protocol-go-sdk/cmd/protoscan/web/dom"
	"gitlab.com/monetha/protocol-go-sdk/cmd/protoscan/web/logging"
	"gitlab.com/monetha/protocol-go-sdk/cmd/protoscan/web/rx"
	"gitlab.com/monetha/protocol-go-sdk/eth"
	"gitlab.com/monetha/protocol-go-sdk/log"
	"gitlab.com/monetha/protocol-go-sdk/passfactory"
)

func main() {
	run("box", logging.Fun)
}

func run(elementId string, lf log.Fun) {
	lf("Initializing protocol scanner...")

	done := make(chan struct{})

	backendURLInput := dom.TextInput().WithClass("form-control" ).WithPlaceholder("Backend URL").WithValue("https://ropsten.infura.io")
	passFactoryAddressInput := dom.TextInput().WithClass("form-control" ).WithPlaceholder("Passport factory address").WithValue("0x87b7Ec2602Da6C9e4D563d788e1e29C064A364a2")
	getPassportListButton := dom.Button("Get passport list").WithClass("btn btn-primary btn-block")
	resultDiv := dom.Div()

	dom.Document.
		GetElementById(elementId).
		AppendChild(
			dom.Div().WithClass("row").WithChildren(
				dom.Div().WithClass("col-3").WithChildren(
					dom.Form().WithChildren(
						dom.Div().WithClass("form-group").WithChildren(
							dom.Label("Backend URL"),
							backendURLInput,
						),
						dom.Div().WithClass("form-group").WithChildren(
							dom.Label("Passport factory address"),
							passFactoryAddressInput,
						),
						getPassportListButton,
					),
				),
				dom.Div().WithClass("col-9").WithChildren(
					resultDiv,
				),
			))

	getPassportCallback := getPassportListButton.OnClick(js.PreventDefault, func(args js.Value) {
		passportFactoryAddressStr := passFactoryAddressInput.Value()
		passportFactoryAddress := common.HexToAddress(passportFactoryAddressStr)

		backendURL := backendURLInput.Value()

		resultDiv.RemoveAllChildren()

		resultStatus := dom.Text("Filtering passports...")
		resultTable := dom.Table().
			WithClass("table table-hover").
			WithHeader(
				dom.Text("Passport address"),
				dom.Text("First owner address"),
				dom.Text("Block number"),
				dom.Text("Transaction hash"),
			)

		resultDiv.WithChildren(
			resultStatus,
			resultTable,
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
				resultDiv.RemoveChild(resultStatus)
			},
			OnNextFun: func(p *passfactory.Passport) {
				lf("next passport", "contract_address", p.ContractAddress.Hex(), "first_owner_address", p.FirstOwner.Hex(), "block_number", p.Raw.BlockNumber, "tx_hash", p.Raw.TxHash.Hex())
				resultTable.AppendRow(
					dom.Text(p.ContractAddress.Hex()),
					dom.Text(p.FirstOwner.Hex()),
					dom.Text(strconv.FormatUint(p.Raw.BlockNumber, 10)),
					dom.Text(p.Raw.TxHash.Hex()),
				)
			},
		})
	})
	defer getPassportCallback.Release()

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
