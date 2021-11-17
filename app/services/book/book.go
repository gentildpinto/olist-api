package book

import (
	"github.com/gentildpinto/olist-api/app/dto"
	"github.com/gentildpinto/olist-api/app/model"
)

func All() (books []model.Book, err error) {
	books, err = (model.Book{}).All()

	return
}

func Create(b dto.Book) (newBook model.Book, err error) {
	for _, author := range b.Authors {
		if _, err = model.FindAuthorByID(author); err != nil {
			return model.Book{}, err
		}
	}

	newBook = model.Book{
		Name:            b.Name,
		Edition:         b.Edition,
		PublicationYear: b.PublicationYear,
	}

	if err = newBook.Create(); err != nil {
		return model.Book{}, err
	}

	authorBook := model.AuthorBook{}

	for _, author := range b.Authors {
		if err = authorBook.Create(author, newBook.Base.ID); err != nil {
			return model.Book{}, err
		}
	}

	return
}
