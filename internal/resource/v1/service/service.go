// Package service provides mid-level interfaces for managing RESTful resources
// so as to decouple high-level interfaces like HTTP API controller with low-level interfaces
// like data storage interfaces
package service

import (
	"github.com/josephzxy/timer_apiserver/internal/resource/v1/store"
)

// ServiceRouter is an interface that provides routes to services
// dedicated to a specific scope of RESTful resource.
type ServiceRouter interface {
	Timer() TimerService
}

type serviceRouter struct {
	storeRouter store.StoreRouter
}

// Timer routes to the service for RESTful resource Timer
func (r *serviceRouter) Timer() TimerService {
	return &timerService{r.storeRouter}
}

// NewRouter returns a concrete value for interface ServiceRouter
func NewRouter(r store.StoreRouter) ServiceRouter {
	return &serviceRouter{r}
}
