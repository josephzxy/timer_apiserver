package timer

import (
	"errors"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/josephzxy/timer_apiserver/internal/resource/v1/model"
	"github.com/josephzxy/timer_apiserver/internal/resource/v1/service"
)

func monkeypatch_validateTimerCoreFunc(ret error) (restore func()) {
	old := validateTimerCoreFunc
	restore = func() { validateTimerCoreFunc = old }
	validateTimerCoreFunc = func(tc *model.TimerCore) error { return ret }
	return
}

func Test_TimerController_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	ginCtx := newTestGinCtxWithReq("PUT", "/v1/timers/test", nil)
	err := errors.New("")

	tests := []struct {
		name        string
		c           *gin.Context
		bindJsonErr error
		validateErr error
		updateErr   error
		http        int
	}{
		{
			"success",
			ginCtx,
			nil, nil, nil, 200,
		},
		{
			"failed to bind json",
			ginCtx,
			err, nil, nil, 400,
		},
		{
			"failed to validate",
			ginCtx,
			nil, err, nil, 400,
		},
		{
			"failed to update",
			ginCtx,
			nil, nil, err, 200,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer monkeyPatch_BindJsonFunc(tt.bindJsonErr)()
			defer monkeypatch_validateTimerCoreFunc(tt.validateErr)()

			mockTimerService := service.NewMockTimerService(ctrl)
			mockTimerService.EXPECT().UpdateByName(gomock.Any(), gomock.Any()).AnyTimes().Return(tt.updateErr)

			mockServiceRouter := service.NewMockServiceRouter(ctrl)
			mockServiceRouter.EXPECT().Timer().AnyTimes().Return(mockTimerService)

			tc := &TimerController{serviceRouter: mockServiceRouter}
			tc.Update(tt.c)
			if tt.updateErr != nil {
				assert.NotEqual(t, tt.c.Writer.Status(), tt.http)
			} else {
				assert.Equal(t, tt.c.Writer.Status(), tt.http)
			}
		})
	}
}
