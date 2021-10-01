package timer

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	pkgerr "github.com/josephzxy/timer_apiserver/internal/pkg/err"
	"github.com/josephzxy/timer_apiserver/internal/resource/v1/model"
	"github.com/josephzxy/timer_apiserver/internal/resource/v1/service"
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

func monkeyPatchTimerValidateFunc(err error) (restore func()) {
	old := validateTimerFunc
	restore = func() { validateTimerFunc = old }
	validateTimerFunc = func(*model.Timer) error { return err }
	return
}

func monkeyPatchBindJsonFunc(err error) (restore func()) {
	old := bindJsonFunc
	restore = func() { bindJsonFunc = old }
	bindJsonFunc = func(c *gin.Context, obj interface{}) error { return err }
	return
}

func Test_TimerController_Create(t *testing.T) {
	ctrl := gomock.NewController(t)

	tests := []struct {
		name             string
		c                *gin.Context
		bindJsonErr      error
		validateTimerErr error
		createTimerErr   error
		http             int
	}{
		{
			"normal",
			newTestGinCtxWithReq("POST", "/v1/timers", map[string]interface{}{
				"name":      "normal",
				"triggerAt": time.Now().Add(time.Hour),
			}),
			nil,
			nil,
			nil,
			200,
		},
		{
			"failed to bind json",
			newTestGinCtxWithReq("POST", "/v1/timers", nil),
			errors.New(""),
			nil,
			nil,
			400,
		},
		{
			"failed to validate timer",
			newTestGinCtxWithReq("POST", "/v1/timers", nil),
			nil,
			errors.New(""),
			nil,
			400,
		},
		{
			"failed to create timer-ErrUnknown",
			newTestGinCtxWithReq("POST", "/v1/timers", nil),
			nil,
			nil,
			errors.New(""),
			500,
		},
		{
			"failed to create timer-ErrTimerAlreadyExists",
			newTestGinCtxWithReq("POST", "/v1/timers", nil),
			nil,
			nil,
			pkgerr.New(pkgerr.ErrTimerAlreadyExists, ""),
			400,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockTimerService := service.NewMockTimerService(ctrl)
			mockTimerService.EXPECT().Create(gomock.Any()).AnyTimes().Return(tt.createTimerErr)

			mockServiceRouter := service.NewMockServiceRouter(ctrl)
			mockServiceRouter.EXPECT().Timer().AnyTimes().Return(mockTimerService)

			defer monkeyPatchTimerValidateFunc(tt.validateTimerErr)()
			defer monkeyPatchBindJsonFunc(tt.bindJsonErr)()

			tc := &TimerController{serviceRouter: mockServiceRouter}
			tc.Create(tt.c)
			assert.Equal(t, tt.c.Writer.Status(), tt.http)
		})
	}
}
