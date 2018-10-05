// Copyright 2018 The QOS Authors

package mtype

type Logger interface {
	Debug(msg string, keyvals ...interface{})
	Info(msg string, keyvals ...interface{})
	Error(msg string, keyvals ...interface{})
}

type Context struct {
	Log Logger
}
