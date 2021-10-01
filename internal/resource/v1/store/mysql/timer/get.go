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
