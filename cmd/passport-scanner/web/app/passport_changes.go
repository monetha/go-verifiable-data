package app

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/monetha/reputation-go-sdk/cmd/passport-scanner/web/rx"
	"github.com/monetha/reputation-go-sdk/eth"
	"github.com/monetha/reputation-go-sdk/facts"
	"github.com/monetha/reputation-go-sdk/log"
	"github.com/monetha/reputation-go-sdk/types/data"
)

type passportChangesObserver struct {
	OnErrorFun     func(err error)
	OnCompletedFun func()
	OnNextFun      func(change *facts.Change)
}

func (o *passportChangesObserver) OnError(err error) {
	o.OnErrorFun(err)
}

func (o *passportChangesObserver) OnCompleted() {
	o.OnCompletedFun()
}

func (o *passportChangesObserver) OnNext(change *facts.Change) {
	o.OnNextFun(change)
}

type passportChangesGetter struct {
	Context    context.Context
	Log        log.Fun
	BackendURL string
}

func (f *passportChangesGetter) GetPassportChangesAsync(passportAddress common.Address, startFromBlock uint64, o *passportChangesObserver) io.Closer {
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

		historian := facts.NewHistorian(e)
		filterOpts := &facts.ChangesFilterOpts{
			Context: ctx,
			Start:   startFromBlock,
		}

		var it *facts.ChangeIterator
		it, err = historian.FilterChanges(filterOpts, passportAddress)

		if err != nil {
			err = fmt.Errorf("Historian.FilterChanges: %v", err)
			return
		}
		defer func() {
			if cErr := it.Close(); err == nil && cErr != nil {
				err = cErr
			}
		}()

		for it.Next() {
			if err = it.Error(); err != nil {
				err = fmt.Errorf("ChangeIterator.Next: %v", err)
				return
			}

			onNext(it.Change)
		}

		return nil
	})
}

type historyItem struct {
	FactProvider common.Address
	Key          [32]byte
	Value        []byte
}

type historyItemObserver struct {
	OnErrorFun     func(err error)
	OnCompletedFun func()
	OnNextFun      func(hi *historyItem)
}

func (o *historyItemObserver) OnError(err error) {
	o.OnErrorFun(err)
}

func (o *historyItemObserver) OnCompleted() {
	o.OnCompletedFun()
}

func (o *historyItemObserver) OnNext(hi *historyItem) {
	o.OnNextFun(hi)
}

func (f *passportChangesGetter) GetHistoryItemAsync(passportAddress common.Address, factType data.Type, txHash common.Hash, o *historyItemObserver) io.Closer {
	backendURL := f.BackendURL
	lf := f.Log
	onNext := o.OnNextFun

	return rx.RunAsyncObserver(f.Context, o, func(ctx context.Context) (err error) {
		client, err := ethclient.Dial(backendURL)
		if err != nil {
			return fmt.Errorf("ethclient.Dial: %v", err)
		}
		defer client.Close()

		e := eth.New(client, lf)

		historian := facts.NewHistorian(e)

		// read history value from transaction
		buf := new(bytes.Buffer)

		switch factType {
		case data.TxData:
			hi, err := historian.GetHistoryItemOfWriteTxData(ctx, passportAddress, txHash)
			if err != nil {
				return fmt.Errorf("Historian.GetHistoryItemOfWriteTxData(): %v", err)
			}
			buf.Write(hi.Data)
			onNext(&historyItem{
				FactProvider: hi.FactProvider,
				Key:          hi.Key,
				Value:        buf.Bytes(),
			})

		case data.String:
			hi, err := historian.GetHistoryItemOfWriteString(ctx, passportAddress, txHash)
			if err != nil {
				return fmt.Errorf("Historian.GetHistoryItemOfWriteString(): %v", err)
			}
			buf.WriteString(hi.Data)
			onNext(&historyItem{
				FactProvider: hi.FactProvider,
				Key:          hi.Key,
				Value:        buf.Bytes(),
			})

		case data.Bytes:
			hi, err := historian.GetHistoryItemOfWriteBytes(ctx, passportAddress, txHash)
			if err != nil {
				return fmt.Errorf("Historian.GetHistoryItemOfWriteBytes(): %v", err)
			}
			buf.Write(hi.Data)
			onNext(&historyItem{
				FactProvider: hi.FactProvider,
				Key:          hi.Key,
				Value:        buf.Bytes(),
			})

		case data.Address:
			hi, err := historian.GetHistoryItemOfWriteAddress(ctx, passportAddress, txHash)
			if err != nil {
				return fmt.Errorf("Historian.GetHistoryItemOfWriteAddress(): %v", err)
			}
			buf.WriteString(hi.Data.String())
			onNext(&historyItem{
				FactProvider: hi.FactProvider,
				Key:          hi.Key,
				Value:        buf.Bytes(),
			})

		case data.Uint:
			hi, err := historian.GetHistoryItemOfWriteUint(ctx, passportAddress, txHash)
			if err != nil {
				return fmt.Errorf("Historian.GetHistoryItemOfWriteUint(): %v", err)
			}
			buf.WriteString(hi.Data.String())
			onNext(&historyItem{
				FactProvider: hi.FactProvider,
				Key:          hi.Key,
				Value:        buf.Bytes(),
			})

		case data.Int:
			hi, err := historian.GetHistoryItemOfWriteInt(ctx, passportAddress, txHash)
			if err != nil {
				return fmt.Errorf("Historian.GetHistoryItemOfWriteInt(): %v", err)
			}
			buf.WriteString(hi.Data.String())
			onNext(&historyItem{
				FactProvider: hi.FactProvider,
				Key:          hi.Key,
				Value:        buf.Bytes(),
			})

		case data.Bool:
			hi, err := historian.GetHistoryItemOfWriteBool(ctx, passportAddress, txHash)
			if err != nil {
				return fmt.Errorf("Historian.GetHistoryItemOfWriteBool(): %v", err)
			}
			buf.WriteString(strconv.FormatBool(hi.Data))
			onNext(&historyItem{
				FactProvider: hi.FactProvider,
				Key:          hi.Key,
				Value:        buf.Bytes(),
			})

		case data.IPFS:
			hi, err := historian.GetHistoryItemOfWriteIPFSHash(ctx, passportAddress, txHash)
			if err != nil {
				return fmt.Errorf("Historian.GetHistoryItemOfWriteIPFSHash(): %v", err)
			}
			buf.WriteString(hi.Hash)
			onNext(&historyItem{
				FactProvider: hi.FactProvider,
				Key:          hi.Key,
				Value:        buf.Bytes(),
			})

		default:
			return fmt.Errorf("unsupported fact type: %v", factType.String())
		}

		return nil
	})
}
