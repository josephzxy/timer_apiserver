package timer

import (
	"github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
	"gorm.io/gorm"

	pkgerr "github.com/josephzxy/timer_apiserver/internal/pkg/err"
	"github.com/josephzxy/timer_apiserver/internal/resource/v1/model"
)

var dbCreateFunc = func(db *gorm.DB, timer *model.Timer) error {
	timer.Alive = true
	// INSERT INTO timer (created_at, updated_at, alive, name, trigger_at)
	// VALUES (NOW(), NOW(), true, ?, ?);
	return db.Create(timer).Error
}

// Create creates a new timer.
func (s *timerStore) Create(timer *model.Timer) error {
	err := dbCreateFunc(s.db, timer)
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
