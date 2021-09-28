package log

import (
	"log"

	"go.uber.org/zap"
)

func init() {
	l, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("failed to get zap production logger: %s", err)
	}
	zap.ReplaceGlobals(l)
}

func Flush() {
	zap.L().Sync()
}
