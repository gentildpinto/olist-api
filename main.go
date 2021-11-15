package main

import (
	"log"

	"github.com/gentildpinto/olist-api/app"
	"github.com/gentildpinto/olist-api/config"
	"github.com/subosito/gotenv"
)

const version = "0.0.1"

func init() {
	gotenv.Load()
}

func main() {
	configuration, err := config.Initialize(version)

	if err != nil {
		log.Fatal(err)
	}

	app.Start(configuration)
}
