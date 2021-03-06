package app

import (
	"github.com/gentildpinto/olist-api/app/controllers"
	echo "github.com/labstack/echo/v4"
)

var (
	welcome_controller = controllers.Welcome
	author_controller  = controllers.Author
	book_controller    = controllers.Book
)

func initRoutes(e *echo.Echo) {
	e.GET("/", welcome_controller.Index())

	// authors
	e.GET("/authors", author_controller.Index())

	// books
	e.GET("/books", book_controller.Index())
	e.POST("/books", book_controller.Create())
	e.GET("/books/:id", book_controller.FindByID())
	e.PUT("/books/:id", book_controller.Update())
	e.DELETE("/books/:id", book_controller.Delete())
}
