package orm

import (
	"os"
	"strconv"
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func New(user, password, host, port, databaseName string) (*gorm.DB, error) {
	mysqlConfig := mysql.Config{
		DSN:                       user + ":" + password + "@tcp(" + host + ":" + port + ")/" + databaseName + "?charset=utf8&parseTime=True&loc=Local",
		DefaultStringSize:         256,
		SkipInitializeWithVersion: false,
	}

	database, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   logger.Default.LogMode(logger.Info),
		CreateBatchSize:                          1000,
		AllowGlobalUpdate:                        true,
	})

	if err != nil {
		return nil, err
	}

	maxDBConnections := 500
	if connections, err := strconv.Atoi(os.Getenv("MAX_DB_CONNECTIONS")); err == nil {
		maxDBConnections = connections
	}

	maxOpenConnections := int(float64(maxDBConnections) * 0.75)
	maxIdleConnections := int(maxDBConnections) - maxOpenConnections

	sqlDB, err := database.DB()

	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(maxOpenConnections)
	sqlDB.SetMaxIdleConns(maxIdleConnections)
	sqlDB.SetConnMaxLifetime(time.Hour * 8)

	return database, err
}
