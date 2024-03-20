package basic

import (
	"fmt"
	"log"

	"github.com/syscrypt/sgwf/log/logger"
)

type basicLogger struct{}

type LogImplInterface interface {
	SetLogger(lg logger.Logger)
	SetFieldLogger(flg logger.FieldLogger)
}

type LogImpl struct {
	Logger      logger.Logger
	FieldLogger logger.FieldLogger
}

func NewBasicLogger() logger.Logger {
	return &basicLogger{}
}

func NewLog() *LogImpl {
	return &LogImpl{
		Logger: NewBasicLogger(),
	}
}

func (l *basicLogger) Infof(format string, v ...any) {
	log.Printf(fmt.Sprintf("info: %s", format), v...)
}

func (l *basicLogger) Warnf(format string, v ...any) {
	log.Printf(fmt.Sprintf("warn: %s", format), v...)
}

func (l *basicLogger) Errorf(format string, v ...any) {
	log.Printf(fmt.Sprintf("error: %s", format), v...)
}

func (l *basicLogger) Debugf(format string, v ...any) {
	log.Printf(fmt.Sprintf("debug: %s", format), v...)
}

func (l *LogImpl) SetLogger(lg logger.Logger) {
	l.Logger = lg
}

func (l *LogImpl) SetFieldLogger(flg logger.FieldLogger) {
	l.FieldLogger = flg
}
