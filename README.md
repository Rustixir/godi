# godi

`godi` is a Go package that extends the Uber Fx framework, providing a simpler and more concise way to manage dependency injection in Go applications. It introduces syntactic sugar to make your code cleaner and more maintainable.

## Features

- **Service Registration**: Easily register services with dependencies.
- **Interface Binding**: Bind implementations to interfaces for flexible dependency management.
- **Lifecycle Hooks**: Run functions after the application starts.
- **Simplified API**: Use a straightforward API to manage your application's dependencies.

## Installation

To install `godi`, use the following command:

```bash
go get github.com/Rustixir/godi
```

## Getting Started

Here's a quick example to get you started with `godi`.

### Define Your Services

```go
package main

import (
	"log/slog"
	"github.com/Rustixir/godi"
)

func NewServiceA(logger *slog.Logger) serviceA {
	return serviceA{logger: logger}
}

type serviceA struct {
	logger *slog.Logger
}

func (s *serviceA) Print() {
	s.logger.Info("serviceA Print: Hello")
}

func NewServiceB(logger *slog.Logger) serviceB {
	return serviceB{logger: logger}
}

type serviceB struct {
	logger *slog.Logger
}

func (s *serviceB) Print() {
	s.logger.Info("serviceB Print: Hello")
}

func NewServiceC(a serviceA, b serviceB, logger *slog.Logger) serviceC {
	return serviceC{ServiceA: a, ServiceB: b, logger: logger}
}

type serviceC struct {
	ServiceA serviceA
	ServiceB serviceB
	logger   *slog.Logger
}

func (c serviceC) Run() {
	c.logger.Info("Works fine")
	c.ServiceA.Print()
	c.ServiceB.Print()
}
```

### Register and Run Services

```go
func main() {
	godi.AddSlog()
	godi.AddService("serviceA", NewServiceA)
	godi.AddService("serviceB", NewServiceB)
	godi.AddService("serviceC", NewServiceC)

	godi.OnAfterStart(func(service serviceC) error {
		service.Run()
		return nil
	})

	godi.Run()
}
```

## API Reference

- `AddService(name string, constructor interface{}, dependencies ...interface{})`: Register a service with optional dependencies.
- `AddIService(name string, constructor interface{}, Interface interface{})`: Register a service and bind it to an interface.
- `OnAfterStart(funcs ...interface{})`: Register functions to execute after the application starts.
- `Run()`: Start the application and execute all registered lifecycle hooks.



## License

This project is licensed under the Apache License 2.0.
