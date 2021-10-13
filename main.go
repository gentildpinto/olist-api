package main

import (
	"log"

	"github.com/gentildpinto/olist-api/app"
	"github.com/gentildpinto/olist-api/config"
)

const version = "0.0.1"

func main() {
	configuration, err := config.Initialize(version)

	if err != nil {
		log.Fatal(err)
	}

	app.Start(configuration)
}
