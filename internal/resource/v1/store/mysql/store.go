package mysql

import (
	"fmt"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/josephzxy/timer_apiserver/internal/resource/v1/store"
)

type MySQLStoreRouter struct {
	db *gorm.DB
}

func (r *MySQLStoreRouter) Timer() store.TimerStore {
	return &MySQLTimerStore{r.db}
}

func NewStoreRouter(cfg *Config) (*MySQLStoreRouter, error) {
	dsn := fmt.Sprintf(
		`%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t&loc=%s`,
		cfg.User,
		cfg.Pwd,
		cfg.Host,
		cfg.Port,
		cfg.Database,
		cfg.Charset,
		cfg.ParseTime,
		cfg.Loc,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		zap.S().Errorw("failed to create mysql connection", "err", err)
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		zap.S().Errorw("failed to configure mysql connection", "err", err)
		return nil, err
	}

	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(cfg.MaxConnLifetime)

	return &MySQLStoreRouter{db}, nil
}
