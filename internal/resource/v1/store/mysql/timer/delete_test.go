package timer

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	"github.com/josephzxy/timer_apiserver/internal/resource/v1/model"
)

func monkeyPatch_dbDeleteByNameFunc(ret error) (restore func()) {
	old := dbDeleteByNameFunc
	restore = func() { dbDeleteByNameFunc = old }
	dbDeleteByNameFunc = func(db *gorm.DB, name string) error { return ret }
	return
}

func monkeyPatch_tsGetByNameFunc(retTimer *model.Timer, retErr error) (restore func()) {
	old := tsGetByNameFunc
	restore = func() { tsGetByNameFunc = old }
	tsGetByNameFunc = func(ts *TimerStore, name string) (*model.Timer, error) { return retTimer, retErr }
	return
}

func Test_TimerStore_DeleteByName(t *testing.T) {
	defer monkeyPatch_dbDeleteByNameFunc(nil)()
	defer monkeyPatch_tsGetByNameFunc(nil, nil)()

	ts := &TimerStore{&gorm.DB{}}
	err := ts.DeleteByName("")
	assert.Nil(t, err)
}

func Test_TimerStore_DeleteByName_tsGetByNameErr(t *testing.T) {
	defer monkeyPatch_dbDeleteByNameFunc(nil)()

	tsGetErr := errors.New("ts get error")
	defer monkeyPatch_tsGetByNameFunc(nil, tsGetErr)()

	ts := &TimerStore{&gorm.DB{}}
	err := ts.DeleteByName("")
	assert.Equal(t, err, tsGetErr)
}

func Test_TimerStore_DeleteByName_dbDeleteByNameErr(t *testing.T) {
	dbDeleteErr := errors.New("db delete error")
	defer monkeyPatch_dbDeleteByNameFunc(dbDeleteErr)()
	defer monkeyPatch_tsGetByNameFunc(nil, nil)()

	ts := &TimerStore{&gorm.DB{}}
	err := ts.DeleteByName("")
	assert.Equal(t, err, dbDeleteErr)
}
