package logger

import (
	"log"
	"os"

	"github.com/sirupsen/logrus"
)

const (
	errorLevel = "ERROR"
	warnLevel  = "WARN"
	infoLevel  = "INFO"
	debugLevel = "DEBUG"
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
	case errorLevel:
		level = logrus.ErrorLevel
	case warnLevel:
		level = logrus.WarnLevel
	case infoLevel:
		level = logrus.InfoLevel
	case debugLevel:
		level = logrus.DebugLevel
	default:
		log.Fatalf("unknown logLevel %s\n", logLevel)
	}

	return level
}
