package config

import (
	"io"
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	dblog "gorm.io/gorm/logger"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
}

// Logger is the logger type
type Logger struct {
	Writer        io.Writer
	DBLogger      dblog.Interface
	ContextLogger *log.Entry
}

func getLogger(ginMode, logPath, appName string) (*Logger, error) {
	var writer io.Writer
	var dbLogLevel dblog.LogLevel
	var dbLoggerLevel log.Level
	if ginMode == "release" {
		// Create creates or truncates the named file. If the file already exists, it is truncated
		fileWriter, err := os.Create(logPath)
		if err != nil {
			return &Logger{}, err
		}
		writer = io.MultiWriter(fileWriter, os.Stderr)
		dbLogLevel = dblog.Silent      // disable gorm log
		dbLoggerLevel = log.ErrorLevel // just a placeholder
	} else {
		writer = os.Stderr
		dbLogLevel = dblog.Info
		dbLoggerLevel = log.DebugLevel
	}

	var dbLogger = &log.Logger{
		Out:       writer,
		Formatter: new(log.TextFormatter),
		Level:     dbLoggerLevel,
	}

	contextLogger := log.WithFields(log.Fields{
		"app_name": appName,
	})

	return &Logger{
		Writer: writer,
		DBLogger: dblog.New(
			dbLogger,
			dblog.Config{
				SlowThreshold: time.Second, // Slow SQL threshold
				LogLevel:      dbLogLevel,  // Log level
				Colorful:      true,        // Enable color
			},
		),
		ContextLogger: contextLogger,
	}, nil
}
