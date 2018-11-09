// +build js,wasm

package app

import (
	"context"
	"fmt"
	"io"
	"strconv"
	"syscall/js"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/monetha/protocol-go-sdk/cmd/protoscan/web/dom"
	"gitlab.com/monetha/protocol-go-sdk/cmd/protoscan/web/rx"
	"gitlab.com/monetha/protocol-go-sdk/eth"
	"gitlab.com/monetha/protocol-go-sdk/log"
	"gitlab.com/monetha/protocol-go-sdk/passfactory"
)

type App struct {
	Log log.Fun

	BackendURLInput         dom.Inp
	PassFactoryAddressInput dom.Inp
	GetPassportListButton   dom.Btn
	PassportListOutputDiv   dom.Elt

	getPassportListRequestCloser io.Closer
	onGetPassportListClickCb     js.Callback
}

func (a *App) SetupOnClickGetPassportList() *App {
	a.onGetPassportListClickCb = a.GetPassportListButton.OnClick(js.PreventDefault, func(args js.Value) {
		if prevGetPassporAsyncCloser := a.getPassportListRequestCloser; prevGetPassporAsyncCloser != nil {
			// cancel previous request
			_ = prevGetPassporAsyncCloser.Close()
		}

		passportFactoryAddressStr := a.PassFactoryAddressInput.Value()
		passportFactoryAddress := common.HexToAddress(passportFactoryAddressStr)

		backendURL := a.BackendURLInput.Value()

		resultStatusDiv := dom.Div().WithChildren(dom.Text("Filtering passports..."))
		resultTable := dom.Table().
			WithClass("table table-hover").
			WithHeader(
				dom.Text("Passport address"),
				dom.Text("First owner address"),
				dom.Text("Block number"),
				dom.Text("Transaction hash"),
			)
		resultDiv := dom.Div().WithChildren(
			dom.Div().WithClass("row").WithChildren(
				resultStatusDiv.WithClass("col-12"),
			),
			dom.Div().WithClass("row").WithChildren(
				dom.Div().WithClass("col-12").WithChildren(resultTable),
			),
		)

		a.PassportListOutputDiv.RemoveAllChildren()
		a.PassportListOutputDiv.AppendChild(resultDiv)

		a.Log("Getting passport list from passport factory...", "backend_url", backendURL, "passport_factory_address", passportFactoryAddress.Hex())

		a.getPassportListRequestCloser = (&passportFilterer{
			Log:        a.Log,
			BackendURL: backendURL,
		}).FilterAsync(passportFactoryAddress, &passportObserverFun{
			OnErrorFun: func(err error) {
				a.Log("passport filtering error", "error", err.Error())
				resultStatusDiv.RemoveAllChildren()
				resultStatusDiv.AppendChild(dom.Text(err.Error()))
			},
			OnCompletedFun: func() {
				a.Log("passport filtering completed")
				resultStatusDiv.Remove()
			},
			OnNextFun: func(p *passfactory.Passport) {
				a.Log("next passport", "contract_address", p.ContractAddress.Hex(), "first_owner_address", p.FirstOwner.Hex(), "block_number", p.Raw.BlockNumber, "tx_hash", p.Raw.TxHash.Hex())
				resultTable.AppendRow(
					dom.Text(p.ContractAddress.Hex()),
					dom.Text(p.FirstOwner.Hex()),
					dom.Text(strconv.FormatUint(p.Raw.BlockNumber, 10)),
					dom.Text(p.Raw.TxHash.Hex()),
				)
			},
		})
	})

	return a
}

func (a *App) Close() error {
	a.onGetPassportListClickCb.Release()

	return nil
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
