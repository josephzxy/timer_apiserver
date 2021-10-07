package timer

import (
	"github.com/gin-gonic/gin"

	resp "github.com/josephzxy/timer_apiserver/internal/restserver/response"
)

// Delete deletes a timer by the name provided in the path parameters.
func (tc *timerController) Delete(c *gin.Context) {
	name := c.Param("name")
	err := tc.serviceRouter.Timer().DeleteByName(name)
	if err != nil {
		resp.WriteResponse(c, err, nil)
		return
	}
	resp.WriteResponse(c, nil, nil)
}
