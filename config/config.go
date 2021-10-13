package config

import (
	"fmt"

	"github.com/gentildpinto/olist-api/app/helpers"
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

	if helpers.ViperEnvVariable("APP_PORT") != "" {
		appPort = helpers.ViperEnvVariable("APP_PORT")
	}

	config = &Configuration{
		Server: &Server{
			Port:         appPort,
			Debug:        helpers.ViperEnvVariable("ENVIRONMENT") != "production" || helpers.ViperEnvVariable("DEBUG") == "true",
			ReadTimeout:  60,
			WriteTimeout: 60,
		},
		Database: &Database{
			Host:         helpers.ViperEnvVariable("DB_HOST"),
			Port:         helpers.ViperEnvVariable("DB_PORT"),
			User:         helpers.ViperEnvVariable("DB_USERNAME"),
			Password:     helpers.ViperEnvVariable("DB_PASSWORD"),
			DatabaseName: helpers.ViperEnvVariable("DB_NAME"),
			LogQueries:   helpers.ViperEnvVariable("ENVIRONMENT") != "production" || helpers.ViperEnvVariable("DEBUG") == "true",
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
		if helpers.ViperEnvVariable(envVar) == "" {
			return logger.Log(fmt.Errorf("missing environment variable: %s", envVar))
		}
	}
	return nil
}
