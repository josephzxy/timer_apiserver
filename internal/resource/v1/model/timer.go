package model

import (
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// TimerCore contains fields that can be specified directly via APIs
type TimerCore struct {
	Name      string    `json:"name" gorm:"index:uniq_name_alive,unique,priority:1" validate:"required"`
	TriggerAt time.Time `json:"triggerAt" gorm:"index:idx_trigger_at" validate:"required,gte=time.Now().Add(time.Minute)"`
}

// ValidateTimerCore validates the given value of TimerCore
func ValidateTimerCore(tc *TimerCore) error {
	return validator.New().Struct(tc)
}

// Timer is the data model for RESTful resource timer.
type Timer struct {
	model
	TimerCore
	Alive     bool           `json:"-" gorm:"index:uniq_name_alive,unique,priority:2"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

// TableName tells gorm the name of the corresponding database table.
func (t *Timer) TableName() string {
	return "timer"
}

// ValidateTimer validates the given value of Timer
func ValidateTimer(t *Timer) error {
	return validator.New().Struct(t)
}
