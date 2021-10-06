package model

import (
	"time"
)

// model is the slightly-updated version of gorm.Model
// It will be managed automatically by gorm
type model struct {
	ID        uint      `json:"-" gorm:"primarykey"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"-"`
}
