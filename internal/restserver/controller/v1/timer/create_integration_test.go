package timer

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/josephzxy/timer_apiserver/internal/resource/v1/service"
	"github.com/josephzxy/timer_apiserver/internal/resource/v1/store/mysql"
)

func Test_timerController_Create_integration(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"success"},
		{"validation error, empty name"},
		{"validation error, empty triggerAt"},
		{"validation error, miss name"},
		{"validation error, miss triggerAt"},
		{"validation error, triggerAt less than a minute after now"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			storeRouter, _ := mysql.NewStoreRouter(&mysql.Config{
				User:            "root",
				Pwd:             "root",
				Host:            "localhost",
				Port:            3306,
				Database:        "test",
				Charset:         "utf8mb4",
				ParseTime:       true,
				Loc:             "Local",
				MaxIdleConns:    100,
				MaxOpenConns:    100,
				MaxConnLifetime: 10 * time.Second,
				LogLevel:        1,
			})
			tc := &timerController{serviceRouter: service.NewRouter(storeRouter)}

			switch tt.name {
			case "success":
				c := newTestGinCtxWithReq("POST", "/v1/timers", map[string]interface{}{
					"name":      "test",
					"triggerAt": time.Now().AddDate(1, 0, 0),
				})
				tc.Create(c)
				assert.Equal(t, c.Writer.Status(), 200)
				assert.True(t, timerExistsInDBByName(testDB, "test"))
				// clean up
				testDB.Exec("DELETE FROM timer")

			case "validation error, empty name":
				c := newTestGinCtxWithReq("POST", "/v1/timers", map[string]interface{}{
					"name":      "",
					"triggerAt": time.Now().AddDate(1, 0, 0),
				})
				tc.Create(c)
				assert.NotEqual(t, c.Writer.Status(), 200)
				assert.False(t, timerExistsInDBByName(testDB, "test"))

			case "validation error, empty triggerAt":
				c := newTestGinCtxWithReq("POST", "/v1/timers", map[string]interface{}{
					"name":      "test",
					"triggerAt": "",
				})
				tc.Create(c)
				assert.NotEqual(t, c.Writer.Status(), 200)
				assert.False(t, timerExistsInDBByName(testDB, "test"))

			case "validation error, miss name":
				c := newTestGinCtxWithReq("POST", "/v1/timers", map[string]interface{}{
					"triggerAt": time.Now().AddDate(1, 0, 0),
				})
				tc.Create(c)
				assert.NotEqual(t, c.Writer.Status(), 200)
				assert.False(t, timerExistsInDBByName(testDB, "test"))

			case "validation error, miss triggerAt":
				c := newTestGinCtxWithReq("POST", "/v1/timers", map[string]interface{}{
					"name": "test",
				})
				tc.Create(c)
				assert.NotEqual(t, c.Writer.Status(), 200)
				assert.False(t, timerExistsInDBByName(testDB, "test"))

			case "validation error, triggerAt less than a minute after now":
				c := newTestGinCtxWithReq("POST", "/v1/timers", map[string]interface{}{
					"name":      "test",
					"triggerAt": time.Now().AddDate(-1, 0, 0),
				})
				tc.Create(c)
				assert.NotEqual(t, c.Writer.Status(), 200)
				assert.False(t, timerExistsInDBByName(testDB, "test"))
			}
		})
	}
}
