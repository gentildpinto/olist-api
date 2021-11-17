package controllers

import (
	"net/http"

	"github.com/gentildpinto/olist-api/app/model"
	"github.com/labstack/echo/v4"
)

var Author = struct {
	Index func() echo.HandlerFunc
}{
	Index: func() echo.HandlerFunc {
		return func(c echo.Context) error {
			authors, err := (model.Author{}).All()

			if err != nil {
				return err
			}

			return c.JSON(http.StatusOK, authors)
		}
	},
}
