package godi

import (
	"go.uber.org/fx"
)

type iocContainer struct {
	modules []fx.Option
	invoker *fx.Option
}

var container iocContainer

func init() {
	container = iocContainer{
		modules: make([]fx.Option, 0),
		invoker: nil,
	}
}

func AddService(name string, constructor interface{}, dependencies ...interface{}) {
	opts := make([]fx.Option, len(dependencies))
	for _, dependency := range dependencies {
		opts = append(opts, fx.Provide(dependency))
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
	container.invoker = &opt
}

func Run() {
	if container.invoker != nil {
		container.modules = append(container.modules, *container.invoker)
	}
	fx.New(container.modules...).Run()
}
