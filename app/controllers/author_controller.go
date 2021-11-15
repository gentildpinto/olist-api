package controllers

import (
	"net/http"

	"github.com/gentildpinto/olist-api/app/model/author"
	"github.com/labstack/echo/v4"
)

var Author = struct {
	Index func() echo.HandlerFunc
}{
	Index: func() echo.HandlerFunc {
		return func(c echo.Context) error {
			authors, err := author.GetAll()

			if err != nil {
				return err
			}

			return c.JSON(http.StatusOK, authors)
		}
	},
}
