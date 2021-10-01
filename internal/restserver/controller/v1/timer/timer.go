package timer

import (
	"github.com/josephzxy/timer_apiserver/internal/resource/v1/service"
)

type TimerController struct {
	serviceRouter service.ServiceRouter
}

func NewController(serviceRouter service.ServiceRouter) *TimerController {
	return &TimerController{serviceRouter}
}
