package gokit

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"sync"
)

var (
	defaultLogger *ZapLogger
	once          sync.Once
)

type (
	CustomLogger struct {
		prefix string
	}
	ZapLogger struct {
		underlieLogger *zap.SugaredLogger
	}
)

func InitLogger(log *zap.Logger) {
	once.Do(func() {
		defaultLogger = &ZapLogger{underlieLogger: log.Sugar()}
	})
}

// new logger
func NewCustomLogger(prefix string) *CustomLogger {
	return &CustomLogger{prefix: "[" + prefix + "] "}
}

// print
func (c *CustomLogger) Print(args ...interface{}) {
	c.Info(args...)
}

// printf
func (c *CustomLogger) Printf(format string, v ...interface{}) {
	c.Infof(format, v)
}

// Debug logger
func (c *CustomLogger) Debug(args ...interface{}) {
	defaultLogger.Debug(append([]interface{}{c.prefix}, args...)...)
}

// Info logger
func (c *CustomLogger) Info(args ...interface{}) {
	defaultLogger.Info(append([]interface{}{c.prefix}, args...)...)
}

// Warn logger
func (c *CustomLogger) Warn(args ...interface{}) {
	defaultLogger.Warn(append([]interface{}{c.prefix}, args...)...)
}

// Error logger
func (c *CustomLogger) Error(args ...interface{}) {
	defaultLogger.Error(append([]interface{}{c.prefix}, args...)...)
}

func (c *CustomLogger) Fatal(args ...interface{}) {
	defaultLogger.Fatal(append([]interface{}{c.prefix}, args...)...)
}

func (c *CustomLogger) Debugf(format string, args ...interface{}) {
	defaultLogger.Debugf(c.prefix+format, args...)
}

func (c *CustomLogger) Infof(format string, args ...interface{}) {
	defaultLogger.Infof(c.prefix+format, args...)
}

func (c *CustomLogger) Warnf(format string, args ...interface{}) {
	defaultLogger.Warnf(c.prefix+format, args...)
}

func (c *CustomLogger) Errorf(format string, args ...interface{}) {
	defaultLogger.Errorf(c.prefix+format, args...)
}

func (c *CustomLogger) Fatalf(format string, args ...interface{}) {
	defaultLogger.Fatalf(c.prefix+format, args...)
}

func (c *CustomLogger) Panicf(format string, args ...interface{}) {
	defaultLogger.Panicf(c.prefix+format, args...)
}

func (c *CustomLogger) DebugFields(msg string, fields ...zapcore.Field) {
	defaultLogger.DebugWithFields(c.prefix+msg, fields...)
}

func (c *CustomLogger) InfoFields(msg string, fields ...zapcore.Field) {
	defaultLogger.InfoWithFields(c.prefix+msg, fields...)
}

func (c *CustomLogger) WarnFields(msg string, fields ...zapcore.Field) {
	defaultLogger.WarnWithFields(c.prefix+msg, fields...)
}

func (c *CustomLogger) ErrorFields(msg string, fields ...zapcore.Field) {
	defaultLogger.ErrorWithFields(c.prefix+msg, fields...)
}

func (c *CustomLogger) PanicFields(msg string, fields ...zapcore.Field) {
	defaultLogger.PanicWithFields(c.prefix+msg, fields...)
}

func (c *CustomLogger) FatalFields(msg string, fields ...zapcore.Field) {
	defaultLogger.FatalWithFields(c.prefix+msg, fields...)
}

// ///////////////////////////////////////////////////

// Debug ZapLogger
func (l *ZapLogger) Debug(args ...interface{}) {
	l.underlieLogger.Debug(args...)
}

// Info logger
func (l *ZapLogger) Info(args ...interface{}) {
	l.underlieLogger.Info(args...)
}

// Warn logger
func (l *ZapLogger) Warn(args ...interface{}) {
	l.underlieLogger.Warn(args...)
}

// Error logger
func (l *ZapLogger) Error(args ...interface{}) {
	l.underlieLogger.Error(args...)
}

func (l *ZapLogger) Fatal(args ...interface{}) {
	l.underlieLogger.Fatal(args...)
}

func (l *ZapLogger) Debugf(format string, args ...interface{}) {
	l.underlieLogger.Debugf(format, args...)
}

func (l *ZapLogger) Infof(format string, args ...interface{}) {
	l.underlieLogger.Infof(format, args...)
}

func (l *ZapLogger) Warnf(format string, args ...interface{}) {
	l.underlieLogger.Warnf(format, args...)
}

func (l *ZapLogger) Errorf(format string, args ...interface{}) {
	l.underlieLogger.Errorf(format, args...)
}

func (l *ZapLogger) Fatalf(format string, args ...interface{}) {
	l.underlieLogger.Fatalf(format, args...)
}

func (l *ZapLogger) Panicf(format string, args ...interface{}) {
	l.underlieLogger.Panicf(format, args...)
}

func (l *ZapLogger) DebugWithFields(msg string, fields ...zap.Field) {
	logger := l.underlieLogger.Desugar()
	logger.Debug(msg, fields...)
}

func (l *ZapLogger) InfoWithFields(msg string, fields ...zap.Field) {
	logger := l.underlieLogger.Desugar()
	logger.Info(msg, fields...)
}

func (l *ZapLogger) WarnWithFields(msg string, fields ...zap.Field) {
	logger := l.underlieLogger.Desugar()
	logger.Warn(msg, fields...)
}

func (l *ZapLogger) ErrorWithFields(msg string, fields ...zap.Field) {
	logger := l.underlieLogger.Desugar()
	logger.Error(msg, fields...)
}

func (l *ZapLogger) PanicWithFields(msg string, fields ...zap.Field) {
	logger := l.underlieLogger.Desugar()
	logger.Panic(msg, fields...)
}

func (l *ZapLogger) FatalWithFields(msg string, fields ...zap.Field) {
	logger := l.underlieLogger.Desugar()
	logger.Fatal(msg, fields...)
}
