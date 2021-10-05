package service

import (
	"github.com/josephzxy/timer_apiserver/internal/resource/v1/model"
	"github.com/josephzxy/timer_apiserver/internal/resource/v1/store"
)

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

func (s *timerService) Create(timer *model.Timer) error {
	return s.storeRouter.Timer().Create(timer)
}

func (s *timerService) GetByName(name string) (*model.Timer, error) {
	return s.storeRouter.Timer().GetByName(name)
}

func (s *timerService) DeleteByName(name string) error {
	return s.storeRouter.Timer().DeleteByName(name)
}

func (s *timerService) UpdateByName(name string, want *model.TimerCore) error {
	return s.storeRouter.Timer().UpdateByName(name, want)
}

func (s *timerService) GetAll() ([]model.Timer, error) {
	return s.storeRouter.Timer().GetAll()
}

func (s *timerService) GetAllPending() ([]model.Timer, error) {
	return s.storeRouter.Timer().GetAllPending()
}
