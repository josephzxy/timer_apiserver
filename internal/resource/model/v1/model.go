package v1

import (
	"time"

	"gorm.io/gorm"
)

// Model is the slightly-updated version of gorm.Model
// It will be managed automatically by gorm
type Model struct {
	ID        uint           `json:"-" gorm:"primarykey"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
