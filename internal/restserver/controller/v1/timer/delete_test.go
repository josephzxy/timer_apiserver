package timer

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/josephzxy/timer_apiserver/internal/resource/v1/service"
)

func Test_timerController_Delete(t *testing.T) {
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
			mockTimerService.EXPECT().DeleteByName(gomock.Any()).AnyTimes().Return(tt.err)

			mockServiceRouter := service.NewMockServiceRouter(ctrl)
			mockServiceRouter.EXPECT().Timer().AnyTimes().Return(mockTimerService)

			tc := &timerController{serviceRouter: mockServiceRouter}
			c := newTestGinCtxWithReq("DELETE", "/v1/timers/name", nil)
			tc.Delete(c)

			switch tt.name {
			case "success":
				assert.Equal(t, 200, c.Writer.Status())
			case "failure":
				assert.NotEqual(t, 200, c.Writer.Status())
			}
		})
	}
}
