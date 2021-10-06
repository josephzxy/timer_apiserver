package util

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PanicIfErr(t *testing.T) {
	tests := []struct {
		name      string
		nonNilErr error
	}{
		{"non nil error 1", errors.New("foo")},
		{"non nil error 2", errors.New("bar")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.PanicsWithError(t, tt.nonNilErr.Error(), func() { PanicIfErr(tt.nonNilErr) })
		})
	}
}

func Test_PanicIfErr_nil(t *testing.T) {
	assert.NotPanics(t, func() { PanicIfErr(nil) })
}

func Test_BatchGoOrErr(t *testing.T) {
	testErr := errors.New("test")
	tests := []struct {
		name        string
		fs          []func() error
		expectedErr error
	}{
		{"0 goroutines", []func() error{}, nil},
		{"1 goroutines, no error", []func() error{
			func() error { return nil },
		}, nil},
		{"1 goroutines, error", []func() error{
			func() error { return testErr },
		}, testErr},
		{"1 goroutines, no error", []func() error{
			func() error { return nil },
			func() error { return nil },
		}, nil},
		{"2 goroutines, 1 error", []func() error{
			func() error { return nil },
			func() error { return testErr },
		}, testErr},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := BatchGoOrErr(tt.fs...)
			assert.Equal(t, err, tt.expectedErr)
		})
	}
}
