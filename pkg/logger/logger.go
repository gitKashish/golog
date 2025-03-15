package logger

import (
	"log"
	"os"
)

// Logger levels
const (
	DEBUG = iota
	INFO
	WARN
	ERROR
	FATAL
)

// Logger is a simple logger interface
type Logger interface {
	Debug(format string, args ...interface{})
	Info(format string, args ...interface{})
	Warn(format string, args ...interface{})
	Error(format string, args ...interface{})
	Fatal(format string, args ...interface{})
	SetLevel(level int)
}

// SimpleLogger is a simple implementation of the Logger interface
type SimpleLogger struct {
	level int
	log   *log.Logger
}

// NewLogger creates a new logger
func NewLogger() *SimpleLogger {
	return &SimpleLogger{
		level: INFO,
		log:   log.New(os.Stdout, "", log.LstdFlags),
	}
}

// Debug logs a debug message
func (l *SimpleLogger) Debug(format string, args ...interface{}) {
	if l.level <= DEBUG {
		l.log.Printf("[DEBUG] "+format, args...)
	}
}

// Info logs an info message
func (l *SimpleLogger) Info(format string, args ...interface{}) {
	if l.level <= INFO {
		l.log.Printf("[INFO] "+format, args...)
	}
}

// Warn logs a warning message
func (l *SimpleLogger) Warn(format string, args ...interface{}) {
	if l.level <= WARN {
		l.log.Printf("[WARN] "+format, args...)
	}
}

// Error logs an error message
func (l *SimpleLogger) Error(format string, args ...interface{}) {
	if l.level <= ERROR {
		l.log.Printf("[ERROR] "+format, args...)
	}
}

// Fatal logs a fatal message and exits
func (l *SimpleLogger) Fatal(format string, args ...interface{}) {
	if l.level <= FATAL {
		l.log.Printf("[FATAL] "+format, args...)
		os.Exit(1)
	}
}

// SetLevel sets the log level
func (l *SimpleLogger) SetLevel(level int) {
	l.level = level
}

// Global logger instance
var globalLogger Logger = NewLogger()

// GetLogger returns the global logger
func GetLogger() Logger {
	return globalLogger
}

// SetLogger sets the global logger
func SetLogger(logger Logger) {
	globalLogger = logger
}

// Debug logs a debug message using the global logger
func Debug(format string, args ...interface{}) {
	globalLogger.Debug(format, args...)
}

// Info logs an info message using the global logger
func Info(format string, args ...interface{}) {
	globalLogger.Info(format, args...)
}

// Warn logs a warning message using the global logger
func Warn(format string, args ...interface{}) {
	globalLogger.Warn(format, args...)
}

// Error logs an error message using the global logger
func Error(format string, args ...interface{}) {
	globalLogger.Error(format, args...)
}

// Fatal logs a fatal message and exits using the global logger
func Fatal(format string, args ...interface{}) {
	globalLogger.Fatal(format, args...)
}

// SetLevel sets the log level of the global logger
func SetLevel(level int) {
	globalLogger.SetLevel(level)
}
