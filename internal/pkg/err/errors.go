// Package err implements a library for the app to handle
// the representation of errors to end users.
package err

import "github.com/pkg/errors"

// WithCode implements interface error as a wrapper of
// an error with AppErrCode.
type WithCode struct {
	code AppErrCode
	error
}

// Code returns the AppErrCode associated.
func (w *WithCode) Code() AppErrCode { return w.code }

// New creates a new WithCode error with the given
// AppErrCode and error message.
func New(c AppErrCode, msg string) *WithCode {
	return &WithCode{
		code:  c,
		error: errors.New(msg),
	}
}
