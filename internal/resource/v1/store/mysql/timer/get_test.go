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

func monkeyPatchDbGetByNameFunc(ret error) (restore func()) {
	old := dbGetByNameFunc
	restore = func() { dbGetByNameFunc = old }
	dbGetByNameFunc = func(db *gorm.DB, name string, timer *model.Timer) error {
		if ret == nil {
			*timer = model.Timer{}
		}
		return ret
	}
	return
}

func monkeyPatchDbGetAllFunc(ret error) (restore func()) {
	old := dbGetAllFunc
	restore = func() { dbGetAllFunc = old }
	dbGetAllFunc = func(db *gorm.DB, timers *[]model.Timer) error {
		if ret == nil {
			*timers = make([]model.Timer, 1)
		}
		return ret
	}
	return
}

func monkeyPatchDbGetAllPendingFunc(ret error) (restore func()) {
	old := dbGetAllPendingFunc
	restore = func() { dbGetAllPendingFunc = old }
	dbGetAllPendingFunc = func(db *gorm.DB, timers *[]model.Timer) error {
		if ret == nil {
			*timers = make([]model.Timer, 1)
		}
		return ret
	}
	return
}

func Test_TimerStore_GetByName(t *testing.T) {
	tests := []struct {
		name  string
		dbErr error
	}{
		{"success", nil},
		{"record not found", gorm.ErrRecordNotFound},
		{"unknown mysql error", &mysql.MySQLError{}},
		{"other error", errors.New("")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer monkeyPatchDbGetByNameFunc(tt.dbErr)()
			ts := &timerStore{&gorm.DB{}}
			timer, err := ts.GetByName("")

			switch tt.name {
			case "success":
				assert.NotNil(t, timer)
				assert.Equal(t, err, tt.dbErr)
			case "record not found":
				assert.Nil(t, timer)
				assert.Equal(t, pkgerr.ErrTimerNotFound, err.(*pkgerr.WithCode).Code())
			case "unknown mysql error":
				assert.Nil(t, timer)
				assert.Equal(t, err.(*pkgerr.WithCode).Code(), pkgerr.ErrDatabase)
			default:
				assert.Nil(t, timer)
				assert.Equal(t, err, tt.dbErr)
			}
		})
	}
}

func Test_TimerStore_GetAll(t *testing.T) {
	tests := []struct {
		name  string
		dbErr error
	}{
		{"success", nil},
		{"unknown mysql error", &mysql.MySQLError{}},
		{"other error", errors.New("")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer monkeyPatchDbGetAllFunc(tt.dbErr)()
			ts := &timerStore{&gorm.DB{}}
			data, err := ts.GetAll()
			switch tt.name {
			case "success":
				assert.Equal(t, err, tt.dbErr)
				assert.NotNil(t, data)
			case "unknown mysql error":
				assert.Equal(t, err.(*pkgerr.WithCode).Code(), pkgerr.ErrDatabase)
				assert.Nil(t, data)
			case "other error":
				assert.Equal(t, err, tt.dbErr)
				assert.Nil(t, data)
			}
		})
	}
}

func Test_TimerStore_GetAllPending(t *testing.T) {
	tests := []struct {
		name  string
		dbErr error
	}{
		{"success", nil},
		{"unknown mysql error", &mysql.MySQLError{}},
		{"other error", errors.New("")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer monkeyPatchDbGetAllPendingFunc(tt.dbErr)()
			ts := &timerStore{&gorm.DB{}}
			data, err := ts.GetAllPending()

			switch tt.name {
			case "success":
				assert.Equal(t, err, tt.dbErr)
				assert.NotNil(t, data)
			case "unknown mysql error":
				assert.Equal(t, err.(*pkgerr.WithCode).Code(), pkgerr.ErrDatabase)
				assert.Nil(t, data)
			case "other error":
				assert.Equal(t, err, tt.dbErr)
				assert.Nil(t, data)
			}
		})
	}
}
