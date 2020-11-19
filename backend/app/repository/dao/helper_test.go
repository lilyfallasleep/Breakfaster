package dao

import (
	"breakfaster/config"

	log "github.com/sirupsen/logrus"
)

func NewDummyConfig() *config.Config {
	dummyConfig := new(config.Config)
	dummyConfig.Logger = &config.Logger{
		ContextLogger: log.WithFields(log.Fields{}),
	}
	return dummyConfig
}
