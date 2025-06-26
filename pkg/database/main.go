package database

import (
	"fmt"

	"github.com/listentogether/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var psql *gorm.DB


func Connect (config *config.Config) (*gorm.DB,error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", config.Database.Host, config.Database.User, config.Database.Password ,config.Database.DbName, config.Database.Port, config.Database.SSLmode ,config.Database.TimeZone)
	var err error = nil
	if (psql != nil) {
		return psql, err
	}
	

	db, dbErr := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if (dbErr != nil){
		return psql, fmt.Errorf("Database connection error")
	}

	psql = db
	return psql, err
}