package timer

import (
	"testing"
	"time"

	"github.com/josephzxy/timer_apiserver/internal/resource/v1/model"
	"github.com/stretchr/testify/assert"
)

func Test_dbGetByNameFunc(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"record exists"},
		{"record not exist"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx := testDB.Begin()
			defer tx.Rollback()
			name, triggerAt := "test", time.Now().Truncate(time.Second)
			tc := &model.TimerCore{Name: name, TriggerAt: triggerAt}
			var fetchedTimer model.Timer

			switch tt.name {
			case "record exists":
				plantTimerOrDie(tx, tc)
				assertTimerExists(t, tx, tc)

				err := dbGetByNameFunc(tx, name, &fetchedTimer)
				assert.Nil(t, err)
				assertTimerNotEmpty(t, &fetchedTimer)

			case "record not exist":
				assertTimerNotExistByName(t, tx, name)
				err := dbGetByNameFunc(tx, name, &fetchedTimer)
				assert.NotNil(t, err)
			}
		})
	}
}
