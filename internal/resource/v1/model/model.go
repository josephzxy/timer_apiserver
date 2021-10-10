// Package model defines data models used in the app
package model

import (
	"time"
)

// Model is the slightly-updated version of gorm.Model
// It will be managed automatically by gorm.
type Model struct {
	ID        uint      `json:"-" gorm:"primarykey"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"-"`
}
