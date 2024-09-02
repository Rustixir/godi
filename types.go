package godi

import (
	"context"
	"go.uber.org/fx"
)

type In = fx.In

type Out = fx.Out

type Lifecycle = fx.Lifecycle

type Option = fx.Option

type Hook func(context.Context) error

func nopHook(context.Context) error { return nil }

func Supply(values ...interface{}) Option {
	return fx.Supply(values...)
}

func Provide(constructors ...interface{}) Option {
	return fx.Provide(constructors...)
}
