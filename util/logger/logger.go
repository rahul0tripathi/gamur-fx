package logger

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewLogger() *zap.SugaredLogger {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	sugar := logger.Sugar()
	return sugar
}

var Module = fx.Provide(func() *zap.SugaredLogger{
	return NewLogger()
})