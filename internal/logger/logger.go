// Package logger is a helper for logs in gogpt
package logger

import (
	f "fmt"
	"gogpt/pkg/adapter"
	"io/ioutil"
	"os"
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/writer"
	"github.com/spf13/cobra"
)

var (
	lock = sync.RWMutex{}
	log  *adapter.Logger
	once sync.Once
)

// LogrusAdapter is used as a passable adapter used to log.
type LogrusAdapter struct{}

// Printf prints out messages for user, appending line break at the end
func (l LogrusAdapter) Printf(fmt string, args ...interface{}) {
	lock.Lock()
	f.Printf(f.Sprintln(fmt), args...)
	lock.Unlock()
}

// Debugf is used for logging debug level logs.
func (l LogrusAdapter) Debugf(fmt string, args ...interface{}) {
	lock.Lock()
	logrus.Debugf(fmt, args...)
	lock.Unlock()
}

// Warnf is used for logging warning level logs.
func (l LogrusAdapter) Warnf(fmt string, args ...interface{}) {
	lock.Lock()
	logrus.Warnf(fmt, args...)
	lock.Unlock()
}

// Infof is used for logging info level logs.
func (l LogrusAdapter) Infof(fmt string, args ...interface{}) {
	lock.Lock()
	logrus.Infof(fmt, args...)
	lock.Unlock()
}

// Errorf is used for logging error level logs.
func (l LogrusAdapter) Errorf(fmt string, args ...interface{}) {
	lock.Lock()
	logrus.Errorf(fmt, args...)
	lock.Unlock()
}

// Error is used for logging error level logs.
func (l LogrusAdapter) Error(args ...interface{}) {
	lock.Lock()
	logrus.Error(args...)
	lock.Unlock()
}

// SetLevelDebug sets logrus' logging level to debug so debug logs are written.
func SetLevelDebug() {
	lock.Lock()
	logrus.SetLevel(logrus.DebugLevel)
	lock.Unlock()
}

// SetLevelInfo sets logrus' logger level to info so debug logs are not written.
func SetLevelInfo() {
	lock.Lock()
	logrus.SetLevel(logrus.InfoLevel)
	lock.Unlock()
}

// GetLevel tells us the current level
func GetLevel() logrus.Level {
	return logrus.GetLevel()
}

// GetInstance returns an instance of adapter.Logger
func GetInstance() *adapter.Logger {
	once.Do(func() {
		logger := adapter.NewLogger(LogrusAdapter{})
		log = &logger

		logrus.SetOutput(ioutil.Discard) // Send all logs to nowhere by default
		logrus.SetFormatter(&logrus.TextFormatter{
			ForceColors:     true,
			TimestampFormat: "2006-01-02 15:04:05",
			FullTimestamp:   true,
		})

		logrus.AddHook(&writer.Hook{ // Send logs with level warning or higher to stderr
			Writer: os.Stderr,
			LogLevels: []logrus.Level{
				logrus.PanicLevel,
				logrus.FatalLevel,
				logrus.ErrorLevel,
				logrus.WarnLevel,
			},
		})
		logrus.AddHook(&writer.Hook{ // Send info and debug logs to stdout
			Writer: os.Stdout,
			LogLevels: []logrus.Level{
				logrus.InfoLevel,
				logrus.DebugLevel,
			},
		})
	})
	return log
}

// SetVerbose sets the verbose flag for the logger
func SetVerbose(cmd *cobra.Command) {
	verbose, err := cmd.Flags().GetBool("verbose")
	if err != nil {
		log.Errorf("Error while getting verbose flag: %v", err)
	}
	if verbose {
		SetLevelDebug()
	}
}
