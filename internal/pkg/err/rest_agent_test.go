package err

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_newSimpleRESTAgent(t *testing.T) {
	type args struct {
		httpStatus int
		msg        string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"400",
			args{400, "Bad request"},
		},
		{
			"404",
			args{404, "Not found"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := newSimpleRESTAgent(tt.args.httpStatus, tt.args.msg)
			assert.IsType(t, &SimpleRESTAgent{}, a)
			assert.Equal(t, tt.args.httpStatus, a.http)
			assert.Equal(t, tt.args.msg, a.msg)
		})
	}
}

func Test_SimpleRESTAgent_HTTPStatus(t *testing.T) {
	tests := []struct {
		name       string
		httpStatus int
	}{
		{"400", 400},
		{"500", 500},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &SimpleRESTAgent{tt.httpStatus, ""}
			assert.Equal(t, tt.httpStatus, a.HTTPStatus())
		})
	}
}

func Test_SimpleRESTAgent_Msg(t *testing.T) {
	tests := []struct {
		name string
		msg  string
	}{
		{"foo", "foo"},
		{"bar", "bar"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &SimpleRESTAgent{msg: tt.msg}
			assert.Equal(t, tt.msg, a.Msg())
		})
	}
}

const (
	errTestOnly1 AppErrCode = iota + 9900
	errTestOnly2
	errTestOnly3
	errTestOnly4
	errTestOnly5
	errTestOnly6
)

func Test_registerRESTAgent(t *testing.T) {
	type args struct {
		code       AppErrCode
		httpStatus int
		msg        string
	}

	tests := []struct {
		name string
		args args
	}{
		{"errTestOnly1", args{
			errTestOnly1, 500, "foo",
		}},
		{"errTestOnly2", args{
			errTestOnly2, 400, "bar",
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := registerRESTAgent(tt.args.code, tt.args.httpStatus, tt.args.msg)
			assert.Nil(t, err)
			assert.Equal(t, &SimpleRESTAgent{tt.args.httpStatus, tt.args.msg}, restAgents[tt.args.code])
		})
	}
}

func Test_registerRESTAgent_HTTPStatusNotAllowed(t *testing.T) {
	type args struct {
		code       AppErrCode
		httpStatus int
		msg        string
	}
	tests := []struct {
		name string
		args args
	}{
		{"not allowed: 509", args{
			errTestOnly3, 509, "foo",
		}},
		{"not allowed: 510", args{
			errTestOnly4, 510, "bar",
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := registerRESTAgent(tt.args.code, tt.args.httpStatus, tt.args.msg)
			assert.NotNil(t, err)
			_, ok := restAgents[tt.args.code]
			assert.False(t, ok)
		})
	}
}

func Test_registerRESTAgent_errorCodeAlreadyRegistered(t *testing.T) {
	type args struct {
		code       AppErrCode
		httpStatus int
		msg        string
	}
	tests := []struct {
		name string
		args args
	}{
		{"ErrUnknown", args{
			ErrUnknown, 500, "foo",
		}},
		{"ErrValidation", args{
			ErrValidation, 400, "bar",
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := registerRESTAgent(tt.args.code, tt.args.httpStatus, tt.args.msg)
			assert.NotNil(t, err)
			agent, ok := restAgents[tt.args.code]
			assert.True(t, ok)
			assert.NotEqual(t, &SimpleRESTAgent{http: tt.args.httpStatus, msg: tt.args.msg}, agent)
		})
	}
}

func Test_GetRESTAgent(t *testing.T) {
	tests := []struct {
		name string
		code AppErrCode
	}{
		{"ErrValidation", ErrValidation},
		{"ErrDatabase", ErrDatabase},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			agent := GetRESTAgentByError(&WithCode{tt.code, nil})
			assert.Equal(t, restAgents[tt.code], agent)
		})
	}
}

func Test_GetRESTAgent_nonWithCodeError(t *testing.T) {
	tests := []struct {
		name string
		err  error
	}{
		{"nil error", nil},
		{"errors.New", errors.New("")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			agent := GetRESTAgentByError(tt.err)
			assert.Equal(t, restAgents[ErrUnknown], agent)
		})
	}
}

func Test_GetRESTAgent_nonRegisteredCode(t *testing.T) {
	tests := []struct {
		name string
		code AppErrCode
	}{
		{"errTestOnly5", errTestOnly5},
		{"errTestOnly6", errTestOnly6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			agent := GetRESTAgentByError(&WithCode{tt.code, nil})
			assert.Equal(t, restAgents[ErrUnknown], agent)
		})
	}
}
