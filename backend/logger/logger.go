package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	ErrLogger   logrus.FieldLogger
	DebugLogger logrus.FieldLogger
}

var logger *Logger

func NewLogger() *Logger {
	debug := logrus.New() // TODO : define log file (using a multiwriter)
	debug.SetLevel(logrus.DebugLevel)
	debug.SetOutput(os.Stdout)

	erl := logrus.New() // TODO : define log file
	if logger == nil {
		logger = &Logger{
			ErrLogger:   erl,
			DebugLogger: debug,
		}
	}

	return logger
}
