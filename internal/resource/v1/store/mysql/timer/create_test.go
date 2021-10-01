package timer

import (
	"errors"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"

	pkgerr "github.com/josephzxy/timer_apiserver/internal/pkg/err"
	"github.com/josephzxy/timer_apiserver/internal/resource/v1/model"
)

func monkeyPatchDbCreateFunc(ret error) (restore func()) {
	old := dbCreateFunc
	restore = func() { dbCreateFunc = old }
	dbCreateFunc = func(db *gorm.DB, value interface{}) error { return ret }
	return
}

func Test_TimerStore_Create(t *testing.T) {
	nonMysqlErr := errors.New("")
	nonSupportedMysqlErr := &mysql.MySQLError{}

	tests := []struct {
		name        string
		dbCreateErr error
		want        error
	}{
		{"normal", nil, nil},
		{"non mysql error", nonMysqlErr, nonMysqlErr},
		{"non supported mysql error", nonSupportedMysqlErr, nonSupportedMysqlErr},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer monkeyPatchDbCreateFunc(tt.dbCreateErr)()
			ts := &TimerStore{&gorm.DB{}}
			got := ts.Create(&model.Timer{})
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_TimerStore_Create_supportedMysqlErr(t *testing.T) {
	mysqlErr := &mysql.MySQLError{Number: 1062, Message: ""}
	defer monkeyPatchDbCreateFunc(mysqlErr)()

	ts := &TimerStore{&gorm.DB{}}
	got := ts.Create(&model.Timer{})
	assert.Equal(t, pkgerr.ErrTimerAlreadyExists, got.(*pkgerr.WithCode).Code())
}
