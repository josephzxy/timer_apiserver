package timer

import (
	"errors"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	pkgerr "github.com/josephzxy/timer_apiserver/internal/pkg/err"
	"github.com/josephzxy/timer_apiserver/internal/resource/v1/service"
)

func Test_timerController_Create(t *testing.T) {
	ctrl := gomock.NewController(t)

	tests := []struct {
		name        string
		c           *gin.Context
		bindJsonErr error
		validateErr error
		createErr   error
		http        int
	}{
		{
			"success",
			newTestGinCtxWithReq("POST", "/v1/timers", nil),
			nil, nil, nil, 200,
		},
		{
			"failed to bind json",
			newTestGinCtxWithReq("POST", "/v1/timers", nil),
			errors.New(""), nil, nil, 400,
		},
		{
			"failed to validate",
			newTestGinCtxWithReq("POST", "/v1/timers", nil),
			nil, errors.New(""), nil, 400,
		},
		{
			"other error",
			newTestGinCtxWithReq("POST", "/v1/timers", nil),
			nil, nil, errors.New(""), 500,
		},
		{
			"record already exists",
			newTestGinCtxWithReq("POST", "/v1/timers", nil),
			nil, nil, pkgerr.New(pkgerr.ErrTimerAlreadyExists, ""), 400,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockTimerService := service.NewMockTimerService(ctrl)
			mockTimerService.EXPECT().Create(gomock.Any()).AnyTimes().Return(tt.createErr)

			mockServiceRouter := service.NewMockServiceRouter(ctrl)
			mockServiceRouter.EXPECT().Timer().AnyTimes().Return(mockTimerService)

			defer monkeyPatch_validateTimerFunc(tt.validateErr)()
			defer monkeyPatch_bindJsonFunc(tt.bindJsonErr)()

			tc := &timerController{serviceRouter: mockServiceRouter}
			tc.Create(tt.c)
			assert.Equal(t, tt.c.Writer.Status(), tt.http)
		})
	}
}
