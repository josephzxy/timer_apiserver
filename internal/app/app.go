package app

import (
	"fmt"
	"time"

	"go.uber.org/zap"

	"github.com/josephzxy/timer_apiserver/internal/resource/service"
	"github.com/josephzxy/timer_apiserver/internal/resource/store/mysql"
	"github.com/josephzxy/timer_apiserver/internal/restserver"
)

type App struct{}

func NewApp() *App {
	return &App{}
}

func (a *App) Run() {
	fmt.Println("Running...")

	mysqlStoreRtr, err := mysql.NewStoreRouter(&mysql.Config{
		User:            "root",
		Pwd:             "root",
		Host:            "localhost",
		Port:            33306,
		Database:        "test",
		Charset:         "utf8mb4",
		ParseTime:       true,
		Loc:             "Local",
		MaxIdleConns:    100,
		MaxOpenConns:    100,
		MaxConnLifetime: 10 * time.Second,
		LogLevel:        1, // silent
	})
	if err != nil {
		zap.S().Fatalw("failed to get mysql store router", "err", err)
	}
	srvRtr := service.NewRouter(mysqlStoreRtr)
	restSrvrCfg := &restserver.Config{
		InsecureServing: restserver.InsecureServingConfig{
			Host: "localhost",
			Port: 8080,
		},
		Mode: "debug",
	}
	restSrvr := restserver.New(restSrvrCfg, srvRtr)

	if err := restSrvr.Start(); err != nil {
		zap.S().Fatalw("error occured during running rest server", "err", err)
	}
}
