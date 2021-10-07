package timer

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/josephzxy/timer_apiserver/internal/resource/v1/model"
)

func Test_dbUpdateByNameFunc(t *testing.T) {
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
			newTc := &model.TimerCore{Name: "test2", TriggerAt: time.Now().Add(time.Hour).Truncate(time.Second)}

			switch tt.name {
			case "record exists":
				plantTimerOrDie(tx, tc)
				assertTimerExists(t, tx, tc)

				err := dbUpdateByNameFunc(tx, name, newTc)
				assert.Nil(t, err)

				assertTimerNotExistByName(t, tx, name)
				assertTimerExists(t, tx, newTc)

			case "record not exist":
				assertTimerNotExistByName(t, tx, name)
				err := dbUpdateByNameFunc(tx, name, newTc)
				assert.NotNil(t, err)
			}
		})
	}
}
