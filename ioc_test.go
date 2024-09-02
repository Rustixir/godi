package godi

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log/slog"
	"testing"
)

func NewServiceA(logger *slog.Logger) serviceA {
	return serviceA{
		logger: logger,
	}
}

type serviceA struct {
	logger *slog.Logger
}

func (s *serviceA) Print() {
	s.logger.Info("serviceA Print: Hello")
}

// -------------------------------------------------------

func NewServiceB(logger *zap.Logger) serviceB {
	return serviceB{
		logger: logger,
	}
}

type serviceB struct {
	logger *zap.Logger
}

func (s *serviceB) Print() {
	s.logger.Info("serviceA Print: Hello")
}

// -------------------------------------------------------

func NewServiceC(a serviceA, b serviceB, logger *slog.Logger) serviceC {
	return serviceC{
		ServiceA: a,
		ServiceB: b,
		logger:   logger,
	}
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

func TestIoc(t *testing.T) {

	AddSlog()
	AddZapDevelopment(
		zap.WithCaller(true),
		zap.WithClock(zapcore.DefaultClock),
	)
	AddService("serviceA", NewServiceA)
	AddService("serviceB", NewServiceB)
	AddService("serviceC", NewServiceC)
	OnAfterStart(func(service serviceC) error {
		service.Run()
		return nil
	})
	Run()
}
