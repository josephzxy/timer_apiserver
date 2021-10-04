package app

import (
	"fmt"

	"go.uber.org/zap"

	"github.com/pkg/errors"
	"github.com/spf13/viper"

	"github.com/josephzxy/timer_apiserver/internal/app/config"
	"github.com/josephzxy/timer_apiserver/internal/resource/v1/service"
	"github.com/josephzxy/timer_apiserver/internal/resource/v1/store/mysql"
	"github.com/josephzxy/timer_apiserver/internal/restserver"
)

type App struct {
	cfg *config.Config
}

func NewApp() *App {
	a := &App{
		cfg: config.NewEmptyConfig(),
	}
	if err := a.loadConfig(); err != nil {
		zap.S().Fatalw("failed to load config for app", "err", err)
	}
	return a
}

func (a *App) loadConfig() error {
	cfgFile := config.CfgFile()
	if cfgFile == "" {
		return fmt.Errorf("config file path must not be empty")
	}
	viper.SetConfigFile(cfgFile)

	if err := viper.ReadInConfig(); err != nil {
		msg := "failed to read config file"
		zap.S().Errorw(msg, "err", err)
		return errors.WithMessage(err, msg)
	}

	if err := viper.Unmarshal(a.cfg); err != nil {
		msg := "failed to unmarshal config"
		zap.S().Errorw(msg, "err", err)
		return errors.WithMessage(err, msg)
	}
	return nil
}

func (a *App) Run() {
	fmt.Println("Running...")

	mysqlStoreRouter, err := mysql.NewStoreRouter(&mysql.Config{
		User:            a.cfg.MySQL.User,
		Pwd:             a.cfg.MySQL.Pwd,
		Host:            a.cfg.MySQL.Host,
		Port:            a.cfg.MySQL.Port,
		Database:        a.cfg.MySQL.Database,
		Charset:         a.cfg.MySQL.Charset,
		ParseTime:       a.cfg.MySQL.ParseTime,
		Loc:             a.cfg.MySQL.Loc,
		MaxIdleConns:    a.cfg.MySQL.MaxIdleConns,
		MaxOpenConns:    a.cfg.MySQL.MaxOpenConns,
		MaxConnLifetime: a.cfg.MySQL.MaxConnLifetime,
		LogLevel:        a.cfg.MySQL.LogLevel, // silent
	})

	if err != nil {
		zap.S().Fatalw("failed to get mysql store router", "err", err)
	}
	serviceRouter := service.NewRouter(mysqlStoreRouter)

	restServer := restserver.New(
		&restserver.Config{
			InsecureServing: restserver.InsecureServingConfig{
				Host: a.cfg.RESTServer.InsecureServing.Host,
				Port: a.cfg.RESTServer.InsecureServing.Port,
			},
			Mode: a.cfg.RESTServer.Mode,
		},
		serviceRouter,
	)

	if err := restServer.Start(); err != nil {
		zap.S().Fatalw("error occured during running rest server", "err", err)
	}
}
