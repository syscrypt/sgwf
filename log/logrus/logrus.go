package logrus

import (
	"io"

	"github.com/syscrypt/sgwf/log/logger"

	log "github.com/sirupsen/logrus"
)

type logrusLogEntry struct {
	entry *log.Entry
}

type logrusLogger struct {
	logger *log.Logger
}

func newLogrusLogger(logger *log.Logger) (logger.FieldLogger, error) {
	return &logrusLogger{
		logger: logger,
	}, nil
}

func (l *logrusLogger) Print(args ...interface{}) {
	l.logger.Print(args...)
}

func (l *logrusLogger) Debug(args ...interface{}) {
	l.logger.Debug(args...)
}

func (l *logrusLogger) Info(args ...interface{}) {
	l.logger.Info(args...)
}

func (l *logrusLogger) Warn(args ...interface{}) {
	l.logger.Warn(args...)
}

func (l *logrusLogger) Warning(args ...interface{}) {
	l.logger.Warning(args...)
}

func (l *logrusLogger) Error(args ...interface{}) {
	l.logger.Error(args...)
}

func (l *logrusLogger) Panic(args ...interface{}) {
	l.logger.Panic(args...)
}

func (l *logrusLogger) Fatal(args ...interface{}) {
	l.logger.Fatal(args...)
}

func (l *logrusLogger) Println(args ...interface{}) {
	l.logger.Println(args...)
}

func (l *logrusLogger) Debugln(args ...interface{}) {
	l.logger.Debugln(args...)
}

func (l *logrusLogger) Infoln(args ...interface{}) {
	l.logger.Infoln(args...)
}

func (l *logrusLogger) Warnln(args ...interface{}) {
	l.logger.Warnln(args...)
}

func (l *logrusLogger) Warningln(args ...interface{}) {
	l.logger.Warningln(args...)
}

func (l *logrusLogger) Errorln(args ...interface{}) {
	l.logger.Errorln(args...)
}

func (l *logrusLogger) Panicln(args ...interface{}) {
	l.logger.Panicln(args...)
}

func (l *logrusLogger) Fatalln(args ...interface{}) {
	l.logger.Fatalln(args...)
}

func (l *logrusLogger) Debugf(template string, args ...interface{}) {
	l.logger.Debugf(template, args...)
}

func (l *logrusLogger) Printf(template string, args ...interface{}) {
	l.logger.Printf(template, args...)
}

func (l *logrusLogger) Infof(template string, args ...interface{}) {
	l.logger.Infof(template, args...)
}

func (l *logrusLogger) Warnf(template string, args ...interface{}) {
	l.logger.Warnf(template, args...)
}

func (l *logrusLogger) Warningf(template string, args ...interface{}) {
	l.logger.Warningf(template, args...)
}

func (l *logrusLogger) Errorf(template string, args ...interface{}) {
	l.logger.Errorf(template, args...)
}

func (l *logrusLogger) Panicf(template string, args ...interface{}) {
	l.logger.Panicf(template, args...)
}

func (l *logrusLogger) Fatalf(template string, args ...interface{}) {
	l.logger.Fatalf(template, args...)
}

func (l *logrusLogger) WithError(err error) logger.FieldLogger {
	return &logrusLogEntry{
		entry: l.logger.WithError(err),
	}
}

func (l *logrusLogger) WithFields(fields logger.Fields) logger.FieldLogger {
	return &logrusLogEntry{
		entry: l.logger.WithFields(convertToLogrusFields(fields)),
	}
}

func (l *logrusLogger) WithField(key string, value interface{}) logger.FieldLogger {
	return &logrusLogEntry{
		entry: l.logger.WithField(key, value),
	}
}

func (l *logrusLogger) SetOutput(writer io.Writer) {
	l.logger.SetOutput(writer)
}

func (l *logrusLogEntry) Print(args ...interface{}) {
	l.entry.Print(args...)
}

func (l *logrusLogEntry) Println(args ...interface{}) {
	l.entry.Println(args...)
}

func (l *logrusLogEntry) Debug(args ...interface{}) {
	l.entry.Debug(args...)
}

func (l *logrusLogEntry) Info(args ...interface{}) {
	l.entry.Info(args...)
}

func (l *logrusLogEntry) Warn(args ...interface{}) {
	l.entry.Warn(args...)
}

func (l *logrusLogEntry) Warning(args ...interface{}) {
	l.entry.Warning(args...)
}

func (l *logrusLogEntry) Error(args ...interface{}) {
	l.entry.Error(args...)
}

func (l *logrusLogEntry) Panic(args ...interface{}) {
	l.entry.Panic(args...)
}

func (l *logrusLogEntry) Fatal(args ...interface{}) {
	l.entry.Fatal(args...)
}

func (l *logrusLogEntry) Debugln(args ...interface{}) {
	l.entry.Debugln(args...)
}

func (l *logrusLogEntry) Infoln(args ...interface{}) {
	l.entry.Infoln(args...)
}

func (l *logrusLogEntry) Warnln(args ...interface{}) {
	l.entry.Warnln(args...)
}

func (l *logrusLogEntry) Warningln(args ...interface{}) {
	l.entry.Warningln(args...)
}

func (l *logrusLogEntry) Errorln(args ...interface{}) {
	l.entry.Errorln(args...)
}

func (l *logrusLogEntry) Panicln(args ...interface{}) {
	l.entry.Panicln(args...)
}

func (l *logrusLogEntry) Fatalln(args ...interface{}) {
	l.entry.Fatalln(args...)
}

func (l *logrusLogEntry) Debugf(template string, args ...interface{}) {
	l.entry.Debugf(template, args...)
}

func (l *logrusLogEntry) Printf(template string, args ...interface{}) {
	l.entry.Printf(template, args...)
}

func (l *logrusLogEntry) Infof(template string, args ...interface{}) {
	l.entry.Infof(template, args...)
}

func (l *logrusLogEntry) Warnf(template string, args ...interface{}) {
	l.entry.Warnf(template, args...)
}

func (l *logrusLogEntry) Warningf(template string, args ...interface{}) {
	l.entry.Warnf(template, args...)
}

func (l *logrusLogEntry) Errorf(template string, args ...interface{}) {
	l.entry.Errorf(template, args...)
}

func (l *logrusLogEntry) Panicf(template string, args ...interface{}) {
	l.entry.Panicf(template, args...)
}

func (l *logrusLogEntry) Fatalf(template string, args ...interface{}) {
	l.entry.Fatalf(template, args...)
}

func (l *logrusLogEntry) WithFields(fields logger.Fields) logger.FieldLogger {
	return &logrusLogEntry{
		entry: l.entry.WithFields(convertToLogrusFields(fields)),
	}
}

func (l *logrusLogEntry) WithError(err error) logger.FieldLogger {
	return &logrusLogEntry{
		entry: l.entry.WithError(err),
	}
}

func (l *logrusLogEntry) WithField(key string, value interface{}) logger.FieldLogger {
	return &logrusLogEntry{
		entry: l.entry.WithField(key, value),
	}
}
func (l *logrusLogEntry) SetOutput(writer io.Writer) {
	l.entry.Logger.SetOutput(writer)
}

func convertToLogrusFields(fields logger.Fields) log.Fields {
	logrusFields := log.Fields{}
	for index, val := range fields {
		logrusFields[index] = val
	}
	return logrusFields
}
