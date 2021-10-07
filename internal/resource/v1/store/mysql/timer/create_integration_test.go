package timer

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/josephzxy/timer_apiserver/internal/resource/v1/model"
)

// Integration test for timerStore.Create

func Test_dbCreateFunc(t *testing.T) {
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
			timer := &model.Timer{TimerCore: *tc}

			switch tt.name {
			case "record exists":
				plantTimerOrDie(tx, tc)
				assertTimerExists(t, tx, tc)
				err := dbCreateFunc(tx, timer)
				assert.NotNil(t, err)
			case "record not exist":
				assertTimerNotExistByName(t, tx, name)
				err := dbCreateFunc(tx, timer)
				assert.Nil(t, err)
				assertTimerExists(t, tx, tc)
			}
		})
	}
}
