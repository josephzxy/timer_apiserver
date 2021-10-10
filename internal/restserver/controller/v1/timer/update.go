package timer

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	pkgerr "github.com/josephzxy/timer_apiserver/internal/pkg/err"
	"github.com/josephzxy/timer_apiserver/internal/resource/v1/model"
	resp "github.com/josephzxy/timer_apiserver/internal/restserver/response"
)

var validateTimerCoreFunc = model.ValidateTimerCore

// Update updates a timer by the name provided in the path parameters
// to the desired state provided in the request body.
func (tc *timerController) Update(c *gin.Context) {
	var want model.TimerCore
	if err := bindJSONFunc(c, &want); err != nil {
		zap.S().Errorw("failed to bind data to model", "err", err)
		resp.WriteResponse(c, pkgerr.New(pkgerr.ErrValidation, err.Error()), nil)

		return
	}

	if err := validateTimerCoreFunc(&want); err != nil {
		zap.S().Errorw("failed to validate data", "err", err, "data", want)
		resp.WriteResponse(c, pkgerr.New(pkgerr.ErrValidation, err.Error()), nil)

		return
	}

	name := c.Param("name")
	if err := tc.serviceRouter.Timer().UpdateByName(name, &want); err != nil {
		resp.WriteResponse(c, err, nil)

		return
	}
	resp.WriteResponse(c, nil, nil)
}
