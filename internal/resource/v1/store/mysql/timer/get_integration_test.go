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

func Test_dbGetAllFunc(t *testing.T) {
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
			var fetchedTimers []model.Timer

			switch tt.name {
			case "record exists":
				plantTimerOrDie(tx, tc)
				assertTimerExists(t, tx, tc)

				err := dbGetAllFunc(tx, &fetchedTimers)
				assert.Nil(t, err)
				assert.NotNil(t, fetchedTimers)

			case "record not exist":
				assertNoTimerExists(t, tx)
				err := dbGetAllFunc(tx, &fetchedTimers)
				assert.Nil(t, err)
				assert.Equal(t, len(fetchedTimers), 0)
			}
		})
	}
}

func Test_dbGetAllPendingFunc(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"record exists, not pending"},
		{"record exists, pending"},
		{"record not exist"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx := testDB.Begin()
			defer tx.Rollback()
			name := "test"
			var fetchedTimers []model.Timer

			switch tt.name {
			case "record exists, not pending":
				triggerAt, _ := time.Parse(time.RFC3339, "1970-01-01T07:59:10+08:00")
				tc := &model.TimerCore{Name: name, TriggerAt: triggerAt}
				plantTimerOrDie(tx, tc)
				assertTimerExists(t, tx, tc)

				err := dbGetAllPendingFunc(tx, &fetchedTimers)
				assert.Nil(t, err)
				assert.Equal(t, len(fetchedTimers), 0)

			case "record exists, pending":
				triggerAt := time.Now().Add(1 * time.Minute).Truncate(time.Second)
				tc := &model.TimerCore{Name: name, TriggerAt: triggerAt}
				plantTimerOrDie(tx, tc)
				assertTimerExists(t, tx, tc)

				err := dbGetAllPendingFunc(tx, &fetchedTimers)
				assert.Nil(t, err)
				assert.Equal(t, len(fetchedTimers), 1)

			case "record not exist":
				assertNoTimerExists(t, tx)
				err := dbGetAllPendingFunc(tx, &fetchedTimers)
				assert.Nil(t, err)
				assert.Equal(t, len(fetchedTimers), 0)
			}
		})
	}
}
