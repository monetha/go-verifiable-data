package app

import (
	"context"
	"fmt"
	"io"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/monetha/protocol-go-sdk/cmd/protoscan/web/rx"
	"gitlab.com/monetha/protocol-go-sdk/eth"
	"gitlab.com/monetha/protocol-go-sdk/facts"
	"gitlab.com/monetha/protocol-go-sdk/log"
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

func (f *passportChangesGetter) GetPassportChangesAsync(passportAddress common.Address, o *passportChangesObserver) io.Closer {
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
		filterOpts := &facts.ChangesFilterOpts{Context: ctx}

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
