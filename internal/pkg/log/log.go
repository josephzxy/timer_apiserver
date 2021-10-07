// Package log provides a simple encapsulation of the logging library "go.uber.org/zap"
package log

import (
	"log"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func init() {
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)
	l, err := cfg.Build()
	if err != nil {
		log.Fatalf("failed to get zap production logger: %s", err)
	}
	zap.ReplaceGlobals(l)
}

// Flush flushes any bufferred log entries.
// It should be called before the program exits
func Flush() error {
	zap.L().Info("flushing log")
	return zap.L().Sync()
}
