package controllers

import (
	"net/http"

	"github.com/gentildpinto/olist-api/app/dto"
	book "github.com/gentildpinto/olist-api/app/services/book"
	"github.com/labstack/echo/v4"
)

var Book = struct {
	Index    func() echo.HandlerFunc
	Create   func() echo.HandlerFunc
	FindByID func() echo.HandlerFunc
	Update   func() echo.HandlerFunc
}{
	Index: func() echo.HandlerFunc {
		return func(c echo.Context) error {
			books, err := book.All()

			if err != nil {
				return err
			}

			return c.JSON(http.StatusOK, books)
		}
	},
	Create: func() echo.HandlerFunc {
		return func(c echo.Context) error {
			bookDto := dto.CreateBook{}

			if err := c.Bind(&bookDto); err != nil {
				return err
			}

			newBook, err := book.Create(bookDto)

			if err != nil {
				return err
			}

			return c.JSON(http.StatusCreated, newBook)
		}
	},
	FindByID: func() echo.HandlerFunc {
		return func(c echo.Context) error {
			id := c.Param("id")

			book, err := book.FindByID(id)

			if err != nil {
				return err
			}

			return c.JSON(http.StatusOK, book)
		}
	},
	Update: func() echo.HandlerFunc {
		return func(c echo.Context) error {
			id := c.Param("id")
			bookDto := dto.UpdateBook{}

			if err := c.Bind(&bookDto); err != nil {
				return err
			}

			updatedBook, err := book.Update(id, bookDto)

			if err != nil {
				return err
			}

			return c.JSON(http.StatusOK, updatedBook)
		}
	},
}
