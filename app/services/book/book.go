package book

import (
	"github.com/gentildpinto/olist-api/app/model"
	"github.com/gentildpinto/olist-api/app/payloads"
)

func All() (books []model.Book, err error) {
	books, err = (model.Book{}).All()

	return
}

func Create(payload payloads.CreateBook) (newBook model.Book, err error) {
	for _, author := range payload.Authors {
		if _, err = (model.Author{}).FindByID(author); err != nil {
			return model.Book{}, err
		}
	}

	newBook = model.Book{
		Name:            payload.Name,
		Edition:         payload.Edition,
		PublicationYear: payload.PublicationYear,
	}

	if err = newBook.Create(); err != nil {
		return model.Book{}, err
	}

	authorBook := model.AuthorBook{}

	for _, author := range payload.Authors {
		if err = authorBook.Create(author, newBook.Base.ID); err != nil {
			return model.Book{}, err
		}
	}

	return
}

func FindByID(id string) (book model.Book, err error) {
	book, err = (model.Book{}).FindByID(id)

	return
}

func Update(id string, payload payloads.UpdateBook) (updatedBook model.Book, err error) {
	updatedBook, err = (model.Book{}).FindByID(id)

	if err != nil {
		return model.Book{}, err
	}

	if err = setPayloadValues(&updatedBook, payload); err != nil {
		return model.Book{}, err
	}

	if err = updatedBook.Update(); err != nil {
		return model.Book{}, err
	}

	return
}

func Delete(id string) (err error) {
	book, err := (model.Book{}).FindByID(id)

	if err != nil {
		return err
	}

	if err = book.Delete(); err != nil {
		return err
	}

	return
}
