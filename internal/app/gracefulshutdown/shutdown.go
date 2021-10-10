// Package gracefulshutdown implements a simple library
// for enabling app-level graceful shutdown on receiving shutdown OS
// signals like SIGTERM and SIGINT.
package gracefulshutdown

import (
	"os"
	"os/signal"
	"sync"
	"syscall"

	"go.uber.org/zap"

	"github.com/josephzxy/timer_apiserver/internal/pkg/log"
)

var posixShutdownSignals = []os.Signal{os.Interrupt, syscall.SIGTERM}

func enable(handler func() error) {
	lis := make(chan os.Signal, 2)
	signal.Notify(lis, posixShutdownSignals...)

	go func() {
		defer log.Flush()
		<-lis
		go func() {
			defer log.Flush()
			zap.L().Info("graceful shutdown started")
			if err := handler(); err != nil {
				zap.S().Panicw("graceful shutdown failed, server exiting", "err", err)
			}
			zap.L().Info("graceful shutdown done, server exiting")
			os.Exit(0)
		}()
		<-lis
		zap.L().Panic("server exits by force by 2 shutdown OS signals")
	}()
}

var once sync.Once

// Enable enables app-level graceful shutdown and invokes the given
// handler once graceful shutdown is triggerred. Note that Enable can ONLY
// be called once. All following calls will be ineffective.
func Enable(handler func() error) {
	did := false
	once.Do(func() {
		enable(handler)
		did = true
	})
	if !did {
		zap.L().Warn("graceful shutdown was attempted to be enabled more than once, which will be ineffective")
	}
}
