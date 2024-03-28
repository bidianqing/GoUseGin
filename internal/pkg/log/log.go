package logger

import (
	"github.com/sirupsen/logrus"
	_ "go.uber.org/zap"
)

var logger *logrus.Logger

func init() {
	if logger == nil {
		logger = logrus.New()
	}

	logger.SetLevel(logrus.InfoLevel)
}

func IsLevelEnabled(level logrus.Level) bool {
	return logger.IsLevelEnabled(level)
}

func Debug(args ...interface{}) {
	logger.Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

func Info(args ...interface{}) {
	logger.Info(args...)
}

func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

func Warn(args ...interface{}) {
	logger.Warn(args...)
}

func Warnf(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}

func Error(args ...interface{}) {
	logger.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}
