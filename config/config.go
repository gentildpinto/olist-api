package config

import (
	"fmt"
	"os"

	"github.com/gentildpinto/olist-api/config/logger"
	"github.com/gentildpinto/olist-api/config/orm"
	"gorm.io/gorm"
)

var (
	version string
	appPort = "80"
	envVars = []string{
		"DB_USERNAME", "DB_PASSWORD", "ENVIRONMENT", "DB_HOST", "DB_PORT", "DB_NAME",
	}
)

type (
	Configuration struct {
		Server   *Server
		Database *Database
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

func Initialize(appVersion string) (config *Configuration, err error) {
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

	config = &Configuration{
		Server: &Server{
			Port:         appPort,
			Debug:        os.Getenv("ENVIRONMENT") != "production" || os.Getenv("DEBUG") == "true",
			ReadTimeout:  60,
			WriteTimeout: 60,
		},
		Database: &Database{
			Host:         os.Getenv("DB_HOST"),
			Port:         os.Getenv("DB_PORT"),
			User:         os.Getenv("DB_USERNAME"),
			Password:     os.Getenv("DB_PASSWORD"),
			DatabaseName: os.Getenv("DB_NAME"),
			LogQueries:   os.Getenv("ENVIRONMENT") != "production" || os.Getenv("DEBUG") == "true",
		},
	}

	config.Database.Db, err = orm.New(config.Database.User, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.DatabaseName)

	if err = logger.Log(err); err != nil {
		return
	}

	return
}

func validateEnvironment() error {
	for _, envVar := range envVars {
		if os.Getenv(envVar) == "" {
			return logger.Log(fmt.Errorf("missing environment variable: %s", envVar))
		}
	}
	return nil
}
