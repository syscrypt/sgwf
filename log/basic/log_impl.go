package basic

import (
	"fmt"
	"io"
	"strings"

	"github.com/syscrypt/sgwf/log/logger"
)

type LogImplInterface interface {
	SetLogger(lg logger.Logger)
	SetFieldLogger(flg logger.FieldLogger)
}

type LogImpl struct {
	Logger      logger.Logger
	FieldLogger logger.FieldLogger
}

func NewLog() *LogImpl {
	return &LogImpl{
		Logger: NewBasicLogger(),
	}
}

func (l *LogImpl) Infof(format string, v ...any) {
	if l.FieldLogger != nil {
		l.FieldLogger.Infof(format, v...)
		return
	}
	l.Logger.Infof(format, v...)
}

func (l *LogImpl) Warnf(format string, v ...any) {
	if l.FieldLogger != nil {
		l.FieldLogger.Warnf(format, v...)
		return
	}
	l.Logger.Warnf(format, v...)
}

func (l *LogImpl) Errorf(format string, v ...any) {
	if l.FieldLogger != nil {
		l.FieldLogger.Errorf(format, v...)
		return
	}
	l.Logger.Errorf(format, v...)
}

func (l *LogImpl) Debugf(format string, v ...any) {
	if l.FieldLogger != nil {
		l.FieldLogger.Debugf(format, v...)
		return
	}
	l.Logger.Debugf(format, v...)
}

func (l *LogImpl) Printf(format string, v ...any) {
	if l.FieldLogger != nil {
		l.FieldLogger.Infof(format, v...)
		return
	}
	l.Logger.Printf(format, v...)
}

func (l *LogImpl) Warningf(format string, v ...any) {
	if l.FieldLogger != nil {
		l.FieldLogger.Warningf(format, v...)
		return
	}
	l.Logger.Warnf(format, v...)
}

func (l *LogImpl) Fatalf(format string, v ...any) {
	if l.FieldLogger != nil {
		l.FieldLogger.Fatalf(format, v...)
		return
	}
	l.Logger.Printf(format, v...)
}

func (l *LogImpl) Panicf(format string, v ...any) {
	if l.FieldLogger != nil {
		l.FieldLogger.Panicf(format, v...)
		return
	}
	l.Logger.Printf(format, v...)
}

func (l *LogImpl) Info(args ...interface{}) {
	if l.FieldLogger != nil {
		l.FieldLogger.Info(args...)
		return
	}
	l.Logger.Infof(buildStringFromArgs(args))
}

func (l *LogImpl) Warn(args ...interface{}) {
	if l.FieldLogger != nil {
		l.FieldLogger.Warn(args)
		return
	}
	l.Logger.Warnf(buildStringFromArgs(args))
}

func (l *LogImpl) Error(args ...interface{}) {
	if l.FieldLogger != nil {
		l.FieldLogger.Error(args)
		return
	}
	l.Logger.Errorf(buildStringFromArgs(args))
}

func (l *LogImpl) Debug(args ...interface{}) {
	if l.FieldLogger != nil {
		l.FieldLogger.Debug(args)
		return
	}
	l.Logger.Debugf(buildStringFromArgs(args))
}

func (l *LogImpl) Print(args ...interface{}) {
	if l.FieldLogger != nil {
		l.FieldLogger.Print(args)
		return
	}
	l.Logger.Printf(buildStringFromArgs(args))
}

func (l *LogImpl) Warning(args ...interface{}) {
	if l.FieldLogger != nil {
		l.FieldLogger.Warning(args...)
		return
	}
	l.Logger.Warnf(buildStringFromArgs(args))
}

func (l *LogImpl) Fatal(args ...interface{}) {
	if l.FieldLogger != nil {
		l.FieldLogger.Fatal(args...)
		return
	}
	l.Logger.Printf(buildStringFromArgs(args))
}

func (l *LogImpl) Panic(args ...interface{}) {
	if l.FieldLogger != nil {
		l.FieldLogger.Panic(args...)
		return
	}
	l.Logger.Printf(buildStringFromArgs(args))
}

func (l *LogImpl) Infoln(args ...interface{}) {
	if l.FieldLogger != nil {
		l.FieldLogger.Infoln(args...)
		return
	}
	l.Logger.Infof(buildStringFromArgs(args))
}

func (l *LogImpl) Warnln(args ...interface{}) {
	if l.FieldLogger != nil {
		l.FieldLogger.Warnln(args)
		return
	}
	l.Logger.Warnf(buildStringFromArgs(args))
}

func (l *LogImpl) Errorln(args ...interface{}) {
	if l.FieldLogger != nil {
		l.FieldLogger.Errorln(args)
		return
	}
	l.Logger.Errorf(buildStringFromArgs(args))
}

func (l *LogImpl) Debugln(args ...interface{}) {
	if l.FieldLogger != nil {
		l.FieldLogger.Debugln(args)
		return
	}
	l.Logger.Debugf(buildStringFromArgs(args))
}

func (l *LogImpl) Println(args ...interface{}) {
	if l.FieldLogger != nil {
		l.FieldLogger.Println(args)
		return
	}
	l.Logger.Printf(buildStringFromArgs(args))
}

func (l *LogImpl) Warningln(args ...interface{}) {
	if l.FieldLogger != nil {
		l.FieldLogger.Warningln(args...)
		return
	}
	l.Logger.Warnf(buildStringFromArgs(args))
}

func (l *LogImpl) Fatalln(args ...interface{}) {
	if l.FieldLogger != nil {
		l.FieldLogger.Fatalln(args...)
		return
	}
	l.Logger.Printf(buildStringFromArgs(args))
}

func (l *LogImpl) Panicln(args ...interface{}) {
	if l.FieldLogger != nil {
		l.FieldLogger.Panicln(args...)
		return
	}
	l.Logger.Printf(buildStringFromArgs(args))
}

func (l *LogImpl) SetOutput(writer io.Writer) {
	if l.FieldLogger == nil {
		return
	}
	l.FieldLogger.SetOutput(writer)
}

func (l *LogImpl) WithField(key string, value interface{}) logger.FieldLogger {
	if l.FieldLogger == nil {
		return l
	}

	return l.FieldLogger.WithField(key, value)
}

func (l *LogImpl) WithFields(fields logger.Fields) logger.FieldLogger {
	if l.FieldLogger == nil {
		return l
	}

	return l.FieldLogger.WithFields(fields)
}

func (l *LogImpl) WithError(err error) logger.FieldLogger {
	if l.FieldLogger == nil {
		return l
	}

	return l.FieldLogger.WithError(err)
}

func (l *LogImpl) SetLogger(lg logger.Logger) {
	l.Logger = lg
}

func (l *LogImpl) SetFieldLogger(flg logger.FieldLogger) {
	l.FieldLogger = flg
}

func buildStringFromArgs(args ...interface{}) string {
	argsStr := make([]string, len(args))
	for idx, arg := range args {
		argsStr[idx] = fmt.Sprintf("%+v", arg)
	}
	return strings.Join(argsStr, " ")
}
