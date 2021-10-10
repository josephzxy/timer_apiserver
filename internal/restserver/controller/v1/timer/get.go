package timer

import (
	"github.com/gin-gonic/gin"

	resp "github.com/josephzxy/timer_apiserver/internal/restserver/response"
)

// Get gets a timer by the name provided in the path parameters.
func (tc *timerController) Get(c *gin.Context) {
	name := c.Param("name")
	timer, err := tc.serviceRouter.Timer().GetByName(name)
	if err != nil {
		resp.WriteResponse(c, err, nil)

		return
	}
	resp.WriteResponse(c, nil, timer)
}

// GetAll gets all timers(deleted timers excluded).
func (tc *timerController) GetAll(c *gin.Context) {
	timers, err := tc.serviceRouter.Timer().GetAll()
	if err != nil {
		resp.WriteResponse(c, err, nil)

		return
	}
	resp.WriteResponse(c, nil, timers)
}
