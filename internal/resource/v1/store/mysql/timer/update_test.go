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

func monkeyPatch_dbUpdateByNameFunc(ret error) (restore func()) {
	old := dbUpdateByNameFunc
	restore = func() { dbUpdateByNameFunc = old }
	dbUpdateByNameFunc = func(db *gorm.DB, name string, want *model.TimerCore) error { return ret }
	return
}

func Test_TimerStore_UpdateByName(t *testing.T) {
	defer monkeyPatch_dbUpdateByNameFunc(nil)()
	ts := &TimerStore{&gorm.DB{}}
	err := ts.UpdateByName("", nil)
	assert.Nil(t, err)
}

func Test_TimerStore_UpdateByName_recordNotFound(t *testing.T) {
	defer monkeyPatch_dbUpdateByNameFunc(gorm.ErrRecordNotFound)()
	ts := &TimerStore{&gorm.DB{}}
	err := ts.UpdateByName("", nil)
	assert.Equal(t, err.(*pkgerr.WithCode).Code(), pkgerr.ErrTimerNotFound)
}

func Test_TimerStore_UpdateByName_timerAlreadyExists(t *testing.T) {
	defer monkeyPatch_dbUpdateByNameFunc(&mysql.MySQLError{Number: 1062})()
	ts := &TimerStore{&gorm.DB{}}
	err := ts.UpdateByName("", nil)
	assert.Equal(t, err.(*pkgerr.WithCode).Code(), pkgerr.ErrTimerAlreadyExists)
}

func Test_TimerStore_UpdateByName_otherErr(t *testing.T) {

	unknownErr := errors.New("")
	unknownMysqlErr := &mysql.MySQLError{}

	tests := []struct {
		name  string
		dbErr error
	}{
		{"unknown error", unknownErr},
		{"unknown mysql error", unknownMysqlErr},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer monkeyPatch_dbUpdateByNameFunc(tt.dbErr)()
			ts := &TimerStore{&gorm.DB{}}
			err := ts.UpdateByName("", nil)
			assert.Equal(t, tt.dbErr, err)
		})
	}
}
