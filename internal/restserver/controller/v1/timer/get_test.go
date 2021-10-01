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
