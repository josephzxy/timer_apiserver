package timer

import (
	"errors"

	"gorm.io/gorm"

	pkgerr "github.com/josephzxy/timer_apiserver/internal/pkg/err"
	"github.com/josephzxy/timer_apiserver/internal/resource/v1/model"
)

var dbGetByNameFunc = func(db *gorm.DB, name string, timer *model.Timer) error {
	return db.Where("name = ?", name).First(timer).Error
}

func (s *TimerStore) GetByName(name string) (*model.Timer, error) {
	var timer model.Timer
	err := dbGetByNameFunc(s.DB, name, &timer)
	if err == nil {
		return &timer, nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, pkgerr.New(pkgerr.ErrTimerNotFound, "")
	}
	return nil, err
}

var dbGetAllFunc = func(db *gorm.DB, timers *[]model.Timer) error {
	return db.Find(timers).Error
}

func (s *TimerStore) GetAll() ([]model.Timer, error) {
	var timers []model.Timer
	err := dbGetAllFunc(s.DB, &timers)
	if err == nil {
		return timers, nil
	}
	return nil, err
}

var dbGetAllPendingFunc = func(db *gorm.DB, timers *[]model.Timer) error {
	// SELECT * FROM timer WHERE alive = true AND trigger_at > NOW();
	return db.Where("alive = ? AND trigger_at > NOW()", true).Find(timers).Error
}

// TODO: test required
func (s *TimerStore) GetAllPending() ([]model.Timer, error) {
	var timers []model.Timer
	err := dbGetAllPendingFunc(s.DB, &timers)
	if err == nil {
		return timers, nil
	}
	return nil, err
}
