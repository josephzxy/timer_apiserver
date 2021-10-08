package timer

import (
	"errors"
	"testing"

	"github.com/go-sql-driver/mysql"
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
	tests := []struct {
		name  string
		dbErr error
	}{
		{"success", nil},
		{"other error", errors.New("")},
		{"unknown mysql error", &mysql.MySQLError{}},
		{"record not found", gorm.ErrRecordNotFound},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer monkeyPatch_dbDeleteByNameFunc(tt.dbErr)()
			ts := &timerStore{&gorm.DB{}}
			err := ts.DeleteByName("")

			switch tt.name {
			case "record not found":
				assert.Equal(t, err.(*pkgerr.WithCode).Code(), pkgerr.ErrTimerNotFound)
			case "unknown mysql error":
				assert.Equal(t, err.(*pkgerr.WithCode).Code(), pkgerr.ErrDatabase)
			default:
				assert.Equal(t, err, tt.dbErr)
			}
		})
	}
}
