package timer

import (
	"errors"

	"go.uber.org/zap"
	"gorm.io/gorm"

	pkgerr "github.com/josephzxy/timer_apiserver/internal/pkg/err"
	"github.com/josephzxy/timer_apiserver/internal/resource/v1/model"
)

var dbDeleteByNameFunc = func(db *gorm.DB, name string) error {
	return db.Transaction(func(tx *gorm.DB) error {
		var timer model.Timer
		if err := tx.Where("name = ?", name).First(&timer).Error; err != nil {
			return err
		}
		if err := tx.Where("name = ?", name).Delete(&timer).Error; err != nil {
			return err
		}
		if err := tx.Exec("UPDATE `timer` SET `alive`=NULL WHERE `id`= ?", timer.ID).Error; err != nil {
			return err
		}
		return nil
	})
}

func (s *TimerStore) DeleteByName(name string) error {
	err := dbDeleteByNameFunc(s.DB, name)
	if err == nil {
		return nil
	}
	zap.S().Errorw("failed to delete timer", "err", err)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return pkgerr.New(pkgerr.ErrTimerNotFound, "")
	}
	return err
}
