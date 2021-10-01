package timer

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/josephzxy/timer_apiserver/internal/resource/v1/service"
)

func Test_TimerController_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockTimerService := service.NewMockTimerService(ctrl)
	mockTimerService.EXPECT().DeleteByName(gomock.Any()).AnyTimes().Return(nil)

	mockServiceRouter := service.NewMockServiceRouter(ctrl)
	mockServiceRouter.EXPECT().Timer().AnyTimes().Return(mockTimerService)

	tc := &TimerController{serviceRouter: mockServiceRouter}
	c := newTestGinCtxWithReq("DELETE", "/v1/timers/name", nil)
	tc.Delete(c)
	assert.Equal(t, 200, c.Writer.Status())
}

func Test_TimerController_Delete_err(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockTimerService := service.NewMockTimerService(ctrl)
	mockTimerService.EXPECT().DeleteByName(gomock.Any()).AnyTimes().Return(errors.New(""))

	mockServiceRouter := service.NewMockServiceRouter(ctrl)
	mockServiceRouter.EXPECT().Timer().AnyTimes().Return(mockTimerService)

	tc := &TimerController{serviceRouter: mockServiceRouter}
	c := newTestGinCtxWithReq("DELETE", "/v1/timers/name", nil)
	tc.Delete(c)
	assert.NotEqual(t, 200, c.Writer.Status())
}
