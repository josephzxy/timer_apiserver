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
	tests := []struct {
		name  string
		dbErr error
	}{
		{"success", nil},
		{"record not found", gorm.ErrRecordNotFound},
		{"record already exists", &mysql.MySQLError{Number: 1062}},
		{"unknown mysql error", &mysql.MySQLError{}},
		{"other error", errors.New("")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer monkeyPatch_dbUpdateByNameFunc(tt.dbErr)()
			ts := &timerStore{&gorm.DB{}}
			err := ts.UpdateByName("", nil)

			switch tt.name {
			case "record not found":
				assert.Equal(t, err.(*pkgerr.WithCode).Code(), pkgerr.ErrTimerNotFound)
			case "record already exists":
				assert.Equal(t, err.(*pkgerr.WithCode).Code(), pkgerr.ErrTimerAlreadyExists)
			default:
				assert.Equal(t, err, tt.dbErr)
			}
		})
	}
}
