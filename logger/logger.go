package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Logger defines a basic abstracted logger
type Logger interface {
	Warningf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Panicf(format string, args ...interface{})
	Printf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	Tracef(format string, args ...interface{})
}

// L defines the logger instance
var L = InitLogger()

// InitLogger init logger
func InitLogger() Logger {
	l := logrus.New()

	l.SetOutput(os.Stdout)

	// initial log formatting; this setting is updated after the daemon configuration is loaded.
	l.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	return l
}
