package timer

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/josephzxy/timer_apiserver/internal/resource/v1/model"
	"github.com/josephzxy/timer_apiserver/internal/resource/v1/service"
)

func Test_timerController_Get(t *testing.T) {
	ctrl := gomock.NewController(t)

	tests := []struct {
		name string
		err  error
	}{
		{"success", nil},
		{"failure", errors.New("")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockTimerService := service.NewMockTimerService(ctrl)
			mockTimerService.EXPECT().GetByName(gomock.Any()).AnyTimes().Return(&model.Timer{}, tt.err)

			mockRouter := service.NewMockRouter(ctrl)
			mockRouter.EXPECT().Timer().AnyTimes().Return(mockTimerService)

			tc := &timerController{serviceRouter: mockRouter}
			c := newTestGinCtxWithReq("GET", "/v1/timers/name", nil)
			tc.Get(c)
			switch tt.name {
			case "success":
				assert.Equal(t, 200, c.Writer.Status())
			case "failure":
				assert.NotEqual(t, 200, c.Writer.Status())
			}
		})
	}
}

func Test_timerController_GetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	tests := []struct {
		name string
		err  error
	}{
		{"success", nil},
		{"failure", errors.New("")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockTimerService := service.NewMockTimerService(ctrl)
			mockTimerService.EXPECT().GetAll().AnyTimes().Return(nil, tt.err)

			mockRouter := service.NewMockRouter(ctrl)
			mockRouter.EXPECT().Timer().AnyTimes().Return(mockTimerService)

			tc := &timerController{serviceRouter: mockRouter}
			c := newTestGinCtxWithReq("GET", "/v1/timers", nil)
			tc.GetAll(c)

			switch tt.name {
			case "success":
				assert.Equal(t, 200, c.Writer.Status())
			case "failure":
				assert.NotEqual(t, 200, c.Writer.Status())
			}
		})
	}
}
