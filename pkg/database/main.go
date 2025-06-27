package database

import (
	"fmt"

	"github.com/listentogether/config/types"
	"github.com/listentogether/database/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

func Connect(databaseConfig *types.Database) (*gorm.DB, error) {
	var err error = nil
	if DBConn != nil {
		fmt.Println("DBConn")
		return DBConn, err
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", databaseConfig.Host, databaseConfig.User, databaseConfig.Password, databaseConfig.DbName, databaseConfig.Port, databaseConfig.SSLmode, databaseConfig.TimeZone)
	db, dbErr := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&models.User{})
	if dbErr != nil {
		return DBConn, fmt.Errorf("database connection error")
	}

	DBConn = db
	return DBConn, err
}
