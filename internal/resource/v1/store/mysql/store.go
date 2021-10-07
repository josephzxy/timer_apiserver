// Package mysql provides MySQL-specific implementations of interface store.StoreRouter
package mysql

import (
	"fmt"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/josephzxy/timer_apiserver/internal/resource/v1/store"
	"github.com/josephzxy/timer_apiserver/internal/resource/v1/store/mysql/timer"
)

type mysqlStoreRouter struct {
	db *gorm.DB
}

// Timer routes to a concrete MySQL-dedicated value of interface store.TimerStore.
func (r *mysqlStoreRouter) Timer() store.TimerStore {
	return timer.NewTimerStore(r.db)
}

// NewStoreRouter creates and configures a MySQL session and returns
// a mysqlStoreRouter with it.
func NewStoreRouter(cfg *Config) (store.StoreRouter, error) {
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
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.LogLevel(cfg.LogLevel)),
	})
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

	return &mysqlStoreRouter{db}, nil
}
