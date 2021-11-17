package dto

import "github.com/gentildpinto/olist-api/app/model"

type (
	Book struct {
		Name            string       `json:"name" validate:"required"`
		Edition         string       `json:"edition" validate:"required"`
		PublicationYear int          `json:"publication_year" validate:"required"`
		Authors         []model.UUID `json:"authors" validate:"required"`
	}
)
