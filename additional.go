package godi

import (
	"go.uber.org/fx"
	"log/slog"
)

func AddSlogWithHandler(handler slog.Handler) {
	AddService("slog-logger", slog.New, fx.Supply(handler))
}

func AddSlog() {
	AddService("slog-logger", slog.Default)
}
