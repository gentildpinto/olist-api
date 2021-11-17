package model

import "github.com/gentildpinto/olist-api/config/logger"

type Author struct {
	Base
	Name  string `json:"name" gorm:"unique"`
	Books []Book `json:"books,omitempty" gorm:"many2many:author_books;"`
}

func (Author) All() (authors []Author, err error) {
	if err = logger.Log(databaseConnection.Preload("Books").Find(&authors).Error); err != nil {
		return []Author{}, err
	}

	return
}

func FindAuthorByID(id UUID) (author Author, err error) {
	if err = logger.Log(databaseConnection.First(&author, id).Error); err != nil {
		return Author{}, err
	}

	return
}
