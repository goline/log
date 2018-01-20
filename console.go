package log

import (
	"fmt"
	"time"

	"github.com/goline/errors"
)

type ConsoleLogger struct {}

func (l *ConsoleLogger) Write(level string, message string, args ...interface{}) error {
	now := time.Now().Format(time.RFC3339)
	msg := fmt.Sprintf(message, args...)
	fmt.Println(fmt.Sprintf("[%s][%s] %s", now, level, msg))
	return nil
}

func (l *ConsoleLogger) Error(message string, args ...interface{}) error {
	return l.Write(errors.LEVEL_ERROR, message, args...)
}

func (l *ConsoleLogger) Info(message string, args ...interface{}) error {
	return l.Write(errors.LEVEL_INFO, message, args...)
}

func (l *ConsoleLogger) Debug(message string, args ...interface{}) error {
	return l.Write(errors.LEVEL_DEBUG, message, args...)
}

func (l *ConsoleLogger) Warn(message string, args ...interface{}) error {
	return l.Write(errors.LEVEL_WARN, message, args...)
}
