package service

import (
	model "github.com/josephzxy/timer_apiserver/internal/resource/model/v1"
	"github.com/josephzxy/timer_apiserver/internal/resource/store"
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
