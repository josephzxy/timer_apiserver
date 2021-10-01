package timer

import (
	"gorm.io/gorm"
)

type TimerStore struct {
	DB *gorm.DB
}
