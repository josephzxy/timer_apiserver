package timer

import (
	"errors"

	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/go-sql-driver/mysql"

	pkgerr "github.com/josephzxy/timer_apiserver/internal/pkg/err"
	"github.com/josephzxy/timer_apiserver/internal/resource/v1/model"
)

var dbDeleteByNameFunc = func(db *gorm.DB, name string) error {
	return db.Transaction(func(tx *gorm.DB) error {
		var timer model.Timer
		// SELECT * FROM timer WHERE name = ? AND deleted_at IS NULL LIMIT 1;
		if err := tx.Where("name = ?", name).Take(&timer).Error; err != nil {
			return err
		}
		// UPDATE timer SET delete_at = NOW() WHERE name = ?;
		if err := tx.Where("name = ?", name).Delete(&timer).Error; err != nil {
			return err
		}
		if err := tx.Exec("UPDATE `timer` SET `alive`=NULL WHERE `id`= ?", timer.ID).Error; err != nil {
			return err
		}

		return nil
	})
}

// DeleteByName deleted a timer by the given name.
func (s *timerStore) DeleteByName(name string) error {
	err := dbDeleteByNameFunc(s.db, name)
	if err == nil {
		return nil
	}
	zap.S().Errorw("failed to delete timer", "err", err, "name", name)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return pkgerr.New(pkgerr.ErrTimerNotFound, "")
	}
	me, ok := err.(*mysql.MySQLError)
	if !ok {
		return err
	}

	return pkgerr.New(pkgerr.ErrDatabase, me.Error())
}
