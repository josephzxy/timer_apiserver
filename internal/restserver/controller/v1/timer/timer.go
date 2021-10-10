// Package timer provides an interface of HTTP controller for RESTful resource Timer
// as well as an implementation.
package timer

import (
	"github.com/gin-gonic/gin"

	"github.com/josephzxy/timer_apiserver/internal/resource/v1/service"
)

// TimerController defines actions for RESTful resource Timer
// on data storage level.
type TimerController interface {
	Create(*gin.Context)
	Get(*gin.Context)
	GetAll(*gin.Context)
	Delete(*gin.Context)
	Update(*gin.Context)
}

type timerController struct {
	serviceRouter service.Router
}

// NewController returns a concrete value of interface TimerController.
func NewController(serviceRouter service.Router) TimerController {
	return &timerController{serviceRouter}
}
