package gokit

import (
	"github.com/sirupsen/logrus"
)

type CustomLogger struct {
	Prefix string
	LogDir string
	*logrus.Logger
}

func NewCustomLogger(prefix string, logDir string) *CustomLogger {
	return &CustomLogger{Prefix: "[" + prefix + "] ", LogDir: logDir, Logger: logrus.New()}
}

func (c *CustomLogger) SetPrefix(prefix string) {
	c.Prefix = "[" + prefix + "] "
}

func (c *CustomLogger) Print(message ...interface{}) {
	c.Logger.Print(c.Prefix, message)
}

func (c *CustomLogger) Printf(format string, message ...interface{}) {
	c.Logger.Printf(c.Prefix+format, message)
}

func (c *CustomLogger) Println(message ...interface{}) {
	c.Logger.Println(c.Prefix, message)
}

// Info level message
func (c *CustomLogger) Info(message ...interface{}) {
	c.Logger.Info(c.Prefix, message)
}

// Infof - formatted message
func (c *CustomLogger) Infof(message string, args ...interface{}) {
	c.Logger.Infof(c.Prefix+message, args...)
}

// InfoFields - message with fields
func (c *CustomLogger) InfoFields(message string, fields map[string]interface{}) {
	c.Logger.WithFields(fields).Info(c.Prefix + message)
}

// Debug level message
func (c *CustomLogger) Debug(message ...interface{}) {
	c.Logger.Debug(c.Prefix, message)
}

// Debugf - formatted message
func (c *CustomLogger) Debugf(message string, args ...interface{}) {
	c.Logger.Debugf(c.Prefix+message, args...)
}

// DebugFields - message with fields
func (c *CustomLogger) DebugFields(message string, fields map[string]interface{}) {
	c.Logger.WithFields(fields).Debug(c.Prefix + message)
}

// Warn level message
func (c *CustomLogger) Warn(message string) {
	c.Logger.Warn(c.Prefix + message)
}

// Warnf - formatted message
func (c *CustomLogger) Warnf(message string, args ...interface{}) {
	c.Logger.Warnf(c.Prefix+message, args...)
}

// WarnFields - message with fields
func (c *CustomLogger) WarnFields(message string, fields map[string]interface{}) {
	c.Logger.WithFields(fields).Warn(c.Prefix + message)
}

// Error level message
func (c *CustomLogger) Error(message ...interface{}) {
	c.Logger.Error(c.Prefix, message)
}

// Errorf - formatted message
func (c *CustomLogger) Errorf(message string, args ...interface{}) {
	c.Logger.Errorf(c.Prefix+message, args...)
}

// ErrorFields - message with fields
func (c *CustomLogger) ErrorFields(message string, fields map[string]interface{}) {
	c.Logger.WithFields(fields).Error(c.Prefix + message)
}

// Fatal level message
func (c *CustomLogger) Fatal(message string) {
	c.Logger.Fatal(c.Prefix + message)
}

// Fatalf - formatted message
func (c *CustomLogger) Fatalf(message string, args ...interface{}) {
	c.Logger.Fatalf(c.Prefix+message, args...)
}

// FatalFields - message with fields
func (c *CustomLogger) FatalFields(message string, fields map[string]interface{}) {
	c.Logger.WithFields(fields).Fatal(c.Prefix + message)
}

// Panic level message
func (c *CustomLogger) Panic(message string) {
	c.Logger.Panic(c.Prefix + message)
}

// Panicf - formatted message
func (c *CustomLogger) Panicf(message string, args ...interface{}) {
	c.Logger.Panicf(c.Prefix+message, args...)
}

// PanicFields - message with fields
func (c *CustomLogger) PanicFields(message string, fields map[string]interface{}) {
	c.Logger.WithFields(fields).Panic(c.Prefix + message)
}

func (c *CustomLogger) SetLevel(level string) error {
	var (
		err      error
		logLevel logrus.Level
	)
	if logLevel, err = logrus.ParseLevel(level); err == nil {
		c.Logger.SetLevel(logLevel)
	}
	return err
}
