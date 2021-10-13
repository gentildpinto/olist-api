package database

import (
	"github.com/gentildpinto/olist-api/app/helpers"
	"github.com/gentildpinto/olist-api/config/logger"
	"gorm.io/gorm"
)

var db *gorm.DB

func Initialize(database *gorm.DB) (err error) {
	db = database

	if helpers.ViperEnvVariable("TRUNCATE_TABLES") == "true" && helpers.ViperEnvVariable("ENVIRONMENT") == "development" {
		if err = logger.Log(truncateTables(db)); err != nil {
			return
		}
	}

	if err = logger.Log(autoMigrateTables(db)); err != nil {
		return
	}

	return
}

func truncateTables(db *gorm.DB) (err error) {
	return
}

func autoMigrateTables(db *gorm.DB) (err error) {
	return
}
