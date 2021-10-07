package timer

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/josephzxy/timer_apiserver/internal/resource/v1/model"
)

func Test_dbDeleteByNameFunc(t *testing.T) {
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

			switch tt.name {
			case "record exists":
				execRaw(tx, "INSERT INTO timer (name, trigger_at) VALUES (?, ?)", name, triggerAt)
				assertTimerExists(t, tx, &model.TimerCore{Name: name, TriggerAt: triggerAt})

				err := dbDeleteByNameFunc(tx, name)
				assert.Nil(t, err)

				assertTimerNotExistByName(t, tx, name)
				var alive *bool
				queryRaw(tx, "SELECT alive FROM timer WHERE name = ?", alive, name)
				assert.Nil(t, alive)
			case "record not exist":
				assertTimerNotExistByName(t, tx, name)
				err := dbDeleteByNameFunc(tx, name)
				assert.NotNil(t, err)
			}
		})
	}
}
