package timer

import (
	"errors"

	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/go-sql-driver/mysql"

	pkgerr "github.com/josephzxy/timer_apiserver/internal/pkg/err"
	"github.com/josephzxy/timer_apiserver/internal/resource/v1/model"
)

var dbGetByNameFunc = func(db *gorm.DB, name string, timer *model.Timer) error {
	// SELECT * FROM timer WHERE name = ? AND deleted_at IS NULL LIMIT 1;
	return db.Where("name = ?", name).Take(timer).Error
}

// GetByName gets a timer by the given name.
func (s *timerStore) GetByName(name string) (*model.Timer, error) {
	var timer model.Timer
	err := dbGetByNameFunc(s.db, name, &timer)
	if err == nil {
		return &timer, nil
	}
	zap.S().Errorw("failed to get timer by name", "err", err, "name", name)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, pkgerr.New(pkgerr.ErrTimerNotFound, "")
	}

	me, ok := err.(*mysql.MySQLError)
	if !ok {
		return nil, err
	}
	return nil, pkgerr.New(pkgerr.ErrDatabase, me.Error())
}

var dbGetAllFunc = func(db *gorm.DB, timers *[]model.Timer) error {
	// SELECT * FROM timer WHERE deleted_at IS NULL;
	return db.Find(timers).Error
}

// GetAll gets all timers(deleted timers excluded).
func (s *timerStore) GetAll() ([]model.Timer, error) {
	var timers []model.Timer
	err := dbGetAllFunc(s.db, &timers)
	if err == nil {
		return timers, nil
	}
	zap.S().Errorw("failed to get all timers", "err", err)
	me, ok := err.(*mysql.MySQLError)
	if !ok {
		return nil, err
	}
	return nil, pkgerr.New(pkgerr.ErrDatabase, me.Error())
}

var dbGetAllPendingFunc = func(db *gorm.DB, timers *[]model.Timer) error {
	// SELECT * FrOM timer WHERE alive = true AND trigger_at > NOW() AND deleted_at IS NULL;
	return db.Where("alive = ? AND trigger_at > NOW()", true).Find(timers).Error
}

// GetAllPending gets all pending timers.
// A timer is "pending" if it is not deleted and not triggerred yet.
func (s *timerStore) GetAllPending() ([]model.Timer, error) {
	var timers []model.Timer
	err := dbGetAllPendingFunc(s.db, &timers)
	if err == nil {
		return timers, nil
	}
	zap.S().Errorw("failed to get all pending timers", "err", err)
	me, ok := err.(*mysql.MySQLError)
	if !ok {
		return nil, err
	}
	return nil, pkgerr.New(pkgerr.ErrDatabase, me.Error())
}
