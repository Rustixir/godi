package godi

import (
	"go.uber.org/fx"
)

type iocContainer struct {
	modules []fx.Option
	invoker []fx.Option
	onStart Hook
	onStop  Hook
}

var container iocContainer

func init() {
	container = iocContainer{
		modules: make([]fx.Option, 0),
		invoker: make([]fx.Option, 0),
		onStart: nopHook,
		onStop:  nopHook,
	}
}

func AddService(name string, constructor interface{}, dependencies ...Option) {
	opts := make([]Option, 0, len(dependencies))
	for _, dependency := range dependencies {
		opts = append(opts, dependency)
	}
	opts = append(opts, fx.Provide(constructor))
	container.modules = append(
		container.modules,
		fx.Module(name, opts...),
	)
}

func AddIService(name string, constructor interface{}, Interface interface{}) {
	container.modules = append(
		container.modules,
		fx.Module(name, fx.Provide(
			fx.Annotate(
				constructor,
				fx.As(Interface),
			),
		)),
	)
}

func OnAfterStart(funcs ...interface{}) {
	opt := fx.Invoke(funcs...)
	container.invoker = append(container.invoker, opt)
}

func AddHook(OnStart Hook, OnStop Hook) {
	container.onStart = OnStart
	container.onStop = OnStop
}

func Run() {
	if len(container.invoker) > 0 {
		container.modules = append(container.modules, container.invoker...)
	}
	container.modules = append(container.modules, fx.Invoke(func(lifecycle fx.Lifecycle) {
		lifecycle.Append(fx.Hook{
			OnStart: container.onStart,
			OnStop:  container.onStop,
		})
	}))
	fx.New(container.modules...).Run()
}
