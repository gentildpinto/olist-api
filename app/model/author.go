package model

import (
	"github.com/gentildpinto/olist-api/app/dto"
	"github.com/gentildpinto/olist-api/app/helpers"
	"github.com/gentildpinto/olist-api/config/logger"
)

type Author struct {
	Base
	Name string `gorm:"unique" json:",omitempty"`
}

func CreateAuthor(data dto.Author) (author Author, err error) {
	id := helpers.GenerateUUID()

	author = Author{
		Base: Base{
			ID: id,
		},
		Name: data.Name,
	}

	if err = logger.Log(db.Create(&author).Error); err != nil {
		return author, err
	}

	return author, err
}
