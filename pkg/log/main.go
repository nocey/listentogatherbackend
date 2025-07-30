package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"time"
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

func New(filePath string, debugLevel DebugLevel) {
	logFile, fileErr := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if fileErr != nil {
		fmt.Printf("Failed to open log file: %v\n", fileErr)
		os.Exit(1)
	}
	logger = &Logger{
		output: logFile,
		level:  debugLevel,
	}
	log.SetOutput(logFile)
}

func (l *Logger) Log(debugLevel DebugLevel, args ...interface{}) {
	if debugLevel > l.level {
		return
	}

	if l.output == nil {
		os.Exit(1)
	}

	timestamp := time.Now().Format("2006-01-02 15:04:05")
	msg := fmt.Sprintf("%s [%s] ", timestamp, debugLevel.String())
	msg += fmt.Sprint(args...) + "\n"
	l.mu.Lock()
	defer l.mu.Unlock()
	l.output.Write([]byte(msg))

}

func Fatal(args ...interface{}) {
	logger.Log(FATAL, args...)
	os.Exit(1)
}

func Error(args ...interface{}) {
	logger.Log(ERROR, args...)
}

func Notice(args ...interface{}) {
	logger.Log(NOTICE, args...)
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
