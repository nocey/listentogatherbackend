package models

import (
	"time"

	"github.com/listentogether/database"
)

type AppConfig struct {
	BaseModel
	Name      string    `json:"name"`
	IsActive  bool      `json:"is_active"`
	Value     string    `json:"value"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (config *AppConfig) GetAll() *[]AppConfig {
	configs := []AppConfig{}
	database.DBConn.Table("app_config").Find(&configs)
	if len(configs) == 0 {
		return nil
	}

	return &configs
}

func (config *AppConfig) GetAllWithInObject() map[string]*string {
	configs := config.GetAll()

	interfaceConfig := make(map[string]*string)
	for i := range *configs {
		config := (*configs)[i]

		if config.Name == "" {
			config.Name = "default"
		}

		interfaceConfig[config.Name] = &config.Value
	}

	return interfaceConfig
}
