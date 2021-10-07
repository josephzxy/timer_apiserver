package model

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
	"gorm.io/gorm"
)

var val *validator.Validate

func init() {
	val = validator.New()
	if err := val.RegisterValidation("notblank", validators.NotBlank); err != nil {
		panic("failed to register validation notblank")
	}
}

// TimerCore contains fields that can be specified directly via APIs
type TimerCore struct {
	Name      string    `json:"name" gorm:"index:uniq_name_alive,unique,priority:1" validate:"required,notblank"`
	TriggerAt time.Time `json:"triggerAt" gorm:"index:idx_trigger_at" validate:"required,gte=time.Now().Add(time.Minute),notblank"`
}

// ValidateTimerCore validates the given value of TimerCore
func ValidateTimerCore(tc *TimerCore) error {
	return val.Struct(tc)
}

// Timer is the data model for RESTful resource timer.
type Timer struct {
	Model
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
	return val.Struct(t)
}
