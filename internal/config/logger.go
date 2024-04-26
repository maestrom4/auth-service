package config

import (
	"strings"

	"github.com/sirupsen/logrus"
)

func InitLogger() {
	var logger = logrus.New()
	logger.Formatter = &logrus.TextFormatter{
		DisableColors:   false,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	}
	logger.Debugln("LogLevel: ", LogLevel)
	logLevelStr := strings.ToUpper(LogLevel)
	logLevel, err := logrus.ParseLevel(logLevelStr)
	if err != nil {
		logLevel = logrus.DebugLevel
	}

	logger.SetLevel(logLevel)
}
