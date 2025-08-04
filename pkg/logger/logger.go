package logger

import (
	"log"
	"os"
)

// Logger interface defines logging methods
type Logger interface {
	Info(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Debug(args ...interface{})
}

// logger implements the Logger interface
type logger struct {
	infoLogger  *log.Logger
	errorLogger *log.Logger
	debugLogger *log.Logger
}

// NewLogger creates a new logger instance
func NewLogger() Logger {
	return &logger{
		infoLogger:  log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		errorLogger: log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
		debugLogger: log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

// Info logs info messages
func (l *logger) Info(args ...interface{}) {
	l.infoLogger.Println(args...)
}

// Error logs error messages
func (l *logger) Error(args ...interface{}) {
	l.errorLogger.Println(args...)
}

// Fatal logs fatal messages and exits
func (l *logger) Fatal(args ...interface{}) {
	l.errorLogger.Fatal(args...)
}

// Debug logs debug messages
func (l *logger) Debug(args ...interface{}) {
	l.debugLogger.Println(args...)
}
