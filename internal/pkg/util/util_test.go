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
