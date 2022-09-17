package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	logrus *logrus.Logger
}

func NewLogger(logLevel string) *Logger {
	log := logrus.New()
	log.Out = os.Stdout
	log.SetLevel(getLogLevelFromString(logLevel))

	return &Logger{
		logrus: log,
	}
}

func (l *Logger) Info(msg string) {
	l.logrus.Info(msg)
}

func (l *Logger) Warn(msg string) {
	l.logrus.Warn(msg)
}

func (l *Logger) Error(err error) {
	l.logrus.Error(err)
}

func getLogLevelFromString(logLevel string) logrus.Level {
	var level logrus.Level
	switch logLevel {
	case "ERROR":
		level = logrus.ErrorLevel
	case "WARN":
		level = logrus.WarnLevel
	case "INFO":
		level = logrus.InfoLevel
	default:
		level = logrus.DebugLevel
	}

	return level
}
