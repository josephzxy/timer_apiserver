package timer

import (
	"testing"
	"time"

	"github.com/josephzxy/timer_apiserver/internal/resource/v1/model"
)

// Integration test for timerStore.Create

func Test_dbCreateFunc(t *testing.T) {
	tx := testDB.Begin()
	defer tx.Rollback()

	timer := &model.Timer{
		TimerCore: model.TimerCore{
			Name:      "test",
			TriggerAt: time.Now().Truncate(time.Second),
		},
	}
	assertTimerNotExistByName(t, tx, timer.Name)
	dbCreateFunc(tx, timer)
	assertTimerExists(t, tx, &timer.TimerCore)
}
