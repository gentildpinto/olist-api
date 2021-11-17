package model

import "github.com/gentildpinto/olist-api/config/logger"

type Book struct {
	Base
	Name            string   `json:"name" gorm:"unique"`
	Edition         string   `json:"edition"`
	PublicationYear int      `json:"publication_year"`
	Authors         []Author `json:"authors,omitempty" gorm:"many2many:author_books;"`
}

func (Book) All() (books []Book, err error) {
	if err = logger.Log(databaseConnection.Preload("Authors").Find(&books).Error); err != nil {
		return []Book{}, err
	}

	return
}

func (b *Book) Create() (err error) {
	if err = logger.Log(databaseConnection.Create(&b).Error); err != nil {
		return
	}

	return
}

func (Book) FindByID(id string) (book Book, err error) {
	uuid, _ := UUIDToBIN(id)

	if err = logger.Log(databaseConnection.Preload("Authors").First(&book, "id = ?", uuid).Error); err != nil {
		return Book{}, err
	}
	return
}
