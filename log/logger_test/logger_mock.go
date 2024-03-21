package logger_test

import (
	"fmt"
)

type LoggerMock struct {
	Buf []string
}

func (l *LoggerMock) Infof(format string, v ...any) {
	inner := fmt.Sprintf(format, v...)
	l.Buf = append(l.Buf, fmt.Sprintf("info: %s", inner))
}

func (l *LoggerMock) Warnf(format string, v ...any) {
	inner := fmt.Sprintf(format, v...)
	l.Buf = append(l.Buf, fmt.Sprintf("warn: %s", inner))
}

func (l *LoggerMock) Errorf(format string, v ...any) {
	inner := fmt.Sprintf(format, v...)
	l.Buf = append(l.Buf, fmt.Sprintf("error: %s", inner))
}

func (l *LoggerMock) Debugf(format string, v ...any) {
	inner := fmt.Sprintf(format, v...)
	l.Buf = append(l.Buf, fmt.Sprintf("debug: %s", inner))
}

func (l *LoggerMock) Printf(format string, v ...any) {
	inner := fmt.Sprintf(format, v...)
	l.Buf = append(l.Buf, fmt.Sprintf("print: %s", inner))
}
