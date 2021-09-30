package store

import "github.com/josephzxy/timer_apiserver/internal/resource/v1/model"

type TimerStore interface {
	Create(*model.Timer) error
}
