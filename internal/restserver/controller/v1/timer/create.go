package timer

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (tc *TimerController) Create(c *gin.Context) {
	// request validation
	// business logic
	// write response
	c.JSON(http.StatusOK, gin.H{"hello": "world"})
}
