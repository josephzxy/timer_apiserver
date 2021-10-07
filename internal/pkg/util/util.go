// Package util provides utility functions
package util

import (
	"context"
	"reflect"
	"runtime"

	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

// PanicIfErr panics only if the given error is non-nil
func PanicIfErr(e error) {
	if e != nil {
		panic(e)
	}
}

// BatchGoOrErr starts multiple goroutines and returns the error
// once a goroutine stops with an error.
// CAUTION: Other goroutines aren't and won't be closed when a goroutine stops with an error.
// In other words, those still-running goroutines "leak"
// It is safe only when the error returned by BatchGoOrEr will eventually quit
// the program
func BatchGoOrErr(fs ...func() error) error {
	waitDone := make(chan struct{}, 1)
	var firstErr error
	eg, ctx := errgroup.WithContext(context.Background())

	for _, f := range fs {
		fn := f
		eg.Go(func() error {
			if err := fn(); err != nil {
				zap.S().Errorw("Function failed in BatchGoOrErr", "func", GetFuncName(fn), "err", err)
				firstErr = err
				return err
			}
			return nil
		})
	}

	go func() {
		_ = eg.Wait()
		waitDone <- struct{}{}
	}()

	select {
	case <-ctx.Done():
		return firstErr
	case <-waitDone:
		return nil
	}
}

// GetFuncName returns the name of the given function
// Reference: https://stackoverflow.com/a/7053871/12023612
func GetFuncName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
