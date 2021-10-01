package service

import (
	"github.com/josephzxy/timer_apiserver/internal/resource/v1/model"
	"github.com/josephzxy/timer_apiserver/internal/resource/v1/store"
)

type TimerService interface {
	Create(*model.Timer) error
}

type timerService struct {
	storeRouter store.StoreRouter
}

func (s *timerService) Create(timer *model.Timer) error {
	return s.storeRouter.Timer().Create(timer)
}
