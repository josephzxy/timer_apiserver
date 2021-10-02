package timer

import (
	"errors"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"

	"github.com/josephzxy/timer_apiserver/internal/resource/v1/model"
	"github.com/josephzxy/timer_apiserver/internal/resource/v1/service"
)

func Test_TimerController_Get(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockTimerService := service.NewMockTimerService(ctrl)
	mockTimerService.EXPECT().GetByName(gomock.Any()).AnyTimes().Return(&model.Timer{}, nil)

	mockServiceRouter := service.NewMockServiceRouter(ctrl)
	mockServiceRouter.EXPECT().Timer().AnyTimes().Return(mockTimerService)

	tc := &TimerController{serviceRouter: mockServiceRouter}
	c := newTestGinCtxWithReq("GET", "/v1/timers/name", nil)
	tc.Get(c)
	assert.Equal(t, 200, c.Writer.Status())
}

func Test_TimerController_Get_err(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockTimerService := service.NewMockTimerService(ctrl)
	mockTimerService.EXPECT().GetByName(gomock.Any()).AnyTimes().Return(nil, errors.New(""))

	mockServiceRouter := service.NewMockServiceRouter(ctrl)
	mockServiceRouter.EXPECT().Timer().AnyTimes().Return(mockTimerService)

	tc := &TimerController{serviceRouter: mockServiceRouter}
	c := newTestGinCtxWithReq("GET", "/v1/timers/name", nil)
	tc.Get(c)
	assert.NotEqual(t, 200, c.Writer.Status())
}

func Test_TimerController_GetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	tests := []struct {
		name   string
		getErr error
	}{
		{"success", nil},
		{"failure", errors.New("")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockTimerService := service.NewMockTimerService(ctrl)
			mockTimerService.EXPECT().GetAll().AnyTimes().Return(nil, tt.getErr)

			mockServiceRouter := service.NewMockServiceRouter(ctrl)
			mockServiceRouter.EXPECT().Timer().AnyTimes().Return(mockTimerService)

			tc := &TimerController{serviceRouter: mockServiceRouter}
			c := newTestGinCtxWithReq("GET", "/v1/timers", nil)
			tc.GetAll(c)
			if tt.getErr != nil {
				assert.NotEqual(t, 200, c.Writer.Status())
			} else {
				assert.Equal(t, 200, c.Writer.Status())
			}
		})
	}
}
