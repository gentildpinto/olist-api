package model

type Book struct {
	Base
	Name            string `json:"name" gorm:"unique"`
	Edition         string `json:"edition"`
	PublicationYear int    `json:"publication_year"`
	// Authors         []author.Author `json:"authors"`
}

func (Book) All() (books []Book, err error) {
	if err = databaseConnection.Find(&books).Error; err != nil {
		return []Book{}, err
	}

	return
}
