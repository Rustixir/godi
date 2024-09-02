package godi

import (
	"go.uber.org/zap"
	"log/slog"
)

func AddSlog() {
	AddService("slog-logger", slog.Default)
}

func AddSlogWithHandler(handler slog.Handler) {
	AddService("slog-logger", func() *slog.Logger {
		return slog.New(handler)
	})
}

func AddZapProduction(options ...zap.Option) {
	AddService("zap-logger-prod", func() (*zap.Logger, error) {
		return zap.NewProduction(options...)
	})
}

func AddZapDevelopment(options ...zap.Option) {
	AddService("zap-logger-dev", func() (*zap.Logger, error) {
		return zap.NewDevelopment(options...)
	})
}
