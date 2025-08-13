package database

import (
	"fmt"
	"log"
	"time"

	"github.com/listentogether/config/types"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DBConn *gorm.DB

func Connect(databaseConfig *types.Database) (*gorm.DB, error) {
	var err error = nil
	if DBConn != nil {
		return DBConn, err
	}
	newLogger := logger.New(
		log.Default(),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      true,
			Colorful:                  true,
		},
	)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", databaseConfig.Host, databaseConfig.User, databaseConfig.Password, databaseConfig.DbName, databaseConfig.Port, databaseConfig.SSLmode, databaseConfig.TimeZone)
	db, dbErr := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if dbErr != nil {
		return DBConn, fmt.Errorf("database connection error")
	}

	DBConn = db
	return DBConn, err
}
