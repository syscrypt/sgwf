package logrus

import (
	ls "github.com/sirupsen/logrus"

	"github.com/syscrypt/sgwf/log/logger"
)

type Formatter string

const (
	JsonFormatter Formatter = "json"
)

func NewLogger(formatter Formatter, logLevel string) logger.FieldLogger {
	log := ls.New()

	if formatter == JsonFormatter {
		log.SetFormatter(&ls.JSONFormatter{})
	}

	switch logLevel {
	case "trace":
		log.SetLevel(ls.TraceLevel)
	case "debug":
		log.SetLevel(ls.DebugLevel)
	case "info":
		log.SetLevel(ls.InfoLevel)
	case "warn":
		log.SetLevel(ls.WarnLevel)
	case "error":
		log.SetLevel(ls.ErrorLevel)
	case "fatal":
		log.SetLevel(ls.FatalLevel)
	default:
		log.SetLevel(ls.InfoLevel)
	}

	ret, _ := newLogrusLogger(log)
	return ret
}
