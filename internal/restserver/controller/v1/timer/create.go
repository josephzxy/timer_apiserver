package timer

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	pkgerr "github.com/josephzxy/timer_apiserver/internal/pkg/err"
	"github.com/josephzxy/timer_apiserver/internal/resource/v1/model"
	resp "github.com/josephzxy/timer_apiserver/internal/restserver/response"
)

var validateTimerFunc = model.ValidateTimer
var bindJSONFunc = func(c *gin.Context, obj interface{}) error {
	return c.ShouldBindJSON(obj)
}

// Create creates a timer by the data provided in the request body.
func (tc *timerController) Create(c *gin.Context) {
	var timer model.Timer
	if err := bindJSONFunc(c, &timer); err != nil {
		zap.S().Errorw("failed to bind data to model", "err", err)
		resp.WriteResponse(c, pkgerr.New(pkgerr.ErrValidation, err.Error()), nil)

		return
	}

	if err := validateTimerFunc(&timer); err != nil {
		zap.S().Errorw("data validation failed", "err", err, "data", timer)
		resp.WriteResponse(c, pkgerr.New(pkgerr.ErrValidation, err.Error()), nil)

		return
	}

	if err := tc.serviceRouter.Timer().Create(&timer); err != nil {
		resp.WriteResponse(c, err, nil)

		return
	}

	resp.WriteResponse(c, nil, timer)
}
