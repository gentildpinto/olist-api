package jobs

import (
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func StartDBQueue(db *gorm.DB) {
	database := db.Session(&gorm.Session{
		Logger: logger.Default.LogMode(logger.Info),
	})
	database = database.Set("role", "batch")
}
