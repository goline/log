package log

// Logger controls log
type Logger interface {
	LogWriter
}

type LogWriter interface {
	// Write logs a message
	Write(level string, message string, args ...interface{}) error
}