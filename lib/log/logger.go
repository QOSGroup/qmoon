// Copyright 2018 The QOS Authors

package log

import "github.com/sirupsen/logrus"

type Logger interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
}

func New() Logger {
	logger:=logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	return logger
}
