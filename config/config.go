package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/gentildpinto/olist-api/config/logger"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var (
	version string
	appPort = "80"
	envVars = []string{
		"DB_USERNAME", "DB_PASSWORD", "ENVIRONMENT", "DB_HOST", "DB_PORT", "DB_NAME",
	}
)

type (
	Config struct {
		Server *Server
	}

	Server struct {
		Port              string
		Debug             bool
		ReadTimeout       int
		WriteTimeout      int
		RequestsPerSecond int
	}

	Database struct {
		Host         string
		Port         string
		User         string
		Password     string
		DatabaseName string
		LogQueries   bool
		Db           *gorm.DB
	}
)

func Initialize(appVersion string) (config *Config, err error) {
	version = appVersion

	err = logger.Initialize(version)

	if err != nil {
		return
	}

	if err = validateEnvironment(); err != nil {
		return
	}

	if os.Getenv("APP_PORT") != "" {
		appPort = os.Getenv("APP_PORT")
	}
	return
}

func validateEnvironment() error {
	if err := checkFileExists(".env"); err != nil {
		if err = logger.Log(godotenv.Load()); err != nil {
			return logger.Log("could not load env file")
		}
	}

	for _, envVar := range envVars {
		if os.Getenv(envVar) == "" {
			return logger.Log(fmt.Errorf("missing environment variable: %s", envVar))
		}
	}
	return nil
}

func checkFileExists(file string) error {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return errors.New(file + " does not exist")
	}
	return nil
}
