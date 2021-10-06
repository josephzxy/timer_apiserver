package timer

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	resp "github.com/josephzxy/timer_apiserver/internal/restserver/response"
)

func (tc *timerController) Get(c *gin.Context) {
	name := c.Param("name")
	timer, err := tc.serviceRouter.Timer().GetByName(name)
	if err != nil {
		zap.S().Errorw("failed to get timer", "err", err, "name", name)
		resp.WriteResponse(c, err, nil)
		return
	}
	resp.WriteResponse(c, nil, timer)
}

func (tc *timerController) GetAll(c *gin.Context) {
	timers, err := tc.serviceRouter.Timer().GetAll()
	if err != nil {
		zap.S().Errorw("failed to get timers", "err", err)
		resp.WriteResponse(c, err, nil)
		return
	}
	resp.WriteResponse(c, nil, timers)
}
