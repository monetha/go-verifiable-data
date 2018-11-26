// +build js,wasm

package app

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"html"
	"io"
	"strconv"
	"strings"
	"syscall/js"

	"github.com/ethereum/go-ethereum/common"
	"gitlab.com/monetha/reputation-go-sdk/cmd/passport-scanner/web/dom"
	"gitlab.com/monetha/reputation-go-sdk/facts"
	"gitlab.com/monetha/reputation-go-sdk/log"
	"gitlab.com/monetha/reputation-go-sdk/passfactory"
	"gitlab.com/monetha/reputation-go-sdk/types/change"
	"gitlab.com/monetha/reputation-go-sdk/types/data"
)

type App struct {
	Log log.Fun

	BackendURLInput dom.Inp

	PassListPassFactoryAddressInput dom.Inp
	GetPassportListButton           dom.Btn
	PassportListOutputDiv           dom.Elt
	getPassportListRequestCloser    io.Closer
	onGetPassportListClickCb        js.Callback

	PassChangesPassAddressInput     dom.Inp
	GetPassportChangesButton        dom.Btn
	PassportChangesOutputDiv        dom.Elt
	getPassportChangesRequestCloser io.Closer
	onGetPassportChangesClickCb     js.Callback

	readHistoryValueCb js.Callback
}

func (a *App) RegisterCallBacks() *App {
	a.setupOnClickGetPassportList()
	a.setupOnClickGetPassportChanges()
	a.setupOnReadHistoryValue()

	return a
}

func (a *App) setupOnClickGetPassportList() *App {
	a.onGetPassportListClickCb = a.GetPassportListButton.OnClick(js.PreventDefault, func(args js.Value) {
		a.cancelGetPassportListRequest()

		passportFactoryAddressStr := a.PassListPassFactoryAddressInput.Value()
		passportFactoryAddress := common.HexToAddress(passportFactoryAddressStr)

		backendURL := a.BackendURLInput.Value()
		networkType := getNetworkType(backendURL)

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
				contractAddress := p.ContractAddress.Hex()
				firstOwnerAddress := p.FirstOwner.Hex()
				blockNumber := strconv.FormatUint(p.Raw.BlockNumber, 10)
				txHash := p.Raw.TxHash.Hex()

				a.Log("next passport", "contract_address", contractAddress, "first_owner_address", firstOwnerAddress, "block_number", blockNumber, "tx_hash", txHash)
				resultTable.AppendRow(
					networkType.CreateAddressLink(contractAddress),
					networkType.CreateAddressLink(firstOwnerAddress),
					networkType.CreateBlockLink(blockNumber),
					networkType.CreateTxLink(txHash),
				)
			},
		})
	})

	return a
}

func (a *App) setupOnClickGetPassportChanges() *App {
	a.onGetPassportChangesClickCb = a.GetPassportChangesButton.OnClick(js.PreventDefault, func(args js.Value) {
		a.cancelGetPassportChangesRequest()

		passportAddressStr := a.PassChangesPassAddressInput.Value()
		passportAddress := common.HexToAddress(passportAddressStr)

		backendURL := a.BackendURLInput.Value()
		networkType := getNetworkType(backendURL)

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
				dom.Text("Value"),
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
				factProviderAddress := ch.FactProvider.Hex()
				key := keyToString(ch.Key)
				dataType := ch.DataType.String()
				changeType := ch.ChangeType.String()
				blockNumber := strconv.FormatUint(ch.Raw.BlockNumber, 10)
				txHash := ch.Raw.TxHash.Hex()

				a.Log("next change",
					"fact_provider", factProviderAddress,
					"key", key,
					"data_type", dataType,
					"change_type", changeType,
					"block_number", blockNumber, "tx_hash", txHash)

				var valueElt dom.Node
				if ch.ChangeType == change.Updated {
					valueElt = dom.Anchor("Download").
						WithClass("btn btn-secondary").
						WithRole("button").
						WithAttribute("href", "#").
						WithAttribute("onClick", getReadHistoryValueCallbackText(backendURL, passportAddressStr, ch.DataType, txHash))
				} else {
					valueElt = dom.Text("–")
				}

				resultTable.AppendRow(
					networkType.CreateAddressLink(factProviderAddress),
					dom.Text(key),
					dom.Text(dataType),
					dom.Text(changeType),
					valueElt,
					networkType.CreateBlockLink(blockNumber),
					networkType.CreateTxLink(txHash),
				)
			},
		})
	})

	return a
}

func getReadHistoryValueCallbackText(backendURL string, passportAddress string, dt data.Type, txHash string) string {
	return fmt.Sprintf("readHistoryValue(this, '%v', '%v', %d, '%v'); return false;", html.EscapeString(backendURL), passportAddress, dt, txHash)
}

func getDownloadValueCallbackText(filename string, content []byte) string {
	return fmt.Sprintf("return Export.createDownloadLink(this, '%v', '%v');", html.EscapeString(filename), base64.StdEncoding.EncodeToString(content))
}

func keyToString(key [32]byte) string {
	return string(bytes.Trim(key[:], "\x00"))
}

func ipfsHashToUrl(hash string) string {
	return fmt.Sprintf("https://ipfs.io/ipfs/%s", html.EscapeString(hash))
}

func (a *App) setupOnReadHistoryValue() {
	a.readHistoryValueCb = js.NewCallback(func(args []js.Value) {
		if len(args) != 5 {
			a.Log("readHistoryValue: unexpected arguments count", "args", args)
			return
		}

		btn := dom.NodeBase{Value: args[0]}.AsElement().WithClassAdded("disabled").AsAnchor()
		backendURL := args[1].String()
		passportAddress := common.HexToAddress(args[2].String())
		dataType := data.Type(args[3].Int())
		txHash := common.HexToHash(args[4].String())

		a.Log("reading history value...", "backend_url", backendURL, "passport_addres", passportAddress.Hex(), "data_type", dataType.String(), "tx_hash", txHash.Hex())
		(&passportChangesGetter{
			Log:        a.Log,
			BackendURL: backendURL,
		}).GetHistoryItemAsync(passportAddress, dataType, txHash, &historyItemObserver{
			OnErrorFun: func(err error) {
				a.Log("reading history error",
					"backend_url", backendURL,
					"passport_address", passportAddress.Hex(),
					"data_type", dataType.String(),
					"tx_hash", txHash.Hex(),
					"error", err)
				dom.Alert(err.Error())
				btn.WithClassRemoved("disabled")
			},
			OnCompletedFun: func() {
				a.Log("reading history completed.",
					"backend_url", backendURL,
					"passport_address", passportAddress.Hex(),
					"data_type", dataType.String(),
					"tx_hash", txHash.Hex())
				btn.WithClassRemoved("disabled")
			},
			OnNextFun: func(hi *historyItem) {
				a.Log("reading history value",
					"backend_url", backendURL,
					"passport_address", passportAddress.Hex(),
					"data_type", dataType.String(),
					"tx_hash", txHash.Hex(),
					"fact_provider", hi.FactProvider.Hex(),
					"key", string(hi.Key[:]))

				switch dataType {
				case data.IPFS:
					hash := string(hi.Value)
					btn.WithClass("btn btn-success").
						WithAttribute("href", ipfsHashToUrl(hash)).
						WithAttribute("target", "_blank").
						WithAttributeRemoved("onClick")
				default:
					filename := txHash.Hex() + "_" + keyToString(hi.Key)
					btn.WithClass("btn btn-success").
						WithAttribute("download", filename).
						WithAttribute("onClick", getDownloadValueCallbackText(filename, hi.Value))
				}
				btn.Call("click")
			},
		})
	})
	js.Global().Set("readHistoryValue", a.readHistoryValueCb)
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

	a.readHistoryValueCb.Release()

	return nil
}

type network interface {
	CreateAddressLink(address string) dom.Elt
	CreateBlockLink(block string) dom.Elt
	CreateTxLink(txHash string) dom.Elt
}

type unknownNetwork struct{}

func (unknownNetwork) CreateAddressLink(address string) dom.Elt { return monospaceText(address) }

func (unknownNetwork) CreateBlockLink(block string) dom.Elt { return monospaceText(block) }

func (unknownNetwork) CreateTxLink(txHash string) dom.Elt { return monospaceText(txHash) }

type ropstenNetwork struct{}

func (ropstenNetwork) CreateAddressLink(address string) dom.Elt {
	return monoAnchor("https://ropsten.etherscan.io/address/%v", address)
}

func (ropstenNetwork) CreateBlockLink(block string) dom.Elt {
	return monoAnchor("https://ropsten.etherscan.io/block/%v", block)
}

func (ropstenNetwork) CreateTxLink(txHash string) dom.Elt {
	return monoAnchor("https://ropsten.etherscan.io/tx/%v", txHash)
}

type mainNetwork struct{}

func (mainNetwork) CreateAddressLink(address string) dom.Elt {
	return monoAnchor("https://etherscan.io/address/%v", address)
}

func (mainNetwork) CreateBlockLink(block string) dom.Elt {
	return monoAnchor("https://etherscan.io/block/%v", block)
}

func (mainNetwork) CreateTxLink(txHash string) dom.Elt {
	return monoAnchor("https://etherscan.io/tx/%v", txHash)
}

func monospaceText(txt string) dom.Elt {
	return dom.Span(txt).WithClass("text-monospace")
}

func monoAnchor(urlFormat, arg string) dom.Elt {
	return dom.
		Anchor("").
		WithAttribute("href", fmt.Sprintf(urlFormat, arg)).
		WithAttribute("target", "_blank").
		WithChildren(monospaceText(arg))
}

func getNetworkType(url string) network {
	if strings.Contains(url, "ropsten") {
		return ropstenNetwork{}
	}
	if strings.Contains(url, "mainnet") {
		return mainNetwork{}
	}
	return unknownNetwork{}
}
