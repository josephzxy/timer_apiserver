package timer

import (
	"github.com/gin-gonic/gin"

	"github.com/josephzxy/timer_apiserver/internal/resource/v1/service"
)

type TimerController interface {
	Create(*gin.Context)
	Get(*gin.Context)
	GetAll(*gin.Context)
	Delete(*gin.Context)
	Update(*gin.Context)
}

type timerController struct {
	serviceRouter service.ServiceRouter
}

func NewController(serviceRouter service.ServiceRouter) TimerController {
	return &timerController{serviceRouter}
}
