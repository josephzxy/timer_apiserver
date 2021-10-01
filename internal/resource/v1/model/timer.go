package model

import (
	"time"

	"github.com/go-playground/validator/v10"
)

// TimerCore contains fields that can be specified directly via APIs
type TimerCore struct {
	Name      string    `json:"name" gorm:"unique" validate:"required"`
	TriggerAt time.Time `json:"triggerAt" gorm:"index" validate:"required,gte=time.Now().Add(time.Minute)"`
}

type Timer struct {
	Model
	TimerCore
}

func (t *Timer) TableName() string {
	return "timer"
}

func ValidateTimer(t *Timer) error {
	return validator.New().Struct(t)
}
