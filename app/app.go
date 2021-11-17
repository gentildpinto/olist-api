package app

import (
	"github.com/gentildpinto/olist-api/app/model"
	"github.com/gentildpinto/olist-api/config"
	"github.com/gentildpinto/olist-api/config/logger"
	"github.com/gentildpinto/olist-api/config/server"
)

func Start(config *config.Configuration) (err error) {
	if err = logger.Log(model.Initialize(config.Database.Db)); err != nil {
		return
	}

	e := server.New()

	initRoutes(e)

	server.Start(e, &server.Config{
		Port:         config.Server.Port,
		ReadTimeout:  config.Server.ReadTimeout,
		WriteTimeout: config.Server.WriteTimeout,
		Debug:        config.Server.Debug,
	})

	return
}
