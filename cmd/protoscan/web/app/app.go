// +build js,wasm

package app

import (
	"io"
	"strconv"
	"syscall/js"

	"github.com/ethereum/go-ethereum/common"
	"gitlab.com/monetha/protocol-go-sdk/cmd/protoscan/web/dom"
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
		a.cancelGetPassportListRequest()

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

		a.getPassportListRequestCloser = (&passportListGetter{
			Log:        a.Log,
			BackendURL: backendURL,
		}).GetPassportListAsync(passportFactoryAddress, &passportListObserver{
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

func (a *App) cancelGetPassportListRequest() {
	if prevGetPassportAsyncCloser := a.getPassportListRequestCloser; prevGetPassportAsyncCloser != nil {
		// cancel previous request
		_ = prevGetPassportAsyncCloser.Close()
	}
}

func (a *App) Close() error {
	a.cancelGetPassportListRequest()
	a.onGetPassportListClickCb.Release()

	return nil
}
