package timer

import (
	"errors"

	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/go-sql-driver/mysql"

	pkgerr "github.com/josephzxy/timer_apiserver/internal/pkg/err"
	"github.com/josephzxy/timer_apiserver/internal/resource/v1/model"
)

type TimerStore struct {
	DB *gorm.DB
}

var dbCreateFunc = func(db *gorm.DB, value interface{}) error {
	return db.Create(value).Error
}

func (s *TimerStore) Create(timer *model.Timer) error {
	err := dbCreateFunc(s.DB, timer)
	if err == nil {
		return nil
	}

	zap.S().Errorw("failed to create timer", "err", err, "data", timer)
	me, ok := err.(*mysql.MySQLError)
	if !ok {
		return err
	}

	if me.Number == 1062 {
		return pkgerr.New(pkgerr.ErrTimerAlreadyExists, me.Error())
	}
	return err
}

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
