package store

import "github.com/josephzxy/timer_apiserver/internal/resource/v1/model"

type TimerStore interface {
	Create(*model.Timer) error
	GetByName(name string) (*model.Timer, error)
	DeleteByName(name string) error
}
