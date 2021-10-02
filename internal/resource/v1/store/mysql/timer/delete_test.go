package timer

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	pkgerr "github.com/josephzxy/timer_apiserver/internal/pkg/err"
)

func monkeyPatch_dbDeleteByNameFunc(ret error) (restore func()) {
	old := dbDeleteByNameFunc
	restore = func() { dbDeleteByNameFunc = old }
	dbDeleteByNameFunc = func(db *gorm.DB, name string) error { return ret }
	return
}

func Test_TimerStore_DeleteByName(t *testing.T) {
	defer monkeyPatch_dbDeleteByNameFunc(nil)()

	ts := &TimerStore{&gorm.DB{}}
	err := ts.DeleteByName("")
	assert.Nil(t, err)
}

func Test_TimerStore_DeleteByName_recordNotFound(t *testing.T) {
	defer monkeyPatch_dbDeleteByNameFunc(gorm.ErrRecordNotFound)()
	ts := &TimerStore{&gorm.DB{}}
	err := ts.DeleteByName("")
	assert.Equal(t, err.(*pkgerr.WithCode).Code(), pkgerr.ErrTimerNotFound)
}

func Test_TimerStore_DeleteByName_otherErr(t *testing.T) {
	dbErr := errors.New("")
	defer monkeyPatch_dbDeleteByNameFunc(dbErr)()
	ts := &TimerStore{&gorm.DB{}}
	err := ts.DeleteByName("")
	assert.Equal(t, err, dbErr)
}
