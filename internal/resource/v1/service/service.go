// Package service provides mid-level interfaces for managing RESTful resources
// so as to decouple high-level interfaces like HTTP API controller with low-level interfaces
// like data storage interfaces
package service

import (
	"github.com/josephzxy/timer_apiserver/internal/resource/v1/store"
)

// Router is an interface that provides routes to services
// dedicated to a specific scope of RESTful resource.
type Router interface {
	Timer() TimerService
}

type router struct {
	storeRouter store.Router
}

// Timer routes to the service for RESTful resource Timer
func (r *router) Timer() TimerService {
	return &timerService{r.storeRouter}
}

// NewRouter returns a concrete value for interface Router
func NewRouter(r store.Router) Router {
	return &router{r}
}
