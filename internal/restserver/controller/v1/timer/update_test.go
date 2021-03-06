package timer

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/josephzxy/timer_apiserver/internal/resource/v1/service"
)

func Test_timerController_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	err := errors.New("")

	tests := []struct {
		name        string
		bindJSONErr error
		validateErr error
		updateErr   error
		http        int
	}{
		{
			"success",
			nil, nil, nil, 200,
		},
		{
			"failed to bind json",
			err, nil, nil, 400,
		},
		{
			"failed to validate",
			nil, err, nil, 400,
		},
		{
			"failed to update",
			nil, nil, err, 200,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer monkeyPatchBindJSONFunc(tt.bindJSONErr)()
			defer monkeyPatchValidateTimerCoreFunc(tt.validateErr)()

			mockTimerService := service.NewMockTimerService(ctrl)
			mockTimerService.EXPECT().UpdateByName(gomock.Any(), gomock.Any()).AnyTimes().Return(tt.updateErr)

			mockRouter := service.NewMockRouter(ctrl)
			mockRouter.EXPECT().Timer().AnyTimes().Return(mockTimerService)

			tc := &timerController{serviceRouter: mockRouter}
			c := newTestGinCtxWithReq("PUT", "/v1/timers/test", nil)
			tc.Update(c)

			switch tt.name {
			case "failed to update":
				assert.NotEqual(t, c.Writer.Status(), tt.http)
			default:
				assert.Equal(t, c.Writer.Status(), tt.http)
			}
		})
	}
}
