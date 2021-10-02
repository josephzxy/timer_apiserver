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
	dbGetByNameFunc = func(db *gorm.DB, name string, timer *model.Timer) error { return ret }
	return
}

func monkeyPatch_dbGetAllFunc(ret error) (restore func()) {
	old := dbGetAllFunc
	restore = func() { dbGetAllFunc = old }
	dbGetAllFunc = func(db *gorm.DB, timers *[]model.Timer) error { return ret }
	return
}

func Test_TimerStore_GetByName(t *testing.T) {
	defer monkeyPatch_dbGetByNameFunc(nil)()
	ts := &TimerStore{&gorm.DB{}}
	timer, err := ts.GetByName("")
	assert.NotNil(t, timer)
	assert.Nil(t, err)
}

func Test_TimerStore_GetByName_recordNotFound(t *testing.T) {
	defer monkeyPatch_dbGetByNameFunc(gorm.ErrRecordNotFound)()
	ts := &TimerStore{&gorm.DB{}}
	timer, err := ts.GetByName("")
	assert.Nil(t, timer)
	assert.NotNil(t, err)
	assert.Equal(t, pkgerr.ErrTimerNotFound, err.(*pkgerr.WithCode).Code())
}

func Test_TimerStore_GetByName_otherErr(t *testing.T) {
	defer monkeyPatch_dbGetByNameFunc(errors.New(""))()
	ts := &TimerStore{&gorm.DB{}}
	timer, err := ts.GetByName("")
	assert.Nil(t, timer)
	assert.NotNil(t, err)
	_, ok := err.(*pkgerr.WithCode)
	assert.False(t, ok)
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
			_, err := ts.GetAll()
			assert.Equal(t, err, tt.dbErr)
		})
	}
}
