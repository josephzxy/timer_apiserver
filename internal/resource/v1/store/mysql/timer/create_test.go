package timer

import (
	"errors"
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	pkgerr "github.com/josephzxy/timer_apiserver/internal/pkg/err"
	"github.com/josephzxy/timer_apiserver/internal/resource/v1/model"
)

func monkeyPatchDbCreateFunc(ret error) (restore func()) {
	old := dbCreateFunc
	restore = func() { dbCreateFunc = old }
	dbCreateFunc = func(db *gorm.DB, timer *model.Timer) error { return ret }
	return
}

func Test_TimerStore_Create(t *testing.T) {
	tests := []struct {
		name  string
		dbErr error
	}{
		{"success", nil},
		{"other error", errors.New("")},
		{"unknown mysql error", &mysql.MySQLError{}},
		{"record already exists", &mysql.MySQLError{Number: 1062, Message: ""}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer monkeyPatchDbCreateFunc(tt.dbErr)()
			ts := &timerStore{&gorm.DB{}}
			err := ts.Create(&model.Timer{})

			switch tt.name {
			case "record already exists":
				assert.Equal(t, pkgerr.ErrTimerAlreadyExists, err.(*pkgerr.WithCode).Code())
			case "unknown mysql error":
				assert.Equal(t, err.(*pkgerr.WithCode).Code(), pkgerr.ErrDatabase)
			default:
				assert.Equal(t, tt.dbErr, err)
			}
		})
	}
}
