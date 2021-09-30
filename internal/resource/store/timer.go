package store

import model "github.com/josephzxy/timer_apiserver/internal/resource/model/v1"

type TimerStore interface {
	Create(*model.Timer) error
}
