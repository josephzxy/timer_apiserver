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
			tc := &model.TimerCore{Name: name, TriggerAt: triggerAt}

			switch tt.name {
			case "record exists":
				plantTimerOrDie(tx, tc)
				assertTimerExists(t, tx, tc)

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
