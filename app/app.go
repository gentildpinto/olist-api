package app

import (
	"github.com/gentildpinto/olist-api/config"
	"github.com/gentildpinto/olist-api/config/database"
	"github.com/gentildpinto/olist-api/config/logger"
	"github.com/gentildpinto/olist-api/config/server"
)

func Start(config *config.Configuration) (err error) {
	if err = logger.Log(database.Initialize(config.Database.Db)); err != nil {
		return
	}

	e := server.New()

	server.Start(e, &server.Config{
		Port:         config.Server.Port,
		ReadTimeout:  config.Server.ReadTimeout,
		WriteTimeout: config.Server.WriteTimeout,
		Debug:        config.Server.Debug,
	})

	return
}
