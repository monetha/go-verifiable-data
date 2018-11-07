package rx

import (
	"context"
	"io"
	"sync"
)

// Observer provides a mechanism for receiving push-based notifications.
type Observer interface {
	// OnError notifies the observer that the provider has experienced an error condition.
	OnError(err error)
	// OnCompleted notifies the observer that the provider has finished sending push-based notifications.
	OnCompleted()
}

// RunAsyncObserver runs `fun` asynchronously and, when it’s finished, calls the appropriate observer method OnError/OnCompleted,
// depending on the result/error.
// The function `fun` should use only OnNext method of observer.
// The return value can be used to terminate the execution of the asynchronous method.
// The Close method of returned value ensures that the asynchronous function has finished its work.
func RunAsyncObserver(ctx context.Context, o Observer, fun func(context.Context) error) io.Closer {
	return RunAsync(ctx, func(ctx context.Context) {
		if err := fun(ctx); err != nil {
			o.OnError(err)
		} else {
			o.OnCompleted()
		}
	})
}

// RunAsync runs `fun` asynchronously.
// The return value can be used to terminate the execution of the asynchronous method.
// The Close method of returned value ensures that the asynchronous function has finished its work.
func RunAsync(ctx context.Context, fun func(context.Context)) io.Closer {
	ctx2, cancel := context.WithCancel(ensureContext(ctx))
	done := make(chan struct{})
	t := &asyncTask{
		cancel: cancel,
		done:   done,
	}

	go func() {
		defer close(done)
		fun(ctx2)
	}()

	return t
}

func ensureContext(ctx context.Context) context.Context {
	if ctx == nil {
		return context.TODO()
	}
	return ctx
}

type asyncTask struct {
	closeOnce sync.Once
	cancel    context.CancelFunc
	done      <-chan struct{}
}

func (s *asyncTask) Close() error {
	s.closeOnce.Do(func() {
		s.cancel()
		<-s.done
	})
	return nil
}
