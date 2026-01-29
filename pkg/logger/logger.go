package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

func init() {
	log = logrus.New()
	log.SetOutput(os.Stdout)
	log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	log.SetLevel(logrus.InfoLevel)
}

// GetLogger retorna la instancia global del logger
func GetLogger() *logrus.Logger {
	return log
}

// Error registra un error
func Error(err error, message string, fields ...map[string]interface{}) {
	entry := log.WithError(err)
	if len(fields) > 0 {
		entry = entry.WithFields(fields[0])
	}
	entry.Error(message)
}

// Info registra información
func Info(message string, fields ...map[string]interface{}) {
	entry := log.WithFields(logrus.Fields{})
	if len(fields) > 0 {
		entry = entry.WithFields(fields[0])
	}
	entry.Info(message)
}

// Warn registra una advertencia
func Warn(message string, fields ...map[string]interface{}) {
	entry := log.WithFields(logrus.Fields{})
	if len(fields) > 0 {
		entry = entry.WithFields(fields[0])
	}
	entry.Warn(message)
}

// Debug registra debug
func Debug(message string, fields ...map[string]interface{}) {
	entry := log.WithFields(logrus.Fields{})
	if len(fields) > 0 {
		entry = entry.WithFields(fields[0])
	}
	entry.Debug(message)
}
