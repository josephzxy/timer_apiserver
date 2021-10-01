package timer

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	resp "github.com/josephzxy/timer_apiserver/internal/restserver/response"
)

func (tc *TimerController) Delete(c *gin.Context) {
	name := c.Param("name")
	err := tc.serviceRouter.Timer().DeleteByName(name)
	if err != nil {
		zap.S().Errorw("failed to delete timer", "err", err, "name", name)
		resp.WriteResponse(c, err, nil)
		return
	}
	resp.WriteResponse(c, nil, nil)
}
