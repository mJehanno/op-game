package logger

import (
	"io"
	"mult-game/backend/conf"
	"os"
	"path"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	ErrLogger   logrus.FieldLogger
	DebugLogger logrus.FieldLogger
	LogFile     *os.File
}

var logger *Logger

func NewLogger() *Logger {
	debug := logrus.New() // TODO : define log file (using a multiwriter)
	debug.SetLevel(logrus.DebugLevel)
	debug.SetOutput(os.Stdout)

	var f *os.File
	erl := logrus.New()

	confFolder, err := conf.GetConfFolder()
	if err != nil {
		erl.WithError(err).Error("failed to get config folder")
	}

	errorPathFile := path.Join(*confFolder, "errors.txt")
	f, err = os.OpenFile(errorPathFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		erl.WithError(err).Error("failed to create error log file")
	}

	erl.SetReportCaller(true)
	errWriter := io.MultiWriter(os.Stdout, f)
	erl.SetOutput(errWriter)
	if logger == nil {
		logger = &Logger{
			ErrLogger:   erl,
			DebugLogger: debug,
			LogFile:     f,
		}
	}

	return logger
}
