package author

import "github.com/gentildpinto/olist-api/app/model"

func All() (authors []model.Author, err error) {
	authors, err = (model.Author{}).All()

	return
}
