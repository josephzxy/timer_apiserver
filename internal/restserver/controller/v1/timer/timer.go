// Package timer provides an interface of HTTP controller for RESTful resource Timer
// as well as an implementation.
package timer

import (
	"github.com/gin-gonic/gin"

	"github.com/josephzxy/timer_apiserver/internal/resource/v1/service"
)

// Controller defines actions for RESTful resource Timer
// on data storage level.
type Controller interface {
	Create(*gin.Context)
	Get(*gin.Context)
	GetAll(*gin.Context)
	Delete(*gin.Context)
	Update(*gin.Context)
}

type timerController struct {
	serviceRouter service.Router
}

// NewController returns a concrete value of interface Controller.
func NewController(serviceRouter service.Router) Controller {
	return &timerController{serviceRouter}
}
