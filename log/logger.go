package log

import (
	syslog "log"
)

type Logger interface {
	Info(a interface{})
	Infof(format string, a ...interface{})
	Error(a interface{})
	Errorf(format string, a ...interface{})
}

var logger Logger

func init() {
	logger = &ElbsLogger{}
}

func SetLogger(l Logger) {
	logger = l
}

type ElbsLogger struct {
}

func (l *ElbsLogger) Info(a interface{}) {
	syslog.Println(a)
}

func (l *ElbsLogger) Infof(format string, a ...interface{}) {
	syslog.Printf(format, a...)
}

func (l *ElbsLogger) Error(a interface{}) {
	syslog.Println(a)
}

func (l *ElbsLogger) Errorf(format string, a ...interface{}) {
	syslog.Printf(format, a...)
}

func Info(a interface{}) {
	logger.Info(a)
}

func Infof(format string, a ...interface{}) {
	logger.Infof(format, a...)
}

func Error(a interface{}) {
	logger.Error(a)
}

func Errorf(format string, a ...interface{}) {
	logger.Errorf(format, a...)
}
