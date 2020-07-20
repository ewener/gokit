package gokit

import (
	"go.uber.org/zap"
	"sync"
)

var (
	defaultLogger = &ZapLogger{logger: zap.L(), sugarLogger: zap.S()}
	once          sync.Once
)

type (
	CustomLogger struct {
		Prefix string
		*ZapLogger
	}
	ZapLogger struct {
		logger      *zap.Logger
		sugarLogger *zap.SugaredLogger
	}
)

func InitLogger(log *zap.Logger) {
	once.Do(func() {
		zap.ReplaceGlobals(log)
	})
}

// new logger
func NewCustomLogger(prefix string) *CustomLogger {
	return &CustomLogger{Prefix: "[" + prefix + "] ", ZapLogger: defaultLogger}
}

// print
func (c *CustomLogger) Print(args ...interface{}) {
	c.Info(args...)
}

// printf
func (c *CustomLogger) Printf(format string, v ...interface{}) {
	c.Infof(format, v)
}

// Debug ZapLogger
func (l *ZapLogger) Debug(args ...interface{}) {
	l.sugarLogger.Debug(args...)
}

// Info logger
func (l *ZapLogger) Info(args ...interface{}) {
	l.sugarLogger.Info(args...)
}

// Warn logger
func (l *ZapLogger) Warn(args ...interface{}) {
	l.sugarLogger.Warn(args...)
}

// Error logger
func (l *ZapLogger) Error(args ...interface{}) {
	l.sugarLogger.Error(args...)
}

func (l *ZapLogger) Fatal(args ...interface{}) {
	l.sugarLogger.Fatal(args...)
}

func (l *ZapLogger) Debugf(format string, args ...interface{}) {
	l.sugarLogger.Debugf(format, args...)
}

func (l *ZapLogger) Infof(format string, args ...interface{}) {
	l.sugarLogger.Infof(format, args...)
}

func (l *ZapLogger) Warnf(format string, args ...interface{}) {
	l.sugarLogger.Warnf(format, args...)
}

func (l *ZapLogger) Errorf(format string, args ...interface{}) {
	l.sugarLogger.Errorf(format, args...)
}

func (l *ZapLogger) Fatalf(format string, args ...interface{}) {
	l.sugarLogger.Fatalf(format, args...)
}

func (l *ZapLogger) Panicf(format string, args ...interface{}) {
	l.sugarLogger.Panicf(format, args...)
}

func (l *ZapLogger) DebugFields(msg string, fields ...zap.Field) {
	l.logger.Debug(msg, fields...)
}

func (l *ZapLogger) InfoFields(msg string, fields ...zap.Field) {
	l.logger.Info(msg, fields...)
}

func (l *ZapLogger) WarnFields(msg string, fields ...zap.Field) {
	l.logger.Warn(msg, fields...)
}

func (l *ZapLogger) ErrorFields(msg string, fields ...zap.Field) {
	l.logger.Error(msg, fields...)
}

func (l *ZapLogger) PanicFields(msg string, fields ...zap.Field) {
	l.logger.Panic(msg, fields...)
}

func (l *ZapLogger) FatalFields(msg string, fields ...zap.Field) {
	l.logger.Fatal(msg, fields...)
}
