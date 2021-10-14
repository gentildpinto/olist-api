package app

import (
	"github.com/gentildpinto/olist-api/app/controllers"
	echo "github.com/labstack/echo/v4"
)

var (
	welcome = controllers.Welcome
)

func initRoutes(e *echo.Echo) {
	e.GET("/", welcome.Index())
}
