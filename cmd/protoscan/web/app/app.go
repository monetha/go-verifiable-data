// +build js,wasm

package app

import (
	"io"
	"strconv"
	"syscall/js"

	"github.com/ethereum/go-ethereum/common"
	"gitlab.com/monetha/protocol-go-sdk/cmd/protoscan/web/dom"
	"gitlab.com/monetha/protocol-go-sdk/facts"
	"gitlab.com/monetha/protocol-go-sdk/log"
	"gitlab.com/monetha/protocol-go-sdk/passfactory"
)

type App struct {
	Log log.Fun

	BackendURLInput dom.Inp

	PassFactoryAddressInput      dom.Inp
	GetPassportListButton        dom.Btn
	PassportListOutputDiv        dom.Elt
	getPassportListRequestCloser io.Closer
	onGetPassportListClickCb     js.Callback

	PassAddressInput                dom.Inp
	GetPassportChangesButton        dom.Btn
	PassportChangesOutputDiv        dom.Elt
	getPassportChangesRequestCloser io.Closer
	onGetPassportChangesClickCb     js.Callback
}

func (a *App) SetupOnClickGetPassportList() *App {
	a.onGetPassportListClickCb = a.GetPassportListButton.OnClick(js.PreventDefault, func(args js.Value) {
		a.cancelGetPassportListRequest()

		passportFactoryAddressStr := a.PassFactoryAddressInput.Value()
		passportFactoryAddress := common.HexToAddress(passportFactoryAddressStr)

		backendURL := a.BackendURLInput.Value()

		resultStatusDiv := dom.Div().
			WithClass("col-12 alert alert-primary").
			WithRole("alert").
			WithChildren(dom.Text("Getting passport list..."))

		resultTable := dom.Table().
			WithClass("table table-hover table-striped").
			WithHeader(
				dom.Text("Passport address"),
				dom.Text("First owner address"),
				dom.Text("Block number"),
				dom.Text("Transaction hash"),
			).
			WithHeaderClass("thead-light")

		resultDiv := dom.Div().WithChildren(
			dom.Div().WithClass("row").WithChildren(
				resultStatusDiv,
			),
			dom.Div().WithClass("row").WithChildren(
				dom.Div().WithClass("col-12 table-responsive").WithChildren(resultTable),
			),
		)

		a.PassportListOutputDiv.RemoveAllChildren()
		a.PassportListOutputDiv.AppendChild(resultDiv)

		a.Log("Getting passport list from passport factory...", "backend_url", backendURL, "passport_factory_address", passportFactoryAddress.Hex())

		a.getPassportListRequestCloser = (&passportListGetter{
			Log:        a.Log,
			BackendURL: backendURL,
		}).GetPassportListAsync(passportFactoryAddress, &passportListObserver{
			OnErrorFun: func(err error) {
				a.Log("passport filtering error", "error", err.Error())
				resultStatusDiv.RemoveAllChildren()
				resultStatusDiv.
					WithClass("col-12 alert alert-danger").
					AppendChild(dom.Text(err.Error()))
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

func (a *App) SetupOnClickGetPassportChanges() *App {
	a.onGetPassportChangesClickCb = a.GetPassportChangesButton.OnClick(js.PreventDefault, func(args js.Value) {
		a.cancelGetPassportChangesRequest()

		passportAddressStr := a.PassAddressInput.Value()
		passportAddress := common.HexToAddress(passportAddressStr)

		backendURL := a.BackendURLInput.Value()

		resultStatusDiv := dom.Div().
			WithClass("col-12 alert alert-primary").
			WithRole("alert").
			WithChildren(dom.Text("Getting passport changes..."))

		resultTable := dom.Table().
			WithClass("table table-hover table-striped").
			WithHeader(
				dom.Text("Fact provider address"),
				dom.Text("Key"),
				dom.Text("Data type"),
				dom.Text("Change type"),
				dom.Text("Block number"),
				dom.Text("Transaction hash"),
			).
			WithHeaderClass("thead-light")

		resultDiv := dom.Div().WithChildren(
			dom.Div().WithClass("row").WithChildren(
				resultStatusDiv,
			),
			dom.Div().WithClass("row").WithChildren(
				dom.Div().WithClass("col-12 table-responsive").WithChildren(resultTable),
			),
		)

		a.PassportChangesOutputDiv.RemoveAllChildren()
		a.PassportChangesOutputDiv.AppendChild(resultDiv)

		a.Log("Getting passport changes...", "backend_url", backendURL, "passport_address", passportAddress.Hex())

		a.getPassportChangesRequestCloser = (&passportChangesGetter{
			Log:        a.Log,
			BackendURL: backendURL,
		}).GetPassportChangesAsync(passportAddress, &passportChangesObserver{
			OnErrorFun: func(err error) {
				a.Log("passport changes error", "error", err.Error())
				resultStatusDiv.RemoveAllChildren()
				resultStatusDiv.
					WithClass("col-12 alert alert-danger").
					AppendChild(dom.Text(err.Error()))
			},
			OnCompletedFun: func() {
				a.Log("passport changes completed")
				resultStatusDiv.Remove()
			},
			OnNextFun: func(ch *facts.Change) {
				a.Log("next change",
					"fact_provider", ch.FactProvider.Hex(),
					"key", string(ch.Key[:]),
					"data_type", ch.DataType,
					"change_type", ch.ChangeType,
					"block_number", ch.Raw.BlockNumber, "tx_hash", ch.Raw.TxHash.Hex())
				resultTable.AppendRow(
					dom.Text(ch.FactProvider.Hex()),
					dom.Text(string(ch.Key[:])),
					dom.Text(ch.DataType.String()),
					dom.Text(ch.ChangeType.String()),
					dom.Text(strconv.FormatUint(ch.Raw.BlockNumber, 10)),
					dom.Text(ch.Raw.TxHash.Hex()),
				)
			},
		})
	})

	return a
}

func (a *App) cancelGetPassportListRequest() {
	if c := a.getPassportListRequestCloser; c != nil {
		// cancel previous request
		_ = c.Close()
	}
}

func (a *App) cancelGetPassportChangesRequest() {
	if c := a.getPassportChangesRequestCloser; c != nil {
		// cancel previous request
		_ = c.Close()
	}
}

func (a *App) Close() error {
	a.cancelGetPassportListRequest()
	a.onGetPassportListClickCb.Release()

	a.cancelGetPassportChangesRequest()
	a.onGetPassportChangesClickCb.Release()

	return nil
}
