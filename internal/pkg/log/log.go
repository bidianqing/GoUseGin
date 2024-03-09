package log

import (
	logger "github.com/sirupsen/logrus"
)

var Logger *logger.Logger = logger.New()

func init() {
	if Logger == nil {
		Logger = logger.New()
	}

	Logger.SetLevel(logger.InfoLevel)
}
