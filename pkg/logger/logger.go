package logger

import (
	"github.com/sirupsen/logrus"
)

type Logger struct{}

func NewLogger(logLevel string) *Logger {
	logrus.SetLevel(getLogLevelFromString(logLevel))
	logrus.SetFormatter(&logrus.JSONFormatter{})

	return &Logger{}
}

func (l *Logger) Info(msg string) {
	logrus.Info(msg)
}

func (l *Logger) Warn(msg string) {
	logrus.Warn(msg)
}

func (l *Logger) Error(msg string) {
	logrus.Error(msg)
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
