package config

import (
	"fmt"
	"os"
	"reflect"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/listentogether/config/types"
)

type Config struct {
	// Server types.ServerConfig `env:"Server"`
	Database types.Database
	// Log string `env:"log"`
	Port        string `env:"PORT"`
	JwtToken    []byte `env:"JWT_AUTH_TOKEN" type:"byte"`
	ConfigError string `env:"CONFIG_ERROR"`
}

var config *Config

func processConfig(t reflect.Value) error {
	if t.Kind() != reflect.Ptr || t.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("Config must be a pointer to a struct")
	}

	elem := t.Elem()
	elemType := t.Elem().Type()

	for i := 0; i < elem.NumField(); i++ {
		field := elem.Field(i)

		if field.Kind() == reflect.Struct {
			if err := processConfig(field.Addr()); err != nil {
				return err
			}
			continue
		}
		tag := elemType.Field(i).Tag.Get("env")

		if tag == "" || tag == "CONFIG_ERROR" {
			continue
		}

		value := os.Getenv(tag)
		if value == "" {
			return fmt.Errorf("Environment variable " + tag + " not set")
		}

		switch field.Kind() {
		case reflect.String:
			field.SetString(value)
		case reflect.Slice:
			if tag := elemType.Field(i).Tag.Get("type"); tag == "byte" {
				field.SetBytes([]byte(value))
			} else {
				field.SetString(value)
			}
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			intValue, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				return fmt.Errorf("invalid value for " + tag + ": " + err.Error())
			}
			field.SetInt(intValue)

		default:
			return fmt.Errorf("Unsupported field type for " + tag + field.Kind().String())
		}
		fmt.Println("Setting", tag, "to", value)
	}
	return nil
}

func Load() (*Config, error) {
	if config != nil {
		return config, nil
	}
	appEnv := os.Getenv("APP_ENV")
	envLocation := os.Getenv("ENV_LOCATION")

	if appEnv == "" {
		appEnv = "development"
	}

	var err error = nil
	if appEnv == "development" {
		err = godotenv.Load("./config/.env." + appEnv)
	} else {
		err = godotenv.Load(envLocation + "/.env." + appEnv)
	}
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}

	config = &Config{}
	if err := processConfig(reflect.ValueOf(config)); err != nil {
		return nil, fmt.Errorf("config load error %s", err)
	}

	if config.ConfigError != "" {
		return nil, fmt.Errorf("config load error: %s", config.ConfigError)
	}
	fmt.Println("Config loaded successfully")

	return config, nil
}
