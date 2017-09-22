package log

// Logger controls log
type Logger interface {
	LogWriter
	LogLeveller
}

type LogWriter interface {
	// Write logs a message
	Write(level string, message string, args ...interface{}) error
}

type LogLeveller interface {
	// Error writes error message to log
	Error(message string, args ...interface{}) error

	// Debug writes debug message to log
	Debug(message string, args ...interface{}) error

	// Warn writes warn message to log
	Warn(message string, args ...interface{}) error

	// Info writes info message to log
	Info(message string, args ...interface{}) error
}
