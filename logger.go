package gokit

import (
	"go.uber.org/zap"
	"sync"
)

var (
	once sync.Once
)

type (
	CustomLogger struct{}
)

func InitLogger(log *zap.Logger) {
	once.Do(func() {
		zap.ReplaceGlobals(log)
	})
}

// print
func (c *CustomLogger) Print(args ...interface{}) {
	zap.S().Info(args...)
}

// printf
func (c *CustomLogger) Printf(format string, v ...interface{}) {
	zap.S().Infof(format, v)
}

// Debug ZapLogger
func (c *CustomLogger) Debug(args ...interface{}) {
	zap.S().Debug(args...)
}

// Info logger
func (c *CustomLogger) Info(args ...interface{}) {
	zap.S().Info(args...)
}

// Warn logger
func (c *CustomLogger) Warn(args ...interface{}) {
	zap.S().Warn(args...)
}

// Error logger
func (c *CustomLogger) Error(args ...interface{}) {
	zap.S().Error(args...)
}

func (c *CustomLogger) Fatal(args ...interface{}) {
	zap.S().Fatal(args...)
}

func (c *CustomLogger) Debugf(format string, args ...interface{}) {
	zap.S().Debugf(format, args...)
}

func (c *CustomLogger) Infof(format string, args ...interface{}) {
	zap.S().Infof(format, args...)
}

func (c *CustomLogger) Warnf(format string, args ...interface{}) {
	zap.S().Warnf(format, args...)
}

func (c *CustomLogger) Errorf(format string, args ...interface{}) {
	zap.S().Errorf(format, args...)
}

func (c *CustomLogger) Fatalf(format string, args ...interface{}) {
	zap.S().Fatalf(format, args...)
}

func (c *CustomLogger) Panicf(format string, args ...interface{}) {
	zap.S().Panicf(format, args...)
}

func (c *CustomLogger) DebugFields(msg string, fields ...zap.Field) {
	zap.L().Debug(msg, fields...)
}

func (c *CustomLogger) InfoFields(msg string, fields ...zap.Field) {
	zap.L().Info(c.Prefix+msg, fields...)
}

func (c *CustomLogger) WarnFields(msg string, fields ...zap.Field) {
	zap.L().Warn(c.Prefix+msg, fields...)
}

func (c *CustomLogger) ErrorFields(msg string, fields ...zap.Field) {
	zap.L().Error(c.Prefix+msg, fields...)
}

func (c *CustomLogger) PanicFields(msg string, fields ...zap.Field) {
	zap.L().Panic(c.Prefix+msg, fields...)
}

func (c *CustomLogger) FatalFields(msg string, fields ...zap.Field) {
	zap.L().Fatal(c.Prefix+msg, fields...)
}
