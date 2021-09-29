package service

import (
	"github.com/josephzxy/timer_apiserver/internal/resource/store"
)

type ServiceRouter interface {
	Timer() TimerService
}

type serviceRouter struct {
	storeRouter store.StoreRouter
}

func (r *serviceRouter) Timer() TimerService {
	return &timerService{r.storeRouter}
}

func NewRouter(r store.StoreRouter) ServiceRouter {
	return &serviceRouter{r}
}
