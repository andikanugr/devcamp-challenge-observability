package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

type Fields map[string]interface{}

type Logger interface {
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	WithField(fields Fields) Logger
}

type logrusLogger struct {
	entry *logrus.Entry
}

func NewLogger() Logger {
	log := logrus.New()
	file, err := os.OpenFile("files/logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.SetOutput(file)
	} else {
		log.Info("Failed to log to file, using default stderr")
	}

	log.SetFormatter(&logrus.JSONFormatter{})

	return &logrusLogger{
		entry: logrus.NewEntry(log),
	}
}

func (l *logrusLogger) Info(args ...interface{}) {
	l.entry.Info(args)
}

func (l *logrusLogger) Infof(format string, args ...interface{}) {
	l.entry.Infof(format, args)
}

func (l *logrusLogger) Error(args ...interface{}) {
	l.entry.Error(args)
}

func (l *logrusLogger) Errorf(format string, args ...interface{}) {
	l.entry.Errorf(format, args)
}

func (l *logrusLogger) WithField(fields Fields) Logger {
	return &logrusLogger{
		entry: l.entry.WithFields(logrus.Fields(fields)),
	}
}
