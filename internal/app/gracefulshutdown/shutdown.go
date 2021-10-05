package gracefulshutdown

import (
	"os"
	"os/signal"
	"sync"
	"syscall"

	"go.uber.org/zap"
)

var posixShutdownSignals = []os.Signal{os.Interrupt, syscall.SIGTERM}

func enable(handler func() error) {
	lis := make(chan os.Signal, 2)
	signal.Notify(lis, posixShutdownSignals...)

	go func() {
		<-lis
		go func() {
			zap.L().Info("starting graceful shutdown")
			if err := handler(); err != nil {
				zap.S().Fatalw("error occurred during graceful shutdown, server exiting", "err", err)
			}
			zap.L().Info("server exiting after graceful shutdown")
			os.Exit(0)
		}()
		<-lis
		zap.L().Fatal("forced shutdown by 2 shutdown OS signals")
	}()
}

var (
	once sync.Once
)

func Enable(handler func() error) {
	did := false
	once.Do(func() {
		enable(handler)
		did = true
	})
	if !did {
		zap.L().Warn("attempt to enable graceful shutdown more than once, which will be ineffective")
	}
}
