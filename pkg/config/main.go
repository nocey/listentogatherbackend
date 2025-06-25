package config

import (
	"fmt"
	"os"
	"reflect"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	// Server types.ServerConfig `env:"Server"`
	// Database string `env:"database"`
	// Log string `env:"log"`
	Port string `env:"PORT"`
	ConfigError string `env:"CONFIG_ERROR"`
}

func (c *Config) Load() Config {
	t := reflect.ValueOf(c)

	if t.Kind() != reflect.Ptr || t.Elem().Kind() != reflect.Struct {
		c.ConfigError = "Config must be a pointer to a struct"
		return *c
	}

	for i := 0; i < t.Elem().NumField(); i++ {
		field := t.Elem().Field(i)
		tag := t.Elem().Type().Field(i).Tag.Get("env")

		if tag == "" || tag == "CONFIG_ERROR" {
			continue
		}

		value := os.Getenv(tag)
		if value == "" {
			c.ConfigError = "Environment variable " + tag + " not set"
			return *c
		}

		switch field.Kind() {
		case reflect.String:
			field.SetString(value)
		case reflect.Int:
			intValue, err := strconv.ParseInt(value,10, 64)
			if err != nil {
				c.ConfigError = "Invalid value for " + tag + ": " + err.Error()
				return *c
			}
			
			field.SetInt(int64(intValue))
		default:
			c.ConfigError = "Unsupported field type for " + tag
			return *c
		}
		fmt.Println("Setting", tag, "to", value)
	}
	return *c
}

func Load() (*Config, error) {
	appEnv := os.Getenv("APP_ENV")
	envLocation := os.Getenv("ENV_LOCATION")

	if appEnv == "" {
		appEnv = "development"
	}

	var err error = nil;
	if appEnv == "development" {
		err = godotenv.Load("./config/.env." + appEnv)
	} else {
		err = godotenv.Load(envLocation + "/.env." + appEnv)
	}
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}

	config := Config{}
	config = config.Load()

	if config.ConfigError != "" {
		return nil, fmt.Errorf("config load error: %s", config.ConfigError)
	}
	fmt.Println("Config loaded successfully")

	return &config, nil
}