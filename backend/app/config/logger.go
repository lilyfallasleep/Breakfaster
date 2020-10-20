package config

import (
	"io"
	"log"
	"os"
	"time"

	dblog "gorm.io/gorm/logger"
)

// Logger is the logger type
type Logger struct {
	Writer   io.Writer
	DBLogger dblog.Interface
}

func getLogger(ginMode, logPath string) (*Logger, error) {
	var writer io.Writer
	var dbLogLevel dblog.LogLevel
	if ginMode == "release" {
		// Create creates or truncates the named file. If the file already exists, it is truncated
		fileWriter, err := os.Create(logPath)
		if err != nil {
			return &Logger{}, err
		}
		writer = io.MultiWriter(fileWriter, os.Stderr)
		dbLogLevel = dblog.Error
	} else {
		writer = os.Stderr
		dbLogLevel = dblog.Info
	}
	return &Logger{
		Writer: writer,
		DBLogger: dblog.New(
			log.New(writer, "\r\n", log.LstdFlags), // io writer
			dblog.Config{
				SlowThreshold: time.Second, // Slow SQL threshold
				LogLevel:      dbLogLevel,  // Log level
				Colorful:      true,        // Enable color
			},
		),
	}, nil
}
