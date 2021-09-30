package timer

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/josephzxy/timer_apiserver/internal/resource/v1/model"
	resp "github.com/josephzxy/timer_apiserver/internal/restserver/response"
)

func (tc *TimerController) Create(c *gin.Context) {
	var timer model.Timer
	if err := c.ShouldBindJSON(&timer); err != nil {
		zap.S().Errorw("failed to bind data to model", "err", err)
		resp.WriteResponse(c, err, nil)
		return
	}

	if err := timer.Validate(); err != nil {
		zap.S().Errorw("data validation failed", "err", err)
		resp.WriteResponse(c, err, nil)
		return
	}

	if err := tc.serviceRouter.Timer().Create(&timer); err != nil {
		zap.S().Errorw("failed to create timer", "err", err)
		resp.WriteResponse(c, err, nil)
		return
	}

	resp.WriteResponse(c, nil, timer)
}
