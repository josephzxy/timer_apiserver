package timer

import (
	"testing"
	"time"

	"github.com/josephzxy/timer_apiserver/internal/resource/v1/model"
	"github.com/stretchr/testify/assert"
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
	err := dbCreateFunc(tx, timer)
	assert.Nil(t, err)
	assertTimerExists(t, tx, &timer.TimerCore)
}
