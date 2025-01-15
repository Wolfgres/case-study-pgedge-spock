package services

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
)

const (
	PathSeparator = string(os.PathSeparator)
)

// Get log.Level from string from config file
func GetLogLevel(level string) logrus.Level {

	var logLevel logrus.Level

	switch strings.ToLower(level) {

	case "debug":
		logLevel = logrus.DebugLevel
	case "info":
		logLevel = logrus.InfoLevel
	case "warning":
		logLevel = logrus.WarnLevel
	case "error":
		logLevel = logrus.ErrorLevel
	case "fatal":
		logLevel = logrus.FatalLevel
	case "panic":
		logLevel = logrus.PanicLevel

	}

	return logLevel
}

// Get a exec path in string
func GetExecPath() string {
	exe, err := os.Executable()

	if err != nil {
		logrus.Error(err)
	}
	s := filepath.Dir(exe)

	if s[0:4] == "/tmp" {
		s = "." + PathSeparator
	}

	return s
}

func parseAnyToInt(value any) (int, error) {
	if intValue, ok := value.(int); ok {
		return intValue, nil
	}
	return 0, fmt.Errorf("no se puede convertir %v a int", value)
}
