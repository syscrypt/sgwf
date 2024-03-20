package logger

type Logger interface {
	Infof(format string, v ...any)
	Warnf(format string, v ...any)
	Errorf(format string, v ...any)
	Debugf(format string, v ...any)
}
