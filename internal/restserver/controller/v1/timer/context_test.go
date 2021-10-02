// context_test.go holds auxiliary functions shared within this package.
package timer

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"

	"github.com/josephzxy/timer_apiserver/internal/resource/v1/model"
)

func newTestGinCtxWithReq(method, url string, body map[string]interface{}) *gin.Context {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	var reqBody io.Reader
	switch body {
	case nil:
		reqBody = nil
	default:
		bodyBytes, _ := json.Marshal(body)
		reqBody = bytes.NewBuffer(bodyBytes)
	}

	c.Request, _ = http.NewRequest(method, url, reqBody)
	c.Request.Header.Set("Content-Type", "application/json")
	return c
}

func monkeyPatch_validateTimerFunc(ret error) (restore func()) {
	old := validateTimerFunc
	restore = func() { validateTimerFunc = old }
	validateTimerFunc = func(*model.Timer) error { return ret }
	return
}

func monkeyPatch_bindJsonFunc(ret error) (restore func()) {
	old := bindJsonFunc
	restore = func() { bindJsonFunc = old }
	bindJsonFunc = func(c *gin.Context, obj interface{}) error { return ret }
	return
}

func monkeypatch_validateTimerCoreFunc(ret error) (restore func()) {
	old := validateTimerCoreFunc
	restore = func() { validateTimerCoreFunc = old }
	validateTimerCoreFunc = func(tc *model.TimerCore) error { return ret }
	return
}
