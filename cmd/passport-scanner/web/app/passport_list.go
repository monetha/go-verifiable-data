package app

import (
	"context"
	"fmt"
	"io"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/monetha/go-verifiable-data/cmd/passport-scanner/web/rx"
	"github.com/monetha/go-verifiable-data/eth"
	"github.com/monetha/go-verifiable-data/log"
	"github.com/monetha/go-verifiable-data/passfactory"
)

type passportListObserver struct {
	OnErrorFun     func(err error)
	OnCompletedFun func()
	OnNextFun      func(passport *passfactory.Passport)
}

func (o *passportListObserver) OnError(err error) {
	o.OnErrorFun(err)
}

func (o *passportListObserver) OnCompleted() {
	o.OnCompletedFun()
}

func (o *passportListObserver) OnNext(passport *passfactory.Passport) {
	o.OnNextFun(passport)
}

type passportListGetter struct {
	Context    context.Context
	Log        log.Fun
	BackendURL string
}

func (f *passportListGetter) GetPassportListAsync(passportFactoryAddress common.Address, startFromBlock uint64, o *passportListObserver) io.Closer {
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
			Start:   startFromBlock,
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
