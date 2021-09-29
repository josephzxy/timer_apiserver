package mysql

import (
	"go.uber.org/zap"
	"gorm.io/gorm"

	model "github.com/josephzxy/timer_apiserver/internal/resource/model/v1"
)

type MySQLTimerStore struct {
	db *gorm.DB
}

func (s *MySQLTimerStore) Create(timer *model.Timer) error {
	err := s.db.Create(&timer).Error
	if err != nil {
		zap.S().Errorw("failed to create timer", "err", err, "data", timer)
		return err
	}
	return nil
}
