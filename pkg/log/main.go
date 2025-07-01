package log

import (
	"fmt"
	"io"
	"os"
	"sync"
)

type DebugLevel int

const (
	FATAL = iota
	ERROR
	WARNING
	NOTICE
	INFO
	DEBUG
)

func (debugLevel *DebugLevel) String() string {
	switch *debugLevel {
	case FATAL:
		return "FATAL"
	case ERROR:
		return "ERROR"
	case WARNING:
		return "WARNING"
	case NOTICE:
		return "NOTICE"
	case INFO:
		return "INFO"
	case DEBUG:
		return "DEBUG"
	default:
		return "UNKNOWN"
	}
}

type Logger struct {
	output io.Writer
	level  DebugLevel
	mu     sync.Mutex
}

var logger = &Logger{}

func Initialize(newLogger *Logger) {
	logger = newLogger
}

func (l *Logger) Log(debugLevel DebugLevel, args ...interface{}) {
	if debugLevel > l.level {
		return
	}

	l.mu.Lock()
	defer l.mu.Unlock()
	l.output.Write([]byte(fmt.Sprintf("[%s]", debugLevel.String())))

}

func Fatal(args ...interface{}) {
	logger.Log(FATAL, args...)
	os.Exit(1)
}

func Error(args ...interface{}) {
	logger.Log(ERROR, args...)
}

func Warning(args ...interface{}) {
	logger.Log(WARNING, args...)
}

func Info(args ...interface{}) {
	logger.Log(INFO, args...)
}

func Debug(args ...interface{}) {
	logger.Log(DEBUG, args...)
}
