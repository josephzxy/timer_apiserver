package timer

import (
	"net/http"

	"github.com/gin-gonic/gin"

	model "github.com/josephzxy/timer_apiserver/internal/resource/model/v1"
)

func (tc *TimerController) Create(c *gin.Context) {
	var timer model.Timer
	if err := c.ShouldBindJSON(&timer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "failed to validate request"})
		return
	}

	if err := timer.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "failed to validate request"})
		return
	}

	if err := tc.serviceRouter.Timer().Create(&timer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, timer)
}
