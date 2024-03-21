package basic

import (
	"fmt"
	"log"

	"github.com/syscrypt/sgwf/log/logger"
)

type basicLogger struct{}

func NewBasicLogger() logger.Logger {
	return &basicLogger{}
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

func (l *basicLogger) Printf(format string, v ...any) {
	log.Printf(format, v...)
}
