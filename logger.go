package gokit

import (
	"go.uber.org/zap"
	"sync"
)

var (
	once sync.Once
)

type (
	CustomLogger struct {
		Prefix string
	}
)

func InitLogger(log *zap.Logger) {
	once.Do(func() {
		zap.ReplaceGlobals(log)
	})
}

// new logger
func NewCustomLogger(prefix string) *CustomLogger {
	return &CustomLogger{Prefix: "[" + prefix + "] "}
}

func Append(prefix string, args ...interface{}) []interface{} {
	return append([]interface{}{prefix}, args...)
}

// print
func (c *CustomLogger) Print(args ...interface{}) {
	zap.S().Info(Append(c.Prefix, args...)...)
}

// printf
func (c *CustomLogger) Printf(format string, v ...interface{}) {
	zap.S().Infof(c.Prefix+format, v)
}

// Debug ZapLogger
func (c *CustomLogger) Debug(args ...interface{}) {
	zap.S().Debug(Append(c.Prefix, args...)...)
}

// Info logger
func (c *CustomLogger) Info(args ...interface{}) {
	zap.S().Info(Append(c.Prefix, args...)...)
}

// Warn logger
func (c *CustomLogger) Warn(args ...interface{}) {
	zap.S().Warn(Append(c.Prefix, args...)...)
}

// Error logger
func (c *CustomLogger) Error(args ...interface{}) {
	zap.S().Error(Append(c.Prefix, args...)...)
}

func (c *CustomLogger) Fatal(args ...interface{}) {
	zap.S().Fatal(Append(c.Prefix, args...)...)
}

func (c *CustomLogger) Debugf(format string, args ...interface{}) {
	zap.S().Debugf(c.Prefix+format, args...)
}

func (c *CustomLogger) Infof(format string, args ...interface{}) {
	zap.S().Infof(c.Prefix+format, args...)
}

func (c *CustomLogger) Warnf(format string, args ...interface{}) {
	zap.S().Warnf(c.Prefix+format, args...)
}

func (c *CustomLogger) Errorf(format string, args ...interface{}) {
	zap.S().Errorf(c.Prefix+format, args...)
}

func (c *CustomLogger) Fatalf(format string, args ...interface{}) {
	zap.S().Fatalf(c.Prefix+format, args...)
}

func (c *CustomLogger) Panicf(format string, args ...interface{}) {
	zap.S().Panicf(c.Prefix+format, args...)
}

func (c *CustomLogger) DebugFields(msg string, fields ...zap.Field) {
	zap.L().Debug(c.Prefix+msg, fields...)
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
