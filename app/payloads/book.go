package payloads

import "github.com/gentildpinto/olist-api/app/model"

type (
	CreateBook struct {
		Name            string       `json:"name" validate:"required"`
		Edition         string       `json:"edition" validate:"required"`
		PublicationYear int          `json:"publication_year" validate:"required"`
		Authors         []model.UUID `json:"authors" validate:"required"`
	}

	UpdateBook struct {
		Name            string `json:"name"`
		Edition         string `json:"edition"`
		PublicationYear int    `json:"publication_year"`
	}
)
