package timer

import (
	"github.com/gin-gonic/gin"
	resp "github.com/josephzxy/timer_apiserver/internal/restserver/response"
	"go.uber.org/zap"
)

func (tc *TimerController) Get(c *gin.Context) {
	name := c.Param("name")
	timer, err := tc.serviceRouter.Timer().GetByName(name)
	if err != nil {
		zap.S().Errorw("faield to get timer", "err", err, "name", name)
		resp.WriteResponse(c, err, nil)
		return
	}
	resp.WriteResponse(c, nil, timer)
}
