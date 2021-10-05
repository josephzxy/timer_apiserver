package timer

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	pkgerr "github.com/josephzxy/timer_apiserver/internal/pkg/err"
	"github.com/josephzxy/timer_apiserver/internal/resource/v1/model"
)

func monkeyPatch_dbGetByNameFunc(ret error) (restore func()) {
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

func monkeyPatch_dbGetAllFunc(ret error) (restore func()) {
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

func monkeyPatch_dbGetAllPendingFunc(ret error) (restore func()) {
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
		{"other error", errors.New("")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer monkeyPatch_dbGetByNameFunc(tt.dbErr)()
			ts := &TimerStore{&gorm.DB{}}
			timer, err := ts.GetByName("")

			switch tt.name {
			case "success":
				assert.NotNil(t, timer)
				assert.Equal(t, err, tt.dbErr)
			case "record not found":
				assert.Nil(t, timer)
				assert.Equal(t, pkgerr.ErrTimerNotFound, err.(*pkgerr.WithCode).Code())
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
		{"failure", errors.New("")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer monkeyPatch_dbGetAllFunc(tt.dbErr)()
			ts := &TimerStore{&gorm.DB{}}
			data, err := ts.GetAll()
			switch tt.name {
			case "success":
				assert.Equal(t, err, tt.dbErr)
				assert.NotNil(t, data)
			case "failure":
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
		{"failure", errors.New("")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer monkeyPatch_dbGetAllPendingFunc(tt.dbErr)()
			ts := &TimerStore{&gorm.DB{}}
			data, err := ts.GetAllPending()

			switch tt.name {
			case "success":
				assert.Equal(t, err, tt.dbErr)
				assert.NotNil(t, data)
			case "failure":
				assert.Equal(t, err, tt.dbErr)
				assert.Nil(t, data)
			}
		})
	}
}
