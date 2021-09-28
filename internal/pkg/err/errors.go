package err

import "github.com/pkg/errors"

type WithCode struct {
	code AppErrCode
	error
}

func (w *WithCode) Code() AppErrCode { return w.code }

func New(c AppErrCode, msg string) *WithCode {
	return &WithCode{
		code:  c,
		error: errors.New(msg),
	}
}
