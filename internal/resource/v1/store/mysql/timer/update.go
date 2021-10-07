package timer

import (
	"errors"

	"github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
	"gorm.io/gorm"

	pkgerr "github.com/josephzxy/timer_apiserver/internal/pkg/err"
	"github.com/josephzxy/timer_apiserver/internal/resource/v1/model"
)

var dbUpdateByNameFunc = func(db *gorm.DB, name string, want *model.TimerCore) error {
	return db.Transaction(func(tx *gorm.DB) error {
		var timer model.Timer
		// SELECT * FROM timer WHERE name = ? AND deleted_at IS NULL LIMIT 1;
		if err := tx.Where("name = ?", name).Take(&timer).Error; err != nil {
			return err
		}
		// https://gorm.io/docs/update.html#Updates-multiple-columns
		// UPDATE timer SET name = ? AND trigger_at = ? WHERE name = ?;
		if err := tx.Model(&timer).Where("name = ?", name).Updates(
			map[string]interface{}{"name": want.Name, "trigger_at": want.TriggerAt},
		).Error; err != nil {
			return err
		}
		return nil
	})
}

// UpdateByName updates a timer by the given name with the given desired state.
func (s *timerStore) UpdateByName(name string, want *model.TimerCore) error {
	err := dbUpdateByNameFunc(s.db, name, want)
	if err == nil {
		return nil
	}
	zap.S().Errorw("failed to update timer", "err", err, "data", want)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return pkgerr.New(pkgerr.ErrTimerNotFound, "")
	}
	if me, ok := err.(*mysql.MySQLError); ok {
		if me.Number == 1062 {
			return pkgerr.New(pkgerr.ErrTimerAlreadyExists, "")
		}
	}
	return err
}
