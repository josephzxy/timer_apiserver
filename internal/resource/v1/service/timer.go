package service

import (
	"github.com/josephzxy/timer_apiserver/internal/resource/v1/model"
	"github.com/josephzxy/timer_apiserver/internal/resource/v1/store"
)

// TimerService defines actions for RESTful resource Timer.
type TimerService interface {
	Create(*model.Timer) error
	GetByName(name string) (*model.Timer, error)
	GetAll() ([]model.Timer, error)
	// A timer will be considered as "pending" if it is created, not deleted, and not triggerred yet.
	GetAllPending() ([]model.Timer, error)
	DeleteByName(name string) error
	UpdateByName(name string, want *model.TimerCore) error
}

type timerService struct {
	storeRouter store.StoreRouter
}

// Create creates a new timer.
func (s *timerService) Create(timer *model.Timer) error {
	return s.storeRouter.Timer().Create(timer)
}

// GetByName gets a timer by the given name.
func (s *timerService) GetByName(name string) (*model.Timer, error) {
	return s.storeRouter.Timer().GetByName(name)
}

// DeleteByName deleted a timer by the given name.
func (s *timerService) DeleteByName(name string) error {
	return s.storeRouter.Timer().DeleteByName(name)
}

// UpdateByName updates a timer by the given name with the given desired state.
func (s *timerService) UpdateByName(name string, want *model.TimerCore) error {
	return s.storeRouter.Timer().UpdateByName(name, want)
}

// GetAll gets all timers(deleted timers excluded).
func (s *timerService) GetAll() ([]model.Timer, error) {
	return s.storeRouter.Timer().GetAll()
}

// GetAllPending gets all pending timers.
// A timer is "pending" if it is not deleted and not triggerred yet.
func (s *timerService) GetAllPending() ([]model.Timer, error) {
	return s.storeRouter.Timer().GetAllPending()
}
