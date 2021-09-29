package err

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleNew() {
	err := New(ErrUnknown, "Internal server error")
	fmt.Println(err)
	fmt.Println(err.Code())

	// Output:
	// Internal server error
	// 1000
}

func TestWithCode_Code(t *testing.T) {
	tests := []struct {
		name string
		code AppErrCode
	}{
		{"ErrUnknown", ErrUnknown},
		{"ErrValidation", ErrValidation},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := &WithCode{tt.code, nil}
			assert.Equal(t, tt.code, err.Code())
		})
	}
}

func Test_New(t *testing.T) {
	type args struct {
		code AppErrCode
		msg  string
	}
	tests := []struct {
		name string
		args args
	}{
		{"ErrUnknown", args{ErrUnknown, "Internal server error"}},
		{"ErrValidation", args{ErrValidation, "Validation failed"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := New(tt.args.code, tt.args.msg)

			assert.IsType(t, &WithCode{}, err)
			assert.Equal(t, err.code, tt.args.code)
			assert.Equal(t, err.Code(), tt.args.code)
			assert.Equal(t, err.Error(), tt.args.msg)
		})
	}
}
