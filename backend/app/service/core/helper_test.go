package core

import (
	"breakfaster/config"
	"breakfaster/service/constant"
	"time"

	log "github.com/sirupsen/logrus"
)

func NewDummyConfig() *config.Config {
	dummyConfig := new(config.Config)
	dummyConfig.Logger = &config.Logger{
		ContextLogger: log.WithFields(log.Fields{}),
	}
	return dummyConfig
}

func GetTestDateTimeInterval() (string, string, time.Time, time.Time) {
	startDate := "2020-10-10"
	endDate := "2020-10-15"
	start, _ := time.ParseInLocation(constant.DateFormat, startDate, time.Local)
	end, _ := time.ParseInLocation(constant.DateFormat, endDate, time.Local)
	return startDate, endDate, start, end
}
