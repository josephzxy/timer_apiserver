package mysql

import (
	"go.uber.org/zap"
	"gorm.io/gorm"

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

	mysqlErr := GetMySQLErr(err)

	var appErrCode pkgerr.AppErrCode
	switch mysqlErr {
	case DUPLICATE_ENTRY:
		appErrCode = pkgerr.ErrTimerAlreadyExists
	default:
		appErrCode = pkgerr.ErrUnknown
	}
	return pkgerr.New(appErrCode, err.Error())
}
