package sum

import (
	"go.uber.org/zap"

)

func NewLoggerWithouTrace()  *zap.SugaredLogger {
	cfg := zap.NewProductionConfig()
	logger, _ := cfg.Build()
	return logger.Sugar()
}