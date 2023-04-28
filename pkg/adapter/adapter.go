// Package adapter packag efor working with configs
package adapter

import "fmt"

// Version is the gogpt overall build version
var Version = "edge"

// Kube represents the kube file used during testing. Leave blank for in cluster.
var Kube = ""

// TestKube can be copied in to Kube for ease of use
var TestKube = "/root/.kube/config"

// Repo is the ACR which will be used when loading workflow images
var Repo = ""

// Logger struct wrapping around an Adapter.
type Logger struct {
	adapter Adapter
}

// Adapter interface for different log levels.
type Adapter interface {
	Printf(string, ...interface{})
	Debugf(string, ...interface{})
	Infof(string, ...interface{})
	Warnf(string, ...interface{})
	Errorf(string, ...interface{})
	Error(...interface{})
}

// SetLogger changes the adapter used by the logger.
func (l *Logger) SetLogger(a Adapter) {
	l.adapter = a
}

// Printf is used for printing out messages for user.
func (l *Logger) Printf(fmt string, args ...interface{}) {
	l.adapter.Printf(fmt, args...)
}

// Debugf is used for logging debug level logs.
func (l *Logger) Debugf(fmt string, args ...interface{}) {
	l.adapter.Debugf(fmt, args...)
}

// Infof is used for logging info level logs.
func (l *Logger) Infof(fmt string, args ...interface{}) {
	l.adapter.Infof(fmt, args...)
}

// Warnf is used for logging warning level logs.
func (l *Logger) Warnf(fmt string, args ...interface{}) {
	l.adapter.Warnf(fmt, args...)
}

// Errorf is used for logging error level logs.
func (l *Logger) Errorf(fmt string, args ...interface{}) {
	l.adapter.Errorf(fmt, args...)
}

// ErrorfRet is used for logging error level logs. Is fluid
func (l *Logger) ErrorfRet(stringFmt string, args ...interface{}) error {
	l.adapter.Errorf(stringFmt, args...)
	return fmt.Errorf(stringFmt, args...)
}

// Error is used for logging error level logs.
func (l *Logger) Error(args ...interface{}) {
	l.adapter.Error(args...)
}

// NewLogger creates a new instance of Logger with passed in adapter.
func NewLogger(a Adapter) Logger {
	return Logger{adapter: a}
}
