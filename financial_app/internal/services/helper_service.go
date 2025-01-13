package services

import (
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/sirupsen/logrus"
)

const (
	PathSeparator = string(os.PathSeparator)
)

var (
	idCounter int        // Contador global de IDs
	idMutex   sync.Mutex // Mutex global para proteger el contador
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

// Genera un ID único de manera segura
func GenerateID() int {
	idMutex.Lock()         // Bloquea el Mutex para evitar accesos concurrentes
	defer idMutex.Unlock() // Asegura que el Mutex se libere después de la función
	idCounter++            // Incrementa el contador global
	return idCounter       // Retorna el nuevo ID
}
