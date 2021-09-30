package mysql

import (
	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/go-sql-driver/mysql"

	pkgerr "github.com/josephzxy/timer_apiserver/internal/pkg/err"
	model "github.com/josephzxy/timer_apiserver/internal/resource/model/v1"
)

type MySQLTimerStore struct {
	db *gorm.DB
}

func (s *MySQLTimerStore) Create(timer *model.Timer) error {
	err := s.db.Create(&timer).Error
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
