package timer

import (
	"gorm.io/gorm"

	"github.com/josephzxy/timer_apiserver/internal/resource/v1/store"
)

type timerStore struct {
	db *gorm.DB
}

func NewTimerStore(db *gorm.DB) store.TimerStore {
	return &timerStore{db}
}
