package logrus

import (
	ls "github.com/sirupsen/logrus"

	"github.com/syscrypt/sgwf/log/logger"
)

func NewLogger(formatter ls.Formatter, logLevel ls.Level) logger.FieldLogger {
	log := ls.New()
	log.SetFormatter(formatter)
	log.SetLevel(logLevel)

	ret, _ := newLogrusLogger(log)
	return ret
}
