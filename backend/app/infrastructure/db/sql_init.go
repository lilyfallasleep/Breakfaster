package db

import (
	c "breakfaster/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// NewDatabaseConnection returns the db connection instance
func NewDatabaseConnection(config *c.Config) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(config.DatabaseDsn), &gorm.Config{
		Logger: config.Logger.DBLogger,
	})
	if err != nil {
		return &gorm.DB{}, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return &gorm.DB{}, err
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(config.DBconfig.MaxIdleConns)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(config.DBconfig.MaxOpenConns)
	config.Logger.ContextLogger.WithField("type", "setup:db").Info("successful SQL connection")
	return db, nil
}
